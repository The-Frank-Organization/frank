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
	Verdicts       []CheckVerdict
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
				ObservedFields:   baseStamps(cand, PredicateResult{ID: verdict.ID, Predicate: Fail}),
				FailingPredicate: "observe-write-allowlist",
			}, record.Rejected
		}
	}

	stamps := baseStamps(cand, verdict)
	for field, value := range verdict.ObservedFields {
		if !bindingField(field) {
			stamps[field] = value
		}
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

func baseStamps(cand record.Record, result PredicateResult) map[string]string {
	outcome := "fail"
	achieved := "E0"
	integrity := "observed"
	if result.Predicate == Pass {
		outcome = "pass"
		achieved = "E1"
	}
	if result.Predicate == Blocked || result.Predicate == Degraded {
		integrity = "self_reported"
	}
	claimRows := []map[string]string{}
	integrityFields := map[string]string{
		"ACTIONS_GIT_REF":        "observed",
		"FINAL_GIT_STATUS_SHORT": "observed",
	}
	if len(result.Verdicts) > 0 {
		achieved = "E0"
		integrityFields = map[string]string{}
		observed, selfReported := false, false
		for _, verdict := range result.Verdicts {
			claimRows = append(claimRows, map[string]string{
				"claim_ref": verdict.ClaimRef,
				"check_id":  verdict.CheckID,
				"outcome":   verdict.Outcome,
			})
			if verdict.Outcome == "pass" && evidenceRank(verdict.RungReached) > evidenceRank(achieved) {
				achieved = verdict.RungReached
			}
			if verdict.Outcome == "pass" || verdict.Outcome == "fail" {
				observed = true
				integrityFields[verdict.ClaimRef] = "observed"
			} else {
				selfReported = true
				integrityFields[verdict.ClaimRef] = "self_reported"
			}
		}
		switch {
		case observed && selfReported:
			integrity = "mixed"
		case observed:
			integrity = "observed"
		default:
			integrity = "self_reported"
		}
	} else {
		claimRows = append(claimRows, map[string]string{
			"claim_ref": "ACTIONS_GIT_REF",
			"check_id":  result.ID,
			"outcome":   outcome,
		})
	}
	targetGap := "not_applicable"
	if target := cand.Headers["EVIDENCE_TARGET"]; target != "" {
		if evidenceRank(achieved) >= evidenceRank(target) {
			targetGap = "met"
		} else {
			targetGap = "target_gt_achieved"
		}
	}
	evidenceIntegrity, _ := json.Marshal(integrityFields)
	encodedClaimRows, _ := json.Marshal(claimRows)
	return map[string]string{
		"achieved_evidence":        achieved,
		"target_gap_result":        targetGap,
		"evidence_integrity":       string(evidenceIntegrity),
		"record_integrity":         integrity,
		"executable_claim_results": string(encodedClaimRows),
		"egress_scan_result":       "not_applicable",
		"degradation_notes":        "",
		"attestation_source":       "conductor",
		"deviated_observed":        "no",
		"bucket_binding_observed":  "no",
	}
}

func bindingField(field string) bool {
	switch field {
	case "achieved_evidence", "target_gap_result", "evidence_integrity", "record_integrity", "executable_claim_results":
		return true
	default:
		return false
	}
}

func evidenceRank(value string) int {
	switch value {
	case "E1":
		return 1
	case "E2":
		return 2
	case "E3":
		return 3
	case "E4":
		return 4
	default:
		return 0
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
