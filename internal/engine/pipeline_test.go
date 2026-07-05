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
	meta := seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload := submitPayload(t, reg, meta, record.Record{
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

func TestSubmitHandlerStampsCurrentSchemaVersion(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload := submitPayload(t, reg, meta, record.Record{
		Envelope: record.Envelope{SchemaVersion: 99},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "schema"},
	})
	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "schema", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.SchemaVersion != 1 {
		t.Fatalf("schema version = %d, want current", rec.Envelope.SchemaVersion)
	}
}

func TestSubmitHandlerAssignsRelayIDAndProjectionIntents(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload := submitPayload(t, reg, meta, record.Record{
		Envelope: record.Envelope{
			RelayID:    "client-picked",
			DispatchID: "dispatch-1",
			From:       "victim.planner",
			To:         "s1-core.planner",
			Role:       "planner",
		},
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "projection"},
		Body:    "hello",
	})
	rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "i-proj", Seat: "s1-core.implementer", Role: "implementer", Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, want accepted", rec.Envelope.DeliveryState)
	}
	if rec.Envelope.RelayID == "" || rec.Envelope.RelayID == "client-picked" {
		t.Fatalf("relay ID was not server assigned: %q", rec.Envelope.RelayID)
	}
	requireIntent(t, intents, store.IntentIndex)
	requireIntent(t, intents, store.IntentRender)
	requireIntent(t, intents, store.IntentMailbox)
}

func TestSubmitHandlerRejectsForbiddenPairGrant(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload := submitPayload(t, reg, meta, record.Record{
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

func TestOperatorVerdictOneShotRunsThroughSubmitHandler(t *testing.T) {
	st, reg := submitDeps(t)
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "gate-1", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes"},
	}, nil); err != nil {
		t.Fatalf("Commit gate: %v", err)
	}
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	handler := engine.SubmitHandler(st, reg, meta)

	firstPayload := submitPayload(t, reg, meta, record.Record{
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "verdict 1", "PARENT_DISPATCH_ID": "gate-1", "resolves_gate": "gate-1"},
	})
	first, _, err := handler(context.Background(), intake.Cmd{IntakeID: "v1", Seat: "operator", Role: "operator", IsOperator: true, Payload: firstPayload})
	if err != nil {
		t.Fatalf("first handler: %v", err)
	}
	if first.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("first state = %s, want accepted", first.Envelope.DeliveryState)
	}
	if _, err := st.Commit(first, nil); err != nil {
		t.Fatalf("commit first: %v", err)
	}

	secondPayload := submitPayload(t, reg, meta, record.Record{
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "verdict 2", "PARENT_DISPATCH_ID": "gate-1", "resolves_gate": "gate-1"},
	})
	second, _, err := handler(context.Background(), intake.Cmd{IntakeID: "v2", Seat: "operator", Role: "operator", IsOperator: true, Payload: secondPayload})
	if err != nil {
		t.Fatalf("second handler: %v", err)
	}
	if second.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("second state = %s, want rejected", second.Envelope.DeliveryState)
	}
}

func requireIntent(t *testing.T, intents []store.Intent, kind string) {
	t.Helper()
	for _, intent := range intents {
		if intent.Kind == kind {
			return
		}
	}
	t.Fatalf("missing %s intent in %#v", kind, intents)
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

func submitPayload(t *testing.T, reg *fieldspec.Registry, meta seat.SeatMeta, rec record.Record) json.RawMessage {
	t.Helper()
	_, digest := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	return mustJSON(t, fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
}
