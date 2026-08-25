package scheduler

import (
	"context"
	"errors"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

type AttemptDecision string

const (
	AttemptCommitted    AttemptDecision = "committed"
	AttemptStaleEpoch   AttemptDecision = "stale_epoch"
	AttemptInvalidTurn  AttemptDecision = "invalid_turn"
	AttemptInvalidLease AttemptDecision = "invalid_lease"
)

type AttemptRequest struct {
	RunID, TurnID, AttemptID, GenerationID, TurnEpoch string
	LogicalSurfaceDigest                              string
	At                                                int64
}

type AttemptResult struct {
	Decision      AttemptDecision
	ParkedUnknown []appipc.ParkedUnknown
}

type attemptEvent struct{ request AttemptRequest }

func (event attemptEvent) RunID() string { return event.request.RunID }

func (scheduler *Scheduler) OpenAttempt(ctx context.Context, request AttemptRequest) (AttemptResult, error) {
	if scheduler == nil || scheduler.applier == nil || request.RunID == "" || request.TurnID == "" || request.AttemptID == "" || len(request.LogicalSurfaceDigest) != 64 {
		return AttemptResult{}, errors.New("scheduler: invalid attempt")
	}
	result, err := scheduler.applier.Apply(ctx, attemptEvent{request: request})
	if err != nil {
		return AttemptResult{}, err
	}
	return result.Value.(AttemptResult), nil
}

func (event attemptEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	var current string
	if err := tx.QueryRowContext(ctx, `SELECT turn_epoch FROM epochs WHERE run_id=?`, request.RunID).Scan(&current); err != nil {
		return applier.Result{}, err
	}
	if unpad(current) != request.TurnEpoch {
		return rejectAttempt(request.AttemptID, AttemptStaleEpoch), nil
	}
	var turnState string
	err := tx.QueryRowContext(ctx, `SELECT state FROM turns WHERE run_id=? AND turn_id=?`, request.RunID, request.TurnID).Scan(&turnState)
	if store.IsNoRows(err) || turnState != "ACTIVE" {
		return rejectAttempt(request.AttemptID, AttemptInvalidTurn), nil
	}
	if err != nil {
		return applier.Result{}, err
	}
	var leases int
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM leases WHERE run_id=? AND generation_id=? AND turn_epoch=? AND state='ACTIVE'`, request.RunID, request.GenerationID, current).Scan(&leases); err != nil {
		return applier.Result{}, err
	}
	if leases != 2 {
		return rejectAttempt(request.AttemptID, AttemptInvalidLease), nil
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO provider_attempts(attempt_id,run_id,turn_id,turn_epoch,state,logical_surface_digest,created_at) VALUES(?,?,?,?,?,?,?)`, request.AttemptID, request.RunID, request.TurnID, current, "OPEN", request.LogicalSurfaceDigest, request.At); err != nil {
		return applier.Result{}, err
	}
	parked, err := snapshotParked(ctx, tx, request.RunID)
	if err != nil {
		return applier.Result{}, err
	}
	body := appipc.AttemptOpenOKBody{AttemptID: request.AttemptID, ParkedUnknown: parked}
	return applier.Result{Value: AttemptResult{Decision: AttemptCommitted, ParkedUnknown: parked}, Emissions: []applier.Emission{{Kind: "attempt_open_ok", Value: body}}}, nil
}

func rejectAttempt(attemptID string, decision AttemptDecision) applier.Result {
	body := appipc.AttemptOpenRejectBody{AttemptID: attemptID, Reason: string(decision)}
	return applier.Result{Value: AttemptResult{Decision: decision}, Emissions: []applier.Emission{{Kind: "attempt_open_reject", Value: body}}, NoMutation: true}
}
