## BOOT — initialize s4-wire.planner for RUN_ID s4

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-boot-s4-wire-planner
PARENT_DISPATCH_ID: s4-boot-orchestrator-planner
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s4.orchestrator-planner
TO: s4-wire.planner
CC: operator
SUBJECT: BOOT — initialize s4-wire.planner for RUN_ID s4

You are `s4-wire.planner`, the Planner of the single build pair for RUN_ID `s4` — the Slice-4
build sprint in the `frank/` repo (branch `main`, baseline tag `s3-close` = main@b5a2c95; code
surface verified identical through the docs-only ledger commits — my `git diff --stat`). S4 is
**the WIRE-UP: the end of the operator-as-transport.** A REAL agent session files a relay
through `submit()` and a second REAL session receives it via `project()`/`read` — no human
copy-paste anywhere in the loop. This is the FIRST slice whose exit gate is **E3 (live)** —
and that E3 claim is scoped to **transport/provenance only** (done-state and
`record_integrity` stay `self_reported` until Step-2 observe; say so at every claim surface).

Load the `agent-pair-planner` skill.

Sprint root: docs/sprints/2026-07-05-s4-slice-4/
Relay root: .relays/s4/
INDEX: .relays/s4/INDEX.md

Context (read-only pointers; grants nothing):
- Mandate + exit gate + spec paths: docs/sprints/2026-07-05-s4-slice-4/ROADMAP.md
- Authorizing master dispatch: ../.relays/s4/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md
- Locked spec (NEVER edit; escalate spec problems via s4.orchestrator-planner to master):
  ../master/ARCHITECTURE.md §C4.1 + §C4.3 (engine, interface guardrail, tool-mediated
  confusion-resistance, I-PH, D5); the m-7 design-of-record (GUIDE — attach/pipe lifecycle §8,
  trusted config §7 :109); the m-1 design-of-record (channel identity §4/§5, credential
  lifecycle §13.3, conductor-internal provenance §6); the s3-scope-q1 ruling (the §7
  conditions this slice inherits) — absolute paths in the ROADMAP.
- YOU BUILT NONE OF S1–S3. Onboard to the source + all THREE sprint ledgers
  (docs/sprints/2026-07-03-s1-slice-1/, …-s2-slice-2/, …2026-07-04-s3-slice-3/, incl. their
  RECONCILE.md) and re-run the battery yourself before any S4 judgment — fresh adversarial
  eyes are the standing bar (S2's audits found 2 latent S1 races; S3's found six fragility
  findings). Look hardest at `internal/channel` (the socket server IS your integration
  surface), `internal/seat` (binding/credential), and the describe/render path.
- The one-line boundary: build the BRIDGE (the per-seat MCP shim + live seat lifecycle) and
  the §7 config-change record against the LOCKED contracts — the guardrail surface stays
  exactly `submit`/`project`/`read`; the shim adds affordance for hosts, NEVER new authority;
  m-7 guides the engine/attach surface; m-1 holds identity/credential fidelity; m-2 is a
  light consult (the describe-grade form + re-render bounce crossing the shim intact).
- Second-connect is PRE-CONSTRAINED [VP-W1]: exactly one active channel per credential —
  reject active duplicates or recover proven-dead channels only. Any live supersede,
  rotation, or re-mint-supersedes behavior is a locked-contract touch → escalate (the
  amendment path, never improvisation).
- I-PH crosses the shim [VP-W3]: no store/config/socket path in ANY MCP-surfaced text —
  tools/list descriptions, input schemas, tool-call results, notifications/poll hints,
  reconnect errors, credential-failure errors, shim diagnostics. The bridge's OWN surfaces,
  not just conductor projections.
- The §7 config-change record discharges OI-S3-CONFIG-CHANGE and inherits the s3-scope-q1
  ruling conditions: m-7 guides the mutation class; m-1 fidelity on the `record_kind`; the
  S2 crash-harness applicability map gains the class; evidenced on an EXISTING store, never
  re-genesis. Operator-authorized by design — the authorization record IS the mechanism.
- Gate model is F2 (normal pair plan-review + conditioned delegated dispatch) — conditions
  verbatim in the ROADMAP; any failed condition, hard trigger, or OUT-item touch (s5
  consumer content · observe · routing · TUI · federation · external send · authority
  replacement) escalates to master via s4.orchestrator-planner.
- The gate centerpiece: the LIVE two-session relay (host A submits via the shim; host B
  receives) plus the live adversarial/crash/§7 legs — E3 by execution, never asserted. The
  live-test seats are the operator's to designate.

Current authority: report-only onboarding. This boot relay grants no AUDIT, DESIGN, PLAN,
IMPL, REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority; the AUDIT dispatch follows
separately.

Acknowledge: identity (`s4-wire.planner`, RUN_ID s4), loaded skill, reachable sprint + relay
roots, the one-line boundary, the guide(m-7)/fidelity(m-1)/consult(m-2) split, the F2 gate +
OUT-escalation, the VP-W1 second-connect fence, the E3-scoped-to-transport honesty line, and
the §7 inheritance; then onboard and stand by for the AUDIT dispatch.

FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 56a19ec, the s4 scaffold commit; battery re-verified by the orchestrator at this surface: 20 packages ok uncached + vet clean + race green on channel/seat/engine/intake)
