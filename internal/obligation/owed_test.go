package obligation_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestOwedItemSubmitProjectsOpenAndDispositionClosesIt(t *testing.T) {
	root := t.TempDir()
	st, reg := owedDeps(t, root)
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	handler := engine.SubmitHandler(st, reg, meta)

	owed := submitAndCommitOwed(t, st, handler, reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":            "SITREP",
			"AUTHORITY":        "report-only",
			"EVIDENCE_TARGET":  "E1",
			"SUBJECT":          "owed",
			"record_kind":      "owed_item",
			"owner":            "s1",
			"source":           "guide",
			"target_surface":   "F11 full classxpoint sweep",
			"disposition_path": "S2 exit gate",
		},
	})
	open := string(mustReadOwed(t, filepath.Join(root, "projections", "owed", "OPEN.md")))
	for _, want := range []string{owed.Envelope.RelayID, "s1", "F11 full classxpoint sweep", "S2 exit gate"} {
		if !strings.Contains(open, want) {
			t.Fatalf("OPEN.md missing %q:\n%s", want, open)
		}
	}
	facts, err := obligation.OpenOwed(st)
	if err != nil {
		t.Fatalf("OpenOwed: %v", err)
	}
	if len(facts) != 1 || facts[0].Key != owed.Envelope.RelayID {
		t.Fatalf("open owed = %+v, want %s", facts, owed.Envelope.RelayID)
	}

	disposition := submitAndCommitOwed(t, st, handler, reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "owed disposition",
			"record_kind":     "owed_disposition",
			"disposes_owed":   owed.Envelope.RelayID,
		},
	})
	if disposition.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("disposition state = %s, want accepted", disposition.Envelope.DeliveryState)
	}
	open = string(mustReadOwed(t, filepath.Join(root, "projections", "owed", "OPEN.md")))
	if strings.Contains(open, owed.Envelope.RelayID) {
		t.Fatalf("OPEN.md still contains disposed owed id:\n%s", open)
	}

	second := submitOwed(t, handler, reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "second disposition",
			"record_kind":     "owed_disposition",
			"disposes_owed":   owed.Envelope.RelayID,
		},
	})
	if second.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("second disposition state = %s, want rejected", second.Envelope.DeliveryState)
	}
}

func TestOwedValidationRejectsUnknownKindsAndUnknownParents(t *testing.T) {
	st, reg := owedDeps(t, t.TempDir())
	meta := seat.SeatMeta{Name: "seat-a", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)

	unknownKind := submitOwed(t, handler, reg, meta, record.Record{
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "bad", "record_kind": "mystery"},
	})
	if unknownKind.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("unknown kind state = %s, want rejected", unknownKind.Envelope.DeliveryState)
	}
	unknownParent := submitOwed(t, handler, reg, meta, record.Record{
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "bad parent", "record_kind": "owed_disposition", "disposes_owed": "missing"},
	})
	if unknownParent.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("unknown parent state = %s, want rejected", unknownParent.Envelope.DeliveryState)
	}
}

func TestOwedItemRejectsNonOperatorSeat(t *testing.T) {
	root := t.TempDir()
	st, reg := owedDeps(t, root)
	meta := seat.SeatMeta{Name: "s2-core.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	owed := submitOwed(t, handler, reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":            "SITREP",
			"AUTHORITY":        "report-only",
			"EVIDENCE_TARGET":  "E1",
			"SUBJECT":          "owed from implementer",
			"record_kind":      "owed_item",
			"owner":            "s2",
			"source":           "review-fold",
			"target_surface":   "operator-channel e2e",
			"disposition_path": "fold report",
		},
	})
	if owed.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("owed state = %s, want rejected", owed.Envelope.DeliveryState)
	}
	if !strings.Contains(owed.Body, "record_kind") {
		t.Fatalf("owed body = %q, want record_kind rejection", owed.Body)
	}
}

func TestRebuildCreatesEmptyOwedProjection(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if err := st.RebuildProjections(); err != nil {
		t.Fatalf("RebuildProjections: %v", err)
	}
	open := string(mustReadOwed(t, filepath.Join(root, "projections", "owed", "OPEN.md")))
	if !strings.Contains(open, "relay_id") {
		t.Fatalf("empty OPEN.md missing header:\n%s", open)
	}
}

func owedDeps(t *testing.T, root string) (*store.Store, *fieldspec.Registry) {
	t.Helper()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	reg, err := fieldspec.Load(filepath.Join("..", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	return st, reg
}

func submitAndCommitOwed(t *testing.T, st *store.Store, handler engine.Handler, reg *fieldspec.Registry, meta seat.SeatMeta, rec record.Record) record.Record {
	t.Helper()
	cand := submitOwed(t, handler, reg, meta, rec)
	if cand.Envelope.DeliveryState == record.Accepted {
		if _, err := st.Commit(cand, store.OwedProjectionIntentsForCandidate(st, cand)); err != nil {
			t.Fatalf("Commit owed candidate: %v", err)
		}
	}
	return cand
}

func submitOwed(t *testing.T, handler engine.Handler, reg *fieldspec.Registry, meta seat.SeatMeta, rec record.Record) record.Record {
	t.Helper()
	_, digest := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	cand, _, err := handler(context.Background(), intake.Cmd{IntakeID: "owed-intake", Seat: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	return cand
}

func mustReadOwed(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return data
}
