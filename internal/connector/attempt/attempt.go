// Package attempt tracks live provider invocations and separates typed
// cancellation intent from raw channel/process loss.
package attempt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/jackli/frank/internal/connector/frame"
	"github.com/jackli/frank/internal/connector/outcome"
	"github.com/jackli/frank/internal/connector/stream"
)

var (
	ErrInvalidIntent    = errors.New("connector: invalid cancel_attempt")
	ErrUnknownAttempt   = errors.New("connector: unknown attempt")
	ErrDuplicateAttempt = errors.New("connector: duplicate attempt")
	ErrStaleEpoch       = errors.New("STALE_EPOCH")
	ErrEpochAhead       = errors.New("EPOCH_AHEAD")
	ErrPartial          = errors.New("connector: invalid cancellation partial")
)

type Intent struct {
	AttemptID string
	TurnEpoch frame.Counter
}

type Cancellation struct {
	Event     stream.Event
	Outcome   outcome.Outcome
	Duplicate bool
}

type Manager struct {
	mu        sync.Mutex
	active    map[string]*Active
	cancelled map[string]cancelledFact
}

type cancelledFact struct {
	intent       Intent
	cancellation Cancellation
}

type Active struct {
	mu      sync.RWMutex
	id      string
	epoch   frame.Counter
	ctx     context.Context
	cancel  context.CancelFunc
	digests outcome.Digests
	invoked bool
	partial string
}

func New() *Manager {
	return &Manager{active: make(map[string]*Active), cancelled: make(map[string]cancelledFact)}
}

func (manager *Manager) Begin(parent context.Context, attemptID string, epoch frame.Counter, frozenCoreDigest, providerLoweredToolsDigest string) (*Active, error) {
	if parent == nil {
		parent = context.Background()
	}
	if attemptID == "" {
		return nil, ErrUnknownAttempt
	}
	digests := &outcome.Digests{FrozenCore: frozenCoreDigest, ProviderLoweredTools: providerLoweredToolsDigest}
	if _, err := outcome.AttemptResult(attemptID, epoch, outcome.Outcome{Kind: outcome.SentCompleted, Digests: digests}); err != nil {
		return nil, err
	}
	manager.mu.Lock()
	defer manager.mu.Unlock()
	if manager.active[attemptID] != nil {
		return nil, ErrDuplicateAttempt
	}
	if _, terminal := manager.cancelled[attemptID]; terminal {
		return nil, ErrDuplicateAttempt
	}
	ctx, cancel := context.WithCancel(parent)
	active := &Active{
		id: attemptID, epoch: epoch, ctx: ctx, cancel: cancel,
		digests: *digests, partial: "none",
	}
	manager.active[attemptID] = active
	return active, nil
}

func (active *Active) Context() context.Context {
	if active == nil {
		return context.Background()
	}
	return active.ctx
}

func (active *Active) MarkInvoked() {
	if active == nil {
		return
	}
	active.mu.Lock()
	active.invoked = true
	active.mu.Unlock()
}

func (active *Active) TryMarkInvoked() bool {
	if active == nil {
		return false
	}
	active.mu.Lock()
	defer active.mu.Unlock()
	select {
	case <-active.ctx.Done():
		return false
	default:
		active.invoked = true
		return true
	}
}

func (active *Active) Invoked() bool {
	if active == nil {
		return false
	}
	active.mu.RLock()
	defer active.mu.RUnlock()
	return active.invoked
}

func (active *Active) SetPartial(partial string) error {
	if partial != "none" && partial != "text" && partial != "tool_call_incomplete" {
		return ErrPartial
	}
	active.mu.Lock()
	active.partial = partial
	active.mu.Unlock()
	return nil
}

func (manager *Manager) Cancel(intent Intent, currentEpoch frame.Counter) (Cancellation, error) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	if terminal, ok := manager.cancelled[intent.AttemptID]; ok {
		if terminal.intent == intent {
			duplicate := terminal.cancellation
			duplicate.Duplicate = true
			return duplicate, nil
		}
		return Cancellation{}, epochError(intent.TurnEpoch, currentEpoch)
	}
	if err := validateEpoch(intent.TurnEpoch, currentEpoch); err != nil {
		return Cancellation{}, err
	}
	active := manager.active[intent.AttemptID]
	if active == nil || active.epoch != intent.TurnEpoch {
		return Cancellation{}, ErrUnknownAttempt
	}
	active.cancel()
	active.mu.RLock()
	invoked, partial := active.invoked, active.partial
	digests := active.digests
	active.mu.RUnlock()
	point := outcome.PreTransport
	if invoked {
		point = outcome.PostInvocation
	}
	cancellation := Cancellation{
		Event: stream.Event{
			Schema: stream.SchemaV2, Kind: stream.Cancelled, AttemptID: intent.AttemptID,
			Partial: partial, FrozenCoreDigest: digests.FrozenCore,
			ProviderLoweredToolsDigest: digests.ProviderLoweredTools,
		},
		Outcome: outcome.Outcome{Kind: outcome.Cancelled, CancelPoint: point, Digests: &digests},
	}
	delete(manager.active, intent.AttemptID)
	manager.cancelled[intent.AttemptID] = cancelledFact{intent: intent, cancellation: cancellation}
	return cancellation, nil
}

// AbortForLoss stops local work but deliberately records no cancellation fact.
func (manager *Manager) AbortForLoss(attemptID string) error {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	active := manager.active[attemptID]
	if active == nil {
		return ErrUnknownAttempt
	}
	active.cancel()
	delete(manager.active, attemptID)
	return nil
}

func (manager *Manager) AbortAllForLoss() {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	for attemptID, active := range manager.active {
		active.cancel()
		delete(manager.active, attemptID)
	}
}

func (manager *Manager) Finish(attemptID string) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	if active := manager.active[attemptID]; active != nil {
		active.cancel()
		delete(manager.active, attemptID)
	}
}

func (manager *Manager) Cancellation(attemptID string) (Cancellation, bool) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	fact, ok := manager.cancelled[attemptID]
	return fact.cancellation, ok
}

func ParseIntent(raw []byte) (Intent, error) {
	var wire struct {
		Type      *string        `json:"type"`
		AttemptID *string        `json:"attempt_id"`
		TurnEpoch *frame.Counter `json:"turn_epoch"`
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&wire); err != nil {
		return Intent{}, fmt.Errorf("%w: %v", ErrInvalidIntent, err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return Intent{}, ErrInvalidIntent
	}
	if wire.Type == nil || *wire.Type != "cancel_attempt" || wire.AttemptID == nil || *wire.AttemptID == "" || wire.TurnEpoch == nil {
		return Intent{}, ErrInvalidIntent
	}
	return Intent{AttemptID: *wire.AttemptID, TurnEpoch: *wire.TurnEpoch}, nil
}

func validateEpoch(presented, current frame.Counter) error {
	if presented < current {
		return ErrStaleEpoch
	}
	if presented > current {
		return ErrEpochAhead
	}
	return nil
}

func epochError(presented, current frame.Counter) error {
	if err := validateEpoch(presented, current); err != nil {
		return err
	}
	return ErrInvalidIntent
}
