package fieldspec_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

func TestRegistryV2MemberParsesAndExposesLockedEnums(t *testing.T) {
	reg := loadRegistry(t)

	if reg.Version != "s10-fieldspec-v8" {
		t.Fatalf("Version = %q, want s10-fieldspec-v8", reg.Version)
	}
	wantProvenance := map[string]string{
		"owner":         "m-2",
		"design_doc_id": "F-S7-R2-COLGRAIN",
		"plan_lock_id":  "s7a-plan-m2",
		"note":          "s7a m-2 pair build under the operator B10 second-application ruling; restores c1 section 5 column-grain fidelity",
	}
	if !reflect.DeepEqual(reg.Provenance, wantProvenance) {
		t.Fatalf("Provenance = %#v, want %#v", reg.Provenance, wantProvenance)
	}
	if len(reg.NamedEnums) != 25 {
		t.Fatalf("named enum count = %d, want 25", len(reg.NamedEnums))
	}
	if len(reg.Fields) != 92 {
		t.Fatalf("row count = %d, want 92", len(reg.Fields))
	}

	assertTokens(t, reg.NamedEnums["PHASE"], []string{
		"AUDIT", "DESIGN", "DESIGN-REVIEW", "PLAN", "PLAN-REVIEW", "IMPL",
		"REVIEW-FOLD", "MERGE-GATE", "LIVE-VERIFY", "SITREP", "RECONCILE",
	})
	assertTokens(t, reg.NamedEnums["AUTHORITY"], []string{
		"read-only", "design-only", "plan-only", "review-only", "implementation",
		"fold-in-only", "merge-gated", "live-verify", "report-only",
	})
	assertTokens(t, reg.NamedEnums["CEREMONY_TIER"], []string{"tiny", "small", "medium", "large", "production-risk"})
	assertTokens(t, reg.NamedEnums["EVIDENCE_TARGET"], []string{"E1", "E2", "E3", "E4"})
	assertTokens(t, reg.NamedEnums["gate_category_A"], []string{
		"merge_to_protected", "irreversible_write", "residual_risk_acceptance", "live_verify_skip",
		"ceremony_downgrade", "authz_security", "product_semantics", "scope_expansion", "routing_escalation",
	})
	assertTokens(t, reg.NamedEnums["gate_category_B"], []string{
		"merge_feature_to_feature", "routing", "sequencing", "scope_within_bounds",
	})
	assertTokens(t, reg.NamedEnums["gate_category"], []string{
		"merge_to_protected", "irreversible_write", "residual_risk_acceptance", "live_verify_skip",
		"ceremony_downgrade", "authz_security", "product_semantics", "scope_expansion",
		"merge_feature_to_feature", "routing", "sequencing", "scope_within_bounds", "routing_escalation", "other",
	})
	assertTokens(t, reg.NamedEnums["achieved_evidence"], []string{"E0", "E1", "E2", "E3", "E4"})
	assertTokens(t, reg.NamedEnums["target_gap_result"], []string{"met", "target_gt_achieved", "not_applicable"})
	assertTokens(t, reg.NamedEnums["evidence_integrity"], []string{"observed", "self_reported"})
	assertTokens(t, reg.NamedEnums["record_integrity"], []string{"observed", "self_reported", "mixed"})
	assertTokens(t, reg.NamedEnums["executable_claim_result"], []string{"pass", "fail", "skipped", "unsafe"})
	assertTokens(t, reg.NamedEnums["egress_scan_result"], []string{"pass", "blocked", "not_applicable"})
	assertTokens(t, reg.NamedEnums["attestation_source"], []string{"conductor", "operator"})
	assertTokens(t, reg.NamedEnums["deviation_reason_code"], []string{"capability_gap", "cost_budget", "latency_budget", "bucket_unavailable", "operator_directive", "experiment", "other"})
	assertTokens(t, reg.NamedEnums["routing_record_kind"], []string{"routing_decision"})
	assertTokens(t, reg.NamedEnums["surface_intent"], []string{"progress", "review_checkpoint", "advisory", "result"})
	assertTokens(t, reg.NamedEnums["grant"], []string{"dispatch-impl", "dispatch-merge"})
	assertTokens(t, reg.NamedEnums["delivery_state"], []string{"accepted", "rejected", "held"})
	assertTokens(t, reg.NamedEnums["dispatch_status"], []string{"read", "awaiting"})
}

func TestRegistryS5MemberContainsRegistryPassRows(t *testing.T) {
	reg := loadRegistry(t)

	achieved := requireField(t, reg, "achieved_evidence")
	if achieved.Type != "enum" ||
		achieved.EnumSet != "achieved_evidence" ||
		achieved.Owner != "system" ||
		achieved.FillConstraints != "observed_value" ||
		achieved.GateReferenceable {
		t.Fatalf("achieved_evidence row = %+v", achieved)
	}
	assertPredicateRaw(t, achieved.RequiredWhen.Raw, `{"all_of":[{"layer_present":"observe"}]}`)
	assertPredicateRaw(t, achieved.VisibleWhen.Raw, `{"all_of":[{"layer_present":"observe"}]}`)

	onTimeout := requireField(t, reg, "on_timeout")
	if onTimeout.Type != "enum" ||
		onTimeout.EnumSet != "" ||
		len(onTimeout.Options) != 0 ||
		onTimeout.Owner != "system" ||
		onTimeout.FillConstraints != "system_only" {
		t.Fatalf("on_timeout row = %+v", onTimeout)
	}

	routingAssignments := requireField(t, reg, "routing_assignments")
	if routingAssignments.Type != "row_array" ||
		routingAssignments.Owner != "seat_scoped_enum" ||
		routingAssignments.FillConstraints != "seat_allowed_values" ||
		!routingAssignments.GateReferenceable ||
		len(routingAssignments.SeatScope) != 0 {
		t.Fatalf("routing_assignments row = %+v", routingAssignments)
	}
	assertTokens(t, routingAssignments.GateReferenceableColumns, []string{"declared_deviated"})
	assertPredicateRaw(t, routingAssignments.VisibleWhen.Raw, `{"all_of":[{"layer_present":"observe"}]}`)

	gateCategoryPick := requireField(t, reg, "gate_category_pick")
	if gateCategoryPick.Type != "enum" ||
		gateCategoryPick.EnumSet != "gate_category" ||
		gateCategoryPick.Owner != "system" ||
		gateCategoryPick.FillConstraints != "computed_result" ||
		gateCategoryPick.GateReferenceable {
		t.Fatalf("gate_category_pick row = %+v", gateCategoryPick)
	}

	resolvesGate := requireField(t, reg, "resolves_gate")
	if resolvesGate.Type != "id_ref" ||
		resolvesGate.Owner != "free_text" ||
		resolvesGate.FillConstraints != "free_text" ||
		resolvesGate.GateReferenceable {
		t.Fatalf("resolves_gate row = %+v", resolvesGate)
	}
	assertPredicateRaw(t, resolvesGate.VisibleWhen.Raw, `{"any_of":[{"seat_is":["operator"]},{"role_in":["operator"]}]}`)

	evidenceTarget := requireField(t, reg, "EVIDENCE_TARGET")
	assertPredicateRaw(t, evidenceTarget.RequiredWhen.Raw, `{"not":{"phase_in":[]}}`)
	actions := requireField(t, reg, "ACTIONS_GIT_REF")
	assertPredicateRaw(t, actions.VisibleWhen.Raw, `{"all_of":[{"layer_present":"observe"}]}`)
	finalStatus := requireField(t, reg, "FINAL_GIT_STATUS_SHORT")
	assertPredicateRaw(t, finalStatus.VisibleWhen.Raw, `{"all_of":[{"layer_present":"observe"}]}`)

	recordKind := requireField(t, reg, "record_kind")
	assertTokens(t, recordKind.SeatScope["*"], []string{"diagnostics"})
	assertTokens(t, recordKind.SeatScope["operator"], []string{"owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint"})
}

func TestRegistryS6TransportBootAndWaiverRows(t *testing.T) {
	reg := loadRegistry(t)

	parentHint := requireField(t, reg, "parent_hint")
	if parentHint.Layer != "header" ||
		parentHint.Owner != "free_text" ||
		parentHint.Type != "id_ref" ||
		parentHint.FillConstraints != "free_text" ||
		parentHint.LineageRole != "none" ||
		parentHint.GateReferenceable {
		t.Fatalf("parent_hint row = %+v", parentHint)
	}

	parentHintHonored := requireField(t, reg, "parent_hint_honored")
	if parentHintHonored.Owner != "system" ||
		parentHintHonored.Type != "bool" ||
		parentHintHonored.EnumSet != "bool" ||
		parentHintHonored.FillConstraints != "computed_result" ||
		parentHintHonored.GateReferenceable {
		t.Fatalf("parent_hint_honored row = %+v", parentHintHonored)
	}

	parentProvenance := requireField(t, reg, "parent_provenance")
	if parentProvenance.Owner != "system" ||
		parentProvenance.Type != "enum" ||
		parentProvenance.EnumSet != "" ||
		parentProvenance.FillConstraints != "computed_result" ||
		parentProvenance.GateReferenceable {
		t.Fatalf("parent_provenance row = %+v", parentProvenance)
	}
	assertTokens(t, parentProvenance.Options, []string{"woken_on", "active_dispatch", "dispatch_root", "hint"})

	routingRefHonored := requireField(t, reg, "routing_ref_honored")
	if routingRefHonored.Owner != "system" ||
		routingRefHonored.Type != "bool" ||
		routingRefHonored.EnumSet != "bool" ||
		routingRefHonored.FillConstraints != "computed_result" ||
		routingRefHonored.GateReferenceable {
		t.Fatalf("routing_ref_honored row = %+v", routingRefHonored)
	}

	for _, tc := range []struct {
		id      string
		wantTyp string
	}{
		{id: "rationale", wantTyp: "text"},
		{id: "waiver_scope", wantTyp: "object"},
		{id: "retracts", wantTyp: "id_ref"},
	} {
		field := requireField(t, reg, tc.id)
		if field.Layer != "header" ||
			field.Owner != "free_text" ||
			field.Type != tc.wantTyp ||
			field.FillConstraints != "free_text" ||
			field.LineageRole != "none" ||
			field.GateReferenceable {
			t.Fatalf("%s row = %+v", tc.id, field)
		}
		assertPredicateRaw(t, field.VisibleWhen.Raw, `{"any_of":[{"seat_is":["operator"]},{"role_in":["operator"]}]}`)
	}

	charterLoaded := requireField(t, reg, "charter_loaded")
	if charterLoaded.Owner != "agent_enum_pick" ||
		charterLoaded.Type != "bool" ||
		charterLoaded.EnumSet != "bool" ||
		charterLoaded.FillConstraints != "seat_allowed_values" ||
		charterLoaded.GateReferenceable {
		t.Fatalf("charter_loaded row = %+v", charterLoaded)
	}
	dispatchStatus := requireField(t, reg, "dispatch_status")
	if dispatchStatus.Owner != "agent_enum_pick" ||
		dispatchStatus.Type != "enum" ||
		dispatchStatus.EnumSet != "dispatch_status" ||
		dispatchStatus.FillConstraints != "seat_allowed_values" ||
		dispatchStatus.GateReferenceable {
		t.Fatalf("dispatch_status row = %+v", dispatchStatus)
	}

	if _, ok := reg.ByID("ORCH_REVIEW_WAIVER"); ok {
		t.Fatalf("ORCH_REVIEW_WAIVER row still present")
	}
	for _, field := range reg.Fields {
		id := strings.ToLower(field.ID)
		if strings.Contains(id, "activation") || strings.Contains(id, "marker") {
			t.Fatalf("activation-marker-like row present: %s", field.ID)
		}
	}

	operatorForm, _ := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "SITREP", "medium", fieldspec.ClosedGrantState)
	seatForm, _ := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: "seat-a", Role: "implementer"}, "SITREP", "medium", fieldspec.ClosedGrantState)
	for _, id := range []string{"rationale", "waiver_scope", "retracts"} {
		if !operatorForm.HasField(id) {
			t.Fatalf("operator form missing %s", id)
		}
		if seatForm.HasField(id) {
			t.Fatalf("non-operator form rendered %s", id)
		}
	}
}

func TestRegistryV2MemberContainsGrillRows(t *testing.T) {
	reg := loadRegistry(t)

	grillRequired, ok := reg.ByID("GRILL_REQUIRED")
	if !ok {
		t.Fatalf("missing GRILL_REQUIRED row")
	}
	if grillRequired.Type != "bool" ||
		grillRequired.Owner != "agent_enum_pick" ||
		grillRequired.FillConstraints != "monotonic" ||
		!grillRequired.GateReferenceable {
		t.Fatalf("GRILL_REQUIRED row = %+v", grillRequired)
	}
	if !contains(grillRequired.Consumers, "lineage_engine") {
		t.Fatalf("GRILL_REQUIRED consumers = %v, want lineage_engine", grillRequired.Consumers)
	}

	grillLock, ok := reg.ByID("GRILL_LOCK_ID")
	if !ok {
		t.Fatalf("missing GRILL_LOCK_ID row")
	}
	if grillLock.Type != "id_ref" ||
		grillLock.LineageRole != "lock_id" ||
		grillLock.GateReferenceable {
		t.Fatalf("GRILL_LOCK_ID row = %+v", grillLock)
	}
	if len(grillLock.RequiredWhen.Raw) == 0 {
		t.Fatalf("GRILL_LOCK_ID missing required_when")
	}
}

func TestRegistryV2MemberContainsOwedRecordRows(t *testing.T) {
	reg := loadRegistry(t)

	for _, id := range []string{"owner", "source", "target_surface", "disposition_path"} {
		field, ok := reg.ByID(id)
		if !ok {
			t.Fatalf("missing %s row", id)
		}
		if field.Layer != "header" ||
			field.Owner != "free_text" ||
			field.Type != "text" ||
			field.FillConstraints != "free_text" ||
			field.LineageRole != "none" ||
			field.GateReferenceable {
			t.Fatalf("%s row = %+v", id, field)
		}
		assertRecordKindRequiredWhen(t, field, "owed_item")
	}

	field, ok := reg.ByID("disposes_owed")
	if !ok {
		t.Fatalf("missing disposes_owed row")
	}
	if field.Layer != "header" ||
		field.Owner != "free_text" ||
		field.Type != "id_ref" ||
		field.FillConstraints != "free_text" ||
		field.LineageRole != "none" ||
		field.GateReferenceable {
		t.Fatalf("disposes_owed row = %+v", field)
	}
	assertRecordKindRequiredWhen(t, field, "owed_disposition")
}

func TestOwedRecordRequiredWhenPredicatesValidate(t *testing.T) {
	reg := loadRegistry(t)
	meta := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	base := map[string]string{
		"PHASE":         "SITREP",
		"AUTHORITY":     "report-only",
		"CEREMONY_TIER": "medium",
		"SUBJECT":       "owed predicates",
	}
	validate := func(headers map[string]string) []fieldspec.Violation {
		for key, value := range base {
			if _, ok := headers[key]; !ok {
				headers[key] = value
			}
		}
		_, digest := reg.Render(fieldspec.RenderEnv{}, meta, headers["PHASE"], headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
		return reg.Validate(record.Record{Headers: headers}, meta, digest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
	}

	for _, missing := range []string{"owner", "source", "target_surface", "disposition_path"} {
		headers := map[string]string{
			"record_kind":      "owed_item",
			"owner":            "s4",
			"source":           "gate",
			"target_surface":   "form",
			"disposition_path": "fold",
		}
		delete(headers, missing)
		assertViolation(t, validate(headers), missing, "required")
	}

	assertViolation(t, validate(map[string]string{"record_kind": "owed_disposition"}), "disposes_owed", "required")
	assertNoFieldViolations(t, validate(map[string]string{}), "owner", "source", "target_surface", "disposition_path", "disposes_owed")
}

func TestRegistryLoadRejectsBadPredicateReferences(t *testing.T) {
	t.Run("non gate referenceable field", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			trigger := fields[0].(map[string]any)
			trigger["required_when"] = map[string]any{"field": "TARGET", "op": "present"}
		}))
		assertLoadReject(t, path, "TRIGGER", "gate-referenceable")
	})

	t.Run("x namespace predicate", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			trigger := fields[0].(map[string]any)
			trigger["required_when"] = map[string]any{"field": "X-BAD", "op": "present"}
		}))
		assertLoadReject(t, path, "TRIGGER", "X-*")
	})

	t.Run("any row over non gate referenceable field", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			trigger := fields[0].(map[string]any)
			trigger["required_when"] = map[string]any{"any_row": "ROWS.path", "op": "present"}
		}))
		assertLoadReject(t, path, "TRIGGER", "gate-referenceable")
	})
}

func TestRegistryLoadRejectsNonAllowlistedRowColumns(t *testing.T) {
	tests := []struct {
		name          string
		predicateKind string
		column        string
	}{
		{name: "required model identity column", predicateKind: "required_when", column: "chosen_model"},
		{name: "visible model identity column", predicateKind: "visible_when", column: "chosen_model"},
		{name: "required non-model column", predicateKind: "required_when", column: "seat"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
				fields := fixture["fields"].([]any)
				trigger := fields[0].(map[string]any)
				rows := fields[2].(map[string]any)
				rows["id"] = "routing_assignments"
				rows["gate_referenceable"] = true
				trigger[tc.predicateKind] = map[string]any{
					"any_row": "routing_assignments." + tc.column,
					"op":      "present",
				}
			}))
			assertLoadReject(
				t,
				path,
				"TRIGGER",
				"non gate-referenceable row field",
				"routing_assignments."+tc.column,
			)
		})
	}
}

func TestRegistryLoadRejectsInvalidRows(t *testing.T) {
	t.Run("model identity cannot be gate referenceable", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			target := fields[1].(map[string]any)
			target["model_identity"] = true
			target["gate_referenceable"] = true
		}))
		assertLoadReject(t, path, "TARGET", "model_identity")
	})

	t.Run("x row cannot have consumers", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			xrow := fields[3].(map[string]any)
			xrow["consumers"] = []any{"form_validator"}
		}))
		assertLoadReject(t, path, "X-BAD", "X-*")
	})

	t.Run("slot_in is reserved for step one", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			trigger := fields[0].(map[string]any)
			trigger["required_when"] = map[string]any{"slot_in": []any{"slot-a"}}
		}))
		assertLoadReject(t, path, "TRIGGER", "slot_in")
	})

	t.Run("unknown owner", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			trigger := fields[0].(map[string]any)
			trigger["owner"] = "mystery"
		}))
		assertLoadReject(t, path, "TRIGGER", "owner")
	})

	t.Run("missing enum set", func(t *testing.T) {
		path := writeRegistryFixture(t, mutateRegistryFixture(t, func(fixture map[string]any) {
			fields := fixture["fields"].([]any)
			trigger := fields[0].(map[string]any)
			trigger["enum_set"] = "missing"
		}))
		assertLoadReject(t, path, "TRIGGER", "enum_set")
	})
}

func assertRecordKindRequiredWhen(t *testing.T, field *fieldspec.FieldSpec, want string) {
	t.Helper()
	var pred struct {
		AllOf []struct {
			RecordKindIn []string `json:"record_kind_in"`
		} `json:"all_of"`
	}
	if err := json.Unmarshal(field.RequiredWhen.Raw, &pred); err != nil {
		t.Fatalf("%s required_when decode %s: %v", field.ID, field.RequiredWhen.Raw, err)
	}
	if len(pred.AllOf) != 1 || len(pred.AllOf[0].RecordKindIn) != 1 || pred.AllOf[0].RecordKindIn[0] != want {
		t.Fatalf("%s required_when = %s, want record_kind_in %s", field.ID, field.RequiredWhen.Raw, want)
	}
}

func assertNoFieldViolations(t *testing.T, violations []fieldspec.Violation, fields ...string) {
	t.Helper()
	for _, violation := range violations {
		for _, field := range fields {
			if violation.Field == field {
				t.Fatalf("unexpected %s violation in %+v", field, violations)
			}
		}
	}
}

func requireField(t *testing.T, reg *fieldspec.Registry, id string) *fieldspec.FieldSpec {
	t.Helper()
	field, ok := reg.ByID(id)
	if !ok {
		t.Fatalf("missing %s row", id)
	}
	return field
}

func assertPredicateRaw(t *testing.T, got json.RawMessage, want string) {
	t.Helper()
	var gotValue any
	if err := json.Unmarshal(got, &gotValue); err != nil {
		t.Fatalf("predicate decode %s: %v", got, err)
	}
	var wantValue any
	if err := json.Unmarshal([]byte(want), &wantValue); err != nil {
		t.Fatalf("want predicate decode %s: %v", want, err)
	}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Fatalf("predicate = %s, want %s", got, want)
	}
}

func assertTokens(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("tokens length = %d, want %d; got %v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("tokens[%d] = %q, want %q; got %v", i, got[i], want[i], got)
		}
	}
}

func assertLoadReject(t *testing.T, path string, wantParts ...string) {
	t.Helper()
	_, err := fieldspec.Load(path)
	if err == nil {
		t.Fatalf("Load(%s) succeeded, want rejection", path)
	}
	for _, want := range wantParts {
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("error %q missing %q", err, want)
		}
	}
	if strings.Contains(err.Error(), filepath.Dir(path)) {
		t.Fatalf("error leaks path %q: %v", filepath.Dir(path), err)
	}
}

func writeRegistryFixture(t *testing.T, fixture map[string]any) string {
	t.Helper()
	data, err := json.Marshal(fixture)
	if err != nil {
		t.Fatalf("marshal fixture: %v", err)
	}
	path := filepath.Join(t.TempDir(), "registry.json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func mutateRegistryFixture(t *testing.T, mutate func(map[string]any)) map[string]any {
	t.Helper()
	fixture := minimalRegistryFixture()
	mutate(fixture)
	return fixture
}

func minimalRegistryFixture() map[string]any {
	return map[string]any{
		"version":    "test",
		"provenance": map[string]any{"owner": "test"},
		"named_enums": map[string]any{
			"bool": []any{"no", "yes"},
			"tiny": []any{"one"},
		},
		"fields": []any{
			map[string]any{
				"id":                 "TRIGGER",
				"layer":              "header",
				"owner":              "agent_enum_pick",
				"type":               "enum",
				"enum_set":           "tiny",
				"gate_referenceable": true,
				"fill_constraints":   "none",
				"lineage_role":       "none",
			},
			map[string]any{
				"id":                 "TARGET",
				"layer":              "header",
				"owner":              "free_text",
				"type":               "text",
				"gate_referenceable": false,
				"fill_constraints":   "none",
				"lineage_role":       "none",
			},
			map[string]any{
				"id":                 "ROWS",
				"layer":              "header",
				"owner":              "agent_enum_pick",
				"type":               "row_array",
				"gate_referenceable": false,
				"fill_constraints":   "none",
				"lineage_role":       "none",
			},
			map[string]any{
				"id":                 "X-BAD",
				"layer":              "header",
				"owner":              "free_text",
				"type":               "text",
				"gate_referenceable": false,
				"fill_constraints":   "none",
				"lineage_role":       "none",
			},
		},
	}
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
