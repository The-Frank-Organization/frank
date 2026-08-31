package fixtures_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/record"
)

func TestS5RegistryDormancySweepAndRecordKindScope(t *testing.T) {
	reg := loadS5Registry(t)
	env := fieldspec.RenderEnv{ConfigDigest: "s5-registry-dormancy"}
	seats := []struct {
		name   string
		meta   fieldspec.SeatMeta
		grants fieldspec.GrantState
	}{
		{"operator", fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, fieldspec.ClosedGrantState},
		{"orchestrator-planner", fieldspec.SeatMeta{Name: "s5.orchestrator-planner", Role: "orchestrator-planner"}, fieldspec.ClosedGrantState},
		{"planner grant off", fieldspec.SeatMeta{Name: "s5-a.planner", Role: "planner"}, fieldspec.ClosedGrantState},
		{"planner grant on", fieldspec.SeatMeta{Name: "s5-a.planner", Role: "planner"}, func(fieldspec.SeatMeta) bool { return true }},
		{"implementer", fieldspec.SeatMeta{Name: "s5-a.implementer", Role: "implementer"}, fieldspec.ClosedGrantState},
		{"orchestrator-reviewer", fieldspec.SeatMeta{Name: "s5.orchestrator-reviewer", Role: "orchestrator-reviewer"}, fieldspec.ClosedGrantState},
	}

	for _, seat := range seats {
		for _, phase := range reg.NamedEnums["PHASE"] {
			for _, tier := range reg.NamedEnums["CEREMONY_TIER"] {
				t.Run(seat.name+"/"+phase+"/"+tier, func(t *testing.T) {
					form, digest := reg.Render(env, seat.meta, phase, tier, seat.grants)
					again, againDigest := reg.Render(env, seat.meta, phase, tier, seat.grants)
					if digest != againDigest {
						t.Fatalf("digest changed across identical render: %s vs %s", digest, againDigest)
					}
					for _, field := range s5StrictDormantFieldNames() {
						if form.HasField(field) || again.HasField(field) {
							t.Fatalf("%s rendered unexpectedly for %s/%s/%s", field, seat.name, phase, tier)
						}
					}
					if seat.meta.IsOperator {
						if !form.HasField("resolves_gate") || !again.HasField("resolves_gate") {
							t.Fatalf("resolves_gate did not render for operator/%s/%s", phase, tier)
						}
					} else if form.HasField("resolves_gate") || again.HasField("resolves_gate") {
						t.Fatalf("resolves_gate rendered unexpectedly for %s/%s/%s", seat.name, phase, tier)
					}
				})
			}
		}
	}

	for _, seat := range seats {
		form, _ := reg.Render(env, seat.meta, "SITREP", "medium", seat.grants)
		if form.OptionAllowed("record_kind", "genesis") {
			t.Fatalf("%s form allows genesis record_kind: %#v", seat.name, form.Fields["record_kind"].Options)
		}
		for _, operatorOnly := range []string{"owed_item", "owed_disposition", "gate_resolution", "disposition"} {
			got := form.OptionAllowed("record_kind", operatorOnly)
			if got != seat.meta.IsOperator {
				t.Fatalf("%s record_kind %s allowed=%v, want %v; options=%#v", seat.name, operatorOnly, got, seat.meta.IsOperator, form.Fields["record_kind"].Options)
			}
		}
	}
}

func TestS5RegistryRequiredLegsAndPredicateControls(t *testing.T) {
	reg := loadS5Registry(t)
	seat := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	base := record.Record{Headers: map[string]string{
		"PHASE":           "SITREP",
		"AUTHORITY":       "report-only",
		"CEREMONY_TIER":   "medium",
		"SUBJECT":         "s5 registry",
		"EVIDENCE_TARGET": "E1",
	}}
	_, digest := reg.Render(fieldspec.RenderEnv{}, seat, base.Headers["PHASE"], base.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	violations := reg.Validate(base, seat, digest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	assertNoRequiredViolations(t, violations, append(s5BlockAFieldNames(), s5BlockCFieldNames()...)...)

	missingSubject := cloneRecord(base)
	delete(missingSubject.Headers, "SUBJECT")
	assertFixtureViolation(t, validateWithRenderDigest(reg, missingSubject, seat), "SUBJECT", "required")

	missingGrillLock := cloneRecord(base)
	missingGrillLock.Headers["GRILL_REQUIRED"] = "yes"
	missingGrillLock.Headers["DESIGN_LOCK_ID"] = "s5-a-registry-design"
	assertFixtureViolation(t, validateWithRenderDigest(reg, missingGrillLock, seat), "GRILL_LOCK_ID", "required")

	missingTarget := cloneRecord(base)
	delete(missingTarget.Headers, "EVIDENCE_TARGET")
	assertFixtureViolation(t, validateWithRenderDigest(reg, missingTarget, seat), "EVIDENCE_TARGET", "required")

	emptyCandidate := record.Record{Headers: map[string]string{}}
	_, emptyDigest := reg.Render(fieldspec.RenderEnv{}, seat, "", "", fieldspec.ClosedGrantState)
	assertFixtureViolation(t, reg.Validate(emptyCandidate, seat, emptyDigest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState), "EVIDENCE_TARGET", "required")

	observeCtx := fieldspec.EvalContext{PresentLayers: fieldspec.DefaultLayers()}
	observeCtx.PresentLayers["observe"] = true
	defaultCtx := fieldspec.EvalContext{PresentLayers: fieldspec.DefaultLayers()}
	for _, id := range append(s5BlockAFieldNames(), s5BlockCFieldNames()...) {
		field := requireS5Field(t, reg, id)
		if id == "surface_intent" {
			if len(field.RequiredWhen.Raw) != 0 || len(field.VisibleWhen.Raw) != 0 {
				t.Fatalf("surface_intent retains Option-B static predicates required=%s visible=%s", field.RequiredWhen.Raw, field.VisibleWhen.Raw)
			}
			continue
		}
		if len(field.VisibleWhen.Raw) == 0 {
			t.Fatalf("%s missing visible_when", id)
		}
		if field.VisibleWhen.Eval(nil, nil, defaultCtx) {
			t.Fatalf("%s visible_when fired in default layers", id)
		}
		if !field.VisibleWhen.Eval(nil, nil, observeCtx) {
			t.Fatalf("%s visible_when did not fire with observe layer", id)
		}
	}

	justified := requireS5Field(t, reg, "justified_deviation")
	deviationCode := requireS5Field(t, reg, "deviation_reason_code")
	yesRows := func(array, field string) []string {
		if array == "routing_assignments" && field == "declared_deviated" {
			return []string{"yes"}
		}
		return nil
	}
	trueRows := func(array, field string) []string {
		if array == "routing_assignments" && field == "declared_deviated" {
			return []string{"true"}
		}
		return nil
	}
	for _, field := range []*fieldspec.FieldSpec{justified, deviationCode} {
		if !field.RequiredWhen.Eval(nil, yesRows, observeCtx) {
			t.Fatalf("%s required_when did not fire on declared_deviated yes", field.ID)
		}
		if field.RequiredWhen.Eval(nil, trueRows, observeCtx) {
			t.Fatalf("%s required_when fired on declared_deviated true", field.ID)
		}
	}
}

func TestS5RegistryRejectsInvalidLayerAndModelPredicate(t *testing.T) {
	reg := loadS5Registry(t)
	if _, err := fieldspec.ParsePredicate(json.RawMessage(`{"layer_present":"model"}`), reg); err == nil {
		t.Fatalf("layer_present model parsed, want rejection")
	}

	raw := loadS5RegistryJSON(t)
	fields := raw["fields"].([]any)
	subject := fields[0].(map[string]any)
	subject["required_when"] = map[string]any{"field": "model_name", "op": "present"}
	path := writeS5RegistryJSON(t, raw)
	_, err := fieldspec.Load(path)
	if err == nil {
		t.Fatalf("Load mutated registry with model_name predicate succeeded, want rejection")
	}
	if !strings.Contains(err.Error(), "non gate-referenceable field") {
		t.Fatalf("mutated registry error = %q, want non gate-referenceable field", err)
	}
}

func TestS5RegistryRawAnnotations(t *testing.T) {
	fields := rawS5FieldsByID(t)
	m2 := "gated to the post-Step-1 consumer fill-layer; NOT observe-owned (owner stays agent_enum_pick/free_text)."
	for _, id := range s5BlockCFieldNames() {
		assertRawAnnotationContains(t, fields, id, m2)
	}
	assertRawAnnotationContains(t, fields, "record_kind", "with m-1 route-back")
	assertRawAnnotationContains(t, fields, "on_timeout", "no value may ever mean auto-approve/auto-resolve")
	for _, id := range []string{"achieved_evidence", "slot_in", "gate_category_pick", "model_name"} {
		assertRawAnnotationContains(t, fields, id, "suppliability guard = the s5-b (h) typed-REJECT validator rule")
	}
}

func loadS5Registry(t *testing.T) *fieldspec.Registry {
	t.Helper()
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	return reg
}

func loadS5RegistryJSON(t *testing.T) map[string]any {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("decode registry: %v", err)
	}
	return raw
}

func writeS5RegistryJSON(t *testing.T, raw map[string]any) string {
	t.Helper()
	data, err := json.Marshal(raw)
	if err != nil {
		t.Fatalf("marshal registry: %v", err)
	}
	path := filepath.Join(t.TempDir(), "registry.json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("write registry: %v", err)
	}
	return path
}

func rawS5FieldsByID(t *testing.T) map[string]map[string]any {
	t.Helper()
	raw := loadS5RegistryJSON(t)
	out := map[string]map[string]any{}
	for _, item := range raw["fields"].([]any) {
		field := item.(map[string]any)
		out[field["id"].(string)] = field
	}
	return out
}

func requireS5Field(t *testing.T, reg *fieldspec.Registry, id string) *fieldspec.FieldSpec {
	t.Helper()
	field, ok := reg.ByID(id)
	if !ok {
		t.Fatalf("missing %s row", id)
	}
	return field
}

func validateWithRenderDigest(reg *fieldspec.Registry, rec record.Record, seat fieldspec.SeatMeta) []fieldspec.Violation {
	_, digest := reg.Render(fieldspec.RenderEnv{}, seat, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	return reg.Validate(rec, seat, digest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
}

func cloneRecord(rec record.Record) record.Record {
	out := record.Record{Envelope: rec.Envelope, Body: rec.Body, Headers: map[string]string{}}
	for key, value := range rec.Headers {
		out.Headers[key] = value
	}
	return out
}

func assertNoRequiredViolations(t *testing.T, violations []fieldspec.Violation, fields ...string) {
	t.Helper()
	for _, violation := range violations {
		if violation.Class != "required" {
			continue
		}
		for _, field := range fields {
			if violation.Field == field {
				t.Fatalf("unexpected required violation for %s in %+v", field, violations)
			}
		}
	}
}

func assertFixtureViolation(t *testing.T, violations []fieldspec.Violation, field, class string) {
	t.Helper()
	for _, violation := range violations {
		if violation.Field == field && violation.Class == class {
			return
		}
	}
	t.Fatalf("missing violation %s/%s in %+v", field, class, violations)
}

func assertRawAnnotationContains(t *testing.T, fields map[string]map[string]any, id, want string) {
	t.Helper()
	field, ok := fields[id]
	if !ok {
		t.Fatalf("missing raw field %s", id)
	}
	annotation, _ := field["annotation"].(string)
	if !strings.Contains(annotation, want) {
		t.Fatalf("%s annotation = %q, want containing %q", id, annotation, want)
	}
}

func s5StrictDormantFieldNames() []string {
	return []string{
		"achieved_evidence", "target_gap_result", "evidence_integrity", "record_integrity",
		"executable_claim_results", "egress_scan_result", "degradation_notes", "attestation_source",
		"authority_class", "deviated_observed", "bucket_binding_observed", "surface_intent",
		"slot_in", "seat_archetype", "authority_ceiling", "capability_prior_snapshot",
		"routing_record_kind", "template_ref", "outcome_feedback_ref", "subject_ref",
		"decision_deadline", "on_timeout", "completed_proof", "away_bridge_eligible", "model_name",
		"routing_assignments", "justified_deviation", "deviation_reason_code", "constraints",
		"plain_language_change", "why_now", "tradeoffs_risks", "recommendation", "choices",
		"gate_category_pick", "ACTIONS_GIT_REF", "FINAL_GIT_STATUS_SHORT",
	}
}

func s5BlockAFieldNames() []string {
	return []string{
		"achieved_evidence", "target_gap_result", "evidence_integrity", "record_integrity",
		"executable_claim_results", "egress_scan_result", "degradation_notes", "attestation_source",
		"authority_class", "deviated_observed", "bucket_binding_observed", "surface_intent",
	}
}

func s5BlockCFieldNames() []string {
	return []string{
		"routing_record_kind", "routing_assignments", "justified_deviation", "deviation_reason_code",
		"constraints", "plain_language_change", "why_now", "tradeoffs_risks", "recommendation", "choices",
	}
}
