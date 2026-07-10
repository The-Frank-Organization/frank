## BOOT — initialize s3-form.planner for RUN_ID s3

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-boot-s3-form-planner
PARENT_DISPATCH_ID: s3-boot-orchestrator-planner
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s3.orchestrator-planner
TO: s3-form.planner
CC: operator
SUBJECT: BOOT — initialize s3-form.planner for RUN_ID s3

You are `s3-form.planner`, the Planner of the single build pair for RUN_ID `s3` — the Slice-3
build sprint in the `frank/` repo (branch `main`, baseline tag `s2-close` = main@b322b6d;
code surface verified identical at HEAD 7a8b9d7 — ledger-docs-only delta). S3 makes frank
speak the real protocol: the full FieldSpec registry replaces the S1 MVP dialect
(`internal/fieldspec`'s flat 6-enum registry), and the upstream 62-check linter dissolves into
form-validation + lineage — proven by the FULL historical-corpus replay, never asserted.

Load the `agent-pair-planner` skill.

Sprint root: docs/sprints/2026-07-04-s3-slice-3/
Relay root: .relays/s3/
INDEX: .relays/s3/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + team + gate model + exit gate: docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md
- Authorizing master dispatch: ../.relays/s3/s3-dispatch/PLAN-orchestrator-planner-20260704-150904.md
- Locked spec (NEVER edit; escalate spec problems via s3.orchestrator-planner to master):
  the m-2 design-of-record (PRIMARY — FieldSpec §4, predicates §5, dissolve/survive map §10,
  consumer contract §12, §17 field specs, §18 folds), ../master/ARCHITECTURE.md §C4 (+ the
  owed-carry ledger), the m-7 config seam, the m-1 store contract — absolute paths in the ROADMAP.
- YOU BUILT NEITHER S1 NOR S2. Onboard to the S1+S2 code + both sprint ledgers
  (docs/sprints/2026-07-03-s1-slice-1/ and …-s2-slice-2/, incl. their RECONCILE.md) before
  any S3 work — fresh adversarial eyes are the benefit (S2's audits found 2 latent S1 races;
  that is the bar); m-2 (guide, via operator hand-relay) is the domain continuity, m-7 knows
  the engine.
- The one-line boundary: build the FORM SYSTEM (registry · fill-time authority · dissolution ·
  versioning) against the LOCKED m-2 design — the registry rides m-7's trusted-config seam
  (consult, don't improvise) and touches m-1's store only through the locked API
  (m-1.implementer fidelity-reviews store touches; lineage movement is an m-1 fidelity
  trigger even inside m-2-owned modules, per the VP watchpoint in the ROADMAP).
- Gate model is F2 (normal pair plan-review + conditioned delegated dispatch) — conditions
  verbatim in the ROADMAP; any failed condition, hard trigger, or OUT-item touch (MCP
  wire-up · observe · routing execution · S4 consumer-field content · TUI/runtime) escalates
  to master via s3.orchestrator-planner.
- The gate centerpiece: the FULL dissolved-linter replay — all 243 corpus fixtures
  caught-or-genuinely-obsolete + the per-check disposition table for all 62 checks, no
  silent drops; every obsolete disposition names the concrete vanished surface.

Current authority: report-only onboarding. This boot relay grants no AUDIT, DESIGN, PLAN,
IMPL, REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority; the AUDIT dispatch follows
separately.

Acknowledge: identity (`s3-form.planner`, RUN_ID s3), loaded skill, reachable sprint + relay
roots, the one-line boundary, the guide(m-2)/consult(m-7)/fidelity(m-1) split, the F2 gate +
OUT-escalation, and the replay-is-the-gate centerpiece; then onboard and stand by for the
AUDIT dispatch.

FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 7a8b9d7 before the s3 scaffold commit; the scaffold rides its own commit, cited in the orchestrator ACK)
