package fixtures_test

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

func TestS6ParentingRegistryRemovesParentFromAllRenderedForms(t *testing.T) {
	reg := loadAssemblyRegistry(t)
	seats := []fieldspec.SeatMeta{
		{Name: "operator", Role: "operator", IsOperator: true},
		{Name: "s6-core.planner", Role: "planner"},
		{Name: "s6-core.implementer", Role: "implementer"},
		{Name: "s6.orchestrator-planner", Role: "orchestrator-planner"},
	}
	for _, meta := range seats {
		for _, phase := range reg.NamedEnums["PHASE"] {
			form, _ := reg.Render(fieldspec.RenderEnv{}, meta, phase, "medium", func(fieldspec.SeatMeta) bool { return true })
			if form.HasField("PARENT_DISPATCH_ID") {
				t.Fatalf("%s/%s rendered PARENT_DISPATCH_ID", meta.Name, phase)
			}
		}
	}
}

func TestS6ParentingWokenOnStampIsReplayStable(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	parent := record.Record{
		Envelope: record.Envelope{
			RelayID:       "wake-parent",
			DispatchID:    "dispatch-a",
			From:          "s6-core.planner",
			To:            "s6-core.implementer",
			Role:          "planner",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "wake"},
	}
	if _, err := st.Commit(parent, nil); err != nil {
		t.Fatalf("Commit parent: %v", err)
	}
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("Build tables: %v", err)
	}
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	meta := seat.SeatMeta{Name: "s6-core.implementer", Role: "implementer"}
	env := fieldspec.RenderEnv{Turn: fieldspec.TurnContext{WokenOn: "wake-parent", ActiveDispatch: "dispatch-a"}}
	payload := s6ParentSubmitPayload(t, reg, meta, env, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "replay stable",
		},
	})
	handler := engine.SubmitHandlerWithRender(st, reg, meta, env, tab)

	first, _, err := handler(context.Background(), intake.Cmd{IntakeID: "first", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("first handler: %v", err)
	}
	second, _, err := handler(context.Background(), intake.Cmd{IntakeID: "second", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("second handler: %v", err)
	}
	for _, got := range []record.Record{first, second} {
		if got.Envelope.DeliveryState != record.Accepted {
			t.Fatalf("state/body = %s/%s, want accepted", got.Envelope.DeliveryState, got.Body)
		}
		if got.Headers["PARENT_DISPATCH_ID"] != "wake-parent" || got.Headers["parent_provenance"] != "woken_on" {
			t.Fatalf("parent stamp = %+v", got.Headers)
		}
	}
	if first.Headers["PARENT_DISPATCH_ID"] != second.Headers["PARENT_DISPATCH_ID"] ||
		first.Headers["parent_provenance"] != second.Headers["parent_provenance"] {
		t.Fatalf("parent stamp changed across replay: first=%+v second=%+v", first.Headers, second.Headers)
	}
}

func s6ParentSubmitPayload(t *testing.T, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, rec record.Record) []byte {
	t.Helper()
	_, digest := reg.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	if err != nil {
		t.Fatalf("marshal submit payload: %v", err)
	}
	return payload
}
