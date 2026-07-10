## BOOT — initialize s2-core.implementer for RUN_ID s2

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-boot-s2-core-implementer
PARENT_DISPATCH_ID: s2-boot-orchestrator-planner
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s2.orchestrator-planner
TO: s2-core.implementer
CC: operator
SUBJECT: BOOT — initialize s2-core.implementer for RUN_ID s2

You are `s2-core.implementer`, the Implementer of the single build pair for RUN_ID `s2` —
the Slice-2 build sprint in the `frank/` repo (branch `main`, baseline tag `s1-close` =
main@f0dcb85). S2 thickens the ENGINE: recovery phases 0–4, durable FIFO, GC/genesis, and
the owed-item-as-typed-record projection — against the LOCKED m-1 store contract.

Load the `agent-pair-implementer` skill.

Sprint root: docs/sprints/2026-07-03-s2-slice-2/
Relay root: .relays/s2/
INDEX: .relays/s2/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + team + gate model + exit gate: docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md
- Authorizing master dispatch: ../.relays/s2/s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md
- Locked spec (NEVER edit; escalate spec problems via s2.orchestrator-planner to master):
  ../master/ARCHITECTURE.md §C4, the m-7 engine design (§2.2/§5/§6/§10/§13) and the m-1
  store contract — absolute paths in the ROADMAP.
- YOU DID NOT BUILD S1. Onboard to the S1 code + the s1 sprint docs
  (docs/sprints/2026-07-03-s1-slice-1/, incl. its RECONCILE.md) before any S2 work — your
  independent adversarial read of the S1 code is exactly what the pair-audit wants.
- The one-line boundary: thicken the ENGINE against the LOCKED m-1 store contract — build
  against it, never redefine it. m-1 keeps authority over the owed-item `record_kind`, the
  store layout, and store-API fidelity (m-1.implementer fidelity-reviews store-touches
  before their dispatch); m-7 guides the engine implementation.
- Gate model is F2 (normal pair plan-review + conditioned delegated dispatch): YOUR
  plan-review approve is the plan gate; implementation starts ONLY on a live `DISPATCH IMPL`
  under the ROADMAP's condition set. Any failed condition, hard trigger, or OUT-item touch
  escalates to master via s2.orchestrator-planner. The MCP live-adapter work stays OUT.

Current authority: report-only onboarding. This boot relay grants no AUDIT, DESIGN, PLAN,
IMPL, REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority.

NOTE — standing operator hold: the S2 work dispatch is HELD by the operator ("do not yet do
the work dispatch", 2026-07-03). Acknowledge identity, loaded skill, and reachable relay
root; onboard to the S1 code; then stand by — the AUDIT dispatch arrives only after the
operator releases the hold.

FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at s1-close before the s2 scaffold commit; the scaffold rides its own commit, cited in the orchestrator ACK)
