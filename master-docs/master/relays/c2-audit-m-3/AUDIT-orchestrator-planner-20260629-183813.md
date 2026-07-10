## Team m-3 — Observation & Evidence (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-audit-m-3
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-3.planner, m-3.implementer
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-3-observation-evidence
OWNER: m-3 (Observation & Evidence)

Phase scope — AUDIT (read-only), opening **Cycle c2**. Inspect source and docs, run safe read-only commands, and produce an independent paired audit with findings. Not in scope: any edits, branches, commits, PRs, scaffolding, or prototype code. This is still frank's research + design phase — no implementation exists or is authorized; IMPL would begin only on a future bare own-line `DISPATCH IMPL` to the Implementer, which is not part of this cycle.

Pair roles & research method:
- m-3.planner (Claude Opus 4.8, high thinking): lead the audit and surface design questions. Use multiple parallel agents, websearch, and a deep-research workflow for the prior-art sweep (observe-before-act / verification gating, evidence ladders + executable claims, egress/DLP content-safety scanning for agent output).
- m-3.implementer (GPT-5.5): run an INDEPENDENT audit — do not mirror the planner — and challenge or answer the design questions. No built-in deep-research skill, so use subagents + websearches.
- Independent paired audit: each member audits separately, then reconciles. The Planner does not spawn, direct, or simulate the Implementer; the two seats are independent operator-relayed sessions.

Domain context. m-3 owns **observation & evidence**: the observe-as-send-gate (the conductor reads lane ground-truth from *outside* the lane before a relay is deliverable), per-phase done-predicates, the evidence ladder (E-tiers), executable claims, and the fail-closed egress/content-safety gate (secrets / PII / model-names). The frank stance: a "done" is only as good as conductor-observed evidence; self-reported completion is *labeled*, never trusted.

You build on the LOCKED c1 contract (do not reopen it — `master/ARCHITECTURE.md`):
- R3 observe-integrity: every observed record carries `evidence_integrity {observed | self_reported}`; the m-3 hook is **observer-only** — a positive write-allowlist that writes only the closed m-3 observed/computed field set + a veto.
- m-1 DI-5 read-vantage: the conductor reads from outside the lane (your isolation guarantee).
- m-2 declares the observe/evidence FieldSpecs; m-3 fills them via the observer hook.

Sources to audit (cross-check the export's distillation against real source):
- The upstream protocol's current state: the upstream protocol release corpus (not vendored) — protocol.md's evidence/observe model, done-predicates, the orchestrator-review/observe gates in `<upstream relay-lint tools>/relay-lint.py`, and any "blocks-before-dispatch" evidence handling. Our LIVE store at `master/relays/` is a running instance.
- `references/jcode` (Rust): the SafetySystem review-queue + any observation/verification gating; wire-level attribution. Under `crates/` and `src/`.
- `references/claude-code` (TS): any verification/observation/permission hooks, tool-result gating, and content checks. Under `mcp-server/` and the agent sources.
- `references/agent-scripts` (Steinberger): the **Owner Decision Brief** + the **egress scan** — directly prior-art for m-3's egress/content-safety gate and evidence-summarization surface.
- Export design intent: the pre-build design-state export (not vendored), the adaptive-routing pillar doc (observe/evidence + the conductor's gate components) and `EXTERNAL-REFERENCES.md`.

Design question to resolve:
What is the minimal observe/evidence primitive that makes "done" **conductor-observed, not lane-asserted** — the observe-as-send-gate placement, the done-predicate + evidence-ladder shape, the executable-claims mechanism, and the fail-closed egress gate — such that it rides existing runtimes in Step 1 (observing lanes we don't yet run) yet extends to the standalone runtime later without re-cutting, and fills the locked c1 `evidence_integrity`/DI-5 seam exactly?

Hard acceptance criteria:
1. A 4-bucket verdict (still-open / already-closed / product-overlapped / recommended-next) on the frank m-3 observe/evidence primitive vs what the upstream protocol and the references already provide.
2. A named statement of the **self-reported-done gap** (where each surveyed system trusts an agent's own "done"/output unobserved) and exactly how observe-as-send + `evidence_integrity` closes it.
3. A design recommendation: the observe-gate API surface (the observer hook against the locked write-allowlist), the evidence ladder (E-tiers + executable claims), and the egress/content-safety gate.
4. A boundary contract naming what m-3 must expose to consumers — m-4 (observed evidence feeding the benchmark routing-quality loop, a later release), m-5 (per-archetype observe invariants), m-6 (gate mechanism + egress for the away-mode external bridge) — so DESIGN locks interfaces consumer-validated.

c2 GUARDRAILS (VP-set, `c2-decomp` 20260629-183247):
- Phase band = AUDIT + DESIGN only. No build / PLAN / IMPL.
- Focus on the LOCKED c1 contract; do **not** reopen m-1/m-2. Audit your own domain's prior art + mechanism.
- m-3↔m-4 SEAM: m-4's benchmark loop (a later release) consumes your observed evidence; a routing record may be an evidenced record. Flag seam findings; a dedicated `c2-*-coord` COORD sub-thread will own the seam at DESIGN (reconciled before any c2 lock). Lighter than c1's PARENT/lineage convergence.
- m-5 SEAM (explicit design risk): m-5 owns per-archetype observe invariants. Name where archetype tags would parameterize your observe gates. c2 cannot lock m-3 without a lock-time m-5 seam disposition — so surface, don't close, concrete archetype/tag semantics.
- m-6 is the warm consumer lens for c2 (engages at consumer-review, after draft designs).

Boundary contract:
- Writes: the observe-gate / evidence-ladder / egress-gate primitive that m-4/m-5/m-6 build on.
- Reads: the upstream protocol's evidence/observe handling; jcode/claude-code/agent-scripts gating + egress prior art; the export design intent; the locked c1 R3/DI-5 seam.
- Target entity: the frank observe/evidence primitive (a design recommendation, not code).
- Downstream consumer: m-4 (benchmark evidence), m-5 (archetype observe invariants), m-6 (gate + egress).
- Proof: E1 source citations (file:line) backing every gap and claim.
- No-consumer action: flag any proposed primitive with no downstream consumer.

AUDIT-FIRST GATES — may reject or narrow:
0. Duplicate/already-built check: does the upstream protocol or a reference already provide an observe-gate / egress scan we should promote rather than rebuild? If so, recommend promote/wire, do not rebuild.
1. If the observe-as-send primitive cannot ride an existing runtime in Step 1 (needs the standalone runtime to observe from outside the lane), say so and route the dependency to the orchestrator.

Out of scope:
- m-4 router/policy internals (sibling audit `c2-audit-m-4`), the m-1/m-2 foundation (locked), the TUI/email-client UX, and any code.

Deliverable:
A file-relay AUDIT report from each seat independently, then a reconciled view: the 4-bucket verdict; the self-reported-done gap statement; the observe/evidence + egress design recommendation; the consumer boundary contract; the m-3↔m-4 + m-5 seam notes; evidence levels (E1 cites); operator-judgment items or none. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT (read-only phase; cwd is not a git repo, so the structured unavailable form is expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
