package fieldspec_test

import (
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

func TestValidateV2SeatScopeAndSystemPayloadIgnore(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}
	rec := validCandidate()
	rec.Headers["AUTHORITY"] = "merge-gated"
	rec.Headers["grant"] = "dispatch-impl"
	rec.Headers["delivery_state"] = "forged"

	violations := reg.Validate(rec, seat, digestFor(t, reg, seat, rec), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	assertViolation(t, violations, "AUTHORITY", "seat-scope")
	assertViolation(t, violations, "grant", "seat-scope")
	assertNoViolation(t, violations, "delivery_state", "enum")
}

func TestValidateV2FullEnumSweep(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	for _, field := range []string{
		"PHASE",
		"AUTHORITY",
		"CEREMONY_TIER",
		"EVIDENCE_TARGET",
		"HUMAN_GATE_REQUIRED",
		"gate_category",
		"DESIGN_RECORD_KIND",
		"DESIGN_REVIEW_VERDICT",
		"grant",
		"record_kind",
	} {
		t.Run(field, func(t *testing.T) {
			rec := validCandidate()
			rec.Headers[field] = "not-a-token"
			violations := reg.Validate(rec, seat, digestFor(t, reg, seat, validCandidate()), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
			assertViolation(t, violations, field, "enum")
		})
	}
}

func TestValidateV2DigestRequiredAndTierBound(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	rec := validCandidate()

	assertViolation(t, reg.Validate(rec, seat, "", fieldspec.RenderEnv{}, fieldspec.ClosedGrantState), "form_digest", "re-render")
	assertViolation(t, reg.Validate(rec, seat, "stale", fieldspec.RenderEnv{}, fieldspec.ClosedGrantState), "form_digest", "re-render")

	_, mediumDigest := reg.Render(fieldspec.RenderEnv{}, seat, rec.Headers["PHASE"], "medium", fieldspec.ClosedGrantState)
	rec.Headers["CEREMONY_TIER"] = "large"
	assertViolation(t, reg.Validate(rec, seat, mediumDigest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState), "form_digest", "re-render")
}

func TestValidateV2TypedCarrierAndGrillDependentRequired(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}

	nonCanonical := validCandidate()
	nonCanonical.Headers["SCOPE_DIFF"] = `[{ "status" : "in", "path" : "README.md" }]`
	violations := reg.Validate(nonCanonical, seat, digestFor(t, reg, seat, nonCanonical), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	assertViolation(t, violations, "SCOPE_DIFF", "canonical-encoding")

	grillMissingLock := validCandidate()
	grillMissingLock.Headers["GRILL_REQUIRED"] = "yes"
	grillMissingLock.Headers["DESIGN_LOCK_ID"] = "design-lock-1"
	violations = reg.Validate(grillMissingLock, seat, digestFor(t, reg, seat, grillMissingLock), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	assertViolation(t, violations, "GRILL_LOCK_ID", "required")

	grillNoDesignLock := validCandidate()
	grillNoDesignLock.Headers["GRILL_REQUIRED"] = "yes"
	violations = reg.Validate(grillNoDesignLock, seat, digestFor(t, reg, seat, grillNoDesignLock), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	assertNoViolation(t, violations, "GRILL_LOCK_ID", "required")
}

func validCandidate() record.Record {
	return record.Record{
		Envelope: record.Envelope{SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "ok",
		},
	}
}

func digestFor(t *testing.T, reg *fieldspec.Registry, seat fieldspec.SeatMeta, rec record.Record) string {
	t.Helper()
	_, digest := reg.Render(fieldspec.RenderEnv{}, seat, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	return digest
}

func assertNoViolation(t *testing.T, violations []fieldspec.Violation, field, class string) {
	t.Helper()
	for _, violation := range violations {
		if violation.Field == field && violation.Class == class {
			t.Fatalf("unexpected violation %s/%s in %+v", field, class, violations)
		}
	}
}
