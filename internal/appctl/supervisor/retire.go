package supervisor

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/f59"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

const MaxConsecutiveFailures = 10

type RetirementBranch string

const (
	RetirementOrdinary       RetirementBranch = "ordinary"
	RetirementParkedCap      RetirementBranch = "parked_cap_terminal"
	RetirementFailureCeiling RetirementBranch = "failure_ceiling_terminal"
)

type RetireRequest struct {
	RunID, TurnID, GenerationID, SuccessorGenerationID string
	At                                                 int64
	CountFailure                                       bool
}

type RetireResult struct {
	Branch       RetirementBranch
	TurnEpoch    string
	Parked       int
	Failures     uint64
	BackoffUntil *int64
	Idempotent   bool
}

type Controller struct{ applier *applier.Host }

func New(host *applier.Host) *Controller { return &Controller{applier: host} }

type retireEvent struct {
	controller *Controller
	request    RetireRequest
}

func (event retireEvent) RunID() string { return event.request.RunID }

func (controller *Controller) Retire(ctx context.Context, request RetireRequest) (RetireResult, error) {
	if controller == nil || controller.applier == nil || request.RunID == "" || request.GenerationID == "" || request.SuccessorGenerationID == "" {
		return RetireResult{}, errors.New("supervisor: invalid retirement")
	}
	result, err := controller.applier.Apply(ctx, retireEvent{controller: controller, request: request})
	if err != nil {
		return RetireResult{}, err
	}
	return result.Value.(RetireResult), nil
}

func (event retireEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	result, mutated, err := retireTx(ctx, tx, event.request)
	if err != nil {
		return applier.Result{}, err
	}
	return applier.Result{Value: result, NoMutation: !mutated}, nil
}

func retireTx(ctx context.Context, tx *store.Tx, request RetireRequest) (RetireResult, bool, error) {
	var workerState, workerEpoch string
	if err := tx.QueryRowContext(ctx, `SELECT state,turn_epoch FROM workers WHERE generation_id=? AND run_id=?`, request.GenerationID, request.RunID).Scan(&workerState, &workerEpoch); err != nil {
		return RetireResult{}, false, err
	}
	var currentStored, failuresStored string
	if err := tx.QueryRowContext(ctx, `SELECT e.turn_epoch,r.consecutive_failures FROM epochs e JOIN runs r ON r.run_id=e.run_id WHERE e.run_id=?`, request.RunID).Scan(&currentStored, &failuresStored); err != nil {
		return RetireResult{}, false, err
	}
	current := unpad(currentStored)
	if workerState == "FAILED" || workerState == "RETIRING" || workerState == "TERMINATED" {
		if relation, _ := compare(workerEpoch, currentStored); relation < 0 {
			parked, _ := parkedCount(ctx, tx, request.RunID)
			failures, _ := parseStored(failuresStored)
			return RetireResult{TurnEpoch: current, Parked: parked, Failures: failures, Idempotent: true}, false, nil
		}
	}
	nextEpoch, err := incrementStored(currentStored)
	if err != nil {
		return RetireResult{}, false, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE workers SET state='FAILED',updated_at=? WHERE generation_id=?`, request.At, request.GenerationID); err != nil {
		return RetireResult{}, false, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE leases SET state='RELEASED',released_at=? WHERE run_id=? AND state='ACTIVE'`, request.At, request.RunID); err != nil {
		return RetireResult{}, false, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE epochs SET turn_epoch=? WHERE run_id=?`, nextEpoch, request.RunID); err != nil {
		return RetireResult{}, false, err
	}
	if _, err := f59.ParkOpen(ctx, tx, request.RunID, request.TurnID, request.At); err != nil {
		return RetireResult{}, false, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE provider_attempts SET state='UNKNOWN_PROVIDER_OUTCOME',updated_at=? WHERE run_id=? AND state IN ('OPEN','STREAMING')`, request.At, request.RunID); err != nil {
		return RetireResult{}, false, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE turns SET state='INTERRUPTED',updated_at=? WHERE run_id=? AND state IN ('ADMITTED','ACTIVE')`, request.At, request.RunID); err != nil {
		return RetireResult{}, false, err
	}
	parked, err := parkedCount(ctx, tx, request.RunID)
	if err != nil {
		return RetireResult{}, false, err
	}
	failures, err := parseStored(failuresStored)
	if err != nil {
		return RetireResult{}, false, err
	}
	if request.CountFailure {
		if failures == math.MaxUint64 {
			return RetireResult{}, false, errors.New("supervisor: failure counter exhausted")
		}
		failures++
	}
	if parked > appipc.MaxParkedRowsPerRun {
		if _, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED',stop_reason='parked_unknown_capacity_exceeded',consecutive_failures=?,backoff_until=NULL,updated_at=? WHERE run_id=?`, pad(failures), request.At, request.RunID); err != nil {
			return RetireResult{}, false, err
		}
		return RetireResult{Branch: RetirementParkedCap, TurnEpoch: unpad(nextEpoch), Parked: parked, Failures: failures}, true, nil
	}
	if failures >= MaxConsecutiveFailures {
		if _, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED',consecutive_failures=?,backoff_until=NULL,updated_at=? WHERE run_id=?`, pad(failures), request.At, request.RunID); err != nil {
			return RetireResult{}, false, err
		}
		return RetireResult{Branch: RetirementFailureCeiling, TurnEpoch: unpad(nextEpoch), Parked: parked, Failures: failures}, true, nil
	}
	var backoff *int64
	if request.CountFailure {
		backoff = backoffDeadline(request.At, failures)
	}
	if _, err := tx.ExecContext(ctx, `UPDATE runs SET consecutive_failures=?,backoff_until=?,updated_at=? WHERE run_id=?`, pad(failures), backoff, request.At, request.RunID); err != nil {
		return RetireResult{}, false, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO workers(generation_id,run_id,turn_epoch,state,created_at) VALUES(?,?,?,?,?)`, request.SuccessorGenerationID, request.RunID, nextEpoch, "ALLOCATED", request.At); err != nil {
		return RetireResult{}, false, err
	}
	return RetireResult{Branch: RetirementOrdinary, TurnEpoch: unpad(nextEpoch), Parked: parked, Failures: failures, BackoffUntil: backoff}, true, nil
}

func parkedCount(ctx context.Context, tx *store.Tx, runID string) (int, error) {
	var count int
	err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM tool_calls WHERE run_id=? AND state IN ('UNKNOWN_TOOL_OUTCOME','PARTIAL_TOOL_EFFECT')`, runID).Scan(&count)
	return count, err
}

func parseStored(value string) (uint64, error) { return appipc.ParseCounter(unpad(value)) }
func pad(value uint64) string {
	result, _ := appipc.PadCounter(appipc.FormatCounter(value))
	return result
}
func unpad(value string) string { result, _ := appipc.UnpadCounter(value); return result }
func incrementStored(value string) (string, error) {
	parsed, err := parseStored(value)
	if err != nil || parsed == math.MaxUint64 {
		return "", fmt.Errorf("supervisor: epoch exhausted")
	}
	return pad(parsed + 1), nil
}
func compare(left, right string) (int, error) {
	l, e := parseStored(left)
	if e != nil {
		return 0, e
	}
	r, e := parseStored(right)
	if e != nil {
		return 0, e
	}
	if l < r {
		return -1, nil
	}
	if l > r {
		return 1, nil
	}
	return 0, nil
}
