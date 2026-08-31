package scheduler

import (
	"context"
	"errors"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type wakeEvent struct {
	runID, relayID string
	at             int64
}

func (event wakeEvent) RunID() string { return event.runID }
func (event wakeEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	var disposition string
	err := tx.QueryRowContext(ctx, `SELECT disposition FROM wake_schedule WHERE relay_id=?`, event.relayID).Scan(&disposition)
	if err == nil {
		return applier.Result{Value: true, NoMutation: true}, nil
	}
	if !store.IsNoRows(err) {
		return applier.Result{}, err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO wake_schedule(relay_id,run_id,disposition,received_at) VALUES(?,?,'PENDING',?)`, event.relayID, event.runID, event.at)
	return applier.Result{Value: false}, err
}
func (scheduler *Scheduler) RecordWake(ctx context.Context, runID, relayID string, at int64) (bool, error) {
	if scheduler == nil || scheduler.applier == nil || runID == "" || relayID == "" {
		return false, errors.New("scheduler: invalid wake")
	}
	result, err := scheduler.applier.Apply(ctx, wakeEvent{runID: runID, relayID: relayID, at: at})
	if err != nil {
		return false, err
	}
	return result.Value.(bool), nil
}

type CancellationRequest struct {
	ID, RunID, TargetKind, TargetID, Epoch string
	At                                     int64
}
type cancellationEvent struct{ request CancellationRequest }

func (event cancellationEvent) RunID() string { return event.request.RunID }
func (event cancellationEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	r := event.request
	if r.TargetKind != "run" && r.TargetKind != "turn" && r.TargetKind != "attempt" {
		return applier.Result{}, errors.New("scheduler: invalid cancellation target")
	}
	if _, err := appipc.ParseCounter(r.Epoch); err != nil {
		return applier.Result{}, err
	}
	var id string
	err := tx.QueryRowContext(ctx, `SELECT cancellation_id FROM cancellations WHERE target_kind=? AND target_id=? AND epoch=?`, r.TargetKind, r.TargetID, pad(r.Epoch)).Scan(&id)
	if err == nil {
		return applier.Result{Value: true, NoMutation: true}, nil
	}
	if !store.IsNoRows(err) {
		return applier.Result{}, err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO cancellations(cancellation_id,run_id,target_kind,target_id,epoch,disposition,requested_at) VALUES(?,?,?,?,?,'PENDING',?)`, r.ID, r.RunID, r.TargetKind, r.TargetID, pad(r.Epoch), r.At)
	return applier.Result{Value: false}, err
}
func (scheduler *Scheduler) RequestCancellation(ctx context.Context, r CancellationRequest) (bool, error) {
	if scheduler == nil || scheduler.applier == nil || r.ID == "" || r.RunID == "" || r.TargetID == "" {
		return false, errors.New("scheduler: invalid cancellation")
	}
	result, err := scheduler.applier.Apply(ctx, cancellationEvent{r})
	if err != nil {
		return false, err
	}
	return result.Value.(bool), nil
}
func pad(wire string) string { value, _ := appipc.PadCounter(wire); return value }
