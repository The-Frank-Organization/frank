## Team m-2 — Forms & Determinism: PROCEED TO DESIGN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c1-design-m-2
PARENT_DISPATCH_ID: c1-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — operator owns schema-carrier / storage / legacy-strictness defaults (grill them); design surfaces operator-judgment items
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-2.implementer, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
BUNDLE_ID: m-2-forms-determinism
OWNER: m-2 (Forms & Determinism)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags product-semantics decisions. Not in scope: source/test edits, branches, commits, PRs, scaffolding, prototype code. This is the frank research+design phase; design-lock is the terminal — no implementation.

Basis: your reconciled c1-audit-m-2 (see master/RECONCILE.md) is APPROVED by the VP (c1-reconcile verdict: approve). Design the schema the audit recommended — but per the VP's load-bearing guardrail, treat the converged audit points as HYPOTHESES to PROVE in design, not as proven facts.

Co-foundational with m-1 (hard boundary): m-2's schema DECLARES the system-owned fields that m-1's store/stamper FILLS (FROM, ROLE, PARENT, timestamps, evidence); system-filled PARENT is exactly where the lineage engine strengthens to forgery-robust. Neither domain design-locks in isolation — coordinate the shared system-field contract with m-1.planner; the orchestrator runs the joint lock after consumer review.

Design questions to resolve (grill the operator on the operator-owned ones):
1. OPERATOR — schema carrier: JSON-Schema core + x-owner / x-seat-scope / x-consumer extensions (mature tooling) vs a bespoke FieldSpec data model (claude-code uses Zod; jcode uses typed Rust records). Pick the single canonical source the tool, courier, and linter all read.
2. OPERATOR — storage backing + markdown role: typed canonical storage as JSON, SQLite row/object, or both; markdown as a rendered view only vs a signed export; and how strict frank is about rejecting legacy hand-authored markdown relays.
3. OPERATOR — schema versioning / evolution: how the courier versions the form, and the compatibility guarantee when a parked lane resumes on an older envelope (cross-cutting m-1 store + m-6 scheduler).
4. required-when predicate limits: confirm it stays a BOUNDED declarative predicate over (PHASE, CEREMONY_TIER, field-values, grant/record kinds), never Turing-complete (CodeAct / determinism).
5. the sanctioned overflow channel: one bounded, never-gate-input free-text escape hatch so unanticipated messages are not blocked and the body does not silently re-grow prose.

Required design alternatives (lay out, with tradeoffs + a recommendation): schema carrier Option A (JSON-Schema + extensions) vs Option B (bespoke FieldSpec). Tradeoff axis: tooling maturity / mechanical-checkability vs expressing field-ownership + seat-scoping + consumer annotation natively.

Hard proof requirements (VP carry-forwards):
- Produce the dissolve/survive TABLE where EVERY dissolved upstream relay-lint check is classified as one of: prose-only (genuinely no semantics), typed-form-validation (survives at fill-time), or cross-relay-lineage (survives in the separate engine). No check may be dropped without a class.
- Specify the field-ownership model (system / seat-scoped-enum / agent-enum-pick / free-text) + the fill-time-authority render semantics (forbidden option absent from the seat's form) + the one canonical schema source.
- Hold the Format-Tax line: reasoning-bearing free-text stays genuinely free (canonical-iff-consumed strictly) so the form does not re-grow the prose problem it dissolved.

Boundary contract — the schema MUST express these consumer fields; name them BEFORE lock (consumer review by m-3/m-4/m-6 is a lock prerequisite):
- m-1: the system-owned field set (FROM, ROLE, PARENT, timestamps) marked owner:system / not lane-fillable.
- m-3: observe/evidence fields (ACTIONS_GIT_REF, FINAL_GIT_STATUS_SHORT, EVIDENCE_TARGET, achieved-evidence, per-phase done-predicate hook) with the target>achieved auto-flag.
- m-4: routing-record fields (role+model-per-dispatch, capability-priors, justified-deviation, reusing the DESIGN_RECORD_KIND enum shape).
- m-6: human-gate/email fields (HUMAN_GATE_REQUIRED as a system monotonic floor — agent may only RAISE; gate→email bucket; the Owner Decision Brief 7-field sub-schema).

Out of scope: m-1 store/identity internals, the observe / routing / human-surface mechanisms, any code.

Relay hygiene (VP note): keep the pair-thread DISPATCH_ID c1-design-m-2; address the design-review request TO m-2.implementer (Template I), not TO the orchestrator.

Deliverable: a design doc (Superpowers brainstorming + design-grill), recorded as DESIGN_DOC_ID c1-design-m-2-forms-determinism, containing — the dissolve/survive table; the field-ownership + fill-time-authority + single-source model; the chosen carrier with rejected alternative; the bounded required-when predicate vocabulary; the sanctioned overflow channel; operator decisions/defaults from the grill folded into a GRILL_LOCK; the named consumer boundary contract; open questions. Then send the design-review request TO m-2.implementer (Template I), and report design-complete to the orchestrator for the joint co-foundational lock (after m-3/m-4/m-6 consumer review). Design Q&A inline; the design lock is file-first.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
