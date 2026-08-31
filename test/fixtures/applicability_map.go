package fixtures

import "github.com/The-Frank-Organization/frank/internal/crashpoint"

func ApplicabilityMap() map[string]map[string]string {
	classes := []string{
		"submit-accept",
		"submit-reject",
		"config-change",
		"held",
		"operator-verdict",
		"park",
		"outbox-enqueue",
		"genesis",
		"quarantine-disposition",
		"gc-marker",
		"owed-item",
		"owed-disposition",
		"seat-mint",
	}
	out := map[string]map[string]string{}
	for _, class := range classes {
		row := map[string]string{}
		for _, name := range crashpoint.Names() {
			row[name] = "clean-completion"
		}
		out[class] = row
	}
	commit := []string{"pre_redo_fsync", "post_redo_fsync", "pre_record_fsync", "post_record_fsync", "pre_rename", "post_rename", "pre_dir_fsync", "post_dir_fsync"}
	recordOnly := []string{"pre_record_fsync", "post_record_fsync", "pre_rename", "post_rename", "pre_dir_fsync", "post_dir_fsync"}
	projection := []string{"pre_projection_write", "post_projection_write"}
	delivery := []string{"pre_delivery_write", "post_delivery_write"}
	recovery := []string{"recovery_post_phase0", "recovery_post_phase1", "recovery_post_phase2", "recovery_post_phase3", "recovery_post_phase3_5", "recovery_post_phase3_6", "recovery_post_phase4"}
	mark := func(class string, groups ...[]string) {
		for _, group := range groups {
			for _, name := range group {
				out[class][name] = "crash-expected"
			}
		}
	}
	mark("submit-accept", commit, projection, delivery)
	mark("submit-reject", commit)
	mark("config-change", commit, projection)
	mark("held", commit, []string{"pre_outcome_reply"})
	mark("operator-verdict", commit, projection, delivery)
	mark("park", commit, projection, delivery)
	mark("outbox-enqueue", commit, projection)
	mark("genesis", recordOnly, projection, recovery)
	mark("quarantine-disposition", commit, projection, []string{"pre_quarantine_evict", "post_quarantine_evict"})
	mark("gc-marker", commit, projection, []string{"post_intake_fsync"}, []string{"pre_segment_rotate", "post_segment_rotate"}, []string{"pre_gc_marker", "post_gc_marker", "pre_gc_unlink", "post_gc_unlink"})
	mark("owed-item", commit, projection)
	mark("owed-disposition", commit, projection)
	mark("seat-mint", commit, projection)
	return out
}
