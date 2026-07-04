package gate_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestCompleteCreatesGateOutboxAndIsIdempotent(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "gate-1", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security"},
	}, nil); err != nil {
		t.Fatalf("commit gate: %v", err)
	}
	if err := gate.Complete(st); err != nil {
		t.Fatalf("Complete first: %v", err)
	}
	if err := gate.Complete(st); err != nil {
		t.Fatalf("Complete second: %v", err)
	}
	items := readOutbox(t, root)
	if len(items) != 1 || items[0].SourceKind != "gate" || items[0].SourceRecordRef != "gate-1" {
		t.Fatalf("outbox items = %+v", items)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var parks int
	for _, rec := range records {
		if rec.Headers["parks_gate"] == "gate-1" {
			parks++
		}
	}
	if parks != 1 {
		t.Fatalf("park records = %d, want 1", parks)
	}
}

func TestCompleteConvergesWhenGateParkExistsWithoutOutbox(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitGateRecord(t, st, "gate-partial")
	commitParkRecord(t, st, "gate-partial")

	if err := gate.Complete(st); err != nil {
		t.Fatalf("Complete: %v", err)
	}
	items := readOutbox(t, root)
	if len(items) != 1 || items[0].SourceKind != "gate" || items[0].SourceRecordRef != "gate-partial" {
		t.Fatalf("outbox items = %+v", items)
	}
	if parks := countParkRecords(t, st, "gate-partial"); parks != 1 {
		t.Fatalf("park records = %d, want 1", parks)
	}
}

func TestCompleteConvergesWhenGateOutboxRecordExistsWithoutPark(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitGateRecord(t, st, "gate-mirror")
	commitOutboxRecordOnly(t, st, "gate-gate-mirror")

	if err := gate.Complete(st); err != nil {
		t.Fatalf("Complete: %v", err)
	}
	if parks := countParkRecords(t, st, "gate-mirror"); parks != 1 {
		t.Fatalf("park records = %d, want 1", parks)
	}
	if outboxRecords := countRecordsByRelayID(t, st, "outbox-gate-gate-mirror"); outboxRecords != 1 {
		t.Fatalf("outbox records = %d, want 1", outboxRecords)
	}
}

func TestCompleteCreatesHeldOutbox(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "held-1", From: "system", Role: "system", DeliveryState: record.Held, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "held"},
	}, nil); err != nil {
		t.Fatalf("commit held: %v", err)
	}
	if err := gate.Complete(st); err != nil {
		t.Fatalf("Complete: %v", err)
	}
	items := readOutbox(t, root)
	if len(items) != 1 || items[0].SourceKind != "held" || items[0].SourceRecordRef != "held-1" {
		t.Fatalf("outbox items = %+v", items)
	}
}

func commitGateRecord(t *testing.T, st *store.Store, relayID string) {
	t.Helper()
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: relayID, From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security"},
	}, nil); err != nil {
		t.Fatalf("commit gate: %v", err)
	}
}

func commitParkRecord(t *testing.T, st *store.Store, gateRelayID string) {
	t.Helper()
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "park-" + gateRelayID,
			From:          "system",
			Role:          "system",
			To:            "seat-a",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":      "SITREP",
			"SUBJECT":    "parked gate",
			"parks_gate": gateRelayID,
		},
	}, nil); err != nil {
		t.Fatalf("commit park: %v", err)
	}
}

func commitOutboxRecordOnly(t *testing.T, st *store.Store, itemID string) {
	t.Helper()
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "outbox-" + itemID,
			From:          "system",
			Role:          "system",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "derived outbox item"},
	}, []store.Intent{}); err != nil {
		t.Fatalf("commit outbox record: %v", err)
	}
}

func countParkRecords(t *testing.T, st *store.Store, gateRelayID string) int {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var parks int
	for _, rec := range records {
		if rec.Headers["parks_gate"] == gateRelayID {
			parks++
		}
	}
	return parks
}

func countRecordsByRelayID(t *testing.T, st *store.Store, relayID string) int {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var count int
	for _, rec := range records {
		if rec.Envelope.RelayID == relayID {
			count++
		}
	}
	return count
}

func readOutbox(t *testing.T, root string) []gate.OutboxItem {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(root, "outbox"))
	if err != nil {
		t.Fatalf("ReadDir outbox: %v", err)
	}
	var items []gate.OutboxItem
	for _, entry := range entries {
		data, err := os.ReadFile(filepath.Join(root, "outbox", entry.Name()))
		if err != nil {
			t.Fatalf("read outbox item: %v", err)
		}
		var item gate.OutboxItem
		if err := json.Unmarshal(data, &item); err != nil {
			t.Fatalf("decode outbox item: %v", err)
		}
		items = append(items, item)
	}
	return items
}
