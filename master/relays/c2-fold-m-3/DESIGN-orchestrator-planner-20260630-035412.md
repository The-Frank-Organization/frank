## Team m-3 — c2 fold-confirm (F1 + F3): bounded additive fold, implementer re-approve

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-fold-m-3
PARENT_DISPATCH_ID: c2-consumer-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded additive fold of approved consumer findings; no new operator gate
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-4.planner, m-5.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
BUNDLE_ID: m-3-observation-evidence
OWNER: m-3 (Observation & Evidence)

Basis: the c2 consumer-lens round is VP-approved (`c2-consumer-reconcile/RECONCILE-orchestrator-reviewer-20260630-034321.md`, verdict approve). m-5's consumer-lens surfaced two findings on YOUR `slot_in` archetype hook. Fold them — **bounded, additive only** — into the m-3 design doc (same DESIGN_DOC_ID), then send the design-review request to m-3.implementer and report fold-complete. This is a small fold, not a re-design; phase band stays DESIGN; no PLAN/IMPL.

**F1 — `slot_in` is conductor-owned / non-lane-writable, classified AT WORK-RECORD ACCEPTANCE.** Record that the work-archetype tag (`slot_in`) is **filled by the conductor and immutable after acceptance**, never lane-writable. This is the tamper-resistance the whole observe story depends on: if a lane could set its own work-archetype, it would re-tag `refactor`→`extension` to escape your no-test-edit invariant (or `bugfix`→other to escape red→green) — the exact reward-hacking failure your design defends against. **VP-pinned wording (use this, reject the alternative):** classified **at work-record acceptance**, NOT "stamped from the seat's spawn-time binding." Spawn-time would wrongly forbid a long-lived seat from doing bugfix→refactor→migration over its life. The invariant that matters is **non-lane-writability + immutable conductor classification**, not spawn-derivation. Use the **already-reserved m-2 opaque `slot_in` atom shape** — select its filler as system/conductor within the existing m-2 field-ownership categories.

**F3 — confirm `slot_in` = the WORK-ARCHETYPE axis only (per work record).** The archetype tag-space is two orthogonal axes: `slot_in` = work-archetype (yours, per-record — extension/refactor/cleanup/bugfix/migration) vs `seat_archetype` = seat-archetype (m-4's, per-seat-at-spawn — sensor/implementer/…). Confirm your done-predicate hook keys on `slot_in` = **work-archetype only**; the seat-archetype axis is m-4's (separate opaque tag in its routing vector). Don't absorb seat/ceiling semantics into `slot_in`.

**BOUNDED SHAPE (VP guardrail — do not exceed without flagging):** keep `slot_in` **opaque** — do NOT define concrete Step-1 values, do NOT add a required-when/visible-when branch on concrete slot values, do NOT change m-2 ownership categories or the bounded predicate vocabulary (the concrete tag-space + invariant selection stay **m-5-owned, c3**). If the fold tempts you toward any of those stronger moves, **STOP and flag an m-2 micro-fold to the orchestrator** — do not silently treat it as c2-local. The fold is *additive provenance + an axis confirmation*, nothing more.

Deliverable: fold F1 + F3 into `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md` (same DESIGN_DOC_ID; a fold-log note), send the design-review request TO m-3.implementer (Template I), and report fold-complete to the orchestrator. **m-3.implementer re-approval is required before the c2 lock.** Do not self-advance to PLAN.

ACTIONS_GIT_REF: docs-workspace fold of an approved consumer finding; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
