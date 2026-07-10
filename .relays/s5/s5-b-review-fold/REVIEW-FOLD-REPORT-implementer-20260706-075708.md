## REVIEW-FOLD report - s5-b MF-1b remainder folded

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
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-075026.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: MF-1b absorb-bounce remainder folded on s5-b-mechanisms
FOLD_SCOPE:
- internal/engine/submit.go -> in
- test/fixtures/s5_gate_raise_test.go -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: s5-b-mechanisms@b80e149; commit b80e149 s5-b: fold absorb reject gate raise; post-commit code worktree status clean

Summary:
- MF-1b fixed: reject cleanup now treats `gate_category_raised` with no `gate_category_pick` as a conductor-added absorb token and deletes `gate_category` before stripping raise headers.
- Fixture added: an absorb candidate with no seat-submitted `gate_category`, detector hit, and bad `PARENT_DISPATCH_ID` now proves the persisted rejected record has no `gate_category`, no `gate_category_raised`, and no `gate_category_pick`.

Files changed:
- internal/engine/submit.go
- test/fixtures/s5_gate_raise_test.go

TDD / focused evidence:
- RED observed before production edit: `go test -count=1 ./test/fixtures -run TestS5GateRaiseDoesNotPersistAddedCategoryOnRejectedAbsorb` failed with `rejected gate_category = "authz_security", want absent`.
- GREEN: `go test -count=1 ./test/fixtures -run TestS5GateRaiseDoesNotPersistAddedCategoryOnRejectedAbsorb` passed.
- GREEN: `go test -count=1 ./test/fixtures -run 'TestS5GateRaise'` passed.

Full verification:
- `git diff --check` passed with no output before commit.
- `go test -count=1 ./test/fixtures -run 'TestS5GateRaise' && go build ./... && go test -count=1 ./... && go vet ./...` passed before commit on the committed tree content.

Acceptance criteria status:
- MF-1b: satisfied.

Out-of-scope preserved:
- No MF-2/MF-3 changes.
- No optional items.
- No T7 work.
- No registered integration-gate ordering changes.
- No merge, push, PR, live activation, or integration-gate work.

Remaining risk:
- The registered integration ordering constraint remains with the integration gate, as the planner noted.

Next requested action:
- s5-b.planner targeted re-check of the MF-1b absorb-bounce leg and battery.
