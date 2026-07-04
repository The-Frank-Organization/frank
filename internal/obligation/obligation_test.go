package obligation_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestCompleteAutoHandlesGateHeldAndQuarantineFacts(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitObligationRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "gate-ob", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security"},
	})
	commitObligationRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "held-ob", From: "system", Role: "system", DeliveryState: record.Held, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "held"},
	})
	if err := os.WriteFile(filepath.Join(root, "records", "quarantined-ob.json"), []byte("not sealed"), 0o644); err != nil {
		t.Fatalf("write corrupt record: %v", err)
	}
	if _, err := st.QuarantineOne("quarantined-ob"); err != nil {
		t.Fatalf("QuarantineOne: %v", err)
	}

	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto: %v", err)
	}
	if countByRelay(t, st, "park-gate-ob") != 1 ||
		countByRelay(t, st, "outbox-gate-gate-ob") != 1 ||
		countByRelay(t, st, "outbox-held-held-ob") != 1 ||
		countByRelay(t, st, "incident-quarantined-ob") != 1 {
		t.Fatalf("completion records missing")
	}

	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto idempotent: %v", err)
	}
	if countByRelay(t, st, "incident-quarantined-ob") != 1 {
		t.Fatalf("incident duplicated")
	}
}

func TestEngineOpenSetUsesRegisteredClass(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitObligationRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "source-1", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "custom_source"},
	})
	engine := obligation.New()
	engine.Register(obligation.Class{
		Name: "custom",
		Sources: func(st *store.Store) ([]obligation.Fact, error) {
			return []obligation.Fact{{Kind: "custom_source", Key: "source-1"}}, nil
		},
		Completed: func(f obligation.Fact, t *obligation.Tables) bool {
			return t.CompletionIndex["custom"][f.Key]
		},
	})
	if err := engine.BuildTables(st); err != nil {
		t.Fatalf("BuildTables: %v", err)
	}
	open := engine.Open("custom")
	if len(open) != 1 || open[0].Key != "source-1" {
		t.Fatalf("open = %+v, want source-1", open)
	}
}

func commitObligationRecord(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
	}
}

func countByRelay(t *testing.T, st *store.Store, relayID string) int {
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
