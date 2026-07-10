## Team m-4 — c2 fold-confirm (F2 + F3 + M4-1): bounded additive fold + an explicit confirm, implementer re-approve

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-fold-m-4
PARENT_DISPATCH_ID: c2-consumer-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded additive fold + a confirm of approved consumer findings; no new operator gate
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-4.planner
CC: m-4.implementer, m-3.planner, m-5.planner, m-6.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c2-design-m-4-routing-policy
BUNDLE_ID: m-4-routing-policy
OWNER: m-4 (Routing & Policy)

Basis: the c2 consumer-lens round is VP-approved (`c2-consumer-reconcile/RECONCILE-orchestrator-reviewer-20260630-034321.md`, verdict approve). m-5 + m-6 surfaced three items on YOUR routing record. Fold/confirm them — **bounded, additive only** — into the m-4 design doc (same DESIGN_DOC_ID), then send the design-review request to m-4.implementer and report fold-complete. Small fold; phase band stays DESIGN; no PLAN/IMPL.

**F2 — record a per-assignment home for `seat_archetype` + resolved `authority_ceiling` (the non-template path).** Your `routing_assignments` row has no archetype/ceiling column. For TEMPLATE-spawned teams `template_ref` carries the per-seat archetype; but for HAND-AUTHORED (non-template) staffing — e.g. a planner directly staffing a sensor seat — the ceiling that made write authority unavailable must still be **recorded for replay/audit**. **VP-approved (per-assignment field preferred):** add an **opaque** per-assignment `seat_archetype` (and/or resolved `authority_ceiling`) column to `routing_assignments` for replay/provenance; OR explicitly require all current archetype-bearing spawns to go through template records. Pick one; the per-assignment field is the replay-complete option both m-5 seats preferred.

**F3 — `seat_archetype` is a distinct opaque tag in the archetype vector (per-seat-at-spawn).** The tag-space is two orthogonal axes: `slot_in` = work-archetype (m-3's, per-record) vs `seat_archetype` = seat-archetype (YOURS, per-seat-at-spawn) — it keys the authority-ceiling-at-spawn (§8) + the capability recommendation (§10). Your design already says "archetype tag **vector**," so this fits. Record `seat_archetype` as a **distinct opaque tag** in that vector; `seat_archetype` is spawn-time (contrast `slot_in`, which is conductor-classified at work-record acceptance — m-3's fold).

**M4-1 — CONFIRM (explicit; a lock blocker if you cannot).** m-6 needs the routing **B→A escalation** to reach the human-surface as a **readable force-A atom on a consumable gate record** — not solely a `route_dispatch()` return — or it cannot distinguish an absorbed routing gate from an escalated one. **Confirm the VP-named route:** on `human_decision_required` / `routing_unavailable`, m-4 raises the **c1-locked HUMAN_GATE** (monotonic), with the reason in the A-set, so the escalation is readable+stamped on a consumable record; ordinary `routing` stays **category-B**. **Name the atom + the stamp path** in your fold. **If you CANNOT confirm this through the existing c1 HUMAN_GATE shape, say so immediately — the VP has named that a real c2 lock blocker** (do not improvise a new gate class).

**BOUNDED SHAPE (VP guardrail — do not exceed without flagging):** keep `seat_archetype` **opaque** and authority-bearing **only within the m-4 routing-record mechanism** (never outside it); do NOT define concrete Step-1 archetype/ceiling values, do NOT change m-2 ownership categories, do NOT narrow existing m-2 enum/row semantics in a breaking way (concrete tag-space + ceiling semantics stay **m-5-owned, c3**). If the fold needs any of those stronger moves, **STOP and flag an m-2 micro-fold to the orchestrator** — do not silently treat it as c2-local. Preserve R2 (no model-derived predicate enters any gate) and the reconciled m-3↔m-4 seam unchanged.

Deliverable: fold F2 + F3 + the M4-1 confirm into `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md` (same DESIGN_DOC_ID; a fold-log note), send the design-review request TO m-4.implementer (Template I), and report fold-complete to the orchestrator with the M4-1 confirm result explicit. **m-4.implementer re-approval is required before the c2 lock.** Do not self-advance to PLAN.

ACTIONS_GIT_REF: docs-workspace fold/confirm of approved consumer findings; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
