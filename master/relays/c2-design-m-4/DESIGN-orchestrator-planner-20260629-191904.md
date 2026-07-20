## Team m-4 — Routing & Policy: PROCEED TO DESIGN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-design-m-4
PARENT_DISPATCH_ID: c2-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — design surfaces operator-judgment items (capability-prior seed; v3.0 record scope); grill them
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-4.planner
CC: m-4.implementer, m-3.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c2-design-m-4-routing-policy
BUNDLE_ID: m-4-routing-policy
OWNER: m-4 (Routing & Policy)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags product-semantics decisions. Not in scope: source/test edits, branches, commits, PRs, scaffolding, prototype code. Design-lock is the terminal — no implementation / PLAN. Note the Step-1/Step-3 split: design the routing **record + policy** now (rides existing runtimes); router **execution** is Step-3 — route that dependency forward, do not design the runtime.

Basis: your reconciled `c2-audit-m-4` is APPROVED by the VP (`c2-reconcile` verdict: approve). Design the governance-record layer your audit recommended — but treat the converged audit points (the `route_dispatch()` strawman, the record FieldSpec, the 3-staged policy) as **HYPOTHESES to PROVE in design, not as proven facts.**

Co-design with m-3 (the seam): coordinate the m-3↔m-4 evidenced-routing-record seam in the **shared COORD thread `c2-design-m3-m4-coord`** (seeded — read it first; cite its current state in your design). **VP load-bearing item:** `routing_decision.deviated` is not a freestanding truth bit — your design owns **how `deviated` is derived against the `capability_prior_snapshot`** (the planner declares it; the conductor observes declared-vs-snapshot). Resolve the derivation in the COORD thread before reporting design-complete.

Design questions to resolve (grill the operator on the operator-owned ones):
1. OPERATOR — **capability-prior seed values**: ship a default seed (Fugu published priors + ours), operator confirms/customizes; config-sourced + operator-configurable (mirrors ratified §J). Grill: confirm the default-seed-then-configurable shape.
2. OPERATOR — **v3.0 record scope**: the record captures the routing decision for human-launched lanes now (Step-1 "automated operator-relay" — operator still picks the lane; the conductor records the decision + justification), execution deferred to Step-3. Grill: confirm this scope at design.

Hard proof requirements (prove, don't assert):
- PROVE the **routing record** closes the implicit-routing gap: a first-class, seat-stamped, lineage-gated `routing_decision` relay → **attributable** (rides m-1 `submit()`, forgery-robust stamped FROM), **auditable** (append-only, lineage-walkable), **overridable** (category-B that escalates to A only on `human_decision_required`). Model stays payload throughout.
- PROVE **R2 preserved BY CONSTRUCTION (the single sharpest point):** the `required_when` predicate atom is the **agent-declared `deviated` boolean**, NOT a comparison of the model value against the snapshot — so **no `model_*` predicate ever enters the schema gate**. The planner declares `deviated`; the conductor observes declared-vs-snapshot as m-3 evidence (the COORD seam). Show the construction explicitly.
- Specify the `route_dispatch()` API — **fail-closed**: no acceptable route emits `human_decision_required` / `routing_unavailable`, and MUST NOT silently fall back to a default model. The capability-prior table — **declared, versioned, and snapshotted into the record** (replay-complete; a live config ref is rejected — an immutable record cannot be reconstructed after the fact). The `routing_decision` FieldSpec (reuse the `DESIGN_RECORD_KIND` *shape*, not its values; `routing_assignments` seat_scoped to planner/orch-planner; `capability_prior_snapshot` system computed_result; `justified_deviation` required_when the declared `deviated` is true + machine-readable reason code; `outcome_feedback_ref` null-reserved for v3.1). The 3-staged policy (prior floor → justified deviation [ships v3.0] → v3.1 outcome-feedback forward hook consuming m-3 evidence).
- identity ≠ authority at the routing layer: the **authority-ceiling-at-spawn caps what the router may assign** — the router refuses to staff a seat the archetype ceiling forbids.

Lock prerequisites (VP — both required before any c2 m-4 lock):
- **m-5 seam disposition (lock prerequisite, NOT optional commentary):** the lock must NOT finalize m-4 archetype-prior semantics without either a narrow m-5 consumer review on the draft design OR an explicit reconcile reservation preserving m-5 ownership of the concrete tag-space + authority-ceiling semantics. Design the prior table to **key on an opaque archetype-tag**; leave the tag + ceiling semantics to m-5.
- **F5 acceptance criterion (no overclaim — material):** state novelty as the **seat-stamped, persisted, auditable routing/deviation artifact** — NOT interpretable routing in general (Routesplain 2511.09373 / Arch-Router 2506.16655 already do that), and NOT non-gradient adaptation in general (C2MAB-V / PILOT bandit routers already do that). The design must concede those priors and locate the contribution precisely (the port of the SR 11-7→SR 26-2 override-register discipline into per-dispatch LLM routing).

Boundary contract — name the consumer fields before lock:
- → m-3: the `routing_decision` record as a possible evidenced record; the benchmark/v3.1 loop consumes m-3 observed evidence via `outcome_feedback_ref`; m-4 declares which routing fields are observed, m-3 owns how (the COORD seam).
- → m-5: archetype tags parameterize the capability-prior lookup; the authority-ceiling caps routable authority (lock prerequisite).
- → m-6: `routing` stays a category-B (orchestrator-absorbed) gate; escalates to A only on `human_decision_required`; expose the routing `gate_category` + the ODB content (recommendation + enumerated model choices) for the rare A-escalation (warm lens at consumer-review).

Out of scope: m-3 observe/evidence internals (sibling `c2-design-m-3`); the locked m-1/m-2 foundation (do not re-litigate R2 — preserve it); router execution (Step-3); the TUI/email-client UX; any code.

Relay hygiene: keep the pair-thread DISPATCH_ID `c2-design-m-4`; address the design-review request TO m-4.implementer (Template I), not TO the orchestrator.

Deliverable: a design doc (Superpowers brainstorming + design-grill), recorded as DESIGN_DOC_ID `c2-design-m-4-routing-policy` under `master/domains/m-4-routing-policy/design/`, containing — the proven routing record with R2-by-construction; the `route_dispatch()` fail-closed interface; the declared/versioned/snapshotted capability-prior table; the routing-record FieldSpec; the 3-staged policy with the v3.1 forward hook; the authority-ceiling cap; the COORD-thread seam resolution cited; operator decisions/defaults folded into a GRILL_LOCK; the named consumer boundary contract; the precise novelty statement + the two honest qualifications; open questions. Then send the design-review request TO m-4.implementer (Template I), and report design-complete to the orchestrator for the c2 lock (after the COORD reconcile + the m-5 disposition). Do not self-advance to PLAN.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
