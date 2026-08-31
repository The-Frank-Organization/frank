package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS6StoreLockSecondConductorRefusesTyped(t *testing.T) {
	root := t.TempDir()
	held, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot held: %v", err)
	}
	defer held.Release()

	_, err = store.AcquireRoot(root)
	var locked store.ErrRootLocked
	if !errors.As(err, &locked) {
		t.Fatalf("second AcquireRoot err = %v, want ErrRootLocked", err)
	}
	if locked.ErrorClass != "root-lock-held" {
		t.Fatalf("ErrorClass = %q, want root-lock-held", locked.ErrorClass)
	}
}

func TestS6SecondFrankProcessRefusesRootLock(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock1 := filepath.Join(os.TempDir(), fmt.Sprintf("frank-lock-1-%d.sock", time.Now().UnixNano()))
	sock2 := filepath.Join(os.TempDir(), fmt.Sprintf("frank-lock-2-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() {
		_ = os.Remove(sock1)
		_ = os.Remove(sock2)
	})
	first, _ := startFrank(t, ctx, bin, root, sock1)
	t.Cleanup(func() {
		if first.Process != nil {
			_ = first.Process.Kill()
		}
		_ = first.Wait()
	})
	waitForSocket(t, sock1)

	var stderr bytes.Buffer
	second := exec.CommandContext(ctx, bin, "-root", root, "-socket", sock2, "-registry", filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	second.Stderr = &stderr
	err := second.Run()
	if err == nil {
		t.Fatalf("second frank process unexpectedly served")
	}
	if !strings.Contains(stderr.String(), "root-lock-held") {
		t.Fatalf("second stderr = %q, want root-lock-held", stderr.String())
	}
	if _, err := os.Stat(sock2); !os.IsNotExist(err) {
		t.Fatalf("second process created socket despite lock refusal: %v", err)
	}
}

func TestS6FrankRootLockTakeoverAfterKilledHolder(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock1 := filepath.Join(os.TempDir(), fmt.Sprintf("frank-lock-takeover-1-%d.sock", time.Now().UnixNano()))
	sock2 := filepath.Join(os.TempDir(), fmt.Sprintf("frank-lock-takeover-2-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() {
		_ = os.Remove(sock1)
		_ = os.Remove(sock2)
	})
	first, _ := startFrank(t, ctx, bin, root, sock1)
	waitForSocket(t, sock1)
	if err := first.Process.Kill(); err != nil {
		t.Fatalf("kill first process: %v", err)
	}
	_ = first.Wait()

	second, _ := startFrank(t, ctx, bin, root, sock2)
	t.Cleanup(func() {
		if second.Process != nil {
			_ = second.Process.Kill()
		}
		_ = second.Wait()
	})
	waitForSocket(t, sock2)
	info := readS6LockInfo(t, root)
	takeover, ok := info["takeover"].(map[string]any)
	if !ok || takeover["event"] != "TAKEOVER" {
		t.Fatalf("lock info missing takeover diagnostic: %#v", info)
	}
	if got := int(takeover["previous_pid"].(float64)); got != first.Process.Pid {
		t.Fatalf("takeover previous_pid = %d, want killed pid %d; info=%#v", got, first.Process.Pid, info)
	}
}

func TestS6FrankRootLockSymlinkAliasRefusesSecondHolder(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	alias := filepath.Join(t.TempDir(), "alias-root")
	if err := os.Symlink(root, alias); err != nil {
		t.Fatalf("Symlink: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock1 := filepath.Join(os.TempDir(), fmt.Sprintf("frank-lock-alias-1-%d.sock", time.Now().UnixNano()))
	sock2 := filepath.Join(os.TempDir(), fmt.Sprintf("frank-lock-alias-2-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() {
		_ = os.Remove(sock1)
		_ = os.Remove(sock2)
	})
	first, _ := startFrank(t, ctx, bin, root, sock1)
	t.Cleanup(func() {
		if first.Process != nil {
			_ = first.Process.Kill()
		}
		_ = first.Wait()
	})
	waitForSocket(t, sock1)

	var stderr bytes.Buffer
	second := exec.CommandContext(ctx, bin, "-root", alias, "-socket", sock2, "-registry", filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	second.Stderr = &stderr
	err := second.Run()
	if err == nil {
		t.Fatalf("alias-root frank process unexpectedly served")
	}
	if !strings.Contains(stderr.String(), "root-lock-held") {
		t.Fatalf("alias second stderr = %q, want root-lock-held", stderr.String())
	}
	if _, err := os.Stat(sock2); !os.IsNotExist(err) {
		t.Fatalf("alias second process created socket despite lock refusal: %v", err)
	}
}

func readS6LockInfo(t *testing.T, root string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, "conductor.lock"))
	if err != nil {
		t.Fatalf("read conductor.lock: %v", err)
	}
	var info map[string]any
	if err := json.Unmarshal(data, &info); err != nil {
		t.Fatalf("decode conductor.lock %s: %v", data, err)
	}
	return info
}
