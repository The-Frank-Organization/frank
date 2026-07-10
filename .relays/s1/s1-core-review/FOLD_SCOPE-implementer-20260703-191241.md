## FOLD_SCOPE - s1-core.implementer accepts bounded REVIEW-FOLD scope for B1-B7 and M1-M7

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-implementer-fold-scope
PARENT_DISPATCH_ID: s1-core-review
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for this fold; merge remains a separate human gate
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@9c1839e
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-review/REVIEW-FOLD-planner-20260703-190326.md
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
SUBJECT: fold scope accepted for required review findings; no edits before this artifact; any OUT row would block implementation

Scope rule:
- Fold blockers B1-B7 first, then must-fixes M1-M7.
- Optional findings O1-O7 may be folded only when touched by the required fix or when they remain within the rows below.
- No merge, PR merge, live-verify, or exit-gate SITREP authority is claimed here.
- If a required fix needs a file outside this table, stop and relay deviation instead of editing.

FOLD_SCOPE:
- internal/engine/loop.go -> in
- internal/engine/submit.go -> in
- internal/channel/server.go -> in
- internal/seat/binding.go -> in
- internal/fsio/fsio.go -> in
- internal/store/store.go -> in
- internal/store/projections.go -> in
- internal/intake/journal.go -> in
- internal/gate/derived.go -> in
- internal/lineage/lineage.go -> in
- internal/fieldspec/fieldspec.go -> in
- internal/recover/recover.go -> in
- cmd/frank/main.go -> in
- test/fixtures/ -> in
- test/replay/classmap.go -> in
- test/replay/replay_test.go -> in
- test/replay/report.md -> in
- test/seatproc/testseat.go -> in
- README.md -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: no tracked edits before this artifact; current tracked tree clean at `main@9c1839e` before fold work begins.
FINAL_GIT_STATUS_SHORT: none - clean tracked tree (`git status --short --branch --untracked-files=all` reported `## main` before this gitignored relay/index write)
Next requested action: implement the bounded fold and report with tests/evidence.
