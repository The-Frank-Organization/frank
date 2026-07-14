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
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

type ApprovalPromptInput struct {
	Seat       string `json:"seat"`
	EntryID    string `json:"entry_id"`
	ClaimRef   string `json:"claim_ref"`
	ContentKey string `json:"content_key"`
}

type ApprovalPrompter struct {
	store  *store.Store
	core   *genericPrompter[observe.ApprovalDecision]
	tables *tables.Live
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
	tab, err := tables.Build(st)
	if err != nil {
		return nil, err
	}
	return &ApprovalPrompter{
		store:  st,
		core:   newGenericPrompter(submitter, "emit-side-effect-approval", observe.ApprovalDecision{Scope: observe.ApprovalDenied}),
		tables: tables.NewLive(tab),
	}, nil
}

func (p *ApprovalPrompter) Prompt(ctx context.Context, request observe.ApprovalRequest) observe.ApprovalDecision {
	if p.entryAllowed(request.Selection.CheckID) {
		return observe.ApprovalDecision{Allowed: true, Scope: observe.ApprovalForEntry}
	}
	input := approvalPromptInput(request)
	gateID := ApprovalGateID(input)
	payload, err := json.Marshal(input)
	if err != nil {
		return observe.ApprovalDecision{Scope: observe.ApprovalDenied}
	}
	return p.core.prompt(ctx, gateID, payload, ApprovalContentHash(input), func() (observe.ApprovalDecision, bool) {
		return p.existingDecision(gateID)
	})
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
	p.recordResolution(gate, resolution)
	p.core.resolve(gateID, decision)
	return nil
}

func (p *ApprovalPrompter) existingDecision(gateID string) (observe.ApprovalDecision, bool) {
	for _, rec := range p.tables.Snapshot().VerdictsByGate[gateID] {
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
	tab := p.tables.Snapshot()
	for _, gate := range tab.ApprovalGates[entryID] {
		for _, rec := range tab.VerdictsByGate[gate.Envelope.RelayID] {
			if rec.Envelope.From != "operator" || rec.Envelope.DeliveryState != record.Accepted {
				continue
			}
			decision, ok := approvalDecisionFromRecord(rec)
			if ok && decision.Allowed && decision.Scope == observe.ApprovalForEntry {
				return true
			}
		}
	}
	return false
}

func (p *ApprovalPrompter) recordResolution(gate, resolution record.Record) {
	tab := p.tables.Snapshot()
	if _, ok := tab.ByRelay[gate.Envelope.RelayID]; !ok {
		tab.OnCommit(gate)
	}
	if _, ok := tab.ByRelay[resolution.Envelope.RelayID]; !ok {
		tab.OnCommit(resolution)
	}
	p.tables.Publish(tab)
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
		rec, err := obligation.SystemOperatorRecord(obligation.SystemOperatorInput{
			RelayID: ApprovalGateID(input), DeliveryState: record.Accepted, SchemaVersion: 1,
			Headers: map[string]string{
				"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "side-effecting check requires live approval",
				"HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security",
				"approval_entry_id": input.EntryID, "approval_claim_ref": input.ClaimRef,
				"approval_seat": input.Seat, "choices": choices,
			},
			Body: string(cmd.Payload),
		})
		return rec, nil, err
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
