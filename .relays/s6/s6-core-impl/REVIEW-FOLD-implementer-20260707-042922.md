## REVIEW-FOLD report - s6-core implementer - s6-core-impl-fold

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
SUBJECT: fold report for panel must-fix families 1-7; re-mint crash-window item remains HELD

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
- .relays/s6/s6-core-impl/REVIEW-FOLD-implementer-20260707-042922.md -> in
- .relays/s6/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF:
branch `s6-transport-impl` at `a8d04b4` (`s6: fold review findings for commit guard lifecycle sweeps and lock takeover`).

FINAL_GIT_STATUS_SHORT:
none - clean tree in `~/frank-s6-impl` after commit `a8d04b4`.

Fold summary:
- Item 1 folded: every non-empty `intake_id` fault/superseded append path now replays an existing outcome before appending; `TestCommitGuardBlocksSecondOutcome` covers the post-commit `AfterAccepted` panic path.
- Item 2 folded: boot admission now runs only while `RenderEnv.PreActive` is true, and active renders omit `charter_loaded` / `dispatch_status`; active boot-shaped ordinary submits accept.
- Item 3 folded: enum floor now scans real `record` delivery-state constants, real registry named enums, and real sealed store record bytes, with a planted-leak scanner proof.
- Item 4 folded: project view params now exercise a real served `frank` socket for `{"view":"audit"}` and `{"view":"roster"}` while preserving exactly `submit,project,read`.
- Item 5 folded: named negatives/legs added for accepted-record auth tag absence, stale non-submit credential behavior, multi-field boot smuggle detail, and lifecycle exactly-once activation refs.
- Item 6 folded: store root lock now records a store-visible `TAKEOVER` diagnostic after holder death; clean release clears diagnostic metadata; internal and `frank` process fixtures cover kill-9 takeover and symlink alias one-winner refusal. The s4-leftover class is treated as subsumed by the process-level alias/refusal and killed-holder takeover legs.
- Item 7 folded in this report: prior report wording is corrected below.

Red-first / evidence notes:
- RED item 1: `go test ./internal/engine -run TestCommitGuardBlocksSecondOutcome -count=1` failed with a rejected `rejected-intake-guard` outcome for `intake-guard`, proving the second-outcome fault path.
- GREEN item 1: same command passed after the loop guard change.
- RED item 2: `go test ./test/fixtures -run TestS6LiveMintBootActivationAndRoster -count=1` failed because the active form still rendered `charter_loaded`.
- GREEN item 2: same command passed after boot admission/render changes.
- RED item 6: `go test ./internal/store -run 'TestAcquireRootTakesOverAfterHolderDeathWithDiagnostic|TestAcquireRootSymlinkAliasHasOneWinner' -count=1` failed because `conductor.lock` had no `takeover` diagnostic after holder death.
- GREEN item 6: store lock subset and `test/fixtures` process lock subset passed after the lock diagnostic change.
- Items 3/4 are test-floor rewrites over real surfaces; no production change was needed. The enum scanner includes an in-test planted leak proof so the test fails if the scanner stops biting.
- Item 5 named negatives were partly already behavior-present; the new named tests passed after the item 1/2 production fixes where applicable and now pin the requested names/legs.

Verification:
- Narrow fold suite passed:
  `go test ./internal/engine -run 'TestCommitGuardBlocksSecondOutcome|TestTagNeverInAcceptedRecords|TestLoopDoesNotPersistAuthGenerationOnAcceptedRecord' -count=1`
  `go test ./internal/tables -run 'TestActivationRecordRefSetExactlyOnce|TestAlreadyActiveBootShapedRecordDoesNotMoveActivationEdge' -count=1`
  `go test ./internal/store -run 'TestAcquireRootExcludesSecondHolderAndReleases|TestAcquireRootTakesOverAfterHolderDeathWithDiagnostic|TestAcquireRootSymlinkAliasHasOneWinner' -count=1`
  `go test ./test/fixtures -run 'TestS6LiveMintBootActivationAndRoster|TestStaleNonSubmitRefusalIsNotLifecycleGating|TestS6SweepEnumFloorAndThreeVerbSurface|TestS6SweepProjectParamsCarryAuditAndRosterWithoutNewVerb|TestS6SecondFrankProcessRefusesRootLock|TestS6FrankRootLockTakeoverAfterKilledHolder|TestS6FrankRootLockSymlinkAliasRefusesSecondHolder' -count=1`
- Full battery passed:
  `go test -count=1 ./... && go vet ./... && go test -race ./internal/engine ./internal/tables ./internal/channel ./internal/store`
- Whitespace check passed:
  `git diff --check`

Corrected scope statement:
The prior completion report's broad wording that scope remained inside the dispatched SCOPE_DIFF rows is corrected: `internal/gc/gc_test.go` was in-plan but out-of-row, was escalated, and was absorbed by `RECONCILE-orchestrator-planner-20260707-040738.md` as the one test-only file; production `internal/gc/*.go` remains OUT-of-row. This fold did not touch `internal/gc/gc_test.go`, so no FOLD_SCOPE row for that file was required.

Held / not touched:
The re-mint crash-window production edit remains HELD exactly as directed. This fold did not touch `completeMissingSeatMintBindings`, `completeSeatMintBinding`, or `seat.Manager` shapes for that held item.

Remaining risk:
- The live operator step-exit and merge gates are still separate downstream gates.
- `FRANK_S6_DOGFOOD_STORE` archive replay remains env/operator supplied; no new live archive evidence is claimed here.
- No PR was opened by this fold report.

Next requested action:
s6-core.planner reviews the fold at `s6-transport-impl@a8d04b4`, then routes any remaining panel/fold or exit-gate work.
