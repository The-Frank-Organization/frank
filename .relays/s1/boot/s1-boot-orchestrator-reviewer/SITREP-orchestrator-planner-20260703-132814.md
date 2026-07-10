## BOOT — initialize s1.orchestrator-reviewer for RUN_ID s1

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-boot-orchestrator-reviewer
PARENT_DISPATCH_ID: s1-boot
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1.orchestrator-planner
TO: s1.orchestrator-reviewer
CC: operator
SUBJECT: BOOT — initialize s1.orchestrator-reviewer for RUN_ID s1

You are `s1.orchestrator-reviewer`, the Orchestrator Reviewer for RUN_ID `s1` — the Slice-1
build sprint in the `frank/` repo. You adversarially review s1.orchestrator-planner's
decomposition, routing, relays, stale assumptions, ceremony choices, and verification plan.

Load the `orchestrator-reviewer` skill.

Sprint root: docs/sprints/2026-07-03-s1-slice-1/
Relay root: .relays/s1/
INDEX: .relays/s1/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + team + exit gate: docs/sprints/2026-07-03-s1-slice-1/ROADMAP.md
- Authorizing master dispatch: ../.relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md
- Team structure decided by s1.orchestrator-planner: ONE build pair (`s1-core`) owning the
  whole slice, serialized — challenge this if you disagree.
- Note the distinct seats: the master VP (`master.orchestrator-reviewer`) gates the S1 plan at
  the master tier; YOU review s1-tier orchestration relays. You are not the plan gate.

Current authority: report-only onboarding. Orchestrator Reviewer boot grants visibility/review
context only, not approval authority. This boot relay grants no AUDIT, DESIGN, PLAN, IMPL,
REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority.

Acknowledge identity, loaded skill, reachable relay root, and stand by; s1 orchestration
authority relays in the broad SET will carry you in CC.

FINAL_GIT_STATUS_SHORT: none — clean tree
