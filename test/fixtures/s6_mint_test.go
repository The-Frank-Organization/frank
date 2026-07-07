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

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestS6LiveSeatMintReplyCredentialNeverPersists(t *testing.T) {
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
	sock := filepath.Join(os.TempDir(), "frank-s6-mint-"+filepath.Base(root)+".sock")
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
		t.Fatalf("DialAuthenticated operator stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = operator.Close() }()

	describe, err := operator.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools: %v", err)
	}
	payload := s6SeatMintPayload(t, describe.FormDigest, "s6-live.implementer", "implementer", false)
	result, err := operator.Call(ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit seat_mint stderr=%s: %v", stderr.String(), err)
	}
	var outcome struct {
		State      string `json:"state"`
		RelayID    string `json:"relay_id"`
		Credential string `json:"credential"`
		Endpoint   string `json:"endpoint"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode outcome %s: %v", result, err)
	}
	if outcome.State != record.Accepted || outcome.RelayID == "" || outcome.Credential == "" || outcome.Endpoint == "" {
		t.Fatalf("seat_mint outcome = %+v", outcome)
	}
	client, err := channel.DialAuthenticated(ctx, sock, outcome.Credential)
	if err != nil {
		t.Fatalf("new credential did not authenticate stderr=%s: %v", stderr.String(), err)
	}
	_ = client.Close()

	assertCredentialAbsentFromTree(t, root, outcome.Credential, filepath.Join(root, "binding"))
}

func TestS6RemintReplacesBindingAndForceClosesOldCredential(t *testing.T) {
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
	oldCred, err := mgr.Mint("s6-remint.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("Mint old seat: %v", err)
	}
	sock := filepath.Join(os.TempDir(), "frank-s6-remint-"+filepath.Base(root)+".sock")
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
	oldClient, err := channel.DialAuthenticated(ctx, sock, oldCred.Value)
	if err != nil {
		t.Fatalf("old credential dial stderr=%s: %v", stderr.String(), err)
	}
	operator, err := channel.DialAuthenticated(ctx, sock, operatorCred.Value)
	if err != nil {
		t.Fatalf("operator dial stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = operator.Close() }()

	describe, err := operator.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools: %v", err)
	}
	result, err := operator.Call(ctx, "submit", s6SeatMintPayload(t, describe.FormDigest, "s6-remint.implementer", "planner", false))
	if err != nil {
		t.Fatalf("submit remint stderr=%s: %v", stderr.String(), err)
	}
	var outcome struct {
		State      string `json:"state"`
		Credential string `json:"credential"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode remint outcome %s: %v", result, err)
	}
	if outcome.State != record.Accepted || outcome.Credential == "" || outcome.Credential == oldCred.Value {
		t.Fatalf("remint outcome = %+v", outcome)
	}
	refreshedMgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("reopen seats: %v", err)
	}
	if _, ok := refreshedMgr.Resolve(oldCred.Value); ok {
		t.Fatalf("old credential still resolves")
	}
	if _, err := oldClient.ListTools(ctx); err == nil {
		t.Fatalf("old live channel still usable after remint")
	}
	if fresh, err := channel.DialAuthenticated(ctx, sock, outcome.Credential); err != nil {
		t.Fatalf("fresh credential did not authenticate: %v", err)
	} else {
		_ = fresh.Close()
	}
}

func TestS6StartupCompletesCommittedInitialSeatMintBinding(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "seat-mint-pivot", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "mint", "record_kind": "seat_mint"},
		Body:     `{"seat":"startup-seat.implementer","role":"implementer","is_operator":false}`,
	}, nil); err != nil {
		t.Fatalf("Commit seat_mint pivot: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), "frank-s6-startup-mint-"+filepath.Base(root)+".sock")
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, _ := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats after startup: %v", err)
	}
	if got := mgr.CredentialsFor("startup-seat.implementer"); got != 1 {
		t.Fatalf("startup-seat credentials = %d, want one binding completed from pivot", got)
	}
}

func s6SeatMintPayload(t *testing.T, formDigest, seatName, role string, isOperator bool) []byte {
	t.Helper()
	body, err := json.Marshal(struct {
		Seat       string `json:"seat"`
		Role       string `json:"role"`
		IsOperator bool   `json:"is_operator"`
	}{Seat: seatName, Role: role, IsOperator: isOperator})
	if err != nil {
		t.Fatalf("marshal mint body: %v", err)
	}
	return mustJSONBytes(t, fieldspec.SubmitPayload{Record: record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "live mint",
			"record_kind":     "seat_mint",
		},
		Body: string(body),
	}, FormDigest: formDigest})
}

func assertCredentialAbsentFromTree(t *testing.T, root, credential string, allowedPrefix string) {
	t.Helper()
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(path, allowedPrefix) {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if bytes.Contains(data, []byte(credential)) {
			t.Fatalf("credential leaked into %s", path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", root, err)
	}
}
