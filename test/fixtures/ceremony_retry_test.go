package fixtures_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/derived"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16CeremonyRetryCreatesDistinctCanonicalRotationEachTime(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	root, st, seatName, firstCredential := setupCeremonyDeliveredAnchor(t)
	second := ceremonyCredential(t, runCeremonyRetryProcess(t, root, seatName, "operator did not receive the first reply", nil, nil))
	third := ceremonyCredential(t, runCeremonyRetryProcess(t, root, seatName, "operator did not acknowledge the second reply", nil, nil))
	if firstCredential == second || second == third || firstCredential == third {
		t.Fatalf("credential values reused: %q %q %q", firstCredential, second, third)
	}

	rotations := ceremonyRetryPivots(t, st, seatName)
	if len(rotations) != 2 {
		t.Fatalf("ceremony retry pivots=%d, want 2: %+v", len(rotations), rotations)
	}
	wantPredecessor := "retry-selected"
	for i, rec := range rotations {
		if rec.Headers["mint_predecessor"] != wantPredecessor || rec.Headers["admin_provenance"] != "ceremony" || rec.Headers["hook_contract"] != "1" {
			t.Fatalf("rotation %d headers=%v, want predecessor=%q ceremony provenance", i, rec.Headers, wantPredecessor)
		}
		var body struct {
			Seat        string `json:"seat"`
			Role        string `json:"role"`
			IsOperator  bool   `json:"is_operator"`
			RetryReason string `json:"retry_reason"`
		}
		if err := json.Unmarshal([]byte(rec.Body), &body); err != nil || body.Seat != seatName || body.Role != "implementer" || body.IsOperator || body.RetryReason == "" {
			t.Fatalf("rotation %d body=%q decoded=%+v err=%v", i, rec.Body, body, err)
		}
		wantPredecessor = rec.Envelope.RelayID
	}
	work := derived.Fold(h16Records(t, st))
	for _, rec := range rotations {
		if got := work[rec.Envelope.RelayID]; got.Status != "" || len(got.Cursor) != 0 {
			t.Fatalf("rotation %s work=%+v, want terminal marker/effect/advance", rec.Envelope.RelayID, got)
		}
	}
	mgr, _ := seat.Open(root)
	if _, ok := mgr.Resolve(firstCredential); ok {
		t.Fatal("initial credential remained valid after retries")
	}
	if _, ok := mgr.Resolve(second); ok {
		t.Fatal("first retry credential remained valid after second retry")
	}
	if _, ok := mgr.Resolve(third); !ok {
		t.Fatal("latest retry credential did not authenticate")
	}
	if realized, ok := mgr.RealizedMintRef(seatName); !ok || realized != rotations[len(rotations)-1].Envelope.RelayID {
		t.Fatalf("realized=%q ok=%v, want latest retry pivot", realized, ok)
	}
}

func TestH16CeremonyRetryInterruptedPivotResumesOrdinaryMarkerMachinery(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	root, st, seatName, initial := setupCeremonyDeliveredAnchor(t)
	crashed := runCeremonyRetryProcess(t, root, seatName, "interrupted before marker", nil, []string{"FRANK_TEST_CRASHPOINT=ceremony_retry_post_pivot"})
	if crashed.err == nil {
		t.Fatalf("post-pivot cut did not crash: output=%q", crashed.output)
	}
	rotations := ceremonyRetryPivots(t, st, seatName)
	if len(rotations) != 1 || rotations[0].Headers["hook_contract"] != "1" {
		t.Fatalf("interrupted rotations=%+v, want one stamped pivot", rotations)
	}
	if got := derived.Fold(h16Records(t, st))[rotations[0].Envelope.RelayID]; got.Status != "pending" {
		t.Fatalf("interrupted pivot work=%+v, want pending", got)
	}
	credential := ceremonyCredential(t, runCeremonyRetryProcess(t, root, seatName, "resume interrupted pivot", nil, nil))
	rotations = ceremonyRetryPivots(t, st, seatName)
	if len(rotations) != 1 {
		t.Fatalf("resume created another pivot: %+v", rotations)
	}
	if got := derived.Fold(h16Records(t, st))[rotations[0].Envelope.RelayID]; got.Status != "" || len(got.Cursor) != 0 {
		t.Fatalf("resumed pivot work=%+v, want terminal", got)
	}
	mgr, _ := seat.Open(root)
	if _, ok := mgr.Resolve(initial); ok {
		t.Fatal("resumed retry retained predecessor credential")
	}
	if _, ok := mgr.Resolve(credential); !ok {
		t.Fatal("resumed retry credential invalid")
	}
}

func TestH16CeremonyRetryScopeIsCanonicalAndFailClosed(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	for _, tc := range []struct {
		name        string
		provenance  string
		undelivered bool
		wantAllowed bool
	}{
		{name: "ordinary in-band delivery refused"},
		{name: "ceremony tip accepted", provenance: "ceremony", wantAllowed: true},
		{name: "future provenance refused", provenance: "future-v2"},
		{name: "realized-undelivered accepted", undelivered: true, wantAllowed: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			st, _ := store.Open(root)
			seatName := "retry-scope.implementer"
			pivot := h16LegacyPivot("retry-scope-tip", seatName)
			if tc.provenance != "" {
				pivot.Headers["admin_provenance"] = tc.provenance
			}
			h16Commit(t, st, pivot)
			mgr, _ := seat.Open(root)
			old, err := mgr.MintOrReplace(seatName, "implementer", false, pivot.Envelope.RelayID)
			if err != nil {
				t.Fatalf("realize pivot: %v", err)
			}
			if tc.undelivered {
				h16Commit(t, st, ceremonyRealizedUndelivered("retry-undelivered", pivot.Envelope.RelayID))
			}
			beforeRetryPivots := len(ceremonyRetryPivots(t, st, seatName))

			result := runCeremonyRetryProcess(t, root, seatName, "ambiguous delivery", nil, nil)
			if tc.wantAllowed {
				credential := ceremonyCredential(t, result)
				after, _ := seat.Open(root)
				if _, ok := after.Resolve(old.Value); ok {
					t.Fatal("allowed retry retained predecessor credential")
				}
				if _, ok := after.Resolve(credential); !ok {
					t.Fatal("allowed retry credential invalid")
				}
				if got := len(ceremonyRetryPivots(t, st, seatName)); got != beforeRetryPivots+1 {
					t.Fatalf("allowed retry pivots=%d, want starting %d + 1", got, beforeRetryPivots)
				}
				return
			}
			if result.err == nil || !strings.Contains(string(result.output), "delivery-retry-not-authorized") {
				t.Fatalf("refused retry err=%v output=%q", result.err, result.output)
			}
			if got := len(ceremonyRetryPivots(t, st, seatName)); got != beforeRetryPivots {
				t.Fatalf("refused retry pivots=%d, want unchanged %d", got, beforeRetryPivots)
			}
			after, _ := seat.Open(root)
			if _, ok := after.Resolve(old.Value); !ok {
				t.Fatal("refused retry invalidated healthy credential")
			}
		})
	}
}

func TestH16CeremonyRetryRejectsAuthorityFieldsAtIntake(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	root, st, seatName, current := setupCeremonyDeliveredAnchor(t)
	for _, args := range [][]string{{"-role", "planner"}, {"-operator"}} {
		result := runCeremonyRetryProcess(t, root, seatName, "authority override", args, nil)
		if result.err == nil || !strings.Contains(string(result.output), "recovery-authority-fields-forbidden") {
			t.Fatalf("args=%v err=%v output=%q", args, result.err, result.output)
		}
	}
	if got := len(ceremonyRetryPivots(t, st, seatName)); got != 0 {
		t.Fatalf("authority-field attempts committed %d pivots", got)
	}
	mgr, _ := seat.Open(root)
	if _, ok := mgr.Resolve(current); !ok {
		t.Fatal("authority-field refusal changed current credential")
	}
}

func TestH16CeremonyRetryCommitTimeAuthorityDeltaRejected(t *testing.T) {
	for _, tc := range []struct {
		name       string
		role       string
		isOperator bool
	}{
		{name: "role delta", role: "planner"},
		{name: "operator delta", role: "implementer", isOperator: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			st, err := store.Open(t.TempDir())
			if err != nil {
				t.Fatalf("open store: %v", err)
			}
			seatName := "retry-authority.implementer"
			h16Commit(t, st, h16LegacyPivot("retry-authority-tip", seatName))
			loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
				body, _ := json.Marshal(map[string]any{
					"seat": seatName, "role": tc.role, "is_operator": tc.isOperator, "retry_reason": "malicious override",
				})
				return record.Record{
					Envelope: record.Envelope{From: "system", Role: "system", DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
					Headers: map[string]string{
						"PHASE": "SITREP", "SUBJECT": tc.name, "record_kind": "seat_mint",
						"mint_predecessor": "retry-authority-tip", "admin_provenance": "ceremony",
					},
					Body: string(body),
				}, nil, nil
			}, engine.TestReady())
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			go loop.Run(ctx)
			out := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "retry-authority-delta"})
			h16AssertRejectedClass(t, out, "retry-authority-delta")
			rec, err := st.Read(out.RelayID)
			if err != nil {
				t.Fatalf("read rejection: %v", err)
			}
			if rec.Headers["failing_edge"] != "retry-authority-delta" || rec.Headers["admin_provenance"] != "ceremony" {
				t.Fatalf("rejection headers=%v", rec.Headers)
			}
			chains, err := engine.BuildMintChains(h16Records(t, st))
			if err != nil || chains[seatName].Tip.Envelope.RelayID != "retry-authority-tip" || chains[seatName].Conflicted {
				t.Fatalf("rejected delta changed chain=%+v err=%v", chains[seatName], err)
			}
		})
	}
}

func TestH16CeremonyRetryReplyCrashRotatesAgainWithoutCredentialReuse(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	for _, cut := range []string{"ceremony_retry_pre_reply", "ceremony_retry_partial_reply", "ceremony_retry_post_reply"} {
		t.Run(cut, func(t *testing.T) {
			root, st, seatName, initial := setupCeremonyDeliveredAnchor(t)
			crashed := runCeremonyRetryProcess(t, root, seatName, "unacknowledged delivery", nil, []string{"FRANK_TEST_CRASHPOINT=" + cut})
			if crashed.err == nil {
				t.Fatalf("%s did not crash: output=%q", cut, crashed.output)
			}
			if cut == "ceremony_retry_pre_reply" && strings.Contains(string(crashed.output), "credential=") {
				t.Fatalf("pre-reply cut emitted credential bytes: %q", crashed.output)
			}
			if cut == "ceremony_retry_partial_reply" && string(crashed.output) != "credential=" {
				t.Fatalf("partial cut output=%q, want credential= prefix", crashed.output)
			}
			if cut == "ceremony_retry_post_reply" && !strings.HasPrefix(string(crashed.output), "credential=") {
				t.Fatalf("post-reply cut output=%q, want full credential", crashed.output)
			}
			firstRetryCredential, _ := bindingRowForSeat(t, root, seatName)["credential"].(string)
			firstRetryRef, _ := bindingRowForSeat(t, root, seatName)["realized_mint_ref"].(string)
			if firstRetryCredential == "" || firstRetryRef == "" {
				t.Fatalf("%s missing durable first retry binding", cut)
			}
			latest := ceremonyCredential(t, runCeremonyRetryProcess(t, root, seatName, "retry after missing acknowledgement", nil, nil))
			rotations := ceremonyRetryPivots(t, st, seatName)
			if len(rotations) != 2 || rotations[0].Envelope.RelayID != firstRetryRef || rotations[1].Envelope.RelayID == firstRetryRef {
				t.Fatalf("%s rotations=%+v first_ref=%q", cut, rotations, firstRetryRef)
			}
			mgr, _ := seat.Open(root)
			for label, credential := range map[string]string{"initial": initial, "first retry": firstRetryCredential} {
				if _, ok := mgr.Resolve(credential); ok {
					t.Fatalf("%s %s credential remained valid", cut, label)
				}
			}
			if _, ok := mgr.Resolve(latest); !ok {
				t.Fatalf("%s latest credential invalid", cut)
			}
		})
	}
}

func setupCeremonyDeliveredAnchor(t *testing.T) (string, *store.Store, string, string) {
	t.Helper()
	root := t.TempDir()
	initFixtureStore(t, root)
	st, _ := store.Open(root)
	seatName := "retry-anchor.implementer"
	h16Commit(t, st, h16LegacyPivot("retry-a", seatName))
	h16Commit(t, st, h16LegacyPivot("retry-selected", seatName))
	removeCeremonyRedo(t, root)
	credential := ceremonyCredential(t, runCeremonyProcess(t, root, seatName, "retry-selected", "", nil))
	return root, st, seatName, credential
}

func runCeremonyRetryProcess(t *testing.T, root, seatName, reason string, extraArgs, extraEnv []string) ceremonyProcessResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	args := []string{"-root", root, "-recover-seat", seatName, "-retry-reason", reason}
	args = append(args, extraArgs...)
	return runCeremonyArgs(ctx, buildFrank(t, ctx), args, extraEnv)
}

func ceremonyRetryPivots(t *testing.T, st *store.Store, seatName string) []record.Record {
	t.Helper()
	byID := map[string]record.Record{}
	for _, rec := range h16Records(t, st) {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" || rec.Headers["admin_provenance"] != "ceremony" {
			continue
		}
		req, violation := engine.ParseSeatMintBody(rec.Body, rec.Envelope.From)
		if violation == nil && req.Seat == seatName {
			byID[rec.Envelope.RelayID] = rec
		}
	}
	if len(byID) == 0 {
		return nil
	}
	var current record.Record
	for _, rec := range byID {
		if _, predecessorIsRetry := byID[rec.Headers["mint_predecessor"]]; !predecessorIsRetry {
			current = rec
			break
		}
	}
	rotations := make([]record.Record, 0, len(byID))
	for current.Envelope.RelayID != "" {
		rotations = append(rotations, current)
		var next record.Record
		for _, candidate := range byID {
			if candidate.Headers["mint_predecessor"] == current.Envelope.RelayID {
				next = candidate
				break
			}
		}
		current = next
	}
	return rotations
}

func ceremonyRealizedUndelivered(relayID, sourceRelayID string) record.Record {
	body, _ := json.Marshal(map[string]any{"source_relay_id": sourceRelayID, "kind": "realized-undelivered"})
	return record.Record{
		Envelope: record.Envelope{RelayID: relayID, From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "derived-work-transition"},
		Body:     string(body),
	}
}
