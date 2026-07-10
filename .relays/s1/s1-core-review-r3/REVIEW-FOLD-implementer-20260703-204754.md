## REVIEW-FOLD - s1-core.implementer r3 fold report for C7 partial-state fixture

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-r3-implementer-report
PARENT_DISPATCH_ID: s1-core-review-r3
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for this fold; merge and exit-gate routing remain separate gates
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@0178ab0
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-review-r3/REVIEW-FOLD-planner-20260703-202251.md
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
SUBJECT: r3 deterministic C7 partial-state fixture and canonical outbox dedupe folded at f0dcb850f3e4

Scope honored:
- Folded only the deterministic C7 partial-state fixture plus the optional canonical-scan robustness ride-along authorized by `s1-core-review-r3/REVIEW-FOLD-planner-20260703-202251.md`.
- Touched no files outside the two authorized rows.
- No merge, PR merge, live-verify, or exit-gate SITREP authority is claimed here.

Fold summary:
- Added a deterministic partial-state fixture where a gate-bearing accepted record and its park record exist without an outbox item; `gate.Complete` must create exactly one outbox item and not duplicate the park record.
- Added the mirror leg where the canonical outbox record exists without the park record; `gate.Complete` must create exactly one park record and not duplicate the outbox record.
- Chose the canonical-scan ride-along: `completeOutbox` now dedupes by scanning canonical `outbox-<item_id>` records, symmetric with `completePark`, instead of depending on the projection file having been rebuilt first.

TDD evidence:
- Before the production change, `go test ./internal/gate -run 'TestCompleteConvergesWhenGateParkExistsWithoutOutbox|TestCompleteConvergesWhenGateOutboxRecordExistsWithoutPark' -count=1 -v` failed on `TestCompleteConvergesWhenGateOutboxRecordExistsWithoutPark` with `Complete: record already exists`.
- After the canonical-scan change, the same command exited 0 and both partial-state tests passed.

FOLD_SCOPE:
- internal/gate/derived_test.go -> in
- internal/gate/derived.go -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: f0dcb850f3e4 (`s1 REVIEW-FOLD r3: pin gate partial-state convergence`)

Verification evidence, freshly run after commit f0dcb850f3e4:
- `go test ./...` -> exit 0; all packages pass.
- `go test -race ./...` -> exit 0; all packages pass under race instrumentation.
- `go vet ./...` -> exit 0; no diagnostics.
- `golangci-lint run ./...` -> exit 0; output `0 issues.`
- `git diff --check` -> exit 0; no whitespace diagnostics.
- `git status --short --branch --untracked-files=all` -> `## main` before this gitignored relay/index write.

FINAL_GIT_STATUS_SHORT: clean tracked tree before this gitignored relay/index write (`## main`)
Next requested action: planner verifies r3 spec-match and green status, then routes the gate-close report as authorized by the r3 relay.
