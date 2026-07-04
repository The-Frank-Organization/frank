package gate_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
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

func TestResolveGateIsOneShotUnderRace(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "gate-race", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes"},
	}, nil); err != nil {
		t.Fatalf("commit gate: %v", err)
	}
	var wg sync.WaitGroup
	results := make(chan string, 2)
	for _, verdict := range []string{"v1", "v2"} {
		wg.Add(1)
		go func(verdict string) {
			defer wg.Done()
			state, err := gate.Resolve(st, "gate-race", verdict)
			if err != nil {
				t.Errorf("Resolve %s: %v", verdict, err)
			}
			results <- state
		}(verdict)
	}
	wg.Wait()
	close(results)
	var accepted, rejected int
	for state := range results {
		if state == record.Accepted {
			accepted++
		}
		if state == record.Rejected {
			rejected++
		}
	}
	if accepted != 1 || rejected != 1 {
		t.Fatalf("accepted=%d rejected=%d, want 1/1", accepted, rejected)
	}
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
