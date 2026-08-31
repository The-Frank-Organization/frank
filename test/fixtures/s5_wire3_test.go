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

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/gate"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

// These binary-path fixtures prove Step-1 detection exactly for S1 + S2 + S3
// plus other->A. S3 is mechanism-wired but input-atom-pending while the pinned
// config leaves target_branch_field unset.
func TestS5Wire3BinaryPathRaisesFromPinnedS1Config(t *testing.T) {
	root := t.TempDir()
	pinned := initWire3FixtureStore(t, root, `"detector":{"a_floor":[{"phase":"SITREP","member":"authz_security"}]}`)
	reg := loadAssemblyRegistry(t)
	cred := mintWire3Seat(t, root, "seat-a", "implementer", false)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-wire3-s1-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)

	client := dialWire3(t, ctx, sock, cred.Value, stderr)
	defer func() { _ = client.Close() }()
	outcome := submitWire3(t, ctx, client, stderr, reg, pinned.Digest, fieldspec.SeatMeta{Name: "seat-a", Role: "implementer"}, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "wire3 s1",
			"gate_category":   "routing",
		},
	})
	got := readWire3Record(t, ctx, client, outcome.RelayID)
	if got.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", got.Envelope.DeliveryState, got.Body)
	}
	if got.Headers["gate_category"] != "authz_security" || got.Headers["gate_category_raised"] != "yes" || got.Headers["gate_category_pick"] != "routing" {
		t.Fatalf("raise headers = %+v", got.Headers)
	}
	items := wire3OutboxItems(t, root)
	if len(items) != 1 || items[0].SourceRecordRef != outcome.RelayID || items[0].GateCategory != "authz_security" {
		t.Fatalf("outbox items = %+v", items)
	}
}

func TestS5Wire3BinaryPathS2ReferenceRaisesAndAbsentConfigIsFailSafeOnly(t *testing.T) {
	root := t.TempDir()
	pinned := initFixtureStore(t, root)
	reg := loadAssemblyRegistry(t)
	seatCred := mintWire3Seat(t, root, "seat-a", "implementer", false)
	operatorCred := mintWire3Seat(t, root, "operator", "operator", true)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-wire3-s2-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)

	seatClient := dialWire3(t, ctx, sock, seatCred.Value, stderr)
	defer func() { _ = seatClient.Close() }()
	operatorClient := dialWire3(t, ctx, sock, operatorCred.Value, stderr)
	defer func() { _ = operatorClient.Close() }()

	boundary := submitWire3(t, ctx, seatClient, stderr, reg, pinned.Digest, fieldspec.SeatMeta{Name: "seat-a", Role: "implementer"}, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "wire3 absent detector",
			"gate_category":   "routing",
		},
	})
	boundaryRead := readWire3Record(t, ctx, seatClient, boundary.RelayID)
	if boundaryRead.Headers["gate_category"] != "routing" || boundaryRead.Headers["gate_category_raised"] != "" {
		t.Fatalf("absent detector mutated headers = %+v", boundaryRead.Headers)
	}

	gateOutcome := submitWire3(t, ctx, operatorClient, stderr, reg, pinned.Digest, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, record.Record{
		Envelope: record.Envelope{To: "operator", DispatchID: "wire3-s2"},
		Headers: map[string]string{
			"PHASE":               "SITREP",
			"AUTHORITY":           "report-only",
			"CEREMONY_TIER":       "medium",
			"EVIDENCE_TARGET":     "E1",
			"SUBJECT":             "wire3 gate",
			"HUMAN_GATE_REQUIRED": "yes",
			"gate_category":       "authz_security",
		},
	})
	gateRead := readWire3Record(t, ctx, operatorClient, gateOutcome.RelayID)
	if gateRead.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("gate state = %s, body=%s", gateRead.Envelope.DeliveryState, gateRead.Body)
	}
	contextOutcome := submitWire3(t, ctx, operatorClient, stderr, reg, pinned.Digest, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, record.Record{
		Envelope: record.Envelope{To: "operator", DispatchID: "wire3-s2"},
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "wire3 s2 active dispatch context",
		},
	})
	contextRead := readWire3Record(t, ctx, operatorClient, contextOutcome.RelayID)
	if contextRead.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("context state = %s, body=%s", contextRead.Envelope.DeliveryState, contextRead.Body)
	}
	verdictOutcome := submitWire3(t, ctx, operatorClient, stderr, reg, pinned.Digest, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, record.Record{
		Envelope: record.Envelope{To: "operator", DispatchID: "wire3-s2"},
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "wire3 s2",
			"record_kind":     "gate_resolution",
			"parent_hint":     gateOutcome.RelayID,
			"resolves_gate":   gateOutcome.RelayID,
			"gate_category":   "routing",
		},
		Body: `{"choice":"approve"}`,
	})
	verdictRead := readWire3Record(t, ctx, operatorClient, verdictOutcome.RelayID)
	if verdictRead.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("verdict state = %s, body=%s", verdictRead.Envelope.DeliveryState, verdictRead.Body)
	}
	if verdictRead.Headers["gate_category"] != "authz_security" || verdictRead.Headers["gate_category_raised"] != "yes" || verdictRead.Headers["gate_category_pick"] != "routing" {
		t.Fatalf("S2 raise headers = %+v", verdictRead.Headers)
	}
}

func TestS5Wire3InvalidAFloorMemberFailsBeforeServing(t *testing.T) {
	root := t.TempDir()
	initWire3FixtureStore(t, root, `"detector":{"a_floor":[{"phase":"SITREP","record_kind":"diagnostics","member":"routing"}]}`)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-wire3-invalid-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	run := exec.CommandContext(ctx, bin, "-root", root, "-socket", sock)
	out, err := run.CombinedOutput()
	if err == nil {
		t.Fatalf("frank unexpectedly served with invalid detector config; output=%s", out)
	}
	if !bytes.Contains(out, []byte("gate_category_A")) {
		t.Fatalf("startup output = %s, want gate_category_A", out)
	}
}

func initWire3FixtureStore(t *testing.T, root, detectorJSON string) *config.Pinned {
	t.Helper()
	sources := writeFixtureConfigSources(t)
	if err := os.WriteFile(sources["engine"], wire3EngineConfigBytes(t, detectorJSON), 0o644); err != nil {
		t.Fatalf("write engine source: %v", err)
	}
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("store Init: %v", err)
	}
	return loadFixturePinned(t, root)
}

func writeWire3EngineConfig(t *testing.T, root, detectorJSON string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(root, "config", "engine.json"), wire3EngineConfigBytes(t, detectorJSON), 0o644); err != nil {
		t.Fatalf("write wire3 engine config: %v", err)
	}
}

func wire3EngineConfigBytes(t *testing.T, detectorJSON string) []byte {
	t.Helper()
	var engine map[string]any
	if err := json.Unmarshal(fixtureEngineConfig(t, false), &engine); err != nil {
		t.Fatalf("decode fixture engine: %v", err)
	}
	var detector map[string]any
	if err := json.Unmarshal([]byte("{"+detectorJSON+"}"), &detector); err != nil {
		t.Fatalf("decode detector config: %v", err)
	}
	engine["detector"] = detector["detector"]
	raw, err := json.Marshal(engine)
	if err != nil {
		t.Fatalf("marshal wire3 engine: %v", err)
	}
	return raw
}

func mintWire3Seat(t *testing.T, root, name, role string, operator bool) seat.Cred {
	t.Helper()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint(name, role, operator)
	if err != nil {
		t.Fatalf("Mint %s: %v", name, err)
	}
	return cred
}

func dialWire3(t *testing.T, ctx context.Context, sock, credential string, stderr *bytes.Buffer) *channel.Client {
	t.Helper()
	client, err := channel.DialAuthenticated(ctx, sock, credential)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	return client
}

func submitWire3(t *testing.T, ctx context.Context, client *channel.Client, stderr *bytes.Buffer, reg *fieldspec.Registry, configDigest string, meta fieldspec.SeatMeta, rec record.Record) liveSubmitOutcome {
	t.Helper()
	payload := submitPayloadBytes(t, reg, configDigest, seat.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec)
	result, err := client.Call(ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit stderr=%s: %v", stderr.String(), err)
	}
	var outcome liveSubmitOutcome
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode submit outcome %s: %v", result, err)
	}
	if outcome.RelayID == "" {
		t.Fatalf("submit outcome missing relay id: %+v", outcome)
	}
	return outcome
}

func readWire3Record(t *testing.T, ctx context.Context, client *channel.Client, relayID string) record.Record {
	t.Helper()
	payload, _ := json.Marshal(map[string]string{"relay_id": relayID})
	result, err := client.Call(ctx, "read", payload)
	if err != nil {
		t.Fatalf("read %s: %v", relayID, err)
	}
	var read struct {
		Record record.Record `json:"record"`
	}
	if err := json.Unmarshal(result, &read); err != nil {
		t.Fatalf("decode read %s: %v", result, err)
	}
	return read.Record
}

func wire3OutboxItems(t *testing.T, root string) []gate.OutboxItem {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(root, "outbox"))
	if err != nil {
		t.Fatalf("ReadDir outbox: %v", err)
	}
	items := make([]gate.OutboxItem, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		var item gate.OutboxItem
		if err := json.Unmarshal(mustReadFile(t, filepath.Join(root, "outbox", entry.Name())), &item); err != nil {
			t.Fatalf("decode outbox item: %v", err)
		}
		items = append(items, item)
	}
	return items
}
