package fixtures_test

import (
	"encoding/json"
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
	catalogPath := filepath.Join(root, "catalog.json")
	if err := os.WriteFile(enginePath, fixtureEngineConfig(t, false), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	registryBytes, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v2 registry fixture: %v", err)
	}
	if err := os.WriteFile(registryPath, registryBytes, 0o644); err != nil {
		t.Fatalf("write registry config: %v", err)
	}
	catalogBytes, err := os.ReadFile(filepath.Join("..", "invariants", "catalog.v1.json"))
	if err != nil {
		t.Fatalf("read catalog fixture: %v", err)
	}
	if err := os.WriteFile(catalogPath, catalogBytes, 0o644); err != nil {
		t.Fatalf("write catalog config: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath, "catalog": catalogPath}
}

func fixtureEngineConfig(t *testing.T, observe bool) []byte {
	t.Helper()
	raw, err := json.Marshal(map[string]any{
		"version": 3, "gc_enabled": false, "segment_rotate_bytes": 4194304,
		"present_layers": map[string]bool{"observe": observe},
		"supply": map[string]any{
			"lane_roots":  map[string]string{"repo": fixtureRepoRoot(t)},
			"lane_vcs":    map[string]string{"repo": "git"},
			"schema_refs": map[string]string{},
			"suites": map[string]any{"dogfood-battery": map[string]any{
				"lane": "repo", "command": "scripts/dogfood-suite.sh", "args": []string{},
				"timeout_class": "suite_bounded", "timeout_seconds": 120,
			}},
		},
	})
	if err != nil {
		t.Fatalf("marshal fixture engine config: %v", err)
	}
	return raw
}

func fixtureRepoRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("fixture repo root: %v", err)
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatalf("canonical fixture repo root: %v", err)
	}
	return root
}
