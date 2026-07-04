package fixtures_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/store"
)

func initFixtureStore(t *testing.T, root string) *config.Pinned {
	t.Helper()
	sources := writeFixtureConfigSources(t)
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("store Init: %v", err)
	}
	return loadFixturePinned(t, root)
}

func loadFixturePinned(t *testing.T, root string) *config.Pinned {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load store pinned config: %v", err)
	}
	return pinned
}

func writeFixtureConfigSources(t *testing.T) map[string]string {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(root, "engine.json")
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(enginePath, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	if err := os.WriteFile(registryPath, []byte(`{"phase":["AUDIT","DESIGN","DESIGN-REVIEW","PLAN","PLAN-REVIEW","IMPL","REVIEW-FOLD","MERGE-GATE","LIVE-VERIFY","SITREP","RECONCILE"],"authority":["read-only","design-only","plan-only","review-only","implementation","fold-in-only","merge-gated","live-verify","report-only"],"ceremony_tier":["tiny","small","medium","large","production-risk"],"evidence_target":["E1","E2","E3","E4"],"gate_category":{"A":["merge_to_protected","irreversible_write","residual_risk_acceptance","live_verify_skip","ceremony_downgrade","authz_security","product_semantics","scope_expansion"],"B":["merge_feature_to_feature","routing","sequencing","scope_within_bounds"]},"grant":["dispatch-impl","dispatch-merge"]}`), 0o644); err != nil {
		t.Fatalf("write registry config: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath}
}
