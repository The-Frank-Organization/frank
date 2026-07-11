package invariants_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type cachedInvariantBinary struct {
	path string
	err  error
}

var frankInvariantBinary cachedInvariantBinary

func TestMain(m *testing.M) {
	if os.Getenv("FRANK_S7_PIVOT_CHILD") == "1" {
		frankInvariantBinary.err = fmt.Errorf("child-mode: no cached binary")
		os.Exit(m.Run())
	}

	dir, err := os.MkdirTemp("", "frank-invariants-")
	if err != nil {
		frankInvariantBinary.err = fmt.Errorf("create invariant binary directory: %w", err)
	} else {
		path := filepath.Join(dir, "frank")
		build := exec.Command("go", "build", "-o", path, filepath.Join("..", "..", "cmd", "frank"))
		if output, buildErr := build.CombinedOutput(); buildErr != nil {
			frankInvariantBinary.err = fmt.Errorf("build frank: %w\n%s", buildErr, output)
		} else {
			frankInvariantBinary.path = path
		}
	}

	code := m.Run()
	if dir != "" {
		if err := os.RemoveAll(dir); err != nil && code == 0 {
			fmt.Fprintf(os.Stderr, "remove invariant binary directory: %v\n", err)
			code = 1
		}
	}
	os.Exit(code)
}
