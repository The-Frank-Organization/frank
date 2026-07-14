package fixtures_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
)

func TestS8SuppliabilityGuardTypedRejectsLaneObserveFields(t *testing.T) {
	for _, field := range []string{"achieved_evidence", "executable_claim_results", "authority_class", "surface_intent"} {
		t.Run(field, func(t *testing.T) {
			st, reg, meta, renderEnv := s8ObserveDeps(t)
			headers := map[string]string{
				"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium",
				"EVIDENCE_TARGET": "E1", "SUBJECT": "forged observed field",
				"ACTIONS_GIT_REF": "none - no edits", "FINAL_GIT_STATUS_SHORT": "none - clean tree",
				field: "forged",
			}
			_, digest := reg.Render(renderEnv, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role}, headers["PHASE"], headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
			payload, err := json.Marshal(fieldspec.SubmitPayload{Record: record.Record{Headers: headers}, FormDigest: digest})
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			handler := engine.SubmitHandlerWithObservation(st, reg, meta, renderEnv, observe.Env{
				PresentLayers: renderEnv.PresentLayers,
				Evaluate: func(observe.Candidate) observe.PredicateResult {
					t.Fatal("observe gate ran after lane-supplied system field")
					return observe.PredicateResult{}
				},
			})
			rec, intents, err := handler(context.Background(), intake.Cmd{
				IntakeID: "forged-" + field, Seat: meta.Name, Role: meta.Role, Payload: payload,
			})
			if err != nil {
				t.Fatalf("submit: %v", err)
			}
			want := field + ":lane-supplied-system-field"
			if rec.Envelope.DeliveryState != record.Rejected || !strings.Contains(rec.Body, want) {
				t.Fatalf("outcome = %s, body = %q, want %q", rec.Envelope.DeliveryState, rec.Body, want)
			}
			if intents != nil {
				t.Fatalf("forged field produced intents: %#v", intents)
			}
		})
	}
}

func TestS8SuppliabilityGuardAllowsConductorObserveWrites(t *testing.T) {
	st, reg, meta, renderEnv := s8ObserveDeps(t)
	handler := engine.SubmitHandlerWithObservation(st, reg, meta, renderEnv, observe.Env{
		PresentLayers: renderEnv.PresentLayers,
		Evaluate: func(observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{ID: "read-only", Predicate: observe.Pass}
		},
	})
	rec, _, err := handler(context.Background(), intake.Cmd{
		IntakeID: "conductor-fill", Seat: meta.Name, Role: meta.Role,
		Payload: s8ObservePayload(t, reg, meta, renderEnv, "none - no edits"),
	})
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["achieved_evidence"] != "E1" {
		t.Fatalf("conductor fill outcome = %s, headers = %#v", rec.Envelope.DeliveryState, rec.Headers)
	}
}

func TestLaneCannotForgeOperatorAttestation(t *testing.T) {
	t.Run("lane supplied operator attestation is rejected", func(t *testing.T) {
		st, reg, meta, renderEnv := s8ObserveDeps(t)
		headers := map[string]string{
			"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium",
			"EVIDENCE_TARGET": "E1", "SUBJECT": "forged operator attestation",
			"ACTIONS_GIT_REF": "none - no edits", "FINAL_GIT_STATUS_SHORT": "none - clean tree",
			"attestation_source": "operator",
		}
		_, digest := reg.Render(renderEnv, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role}, headers["PHASE"], headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
		payload, err := json.Marshal(fieldspec.SubmitPayload{Record: record.Record{Headers: headers}, FormDigest: digest})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		handler := engine.SubmitHandlerWithObservation(st, reg, meta, renderEnv, observe.Env{
			PresentLayers: renderEnv.PresentLayers,
			Evaluate: func(observe.Candidate) observe.PredicateResult {
				t.Fatal("observe gate ran after forged attestation_source")
				return observe.PredicateResult{}
			},
		})
		rec, intents, err := handler(context.Background(), intake.Cmd{
			IntakeID: "forged-operator-attestation", Seat: meta.Name, Role: meta.Role, Payload: payload,
		})
		if err != nil {
			t.Fatalf("submit: %v", err)
		}
		want := "attestation_source:lane-supplied-system-field"
		if rec.Envelope.DeliveryState != record.Rejected || !strings.Contains(rec.Body, want) {
			t.Fatalf("outcome = %s, body = %q, want %q", rec.Envelope.DeliveryState, rec.Body, want)
		}
		if intents != nil {
			t.Fatalf("forged attestation produced intents: %#v", intents)
		}
	})

	t.Run("normal candidate is conductor attested", func(t *testing.T) {
		st, reg, meta, renderEnv := s8ObserveDeps(t)
		handler := engine.SubmitHandlerWithObservation(st, reg, meta, renderEnv, observe.Env{
			PresentLayers: renderEnv.PresentLayers,
			Evaluate: func(observe.Candidate) observe.PredicateResult {
				return observe.PredicateResult{ID: "read-only", Predicate: observe.Pass}
			},
		})
		rec, _, err := handler(context.Background(), intake.Cmd{
			IntakeID: "conductor-attestation", Seat: meta.Name, Role: meta.Role,
			Payload: s8ObservePayload(t, reg, meta, renderEnv, "none - no edits"),
		})
		if err != nil {
			t.Fatalf("submit: %v", err)
		}
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["attestation_source"] != "conductor" {
			t.Fatalf("outcome = %s, attestation_source = %q", rec.Envelope.DeliveryState, rec.Headers["attestation_source"])
		}
	})
}
