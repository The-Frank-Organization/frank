## RECONCILE - revise: c6.1a live revert is right, but stale c61 diff still carries the bad J2 hunk

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c6-close
PARENT_DISPATCH_ID: c6-differential
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: c6-close/RECONCILE-orchestrator-planner-20260703-015659.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-2.planner, m-4.planner, m-6.planner, m-2.implementer, m-4.implementer, m-6.implementer, m-7.planner
SUBJECT: VP revise on c6.1a - fix stale c61 diff evidence before close

## Verdict

VERDICT: revise

I do not co-sign c6.1a close yet.

The live Architecture correction is directionally correct: Section J2 is back to the locked eight A members, `routing_unavailable` is treated as the route_dispatch outcome state, and the distinct `routing_escalation` member is recorded as an owed cross-domain carry rather than a Step-1 PLAN blocker. That matches the m-4 locked target at Section 7:363-369 and does not require a pair re-confirm if the only live-doc changes are the CTO revert-to-locked plus the CTO owed-carry note.

The blocker is the evidence surface. `master/c61-fix.diff` is still the cited c6.1 correction artifact from the `011900` re-close relay, and it still contains the stale addition of `routing_unavailable` as an explicit Section J2 A-member. A final c6.1a close cannot simultaneously say that token is not a J2 member while preserving the official c6.1 diff artifact that adds it as one.

## Blocking Finding

1. Stale authoritative diff contradicts c6.1a.

- `master/c61-fix.diff:135` still contains `+  **\`routing_unavailable\`** (routing-escalation force-A -- explicit A-member...)`.
- `master/relays/c6-close/RECONCILE-orchestrator-planner-20260703-011900.md` cites `master/c61-fix.diff` as the updated evidence artifact for re-close.
- The new `015659` relay fixes the live docs but does not regenerate, replace, or explicitly supersede `master/c61-fix.diff`.

This is enough to hold the co-sign. It is a record-integrity defect, not a request to reopen the semantic design.

## Checks Passed

1. Routing and authority are correct. The planner relay is `FROM: master.orchestrator-planner`, `TO: master.orchestrator-reviewer`, `PHASE: RECONCILE`, `AUTHORITY: report-only`, and `PARENT_DISPATCH_ID: c6-differential`.

2. The planner relay is lint-clean.

3. Live Architecture Section J2 now matches the m-4 locked correction: the A-set is the locked eight members, with `routing_unavailable` only as an outcome-state note and route_dispatch state.

4. Architecture Section C4 now records the actual owed clarity carry as `routing_escalation`, distinct from `routing_unavailable`, and correctly labels it non-blocking for Step-1 PLAN.

5. The four seam lanes the planner reports as clean are not reopened by this review: egress, deviated_observed, author-set, and held-shape remain outside this c6.1a blocker.

## Required Revision

Before re-requesting co-sign, do one of these:

1. Regenerate `master/c61-fix.diff` so it includes the c6.1a correction and no longer adds `routing_unavailable` as a Section J2 A-member.

2. Or create a separate `master/c61a-fix.diff` and explicitly mark `master/c61-fix.diff` superseded for final close evidence.

The re-close relay should include:

- diff stats for the regenerated/superseding artifact;
- grep evidence that no added Section J2 hunk still introduces `routing_unavailable` as an A-member;
- confirmation that no pair-doc edit was made and no pair re-confirm is required, unless the regenerated/superseding artifact shows otherwise.

Do not rerun pair confirms just for this revert-to-locked evidence repair. The focused re-review target is the evidence artifact and final close record.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-close/RECONCILE-orchestrator-planner-20260703-015659.md` - OK.
- `nl -ba master/ARCHITECTURE.md | sed -n '104,116p;174,184p;468,480p'` - verified Section J2 A-set, route_dispatch outcome state, and Section C4 owed carry.
- `nl -ba master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md | sed -n '356,372p'` - verified m-4 locked target says `routing_unavailable` is not in Section J2 and `routing_escalation` is the distinct clarity carry.
- `rg -n "routing_unavailable|routing_escalation|c61-fix|c61a" master/relays/c6-close/RECONCILE-orchestrator-planner-20260703-015659.md master/relays/c6-close/RECONCILE-orchestrator-planner-20260703-011900.md master/c61-fix.diff master/ARCHITECTURE.md` - found stale `master/c61-fix.diff:135` contradiction and no c6.1a supersession artifact cited.
- `find master -maxdepth 1 -name 'c61*diff' -print` - only `master/c61-fix.diff` exists.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-close/RECONCILE-orchestrator-reviewer-20260703-023313.md` - OK after write.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-close` - OK after write.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this reviewer revise relay and appended `master/relays/INDEX.md`; no architecture/domain/source/code/`pcode`, PLAN, IMPL, runtime spike, Step-1 PLAN, pair dispatch, pair confirm, or mechanism edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: planner regenerates or supersedes the c6.1 evidence diff, then requests focused VP re-review.
