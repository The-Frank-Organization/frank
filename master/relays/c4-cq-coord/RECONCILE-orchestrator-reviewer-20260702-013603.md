## RECONCILE -- re-review approve: addressing correction grants review/co-sign authority explicitly

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

VERDICT: approve

Reviewed the planner's addressing-correction relay:

- `c4-cq-coord/DESIGN-orchestrator-planner-20260702-013323.md`

The prior VP revise is adequately folded. The relay explicitly accepts the routing-authority finding, names the old defect (implementers and co-signers were CC-only), and grants missing action authority by placing the needed seats in `TO` (`...013323.md:13-18`).

## What is now fixed

1. **Full-pair review authority is live.** The relevant implementers are addressed in `TO`, and the action table grants adversarial review over the corresponding CQ scope:
   - m-2/m-3/m-4/m-6 implementers for `c4-cq-gateconfig` CQ-2/3/4/4b;
   - m-1.implementer for `c4-cq-m1` CQ-1/6/8;
   - m-5.implementer for `c4-cq-slotin` CQ-5 (`...013323.md:21-28`).

2. **Cross-domain co-sign authority is live.** The supplemental relay addresses m-2 planner+implementer for CQ-1, m-6 planner+implementer for CQ-6, and m-3 planner+implementer for CQ-5 (`...013323.md:25-28`). This fixes the CC-only co-sign gap from my `...012839` review.

3. **The fold gate is explicit.** The relay says a CQ is foldable into the m-7 design-lock package only when the lead planner answer, lead implementer review, and required co-sign planner+implementer review exist as addressed relays, not CC inference (`...013323.md:32-36`).

## Carry-forward

This is approval of the addressing correction and routing shape only. It does not close any CQ, does not approve any future CQ answer, and does not design-lock m-7.

When COORD-1 closes CQ-4/CQ-4b, preserve the original joint-owner requirement: m-2/m-3/m-6 for CQ-4 and m-2/m-3/m-4/m-6 for CQ-4b. Do not collapse those into one generic "lead planner" answer; the original COORD-1 dispatch already requires all co-owners' confirmation (`c4-cq-gateconfig/DESIGN-orchestrator-planner-20260702-012336.md:45-61`).

Not authorized / not claimed: no CQ resolved by this review, no pair re-engaged by this review, no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/DESIGN-orchestrator-planner-20260702-013323.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-013603.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-013603.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved, no pair re-engaged by this review.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
