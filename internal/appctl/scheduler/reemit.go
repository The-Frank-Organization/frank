package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type reemitActiveEvent struct{ runID string }

func (event reemitActiveEvent) RunID() string { return event.runID }

// ReemitActive reconstructs a committed ACTIVE turn_open solely from its
// durable admission sources and sends it through the applier's post-commit
// emission path without mutating state or consuming a wake row again.
func (scheduler *Scheduler) ReemitActive(ctx context.Context, runID string) ([]appipc.TurnOpenBody, error) {
	if scheduler == nil || scheduler.applier == nil || runID == "" {
		return nil, errors.New("scheduler: invalid turn_open recovery")
	}
	result, err := scheduler.applier.Apply(ctx, reemitActiveEvent{runID: runID})
	if err != nil {
		return nil, err
	}
	return result.Value.([]appipc.TurnOpenBody), nil
}

func (event reemitActiveEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	var turnID, epoch, encodedRef, storedDisposition, createAuthID, runPhase, runSessionLogPath string
	var predecessor *string
	var snapshot []byte
	err := tx.QueryRowContext(ctx, `SELECT t.turn_id,t.turn_epoch,t.admission_ref,t.run_disposition,t.create_auth_id,t.predecessor_turn_id,t.resume_snapshot,r.run_phase,r.session_log_path
		FROM turns t JOIN runs r ON r.run_id=t.run_id WHERE t.run_id=? AND t.state='ACTIVE'`, event.runID).
		Scan(&turnID, &epoch, &encodedRef, &storedDisposition, &createAuthID, &predecessor, &snapshot, &runPhase, &runSessionLogPath)
	if store.IsNoRows(err) {
		return applier.Result{Value: []appipc.TurnOpenBody{}, NoMutation: true}, nil
	}
	if err != nil {
		return applier.Result{}, err
	}
	if runPhase != "create_authorized" && runPhase != "established" {
		return applier.Result{}, errors.New("scheduler: active turn lacks durable create authorization")
	}
	var admissionRef appipc.AdmissionRef
	if err := json.Unmarshal([]byte(encodedRef), &admissionRef); err != nil {
		return applier.Result{}, err
	}
	parked, err := persistedDisclosures(ctx, tx, turnID)
	if err != nil {
		return applier.Result{}, err
	}
	body := appipc.TurnOpenBody{
		TurnID: turnID, AdmissionRef: admissionRef, ParkedUnknown: parked,
		RunDisposition: storedDisposition, CreateAuthID: createAuthID, SessionLogPath: runSessionLogPath,
	}
	if predecessor == nil {
		if runSessionLogPath == "" {
			return applier.Result{}, errors.New("scheduler: active genesis lacks durable session log path")
		}
		body.RunDisposition = appipc.RunDispositionResume
	} else {
		var persisted struct {
			SessionLogPath     string                    `json:"session_log_path"`
			SettlementManifest appipc.SettlementManifest `json:"settlement_manifest"`
		}
		if len(snapshot) == 0 {
			return applier.Result{}, errors.New("scheduler: continuation resume snapshot is absent")
		}
		if err := json.Unmarshal(snapshot, &persisted); err != nil || persisted.SessionLogPath == "" {
			return applier.Result{}, fmt.Errorf("scheduler: continuation resume snapshot is invalid: %v", err)
		}
		body.RunDisposition = appipc.RunDispositionResume
		body.SessionLogPath = persisted.SessionLogPath
		body.SettlementManifest = &persisted.SettlementManifest
		body.PredecessorTurnID = predecessor
	}
	bodies := []appipc.TurnOpenBody{body}
	return applier.Result{Value: bodies, Emissions: []applier.Emission{{Kind: "turn_open", Value: body}}, NoMutation: true}, nil
}

func persistedDisclosures(ctx context.Context, tx *store.Tx, turnID string) ([]appipc.ParkedUnknown, error) {
	rows, err := tx.QueryContext(ctx, `SELECT source_turn_id,tool_call_id,ticket_id,state_at_disclosure,canonical_tool_name,canonical_args_digest
		FROM turn_disclosures WHERE disclosing_turn_id=? ORDER BY tool_call_id,ticket_id`, turnID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]appipc.ParkedUnknown, 0)
	for rows.Next() {
		var row appipc.ParkedUnknown
		if err := rows.Scan(&row.TurnID, &row.ToolCallID, &row.TicketID, &row.State, &row.CanonicalToolName, &row.CanonicalArgsDigest); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, rows.Err()
}
