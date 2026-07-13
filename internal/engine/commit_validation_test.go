package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestCommitRevalidationSweepsSnapshotDependentRecordKinds(t *testing.T) {
	t.Run("owed disposition", func(t *testing.T) {
		st, err := store.Open(t.TempDir())
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		commitValidationRecord(t, st, record.Record{
			Envelope: record.Envelope{RelayID: "owed-target", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers: map[string]string{
				"record_kind": "owed_item", "owner": "s10", "source": "test",
				"target_surface": "fold", "disposition_path": "done",
			},
		})
		commitValidationRecord(t, st, record.Record{
			Envelope: record.Envelope{RelayID: "owed-first", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"record_kind": "owed_disposition", "disposes_owed": "owed-target"},
		})
		candidate := record.Record{
			Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"record_kind": "owed_disposition", "disposes_owed": "owed-target"},
		}
		got, _, err := revalidateAtCommit(st, candidate, nil)
		if err != nil || got.Envelope.DeliveryState != record.Rejected || !strings.Contains(got.Body, "disposes_owed:already-resolved") {
			t.Fatalf("commit revalidation = %#v, %v", got, err)
		}
	})

	t.Run("waiver retraction", func(t *testing.T) {
		st, err := store.Open(t.TempDir())
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		commitValidationRecord(t, st, record.Record{
			Envelope: record.Envelope{RelayID: "waiver-target", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"waiver_scope": "one-gate"},
		})
		commitValidationRecord(t, st, record.Record{
			Envelope: record.Envelope{RelayID: "waiver-first", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"record_kind": "waiver_retraction", "retracts": "waiver-target"},
		})
		candidate := record.Record{
			Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"record_kind": "waiver_retraction", "retracts": "waiver-target"},
		}
		got, _, err := revalidateAtCommit(st, candidate, nil)
		if err != nil || got.Envelope.DeliveryState != record.Rejected || !strings.Contains(got.Body, "retracts:already-resolved") {
			t.Fatalf("commit revalidation = %#v, %v", got, err)
		}
	})

	t.Run("seat remint remains latest-wins", func(t *testing.T) {
		st, err := store.Open(t.TempDir())
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		commitValidationRecord(t, st, record.Record{
			Envelope: record.Envelope{RelayID: "seat-mint-first", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"record_kind": "seat_mint"},
			Body:     `{"seat":"seat-a","role":"implementer","is_operator":false}`,
		})
		candidate := record.Record{
			Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"record_kind": "seat_mint"},
			Body:     `{"seat":"seat-a","role":"implementer","is_operator":false}`,
		}
		got, _, err := revalidateAtCommit(st, candidate, nil)
		if err != nil || got.Envelope.DeliveryState != record.Accepted {
			t.Fatalf("latest-wins remint revalidation = %#v, %v", got, err)
		}
	})

	t.Run("config is reread", func(t *testing.T) {
		root := t.TempDir()
		enginePath := filepath.Join(t.TempDir(), "engine.json")
		if err := os.WriteFile(enginePath, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
			t.Fatalf("write engine source: %v", err)
		}
		if err := store.Init(root, map[string]string{
			"engine": enginePath, "fieldspec": filepath.Join("..", "fieldspec", "registry.json"),
		}); err != nil {
			t.Fatalf("Init: %v", err)
		}
		st, err := store.Open(root)
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		if err := os.WriteFile(filepath.Join(root, "config", "engine.json"), []byte(`{"broken":`), 0o644); err != nil {
			t.Fatalf("corrupt pinned engine: %v", err)
		}
		candidate := record.Record{
			Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers: map[string]string{
				"record_kind": "config_change", "member": "engine", "new_digest": "stale-handler-digest",
			},
			Body: `{"gc_enabled":true,"segment_rotate_bytes":4194304}`,
		}
		got, _, err := revalidateAtCommit(st, candidate, nil)
		if err != nil || got.Envelope.DeliveryState != record.Rejected || !strings.Contains(got.Body, "member:config-read-error") {
			t.Fatalf("config commit revalidation = %#v, %v", got, err)
		}
	})
}

func commitValidationRecord(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, []store.Intent{}); err != nil {
		t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
	}
}
