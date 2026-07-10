## FOLD_SCOPE - s1-core.implementer accepts r3 narrow fold scope for C7 partial-state fixture and canonical outbox dedupe

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-r3-implementer-fold-scope
PARENT_DISPATCH_ID: s1-core-review-r3
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for this fold; merge remains a separate human gate
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
SUBJECT: r3 fold scope accepted for deterministic C7 partial-state fixture plus canonical-scan ride-along; no edits before this artifact

Scope rule:
- Fold only the deterministic C7 partial-state fixture from `s1-core-review-r3/REVIEW-FOLD-planner-20260703-202251.md`.
- Include the optional canonical-record-scan ride-along for `completeOutbox` because it removes the projection rebuild ordering dependency.
- Any needed file outside this table requires a deviation relay before editing.
- No merge, PR merge, live-verify, or exit-gate SITREP authority is claimed here.

FOLD_SCOPE:
- internal/gate/derived_test.go -> in
- internal/gate/derived.go -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: no tracked edits before this artifact; current tracked tree clean at `main@0178ab0` before r3 fold work begins.
FINAL_GIT_STATUS_SHORT: none - clean tracked tree (`git status --short --branch --untracked-files=all` reported `## main` before this gitignored relay/index write)
Next requested action: implement the C7 partial-state fixture and canonical outbox dedupe, then report with E2 verification.
