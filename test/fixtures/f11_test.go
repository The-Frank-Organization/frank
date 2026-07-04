package fixtures_test

import (
	"slices"
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
	} {
		if !slices.Contains(names, want) {
			t.Fatalf("crashpoint registry missing %s in %v", want, names)
		}
	}
}
