package f59

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/manifest"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

type issueEvent struct {
	host    *Host
	request IssueRequest
}

func (event issueEvent) RunID() string { return event.request.RunID }

func (host *Host) Issue(ctx context.Context, request IssueRequest) (Decision, error) {
	if host == nil || host.applier == nil || request.RunID == "" || request.TurnID == "" || request.ToolCallID == "" || request.GenerationID == "" || !validDigest(request.CanonicalArgsDigest) {
		return Decision{}, ErrInvalidRequest
	}
	if _, err := compareCounters(request.TurnEpoch, request.TurnEpoch); err != nil {
		return Decision{}, ErrInvalidRequest
	}
	result, err := host.applier.Apply(ctx, issueEvent{host: host, request: request})
	if err != nil {
		return Decision{}, err
	}
	decision, ok := result.Value.(Decision)
	if !ok {
		return Decision{}, errors.New("f59: invalid issue result")
	}
	return decision, nil
}

func (event issueEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	if replay, found, err := readReplay(ctx, tx, request); err != nil {
		return applier.Result{}, err
	} else if found {
		return noMutation(replay), nil
	}

	runState, frozen, found, err := loadRun(ctx, tx, request.RunID, event.host.config.Gate)
	if err != nil {
		return applier.Result{}, err
	}
	if !found {
		return noMutation(Decision{Kind: AuthorizeReject, Reason: RunNotAdmitted}), nil
	}
	currentEpoch, err := currentEpoch(ctx, tx, request.RunID)
	if err != nil {
		return applier.Result{}, err
	}
	if request.TurnEpoch != currentEpoch {
		return noMutation(Decision{Kind: StaleEpoch}), nil
	}

	turnState, turnExists, err := readTurnState(ctx, tx, request.RunID, request.TurnID)
	if err != nil {
		return applier.Result{}, err
	}
	count, err := authorizationCount(ctx, tx, request.RunID, request.TurnID)
	if err != nil {
		return applier.Result{}, err
	}
	if runState != "ADMITTED" && runState != "ACTIVE" {
		return event.classifiedDenial(ctx, tx, turnExists, count, "run_not_admitted", Decision{Kind: AuthorizeReject, Reason: RunNotAdmitted})
	}
	if !turnExists || turnState != "ACTIVE" {
		return event.classifiedDenial(ctx, tx, turnExists, count, "turn_inactive", Decision{Kind: AuthorizeReject, Reason: TurnInactive})
	}
	validLease, err := holdsBothLeases(ctx, tx, request.RunID, request.GenerationID, request.TurnEpoch)
	if err != nil {
		return applier.Result{}, err
	}
	if !validLease {
		if event.host.config.LeaseInvalid == nil {
			return applier.Result{}, ErrLeaseFaultUnwired
		}
		if count < appipc.MaxToolCallsPerTurn {
			if _, err := event.insertDenial(ctx, tx, "lease_invalid"); err != nil {
				return applier.Result{}, err
			}
		}
		fault := LeaseFault{RunID: request.RunID, TurnID: request.TurnID, GenerationID: request.GenerationID, TurnEpoch: request.TurnEpoch, At: request.IssuedAt}
		if err := event.host.config.LeaseInvalid(ctx, tx, fault); err != nil {
			return applier.Result{}, err
		}
		return applier.Result{Value: Decision{Kind: AuthorizeReject, Reason: LeaseInvalid}}, nil
	}
	if count >= appipc.MaxToolCallsPerTurn {
		return noMutation(Decision{Kind: AuthorizeReject, Reason: TurnBudgetExhausted}), nil
	}
	if err := event.host.config.Gate.Authorize(frozen, request.CanonicalToolName); err != nil {
		return event.classifiedDenial(ctx, tx, true, count, "denied_above_set", Decision{Kind: DeniedAboveSet})
	}
	descriptor, encoded, err := BuildDescriptor(frozen.Manifest, request.CanonicalToolName, request.CanonicalArgsDigest, request.Operands)
	if err != nil {
		return event.classifiedDenial(ctx, tx, true, count, "denied_above_set", Decision{Kind: DeniedAboveSet})
	}
	ticketID, err := event.host.nextTicketID()
	if err != nil {
		return applier.Result{}, err
	}
	if err := insertAuthorization(ctx, tx, request, ticketID, "ISSUED", nil, encoded, descriptor); err != nil {
		return applier.Result{}, err
	}
	return applier.Result{Value: Decision{Kind: TicketGranted, TicketID: ticketID}}, nil
}

func (event issueEvent) classifiedDenial(ctx context.Context, tx *store.Tx, turnExists bool, count int, reason string, decision Decision) (applier.Result, error) {
	if !turnExists || count >= appipc.MaxToolCallsPerTurn {
		return noMutation(decision), nil
	}
	_, err := event.insertDenial(ctx, tx, reason)
	if err != nil {
		return applier.Result{}, err
	}
	return applier.Result{Value: decision}, nil
}

func (event issueEvent) insertDenial(ctx context.Context, tx *store.Tx, reason string) (string, error) {
	ticketID, err := event.host.nextTicketID()
	if err != nil {
		return "", err
	}
	descriptor := EffectDescriptor{Action: event.request.CanonicalToolName, CanonicalArgsDigest: event.request.CanonicalArgsDigest,
		CanonicalResource: cloneString(event.request.Operands.CanonicalResource), CWD: cloneString(event.request.Operands.CWD)}
	encoded, err := appipc.MarshalJCS(map[string]any{"action": event.request.CanonicalToolName, "canonical_args_digest": event.request.CanonicalArgsDigest,
		"canonical_resource": event.request.Operands.CanonicalResource, "cwd": event.request.Operands.CWD})
	if err != nil {
		return "", err
	}
	if err := insertAuthorization(ctx, tx, event.request, ticketID, "VOID", &reason, encoded, descriptor); err != nil {
		return "", err
	}
	return ticketID, nil
}

func (host *Host) nextTicketID() (string, error) {
	if host.config.NewTicketID == nil {
		return "", errors.New("f59: ticket id source is required")
	}
	id := host.config.NewTicketID()
	if id == "" {
		return "", errors.New("f59: ticket id source returned empty id")
	}
	return id, nil
}

func insertAuthorization(ctx context.Context, tx *store.Tx, request IssueRequest, ticketID, state string, voidReason *string, encoded []byte, descriptor EffectDescriptor) error {
	var resource, cwd any
	if descriptor.CanonicalResource != nil {
		resource = *descriptor.CanonicalResource
	}
	if descriptor.CWD != nil {
		cwd = *descriptor.CWD
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO tool_authorizations
		(ticket_id,run_id,turn_id,tool_call_id,turn_epoch,state,void_reason,canonical_tool_name,canonical_args_digest,canonical_resource,cwd,effect_descriptor,issued_at)
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`, ticketID, request.RunID, request.TurnID, request.ToolCallID, padCounter(request.TurnEpoch), state,
		voidReason, request.CanonicalToolName, request.CanonicalArgsDigest, resource, cwd, encoded, request.IssuedAt); err != nil {
		return fmt.Errorf("f59: insert authorization: %w", err)
	}
	return nil
}

func readReplay(ctx context.Context, tx *store.Tx, request IssueRequest) (Decision, bool, error) {
	var ticketID, state, name, argsDigest, epoch string
	var reason, resource, cwd *string
	err := tx.QueryRowContext(ctx, `SELECT ticket_id,state,void_reason,canonical_tool_name,canonical_args_digest,turn_epoch,canonical_resource,cwd
		FROM tool_authorizations WHERE run_id=? AND turn_id=? AND tool_call_id=?`, request.RunID, request.TurnID, request.ToolCallID).
		Scan(&ticketID, &state, &reason, &name, &argsDigest, &epoch, &resource, &cwd)
	if store.IsNoRows(err) {
		return Decision{}, false, nil
	}
	if err != nil {
		return Decision{}, false, err
	}
	if name != request.CanonicalToolName || argsDigest != request.CanonicalArgsDigest || unpadCounter(epoch) != request.TurnEpoch || !equalNullable(resource, request.Operands.CanonicalResource) || !equalNullable(cwd, request.Operands.CWD) {
		return Decision{Kind: IdentityMismatch}, true, nil
	}
	switch state {
	case "ISSUED":
		return Decision{Kind: TicketGranted, TicketID: ticketID}, true, nil
	case "VOID":
		if reason == nil {
			return Decision{}, false, fmt.Errorf("f59: stored VOID row has no reason")
		}
		switch *reason {
		case "run_not_admitted":
			return Decision{Kind: AuthorizeReject, Reason: RunNotAdmitted}, true, nil
		case "turn_inactive", "expired":
			return Decision{Kind: AuthorizeReject, Reason: TurnInactive}, true, nil
		case "lease_invalid":
			return Decision{Kind: AuthorizeReject, Reason: LeaseInvalid}, true, nil
		case "denied_above_set":
			return Decision{Kind: DeniedAboveSet}, true, nil
		default:
			return Decision{}, false, fmt.Errorf("f59: invalid stored VOID reason %q", *reason)
		}
	case "CONSUMED", "OUTCOME_RECORDED", "UNKNOWN_TOOL_OUTCOME":
		return Decision{Kind: DuplicateRequest}, true, nil
	default:
		return Decision{}, false, fmt.Errorf("f59: invalid stored ticket state %q", state)
	}
}

func loadRun(ctx context.Context, tx *store.Tx, runID string, gate manifest.Gate) (string, manifest.Frozen, bool, error) {
	var state, digest string
	var encoded []byte
	err := tx.QueryRowContext(ctx, `SELECT state,manifest_bytes,run_manifest_digest FROM runs WHERE run_id=?`, runID).Scan(&state, &encoded, &digest)
	if store.IsNoRows(err) {
		return "", manifest.Frozen{}, false, nil
	}
	if err != nil {
		return "", manifest.Frozen{}, false, err
	}
	frozen, err := manifest.DecodeFrozen(encoded, digest, gate)
	return state, frozen, true, err
}

func readTurnState(ctx context.Context, tx *store.Tx, runID, turnID string) (string, bool, error) {
	var state string
	err := tx.QueryRowContext(ctx, `SELECT state FROM turns WHERE run_id=? AND turn_id=?`, runID, turnID).Scan(&state)
	if store.IsNoRows(err) {
		return "", false, nil
	}
	return state, err == nil, err
}

func currentEpoch(ctx context.Context, tx *store.Tx, runID string) (string, error) {
	var stored string
	if err := tx.QueryRowContext(ctx, `SELECT turn_epoch FROM epochs WHERE run_id=?`, runID).Scan(&stored); err != nil {
		return "", err
	}
	return unpadCounter(stored), nil
}

func authorizationCount(ctx context.Context, tx *store.Tx, runID, turnID string) (int, error) {
	var count int
	err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM tool_authorizations WHERE run_id=? AND turn_id=?`, runID, turnID).Scan(&count)
	return count, err
}

func holdsBothLeases(ctx context.Context, tx *store.Tx, runID, generationID, epoch string) (bool, error) {
	var count int
	err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM leases WHERE run_id=? AND generation_id=? AND turn_epoch=? AND state='ACTIVE' AND lease_kind IN ('worker','turn')`, runID, generationID, padCounter(epoch)).Scan(&count)
	return count == 2, err
}

func noMutation(decision Decision) applier.Result {
	return applier.Result{Value: decision, NoMutation: true}
}

func equalNullable(stored *string, request *string) bool {
	if request == nil {
		return stored == nil
	}
	return stored != nil && *stored == *request
}

func padCounter(value string) string {
	padded, err := appipc.PadCounter(value)
	if err != nil {
		return ""
	}
	return padded
}

func unpadCounter(value string) string {
	wire, err := appipc.UnpadCounter(value)
	if err != nil {
		return ""
	}
	return wire
}
