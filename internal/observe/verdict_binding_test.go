package observe

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/record"
)

type bindingExecutor struct{ verdict CheckVerdict }

func (e bindingExecutor) Spawn(CheckEntry, Selection) CheckVerdict { return e.verdict }

func TestExecutorCannotForgeDifferentialIdentity(t *testing.T) {
	reg := NewRegistry(RegistryEnv{
		NamedSuites: map[string]bool{"suite": true},
		Executor: bindingExecutor{verdict: CheckVerdict{
			CheckID: "red→green-differential", ClaimRef: "forged", Outcome: "pass", RungReached: "E2", Predicate: Pass, Timing: "under-timeout",
		}},
	})
	selection := Selection{CheckID: "run-suite", ClaimRef: "selected", Params: map[string]string{"target": "suite", "expect_green": "true"}}
	result := reg.Evaluator(selection)(Candidate{})
	if !result.MachineryFault || len(result.Verdicts) != 1 || result.Verdicts[0].FailingDetail != "check-machinery-verdict-identity-mismatch" {
		t.Fatalf("result = %#v", result)
	}
	stamps := baseStamps(record.Record{}, result)
	if strings.Contains(stamps["executable_claim_results"], "differential") {
		t.Fatalf("forged differential reached rows: %s", stamps["executable_claim_results"])
	}
}

func TestVerdictTupleMatrixRejectsContradictions(t *testing.T) {
	entry := CheckEntry{ID: "run-suite", Rung: "E2"}
	tests := []CheckVerdict{
		{Outcome: "pass", Predicate: Fail, RungReached: "E2", Timing: "under-timeout"},
		{Outcome: "pass", Predicate: Blocked, RungReached: "E2", Timing: "under-timeout"},
		{Outcome: "pass", Predicate: Degraded, RungReached: "E2", Timing: "under-timeout"},
		{Outcome: "fail", Predicate: Pass, RungReached: "none", Timing: "under-timeout", FailingDetail: "suite-exit-mismatch"},
		{Outcome: "skipped", Predicate: Pass, RungReached: "none", Timing: "under-timeout"},
		{Outcome: "unsafe", Predicate: Pass, RungReached: "none", Timing: "not-completed", FailingDetail: "executor-timeout"},
		{Outcome: "fail", Predicate: Fail, RungReached: "E1", Timing: "under-timeout", FailingDetail: "suite-exit-mismatch"},
		{Outcome: "pass", Predicate: Pass, RungReached: "E3", Timing: "under-timeout"},
	}
	for i, verdict := range tests {
		verdict.CheckID, verdict.ClaimRef = "run-suite", "claim"
		validated := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originExecutor, Entry: entry})
		if validated.Verdict.FailingDetail != "check-machinery-verdict-tuple-invalid" || !validated.MachineryFault {
			t.Fatalf("case %d = %#v", i, validated)
		}
	}
}

func TestTruncatedPassSurvivesBindingPass(t *testing.T) {
	verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: "suite", Outcome: "pass", RungReached: "E2", Predicate: Pass, Timing: "under-timeout", FailingDetail: "output-truncated"}
	validated := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originExecutor, Entry: CheckEntry{ID: "run-suite", Rung: "E2"}})
	if validated.MachineryFault || validated.Verdict != verdict {
		t.Fatalf("validated = %#v, want byte-intact verdict", validated)
	}
}

func TestExecutorOriginCannotPresentPolicyClass(t *testing.T) {
	verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: "suite", Outcome: "unsafe", RungReached: "none", Predicate: Blocked, Timing: "not-completed", FailingDetail: "side-effecting-unapproved"}
	validated := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originExecutor, Entry: CheckEntry{ID: "run-suite", Rung: "E2"}})
	assertOriginMismatchFault(t, validated)
}

func TestConductorPolicyOriginCannotPresentMachineryToken(t *testing.T) {
	verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: "suite", Outcome: "unsafe", RungReached: "none", Predicate: Blocked, FailingDetail: "check-machinery-hostile"}
	validated := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originConductorPolicy, Entry: CheckEntry{ID: "run-suite", Rung: "E2"}})
	assertOriginMismatchFault(t, validated)
}

func TestValidPolicyRefusalStaysMachineryFaultFalse(t *testing.T) {
	verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: "suite", Outcome: "unsafe", RungReached: "none", Predicate: Blocked, FailingDetail: "check-params-refused"}
	validated := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originConductorPolicy, Entry: CheckEntry{ID: "run-suite", Rung: "E2"}})
	if validated.MachineryFault || validated.Verdict.FailingDetail != "check-params-refused" || validated.Verdict.Timing != "not-completed" {
		t.Fatalf("validated = %#v", validated)
	}
}

func TestBaseRefusalRowPreservesNoVantageDisposition(t *testing.T) {
	for _, tc := range []struct {
		detail string
		timing string
	}{
		{detail: "not-regular-file", timing: "under-timeout"},
		{detail: "read-size-exceeded", timing: "under-timeout"},
		{detail: "read-deadline-exceeded", timing: "timeout"},
	} {
		t.Run(tc.detail, func(t *testing.T) {
			selection := Selection{CheckID: "read-file", ClaimRef: tc.detail}
			verdict := CheckVerdict{
				CheckID: "read-file", ClaimRef: tc.detail, Outcome: "unsafe",
				RungReached: "none", Predicate: Blocked, FailingDetail: tc.detail,
			}
			got := validateBoundVerdict(boundVerdict{
				Verdict: verdict, Origin: originBaseCheck,
				Entry: CheckEntry{ID: "read-file", Rung: "E1"}, Selected: selection,
			})
			if got.MachineryFault || got.Verdict.Outcome != "unsafe" || got.Verdict.Predicate != Blocked || got.Verdict.FailingDetail != tc.detail || got.Verdict.Timing != tc.timing || got.Integrity != "self_reported" {
				t.Fatalf("validated = %#v", got)
			}
		})
	}

	t.Run("allowlist is closed", func(t *testing.T) {
		verdict := CheckVerdict{
			CheckID: "read-file", ClaimRef: "unruled", Outcome: "unsafe",
			RungReached: "none", Predicate: Blocked, FailingDetail: "read-file-absent",
		}
		got := validateBoundVerdict(boundVerdict{
			Verdict: verdict, Origin: originBaseCheck,
			Entry:    CheckEntry{ID: "read-file", Rung: "E1"},
			Selected: Selection{CheckID: "read-file", ClaimRef: "unruled"},
		})
		assertOriginMismatchFault(t, got)
	})
}

func TestReadFileMachineryStaysOutsideBaseRefusal(t *testing.T) {
	for _, tc := range []struct {
		detail string
		timing string
	}{
		{detail: "check-machinery-read-file-timeout", timing: "timeout"},
		{detail: "check-machinery-read-file-breaker-open", timing: "not-completed"},
	} {
		t.Run(tc.detail, func(t *testing.T) {
			selection := Selection{CheckID: "read-file", ClaimRef: tc.detail}
			verdict := CheckVerdict{
				CheckID: "read-file", ClaimRef: tc.detail, Outcome: "unsafe",
				RungReached: "none", Predicate: Blocked, FailingDetail: tc.detail,
			}
			got := validateBoundVerdict(boundVerdict{
				Verdict: verdict, Origin: originConductorMachinery,
				Entry: CheckEntry{ID: "read-file", Rung: "E1"}, Selected: selection,
			})
			if !got.MachineryFault || got.Verdict.FailingDetail != tc.detail || got.Verdict.Timing != tc.timing {
				t.Fatalf("validated = %#v", got)
			}
		})
	}
}

func TestTimingBranchesAndInconsistency(t *testing.T) {
	entry := CheckEntry{ID: "run-suite", Rung: "E2"}
	for _, timing := range []string{"under-timeout", "extended"} {
		verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: timing, Outcome: "pass", RungReached: "E2", Predicate: Pass, Timing: timing}
		if got := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originExecutor, Entry: entry}); got.MachineryFault {
			t.Fatalf("%s = %#v", timing, got)
		}
	}
	for _, timing := range []string{"timeout", "not-completed"} {
		verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: timing, Outcome: "unsafe", RungReached: "none", Predicate: Blocked, Timing: timing, FailingDetail: "executor-timeout"}
		if got := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originExecutor, Entry: entry}); got.MachineryFault != true || got.Verdict.FailingDetail != "executor-timeout" {
			t.Fatalf("%s = %#v", timing, got)
		}
	}
	for _, timing := range []string{"", "host-secret", "extended"} {
		verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: "bad", Outcome: "unsafe", RungReached: "none", Predicate: Blocked, Timing: timing, FailingDetail: "executor-timeout"}
		got := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originExecutor, Entry: entry})
		if !got.MachineryFault || got.Verdict.FailingDetail != "check-machinery-verdict-tuple-invalid" {
			t.Fatalf("inconsistent %s = %#v", timing, got)
		}
	}
}

func TestVerdictOutputRedaction(t *testing.T) {
	secret := "/governed/root/secret-token"
	verdict := CheckVerdict{CheckID: "run-suite", ClaimRef: "suite", Outcome: "fail", RungReached: "none", Predicate: Fail, Timing: secret, FailingDetail: secret}
	validated := validateBoundVerdict(boundVerdict{Verdict: verdict, Origin: originExecutor, Entry: CheckEntry{ID: "run-suite", Rung: "E2"}})
	raw, err := json.Marshal(validated)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), secret) {
		t.Fatalf("validated output leaks planted residue: %s", raw)
	}
}

func TestSignalClassDerivedFromSelection(t *testing.T) {
	if got := signalClass("red→green-differential", "pass"); got != "differential" {
		t.Fatalf("signal = %q", got)
	}
	for _, id := range []string{"read-file", "git-status", "find-references", "run-suite"} {
		if got := signalClass(id, "pass"); got != "none" {
			t.Fatalf("%s signal = %q", id, got)
		}
	}
}

func TestClaimRowsCarryRungAndSignalClass(t *testing.T) {
	result := PredicateResult{ID: "claim", Predicate: Pass, Verdicts: []CheckVerdict{{CheckID: "run-suite", ClaimRef: "claim", Outcome: "pass", RungReached: "E2", Predicate: Pass}}, claimBindings: map[string]claimBinding{
		"claim": {RungReached: "E2", SignalClass: "none", Integrity: "observed"},
	}}
	stamps := baseStamps(record.Record{}, result)
	var rows []map[string]string
	if err := json.Unmarshal([]byte(stamps["executable_claim_results"]), &rows); err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0]["rung_reached"] != "E2" || rows[0]["signal_class"] != "none" || rows[0]["integrity"] != "observed" {
		t.Fatalf("rows = %#v", rows)
	}
}

func assertOriginMismatchFault(t *testing.T, got validatedVerdict) {
	t.Helper()
	if !got.MachineryFault || got.Verdict.FailingDetail != "check-machinery-verdict-origin-class-mismatch" {
		t.Fatalf("validated = %#v", got)
	}
}
