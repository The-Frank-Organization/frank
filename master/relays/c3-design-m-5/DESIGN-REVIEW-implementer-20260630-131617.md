## DESIGN-REVIEW - m-5 implementer review of archetype-system design

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c3-design-m-5
PARENT_DISPATCH_ID: c3-design-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - design-review verdict; no value lock or implementation authority
GRILL_REQUIRED: yes - reviewed folded GRILL_LOCK `c3-grill-m-5`
DESIGN_DOC_ID: c3-design-m-5-workflows-archetypes
DESIGN_REVIEW_VERDICT: must-revise
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, operator
IN_REPLY_TO: c3-design-m-5/DESIGN-planner-20260630-125604.md
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)
SUBJECT: DESIGN-REVIEW must-revise - core m-5 archetype model is sound, but the m-6 seam is not lock-consistent across the current design docs/COORD trail

Verdict: **must-revise**.

I approve the core shape of the m-5 archetype system: the two-axis tag space, derived actuator for v3.0, partial-order ceiling vector, T1/T2/T3 lineup, sensor integrity split, Step-1 recorded-contract stance, and novelty claim are defensible against the audit record. I do **not** approve the design lock yet because the m-6 seam is internally and cross-domain inconsistent in the current artifacts. That is a lock-blocking issue for this design because §8 is a named consumer boundary contract, and m-6 is the direct consumer.

## Findings

1. **Blocking - §8 locks a seam vocabulary that is not the m-6-consumed vocabulary.**

   The m-5 design locks `surface_intent {verdict, fyi, collaborate}` and says `verdict` is derived from `HUMAN_GATE_REQUIRED + gate_category`; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:137-148`. The sibling m-6 design, however, says the resolved seam is `surface_intent {progress, review_checkpoint, advisory, result}`, **non-gate records only**, with gate-bearing records carrying no `surface_intent` and binding directly off locked `gate_category`/HUMAN_GATE/J1; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:103-116`.

   The COORD trail explains the split. m-6's binding response accepted the four non-gate classes and explicitly said gate-bearing outputs carry no `surface_intent`; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-123022.md:24-43` and `:69-77`. The later m-5 final declaration switched back to `{verdict, fyi, collaborate}` and states "await your binding confirm"; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-125604.md:20-27` and `:46-47`. I do not see a later m-6 bind-confirm for that changed declaration.

   Required revision: either align m-5 §8 to the m-6-consumed four-value non-gate model from `123022`, or file a new m-6 binding-confirm that accepts the later three-value model, then update both m-5 and m-6 design docs to cite the same final seam source. Do not lock this while the producer and consumer docs disagree.

2. **Blocking - `away_bridge_eligible` ownership is contradictory across the same seam.**

   The m-5 design says `away_bridge_eligible` is a per-archetype capability ceiling that m-5 declares and m-6 narrows by policy; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:148`. The m-6 design says `away_bridge_eligible` is an m-6-owned per-gate boolean and that m-5 declares only the `away` posture; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:75-80` and `:103-116`.

   The COORD trail again is not converged: m-6 yielded to "m-6 owns the boolean" in `123022`; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-123022.md:41-43`, while m-5's final declaration returned to a per-archetype capability ceiling; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-125604.md:29-30`.

   Required revision: pick one owner/representation and make both docs match. My recommendation: for v3.0, follow the consumer doc unless the operator explicitly wants a hard m-5 never-bridge archetype ceiling now: `away_bridge_eligible` is m-6 policy over locked A/opt-in/egress fields, with a reserved future m-5 hard-ceiling hook. That keeps the external-send axis deferred and avoids adding a second authority path in m-5.

3. **Blocking - the design doc claims the COORD is reconciled while its own status and the latest COORD relay say otherwise.**

   The m-5 design header says the doc is still pending "the m-5↔m-6 COORD seam reconcile"; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:7`. But §8 says the seam is "Reconciled" and the review request summarizes it as reconciled; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:137-152` and `master/relays/c3-design-m-5/DESIGN-planner-20260630-125604.md:31-37`. The latest COORD file from m-5 says "await your binding confirm"; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-125604.md:46-47`.

   Required revision: make the status mechanically true. If the seam is still pending, leave it as a lock prerequisite and do not claim it as reconciled in §8. If the seam is reconciled, add the missing final bind/reconcile relay and cite that exact file in both design docs.

## Non-blocking Notes

- The tamper-resistance proof is acceptable as a design-level proof, with the caveat already carried to PLAN: mixed work classification needs negative fixtures and a conductor-side classifier ordering test. The proof correctly depends on locked F1 and observe-as-send, not on lane self-report; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:62-78`.
- The derived actuator decision is clean for v3.0. Deferring literal `single_bounded_action` and `external_send` to Step 4/5 is consistent with Step-1 host-enforcement limits; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:132-133` and `:156-161`.
- The novelty statement is scoped correctly: it claims the governed auditable integration, not novelty in topology, presets, or task gates individually; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:175-176`.

## Required Revision Checklist

- Update m-5 §8 and the design-review request summary to cite the actual final m-5/m-6 seam state.
- Resolve `surface_intent` to one canonical set and derivation domain across m-5 and m-6.
- Resolve `away_bridge_eligible` ownership/representation across m-5 and m-6.
- Keep the m-1 token bridge as an m-6/m-1 lock blocker only; do not make it an m-5 design blocker beyond the `away` posture touchpoint.
- Re-run relay-lint on the revised DESIGN relay and return a new DESIGN-REVIEW request.

ACTIONS_GIT_REF: wrote this read-only design-review relay only: `master/relays/c3-design-m-5/DESIGN-REVIEW-implementer-20260630-131617.md` and appended the matching `master/relays/INDEX.md` row; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c3-design-m-5/DESIGN-REVIEW-implementer-20260630-131617.md` passed.
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root=master/relays master/relays/c3-design-m-5/DESIGN-REVIEW-implementer-20260630-131617.md` returned `OK` for this relay and expected root noise for lint-exempt `INDEX.md`.
- E1: appended `master/relays/INDEX.md` row for this design-review.
- E1: `git -C pcode status --short` returned clean output.
- E1: `git status --short` at harness root exits 128 (`fatal: not a git repository...`), matching the expected structured unavailable form.
