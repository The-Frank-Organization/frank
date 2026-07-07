package engine_test

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
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestSubmitHandlerStampsAndAcceptsValidCandidate(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload := submitPayload(t, reg, meta, record.Record{
		Envelope: record.Envelope{RelayID: "candidate-1", From: "victim.planner", Role: "planner"},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "EVIDENCE_TARGET": "E1", "SUBJECT": "ok"},
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
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "EVIDENCE_TARGET": "E1", "SUBJECT": "projection"},
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

func TestSubmitHandlerAcceptsOfferedRecordKindsWithoutEngineMembershipSwitch(t *testing.T) {
	cases := []struct {
		name       string
		meta       seat.SeatMeta
		recordKind string
	}{
		{
			name:       "operator disposition",
			meta:       seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true},
			recordKind: "disposition",
		},
		{
			name:       "plain diagnostics",
			meta:       seat.SeatMeta{Name: "seat-a", Role: "implementer"},
			recordKind: "diagnostics",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			st, reg := submitDeps(t)
			handler := engine.SubmitHandler(st, reg, tc.meta)
			payload := submitPayload(t, reg, tc.meta, record.Record{
				Headers: map[string]string{
					"PHASE":           "SITREP",
					"AUTHORITY":       "report-only",
					"CEREMONY_TIER":   "medium",
					"EVIDENCE_TARGET": "E1",
					"SUBJECT":         tc.name,
					"record_kind":     tc.recordKind,
				},
			})

			rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "kind-" + tc.recordKind, Seat: tc.meta.Name, Role: tc.meta.Role, IsOperator: tc.meta.IsOperator, Payload: payload})
			if err != nil {
				t.Fatalf("handler: %v", err)
			}
			if rec.Envelope.DeliveryState != record.Accepted {
				t.Fatalf("state = %s, body = %q; want accepted", rec.Envelope.DeliveryState, rec.Body)
			}
		})
	}
}

func TestSubmitHandlerRejectsGenesisAtSeatScopeNotEngineUnknown(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "seat-a", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload := submitPayload(t, reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "genesis",
			"record_kind":     "genesis",
		},
	})

	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "kind-genesis", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state = %s, want rejected", rec.Envelope.DeliveryState)
	}
	if !strings.Contains(rec.Body, "record_kind:seat-scope") || strings.Contains(rec.Body, "record_kind:unknown") {
		t.Fatalf("body = %q, want seat-scope rejection without engine unknown", rec.Body)
	}
}

func TestSubmitHandlerRejectsUnknownRecordKindAtMembership(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "seat-a", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload := submitPayload(t, reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "unknown",
			"record_kind":     "not_a_record_kind",
		},
	})

	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "kind-unknown", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state = %s, want rejected", rec.Envelope.DeliveryState)
	}
	if !strings.Contains(rec.Body, "record_kind:enum") || strings.Contains(rec.Body, "record_kind:unknown") {
		t.Fatalf("body = %q, want enum membership rejection without engine unknown", rec.Body)
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
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "EVIDENCE_TARGET": "E1", "SUBJECT": "verdict 1", "PARENT_DISPATCH_ID": "gate-1", "resolves_gate": "gate-1"},
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
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "EVIDENCE_TARGET": "E1", "SUBJECT": "verdict 2", "PARENT_DISPATCH_ID": "gate-1", "resolves_gate": "gate-1"},
	})
	second, _, err := handler(context.Background(), intake.Cmd{IntakeID: "v2", Seat: "operator", Role: "operator", IsOperator: true, Payload: secondPayload})
	if err != nil {
		t.Fatalf("second handler: %v", err)
	}
	if second.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("second state = %s, want rejected", second.Envelope.DeliveryState)
	}
}

func TestSubmitHandlerBuildsOwedProjectionFromProvidedTable(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	owed := record.Record{
		Envelope: record.Envelope{RelayID: "owed-base", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":            "SITREP",
			"AUTHORITY":        "report-only",
			"SUBJECT":          "owed",
			"record_kind":      "owed_item",
			"owner":            "s1",
			"source":           "review",
			"target_surface":   "live table",
			"disposition_path": "fold",
		},
	}
	if _, err := st.Commit(owed, nil); err != nil {
		t.Fatalf("Commit owed: %v", err)
	}
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("Build tables: %v", err)
	}
	if err := os.WriteFile(filepath.Join(st.Root, "records", "owed-base.json"), []byte(`{"corrupt":true}`), 0o644); err != nil {
		t.Fatalf("corrupt owed record: %v", err)
	}

	handler := engine.SubmitHandler(st, reg, meta, tab)
	payload := submitPayload(t, reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "owed disposition",
			"record_kind":     "owed_disposition",
			"disposes_owed":   "owed-base",
		},
	})
	rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "owed-disposition", Seat: "operator", Role: "operator", IsOperator: true, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, want accepted: %s", rec.Envelope.DeliveryState, rec.Body)
	}
	for _, intent := range intents {
		if intent.Kind == store.IntentRender && intent.Path == filepath.Join("owed", "OPEN.md") {
			return
		}
	}
	t.Fatalf("missing owed OPEN.md render intent in %#v", intents)
}

func TestSubmitHandlerRejectsParentOutsideRenderedCandidateSet(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "s3-form.implementer", Role: "implementer"}
	tab := tables.New()
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "visible-but-unrelated", DispatchID: "unrelated", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "PLAN", "SUBJECT": "visible"},
	})
	env := fieldspec.RenderEnv{ParentCandidates: func(fieldspec.SeatMeta) ([]string, string) {
		return []string{"allowed-parent"}, "allowed-parent"
	}}
	handler := engine.SubmitHandlerWithRender(st, reg, meta, env, tab)
	payload := submitPayloadWithEnv(t, reg, meta, env, record.Record{
		Headers: map[string]string{
			"PHASE":              "PLAN",
			"AUTHORITY":          "plan-only",
			"SUBJECT":            "outside parent",
			"PARENT_DISPATCH_ID": "visible-but-unrelated",
		},
	})
	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "outside-parent", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state = %s, want rejected", rec.Envelope.DeliveryState)
	}
	if !strings.Contains(rec.Body, "outside-active-lineage") {
		t.Fatalf("body = %s, want outside-active-lineage", rec.Body)
	}
}

func TestSubmitHandlerRejectsStalePositiveParentAfterRender(t *testing.T) {
	st, reg := submitDeps(t)
	meta := seat.SeatMeta{Name: "s3-form.implementer", Role: "implementer"}
	env := fieldspec.RenderEnv{ParentCandidates: func(fieldspec.SeatMeta) ([]string, string) {
		return []string{"parent-at-render"}, "parent-at-render"
	}}
	payload := submitPayloadWithEnv(t, reg, meta, env, record.Record{
		Headers: map[string]string{
			"PHASE":              "PLAN",
			"AUTHORITY":          "plan-only",
			"EVIDENCE_TARGET":    "E1",
			"SUBJECT":            "stale parent",
			"PARENT_DISPATCH_ID": "parent-at-render",
		},
	})

	handler := engine.SubmitHandlerWithRender(st, reg, meta, env, tables.New())
	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "stale-parent", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state = %s, want rejected", rec.Envelope.DeliveryState)
	}
	if !strings.Contains(rec.Body, lineage.ParentUnknownRecompose) {
		t.Fatalf("body = %s, want parent substrate bounce", rec.Body)
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

func submitPayloadWithEnv(t *testing.T, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, rec record.Record) json.RawMessage {
	t.Helper()
	_, digest := reg.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	return mustJSON(t, fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
}
