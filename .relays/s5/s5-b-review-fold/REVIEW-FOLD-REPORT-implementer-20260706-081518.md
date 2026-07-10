## REVIEW-FOLD report - s5-b MF-1c other-pick remainder folded

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
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-081048.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: MF-1c other-pick bounce remainder folded on s5-b-mechanisms
FOLD_SCOPE:
- internal/fieldspec/validate.go -> in
- test/fixtures/s5_gate_raise_test.go -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: s5-b-mechanisms@2c3ccf5; commit 2c3ccf5 s5-b: fold other pick reject provenance; post-commit code worktree status clean

Summary:
- MF-1c fixed: the `other` pick raise branch now stamps `gate_category_pick: other`, making raised-without-pick unique to conductor-added absorbs.
- Accepted `other` records now carry `gate_category_raised: yes` plus `gate_category_pick: other`.
- Rejected `other` bounces now restore the seat-submitted `gate_category: other` while stripping computed raise headers.

Files changed:
- internal/fieldspec/validate.go
- test/fixtures/s5_gate_raise_test.go

TDD / focused evidence:
- RED observed before production edit: `go test -count=1 ./test/fixtures -run 'TestS5GateRaiseOtherUsesYesByteAndRegistryToken|TestS5GateRaiseDoesNotDropOtherPickOnRejectedRecord'` failed with `gate_category_pick = "", want other` and `rejected gate_category = "", want other`.
- GREEN: `go test -count=1 ./test/fixtures -run 'TestS5GateRaiseOtherUsesYesByteAndRegistryToken|TestS5GateRaiseDoesNotDropOtherPickOnRejectedRecord'` passed.
- GREEN: `go test -count=1 ./test/fixtures -run 'TestS5GateRaise'` passed.

Full verification:
- `git diff --check` passed with no output before commit.
- `go build ./... && go test -count=1 ./... && go vet ./...` passed before commit on the committed tree content.

Acceptance criteria status:
- MF-1c: satisfied.

Out-of-scope preserved:
- No MF-2/MF-3 changes.
- No optional items.
- No T7 work.
- No registered integration-gate ordering changes.
- No merge, push, PR, live activation, or integration-gate work.

Remaining risk:
- The registered integration ordering constraint remains with the integration gate.

Next requested action:
- s5-b.planner targeted re-check of the MF-1c other-pick accept and bounce legs, then pair review-complete report if green.
