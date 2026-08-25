## FOLD_SCOPE — R14 F7 self-contained missing-module naming probe

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-fold-scope-7
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded test-only successor under rows 12/13/14; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260822-024149
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260822-024017.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer, m-3.planner, m-7.planner
SUBJECT: pre-edit scope — R14 F7 replaces the consuming-root-dependent naming probe with one synthetic closure-invariant probe; all-in

FOLD_SCOPE:
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

The source worktree is clean at `s14-m8-connector@ebc9fbe4d8eadfb81d8d8e9c9fa7ee0b58c30d86`. The RED evidence is the old test's closure dependency: drive the unchanged probe against an s13-shaped consuming closure, or an equivalent isolated temporary construction, and record that module-graph traversal names a different closure module before `golang.org/x/text`. Then change only the named test so its own synthetic `go.mod`, import, and expected diagnostic all name a fabricated absent module path under `GOPROXY=off`.

The production change that would make the successor test fail is a regression in `Executor.Spawn` that stops propagating the subprocess's real offline module diagnostic or omits its private retained tail; the expected fabricated path is a hand-written literal independent of executor helpers. The accepted `executor.go`, s8 fixture, R8 fixture, cache shape, offline flags, and R10 bare-token law remain byte-exact.

At final bytes, the fold report will restate a closure statement for every test in `executor_test.go`: each depends on nothing from the consuming root, or only on the root closure in a closure-agnostic way. Any need outside the two rows stops before edit and routes a deviation.

ACTIONS_GIT_REF: governance-only pre-edit barrier — this FOLD_SCOPE relay plus one append-only live-EOF s14 INDEX row; source remains clean at `s14-m8-connector@ebc9fbe4d8eadfb81d8d8e9c9fa7ee0b58c30d86`; no source or test edit yet
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@ebc9fbe4d8eadfb81d8d8e9c9fa7ee0b58c30d86`
Next requested action: s14.implementer records the closure-dependent RED, authors the one-test synthetic successor, runs the exact-commit E2 battery, and returns REVIEW-FOLD for delta re-verdict. No merge authority is requested or implied.
