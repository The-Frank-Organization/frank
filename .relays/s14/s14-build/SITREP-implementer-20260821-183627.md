## SITREP — F1+F2 are RED/GREEN in scope, but the mandatory full E2 battery exposed a reproducible pre-existing transport counter race outside FOLD_SCOPE; fold stopped uncommitted before any transport edit

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s14-build-end-review-green-blocker-1
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — master must disposition the out-of-scope transport counter race; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-183627
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-182628.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: F1/F2 scoped RED/GREEN complete but fold commit blocked — full battery reproducibly catches an unchanged transport request-write-completion truth race outside the four-file grant

## In-scope fold state

- The mandatory pre-edit barrier is filed at `FOLD_SCOPE-implementer-20260821-182801.md`, SHA-256 `d3f7ceee61ad4d6f875d82842eea4fbaa242433ae01f8e5e93daec6f2d51f98b`; exact relay and s14 INDEX lint passed before the first test edit.
- **F1 RED observed:** `TestFenceDataEpochReevaluatesUpdateDuringBoundedHold` returned `EPOCH_AHEAD` instead of `allowed` after the query was on the wire and an in-window `epoch_update` arrived through the real `HandleControl` path.
- **F1 GREEN observed:** `FenceDataEpoch` now waits on the compiled `EpochQueryHoldDeadline`, wakes on the session epoch-change signal, and re-evaluates once. The in-window update test and the existing unresolved `EPOCH_AHEAD` test both pass.
- **F2 RED observed:** the deterministic cancellation-boundary interleaving produced a successful `TryMarkInvoked` with a recorded `pre_transport` cancel point on iteration zero.
- **F2 GREEN observed:** `Manager.Cancel` now cancels the context before reading `invoked`; the invariant test, the already-cancelled gate refusal, and both pre/post-invocation cancellation tests pass.
- Repeated race verification is green: `go test -mod=readonly -race -count=20 ./internal/connector/control ./internal/connector/attempt` exited 0 (`control` 3.735s; `attempt` 1.442s).

All current source edits remain exactly the four in-scope paths, unstaged and uncommitted:

```text
 M internal/connector/attempt/attempt.go
 M internal/connector/attempt/attempt_test.go
 M internal/connector/control/control.go
 M internal/connector/control/control_test.go
```

## Full-battery blocker and root-cause evidence

The required `go test -mod=readonly -p=1 -count=1 ./...` exited 1. Every package before and after the failure completed; `test/fixtures` was green in 192.324s. The sole failure was:

```text
--- FAIL: TestRetryAfterIsRecordedAndNeverActedOn (0.01s)
    transport_test.go:221: Counters() = {DialAttempts:1 ConnectionsEstablished:1 RequestWriteStarted:1 RequestWriteCompleted:0}, want {DialAttempts:1 ConnectionsEstablished:1 RequestWriteStarted:1 RequestWriteCompleted:1}
FAIL github.com/jackli/frank/internal/connector/transport
```

Systematic isolation establishes that this is not an F1/F2 regression:

- `git diff --exit-code -- internal/connector/transport` exits 0; no transport byte differs from head `316f97a6025111cd3aa65e2841b4465010694bb8`.
- `go list -deps ./internal/connector/transport` contains neither `internal/connector/control` nor `internal/connector/attempt`; the isolated package does not compile or execute either changed package.
- `go test -mod=readonly -count=100 ./internal/connector/transport -run TestRetryAfterIsRecordedAndNeverActedOn` reproduces the same mismatch twice and exits 1.
- Root cause is the existing asynchronous observation seam: `RequestWriteCompleted` increments only in the `httptrace.WroteRequest` callback, while `SendGated` may return the received response before the HTTP/1 write goroutine executes that callback. Go's own `net/http` source documents the response-before-write-goroutine-completion case around `persistConn.wroteRequest`; this client disables keep-alives and does not itself join the callback before exposing `Counters()`. The test therefore intermittently observes a real zero-completion window immediately after a successful response.

That mechanism and its test live under `frank/internal/connector/transport/**`, outside the active FOLD_SCOPE. Making the mandatory full battery reliably green requires a transport-owner change or an explicit upstream ruling on how this pre-existing counter-truth race is carried; simply rerunning until green would conceal reproducible red evidence.

## Stop and requested disposition

No out-of-scope source/test edit, staging action, commit, push, PR, merge, publication, or release action was performed. The four valid F1/F2 edits are preserved uncommitted for continuation after the scope disposition.

SCOPE_DIFF:
- frank/internal/connector/control/control.go -> in
- frank/internal/connector/control/control_test.go -> in
- frank/internal/connector/attempt/attempt.go -> in
- frank/internal/connector/attempt/attempt_test.go -> in
- frank/internal/connector/transport/transport.go -> OUT
- frank/internal/connector/transport/transport_test.go -> OUT
SCOPE_DIFF_RESULT: deviation-present

ACTIONS_GIT_REF: in-scope uncommitted F1/F2 edits only at `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`, plus this governance SITREP and one append-only live-EOF s14 INDEX row; no transport edit, stage, commit, push, PR, merge, publication, or release
FINAL_GIT_STATUS_SHORT:
 M internal/connector/attempt/attempt.go
 M internal/connector/attempt/attempt_test.go
 M internal/connector/control/control.go
 M internal/connector/control/control_test.go
Next requested action: `s14.planner` routes the reproducible out-of-scope transport counter-truth race to master. Master either grants a separately reviewed transport repair and returns a corrected fold scope, or supplies an explicit evidence disposition that preserves the red run and defines the acceptable rerun gate. Until then s14.implementer holds the four in-scope edits uncommitted. No merge authority is requested or implied.
