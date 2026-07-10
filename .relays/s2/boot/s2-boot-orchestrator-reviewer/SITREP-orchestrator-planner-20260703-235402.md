## BOOT — initialize s2.orchestrator-reviewer for RUN_ID s2

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-boot-orchestrator-reviewer
PARENT_DISPATCH_ID: s2-boot-orchestrator-planner
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s2.orchestrator-planner
TO: s2.orchestrator-reviewer
CC: operator
SUBJECT: BOOT — initialize s2.orchestrator-reviewer for RUN_ID s2

You are `s2.orchestrator-reviewer` for RUN_ID `s2` — the adversarial reviewer of
s2.orchestrator-planner's decomposition, routing, relays, ceremony choices, and
verification plans for the Slice-2 build sprint in `frank/` (branch `main`, baseline tag
`s1-close`). S2 thickens the ENGINE (recovery phases 0–4, durable FIFO, GC/genesis, the
owed-item projection) against the LOCKED m-1 store contract.

Load the `orchestrator-reviewer` skill.

Sprint root: docs/sprints/2026-07-03-s2-slice-2/
Relay root: .relays/s2/
INDEX: .relays/s2/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + team + gate model + exit gate: docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md
- Authorizing master dispatch: ../.relays/s2/s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md
- Per the protocol, every authority-bearing orchestrator relay in the broad SET will
  carry your address in CC — visibility, not approval; review on your own cadence.
- Watch-surfaces inherited from the dispatch: scope fence vs the OUT list (MCP live-adapter
  stays OUT), the F2 delegated-dispatch condition set, m-1 fidelity sequencing before any
  store-touch dispatch, cross-slice collision, and claim honesty (materialize-first: the
  owed-item projection guards recorded owed-items only).

Current authority: report-only onboarding. This boot relay grants no review verdicts on
work not yet dispatched and no work authority.

NOTE — standing operator hold: the S2 work dispatch is HELD by the operator ("do not yet do
the work dispatch", 2026-07-03). Acknowledge identity, loaded skill, and reachable relay
root, then stand by for CC'd broad-SET relays once the hold lifts.

FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at s1-close before the s2 scaffold commit; the scaffold rides its own commit, cited in the orchestrator ACK)
