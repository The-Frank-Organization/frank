// Package executor implements the worker's single F59 tool settlement point.
package executor

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/jackli/frank/internal/worker/jcs"
)

// Code is a stable attempt disposition returned by the executor.
type Code string

const (
	CodeAuthorizeRejected Code = "AUTHORIZE_REJECT"
	CodeStaleEpoch        Code = "STALE_EPOCH"
	CodeDeniedAboveSet    Code = "DENIED_ABOVE_SET"
	CodeDuplicateRequest  Code = "DUPLICATE_REQUEST"
	CodeIdentityMismatch  Code = "IDENTITY_MISMATCH"
	CodeDuplicateConsume  Code = "DUPLICATE_CONSUME"
	CodeProtocolFault     Code = "PROTOCOL_FAULT"
)

// Error is a typed, fail-closed settlement error.
type Error struct {
	Code   Code
	Detail string
	Cause  error
}

func (err *Error) Error() string {
	if err.Detail != "" {
		return fmt.Sprintf("executor: %s: %s", err.Code, err.Detail)
	}
	return fmt.Sprintf("executor: %s", err.Code)
}

func (err *Error) Unwrap() error { return err.Cause }

// ArgumentSource supplies a fresh snapshot at each F84 derivation point. The
// implementation deliberately does not trust the snapshot taken at Prepare.
type ArgumentSource interface {
	Snapshot() []byte
}

// Identity is the exact three-field tool invocation identity.
type Identity struct {
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
	TurnEpoch           uint64 `json:"turn_epoch"`
}

// FullIdentity adds the run/turn/call keys frozen into an issued ticket.
type FullIdentity struct {
	RunID      string `json:"run_id"`
	TurnID     string `json:"turn_id"`
	ToolCallID string `json:"tool_call_id"`
	Identity
}

// TicketState names the m-10 states consumed by worker boundary fixtures.
// The worker never mutates these states itself.
type TicketState string

const (
	TicketIssued             TicketState = "ISSUED"
	TicketConsumed           TicketState = "CONSUMED"
	TicketOutcomeRecorded    TicketState = "OUTCOME_RECORDED"
	TicketVoid               TicketState = "VOID"
	TicketUnknownToolOutcome TicketState = "UNKNOWN_TOOL_OUTCOME"
)

// PreparedCall is inert work plus its construction-time identity. It exposes
// no invocation operation; only Executor can advance it through authority.
type PreparedCall struct {
	identity FullIdentity
	source   ArgumentSource
}

// Prepare freezes identity #1 over a canonical copy of the complete arguments.
func Prepare(runID, turnID, toolCallID, canonicalToolName string, turnEpoch uint64, source ArgumentSource) (*PreparedCall, error) {
	if runID == "" || turnID == "" || toolCallID == "" || canonicalToolName == "" {
		return nil, errors.New("executor: prepared call identity contains an empty field")
	}
	if source == nil {
		return nil, errors.New("executor: argument source is absent")
	}
	identity, _, err := deriveIdentity(canonicalToolName, turnEpoch, source)
	if err != nil {
		return nil, fmt.Errorf("executor: freeze arguments: %w", err)
	}
	return &PreparedCall{
		identity: FullIdentity{
			RunID:      runID,
			TurnID:     turnID,
			ToolCallID: toolCallID,
			Identity:   identity,
		},
		source: source,
	}, nil
}

// FrozenIdentity returns construction-time identity #1 by value.
func (call *PreparedCall) FrozenIdentity() FullIdentity { return call.identity }

// AuthorizeCode is the closed authority issue result domain.
type AuthorizeCode string

const (
	AuthorizeGranted          AuthorizeCode = "ticket_granted"
	AuthorizeRejected         AuthorizeCode = "authorize_reject"
	AuthorizeStaleEpoch       AuthorizeCode = "STALE_EPOCH"
	AuthorizeDeniedAboveSet   AuthorizeCode = "DENIED_ABOVE_SET"
	AuthorizeDuplicateRequest AuthorizeCode = "DUPLICATE_REQUEST"
	AuthorizeIdentityMismatch AuthorizeCode = "IDENTITY_MISMATCH"
)

// AuthorizeRejectReason is the closed lifecycle-rejection domain.
type AuthorizeRejectReason string

const (
	RejectRunNotAdmitted      AuthorizeRejectReason = "run_not_admitted"
	RejectTurnInactive        AuthorizeRejectReason = "turn_inactive"
	RejectLeaseInvalid        AuthorizeRejectReason = "lease_invalid"
	RejectTurnBudgetExhausted AuthorizeRejectReason = "turn_budget_exhausted"
)

var lifecycleRejectReasons = [...]AuthorizeRejectReason{
	RejectRunNotAdmitted,
	RejectTurnInactive,
	RejectLeaseInvalid,
	RejectTurnBudgetExhausted,
}

// LifecycleRejectReasons returns the closed four-token lifecycle domain.
func LifecycleRejectReasons() []AuthorizeRejectReason {
	return append([]AuthorizeRejectReason(nil), lifecycleRejectReasons[:]...)
}

type AuthorizeRequest struct {
	Identity FullIdentity `json:"identity"`
}

type AuthorizeReply struct {
	Code         AuthorizeCode         `json:"code"`
	TicketID     string                `json:"ticket_id,omitempty"`
	RejectReason AuthorizeRejectReason `json:"reason,omitempty"`
}

func (reply AuthorizeReply) Validate() error {
	switch reply.Code {
	case AuthorizeGranted:
		if reply.TicketID == "" || reply.RejectReason != "" {
			return errors.New("ticket_granted requires only ticket_id")
		}
	case AuthorizeRejected:
		if reply.TicketID != "" || !validRejectReason(reply.RejectReason) {
			return errors.New("authorize_reject requires one lifecycle reason")
		}
	case AuthorizeStaleEpoch, AuthorizeDeniedAboveSet, AuthorizeDuplicateRequest, AuthorizeIdentityMismatch:
		if reply.TicketID != "" || reply.RejectReason != "" {
			return fmt.Errorf("%s carries unexpected fields", reply.Code)
		}
	default:
		return fmt.Errorf("unknown authorize code %q", reply.Code)
	}
	return nil
}

func validRejectReason(reason AuthorizeRejectReason) bool {
	for _, candidate := range lifecycleRejectReasons {
		if reason == candidate {
			return true
		}
	}
	return false
}

// ConsumeCode is the closed ticket-consumption result domain.
type ConsumeCode string

const (
	ConsumeOK               ConsumeCode = "consume_ok"
	ConsumeStaleEpoch       ConsumeCode = "STALE_EPOCH"
	ConsumeDuplicate        ConsumeCode = "DUPLICATE_CONSUME"
	ConsumeIdentityMismatch ConsumeCode = "IDENTITY_MISMATCH"
)

type ConsumeRequest struct {
	TicketID            string `json:"ticket_id"`
	TurnEpoch           uint64 `json:"turn_epoch"`
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
}

type ConsumeReply struct {
	Code ConsumeCode `json:"code"`
}

func (reply ConsumeReply) Validate() error {
	switch reply.Code {
	case ConsumeOK, ConsumeStaleEpoch, ConsumeDuplicate, ConsumeIdentityMismatch:
		return nil
	default:
		return fmt.Errorf("unknown consume code %q", reply.Code)
	}
}

type Outcome string

const (
	OutcomeExecuted                 Outcome = "executed"
	OutcomeNotInvokedIntegrityFault Outcome = "not_invoked_integrity_fault"
)

type IntegrityEvidence struct {
	Expected Identity `json:"expected_identity"`
	Observed Identity `json:"observed_identity"`
}

type OutcomeRecord struct {
	TicketID           string             `json:"ticket_id"`
	Outcome            Outcome            `json:"outcome"`
	InvocationIdentity *Identity          `json:"invocation_identity,omitempty"`
	IntegrityEvidence  *IntegrityEvidence `json:"integrity_evidence,omitempty"`
}

func (record OutcomeRecord) Validate() error {
	if record.TicketID == "" {
		return errors.New("outcome ticket_id is empty")
	}
	switch record.Outcome {
	case OutcomeExecuted:
		if record.InvocationIdentity == nil || record.IntegrityEvidence != nil {
			return errors.New("executed requires only invocation_identity")
		}
	case OutcomeNotInvokedIntegrityFault:
		if record.InvocationIdentity != nil || record.IntegrityEvidence == nil {
			return errors.New("not_invoked_integrity_fault requires only integrity_evidence")
		}
		if record.IntegrityEvidence.Expected == record.IntegrityEvidence.Observed {
			return errors.New("integrity identities must differ")
		}
	default:
		return fmt.Errorf("unknown outcome %q", record.Outcome)
	}
	return nil
}

// Authority is m-10's synchronous authorize/consume seam and asynchronous,
// no-reply outcome-send seam. RecordOutcome returns only a local send failure.
type Authority interface {
	Authorize(context.Context, AuthorizeRequest) (AuthorizeReply, error)
	Consume(context.Context, ConsumeRequest) (ConsumeReply, error)
	RecordOutcome(context.Context, OutcomeRecord) error
}

type Invocation struct {
	Identity  Identity
	Arguments []byte
}

type Invoker interface {
	Invoke(context.Context, Invocation) (any, error)
}

type Result struct {
	TicketID           string
	Outcome            Outcome
	InvocationIdentity Identity
	Value              any
}

// Executor owns the only path from an inert parsed call to tool invocation.
// Its mutex serializes the entire settlement, not merely backend execution.
type Executor struct {
	authority Authority
	invoker   Invoker
	mu        sync.Mutex
}

func New(authority Authority, invoker Invoker) (*Executor, error) {
	if authority == nil || invoker == nil {
		return nil, errors.New("executor: authority and invoker are required")
	}
	return &Executor{authority: authority, invoker: invoker}, nil
}

func (executor *Executor) Execute(ctx context.Context, call *PreparedCall) (Result, error) {
	executor.mu.Lock()
	defer executor.mu.Unlock()

	if call == nil {
		return Result{}, &Error{Code: CodeProtocolFault, Detail: "prepared call is absent"}
	}
	if err := ctx.Err(); err != nil {
		return Result{}, err
	}
	authorizeReply, err := executor.authority.Authorize(ctx, AuthorizeRequest{Identity: call.identity})
	if err != nil {
		return Result{}, err
	}
	if err := authorizeReply.Validate(); err != nil {
		return Result{}, &Error{Code: CodeProtocolFault, Detail: err.Error(), Cause: err}
	}
	if authorizeReply.Code != AuthorizeGranted {
		return Result{}, authorizeError(authorizeReply)
	}
	result := Result{TicketID: authorizeReply.TicketID}
	if err := ctx.Err(); err != nil {
		return result, err
	}

	// F84 derivation #2 is fresh and its exact four-field identity goes to the
	// authority. The authority, not stale worker memory, performs the match.
	preConsume, _, err := deriveIdentity(call.identity.CanonicalToolName, call.identity.TurnEpoch, call.source)
	if err != nil {
		return result, &Error{Code: CodeIdentityMismatch, Detail: "pre-consume arguments are no longer valid", Cause: err}
	}
	consumeReply, err := executor.authority.Consume(ctx, ConsumeRequest{
		TicketID:            result.TicketID,
		TurnEpoch:           preConsume.TurnEpoch,
		CanonicalToolName:   preConsume.CanonicalToolName,
		CanonicalArgsDigest: preConsume.CanonicalArgsDigest,
	})
	if err != nil {
		return result, err
	}
	if err := consumeReply.Validate(); err != nil {
		return result, &Error{Code: CodeProtocolFault, Detail: err.Error(), Cause: err}
	}
	if consumeReply.Code != ConsumeOK {
		return result, consumeError(consumeReply)
	}
	if err := ctx.Err(); err != nil {
		return result, err
	}

	// F84 derivation #3 is fresh after consume and immediately before invoke.
	preInvoke, canonicalArguments, err := deriveIdentity(call.identity.CanonicalToolName, call.identity.TurnEpoch, call.source)
	if err != nil {
		return result, &Error{Code: CodeIdentityMismatch, Detail: "pre-invocation arguments are no longer valid", Cause: err}
	}
	if preInvoke != call.identity.Identity {
		record := OutcomeRecord{
			TicketID: result.TicketID,
			Outcome:  OutcomeNotInvokedIntegrityFault,
			IntegrityEvidence: &IntegrityEvidence{
				Expected: call.identity.Identity,
				Observed: preInvoke,
			},
		}
		if err := executor.authority.RecordOutcome(context.WithoutCancel(ctx), record); err != nil {
			return result, err
		}
		result.Outcome = OutcomeNotInvokedIntegrityFault
		return result, nil
	}

	invocation := Invocation{Identity: preInvoke, Arguments: canonicalArguments}
	value, invocationErr := executor.invoker.Invoke(ctx, invocation)
	record := OutcomeRecord{
		TicketID:           result.TicketID,
		Outcome:            OutcomeExecuted,
		InvocationIdentity: &preInvoke,
	}
	recordErr := executor.authority.RecordOutcome(context.WithoutCancel(ctx), record)
	result.Outcome = OutcomeExecuted
	result.InvocationIdentity = preInvoke
	result.Value = value
	return result, errors.Join(invocationErr, recordErr)
}

func deriveIdentity(name string, epoch uint64, source ArgumentSource) (Identity, []byte, error) {
	snapshot := append([]byte(nil), source.Snapshot()...)
	canonical, err := jcs.Canonicalize(snapshot)
	if err != nil {
		return Identity{}, nil, err
	}
	digest, err := jcs.Digest(canonical)
	if err != nil {
		return Identity{}, nil, err
	}
	return Identity{CanonicalToolName: name, CanonicalArgsDigest: digest, TurnEpoch: epoch}, canonical, nil
}

func authorizeError(reply AuthorizeReply) error {
	switch reply.Code {
	case AuthorizeRejected:
		return &Error{Code: CodeAuthorizeRejected, Detail: string(reply.RejectReason)}
	case AuthorizeStaleEpoch:
		return &Error{Code: CodeStaleEpoch}
	case AuthorizeDeniedAboveSet:
		return &Error{Code: CodeDeniedAboveSet}
	case AuthorizeDuplicateRequest:
		return &Error{Code: CodeDuplicateRequest}
	case AuthorizeIdentityMismatch:
		return &Error{Code: CodeIdentityMismatch}
	default:
		return &Error{Code: CodeProtocolFault, Detail: "unhandled authorize reply"}
	}
}

func consumeError(reply ConsumeReply) error {
	switch reply.Code {
	case ConsumeStaleEpoch:
		return &Error{Code: CodeStaleEpoch}
	case ConsumeDuplicate:
		return &Error{Code: CodeDuplicateConsume}
	case ConsumeIdentityMismatch:
		return &Error{Code: CodeIdentityMismatch}
	default:
		return &Error{Code: CodeProtocolFault, Detail: "unhandled consume reply"}
	}
}
