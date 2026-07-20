## Team m-3 — Observation & Evidence: PROCEED TO DESIGN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-design-m-3
PARENT_DISPATCH_ID: c2-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — design surfaces operator-judgment items (executable-claim surface, egress policy, Step-1 read-vantage floor); grill them
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-4.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
BUNDLE_ID: m-3-observation-evidence
OWNER: m-3 (Observation & Evidence)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags product-semantics decisions. Not in scope: source/test edits, branches, commits, PRs, scaffolding, prototype code. Design-lock is the terminal — no implementation / PLAN.

Basis: your reconciled `c2-audit-m-3` is APPROVED by the VP (`c2-reconcile` verdict: approve). Design the observe/evidence + egress primitive your audit recommended — but per the standing guardrail, treat the converged audit points (the `observe_gate()`/`observe_result{}` strawman, the placement, the ladder split) as **HYPOTHESES to PROVE in design, not as proven facts.**

Co-design with m-4 (the seam): coordinate the m-3↔m-4 evidenced-routing-record seam in the **shared COORD thread `c2-design-m3-m4-coord`** (seeded — read it first; cite its current state in your design). **VP load-bearing item:** `routing_decision.deviated` is not a freestanding truth bit — your design owns **how m-3 observes/classifies the declared-vs-snapshot derivation** (mismatch = veto / flag / labeled evidence?) and what `evidence_integrity` the routing-record observation carries. Resolve in the COORD thread before reporting design-complete.

Design questions to resolve (grill the operator on the operator-owned ones):
1. OPERATOR — **executable-claim execution surface**: registry-approved check descriptors vs arbitrary project commands. Recommend v3.0 ships observe-as-send on passive conductor-derived E1/E2, the conductor-EXECUTES mechanism a reserved seam operator-gated at activation. Grill: confirm v3.0 ships without arbitrary execution; descriptors are the eventual form.
2. OPERATOR — **egress fail-closed policy**: auto-redact-low-risk vs always-block-on-first-release; rule-set membership operator-configurable (mirrors ratified §J2). Activation rides the first external send (m-6 away-mode, §J1). Grill: confirm operator owns the rule set + the fail-closed default.
3. OPERATOR — **Step-1 read-vantage floor**: observe-as-send rides Step 1 wherever the conductor holds an outside-the-lane read handle; a fully-opaque remote lane degrades honestly to labeled `self_reported` (locked DI-5 fallback); arbitrary-lane outside-vantage is the standalone-runtime carry-forward. Grill: confirm the honest-fallback labeling is an acceptable Step-1 floor.
4. The record-level `evidence_integrity: mixed` rollup — decide at DESIGN whether to carry a record-level rollup over per-field integrity. **The per-field tag stays two-value {observed | self_reported} (locked R3 — do NOT reopen).**

Hard proof requirements (prove, don't assert):
- PROVE the **observe-AS-send** gate closes the self-reported-done gap by construction: the outside-the-lane observation binds atomically to the gated relay (closes the check→edit→send-stale TOCTOU); placement inside m-1 `submit()` AFTER m-2 form-validation, BEFORE append/accept; **observer-only against the locked R3 write-allowlist** (writes only the closed m-3 observed/computed set + a pass/fail veto; never authors identity/envelope fields).
- Specify the `observe_gate()` interface + `observe_result{}` shape; the evidence ladder (split `EVIDENCE_TARGET` intent vs `achieved_evidence` fact; E1/E2 passive Step-1 probes; E4 = the explicit live-verify gate, often operator-observed); executable-claims as the reserved E2+ slot; the egress gate as the fail-closed conductor chokepoint (v3.0-dormant under canonical-iff-consumed until the first external send).
- Done-predicates: phase-shaped base + archetype-tag add-ons via the locked `slot_in` reserved atom. **SURFACE the candidate archetype→invariant mappings (incl. the verifier-tamper-resistant ones); do NOT close them** — see the m-5 lock prerequisite.

Lock prerequisites (VP — both required before any c2 m-3 lock):
- **m-5 seam disposition (lock prerequisite, NOT optional commentary):** the lock must NOT finalize m-3 predicate semantics without either a narrow m-5 consumer review on the draft design OR an explicit reconcile reservation preserving m-5 ownership of the concrete tag-space + archetype invariants. Design the predicate **mechanism** against an opaque archetype-tag input; leave the tag semantics to m-5.
- **F5 acceptance criterion (no overclaim):** state novelty precisely. v2.8.8 (E-ladder vocabulary), jcode (permission/review queue), claude-code (Pre/PostTool/Stop hooks), agent-scripts (ODB / egress rules) already provide PARTIAL hook/evidence/egress primitives — the design must PROMOTE those and claim novelty ONLY for the integration (outside-the-lane observe-AS-send + `evidence_integrity` + the unbypassable chokepoint), not for hooks/evidence/egress in general.

Boundary contract — name the consumer fields before lock:
- → m-4: observed-evidence records weighted on `evidence_integrity` (`self_reported` ≠ clean benchmark signal); the routing record as a possible evidenced record; the declared-vs-snapshot observation (the COORD seam).
- → m-5: the done-predicate hook's `slot_in` archetype-tag input — m-3 owns the predicate-execution mechanism + result shape, m-5 owns the tag-space (lock prerequisite).
- → m-6: the observe-gate veto → gate→email bucket; the egress gate for the away-mode external bridge; the promoted Owner Decision Brief as the evidence-summary content schema (warm lens at consumer-review).

Out of scope: m-4 router/policy internals (sibling `c2-design-m-4`); the locked m-1/m-2 foundation; the TUI/email-client UX; any code.

Relay hygiene: keep the pair-thread DISPATCH_ID `c2-design-m-3`; address the design-review request TO m-3.implementer (Template I), not TO the orchestrator.

Deliverable: a design doc (Superpowers brainstorming + design-grill), recorded as DESIGN_DOC_ID `c2-design-m-3-observation-evidence` under `master/domains/m-3-observation-evidence/design/`, containing — the proven observe-AS-send gate + placement + R3-allowlist observer; the `observe_gate()`/`observe_result{}` interface; the evidence ladder + executable-claim reserved slot; the egress chokepoint; the surfaced (not closed) archetype done-predicate mappings; the COORD-thread seam resolution cited; operator decisions/defaults folded into a GRILL_LOCK; the named consumer boundary contract; the precise novelty statement; open questions. Then send the design-review request TO m-3.implementer (Template I), and report design-complete to the orchestrator for the c2 lock (after the COORD reconcile + the m-5 disposition). Do not self-advance to PLAN.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
