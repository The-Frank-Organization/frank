// Package observe owns conductor-side observation and its positive write
// allowlist. Gate computes stamps and a terminal; it never mutates or delivers
// the candidate itself.
package observe

import (
	"encoding/json"

	"github.com/jackli/frank/internal/record"
)

const (
	Pass     = "pass"
	Fail     = "fail"
	Blocked  = "blocked"
	Degraded = "degraded"
)

type Candidate struct {
	Record    record.Record
	Seat      string
	Phase     string
	Authority string
}

type PredicateResult struct {
	ID             string
	Predicate      string
	ObservedFields map[string]string
}

type Env struct {
	PresentLayers map[string]bool
	Evaluate      func(Candidate) PredicateResult
}

type ObserveResult struct {
	PredicateResult  string
	Veto             string
	ObservedFields   map[string]string
	FailingPredicate string
}

var writableFields = map[string]bool{
	"ACTIONS_GIT_REF":          true,
	"FINAL_GIT_STATUS_SHORT":   true,
	"achieved_evidence":        true,
	"target_gap_result":        true,
	"evidence_integrity":       true,
	"record_integrity":         true,
	"executable_claim_results": true,
	"egress_scan_result":       true,
	"degradation_notes":        true,
	"attestation_source":       true,
	"deviated_observed":        true,
	"bucket_binding_observed":  true,
}

func Gate(cand record.Record, seat, phase, authority string, env Env) (ObserveResult, string) {
	if !env.PresentLayers["observe"] {
		return ObserveResult{}, record.Accepted
	}

	view := Candidate{
		Record:    cloneRecord(cand),
		Seat:      seat,
		Phase:     phase,
		Authority: authority,
	}
	verdict := PredicateResult{ID: "observe-unavailable", Predicate: Blocked}
	if env.Evaluate != nil {
		verdict = env.Evaluate(view)
	}
	if verdict.ID == "" {
		verdict.ID = "phase-done"
	}

	for field := range verdict.ObservedFields {
		if !writableFields[field] {
			return ObserveResult{
				PredicateResult:  Fail,
				Veto:             "block_delivery",
				ObservedFields:   baseStamps(cand, verdict.ID, Fail),
				FailingPredicate: "observe-write-allowlist",
			}, record.Rejected
		}
	}

	stamps := baseStamps(cand, verdict.ID, verdict.Predicate)
	for field, value := range verdict.ObservedFields {
		stamps[field] = value
	}
	result := ObserveResult{PredicateResult: verdict.Predicate, ObservedFields: stamps}
	switch verdict.Predicate {
	case Pass:
		return result, record.Accepted
	case Fail:
		result.Veto = "block_delivery"
		result.FailingPredicate = verdict.ID
		return result, record.Rejected
	case Blocked, Degraded:
		result.Veto = "block_delivery"
		result.FailingPredicate = verdict.ID
		return result, record.Rejected
	default:
		result.PredicateResult = Fail
		result.Veto = "block_delivery"
		result.FailingPredicate = "observe-predicate-result"
		return result, record.Rejected
	}
}

func baseStamps(cand record.Record, checkID, predicate string) map[string]string {
	outcome := "fail"
	achieved := "E0"
	integrity := "observed"
	if predicate == Pass {
		outcome = "pass"
		achieved = "E1"
	}
	if predicate == Blocked || predicate == Degraded {
		integrity = "self_reported"
	}
	targetGap := "not_applicable"
	if target := cand.Headers["EVIDENCE_TARGET"]; target != "" {
		if target == achieved {
			targetGap = "met"
		} else {
			targetGap = "target_gt_achieved"
		}
	}
	evidenceIntegrity, _ := json.Marshal(map[string]string{
		"ACTIONS_GIT_REF":        "observed",
		"FINAL_GIT_STATUS_SHORT": "observed",
	})
	claimRows, _ := json.Marshal([]map[string]string{{
		"claim_ref": "ACTIONS_GIT_REF",
		"check_id":  checkID,
		"outcome":   outcome,
	}})
	return map[string]string{
		"achieved_evidence":        achieved,
		"target_gap_result":        targetGap,
		"evidence_integrity":       string(evidenceIntegrity),
		"record_integrity":         integrity,
		"executable_claim_results": string(claimRows),
		"egress_scan_result":       "not_applicable",
		"degradation_notes":        "",
		"attestation_source":       "conductor",
		"deviated_observed":        "no",
		"bucket_binding_observed":  "no",
	}
}

func cloneRecord(in record.Record) record.Record {
	out := in
	out.Headers = cloneMap(in.Headers)
	out.XFields = cloneMap(in.XFields)
	return out
}

func cloneMap(in map[string]string) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}
