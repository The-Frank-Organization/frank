package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestS6IPHSeatMintReplyCarveOutsScoped(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	operatorCred, err := mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	sock := filepath.Join(os.TempDir(), "frank-s6-iph-"+filepath.Base(root)+".sock")
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)
	operator, err := channel.DialAuthenticated(ctx, sock, operatorCred.Value)
	if err != nil {
		t.Fatalf("operator dial stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = operator.Close() }()

	mint := submitS6LiveMint(t, ctx, operator, "iph-seat.implementer", "implementer", false)
	replyBytes, _ := json.Marshal(mint)
	if !bytes.Contains(replyBytes, []byte(mint.Credential)) || !bytes.Contains(replyBytes, []byte(mint.Endpoint)) {
		t.Fatalf("seat_mint reply missing credential/endpoint carve-outs: %s", replyBytes)
	}
	assertCredentialAbsentFromTree(t, root, mint.Credential, filepath.Join(root, "binding"))

	readRaw, err := operator.Call(ctx, "read", mustJSONBytes(t, map[string]string{"relay_id": mint.RelayID}))
	if err != nil {
		t.Fatalf("read mint record: %v", err)
	}
	for _, forbidden := range []string{mint.Credential, mint.Endpoint, root, filepath.Join(root, "records")} {
		if forbidden != "" && strings.Contains(string(readRaw), forbidden) {
			t.Fatalf("read surface leaked %q in %s", forbidden, readRaw)
		}
	}
	rosterRaw, err := operator.Call(ctx, "project", mustJSONBytes(t, map[string]string{"view": "roster"}))
	if err != nil {
		t.Fatalf("roster: %v", err)
	}
	if strings.Contains(string(rosterRaw), mint.Credential) || strings.Contains(string(rosterRaw), mint.Endpoint) {
		t.Fatalf("roster leaked carve-out material: %s", rosterRaw)
	}
}

func TestS6IPHNewPayloadFamiliesPathFree(t *testing.T) {
	root := t.TempDir()
	outputs := []string{
		bounce.Format(fieldspec.Violation{Field: "AUTHORITY", Class: "non-boot-before-active", Reason: filepath.Join(root, "records", "leak")}),
		(store.ErrRootLocked{ErrorClass: "root-lock-held", HolderPID: 123, HolderStarted: "2026-07-07T00:00:00Z"}).Error(),
		`{"seat":"seat-a","activation_state":"minted","bound_now":false,"role":"implementer","minted_at":"relay-a","activation_record_ref":"","last_accepted_at":""}`,
		`{"parent_hint_honored":"no","parent_provenance":"dispatch_root"}`,
	}
	for _, output := range outputs {
		for _, forbidden := range []string{root, "/records/", "/staging/", "/outbox/", "/binding/", "seats.json"} {
			if strings.Contains(output, forbidden) {
				t.Fatalf("new payload family leaked %q in %s", forbidden, output)
			}
		}
	}
}
