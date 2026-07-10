## RECONCILE -- master.orchestrator-reviewer / c2 consumer reconcile review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c2-consumer-reconcile
PARENT_DISPATCH_ID: c2-consumer-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- approve bounded fold-confirm routing; no c2 lock/PLAN/IMPL authority granted
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-4.planner, m-5.planner, m-6.planner

Verdict: approve.

Scope reviewed:
- `master/relays/c2-consumer-reconcile/RECONCILE-orchestrator-planner-20260629-230608.md`
- `master/relays/c2-consumer-review-m-5/AUDIT-planner-20260629-213422.md`
- `master/relays/c2-consumer-review-m-5/AUDIT-implementer-20260629-213807.md`
- `master/relays/c2-consumer-review-m-6/AUDIT-planner-20260629-214112.md`
- `master/relays/c2-consumer-review-m-6/AUDIT-implementer-20260629-213803.md`
- `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md`
- `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md`
- `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md`
- `master/ARCHITECTURE.md`

Finding 1 -- consumer clearance is substantively supported.

m-5's two independent passes converge on `FITS-with-folds`: the opaque-tag seam is viable, but c2 should record the two-axis split, non-lane-writable work-archetype provenance, a recorded home for seat-archetype / authority ceiling, and the sensor integrity split. m-6's two independent passes converge on "sufficient" for the human-surface reader/writer check, with M4-1 as a confirm item rather than a blocker. The pairs did not file intra-pair reconcile relays, which is not ideal relative to the dispatch, but the four source passes are visible, non-conflicting on the lock-critical points, and the planner's synthesis preserves the sharper implementer corrections. I do not require another consumer round before fold-confirm.

Finding 2 -- F3 two-axis tag split is the right lock-time clarification.

The split between `work_archetype` / `slot_in` per work record and `seat_archetype` per seat/spawn is necessary to avoid overloading one atom with both m-3 done-predicate semantics and m-4 authority-ceiling/routing semantics. This is compatible with m-4's existing "archetype tag vector" language and with c1's `slot_in` reservation, provided the lock keeps both values opaque and reserves concrete tag-space to m-5/c3.

Finding 3 -- F1 is additive only with the implementer wording, not the original spawn-derived wording.

Approve the planner reconcile's corrected provenance split: `seat_archetype` is spawn-time; `work_archetype` / `slot_in` is conductor-owned and classified at work-record acceptance. Do not fold the earlier m-5 planner wording that `slot_in` is always stamped from the seat's spawn-time binding. That would incorrectly prevent long-lived seats from moving across bugfix/refactor/migration work types. The invariant that matters for c2 is non-lane-writability and immutable conductor classification, not spawn-time derivation.

Finding 4 -- the two m-2-adjacent folds do not reopen c1 if kept inside the bounded shape below.

Approved as additive:
- `seat_archetype` / resolved `authority_ceiling` may be added as opaque replay/provenance on the m-4 `routing_decision` assignment row, or m-4 may explicitly require all archetype-bearing spawns at Step-1 to go through template records. The per-assignment field is the better replay-complete option.
- `slot_in` may be recorded as conductor-owned / non-lane-writable work-record classification, using the already-reserved m-2 atom shape.

Not approved without an m-2 micro-fold:
- defining concrete Step-1 `slot_in` values;
- adding a required-when / visible-when branch on concrete slot values;
- changing m-2 ownership categories or the bounded predicate vocabulary;
- making `seat_archetype` authority-bearing outside the m-4 routing-record mechanism;
- narrowing existing m-2 enum/row semantics in a breaking way.

If m-3 or m-4 needs any of those stronger moves, route a flagged m-2 micro-fold instead of silently treating it as c2-local.

Finding 5 -- M4-1 is correctly routed as a confirm, but the confirm must be explicit before c2 lock.

m-6 can consume routing B-to-A escalation if the escalation is a readable force-A atom on a consumable gate record. The natural route is the c1-locked HUMAN_GATE mechanism: `human_decision_required` / `routing_unavailable` raises the monotonic human gate with a reason in the A set, while ordinary `routing` remains category B. m-4 should confirm the atom and stamp path in its fold-confirm. If m-4 cannot confirm this through the existing HUMAN_GATE shape, that becomes a real lock blocker.

Finding 6 -- sensor/template outputs are lock-time surface, not a c2 full m-5 lock.

Approve recording Solo / Adversarial Pair / Sensor as the Step-1 proposed lineup and recording the sensor archetype split: answer content is advisory/self-reported and never gate-bearing, while observable runtime metadata such as tool-blocked, no source actions, and one-turn closure can be stamped as observed where the conductor exposes those facts. The concrete template schema, exact ceilings, invariant maps, default gate composition, and full m-5 archetype semantics remain c3-owned.

Approved next action:
- Route the brief m-3 fold-confirm for F1 and F3.
- Route the brief m-4 fold-confirm for F2, F3, and M4-1.
- Require each fold-confirm to receive the relevant implementer approval before the c2 lock.
- Then bring the c2 lock relay, including the m-5 reservation-of-the-rest and the operator-directive provenance item already called out by the planner.

Not authorized:
- no PLAN;
- no IMPL;
- no merge;
- no pcode/source changes;
- no c2 lock before fold-confirm returns are visible and reconciled;
- no full m-5 archetype-system design lock in c2.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
