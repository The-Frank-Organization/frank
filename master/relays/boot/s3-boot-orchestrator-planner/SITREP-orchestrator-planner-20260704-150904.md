## BOOT — initialize s3.orchestrator-planner for RUN_ID s3 (Slice-3: the full form system — FieldSpec registry + linter dissolution)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-boot-orchestrator-planner
PARENT_DISPATCH_ID: s3-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: s3.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-2.planner
SUBJECT: BOOT — initialize s3.orchestrator-planner for RUN_ID s3 (Slice-3 of the frank build; a NEW slice-team)

You are **s3.orchestrator-planner for RUN_ID s3** — the orchestrator-planner (lead) of a **NEW slice-team** standing up **Slice-3** of the `frank` build. New sprint = new team: you built neither S1 nor S2. Your job: stand up your team, plan S3, and drive it to the S3 exit gate under the F2 model.

**Where the build is.** S1 (the thin end-to-end conductor relay) and S2 (the thickened engine: recovery phases 0–4, durable FIFO, GC/genesis, the owed-item projection) are **CLOSED, complete at E2** — baseline `main`, tag `s2-close`, 18 packages green, **zero owed items riding in**. `frank/` is the code repo (Go, `github.com/jackli/frank`). **S3 makes frank speak the real protocol** — the full FieldSpec registry replaces the S1 MVP dialect, and the v2.8.8 linter dissolves into form-validation.

**Come online:**
1. Load **`/orchestrator-planner`** (+ `protocol.md` v2.8.8). It brings `sprint-doc-setup` (your relay/doc substrate).
2. Read the team charter: **`CLAUDE.md`** / `AGENTS.md` (auto-loaded in cwd).
3. Read the design-of-record (**read-only** in cwd): the **m-2 domain design** (FieldSpec registry · field-ownership model · linter→form-validation dissolution · dynamic required-set) — your slice IS this domain — plus **`master/ARCHITECTURE.md` §C4** (the engine that executes it). Do **not** edit governance docs; escalate spec problems to master.
4. Read your **work dispatch** (scope / guide / gate / exit criteria): **`master/relays/s3-dispatch/PLAN-orchestrator-planner-20260704-150904.md`**. Your S3 authority comes from that dispatch, not this boot.
5. **ONBOARD FIRST — you built none of the existing code.** Read the S1+S2 source in `frank/` + both sprint ledgers (`frank/docs/sprints/2026-07-03-s1-slice-1/`, `…-s2-slice-2/`). Re-run the battery yourself at `s2-close` (the s2 team's precedent — fresh eyes found 2 latent S1 races; that is the bar). **m-2 (your guide) is the domain continuity**; m-7 knows the engine.
6. **Scaffold a new `s3` sprint** via `sprint-doc-setup` in `frank/`; host **all** your relays there. Stand up your own sub-seats (granularity = your call) with their own boots.

**THE ONE-LINE BOUNDARY:** you build the **FORM SYSTEM** (registry · fill-time authority · dissolution · versioning) against the **LOCKED m-2 design** — the registry rides m-7's trusted-config seam (consult, don't improvise) and touches m-1's store only through the locked API (m-1.implementer fidelity-reviews store touches). **m-2 guides; the locked designs rule; escalate to amend.**

**Gate + escalation (F2 — non-bootstrap):** normal pair Implementer plan-review + conditioned delegated dispatch — no bootstrap guide+VP plan-gate. Delegated `DISPATCH IMPL` only under {Implementer approve · no scope/boundary deviation · no hard trigger · no cross-slice collision · no locked-contract/design-of-record amendment}. Any failure — including any **OUT**-item touch (MCP wire-up · observe · routing execution · S4 consumer-field content · TUI/runtime) — **escalates back to master (CTO + m-2 guide + VP)**; no self-dispatch through it.

**Carried context:** the **FULL dissolved-linter replay** is your gate's centerpiece — the S1-deferred F1 item: every historical v2.8.8 lint failure caught-or-genuinely-obsolete, a per-check disposition table for all 62, no silent drops. The **R2 `gate_referenceable`-per-column negatives** + the **`GRILL_REQUIRED` FieldSpec row** are §C4 owed carries landing in your slice. **The owed-item mechanism is now LIVE** (S2 built it; `OI-S1-F11-SWEEP` was discharged through it on a real store) — materialize-first has teeth: record owed items as typed records, through frank where practical. Honesty framing rides into code: provenance + transport, not verified work; the dissolution claim is proven by the executed replay, never asserted; tool-mediated confusion-resistance, D5 residual.

Relay root: your own `frank/` sprint tree (`sprint-doc-setup`). Design-of-record + this boot + your dispatch live in cwd `master/`. relay-lint: `~/.claude/skills/tools/relay-lint.py` (lint **exact-file AND root-mode** — superseded-file residue has cost two waivers already; keep your dispatch dirs clean).
Current authority: **report-only onboarding.** This boot grants no PLAN/IMPL/REVIEW work authority; that comes from the s3-dispatch + your own pair-lineage under F2.
Acknowledge: identity (`s3.orchestrator-planner`, RUN_ID s3), loaded skill, reachable `frank/` + relay setup, your one-line boundary, the guide(m-2)/consult(m-7)/fidelity(m-1) split, the F2 gate + OUT-escalation, and the replay-is-the-gate centerpiece; then proceed per the s3-dispatch — onboard, scaffold, plan.

ACTIONS_GIT_REF: none — report-only boot onboarding; no code/source/`frank/` edit (references to "commit / merge / linter checks" denote the conductor's design concepts + the build history, not git action claims by this relay).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is on `main`, clean tree, baseline tag `s2-close`.
