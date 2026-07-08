package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
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
	if got, ok := mgr.RealizedMintRef("startup-seat.implementer"); !ok || got != "seat-mint-pivot" {
		t.Fatalf("startup-seat realized pivot = %q, %v; want seat-mint-pivot, true", got, ok)
	}
}

func TestS6RemintCrashBeforeBindingReplacementRecoversBeforeServe(t *testing.T) {
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
	oldCred, err := mgr.Mint("s6-remint-crash.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("Mint old seat: %v", err)
	}
	sock := filepath.Join(os.TempDir(), "frank-s6-remint-crash-"+filepath.Base(root)+".sock")
	t.Cleanup(func() { _ = os.Remove(sock) })
	counter := filepath.Join(root, "rename-counter.log")

	crashCmd, crashStderr := startFrankWithEnv(t, ctx, bin, root, sock, []string{
		"FRANK_TEST_CRASHPOINT=post_rename:3",
		"FRANK_TEST_RENAME_COUNTER=" + counter,
	})
	waitForSocket(t, sock)
	operator, err := channel.DialAuthenticated(ctx, sock, operatorCred.Value)
	if err != nil {
		t.Fatalf("operator dial before crash stderr=%s: %v", crashStderr.String(), err)
	}
	describe, err := operator.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools before crash: %v", err)
	}
	_, err = operator.Call(ctx, "submit", s6SeatMintPayload(t, describe.FormDigest, "s6-remint-crash.implementer", "planner", false))
	if err == nil {
		t.Fatalf("remint submit unexpectedly returned before crash")
	}
	_ = operator.Close()
	waitErr := crashCmd.Wait()
	if waitErr == nil {
		t.Fatalf("crash process exited cleanly; stderr=%s", crashStderr.String())
	}
	assertSIGKILL(t, waitErr)
	pivot := onlySeatMintPivot(t, root, "s6-remint-crash.implementer")
	assertLastRename(t, counter, filepath.Join("records", pivot+".json"))
	_ = os.Remove(sock)

	restarted, restartStderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		if restarted.Process != nil {
			_ = restarted.Process.Kill()
		}
		_ = restarted.Wait()
	})
	waitForSocket(t, sock)
	if old, err := channel.DialAuthenticated(ctx, sock, oldCred.Value); err == nil {
		_ = old.Close()
		t.Fatalf("old credential authenticated on first post-restart attempt; stderr=%s", restartStderr.String())
	}
	row := bindingRowForSeat(t, root, "s6-remint-crash.implementer")
	if got, _ := row["realized_mint_ref"].(string); got != pivot {
		t.Fatalf("realized_mint_ref = %q, want latest pivot %q; row=%#v", got, pivot, row)
	}
	assertStringAbsentFromTree(t, root, "realized_mint_ref", filepath.Join(root, "binding"))
	assertCredentialAbsentFromTree(t, root, oldCred.Value, filepath.Join(root, "binding"))
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

func startFrankWithEnv(t *testing.T, ctx context.Context, bin, root, sock string, env []string) (*exec.Cmd, *bytes.Buffer) {
	t.Helper()
	cmd := exec.CommandContext(ctx, bin, "-root", root, "-socket", sock, "-registry", filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	cmd.Env = append(os.Environ(), env...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("start frank: %v", err)
	}
	return cmd, &stderr
}

func onlySeatMintPivot(t *testing.T, root, seatName string) string {
	t.Helper()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var pivot string
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" {
			continue
		}
		var body struct {
			Seat string `json:"seat"`
		}
		if json.Unmarshal([]byte(rec.Body), &body) != nil || body.Seat != seatName {
			continue
		}
		if pivot != "" {
			t.Fatalf("multiple seat_mint pivots for %s: %s and %s", seatName, pivot, rec.Envelope.RelayID)
		}
		pivot = rec.Envelope.RelayID
	}
	if pivot == "" {
		t.Fatalf("missing accepted seat_mint pivot for %s", seatName)
	}
	return pivot
}

func bindingRowForSeat(t *testing.T, root, seatName string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, "binding", "seats.json"))
	if err != nil {
		t.Fatalf("read binding table: %v", err)
	}
	var table struct {
		Seats map[string]map[string]any `json:"seats"`
	}
	if err := json.Unmarshal(data, &table); err != nil {
		t.Fatalf("decode binding table %s: %v", data, err)
	}
	row := table.Seats[seatName]
	if row == nil {
		t.Fatalf("binding table missing seat %s: %s", seatName, data)
	}
	return row
}

func assertLastRename(t *testing.T, counter, want string) {
	t.Helper()
	data, err := os.ReadFile(counter)
	if err != nil {
		t.Fatalf("read rename counter: %v", err)
	}
	lines := strings.Fields(string(data))
	if len(lines) == 0 {
		t.Fatalf("rename counter empty")
	}
	if got := lines[len(lines)-1]; got != want {
		t.Fatalf("last rename = %q, want %q; all=%q", got, want, string(data))
	}
}

func assertStringAbsentFromTree(t *testing.T, root, forbidden, allowedPrefix string) {
	t.Helper()
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || strings.HasPrefix(path, allowedPrefix) {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if bytes.Contains(data, []byte(forbidden)) {
			t.Fatalf("%q leaked into %s", forbidden, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", root, err)
	}
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
