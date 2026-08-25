package scheduler

import (
	"context"
	"errors"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

type TerminalDecision string

const (
	TerminalRecorded  TerminalDecision = "terminal_recorded"
	TerminalDuplicate TerminalDecision = "duplicate"
	TerminalConflict  TerminalDecision = "conflicting_report"
	TerminalStale     TerminalDecision = "stale_epoch"
	TerminalUnknown   TerminalDecision = "unknown_turn"
)

type TerminalRequest struct {
	RunID, TurnID, TurnEpoch, Terminal string
	At                                 int64
}
type terminalEvent struct{ request TerminalRequest }

func (event terminalEvent) RunID() string { return event.request.RunID }
func (scheduler *Scheduler) RecordTerminal(ctx context.Context, r TerminalRequest) (TerminalDecision, error) {
	if scheduler == nil || scheduler.applier == nil || r.RunID == "" || r.TurnID == "" {
		return "", errors.New("scheduler: invalid terminal")
	}
	result, err := scheduler.applier.Apply(ctx, terminalEvent{r})
	if err != nil {
		return "", err
	}
	return result.Value.(TerminalDecision), nil
}
func (event terminalEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	r := event.request
	var current, turnEpoch, state string
	if err := tx.QueryRowContext(ctx, `SELECT turn_epoch FROM epochs WHERE run_id=?`, r.RunID).Scan(&current); err != nil {
		return applier.Result{}, err
	}
	if err := tx.QueryRowContext(ctx, `SELECT turn_epoch,state FROM turns WHERE run_id=? AND turn_id=?`, r.RunID, r.TurnID).Scan(&turnEpoch, &state); store.IsNoRows(err) {
		return terminalResult(TerminalUnknown, true), nil
	} else if err != nil {
		return applier.Result{}, err
	}
	if unpad(current) != r.TurnEpoch || unpad(turnEpoch) != r.TurnEpoch {
		return terminalResult(TerminalStale, true), nil
	}
	want, ok := terminalState(r.Terminal)
	if !ok {
		return applier.Result{}, errors.New("scheduler: unknown terminal")
	}
	if state == want {
		return terminalResult(TerminalDuplicate, true), nil
	}
	if state != "ACTIVE" && state != "ADMITTED" && state != "INTERRUPTED" {
		return terminalResult(TerminalConflict, true), nil
	}
	if _, err := tx.ExecContext(ctx, `UPDATE turns SET state=?,updated_at=? WHERE turn_id=?`, want, r.At, r.TurnID); err != nil {
		return applier.Result{}, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE leases SET state='RELEASED',released_at=? WHERE run_id=? AND lease_kind='turn' AND state='ACTIVE'`, r.At, r.RunID); err != nil {
		return applier.Result{}, err
	}
	if want == "CANCELLED" {
		if _, err := tx.ExecContext(ctx, `UPDATE cancellations SET disposition='COMPLETED',resolved_at=? WHERE run_id=? AND target_kind='turn' AND target_id=? AND epoch=? AND disposition IN ('PENDING','ACKNOWLEDGED')`, r.At, r.RunID, r.TurnID, current); err != nil {
			return applier.Result{}, err
		}
	}
	return terminalResult(TerminalRecorded, false), nil
}
func terminalResult(decision TerminalDecision, noMutation bool) applier.Result {
	return applier.Result{Value: decision, NoMutation: noMutation}
}
func terminalState(terminal string) (string, bool) {
	switch terminal {
	case "turn_completed":
		return "COMPLETED", true
	case "turn_cancelled":
		return "CANCELLED", true
	case "turn_refused", "turn_denied", "turn_failed", "turn_exhausted":
		return "FAILED", true
	default:
		return "", false
	}
}

var _ = appipc.TurnTerminalBody{}
