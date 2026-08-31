// Package settle produces payload-free continuation settlement records.
package settle

import (
	"context"
	"errors"
	"sort"

	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type Producer struct{}

func (Producer) Produce(ctx context.Context, tx *store.Tx, runID, predecessorTurnID, producedForTurnID string) (appipc.SettlementManifest, []byte, error) {
	if tx == nil || runID == "" || predecessorTurnID == "" || producedForTurnID == "" {
		return appipc.SettlementManifest{}, nil, errors.New("settle: invalid production request")
	}
	ancestry, err := walkAncestry(ctx, tx, runID, predecessorTurnID)
	if err != nil {
		return appipc.SettlementManifest{}, nil, err
	}
	entries := make([]appipc.SettlementEntry, 0)
	for _, turnID := range ancestry {
		turnEntries, err := produceTurn(ctx, tx, runID, turnID)
		if err != nil {
			return appipc.SettlementManifest{}, nil, err
		}
		entries = append(entries, turnEntries...)
	}
	manifest := appipc.SettlementManifest{Version: 1, RunID: runID, ProducedForTurnID: producedForTurnID, Entries: entries}
	encoded, err := appipc.MarshalJCS(manifest)
	return manifest, encoded, err
}

func walkAncestry(ctx context.Context, tx *store.Tx, runID, turnID string) ([]string, error) {
	reversed := make([]string, 0)
	seen := make(map[string]struct{})
	for turnID != "" {
		if _, cycle := seen[turnID]; cycle {
			return nil, errors.New("settle: continuation ancestry cycle")
		}
		seen[turnID] = struct{}{}
		reversed = append(reversed, turnID)
		var predecessor *string
		if err := tx.QueryRowContext(ctx, `SELECT predecessor_turn_id FROM turns WHERE run_id=? AND turn_id=?`, runID, turnID).Scan(&predecessor); err != nil {
			return nil, err
		}
		if predecessor == nil {
			break
		}
		turnID = *predecessor
	}
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed[left], reversed[right] = reversed[right], reversed[left]
	}
	return reversed, nil
}

func produceTurn(ctx context.Context, tx *store.Tx, runID, turnID string) ([]appipc.SettlementEntry, error) {
	tools, err := toolEntries(ctx, tx, runID, turnID)
	if err != nil {
		return nil, err
	}
	orphans, err := orphanEntries(ctx, tx, runID, turnID)
	if err != nil {
		return nil, err
	}
	tools = append(tools, orphans...)
	sort.Slice(tools, func(i, j int) bool { return *tools[i].ToolCallID < *tools[j].ToolCallID })
	providers, err := providerEntries(ctx, tx, runID, turnID)
	if err != nil {
		return nil, err
	}
	return append(tools, providers...), nil
}

func toolEntries(ctx context.Context, tx *store.Tx, runID, turnID string) ([]appipc.SettlementEntry, error) {
	rows, err := tx.QueryContext(ctx, `SELECT tool_call_id,state,canonical_args_digest FROM tool_calls WHERE run_id=? AND turn_id=? ORDER BY tool_call_id`, runID, turnID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]appipc.SettlementEntry, 0)
	for rows.Next() {
		var id, state, digest string
		if err := rows.Scan(&id, &state, &digest); err != nil {
			return nil, err
		}
		class := ""
		switch state {
		case "EXECUTED":
			class = "settled_with_content"
		case "NOT_INVOKED_INTEGRITY_FAULT":
			class = "determinate_no_resume"
		case "UNKNOWN_TOOL_OUTCOME", "PARTIAL_TOOL_EFFECT":
			class = "uncertain"
		default:
			return nil, errors.New("settle: nonterminal tool row in producer domain")
		}
		idCopy, digestCopy := id, digest
		entries = append(entries, appipc.SettlementEntry{Kind: "tool", Class: class, RunID: runID, TurnID: turnID, ToolCallID: &idCopy, ArgsDigest: &digestCopy, Terminal: state})
	}
	return entries, rows.Err()
}

func orphanEntries(ctx context.Context, tx *store.Tx, runID, turnID string) ([]appipc.SettlementEntry, error) {
	rows, err := tx.QueryContext(ctx, `SELECT a.tool_call_id,a.state,a.void_reason,a.canonical_tool_name,a.canonical_args_digest FROM tool_authorizations a WHERE a.run_id=? AND a.turn_id=? AND NOT EXISTS(SELECT 1 FROM tool_calls c WHERE c.run_id=a.run_id AND c.turn_id=a.turn_id AND c.tool_call_id=a.tool_call_id) ORDER BY a.tool_call_id`, runID, turnID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]appipc.SettlementEntry, 0)
	for rows.Next() {
		var id, state, name, digest string
		var reason *string
		if err := rows.Scan(&id, &state, &reason, &name, &digest); err != nil {
			return nil, err
		}
		if state != "VOID" || reason == nil {
			return nil, errors.New("settle: non-VOID authorization orphan")
		}
		class := "determinate_no_resume"
		if *reason == "expired" && isRelay(name) {
			class = "uncertain"
		}
		idCopy, digestCopy, reasonCopy := id, digest, *reason
		entries = append(entries, appipc.SettlementEntry{Kind: "tool", Class: class, RunID: runID, TurnID: turnID, ToolCallID: &idCopy, ArgsDigest: &digestCopy, Terminal: "VOID", VoidReason: &reasonCopy})
	}
	return entries, rows.Err()
}

func providerEntries(ctx context.Context, tx *store.Tx, runID, turnID string) ([]appipc.SettlementEntry, error) {
	rows, err := tx.QueryContext(ctx, `SELECT p.attempt_id,p.state,p.cancel_point,EXISTS(SELECT 1 FROM content_ready_receipts r WHERE r.run_id=p.run_id AND r.turn_id=p.turn_id AND r.attempt_id=p.attempt_id) FROM provider_attempts p WHERE p.run_id=? AND p.turn_id=? ORDER BY p.attempt_id`, runID, turnID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]appipc.SettlementEntry, 0)
	for rows.Next() {
		var id, state string
		var cancel *string
		var receipt bool
		if err := rows.Scan(&id, &state, &cancel, &receipt); err != nil {
			return nil, err
		}
		class := ""
		switch state {
		case "COMPLETED":
			if receipt {
				class = "settled_with_content"
			} else {
				class = "uncertain"
			}
		case "REJECTED_LOCAL", "CANCELLED":
			class = "determinate_no_resume"
		case "UNKNOWN_PROVIDER_OUTCOME", "PARTIAL_STREAM":
			class = "uncertain"
		default:
			return nil, errors.New("settle: nonterminal provider row in producer domain")
		}
		idCopy := id
		entries = append(entries, appipc.SettlementEntry{Kind: "provider", Class: class, RunID: runID, TurnID: turnID, AttemptID: &idCopy, Terminal: state, CancelPoint: cancel})
	}
	return entries, rows.Err()
}

func isRelay(name string) bool {
	return name == "relay.submit" || name == "relay.project" || name == "relay.read"
}
