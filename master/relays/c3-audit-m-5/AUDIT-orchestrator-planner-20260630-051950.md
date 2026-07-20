## Team m-5 — Workflows & Archetypes (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c3-audit-m-5
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-5.planner, m-5.implementer
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)

Phase scope — AUDIT (read-only), opening **Cycle c3** (the **final Step-0 design cycle** — completes the six-domain design-of-record). Inspect source and docs, run safe read-only commands, produce an independent paired audit with findings. Not in scope: any edits, branches, commits, PRs, scaffolding, or prototype code. Still the v3 research + design phase — no implementation exists or is authorized.

**Focused audit (VP-set, `c3-decomp` 20260630-051448 F4).** You already ran the c2 narrow consumer-lens (the m-3↔m-4 seam + your three reserved proposals). **Do not redo that as if fresh.** This audit is focused on two things: (a) the workflow/archetype **prior-art sweep** you have *not* done as your own domain, and (b) consolidating your **c3-reserved decisions** into auditable form — what's settled vs still-open from the c2 proposals, ready to bind at c3 design-lock.

Pair roles & research method:
- m-5.planner (Claude Opus 4.8, high thinking): lead the audit; surface design questions. Use parallel agents + websearch + a deep-research workflow for the prior-art sweep (workflow/archetype presets, agent-topology graphs, behavioral-mode templates, side-question/sensor patterns).
- m-5.implementer (GPT-5.5): run an INDEPENDENT audit — do not mirror the planner — and challenge/answer the design questions. Use subagents + websearches.
- Independent paired audit: each member audits separately, then reconciles. The Planner does not spawn, direct, or simulate the Implementer; the two seats are independent operator-relayed sessions.

Domain context. m-5 owns **workflows & archetypes**: expansion-slot presets (topology + gate-set + **human-mode**), the **tag-space**, per-archetype **observe invariants**, **authority-ceiling-at-spawn**, and **sensor/actuator** archetypes. The archetype system is where the c2-reserved opaque atoms become concrete, consumable semantics.

You build on the LOCKED c1+c2 contract (do **not** reopen it — `master/ARCHITECTURE.md` §1–§C2):
- m-2 declares `slot_in` (work-archetype) + `seat_archetype` as **opaque atoms reserved to you** — c3 binds their concrete vocabulary.
- m-3 owns the observe-gate mechanism + classifies `slot_in` **at work-record acceptance** (non-lane-writable, F1). Your per-archetype observe invariants parameterize *which* done-predicate/gate-set an archetype runs.
- m-4 owns routing + the `seat_archetype` capability-prior key + the **GL-4 routing-template record mechanism** (pre-filled `routing_decision` + `template_ref`, no-bypass). Your template *structures/lineup* feed it.
- M4-1: routing B→A escalation rides the c1 monotonic HUMAN_GATE routing-raise (no new gate class).

Sources to audit (cross-check the export's distillation against real source):
- v2.8.8 current state: `extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/` — expansion-slot presets, workflow/sprint archetypes, any tag/topology + per-archetype gate handling in protocol.md + the role skills. Our LIVE team (`master/`) is itself a running orchestrator-archetype instance.
- `references/codex/codex-rs/` — **`collaboration-mode-templates`** (markdown single-agent behavioral presets: PLAN/EXECUTE/PAIR_PROGRAMMING — NOT multi-agent topologies) and **`agent-graph-store`** (MUTABLE spawn-topology graph: upsert, Open/Closed, single-parent — the *opposite* mutability stance to our append-only lineage). Cross-ref `references/codex-notes.md` §3.
- `references/claude-code/` — subagent/agent-type presets + `src/utils/sideQuestion.ts` + `src/commands/btw` (the read-only parallel forked side-question — your sensor archetype's direct prior art). Cross-ref `references/jcode-ux-notes.md`.
- Export design intent: `extracted/agentic-dev-team-skills-v3-export/v3-design/` (the adaptive-routing pillar + any archetype/workflow/expansion-slot design intent).
- Your own c2 reserved proposals: `master/relays/c2-consumer-review-m-5/` + the locked m-3/m-4 design docs (the `slot_in` / `seat_archetype` / GL-4 / sensor seam you surfaced).

Design question to resolve:
What is the minimal **archetype system** — the concrete tag-space (`slot_in` work-archetypes + `seat_archetypes`), per-archetype **invariant composition** (observe invariants + default gate-set + authority-ceiling-at-spawn + routing prior), and the **template lineup** (T1 Solo / T2 Adversarial Pair / T3 Sensor) — that binds the c2-reserved opaque atoms into concrete m-3/m-4-consumable semantics, rides existing runtimes in Step 1, and extends to the standalone runtime (Steps 4–5) without re-cutting?

Hard acceptance criteria:
1. A 4-bucket verdict (still-open / already-closed / product-overlapped / recommended-next) on the v3 m-5 archetype system vs what v2.8.8, codex (`collaboration-mode-templates`, `agent-graph-store`), and claude-code (subagents, `sideQuestion`) already provide.
2. The **concrete tag-space** proposal — `slot_in` work-archetype values + `seat_archetype` values — *surfaced* as the c3 design input (locked in DESIGN, not audit; F4 / "not authorized: no concrete archetype value lock outside m-5-owned c3 design").
3. **Per-archetype invariant composition**: observe invariants (m-3), default gate-set (m-2/m-3), authority-ceiling-at-spawn semantics, routing prior (m-4 `seat_archetype`).
4. The **template lineup** (T1/T2/T3) structures from the c2 GL-4 proposal — topology / seats / panes / gate-set / read-only-ness for each shipped template.
5. The **sensor/actuator** archetype design (the c2 sensor sketch → full): read-only ceiling, tool-blocked, single-turn; the interjection side-question home; the content=self_reported vs metadata=observed integrity split.
6. A boundary contract: what m-5 exposes to **m-6** — the per-archetype **human-mode vocabulary** + the **interjection surface contract** (Seam A/B).

c3 GUARDRAILS (VP-set, `c3-decomp` 20260630-051448):
- Phase band = AUDIT + DESIGN only. No build / PLAN / IMPL.
- Focus on the LOCKED c1+c2 contract; do **not** reopen m-1..m-4. Audit your own domain prior art + bind your reserved decisions.
- **Pair-artifact requirement (F4):** BOTH m-5.planner and m-5.implementer return an audit artifact, OR one explicitly reconciled pair artifact. No single-seat audit — c2 showed missing pair-reconcile relays force avoidable orchestrator inference.
- **m-5↔m-6 SEAM (F2 — declare-before-bind):** you OWN the human-mode vocabulary + archetype/sensor semantics; m-6 BINDS surface behavior to them. The DESIGN-phase COORD thread must surface your vocabulary **explicitly before** m-6 binds. In audit, name the vocabulary surface you will declare; do not wait on m-6.
- No concrete archetype value **LOCK** in audit — surface for the c3 design-lock (your owned decision).
- `GRILL_REQUIRED: yes` will be set at DESIGN (archetype semantics + authority-ceiling are cross-domain + hard-to-reverse).

Boundary contract:
- Writes: the archetype system (tag-space + invariant composition + template lineup + sensor/actuator) that m-3/m-4 consume and m-6 binds.
- Reads: v2.8.8 expansion-slots/workflows; codex templates/graph; claude-code subagents/sideQuestion; the v3 export; your c2 reserved proposals + the locked m-3/m-4 docs.
- Target entity: the v3 archetype system (a design recommendation, not code).
- Downstream consumer: m-6 (human-mode vocabulary + interjection surface). m-3/m-4 are locked suppliers you parameterize, not reopen.
- Proof: E1 source citations (file:line) backing every gap and claim.
- No-consumer action: flag any proposed archetype/template with no downstream consumer or no m-3/m-4 mechanism behind it.

AUDIT-FIRST GATES — may reject or narrow:
0. Duplicate/already-built check: does codex `collaboration-mode-templates` / `agent-graph-store` or v2.8.8 expansion-slots already provide a preset/topology mechanism to **promote** rather than rebuild? If so, recommend promote/wire (noting the append-only-vs-mutable lineage stance), do not rebuild.
1. If an archetype invariant cannot ride an existing runtime in Step 1 (needs the standalone runtime to enforce a ceiling/topology), say so and route the dependency to the orchestrator (it is a later-step build concern, not a c3 design blocker).

Out of scope:
- m-6 human-surface internals (sibling audit `c3-audit-m-6`), the locked m-1..m-4 foundation/runtime-intelligence, the TUI/email-client UX, and any code.

Deliverable:
A file-relay AUDIT report from each seat independently, then a reconciled view: the 4-bucket verdict; the concrete tag-space proposal (surfaced); the per-archetype invariant composition; the T1/T2/T3 template structures; the sensor/actuator design; the m-6 boundary contract (human-mode vocabulary + interjection surface); the m-5↔m-6 seam notes; evidence levels (E1 cites); operator-judgment items or none. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT (read-only phase; cwd is not a git repo, so the structured unavailable form is expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
