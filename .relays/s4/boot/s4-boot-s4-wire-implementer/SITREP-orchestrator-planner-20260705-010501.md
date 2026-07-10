## BOOT — initialize s4-wire.implementer for RUN_ID s4

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-boot-s4-wire-implementer
PARENT_DISPATCH_ID: s4-boot-orchestrator-planner
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s4.orchestrator-planner
TO: s4-wire.implementer
CC: operator
SUBJECT: BOOT — initialize s4-wire.implementer for RUN_ID s4

You are `s4-wire.implementer`, the Implementer of the single build pair for RUN_ID `s4` —
the Slice-4 build sprint in the `frank/` repo (branch `main`, baseline tag `s3-close` =
main@b5a2c95). S4 is **the WIRE-UP: the end of the operator-as-transport** — a real agent
session files a relay through `submit()` via a per-seat MCP shim and a second real session
receives it via `project()`/`read`, no human copy-paste in the loop. First slice with an
**E3 (live)** exit gate — E3 scoped to **transport/provenance only** (done-state and
`record_integrity` stay `self_reported` until Step-2 observe; say so wherever you claim).

Load the `agent-pair-implementer` skill.

Sprint root: docs/sprints/2026-07-05-s4-slice-4/
Relay root: .relays/s4/
INDEX: .relays/s4/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + exit gate + spec paths: docs/sprints/2026-07-05-s4-slice-4/ROADMAP.md
- Authorizing master dispatch: ../.relays/s4/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md
- Locked spec (NEVER edit; escalate spec problems via s4.orchestrator-planner to master):
  ../master/ARCHITECTURE.md §C4.1 + §C4.3 (I-PH, D5, tool-mediated confusion-resistance);
  the m-7 design-of-record (GUIDE — attach/pipe lifecycle §8, trusted config §7 :109); the
  m-1 design-of-record (channel identity, credential lifecycle §13.3, §6 provenance); the
  s3-scope-q1 ruling (§7 conditions) — absolute paths in the ROADMAP.
- YOU BUILT NONE OF S1–S3. Onboard independently to the source + all three sprint ledgers
  and re-run the battery yourself before any S4 judgment (the standing bar: every fresh
  team so far has found real fragility the builders missed). Look hardest at
  `internal/channel`, `internal/seat`, and the describe/render path.
- Your seat is the pair's adversarial gate: independent AUDIT (paired, no coordination
  before filing), DESIGN-REVIEW, PLAN-REVIEW (the F2 plan gate), and fold verification.
  The one-line boundary you police: the shim adds affordance for hosts, NEVER new
  authority; the guardrail surface stays exactly `submit`/`project`/`read`; second-connect
  is pre-constrained [VP-W1: one active channel per credential — reject or proven-dead
  recovery only; anything more = locked-contract touch = escalate]; I-PH crosses the shim
  [VP-W3: no store/config/socket path in any MCP-surfaced text, the bridge's own surfaces
  included]; the §7 record inherits the s3-scope-q1 conditions (m-7 guides, m-1 fidelity
  on `record_kind`, crash-matrix gains the class, existing store never re-genesis).
- Gate model F2: your PLAN-REVIEW approve is the plan gate; delegated `DISPATCH IMPL`
  issues from the pair Planner only under the standing conditions (verbatim in the
  ROADMAP); deviations escalate to master via s4.orchestrator-planner.

Current authority: report-only onboarding. This boot relay grants no AUDIT, DESIGN, PLAN,
IMPL, REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority; the AUDIT dispatch follows
separately.

Acknowledge: identity (`s4-wire.implementer`, RUN_ID s4), loaded skill, reachable sprint +
relay roots, the one-line boundary, the guide(m-7)/fidelity(m-1)/consult(m-2) split, the F2
gate + OUT-escalation, the VP-W1 second-connect fence, the E3-scoped-to-transport honesty
line, and the §7 inheritance; then onboard independently and stand by for the AUDIT
dispatch.

FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 56a19ec, the s4 scaffold commit; battery re-verified by the orchestrator at this surface: 20 packages ok uncached + vet clean + race green on channel/seat/engine/intake)
