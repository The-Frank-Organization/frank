// Package provider consumes the normalized m-8 attempt stream without owning
// provider translation, credentials, routing, or retry policy.
package provider

import (
	"context"
	"encoding/json"
	"errors"
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
	TurnEpoch     uint64
	ProviderLane  string
	OpaqueRequest json.RawMessage
}

type Event struct {
	Kind   string
	Opaque json.RawMessage
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
	Cancel(context.Context, string, uint64) (Disposition, error)
}

type Cycle struct {
	gate      Gate
	connector Connector
	epoch     uint64
}

func New(gate Gate, connector Connector, epoch uint64) (*Cycle, error) {
	if gate == nil || connector == nil {
		return nil, errors.New("provider: gate and connector are required")
	}
	return &Cycle{gate: gate, connector: connector, epoch: epoch}, nil
}

func (cycle *Cycle) Run(ctx context.Context, request Request) (Outcome, error) {
	if request.AttemptID == "" || request.TurnID == "" || request.ProviderLane == "" || request.TurnEpoch != cycle.epoch {
		return Outcome{Disposition: StaleEpoch}, nil
	}
	if err := cycle.gate.AttemptOpen(ctx, request); err != nil {
		return Outcome{}, err
	}
	disposition, items, err := cycle.connector.Attempt(ctx, request)
	if err != nil {
		return Outcome{Disposition: StreamLost}, err
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

func (cycle *Cycle) Cancel(ctx context.Context, attemptID string, epoch uint64) (Outcome, error) {
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
