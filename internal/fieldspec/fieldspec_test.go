package fieldspec_test

import (
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

func TestRenderA2EnumAndGrantVisibility(t *testing.T) {
	reg := loadRegistry(t)

	env := fieldspec.RenderEnv{}
	pairForm, _ := reg.Render(env, fieldspec.SeatMeta{Name: "s1-core", Role: "implementer"}, "MERGE-GATE", "medium", fieldspec.ClosedGrantState)
	if pairForm.HasField("grant") {
		t.Fatalf("pair form unexpectedly renders grant")
	}
	if pairForm.OptionAllowed("AUTHORITY", "merge-gated") {
		t.Fatalf("pair form unexpectedly allows merge-gated")
	}

	operatorForm, _ := reg.Render(env, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "MERGE-GATE", "medium", fieldspec.ClosedGrantState)
	if !operatorForm.HasField("grant") {
		t.Fatalf("operator merge-gate form did not render grant")
	}
	if !operatorForm.OptionAllowed("grant", "dispatch-merge") {
		t.Fatalf("operator merge-gate form did not allow dispatch-merge")
	}

	orchestratorForm, _ := reg.Render(env, fieldspec.SeatMeta{Name: "s1.orchestrator-planner", Role: "orchestrator-planner"}, "IMPL", "medium", fieldspec.ClosedGrantState)
	if !orchestratorForm.HasField("grant") || !orchestratorForm.OptionAllowed("grant", "dispatch-impl") {
		t.Fatalf("orchestrator form missing dispatch-impl grant: %#v", orchestratorForm.Fields["grant"])
	}
}

func TestValidateB3AndStaleDigest(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s1-core", Role: "implementer"}
	_, digest := reg.Render(fieldspec.RenderEnv{}, seat, "SITREP", "medium", fieldspec.ClosedGrantState)

	rec := record.Record{
		Envelope: record.Envelope{From: "s1-core.implementer", Role: "implementer", SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":     "NOPE",
			"AUTHORITY": "merge-gated",
		},
	}
	violations := reg.Validate(rec, seat, digest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	assertViolation(t, violations, "SUBJECT", "required")
	assertViolation(t, violations, "PHASE", "enum")
	assertViolation(t, violations, "AUTHORITY", "seat-scope")

	valid := record.Record{
		Envelope: record.Envelope{From: "s1-core.implementer", Role: "implementer", SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "ok"},
	}
	assertViolation(t, reg.Validate(valid, seat, "stale", fieldspec.RenderEnv{}, fieldspec.ClosedGrantState), "form_digest", "re-render")
}

func TestGateCategoryEnumAndRaiseOnly(t *testing.T) {
	reg := loadRegistry(t)

	if class, raised := reg.ClassifyGateCategory("irreversible_write", false); class != "A" || raised {
		t.Fatalf("irreversible_write class=%s raised=%v, want A false", class, raised)
	}
	if class, raised := reg.ClassifyGateCategory("routing", false); class != "B" || raised {
		t.Fatalf("routing class=%s raised=%v, want B false", class, raised)
	}
	if class, raised := reg.ClassifyGateCategory("unlisted", false); class != "A" || !raised {
		t.Fatalf("unlisted class=%s raised=%v, want A true", class, raised)
	}
	if class, raised := reg.ClassifyGateCategory("routing", true); class != "A" || !raised {
		t.Fatalf("known-A routing class=%s raised=%v, want A true", class, raised)
	}
	if class, raised := reg.ClassifyGateCategory("authz_security", true); class != "A" || raised {
		t.Fatalf("known-A authz_security class=%s raised=%v, want A false", class, raised)
	}
}

func TestValidateGateCategoryRaiseUsesKnownADetector(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s5-b.implementer", Role: "implementer"}
	env := fieldspec.RenderEnv{KnownA: func(record.Record, map[string]string) (string, bool) {
		return "authz_security", true
	}}
	rec := validCandidate()
	rec.Headers["gate_category"] = "routing"

	violations := reg.Validate(rec, seat, digestForEnv(t, reg, seat, env, rec), env, fieldspec.ClosedGrantState)
	if len(violations) != 0 {
		t.Fatalf("violations = %+v, want none", violations)
	}
	if rec.Headers["gate_category"] != "authz_security" {
		t.Fatalf("gate_category = %q, want detector A member", rec.Headers["gate_category"])
	}
	if rec.Headers["gate_category_raised"] != "yes" {
		t.Fatalf("gate_category_raised = %q, want yes", rec.Headers["gate_category_raised"])
	}
	if rec.Headers["gate_category_pick"] != "routing" {
		t.Fatalf("gate_category_pick = %q, want original B pick", rec.Headers["gate_category_pick"])
	}
}

func TestValidateGateCategoryRaiseHandlesAbsentPick(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s5-b.implementer", Role: "implementer"}
	env := fieldspec.RenderEnv{KnownA: func(record.Record, map[string]string) (string, bool) {
		return "authz_security", true
	}}
	rec := validCandidate()

	violations := reg.Validate(rec, seat, digestForEnv(t, reg, seat, env, rec), env, fieldspec.ClosedGrantState)
	if len(violations) != 0 {
		t.Fatalf("violations = %+v, want none", violations)
	}
	if rec.Headers["gate_category"] != "authz_security" {
		t.Fatalf("gate_category = %q, want detector A member", rec.Headers["gate_category"])
	}
	if rec.Headers["gate_category_raised"] != "yes" {
		t.Fatalf("gate_category_raised = %q, want yes", rec.Headers["gate_category_raised"])
	}
	if _, ok := rec.Headers["gate_category_pick"]; ok {
		t.Fatalf("gate_category_pick unexpectedly present: %q", rec.Headers["gate_category_pick"])
	}
}

func TestValidateGateCategoryDoesNotLowerAOrIndexFloorB(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s5-b.implementer", Role: "implementer"}
	env := fieldspec.RenderEnv{
		KnownA: func(record.Record, map[string]string) (string, bool) {
			return "routing", true
		},
		MonotonicFloors: map[string]string{"gate_category": "authz_security"},
	}
	rec := validCandidate()
	rec.Headers["gate_category"] = "authz_security"
	violations := reg.Validate(rec, seat, digestForEnv(t, reg, seat, env, rec), env, fieldspec.ClosedGrantState)
	if len(violations) != 0 {
		t.Fatalf("A-pick violations = %+v, want none", violations)
	}
	if rec.Headers["gate_category"] != "authz_security" || rec.Headers["gate_category_raised"] != "" {
		t.Fatalf("A-pick mutated headers = %+v", rec.Headers)
	}

	bPick := validCandidate()
	bPick.Headers["gate_category"] = "routing"
	noDetector := fieldspec.RenderEnv{MonotonicFloors: map[string]string{"gate_category": "authz_security"}}
	violations = reg.Validate(bPick, seat, digestForEnv(t, reg, seat, noDetector, bPick), noDetector, fieldspec.ClosedGrantState)
	assertNoViolation(t, violations, "gate_category", "monotonic-floor")
	if bPick.Headers["gate_category"] != "routing" || bPick.Headers["gate_category_raised"] != "" {
		t.Fatalf("B-pick without detector mutated headers = %+v", bPick.Headers)
	}
}

func TestValidateGateCategoryOtherStampUsesYesByte(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s5-b.implementer", Role: "implementer"}
	rec := validCandidate()
	rec.Headers["gate_category"] = "other"

	violations := reg.Validate(rec, seat, digestFor(t, reg, seat, rec), fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	if len(violations) != 0 {
		t.Fatalf("violations = %+v, want none", violations)
	}
	if rec.Headers["gate_category"] != "other" {
		t.Fatalf("gate_category = %q, want other", rec.Headers["gate_category"])
	}
	if rec.Headers["gate_category_raised"] != "yes" {
		t.Fatalf("gate_category_raised = %q, want yes", rec.Headers["gate_category_raised"])
	}
}

func loadRegistry(t *testing.T) *fieldspec.Registry {
	t.Helper()
	reg, err := fieldspec.Load(filepath.Join("registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	return reg
}

func assertViolation(t *testing.T, violations []fieldspec.Violation, field, class string) {
	t.Helper()
	for _, violation := range violations {
		if violation.Field == field && violation.Class == class {
			return
		}
	}
	t.Fatalf("missing violation %s/%s in %+v", field, class, violations)
}

func digestForEnv(t *testing.T, reg *fieldspec.Registry, seat fieldspec.SeatMeta, env fieldspec.RenderEnv, rec record.Record) string {
	t.Helper()
	_, digest := reg.Render(env, seat, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	return digest
}
