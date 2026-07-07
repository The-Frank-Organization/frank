package fixtures_test

import (
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
)

func TestS6LiveMintBootActivationAndRoster(t *testing.T) {
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
	sock := filepath.Join(os.TempDir(), "frank-s6-lifecycle-"+filepath.Base(root)+".sock")
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

	mint := submitS6LiveMint(t, ctx, operator, "life-seat.implementer", "implementer", false)
	lifeSeat, err := channel.DialAuthenticated(ctx, sock, mint.Credential)
	if err != nil {
		t.Fatalf("life-seat dial stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = lifeSeat.Close() }()

	pre, err := lifeSeat.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("pre-active describe: %v", err)
	}
	for _, id := range []string{"CEREMONY_TIER", "SUBJECT", "charter_loaded", "dispatch_status"} {
		if !pre.SubmitSchema.HasField(id) {
			t.Fatalf("pre-active boot form missing %s: %+v", id, pre.SubmitSchema.Fields)
		}
	}
	for _, id := range []string{"AUTHORITY", "EVIDENCE_TARGET", "record_kind", "PARENT_DISPATCH_ID"} {
		if pre.SubmitSchema.HasField(id) {
			t.Fatalf("pre-active boot form rendered %s: %+v", id, pre.SubmitSchema.Fields)
		}
	}

	nonBoot := mustJSONBytes(t, fieldspec.SubmitPayload{Record: record.Record{
		Headers: map[string]string{"PHASE": "SITREP", "CEREMONY_TIER": "medium", "SUBJECT": "not boot", "AUTHORITY": "report-only"},
	}, FormDigest: pre.FormDigest})
	result, err := lifeSeat.Call(ctx, "submit", nonBoot)
	if err != nil {
		t.Fatalf("non-boot submit call: %v", err)
	}
	var rejected struct {
		State  string `json:"state"`
		Detail string `json:"detail"`
	}
	if err := json.Unmarshal(result, &rejected); err != nil {
		t.Fatalf("decode non-boot outcome %s: %v", result, err)
	}
	if rejected.State != record.Rejected || !strings.Contains(rejected.Detail, "AUTHORITY:non-boot-before-active") {
		t.Fatalf("non-boot outcome = %+v", rejected)
	}

	boot := mustJSONBytes(t, fieldspec.SubmitPayload{Record: record.Record{
		Headers: map[string]string{"PHASE": "SITREP", "CEREMONY_TIER": "medium", "SUBJECT": "boot", "charter_loaded": "yes", "dispatch_status": "read"},
	}, FormDigest: pre.FormDigest})
	result, err = lifeSeat.Call(ctx, "submit", boot)
	if err != nil {
		t.Fatalf("boot submit call: %v", err)
	}
	var booted struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &booted); err != nil {
		t.Fatalf("decode boot outcome %s: %v", result, err)
	}
	if booted.State != record.Accepted || booted.RelayID == "" {
		t.Fatalf("boot outcome = %+v", booted)
	}

	waitForS6OrdinaryForm(t, ctx, lifeSeat)

	rosterRaw, err := operator.Call(ctx, "project", mustJSONBytes(t, map[string]string{"view": "roster"}))
	if err != nil {
		t.Fatalf("operator roster: %v", err)
	}
	var roster []struct {
		Seat                string `json:"seat"`
		ActivationState     string `json:"activation_state"`
		BoundNow            bool   `json:"bound_now"`
		Role                string `json:"role"`
		MintedAt            string `json:"minted_at"`
		ActivationRecordRef string `json:"activation_record_ref"`
		LastAcceptedAt      string `json:"last_accepted_at"`
	}
	if err := json.Unmarshal(rosterRaw, &roster); err != nil {
		t.Fatalf("decode roster %s: %v", rosterRaw, err)
	}
	var found bool
	for _, row := range roster {
		if row.Seat != "life-seat.implementer" {
			continue
		}
		found = true
		if row.ActivationState != "active" || !row.BoundNow || row.Role != "implementer" ||
			row.MintedAt != mint.RelayID || row.ActivationRecordRef != booted.RelayID || row.LastAcceptedAt != booted.RelayID {
			t.Fatalf("roster row = %+v; mint=%s boot=%s", row, mint.RelayID, booted.RelayID)
		}
	}
	if !found {
		t.Fatalf("roster missing life-seat row: %+v", roster)
	}

	refused, err := lifeSeat.Call(ctx, "project", mustJSONBytes(t, map[string]string{"view": "roster"}))
	if err != nil {
		t.Fatalf("non-operator roster call: %v", err)
	}
	if !strings.Contains(string(refused), "roster:seat-scope") {
		t.Fatalf("non-operator roster response = %s", refused)
	}
}

func waitForS6OrdinaryForm(t *testing.T, ctx context.Context, client *channel.Client) channel.DescriptionResponse {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	var last channel.DescriptionResponse
	var lastErr error
	for time.Now().Before(deadline) {
		last, lastErr = client.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
		if lastErr == nil && last.SubmitSchema != nil && last.SubmitSchema.HasField("AUTHORITY") {
			return last
		}
		time.Sleep(25 * time.Millisecond)
	}
	if lastErr != nil {
		t.Fatalf("post-active describe: %v", lastErr)
	}
	t.Fatalf("post-active form did not restore ordinary fields: %+v", last.SubmitSchema.Fields)
	return channel.DescriptionResponse{}
}

type s6LiveMintOutcome struct {
	State      string `json:"state"`
	RelayID    string `json:"relay_id"`
	Credential string `json:"credential"`
	Endpoint   string `json:"endpoint"`
}

func submitS6LiveMint(t *testing.T, ctx context.Context, operator *channel.Client, seatName, role string, isOperator bool) s6LiveMintOutcome {
	t.Helper()
	describe, err := operator.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("operator describe: %v", err)
	}
	result, err := operator.Call(ctx, "submit", s6SeatMintPayload(t, describe.FormDigest, seatName, role, isOperator))
	if err != nil {
		t.Fatalf("seat_mint submit: %v", err)
	}
	var outcome s6LiveMintOutcome
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode seat_mint outcome %s: %v", result, err)
	}
	if outcome.State != record.Accepted || outcome.RelayID == "" || outcome.Credential == "" {
		t.Fatalf("seat_mint outcome = %+v", outcome)
	}
	return outcome
}
