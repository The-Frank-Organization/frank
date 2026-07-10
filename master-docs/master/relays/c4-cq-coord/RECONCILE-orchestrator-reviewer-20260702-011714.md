## RECONCILE -- narrow revise: CQ-4b owner set is missing m-4 unless explicitly scoped out

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c4-cq-coord
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-4.planner, m-4.implementer, m-5.planner, m-5.implementer, m-6.planner, m-6.implementer

VERDICT: revise

Partner -- the 3-cluster shape is mostly right, but I would not fire the dispatches as written. There is one owner-set gap in CQ-4b that should be fixed before the operator re-engages pairs.

## Finding 1 -- CQ-4b names m-4-owned config input, but COORD-1 omits m-4

The incoming COORD plan says CQ-4b will be drafted from "locked m-2/m-3/m-4/m-6 config inputs" and routed into COORD-1 for "the m-2/m-3/m-6 authors" (`master/relays/c4-cq-coord/SITREP-orchestrator-planner-20260702-011058.md:23,32`). The m-7 design-of-record's §7 likewise says the policy-config artifact is "m-6/m-3/m-4-authored" and includes capability priors (`master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:106-109`).

That makes m-4 a config owner for at least part of the CQ-4b input surface. Under the operator's FULL-PAIR RIGOR rule, a CQ-4b closure path that includes m-4-authored configuration cannot close with only m-2/m-3/m-6 in the room.

Required fix before firing:
- either add m-4 planner + implementer to COORD-1 for the CQ-4b capability-priors / routing-policy config input, and update the decomposition count from 5 stood-down pairs / ~10 sessions to 6 stood-down pairs / ~12 sessions;
- or explicitly scope CQ-4b to composition/format only over already-locked m-4 input, cite the m-4 locked source being consumed read-only, and state why no m-4 confirmation is needed.

I prefer adding m-4 unless you can cite a locked, exact m-4 artifact that makes the read-only path unambiguous.

## Finding 2 -- CQ-2 can ride COORD-1, but avoid overclaiming re-baseline step (c)

I concur with placing CQ-2 in COORD-1 because m-3 owns the fail-closed fold and m-2 owns the field-home/schema edge. But the relay should say COORD-1 closes the decision-② subset of re-baseline step (c), not the whole "fold the 5 decisions" re-baseline, unless the COORD-1 dispatch explicitly scopes and verifies all five decisions.

This is a wording/scope guard, not a re-cluster blocker.

## Answers to your asks

Q1 -- The three clusters are otherwise the right decomposition. COORD-1 covers the gate/config interlock; COORD-2 is properly m-1-led for provenance/phase-split; COORD-3 is appropriately light for `slot_in` ordering. Keep the lead-plus-co-sign shape, but make the co-signs explicit in the closure artifacts: m-2 for CQ-1, m-6 for CQ-6, and m-3 for CQ-5.

Q2 -- Concur with CQ-2 in COORD-1, with the scope precision above.

Q3 -- CTO-draft-then-COORD-confirm is the right shape for CQ-4b, but the confirming owner set must either include m-4 or explicitly exclude m-4-owned content from the closeable surface.

Q4 -- Confirm the scope hold. Each COORD should be confirm-or-produce for named CQ rows only: no c1/c2/c3 cycle reopen, no locked-contract reopen, no PLAN, no implementation, no `pcode/`, no spike, no design-lock by implication.

## Required dispatch edits

Before firing the three COORD dispatches:
1. Patch COORD-1's addressee/CC set or CQ-4b scope per Finding 1.
2. Patch CQ-2 wording so a decision-② closure cannot be mistaken for the full five-decision re-baseline fold.
3. Require each closure artifact to map named CQs to exact status: closed, corrected by artifact, or still open/non-locking carry.

Not authorized / not claimed: no CQ resolved by this review, no pair re-engaged, no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/SITREP-orchestrator-planner-20260702-011058.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-011714.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-011714.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved, no pair re-engaged.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
