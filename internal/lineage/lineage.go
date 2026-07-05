package lineage

import (
	"slices"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/tables"
)

const (
	ParentUnknownRecompose = "parent-unknown-recompose"
	ParentInvalidDeadEdge  = "parent-invalid-dead-edge"
)

type Bounce struct {
	Edge string
	Kind string
}

func AuthorityBearing(cand record.Record, meta seat.SeatMeta) bool {
	if cand.Headers["grant"] != "" {
		return true
	}
	if cand.Headers["HUMAN_GATE_REQUIRED"] == "yes" {
		return true
	}
	if isAGateCategory(cand.Headers["gate_category"]) {
		return true
	}
	if slices.Contains([]string{"PLAN", "IMPL", "REVIEW-FOLD", "MERGE-GATE", "LIVE-VERIFY"}, cand.Headers["PHASE"]) {
		return true
	}
	if slices.Contains([]string{"implementation", "merge-gated", "live-verify", "fold-in-only"}, cand.Headers["AUTHORITY"]) {
		return true
	}
	if meta.Role == "orchestrator-planner" && !slices.Contains([]string{"SITREP", "RECONCILE"}, cand.Headers["PHASE"]) {
		return true
	}
	return false
}

func Check(cand record.Record, t *tables.T) *Bounce {
	if !AuthorityBearing(cand, seat.SeatMeta{Name: cand.Envelope.From, Role: cand.Envelope.Role}) {
		return nil
	}
	parent := cand.Headers["PARENT_DISPATCH_ID"]
	if parent == "" {
		return nil
	}
	if t == nil {
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentUnknownRecompose}
	}
	if rec, ok := t.ByRelay[parent]; ok {
		if rec.Envelope.DeliveryState == record.Accepted {
			return nil
		}
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentInvalidDeadEdge}
	}
	for _, rec := range t.ByDispatch[parent] {
		if rec.Envelope.DeliveryState == record.Accepted {
			return nil
		}
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentInvalidDeadEdge}
	}
	return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentUnknownRecompose}
}

func isAGateCategory(value string) bool {
	return value == "other" || slices.Contains([]string{
		"merge_to_protected",
		"irreversible_write",
		"residual_risk_acceptance",
		"live_verify_skip",
		"ceremony_downgrade",
		"authz_security",
		"product_semantics",
		"scope_expansion",
	}, value)
}
