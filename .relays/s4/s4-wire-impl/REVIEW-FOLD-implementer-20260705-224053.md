## REVIEW-FOLD - s4-wire-impl F-GATE-2 fold scope before edits

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-impl
RUN_ID: s4
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no new gate; orchestrator-routed hygiene items remain out of fold scope
DESIGN_LOCK_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s4-grill-s4-wire
PLAN_LOCK_ID: s4-slice-4-plan
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator, m-2.planner
IN_REPLY_TO: s4-wire-impl/SITREP-planner-20260705-223638.md
BASE: s4-wire-impl@605b3ef
BRANCH: s4-wire-impl
SUBJECT: FOLD_SCOPE before F-GATE-2 edits - five owed FieldSpec rows plus rendered-form and validate fixtures

FOLD_SCOPE:
- internal/fieldspec/registry.json -> in
- internal/fieldspec/registry_test.go -> in
- test/fixtures/s4_shim_test.go -> in
FOLD_SCOPE_RESULT: all-in

Scope notes:
- `internal/fieldspec/registry.json` is in for the five confirmed FieldSpec rows: `owner`, `source`, `target_surface`, `disposition_path`, and `disposes_owed`.
- `internal/fieldspec/registry_test.go` is in for row-shape and `required_when`/validate predicate fixtures.
- `test/fixtures/s4_shim_test.go` is in for the describe-grade rendered-form fixture.
- The orchestrator-routed hygiene items in the planner relay are not fold content and are not touched here.

ACTIONS_GIT_REF: relay-substrate write only before code edits; branch remains `s4-wire-impl@605b3ef` at scope filing time.
FINAL_GIT_STATUS_SHORT: code worktree clean before fold edits.
