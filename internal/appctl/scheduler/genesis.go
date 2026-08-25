package scheduler

import (
	"context"
	"errors"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

type GenesisCommittedDecision string

const (
	GenesisCommittedIdentityFault GenesisCommittedDecision = "identity_fault"
	GenesisCommittedDuplicate     GenesisCommittedDecision = "established_duplicate"
	GenesisCommittedUnexpected    GenesisCommittedDecision = "created_unexpected_ack"
	GenesisCommittedStale         GenesisCommittedDecision = "stale_sender"
	GenesisCommittedEstablished   GenesisCommittedDecision = "established"
)

type GenesisCommittedRequest struct {
	RunID, BoundRunID, GenerationID, TurnEpoch string
	At                                         int64
}

type genesisCommittedEvent struct{ request GenesisCommittedRequest }

func (event genesisCommittedEvent) RunID() string { return event.request.RunID }

func (scheduler *Scheduler) RecordGenesisCommitted(ctx context.Context, request GenesisCommittedRequest) (GenesisCommittedDecision, error) {
	if scheduler == nil || scheduler.applier == nil || request.RunID == "" || request.BoundRunID == "" || request.GenerationID == "" {
		return "", errors.New("scheduler: malformed genesis_committed")
	}
	if _, err := appipc.ParseCounter(request.TurnEpoch); err != nil {
		return "", errors.New("scheduler: malformed genesis_committed")
	}
	result, err := scheduler.applier.Apply(ctx, genesisCommittedEvent{request: request})
	if err != nil {
		return "", err
	}
	return result.Value.(GenesisCommittedDecision), nil
}

func (event genesisCommittedEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	var phase string
	err := tx.QueryRowContext(ctx, `SELECT run_phase FROM runs WHERE run_id=?`, request.RunID).Scan(&phase)
	if store.IsNoRows(err) {
		return genesisCommittedNoMutation(GenesisCommittedIdentityFault), nil
	}
	if err != nil {
		return applier.Result{}, err
	}
	if request.RunID != request.BoundRunID {
		return genesisCommittedNoMutation(GenesisCommittedIdentityFault), nil
	}
	switch phase {
	case "established":
		return genesisCommittedNoMutation(GenesisCommittedDuplicate), nil
	case "created":
		return genesisCommittedNoMutation(GenesisCommittedUnexpected), nil
	case "create_authorized":
	default:
		return applier.Result{}, errors.New("scheduler: invalid durable run phase")
	}
	var currentGeneration, currentEpoch string
	err = tx.QueryRowContext(ctx, `SELECT l.generation_id,l.turn_epoch FROM leases l
		WHERE l.run_id=? AND l.lease_kind='worker' AND l.state='ACTIVE'`, request.RunID).Scan(&currentGeneration, &currentEpoch)
	if store.IsNoRows(err) {
		return genesisCommittedNoMutation(GenesisCommittedStale), nil
	}
	if err != nil {
		return applier.Result{}, err
	}
	if currentGeneration != request.GenerationID || unpad(currentEpoch) != request.TurnEpoch {
		return genesisCommittedNoMutation(GenesisCommittedStale), nil
	}
	result, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='established',updated_at=? WHERE run_id=? AND run_phase='create_authorized'`, request.At, request.RunID)
	if err != nil {
		return applier.Result{}, err
	}
	changed, _ := result.RowsAffected()
	if changed != 1 {
		return applier.Result{}, errors.New("scheduler: genesis_committed phase race")
	}
	return applier.Result{Value: GenesisCommittedEstablished}, nil
}

func genesisCommittedNoMutation(decision GenesisCommittedDecision) applier.Result {
	return applier.Result{Value: decision, NoMutation: true}
}
