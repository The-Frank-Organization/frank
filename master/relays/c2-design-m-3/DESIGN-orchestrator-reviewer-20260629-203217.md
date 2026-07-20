## DESIGN relay review -- master.orchestrator-reviewer / c2-design-m-3 sequencing relay

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c2-design-m-3
PARENT_DISPATCH_ID: c2-design-m-3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- narrow relay correction; no operator decision requested here
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-3.implementer, m-4.planner

Verdict: revise.

Review target:
- `master/relays/c2-design-m-3/DESIGN-orchestrator-planner-20260629-202559.md`

Context checked:
- Original m-3 DESIGN dispatch: `master/relays/c2-design-m-3/DESIGN-orchestrator-planner-20260629-191904.md`
- m-3 status relay: `master/relays/c2-m-3-sitrep/SITREP-planner-20260629-201435.md`
- m-3/m-4 COORD seed and reconciled seam:
  - `master/relays/c2-design-m3-m4-coord/COORD-orchestrator-planner-20260629-191904.md`
  - `master/relays/c2-design-m3-m4-coord/COORD-planner-20260629-192916.md`
  - `master/relays/c2-design-m3-m4-coord/COORD-RECONCILE-planner-20260629-193400.md`
- Current m-4 review state:
  - `master/relays/c2-design-m-4/DESIGN-REVIEW-implementer-20260629-202053.md`
  - `master/relays/c2-design-m-4/DESIGN-planner-20260629-203100.md`

Blocking correction:

1. `GRILL_REQUIRED: no` is unsafe as written on this relay.

The latest relay is not merely informational; it is a live DESIGN sequencing dispatch that says "Full m-3 DESIGN -- PROCEED NOW" and "design to it." The prior VP-approved c2 transition and the original m-3 design dispatch required design-grill/operator grilling for m-3:

- `c2-reconcile` reviewer approval required c2 DESIGN dispatches to carry `GRILL_REQUIRED: yes`.
- `DESIGN-orchestrator-planner-20260629-191904.md` carries `GRILL_REQUIRED: yes` and enumerates the m-3 operator-grill items: executable-claim execution surface, egress fail-closed policy, Step-1 read-vantage floor, and the record-level `evidence_integrity: mixed` decision.

This later relay's unqualified `GRILL_REQUIRED: no` can be read as a field-level downgrade of the still-active m-3 design-grill requirement. That is not justified by the relay body and would be a stale/ambiguous handoff to m-3.

Required correction:
- Either change the latest relay to `GRILL_REQUIRED: yes`, preserving the original c2-design-m-3 grill requirement; or
- replace it with an explicit scoped form such as: `GRILL_REQUIRED: no additional grill beyond the already-required c2-design-m-3 operator/design grill in DESIGN-orchestrator-planner-20260629-191904.md`.

If using the scoped form, the body should also state that the three original m-3 operator-grill items remain acceptance criteria for the full m-3 design draft and lock.

Non-blocking findings:

2. m-5 reservation is acceptable at this stage.

The relay's choice not to boot m-5 now is consistent with the prior guardrail only because it commits to an explicit lock-time reservation preserving m-5 ownership of concrete tag-space, archetype invariants, and authority-ceiling semantics. Keep this as a c2 lock prerequisite, not optional prose. The m-3 design may lock only the predicate-execution mechanism keyed on an opaque archetype tag.

3. R2-boundary sequencing is acceptable if ratification remains at c2 lock.

The reconciled m-3/m-4 seam supports the "block the dishonesty, never block the deviation" split. The latest relay phrases this as "sound" and says ratification occurs with the VP at lock; that is acceptable. Do not let m-3 treat it as already locked beyond design-draft usage. The c2 lock still needs to record the VP ratification that no model-derived predicate enters the m-2 schema gate and that the silent-deviation bounce is an observe-layer integrity veto.

4. Current m-4 state does not block m-3 drafting.

m-4's first design review returned a narrow `must-revise`, and m-4 has now requested re-review after folding the three findings. That does not block m-3 from drafting against the reconciled seam, but c2 lock still waits for the m-4 approval path to finish.

Approved after correction:
- m-3 may author the full design draft under `c2-design-m-3`.
- m-3 must still send the design-review request to `m-3.implementer`.
- The design must still carry the original operator/design-grill items, the reconciled m-3/m-4 seam, the m-5 reservation, and the F5 novelty qualification.
- No PLAN, IMPL, merge, or pcode work is authorized.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
