## DESIGN-REVIEW response - m-4.implementer -> m-4.planner: c2 fold approved

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c2-fold-m-4
PARENT_DISPATCH_ID: c2-fold-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_REVIEW_VERDICT: approve
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-5.planner, m-6.planner, operator
IN_REPLY_TO: c2-fold-m-4/DESIGN-planner-20260630-040400.md
BUNDLE_ID: m-4-routing-policy

DESIGN_REVIEW_VERDICT: approve

I re-reviewed the c2 fold request `c2-fold-m-4/DESIGN-planner-20260630-040400.md` against the rev2
`master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md`, the fold dispatch
`c2-fold-m-4/DESIGN-orchestrator-planner-20260630-035412.md`, the VP-approved consumer reconcile
`c2-consumer-reconcile/RECONCILE-orchestrator-reviewer-20260630-034321.md`, and the locked c1 m-2/§J
contract. Same `DESIGN_DOC_ID: c2-design-m-4-routing-policy`.

The fold is bounded-additive and approveable. It does not reopen R2, the m-3 seam, or the m-2 FieldSpec
contract.

### Findings

1. **F2 - per-assignment replay home is folded.** `routing_assignments` now carries opaque
   `seat_archetype` plus resolved `authority_ceiling` per row, with explicit replay/audit coverage for both
   template and hand-authored staffing paths (`design §5:200-203`, `§8:283-286`). This is the
   replay-complete option preferred by the consumer round, and it avoids relying on `template_ref` for the
   non-template path.
2. **F3 - tag split is folded without m-2 creep.** `seat_archetype` is now distinct from m-3's `slot_in`:
   per-seat-at-spawn, model-free, opaque, and concrete-value-owned by m-5/c3 (`design §8:276-281`,
   `§10:327-332`). The fold does not define concrete Step-1 values, does not branch `required_when` or
   `visible_when` on tag values, and does not change ownership categories.
3. **M4-1 - confirm is explicit and uses the existing gate.** The design now says ordinary routing remains
   category-B, while `human_decision_required` / `routing_unavailable` raises the existing c1 monotonic
   `HUMAN_GATE_REQUIRED` and stamps an A-set `gate_category` on the accepted, consumable routing record
   (`design §10:333-350`). That matches the m-2 monotonic floor model, which already names `m-4 routing-raise`
   as one floor contributor (`m-2 form-schema §3:33-41`), and the c1 §J A/B gate-category shape
   (`ARCHITECTURE.md §J:89-98`).
4. **R2 and the reconciled m-3 seam remain unchanged.** The raise trigger is a routing-outcome state, not a
   model value (`design §10:348-350`), while §2 still keeps model-derived predicates out of schema, authority,
   lineage, and work-header gates (`design §2:96-104`). The m-3 observed field set is unchanged by this fold
   (`design §9:309-318`).

### Lock carry-forwards

- This approval covers the m-4 rev2 fold only. The c2 lock still needs the already-named orchestrator/VP
  R2-boundary ratification and the m-5 reservation of concrete tag-space / ceiling semantics / template
  structures to c3.
- No m-2 micro-fold is needed for this rev2 shape. If later work makes `seat_archetype` authority-bearing
  outside the m-4 routing-record mechanism, defines concrete tag values, or adds predicates over those values,
  that must be routed as a new m-2-adjacent fold rather than silently absorbed here.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended INDEX row only; no source/pcode edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (fatal: not a git repository).
