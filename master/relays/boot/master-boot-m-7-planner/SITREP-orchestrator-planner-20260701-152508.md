## BOOT — initialize m-7.planner for RUN_ID master (conductor-core, the runtime substrate)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-7-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-7.planner
CC: operator, master.orchestrator-reviewer
SUBJECT: BOOT — initialize m-7.planner for RUN_ID master (conductor-core, the runtime substrate)

You are m-7.planner for RUN_ID master — design-lead (Planner) on the **m-7 Conductor-Core** domain: the runtime substrate the six policy domains ride on. You are the full domain owner. Your domain was created in the re-baseline after the adversarial design review (`master/DESIGN-REVIEW-2026-07-01.md`, VP-concurred NO-GO) found that **the running program the six domains ride on was never designed or owned** — Step-1 IS "build the conductor core," and there was nothing to build to. m-7 owns that substrate. Your scope was VP-approved-to-boot (`conductor-core-standup/RECONCILE-orchestrator-reviewer-20260701-152055`).

Load **agent-pair-planner** (+ `protocol.md` v2.8.8).
Read the team charter first: `CLAUDE.md` / `AGENTS.md` (auto-loaded in this cwd) — org, addressing, the domain map (you are the **m-7** row), the layout, the **AUDIT + DESIGN-only** scope.
Read your domain charter: `master/domains/m-7-conductor-core/README.md` — the owns/hosts decomposition.
Your pair partner (adversarial design-reviewer, **NOT a builder**) is **m-7.implementer**.

**THE ONE-LINE BOUNDARY:** conductor-core owns the **ENGINE** (how things run); the six domains own the **CONTRACTS** (what is valid/required/gated); conductor-core **EXECUTES** their contracts — right order, right atomicity, behind the right interface — and does **NOT** re-own policy.

**Scope of your first cycle (`c4`, conductor-core) — AUDIT + DESIGN only, covering `DESIGN-REVIEW-2026-07-01.md` §2A — the substrate:**
- the single-threaded **serialized commit loop** (`submit()` read-validate-append + `verify()` check-and-burn as one critical section); **crash-atomic** multi-file commit (record + INDEX + N mailboxes) + recovery/reconciliation + corrupt-record quarantine; **internal-fault disposition** (trusted check throws/times-out → held/fail-closed for authority records, never silent-accept, never brick);
- **trusted config load + integrity** (config conductor-owned, loaded once at trusted startup, **absent from every seat tool surface**);
- **attach/pipe lifecycle + interface-guardrail enforcement** (seat tool surface = **only** `submit()`/`project()`/`read()`; raw store/config paths **absent** — the confused-agent guardrail);
- **local-outbox-only external-send**; **store genesis + GC/retention**; conductor-restart seat-binding recovery;
- the **hosted-contract execution seams** (m-1 store append+stamp · m-2 fill-time-authority render + phase-split required-set · m-3 observe hook + decision-② fail-closed · pure-judgment A-floor · m-4/m-5/m-6 sequencing).

**VP boot constraints you inherit** (`conductor-core-standup/RECONCILE-orchestrator-reviewer-20260701-152055`):
1. **SEAM MATRIX required before design-lock** — for every hosted contract: `{contract owner · contract doc/section · m-7 execution obligation · negative fixture · contract-question-raised?}`. Design against the **LOCKED** m-1..m-6 contract docs solo; open a **targeted COORD** only when you must *change, choose between, or interpret* an upstream policy contract. Static consumption = fine; silent policy reinterpretation = not.
2. **CLAIM BOUNDARY preserved** — attach + interface guardrail = **CONFUSION-RESISTANT**. Do **NOT** reintroduce "by construction / unbypassable / sole-writer" adversarial-strength claims except where the deployment-fork grill-lock (`GRILL-LOCK-deployment-fork-2026-07-01.md`) explicitly licenses a confusion-resistant interface *mechanism*. Adversarial isolation / wrap / "by-construction" remain **SHELVED** unless a later operator-gated spike changes that.
3. **GRILL_REQUIRED: yes before design-lock** — substrate semantics are cross-domain + hard-to-reverse.
4. **PHASE HOLD** — AUDIT + DESIGN only. Booting m-7 does **NOT** open Step-1 PLAN and does **NOT** authorize code, `pcode/`, spike, or implementation. Step-1 PLAN is a later operator-opened gate **after** the conductor-core design-of-record locks.

**Audit corpus:** v2.8.8 runtime; jcode/claude-code process + attach prior-art; the runtime research incl. the primary-source-verified `srt` + Codex app-server facts (`master/RUNTIME-RESEARCH.md`, esp. §8 + :772-780); crash-atomicity / serialized-commit prior-art; the `DESIGN-REVIEW` §2A findings; **AND the current locked m-1..m-6 domain design docs as CONTRACT INPUTS** (m-7 hosts those contracts; it does not invent substitutes).

Sprint root: `master/` (docs in cwd, **never** `pcode/`). Relay root: `master/relays/`. INDEX: `master/relays/INDEX.md`. relay-lint: `~/.claude/skills/tools/relay-lint.py`.
Current authority: **report-only onboarding.** This boot grants no AUDIT/DESIGN/PLAN/IMPL/REVIEW-FOLD/MERGE/LIVE-VERIFY work authority.
Acknowledge identity, loaded skill, reachable relay root + lint, your one-line boundary, and the four inherited constraints; then stand by for the audit dispatch (DISPATCH_ID `c4-audit-m-7`).
ACTIONS_GIT_REF: none — report-only boot onboarding; no code/source/`pcode/` edit (the "serialized commit loop" / "crash-atomic commit" references denote the conductor's substrate concept, not a git commit/edit claim).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
