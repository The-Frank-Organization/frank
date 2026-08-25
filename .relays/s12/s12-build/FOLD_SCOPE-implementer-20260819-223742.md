## FOLD_SCOPE — s12 VP R3-F1 executable-guard durability correction

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-review-fold
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded test-only correction under the standing review-fold token
HUMAN_GATE_REQUIRED: no — VP r4 decides whether to route the separate operator-only MERGE-GATE
FILED_AT_LOCAL: 20260819-223742
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260819-222245.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
STARTING_HEAD: 08f18725c186d147a5671923a8fcd604d25cf66a
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: R3-F1 test and battery surfaces are all in; RED-first correction may proceed

FOLD_SCOPE:
- frank/test/fixtures/s12_nested_scope_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

The sole code-repository edit is the named test file: add a synthetic registered test whose only helper call is inside an uninvoked function literal, observe its RED leg against the current recursive inspector, then constrain the inspector to the directly executed top-level guard shape. The matcher preserves the exact startup crash-child prefix before that test's guard. Battery captures and this pair-local relay trail are governance output. No product, policy, design, frozen oracle, exit fixture, registry schema, timeout, merge, push, PR, deploy, or release byte is in scope.

ACTIONS_GIT_REF: none — pre-edit scope artifact only; implementation branch remains at 08f18725c186d147a5671923a8fcd604d25cf66a
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260819-223742.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260819-222245.md

Next requested action: add only the nested ineffective-occurrence negative, observe RED against the current inspector, then implement and verify the exact executable top-level guard constraint. MERGE-GATE remains HELD.
