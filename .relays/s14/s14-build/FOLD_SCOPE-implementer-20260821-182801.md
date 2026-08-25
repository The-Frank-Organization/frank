## FOLD_SCOPE — s14 end-review F1+F2 only, four code/test paths plus s14 fold evidence

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-fold-scope-1
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair-internal bounded fold; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-182801
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-182628.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: pre-edit scope for F1 EPOCH_AHEAD hold/re-evaluate and F2 cancel-before-classify — all named paths in, no deviation

FOLD_SCOPE:
- frank/internal/connector/control/control.go -> in
- frank/internal/connector/control/control_test.go -> in
- frank/internal/connector/attempt/attempt.go -> in
- frank/internal/connector/attempt/attempt_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

The fold is limited to the two findings in `REVIEW-FOLD-planner-20260821-182628.md`. F1 will first add a behavioral negative proving an in-window epoch update changes the re-evaluated disposition while an unresolved query remains `EPOCH_AHEAD`; only then will the minimal bounded hold and one-time re-evaluation land. F2 will first add a race-driven negative proving `pre_transport` cannot coexist with a successful invocation gate; only then will cancellation move before classification. Each finding receives targeted RED/GREEN evidence before the full E2 battery.

Any discovered need outside the five rows above stops the fold before that edit and routes a deviation. No locked design, dependency surface, provider/credential boundary, executor/observe byte, merge gate, or release surface is in scope.

ACTIONS_GIT_REF: governance-only pre-edit barrier — this FOLD_SCOPE relay plus one append-only live-EOF s14 INDEX row; source remains untouched and clean at `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`
Next requested action: s14.implementer now performs exactly the scoped RED-first F1 and F2 folds, targeted verification after each, the full E2 battery, bounded commit(s), and a REVIEW-FOLD report. No merge authority is requested or implied.
