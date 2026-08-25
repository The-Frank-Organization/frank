package conduct

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	AttachOK            = "attach-ok"
	AttachSuspended     = "broker:attach-suspended"
	AttachTupleMismatch = "broker:attach-tuple-mismatch"
)

var (
	ErrAttachDeadline   = errors.New("conduct: attach deadline exceeded")
	ErrGenerationFenced = errors.New("conduct: generation fenced")
)

type AttachTuple struct {
	RunID        string `json:"run_id"`
	GenerationID string `json:"generation_id"`
	TurnEpoch    uint64 `json:"turn_epoch"`
}

type AttachReply struct {
	Result    string
	Transport Transport
}

type Attacher interface {
	Attach(context.Context, AttachTuple) (AttachReply, error)
}

type Control interface {
	AttachResult(context.Context, AttachTuple, string) error
	WakeForward(context.Context, string) error
}

type Receiver struct {
	attacher Attacher
	control  Control
	deadline time.Duration
	backoff  time.Duration

	mu     sync.Mutex
	client *Client
}

func NewReceiver(attacher Attacher, control Control, deadline, backoff time.Duration) (*Receiver, error) {
	if attacher == nil || control == nil || deadline <= 0 || backoff <= 0 {
		return nil, errors.New("conduct: receiver dependencies and positive bounds are required")
	}
	return &Receiver{attacher: attacher, control: control, deadline: deadline, backoff: backoff}, nil
}

// Attach acquires only an already-authorized connection-scoped capability. It
// has no credential input or credential-resolution callback.
func (receiver *Receiver) Attach(ctx context.Context, tuple AttachTuple) (*Client, error) {
	bounded, cancel := context.WithTimeout(ctx, receiver.deadline)
	defer cancel()
	for {
		reply, err := receiver.attacher.Attach(bounded, tuple)
		if err != nil {
			if errors.Is(bounded.Err(), context.DeadlineExceeded) {
				return nil, ErrAttachDeadline
			}
			return nil, err
		}
		if err := receiver.control.AttachResult(bounded, tuple, reply.Result); err != nil {
			return nil, err
		}
		switch reply.Result {
		case AttachOK:
			client, err := New(reply.Transport)
			if err != nil {
				return nil, err
			}
			receiver.mu.Lock()
			previous := receiver.client
			receiver.client = client
			receiver.mu.Unlock()
			if previous != nil {
				_ = previous.Close()
			}
			return client, nil
		case AttachSuspended:
			timer := time.NewTimer(receiver.backoff)
			select {
			case <-bounded.Done():
				timer.Stop()
				return nil, ErrAttachDeadline
			case <-timer.C:
			}
		case AttachTupleMismatch:
			return nil, ErrGenerationFenced
		default:
			return nil, fmt.Errorf("conduct: unknown attach result %q", reply.Result)
		}
	}
}

// Reattach replaces dead connection-scoped material at the same caller-owned
// tuple, then performs durable mailbox rediscovery before normal work resumes.
func (receiver *Receiver) Reattach(ctx context.Context, tuple AttachTuple) ([]json.RawMessage, error) {
	client, err := receiver.Attach(ctx, tuple)
	if err != nil {
		return nil, err
	}
	return Rediscover(ctx, client)
}

// Rediscover treats project/read as the durable truth and keeps no local
// delivery or repair ledger.
func Rediscover(ctx context.Context, client *Client) ([]json.RawMessage, error) {
	if client == nil {
		return nil, errors.New("conduct: client is absent")
	}
	projected, err := client.Relay(ctx, "relay.project", json.RawMessage(`{}`))
	if err != nil {
		return nil, err
	}
	var relayIDs []string
	if err := json.Unmarshal(projected, &relayIDs); err != nil {
		return nil, fmt.Errorf("conduct: project response: %w", err)
	}
	records := make([]json.RawMessage, 0, len(relayIDs))
	for _, relayID := range relayIDs {
		arguments, _ := json.Marshal(map[string]string{"relay_id": relayID})
		record, err := client.Relay(ctx, "relay.read", arguments)
		if err != nil {
			return nil, err
		}
		records = append(records, append(json.RawMessage(nil), record...))
	}
	return records, nil
}

// ForwardNextPush forwards only the advisory relay identity; the caller's
// durable rediscovery loop remains the delivery mechanism.
func (receiver *Receiver) ForwardNextPush(ctx context.Context) error {
	receiver.mu.Lock()
	client := receiver.client
	receiver.mu.Unlock()
	if client == nil {
		return errors.New("conduct: client is absent")
	}
	frame, err := client.NextPush(ctx)
	if err != nil {
		return err
	}
	var nudge struct {
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(frame, &nudge); err != nil || nudge.RelayID == "" {
		return errors.New("conduct: malformed push")
	}
	return receiver.control.WakeForward(ctx, nudge.RelayID)
}

type ErrorDisposition string

const (
	DispositionFence      ErrorDisposition = "fence"
	DispositionHold       ErrorDisposition = "hold"
	DispositionReinvoke   ErrorDisposition = "reinvoke"
	DispositionRediscover ErrorDisposition = "rediscover"
	DispositionReconnect  ErrorDisposition = "reconnect"
	DispositionTerminal   ErrorDisposition = "terminal"
)

func DispositionForError(err error, submit bool) ErrorDisposition {
	if err == nil {
		return DispositionTerminal
	}
	message := err.Error()
	switch {
	case strings.Contains(message, "broker:stale-epoch"), strings.Contains(message, AttachTupleMismatch):
		return DispositionFence
	case strings.Contains(message, "broker:suspended"), strings.Contains(message, "broker:preparing"), strings.Contains(message, AttachSuspended):
		return DispositionHold
	case strings.Contains(message, "broker:record-unavailable"):
		if submit {
			return DispositionRediscover
		}
		return DispositionReinvoke
	case strings.Contains(message, "broker:unknown-outcome"):
		return DispositionRediscover
	case strings.Contains(message, "connection-lost"), strings.Contains(message, "closed network connection"):
		return DispositionReconnect
	default:
		return DispositionTerminal
	}
}
