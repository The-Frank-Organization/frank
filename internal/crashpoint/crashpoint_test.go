package crashpoint_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/crashpoint"
)

func TestNamesAndNoopHitWithoutEnv(t *testing.T) {
	want := []string{
		"post_intake_fsync",
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
		"pre_quarantine_evict",
		"post_quarantine_evict",
		"pre_segment_rotate",
		"post_segment_rotate",
		"recovery_post_phase0",
		"recovery_post_phase1",
		"recovery_post_phase2",
		"recovery_post_phase3",
		"recovery_post_phase3_5",
		"recovery_post_phase3_6",
		"recovery_post_phase4",
		"stale_reissue_after_held",
		"pre_gc_marker",
		"post_gc_marker",
		"pre_gc_unlink",
		"post_gc_unlink",
	}
	if got := crashpoint.Names(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Names() = %v, want %v", got, want)
	}
	crashpoint.Hit("pre_rename")
}

func TestHitTraceAppendsNamesWithoutArmingCrashpoint(t *testing.T) {
	trace := filepath.Join(t.TempDir(), "hits.log")
	t.Setenv("FRANK_TEST_HIT_TRACE", trace)

	crashpoint.Hit("pre_rename")
	crashpoint.Hit("post_rename")

	data, err := os.ReadFile(trace)
	if err != nil {
		t.Fatalf("read trace: %v", err)
	}
	if got, want := string(data), "pre_rename\npost_rename\n"; got != want {
		t.Fatalf("trace = %q, want %q", got, want)
	}
}
