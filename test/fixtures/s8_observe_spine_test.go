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
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestS8ObserveGatePassesAndStampsInsideSubmit(t *testing.T) {
	st, reg, meta, renderEnv := s8ObserveDeps(t)
	predicateSaw := record.Record{}
	obsEnv := observe.Env{
		PresentLayers: renderEnv.PresentLayers,
		Evaluate: func(cand observe.Candidate) observe.PredicateResult {
			predicateSaw = cand.Record
			return observe.PredicateResult{ID: "git-ref-exists", Predicate: observe.Pass}
		},
	}
	handler := engine.SubmitHandlerWithObservation(st, reg, meta, renderEnv, obsEnv)
	rec, intents, err := handler(context.Background(), intake.Cmd{
		IntakeID: "observe-pass",
		Seat:     meta.Name,
		Role:     meta.Role,
		Payload:  s8ObservePayload(t, reg, meta, renderEnv, "branch@abc123"),
	})
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body = %s", rec.Envelope.DeliveryState, rec.Body)
	}
	if predicateSaw.Envelope.From != meta.Name || predicateSaw.Envelope.RelayID == "" {
		t.Fatalf("predicate did not see conductor-stamped candidate: %+v", predicateSaw.Envelope)
	}
	if predicateSaw.Headers["authority_class"] != "no" {
		t.Fatalf("observe gate did not read step-3 authority_class: %#v", predicateSaw.Headers)
	}
	if rec.Headers["achieved_evidence"] != "E1" || rec.Headers["record_integrity"] != "observed" || rec.Headers["attestation_source"] != "conductor" {
		t.Fatalf("observe stamps absent: %#v", rec.Headers)
	}
	if rec.Headers["surface_intent"] != "progress" {
		t.Fatalf("step-4.5 surface_intent = %q, want progress", rec.Headers["surface_intent"])
	}
	if len(intents) == 0 {
		t.Fatalf("accepted candidate has no delivery intents")
	}
}

func TestS8ObserveFalseActionClaimCommitsRejectedWithoutDelivery(t *testing.T) {
	st, reg, meta, renderEnv := s8ObserveDeps(t)
	obsEnv := observe.Env{
		PresentLayers: renderEnv.PresentLayers,
		Evaluate: func(observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{ID: "git-ref-exists", Predicate: observe.Fail}
		},
	}
	handler := engine.SubmitHandlerWithObservation(st, reg, meta, renderEnv, obsEnv)
	rec, intents, err := handler(context.Background(), intake.Cmd{
		IntakeID: "observe-fail",
		Seat:     meta.Name,
		Role:     meta.Role,
		Payload:  s8ObservePayload(t, reg, meta, renderEnv, "branch@missing"),
	})
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	if rec.Envelope.DeliveryState != record.Rejected || !strings.Contains(rec.Body, "git-ref-exists") {
		t.Fatalf("false predicate outcome = %s, body = %q", rec.Envelope.DeliveryState, rec.Body)
	}
	if intents != nil {
		t.Fatalf("rejected candidate produced delivery intents: %#v", intents)
	}
	relayID, err := st.Commit(rec, intents)
	if err != nil {
		t.Fatalf("commit terminal evidence: %v", err)
	}
	committed, err := st.Read(relayID)
	if err != nil {
		t.Fatalf("read terminal evidence: %v", err)
	}
	if committed.Envelope.DeliveryState != record.Rejected || committed.Headers["achieved_evidence"] != "E0" || committed.Headers["record_integrity"] != "observed" {
		t.Fatalf("committed terminal evidence = %#v", committed)
	}
	projected, err := st.Project("recipient.planner")
	if err != nil {
		t.Fatalf("project recipient: %v", err)
	}
	if len(projected) != 0 {
		t.Fatalf("rejected candidate delivered: %#v", projected)
	}
}

func TestS8ObserveGatePositiveAllowlistRejectsIdentityWrite(t *testing.T) {
	cand := record.Record{
		Envelope: record.Envelope{From: "seat-a", Role: "implementer", RelayID: "relay-a"},
		Headers: map[string]string{
			"PHASE": "SITREP", "AUTHORITY": "report-only", "slot_in": "opaque", "authority_class": "no",
		},
	}
	result, terminal := observe.Gate(cand, "seat-a", "SITREP", "report-only", observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{
				ID: "identity-write", Predicate: observe.Pass,
				ObservedFields: map[string]string{"FROM": "forged"},
			}
		},
	})
	if terminal != record.Rejected || result.FailingPredicate != "observe-write-allowlist" {
		t.Fatalf("terminal = %q, result = %#v", terminal, result)
	}
	if cand.Envelope.From != "seat-a" || cand.Headers["slot_in"] != "opaque" || cand.Headers["authority_class"] != "no" {
		t.Fatalf("gate mutated read-only candidate: %#v", cand)
	}
}

func s8ObserveDeps(t *testing.T) (*store.Store, *fieldspec.Registry, seat.SeatMeta, fieldspec.RenderEnv) {
	t.Helper()
	pinned := s8LoadPinnedSources(t, s8ConfigSources(t, true))
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	meta := seat.SeatMeta{Name: "s8.implementer", Role: "implementer"}
	env := fieldspec.RenderEnv{ConfigDigest: pinned.Digest, PresentLayers: map[string]bool{"store": true, "form": true, "lineage": true, "observe": true}}
	return st, pinned.Registry, meta, env
}

func s8ObservePayload(t *testing.T, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, gitRef string) []byte {
	t.Helper()
	headers := map[string]string{
		"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium",
		"EVIDENCE_TARGET": "E1", "SUBJECT": "observed submit", "ACTIONS_GIT_REF": gitRef,
		"FINAL_GIT_STATUS_SHORT": "none - clean tree",
	}
	_, digest := reg.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, headers["PHASE"], headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: record.Record{Headers: headers, Body: "done"}, FormDigest: digest})
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}
	return payload
}
