## REVIEW-FOLD scope — T7 E1 machinery-fault classification

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s8-build-t7-fold
PARENT_DISPATCH_ID: s8-build-t8-review-verdict
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded must-revise fold under the standing s8-build-impl token; no merge authority
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T7
IN_REPLY_TO: /Users/jack/Programming/harness/master/relays/s8-build-t8-review/SITREP-planner-20260711-231500.md
FROM: s8.implementer
TO: s8.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-7.planner, m-2.planner
SUBJECT: pre-edit fold scope for the directly addressed T7 E1 machinery-fault classification findings F1-F3
FOLD_SCOPE:
- internal/observe/checks_base.go -> in
- internal/observe/registry.go -> in
- test/fixtures/s8_decision2_test.go -> in
- test/fixtures/s8_check_registry_e1_test.go -> in
- .relays/s8/ -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: no code edits yet; scope artifact only; pre-edit worktree clean at s8-observe-spine@59f01df
FINAL_GIT_STATUS_SHORT: code worktree clean at s8-observe-spine@59f01df before fold edits

Scope basis:
- The planner's F1-F3 findings name one seam: E1 check machinery classification versus observed absence.
- The two production files and two fixture files are the exact fold surfaces named by the verdict; `.relays/s8/` is included only for the required durable scope and return artifacts.
- T8 is approved and untouched. T9/T10 remain outside this fold and hard-gated on m-7 catalog bytes.
