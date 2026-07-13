package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

type ApprovalPromptInput struct {
	Seat       string `json:"seat"`
	EntryID    string `json:"entry_id"`
	ClaimRef   string `json:"claim_ref"`
	ContentKey string `json:"content_key"`
}

type ApprovalPrompter struct {
	store     *store.Store
	submitter resummonSubmitter
	mu        sync.Mutex
	pending   map[string]chan observe.ApprovalDecision
}

type ApprovalRouter struct {
	mu     sync.RWMutex
	target *ApprovalPrompter
}

func NewApprovalRouter() *ApprovalRouter {
	return &ApprovalRouter{}
}

func (r *ApprovalRouter) Bind(target *ApprovalPrompter) {
	r.mu.Lock()
	r.target = target
	r.mu.Unlock()
}

func (r *ApprovalRouter) Prompt(ctx context.Context, request observe.ApprovalRequest) observe.ApprovalDecision {
	r.mu.RLock()
	target := r.target
	r.mu.RUnlock()
	if target == nil {
		return observe.ApprovalDecision{Scope: observe.ApprovalDenied}
	}
	return target.Prompt(ctx, request)
}

func NewApprovalPrompter(st *store.Store, submitter resummonSubmitter) (*ApprovalPrompter, error) {
	if st == nil {
		return nil, fmt.Errorf("approval store required")
	}
	if submitter == nil {
		return nil, fmt.Errorf("approval submitter required")
	}
	return &ApprovalPrompter{store: st, submitter: submitter, pending: map[string]chan observe.ApprovalDecision{}}, nil
}

func (p *ApprovalPrompter) Prompt(ctx context.Context, request observe.ApprovalRequest) observe.ApprovalDecision {
	if p.entryAllowed(request.Selection.CheckID) {
		return observe.ApprovalDecision{Allowed: true, Scope: observe.ApprovalForEntry}
	}
	input := approvalPromptInput(request)
	gateID := ApprovalGateID(input)
	decision := make(chan observe.ApprovalDecision, 1)
	p.mu.Lock()
	p.pending[gateID] = decision
	p.mu.Unlock()
	defer func() {
		p.mu.Lock()
		delete(p.pending, gateID)
		p.mu.Unlock()
	}()
	payload, err := json.Marshal(input)
	if err != nil {
		return observe.ApprovalDecision{Scope: observe.ApprovalDenied}
	}
	reply, _, err := p.submitter.Submit(ctx, intake.Cmd{
		Seat: "system", Role: "system", Verb: "emit-side-effect-approval", Payload: payload, ContentHash: ApprovalContentHash(input),
	})
	if err != nil {
		return observe.ApprovalDecision{Scope: observe.ApprovalDenied}
	}
	select {
	case <-ctx.Done():
		return observe.ApprovalDecision{Scope: observe.ApprovalDenied}
	case outcome := <-reply:
		if outcome.State != record.Accepted || outcome.RelayID != gateID {
			return observe.ApprovalDecision{Scope: observe.ApprovalDenied}
		}
	}
	if existing, ok := p.existingDecision(gateID); ok {
		return existing
	}
	select {
	case <-ctx.Done():
		return observe.ApprovalDecision{Scope: observe.ApprovalDenied}
	case picked := <-decision:
		return picked
	}
}

func (p *ApprovalPrompter) Apply(resolution record.Record) error {
	gateID := resolution.Headers["resolves_gate"]
	if gateID == "" || resolution.Envelope.DeliveryState != record.Accepted || resolution.Envelope.From != "operator" {
		return nil
	}
	gate, err := p.store.Read(gateID)
	if err != nil || gate.Headers["approval_entry_id"] == "" {
		return nil
	}
	decision, ok := approvalDecisionFromRecord(resolution)
	if !ok {
		return nil
	}
	p.mu.Lock()
	pending := p.pending[gateID]
	p.mu.Unlock()
	if pending != nil {
		select {
		case pending <- decision:
		default:
		}
	}
	return nil
}

func (p *ApprovalPrompter) existingDecision(gateID string) (observe.ApprovalDecision, bool) {
	records, err := p.store.Records()
	if err != nil {
		return observe.ApprovalDecision{}, false
	}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Envelope.From != "operator" || rec.Headers["resolves_gate"] != gateID {
			continue
		}
		if decision, ok := approvalDecisionFromRecord(rec); ok {
			return decision, true
		}
	}
	return observe.ApprovalDecision{}, false
}

func (p *ApprovalPrompter) entryAllowed(entryID string) bool {
	records, err := p.store.Records()
	if err != nil {
		return false
	}
	gates := make(map[string]record.Record)
	for _, rec := range records {
		if rec.Headers["approval_entry_id"] != "" {
			gates[rec.Envelope.RelayID] = rec
		}
	}
	for _, rec := range records {
		gate := gates[rec.Headers["resolves_gate"]]
		if gate.Headers["approval_entry_id"] != entryID || rec.Envelope.From != "operator" || rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		decision, ok := approvalDecisionFromRecord(rec)
		if ok && decision.Allowed && decision.Scope == observe.ApprovalForEntry {
			return true
		}
	}
	return false
}

func ApprovalHandler(next Handler) Handler {
	return func(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		if cmd.Verb != "emit-side-effect-approval" {
			if next == nil {
				return record.Record{}, nil, fmt.Errorf("unsupported internal command %q", cmd.Verb)
			}
			return next(ctx, cmd)
		}
		if cmd.Seat != "system" || cmd.Role != "system" {
			return record.Record{}, nil, fmt.Errorf("approval command requires system emitter")
		}
		var input ApprovalPromptInput
		if err := json.Unmarshal(cmd.Payload, &input); err != nil {
			return record.Record{}, nil, fmt.Errorf("decode approval: %w", err)
		}
		if err := validateApprovalPromptInput(input); err != nil {
			return record.Record{}, nil, err
		}
		if cmd.ContentHash != ApprovalContentHash(input) {
			return record.Record{}, nil, fmt.Errorf("approval content hash mismatch")
		}
		choices, err := fieldspec.CanonicalMarshal([]map[string]string{
			{"label": "Allow once", "value": string(observe.ApprovalOnce)},
			{"label": "Allow for entry", "value": string(observe.ApprovalForEntry)},
			{"label": "Deny", "value": string(observe.ApprovalDenied)},
		})
		if err != nil {
			return record.Record{}, nil, err
		}
		to, err := fieldspec.EncodeAddressList([]string{"operator"})
		if err != nil {
			return record.Record{}, nil, err
		}
		return record.Record{
			Envelope: record.Envelope{
				RelayID: ApprovalGateID(input), From: "system", To: "operator", Role: "system",
				DeliveryState: record.Accepted, SchemaVersion: 1,
			},
			Headers: map[string]string{
				"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "side-effecting check requires live approval",
				"TO": to, "HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security",
				"approval_entry_id": input.EntryID, "approval_claim_ref": input.ClaimRef,
				"approval_seat": input.Seat, "choices": choices,
			},
			Body: string(cmd.Payload),
		}, nil, nil
	}
}

func ApprovalContentHash(input ApprovalPromptInput) string {
	sum := sha256.Sum256([]byte(input.ContentKey))
	return "approval:" + hex.EncodeToString(sum[:])
}

func ApprovalGateID(input ApprovalPromptInput) string {
	return "approval-" + ApprovalContentHash(input)[len("approval:"):]
}

func approvalPromptInput(request observe.ApprovalRequest) ApprovalPromptInput {
	selection := request.Selection
	keyBytes, _ := json.Marshal(struct {
		Seat            string `json:"seat"`
		EntryID         string `json:"entry_id"`
		ClaimRef        string `json:"claim_ref"`
		CandidateDigest string `json:"candidate_digest"`
	}{Seat: selection.Seat, EntryID: selection.CheckID, ClaimRef: selection.ClaimRef, CandidateDigest: selection.CandidateDigest})
	sum := sha256.Sum256(keyBytes)
	return ApprovalPromptInput{
		Seat: selection.Seat, EntryID: selection.CheckID, ClaimRef: selection.ClaimRef, ContentKey: hex.EncodeToString(sum[:]),
	}
}

func approvalDecisionFromRecord(rec record.Record) (observe.ApprovalDecision, bool) {
	reply, violation := ParseODBReply(rec.Body)
	if violation != nil {
		return observe.ApprovalDecision{}, false
	}
	scope := observe.ApprovalScope(reply.Choice)
	switch scope {
	case observe.ApprovalOnce, observe.ApprovalForEntry:
		return observe.ApprovalDecision{Allowed: true, Scope: scope}, true
	case observe.ApprovalDenied:
		return observe.ApprovalDecision{Scope: scope}, true
	default:
		return observe.ApprovalDecision{}, false
	}
}

func validateApprovalPromptInput(input ApprovalPromptInput) error {
	if input.EntryID == "" || input.ClaimRef == "" || input.ContentKey == "" {
		return fmt.Errorf("approval entry, claim, and key required")
	}
	return nil
}
