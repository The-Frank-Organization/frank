// Package recovery reconstructs the app-control lifecycle exclusively from
// committed state and republishes every current durable broker tuple.
package recovery

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/brokerclient"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appctl/supervisor"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type Case string

const (
	CaseLeased   Case = "leased_without_retirement"
	CaseRetired  Case = "retirement_committed"
	CasePreLease Case = "prelease_candidate"
	CaseInitial  Case = "initial_run"
)

const maxProposalAttempts = 3

type Proposer interface {
	Propose(context.Context, string, appipc.EpochStateBody) (brokerclient.FoldResult, error)
}

type IDFunc func() string

type Engine struct {
	applier    *applier.Host
	supervisor *supervisor.Controller
	proposer   Proposer
	nextID     IDFunc
	now        func() int64
}

type Outcome struct {
	RunID            string
	Case             Case
	Tuple            appipc.EpochStateBody
	ProposalAttempts int
	AssignOpen       bool
}

func New(host *applier.Host, proposer Proposer, nextID IDFunc, now func() int64) *Engine {
	if now == nil {
		now = func() int64 { return time.Now().UnixNano() }
	}
	return &Engine{applier: host, supervisor: supervisor.New(host), proposer: proposer, nextID: nextID, now: now}
}

func (engine *Engine) Run(ctx context.Context) ([]Outcome, error) {
	if engine == nil || engine.applier == nil || engine.proposer == nil || engine.nextID == nil {
		return nil, errors.New("recovery: incomplete dependencies")
	}
	states, err := engine.load(ctx)
	if err != nil {
		return nil, err
	}
	outcomes := make([]Outcome, 0, len(states))
	for _, state := range states {
		outcome, err := engine.recoverRun(ctx, state)
		if err != nil {
			return outcomes, fmt.Errorf("recovery: run %q: %w", state.runID, err)
		}
		outcomes = append(outcomes, outcome)
	}
	return outcomes, nil
}

type durableState struct {
	runID, epoch, stateSeq                 string
	generationID, workerState, workerEpoch string
	activeLease                            bool
}

func (engine *Engine) load(ctx context.Context) ([]durableState, error) {
	value, err := engine.applier.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		rows, err := snapshot.QueryContext(ctx, `SELECT r.run_id,e.turn_epoch,e.state_seq FROM runs r JOIN epochs e ON e.run_id=r.run_id WHERE r.state NOT IN ('COMPLETED','FAILED','CANCELLED') ORDER BY r.run_id`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var states []durableState
		for rows.Next() {
			var state durableState
			if err := rows.Scan(&state.runID, &state.epoch, &state.stateSeq); err != nil {
				return nil, err
			}
			states = append(states, state)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		for index := range states {
			state := &states[index]
			err := snapshot.QueryRowContext(ctx, `SELECT generation_id,state,turn_epoch FROM workers WHERE run_id=? AND turn_epoch=? ORDER BY created_at DESC,generation_id DESC LIMIT 1`, state.runID, state.epoch).Scan(&state.generationID, &state.workerState, &state.workerEpoch)
			if store.IsNoRows(err) {
				err = snapshot.QueryRowContext(ctx, `SELECT generation_id,state,turn_epoch FROM workers WHERE run_id=? ORDER BY created_at DESC,generation_id DESC LIMIT 1`, state.runID).Scan(&state.generationID, &state.workerState, &state.workerEpoch)
			}
			if err != nil && !store.IsNoRows(err) {
				return nil, err
			}
			var count int
			if err := snapshot.QueryRowContext(ctx, `SELECT COUNT(*) FROM leases WHERE run_id=? AND lease_kind='worker' AND state='ACTIVE'`, state.runID).Scan(&count); err != nil {
				return nil, err
			}
			state.activeLease = count == 1
		}
		return states, nil
	}))
	if err != nil {
		return nil, err
	}
	return value.([]durableState), nil
}

func (engine *Engine) recoverRun(ctx context.Context, state durableState) (Outcome, error) {
	classification := classify(state)
	var generationID string
	switch classification {
	case CaseLeased:
		generationID = engine.nextID()
		if generationID == "" {
			return Outcome{}, errors.New("empty successor generation")
		}
		result, err := engine.supervisor.Retire(ctx, supervisor.RetireRequest{
			RunID: state.runID, GenerationID: state.generationID, SuccessorGenerationID: generationID,
			At: engine.now(), CountFailure: false,
		})
		if err != nil {
			return Outcome{}, err
		}
		if result.Branch != supervisor.RetirementOrdinary {
			return Outcome{}, fmt.Errorf("retirement reached terminal branch %q", result.Branch)
		}
	case CasePreLease:
		generationID = engine.nextID()
		if generationID == "" {
			return Outcome{}, errors.New("empty washout successor generation")
		}
		if _, err := engine.supervisor.Washout(ctx, supervisor.WashoutRequest{
			RunID: state.runID, GenerationID: state.generationID, SuccessorGenerationID: generationID, At: engine.now(),
		}); err != nil {
			return Outcome{}, err
		}
	case CaseRetired, CaseInitial:
		generationID = engine.nextID()
		if generationID == "" {
			return Outcome{}, errors.New("empty recovery generation")
		}
		if _, err := engine.applier.Apply(ctx, allocateEvent{runID: state.runID, generationID: generationID, epoch: state.epoch, at: engine.now()}); err != nil {
			return Outcome{}, err
		}
	default:
		return Outcome{}, errors.New("unclassified durable state")
	}
	tuple, err := engine.tuple(ctx, state.runID, generationID)
	if err != nil {
		return Outcome{}, err
	}
	outcome := Outcome{RunID: state.runID, Case: classification, Tuple: tuple}
	for attempt := 1; attempt <= maxProposalAttempts; attempt++ {
		outcome.ProposalAttempts = attempt
		correlation := fmt.Sprintf("recovery-%s-%d", state.runID, attempt)
		fold, err := engine.proposer.Propose(ctx, correlation, tuple)
		if err != nil {
			return Outcome{}, err
		}
		switch fold.Action {
		case brokerclient.OpenAssign:
			outcome.AssignOpen = true
			return outcome, nil
		case brokerclient.Repropose, brokerclient.AwaitEventOrDeadline:
			continue
		default:
			return Outcome{}, fmt.Errorf("proposal failed closed with action %q", fold.Action)
		}
	}
	return Outcome{}, errors.New("proposal retry bound exhausted")
}

func classify(state durableState) Case {
	if state.activeLease && state.workerState == "LEASED" && state.workerEpoch == state.epoch {
		return CaseLeased
	}
	switch state.workerState {
	case "ALLOCATED", "SPAWNING", "READY":
		if state.workerEpoch == state.epoch {
			return CasePreLease
		}
	case "FAILED", "RETIRING", "TERMINATED":
		return CaseRetired
	case "":
		return CaseInitial
	}
	return ""
}

type allocateEvent struct {
	runID, generationID, epoch string
	at                         int64
}

func (event allocateEvent) RunID() string { return event.runID }
func (event allocateEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	_, err := tx.ExecContext(ctx, `INSERT INTO workers(generation_id,run_id,turn_epoch,state,created_at) VALUES(?,?,?,?,?)`, event.generationID, event.runID, event.epoch, "ALLOCATED", event.at)
	return applier.Result{}, err
}

func (engine *Engine) tuple(ctx context.Context, runID, generationID string) (appipc.EpochStateBody, error) {
	value, err := engine.applier.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var epoch, stateSeq string
		if err := snapshot.QueryRowContext(ctx, `SELECT turn_epoch,state_seq FROM epochs WHERE run_id=?`, runID).Scan(&epoch, &stateSeq); err != nil {
			return nil, err
		}
		wireEpoch, err := appipc.UnpadCounter(epoch)
		if err != nil {
			return nil, err
		}
		wireSeq, err := appipc.UnpadCounter(stateSeq)
		if err != nil {
			return nil, err
		}
		return appipc.EpochStateBody{RunID: runID, GenerationID: generationID, TurnEpoch: wireEpoch, LeaseState: appipc.LeaseUnleased, StateSeq: wireSeq}, nil
	}))
	if err != nil {
		return appipc.EpochStateBody{}, err
	}
	return value.(appipc.EpochStateBody), nil
}
