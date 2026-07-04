package fieldspec_test

import (
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

func TestRenderA2EnumAndGrantVisibility(t *testing.T) {
	reg := loadRegistry(t)

	pairForm, _ := reg.Render(fieldspec.SeatMeta{Name: "s1-core", Role: "implementer"}, "MERGE-GATE", "medium")
	if pairForm.HasField("grant") {
		t.Fatalf("pair form unexpectedly renders grant")
	}
	if pairForm.OptionAllowed("AUTHORITY", "merge-gated") {
		t.Fatalf("pair form unexpectedly allows merge-gated")
	}

	operatorForm, _ := reg.Render(fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "MERGE-GATE", "medium")
	if !operatorForm.HasField("grant") {
		t.Fatalf("operator merge-gate form did not render grant")
	}
	if !operatorForm.OptionAllowed("grant", "dispatch-merge") {
		t.Fatalf("operator merge-gate form did not allow dispatch-merge")
	}

	orchestratorForm, _ := reg.Render(fieldspec.SeatMeta{Name: "s1.orchestrator-planner", Role: "orchestrator-planner"}, "IMPL", "medium")
	if !orchestratorForm.HasField("grant") || !orchestratorForm.OptionAllowed("grant", "dispatch-impl") {
		t.Fatalf("orchestrator form missing dispatch-impl grant: %#v", orchestratorForm.Fields["grant"])
	}
}

func TestValidateB3AndStaleDigest(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s1-core", Role: "implementer"}
	_, digest := reg.Render(seat, "SITREP", "medium")

	rec := record.Record{
		Envelope: record.Envelope{From: "s1-core.implementer", Role: "implementer", SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":     "NOPE",
			"AUTHORITY": "merge-gated",
		},
	}
	violations := reg.Validate(rec, seat, digest)
	assertViolation(t, violations, "SUBJECT", "required")
	assertViolation(t, violations, "PHASE", "enum")
	assertViolation(t, violations, "AUTHORITY", "seat-scope")

	valid := record.Record{
		Envelope: record.Envelope{From: "s1-core.implementer", Role: "implementer", SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "ok"},
	}
	assertViolation(t, reg.Validate(valid, seat, "stale"), "form_digest", "re-render")
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
