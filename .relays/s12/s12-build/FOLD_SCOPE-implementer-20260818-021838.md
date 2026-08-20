## FOLD_SCOPE — s12 H-16 rev21 end-review correction for the pre-serve evidence fold and Class-G dirty barrier

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-impl
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded implementation-fidelity correction against the frozen rev21 design
HUMAN_GATE_REQUIRED: no — this fold runs under the standing IMPL token; the operator MERGE-GATE remains held
FILED_AT_LOCAL: 20260818-021838
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260818-021247.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
STARTING_HEAD: 692a3af9c29535cdbf4fe81cd6f316ea12d50bdd
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
SUBJECT: all routed F1/F2 correction rows are in; RED-first implementation may proceed under the existing token

FOLD_SCOPE:
- frank/internal/engine/loop.go -> in
- frank/cmd/frank/main.go -> in
- frank/test/fixtures/main_assembly_test.go -> in
- frank/test/fixtures/h16_startup_evidence_test.go -> in
- frank/test/fixtures/h16_classg_barrier_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

The production change is confined to the existing startup/Class-G seams named by the planner. The `main_assembly_test.go` row permits only recognition of the new process-crash child mode; the two new fixture files carry the routed regressions. Forward captures under the battery directory will be self-describing with command, exit status, source branch and tip, timestamps, and contention conditions. No frozen oracle, exit fixture, INV-CATALOG law, design, plan, registry, S8-precedence, H-12, merge, push, PR, deploy, or release byte is in scope.

ACTIONS_GIT_REF: none — pre-edit scope artifact only; implementation branch remains clean at 692a3af9c29535cdbf4fe81cd6f316ea12d50bdd
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260818-021838.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260818-021247.md

Next requested action: execute the two routed regressions RED-first, implement only the matching runtime corrections, verify under an explicitly uncontended quiet window, and return the refreshed evidence set to s12.planner. MERGE-GATE remains HELD.
