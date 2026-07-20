## Team m-6 — Human Surface & Scheduler (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c3-audit-m-6
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-6.planner, m-6.implementer
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler)

Phase scope — AUDIT (read-only), opening **Cycle c3** (the **final Step-0 design cycle** — completes the six-domain design-of-record). Inspect source and docs, run safe read-only commands, produce an independent paired audit with findings. Not in scope: any edits, branches, commits, PRs, scaffolding, or prototype code. Still the v3 research + design phase — no implementation exists or is authorized.

**Full domain audit (VP-set, `c3-decomp` 20260630-051448 F4).** m-6 has **no design-of-record** — you ran the c1 + c2 consumer-lenses against the locked *foundation/runtime* contracts but never designed your own domain. This is your **full domain audit**: human-gate surfaces, the Owner Decision Brief, email/meeting collaboration, scheduler park/wake, and the away-mode bridge.

Pair roles & research method:
- m-6.planner (Claude Opus 4.8, high thinking): lead the audit; surface design questions. Use parallel agents + websearch + a deep-research workflow for the prior-art sweep (human-in-the-loop approval surfaces, decision-brief/digest patterns, email/notification-governance TUIs, park/wake schedulers + away-mode bridges).
- m-6.implementer (GPT-5.5): run an INDEPENDENT audit — do not mirror the planner — and challenge/answer the design questions. Use subagents + websearches.
- Independent paired audit: each member audits separately, then reconciles. The Planner does not spawn, direct, or simulate the Implementer; the two seats are independent operator-relayed sessions.

Domain context. m-6 owns the **human surface & scheduler**: email-governance + meeting-collaboration **surfaces**, **gate→email buckets**, the **Owner Decision Brief** (ODB), and **scheduler park/wake** — the relay/inbox governance graph as first-class operator-facing comms.

You build on the LOCKED c1+c2 contract (do **not** reopen it — `master/ARCHITECTURE.md` §1–§C2 + §J):
- m-1 addressing graph (who a gate routes to) + the §J operator-judgment defaults (J1 `on_timeout = hold_and_resummon`, never auto-approve; J2 `gate_category` default set, **operator-configurable** — a forward requirement that lands here).
- m-2 HUMAN_GATE fields + `gate_category` (A/B set) + the **monotonic HUMAN_GATE floor**.
- m-3 egress/content-safety gate (the dormant fail-closed chokepoint, sole external sender) — your away-mode bridge's *first external send* trips it (§J1 away-mode).
- m-4 routing — for the interjection **side-question** archetype you host.

Sources to audit (cross-check the export's distillation against real source):
- `references/agent-scripts/` (Steinberger): the **Owner Decision Brief** + the egress/away surfaces — direct prior art for your ODB + decision-capture path.
- v2.8.8 current state: `extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/` — email-governance + meeting-collaboration surfaces, gate handling, any human-approval/digest surface in protocol.md + the role skills. Our LIVE relay store (`master/relays/`) is a running governance-graph instance the operator hand-relays today (the pain you remove).
- `references/jcode/` + `references/claude-code/` — human-surface / permission / notification UX, side-panel + approval flows. **Apply the negative-look** `references/jcode-ux-notes.md` (GUI noise/filler; the bad `/btw` side-panel) and the **positive-look** `references/codex-notes.md` (the look/feel brief).
- Export design intent: `extracted/agentic-dev-team-skills-v3-export/v3-design/` + `EXTERNAL-REFERENCES.md` (any human-surface / scheduler / email-governance design intent).
- The §J forward requirements: `master/ARCHITECTURE.md` §J — customizable `gate_category` membership / A·B map, the opt-in away-mode external-inbox bridge (egress-gated).

Design question to resolve:
What is the minimal **human-surface + scheduler** design — the gate→email **bucket taxonomy**, the **Owner Decision Brief**, the **park/wake** scheduler, and the egress-gated **away-mode bridge** — that turns m-2 HUMAN_GATE / §J gates into an operator-facing governance surface, rides existing runtimes in Step 1 (local-only mechanism, Step 2), and extends to the full email-client UX (Step 4) without re-cutting?

Hard acceptance criteria:
1. A 4-bucket verdict (still-open / already-closed / product-overlapped / recommended-next) on the v3 m-6 human-surface + scheduler vs what v2.8.8, agent-scripts (ODB / egress), and jcode/claude-code already provide.
2. The **gate→email bucket taxonomy**: m-2 `gate_category` (A/B, §J) → buckets → per-bucket surface behavior + the operator **decision-capture/return path**.
3. The **Owner Decision Brief** design (port agent-scripts' ODB): what a gate presents to the human + the decision-capture path (the J1 `hold_and_resummon` surface — never auto-approve).
4. The **scheduler park/wake** design: park a gated lane, wake-on-reply; the **away-mode external-inbox bridge** (egress-gated via m-3; §J1 away-mode = the egress chokepoint's first external send).
5. The **email-governance + meeting-collaboration** surface design-of-record (mechanism = Step 2, UX = Step 4 — design the contracts now). Explicitly apply the jcode negative-look (no GUI noise/filler) + the codex positive-look.
6. A boundary contract: what m-6 **consumes** from m-5 — the per-archetype **human-mode vocabulary** (declare-before-bind, F2) — and **hosts** — the **interjection surface** (steer/side-question/interrupt; m-5 sensor archetype + m-4 routing). The Seam A/B deliverable.

c3 GUARDRAILS (VP-set, `c3-decomp` 20260630-051448):
- Phase band = AUDIT + DESIGN only. No build / PLAN / IMPL.
- Focus on the LOCKED c1+c2 contract; do **not** reopen m-1..m-4. Audit your own domain prior art + design your surface.
- **Pair-artifact requirement (F4):** BOTH m-6.planner and m-6.implementer return an audit artifact, OR one explicitly reconciled pair artifact. No single-seat audit — c2 showed missing pair-reconcile relays force avoidable orchestrator inference.
- **m-5↔m-6 SEAM (F2 — declare-before-bind):** m-5 OWNS the human-mode vocabulary; you BIND surface behavior to it. **Do not pre-bind** surface behavior before m-5 declares the vocabulary — in audit, name *where* you will consume the vocabulary + host the interjection surface, and flag the binding points for the DESIGN-phase COORD thread.
- No surface **LOCK** in audit — surface for the c3 design-lock.
- `GRILL_REQUIRED: yes` will be set at DESIGN (human-gate/egress surfaces are cross-domain + hard-to-reverse).

Boundary contract:
- Writes: the human-surface + scheduler design (bucket taxonomy + ODB + park/wake + away-mode bridge + interjection surface) — the operator-facing governance graph.
- Reads: agent-scripts ODB/egress; v2.8.8 email-governance/meeting surfaces; jcode/claude-code human-surface UX (+ the two reference notes); the v3 export; the locked m-1/m-2 §J fields; the locked m-3 egress + m-4 routing.
- Target entity: the v3 human-surface + scheduler design (a design recommendation, not code).
- Downstream consumer: none below you (terminal design domain); the runtime/product seats (m-7..m-12) are a later step. You **consume** m-5 (human-mode vocabulary) — flag any binding point with no m-5 vocabulary behind it.
- Proof: E1 source citations (file:line) backing every gap and claim.
- No-consumer action: flag any proposed surface/bucket with no gate behind it (m-2/§J) or no mechanism (m-1 addressing, m-3 egress, m-4 routing).

AUDIT-FIRST GATES — may reject or narrow:
0. Duplicate/already-built check: does agent-scripts' ODB or a v2.8.8 governance surface already provide a decision-brief/bucket surface to **promote** rather than rebuild? If so, recommend promote/wire, do not rebuild.
1. Local-first / egress fail-closed (`ROADMAP.md` cross-cutting): the Step-1/2 mechanism is **local-only** (UX over the *local* relay store; no external send); the away-mode bridge's first external send is **egress-gated** (m-3). If any surface needs external send before the egress gate exists, say so and route the dependency.

Out of scope:
- m-5 archetype internals (sibling audit `c3-audit-m-5`), the locked m-1..m-4, the full email-client UX *build* (Step 4 — design the contract now, don't build), and any code.

Deliverable:
A file-relay AUDIT report from each seat independently, then a reconciled view: the 4-bucket verdict; the gate→email bucket taxonomy; the Owner Decision Brief design; the park/wake + away-mode scheduler design; the email-governance/meeting surface design-of-record (with the jcode/codex looks applied); the m-5 boundary contract (human-mode vocabulary consume-point + interjection surface); the m-5↔m-6 seam notes; evidence levels (E1 cites); operator-judgment items or none. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT (read-only phase; cwd is not a git repo, so the structured unavailable form is expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
