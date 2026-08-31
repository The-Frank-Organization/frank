// Package provider consumes the normalized m-8 attempt stream without owning
// provider translation, credentials, routing, or retry policy.
package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/worker/wire"
)

type Disposition string

const (
	Completed       Disposition = "completed"
	TransportFailed Disposition = "transport_failed"
	EgressDenied    Disposition = "egress_denied"
	RejectedLocal   Disposition = "rejected_local"
	CancelledPre    Disposition = "cancelled_pre_transport"
	CancelledPost   Disposition = "cancelled_post_invocation"
	StreamLost      Disposition = "stream_lost"
	StaleEpoch      Disposition = "STALE_EPOCH"
	EpochAhead      Disposition = "EPOCH_AHEAD"
)

type RejectReason string

const (
	MalformedRequest       RejectReason = "malformed_request"
	LaneCapabilityMismatch RejectReason = "lane_capability_mismatch"
	ReplayScopeViolation   RejectReason = "replay_scope_violation"
	InternalIntegrityFault RejectReason = "internal_integrity_fault"
)

type StreamEnd string

const (
	StreamCompleted StreamEnd = "stream_completed"
	StreamFailed    StreamEnd = "stream_failed"
	StreamCancelled StreamEnd = "stream_cancelled"
	StreamLostEnd   StreamEnd = "stream_lost"
)

type Request struct {
	AttemptID     string
	TurnID        string
	TurnEpoch     string
	ProviderLane  string
	OpaqueRequest json.RawMessage
}

type Event struct {
	Kind   string
	Opaque json.RawMessage
}

const NormalizedEventSchemaV2 = "m8.provider_event.v2"

type ToolCall struct {
	ID            string
	CanonicalName string
	Arguments     json.RawMessage
}

// ToolCallFromEvents applies the uniform v2 normalized-event gate and lowers
// the single complete tool_call_end, if present. Opaque non-event items remain
// opaque and are not rejected merely because their bytes are not JSON.
func ToolCallFromEvents(events []Event) (*ToolCall, error) {
	var result *ToolCall
	for _, event := range events {
		var header struct {
			Schema string          `json:"schema"`
			Kind   string          `json:"kind"`
			Type   string          `json:"type"`
			Body   json.RawMessage `json:"body"`
		}
		if err := json.Unmarshal(event.Opaque, &header); err != nil || header.Schema == "" {
			continue
		}
		if header.Schema != NormalizedEventSchemaV2 {
			return nil, fmt.Errorf("provider: unsupported normalized event schema %q", header.Schema)
		}
		kind := header.Kind
		if kind == "" {
			kind = header.Type
		}
		if kind != "tool_call_end" {
			continue
		}
		if result != nil {
			return nil, errors.New("provider: multiple complete tool calls in one turn")
		}
		var body struct {
			ToolCallID string          `json:"tool_call_id"`
			Name       string          `json:"name"`
			Arguments  json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(header.Body, &body); err != nil || body.ToolCallID == "" || body.Name == "" || len(body.Arguments) == 0 {
			return nil, errors.New("provider: malformed tool_call_end")
		}
		result = &ToolCall{ID: body.ToolCallID, CanonicalName: strings.ToLower(body.Name), Arguments: append(json.RawMessage(nil), body.Arguments...)}
	}
	return result, nil
}

type Outcome struct {
	Disposition Disposition
	Events      []Event
	StreamEnd   *StreamEnd
	WireStarted bool
}

type Gate interface {
	AttemptOpen(context.Context, Request) error
	RecordStreamEnd(context.Context, string, StreamEnd) error
}

type Connector interface {
	Attempt(context.Context, Request) (Disposition, []json.RawMessage, error)
	Cancel(context.Context, string, string) (Disposition, error)
}

type Cycle struct {
	gate      Gate
	connector Connector
	epoch     string
}

func New(gate Gate, connector Connector, epoch string) (*Cycle, error) {
	if gate == nil || connector == nil {
		return nil, errors.New("provider: gate and connector are required")
	}
	if _, err := wire.ParseCounter(epoch); err != nil {
		return nil, fmt.Errorf("provider: turn epoch: %w", err)
	}
	return &Cycle{gate: gate, connector: connector, epoch: epoch}, nil
}

func (cycle *Cycle) Run(ctx context.Context, request Request) (Outcome, error) {
	requestEpoch, err := wire.ParseCounter(request.TurnEpoch)
	if err != nil {
		return Outcome{}, fmt.Errorf("provider: request turn epoch: %w", err)
	}
	if request.AttemptID == "" || request.TurnID == "" || request.ProviderLane == "" {
		return Outcome{}, errors.New("provider: attempt id, turn id, and provider lane are required")
	}
	cycleEpoch, _ := wire.ParseCounter(cycle.epoch)
	if requestEpoch < cycleEpoch {
		return Outcome{Disposition: StaleEpoch}, nil
	}
	if requestEpoch > cycleEpoch {
		return Outcome{Disposition: EpochAhead}, nil
	}
	if err := cycle.gate.AttemptOpen(ctx, request); err != nil {
		return Outcome{}, err
	}
	disposition, items, err := cycle.connector.Attempt(ctx, request)
	if err != nil {
		return Outcome{}, err
	}
	disposition, err = normalizeDisposition(disposition)
	if err != nil {
		return Outcome{}, err
	}
	outcome := Outcome{Disposition: disposition}
	for _, item := range items {
		outcome.Events = append(outcome.Events, Event{Kind: "response_item", Opaque: append(json.RawMessage(nil), item...)})
	}
	var end StreamEnd
	switch disposition {
	case Completed:
		outcome.WireStarted, end = true, StreamCompleted
	case TransportFailed:
		outcome.WireStarted, end = true, StreamFailed
	case CancelledPost:
		outcome.WireStarted, end = true, StreamCancelled
	case StreamLost:
		outcome.WireStarted, end = true, StreamLostEnd
	case EgressDenied, RejectedLocal, CancelledPre, StaleEpoch, EpochAhead:
		return outcome, nil
	default:
		return Outcome{}, errors.New("provider: unknown attempt disposition")
	}
	outcome.StreamEnd = &end
	if err := cycle.gate.RecordStreamEnd(ctx, request.AttemptID, end); err != nil {
		return outcome, err
	}
	return outcome, nil
}

func normalizeDisposition(disposition Disposition) (Disposition, error) {
	switch disposition {
	case Completed, "sent_completed":
		return Completed, nil
	case EgressDenied, "denied", "denied(policy-unavailable)":
		return EgressDenied, nil
	case RejectedLocal:
		return RejectedLocal, nil
	case TransportFailed:
		return TransportFailed, nil
	case StreamLost, "unknown":
		return StreamLost, nil
	case CancelledPre, "cancelled", "cancelled(pre_transport)":
		return CancelledPre, nil
	case CancelledPost, "cancelled(post_invocation)":
		return CancelledPost, nil
	case StaleEpoch, EpochAhead:
		return disposition, nil
	}
	const rejectedPrefix = "rejected_local("
	text := string(disposition)
	if strings.HasPrefix(text, rejectedPrefix) && strings.HasSuffix(text, ")") {
		reason := RejectReason(strings.TrimSuffix(strings.TrimPrefix(text, rejectedPrefix), ")"))
		switch reason {
		case MalformedRequest, LaneCapabilityMismatch, ReplayScopeViolation, InternalIntegrityFault:
			return RejectedLocal, nil
		}
	}
	return "", errors.New("provider: unknown attempt disposition")
}

func (cycle *Cycle) Cancel(ctx context.Context, attemptID string, epoch string) (Outcome, error) {
	if _, err := wire.ParseCounter(epoch); err != nil {
		return Outcome{}, fmt.Errorf("provider: cancellation turn epoch: %w", err)
	}
	if attemptID == "" || epoch != cycle.epoch {
		return Outcome{Disposition: StaleEpoch}, nil
	}
	disposition, err := cycle.connector.Cancel(ctx, attemptID, epoch)
	if err != nil {
		return Outcome{}, err
	}
	outcome := Outcome{Disposition: disposition}
	if disposition == CancelledPost {
		end := StreamCancelled
		outcome.WireStarted, outcome.StreamEnd = true, &end
		if err := cycle.gate.RecordStreamEnd(ctx, attemptID, end); err != nil {
			return outcome, err
		}
	} else if disposition != CancelledPre {
		return Outcome{}, errors.New("provider: invalid cancellation cut")
	}
	return outcome, nil
}
