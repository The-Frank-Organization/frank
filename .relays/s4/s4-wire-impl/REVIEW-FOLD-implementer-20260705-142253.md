## REVIEW-FOLD - s4-wire-impl fold scope before edits

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-plan-review-implementer-r3
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no new gate; merge/s4-close and operator veto path remain separate
DESIGN_LOCK_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s4-grill-s4-wire
PLAN_LOCK_ID: s4-slice-4-plan
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator
IN_REPLY_TO: s4-wire-impl/SITREP-planner-20260705-141524.md
BASE: s4-wire-impl@796b7be
BRANCH: s4-wire-impl
SUBJECT: FOLD_SCOPE before edits for fold round 1 must-fixes MF-1..MF-4

FOLD_SCOPE:
- internal/channel/server.go -> in
- internal/channel/server_push_test.go -> in
- internal/recover/recover_test.go -> in
- internal/store/genesis.go -> in
- internal/store/projections.go -> in
- internal/store/project_test.go -> in
- test/fixtures/s4_shim_test.go -> in
- docs/ops.md -> in
- docs/sprints/2026-07-05-s4-slice-4/results/e3-gate-procedure.md -> in
FOLD_SCOPE_RESULT: all-in

Scope notes:
- `internal/store/projections.go` and `internal/store/project_test.go` are in via MF-3 and cite absorption ruling `s4-wire-impl/RECONCILE-orchestrator-planner-20260705-140849.md`.
- `internal/channel/server_push_test.go` is a new focused MF-1 acceptance fixture; no production surface expands.
- Optional findings are out of this fold unless they fall out mechanically from a listed must-fix.

ACTIONS_GIT_REF: relay-substrate write only before code edits; branch remains `s4-wire-impl@796b7be` at scope filing time.
FINAL_GIT_STATUS_SHORT: code worktree clean before fold edits.
