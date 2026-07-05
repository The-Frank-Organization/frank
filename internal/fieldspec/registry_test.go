package fieldspec_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
)

func TestRegistryV2MemberParsesAndExposesLockedEnums(t *testing.T) {
	reg := loadRegistry(t)

	if reg.Version == "" {
		t.Fatalf("Version empty")
	}
	if reg.Provenance["owner"] == "" {
		t.Fatalf("Provenance owner empty: %#v", reg.Provenance)
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
		"ceremony_downgrade", "authz_security", "product_semantics", "scope_expansion",
	})
	assertTokens(t, reg.NamedEnums["gate_category_B"], []string{
		"merge_feature_to_feature", "routing", "sequencing", "scope_within_bounds",
	})
	assertTokens(t, reg.NamedEnums["gate_category"], []string{
		"merge_to_protected", "irreversible_write", "residual_risk_acceptance", "live_verify_skip",
		"ceremony_downgrade", "authz_security", "product_semantics", "scope_expansion",
		"merge_feature_to_feature", "routing", "sequencing", "scope_within_bounds", "other",
	})
	assertTokens(t, reg.NamedEnums["grant"], []string{"dispatch-impl", "dispatch-merge"})
	assertTokens(t, reg.NamedEnums["delivery_state"], []string{"accepted", "rejected", "held"})
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
