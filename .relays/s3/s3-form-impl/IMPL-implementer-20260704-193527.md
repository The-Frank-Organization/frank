## IMPL - s3-form implementation blocker before source edits

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s3-form-impl-implementer-blocker
PARENT_DISPATCH_ID: s3-form-impl
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - planner/orchestrator fold needed before implementation can resume; merge stays the operator's separate S3-close gate
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s3-grill-s3-form
DESIGN_DOC_ID: s3-slice-3-design
DESIGN_LOCK_ID: s3-slice-3-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s3-slice-3-plan
IN_REPLY_TO: s3-form-impl/IMPL-planner-20260704-191116.md
FROM: s3-form.implementer
TO: s3-form.planner
CC: s3.orchestrator-planner, operator
SUBJECT: IMPL blocker - r4 plan still says root README.md is not in the plan file list while Task 12 and SCOPE_DIFF include README.md

Summary:
- I accepted the addressed implementation dispatch mechanically: it is under `.relays/s3`, addressed to `s3-form.implementer`, parented to the approving PLAN-REVIEW, and contains the literal implementation token in the dispatch relay.
- I created the isolated implementation worktree requested by the dispatch: `~/.config/superpowers/worktrees/frank/s3-form-impl` on branch `s3-form-impl` at `main@8750fcd`.
- Baseline verification in that worktree passed before any source edits.
- I stopped before Task 1 because the locked r4 plan has a material scope contradiction around root `README.md`. I made no source, test, sprint-plan, README, branch commit, PR, merge, or prototype changes.

## Blocker

The r4 plan carries the README fence ruling into the revision note, file list, and Task 12, but leaves the older global scope bullet active.

Evidence:
- The r4 revision note says root `README.md` enters the file list under the fence ruling (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:5`).
- The global constraints say every task inherits the blockers, then state: root `README.md` is NOT in this plan's file list and the fresh-store qualifier README line remains an ASK to the orchestrator (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:10-23`).
- The same plan's file-structure block includes `README.md` as a Task 12 file (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:62-86`).
- Task 12 explicitly requires the one bounded README fresh-store claim delta (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:183-188`).
- The orchestrator ruling does approve exactly that bounded README delta and requires the README row to enter the plan file list plus SCOPE_DIFF (`.relays/s3/s3-form-plan/SITREP-orchestrator-planner-20260704-190728.md:20-29`).
- The dispatch SCOPE_DIFF includes `README.md -> in` and says it is in-fence by the explicit ruling (`.relays/s3/s3-form-impl/IMPL-planner-20260704-191116.md:26-29`, `.relays/s3/s3-form-impl/IMPL-planner-20260704-191116.md:70-75`).

Why this blocks implementation:
- The plan says global constraints are inherited by every task and violations are blockers.
- Task 12 cannot be performed without touching `README.md`.
- Touching `README.md` would violate the still-active global bullet as written, even though later plan sections and the dispatch say the opposite.
- I should not choose which contradictory authority wins during implementation. The pair Planner/orchestrator needs to fold the stale global line so the locked plan and dispatch fence agree.

Required fold:
- Update the global scope bullet in the plan so it no longer says root `README.md` is absent from the plan file list or still awaiting the ASK.
- Preserve the narrow scope: exactly one bounded claim-honest README fresh-store delta under the ruling conditions, no broader README work.
- Reissue or confirm the implementation dispatch after that fold.

## Work performed

- Created isolated worktree branch `s3-form-impl` from `main@8750fcd`.
- Ran setup: `go mod download`.
- Ran baseline: `go test -count=1 ./...`.
- Reviewed the locked r4 plan, dispatch relay, and README fence ruling.
- Wrote this blocker relay and index row only.

PR: none.

Plan lock:
- `s3-slice-3-plan`, r4 at `main@8750fcd`.

Files changed:
- No tracked source/test/docs files changed.
- Gitignored relay artifact added: `.relays/s3/s3-form-impl/IMPL-implementer-20260704-193527.md`.
- Gitignored relay index row appended: `.relays/s3/INDEX.md`.

Acceptance criteria status:
- Not started. Implementation halted before Task 1 due to the scope contradiction above.

Boundary contract proof:
- No files from the implementation SCOPE_DIFF were edited.
- No file outside SCOPE_DIFF was edited.
- The worktree branch has no commits beyond `8750fcd`.
- No hard m-1 trigger was intentionally tripped; this is a plan/scope consistency blocker.

Tests/verification:
- Dispatch relay lint: `OK .relays/s3/s3-form-impl/IMPL-planner-20260704-191116.md`.
- Worktree setup: `go mod download` -> no module dependencies to download.
- Baseline from `~/.config/superpowers/worktrees/frank/s3-form-impl`: `go test -count=1 ./...` passed.
- Worktree status before this relay write: clean at `s3-form-impl@8750fcd`.
- Main relay checkout status before this relay write: clean at `main@8750fcd`.

Evidence levels:
- E1: file/line contradiction evidence above.
- E2: baseline test evidence above.

Out-of-scope preserved:
- No `../master` or `../extracted` writes.
- No config-change record work.
- No routing execution, observe/evidence Step-2, S4 consumer-field content, TUI/runtime, or locked-design amendment.
- No merge, push, PR, or S3-close claim.

Remaining risk:
- Low once folded: the required change appears to be a stale global-plan sentence, not a source-code concern.
- High if ignored: Task 12 would require knowingly violating a global constraint.

Next requested action:
- Planner should fold the stale global scope bullet or get an explicit orchestrator ruling that the later r4 README sections supersede it, then reissue/confirm implementation dispatch.

ACTIONS_GIT_REF: created isolated worktree branch `s3-form-impl` at `8750fcd`; no source/test/docs commits; wrote gitignored relay `.relays/s3/s3-form-impl/IMPL-implementer-20260704-193527.md` plus `.relays/s3/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: main relay checkout none - clean tree at `main@8750fcd`; implementation worktree none - clean tree at `s3-form-impl@8750fcd`
