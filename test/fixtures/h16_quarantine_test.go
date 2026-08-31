package fixtures_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16ResolveTypedQuarantineDoesNotLeakToNonmatchingCredential(t *testing.T) {
	mgr, err := seat.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	cred, err := mgr.MintOrReplace("quarantined.implementer", "implementer", false, "pivot-a")
	if err != nil {
		t.Fatalf("mint: %v", err)
	}
	mgr.PublishQuarantine(map[string]bool{"quarantined.implementer": true})
	if _, refusal := mgr.ResolveTyped(cred.Value); refusal != "auth:seat-quarantined" {
		t.Fatalf("matching refusal=%q", refusal)
	}
	if _, refusal := mgr.ResolveTyped("not-a-real-credential"); refusal != "auth:invalid-credential" {
		t.Fatalf("nonmatching refusal leaked quarantine=%q", refusal)
	}
	mgr.PublishQuarantine(nil)
	if meta, refusal := mgr.ResolveTyped(cred.Value); refusal != "" || meta.Name != "quarantined.implementer" {
		t.Fatalf("clear result meta=%+v refusal=%q", meta, refusal)
	}
}

func TestH16ConflictQuarantinePublishedBeforeServeOtherSeatsStillAuthenticate(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	h16Commit(t, st, h16LegacyPivot("conflict-a", "conflict.implementer"))
	h16Commit(t, st, h16LegacyPivot("conflict-b", "conflict.implementer"))
	if err := os.RemoveAll(filepath.Join(root, "journal", "redo")); err != nil {
		t.Fatalf("remove redo: %v", err)
	}
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open seats: %v", err)
	}
	conflicted, err := mgr.MintOrReplace("conflict.implementer", "implementer", false, "conflict-a")
	if err != nil {
		t.Fatalf("mint conflicted: %v", err)
	}
	healthy, err := mgr.Mint("healthy.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("mint healthy: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	sock := filepath.Join(os.TempDir(), "frank-h16-quarantine-"+filepath.Base(root)+".sock")
	cmd, stderr := startFrank(t, ctx, buildFrank(t, ctx), root, sock)
	defer func() {
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
		_ = os.Remove(sock)
	}()
	waitForSocket(t, sock)
	if client, err := channel.DialAuthenticated(ctx, sock, conflicted.Value); err == nil {
		_ = client.Close()
		t.Fatal("conflicted credential authenticated")
	} else if !strings.Contains(err.Error(), "auth:seat-quarantined") {
		t.Fatalf("conflicted refusal=%v stderr=%s", err, stderr.String())
	}
	if _, err := channel.DialAuthenticated(ctx, sock, "not-real"); err == nil || !strings.Contains(err.Error(), "auth:invalid-credential") {
		t.Fatalf("nonmatching refusal=%v", err)
	}
	if client, err := channel.DialAuthenticated(ctx, sock, healthy.Value); err != nil {
		t.Fatalf("healthy seat blocked: %v stderr=%s", err, stderr.String())
	} else {
		_ = client.Close()
	}
}
