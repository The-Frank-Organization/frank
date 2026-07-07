package fixtures_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/store"
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
