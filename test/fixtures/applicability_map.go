package fixtures

import "github.com/jackli/frank/internal/crashpoint"

func ApplicabilityMap() map[string]map[string]string {
	classes := []string{
		"submit-accept",
		"submit-reject",
		"held",
		"operator-verdict",
		"park",
		"outbox-enqueue",
		"genesis",
		"quarantine-disposition",
		"gc-marker",
		"owed-item",
		"owed-disposition",
	}
	out := map[string]map[string]string{}
	for _, class := range classes {
		row := map[string]string{}
		for _, name := range crashpoint.Names() {
			row[name] = "clean-completion"
		}
		out[class] = row
	}
	for _, name := range []string{"pre_rename", "post_rename", "pre_record_fsync", "post_record_fsync", "pre_redo_fsync", "post_redo_fsync"} {
		out["submit-accept"][name] = "crash-expected"
		out["submit-reject"][name] = "crash-expected"
		out["held"][name] = "crash-expected"
	}
	for _, name := range []string{"pre_quarantine_evict", "post_quarantine_evict"} {
		out["quarantine-disposition"][name] = "crash-expected"
	}
	for _, name := range []string{"pre_gc_marker", "post_gc_marker", "pre_gc_unlink", "post_gc_unlink"} {
		out["gc-marker"][name] = "crash-expected"
	}
	for _, name := range []string{"recovery_post_phase0", "recovery_post_phase1", "recovery_post_phase2", "recovery_post_phase3", "recovery_post_phase3_5", "recovery_post_phase3_6", "recovery_post_phase4"} {
		out["genesis"][name] = "crash-expected"
	}
	out["owed-item"]["pre_rename"] = "crash-expected"
	out["owed-disposition"]["pre_rename"] = "crash-expected"
	return out
}
