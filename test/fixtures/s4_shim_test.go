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

func TestCCRecipientMailboxedNudgedOnceAndPathClean(t *testing.T) {
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

	relayID := h.submit(t, sender, s4Relay("seat-b", `["seat-c","seat-b","seat-c"]`, "cc-dedupe"))
	for seatName, client := range map[string]*channel.Client{"seat-b": seatB, "seat-c": seatC} {
		frame := expectFixtureNudge(t, h.ctx, client, "delivery-nudge", relayID)
		for _, forbidden := range [][]byte{[]byte(h.root), []byte(h.sock), []byte("binding"), []byte("seats.json"), []byte(credA.Value), []byte(credB.Value), []byte(credC.Value), []byte("wake")} {
			if bytes.Contains(frame, forbidden) {
				t.Fatalf("%s nudge leaked %q in %s", seatName, forbidden, frame)
			}
		}
		expectNoFixturePush(t, client)
		assertFixtureProject(t, h.ctx, client, seatName, relayID)
		assertFixtureMailbox(t, h.root, seatName, relayID)
	}
	expectNoFixturePush(t, sender)
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

func TestFrankMCPInitializeEnvelopeDeclaresServerInfoVersion(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	mcpBin := buildFrankMCP(t, ctx)
	out, errOut := runFrankMCP(t, ctx, mcpBin, "", "", `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`+"\n")
	if errOut != "" {
		t.Fatalf("initialize stderr = %s", errOut)
	}
	var envelope struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  struct {
			ProtocolVersion string `json:"protocolVersion"`
			Capabilities    struct {
				Tools map[string]bool `json:"tools"`
			} `json:"capabilities"`
			ServerInfo struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"serverInfo"`
		} `json:"result"`
		Error *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(bytes.TrimSpace([]byte(out)), &envelope); err != nil {
		t.Fatalf("decode initialize envelope %q: %v", out, err)
	}
	if envelope.JSONRPC != "2.0" || envelope.ID != 1 || envelope.Error != nil {
		t.Fatalf("initialize envelope = %+v", envelope)
	}
	if envelope.Result.ProtocolVersion == "" {
		t.Fatalf("initialize missing protocolVersion: %+v", envelope.Result)
	}
	if envelope.Result.Capabilities.Tools == nil || !envelope.Result.Capabilities.Tools["listChanged"] {
		t.Fatalf("initialize tools capabilities = %#v", envelope.Result.Capabilities.Tools)
	}
	if envelope.Result.ServerInfo.Name == "" || envelope.Result.ServerInfo.Version == "" {
		t.Fatalf("initialize serverInfo = %+v", envelope.Result.ServerInfo)
	}
}

func TestOperatorDescribeToolsIncludesOwedRecordFields(t *testing.T) {
	h := newS4ShimHarness(t)
	cred := h.mint(t, "operator", "operator")
	h.start(t)

	operator := h.dial(t, cred)
	defer func() { _ = operator.Close() }()
	describe, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools stderr=%s: %v", h.stderr.String(), err)
	}
	if describe.SubmitSchema == nil || describe.FormDigest == "" {
		t.Fatalf("describe missing form: %+v", describe)
	}
	if !describe.SubmitSchema.OptionAllowed("record_kind", "owed_item") ||
		!describe.SubmitSchema.OptionAllowed("record_kind", "owed_disposition") {
		t.Fatalf("record_kind options = %+v", describe.SubmitSchema.Fields["record_kind"].Options)
	}
	for _, want := range []struct {
		Name string
		Type string
	}{
		{Name: "owner", Type: "text"},
		{Name: "source", Type: "text"},
		{Name: "target_surface", Type: "text"},
		{Name: "disposition_path", Type: "text"},
		{Name: "disposes_owed", Type: "id_ref"},
	} {
		field, ok := describe.SubmitSchema.Fields[want.Name]
		if !ok {
			t.Fatalf("describe form missing %s; fields=%v", want.Name, describe.SubmitSchema.Fields)
		}
		if field.Type != want.Type {
			t.Fatalf("%s field type = %q, want %q", want.Name, field.Type, want.Type)
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
		"PHASE":           "SITREP",
		"AUTHORITY":       "report-only",
		"CEREMONY_TIER":   "medium",
		"EVIDENCE_TARGET": "E1",
		"SUBJECT":         subject,
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

func assertFixtureProject(t *testing.T, ctx context.Context, client *channel.Client, seatName, relayID string) {
	t.Helper()
	data, err := client.Call(ctx, "project", nil)
	if err != nil {
		t.Fatalf("project %s: %v", seatName, err)
	}
	var got []string
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("decode project %s: %v", data, err)
	}
	if len(got) != 1 || got[0] != relayID {
		t.Fatalf("project %s = %v, want [%s]", seatName, got, relayID)
	}
}

func assertFixtureMailbox(t *testing.T, root, seatName, relayID string) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, "mailboxes", seatName+".jsonl"))
	if err != nil {
		t.Fatalf("read mailbox %s: %v", seatName, err)
	}
	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	if len(lines) != 1 || string(lines[0]) != relayID {
		t.Fatalf("mailbox %s = %q, want exactly %s", seatName, data, relayID)
	}
}
