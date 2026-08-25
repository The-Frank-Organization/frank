package f59

import (
	"context"
	"errors"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
)

type expireEvent struct {
	runID  string
	turnID string
	at     int64
}

func (event expireEvent) RunID() string { return event.runID }

func (host *Host) Expire(ctx context.Context, runID, turnID string, at int64) (int, error) {
	if host == nil || host.applier == nil || runID == "" || turnID == "" {
		return 0, ErrInvalidRequest
	}
	result, err := host.applier.Apply(ctx, expireEvent{runID: runID, turnID: turnID, at: at})
	if err != nil {
		return 0, err
	}
	count, ok := result.Value.(int)
	if !ok {
		return 0, errors.New("f59: invalid expiry result")
	}
	return count, nil
}

func (event expireEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	count, err := ParkOpen(ctx, tx, event.runID, event.turnID, event.at)
	if err != nil {
		return applier.Result{}, err
	}
	if count == 0 {
		return applier.Result{Value: 0, NoMutation: true}, nil
	}
	return applier.Result{Value: count}, nil
}

// ParkOpen performs the D.4 pre-consume/post-consume crash split inside an
// existing lifecycle transaction. T8 calls this in the retirement commit.
func ParkOpen(ctx context.Context, tx *store.Tx, runID, turnID string, at int64) (int, error) {
	issued, err := tx.ExecContext(ctx, `UPDATE tool_authorizations SET state='VOID',void_reason='expired'
		WHERE run_id=? AND turn_id=? AND state='ISSUED'`, runID, turnID)
	if err != nil {
		return 0, err
	}
	consumed, err := tx.ExecContext(ctx, `INSERT INTO tool_calls
		(tool_call_id,run_id,turn_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,created_at,updated_at)
		SELECT tool_call_id,run_id,turn_id,turn_epoch,'UNKNOWN_TOOL_OUTCOME',canonical_tool_name,canonical_args_digest,issued_at,?
		FROM tool_authorizations WHERE run_id=? AND turn_id=? AND state='CONSUMED'`, at, runID, turnID)
	if err != nil {
		return 0, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE tool_authorizations SET state='UNKNOWN_TOOL_OUTCOME'
		WHERE run_id=? AND turn_id=? AND state='CONSUMED'`, runID, turnID); err != nil {
		return 0, err
	}
	issuedCount, _ := issued.RowsAffected()
	consumedCount, _ := consumed.RowsAffected()
	return int(issuedCount + consumedCount), nil
}
