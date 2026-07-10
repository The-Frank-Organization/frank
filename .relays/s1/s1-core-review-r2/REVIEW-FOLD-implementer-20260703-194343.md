## REVIEW-FOLD - s1-core.implementer r2 fold report for RB1 RB2 RM1

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-r2-implementer-report
PARENT_DISPATCH_ID: s1-core-review-r2
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for this fold; merge and exit-gate SITREP remain separate gates
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@7f66057
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-review-r2/REVIEW-FOLD-planner-20260703-193332.md
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
SUBJECT: r2 RB1 RB2 RM1 folded at 80c5df5e626a with fresh E2 verification

Scope honored:
- Folded only RB1, RB2, and RM1 from `s1-core-review-r2/REVIEW-FOLD-planner-20260703-193332.md`.
- Did not edit the round-2 verified-folded items outside the authorized rows.
- No merge, PR merge, live-verify, or exit-gate SITREP authority is claimed here.

Fold summary:
- RB1: added real child-process SIGKILL crash fixtures around store/submit/park/outbox mutation paths, recovery, staging cleanup, duplicate relay detection, and canonical rename counting through fsio instrumentation.
- RB1/F9: added a real crash intake fixture proving only outcome-less intake records are returned in arrival order after restart.
- RB2: added pending-delivery derivation from mailboxes plus authenticated reconnect nudge queuing; the binary now queues a recovery nudge for existing mailbox deliveries at startup.
- RM1: added assembled-server coverage for `tools/list`, served tool descriptions, and captured push frames against path-family leakage.

FOLD_SCOPE:
- cmd/frank/main.go -> in
- internal/channel/server.go -> in
- internal/fsio/fsio.go -> in
- internal/store/store.go -> in
- test/fixtures/f11_test.go -> in
- test/fixtures/iph_test.go -> in
- test/fixtures/main_assembly_test.go -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: 80c5df5e626a (`s1 REVIEW-FOLD r2: crash matrix and recovery wake residues`)

Verification evidence, freshly run after commit 80c5df5e626a:
- `go test ./...` -> exit 0; all packages pass.
- `go test -race ./...` -> exit 0; all packages pass under race instrumentation.
- `go vet ./...` -> exit 0; no diagnostics.
- `golangci-lint run ./...` -> exit 0; output `0 issues.`
- `git diff --check` -> exit 0; no whitespace diagnostics.
- `git status --short --branch --untracked-files=all` -> `## main` before this gitignored relay/index write.

FINAL_GIT_STATUS_SHORT: clean tracked tree before this gitignored relay/index write (`## main`)
Next requested action: planner re-checks RB1, RB2, and RM1 only; no merge or exit-gate authority is included in this relay.
