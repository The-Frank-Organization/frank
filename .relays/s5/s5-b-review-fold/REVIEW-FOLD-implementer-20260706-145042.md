## REVIEW-FOLD scope - s5-b MF-5 tree-invariant registry fixtures

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-b-review-fold
PARENT_DISPATCH_ID: s5-b-impl-fold
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-mechanisms
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-144638.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: FOLD_SCOPE before MF-5 tree-invariant §7 fixture edit
FOLD_SCOPE:
- test/fixtures/s5_config_change_test.go -> in
- test/fixtures/testdata/s5_pre_registry.json -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: no code edits yet; scope artifact only; pre-edit code worktree clean at s5-b-mechanisms@78bda2e
FINAL_GIT_STATUS_SHORT: code worktree clean at s5-b-mechanisms@78bda2e before fold edit

Scope basis:
- MF-5 names the existing `test/fixtures/s5_config_change_test.go` and a new `test/fixtures/testdata/` file as the only allowed surface.
- No production files, registry bytes, cmd wiring, or wider test surfaces are in scope.
