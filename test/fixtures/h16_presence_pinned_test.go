package fixtures_test

import (
	"path/filepath"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/record"
)

func TestH16PresencePinnedRejectsPresentEmptySystemHeader(t *testing.T) {
	reg := loadH16Registry(t)
	seat := fieldspec.SeatMeta{Name: "s12.implementer", Role: "implementer"}
	rec := h16PresenceCandidate()
	rec.Headers["delivery_state"] = ""
	violations := reg.Validate(rec, seat, h16PresenceDigest(reg, seat, rec), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	assertH16Violation(t, violations, "delivery_state", "system-owned", "delivery_state is system-owned")
}

func TestH16PresencePinnedSharedImpactPopulation(t *testing.T) {
	reg := loadH16Registry(t)
	assertH16SystemHeaderPopulation(t, reg, 34)
}

func TestH16PresencePinnedEnvelopeRowsRemainOutsideHeaderPopulation(t *testing.T) {
	reg := loadH16Registry(t)
	seat := fieldspec.SeatMeta{Name: "s12.implementer", Role: "implementer"}
	for _, id := range []string{"FROM", "ROLE", "relay_id"} {
		t.Run(id, func(t *testing.T) {
			spec, ok := reg.ByID(id)
			if !ok {
				t.Fatalf("registry missing %s", id)
			}
			if spec.Layer == "header" {
				t.Fatalf("%s layer=%q, want non-header", id, spec.Layer)
			}
			rec := h16PresenceCandidate()
			rec.Headers[id] = ""
			violations := reg.Validate(rec, seat, h16PresenceDigest(reg, seat, rec), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
			assertH16NoViolation(t, violations, id, "system-owned")
		})
	}
}

func TestH16PresencePinnedObservationProducerRetainsEmptyDegradationNotes(t *testing.T) {
	candidate := record.Record{Headers: map[string]string{"authority_class": "no"}}
	completed, violation := engine.CompleteObserved(candidate, map[string]string{
		"achieved_evidence":        "E1",
		"evidence_integrity":       `{}`,
		"executable_claim_results": `[]`,
		"egress_scan_result":       "not_applicable",
		"degradation_notes":        "",
		"attestation_source":       "conductor",
		"target_gap_result":        "met",
		"record_integrity":         "observed",
		"deviated_observed":        "no",
		"bucket_binding_observed":  "no",
	})
	if violation != nil {
		t.Fatalf("CompleteObserved violation=%+v", violation)
	}
	if value, present := completed.Headers["degradation_notes"]; !present || value != "" {
		t.Fatalf("degradation_notes present=%v value=%q, want retained present-empty producer output", present, value)
	}
}

func assertH16SystemHeaderPopulation(t *testing.T, reg *fieldspec.Registry, wantCount int) {
	t.Helper()
	population := make([]fieldspec.FieldSpec, 0, wantCount)
	for _, spec := range reg.Fields {
		if spec.Layer == "header" &&
			(spec.Owner == "system" || spec.Owner == "computed" || spec.FillConstraints == "system_only" || spec.FillConstraints == "computed_result") {
			population = append(population, spec)
		}
	}
	if len(population) != wantCount {
		t.Fatalf("system/computed header population=%d, want %d", len(population), wantCount)
	}

	seat := fieldspec.SeatMeta{Name: "s12.implementer", Role: "implementer"}
	for _, spec := range population {
		t.Run(spec.ID, func(t *testing.T) {
			rec := h16PresenceCandidate()
			rec.Headers[spec.ID] = ""
			violations := reg.Validate(rec, seat, h16PresenceDigest(reg, seat, rec), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
			assertH16Violation(t, violations, spec.ID, "system-owned", spec.ID+" is system-owned")
		})
	}
}

func loadH16Registry(t *testing.T) *fieldspec.Registry {
	t.Helper()
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("load registry: %v", err)
	}
	return reg
}

func h16PresenceCandidate() record.Record {
	return record.Record{
		Envelope: record.Envelope{SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":               "SITREP",
			"AUTHORITY":           "report-only",
			"CEREMONY_TIER":       "medium",
			"EVIDENCE_TARGET":     "E1",
			"HUMAN_GATE_REQUIRED": "no",
			"SUBJECT":             "h16 presence-pinned candidate",
		},
	}
}

func h16PresenceDigest(reg *fieldspec.Registry, seat fieldspec.SeatMeta, rec record.Record) string {
	_, digest := reg.Render(fieldspec.RenderEnv{}, seat, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	return digest
}

func assertH16Violation(t *testing.T, violations []fieldspec.Violation, field, class, reason string) {
	t.Helper()
	for _, violation := range violations {
		if violation.Field == field && violation.Class == class {
			if violation.Reason != reason {
				t.Fatalf("%s/%s reason=%q, want %q", field, class, violation.Reason, reason)
			}
			return
		}
	}
	t.Fatalf("missing violation %s/%s in %+v", field, class, violations)
}

func assertH16NoViolation(t *testing.T, violations []fieldspec.Violation, field, class string) {
	t.Helper()
	for _, violation := range violations {
		if violation.Field == field && violation.Class == class {
			t.Fatalf("unexpected violation %s/%s in %+v", field, class, violations)
		}
	}
}
