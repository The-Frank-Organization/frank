package observe

import "strings"

type verdictOrigin string

const (
	originConductorConfig    verdictOrigin = "conductor-config"
	originConductorPolicy    verdictOrigin = "conductor-policy"
	originConductorMachinery verdictOrigin = "conductor-machinery"
	originBaseCheck          verdictOrigin = "base-check"
	originExecutor           verdictOrigin = "executor"
)

type boundVerdict struct {
	Verdict  CheckVerdict
	Origin   verdictOrigin
	Entry    CheckEntry
	Selected Selection
}

type validatedVerdict struct {
	Verdict        CheckVerdict
	MachineryFault bool
	SignalClass    string
	Integrity      string
}

type claimBinding struct {
	RungReached string
	SignalClass string
	Integrity   string
}

func validateBoundVerdict(bound boundVerdict) validatedVerdict {
	expectedID := bound.Selected.CheckID
	if expectedID == "" {
		expectedID = bound.Entry.ID
	}
	expectedClaim := bound.Selected.ClaimRef
	if expectedClaim == "" {
		expectedClaim = bound.Verdict.ClaimRef
	}
	if bound.Selected.CheckID != "" && (bound.Verdict.CheckID != expectedID || bound.Verdict.ClaimRef != expectedClaim) {
		return verdictBindingFault(bound, "check-machinery-verdict-identity-mismatch")
	}

	verdict := bound.Verdict
	if verdict.RungReached == "" {
		verdict.RungReached = "none"
	}
	switch bound.Origin {
	case originConductorConfig, originConductorPolicy:
		verdict.Timing = "not-completed"
	case originConductorMachinery:
		if verdict.Timing == "timeout" || strings.Contains(verdict.FailingDetail, "timeout") {
			verdict.Timing = "timeout"
		} else {
			verdict.Timing = "not-completed"
		}
	case originBaseCheck:
		if timing, ok := baseRefusalTiming(verdict.FailingDetail); ok {
			verdict.Timing = timing
		} else {
			verdict.Timing = "under-timeout"
		}
	}
	bound.Verdict = verdict
	if !validOriginTuple(bound) {
		return verdictBindingFault(bound, "check-machinery-verdict-tuple-invalid")
	}
	if !validOriginClass(bound.Origin, verdict) {
		return verdictBindingFault(bound, "check-machinery-verdict-origin-class-mismatch")
	}
	machinery := bound.Origin == originConductorConfig || bound.Origin == originConductorMachinery || bound.Origin == originExecutor && strings.HasPrefix(verdict.FailingDetail, "executor-")
	integrity := "self_reported"
	if verdict.Outcome == "pass" || verdict.Outcome == "fail" {
		integrity = "observed"
	}
	return validatedVerdict{
		Verdict: verdict, MachineryFault: machinery,
		SignalClass: signalClass(expectedID, verdict.Outcome), Integrity: integrity,
	}
}

func validOriginTuple(bound boundVerdict) bool {
	v := bound.Verdict
	none := v.RungReached == "none"
	switch bound.Origin {
	case originConductorConfig, originConductorPolicy:
		return v.Outcome == "unsafe" && v.Predicate == Blocked && none && v.Timing == "not-completed"
	case originConductorMachinery:
		return (v.Outcome == "unsafe" || v.Outcome == "skipped") && v.Predicate == Blocked && none && (v.Timing == "timeout" || v.Timing == "not-completed")
	case originBaseCheck:
		return v.Timing == "under-timeout" && (v.Outcome == "pass" && v.Predicate == Pass && v.RungReached == bound.Entry.Rung || v.Outcome == "fail" && v.Predicate == Fail && none || v.Outcome == "skipped" && v.Predicate == Degraded && none) ||
			(v.Timing == "under-timeout" || v.Timing == "timeout") && v.Outcome == "unsafe" && v.Predicate == Blocked && none
	case originExecutor:
		completedTiming := v.Timing == "under-timeout" || v.Timing == "extended"
		switch v.Outcome {
		case "pass":
			return v.Predicate == Pass && v.RungReached == "E2" && completedTiming
		case "fail":
			return v.Predicate == Fail && none && completedTiming
		case "unsafe":
			if v.Predicate != Blocked || !none || v.Timing != "timeout" && v.Timing != "not-completed" {
				return false
			}
			return v.Timing != "timeout" || v.FailingDetail == "executor-timeout"
		}
	}
	return false
}

func validOriginClass(origin verdictOrigin, verdict CheckVerdict) bool {
	detail := verdict.FailingDetail
	switch origin {
	case originConductorConfig:
		return detail == "lane-ungoverned" || detail == "schema-ref-unknown"
	case originConductorPolicy:
		return detail == "check-params-refused" || detail == "side-effecting-unapproved" || detail == "side-effecting-execution-refused"
	case originConductorMachinery:
		return strings.HasPrefix(detail, "check-machinery-")
	case originBaseCheck:
		switch verdict.Outcome {
		case "pass":
			return detail == ""
		case "fail":
			return detail == "read-file-mismatch" || detail == "read-file-absent" || detail == "git-status-mismatch" || detail == "find-references-count-nonzero"
		case "skipped":
			return detail == "scan-domain-incomplete-undecodable"
		case "unsafe":
			_, ok := baseRefusalTiming(detail)
			return ok
		}
	case originExecutor:
		switch verdict.Outcome {
		case "pass":
			return detail == "" || detail == "output-truncated"
		case "fail":
			return detail == "suite-exit-mismatch"
		case "unsafe":
			return strings.HasPrefix(detail, "executor-")
		}
	}
	return false
}

func baseRefusalTiming(detail string) (string, bool) {
	switch detail {
	case "not-regular-file", "read-size-exceeded":
		return "under-timeout", true
	case "read-deadline-exceeded":
		return "timeout", true
	default:
		return "", false
	}
}

func verdictBindingFault(bound boundVerdict, detail string) validatedVerdict {
	checkID := bound.Selected.CheckID
	if checkID == "" {
		checkID = bound.Entry.ID
	}
	if checkID == "" {
		checkID = bound.Verdict.CheckID
	}
	claimRef := bound.Selected.ClaimRef
	if claimRef == "" {
		claimRef = bound.Verdict.ClaimRef
	}
	return validatedVerdict{
		Verdict:        CheckVerdict{CheckID: checkID, ClaimRef: claimRef, Outcome: "unsafe", RungReached: "none", Predicate: Blocked, Timing: "not-completed", FailingDetail: detail},
		MachineryFault: true, SignalClass: "none", Integrity: "self_reported",
	}
}

func signalClass(selectedEntryID, outcome string) string {
	if selectedEntryID == "red→green-differential" && outcome == "pass" {
		return "differential"
	}
	return "none"
}
