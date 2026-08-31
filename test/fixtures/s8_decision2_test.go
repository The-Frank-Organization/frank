package fixtures_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/observe"
	"github.com/The-Frank-Organization/frank/internal/record"
)

func TestS8Decision2NoVantageDisposition(t *testing.T) {
	for _, tc := range []struct {
		name      string
		authority string
		verdicts  []observe.CheckVerdict
		want      string
		integrity string
		escalate  bool
	}{
		{name: "authority self reported holds", authority: "yes", want: record.Held, integrity: "self_reported", escalate: true},
		{name: "non authority self reported accepts", authority: "no", want: record.Accepted, integrity: "self_reported"},
		{name: "authority mixed holds", authority: "yes", want: record.Held, integrity: "mixed", escalate: true, verdicts: []observe.CheckVerdict{
			{CheckID: "read-file", ClaimRef: "visible", Outcome: "pass", RungReached: "E1", Predicate: observe.Pass},
			{CheckID: "opaque", ClaimRef: "hidden", Outcome: "unsafe", RungReached: "none", Predicate: observe.Blocked},
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cand := record.Record{Headers: map[string]string{"authority_class": tc.authority, "EVIDENCE_TARGET": "E1"}}
			result, terminal := observe.Gate(cand, "seat-a", "SITREP", "report-only", observe.Env{
				PresentLayers: map[string]bool{"observe": true},
				Evaluate: func(observe.Candidate) observe.PredicateResult {
					return observe.PredicateResult{ID: "no-vantage", Predicate: observe.Blocked, Verdicts: tc.verdicts}
				},
			})
			if terminal != tc.want || result.Escalate != tc.escalate {
				t.Fatalf("terminal = %q, escalate = %v, result = %#v", terminal, result.Escalate, result)
			}
			if result.ObservedFields["record_integrity"] != tc.integrity {
				t.Fatalf("record_integrity = %q, want %q", result.ObservedFields["record_integrity"], tc.integrity)
			}
			if result.ObservedFields["degradation_notes"] != "observation-unavailable" {
				t.Fatalf("degradation_notes = %q", result.ObservedFields["degradation_notes"])
			}
		})
	}
}

func TestS8Decision2ObservedFailRejectsBothClasses(t *testing.T) {
	for _, authority := range []string{"no", "yes"} {
		cand := record.Record{Headers: map[string]string{"authority_class": authority}}
		result, terminal := observe.Gate(cand, "seat-a", "IMPL", "implementation", observe.Env{
			PresentLayers: map[string]bool{"observe": true},
			Evaluate: func(observe.Candidate) observe.PredicateResult {
				return observe.PredicateResult{ID: "suite-green", Predicate: observe.Fail}
			},
		})
		if terminal != record.Rejected || result.FailingPredicate != "suite-green" || result.Escalate {
			t.Fatalf("authority %s terminal = %q, result = %#v", authority, terminal, result)
		}
	}
}

func TestS8Decision2MachineryFaultNeverUsesNoVantageAcceptance(t *testing.T) {
	reg := observe.NewRegistry(observe.RegistryEnv{
		NamedSuites: map[string]bool{"slow": true}, Executor: s8TimeoutExecutor{},
	})
	evaluate := reg.Evaluator(observe.Selection{
		CheckID: "run-suite", ClaimRef: "suite-timeout",
		Params: map[string]string{"target": "slow", "expect_green": "true"},
	})
	for _, tc := range []struct {
		authority string
		want      string
		escalate  bool
	}{
		{authority: "yes", want: record.Held, escalate: true},
		{authority: "no", want: record.Rejected},
	} {
		cand := record.Record{Headers: map[string]string{"authority_class": tc.authority, "EVIDENCE_TARGET": "E2"}}
		result, terminal := observe.Gate(cand, "seat-a", "IMPL", "implementation", observe.Env{
			PresentLayers: map[string]bool{"observe": true},
			Evaluate:      evaluate,
		})
		if terminal != tc.want || result.Escalate != tc.escalate || result.FailureClass != "observe-machinery-fault" {
			t.Fatalf("authority %s terminal = %q, result = %#v", tc.authority, terminal, result)
		}
		if terminal == record.Accepted {
			t.Fatalf("machinery fault accepted with label: %#v", result)
		}
	}
}

func TestS8Decision2E1MachineryFaultsAndObservedAbsence(t *testing.T) {
	lane := t.TempDir()
	if err := os.Mkdir(filepath.Join(lane, "not-a-file"), 0o755); err != nil {
		t.Fatalf("mkdir read-file machinery fixture: %v", err)
	}
	reg := observe.NewRegistry(observe.RegistryEnv{
		Lanes:         map[string]string{"lane-a": lane},
		GitExecutable: s8TimeoutGit(t),
		ReadTimeout:   20 * time.Millisecond,
	})
	for _, tc := range []struct {
		name                 string
		selection            observe.Selection
		nonAuthorityTerminal string
		authorityClass       string
		nonAuthorityClass    string
		degradation          string
		refusalDetail        string
	}{
		{
			name:                 "git status timeout",
			nonAuthorityTerminal: record.Rejected,
			authorityClass:       "observe-machinery-fault",
			nonAuthorityClass:    "observe-machinery-fault",
			degradation:          "observe-machinery-fault",
			selection: observe.Selection{CheckID: "git-status", ClaimRef: "git-timeout", Params: map[string]string{
				"lane_ref": "lane-a", "expect": "clean",
			}},
		},
		{
			name:                 "read file IO error",
			nonAuthorityTerminal: record.Accepted,
			authorityClass:       "observation-unavailable",
			degradation:          "observation-unavailable",
			refusalDetail:        "not-regular-file",
			selection: observe.Selection{CheckID: "read-file", ClaimRef: "read-error", Params: map[string]string{
				"lane_ref": "lane-a", "path": "not-a-file", "expect": "line:x",
			}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.refusalDetail != "" {
				if verdict := reg.Run(tc.selection); verdict.FailingDetail != tc.refusalDetail {
					t.Fatalf("refusal detail = %q, want %q: %#v", verdict.FailingDetail, tc.refusalDetail, verdict)
				}
			}
			_, controlTerminal := observe.Gate(record.Record{Headers: map[string]string{"authority_class": "no"}}, "seat-a", "IMPL", "implementation", observe.Env{
				PresentLayers: map[string]bool{"observe": true},
				Evaluate: func(observe.Candidate) observe.PredicateResult {
					return observe.PredicateResult{ID: "no-vantage-control", Predicate: observe.Blocked}
				},
			})
			if controlTerminal != record.Accepted {
				t.Fatalf("no-vantage control terminal = %q", controlTerminal)
			}

			for _, authority := range []string{"yes", "no"} {
				cand := record.Record{Headers: map[string]string{"authority_class": authority, "EVIDENCE_TARGET": "E1"}}
				result, terminal := observe.Gate(cand, "seat-a", "IMPL", "implementation", observe.Env{
					PresentLayers: map[string]bool{"observe": true}, Evaluate: reg.Evaluator(tc.selection),
				})
				want := tc.nonAuthorityTerminal
				wantClass := tc.nonAuthorityClass
				wantEscalate := false
				if authority == "yes" {
					want = record.Held
					wantClass = tc.authorityClass
					wantEscalate = true
				}
				if terminal != want || result.Escalate != wantEscalate || result.FailureClass != wantClass || result.ObservedFields["degradation_notes"] != tc.degradation {
					t.Fatalf("authority %s terminal = %q, result = %#v", authority, terminal, result)
				}
			}
		})
	}

	missing := observe.Selection{CheckID: "read-file", ClaimRef: "missing", Params: map[string]string{
		"lane_ref": "lane-a", "path": "missing.txt", "expect": "line:x",
	}}
	result, terminal := observe.Gate(record.Record{Headers: map[string]string{"authority_class": "no"}}, "seat-a", "IMPL", "implementation", observe.Env{
		PresentLayers: map[string]bool{"observe": true}, Evaluate: reg.Evaluator(missing),
	})
	if terminal != record.Rejected || result.FailureClass != "observed-false" || result.FailingPredicate != "missing" {
		t.Fatalf("missing-file terminal = %q, result = %#v", terminal, result)
	}
}

type s8TimeoutExecutor struct{}

func (s8TimeoutExecutor) Spawn(observe.CheckEntry, observe.Selection) observe.CheckVerdict {
	return observe.CheckVerdict{
		CheckID: "run-suite", ClaimRef: "suite-timeout", Outcome: "unsafe",
		RungReached: "none", Predicate: observe.Blocked, Timing: "timeout", FailingDetail: "executor-timeout",
	}
}
