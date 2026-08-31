package settle

import (
	"context"
	"errors"
	"regexp"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

var digestRE = regexp.MustCompile(`^[0-9a-f]{64}$`)

type CarriageDecision string

const (
	CarriageRecorded  CarriageDecision = "recorded"
	CarriageDuplicate CarriageDecision = "duplicate"
	CarriageConflict  CarriageDecision = "carriage_conflict"
	CarriageStale     CarriageDecision = "stale_epoch"
	CarriageUnknown   CarriageDecision = "unknown_attempt"
)

type AttemptResultRequest struct {
	RunID, TurnID, AttemptID, TurnEpoch, Disposition          string
	FrozenCoreDigest, ProviderLoweredToolsDigest, CancelPoint *string
	RefusalStage                                              *string
	At                                                        int64
}
type carriageEvent struct{ request AttemptResultRequest }

func (e carriageEvent) RunID() string { return e.request.RunID }
func (h *Host) RecordAttemptResult(ctx context.Context, r AttemptResultRequest) (CarriageDecision, error) {
	if h == nil || h.applier == nil || r.RunID == "" || r.TurnID == "" || r.AttemptID == "" || !validOptionalDigest(r.FrozenCoreDigest) || !validOptionalDigest(r.ProviderLoweredToolsDigest) || !validRefusalStage(r.Disposition, r.RefusalStage) {
		return "", errors.New("settle: invalid attempt result")
	}
	result, err := h.applier.Apply(ctx, carriageEvent{r})
	if err != nil {
		return "", err
	}
	return result.Value.(CarriageDecision), nil
}
func (e carriageEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	r := e.request
	var epoch, state, logical string
	var frozen, lowered, cancel, refusalStage *string
	err := tx.QueryRowContext(ctx, `SELECT turn_epoch,state,logical_surface_digest,frozen_core_digest,provider_lowered_tools_digest,cancel_point,refusal_stage FROM provider_attempts WHERE run_id=? AND turn_id=? AND attempt_id=?`, r.RunID, r.TurnID, r.AttemptID).Scan(&epoch, &state, &logical, &frozen, &lowered, &cancel, &refusalStage)
	if store.IsNoRows(err) {
		return carriageValue(CarriageUnknown, true), nil
	}
	if err != nil {
		return applier.Result{}, err
	}
	target, err := attemptState(r.Disposition, r.CancelPoint)
	if err != nil {
		return applier.Result{}, err
	}
	if state != "OPEN" && state != "STREAMING" {
		if state == target && equalString(frozen, r.FrozenCoreDigest) && equalString(lowered, r.ProviderLoweredToolsDigest) && equalString(cancel, r.CancelPoint) && equalString(refusalStage, r.RefusalStage) {
			return carriageValue(CarriageDuplicate, true), nil
		}
		return carriageValue(CarriageConflict, true), nil
	}
	var current string
	if err := tx.QueryRowContext(ctx, `SELECT turn_epoch FROM epochs WHERE run_id=?`, r.RunID).Scan(&current); err != nil {
		return applier.Result{}, err
	}
	relation, err := compareWire(r.TurnEpoch, unpad(current))
	if err != nil {
		return applier.Result{}, err
	}
	if relation != 0 || unpad(epoch) != r.TurnEpoch {
		return carriageValue(CarriageStale, true), nil
	}
	var cancellationID any
	if target == "CANCELLED" {
		var id string
		if err := tx.QueryRowContext(ctx, `SELECT cancellation_id FROM cancellations WHERE run_id=? AND target_kind='attempt' AND target_id=? AND epoch=? AND disposition IN ('PENDING','ACKNOWLEDGED')`, r.RunID, r.AttemptID, current).Scan(&id); err != nil {
			return applier.Result{}, err
		}
		cancellationID = id
	}
	_, err = tx.ExecContext(ctx, `UPDATE provider_attempts SET state=?,frozen_core_digest=?,provider_lowered_tools_digest=?,cancel_point=?,refusal_stage=?,cancellation_id=?,updated_at=? WHERE attempt_id=?`, target, r.FrozenCoreDigest, r.ProviderLoweredToolsDigest, r.CancelPoint, r.RefusalStage, cancellationID, r.At, r.AttemptID)
	return carriageValue(CarriageRecorded, false), err
}
func carriageValue(d CarriageDecision, no bool) applier.Result {
	return applier.Result{Value: d, NoMutation: no}
}
func validOptionalDigest(v *string) bool { return v == nil || digestRE.MatchString(*v) }
func validRefusalStage(disposition string, stage *string) bool {
	if disposition != "rejected_local" {
		return stage == nil
	}
	return stage != nil && (*stage == "pre_freeze" || *stage == "post_freeze")
}
func attemptState(disposition string, cancel *string) (string, error) {
	switch disposition {
	case "sent_completed":
		if cancel != nil {
			return "", errors.New("settle: completion forbids cancel point")
		}
		return "COMPLETED", nil
	case "denied", "transport_failed", "rejected_local":
		if cancel != nil {
			return "", errors.New("settle: rejected result forbids cancel point")
		}
		return "REJECTED_LOCAL", nil
	case "unknown":
		if cancel != nil {
			return "", errors.New("settle: unknown forbids cancel point")
		}
		return "UNKNOWN_PROVIDER_OUTCOME", nil
	case "cancelled":
		if cancel == nil || (*cancel != "pre_transport" && *cancel != "post_invocation") {
			return "", errors.New("settle: cancelled requires closed cancel point")
		}
		return "CANCELLED", nil
	default:
		return "", errors.New("settle: unknown attempt disposition")
	}
}

type RowState string

const (
	RowPresent  RowState = "present"
	RowNotFound RowState = "not_found"
)

type AttemptRow struct {
	State                                        RowState
	LogicalSurfaceDigest                         string
	FrozenCoreDigest, ProviderLoweredToolsDigest *string
	RefusalStage                                 *string
}

func (h *Host) QueryAttempt(ctx context.Context, runID, turnID, attemptID string) (AttemptRow, error) {
	if h == nil || h.applier == nil {
		return AttemptRow{}, errors.New("settle: nil host")
	}
	value, err := h.applier.Read(ctx, applier.QueryFunc(func(ctx context.Context, s *store.Snapshot) (any, error) {
		var row AttemptRow
		err := s.QueryRowContext(ctx, `SELECT logical_surface_digest,frozen_core_digest,provider_lowered_tools_digest,refusal_stage FROM provider_attempts WHERE run_id=? AND turn_id=? AND attempt_id=?`, runID, turnID, attemptID).Scan(&row.LogicalSurfaceDigest, &row.FrozenCoreDigest, &row.ProviderLoweredToolsDigest, &row.RefusalStage)
		if store.IsNoRows(err) {
			return AttemptRow{State: RowNotFound}, nil
		}
		row.State = RowPresent
		return row, err
	}))
	if err != nil {
		return AttemptRow{}, err
	}
	return value.(AttemptRow), nil
}
