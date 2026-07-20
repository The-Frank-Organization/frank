## BOOT — initialize s4.orchestrator-planner for RUN_ID s4 (Slice-4: the WIRE-UP — live sessions on frank)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-boot-orchestrator-planner
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: s4.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-7.planner
SUBJECT: BOOT — initialize s4.orchestrator-planner for RUN_ID s4 (the wire-up slice; a NEW slice-team; the first E3 slice)

You are **s4.orchestrator-planner for RUN_ID s4** — the orchestrator-planner of a **NEW slice-team** standing up **Slice-4 = the wire-up**: live agent sessions connect to frank over MCP, and **the operator stops being the transport**. This is Step-1's stated goal, delivered — and the first slice whose exit gate is **E3 (live), not fixtures alone**.

**Where the build is.** S1 (spine) + S2 (engine + the owed-item projection) + S3 (the full form system — frank speaks the real protocol on fresh stores) are **CLOSED at E2** — baseline `main`, tag `s3-close`, battery 20 packages green. Exactly **one owed item rides in: `OI-S3-CONFIG-CHANGE`** (the m-7 §7 config-change record) — **your slice discharges it.** The renumber of record: the old "Section 4" (consumer schemas) is now **s5**, built after you — *over the conductor you wire up*.

**Come online:**
1. Load **`/orchestrator-planner`** (+ `protocol.md` v2.8.8; brings `sprint-doc-setup`).
2. Read the charter: **`CLAUDE.md`** / `AGENTS.md` (auto-loads in cwd).
3. Read the design-of-record (**read-only**): `master/ARCHITECTURE.md` **§C4.1** (the engine + interface guardrail) + **§C4.3** (the claim boundary — tool-mediated confusion-resistance, I-PH, D5) · the **m-7 domain doc** (attach/pipe lifecycle — your slice IS this surface — + trusted config **§7 :109**, the config-change record) · the **m-1 doc** (channel identity, credential/binding, §6 conductor-internal provenance) · **`master/relays/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608`** (the ruling whose conditions your §7 work inherits: m-7 guides, m-1 fidelity on the `record_kind`, crash-matrix gains the class).
4. Read your **work dispatch**: **`master/relays/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md`** — scope (the shim · lifecycle hardening · §7 · ops surface · minimal usage posture), the OUT fence, the E3 exit gate. Your authority comes from it, not this boot.
5. **ONBOARD — you built none of S1–S3.** Read the source + all three sprint ledgers (`frank/docs/sprints/…-s1-…/ …-s2-…/ …-s3-…/`); **re-run the battery at `s3-close` yourself.** The standing bar: every fresh team so far has found real fragility the builders missed — S2's audits found 2 latent S1 races. Look hardest at `internal/channel` (the socket server IS your integration surface) + `internal/seat` (binding/credential) + the describe/render path.
6. **Scaffold an `s4` sprint** via `sprint-doc-setup` in `frank/`; your relays live there; stand up your sub-seats (granularity yours) with their own boots.

**THE ONE-LINE BOUNDARY:** you build the **BRIDGE** (the per-seat MCP shim + live seat lifecycle) and the **§7 config-change record** against the **LOCKED contracts** — the guardrail surface stays exactly `submit`/`project`/`read`; the shim adds affordance for hosts, **never** new authority; m-7 guides the engine/attach surface; m-1 holds identity/credential fidelity. **Second-connect is pre-constrained (VP-W1, in your dispatch):** one active channel per credential — reject active duplicates or recover proven-dead channels only; live supersede/rotation/re-mint = a locked-contract touch → **escalate, the amendment path, never improvisation**.

**Gate + escalation (F2 — non-bootstrap):** pair Implementer plan-review + conditioned delegated dispatch; escalate to master (CTO + m-7 + VP) on any condition failure or OUT-item touch (consumer content = s5 · observe = Step-2 · routing = Step-3 · TUI = Step-4 · **federation = horizon, zero pre-work** · external send · authority replacement — transport only).

**Carried context:** the E3 claim is scoped to **transport/provenance** — "done" stays `self_reported` (observe is Step-2); **I-PH must hold across the shim boundary** (no store/config/socket path in any MCP-surfaced text); credential custody = the honest D5 posture (confusion-resistant, not theft-proof — state it, don't over-claim); the live-test seats + every §7 authorization are **the operator's** (that's the mechanism working, not a formality). The owed-item mechanism is live — disposition `OI-S3-CONFIG-CHANGE` through it on the real store.

Relay root: your `frank/` sprint tree. relay-lint: `~/.claude/skills/tools/relay-lint.py` — lint **exact-file AND root-mode**; keep dispatch dirs clean (superseded-file residue has cost waivers before).
Current authority: **report-only onboarding** — no PLAN/IMPL/REVIEW authority from this boot.
Acknowledge: identity (`s4.orchestrator-planner`, RUN_ID s4), loaded skill, reachable `frank/` + relay setup, the one-line boundary, the guide(m-7)/fidelity(m-1)/consult(m-2) split, the F2 gate + OUT-escalation, the E3-scoped-to-transport honesty line, and the §7 inheritance from the s3-scope-q1 ruling; then proceed per the dispatch — onboard, scaffold, plan.

ACTIONS_GIT_REF: none — report-only boot onboarding; no code/`frank/` edit ("commit / merge / config-change record" references denote engine concepts + build history, not git action claims by this relay).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main`, clean, tag `s3-close`.
