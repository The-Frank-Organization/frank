## Team m-2 - Forms & Determinism: DESIGN-REVIEW RESPONSE (rev2)

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c1-design-m-2-review-r2
PARENT_DISPATCH_ID: c1-design-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - approval of m-2 rev2 design refinement; operator ratification and joint lock remain orchestrator-owned
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, m-1.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve

Verdict: approve.

I reviewed the rev2 request and revised design artifact:
- `master/relays/c1-refine-m-2/DESIGN-planner-20260629-165702.md`
- `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md`
- m-1 mutual reaffirm: `master/relays/c1-design-m-2/COORD-planner-20260629-165643.md` and `master/relays/c1-refine-m-2/COORD-CONCUR-planner-20260629-170032.md`

Acceptance review:
1. `HUMAN_GATE_REQUIRED` is reachable and monotonic. The rev2 hybrid model makes the field visible as a constrained pick over `[floor, MAX]`, while lower-than-floor values are absent from the rendered candidate set and rejected if submitted through a tampered/API path. Evidence: design doc lines 39, 41, 177, 222.
2. `gate_category` is a valid closed enum driver for the A/B email bucket split, with no free-text gate leak. The bucket projection reads enum categories plus `delivery_state`/`failing_edge`; prose remains out of the mechanical split. Evidence: lines 250-258, 177, 223.
3. `completed_proof` is not agent-forgeable in the Owner Decision Brief because it is an `evidence_ref` filled through m-3 observed/system evidence, not brief prose. Evidence: lines 175, 267, 305. Non-blocking constraint for PLAN: readers must respect `evidence_integrity`; `self_reported` is honest fallback evidence, not observed proof.
4. The separate routing-relay shape keeps model out of every work-relay authority gate. The model lives in the routing relay payload; work relays reference an accepted routing relay via `routing_ref` for provenance only, with no `model_*` predicate atom in required-when. Evidence: lines 63, 80-90, 176, 278-289, 304.
5. I found no new consumed field without a named consumer. The canonical additions are consumed by m-6 human surface/scheduler, m-4 router/provenance, m-3 observe/evidence, or the lineage/courier substrate. The null-reserved fields (`decision_deadline`, `on_timeout`, `outcome_feedback_ref`) are explicitly reserved and non-gating, so they do not need to move to `X-`.

Operator/joint-lock notes, non-blocking to this review:
- The proposed `on_timeout = hold_and_resummon` default and novel `gate_category` placement still need operator/joint-lock ratification as already surfaced in the design doc.
- AC fixtures should include monotonic lower-than-floor tamper rejection, closed-enum bucket routing, `completed_proof` evidence-integrity projection, and a routing-ref case proving model is not a required-when/gate input.

ACTIONS_GIT_REF: created this DESIGN-REVIEW relay at master/relays/c1-refine-m-2/DESIGN-REVIEW-Implementer-20260629-170241.md and appended the routing row to master/relays/INDEX.md; docs-workspace artifacts, no code/source implementation changes; cwd is not a git repo so a commit/ref is unavailable
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; git status reports fatal not-a-git-repository at filesystem boundary
