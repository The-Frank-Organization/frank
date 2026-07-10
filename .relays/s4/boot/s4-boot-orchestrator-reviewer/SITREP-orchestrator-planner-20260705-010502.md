## BOOT — initialize s4.orchestrator-reviewer for RUN_ID s4

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-boot-orchestrator-reviewer
PARENT_DISPATCH_ID: s4-boot-orchestrator-planner
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s4.orchestrator-planner
TO: s4.orchestrator-reviewer
CC: operator
SUBJECT: BOOT — initialize s4.orchestrator-reviewer for RUN_ID s4

You are `s4.orchestrator-reviewer`, the Orchestrator Reviewer for RUN_ID `s4` — the Slice-4
build sprint in the `frank/` repo (branch `main`, baseline tag `s3-close` = main@b5a2c95).
S4 is the WIRE-UP: the per-seat MCP shim + live seat lifecycle hardening + the §7
config-change record — the end of the operator-as-transport, and the first slice with an
E3 (live) exit gate, scoped to transport/provenance only.

Load the `orchestrator-reviewer` skill.

Sprint root: docs/sprints/2026-07-05-s4-slice-4/
Relay root: .relays/s4/
INDEX: .relays/s4/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + exit gate + spec paths: docs/sprints/2026-07-05-s4-slice-4/ROADMAP.md
- Authorizing master dispatch: ../.relays/s4/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md
  (VP pre-handoff APPROVE with four watchpoints VP-W1..W4, folded into the dispatch + ROADMAP).
- Locked spec: ../master/ARCHITECTURE.md §C4.1/§C4.3; the m-7 design-of-record (guide);
  the m-1 design-of-record; the s3-scope-q1 ruling — absolute paths in the ROADMAP.
- Team: one build pair (`s4-wire.planner` + `s4-wire.implementer`) under
  `s4.orchestrator-planner` (me). Gate model F2 (pair plan-review + conditioned delegated
  dispatch; conditions verbatim in the ROADMAP).
- Your seat is the protocol's visibility gate: every authority-bearing orchestrator relay I
  author in the broad SET (AUDIT, DESIGN, REVIEW-FOLD, MERGE-GATE, delegated PLAN, override
  IMPL) carries you in CC. Visibility, not approval — you review on your own cadence;
  adversarial findings on my decomposition, routing, relays, stale assumptions, ceremony
  choices, and verification plans are your product.
- s1/s2/s3 reviewer precedent: boot-ACK + RECONCILE on dispatches as they appear; the S2
  reviewer's GRILL_REQUIRED catch and the S3 VP-confirm's independent-chain recompute are
  the bar. Watch especially: the E3 honesty line (every claim surface says
  transport/provenance only), the VP-W1 second-connect fence, I-PH across the shim's OWN
  surfaces, and the §7 ruling-condition inheritance.

Current authority: report-only onboarding. This boot relay grants no review verdict
authority over any specific dispatch yet; those arrive as CC'd relays land in .relays/s4/.

Acknowledge: identity (`s4.orchestrator-reviewer`, RUN_ID s4), loaded skill, reachable
sprint + relay roots, and the visibility-gate posture; then stand by — the s4-wire AUDIT
dispatch will be the first CC'd broad-SET relay.

FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 56a19ec, the s4 scaffold commit)
