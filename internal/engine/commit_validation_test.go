package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
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

	t.Run("mint chain anchor rejects a chain resolved after handler snapshot", func(t *testing.T) {
		st, err := store.Open(t.TempDir())
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		commitValidationRecord(t, st, record.Record{
			Envelope: record.Envelope{RelayID: "resolved-pivot", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"record_kind": "seat_mint"},
			Body:     `{"seat":"anchor-seat","role":"implementer","is_operator":false}`,
		})
		candidate := MintChainAnchorRecord("anchor-seat", "resolved-pivot")
		candidate.Envelope.From = "operator"
		candidate.Envelope.Role = "operator"
		got, _, err := revalidateAtCommit(st, candidate, nil)
		if err != nil || got.Envelope.DeliveryState != record.Rejected || got.Headers["failing_edge"] != "anchor-target-resolved" || !strings.Contains(got.Body, "anchor-target-resolved") {
			t.Fatalf("anchor commit revalidation = %#v, %v", got, err)
		}
	})

	t.Run("mint chain anchor must resolve the live linked branches", func(t *testing.T) {
		for _, tc := range []struct {
			name      string
			selects   string
			wantState string
			wantClass string
		}{
			{name: "selected root owns linked branch", selects: "legacy-a", wantState: record.Accepted},
			{name: "unselected root owns linked branch", selects: "legacy-b", wantState: record.Rejected, wantClass: "selects:unknown-target"},
		} {
			t.Run(tc.name, func(t *testing.T) {
				st, err := store.Open(t.TempDir())
				if err != nil {
					t.Fatalf("Open: %v", err)
				}
				for _, rec := range []record.Record{
					{Envelope: record.Envelope{RelayID: "legacy-a", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"record_kind": "seat_mint"}, Body: `{"seat":"branch-seat","role":"implementer","is_operator":false}`},
					{Envelope: record.Envelope{RelayID: "legacy-b", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"record_kind": "seat_mint"}, Body: `{"seat":"branch-seat","role":"implementer","is_operator":false}`},
					{Envelope: record.Envelope{RelayID: "linked-child", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"record_kind": "seat_mint", "mint_predecessor": "legacy-a"}, Body: `{"seat":"branch-seat","role":"implementer","is_operator":false}`},
				} {
					commitValidationRecord(t, st, rec)
				}
				candidate := MintChainAnchorRecord("branch-seat", tc.selects)
				candidate.Envelope.From = "operator"
				candidate.Envelope.Role = "operator"
				tab, err := tables.Build(st)
				if err != nil {
					t.Fatalf("tables.Build: %v", err)
				}
				admissionCandidate := candidate
				admissionCandidate.Envelope.DeliveryState = ""
				admission := validateRecordKind(tab, admissionCandidate, seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true})
				got, _, err := revalidateAtCommit(st, candidate, nil)
				if err != nil || got.Envelope.DeliveryState != tc.wantState {
					t.Fatalf("commit state=%s body=%q admission=%+v err=%v, want %s", got.Envelope.DeliveryState, got.Body, admission, err, tc.wantState)
				}
				if tc.wantClass == "" {
					if admission != nil {
						t.Fatalf("admission violation=%+v, want nil", admission)
					}
				} else if admission == nil || !strings.Contains(got.Body, tc.wantClass) {
					t.Fatalf("admission=%+v commit body=%q, want %s", admission, got.Body, tc.wantClass)
				}
			})
		}
	})
}

func commitValidationRecord(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, []store.Intent{}); err != nil {
		t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
	}
}
