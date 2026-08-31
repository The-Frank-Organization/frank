package fieldspec_test

import (
	"encoding/json"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/fieldspec"
)

func TestPredicateAtomsAndCombinators(t *testing.T) {
	reg := loadRegistry(t)
	fields := map[string]string{
		"AUTHORITY":       "implementation",
		"GRILL_REQUIRED":  "yes",
		"GRILL_LOCK_ID":   "s3-grill-s3-form",
		"ACTIONS_GIT_REF": "branch@sha",
	}
	rows := func(array, field string) []string {
		if array == "routing_assignments" && field == "declared_deviated" {
			return []string{"no", "yes"}
		}
		return nil
	}
	ctx := fieldspec.EvalContext{
		Seat:          fieldspec.SeatMeta{Name: "s3-form.implementer", Role: "implementer"},
		Phase:         "IMPL",
		Tier:          "large",
		PresentLayers: fieldspec.DefaultLayers(),
	}

	cases := []struct {
		name string
		raw  string
		want bool
	}{
		{"phase true", `{"phase_in":["IMPL","PLAN"]}`, true},
		{"phase false", `{"phase_in":["AUDIT"]}`, false},
		{"authority true", `{"authority_in":["implementation"]}`, true},
		{"authority false", `{"authority_in":["read-only"]}`, false},
		{"seat true", `{"seat_is":["s3-form.implementer"]}`, true},
		{"role false", `{"role_in":["planner"]}`, false},
		{"tier gte true", `{"ceremony_tier_gte":"medium"}`, true},
		{"tier gte false", `{"ceremony_tier_gte":"production-risk"}`, false},
		{"field equals", `{"field":"GRILL_REQUIRED","op":"==","value":"yes"}`, true},
		{"field in", `{"field":"GRILL_REQUIRED","op":"in","value":["yes"]}`, true},
		{"field present", `{"field":"GRILL_REQUIRED","op":"present"}`, true},
		{"any row equals", `{"any_row":"routing_assignments.declared_deviated","op":"==","value":"yes"}`, true},
		{"layer default", `{"layer_present":"lineage"}`, true},
		{"layer absent", `{"layer_present":"observe"}`, false},
		{"all of", `{"all_of":[{"phase_in":["IMPL"]},{"field":"GRILL_REQUIRED","op":"==","value":"yes"}]}`, true},
		{"any of", `{"any_of":[{"phase_in":["AUDIT"]},{"role_in":["implementer"]}]}`, true},
		{"not", `{"not":{"phase_in":["AUDIT"]}}`, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pred := parsePredicate(t, reg, tc.raw)
			if got := pred.Eval(fields, rows, ctx); got != tc.want {
				t.Fatalf("Eval(%s) = %v, want %v", tc.raw, got, tc.want)
			}
		})
	}
}

func TestPredicateObserveLayerContext(t *testing.T) {
	reg := loadRegistry(t)
	pred := parsePredicate(t, reg, `{"all_of":[{"claims_action":true},{"layer_present":"observe"}]}`)
	fields := map[string]string{"ACTIONS_GIT_REF": "branch@sha"}
	ctx := fieldspec.EvalContext{PresentLayers: fieldspec.DefaultLayers()}

	if pred.Eval(fields, nilRows, ctx) {
		t.Fatalf("observe predicate fired in default Step-1 layers")
	}
	ctx.PresentLayers["observe"] = true
	if !pred.Eval(fields, nilRows, ctx) {
		t.Fatalf("observe predicate did not fire with observe layer present")
	}
}

func TestPredicateParseRejectsInvalidAtoms(t *testing.T) {
	reg := loadRegistry(t)
	for _, raw := range []string{
		`{"field":"gate_category_raised","op":"present"}`,
		`{"field":"GRILL_REQUIRED","op":"contains","value":"yes"}`,
		`{"unknown_atom":"x"}`,
	} {
		t.Run(raw, func(t *testing.T) {
			var msg json.RawMessage = []byte(raw)
			if _, err := fieldspec.ParsePredicate(msg, reg); err == nil {
				t.Fatalf("ParsePredicate(%s) succeeded, want error", raw)
			}
		})
	}
}

func parsePredicate(t *testing.T, reg *fieldspec.Registry, raw string) *fieldspec.Predicate {
	t.Helper()
	pred, err := fieldspec.ParsePredicate(json.RawMessage(raw), reg)
	if err != nil {
		t.Fatalf("ParsePredicate(%s): %v", raw, err)
	}
	return pred
}

func nilRows(array, field string) []string { return nil }
