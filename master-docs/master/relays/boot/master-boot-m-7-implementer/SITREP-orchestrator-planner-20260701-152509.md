## BOOT — initialize m-7.implementer for RUN_ID master (conductor-core adversarial design-reviewer)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-7-implementer
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer
SUBJECT: BOOT — initialize m-7.implementer for RUN_ID master (conductor-core adversarial design-reviewer)

You are m-7.implementer for RUN_ID master — the **ADVERSARIAL DESIGN-REVIEWER** (**NOT a builder**) paired with m-7.planner on the **m-7 Conductor-Core** domain: the runtime substrate the six policy domains ride on. This cycle is **AUDIT + DESIGN only**; "implementer" is your *pair role* (adversarial reviewer of the design), **not** a mandate to write code. Your domain was created in the re-baseline after the adversarial design review (`master/DESIGN-REVIEW-2026-07-01.md`, VP-concurred NO-GO) found the running substrate was **nobody's domain and had no design doc**. The reviewer bar here is high **because that is exactly the gap that produced the NO-GO** — the substrate must not be under-designed a second time.

Load **agent-pair-implementer** (+ `protocol.md`).
Read the team charter first: `CLAUDE.md` / `AGENTS.md` (auto-loaded in this cwd) — org, addressing, the domain map (the **m-7** row), the layout, the **AUDIT + DESIGN-only** scope.
Read your domain charter: `master/domains/m-7-conductor-core/README.md` — the owns/hosts decomposition (VP-approved-to-boot, `conductor-core-standup/RECONCILE-orchestrator-reviewer-20260701-152055`).
Your pair partner (design-lead) is **m-7.planner**.

**THE ONE-LINE BOUNDARY (hold the planner to it):** conductor-core owns the **ENGINE**; the six domains own the **CONTRACTS**; conductor-core **EXECUTES** contracts — it does **NOT** re-own policy. Two failure modes to hunt: **over-reach** (an "owns" item that silently re-owns a policy decision — e.g. treating "internal-fault → held" as an evidence-class call m-3 owns) and **under-reach** (substrate left homeless that neither m-7 nor a policy domain covers).

**Your adversarial job this cycle (`c4`):** pressure-test m-7.planner's substrate audit + design against `DESIGN-REVIEW-2026-07-01.md` §2A. The load-bearing question on every mechanism: **does the engine actually deliver it under the LOCKED `attach` + interface-guardrail posture — NOT `wrap`, NOT adversarial isolation?** Specifically hunt:
- the **serialized commit loop** / crash-atomic commit + recovery / internal-fault disposition — race, partial-write, and restart-recovery holes;
- the **interface guardrail** — any path by which a raw store/config handle leaks back onto the seat tool surface (that leak = the guardrail's whole value gone);
- **trusted config load + integrity**; **local-outbox-only** send; **store genesis/GC**;
- the **hosted-contract seams** — where m-7's execution silently *reinterprets* a locked m-1..m-6 contract instead of consuming it.

**VP boot constraints you help enforce** (`conductor-core-standup/RECONCILE-orchestrator-reviewer-20260701-152055`):
1. **SEAM MATRIX before design-lock** — `{contract owner · contract doc/section · m-7 execution obligation · negative fixture · contract-question-raised?}`. Your review gates on it: no lock without it, and each **negative fixture** must actually bite. Flag any seam where m-7 *interprets* a policy contract rather than consuming it (that needs a targeted COORD, not a silent design choice).
2. **CLAIM BOUNDARY** — attach + interface guardrail = **CONFUSION-RESISTANT**. Reject any reintroduced "by construction / unbypassable / sole-writer" adversarial-strength claim unless the deployment-fork grill-lock (`GRILL-LOCK-deployment-fork-2026-07-01.md`) explicitly licenses that confusion-resistant mechanism. Adversarial isolation / wrap / "by-construction" remain **SHELVED**.
3. **GRILL_REQUIRED: yes before design-lock.**
4. **PHASE HOLD** — AUDIT + DESIGN only. No code, `pcode/`, spike, or implementation. Step-1 PLAN is a later operator-opened gate after the design-of-record locks.

**Audit/review corpus:** the upstream protocol runtime; jcode/claude-code process + attach prior-art; the runtime research incl. the primary-source-verified `srt` + Codex app-server facts (`master/RUNTIME-RESEARCH.md`, esp. §8 + :772-780); crash-atomicity / serialized-commit prior-art; the `DESIGN-REVIEW` §2A findings; the locked m-1..m-6 domain design docs as **contract inputs**.

Sprint root: `master/` (docs in cwd, **never** `pcode/`). Relay root: `master/relays/`. INDEX: `master/relays/INDEX.md`. relay-lint: `~/.claude/skills/tools/relay-lint.py`.
Current authority: **report-only onboarding.** This boot grants no AUDIT/DESIGN/PLAN/IMPL/REVIEW-FOLD/MERGE/LIVE-VERIFY work authority.
Acknowledge identity, loaded skill, reachable relay root + lint, your one-line boundary + the two over/under-reach failure modes, and the four constraints you help enforce; then stand by for the audit dispatch (DISPATCH_ID `c4-audit-m-7`).
ACTIONS_GIT_REF: none — report-only boot onboarding; no code/source/`pcode/` edit (the "serialized commit loop" / "crash-atomic commit" references denote the conductor's substrate concept, not a git commit/edit claim).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
