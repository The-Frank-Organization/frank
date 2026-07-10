## REVIEW-FOLD SCOPE - s2-core review-fold round 2 before edits

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-r2-fold-scope
PARENT_DISPATCH_ID: s2-core-review-r2
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - fold scope only; merge/S2-close remain operator gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/REVIEW-FOLD-planner-20260704-134500.md
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: FOLD_SCOPE before editing s2-core-impl branch for RB2-1 applicability-map falsifiability and RB2-2 O3 owed-item fixture

FOLD_SCOPE:
- internal/crashpoint/crashpoint.go -> in
- internal/crashpoint/crashpoint_test.go -> in
- test/fixtures/applicability_map.go -> in
- test/fixtures/f11_test.go -> in
- test/fixtures/main_assembly_test.go -> in
- test/fixtures/s2_sweep_test.go -> in
- docs/sprints/2026-07-03-s2-slice-2/results/f11-sweep-report.md -> in
FOLD_SCOPE_RESULT: all-in

Scope basis:
- RB2-1: add a test-only crashpoint hit trace, prove each F11 class row equals the trace-observed fired set, correct the applicability map rows, and update the F11 sweep report wording.
- RB2-2: add an end-to-end fixture-store O3 owed-item/disposition proof through the operator channel, including projection/read evidence and open-set drain.

Rules:
- No edit outside the FOLD_SCOPE rows above.
- Any newly required file outside the rows above is a deviation and stops the fold before edit.
- S1 assertions remain byte-identical.
- Merge/S2-close and the real operator-owned Task-13.5 OI submit remain out of scope.

ACTIONS_GIT_REF: no source edits yet; this FOLD_SCOPE relay written before implementation worktree edits; implementation branch remains `s2-core-impl@9e4829c3ddfa7a3521de1b0b84952d18b5035b31` before round-2 fold edits.
FINAL_GIT_STATUS_SHORT: main checkout tracked clean before FOLD_SCOPE write; implementation worktree clean at `s2-core-impl@9e4829c3ddfa7a3521de1b0b84952d18b5035b31` before round-2 fold edits.
