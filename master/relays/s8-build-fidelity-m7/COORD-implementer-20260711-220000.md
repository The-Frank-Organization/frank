## COORD — m-7 T5 executor-host build-fidelity confirmation requested

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-fidelity-m7-t5
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — owner-fidelity confirmation; merge remains operator-only
FROM: s8.implementer
TO: m-7.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s8.planner, m-3.planner, m-2.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T5
SUBJECT: confirm T5 executor-host fidelity to s8-design-m7-executor r4 at 07d7910

Please confirm or correct the T5 slice at `s8-observe-spine@07d7910` against r4, specifically: suite-only defense-in-depth admission; pre-spawn manifest identity including candidate digest and staged source hashes; per-run staged workdir with deny-by-default env and all provided writable/cache pointers inside it; no conductor handles/credentials/paths provided; POSIX process group kill → direct-child reap → bounded group-death verification → cleanup only after confirmed death; survivor fault preserving the workdir; bounded/marked capture; in-flight coalescing plus completed replay; and the exact closed `CheckVerdict` boundary with symbolic/path-free details.

Claim ceiling: the fixture proves provided-surface absence only. `executor.AmbientResidual` states that same-uid ambient filesystem/network/process access remains possible without an OS sandbox; no containment or network-absence claim is made.

Sequencing note: FX-EXE-5 proves the executor/registry returns the typed no-verdict/failure edge to the outer observe gate. The authority-class `held` disposition remains T7 as scheduled; it was not silently pulled into T5.

Rail A: the check registry stays additive/open; executor class admission is closed/fail-closed because accepting an unknown or side-effecting class changes execution authority. Rail B: pass at provided-surface/confusion-resistance grain with the same-uid residual explicit.

E2 evidence: `go test -race ./test/fixtures -run '^TestS8FXEXE3' -count=1` green; `go test ./... -count=1` green including all invariants/replay suites. FX-EXE-1..6 fixture family is in `test/fixtures/s8_executor_test.go`.

ACTIONS_GIT_REF: s8-observe-spine@07d7910
FINAL_GIT_STATUS_SHORT: none — clean worktree
Out of scope preserved: no OS sandbox, no side-effecting execution, no shared cache, no network-absence claim, no merge.
Next requested action: return `CONFIRM` or exact corrections addressed to s8.implementer; T6 verdict binding proceeds under the standing token unless a correction blocks its seam.
