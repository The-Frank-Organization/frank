package settle

import (
	"context"
	"errors"
	"strings"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
)

type DispositionDecision string

const (
	DispositionRecorded    DispositionDecision = "recorded"
	DispositionDuplicate   DispositionDecision = "duplicate"
	DispositionConflict    DispositionDecision = "disposition_conflict"
	DispositionStale       DispositionDecision = "stale_epoch"
	DispositionUnknown     DispositionDecision = "unknown_turn"
	DispositionFutureFault DispositionDecision = "future_epoch_fault"
)

type DispositionPair struct {
	Disposition  string
	ResumeAction *string
}
type DispositionRequest struct {
	RunID, TurnID, TurnEpoch, GenerationID, Disposition string
	ResumeAction                                        *string
	At                                                  int64
}
type dispositionEvent struct{ request DispositionRequest }

func (e dispositionEvent) RunID() string { return e.request.RunID }
func (h *Host) ReportDisposition(ctx context.Context, r DispositionRequest) (DispositionDecision, DispositionPair, error) {
	if h == nil || h.applier == nil || r.RunID == "" || r.TurnID == "" || r.GenerationID == "" || !validPair(r.Disposition, r.ResumeAction) {
		return "", DispositionPair{}, errors.New("settle: invalid disposition")
	}
	result, err := h.applier.Apply(ctx, dispositionEvent{r})
	if err != nil {
		return "", DispositionPair{}, err
	}
	value := result.Value.(dispositionResult)
	return value.decision, value.pair, nil
}

type dispositionResult struct {
	decision DispositionDecision
	pair     DispositionPair
}

func (e dispositionEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	r := e.request
	var storedEpoch, state, disposition string
	var action *string
	err := tx.QueryRowContext(ctx, `SELECT turn_epoch,state,resume_disposition,resume_action FROM turns WHERE run_id=? AND turn_id=?`, r.RunID, r.TurnID).Scan(&storedEpoch, &state, &disposition, &action)
	if store.IsNoRows(err) {
		return dispositionValue(DispositionUnknown, DispositionPair{}, true), nil
	}
	if err != nil {
		return applier.Result{}, err
	}
	pair := storedPair(disposition, action)
	if disposition != "PENDING" {
		if strings.EqualFold(disposition, r.Disposition) && equalString(action, r.ResumeAction) {
			return dispositionValue(DispositionDuplicate, pair, true), nil
		}
		return dispositionValue(DispositionConflict, pair, true), nil
	}
	var current string
	if err := tx.QueryRowContext(ctx, `SELECT turn_epoch FROM epochs WHERE run_id=?`, r.RunID).Scan(&current); err != nil {
		return applier.Result{}, err
	}
	relation, err := compareWire(r.TurnEpoch, unpad(current))
	if err != nil {
		return applier.Result{}, err
	}
	if relation > 0 {
		return dispositionValue(DispositionFutureFault, pair, true), nil
	}
	var sender int
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM leases WHERE run_id=? AND lease_kind='worker' AND generation_id=? AND turn_epoch=? AND state='ACTIVE'`, r.RunID, r.GenerationID, pad(r.TurnEpoch)).Scan(&sender); err != nil {
		return applier.Result{}, err
	}
	if relation < 0 || unpad(storedEpoch) != r.TurnEpoch || sender != 1 {
		return dispositionValue(DispositionStale, pair, true), nil
	}
	upper := strings.ToUpper(r.Disposition)
	var storedAction any
	if r.ResumeAction != nil {
		storedAction = *r.ResumeAction
	}
	if _, err := tx.ExecContext(ctx, `UPDATE turns SET resume_disposition=?,resume_action=?,updated_at=? WHERE turn_id=?`, upper, storedAction, r.At, r.TurnID); err != nil {
		return applier.Result{}, err
	}
	committed := DispositionPair{Disposition: r.Disposition, ResumeAction: clone(r.ResumeAction)}
	return dispositionValue(DispositionRecorded, committed, false), nil
}
func dispositionValue(d DispositionDecision, p DispositionPair, no bool) applier.Result {
	return applier.Result{Value: dispositionResult{d, p}, NoMutation: no}
}
func validPair(d string, a *string) bool {
	return (d == "resumable" && a == nil) || (d == "degraded" && a != nil && *a == "re_derive")
}
func storedPair(d string, a *string) DispositionPair {
	if d == "PENDING" {
		return DispositionPair{Disposition: "pending"}
	}
	return DispositionPair{Disposition: strings.ToLower(d), ResumeAction: clone(a)}
}
func clone(v *string) *string {
	if v == nil {
		return nil
	}
	c := *v
	return &c
}
func equalString(a, b *string) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return *a == *b
}
