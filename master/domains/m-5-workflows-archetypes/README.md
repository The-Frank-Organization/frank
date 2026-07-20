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
m-5 design-of-record = Step 0 (COMPLETE; current milestone = Step-3 DESIGN). m-5 *product feature* (workflows / recursion / nested teams) = Step 5.

## Status
- **c3 audit RECONCILED (F4 ✓).** Independent passes (`053308` planner + `053116` implementer) + two convergent
  pair-reconciles (`120326` planner + `120346` implementer). Verdict **still-open / net-new binding over a 4-mechanism
  PROMOTE base**; the five items **structure-converged + values carried to the DESIGN-grill**: actuator → derived
  `actuator_class` (no literal seat value v3.0); read-only work-archetypes surfaced (`research_synthesis`/`qa_review`/
  `docs_chore`); human-mode → **two-layer** `human_mode` posture × `surface_intent` (the m-5↔m-6 COORD content); ceiling
  → **partial-order/vector** (not a total ladder); naming → `lower_snake_case`. Joint audit-reconcile relayed to the VP
  (`c3-reconcile/RECONCILE-orchestrator-planner-20260630-120832.md`); on VP approve → c3 DESIGN (GRILL-gated).
  (c2 narrow consumer-lens — see Engagement above.)
- c2 narrow consumer-lens delivered (`c2-consumer-review-m-5`, 20260629) — reserved to c3 (above).

## Layout
- `audit/` — consumer-review + c3 AUDIT artifacts.  `design/` — c3 DESIGN docs + grill locks.

## Step-3 reframe delta (RATIFIED 2026-07-15 — `master/STEP-3-ARCH-AMENDMENT.md` @ `2d240eb6…`)
m-5 stays the **SOLE policy owner** of the authority ceiling. The ceiling **ENFORCEMENT HOST moves to m-10** (app-side) via the **m-5-authored ceiling-host amendment** — the *coordinated first stage* (interface-locks with the m-10 boundary design **before** any m-8/m-9 consumer lock). The ceiling **artifact interface** — source/writer/schema-home, **immutable binding to `run_id` + worker identity**, the m-10 read/load path, **fail-closed when absent/stale** — has its **bytes PINNED in the canonical contract (`643dd7c2…`) and pair-approved, but is NON-CONSUMABLE until the Master+VP interface-lock** (the *interface-lock event*, not the bytes, is what is not yet issued). m-5 does **NOT** lose policy ownership; m-10 is enforcement host only.

**PENDING / NON-CONSUMABLE (VP F20).** This ceiling-host amendment is a **PROPOSAL** until (a) the m-5 pair review closes AND (b) the Master+VP first-stage join issues the one canonical **ceiling-interface-lock** event. Until that lock: **the locked m-5 enforcement text (conductor / host-config enforcement, `…/2026-06-30-v3-archetype-system-design.md:158-174`) REMAINS OPERATIVE** — this amendment does **not** silently rewrite it — and **no m-10/m-9 consumer may consume the proposed interface**. The re-dispatch stage is the **coordinated first stage** (`step3-amend-m5-ceiling`). **Live state:** the amendment is **DESIGN-COMPLETE + implementer-approved (PROVISIONAL, not a lock)** — one canonical ceiling-artifact contract `design/2026-07-15-ceiling-artifact-contract.md` @ SHA-256 `643dd7c2…`, GRILL_LOCK `step3-amend-m5-ceiling-grill` closed, report-only (SITREP `…-091000`, no self-declared lock). The **Master+VP first-stage interface-lock is the gate** — but it is NOT yet reachable: **m-10 has only hash-confirmed the bytes in COORD `091500` + received the direct clarification `092000`; it has returned NO DESIGN, GRILL_LOCK, implementer DESIGN-REVIEW, or report-only completion SITREP** (VP F28 — COORD/hash convergence is not an approved m-10 artifact, so this amendment's own SITREP `092000` readiness-to-lock claim is SUPERSEDED). The still-owed sequence: m-10 DESIGN + GRILL_LOCK → implementer child review → planner report-only SITREP → THEN Master+VP reconcile over both approved artifacts. The `config_generation` (m-7/m-1) app-side read-path: **owner legs discharged; OPERATOR ELECTED BRANCH B** (positive Step-3 tools, 2026-07-15, after VP `144149` F33). The owner returns proved **no packet-compliant live-freshness read exists** ⇒ under the unchanged contract `643dd7c2…` the packet-preserving reading is **deny-all** (`tool→none`) — NOT the pinned positive ceiling an earlier master relay mislabeled (VP F33). **STOOD DOWN for the MVP (operator, 2026-07-15 — MVP scope reframe).** The operator re-cut the MVP to **defer the entire permission/authority system to Step-4**: the MVP ships a **built-but-EMPTY permission seam** (authorization = a trivial static run-manifest allow-list — no config-derived ceiling, no `config_generation` freshness). This **dissolves the seam-13 knot** and **stands down the m-5 MVP amendment** — the `145500` permissive-tools authoring task and the mis-framed `133500` nod are **WITHDRAWN** (`step3-amend-m5-ceiling/…-152000`); m-5 authors **no MVP amendment**. **m-5's ceiling contract `643dd7c2…` (untouched) + `config_generation` freshness + the per-role ceiling are the STEP-4 basis** — they plug into the m-10 enforcement seam when Step-4 opens. The governance-model refinement (`FRANK-HARDENING-BACKLOG.md` H-11) + item-(b) remain m-5 **Step-4 backlog**. m-5 holds; no MVP work owed.

**MVP amendment RATIFIED (2026-07-16, r7 @ `2f75f2a1…`):** the stand-down above is now operative ratified text (§0: "the entire permission/authority system → Step-4"; §4: the MVP seam = the operator-fixed 8-name dispatch constant, m-10-hosted). **No MVP work owed from m-5**; the untouched ceiling contract `643dd7c2…` + `config_generation` freshness + per-role permissions remain the **Step-4 basis** plugging into the same m-10 seam (+ H-11 backlog).
