package settle

import (
	"context"
	"errors"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type Host struct{ applier *applier.Host }

func New(host *applier.Host) *Host { return &Host{applier: host} }

type ReceiptDecision string

const (
	ReceiptRecorded    ReceiptDecision = "recorded"
	ReceiptDuplicate   ReceiptDecision = "duplicate"
	ReceiptConflict    ReceiptDecision = "receipt_conflict"
	ReceiptStale       ReceiptDecision = "stale_epoch"
	ReceiptUnknown     ReceiptDecision = "unknown_attempt"
	ReceiptFutureFault ReceiptDecision = "future_epoch_fault"
)

type ContentReadyRequest struct {
	RunID, TurnID, AttemptID, RoundIdentity, SeqHWM, GenerationID, TurnEpoch string
	At                                                                       int64
}
type receiptEvent struct{ request ContentReadyRequest }

func (e receiptEvent) RunID() string { return e.request.RunID }
func (h *Host) RecordContentReady(ctx context.Context, r ContentReadyRequest) (ReceiptDecision, error) {
	if h == nil || h.applier == nil || r.RunID == "" || r.TurnID == "" || r.AttemptID == "" || r.RoundIdentity == "" || r.GenerationID == "" {
		return "", errors.New("settle: invalid receipt")
	}
	if _, err := appipc.ParseCounter(r.SeqHWM); err != nil {
		return "", err
	}
	if _, err := appipc.ParseCounter(r.TurnEpoch); err != nil {
		return "", err
	}
	result, err := h.applier.Apply(ctx, receiptEvent{r})
	if err != nil {
		return "", err
	}
	return result.Value.(ReceiptDecision), nil
}
func (e receiptEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	r := e.request
	var round, seq string
	err := tx.QueryRowContext(ctx, `SELECT round_identity,seq_hwm FROM content_ready_receipts WHERE run_id=? AND turn_id=? AND attempt_id=?`, r.RunID, r.TurnID, r.AttemptID).Scan(&round, &seq)
	if err == nil {
		if round == r.RoundIdentity && unpad(seq) == r.SeqHWM {
			return receiptResult(ReceiptDuplicate, true), nil
		}
		return receiptResult(ReceiptConflict, true), nil
	}
	if !store.IsNoRows(err) {
		return applier.Result{}, err
	}
	var count int
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM provider_attempts WHERE run_id=? AND turn_id=? AND attempt_id=?`, r.RunID, r.TurnID, r.AttemptID).Scan(&count); err != nil {
		return applier.Result{}, err
	}
	if count != 1 {
		return receiptResult(ReceiptUnknown, true), nil
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
		return receiptResult(ReceiptFutureFault, true), nil
	}
	var sender int
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM leases WHERE run_id=? AND lease_kind='worker' AND generation_id=? AND turn_epoch=? AND state='ACTIVE'`, r.RunID, r.GenerationID, pad(r.TurnEpoch)).Scan(&sender); err != nil {
		return applier.Result{}, err
	}
	if relation < 0 || sender != 1 {
		return receiptResult(ReceiptStale, true), nil
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO content_ready_receipts(run_id,turn_id,attempt_id,round_identity,seq_hwm,generation_id,committed_at) VALUES(?,?,?,?,?,?,?)`, r.RunID, r.TurnID, r.AttemptID, r.RoundIdentity, pad(r.SeqHWM), r.GenerationID, r.At)
	return receiptResult(ReceiptRecorded, false), err
}
func receiptResult(d ReceiptDecision, no bool) applier.Result {
	return applier.Result{Value: d, NoMutation: no}
}
func pad(w string) string   { v, _ := appipc.PadCounter(w); return v }
func unpad(s string) string { v, _ := appipc.UnpadCounter(s); return v }
func compareWire(a, b string) (int, error) {
	x, e := appipc.ParseCounter(a)
	if e != nil {
		return 0, e
	}
	y, e := appipc.ParseCounter(b)
	if e != nil {
		return 0, e
	}
	if x < y {
		return -1, nil
	}
	if x > y {
		return 1, nil
	}
	return 0, nil
}
