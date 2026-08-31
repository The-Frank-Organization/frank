package fixtures_test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/derived"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

const (
	h16StartupMintSeat  = "s12-startup.implementer"
	h16StartupMintPivot = "h16-startup-mint-pivot"
)

func TestH16StartupFoldsRealizedMintEvidenceBeforeServing(t *testing.T) {
	if os.Getenv("FRANK_H16_MINT_CRASH_CHILD") == "1" {
		h16CrashAfterMintPersistence(t)
		return
	}
	if s12SkipOuterRunOnly(t) {
		return
	}

	t.Run("proof failure blocks channel", func(t *testing.T) {
		root, _ := h16PrepareRealizedMintCrashCut(t)
		recordsDir := filepath.Join(root, "records")
		info, err := os.Stat(recordsDir)
		if err != nil {
			t.Fatalf("stat records directory: %v", err)
		}
		if err := os.Chmod(recordsDir, 0o500); err != nil {
			t.Fatalf("make records directory read-only: %v", err)
		}
		defer func() { _ = os.Chmod(recordsDir, info.Mode().Perm()) }()

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		sock := filepath.Join(os.TempDir(), "frank-h16-fold-blocked-"+time.Now().Format("150405.000000000")+".sock")
		defer func() { _ = os.Remove(sock) }()
		cmd, stderr := startFrank(t, ctx, buildFrank(t, ctx), root, sock)
		exited := make(chan error, 1)
		go func() { exited <- cmd.Wait() }()
		deadline := time.NewTimer(5 * time.Second)
		defer deadline.Stop()
		poll := time.NewTicker(10 * time.Millisecond)
		defer poll.Stop()
		for {
			select {
			case err := <-exited:
				if err == nil {
					t.Fatalf("startup unexpectedly succeeded while evidence transition could not commit")
				}
				if _, statErr := os.Stat(sock); !os.IsNotExist(statErr) {
					t.Fatalf("channel exists after failed evidence fold: stat=%v stderr=%s", statErr, stderr.String())
				}
				return
			case <-poll.C:
				if _, err := os.Stat(sock); err == nil {
					_ = cmd.Process.Kill()
					<-exited
					t.Fatalf("channel opened while realized-mint evidence transition could not commit; stderr=%s", stderr.String())
				}
			case <-deadline.C:
				_ = cmd.Process.Kill()
				<-exited
				t.Fatalf("startup neither failed nor opened a channel; stderr=%s", stderr.String())
			}
		}
	})

	t.Run("normal restart folds once without remint", func(t *testing.T) {
		root, bindingBefore := h16PrepareRealizedMintCrashCut(t)
		ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
		defer cancel()
		bin := buildFrank(t, ctx)
		sock := filepath.Join(os.TempDir(), "frank-h16-fold-"+time.Now().Format("150405.000000000")+".sock")
		defer func() { _ = os.Remove(sock) }()

		cmd, stderr := startFrank(t, ctx, bin, root, sock)
		waitForSocket(t, sock)
		st, err := store.Open(root)
		if err != nil {
			t.Fatalf("open restarted store: %v", err)
		}
		assertRealizedUndeliveredTransition(t, st, h16StartupMintPivot, 1)
		if status := engineDerivedStatus(t, st, h16StartupMintPivot); status != "failed" {
			t.Fatalf("startup evidence fold=%q, want failed", status)
		}
		if after := mustReadFile(t, filepath.Join(root, "binding", "seats.json")); !bytes.Equal(after, bindingBefore) {
			t.Fatalf("startup reminted realized pivot\nbefore=%s\nafter=%s", bindingBefore, after)
		}
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
		_ = os.Remove(sock)

		ctx2, cancel2 := context.WithTimeout(context.Background(), 12*time.Second)
		defer cancel2()
		cmd2, stderr2 := startFrank(t, ctx2, bin, root, sock)
		defer func() {
			_ = cmd2.Process.Kill()
			_ = cmd2.Wait()
		}()
		waitForSocket(t, sock)
		assertRealizedUndeliveredTransition(t, st, h16StartupMintPivot, 1)
		if after := mustReadFile(t, filepath.Join(root, "binding", "seats.json")); !bytes.Equal(after, bindingBefore) {
			t.Fatalf("idempotent restart reminted realized pivot\nbefore=%s\nafter=%s\nstderr=%s/%s", bindingBefore, after, stderr.String(), stderr2.String())
		}
	})
}

func h16PrepareRealizedMintCrashCut(t *testing.T) (string, []byte) {
	t.Helper()
	root := t.TempDir()
	initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open crash-cut store: %v", err)
	}
	source := record.Record{
		Envelope: record.Envelope{RelayID: h16StartupMintPivot, From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":            "SITREP",
			"SUBJECT":          "H16 startup evidence crash cut",
			"record_kind":      "seat_mint",
			"hook_contract":    derived.HookContractV1,
			"mint_predecessor": "genesis",
		},
		Body: `{"seat":"` + h16StartupMintSeat + `","role":"implementer","is_operator":false}`,
	}
	if _, err := st.Commit(source, nil); err != nil {
		t.Fatalf("commit mint source: %v", err)
	}
	if _, err := st.Commit(derived.AttemptRecord(h16StartupMintPivot, "mint", "none"), nil); err != nil {
		t.Fatalf("commit open mint marker: %v", err)
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestH16StartupFoldsRealizedMintEvidenceBeforeServing$", "-test.count=1")
	cmd.Env = append(os.Environ(),
		"FRANK_H16_MINT_CRASH_CHILD=1",
		"FRANK_H16_MINT_CRASH_ROOT="+root,
	)
	err = cmd.Run()
	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("mint child did not crash: %v", err)
	}
	status, ok := exitErr.Sys().(syscall.WaitStatus)
	if !ok || !status.Signaled() || status.Signal() != syscall.SIGKILL {
		t.Fatalf("mint child exit=%v, want SIGKILL", err)
	}
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open realized binding: %v", err)
	}
	if realized, ok := mgr.RealizedMintRef(h16StartupMintSeat); !ok || realized != h16StartupMintPivot {
		t.Fatalf("realized_mint_ref=(%q,%v), want (%q,true)", realized, ok, h16StartupMintPivot)
	}
	if status := engineDerivedStatus(t, st, h16StartupMintPivot); status != "unknown" {
		t.Fatalf("crash-cut fold=%q, want unknown before startup", status)
	}
	return root, mustReadFile(t, filepath.Join(root, "binding", "seats.json"))
}

func h16CrashAfterMintPersistence(t *testing.T) {
	t.Helper()
	mgr, err := seat.Open(os.Getenv("FRANK_H16_MINT_CRASH_ROOT"))
	if err != nil {
		t.Fatalf("child open seat manager: %v", err)
	}
	if _, err := mgr.MintOrReplace(h16StartupMintSeat, "implementer", false, h16StartupMintPivot); err != nil {
		t.Fatalf("child MintOrReplace: %v", err)
	}
	if err := syscall.Kill(os.Getpid(), syscall.SIGKILL); err != nil {
		t.Fatalf("child SIGKILL: %v", err)
	}
	t.Fatal("child survived SIGKILL")
}
