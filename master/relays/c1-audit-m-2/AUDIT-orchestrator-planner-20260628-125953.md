## Team m-2 — Forms & Determinism (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-audit-m-2
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-2.planner, m-2.implementer
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-2-forms-determinism
OWNER: m-2 (Forms & Determinism)

Phase scope — AUDIT (read-only). Inspect source and docs, run safe read-only commands, and produce an independent paired audit with findings. Not in scope: any edits, branches, commits, PRs, scaffolding, or prototype code. This is the v3 research + design phase — no implementation exists yet and none is authorized; implementation would begin only on a future bare own-line `DISPATCH IMPL` addressed to the Implementer, which is not part of this cycle.

Pair roles & research method:
- m-2.planner (Claude Opus 4.8, high thinking): lead the audit and surface design questions. Use multiple parallel agents, websearch, and a deep-research workflow for the prior-art sweep (typed/structured inter-agent message formats, schema-validation-as-gate, declarative form / field-ownership models).
- m-2.implementer (GPT-5.5): run an INDEPENDENT audit — do not mirror the planner — and challenge or answer design questions. You lack a built-in deep-research skill, so use subagents and websearches for your research.
- Independent paired audit: each member audits separately, then reconciles. The Planner does not spawn, direct, or simulate the Implementer; the two seats are independent operator-relayed sessions.

Domain context. m-2 owns the relay ENVELOPE: the declarative field schema (field, owner, type, required-when, enum-set), the field-ownership model (system-filled / seat-scoped-enum / agent-enum-pick / free-text), and the linter's refactor from prose-ambiguity defenses into form-validation (intra-relay) plus a lineage engine (cross-relay). The thesis: relays stop being hand-authored markdown and become a form the courier renders from typed input — "fill the form, don't write the prose" — dissolving most of relay-lint.py's prose machinery and enforcing authority at FILL (a forbidden option is simply absent from your seat's form) rather than post-hoc.

Sources to audit (cross-check the export's distillation against the real source):
- v2.8.8 current state: extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/ — tools/relay-lint.py (the ~1448-line linter, the prose-ambiguity defenses to be dissolved) and protocol.md (the ~50-field header schema across its field clusters), plus the three locked field-ownership decisions described in the export pillar note.
- references/agent-scripts: the JSON-only structured contracts and the committer reset-then-stage-explicit pattern (constrain-at-the-tool, no FROM param).
- references/jcode and references/claude-code: how each types its inter-agent messages (Rust structs / TS types) and where validation happens.
- Export design intent: extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md (the structured-form / determinism model, the "canonical iff a mechanical consumer reads it" rule, the three field-ownership locks, and the CodeAct "forms beat code for the control plane" reasoning).

Design question to resolve:
What is the minimal declarative field-schema plus field-ownership model that the tool, courier, and linter all read from one source — dynamic by phase and ceremony-tier — such that authority is enforced at FILL rather than post-hoc, and the schema can express its consumers' fields (observe, routing, human-gate) from day one?

Hard acceptance criteria:
1. A 4-bucket verdict on the v3 form schema + field-ownership model vs what v2.8.8's linter and protocol header already encode.
2. A map of which v2.8.8 prose-ambiguity defenses (strip_fenced/inline, ambiguous-continuation, detached-row, duplicate-block, bare unfenced DISPATCH-token machinery) become moot once the envelope is typed, and which lint logic survives as form-validation plus a lineage engine.
3. A design recommendation (the seed of the Step-1 interface sketch): the field-schema representation (field, owner, type, required-when, enum-set) and how the form is rendered and validated.
4. A boundary contract: the schema MUST express m-3 observe-as-send fields (observable done-predicates, evidence ladder, system-filled evidence), m-4 routing-record fields (role+model dispatch, static priors, justified-deviation), and m-6 human-gate/email fields (HUMAN_GATE monotonic floor, gate→email buckets) — so consumer review can validate it before lock.

Boundary contract:
- Writes: the declarative field-schema + form-validation contract that m-1 stores and m-3/m-4/m-6 consume.
- Reads: the v2.8.8 linter + protocol schema; agent-scripts JSON contracts; jcode/claude-code message typing; the export form model.
- Target entity: the v3 form schema + field-ownership model (a design recommendation, not code).
- Downstream consumer: m-1 (store persists the form), m-3/m-4/m-6 (their fields live in the schema).
- Proof: E1 source citations backing every claim.
- No-consumer action: flag any field or rule with no consumer.

AUDIT-FIRST GATES — may reject or narrow:
0. Duplicate/already-built check: the v2.8.8 protocol header and relay-lint ALREADY encode ~50 fields plus field grammar — audit how much is promote-and-formalize vs net-new, and do not rebuild what survives.
1. If a typed envelope cannot preserve a v2.8.8 gate (for example a lineage check), name the gate and route the dependency to the orchestrator.

Out of scope:
- m-1 store/identity internals (sibling audit c1-audit-m-1), the observe / routing / human-surface internals (their own domains), and any code.

Deliverable:
A file-relay AUDIT report from each seat independently, then a reconciled view: the 4-bucket verdict; the dissolve-vs-survive map of the v2.8.8 lint logic; the form-schema design recommendation; the consumer boundary contract; evidence levels (E1 source cites); operator-judgment items or none. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT (read-only phase; cwd is not a git repo, so the structured unavailable form is expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
