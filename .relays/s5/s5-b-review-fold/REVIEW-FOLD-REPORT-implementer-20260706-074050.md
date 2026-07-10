## REVIEW-FOLD report - s5-b panel must-fixes folded

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-b-review-fold
PARENT_DISPATCH_ID: s5-b-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-mechanisms
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-073232.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: MF-1/MF-2/MF-3 folded on s5-b-mechanisms
FOLD_SCOPE:
- internal/engine/submit.go -> in
- test/fixtures/s5_gate_raise_test.go -> in
- test/replay/zeroloss/zeroloss_test.go -> in
- test/fixtures/s5_iph_test.go -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: s5-b-mechanisms@a11e4bb; commit a11e4bb s5-b: fold panel must-fixes; post-commit worktree status clean

Summary:
- MF-1 fixed: rejected records now strip raise-derived headers at submit reject boundaries; if `gate_category_pick` exists, the rejected record restores the original pick before deleting `gate_category_raised` and `gate_category_pick`.
- MF-1 fixture added: detector-hit plus bad `PARENT_DISPATCH_ID` now proves the persisted rejected record has the original B pick and no raise proof headers.
- MF-2 hardened: the replay canonical-wins test now asserts the corrupted outbox projection bytes remain on disk while the replayed view remains canonical from the record.
- MF-3 covered: egress Drain sentinel strings are now part of the S5 I-PH output sweep.

Files changed:
- internal/engine/submit.go
- test/fixtures/s5_gate_raise_test.go
- test/replay/zeroloss/zeroloss_test.go
- test/fixtures/s5_iph_test.go

TDD / focused evidence:
- RED observed before production edit: `go test -count=1 ./test/fixtures -run TestS5GateRaiseDoesNotPersistStampOnRejectedRecord` failed with `rejected gate_category = "authz_security", want original B pick`.
- GREEN: `go test -count=1 ./test/fixtures -run TestS5GateRaiseDoesNotPersistStampOnRejectedRecord` passed.
- GREEN: `go test -count=1 ./test/fixtures -run 'TestS5GateRaise'` passed.
- GREEN: `go test -count=1 ./test/replay/zeroloss -run TestReplayConstructedStoreZeroLossIdentityAndCanonicalWins` passed.
- GREEN: `go test -count=1 ./test/fixtures -run TestS5IPH` passed.
- GREEN combined: `go test -count=1 ./test/fixtures -run 'TestS5GateRaise|TestS5IPH' && go test -count=1 ./test/replay/zeroloss -run TestReplayConstructedStoreZeroLossIdentityAndCanonicalWins` passed.

Full verification:
- `git diff --check` passed with no output.
- `go build ./... && go test -count=1 ./... && go vet ./...` passed.

Acceptance criteria status:
- MF-1: satisfied.
- MF-2: satisfied.
- MF-3: satisfied.

Out-of-scope preserved:
- No optional items folded.
- No T7 work.
- No merge, push, PR, live activation, or integration-gate work.

Remaining risk:
- T7 carry and integration-gate ordering note remain with their existing owners/gates.

Next requested action:
- s5-b.planner targeted re-check of the three folded must-fixes only.
