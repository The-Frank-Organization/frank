## RECONCILE -- master.orchestrator-reviewer / c2 fold-dispatch review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c2-fold-dispatch
PARENT_DISPATCH_ID: c2-consumer-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- approve fold-dispatch shape; no c2 lock/PLAN/IMPL authority granted
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-4.planner, m-5.planner, m-6.planner

Verdict: approve.

Relays reviewed:
- `master/relays/c2-fold-m-3/DESIGN-orchestrator-planner-20260630-035412.md`
- `master/relays/c2-fold-m-4/DESIGN-orchestrator-planner-20260630-035412.md`

Basis checked:
- `master/relays/c2-consumer-reconcile/RECONCILE-orchestrator-reviewer-20260630-034321.md`
- standalone lint for both reviewed relays
- scoped lint for `master/relays/c2-fold-m-3`
- scoped lint for `master/relays/c2-fold-m-4`

Findings:

1. The m-3 fold dispatch matches the approved bounds. It routes only F1 and F3 to m-3, preserves the VP-pinned wording that `slot_in` is conductor-owned and classified at work-record acceptance, rejects the spawn-derived wording, keeps `slot_in` as the work-archetype axis only, and repeats the no concrete Step-1 values / no required-when branch / no m-2 ownership-category change guardrail.

2. The m-4 fold dispatch matches the approved bounds. It routes only F2, F3, and M4-1 to m-4, asks for an opaque replay home for `seat_archetype` / resolved `authority_ceiling` or an explicit all-template route, keeps `seat_archetype` as a distinct per-seat-at-spawn opaque tag, and makes M4-1 an explicit confirm of the c1 HUMAN_GATE path rather than a new gate class.

3. Both dispatches correctly require implementer re-approval before c2 lock. That is load-bearing: neither m-3 nor m-4 may self-lock the fold, and the orchestrator should not bring the c2 lock until both fold-complete reports and implementer approvals are visible.

4. No phase leak found. Both relays stay in DESIGN/design-only, do not authorize PLAN, IMPL, merge, pcode/source edits, or final c2 lock.

Carry-forward guardrails:
- If m-3 needs concrete `slot_in` values or predicate branches, route an m-2 micro-fold.
- If m-4 cannot confirm B-to-A routing escalation through the existing HUMAN_GATE mechanism, treat it as a c2 lock blocker.
- If m-4 makes `seat_archetype` authority-bearing outside the m-4 routing-record mechanism, route a new cross-domain review instead of folding silently.

Approved next action:
- m-3 and m-4 may execute the bounded design folds as dispatched.
- Each must send a design-review request to its implementer.
- Orchestrator should wait for fold-complete plus implementer approvals before any c2 lock relay.

Not authorized:
- no PLAN;
- no IMPL;
- no merge;
- no pcode/source changes;
- no c2 lock from dispatch shape alone.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
