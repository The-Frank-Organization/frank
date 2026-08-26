package f59

import (
	"context"
	"reflect"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
)

type outcomeEvent struct {
	runID        string
	generationID string
	request      OutcomeRequest
}

func (event outcomeEvent) RunID() string { return event.runID }

func (host *Host) RecordOutcome(ctx context.Context, channel ChannelIdentity, request OutcomeRequest) (Decision, error) {
	if host == nil || host.applier == nil || channel.GenerationID == "" || !validOutcomeShape(request) {
		return Decision{Kind: NoReply, Fault: true}, nil
	}
	runID, found, err := host.ticketRunID(ctx, request.TicketID)
	if err != nil {
		return Decision{}, err
	}
	if !found {
		return Decision{Kind: NoReply, Fault: true}, nil
	}
	result, err := host.applier.Apply(ctx, outcomeEvent{runID: runID, generationID: channel.GenerationID, request: request})
	if err != nil {
		return Decision{}, err
	}
	return result.Value.(Decision), nil
}

func (event outcomeEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	row, found, err := readOutcomeRow(ctx, tx, request.TicketID)
	if err != nil {
		return applier.Result{}, err
	}
	if !found {
		return noMutation(Decision{Kind: NoReply, Fault: true}), nil
	}
	current, err := currentEpoch(ctx, tx, row.RunID)
	if err != nil {
		return applier.Result{}, err
	}
	relation, err := compareCounters(request.TurnEpoch, current)
	if err != nil || relation > 0 {
		return noMutation(Decision{Kind: NoReply, Fault: true}), nil
	}
	if row.State == "OUTCOME_RECORDED" && outcomeEquivalent(row, request) {
		return noMutation(Decision{Kind: NoReply, Idempotent: true}), nil
	}
	lease, err := holdsBothLeases(ctx, tx, row.RunID, event.generationID, current)
	if err != nil {
		return applier.Result{}, err
	}
	if relation < 0 || !lease {
		return noMutation(Decision{Kind: NoReply, Dropped: true}), nil
	}
	if row.State == "OUTCOME_RECORDED" {
		return noMutation(Decision{Kind: NoReply}), nil
	}
	if row.State == "UNKNOWN_TOOL_OUTCOME" {
		return noMutation(Decision{Kind: NoReply, Fault: true}), nil
	}
	if row.State != "CONSUMED" {
		return noMutation(Decision{Kind: NoReply, Fault: true}), nil
	}
	expected := Identity{CanonicalToolName: row.Name, CanonicalArgsDigest: row.ArgsDigest, TurnEpoch: unpadCounter(row.Epoch)}
	if request.Outcome == Executed {
		if !reflect.DeepEqual(*request.InvocationIdentity, expected) {
			return noMutation(Decision{Kind: NoReply, Fault: true}), nil
		}
		identity := request.InvocationIdentity
		if _, err := tx.ExecContext(ctx, `INSERT INTO tool_calls
			(tool_call_id,run_id,turn_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,invocation_tool_name,invocation_args_digest,created_at,updated_at)
			VALUES(?,?,?,?,?,?,?,?,?,?,?)`, row.ToolCallID, row.RunID, row.TurnID, row.Epoch, "EXECUTED", row.Name, row.ArgsDigest,
			identity.CanonicalToolName, identity.CanonicalArgsDigest, row.IssuedAt, request.RecordedAt); err != nil {
			return applier.Result{}, err
		}
	} else {
		evidence := request.IntegrityEvidence
		if !reflect.DeepEqual(evidence.Expected, expected) || reflect.DeepEqual(evidence.Expected, evidence.Observed) || !validIdentity(evidence.Observed) {
			return noMutation(Decision{Kind: NoReply, Fault: true}), nil
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO tool_calls
			(tool_call_id,run_id,turn_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,expected_tool_name,expected_args_digest,expected_turn_epoch,
			 observed_tool_name,observed_args_digest,observed_turn_epoch,created_at,updated_at)
			VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, row.ToolCallID, row.RunID, row.TurnID, row.Epoch, "NOT_INVOKED_INTEGRITY_FAULT", row.Name, row.ArgsDigest,
			evidence.Expected.CanonicalToolName, evidence.Expected.CanonicalArgsDigest, padCounter(evidence.Expected.TurnEpoch), evidence.Observed.CanonicalToolName,
			evidence.Observed.CanonicalArgsDigest, padCounter(evidence.Observed.TurnEpoch), row.IssuedAt, request.RecordedAt); err != nil {
			return applier.Result{}, err
		}
	}
	if _, err := tx.ExecContext(ctx, `UPDATE tool_authorizations SET state='OUTCOME_RECORDED',outcome_ref=? WHERE ticket_id=? AND state='CONSUMED'`, row.ToolCallID, request.TicketID); err != nil {
		return applier.Result{}, err
	}
	return applier.Result{Value: Decision{Kind: NoReply}}, nil
}

type storedOutcome struct {
	RunID, TurnID, ToolCallID, State, Epoch, Name, ArgsDigest string
	IssuedAt, ConsumedAt                                      int64
	InvocationName, InvocationDigest                          *string
	ExpectedName, ExpectedDigest, ExpectedEpoch               *string
	ObservedName, ObservedDigest, ObservedEpoch               *string
}

func readOutcomeRow(ctx context.Context, tx *store.Tx, ticketID string) (storedOutcome, bool, error) {
	var row storedOutcome
	err := tx.QueryRowContext(ctx, `SELECT a.run_id,a.turn_id,a.tool_call_id,a.state,a.turn_epoch,a.canonical_tool_name,a.canonical_args_digest,a.issued_at,COALESCE(a.consumed_at,a.issued_at),
		c.invocation_tool_name,c.invocation_args_digest,c.expected_tool_name,c.expected_args_digest,c.expected_turn_epoch,c.observed_tool_name,c.observed_args_digest,c.observed_turn_epoch
		FROM tool_authorizations a LEFT JOIN tool_calls c ON c.tool_call_id=a.tool_call_id WHERE a.ticket_id=?`, ticketID).Scan(
		&row.RunID, &row.TurnID, &row.ToolCallID, &row.State, &row.Epoch, &row.Name, &row.ArgsDigest, &row.IssuedAt, &row.ConsumedAt,
		&row.InvocationName, &row.InvocationDigest, &row.ExpectedName, &row.ExpectedDigest, &row.ExpectedEpoch, &row.ObservedName, &row.ObservedDigest, &row.ObservedEpoch)
	if store.IsNoRows(err) {
		return storedOutcome{}, false, nil
	}
	return row, err == nil, err
}

func outcomeEquivalent(row storedOutcome, request OutcomeRequest) bool {
	if unpadCounter(row.Epoch) != request.TurnEpoch {
		return false
	}
	if request.Outcome == Executed && request.InvocationIdentity != nil {
		return pointerEquals(row.InvocationName, request.InvocationIdentity.CanonicalToolName) && pointerEquals(row.InvocationDigest, request.InvocationIdentity.CanonicalArgsDigest)
	}
	if request.Outcome == NotInvokedIntegrityFault && request.IntegrityEvidence != nil {
		evidence := request.IntegrityEvidence
		return nullableEquals(row.ExpectedName, evidence.Expected.CanonicalToolName) && nullableEquals(row.ExpectedDigest, evidence.Expected.CanonicalArgsDigest) && nullableEquals(row.ExpectedEpoch, padCounter(evidence.Expected.TurnEpoch)) &&
			nullableEquals(row.ObservedName, evidence.Observed.CanonicalToolName) && nullableEquals(row.ObservedDigest, evidence.Observed.CanonicalArgsDigest) && nullableEquals(row.ObservedEpoch, padCounter(evidence.Observed.TurnEpoch))
	}
	return false
}

func validOutcomeShape(request OutcomeRequest) bool {
	if request.TicketID == "" {
		return false
	}
	switch request.Outcome {
	case Executed:
		return request.InvocationIdentity != nil && request.IntegrityEvidence == nil && validIdentity(*request.InvocationIdentity)
	case NotInvokedIntegrityFault:
		return request.InvocationIdentity == nil && request.IntegrityEvidence != nil && validIdentity(request.IntegrityEvidence.Expected) && validIdentity(request.IntegrityEvidence.Observed)
	default:
		return false
	}
}

func validIdentity(identity Identity) bool {
	if identity.CanonicalToolName == "" || !validDigest(identity.CanonicalArgsDigest) {
		return false
	}
	_, err := compareCounters(identity.TurnEpoch, identity.TurnEpoch)
	return err == nil
}

func nullableEquals(value *string, expected string) bool { return pointerEquals(value, expected) }

func pointerEquals(value *string, expected string) bool { return value != nil && *value == expected }
