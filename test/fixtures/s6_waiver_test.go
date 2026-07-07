package fixtures_test

import (
	"context"
	"encoding/json"
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

func TestS6WaiverRowsAbsentFromNonOperatorRenderAndRejectedOnSubmit(t *testing.T) {
	reg := loadAssemblyRegistry(t)
	plainSeat := fieldspec.SeatMeta{Name: "s6-core.implementer", Role: "implementer"}
	operatorSeat := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	for _, id := range []string{"rationale", "waiver_scope", "retracts"} {
		plainForm, _ := reg.Render(fieldspec.RenderEnv{}, plainSeat, "SITREP", "medium", fieldspec.ClosedGrantState)
		if plainForm.HasField(id) {
			t.Fatalf("non-operator form rendered %s", id)
		}
		operatorForm, _ := reg.Render(fieldspec.RenderEnv{}, operatorSeat, "SITREP", "medium", fieldspec.ClosedGrantState)
		if !operatorForm.HasField(id) {
			t.Fatalf("operator form omitted %s", id)
		}
	}

	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	meta := seat.SeatMeta{Name: "s6-core.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	payload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":        "SITREP",
			"AUTHORITY":    "report-only",
			"SUBJECT":      "bad waiver rows",
			"rationale":    "not operator",
			"waiver_scope": `{"kind":"run"}`,
		},
	}))
	got, _, err := handler(context.Background(), intake.Cmd{IntakeID: "non-operator-waiver", Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if got.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state/body = %s/%s, want rejected", got.Envelope.DeliveryState, got.Body)
	}
	if !strings.Contains(got.Body, "waiver_scope:seat-scope") {
		t.Fatalf("body = %q, want waiver_scope:seat-scope", got.Body)
	}
}

func TestS6ScopedWaiverPassesOnlyInScopeAndRetractionReArms(t *testing.T) {
	cand := record.Record{
		Envelope: record.Envelope{RelayID: "candidate-d1", DispatchID: "dispatch-1", From: "s6.orchestrator-planner", Role: "orchestrator-planner"},
		Headers:  map[string]string{"PHASE": "DESIGN", "TO": "s6-core.planner"},
	}
	meta := seat.SeatMeta{Name: "s6.orchestrator-planner", Role: "orchestrator-planner"}
	tab := tables.New()
	eng := lineage.Engine{T: tab}
	if bounce := eng.Check(cand, meta); bounce == nil || bounce.Kind != lineage.ReviewerVisibilityMissing {
		t.Fatalf("missing waiver bounce = %+v", bounce)
	}

	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "waiver-d1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "waiver", "rationale": "narrow", "waiver_scope": `{"kind":"dispatch","dispatch_id":"dispatch-1"}`},
	})
	if bounce := eng.Check(cand, meta); bounce != nil {
		t.Fatalf("in-scope waiver bounced: %+v", bounce)
	}
	outOfScope := cand
	outOfScope.Envelope.DispatchID = "dispatch-2"
	if bounce := eng.Check(outOfScope, meta); bounce == nil || bounce.Kind != lineage.ReviewerVisibilityMissing {
		t.Fatalf("out-of-scope waiver bounce = %+v", bounce)
	}

	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "retract-d1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "retract", "record_kind": "waiver_retraction", "retracts": "waiver-d1"},
	})
	if bounce := eng.Check(cand, meta); bounce == nil || bounce.Kind != lineage.ReviewerVisibilityMissing {
		t.Fatalf("post-retraction bounce = %+v", bounce)
	}
}

func TestS6LegacyUnscopedWaiverRunWideUntilRetracted(t *testing.T) {
	cand := record.Record{
		Envelope: record.Envelope{RelayID: "candidate", DispatchID: "dispatch-1", From: "s6.orchestrator-planner", Role: "orchestrator-planner"},
		Headers:  map[string]string{"PHASE": "DESIGN", "TO": "s6-core.planner"},
	}
	meta := seat.SeatMeta{Name: "s6.orchestrator-planner", Role: "orchestrator-planner"}
	tab := tables.New()
	eng := lineage.Engine{T: tab}
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "legacy-waiver", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "legacy", "ORCH_REVIEW_WAIVER": "operator approved no reviewer"},
	})
	if bounce := eng.Check(cand, meta); bounce != nil {
		t.Fatalf("legacy waiver bounced: %+v", bounce)
	}
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "legacy-retract", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "legacy retract", "record_kind": "waiver_retraction", "retracts": "legacy-waiver"},
	})
	if bounce := eng.Check(cand, meta); bounce == nil || bounce.Kind != lineage.ReviewerVisibilityMissing {
		t.Fatalf("retracted legacy waiver bounce = %+v", bounce)
	}
}

func TestS6WaiverRetractionLayer3SubmitChecks(t *testing.T) {
	reg := loadAssemblyRegistry(t)
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	handler := engine.SubmitHandler(st, reg, meta)

	waiverPayload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":        "SITREP",
			"AUTHORITY":    "report-only",
			"SUBJECT":      "waiver",
			"rationale":    "narrow",
			"waiver_scope": `{"kind":"run"}`,
		},
	}))
	waiver, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "waiver", Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: waiverPayload})
	if err != nil {
		t.Fatalf("waiver handler: %v", err)
	}
	if waiver.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("waiver state/body = %s/%s", waiver.Envelope.DeliveryState, waiver.Body)
	}
	if _, err := st.Commit(waiver, intents); err != nil {
		t.Fatalf("Commit waiver: %v", err)
	}

	retractPayload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":       "SITREP",
			"AUTHORITY":   "report-only",
			"SUBJECT":     "retract",
			"record_kind": "waiver_retraction",
			"retracts":    waiver.Envelope.RelayID,
		},
	}))
	retract, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "retract", Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: retractPayload})
	if err != nil {
		t.Fatalf("retract handler: %v", err)
	}
	if retract.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("retract state/body = %s/%s", retract.Envelope.DeliveryState, retract.Body)
	}
	if _, err := st.Commit(retract, intents); err != nil {
		t.Fatalf("Commit retract: %v", err)
	}

	dupe, _, err := handler(context.Background(), intake.Cmd{IntakeID: "retract-dupe", Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: retractPayload})
	if err != nil {
		t.Fatalf("dupe handler: %v", err)
	}
	if dupe.Envelope.DeliveryState != record.Rejected || !strings.Contains(dupe.Body, "retracts:already-resolved") {
		t.Fatalf("dupe state/body = %s/%s, want already-resolved", dupe.Envelope.DeliveryState, dupe.Body)
	}

	unknownPayload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, record.Record{
		Headers: map[string]string{
			"PHASE":       "SITREP",
			"AUTHORITY":   "report-only",
			"SUBJECT":     "unknown retract",
			"record_kind": "waiver_retraction",
			"retracts":    "missing-waiver",
		},
	}))
	unknown, _, err := handler(context.Background(), intake.Cmd{IntakeID: "retract-unknown", Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: unknownPayload})
	if err != nil {
		t.Fatalf("unknown handler: %v", err)
	}
	if unknown.Envelope.DeliveryState != record.Rejected || !strings.Contains(unknown.Body, lineage.ParentUnknownRecompose) {
		t.Fatalf("unknown state/body = %s/%s, want unknown waiver", unknown.Envelope.DeliveryState, unknown.Body)
	}
}
