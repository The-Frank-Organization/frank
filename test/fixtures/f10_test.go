package fixtures_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/fsio"
)

func TestF10AtomicWriteCrashWindows(t *testing.T) {
	if os.Getenv("FRANK_F10_CHILD") == "1" {
		root := os.Getenv("FRANK_F10_ROOT")
		if err := fsio.WriteFileAtomic(root, "records/one.json", []byte("hello")); err != nil {
			panic(err)
		}
		return
	}

	for _, tc := range []struct {
		name       string
		crashpoint string
		wantFile   bool
	}{
		{name: "before rename", crashpoint: "pre_rename", wantFile: false},
		{name: "after rename before dir fsync", crashpoint: "pre_dir_fsync", wantFile: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			cmd := exec.Command(os.Args[0], "-test.run=TestF10AtomicWriteCrashWindows")
			cmd.Env = append(os.Environ(),
				"FRANK_F10_CHILD=1",
				"FRANK_F10_ROOT="+root,
				"FRANK_TEST_CRASHPOINT="+tc.crashpoint,
			)
			err := cmd.Run()
			if err == nil {
				t.Fatalf("child did not crash")
			}
			path := filepath.Join(root, "records", "one.json")
			_, statErr := os.Stat(path)
			if tc.wantFile && statErr != nil {
				t.Fatalf("expected committed file after %s: %v", tc.crashpoint, statErr)
			}
			if !tc.wantFile && !os.IsNotExist(statErr) {
				t.Fatalf("expected no committed file after %s, stat err=%v", tc.crashpoint, statErr)
			}
		})
	}
}
