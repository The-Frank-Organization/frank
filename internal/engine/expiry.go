package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/observe"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

type ExpiryPromptInput struct {
	Seat       string `json:"seat"`
	CheckID    string `json:"check_id"`
	ClaimRef   string `json:"claim_ref"`
	ContentKey string `json:"content_key"`
	SoftMillis int64  `json:"soft_millis"`
	HardMillis int64  `json:"hard_millis"`
}

type ExpiryPrompter struct {
	store  *store.Store
	core   *genericPrompter[observe.ExpiryDecision]
	tables *tables.Live
}

type ExpiryRouter struct {
	mu     sync.RWMutex
	target *ExpiryPrompter
	ready  chan struct{}
	once   sync.Once
}

func NewExpiryRouter() *ExpiryRouter {
	return &ExpiryRouter{ready: make(chan struct{})}
}

func (r *ExpiryRouter) Bind(target *ExpiryPrompter) {
	r.mu.Lock()
	r.target = target
	r.mu.Unlock()
	r.once.Do(func() { close(r.ready) })
}

func (r *ExpiryRouter) Prompt(ctx context.Context, request observe.ExpiryRequest) observe.ExpiryDecision {
	select {
	case <-ctx.Done():
		return observe.ExpiryDecision{Action: observe.ExpiryKill}
	case <-r.ready:
	}
	r.mu.RLock()
	target := r.target
	r.mu.RUnlock()
	if target == nil {
		return observe.ExpiryDecision{Action: observe.ExpiryKill}
	}
	return target.Prompt(ctx, request)
}

func NewExpiryPrompter(st *store.Store, submitter resummonSubmitter) (*ExpiryPrompter, error) {
	if st == nil {
		return nil, fmt.Errorf("expiry store required")
	}
	if submitter == nil {
		return nil, fmt.Errorf("expiry submitter required")
	}
	tab, err := tables.Build(st)
	if err != nil {
		return nil, err
	}
	return &ExpiryPrompter{
		store:  st,
		core:   newGenericPrompter(submitter, "emit-executor-expiry", observe.ExpiryDecision{Action: observe.ExpiryKill}),
		tables: tables.NewLive(tab),
	}, nil
}

func (p *ExpiryPrompter) Prompt(ctx context.Context, request observe.ExpiryRequest) observe.ExpiryDecision {
	input := expiryPromptInput(request)
	gateID := ExpiryGateID(input)
	payload, err := json.Marshal(input)
	if err != nil {
		return observe.ExpiryDecision{Action: observe.ExpiryKill}
	}
	return p.core.prompt(ctx, gateID, payload, ExpiryContentHash(input), func() (observe.ExpiryDecision, bool) {
		return p.existingDecision(gateID)
	})
}

func (p *ExpiryPrompter) existingDecision(gateID string) (observe.ExpiryDecision, bool) {
	for _, rec := range p.tables.Snapshot().VerdictsByGate[gateID] {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Envelope.From != "operator" || rec.Headers["resolves_gate"] != gateID {
			continue
		}
		reply, violation := ParseODBReply(rec.Body)
		if violation != nil {
			continue
		}
		action := observe.ExpiryAction(reply.Choice)
		if action == observe.ExpiryKill || action == observe.ExpiryExtend {
			return observe.ExpiryDecision{Action: action}, true
		}
	}
	return observe.ExpiryDecision{}, false
}

func (p *ExpiryPrompter) Apply(resolution record.Record) error {
	gateID := resolution.Headers["resolves_gate"]
	if gateID == "" || resolution.Envelope.DeliveryState != record.Accepted || resolution.Envelope.From != "operator" {
		return nil
	}
	gate, err := p.store.Read(gateID)
	if err != nil {
		return nil
	}
	if gate.Headers["expiry_check_id"] == "" {
		return nil
	}
	reply, violation := ParseODBReply(resolution.Body)
	if violation != nil {
		return nil
	}
	action := observe.ExpiryAction(reply.Choice)
	if action != observe.ExpiryKill && action != observe.ExpiryExtend {
		return nil
	}
	p.recordResolution(gate, resolution)
	p.core.resolve(gateID, observe.ExpiryDecision{Action: action})
	return nil
}

func (p *ExpiryPrompter) recordResolution(gate, resolution record.Record) {
	tab := p.tables.Snapshot()
	if _, ok := tab.ByRelay[gate.Envelope.RelayID]; !ok {
		tab.OnCommit(gate)
	}
	if _, ok := tab.ByRelay[resolution.Envelope.RelayID]; !ok {
		tab.OnCommit(resolution)
	}
	p.tables.Publish(tab)
}

func ExpiryHandler(next Handler) Handler {
	return func(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		if cmd.Verb != "emit-executor-expiry" {
			if next == nil {
				return record.Record{}, nil, fmt.Errorf("unsupported internal command %q", cmd.Verb)
			}
			return next(ctx, cmd)
		}
		if cmd.Seat != "system" || cmd.Role != "system" {
			return record.Record{}, nil, fmt.Errorf("expiry command requires system emitter")
		}
		var input ExpiryPromptInput
		if err := json.Unmarshal(cmd.Payload, &input); err != nil {
			return record.Record{}, nil, fmt.Errorf("decode expiry: %w", err)
		}
		if err := validateExpiryPromptInput(input); err != nil {
			return record.Record{}, nil, err
		}
		if cmd.ContentHash != ExpiryContentHash(input) {
			return record.Record{}, nil, fmt.Errorf("expiry content hash mismatch")
		}
		choices, err := fieldspec.CanonicalMarshal([]map[string]string{
			{"label": "Kill", "value": string(observe.ExpiryKill)},
			{"label": "Extend", "value": string(observe.ExpiryExtend)},
		})
		if err != nil {
			return record.Record{}, nil, err
		}
		rec, err := obligation.SystemOperatorRecord(obligation.SystemOperatorInput{
			RelayID: ExpiryGateID(input), DeliveryState: record.Accepted, SchemaVersion: 1,
			Headers: map[string]string{
				"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "long-running check requires disposition",
				"HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security",
				"expiry_check_id": input.CheckID, "expiry_claim_ref": input.ClaimRef,
				"expiry_seat": input.Seat, "choices": choices,
			},
			Body: string(cmd.Payload),
		})
		return rec, nil, err
	}
}

func ExpiryContentHash(input ExpiryPromptInput) string {
	return "expiry:" + promptContentDigest(input.ContentKey)
}

func ExpiryGateID(input ExpiryPromptInput) string {
	return "expiry-" + promptContentDigest(input.ContentKey)
}

func expiryPromptInput(request observe.ExpiryRequest) ExpiryPromptInput {
	keyBytes, _ := json.Marshal(struct {
		Seat            string `json:"seat"`
		CheckID         string `json:"check_id"`
		ClaimRef        string `json:"claim_ref"`
		CandidateDigest string `json:"candidate_digest"`
	}{
		Seat: request.Selection.Seat, CheckID: request.Selection.CheckID,
		ClaimRef: request.Selection.ClaimRef, CandidateDigest: request.Selection.CandidateDigest,
	})
	sum := sha256.Sum256(keyBytes)
	return ExpiryPromptInput{
		Seat: request.Selection.Seat, CheckID: request.Selection.CheckID, ClaimRef: request.Selection.ClaimRef,
		ContentKey: hex.EncodeToString(sum[:]), SoftMillis: request.SoftExpiry.Milliseconds(), HardMillis: request.HardCeiling.Milliseconds(),
	}
}

func validateExpiryPromptInput(input ExpiryPromptInput) error {
	if input.CheckID == "" || input.ClaimRef == "" || input.ContentKey == "" || input.SoftMillis <= 0 || input.HardMillis < input.SoftMillis {
		return fmt.Errorf("expiry check, claim, key, and monotonic bounds required")
	}
	return nil
}
