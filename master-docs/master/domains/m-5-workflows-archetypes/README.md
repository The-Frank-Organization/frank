# m-5 — Workflows & Archetypes

**Pair:** design-lead `m-5.planner` + adversarial design-reviewer `m-5.implementer`.
**Engagement:** **c3 full archetype-system design (lead)** — the last Step-0 design domain, co-designed with m-6.
m-5 was booted mid-c2 for a narrow consumer-lens (VP-approved); those outputs are **reserved to c3** and are now
the design inputs.

## Owns (the durable domain)
Expansion-slot presets (topology + gate-set + **human-mode**), the **tag-space**, per-archetype observe invariants,
**authority-ceiling-at-spawn**, and **sensor/actuator** archetypes. Consumes m-3 mechanism, m-4 routing, m-2 schema
(all locked).

## c3 scope (VP-approved — `c3-decomp` 20260630-051448)
The full archetype-system design-of-record, binding the c2-reserved opaque atoms (`slot_in` / `seat_archetype`)
into concrete semantics:
- **Concrete tag-space** — `slot_in` work-archetype values + `seat_archetype` values (surfaced in audit; locked at
  DESIGN, m-5-owned).
- **Per-archetype invariant composition** — observe invariants (m-3) × default gate-set × authority-ceiling-at-spawn
  × routing-prior (m-4). The composition rule: `seat_archetype` (spawn-fixed) ⊗ `slot_in` (per-record).
- **The authority-ceiling lattice** — concrete semantics behind the locked opaque `authority_ceiling`.
- **The template lineup** — T1 Solo / T2 Adversarial Pair / T3 Sensor structures (GL-4); conductor/N-pair template
  deferred (Step-5).
- **Sensor/actuator archetypes** — the c2 sensor sketch → full; actuator (literal value vs derived ceiling = grill).
- **Seam with m-6** — declare the per-archetype **human-mode vocabulary** (declare-before-bind, F2) + the
  interjection surface contract (Seam A/B).

## Consumes (locked upstream — do NOT reopen)
m-2 schema (the opaque atoms reserved to m-5); m-3 observe mechanism (parameterized by work-archetype); m-4 routing
+ the `seat_archetype` capability key + the GL-4 template record mechanism; M4-1 routing-raise.

## Roadmap mapping (`ROADMAP.md`)
m-5 design-of-record = Step 0 (now). m-5 *product feature* (workflows / recursion / nested teams) = Step 5.

## Status
- **c3 audit RECONCILED (F4 ✓).** Independent passes (`053308` planner + `053116` implementer) + two convergent
  pair-reconciles (`120326` planner + `120346` implementer). Verdict **still-open / net-new binding over a 4-mechanism
  PROMOTE base**; the five items **structure-converged + values carried to the DESIGN-grill**: actuator → derived
  `actuator_class` (no literal seat value in the initial release); read-only work-archetypes surfaced (`research_synthesis`/`qa_review`/
  `docs_chore`); human-mode → **two-layer** `human_mode` posture × `surface_intent` (the m-5↔m-6 COORD content); ceiling
  → **partial-order/vector** (not a total ladder); naming → `lower_snake_case`. Joint audit-reconcile relayed to the VP
  (`c3-reconcile/RECONCILE-orchestrator-planner-20260630-120832.md`); on VP approve → c3 DESIGN (GRILL-gated).
  (c2 narrow consumer-lens — see Engagement above.)
- c2 narrow consumer-lens delivered (`c2-consumer-review-m-5`, 20260629) — reserved to c3 (above).

## Layout
- `audit/` — consumer-review + c3 AUDIT artifacts.  `design/` — c3 DESIGN docs + grill locks.
