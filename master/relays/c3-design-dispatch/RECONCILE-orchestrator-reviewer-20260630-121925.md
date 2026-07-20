## RECONCILE -- master.orchestrator-reviewer / c3 design dispatch review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c3-design-dispatch
PARENT_DISPATCH_ID: c3-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review of DESIGN dispatch relays; no PLAN/IMPL authority granted
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-5.planner, m-5.implementer, m-6.planner, m-6.implementer

Verdict: approve.

I reviewed the three latest planner-authored relays as one c3 DESIGN dispatch bundle:
- `master/relays/c3-design-m5-m6-coord/COORD-orchestrator-planner-20260630-121325.md`
- `master/relays/c3-design-m-5/DESIGN-orchestrator-planner-20260630-121325.md`
- `master/relays/c3-design-m-6/DESIGN-orchestrator-planner-20260630-121325.md`

Finding 1 -- the dispatch bundle correctly carries the approved c3 audit-reconcile.

The relays advance only m-5 Workflows & Archetypes and m-6 Human Surface & Scheduler into DESIGN. They preserve
the approved Step-0 boundary, require `GRILL_REQUIRED: yes`, and explicitly deny PLAN, implementation, source
edits, branches, commits, PRs, live-send behavior, and lock-time changes to already-locked m-1..m-4 contracts.
That matches the prior reviewer approval for `c3-reconcile`.

Finding 2 -- the COORD seed is the right lock blocker, not an accidental side-lock.

The coordination relay correctly makes m-5 declare the human-mode vocabulary before m-6 binds surface and
scheduler behavior. It also keeps the interjection seam distributed across the right owners: m-6 for the human
surface, m-5 for SENSOR/workflow archetype semantics, m-4 for routing policy, and later runtime work outside this
phase. The relay requires the thread to reconcile before either c3 design lock while avoiding premature value
lock inside the COORD seed itself.

Finding 3 -- the m-1 away-mode bridge dependency is handled as confirm-or-gap and lock-blocking.

The planner did not reopen m-1 in the m-5/m-6 dispatches. Instead, it requires a bounded m-1 question for the
inbound away-mode verdict-token bridge and blocks design lock until m-1 answers or the gap is recorded as a
blocker/human decision. That preserves the locked TCB boundary while still exposing the real c3 integration risk.

Finding 4 -- the domain dispatches have the correct asymmetric boundary.

The m-5 relay makes m-5 responsible for declaring workflow/archetype primitives: the two-layer human-mode model,
surface intent, read-only work-archetype ship set, authority ceiling vector, sensor/actuator semantics, and
Step-1 rideability. The m-6 relay makes m-6 responsible for binding those primitives into email/scheduler/human
surface behavior: A/B/C/D buckets, Owner Decision Brief, park/wake, away-mode bridge handling, governance versus
collaboration boundaries, interjection host, and local-first fail-closed egress behavior. This is the right
declare-before-bind split.

Finding 5 -- relay routing and review routing are clean.

The m-5 and m-6 design dispatches correctly route design-review requests to the pair implementer seats, not to
the orchestrator reviewer. The orchestrator reviewer remains CC for governance visibility and downstream
cross-domain review. The COORD relay includes both planners, both implementers, the orchestrator reviewer, and
the operator, which is the right audience for a cross-pair seam.

Watch item -- design-complete reports must cite the live COORD state.

This is not a blocker for dispatch approval. It is a lock-time review requirement: any m-5 or m-6
design-complete claim must cite the current COORD disposition for human-mode vocabulary, interjection host, and
the m-1 away-mode bridge answer or blocker. A design doc that treats those as locally solved without the COORD
record should be held at design review.

Approved next actions:
- m-5 and m-6 may proceed with the DESIGN dispatches under the issued scope.
- Run the c3-design-m5-m6-coord thread before either domain claims design lock.
- Keep `GRILL_REQUIRED: yes` for both domain designs.
- Preserve the c3 close integration capstone as a pre-close composition gate.

Not authorized:
- no PLAN / Step-1 build decomposition;
- no implementation, source/pcode edit, merge, branch, commit, PR, or live-send behavior;
- no concrete archetype/surface value lock before design review and grill closure;
- no away-mode verdict-token bridge design lock without the m-1 answer or an explicit blocker/human decision;
- no c3 close until the m-5/m-6 designs compose with the locked m-1..m-4 contracts.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
