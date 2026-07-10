## Team m-5 — Workflows & Archetypes: PROCEED TO DESIGN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c3-design-m-5
PARENT_DISPATCH_ID: c3-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — design surfaces operator-judgment items (the grill agenda below); grill them
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-5.planner
CC: m-5.implementer, m-6.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c3-design-m-5-workflows-archetypes
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags product-semantics decisions. Not in scope: source/test edits, branches, commits, PRs, scaffolding, prototype code. Design-lock is the terminal — no implementation / PLAN.

Basis: your reconciled `c3-audit-m-5` is APPROVED by the VP (`c3-reconcile` verdict: approve). Design the **archetype system** your audit surfaced — but per the standing guardrail, treat the surfaced candidates (the tag-space, the composition rule, T1/T2/T3, sensor/actuator) as **HYPOTHESES to PROVE and LOCK in design, not as proven facts.** This is the Step-1 archetype-system lock that binds the c2-reserved opaque atoms (`slot_in`/`seat_archetype`).

Co-design with m-6 (the seam): coordinate in the shared COORD thread **`c3-design-m5-m6-coord`** (seeded — read it first; cite its current state in your design). **The load-bearing seam item: you DECLARE the human-mode two-layer vocabulary (posture × surface_intent) BEFORE m-6 binds surface behavior to it (declare-before-bind, VP F2).** Resolve the seam in the COORD before reporting design-complete.

Design questions / grill agenda (your reconcile's five items + the sensor routing-prior — grill and LOCK them, m-5-owned):
1. **Actuator** — lock `actuator_class` as a **derived** mutating-ceiling class (NOT a literal `seat_archetype` value) for Step-1; record a future `single_bounded_action` literal seat as a Step-4/5 runtime-enforcement carry. Grill: confirm derived-for-Step-1.
2. **Read-only work-archetype ship-set** — which of `research_synthesis` / `qa_review` / `docs_chore` LOCK as Step-1 `slot_in` values vs mark Step-5. Grill: pick the minimal set m-3/m-6 need for the Step-0 architecture-of-record.
3. **Human-mode vocabulary** — the two-layer (posture × `surface_intent`) value-sets; resolve the `operator_gate`/`hold_and_resummon` placement (m-5 value vs locked-mechanism reference), `away_bridge_eligible`-as-flag, and one-field-vs-two. **Declare in the COORD** (declare-before-bind).
4. **Ceiling vector** — the exact authority-ceiling axes (candidate: read · tool · write · dispatch/route · external-send · merge · human-verdict) + per-axis monotonic tightening rules. Grill the minimal Step-1 dimension set (partial order, not a total ladder).
5. **Naming** — lock the canonical `lower_snake_case` roster for `slot_in` + `seat_archetype`.
+ **Sensor routing-prior** (locked `fast-cheap` vs inherit-parent-for-cache) — minor, m-4-boundary; grill-confirm.

Hard proof requirements (prove, don't assert):
- PROVE the **unifying binding** — archetype → {topology + gate-set + authority-ceiling-at-spawn + observe-invariants + routing-prior} as ONE declared, seat-stamped, append-only, auditable unit. This integration is the net-new contribution; PROMOTE the four mechanisms (the upstream panel preset+selection + AUTHORITY/tier enum; codex presets-as-data + graph-query-as-append-only-projection; claude-code agent-type shape + sideQuestion sensor-invariant), do not rebuild them.
- The **two-axis composition**: `seat_archetype` (spawn-fixed) ⊗ `slot_in` (per-work-record, conductor-classified-at-acceptance — locked F1); the work-invariant families (the tamper-resistant `refactor`-no-test-edits + `bugfix`-red→green depend on F1 — prove they cannot be escaped by lane re-tagging).
- The **T1/T2/T3** template structures (schema + seats + topology + gate-set + read-only-ness + model-slots); conductor/N-pair template **DEFERRED to Step-5** (no Step-1 consumer).
- The **sensor (full)** + `actuator_class`; the **read-only→write hard-gated boundary** (sensors emit into a separately-spawned actuator, no in-place upgrade).
- **Step-1-rideable** as a recorded governance contract; ceiling enforcement **tiered** (best-effort host config Step-1 — real on claude-code's tool-allowlist; conductor-uniform at the standalone runtime Steps 4-5; **no re-cut** — F2 records the ceiling per-assignment). Route the standalone-enforcement dependency to the orchestrator, not a c3 blocker.

Guardrails (VP):
- **No value lock outside this DESIGN grill** — the lock is here, m-5-owned, under `GRILL_REQUIRED: yes`.
- **Do NOT reopen locked m-1..m-4** — design AGAINST the locked m-2 schema, m-3 observe mechanism, m-4 routing + GL-4 template record + `seat_archetype` key, M4-1. They are suppliers you parameterize.
- **No m-2 `required_when`/`visible_when` over concrete tag values** (C2.4 reserved opaque atoms only — no m-2 micro-fold).
- **Declare-before-bind**: declare the human-mode vocabulary in the COORD before m-6 binds.

Boundary contract — name the consumer fields before lock:
- → m-6: the human-mode two-layer vocabulary (posture × `surface_intent`) + the interjection surface contract (the sensor archetype) — resolved in the COORD.
- m-3/m-4 are locked suppliers you parameterize (`slot_in` → m-3 observe-invariant selection; `seat_archetype` → m-4 ceiling + routing-prior; template → m-4 record + conductor pane-spawn) — not reopen.

Out of scope: m-6 surface internals (sibling `c3-design-m-6`); the locked m-1..m-4 foundation/runtime-intelligence; the TUI/email-client UX; any code.

Relay hygiene: keep the pair-thread DISPATCH_ID `c3-design-m-5`; address the design-review request TO m-5.implementer (not the orchestrator).

Deliverable: a design doc (Superpowers brainstorming + design-grill), recorded as DESIGN_DOC_ID `c3-design-m-5-workflows-archetypes` under `master/domains/m-5-workflows-archetypes/design/`, containing — the proven unifying binding; the LOCKED tag-space + composition + ceiling vector + T1/T2/T3 + sensor/`actuator_class`; the COORD-thread seam resolution (the declared human-mode vocabulary) cited; operator decisions/defaults folded into a GRILL_LOCK; the named consumer boundary contract; the precise novelty statement; open questions. Then send the design-review request TO m-5.implementer, and report design-complete to the orchestrator for the c3 lock (after the COORD reconcile). Do not self-advance to PLAN.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
