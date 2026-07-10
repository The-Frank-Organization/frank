## REVIEW-FOLD SCOPE - s2-core review-fold round 3 before edits

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-r3-fold-scope
PARENT_DISPATCH_ID: s2-core-review-r3
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - fold scope only; merge/S2-close and the Task-13.5 submit remain operator gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/REVIEW-FOLD-planner-20260704-143700.md
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: FOLD_SCOPE before editing s2-core-impl branch for sanctioned -mint admin-time flag

FOLD_SCOPE:
- cmd/frank/main.go -> in
- test/fixtures/main_assembly_test.go -> in
FOLD_SCOPE_RESULT: all-in

Scope basis:
- The orchestrator sanction relay `s2-core-impl/SITREP-orchestrator-planner-20260704-143000.md` rules option (a) in as a bounded fence amendment and rejects option (b).
- The planner round-3 review-fold relay requires one fold-sized commit adding a conductor-internal `-mint <seat>` admin-time flag with red-first fixture coverage.
- The test surface is the existing binary assembly fixture because the required behavior is an init/admin-time `cmd/frank` CLI path plus successful authenticated connect using the minted credential.

Rules:
- No edit outside the FOLD_SCOPE rows above.
- Any newly required file outside the rows above is a deviation and stops the fold before edit.
- No binding-table shape change.
- `-mint` is admin-time only, never a seat-facing verb, never in rendered tool registries.
- The real Task-13.5 OI-S1-F11-SWEEP submit remains operator-executed and out of scope.

ACTIONS_GIT_REF: no source edits yet; this FOLD_SCOPE relay written before implementation worktree edits; implementation branch remains `s2-core-impl@16342e0ce8fd28791fac21f261cc7404cab30d9b` before round-3 fold edits.
FINAL_GIT_STATUS_SHORT: main checkout tracked clean before FOLD_SCOPE write; implementation worktree clean at `s2-core-impl@16342e0ce8fd28791fac21f261cc7404cab30d9b` before round-3 fold edits.
