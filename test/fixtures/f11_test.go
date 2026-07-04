package fixtures_test

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/crashpoint"
)

func TestF11CrashpointRegistryCoversS1MutationBoundaries(t *testing.T) {
	names := crashpoint.Names()
	for _, want := range []string{
		"pre_record_fsync",
		"post_record_fsync",
		"pre_rename",
		"post_rename",
		"pre_dir_fsync",
		"post_dir_fsync",
		"pre_redo_fsync",
		"post_redo_fsync",
		"pre_projection_write",
		"post_projection_write",
		"pre_delivery_write",
		"post_delivery_write",
		"pre_outcome_reply",
	} {
		if !slices.Contains(names, want) {
			t.Fatalf("crashpoint registry missing %s in %v", want, names)
		}
	}
}

func TestF11CrashpointRegistryNamesAreLiveInSource(t *testing.T) {
	source := readSource(t,
		filepath.Join("..", "..", "internal", "fsio", "fsio.go"),
		filepath.Join("..", "..", "internal", "store", "store.go"),
		filepath.Join("..", "..", "internal", "store", "projections.go"),
		filepath.Join("..", "..", "internal", "intake", "journal.go"),
		filepath.Join("..", "..", "internal", "engine", "loop.go"),
	)
	for _, name := range crashpoint.Names() {
		needle := `crashpoint.Hit("` + name + `")`
		if !strings.Contains(source, needle) {
			t.Fatalf("registered crashpoint %s has no live Hit site", name)
		}
	}
}

func readSource(t *testing.T, paths ...string) string {
	t.Helper()
	var b strings.Builder
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		b.Write(data)
		b.WriteByte('\n')
	}
	return b.String()
}
