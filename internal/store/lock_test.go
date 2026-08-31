package store_test

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestAcquireRootExcludesSecondHolderAndReleases(t *testing.T) {
	root := t.TempDir()
	first, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot first: %v", err)
	}
	defer first.Release()

	_, err = store.AcquireRoot(root)
	var locked store.ErrRootLocked
	if !errors.As(err, &locked) {
		t.Fatalf("second AcquireRoot err = %v, want ErrRootLocked", err)
	}
	if locked.HolderPID == 0 {
		t.Fatalf("locked error missing holder pid: %+v", locked)
	}

	if err := first.Release(); err != nil {
		t.Fatalf("Release: %v", err)
	}
	second, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot after release: %v", err)
	}
	defer second.Release()
}

func TestAcquireRootTakesOverAfterHolderDeathWithDiagnostic(t *testing.T) {
	if os.Getenv("FRANK_STORE_LOCK_HELPER") == "1" {
		root := os.Getenv("FRANK_STORE_LOCK_ROOT")
		lock, err := store.AcquireRoot(root)
		if err != nil {
			t.Fatalf("helper AcquireRoot: %v", err)
		}
		defer lock.Release()
		select {}
	}

	root := t.TempDir()
	helper := exec.Command(os.Args[0], "-test.run=TestAcquireRootTakesOverAfterHolderDeathWithDiagnostic")
	helper.Env = append(os.Environ(), "FRANK_STORE_LOCK_HELPER=1", "FRANK_STORE_LOCK_ROOT="+root)
	if err := helper.Start(); err != nil {
		t.Fatalf("start helper: %v", err)
	}
	t.Cleanup(func() {
		if helper.Process != nil {
			_ = helper.Process.Kill()
		}
		_ = helper.Wait()
	})
	waitForLockPID(t, root, helper.Process.Pid)
	if err := helper.Process.Kill(); err != nil {
		t.Fatalf("kill helper: %v", err)
	}
	_ = helper.Wait()

	taken, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot after holder death: %v", err)
	}
	defer taken.Release()
	info := readLockInfoForTest(t, root)
	takeover, ok := info["takeover"].(map[string]any)
	if !ok {
		t.Fatalf("lock info missing takeover diagnostic: %#v", info)
	}
	if takeover["event"] != "TAKEOVER" || int(takeover["previous_pid"].(float64)) != helper.Process.Pid {
		t.Fatalf("takeover diagnostic = %#v, want previous pid %d", takeover, helper.Process.Pid)
	}
}

func TestAcquireRootSymlinkAliasHasOneWinner(t *testing.T) {
	realRoot := t.TempDir()
	aliasRoot := filepath.Join(t.TempDir(), "alias-root")
	if err := os.Symlink(realRoot, aliasRoot); err != nil {
		t.Fatalf("Symlink: %v", err)
	}
	first, err := store.AcquireRoot(realRoot)
	if err != nil {
		t.Fatalf("AcquireRoot real: %v", err)
	}
	defer first.Release()

	_, err = store.AcquireRoot(aliasRoot)
	var locked store.ErrRootLocked
	if !errors.As(err, &locked) {
		t.Fatalf("AcquireRoot alias err = %v, want ErrRootLocked", err)
	}
}

func waitForLockPID(t *testing.T, root string, pid int) {
	t.Helper()
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		info := readLockInfoForTest(t, root)
		if got, ok := info["pid"].(float64); ok && int(got) == pid {
			return
		}
		time.Sleep(25 * time.Millisecond)
	}
	t.Fatalf("lock file did not record helper pid %d; got %#v", pid, readLockInfoForTest(t, root))
}

func readLockInfoForTest(t *testing.T, root string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, "conductor.lock"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		t.Fatalf("read lock info: %v", err)
	}
	if len(data) == 0 {
		return nil
	}
	var info map[string]any
	if err := json.Unmarshal(data, &info); err != nil {
		t.Fatalf("decode lock info %s: %v", data, err)
	}
	return info
}
