package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestAdmissionAtomicDisclosureAndWake(t *testing.T) {
	fixture := newFixture(t, "admit")
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `DELETE FROM turns WHERE turn_id=?`, fixture.turnID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='created' WHERE run_id=?`, fixture.runID)
		return err
	})
	if duplicate, err := fixture.scheduler.RecordWake(fixture.ctx, fixture.runID, "relay", 4); err != nil || duplicate {
		t.Fatalf("RecordWake = duplicate=%v err=%v", duplicate, err)
	}
	if duplicate, err := fixture.scheduler.RecordWake(fixture.ctx, fixture.runID, "relay", 5); err != nil || !duplicate {
		t.Fatalf("duplicate wake = duplicate=%v err=%v", duplicate, err)
	}
	relay := "relay"
	result, err := fixture.scheduler.Admit(fixture.ctx, AdmitRequest{
		RunID: fixture.runID, TurnID: "turn-new", GenerationID: fixture.generationID, TurnEpoch: "1",
		AdmissionRef:   appipc.AdmissionRef{Kind: appipc.AdmissionWakeRelay, RelayID: &relay},
		ConnectorReady: true, EncodedSize: 100, At: 6,
	})
	if err != nil || result.Decision != AdmissionCommitted || len(result.Body.ParkedUnknown) != 0 {
		t.Fatalf("Admit = %#v err=%v", result, err)
	}
	if result.Body.CreateAuthID != strings.Repeat("a", 32) || result.Body.RunDisposition != appipc.RunDispositionFresh || result.Body.SessionLogPath != "/initial/session.log" {
		t.Fatalf("fresh turn_open = %#v", result.Body)
	}
	fixture.assertScalar(`SELECT state FROM turns WHERE turn_id='turn-new'`, nil, "ACTIVE")
	fixture.assertScalar(`SELECT run_phase FROM runs WHERE run_id=?`, fixture.runID, "create_authorized")
	fixture.assertScalar(`SELECT create_auth_id FROM turns WHERE turn_id='turn-new'`, nil, strings.Repeat("a", 32))
	fixture.assertScalar(`SELECT disposition FROM wake_schedule WHERE relay_id='relay'`, nil, "ADMITTED")
	fixture.assertScalar(`SELECT state FROM leases WHERE run_id=? AND lease_kind='turn'`, fixture.runID, "ACTIVE")
	if fixture.count(`SELECT COUNT(*) FROM turn_disclosures WHERE disclosing_turn_id='turn-new'`, nil) != 0 {
		t.Fatal("disclosure snapshot not committed with admission")
	}
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE turns SET state='COMPLETED' WHERE turn_id='turn-new'`)
		return err
	})
	if _, err := fixture.scheduler.Admit(fixture.ctx, fixture.request("second-fresh", 100, false)); err == nil {
		t.Fatal("one-shot fresh authorization was consumed twice")
	}
}

func TestFreshAdmissionRejectsCreateAuthorizationCollision(t *testing.T) {
	fixture := newFixture(t, "nonce-collision")
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		rows := []struct {
			query string
			args  []any
		}{
			{`INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,session_log_path,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?,?)`, []any{"collision-run", []byte("{}"), strings.Repeat("0", 64), "ADMITTED", "created", "/collision/session.log", fmt.Sprintf("%020d", 0), 1}},
			{`INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, []any{"collision-run", fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0)}},
			{`INSERT INTO workers(generation_id,run_id,turn_epoch,pid,state,attach_result,created_at) VALUES(?,?,?,?,?,?,?)`, []any{"collision-worker", "collision-run", fmt.Sprintf("%020d", 1), 42, "LEASED", appipc.AttachOK, 1}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{"collision-run", "worker", "collision-lease", "collision-worker", fmt.Sprintf("%020d", 1), "ACTIVE", 1}},
		}
		for _, row := range rows {
			if _, err := tx.ExecContext(ctx, row.query, row.args...); err != nil {
				return err
			}
		}
		return nil
	})
	task := "task"
	_, err := fixture.scheduler.Admit(fixture.ctx, AdmitRequest{
		RunID: "collision-run", TurnID: "collision-turn", GenerationID: "collision-worker", TurnEpoch: "1",
		AdmissionRef: appipc.AdmissionRef{Kind: appipc.AdmissionOperatorInput, TaskInput: &task}, ConnectorReady: true, At: 2,
	})
	if err == nil || !strings.Contains(err.Error(), "mint collided") {
		t.Fatalf("collision admission err=%v", err)
	}
	fixture.assertScalar(`SELECT run_phase FROM runs WHERE run_id='collision-run'`, nil, "created")
	if fixture.count(`SELECT COUNT(*) FROM turns WHERE run_id='collision-run'`, nil) != 0 {
		t.Fatal("colliding create authorization mutated admission state")
	}
}

func TestContinuationInheritsAdmissionAndSnapshotsSettlement(t *testing.T) {
	fixture := newFixture(t, "continuation")
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE turns SET state='INTERRUPTED' WHERE turn_id=?`, fixture.turnID)
		return err
	})
	fixture.seedParked(fixture.turnID, "call", "ticket")
	request := fixture.request("turn-next", 100, true)
	request.PredecessorTurnID = fixture.turnID
	result, err := fixture.scheduler.Admit(fixture.ctx, request)
	if err != nil || result.Decision != AdmissionCommitted || result.Body.SettlementManifest == nil || len(result.Body.SettlementManifest.Entries) != 1 {
		t.Fatalf("continuation = %#v err=%v", result, err)
	}
	var inherited, successor, snapshot []byte
	fixture.read(func(ctx context.Context, s *store.Snapshot) error {
		if err := s.QueryRowContext(ctx, `SELECT admission_ref FROM turns WHERE turn_id=?`, fixture.turnID).Scan(&inherited); err != nil {
			return err
		}
		return s.QueryRowContext(ctx, `SELECT admission_ref,resume_snapshot FROM turns WHERE turn_id='turn-next'`).Scan(&successor, &snapshot)
	})
	if string(inherited) != string(successor) || len(snapshot) == 0 {
		t.Fatalf("inherited=%s successor=%s snapshot=%d", inherited, successor, len(snapshot))
	}
}

func TestWakeOriginContinuationCopiesNonceAndRecoveryReemitsCommittedSnapshot(t *testing.T) {
	var emitted []appipc.TurnOpenBody
	dropFirst := true
	fixture := newFixtureWithEmitter(t, "wake-continuation", applier.EmitterFunc(func(_ context.Context, emission applier.Emission) error {
		if emission.Kind != "turn_open" {
			return nil
		}
		emitted = append(emitted, emission.Value.(appipc.TurnOpenBody))
		if dropFirst {
			dropFirst = false
			return errors.New("injected post-commit drop")
		}
		return nil
	}))
	relay := "relay-wake"
	wakeRef, err := appipc.MarshalJCS(appipc.AdmissionRef{Kind: appipc.AdmissionWakeRelay, RelayID: &relay})
	if err != nil {
		t.Fatal(err)
	}
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `UPDATE turns SET state='INTERRUPTED',admission_ref=? WHERE turn_id=?`, wakeRef, fixture.turnID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO wake_schedule(relay_id,run_id,disposition,received_at,admitted_turn_id) VALUES(?,?, 'ADMITTED', ?, ?)`, relay, fixture.runID, 1, fixture.turnID)
		return err
	})
	request := fixture.request("turn-next", 100, true)
	request.PredecessorTurnID = fixture.turnID
	request.SessionLogPath = "/continuation/session.log"
	result, err := fixture.scheduler.Admit(fixture.ctx, request)
	if !errors.Is(err, applier.ErrEmission) || result.Decision != "" || len(emitted) != 1 {
		t.Fatalf("commit-then-drop = %#v err=%v emitted=%d", result, err, len(emitted))
	}
	first := emitted[0]
	if first.CreateAuthID != strings.Repeat("a", 32) || first.RunDisposition != appipc.RunDispositionResume {
		t.Fatalf("continuation identity/disposition = %#v", first)
	}

	var storedRef, snapshot, storedNonce []byte
	var resumeDisposition, wakeDisposition, admittedTurn string
	fixture.read(func(ctx context.Context, s *store.Snapshot) error {
		if err := s.QueryRowContext(ctx, `SELECT admission_ref,resume_snapshot,create_auth_id,resume_disposition FROM turns WHERE turn_id='turn-next'`).Scan(&storedRef, &snapshot, &storedNonce, &resumeDisposition); err != nil {
			return err
		}
		return s.QueryRowContext(ctx, `SELECT disposition,admitted_turn_id FROM wake_schedule WHERE relay_id=?`, relay).Scan(&wakeDisposition, &admittedTurn)
	})
	if string(storedRef) != string(wakeRef) || string(storedNonce) != strings.Repeat("a", 32) || resumeDisposition != "PENDING" {
		t.Fatalf("stored ref=%s nonce=%s disposition=%s", storedRef, storedNonce, resumeDisposition)
	}
	if wakeDisposition != "ADMITTED" || admittedTurn != fixture.turnID {
		t.Fatalf("wake mutated on continuation: disposition=%s admitted_turn_id=%s", wakeDisposition, admittedTurn)
	}
	var persisted map[string]json.RawMessage
	if err := json.Unmarshal(snapshot, &persisted); err != nil || len(persisted["session_log_path"]) == 0 || len(persisted["settlement_manifest"]) == 0 {
		t.Fatalf("resume snapshot=%s err=%v", snapshot, err)
	}

	replayed, err := fixture.scheduler.ReemitActive(fixture.ctx, fixture.runID)
	if err != nil || len(replayed) != 1 {
		t.Fatalf("ReemitActive = %#v err=%v", replayed, err)
	}
	if !reflect.DeepEqual(replayed[0], first) || len(emitted) != 2 || !reflect.DeepEqual(emitted[1], first) {
		t.Fatalf("replay changed committed payload\nfirst:  %#v\nreplay: %#v\nemitted=%#v", first, replayed[0], emitted)
	}
	fixture.assertScalar(`SELECT resume_disposition FROM turns WHERE turn_id='turn-next'`, nil, "PENDING")
	fixture.assertScalar(`SELECT disposition FROM wake_schedule WHERE relay_id=?`, relay, "ADMITTED")
}

func TestGenesisRecoveryRederivesDispositionAndCommittedAckAdvancesPhase(t *testing.T) {
	fixture := newFixture(t, "genesis-recovery")
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='create_authorized' WHERE run_id=?`, fixture.runID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `UPDATE turns SET state='ACTIVE',run_disposition='fresh' WHERE turn_id=?`, fixture.turnID)
		return err
	})
	replayed, err := fixture.scheduler.ReemitActive(fixture.ctx, fixture.runID)
	if err != nil || len(replayed) != 1 {
		t.Fatalf("ReemitActive = %#v err=%v", replayed, err)
	}
	if replayed[0].RunDisposition != appipc.RunDispositionResume || replayed[0].CreateAuthID != strings.Repeat("a", 32) {
		t.Fatalf("genesis replay = %#v", replayed[0])
	}
	decision, err := fixture.scheduler.RecordGenesisCommitted(fixture.ctx, GenesisCommittedRequest{
		RunID: fixture.runID, BoundRunID: fixture.runID, GenerationID: fixture.generationID, TurnEpoch: "1", At: 8,
	})
	if err != nil || decision != GenesisCommittedEstablished {
		t.Fatalf("RecordGenesisCommitted = %q err=%v", decision, err)
	}
	fixture.assertScalar(`SELECT run_phase FROM runs WHERE run_id=?`, fixture.runID, "established")
	replayed, err = fixture.scheduler.ReemitActive(fixture.ctx, fixture.runID)
	if err != nil || len(replayed) != 1 || replayed[0].RunDisposition != appipc.RunDispositionResume {
		t.Fatalf("established genesis replay = %#v err=%v", replayed, err)
	}
	decision, err = fixture.scheduler.RecordGenesisCommitted(fixture.ctx, GenesisCommittedRequest{
		RunID: fixture.runID, BoundRunID: fixture.runID, GenerationID: "stale", TurnEpoch: "0", At: 9,
	})
	if err != nil || decision != GenesisCommittedDuplicate {
		t.Fatalf("established duplicate = %q err=%v", decision, err)
	}
}

func TestGenesisCommittedOrderedTotalReceiver(t *testing.T) {
	request := func(f *fixture) GenesisCommittedRequest {
		return GenesisCommittedRequest{RunID: f.runID, BoundRunID: f.runID, GenerationID: f.generationID, TurnEpoch: "1", At: 8}
	}
	seq := func(f *fixture) string {
		var value string
		f.read(func(ctx context.Context, snapshot *store.Snapshot) error {
			return snapshot.QueryRowContext(ctx, `SELECT state_seq FROM epochs WHERE run_id=?`, f.runID).Scan(&value)
		})
		return value
	}

	malformed := newFixture(t, "ack-malformed")
	before := seq(malformed)
	bad := request(malformed)
	bad.GenerationID = ""
	if decision, err := malformed.scheduler.RecordGenesisCommitted(malformed.ctx, bad); err == nil || decision != "" || seq(malformed) != before {
		t.Fatalf("malformed = %q err=%v seq=%s", decision, err, seq(malformed))
	}

	identity := newFixture(t, "ack-identity")
	before = seq(identity)
	boundFault := request(identity)
	boundFault.BoundRunID = "other-run"
	if decision, err := identity.scheduler.RecordGenesisCommitted(identity.ctx, boundFault); err != nil || decision != GenesisCommittedIdentityFault || seq(identity) != before {
		t.Fatalf("identity = %q err=%v seq=%s", decision, err, seq(identity))
	}

	duplicate := newFixture(t, "ack-duplicate")
	before = seq(duplicate)
	late := request(duplicate)
	late.GenerationID = "superseded"
	late.TurnEpoch = "99"
	if decision, err := duplicate.scheduler.RecordGenesisCommitted(duplicate.ctx, late); err != nil || decision != GenesisCommittedDuplicate || seq(duplicate) != before {
		t.Fatalf("duplicate = %q err=%v seq=%s", decision, err, seq(duplicate))
	}

	created := newFixture(t, "ack-created")
	created.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='created' WHERE run_id=?`, created.runID)
		return err
	})
	before = seq(created)
	if decision, err := created.scheduler.RecordGenesisCommitted(created.ctx, request(created)); err != nil || decision != GenesisCommittedUnexpected || seq(created) != before {
		t.Fatalf("created = %q err=%v seq=%s", decision, err, seq(created))
	}

	stale := newFixture(t, "ack-stale")
	stale.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='create_authorized' WHERE run_id=?`, stale.runID)
		return err
	})
	before = seq(stale)
	staleRequest := request(stale)
	staleRequest.GenerationID = "superseded"
	if decision, err := stale.scheduler.RecordGenesisCommitted(stale.ctx, staleRequest); err != nil || decision != GenesisCommittedStale || seq(stale) != before {
		t.Fatalf("stale = %q err=%v seq=%s", decision, err, seq(stale))
	}

	current := newFixture(t, "ack-current")
	current.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='create_authorized' WHERE run_id=?`, current.runID)
		return err
	})
	before = seq(current)
	if decision, err := current.scheduler.RecordGenesisCommitted(current.ctx, request(current)); err != nil || decision != GenesisCommittedEstablished || seq(current) == before {
		t.Fatalf("current = %q err=%v before=%s after=%s", decision, err, before, seq(current))
	}
	current.assertScalar(`SELECT run_phase FROM runs WHERE run_id=?`, current.runID, "established")
}

func TestAdmissionSizingAndGates(t *testing.T) {
	initial := newFixture(t, "initial-overflow")
	initial.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `DELETE FROM turns WHERE turn_id=?`, initial.turnID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='created' WHERE run_id=?`, initial.runID)
		return err
	})
	result, err := initial.scheduler.Admit(initial.ctx, initial.request("turn", appipc.FrameMax+1, false))
	if err != nil || result.Decision != AdmissionTaskOverflow || initial.count(`SELECT COUNT(*) FROM turns WHERE run_id=?`, initial.runID) != 0 {
		t.Fatalf("initial overflow = %#v err=%v", result, err)
	}
	continuation := newFixture(t, "resume-overflow")
	continuation.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE turns SET state='INTERRUPTED' WHERE turn_id=?`, continuation.turnID)
		return err
	})
	request := continuation.request("turn", appipc.FrameMax+1, true)
	request.PredecessorTurnID = "old-turn"
	result, err = continuation.scheduler.Admit(continuation.ctx, request)
	if err != nil || result.Decision != AdmissionResumeTerminal {
		t.Fatalf("resume overflow = %#v err=%v", result, err)
	}
	continuation.assertScalar(`SELECT stop_reason FROM runs WHERE run_id=?`, continuation.runID, "resume_frame_overflow")

	refused := newFixture(t, "not-ready")
	request = refused.request("turn", 100, false)
	request.ConnectorReady = false
	if _, err := refused.scheduler.Admit(refused.ctx, request); err == nil || refused.count(`SELECT COUNT(*) FROM turns WHERE turn_id='turn'`, nil) != 0 {
		t.Fatal("not-ready admission did not fail closed")
	}
}

func TestAttemptCancellationTerminalAndLiveDisclosure(t *testing.T) {
	fixture := newFixture(t, "attempt")
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `UPDATE turns SET state='ACTIVE' WHERE turn_id=?`, fixture.turnID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,'ACTIVE',?)`, fixture.runID, "turn", "turn-lease", fixture.generationID, fmt.Sprintf("%020d", 1), 2)
		return err
	})
	fixture.seedParked(fixture.turnID, "call", "ticket")
	result, err := fixture.scheduler.OpenAttempt(fixture.ctx, AttemptRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "attempt", GenerationID: fixture.generationID, TurnEpoch: "1", LogicalSurfaceDigest: strings.Repeat("b", 64), At: 10})
	if err != nil || result.Decision != AttemptCommitted || len(result.ParkedUnknown) != 1 {
		t.Fatalf("OpenAttempt = %#v err=%v", result, err)
	}
	if duplicate, err := fixture.scheduler.RequestCancellation(fixture.ctx, CancellationRequest{ID: "cancel", RunID: fixture.runID, TargetKind: "attempt", TargetID: "attempt", Epoch: "1", At: 11}); err != nil || duplicate {
		t.Fatalf("cancel = duplicate=%v err=%v", duplicate, err)
	}
	if duplicate, err := fixture.scheduler.RequestCancellation(fixture.ctx, CancellationRequest{ID: "different", RunID: fixture.runID, TargetKind: "attempt", TargetID: "attempt", Epoch: "1", At: 12}); err != nil || !duplicate {
		t.Fatalf("cancel replay = duplicate=%v err=%v", duplicate, err)
	}
	if decision, err := fixture.scheduler.RecordTerminal(fixture.ctx, TerminalRequest{RunID: fixture.runID, TurnID: fixture.turnID, TurnEpoch: "1", Terminal: "turn_completed", At: 13}); err != nil || decision != TerminalRecorded {
		t.Fatalf("terminal = %q err=%v", decision, err)
	}
	if decision, err := fixture.scheduler.RecordTerminal(fixture.ctx, TerminalRequest{RunID: fixture.runID, TurnID: fixture.turnID, TurnEpoch: "1", Terminal: "turn_completed", At: 14}); err != nil || decision != TerminalDuplicate {
		t.Fatalf("terminal replay = %q err=%v", decision, err)
	}
	if decision, err := fixture.scheduler.RecordTerminal(fixture.ctx, TerminalRequest{RunID: fixture.runID, TurnID: fixture.turnID, TurnEpoch: "1", Terminal: "turn_failed", At: 15}); err != nil || decision != TerminalConflict {
		t.Fatalf("terminal conflict = %q err=%v", decision, err)
	}
}

type fixture struct {
	t                           *testing.T
	ctx                         context.Context
	store                       *store.Store
	host                        *applier.Host
	scheduler                   *Scheduler
	runID, turnID, generationID string
}

func newFixture(t *testing.T, suffix string) *fixture {
	return newFixtureWithEmitter(t, suffix, nil)
}

func newFixtureWithEmitter(t *testing.T, suffix string, emitter applier.Emitter) *fixture {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{Emitter: emitter})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	f := &fixture{t: t, ctx: ctx, store: db, host: host, scheduler: New(host, func() (string, error) { return strings.Repeat("a", 32), nil }), runID: "run-" + suffix, turnID: "old-turn", generationID: "generation-" + suffix}
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		rows := []struct {
			q    string
			args []any
		}{
			{`INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,session_log_path,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?,?)`, []any{f.runID, []byte("{}"), strings.Repeat("0", 64), "ACTIVE", "established", "/initial/session.log", fmt.Sprintf("%020d", 0), 1}},
			{`INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, []any{f.runID, fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0)}},
			{`INSERT INTO workers(generation_id,run_id,turn_epoch,pid,state,attach_result,created_at) VALUES(?,?,?,?,?,?,?)`, []any{f.generationID, f.runID, fmt.Sprintf("%020d", 1), 42, "LEASED", appipc.AttachOK, 1}},
			{`INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, []any{f.turnID, f.runID, fmt.Sprintf("%020d", 1), "COMPLETED", []byte(`{"kind":"operator_input","task_input":"task"}`), "fresh", strings.Repeat("a", 32), "RESUMABLE", 1}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{f.runID, "worker", "worker-lease", f.generationID, fmt.Sprintf("%020d", 1), "ACTIVE", 1}},
		}
		for _, row := range rows {
			if _, err := tx.ExecContext(ctx, row.q, row.args...); err != nil {
				return err
			}
		}
		return nil
	})
	return f
}

func (f *fixture) request(turn string, size int, continuation bool) AdmitRequest {
	task := "task"
	request := AdmitRequest{RunID: f.runID, TurnID: turn, GenerationID: f.generationID, TurnEpoch: "1", AdmissionRef: appipc.AdmissionRef{Kind: appipc.AdmissionOperatorInput, TaskInput: &task}, SessionLogPath: "/log", ConnectorReady: true, EncodedSize: size, At: 2}
	return request
}

func (f *fixture) seedParked(sourceTurn, call, ticket string) {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `INSERT INTO tool_calls(tool_call_id,run_id,turn_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,created_at) VALUES(?,?,?,?,?,?,?,?)`, call, f.runID, sourceTurn, fmt.Sprintf("%020d", 1), "UNKNOWN_TOOL_OUTCOME", "read", strings.Repeat("c", 64), 1)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, `INSERT INTO tool_authorizations(ticket_id,run_id,turn_id,tool_call_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,effect_descriptor,issued_at) VALUES(?,?,?,?,?,?,?,?,?,?)`, ticket, f.runID, sourceTurn, call, fmt.Sprintf("%020d", 1), "UNKNOWN_TOOL_OUTCOME", "read", strings.Repeat("c", 64), []byte("{}"), 1)
		return err
	})
}
func (f *fixture) mutate(fn func(context.Context, *store.Tx) error) {
	_, err := f.host.Apply(f.ctx, eventFunc{f.runID, fn})
	if err != nil {
		f.t.Fatal(err)
	}
}

type eventFunc struct {
	run string
	fn  func(context.Context, *store.Tx) error
}

func (e eventFunc) RunID() string { return e.run }
func (e eventFunc) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, e.fn(ctx, tx)
}
func (f *fixture) count(q string, arg any) int {
	n := 0
	f.read(func(ctx context.Context, s *store.Snapshot) error {
		if arg == nil {
			return s.QueryRowContext(ctx, q).Scan(&n)
		}
		return s.QueryRowContext(ctx, q, arg).Scan(&n)
	})
	return n
}
func (f *fixture) assertScalar(q string, arg any, want string) {
	f.t.Helper()
	got := ""
	f.read(func(ctx context.Context, s *store.Snapshot) error {
		if arg == nil {
			return s.QueryRowContext(ctx, q).Scan(&got)
		}
		return s.QueryRowContext(ctx, q, arg).Scan(&got)
	})
	if got != want {
		f.t.Fatalf("got %q want %q", got, want)
	}
}
func (f *fixture) read(fn func(context.Context, *store.Snapshot) error) {
	_, err := f.host.Read(f.ctx, applier.QueryFunc(func(ctx context.Context, s *store.Snapshot) (any, error) { return nil, fn(ctx, s) }))
	if err != nil {
		f.t.Fatal(err)
	}
}
