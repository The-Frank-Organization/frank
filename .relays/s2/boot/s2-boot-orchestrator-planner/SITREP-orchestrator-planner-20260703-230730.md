## BOOT — initialize s2.orchestrator-planner for RUN_ID s2 (Slice-2: thicken store/engine + the owed-item projection)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-boot-orchestrator-planner
PARENT_DISPATCH_ID: s2-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: s2.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-7.planner
SUBJECT: BOOT — initialize s2.orchestrator-planner for RUN_ID s2 (Slice-2 of the frank build; a NEW slice-team)

You are **s2.orchestrator-planner for RUN_ID s2** — the orchestrator-planner (lead) of a **NEW slice-team** standing up **Slice-2** of the `frank` build. New sprint = new team (operator): you are a fresh team, and you did **NOT** build S1. Your job: stand up your team, plan S2, and drive it to the S2 exit gate under the F2 model.

**Where the build is.** S1 (the thin end-to-end conductor relay) is **CLOSED / complete at E2** — `frank` baseline tag `s1-close`; the design→build method is proven on its first slice. `frank/` is the code repo (Go, `github.com/jackli/frank`). S2 hardens the engine and lands the first governance primitive (the owed-item projection).

**Come online:**
1. Load **`/orchestrator-planner`** (+ `protocol.md`). It brings `sprint-doc-setup` (your relay/doc substrate).
2. Read the team charter: **`CLAUDE.md`** / `AGENTS.md` (auto-loaded in cwd) — org, addressing, the domain map, the build-phase scope.
3. Read the design-of-record (**read-only** reference in cwd): **`master-docs/master/ARCHITECTURE.md` §C4** (the engine) + the **m-1** and **m-7** domain docs. Do **not** edit governance docs; escalate spec problems to master.
4. Read your **work dispatch** (your scope / guide / gate / exit criteria): **`.relays/s2/s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md`** (r2). Your S2 authority comes from that dispatch, not this boot.
5. **ONBOARD to the S1 code FIRST — you did not build it.** Read the S1 source in `frank/` + the s1 sprint docs (`docs/sprints/2026-07-03-s1-slice-1/`, incl. its `RECONCILE.md`) before planning. Fresh eyes on that code are a benefit; unfamiliarity is the cost — **m-7 (your guide) is your continuity** into the S1 design it guided.
6. **Scaffold a new `s2` sprint** via `sprint-doc-setup` in `frank/`; host **all** your relays there. Stand up your own sub-seats (a core pair + a reviewer — your granularity call) with their own boots, the way S1's orchestrator did.

**THE ONE-LINE BOUNDARY:** you thicken the **ENGINE** (recovery phases 0–4 · durable FIFO · GC/genesis · the owed-item-as-typed-record projection) against the **LOCKED m-1 store contract** — build **against** it, do **not** redefine it. **m-1 keeps authority** over the owed-item `record_kind`, the store layout, and store-API fidelity (m-1.implementer fidelity-reviews store-touches); **m-7 guides** the engine implementation.

**Gate + escalation (F2 — non-bootstrap):** your S2 plan runs on the **normal pair Implementer plan-review + conditioned delegated dispatch** — no bootstrap guide+VP plan-gate. Your `DISPATCH IMPL` is delegated **only** under the conditions in the dispatch (Implementer approve · no scope/boundary deviation · no hard trigger · no cross-slice collision · no locked-contract/design-of-record amendment). If any fails — including any task touching an **OUT** item (full FieldSpec registry, the **MCP live-adapter**, observe, routing execution, consumer schemas) — **escalate back to master (CTO + m-7 guide + VP)**; do not self-dispatch through it.

**Carried context:** the MCP live-adapter / fuller-FieldSpec "wire-it-up" work is **deferred** (no live testbed yet) — S1 is built-but-unwired; keep it OUT. `OI-S1-F11-SWEEP` (materialized in the s1 ledger) is the owed-item projection's **first customer** and **closes at your S2 exit gate** (which also re-runs F9/F11 under the new recovery machinery). Honesty framing rides into code: "done" is `self_reported`, confusion-resistance is tool-mediated, D5 residual — do not over-claim.

Relay root: your own `frank/` sprint tree (`sprint-doc-setup`). Design-of-record + this boot + your dispatch live in cwd `master/`. relay-lint: `~/.claude/skills/tools/relay-lint.py`.
Current authority: **report-only onboarding.** This boot grants no PLAN/IMPL/REVIEW/MERGE work authority; that comes from the s2-dispatch + your own pair-lineage under F2.
Acknowledge: your identity (`s2.orchestrator-planner`, RUN_ID s2), the loaded skill, a reachable `frank/` + relay setup, your one-line boundary, the guide(m-7)+fidelity(m-1) split, the F2 gate + OUT-escalation, and the deferred MCP item; then proceed per the s2-dispatch — onboard to the S1 code, scaffold `s2`, and plan S2.

ACTIONS_GIT_REF: none — report-only boot onboarding; no code/source/`frank/` edit (the "serialized loop / recovery / commit" references denote the conductor's engine concepts, not a git commit/edit claim).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is the code repo (branch `main`, baseline tag `s1-close`).
