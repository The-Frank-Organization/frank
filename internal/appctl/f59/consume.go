package f59

import (
	"context"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

type consumeEvent struct {
	runID        string
	generationID string
	request      ConsumeRequest
}

func (event consumeEvent) RunID() string { return event.runID }

func (host *Host) Consume(ctx context.Context, channel ChannelIdentity, request ConsumeRequest) (Decision, error) {
	if host == nil || host.applier == nil || request.TicketID == "" || channel.GenerationID == "" || !validDigest(request.CanonicalArgsDigest) {
		return Decision{}, ErrInvalidRequest
	}
	runID, found, err := host.ticketRunID(ctx, request.TicketID)
	if err != nil {
		return Decision{}, err
	}
	if !found {
		return Decision{Kind: NoReply, Fault: true}, nil
	}
	result, err := host.applier.Apply(ctx, consumeEvent{runID: runID, generationID: channel.GenerationID, request: request})
	if err != nil {
		return Decision{}, err
	}
	decision, ok := result.Value.(Decision)
	if !ok {
		return Decision{}, ErrInvalidRequest
	}
	return decision, nil
}

func (event consumeEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	var state, storedEpoch, name, argsDigest, runID string
	err := tx.QueryRowContext(ctx, `SELECT state,turn_epoch,canonical_tool_name,canonical_args_digest,run_id FROM tool_authorizations WHERE ticket_id=?`, request.TicketID).
		Scan(&state, &storedEpoch, &name, &argsDigest, &runID)
	if store.IsNoRows(err) {
		return noMutation(Decision{Kind: NoReply, Fault: true}), nil
	}
	if err != nil {
		return applier.Result{}, err
	}
	current, err := currentEpoch(ctx, tx, runID)
	if err != nil {
		return applier.Result{}, err
	}
	relation, err := compareCounters(request.TurnEpoch, current)
	if err != nil {
		return noMutation(Decision{Kind: NoReply, Fault: true}), nil
	}
	if relation > 0 {
		return noMutation(Decision{Kind: NoReply, Fault: true}), nil
	}
	lease, err := holdsBothLeases(ctx, tx, runID, event.generationID, current)
	if err != nil {
		return applier.Result{}, err
	}
	if relation < 0 || !lease {
		return noMutation(Decision{Kind: StaleEpoch}), nil
	}
	if state != "ISSUED" {
		return noMutation(Decision{Kind: DuplicateConsume}), nil
	}
	if unpadCounter(storedEpoch) != request.TurnEpoch || name != request.CanonicalToolName || argsDigest != request.CanonicalArgsDigest {
		return noMutation(Decision{Kind: IdentityMismatch}), nil
	}
	if _, err := tx.ExecContext(ctx, `UPDATE tool_authorizations SET state='CONSUMED',consumed_at=? WHERE ticket_id=? AND state='ISSUED'`, request.ConsumedAt, request.TicketID); err != nil {
		return applier.Result{}, err
	}
	return applier.Result{Value: Decision{Kind: ConsumeOK}}, nil
}

func (host *Host) ticketRunID(ctx context.Context, ticketID string) (string, bool, error) {
	value, err := host.applier.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var runID string
		err := snapshot.QueryRowContext(ctx, `SELECT run_id FROM tool_authorizations WHERE ticket_id=?`, ticketID).Scan(&runID)
		if store.IsNoRows(err) {
			return "", nil
		}
		return runID, err
	}))
	if err != nil {
		return "", false, err
	}
	runID := value.(string)
	return runID, runID != "", nil
}

func compareCounters(left, right string) (int, error) {
	l, err := appipc.ParseCounter(left)
	if err != nil {
		return 0, err
	}
	r, err := appipc.ParseCounter(right)
	if err != nil {
		return 0, err
	}
	if l < r {
		return -1, nil
	}
	if l > r {
		return 1, nil
	}
	return 0, nil
}
