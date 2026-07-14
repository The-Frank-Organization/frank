package fixtures_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/executor"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
)

func TestS8AdversarialSlotInImmutableAgainstRetagEscape(t *testing.T) {
	cand := record.Record{Headers: map[string]string{
		"slot_in": "original-archetype", "authority_class": "no", "EVIDENCE_TARGET": "E1",
	}}
	result, terminal := observe.Gate(cand, "s8.implementer", "IMPL", "implementation", observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{
				ID: "retag-to-escape", Predicate: observe.Pass,
				ObservedFields: map[string]string{"slot_in": "escaped-archetype"},
			}
		},
	})
	if terminal != record.Rejected || result.FailingPredicate != "observe-write-allowlist" || result.FailureClass != "observe-machinery-fault" {
		t.Fatalf("retag escape terminal = %q, result = %#v", terminal, result)
	}
	if cand.Headers["slot_in"] != "original-archetype" {
		t.Fatalf("slot_in mutated through observe result: %#v", cand.Headers)
	}
}

func TestS8AdversarialExecutorIsolationAndVerdictIPHIntegrated(t *testing.T) {
	sourceRootSentinel := t.TempDir()
	tempRootSentinel := t.TempDir()
	commandSentinel := "descriptor-command-sentinel.sh"
	targetSentinel := "descriptor-target-sentinel"
	argumentSentinel := "descriptor-argument-sentinel"
	timeoutClassSentinel := "suite_bounded"
	timeoutSentinel := 1731 * time.Millisecond
	s8WriteExecutable(t, sourceRootSentinel, commandSentinel, `#!/bin/sh
set -eu
[ "$HOME" = "$PWD" ]
[ "$TMPDIR" = "$PWD/.tmp" ]
for key in FRANK_ROOT FRANK_SOCKET FRANK_CREDENTIAL SIGNING_KEY CONFIG_PATH OUTBOX_PATH; do
  eval "value=\${$key-}"
  [ -z "$value" ]
done
`)
	host := executor.New(executor.Config{
		TempRoot: tempRootSentinel,
		Suites: map[string]executor.Suite{
			targetSentinel: {
				SourceDir: sourceRootSentinel, Command: commandSentinel, Args: []string{argumentSentinel},
				TimeoutClass: timeoutClassSentinel, Timeout: timeoutSentinel,
			},
		},
	})
	registry := observe.NewRegistry(observe.RegistryEnv{
		NamedSuites: map[string]bool{targetSentinel: true}, Executor: host,
	})
	claims := s8Claims(t, map[string]string{
		"claim_ref": "isolation-probe", "check_id": "run-suite",
		"params": s8ClaimParams(t, map[string]string{"target": targetSentinel, "expect_green": "true"}),
	})
	cand := record.Record{Headers: map[string]string{"authority_class": "no", "EVIDENCE_TARGET": "E2"}}
	result, terminal := observe.Gate(cand, "s8.implementer", "IMPL", "implementation", observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(candidate observe.Candidate) observe.PredicateResult {
			return registry.EvaluateClaims(claims, candidate)
		},
	})
	if terminal != record.Accepted || result.ObservedFields["achieved_evidence"] != "E2" {
		t.Fatalf("integrated isolation terminal = %q, result = %#v", terminal, result)
	}
	var rows []map[string]string
	if err := json.Unmarshal([]byte(result.ObservedFields["executable_claim_results"]), &rows); err != nil {
		t.Fatalf("decode conductor-produced result row: %v", err)
	}
	wantRow := map[string]string{
		"check_id": "run-suite", "claim_ref": "isolation-probe", "outcome": "pass",
		"rung_reached": "E2", "signal_class": "none", "integrity": "observed",
	}
	if len(rows) != 1 || len(rows[0]) != len(wantRow) {
		t.Fatalf("conductor-produced result rows = %#v, want one exact thickened row", rows)
	}
	for key, want := range wantRow {
		if got := rows[0][key]; got != want {
			t.Fatalf("conductor-produced result row %s = %q, want %q; row=%#v", key, got, want, rows[0])
		}
	}
	serialized, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal integrated verdict: %v", err)
	}
	for _, forbidden := range []string{
		sourceRootSentinel, tempRootSentinel, commandSentinel, targetSentinel, argumentSentinel,
		timeoutClassSentinel, timeoutSentinel.String(),
		"FRANK_ROOT", "FRANK_SOCKET", "FRANK_CREDENTIAL", "SIGNING_KEY", "CONFIG_PATH", "OUTBOX_PATH",
	} {
		if strings.Contains(string(serialized), forbidden) {
			t.Fatalf("seat-visible verdict leaked %q: %s", forbidden, serialized)
		}
	}
	if strings.Contains(strings.ToLower(executor.AmbientResidual), "by construction") || !strings.Contains(executor.AmbientResidual, "same-uid") || !strings.Contains(executor.AmbientResidual, "without an OS sandbox") {
		t.Fatalf("same-uid residual was not bounded honestly: %q", executor.AmbientResidual)
	}
}

func TestS8AdversarialTruncationMarkerDistinctFromFailureClasses(t *testing.T) {
	source := t.TempDir()
	s8WriteExecutable(t, source, "pass-loud.sh", "#!/bin/sh\nprintf 'abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz'\n")
	s8WriteExecutable(t, source, "fail-loud.sh", "#!/bin/sh\nprintf 'abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz'\nexit 2\n")
	s8WriteExecutable(t, source, "timeout-loud.sh", "#!/bin/sh\nprintf 'abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz'\nsleep 30\n")
	host := executor.New(executor.Config{
		TempRoot: t.TempDir(), OutputLimit: 16,
		Suites: map[string]executor.Suite{
			"pass-loud":    {SourceDir: source, Command: "pass-loud.sh", TimeoutClass: "suite_bounded", Timeout: 2 * time.Second},
			"fail-loud":    {SourceDir: source, Command: "fail-loud.sh", TimeoutClass: "suite_bounded", Timeout: 2 * time.Second},
			"timeout-loud": {SourceDir: source, Command: "timeout-loud.sh", TimeoutClass: "suite_bounded", Timeout: 40 * time.Millisecond},
		},
	})
	entry := observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}
	for _, tc := range []struct {
		name, target, outcome, predicate, detail string
	}{
		{name: "truncated pass", target: "pass-loud", outcome: "pass", predicate: observe.Pass, detail: "output-truncated"},
		{name: "truncated failure", target: "fail-loud", outcome: "fail", predicate: observe.Fail, detail: "suite-exit-mismatch"},
		{name: "truncated timeout", target: "timeout-loud", outcome: "unsafe", predicate: observe.Blocked, detail: "executor-timeout"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			verdict := host.Spawn(entry, observe.Selection{
				CheckID: "run-suite", ClaimRef: tc.target,
				Params: map[string]string{"target": tc.target, "expect_green": "true"},
			})
			s8AssertAdversarialVerdict(t, verdict, tc.outcome, tc.predicate, tc.detail)
		})
	}
}

func s8AssertAdversarialVerdict(t *testing.T, verdict observe.CheckVerdict, outcome, predicate, detail string) {
	t.Helper()
	if verdict.Outcome != outcome || verdict.Predicate != predicate || verdict.FailingDetail != detail {
		t.Fatalf("verdict = %#v, want outcome=%q predicate=%q detail=%q", verdict, outcome, predicate, detail)
	}
}
