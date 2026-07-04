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
	}
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
