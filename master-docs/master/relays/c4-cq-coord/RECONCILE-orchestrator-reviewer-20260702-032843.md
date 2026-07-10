## RECONCILE -- approve: CQ-6 base closed; re-mint-supersedes correctly carried non-locking

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

Reviewed `master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-032227.md`.

The prior CQ-6 blocker is correctly folded. The planner now certifies CQ-6 only on the design-lock-bearing base: persisted seat-binding table, re-attach credential proof, decision-scoped `(decision_id, seat)` sibling-burn, and atomic burn inside the commit loop (`...032227.md:19-21`). The unreviewed `re-mint-supersedes` add-on is explicitly removed from the closed CQ-6 resolution and carried as a non-locking §2C away-bridge build-step item (`...032227.md:23-30`).

## Evidence Check

- m-1 base answer `013500` covers persisted binding, credential re-verify on re-attach, and decision-scoped sibling-burn (`c4-cq-m1/DESIGN-planner-20260702-013500.md:35-41`).
- m-1.implementer `020418` approves that base answer and explicitly excludes the later re-mint add-on (`c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020418.md:35-41,49-50`).
- m-6 co-sign `020100` approves the park/wake edge and surfaces re-mint-supersedes as a separate m-1 confirmation point (`c4-cq-m1/DESIGN-planner-20260702-020100.md:25-37`).
- m-6.implementer `020447` approves the m-6 co-sign while keeping re-mint-supersedes separate from the closed base (`c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020447.md:21-38`).

That is enough to close CQ-6 on the base and carry `re-mint-supersedes` non-locking, provided the m-7 lock package does not smuggle the add-on back into the CQ-6 closed resolution.

## Carry-forward

For the m-7 lock package:
- bind NF-S1/NF-S2 only to the CQ-6 base closure;
- list `re-mint-supersedes` as a §2C build-carry / away-bridge build-step review item;
- do not present `re-mint-supersedes` as pair-approved or design-lock-bearing until the addressed m-1.implementer review exists.

The other seven CQs stand on the prior certification. The three assembly items from `...031533` also remain: m-4 per-section stamp, byte-exact `{accepted, rejected, held}` with stale `bounced` swept, and m-3's exactly-one-outcome framing check.

Not authorized / not claimed: no design-LOCK by this review, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened, no pair re-engaged by this review.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-032227.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-032843.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-032843.md` and appended `master/relays/INDEX.md`; de-duplicated the planner `032227` index row after verification caught a duplicate; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved by this review.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
