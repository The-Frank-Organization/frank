package recovery

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/brokerclient"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestRecoveryMatrixAlwaysReproposesDurableTuple(t *testing.T) {
	tests := []struct {
		name       string
		seed       func(*fixture)
		wantCase   Case
		wantEpoch  string
		wantFailed string
	}{
		{name: "a leased generation without retirement", seed: func(f *fixture) { f.seedWorker("old", "1", "LEASED", true) }, wantCase: CaseLeased, wantEpoch: "2", wantFailed: "old"},
		{name: "b committed retirement without successor", seed: func(f *fixture) { f.seedWorker("old", "1", "FAILED", false); f.setEpoch("2") }, wantCase: CaseRetired, wantEpoch: "2"},
		{name: "c prelease candidate", seed: func(f *fixture) { f.seedWorker("candidate", "1", "READY", false) }, wantCase: CasePreLease, wantEpoch: "1", wantFailed: "candidate"},
		{name: "d initial run", seed: func(*fixture) {}, wantCase: CaseInitial, wantEpoch: "1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fixture := newFixture(t)
			test.seed(fixture)
			proposer := &scriptedProposer{actions: []brokerclient.FoldAction{brokerclient.Repropose, brokerclient.OpenAssign}}
			ids := &idSource{}
			engine := New(fixture.host, proposer, ids.Next, func() int64 { return 50 })
			outcomes, err := engine.Run(fixture.ctx)
			if err != nil || len(outcomes) != 1 {
				t.Fatalf("Run outcomes=%#v err=%v", outcomes, err)
			}
			outcome := outcomes[0]
			if outcome.Case != test.wantCase || outcome.ProposalAttempts != 2 || !outcome.AssignOpen {
				t.Fatalf("outcome=%#v", outcome)
			}
			if outcome.Tuple.TurnEpoch != test.wantEpoch || outcome.Tuple.StateSeq == "" || outcome.Tuple.GenerationID == "" {
				t.Fatalf("tuple=%#v", outcome.Tuple)
			}
			if len(proposer.tuples) != 2 || proposer.tuples[0] != proposer.tuples[1] || proposer.tuples[0] != outcome.Tuple {
				t.Fatalf("re-proposals=%#v outcome=%#v", proposer.tuples, outcome.Tuple)
			}
			if test.wantFailed != "" {
				fixture.assertWorkerState(test.wantFailed, "FAILED")
			}
		})
	}
}

func TestTerminalRunsDoNotRestartOrPropose(t *testing.T) {
	fixture := newFixture(t)
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED' WHERE run_id=?`, fixture.runID)
		return err
	})
	proposer := &scriptedProposer{}
	outcomes, err := New(fixture.host, proposer, (&idSource{}).Next, func() int64 { return 50 }).Run(fixture.ctx)
	if err != nil || len(outcomes) != 0 || len(proposer.tuples) != 0 {
		t.Fatalf("outcomes=%#v proposals=%#v err=%v", outcomes, proposer.tuples, err)
	}
}

func TestProposalRefusalFailsClosed(t *testing.T) {
	fixture := newFixture(t)
	proposer := &scriptedProposer{actions: []brokerclient.FoldAction{brokerclient.InvariantFault}}
	_, err := New(fixture.host, proposer, (&idSource{}).Next, func() int64 { return 50 }).Run(fixture.ctx)
	if err == nil || !strings.Contains(err.Error(), "proposal") {
		t.Fatalf("error=%v", err)
	}
}

type scriptedProposer struct {
	actions []brokerclient.FoldAction
	tuples  []appipc.EpochStateBody
}

func (proposer *scriptedProposer) Propose(_ context.Context, _ string, tuple appipc.EpochStateBody) (brokerclient.FoldResult, error) {
	proposer.tuples = append(proposer.tuples, tuple)
	if len(proposer.actions) == 0 {
		return brokerclient.FoldResult{}, fmt.Errorf("no scripted proposal")
	}
	action := proposer.actions[0]
	proposer.actions = proposer.actions[1:]
	return brokerclient.FoldResult{Action: action, Loud: action == brokerclient.InvariantFault}, nil
}

type idSource struct{ next int }

func (source *idSource) Next() string {
	source.next++
	return fmt.Sprintf("generation-%d", source.next)
}

type fixture struct {
	t     *testing.T
	ctx   context.Context
	db    *store.Store
	host  *applier.Host
	runID string
}

func newFixture(t *testing.T) *fixture {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	fixture := &fixture{t: t, ctx: ctx, db: db, host: host, runID: "run"}
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, fixture.runID, []byte("{}"), strings.Repeat("0", 64), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, fixture.runID, fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0))
		return err
	})
	return fixture
}

func (fixture *fixture) seedWorker(generationID, epoch, state string, leased bool) {
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `INSERT INTO workers(generation_id,run_id,turn_epoch,state,created_at) VALUES(?,?,?,?,?)`, generationID, fixture.runID, pad(epoch), state, 2); err != nil {
			return err
		}
		if leased {
			_, err := tx.ExecContext(ctx, `INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, fixture.runID, "worker", "lease", generationID, pad(epoch), "ACTIVE", 2)
			return err
		}
		return nil
	})
}

func (fixture *fixture) setEpoch(epoch string) {
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE epochs SET turn_epoch=? WHERE run_id=?`, pad(epoch), fixture.runID)
		return err
	})
}

func (fixture *fixture) assertWorkerState(generationID, want string) {
	fixture.t.Helper()
	var got string
	fixture.read(func(ctx context.Context, snapshot *store.Snapshot) error {
		return snapshot.QueryRowContext(ctx, `SELECT state FROM workers WHERE generation_id=?`, generationID).Scan(&got)
	})
	if got != want {
		fixture.t.Fatalf("worker %s state=%s want=%s", generationID, got, want)
	}
}

func (fixture *fixture) mutate(fn func(context.Context, *store.Tx) error) {
	_, err := fixture.host.Apply(fixture.ctx, recoveryEvent{runID: fixture.runID, fn: fn})
	if err != nil {
		fixture.t.Fatal(err)
	}
}

func (fixture *fixture) read(fn func(context.Context, *store.Snapshot) error) {
	_, err := fixture.host.Read(fixture.ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) { return nil, fn(ctx, snapshot) }))
	if err != nil {
		fixture.t.Fatal(err)
	}
}

type recoveryEvent struct {
	runID string
	fn    func(context.Context, *store.Tx) error
}

func (event recoveryEvent) RunID() string { return event.runID }
func (event recoveryEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, event.fn(ctx, tx)
}

func pad(value string) string {
	padded, _ := appipc.PadCounter(value)
	return padded
}
