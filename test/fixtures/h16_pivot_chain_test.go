package fixtures_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/derived"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16MintPredecessorStampedAndMismatchRejectedAtCommit(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		headers := map[string]string{"PHASE": "SITREP", "SUBJECT": cmd.IntakeID, "record_kind": "seat_mint"}
		if cmd.IntakeID == "mint-mismatch" {
			headers["mint_predecessor"] = "not-the-tip"
		}
		return record.Record{
			Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers:  headers,
			Body:     `{"seat":"chain-seat.implementer","role":"implementer","is_operator":false}`,
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	first := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "mint-first", Seat: "operator", Role: "operator", IsOperator: true})
	firstRec, err := st.Read(first.RelayID)
	if err != nil {
		t.Fatalf("read first pivot: %v", err)
	}
	if firstRec.Envelope.DeliveryState != record.Accepted || firstRec.Headers["mint_predecessor"] != "genesis" {
		t.Fatalf("first pivot=%+v, want accepted genesis predecessor", firstRec)
	}

	mismatch := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "mint-mismatch", Seat: "operator", Role: "operator", IsOperator: true})
	h16AssertRejectedClass(t, mismatch, "mint_predecessor:mint-predecessor-mismatch")
	mismatchRec, err := st.Read(mismatch.RelayID)
	if err != nil {
		t.Fatalf("read mismatch: %v", err)
	}
	if mismatchRec.Headers["failing_edge"] != "mint-predecessor-mismatch" || mismatchRec.Headers["mint_predecessor"] != "not-the-tip" {
		t.Fatalf("mismatch headers=%v", mismatchRec.Headers)
	}
}

func TestH16CanonicalTipAndCallerlessRepairRespectUnresolvedMarker(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	seatName := "h16-chain.implementer"
	pivotA := h16PivotRecord("z-pivot-a", seatName, "genesis")
	pivotB := h16PivotRecord("a-pivot-b", seatName, pivotA.Envelope.RelayID)
	h16Commit(t, st, pivotA)
	h16Commit(t, st, pivotB)
	marker := derived.AttemptRecord(pivotB.Envelope.RelayID, "mint", "none")
	marker.Envelope.RelayID = "m-unresolved-marker"
	h16Commit(t, st, marker)

	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open seats: %v", err)
	}
	old, err := mgr.MintOrReplace(seatName, "implementer", false, pivotA.Envelope.RelayID)
	if err != nil {
		t.Fatalf("realize A: %v", err)
	}
	if err := os.RemoveAll(filepath.Join(root, "redo")); err != nil {
		t.Fatalf("remove redo: %v", err)
	}
	if err := os.RemoveAll(filepath.Join(root, "projections")); err != nil {
		t.Fatalf("remove projections: %v", err)
	}

	bin := buildFrank(t, context.Background())
	h16StartAndStopFrank(t, bin, root, "blocked")
	afterBlocked, err := seat.Open(root)
	if err != nil {
		t.Fatalf("reopen seats after blocked startup: %v", err)
	}
	if realized, ok := afterBlocked.RealizedMintRef(seatName); !ok || realized != pivotA.Envelope.RelayID {
		t.Fatalf("unresolved marker did not block repair: realized=%q ok=%v", realized, ok)
	}
	if _, ok := afterBlocked.Resolve(old.Value); !ok {
		t.Fatal("A credential invalidated while B marker unresolved")
	}

	resolution := h16DerivedRecord("n-marker-resolution", "attempt_resolution", map[string]any{
		"resolves": marker.Envelope.RelayID, "disposition": "effect-confirmed-unrealized", "evidence_ref": "E-restart",
	})
	resolution.Envelope.From = "operator"
	resolution.Envelope.Role = "operator"
	h16Commit(t, st, resolution)
	h16StartAndStopFrank(t, bin, root, "repair")

	afterRepair, err := seat.Open(root)
	if err != nil {
		t.Fatalf("reopen seats after repair: %v", err)
	}
	if realized, ok := afterRepair.RealizedMintRef(seatName); !ok || realized != pivotB.Envelope.RelayID {
		t.Fatalf("canonical B not realized: realized=%q ok=%v", realized, ok)
	}
	if _, ok := afterRepair.Resolve(old.Value); ok {
		t.Fatal("A credential survived B realization")
	}
	credentialAfterRepair, _ := bindingRowForSeat(t, root, seatName)["credential"].(string)
	h16StartAndStopFrank(t, bin, root, "matching-ref")
	credentialAfterMatchingRestart, _ := bindingRowForSeat(t, root, seatName)["credential"].(string)
	if credentialAfterMatchingRestart == "" || credentialAfterMatchingRestart != credentialAfterRepair {
		t.Fatalf("matching realized ref rotated on restart: before=%q after=%q", credentialAfterRepair, credentialAfterMatchingRestart)
	}
	if pivotB.Envelope.RelayID >= pivotA.Envelope.RelayID {
		t.Fatalf("fixture lost adversarial relay order: B=%q A=%q", pivotB.Envelope.RelayID, pivotA.Envelope.RelayID)
	}
}

func TestH16RecoveryProcessorRescansAndRepairsMarkerlessPivot(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	reg := loadH16Registry(t)
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	rec := h16PresenceCandidate()
	rec.Headers["SUBJECT"] = "recovery processor mint"
	rec.Headers["record_kind"] = "seat_mint"
	rec.Body = `{"seat":"h16-recovery.implementer","role":"implementer","is_operator":false}`
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load pinned config: %v", err)
	}
	renderEnv := fieldspec.RenderEnv{ConfigDigest: pinned.Digest, PresentLayers: config.PresentLayers(pinned)}
	_, formDigest := reg.Render(renderEnv, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, "SITREP", "medium", fieldspec.ClosedGrantState)
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("open journal: %v", err)
	}
	intakeID, err := journal.Append(intake.Cmd{
		Seat: "operator", Role: "operator", IsOperator: true, Verb: "submit",
		Payload: mustJSONBytes(t, fieldspec.SubmitPayload{Record: rec, FormDigest: formDigest}),
	})
	if err != nil {
		t.Fatalf("append recovery command: %v", err)
	}
	bin := buildFrank(t, context.Background())
	h16StartAndStopFrank(t, bin, root, "recovery-processor")
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open recovered store: %v", err)
	}
	var pivot record.Record
	for _, stored := range h16Records(t, st) {
		if stored.Envelope.IntakeID == intakeID {
			pivot = stored
			break
		}
	}
	if pivot.Envelope.DeliveryState != record.Accepted || pivot.Headers["mint_predecessor"] != "genesis" {
		t.Fatalf("recovery pivot=%+v", pivot)
	}
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open recovered binding: %v", err)
	}
	if realized, ok := mgr.RealizedMintRef("h16-recovery.implementer"); !ok || realized != pivot.Envelope.RelayID {
		t.Fatalf("recovery processor did not repair markerless pivot: realized=%q ok=%v pivot=%q", realized, ok, pivot.Envelope.RelayID)
	}
}

func TestH16MintChainConflictShapesFailClosedOrderFree(t *testing.T) {
	seatName := "h16-conflict.implementer"
	for _, tc := range []struct {
		name    string
		records []record.Record
	}{
		{name: "shared predecessor", records: []record.Record{
			h16PivotRecord("root", seatName, "genesis"),
			h16PivotRecord("left", seatName, "root"),
			h16PivotRecord("right", seatName, "root"),
		}},
		{name: "broken link", records: []record.Record{
			h16PivotRecord("broken", seatName, "missing"),
		}},
		{name: "cycle", records: []record.Record{
			h16PivotRecord("cycle-a", seatName, "cycle-b"),
			h16PivotRecord("cycle-b", seatName, "cycle-a"),
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			forward, err := engine.BuildMintChains(tc.records)
			if err != nil || !forward[seatName].Conflicted || forward[seatName].Tip.Envelope.RelayID != "" {
				t.Fatalf("forward chain=%+v err=%v", forward[seatName], err)
			}
			reverse := append([]record.Record(nil), tc.records...)
			for left, right := 0, len(reverse)-1; left < right; left, right = left+1, right-1 {
				reverse[left], reverse[right] = reverse[right], reverse[left]
			}
			backward, err := engine.BuildMintChains(reverse)
			if err != nil || backward[seatName].Conflicted != forward[seatName].Conflicted || backward[seatName].Tip.Envelope.RelayID != "" {
				t.Fatalf("reverse chain=%+v err=%v", backward[seatName], err)
			}
		})
	}
}

func h16PivotRecord(relayID, seatName, predecessor string) record.Record {
	rec := record.Record{
		Envelope: record.Envelope{RelayID: relayID, From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE": "SITREP", "SUBJECT": relayID, "record_kind": "seat_mint", "mint_predecessor": predecessor,
		},
		Body: `{"seat":"` + seatName + `","role":"implementer","is_operator":false}`,
	}
	derived.Stamp(&rec)
	return rec
}

func h16StartAndStopFrank(t *testing.T, bin, root, suffix string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	sock := filepath.Join(os.TempDir(), "frank-h16-chain-"+suffix+"-"+filepath.Base(root)+".sock")
	_ = os.Remove(sock)
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	waitForSocket(t, sock)
	cancel()
	if cmd.Process != nil {
		_ = cmd.Process.Kill()
	}
	_ = cmd.Wait()
	_ = os.Remove(sock)
	if text := stderr.String(); strings.Contains(text, "panic") {
		t.Fatalf("frank startup panic: %s", text)
	}
}
