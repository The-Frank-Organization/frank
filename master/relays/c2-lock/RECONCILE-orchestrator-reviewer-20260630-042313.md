## RECONCILE -- master.orchestrator-reviewer / c2 lock review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c2-lock
PARENT_DISPATCH_ID: c2-lock
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- revise c2 design-lock text before co-sign; no PLAN/IMPL authority granted
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-4.planner, m-5.planner, m-6.planner

Verdict: revise.

Scope reviewed:
- `master/relays/c2-lock/RECONCILE-orchestrator-planner-20260630-041627.md`
- `master/relays/c2-lock-prep/RECONCILE-orchestrator-reviewer-20260629-212213.md`
- `master/relays/c2-consumer-reconcile/RECONCILE-orchestrator-reviewer-20260630-034321.md`
- `master/relays/c2-fold-dispatch/RECONCILE-orchestrator-reviewer-20260630-035726.md`
- `master/relays/c2-fold-m-3/DESIGN-planner-20260630-040131.md`
- `master/relays/c2-fold-m-3/DESIGN-REVIEW-implementer-20260630-040633.md`
- `master/relays/c2-fold-m-4/DESIGN-planner-20260630-040400.md`
- `master/relays/c2-fold-m-4/DESIGN-REVIEW-implementer-20260630-040641.md`
- `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md`
- `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md`
- `master/ARCHITECTURE.md`

Finding 1 -- the c2 evidence chain is otherwise green.

m-3 and m-4 both have approved r1 designs and approved fold re-reviews. The m-3 fold keeps `slot_in`
conductor-classified at work-record acceptance, immutable, non-lane-writable, and outside lane control.
The m-4 fold records opaque `seat_archetype` plus resolved `authority_ceiling` per assignment, and M4-1
is confirmed through the existing c1 monotonic `HUMAN_GATE` "routing-raise" path. The consumer folds remain
bounded-additive, and the planner lock relay preserves the prior provenance caveat by saying GL-4 / m-5
narrow scope is operator-directed by current session context rather than asserting a missing `FROM: operator`
relay.

Finding 2 -- lock-blocker: the design-of-record still names concrete archetype values while claiming none
are locked.

The lock relay correctly says the c2 lock does not define concrete Step-1 values. But the integrated
architecture and m-3 rev2 doc still contain value lists close enough to the lock text that a later reader
could treat them as the locked tag-space:

- `master/ARCHITECTURE.md:181-183` says `slot_in` is "extension/refactor/cleanup/bugfix/migration" and
  `seat_archetype` is "sensor/implementer/..." before `master/ARCHITECTURE.md:189-190` says both atoms stay
  opaque and there are "No concrete Step-1 values."
- `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:88-98` presents
  candidate mappings as surfaced/not closed, but `:102` repeats the same `extension/refactor/cleanup/bugfix/migration`
  list inside the folded axis definition, immediately before `:103` says no concrete values are defined.

This is not a mechanical lint error, but it is a design-lock ambiguity. C2's no-m-2-micro-fold conclusion
depends on `slot_in` and `seat_archetype` remaining opaque and on concrete tag-space / invariant selection /
ceiling semantics staying m-5-owned in c3. The current wording lets the lock both reserve and name the values.

Required revision before VP co-sign:
- In `master/ARCHITECTURE.md` C2.4, remove the concrete parenthetical value lists or explicitly label them
  as non-locking examples/candidates only, with the opaque c3 reservation as the operative lock rule.
- In the m-3 design doc's folded §5.1 axis sentence, remove the concrete `slot_in` value list or explicitly
  mark it as non-locking candidate vocabulary inherited from the surfaced table, not Step-1 locked values.
- Keep the m-5/c3 reservation intact: concrete tag-space, invariant maps, default per-archetype gate composition,
  routing-template structures/lineup, and authority-ceiling semantics remain outside the c2 lock.

Finding 3 -- no broader reopen is required if the revision is only a de-locking clarification.

If the next relay only clarifies that the listed archetype values are examples/candidates and not locked
Step-1 values, I do not require another consumer round or implementer re-review. The underlying folds,
M4-1 confirm, R2 ratification, m-6 reader/writer clearance, m-5 reservation, and AUDIT+DESIGN phase boundary
are sufficient after that textual correction.

Approved next action:
- Patch the two lock-text ambiguities above.
- Return a narrow c2-lock follow-up relay for co-sign.

Not authorized:
- no c2 close on this relay;
- no PLAN;
- no IMPL;
- no merge;
- no pcode/source changes;
- no full m-5 archetype-system lock in c2.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
