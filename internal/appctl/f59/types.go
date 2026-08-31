// Package f59 implements the app-control plane's durable one-shot tool ticket.
package f59

import (
	"context"
	"errors"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/manifest"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

var (
	ErrInvalidDescriptor = errors.New("f59: invalid effect descriptor")
	ErrInvalidRequest    = errors.New("f59: invalid request")
	ErrLeaseFaultUnwired = errors.New("f59: lease-invalid retirement is not wired")
)

type Kind string

const (
	TicketGranted    Kind = "ticket_granted"
	AuthorizeReject  Kind = "authorize_reject"
	StaleEpoch       Kind = "STALE_EPOCH"
	DeniedAboveSet   Kind = "DENIED_ABOVE_SET"
	DuplicateRequest Kind = "DUPLICATE_REQUEST"
	IdentityMismatch Kind = "IDENTITY_MISMATCH"
	ConsumeOK        Kind = "consume_ok"
	DuplicateConsume Kind = "DUPLICATE_CONSUME"
	NoReply          Kind = "NO_REPLY"
)

type RejectReason string

const (
	RunNotAdmitted      RejectReason = "run_not_admitted"
	TurnInactive        RejectReason = "turn_inactive"
	LeaseInvalid        RejectReason = "lease_invalid"
	TurnBudgetExhausted RejectReason = "turn_budget_exhausted"
)

type Decision struct {
	Kind       Kind
	Reason     RejectReason
	TicketID   string
	Fault      bool
	Dropped    bool
	Idempotent bool
}

type Operands struct {
	CanonicalResource *string
	CWD               *string
}

type IssueRequest struct {
	RunID               string
	TurnID              string
	TurnEpoch           string
	ToolCallID          string
	CanonicalToolName   string
	CanonicalArgsDigest string
	GenerationID        string
	Operands            Operands
	IssuedAt            int64
}

type ConsumeRequest struct {
	TicketID            string
	TurnEpoch           string
	CanonicalToolName   string
	CanonicalArgsDigest string
	ConsumedAt          int64
}

// ChannelIdentity is assigned by the authenticated CTRL-W receiver. It is
// deliberately separate from worker-authored request bodies.
type ChannelIdentity struct{ GenerationID string }

type Identity struct {
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
	TurnEpoch           string `json:"turn_epoch"`
}

type IntegrityEvidence struct {
	Expected Identity `json:"expected_identity"`
	Observed Identity `json:"observed_identity"`
}

type Outcome string

const (
	Executed                 Outcome = "executed"
	NotInvokedIntegrityFault Outcome = "not_invoked_integrity_fault"
)

type OutcomeRequest struct {
	TicketID           string
	TurnEpoch          string
	Outcome            Outcome
	InvocationIdentity *Identity
	IntegrityEvidence  *IntegrityEvidence
	RecordedAt         int64
}

type Config struct {
	Gate        manifest.Gate
	NewTicketID func() string
	// LeaseInvalid runs the T8-owned complete retirement effects in the same
	// transaction as the optional VOID/lease_invalid row.
	LeaseInvalid func(context.Context, *store.Tx, LeaseFault) error
}

type LeaseFault struct {
	RunID, TurnID, GenerationID, TurnEpoch string
	At                                     int64
}

type Host struct {
	applier *applier.Host
	config  Config
}

func New(host *applier.Host, config Config) *Host {
	return &Host{applier: host, config: config}
}
