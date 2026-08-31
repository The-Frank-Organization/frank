package supervisor

import (
	"context"
	"errors"
	"math"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/f59"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type WashoutRequest struct {
	RunID, GenerationID, SuccessorGenerationID string
	At                                         int64
}

type WashoutResult struct {
	TurnEpoch    string
	Failures     uint64
	BackoffUntil *int64
	Terminal     bool
	Idempotent   bool
}

type washoutEvent struct{ request WashoutRequest }

func (event washoutEvent) RunID() string { return event.request.RunID }

func (controller *Controller) Washout(ctx context.Context, request WashoutRequest) (WashoutResult, error) {
	if controller == nil || controller.applier == nil || request.RunID == "" || request.GenerationID == "" || request.SuccessorGenerationID == "" {
		return WashoutResult{}, errors.New("supervisor: invalid washout")
	}
	result, err := controller.applier.Apply(ctx, washoutEvent{request: request})
	if err != nil {
		return WashoutResult{}, err
	}
	return result.Value.(WashoutResult), nil
}

func (event washoutEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	var state, epoch, failuresStored string
	err := tx.QueryRowContext(ctx, `SELECT w.state,w.turn_epoch,r.consecutive_failures FROM workers w JOIN runs r ON r.run_id=w.run_id WHERE w.generation_id=? AND w.run_id=?`, request.GenerationID, request.RunID).Scan(&state, &epoch, &failuresStored)
	if err != nil {
		return applier.Result{}, err
	}
	if state == "FAILED" || state == "TERMINATED" {
		failures, _ := parseStored(failuresStored)
		return applier.Result{Value: WashoutResult{TurnEpoch: unpad(epoch), Failures: failures, Idempotent: true}, NoMutation: true}, nil
	}
	if state != "ALLOCATED" && state != "SPAWNING" && state != "READY" {
		return applier.Result{}, errors.New("supervisor: washout requires a pre-lease worker")
	}
	failures, err := parseStored(failuresStored)
	if err != nil || failures == math.MaxUint64 {
		return applier.Result{}, errors.New("supervisor: failure counter exhausted")
	}
	failures++
	if _, err := tx.ExecContext(ctx, `UPDATE workers SET state='FAILED',updated_at=? WHERE generation_id=?`, request.At, request.GenerationID); err != nil {
		return applier.Result{}, err
	}
	if failures >= MaxConsecutiveFailures {
		if _, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED',consecutive_failures=?,backoff_until=NULL,updated_at=? WHERE run_id=?`, pad(failures), request.At, request.RunID); err != nil {
			return applier.Result{}, err
		}
		return applier.Result{Value: WashoutResult{TurnEpoch: unpad(epoch), Failures: failures, Terminal: true}}, nil
	}
	backoff := backoffDeadline(request.At, failures)
	if _, err := tx.ExecContext(ctx, `UPDATE runs SET consecutive_failures=?,backoff_until=?,updated_at=? WHERE run_id=?`, pad(failures), backoff, request.At, request.RunID); err != nil {
		return applier.Result{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO workers(generation_id,run_id,turn_epoch,state,created_at) VALUES(?,?,?,?,?)`, request.SuccessorGenerationID, request.RunID, epoch, "ALLOCATED", request.At); err != nil {
		return applier.Result{}, err
	}
	return applier.Result{Value: WashoutResult{TurnEpoch: unpad(epoch), Failures: failures, BackoffUntil: backoff}}, nil
}

type runEvent struct {
	runID string
	at    int64
	kind  string
}

func (event runEvent) RunID() string { return event.runID }
func (event runEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	var result error
	switch event.kind {
	case "reset":
		_, result = tx.ExecContext(ctx, `UPDATE runs SET consecutive_failures=?,backoff_until=NULL,updated_at=? WHERE run_id=?`, pad(0), event.at, event.runID)
	case "cancel":
		_, result = tx.ExecContext(ctx, `UPDATE runs SET state='CANCELLED',backoff_until=NULL,updated_at=? WHERE run_id=? AND state NOT IN ('COMPLETED','FAILED','CANCELLED')`, event.at, event.runID)
	default:
		result = errors.New("supervisor: invalid run event")
	}
	return applier.Result{}, result
}

func (controller *Controller) ResetFailures(ctx context.Context, runID string, at int64) error {
	if controller == nil || controller.applier == nil || runID == "" {
		return errors.New("supervisor: invalid reset")
	}
	_, err := controller.applier.Apply(ctx, runEvent{runID: runID, at: at, kind: "reset"})
	return err
}

func (controller *Controller) CancelRun(ctx context.Context, runID string, at int64) error {
	if controller == nil || controller.applier == nil || runID == "" {
		return errors.New("supervisor: invalid cancel")
	}
	_, err := controller.applier.Apply(ctx, runEvent{runID: runID, at: at, kind: "cancel"})
	return err
}

func (controller *Controller) LeaseInvalidHandler(nextGenerationID func() string) func(context.Context, *store.Tx, f59.LeaseFault) error {
	return func(ctx context.Context, tx *store.Tx, fault f59.LeaseFault) error {
		if nextGenerationID == nil {
			return errors.New("supervisor: successor id source is required")
		}
		_, _, err := retireTx(ctx, tx, RetireRequest{
			RunID: fault.RunID, TurnID: fault.TurnID, GenerationID: fault.GenerationID,
			SuccessorGenerationID: nextGenerationID(), At: fault.At, CountFailure: true,
		})
		return err
	}
}

func backoffDeadline(at int64, failures uint64) *int64 {
	if failures < 2 {
		return nil
	}
	value := at + int64(1<<(failures-2))
	return &value
}

type EpochRelation string

const (
	EpochCurrent EpochRelation = "current"
	EpochStale   EpochRelation = "stale"
	EpochFuture  EpochRelation = "future_fault"
)

func ClassifyEpoch(presented, current string) (EpochRelation, error) {
	left, err := appipc.ParseCounter(presented)
	if err != nil {
		return "", err
	}
	right, err := appipc.ParseCounter(current)
	if err != nil {
		return "", err
	}
	if left < right {
		return EpochStale, nil
	}
	if left > right {
		return EpochFuture, nil
	}
	return EpochCurrent, nil
}

func ValidateAttachResult(result string) error {
	switch result {
	case appipc.AttachOK, appipc.AttachSuspended, appipc.AttachTupleMismatch:
		return nil
	default:
		return errors.New("supervisor: unknown attach result")
	}
}
