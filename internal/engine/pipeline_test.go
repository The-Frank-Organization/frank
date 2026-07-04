package engine_test

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestSubmitHandlerStampsAndAcceptsValidCandidate(t *testing.T) {
	st, reg := submitDeps(t)
	handler := engine.SubmitHandler(st, reg, seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"})
	payload := mustJSON(t, record.Record{
		Envelope: record.Envelope{RelayID: "candidate-1", From: "victim.planner", Role: "planner"},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "ok"},
		Body:     "hello",
	})
	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "i1", Seat: "s1-core.implementer", Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, want accepted", rec.Envelope.DeliveryState)
	}
	if rec.Envelope.From != "s1-core.implementer" || rec.Envelope.Role != "implementer" {
		t.Fatalf("identity not stamped: %+v", rec.Envelope)
	}
}

func TestSubmitHandlerRejectsForbiddenPairGrant(t *testing.T) {
	st, reg := submitDeps(t)
	handler := engine.SubmitHandler(st, reg, seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"})
	payload := mustJSON(t, record.Record{
		Envelope: record.Envelope{RelayID: "candidate-2", From: "s1-core.implementer", Role: "implementer"},
		Headers: map[string]string{
			"PHASE":     "SITREP",
			"AUTHORITY": "merge-gated",
			"grant":     "dispatch-impl",
			"SUBJECT":   "bad",
		},
	})
	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "i2", Seat: "s1-core.implementer", Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state = %s, want rejected", rec.Envelope.DeliveryState)
	}
	if rec.Envelope.From != "s1-core.implementer" {
		t.Fatalf("rejected record not stamped: %+v", rec.Envelope)
	}
}

func submitDeps(t *testing.T) (*store.Store, *fieldspec.Registry) {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	reg, err := fieldspec.Load(filepath.Join("..", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	return st, reg
}

func mustJSON(t *testing.T, v any) json.RawMessage {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return data
}
