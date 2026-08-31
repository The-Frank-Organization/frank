package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"sync"

	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
)

func promptContentDigest(contentKey string) string {
	sum := sha256.Sum256([]byte(contentKey))
	return hex.EncodeToString(sum[:])
}

type promptPending[D any] struct {
	done     chan struct{}
	once     sync.Once
	decision D
}

func newPromptPending[D any]() *promptPending[D] {
	return &promptPending[D]{done: make(chan struct{})}
}

func (p *promptPending[D]) resolve(decision D) {
	p.once.Do(func() {
		p.decision = decision
		close(p.done)
	})
}

type genericPrompter[D any] struct {
	submitter resummonSubmitter
	verb      string
	fallback  D
	mu        sync.Mutex
	pending   map[string]*promptPending[D]
}

func newGenericPrompter[D any](submitter resummonSubmitter, verb string, fallback D) *genericPrompter[D] {
	return &genericPrompter[D]{
		submitter: submitter,
		verb:      verb,
		fallback:  fallback,
		pending:   map[string]*promptPending[D]{},
	}
}

func (p *genericPrompter[D]) prompt(
	ctx context.Context,
	gateID string,
	payload []byte,
	contentHash string,
	existing func() (D, bool),
) D {
	p.mu.Lock()
	pending := p.pending[gateID]
	owner := pending == nil
	if owner {
		pending = newPromptPending[D]()
		p.pending[gateID] = pending
	}
	p.mu.Unlock()
	if !owner {
		return p.await(ctx, pending)
	}
	defer func() {
		p.mu.Lock()
		if p.pending[gateID] == pending {
			delete(p.pending, gateID)
		}
		p.mu.Unlock()
	}()
	fail := func() D {
		pending.resolve(p.fallback)
		return p.fallback
	}
	reply, _, err := p.submitter.Submit(ctx, intake.Cmd{
		Seat: "system", Role: "system", Verb: p.verb, Payload: payload, ContentHash: contentHash,
	})
	if err != nil {
		return fail()
	}
	select {
	case <-ctx.Done():
		return fail()
	case outcome := <-reply:
		if outcome.State != record.Accepted || outcome.RelayID != gateID {
			return fail()
		}
	}
	if decision, ok := existing(); ok {
		pending.resolve(decision)
		return decision
	}
	return p.await(ctx, pending)
}

func (p *genericPrompter[D]) resolve(gateID string, decision D) {
	p.mu.Lock()
	pending := p.pending[gateID]
	p.mu.Unlock()
	if pending != nil {
		pending.resolve(decision)
	}
}

func (p *genericPrompter[D]) await(ctx context.Context, pending *promptPending[D]) D {
	select {
	case <-ctx.Done():
		return p.fallback
	case <-pending.done:
		return pending.decision
	}
}
