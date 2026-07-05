package lineage

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/tables"
)

const (
	ParentUnknownRecompose    = "parent-unknown-recompose"
	ParentInvalidDeadEdge     = "parent-invalid-dead-edge"
	ParentNotAddressed        = "parent-not-addressed"
	MergeGrantMissing         = "merge-grant-missing"
	ScopeFlipDrift            = "scope-flip-drift"
	DesignReviewMissing       = "design-review-missing"
	GrantParentMissing        = "grant-parent-missing"
	ReviewerVisibilityMissing = "reviewer-visibility-missing"
)

type Bounce struct {
	Edge string
	Kind string
}

type Engine struct {
	Reg *fieldspec.Registry
	T   *tables.T
}

type TurnContext struct {
	WokenOn        string
	ActiveDispatch string
}

func AuthorityBearing(cand record.Record, meta seat.SeatMeta) bool {
	return (&Engine{}).AuthorityBearing(cand, meta)
}

func (e *Engine) AuthorityBearing(cand record.Record, meta seat.SeatMeta) bool {
	if cand.Headers["grant"] != "" {
		return true
	}
	if cand.Headers["HUMAN_GATE_REQUIRED"] == "yes" {
		return true
	}
	if e.gateCategoryAuthorityBearing(cand.Headers["gate_category"]) {
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

func (e *Engine) gateCategoryAuthorityBearing(value string) bool {
	if value == "" {
		return false
	}
	if e.Reg == nil {
		return true
	}
	return e.Reg.GateCategoryAuthorityBearing(value)
}

func Check(cand record.Record, t *tables.T) *Bounce {
	return (&Engine{T: t}).Check(cand, seat.SeatMeta{Name: cand.Envelope.From, Role: cand.Envelope.Role})
}

func (e *Engine) Check(cand record.Record, meta seat.SeatMeta) *Bounce {
	if !e.AuthorityBearing(cand, meta) {
		return nil
	}
	if bounce := e.checkDesignReviewChain(cand, meta); bounce != nil {
		return bounce
	}
	if bounce := e.checkPairDispatchGrant(cand, meta); bounce != nil {
		return bounce
	}
	if bounce := e.checkReviewerVisibility(cand, meta); bounce != nil {
		return bounce
	}
	if bounce := e.checkParentSubstrate(cand); bounce != nil {
		return bounce
	}
	if bounce := e.checkIMPLAddressee(cand, meta); bounce != nil {
		return bounce
	}
	if bounce := e.checkMergeClaim(cand); bounce != nil {
		return bounce
	}
	if bounce := e.checkScopeFlip(cand); bounce != nil {
		return bounce
	}
	return nil
}

func (e *Engine) checkDesignReviewChain(cand record.Record, meta seat.SeatMeta) *Bounce {
	lockID := cand.Headers["DESIGN_LOCK_ID"]
	kind := cand.Headers["DESIGN_RECORD_KIND"]
	if lockID == "" || kind == "" {
		return nil
	}
	switch kind {
	case "design-doc":
		parent := cand.Headers["PARENT_DISPATCH_ID"]
		if parent == "" {
			return &Bounce{Edge: "DESIGN_LOCK_ID", Kind: DesignReviewMissing}
		}
		for _, review := range e.parents(parent) {
			if !accepted(review) || review.Headers["PHASE"] != "DESIGN-REVIEW" || review.Headers["DESIGN_DOC_ID"] != lockID || !approving(review) {
				continue
			}
			for _, design := range e.parents(review.Headers["PARENT_DISPATCH_ID"]) {
				if accepted(design) && design.Headers["PHASE"] == "DESIGN" && design.Headers["DESIGN_DOC_ID"] == lockID && design.Envelope.From == cand.Envelope.From {
					return nil
				}
			}
		}
		return &Bounce{Edge: "DESIGN_LOCK_ID", Kind: DesignReviewMissing}
	case "audit-record":
		for _, rec := range e.records() {
			if accepted(rec) && rec.Headers["PHASE"] == "DESIGN" && rec.Headers["DESIGN_DOC_ID"] == lockID && rec.Envelope.From == cand.Envelope.From {
				return &Bounce{Edge: "DESIGN_RECORD_KIND", Kind: DesignReviewMissing}
			}
		}
	case "direct-override":
		if !(meta.IsOperator || meta.Role == "operator" || meta.Role == "orchestrator-planner" || meta.Name == "operator" || meta.Name == "orchestrator") {
			return &Bounce{Edge: "DESIGN_RECORD_KIND", Kind: DesignReviewMissing}
		}
	}
	return nil
}

func (e *Engine) checkPairDispatchGrant(cand record.Record, meta seat.SeatMeta) *Bounce {
	if cand.Headers["grant"] != "dispatch-impl" || !isPairPlannerMeta(meta, cand) {
		return nil
	}
	parent := cand.Headers["PARENT_DISPATCH_ID"]
	if parent == "" {
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: GrantParentMissing}
	}
	target := cand.Envelope.To
	if target == "" {
		target = cand.Headers["TO"]
	}
	for _, review := range e.parents(parent) {
		if !accepted(review) || review.Headers["PHASE"] != "PLAN-REVIEW" || !approving(review) {
			continue
		}
		for _, plan := range e.parents(review.Headers["PARENT_DISPATCH_ID"]) {
			if accepted(plan) && plan.Headers["PHASE"] == "PLAN" && plan.Envelope.From == cand.Envelope.From && (target == "" || addressedTo(plan, target)) {
				return nil
			}
		}
	}
	return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: GrantParentMissing}
}

func (e *Engine) checkReviewerVisibility(cand record.Record, meta seat.SeatMeta) *Bounce {
	if !isOrchestratorPlanner(meta, cand) || !broadSet(cand) {
		return nil
	}
	reviewer := reviewerFor(meta.Name)
	if reviewer == "" || addressedTo(cand, reviewer) || addressedInHeader(cand.Headers["CC"], reviewer) {
		return nil
	}
	for _, rec := range e.records() {
		if accepted(rec) && rec.Headers["ORCH_REVIEW_WAIVER"] != "" && (rec.Envelope.From == "operator" || rec.Envelope.Role == "operator") {
			return nil
		}
	}
	return &Bounce{Edge: "ORCH_REVIEW_WAIVER", Kind: ReviewerVisibilityMissing}
}

func (e *Engine) checkParentSubstrate(cand record.Record) *Bounce {
	parent := cand.Headers["PARENT_DISPATCH_ID"]
	if parent == "" {
		return nil
	}
	if e.T == nil {
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentUnknownRecompose}
	}
	if rec, ok := e.T.ByRelay[parent]; ok {
		if rec.Envelope.DeliveryState == record.Accepted {
			return nil
		}
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentInvalidDeadEdge}
	}
	for _, rec := range e.T.ByDispatch[parent] {
		if rec.Envelope.DeliveryState == record.Accepted {
			return nil
		}
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentInvalidDeadEdge}
	}
	return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentUnknownRecompose}
}

func (e *Engine) checkIMPLAddressee(cand record.Record, meta seat.SeatMeta) *Bounce {
	if cand.Headers["PHASE"] != "IMPL" || !claimsAction(cand) {
		return nil
	}
	parent := cand.Headers["PARENT_DISPATCH_ID"]
	if parent == "" {
		return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentUnknownRecompose}
	}
	for _, rec := range e.parents(parent) {
		if addressedTo(rec, meta.Name) {
			return nil
		}
	}
	return &Bounce{Edge: "PARENT_DISPATCH_ID", Kind: ParentNotAddressed}
}

func (e *Engine) checkMergeClaim(cand record.Record) *Bounce {
	if cand.Headers["grant"] != "dispatch-merge" {
		return nil
	}
	for _, rec := range e.records() {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Envelope.DispatchID != cand.Envelope.DispatchID {
			continue
		}
		if rec.Headers["PHASE"] == "MERGE-GATE" && rec.Headers["grant"] == "dispatch-merge" {
			return nil
		}
	}
	return &Bounce{Edge: "grant", Kind: MergeGrantMissing}
}

func (e *Engine) checkScopeFlip(cand record.Record) *Bounce {
	if cand.Headers["ROW_TRUTH_CHECK"] != "required" {
		return nil
	}
	now := scopeRows(cand)
	if len(now) == 0 {
		return nil
	}
	for _, rec := range e.records() {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Envelope.DispatchID != cand.Envelope.DispatchID {
			continue
		}
		for path, oldState := range scopeRows(rec) {
			if oldState == "OUT" && now[path] == "IN" {
				return &Bounce{Edge: "ROW_TRUTH_CHECK", Kind: ScopeFlipDrift}
			}
		}
	}
	return nil
}

func (e *Engine) parents(parent string) []record.Record {
	if e.T == nil {
		return nil
	}
	if rec, ok := e.T.ByRelay[parent]; ok {
		return []record.Record{rec}
	}
	return e.T.ByDispatch[parent]
}

func (e *Engine) records() []record.Record {
	if e.T == nil {
		return nil
	}
	return e.T.Records
}

func claimsAction(cand record.Record) bool {
	return cand.Headers["ACTIONS_GIT_REF"] != "" || cand.Headers["FINAL_GIT_STATUS_SHORT"] != "" || strings.Contains(strings.ToLower(cand.Body), "commit")
}

func addressedTo(rec record.Record, seatName string) bool {
	if rec.Envelope.To == seatName {
		return true
	}
	for _, value := range strings.Split(rec.Headers["TO"], ",") {
		if strings.TrimSpace(value) == seatName {
			return true
		}
	}
	return false
}

func scopeRows(rec record.Record) map[string]string {
	out := map[string]string{}
	for _, field := range []string{"SCOPE_DIFF", "FOLD_SCOPE"} {
		raw := rec.Headers[field]
		if raw == "" {
			continue
		}
		var rows []map[string]string
		if err := json.Unmarshal([]byte(raw), &rows); err != nil {
			continue
		}
		for _, row := range rows {
			path := row["path"]
			state := strings.ToUpper(row["state"])
			if state == "" {
				state = strings.ToUpper(row["status"])
			}
			if path != "" && state != "" {
				out[path] = state
			}
		}
	}
	return out
}

func RealGrantState(t *tables.T) fieldspec.GrantState {
	return func(seat fieldspec.SeatMeta) bool {
		if t == nil {
			return false
		}
		for _, review := range t.Records {
			if !accepted(review) || review.Headers["PHASE"] != "PLAN-REVIEW" || !approving(review) {
				continue
			}
			parent := review.Headers["PARENT_DISPATCH_ID"]
			if parent == "" {
				continue
			}
			for _, plan := range (&Engine{T: t}).parents(parent) {
				target := plan.Envelope.To
				if target == "" {
					target = plan.Headers["TO"]
				}
				if target != "" && accepted(plan) && plan.Headers["PHASE"] == "PLAN" && plan.Envelope.From == seat.Name && addressedTo(plan, target) {
					return true
				}
			}
		}
		return false
	}
}

func ActiveLineageCandidates(t *tables.T, turn TurnContext) fieldspec.ParentCandidates {
	return func(fieldspec.SeatMeta) ([]string, string) {
		seen := map[string]bool{}
		var candidates []string
		add := func(value string) {
			if value == "" || seen[value] {
				return
			}
			seen[value] = true
			candidates = append(candidates, value)
		}
		add(turn.WokenOn)
		add(turn.ActiveDispatch)
		if t != nil && turn.ActiveDispatch != "" {
			for _, rec := range t.ByDispatch[turn.ActiveDispatch] {
				if rec.Envelope.DeliveryState == record.Accepted {
					add(rec.Envelope.RelayID)
				}
			}
		}
		dflt := ""
		if len(candidates) > 0 {
			dflt = candidates[0]
		}
		return candidates, dflt
	}
}

func approving(rec record.Record) bool {
	for _, field := range []string{"PLAN_REVIEW_VERDICT", "DESIGN_REVIEW_VERDICT", "verdict"} {
		value := rec.Headers[field]
		if value == "approve" || value == "approved" {
			return true
		}
	}
	return false
}

func accepted(rec record.Record) bool {
	return rec.Envelope.DeliveryState == record.Accepted
}

func isPairPlannerMeta(meta seat.SeatMeta, cand record.Record) bool {
	role := meta.Role
	if role == "" {
		role = cand.Envelope.Role
	}
	name := meta.Name
	if name == "" {
		name = cand.Envelope.From
	}
	return role == "planner" || strings.HasSuffix(name, ".planner")
}

func isOrchestratorPlanner(meta seat.SeatMeta, cand record.Record) bool {
	role := meta.Role
	if role == "" {
		role = cand.Envelope.Role
	}
	name := meta.Name
	if name == "" {
		name = cand.Envelope.From
	}
	return role == "orchestrator-planner" || strings.HasSuffix(name, ".orchestrator-planner")
}

func broadSet(cand record.Record) bool {
	switch cand.Headers["PHASE"] {
	case "AUDIT", "DESIGN", "REVIEW-FOLD", "MERGE-GATE":
		return true
	case "PLAN":
		return cand.Headers["grant"] != ""
	case "IMPL":
		return cand.Headers["grant"] == "dispatch-impl"
	default:
		return false
	}
}

func reviewerFor(planner string) string {
	if strings.HasSuffix(planner, ".orchestrator-planner") {
		return strings.TrimSuffix(planner, ".orchestrator-planner") + ".orchestrator-reviewer"
	}
	return ""
}

func addressedInHeader(header, seatName string) bool {
	for _, value := range strings.Split(header, ",") {
		if strings.TrimSpace(value) == seatName {
			return true
		}
	}
	return false
}
