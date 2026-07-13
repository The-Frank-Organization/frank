package fixtures_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
)

func TestS8VerdictBindingComputesE2RowsAndMetTarget(t *testing.T) {
	cand := record.Record{Headers: map[string]string{"EVIDENCE_TARGET": "E2"}}
	result, terminal := observe.Gate(cand, "seat-a", "IMPL", "implementation", observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{
				ID: "suite-green", Predicate: observe.Pass,
				Verdicts: []observe.CheckVerdict{{
					CheckID: "run-suite", ClaimRef: "suite-green", Outcome: "pass",
					RungReached: "E2", Predicate: observe.Pass, Timing: "under-timeout",
				}},
			}
		},
	})
	if terminal != record.Accepted {
		t.Fatalf("terminal = %q, result = %#v", terminal, result)
	}
	if result.ObservedFields["achieved_evidence"] != "E2" || result.ObservedFields["target_gap_result"] != "met" {
		t.Fatalf("computed evidence fields = %#v", result.ObservedFields)
	}
	var rows []map[string]string
	if err := json.Unmarshal([]byte(result.ObservedFields["executable_claim_results"]), &rows); err != nil {
		t.Fatalf("decode claim rows: %v", err)
	}
	if len(rows) != 1 || rows[0]["check_id"] != "run-suite" || rows[0]["claim_ref"] != "suite-green" || rows[0]["outcome"] != "pass" {
		t.Fatalf("claim rows = %#v", rows)
	}
	var integrity map[string]string
	if err := json.Unmarshal([]byte(result.ObservedFields["evidence_integrity"]), &integrity); err != nil {
		t.Fatalf("decode integrity map: %v", err)
	}
	if len(integrity) != 1 || integrity["suite-green"] != "observed" {
		t.Fatalf("integrity map = %#v", integrity)
	}
}

func TestS8VerdictBindingComputesTargetGapFromPassingRung(t *testing.T) {
	cand := record.Record{Headers: map[string]string{"EVIDENCE_TARGET": "E2"}}
	result, terminal := observe.Gate(cand, "seat-a", "SITREP", "report-only", observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{
				ID: "read", Predicate: observe.Pass,
				Verdicts: []observe.CheckVerdict{{
					CheckID: "read-file", ClaimRef: "artifact", Outcome: "pass",
					RungReached: "E1", Predicate: observe.Pass, Timing: "under-timeout",
				}},
			}
		},
	})
	if terminal != record.Accepted || result.ObservedFields["achieved_evidence"] != "E1" || result.ObservedFields["target_gap_result"] != "target_gt_achieved" {
		t.Fatalf("target-gap binding = terminal %q, fields %#v", terminal, result.ObservedFields)
	}
}

func TestS8ForgedExecutorFieldWriteIsInert(t *testing.T) {
	cand := record.Record{Headers: map[string]string{"EVIDENCE_TARGET": "E2"}}
	result, terminal := observe.Gate(cand, "seat-a", "IMPL", "implementation", observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{
				ID: "suite-green", Predicate: observe.Pass,
				ObservedFields: map[string]string{"achieved_evidence": "E4"},
				Verdicts: []observe.CheckVerdict{{
					CheckID: "run-suite", ClaimRef: "suite-green", Outcome: "pass",
					RungReached: "E2", Predicate: observe.Pass, Timing: "under-timeout",
					FailingDetail: "/canonical/store/secret",
				}},
			}
		},
	})
	if terminal != record.Accepted {
		t.Fatalf("terminal = %q, result = %#v", terminal, result)
	}
	if result.ObservedFields["achieved_evidence"] != "E2" {
		t.Fatalf("forged field write took effect: %#v", result.ObservedFields)
	}
	if bytes, _ := json.Marshal(result); string(bytes) == "" || strings.Contains(string(bytes), "/canonical") || strings.Contains(string(bytes), "secret") {
		t.Fatalf("unredacted executor detail crossed boundary: %s", bytes)
	}
}
