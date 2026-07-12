// Package observe owns conductor-side observation and its positive write
// allowlist. Gate computes stamps and a terminal; it never mutates or delivers
// the candidate itself.
package observe

import (
	"encoding/json"
	"sort"

	"github.com/jackli/frank/internal/fieldspec"
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
	MachineryFault bool
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
	FailureClass     string
	Escalate         bool
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

// LaneSuppliedSystemField applies the header-layer face of the DEF-2 guard.
// Its caller is the authenticated lane submit path; conductor-produced fields
// are added only after this admission check.
func LaneSuppliedSystemField(reg *fieldspec.Registry, cand record.Record) string {
	fields := make([]string, 0, len(cand.Headers))
	for field, value := range cand.Headers {
		if value != "" {
			fields = append(fields, field)
		}
	}
	sort.Strings(fields)
	for _, field := range fields {
		spec, ok := reg.ByID(field)
		if !ok || spec.Layer != "header" {
			continue
		}
		if spec.Owner == "system" || spec.Owner == "computed" || spec.FillConstraints == "system_only" || spec.FillConstraints == "computed_result" {
			return field
		}
	}
	return ""
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
			verdict = PredicateResult{ID: "observe-write-allowlist", Predicate: Blocked, MachineryFault: true}
			break
		}
	}

	stamps := baseStamps(cand, verdict)
	for field, value := range verdict.ObservedFields {
		if !bindingField(field) {
			stamps[field] = value
		}
	}
	result := ObserveResult{PredicateResult: verdict.Predicate, ObservedFields: stamps}
	integrity := stamps["record_integrity"]
	if verdict.MachineryFault {
		return machineryFaultDisposition(cand, verdict, result)
	}
	switch verdict.Predicate {
	case Pass:
		if integrity == "self_reported" || integrity == "mixed" {
			return noVantageDisposition(cand, verdict, result)
		}
		return result, record.Accepted
	case Fail:
		result.Veto = "block_delivery"
		result.FailingPredicate = verdict.ID
		result.FailureClass = "observed-false"
		return result, record.Rejected
	case Blocked, Degraded:
		return noVantageDisposition(cand, verdict, result)
	default:
		result.PredicateResult = Fail
		result.Veto = "block_delivery"
		result.FailingPredicate = "observe-predicate-result"
		result.FailureClass = "observed-false"
		return result, record.Rejected
	}
}

func noVantageDisposition(cand record.Record, verdict PredicateResult, result ObserveResult) (ObserveResult, string) {
	result.ObservedFields["degradation_notes"] = "observation-unavailable"
	if cand.Headers["authority_class"] == "yes" {
		result.Veto = "block_delivery"
		result.FailingPredicate = verdict.ID
		result.FailureClass = "observation-unavailable"
		result.Escalate = true
		return result, record.Held
	}
	return result, record.Accepted
}

func machineryFaultDisposition(cand record.Record, verdict PredicateResult, result ObserveResult) (ObserveResult, string) {
	result.Veto = "block_delivery"
	result.FailingPredicate = verdict.ID
	result.FailureClass = "observe-machinery-fault"
	result.ObservedFields["degradation_notes"] = "observe-machinery-fault"
	if cand.Headers["authority_class"] == "yes" {
		result.Escalate = true
		return result, record.Held
	}
	return result, record.Rejected
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
