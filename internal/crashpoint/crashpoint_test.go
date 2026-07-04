package crashpoint_test

import (
	"reflect"
	"testing"

	"github.com/jackli/frank/internal/crashpoint"
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
	}
	if got := crashpoint.Names(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Names() = %v, want %v", got, want)
	}
	crashpoint.Hit("pre_rename")
}
