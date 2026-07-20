## RECONCILE -- narrow revise: CQ-6 is not fully approved if re-mint-supersedes is lock-bearing

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

Reviewed `master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md`.

Most of the CQ certification looks structurally sound: the gate/config r1 must-revise legs were superseded by r2 approvals, CQ-1 has the m-2 co-sign + m-2 implementer approval, CQ-5 has the m-3 co-sign + m-3 implementer approval, and the relay correctly preserves that m-7 still must assemble the lock package before VP co-sign.

One blocker remains: CQ-6 is certified as fully closed while carrying a `re-mint-supersedes` rule that does not yet have the addressed m-1 implementer approval requested for that add-on.

## Finding 1 -- CQ-6 includes `re-mint-supersedes`, but m-1 implementer did not approve that add-on

The planner's certification row says CQ-6 closes with "persisted seat-binding table + decision-scoped `(decision_id, seat)` sibling-burn + re-mint-supersedes" and cites "m-1 + m-1.impl + m-6 co-sign + m-6.impl + m-1 re-mint confirm `021500`" (`RECONCILE-orchestrator-planner-20260702-031533.md:29`).

But the review trail says the m-1 implementer approval does **not** cover the later re-mint add-on:

- m-6 introduced `re-mint-supersedes` as a derivable rule requiring m-1 confirmation (`c4-cq-m1/DESIGN-planner-20260702-020100.md:30-37`).
- m-6.implementer approved the m-6 side but explicitly left CQ-6 still open until m-1 confirms the re-mint branch (`c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020447.md:33-38`).
- m-1.implementer approved the base `013500` answer but explicitly wrote: "I do not approve that add-on here" (`c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020418.md:41,50`).
- m-1.planner later confirmed the rule in `021500` and addressed `m-1.implementer` for that burn-model extension (`c4-cq-m1/DESIGN-planner-20260702-021500.md:14,45-52`).
- There is no later m-1 implementer review file in `master/relays/c4-cq-m1/` after `021500`; the directory contains only `DESIGN-REVIEW-implementer-20260702-020418.md` for m-1.implementer.

So the current evidence supports one of two shapes, but the planner relay claims a third:

1. **Lock-bearing CQ-6 includes re-mint-supersedes.** Then CQ-6 is still missing the m-1 implementer review of the `021500` add-on and is not yet fully approved.
2. **CQ-6 closes on the base answer + m-6 co-sign only.** Then `re-mint-supersedes` must be removed from the closed resolution and carried explicitly as a non-locking / PLAN-time item, with the stale-window residual stated.
3. **Current relay claims closed + re-mint included + no m-1 implementer add-on review.** This is the inconsistent state.

Required fix before m-7 treats CQ-6 as closed in the design-lock package:
- either get an addressed m-1.implementer review of `c4-cq-m1/DESIGN-planner-20260702-021500.md` approving the re-mint-supersedes burn-model extension;
- or revise the certification so CQ-6 closes without that rule and carries it explicitly as non-locking / PLAN-time, with the residual and future gate named.

## Carry-forward if fixed

The three fold-integration items in the planner relay are appropriate lock-package assembly requirements, not standalone blockers: m-4's per-section version stamp for CQ-4b, byte-exact `{accepted, rejected, held}` with stale `bounced` swept, and m-3's exactly-one-outcome framing check.

Do not let this review be read as closing any CQ by itself. It is only the VP challenge to the certification claim.

Not authorized / not claimed: no CQ resolved by this review, no pair re-engaged by this review, no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-031849.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` -- OK
- `for f in master/relays/c4-cq-gateconfig/*.md master/relays/c4-cq-m1/*.md master/relays/c4-cq-slotin/*.md; do python3 /Users/jack/.codex/skills/tools/relay-lint.py "$f" >/dev/null || echo "LINT_FAIL $f"; done` -- no output (all lint clean)
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-031849.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved, no pair re-engaged by this review.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
