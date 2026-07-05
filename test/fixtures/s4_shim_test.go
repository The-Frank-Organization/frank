package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestNudgeRecipientOnly(t *testing.T) {
	h := newS4ShimHarness(t)
	credA := h.mint(t, "seat-a", "implementer")
	credB := h.mint(t, "seat-b", "planner")
	h.start(t)

	sender := h.dial(t, credA)
	defer func() { _ = sender.Close() }()
	recipient := h.dial(t, credB)
	defer func() { _ = recipient.Close() }()

	relayID := h.submit(t, sender, s4Relay("seat-b", "", "recipient-only"))
	expectNoFixturePush(t, sender)
	expectFixtureNudge(t, h.ctx, recipient, "delivery-nudge", relayID)
}

func TestNudgeAllRecipients(t *testing.T) {
	h := newS4ShimHarness(t)
	credA := h.mint(t, "seat-a", "implementer")
	credB := h.mint(t, "seat-b", "planner")
	credC := h.mint(t, "seat-c", "reviewer")
	h.start(t)

	sender := h.dial(t, credA)
	defer func() { _ = sender.Close() }()
	seatB := h.dial(t, credB)
	defer func() { _ = seatB.Close() }()
	seatC := h.dial(t, credC)
	defer func() { _ = seatC.Close() }()

	relayID := h.submit(t, sender, s4Relay("seat-b", `["seat-c"]`, "all-recipients"))
	expectFixtureNudge(t, h.ctx, seatB, "delivery-nudge", relayID)
	expectFixtureNudge(t, h.ctx, seatC, "delivery-nudge", relayID)
}

func TestOfflineRecipientNudgedOnReconnect(t *testing.T) {
	h := newS4ShimHarness(t)
	credA := h.mint(t, "seat-a", "implementer")
	credB := h.mint(t, "seat-b", "planner")
	h.start(t)

	sender := h.dial(t, credA)
	defer func() { _ = sender.Close() }()
	h.submit(t, sender, s4Relay("seat-b", "", "offline-recipient"))
	expectNoFixturePush(t, sender)

	recipient := h.dial(t, credB)
	defer func() { _ = recipient.Close() }()
	expectFixtureNudge(t, h.ctx, recipient, "recovery-nudge", "")
}

func TestNoCrossSeatMetadata(t *testing.T) {
	h := newS4ShimHarness(t)
	credB := h.mint(t, "seat-b", "planner")
	_ = h.mint(t, "seat-c", "reviewer")
	st, err := store.Open(h.root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "relay-b", DispatchID: "d", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "SUBJECT": "b"},
	}, nil); err != nil {
		t.Fatalf("commit relay-b: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "relay-c", DispatchID: "d", From: "seat-a", To: "seat-c", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "SUBJECT": "c"},
	}, nil); err != nil {
		t.Fatalf("commit relay-c: %v", err)
	}
	h.start(t)

	recipient := h.dial(t, credB)
	defer func() { _ = recipient.Close() }()
	frame := expectFixtureNudge(t, h.ctx, recipient, "recovery-nudge", "")
	for _, forbidden := range [][]byte{[]byte("seats"), []byte("seat-a"), []byte("seat-c"), []byte("relay-c")} {
		if bytes.Contains(frame, forbidden) {
			t.Fatalf("recovery nudge leaked %q in %s", forbidden, frame)
		}
	}
}

type s4ShimHarness struct {
	root   string
	ctx    context.Context
	cancel context.CancelFunc
	bin    string
	sock   string
	mgr    *seat.Manager
	stderr *bytes.Buffer
	cmd    *exec.Cmd
}

func newS4ShimHarness(t *testing.T) *s4ShimHarness {
	t.Helper()
	root := t.TempDir()
	initFixtureStore(t, root)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	h := &s4ShimHarness{
		root:   root,
		ctx:    ctx,
		cancel: cancel,
		bin:    buildFrank(t, ctx),
		sock:   filepath.Join(os.TempDir(), fmt.Sprintf("frank-s4-%d.sock", time.Now().UnixNano())),
		mgr:    mgr,
	}
	t.Cleanup(func() {
		cancel()
		if h.cmd != nil && h.cmd.Process != nil {
			_ = h.cmd.Process.Kill()
			_ = h.cmd.Wait()
		}
		_ = os.Remove(h.sock)
	})
	return h
}

func (h *s4ShimHarness) mint(t *testing.T, name, role string) seat.Cred {
	t.Helper()
	cred, err := h.mgr.Mint(name, role, false)
	if err != nil {
		t.Fatalf("Mint %s: %v", name, err)
	}
	return cred
}

func (h *s4ShimHarness) start(t *testing.T) {
	t.Helper()
	if h.cmd != nil {
		t.Fatalf("s4 shim already started")
	}
	h.cmd, h.stderr = startFrank(t, h.ctx, h.bin, h.root, h.sock)
	waitForSocket(t, h.sock)
}

func (h *s4ShimHarness) stop(t *testing.T) {
	t.Helper()
	if h.cmd == nil {
		return
	}
	if h.cmd.Process != nil {
		_ = h.cmd.Process.Kill()
	}
	_ = h.cmd.Wait()
	h.cmd = nil
	h.stderr = nil
	_ = os.Remove(h.sock)
}

func (h *s4ShimHarness) dial(t *testing.T, cred seat.Cred) *channel.Client {
	t.Helper()
	client, err := channel.DialAuthenticated(h.ctx, h.sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", h.stderr.String(), err)
	}
	return client
}

func (h *s4ShimHarness) submit(t *testing.T, client *channel.Client, rec record.Record) string {
	t.Helper()
	describe, err := client.DescribeTools(h.ctx, channel.DescribeRequest{Phase: rec.Headers["PHASE"], Tier: rec.Headers["CEREMONY_TIER"]})
	if err != nil {
		t.Fatalf("DescribeTools stderr=%s: %v", h.stderr.String(), err)
	}
	payload := mustJSONBytes(t, fieldspec.SubmitPayload{Record: rec, FormDigest: describe.FormDigest})
	result, err := client.Call(h.ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit stderr=%s: %v", h.stderr.String(), err)
	}
	var outcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode submit outcome %s: %v", result, err)
	}
	if outcome.State != record.Accepted || outcome.RelayID == "" {
		t.Fatalf("submit outcome = %+v", outcome)
	}
	return outcome.RelayID
}

func s4Relay(to, cc, subject string) record.Record {
	headers := map[string]string{
		"PHASE":         "SITREP",
		"AUTHORITY":     "report-only",
		"CEREMONY_TIER": "medium",
		"SUBJECT":       subject,
	}
	if cc != "" {
		headers["CC"] = cc
	}
	return record.Record{
		Envelope: record.Envelope{To: to, DispatchID: "s4-nudge"},
		Headers:  headers,
		Body:     "wake",
	}
}

func expectFixtureNudge(t *testing.T, ctx context.Context, client *channel.Client, kind, relayID string) []byte {
	t.Helper()
	frame, err := client.NextPush(ctx)
	if err != nil {
		t.Fatalf("NextPush %s: %v", kind, err)
	}
	var msg map[string]any
	if err := json.Unmarshal(frame, &msg); err != nil {
		t.Fatalf("decode nudge %s: %v", frame, err)
	}
	if msg["kind"] != kind {
		t.Fatalf("nudge kind = %v, want %s; frame=%s", msg["kind"], kind, frame)
	}
	if relayID != "" && msg["relay_id"] != relayID {
		t.Fatalf("nudge relay_id = %v, want %s; frame=%s", msg["relay_id"], relayID, frame)
	}
	if relayID == "" {
		if _, ok := msg["relay_id"]; ok {
			t.Fatalf("recovery nudge included relay_id: %s", frame)
		}
	}
	return frame
}

func expectNoFixturePush(t *testing.T, client *channel.Client) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if frame, err := client.NextPush(ctx); err == nil {
		t.Fatalf("unexpected push: %s", frame)
	}
}
