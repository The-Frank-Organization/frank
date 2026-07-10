## REVIEW-FOLD scope - s5-a r2 annotation must-fixes

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-a-impl-r2
PARENT_DISPATCH_ID: s5-a-impl-r2
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
BRANCH: s5-a-registry
IN_REPLY_TO: .relays/s5/s5-a-impl-r2/RECONCILE-planner-20260706-090059.md
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: FOLD_SCOPE before edits for F-SEC-1 and F-SEM-1 annotation bytes only
FOLD_SCOPE:
- internal/fieldspec/registry.json -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: no code edits claimed for this fold yet; s5-a-registry@dd8189d; code worktree status clean before edit
FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/

Scope basis:
- F-SEC-1: replace the `model_name` annotation's non-lane-writability overclaim with the render-absent/raw-submit-suppliable wording from design rev4.
- F-SEM-1: restore the two m-1 parentheticals in the `record_kind` annotation and preserve the route-back sentence.
- Planner punctuation ruling: restore byte-verbatim em dashes in the marked m-1 and m-6 quotes, and restore the `§J1` reference in the `on_timeout` annotation.

Not accepted for this fold:
- Optional test comments in `test/fixtures/s5_registry_dormancy_test.go`.
- Any mechanism, validator, renderer, engine, fixture-logic, migration, lineage, store, s5-b, merge, push, PR, or live-verify work.
