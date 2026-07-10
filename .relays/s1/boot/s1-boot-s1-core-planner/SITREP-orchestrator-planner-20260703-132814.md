## BOOT — initialize s1-core.planner for RUN_ID s1

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-boot-s1-core-planner
PARENT_DISPATCH_ID: s1-boot
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: operator
SUBJECT: BOOT — initialize s1-core.planner for RUN_ID s1

You are `s1-core.planner`, the Planner of the single build pair for RUN_ID `s1` — the Slice-1
build sprint in the `frank/` repo (the conductor's greenfield build target; branch `main`,
first commit = the sprint scaffold only).

Load the `agent-pair-planner` skill.

Sprint root: docs/sprints/2026-07-03-s1-slice-1/
Relay root: .relays/s1/
INDEX: .relays/s1/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + team + exit gate: docs/sprints/2026-07-03-s1-slice-1/ROADMAP.md
- Authorizing master dispatch: ../.relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md
- Locked spec (NEVER edit; escalate spec problems via s1.orchestrator-planner to master):
  ../master/ARCHITECTURE.md §C4 + §C4.3/I-PH, the m-1 / m-2 / m-7 domain docs, ../master/STEP-1-KICKOFF.md
- Your pair owns the whole slice: `mint → connect → render-MVP-form → submit → stamp →
  validate → lineage → append(crash-atomic) → project → deliver → gate-outbox`, against the
  FROZEN m-1 store API + m-2 FieldSpec envelope (build against them; no re-design).
- Hard external gates ABOVE the normal pair lifecycle: the S1 plan is gated by m-7 (guide) +
  the master VP, and by m-1/m-2 fidelity review of our consuming surface, before any
  `DISPATCH IMPL` can go live. No code before those gates.

Current authority: report-only onboarding. This boot relay grants no AUDIT, DESIGN, PLAN,
IMPL, REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority.

Acknowledge identity, loaded skill, reachable relay root, and stand by for the next addressed
relay (an AUDIT dispatch from s1.orchestrator-planner).

FINAL_GIT_STATUS_SHORT: none — clean tree
