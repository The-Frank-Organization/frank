package crashpoint

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

var names = []string{
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
	"pre_gc_marker",
	"post_gc_marker",
	"pre_gc_unlink",
	"post_gc_unlink",
}

var (
	mu     sync.Mutex
	counts = map[string]int{}
)

func Names() []string {
	out := make([]string, len(names))
	copy(out, names)
	return out
}

func Hit(name string) {
	traceHit(name)
	target, nth := target()
	if target == "" || target != name {
		return
	}
	mu.Lock()
	counts[name]++
	count := counts[name]
	mu.Unlock()
	if count == nth {
		_ = syscall.Kill(os.Getpid(), syscall.SIGKILL)
		select {}
	}
}

func traceHit(name string) {
	path := os.Getenv("FRANK_TEST_HIT_TRACE")
	if path == "" {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return
	}
	_, _ = f.WriteString(name + "\n")
	_ = f.Close()
}

func target() (string, int) {
	raw := os.Getenv("FRANK_TEST_CRASHPOINT")
	if raw == "" {
		return "", 0
	}
	name, suffix, ok := strings.Cut(raw, ":")
	if !ok {
		return name, 1
	}
	n, err := strconv.Atoi(suffix)
	if err != nil || n < 1 {
		return name, 1
	}
	return name, n
}
