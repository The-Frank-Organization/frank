## FOLD_SCOPE - s1-core.implementer accepts r2 narrow fold scope for RB1 RB2 RM1

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-r2-implementer-fold-scope
PARENT_DISPATCH_ID: s1-core-review-r2
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for this fold; merge remains a separate human gate
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
SUBJECT: r2 fold scope accepted for RB1 RB2 RM1 only; no edits before this artifact

Scope rule:
- Fold only RB1, RB2, and RM1 from `s1-core-review-r2/REVIEW-FOLD-planner-20260703-193332.md`.
- Do not re-touch findings marked verified folded unless the r2 scope row below is required for RB1, RB2, or RM1.
- Any needed file outside this table requires a deviation relay before editing.
- No merge, PR merge, live-verify, or exit-gate SITREP authority is claimed here.

FOLD_SCOPE:
- test/fixtures/ -> in
- test/replay/ -> in
- internal/recover/recover.go -> in
- internal/fsio/fsio.go -> in
- internal/store/store.go -> in
- internal/engine/loop.go -> in
- internal/channel/server.go -> in
- cmd/frank/main.go -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: no tracked edits before this artifact; current tracked tree clean at `main@7f66057` before r2 fold work begins.
FINAL_GIT_STATUS_SHORT: none - clean tracked tree (`git status --short --branch --untracked-files=all` reported `## main` before this gitignored relay/index write)
Next requested action: implement RB1, RB2, and RM1 only, then report with E2 verification.
