package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16CeremonyConflictShapesRetainInBandOperatorPaths(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	for _, targetOperator := range []bool{false, true} {
		t.Run(fmt.Sprintf("target_operator=%v", targetOperator), func(t *testing.T) {
			root := t.TempDir()
			pinned := initFixtureStore(t, root)
			st, _ := store.Open(root)
			seatName := fmt.Sprintf("in-band-target-%v", targetOperator)
			h16Commit(t, st, ceremonyPivot("in-band-a", seatName, "implementer", targetOperator))
			h16Commit(t, st, ceremonyPivot("in-band-b", seatName, "operator", targetOperator))
			removeCeremonyRedo(t, root)

			operatorName := "recovery-operator"
			h16Commit(t, st, record.Record{
				Envelope: record.Envelope{RelayID: "operator-active", From: operatorName, Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
				Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "operator active"},
				Body:     "active",
			})
			mgr, _ := seat.Open(root)
			operatorCredential, err := mgr.Mint(operatorName, "operator", true)
			if err != nil {
				t.Fatalf("mint other operator: %v", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
			defer cancel()
			socket := filepath.Join(os.TempDir(), fmt.Sprintf("frank-ceremony-in-band-%d.sock", time.Now().UnixNano()))
			t.Cleanup(func() { _ = os.Remove(socket) })
			cmd, stderr := startFrank(t, ctx, buildFrank(t, ctx), root, socket)
			defer func() {
				if cmd.Process != nil {
					_ = cmd.Process.Kill()
				}
				_ = cmd.Wait()
			}()
			waitForSocket(t, socket)
			client, err := channel.DialAuthenticated(ctx, socket, operatorCredential.Value)
			if err != nil {
				t.Fatalf("dial other operator: %v stderr=%s", err, stderr.String())
			}
			defer client.Close()
			reg, err := fieldspec.Load(store.StoreRootConfigPaths(root)["fieldspec"])
			if err != nil {
				t.Fatalf("load registry: %v", err)
			}
			body, _ := json.Marshal(map[string]string{"seat": seatName, "selects": "in-band-b"})
			rec := record.Record{
				Envelope: record.Envelope{To: seatName},
				Headers: map[string]string{
					"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1",
					"SUBJECT": "in-band recovery anchor", "record_kind": "mint-chain-anchor",
				},
				Body: string(body),
			}
			payload := submitPayloadBytes(t, reg, pinned.Digest, seat.SeatMeta{Name: operatorName, Role: "operator", IsOperator: true}, rec)
			raw, err := client.Call(ctx, "submit", payload)
			if err != nil {
				t.Fatalf("submit anchor: %v stderr=%s", err, stderr.String())
			}
			var outcome struct {
				State string `json:"state"`
			}
			if err := json.Unmarshal(raw, &outcome); err != nil || outcome.State != record.Accepted {
				t.Fatalf("anchor outcome=%s decoded=%+v err=%v", raw, outcome, err)
			}
			stored, err := st.Records()
			if err != nil {
				t.Fatalf("read records: %v", err)
			}
			chains, err := engine.BuildMintChains(stored)
			if err != nil || chains[seatName].Conflicted || chains[seatName].Tip.Envelope.RelayID != "in-band-b" {
				t.Fatalf("in-band chain=%+v err=%v", chains[seatName], err)
			}
		})
	}
}

func TestH16CeremonyConflictSelectsOnceRealizesAndRefusesReplay(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	root := t.TempDir()
	initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	seatName := "ceremony-conflict.implementer"
	h16Commit(t, st, h16LegacyPivot("ceremony-a", seatName))
	h16Commit(t, st, h16LegacyPivot("ceremony-b", seatName))
	removeCeremonyRedo(t, root)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open bindings: %v", err)
	}
	old, err := mgr.MintOrReplace(seatName, "implementer", false, "ceremony-a")
	if err != nil {
		t.Fatalf("seed old binding: %v", err)
	}

	first := runCeremonyProcess(t, root, seatName, "ceremony-b", "", nil)
	credential := ceremonyCredential(t, first)
	anchors := ceremonyAnchorRecords(t, st, seatName)
	if len(anchors) != 1 || anchors[0].Envelope.DeliveryState != record.Accepted || anchors[0].Headers["admin_provenance"] != "ceremony" {
		t.Fatalf("anchors=%+v, want one accepted ceremony anchor", anchors)
	}
	var body struct {
		Selects string `json:"selects"`
	}
	if err := json.Unmarshal([]byte(anchors[0].Body), &body); err != nil || body.Selects != "ceremony-b" {
		t.Fatalf("anchor body=%q err=%v", anchors[0].Body, err)
	}
	after, err := seat.Open(root)
	if err != nil {
		t.Fatalf("reopen bindings: %v", err)
	}
	if realized, ok := after.RealizedMintRef(seatName); !ok || realized != "ceremony-b" {
		t.Fatalf("realized=%q ok=%v, want ceremony-b", realized, ok)
	}
	if _, ok := after.Resolve(old.Value); ok {
		t.Fatal("old credential authenticated after durable realization")
	}
	if meta, ok := after.Resolve(credential); !ok || meta.Name != seatName {
		t.Fatalf("returned credential meta=%+v ok=%v", meta, ok)
	}

	before := bindingRowForSeat(t, root, seatName)["credential"]
	replay := runCeremonyProcess(t, root, seatName, "", "", nil)
	if replay.err == nil || !strings.Contains(string(replay.output), "recovery-target-not-quarantined") {
		t.Fatalf("replay err=%v output=%q", replay.err, replay.output)
	}
	if got := ceremonyAnchorRecords(t, st, seatName); len(got) != 1 {
		t.Fatalf("replay anchors=%d, want 1", len(got))
	}
	if afterCredential := bindingRowForSeat(t, root, seatName)["credential"]; afterCredential != before {
		t.Fatalf("replay rotated credential: before=%v after=%v", before, afterCredential)
	}
}

func TestH16CeremonyStateActionMatrixResolvedBranches(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	t.Run("unique tip unrealized repairs without selector", func(t *testing.T) {
		root := t.TempDir()
		initFixtureStore(t, root)
		st, _ := store.Open(root)
		seatName := "ceremony-unrealized.implementer"
		h16Commit(t, st, h16LegacyPivot("unique-tip", seatName))
		result := runCeremonyProcess(t, root, seatName, "", "", nil)
		credential := ceremonyCredential(t, result)
		if anchors := ceremonyAnchorRecords(t, st, seatName); len(anchors) != 0 {
			t.Fatalf("unrealized unique tip authored anchors=%+v", anchors)
		}
		mgr, err := seat.Open(root)
		if err != nil {
			t.Fatalf("open bindings: %v", err)
		}
		if realized, ok := mgr.RealizedMintRef(seatName); !ok || realized != "unique-tip" {
			t.Fatalf("realized=%q ok=%v", realized, ok)
		}
		if _, ok := mgr.Resolve(credential); !ok {
			t.Fatal("returned credential did not authenticate")
		}
	})

	for _, realized := range []bool{false, true} {
		t.Run(fmt.Sprintf("selector against resolved realized=%v", realized), func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			st, _ := store.Open(root)
			seatName := fmt.Sprintf("ceremony-resolved-%v.implementer", realized)
			h16Commit(t, st, h16LegacyPivot("resolved-tip", seatName))
			if realized {
				mgr, _ := seat.Open(root)
				if _, err := mgr.MintOrReplace(seatName, "implementer", false, "resolved-tip"); err != nil {
					t.Fatalf("realize: %v", err)
				}
			}
			result := runCeremonyProcess(t, root, seatName, "resolved-tip", "", nil)
			if result.err == nil || !strings.Contains(string(result.output), "anchor-target-resolved") {
				t.Fatalf("resolved selector err=%v output=%q", result.err, result.output)
			}
			anchors := ceremonyAnchorRecords(t, st, seatName)
			if len(anchors) != 1 || anchors[0].Envelope.DeliveryState != record.Rejected || anchors[0].Headers["failing_edge"] != "anchor-target-resolved" {
				t.Fatalf("resolved anomaly=%+v", anchors)
			}
		})
	}

	t.Run("resolved realized without selector refuses", func(t *testing.T) {
		root := t.TempDir()
		initFixtureStore(t, root)
		st, _ := store.Open(root)
		seatName := "ceremony-healthy.implementer"
		h16Commit(t, st, h16LegacyPivot("healthy-tip", seatName))
		mgr, _ := seat.Open(root)
		credential, err := mgr.MintOrReplace(seatName, "implementer", false, "healthy-tip")
		if err != nil {
			t.Fatalf("realize: %v", err)
		}
		result := runCeremonyProcess(t, root, seatName, "", "", nil)
		if result.err == nil || !strings.Contains(string(result.output), "recovery-target-not-quarantined") {
			t.Fatalf("healthy recovery err=%v output=%q", result.err, result.output)
		}
		after, _ := seat.Open(root)
		if _, ok := after.Resolve(credential.Value); !ok {
			t.Fatal("healthy refusal changed current credential")
		}
		if anchors := ceremonyAnchorRecords(t, st, seatName); len(anchors) != 0 {
			t.Fatalf("healthy refusal authored anchors=%+v", anchors)
		}
	})
}

func TestH16CeremonyCustodyAuthorityIgnoresDisputedRoleBits(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	for _, bindingRef := range []string{"role-worker", "role-operator"} {
		t.Run(bindingRef, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			st, _ := store.Open(root)
			seatName := "ceremony-role-flip"
			h16Commit(t, st, ceremonyPivot("role-worker", seatName, "implementer", false))
			h16Commit(t, st, ceremonyPivot("role-operator", seatName, "orchestrator-planner", true))
			removeCeremonyRedo(t, root)
			mgr, _ := seat.Open(root)
			role, operator := "implementer", false
			if bindingRef == "role-operator" {
				role, operator = "orchestrator-planner", true
			}
			if _, err := mgr.MintOrReplace(seatName, role, operator, bindingRef); err != nil {
				t.Fatalf("seed binding: %v", err)
			}
			result := runCeremonyProcess(t, root, seatName, "role-operator", "", nil)
			credential := ceremonyCredential(t, result)
			after, _ := seat.Open(root)
			meta, ok := after.Resolve(credential)
			if !ok || meta.Role != "orchestrator-planner" || !meta.IsOperator {
				t.Fatalf("selected canonical authority not realized: meta=%+v ok=%v", meta, ok)
			}
		})
	}
}

func TestH16CeremonyLockFirstAliasAndSocketDiagnostic(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	root := t.TempDir()
	initFixtureStore(t, root)
	st, _ := store.Open(root)
	seatName := "ceremony-lock.implementer"
	h16Commit(t, st, h16LegacyPivot("lock-tip", seatName))
	bindingPath := filepath.Join(root, "binding", "seats.json")
	beforeBinding := readOptionalFile(t, bindingPath)
	beforeRecords := len(h16Records(t, st))

	socket := filepath.Join(os.TempDir(), fmt.Sprintf("frank-ceremony-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(socket) })
	listener, err := net.Listen("unix", socket)
	if err != nil {
		t.Fatalf("listen decoy: %v", err)
	}
	defer listener.Close()
	held, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("hold root: %v", err)
	}
	result := runCeremonyProcess(t, root, seatName, "", socket, nil)
	if result.err == nil || !strings.Contains(string(result.output), "root-lock-held") || strings.Contains(string(result.output), "socket appears live") {
		t.Fatalf("lock precedence err=%v output=%q", result.err, result.output)
	}
	if err := held.Release(); err != nil {
		t.Fatalf("release: %v", err)
	}
	if after := readOptionalFile(t, bindingPath); !bytes.Equal(after, beforeBinding) {
		t.Fatalf("held-root refusal changed binding\nbefore=%s\nafter=%s", beforeBinding, after)
	}
	if after := len(h16Records(t, st)); after != beforeRecords {
		t.Fatalf("held-root refusal changed records: before=%d after=%d", beforeRecords, after)
	}

	alias := filepath.Join(t.TempDir(), "alias")
	if err := os.Symlink(root, alias); err != nil {
		t.Fatalf("symlink alias: %v", err)
	}
	held, err = store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("hold root for alias: %v", err)
	}
	aliasResult := runCeremonyProcess(t, alias, seatName, "", "", nil)
	_ = held.Release()
	if aliasResult.err == nil || !strings.Contains(string(aliasResult.output), "root-lock-held") {
		t.Fatalf("alias err=%v output=%q", aliasResult.err, aliasResult.output)
	}

	liveDiagnostic := runCeremonyProcess(t, root, seatName, "", socket, nil)
	if liveDiagnostic.err != nil || !strings.Contains(string(liveDiagnostic.output), "warning: conductor socket appears live") {
		t.Fatalf("post-lock diagnostic err=%v output=%q", liveDiagnostic.err, liveDiagnostic.output)
	}
}

func TestH16CeremonyConcurrentProcessesHaveOneLockWinner(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	root := t.TempDir()
	initFixtureStore(t, root)
	st, _ := store.Open(root)
	seatName := "ceremony-race.implementer"
	h16Commit(t, st, h16LegacyPivot("race-a", seatName))
	h16Commit(t, st, h16LegacyPivot("race-b", seatName))
	removeCeremonyRedo(t, root)
	writeLargeBindingFixture(t, root, 40_000)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	start := make(chan struct{})
	results := make([]ceremonyProcessResult, 2)
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-start
			results[i] = runCeremonyBinary(ctx, bin, root, seatName, "race-b", "", nil)
		}(i)
	}
	close(start)
	wg.Wait()
	successes, locked := 0, 0
	for _, result := range results {
		if result.err == nil {
			successes++
		} else if strings.Contains(string(result.output), "root-lock-held") {
			locked++
		}
	}
	if successes != 1 || locked != 1 {
		t.Fatalf("race results=%+v, want one success and one root-lock-held", results)
	}
	if anchors := ceremonyAnchorRecords(t, st, seatName); len(anchors) != 1 {
		t.Fatalf("race anchors=%d, want 1", len(anchors))
	}
}

func TestH16CeremonyCrashCutsConvergeWithoutDuplicateAnchor(t *testing.T) {
	if s12SkipOuterRunOnly(t) {
		return
	}
	for _, crashpoint := range []string{"ceremony_post_anchor", "ceremony_post_binding", "ceremony_pre_reply"} {
		t.Run(crashpoint, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			st, _ := store.Open(root)
			seatName := "ceremony-crash.implementer"
			h16Commit(t, st, h16LegacyPivot("crash-a", seatName))
			h16Commit(t, st, h16LegacyPivot("crash-b", seatName))
			removeCeremonyRedo(t, root)
			crashed := runCeremonyProcess(t, root, seatName, "crash-b", "", []string{"FRANK_TEST_CRASHPOINT=" + crashpoint})
			if crashed.err == nil {
				t.Fatalf("%s did not crash: output=%q", crashpoint, crashed.output)
			}

			retry := runCeremonyProcess(t, root, seatName, "", "", nil)
			if crashpoint == "ceremony_post_anchor" {
				ceremonyCredential(t, retry)
			} else if retry.err == nil || !strings.Contains(string(retry.output), "recovery-target-not-quarantined") {
				t.Fatalf("%s retry err=%v output=%q", crashpoint, retry.err, retry.output)
			}
			if anchors := ceremonyAnchorRecords(t, st, seatName); len(anchors) != 1 {
				t.Fatalf("%s anchors=%d, want 1", crashpoint, len(anchors))
			}
			mgr, _ := seat.Open(root)
			if realized, ok := mgr.RealizedMintRef(seatName); !ok || realized != "crash-b" {
				t.Fatalf("%s realized=%q ok=%v", crashpoint, realized, ok)
			}
		})
	}
}

type ceremonyProcessResult struct {
	output []byte
	err    error
}

func runCeremonyProcess(t *testing.T, root, seatName, selects, socket string, extraEnv []string) ceremonyProcessResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return runCeremonyBinary(ctx, buildFrank(t, ctx), root, seatName, selects, socket, extraEnv)
}

func runCeremonyBinary(ctx context.Context, bin, root, seatName, selects, socket string, extraEnv []string) ceremonyProcessResult {
	args := []string{"-root", root, "-recover-seat", seatName}
	if selects != "" {
		args = append(args, "-select", selects)
	}
	if socket != "" {
		args = append(args, "-socket", socket)
	}
	return runCeremonyArgs(ctx, bin, args, extraEnv)
}

func runCeremonyArgs(ctx context.Context, bin string, args, extraEnv []string) ceremonyProcessResult {
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Env = append(os.Environ(), extraEnv...)
	output, err := cmd.CombinedOutput()
	return ceremonyProcessResult{output: output, err: err}
}

func ceremonyCredential(t *testing.T, result ceremonyProcessResult) string {
	t.Helper()
	if result.err != nil {
		t.Fatalf("ceremony err=%v output=%q", result.err, result.output)
	}
	for _, line := range strings.Split(string(result.output), "\n") {
		if value, ok := strings.CutPrefix(line, "credential="); ok && value != "" {
			return value
		}
	}
	t.Fatalf("ceremony output=%q, want credential", result.output)
	return ""
}

func ceremonyAnchorRecords(t *testing.T, st *store.Store, seatName string) []record.Record {
	t.Helper()
	var anchors []record.Record
	for _, rec := range h16Records(t, st) {
		if rec.Headers["record_kind"] != "mint-chain-anchor" {
			continue
		}
		var body struct {
			Seat string `json:"seat"`
		}
		if json.Unmarshal([]byte(rec.Body), &body) == nil && body.Seat == seatName {
			anchors = append(anchors, rec)
		}
	}
	return anchors
}

func ceremonyPivot(relayID, seatName, role string, isOperator bool) record.Record {
	body, _ := json.Marshal(map[string]any{"seat": seatName, "role": role, "is_operator": isOperator})
	return record.Record{
		Envelope: record.Envelope{RelayID: relayID, From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": relayID, "record_kind": "seat_mint"},
		Body:     string(body),
	}
}

func removeCeremonyRedo(t *testing.T, root string) {
	t.Helper()
	if err := os.RemoveAll(filepath.Join(root, "journal", "redo")); err != nil {
		t.Fatalf("remove redo: %v", err)
	}
}
