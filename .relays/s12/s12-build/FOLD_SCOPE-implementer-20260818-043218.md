## FOLD_SCOPE — s12 m-3 flake ruling, test-only nested narrowing and executor diagnosis

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-review-fold
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded test-only correction under the ratified m-3 ruling
HUMAN_GATE_REQUIRED: no — this fold runs under the standing IMPL token; the operator MERGE-GATE remains held
FILED_AT_LOCAL: 20260818-043218
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260818-042436.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
STARTING_HEAD: 8ca99650019f2ddd12d47237e1abb32fbd5895e8
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: all routed test-only rows are in; diagnosis may proceed before the licensed remedies

FOLD_SCOPE:
- frank/test/fixtures/s8_adversarial_test.go -> in
- frank/test/fixtures/s8_exit_gate_test.go -> in
- frank/test/fixtures/s12_nested_scope_test.go -> in
- frank/test/fixtures/ceremony_test.go -> in
- frank/test/fixtures/ceremony_retry_test.go -> in
- frank/test/fixtures/h16_startup_evidence_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

The first edit is a bounded diagnostic on the existing real executor probe; its purpose is to expose the executor verdict under load before any remedy. The licensed remedy may then change only test fixture behavior: one authoritative executable outer-run-only registry and same environment guard, guards on the enumerated crash-cut/ceremony/startup tests, the exit-gate reachability witness, and a unique load-safe timeout sentinel that remains in the forbidden-leak assertion. Forward captures under the battery directory will bind command, exit, branch tip, timestamp, contention, nested-leg time, and numeric margins against the unchanged 150-second test context and unchanged 120-second executor cap. Product code, policy values, scripts, merge, push, PR, deploy, and release bytes are out of scope.

ACTIONS_GIT_REF: none — pre-edit scope artifact only; implementation branch remains clean at 8ca99650019f2ddd12d47237e1abb32fbd5895e8
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260818-043218.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260818-042436.md

Next requested action: run the bounded load diagnosis, honor the block branch if no executor detail is observed, otherwise implement only the licensed test remedies and verify in a quiet window. MERGE-GATE remains HELD.
