package fixtures_test

import (
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
)

func TestS8PipelineStep3ComputesAuthorityClassFromLockedInputs(t *testing.T) {
	reg := &fieldspec.Registry{GateCategory: map[string][]string{
		"A": {"authz_security"},
		"B": {"routing"},
	}}
	meta := seat.SeatMeta{Name: "s8.implementer", Role: "implementer"}
	for name, tc := range map[string]struct {
		rec  record.Record
		want string
	}{
		"plain sitrep": {rec: record.Record{Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only"}}, want: "no"},
		"impl phase":   {rec: record.Record{Headers: map[string]string{"PHASE": "IMPL", "AUTHORITY": "implementation"}}, want: "yes"},
		"a gate":       {rec: record.Record{Headers: map[string]string{"PHASE": "SITREP", "gate_category": "authz_security"}}, want: "yes"},
		"b gate":       {rec: record.Record{Headers: map[string]string{"PHASE": "SITREP", "gate_category": "routing"}}, want: "no"},
	} {
		t.Run(name, func(t *testing.T) {
			if got := engine.AuthorityClass(reg, tc.rec, meta); got != tc.want {
				t.Fatalf("authority_class = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestS8PipelineStep45DerivesSurfaceIntentByProfile(t *testing.T) {
	for name, tc := range map[string]struct {
		headers     map[string]string
		wantSurface bool
	}{
		"non-gate defaults progress": {headers: map[string]string{"authority_class": "no"}, wantSurface: true},
		"category gate has none":     {headers: map[string]string{"authority_class": "no", "gate_category": "routing"}},
		"human gate has none":        {headers: map[string]string{"authority_class": "yes", "HUMAN_GATE_REQUIRED": "yes"}},
	} {
		t.Run(name, func(t *testing.T) {
			rec, violation := engine.CompleteObserved(record.Record{Headers: tc.headers}, s8ObserveManifest())
			if violation != nil {
				t.Fatalf("completeness violation: %#v", violation)
			}
			if tc.wantSurface && rec.Headers["surface_intent"] != "progress" {
				t.Fatalf("surface_intent = %q, want progress", rec.Headers["surface_intent"])
			}
			if !tc.wantSurface {
				if _, present := rec.Headers["surface_intent"]; present {
					t.Fatalf("gate-bearing record has surface_intent: %#v", rec.Headers)
				}
			}
		})
	}
}

func TestS8PipelineStep45RejectsProducerOutsideManifest(t *testing.T) {
	for _, field := range []string{"authority_class", "surface_intent", "FROM"} {
		writes := s8ObserveManifest()
		writes[field] = "forged"
		_, violation := engine.CompleteObserved(record.Record{Headers: map[string]string{"authority_class": "no"}}, writes)
		if violation == nil || violation.Field != field || violation.Class != "producer-manifest" {
			t.Fatalf("field %s violation = %#v", field, violation)
		}
	}
}

func TestS8PipelineStep45RejectsIncompleteObserveManifest(t *testing.T) {
	writes := s8ObserveManifest()
	delete(writes, "record_integrity")
	_, violation := engine.CompleteObserved(record.Record{Headers: map[string]string{"authority_class": "no"}}, writes)
	if violation == nil || violation.Field != "record_integrity" || violation.Class != "producer-incomplete" {
		t.Fatalf("violation = %#v", violation)
	}
}

func s8ObserveManifest() map[string]string {
	return map[string]string{
		"achieved_evidence": "E1", "evidence_integrity": `{}`,
		"executable_claim_results": `[]`, "egress_scan_result": "not_applicable",
		"degradation_notes": "", "attestation_source": "conductor",
		"target_gap_result": "met", "record_integrity": "observed",
		"deviated_observed": "no", "bucket_binding_observed": "no",
	}
}
