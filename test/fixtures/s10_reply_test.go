package fixtures_test

import (
	"context"
	"encoding/json"
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

func TestS10OperatorReplyValidatesFrozenChoiceAndResolvesOnlyOnce(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	commitS10ParkRecord(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "reply-gate",
			From:          "seat-a",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":               "SITREP",
			"SUBJECT":             "reply gate",
			"HUMAN_GATE_REQUIRED": "yes",
		},
	})
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto: %v", err)
	}

	nonOperator := submitS10Reply(t, st, reg, seat.SeatMeta{Name: "seat-b", Role: "implementer"}, "approve", "reply-non-operator")
	if nonOperator.Envelope.DeliveryState != record.Rejected || !strings.Contains(nonOperator.Body, "record_kind:seat-scope") {
		t.Fatalf("non-operator reply = %s/%q, want rejected seat-scope", nonOperator.Envelope.DeliveryState, nonOperator.Body)
	}

	outOfSet := submitS10Reply(t, st, reg, seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "outside-frozen-set", "reply-out-of-set")
	if outOfSet.Envelope.DeliveryState != record.Rejected || !strings.Contains(outOfSet.Body, "choices:enum") {
		t.Fatalf("out-of-set reply = %s/%q, want rejected choices:enum", outOfSet.Envelope.DeliveryState, outOfSet.Body)
	}

	valid := submitS10Reply(t, st, reg, seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "approve", "reply-valid")
	if valid.Envelope.DeliveryState != record.Accepted || valid.Envelope.From != "operator" || valid.Envelope.To != "seat-a" {
		t.Fatalf("valid reply = %+v, want accepted operator-FROM wake to seat-a", valid)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto after reply: %v", err)
	}
	assertS10RelayCount(t, st, "odb-"+valid.Envelope.RelayID, 0)
	assertS10RelayCount(t, st, "park-"+valid.Envelope.RelayID, 0)
	if state := s10GateState(t, st, "reply-gate"); state != "replied_pending_validation" {
		t.Fatalf("gate state = %q, want replied_pending_validation", state)
	}

	second := submitS10Reply(t, st, reg, seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "reject", "reply-second")
	if second.Envelope.DeliveryState != record.Rejected || !strings.Contains(second.Body, "resolves_gate:already-resolved") {
		t.Fatalf("second reply = %s/%q, want rejected already-resolved", second.Envelope.DeliveryState, second.Body)
	}
}

func submitS10Reply(t *testing.T, st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, choice, intakeID string) record.Record {
	t.Helper()
	replyBody, err := json.Marshal(map[string]string{"choice": choice})
	if err != nil {
		t.Fatalf("Marshal reply: %v", err)
	}
	candidate := record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "operator gate resolution",
			"record_kind":     "gate_resolution",
			"parent_hint":     "reply-gate",
			"resolves_gate":   "reply-gate",
			"gate_category":   "authz_security",
		},
		Body: string(replyBody),
	}
	_, digest := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, "SITREP", "medium", fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: candidate, FormDigest: digest})
	if err != nil {
		t.Fatalf("Marshal payload: %v", err)
	}
	rec, intents, err := engine.SubmitHandler(st, reg, meta)(context.Background(), intake.Cmd{
		IntakeID:   intakeID,
		Seat:       meta.Name,
		Role:       meta.Role,
		IsOperator: meta.IsOperator,
		Payload:    payload,
	})
	if err != nil {
		t.Fatalf("SubmitHandler: %v", err)
	}
	if _, err := st.Commit(rec, intents); err != nil {
		t.Fatalf("Commit reply: %v", err)
	}
	return rec
}

func s10GateState(t *testing.T, st *store.Store, gateRef string) string {
	t.Helper()
	tables, err := obligation.BuildTables(st)
	if err != nil {
		t.Fatalf("BuildTables: %v", err)
	}
	for _, rec := range tables.Records {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["resolves_gate"] == gateRef {
			return "replied_pending_validation"
		}
	}
	if tables.ParkedLanes[gateRef] {
		return "parked_waiting_human"
	}
	return "active"
}
