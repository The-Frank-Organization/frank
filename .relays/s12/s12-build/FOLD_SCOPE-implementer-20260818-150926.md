## FOLD_SCOPE — s12 VP R2 durability binding and unchanged-boundary reproduction

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-review-fold
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded test-only correction and reproduction under the standing token
HUMAN_GATE_REQUIRED: no — m-3 owns any non-firing reproduction disposition; the operator MERGE-GATE remains held
FILED_AT_LOCAL: 20260818-150926
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260818-145648.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
STARTING_HEAD: 1ac14a019f4122a7e43a7f2d17c465ba331e1985
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: all R2 rows are in; AST binding RED and fixed-budget historical reproduction may proceed

FOLD_SCOPE:
- frank/test/fixtures/s12_nested_scope_test.go -> in
- frank/test/fixtures/s12_historical_boundary_repro_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

`s12_nested_scope_test.go` is the only landed source surface: it will gain a Go-AST inspector binding every registry name to a `s12SkipOuterRunOnly(t)` call in that named test body, plus a synthetic missing-call negative proof observed RED before the implementation. `s12_historical_boundary_repro_test.go` is a disclosed scratch harness pinned to the unchanged historical 1731ms live budget; it will be removed before staging and will not land. The battery directory carries the immutable predeclared reproduction budget and forward-bound results. No shipped timeout value, other test fixture, product, policy, script, design, plan, merge, push, PR, deploy, or release byte is in scope.

ACTIONS_GIT_REF: none — pre-edit scope artifact only; implementation branch remains clean at 1ac14a019f4122a7e43a7f2d17c465ba331e1985
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260818-150926.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260818-145648.md

Next requested action: prove the missing-call case RED, implement the deterministic AST binding, then run the predeclared unchanged-1731ms reproduction budget without tuning conditions. MERGE-GATE remains HELD.
