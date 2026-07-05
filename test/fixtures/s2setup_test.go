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
	registryBytes, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v2 registry fixture: %v", err)
	}
	if err := os.WriteFile(registryPath, registryBytes, 0o644); err != nil {
		t.Fatalf("write registry config: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath}
}
