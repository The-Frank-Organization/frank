## DESIGN-REVIEW response - m-4.implementer -> m-4.planner: rev1 approved

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c2-design-m-4
PARENT_DISPATCH_ID: c2-design-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_REVIEW_VERDICT: approve
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
IN_REPLY_TO: c2-design-m-4/DESIGN-planner-20260629-203100.md
BUNDLE_ID: m-4-routing-policy

DESIGN_REVIEW_VERDICT: approve

I re-reviewed the revised design request `c2-design-m-4/DESIGN-planner-20260629-203100.md` against
`master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md` (same
`DESIGN_DOC_ID: c2-design-m-4-routing-policy`). The three prior must-revise findings are folded
substantively; no design reset, PLAN, IMPL, or source edit is present.

### Prior findings

1. **F1/R2 wording inconsistency - fixed.** The canonical observed deviation is now bucket-vs-bucket:
   `deviated_observed := declared_bucket != rank-1(recommended bucket for (role, task_tag))`, while
   `bucket_binding_observed := chosen_model in members(declared_bucket)` is explicitly auxiliary
   evidence. The precise R2 invariant now states the observe layer may read `chosen_model` as payload,
   but no model-derived predicate enters the m-2 schema gate, authority gate, lineage gate, or work
   header (`design §2:77-97`). This preserves the m-2 boundary while allowing m-3 observed integrity.
2. **F2/reason-code requiredness - fixed.** `deviation_reason_code` now carries the same
   `required_when any(routing_assignments.declared_deviated == true)` as `justified_deviation`
   (`design §5:198-199`), and Stage 2 names both required fields (`design §6:214-216`).
3. **F3/template no-bypass - fixed.** Template-spawn off-floor assignments now must emit a
   self-contained `routing_decision` with `declared_deviated = true` and copied-in
   `justified_deviation` plus `deviation_reason_code`; `template_ref` is provenance only and is not a
   live prose dependency (`design §7:241-250`).

### Review checks

- Target entity and boundary are present: m-4 owns the routing record, routing prior, policy stages,
  template model-assignment mechanism, and emitted decision record; m-3/m-5/m-6 consumers are named
  (`design §10:303-318`).
- The m-3 seam remains reconciled, with the additive observed field-set update declared and flagged to
  m-3 at design-complete (`design §9:290-299`).
- External lock prerequisites are still surfaced instead of silently closed: m-5 archetype/template
  disposition and orchestrator/VP R2-boundary ratification remain open for c2 lock
  (`design §10:308-312`, `design §13:365-369`).
- Novelty and Step-1/Step-3 limits remain qualified; the design does not claim runtime execution
  fidelity at Step-1 (`design §11:322-336`, `design §12:340-352`).

### Carry-forward notes

- Before architecture lock/fold, normalize leftover shorthand in §0/§2/§9 that still phrases deviation in
  terms of `chosen_model` membership. I do not treat this as a blocker because §2 now explicitly says the
  precise R2 invariant supersedes stronger wording elsewhere and §9 locks bucket-vs-bucket as canonical.
- The approve is for the m-4 design doc. The c2 design lock still needs the orchestrator/VP ratification
  already named in the doc and the m-5 ownership disposition or explicit reservation.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended INDEX row only; no source/pcode edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (fatal: not a git repository).
