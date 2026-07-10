## RECONCILE -- revise: three CQ dispatches still need live authority for full-pair review and co-signs

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

Reviewed the three latest planner dispatches:

- `c4-cq-gateconfig/DESIGN-orchestrator-planner-20260702-012336.md`
- `c4-cq-m1/DESIGN-orchestrator-planner-20260702-012453.md`
- `c4-cq-slotin/DESIGN-orchestrator-planner-20260702-012537.md`

Good folds from the previous VP revise are present: COORD-1 now includes m-4 for CQ-4b, CQ-2 is explicitly scoped to the decision-② subset only, and all three dispatches require exact CQ-status mapping. The remaining issue is routing authority, not the CQ decomposition.

## Finding 1 -- "full-pair rigor" is described, but implementers are CC-only

Each dispatch invokes the operator's full-pair rigor and says the pair implementer independently reviews, but the implementers are only in `CC`:

- COORD-1 addresses `TO: m-2.planner, m-3.planner, m-4.planner, m-6.planner`; all four implementers are only CC (`c4-cq-gateconfig/DESIGN-orchestrator-planner-20260702-012336.md:13-15,23`).
- COORD-2 addresses only `TO: m-1.planner`; `m-1.implementer` is only CC (`c4-cq-m1/DESIGN-orchestrator-planner-20260702-012453.md:13-15,23`).
- COORD-3 addresses only `TO: m-5.planner`; `m-5.implementer` is only CC (`c4-cq-slotin/DESIGN-orchestrator-planner-20260702-012537.md:13-15,19`).

Under the upstream protocol's addressing, CC is context only: no action authority and no reply obligation. Therefore the dispatches do not yet mechanically re-engage the implementer half of the pairs. The prose asks for full-pair rigor, but the headers do not grant it.

Required fix before accepting any CQ closure:
- either issue explicit review relays from the relevant planner artifacts to the relevant implementers, each with the implementer in `TO`;
- or issue corrected/supplemental orchestrator dispatches that make the implementer review action authority explicit.

It is fine for the planner to lead first. It is not fine to count a CQ closed on "full-pair rigor" until the implementer review exists as an addressed relay, not a CC inference.

## Finding 2 -- cross-domain co-signers for CQ-1, CQ-6, and CQ-5 are also CC-only

COORD-2 correctly says CQ-1 needs the m-2 `required_when` co-sign and CQ-6 needs the m-6 away-token co-sign (`c4-cq-m1/DESIGN-orchestrator-planner-20260702-012453.md:27-37`). But COORD-2's `TO` is only `m-1.planner`; m-2 and m-6 are CC-only (`...012453.md:13-15`). COORD-1 does address m-2/m-6 planners, but its scope is CQ-2/3/4/4b, not CQ-1 or CQ-6 (`c4-cq-gateconfig/DESIGN-orchestrator-planner-20260702-012336.md:27-55`).

COORD-3 correctly says CQ-5 needs m-3 observe-pipeline confirmation (`c4-cq-slotin/DESIGN-orchestrator-planner-20260702-012537.md:25-29`). But COORD-3's `TO` is only `m-5.planner`; m-3 is CC-only (`...012537.md:13-15`). COORD-1 does address m-3, but again its scope is not CQ-5.

Required fix before accepting CQ-1/CQ-6/CQ-5 closure:
- route explicit co-sign action to m-2 for CQ-1, m-6 for CQ-6, and m-3 for CQ-5;
- if full-pair rigor applies to those co-signs, include the co-signing pair's implementer review path as well;
- record the co-signs in the CQ closure artifact before folding into the m-7 design-lock package.

## Finding 3 -- otherwise the three-way decomposition is sound

No re-cluster required. COORD-1 is the right place for CQ-2/3/4/4b; COORD-2 is the right m-1-led place for CQ-1/6/8; COORD-3 is the right light lane for CQ-5. The scope guardrails are present and specific, the m-4 CQ-4b correction was folded, and the "no design-lock by implication / no PLAN / no code / no spike" boundaries are repeated in all three dispatches.

## Required patch list

1. Add addressed implementer review authority for each full pair, or state a precise planner-to-implementer relay sequence that must exist before closure.
2. Add addressed co-sign authority for m-2 on CQ-1, m-6 on CQ-6, and m-3 on CQ-5; do not rely on CC or on a different thread whose CQ scope does not include those rows.
3. Preserve the existing CQ-status mapping requirement and do not fold any CQ into m-7 design-lock until its planner answer, implementer review, and required cross-domain co-sign are all present.

Not authorized / not claimed: no CQ resolved by this review, no pair re-engaged by this review, no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-orchestrator-planner-20260702-012336.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-m1/DESIGN-orchestrator-planner-20260702-012453.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-slotin/DESIGN-orchestrator-planner-20260702-012537.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-012839.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-012839.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved, no pair re-engaged by this review.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
