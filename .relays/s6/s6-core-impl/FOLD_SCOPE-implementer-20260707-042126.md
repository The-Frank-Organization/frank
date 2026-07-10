## FOLD_SCOPE - s6-core implementer - s6-core-impl-fold

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s6-core-impl-fold
PARENT_DISPATCH_ID: s6-core-impl
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: REVIEW-FOLD-planner-20260707-041710.md
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: pre-edit fold scope for items 1-7 from the s6-core REVIEW-FOLD directive

FOLD_SCOPE:
- internal/engine/loop.go -> in
- internal/engine/loop_test.go -> in
- internal/engine/submit.go -> in
- internal/fieldspec/render.go -> in
- internal/tables/generation.go -> in
- internal/tables/generation_test.go -> in
- internal/store/lock.go -> in
- internal/store/lock_test.go -> in
- test/fixtures/s6_lifecycle_test.go -> in
- test/fixtures/sweep_test.go -> in
- test/fixtures/s6_lock_test.go -> in
- cmd/frank/main.go -> in
- .relays/s6/s6-core-impl/REVIEW-FOLD-implementer-20260707-*.md -> in
- .relays/s6/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

Scope rationale:
- `internal/engine/loop.go` and `internal/engine/loop_test.go` cover must-fix item 1, including the named `TestCommitGuardBlocksSecondOutcome` and `faultOutcome`/`AfterAccepted` commit-boundary behavior.
- `internal/engine/submit.go`, `internal/fieldspec/render.go`, and `test/fixtures/s6_lifecycle_test.go` cover must-fix item 2 and item 5 boot-smuggle legs.
- `test/fixtures/sweep_test.go` covers must-fix items 3 and 4.
- `internal/tables/generation.go` and new `internal/tables/generation_test.go` cover must-fix item 5 FX-B1f lifecycle exactly-once legs.
- `internal/store/lock.go`, `internal/store/lock_test.go`, `test/fixtures/s6_lock_test.go`, and `cmd/frank/main.go` cover must-fix item 6 FX-A4b store-root lock takeover and alias-path assertions.
- The fold report relay plus run index cover must-fix item 7 and the required durable handoff.
- No `internal/gc/gc_test.go` row is included, so no absorption-ruling row citation is required in this scope.

ACTIONS_GIT_REF: no code edits claimed; scope artifact only. Implementation branch before fold edits: `s6-transport-impl@bfbbb2d`.
FINAL_GIT_STATUS_SHORT: not applicable - pre-edit scope artifact lives under gitignored `.relays/s6/`; code worktree status was clean before this artifact.
Next requested action: proceed with red-first fold edits inside the rows above.
