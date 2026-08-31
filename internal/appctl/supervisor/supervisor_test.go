package supervisor

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestRetirementTransactionParksAndMintsOnce(t *testing.T) {
	fixture := newSupervisorFixture(t, "ordinary")
	fixture.seedOpenRows()
	result, err := fixture.controller.Retire(fixture.ctx, RetireRequest{RunID: fixture.runID, TurnID: fixture.turnID, GenerationID: fixture.generationID, SuccessorGenerationID: "generation-next", At: 10, CountFailure: true})
	if err != nil || result.Branch != RetirementOrdinary || result.TurnEpoch != "2" || result.Parked != 1 {
		t.Fatalf("Retire = %#v err=%v", result, err)
	}
	fixture.assertScalar(`SELECT state FROM workers WHERE generation_id=?`, fixture.generationID, "FAILED")
	fixture.assertScalar(`SELECT state FROM workers WHERE generation_id=?`, "generation-next", "ALLOCATED")
	fixture.assertScalar(`SELECT state FROM turns WHERE turn_id=?`, fixture.turnID, "INTERRUPTED")
	fixture.assertScalar(`SELECT state FROM provider_attempts WHERE attempt_id='attempt-1'`, nil, "UNKNOWN_PROVIDER_OUTCOME")
	fixture.assertScalar(`SELECT state FROM tool_authorizations WHERE tool_call_id='call-issued'`, nil, "VOID")
	fixture.assertScalar(`SELECT state FROM tool_authorizations WHERE tool_call_id='call-consumed'`, nil, "UNKNOWN_TOOL_OUTCOME")
	fixture.assertScalar(`SELECT state FROM tool_calls WHERE tool_call_id='call-consumed'`, nil, "UNKNOWN_TOOL_OUTCOME")
	fixture.assertScalar(`SELECT turn_epoch FROM epochs WHERE run_id=?`, fixture.runID, storeCounter(2))

	replay, err := fixture.controller.Retire(fixture.ctx, RetireRequest{RunID: fixture.runID, TurnID: fixture.turnID, GenerationID: fixture.generationID, SuccessorGenerationID: "different", At: 11, CountFailure: true})
	if err != nil || !replay.Idempotent || fixture.count(`SELECT COUNT(*) FROM workers WHERE run_id=?`, fixture.runID) != 2 {
		t.Fatalf("retirement replay = %#v err=%v", replay, err)
	}
}

func TestRetirementTenthFailureAndParkedCapAreDistinctTerminals(t *testing.T) {
	tenth := newSupervisorFixture(t, "tenth")
	tenth.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET consecutive_failures=? WHERE run_id=?`, storeCounter(9), tenth.runID)
		return err
	})
	result, err := tenth.controller.Retire(tenth.ctx, RetireRequest{RunID: tenth.runID, TurnID: tenth.turnID, GenerationID: tenth.generationID, SuccessorGenerationID: "forbidden", At: 10, CountFailure: true})
	if err != nil || result.Branch != RetirementFailureCeiling || tenth.count(`SELECT COUNT(*) FROM workers WHERE generation_id='forbidden'`, nil) != 0 {
		t.Fatalf("tenth failure = %#v err=%v", result, err)
	}
	tenth.assertScalar(`SELECT state FROM runs WHERE run_id=?`, tenth.runID, "FAILED")

	capFixture := newSupervisorFixture(t, "cap")
	capFixture.seedParked(appipc.MaxParkedRowsPerRun)
	capFixture.seedConsumed("cap-crossing")
	result, err = capFixture.controller.Retire(capFixture.ctx, RetireRequest{RunID: capFixture.runID, TurnID: capFixture.turnID, GenerationID: capFixture.generationID, SuccessorGenerationID: "forbidden-cap", At: 10})
	if err != nil || result.Branch != RetirementParkedCap || result.Parked != appipc.MaxParkedRowsPerRun+1 {
		t.Fatalf("parked cap = %#v err=%v", result, err)
	}
	capFixture.assertScalar(`SELECT stop_reason FROM runs WHERE run_id=?`, capFixture.runID, "parked_unknown_capacity_exceeded")
	if capFixture.count(`SELECT COUNT(*) FROM workers WHERE generation_id='forbidden-cap'`, nil) != 0 || capFixture.count(`SELECT COUNT(*) FROM tool_calls WHERE run_id=? AND state='UNKNOWN_TOOL_OUTCOME'`, capFixture.runID) != appipc.MaxParkedRowsPerRun+1 {
		t.Fatal("cap terminal truncated rows or allocated a successor")
	}
}

func TestWorkerMachineWashoutAndSpawnBoundaries(t *testing.T) {
	machine := NewMachine(WorkerAllocated)
	for _, next := range []WorkerState{WorkerSpawning, WorkerReady, WorkerLeased, WorkerRetiring, WorkerTerminated} {
		if err := machine.Transition(next); err != nil {
			t.Fatalf("transition to %s: %v", next, err)
		}
	}
	if err := NewMachine(WorkerAllocated).Transition(WorkerLeased); err == nil {
		t.Fatal("illegal lifecycle transition accepted")
	}

	env := SanitizeEnv(map[string]string{"PATH": "/bin", "LANG": "C", "PROVIDER_SECRET": "must-not-pass"}, []string{"PATH", "LANG"})
	if len(env) != 2 || env[0] != "LANG=C" || env[1] != "PATH=/bin" {
		t.Fatalf("sanitized env = %v", env)
	}
	runtimeDir := filepath.Join(t.TempDir(), "run")
	if err := PrepareRuntimeDir(runtimeDir); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(runtimeDir)
	if err != nil || info.Mode().Perm() != 0o700 {
		t.Fatalf("runtime dir mode = %v err=%v", info.Mode().Perm(), err)
	}

	pair, err := NewSocketPair()
	if err != nil {
		t.Fatal(err)
	}
	defer pair.Parent.Close()
	if err := pair.Child.Close(); err != nil {
		t.Fatal(err)
	}
	buffer := make([]byte, 1)
	if _, err := pair.Parent.Read(buffer); err != io.EOF {
		t.Fatalf("channel close = %v, want EOF", err)
	}

	death, err := NewDeathPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer death.Child.Close()
	if err := death.Parent.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := death.Child.Read(buffer); err != io.EOF {
		t.Fatalf("death pipe close = %v, want EOF", err)
	}
	command, err := BuildCommand(ProcessSpec{Path: "/bin/worker", Args: []string{"--serve"}, Dir: runtimeDir, Env: env})
	if err != nil || command.Path != "/bin/worker" || command.Dir != runtimeDir || len(command.Env) != 2 {
		t.Fatalf("BuildCommand = %#v err=%v", command, err)
	}
}

func TestWashoutResetCancelHealthAndTermination(t *testing.T) {
	fixture := newSupervisorFixture(t, "washout")
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE workers SET state='SPAWNING' WHERE generation_id=?`, fixture.generationID)
		return err
	})
	result, err := fixture.controller.Washout(fixture.ctx, WashoutRequest{
		RunID: fixture.runID, GenerationID: fixture.generationID,
		SuccessorGenerationID: "washout-next", At: 20,
	})
	if err != nil || result.TurnEpoch != "1" || result.Failures != 1 {
		t.Fatalf("Washout = %#v err=%v", result, err)
	}
	fixture.assertScalar(`SELECT turn_epoch FROM workers WHERE generation_id='washout-next'`, nil, storeCounter(1))
	if err := fixture.controller.ResetFailures(fixture.ctx, fixture.runID, 21); err != nil {
		t.Fatal(err)
	}
	fixture.assertScalar(`SELECT consecutive_failures FROM runs WHERE run_id=?`, fixture.runID, storeCounter(0))
	if err := fixture.controller.CancelRun(fixture.ctx, fixture.runID, 22); err != nil {
		t.Fatal(err)
	}
	fixture.assertScalar(`SELECT state FROM runs WHERE run_id=?`, fixture.runID, "CANCELLED")

	health := NewHealth(time.Second, 2)
	health.Pong(time.Unix(10, 0))
	if health.Missed(time.Unix(11, 0)) || !health.Missed(time.Unix(13, 0)) {
		t.Fatal("health deadline did not distinguish live from missed")
	}
	if got := ValidateAttachResult("attach-ok"); got != nil || ValidateAttachResult("unknown") == nil {
		t.Fatalf("attach closed set: valid=%v unknown=%v", got, ValidateAttachResult("unknown"))
	}
	if relation, err := ClassifyEpoch("1", "2"); err != nil || relation != EpochStale {
		t.Fatalf("ClassifyEpoch = %q err=%v", relation, err)
	}

	child := &fakeChild{}
	if err := Terminate(context.Background(), child, time.Millisecond); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(child.steps) != "[shutdown terminated]" {
		t.Fatalf("graceful ladder = %v", child.steps)
	}
	child = &fakeChild{waitErrors: 2}
	if err := Terminate(context.Background(), child, time.Millisecond); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(child.steps) != "[shutdown wait signal:terminated wait signal:killed terminated]" {
		t.Fatalf("forced ladder = %v", child.steps)
	}
}

type fakeChild struct {
	steps      []string
	waitErrors int
}

func (child *fakeChild) Shutdown(context.Context) error {
	child.steps = append(child.steps, "shutdown")
	return nil
}
func (child *fakeChild) Signal(signal os.Signal) error {
	child.steps = append(child.steps, "signal:"+signal.String())
	return nil
}
func (child *fakeChild) Wait(context.Context) error {
	if child.waitErrors > 0 {
		child.waitErrors--
		child.steps = append(child.steps, "wait")
		return context.DeadlineExceeded
	}
	child.steps = append(child.steps, "terminated")
	return nil
}

type supervisorFixture struct {
	t                           *testing.T
	ctx                         context.Context
	store                       *store.Store
	applier                     *applier.Host
	controller                  *Controller
	runID, turnID, generationID string
}

func newSupervisorFixture(t *testing.T, suffix string) *supervisorFixture {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	f := &supervisorFixture{t: t, ctx: ctx, store: db, applier: host, runID: "run-" + suffix, turnID: "turn-" + suffix, generationID: "generation-" + suffix}
	f.mutateWithRun(f.runID, func(ctx context.Context, tx *store.Tx) error {
		stmts := []struct {
			q string
			a []any
		}{
			{`INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, []any{f.runID, []byte("{}"), fmt.Sprintf("%064d", 0), "ACTIVE", "established", storeCounter(0), 1}},
			{`INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, []any{f.runID, storeCounter(1), storeCounter(0)}},
			{`INSERT INTO workers(generation_id,run_id,turn_epoch,pid,state,created_at) VALUES(?,?,?,?,?,?)`, []any{f.generationID, f.runID, storeCounter(1), 42, "LEASED", 1}},
			{`INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, []any{f.turnID, f.runID, storeCounter(1), "ACTIVE", []byte("task"), "fresh", strings.Repeat("a", 32), "PENDING", 1}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{f.runID, "worker", "wl", f.generationID, storeCounter(1), "ACTIVE", 1}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{f.runID, "turn", "tl", f.generationID, storeCounter(1), "ACTIVE", 1}},
		}
		for _, s := range stmts {
			if _, err := tx.ExecContext(ctx, s.q, s.a...); err != nil {
				return err
			}
		}
		return nil
	})
	f.controller = New(host)
	return f
}

func (f *supervisorFixture) seedOpenRows() {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		if err := insertAuth(ctx, tx, f, "call-issued", "ticket-issued", "ISSUED", nil); err != nil {
			return err
		}
		if err := insertAuth(ctx, tx, f, "call-consumed", "ticket-consumed", "CONSUMED", nil); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO provider_attempts(attempt_id,run_id,turn_id,turn_epoch,state,logical_surface_digest,created_at) VALUES(?,?,?,?,?,?,?)`, "attempt-1", f.runID, f.turnID, storeCounter(1), "OPEN", fmt.Sprintf("%064d", 1), 1)
		return err
	})
}

func (f *supervisorFixture) seedConsumed(id string) {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		return insertAuth(ctx, tx, f, id, "ticket-"+id, "CONSUMED", nil)
	})
}
func (f *supervisorFixture) seedParked(n int) {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		for i := 0; i < n; i++ {
			id := fmt.Sprintf("park-%04d", i)
			d := fmt.Sprintf("%064x", i+1)
			if _, err := tx.ExecContext(ctx, `INSERT INTO tool_calls(tool_call_id,run_id,turn_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,created_at) VALUES(?,?,?,?,?,?,?,?)`, id, f.runID, f.turnID, storeCounter(1), "UNKNOWN_TOOL_OUTCOME", "read", d, 1); err != nil {
				return err
			}
		}
		return nil
	})
}

func insertAuth(ctx context.Context, tx *store.Tx, f *supervisorFixture, call, ticket, state string, reason any) error {
	_, err := tx.ExecContext(ctx, `INSERT INTO tool_authorizations(ticket_id,run_id,turn_id,tool_call_id,turn_epoch,state,void_reason,canonical_tool_name,canonical_args_digest,effect_descriptor,issued_at,consumed_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`, ticket, f.runID, f.turnID, call, storeCounter(1), state, reason, "read", fmt.Sprintf("%064x", len(call)), []byte("{}"), 1, func() any {
		if state == "CONSUMED" {
			return 2
		}
		return nil
	}())
	return err
}

func (f *supervisorFixture) mutate(fn func(context.Context, *store.Tx) error) {
	f.mutateWithRun(f.runID, fn)
}
func (f *supervisorFixture) mutateWithRun(runID string, fn func(context.Context, *store.Tx) error) {
	_, err := f.applier.Apply(f.ctx, testEvent{runID, fn})
	if err != nil {
		f.t.Fatal(err)
	}
}

type testEvent struct {
	runID string
	fn    func(context.Context, *store.Tx) error
}

func (e testEvent) RunID() string { return e.runID }
func (e testEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, e.fn(ctx, tx)
}
func (f *supervisorFixture) assertScalar(q string, arg any, want string) {
	f.t.Helper()
	var got string
	f.read(func(ctx context.Context, s *store.Snapshot) error {
		var row interface{ Scan(...any) error }
		_ = row
		if arg == nil {
			return s.QueryRowContext(ctx, q).Scan(&got)
		}
		return s.QueryRowContext(ctx, q, arg).Scan(&got)
	})
	if got != want {
		f.t.Fatalf("scalar=%q want=%q query=%s", got, want, q)
	}
}
func (f *supervisorFixture) count(q string, arg any) int {
	f.t.Helper()
	n := 0
	f.read(func(ctx context.Context, s *store.Snapshot) error {
		if arg == nil {
			return s.QueryRowContext(ctx, q).Scan(&n)
		}
		return s.QueryRowContext(ctx, q, arg).Scan(&n)
	})
	return n
}
func (f *supervisorFixture) read(fn func(context.Context, *store.Snapshot) error) {
	_, err := f.applier.Read(f.ctx, applier.QueryFunc(func(ctx context.Context, s *store.Snapshot) (any, error) { return nil, fn(ctx, s) }))
	if err != nil {
		f.t.Fatal(err)
	}
}
func storeCounter(v uint64) string { return fmt.Sprintf("%020d", v) }
