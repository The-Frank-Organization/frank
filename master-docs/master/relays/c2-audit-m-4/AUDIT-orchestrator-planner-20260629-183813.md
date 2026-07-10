## Team m-4 — Routing & Policy (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-audit-m-4
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-4.planner, m-4.implementer
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-4-routing-policy
OWNER: m-4 (Routing & Policy)

Phase scope — AUDIT (read-only), part of **Cycle c2** (sibling of `c2-audit-m-3`). Inspect source and docs, run safe read-only commands, and produce an independent paired audit with findings. Not in scope: any edits, branches, commits, PRs, scaffolding, or prototype code. Still the frank research + design phase — no implementation exists or is authorized; IMPL would begin only on a future bare own-line `DISPATCH IMPL` to the Implementer, not part of this cycle.

Pair roles & research method:
- m-4.planner (Claude Opus 4.8, high thinking): lead the audit and surface design questions. Use multiple parallel agents, websearch, and a deep-research workflow for the prior-art sweep (model-routing / model-selection policies, capability-prior / model-card-driven dispatch, justified-deviation + audit-record patterns, LLM-routing benchmarks).
- m-4.implementer (GPT-5.5): run an INDEPENDENT audit — do not mirror the planner — and challenge or answer the design questions. No built-in deep-research skill, so use subagents + websearches.
- Independent paired audit: each member audits separately, then reconciles. The Planner does not spawn, direct, or simulate the Implementer; the two seats are independent operator-relayed sessions.

Domain context. m-4 owns **routing & policy** — the conductor's thesis. The model→seat router, capability priors, the routing record, justified deviation, and the benchmark + later-release feedback loop. Locked high-level decisions you build on: altitude **B** (the planner emits role+model per dispatch), policy **3-staged** (capability priors + justified deviation; benchmark feedback arrives in a later release). The differentiator: routing is a first-class, recorded, *justifiable* governance decision — not implicit config.

You build on the LOCKED c1 contract (do not reopen it — `master/ARCHITECTURE.md`):
- R2: a routing decision is a **separate seat-stamped routing relay** (m-2 FieldSpec + record-kind + accepted semantics); m-1 admits the accepted routing relay into the conductor-derived `parent_picker` candidate set for the dispatch it routes. The dispatch references it as **provenance/bookkeeping** — **model is NEVER a gate input** (no `model_*` predicate may enter the schema gate). This is load-bearing; your design must preserve it.
- §5 identity ≠ authority (ratified): m-1 owns *who* (the stamp); **m-4 owns *what a stamped seat may do*** — routing/policy keyed to the stamp (anti-confused-deputy).
- §J: `routing` is a category-B (orchestrator-absorbed) `gate_category`; model = payload.

Sources to audit (cross-check the export's distillation against real source):
- Export design intent — **READ FIRST**: the pre-build design-state export (not vendored) — the adaptive-routing pillar (the m-4 spec seed: altitude B, capability priors, the 3-staged policy, justified deviation, the benchmark loop in a later release) and `EXTERNAL-REFERENCES.md`.
- the stock protocol current state: the upstream protocol release corpus (not vendored) — what routing/model-selection (if any) the stock protocol has; expect this to be a GAP (single-model relay) and to confirm the still-open verdict. Our LIVE store at `master/relays/` shows hand-routing (role+model chosen per dispatch by the operator).
- `references/jcode` (Rust): **MultiProvider** per-seat routing, wire-level model attribution, and the swarm action that assigns model/role — the closest prior art to an actual router. Under `crates/` and `src/`.
- `references/claude-code` (TS): `agentType`/subagent routing, the model-selection surface, and any per-task model override. Under `mcp-server/` and the agent sources.
- `references/agent-scripts` (Steinberger): any model-selection / escalation logic.

Design question to resolve:
What is the minimal routing/policy primitive that makes model→seat routing a **first-class, recorded, justifiable governance decision** — altitude B (role+model per dispatch), the 3-staged policy (capability priors → justified deviation → benchmark feedback in a later release), and the routing record's shape — such that it preserves the locked R2 seam (routing = separate seat-stamped relay; model never a gate input) and the identity≠authority boundary, rides existing runtimes in Step 1 (recording decisions for runtimes we don't yet drive) and extends to executing the record on our own runtime later (Step 3) without re-cutting?

Hard acceptance criteria:
1. A 4-bucket verdict (still-open / already-closed / product-overlapped / recommended-next) on the frank m-4 routing/policy primitive vs what the stock protocol and the references already provide.
2. A named statement of the **implicit-routing gap** (where each surveyed system selects a model/role with no recorded, justifiable governance decision) and exactly how the routing record + justified deviation closes it.
3. A design recommendation: the router API surface, the capability-prior representation, the routing-record schema (as a consumer of the locked m-2 FieldSpec/R2 shape), and the 3-staged policy mechanism — with the benchmark loop (a later release) as a forward hook, not Step-0 closure.
4. A boundary contract naming what m-4 must expose to consumers — m-3 (the routing record as a possible evidenced record + the benchmark consuming observed evidence), m-5 (archetype tags ↔ routing), m-6 (routing categories feeding gate→email buckets) — so DESIGN locks interfaces consumer-validated.

c2 GUARDRAILS (VP-set, `c2-decomp` 20260629-183247):
- Phase band = AUDIT + DESIGN only. No build / PLAN / IMPL.
- Focus on the LOCKED c1 contract; do **not** reopen m-1/m-2 or re-litigate R2 (preserve it). Audit your own domain's prior art + mechanism.
- m-3↔m-4 SEAM: your benchmark loop (a later release) consumes m-3 observed evidence; a routing record may be an evidenced record. Flag seam findings; a dedicated `c2-*-coord` COORD sub-thread owns the seam at DESIGN (reconciled before any c2 lock).
- m-5 SEAM (explicit design risk): m-5 owns archetype tag-space + authority-ceiling-at-spawn. Name where archetype tags would parameterize routing/policy. c2 cannot lock m-4 without a lock-time m-5 seam disposition — surface, don't close, concrete archetype/tag semantics.
- m-6 is the warm consumer lens for c2 (engages at consumer-review, after draft designs).

Boundary contract:
- Writes: the router + routing-record + policy primitive that m-3/m-5/m-6 build on.
- Reads: the export adaptive-routing pillar; the stock protocol routing gap; jcode MultiProvider + claude-code/agent-scripts model-selection; the locked c1 R2/§5/§J seam.
- Target entity: the frank routing/policy primitive (a design recommendation, not code).
- Downstream consumer: m-3 (evidenced routing record + benchmark evidence seam), m-5 (archetype routing), m-6 (routing categories → buckets).
- Proof: E1 source citations (file:line) backing every gap and claim.
- No-consumer action: flag any proposed primitive with no downstream consumer.

AUDIT-FIRST GATES — may reject or narrow:
0. Duplicate/already-built check: does jcode's MultiProvider (or another reference) already provide a router we should promote/wire rather than rebuild? If so, recommend promote/wire and identify only the governance-record gap on top.
1. If the routing-record primitive cannot ride an existing runtime in Step 1 (the record can be written, but execution needs our own runtime), say so and route the execution dependency to the orchestrator (it maps to Step 3 per `ROADMAP.md`).

Out of scope:
- m-3 observe/evidence internals (sibling audit `c2-audit-m-3`), the m-1/m-2 foundation (locked), the TUI/email-client UX, and any code. Router *execution* is Step 3 — design the record + policy now, not the runtime.

Deliverable:
A file-relay AUDIT report from each seat independently, then a reconciled view: the 4-bucket verdict; the implicit-routing gap statement; the router + routing-record + 3-staged-policy design recommendation; the consumer boundary contract; the m-3↔m-4 + m-5 seam notes; evidence levels (E1 cites); operator-judgment items or none. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT (read-only phase; cwd is not a git repo, so the structured unavailable form is expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
