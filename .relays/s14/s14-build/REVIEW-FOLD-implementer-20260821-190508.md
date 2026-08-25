## REVIEW-FOLD — F1+F2+F3 folded at `da491c1c`; all RED/GREEN legs and two full serialized batteries green; returning for planner re-verdict

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair re-verdict is next; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-190508
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-184809.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: F1 bounded epoch hold, F2 cancel-before-classify, F3 below-encoder completion count folded at da491c1c — E2 green, clean tree, re-verdict requested

Summary: all three end-review findings are folded in one bounded commit. F1 now holds an above-current DATA-P epoch for the compiled `EpochQueryHoldDeadline`, wakes on the m-10 epoch-change signal, and re-evaluates once. F2 cancels the attempt context before classifying `invoked`, closing the successful-gate/pre-transport window. F3 moves write-completion counting from the asynchronous `httptrace.WroteRequest` callback to the instrumented connection wrapper's first full successful write return, exactly once per fresh connection; the counter vocabulary, no-replay pins, and fixture cut vectors are unchanged.

PR: none — this governed lane uses the existing `s14-m8-connector` branch; no push or PR action was authorized or performed.
Plan lock: successor fold grant `REVIEW-FOLD-planner-20260821-184809.md` at SHA-256 `c6f5b2992ef3e5e2c1819b1855a5532640b6312420f1d3211cdbbaa4b1d8b7f0`; successor pre-edit scope `FOLD_SCOPE-implementer-20260821-184953.md` at SHA-256 `e43ed11de27b267d7ca2df3c665c7821a7dbd6eb5e9e174dfa5ec42fa21a0ce4`.
Files changed: `internal/connector/control/control.go`, `internal/connector/control/control_test.go`, `internal/connector/attempt/attempt.go`, `internal/connector/attempt/attempt_test.go`, `internal/connector/transport/transport.go`. The authorized `transport_test.go` path required no byte change because its existing counter assertion supplied the F3 regression RED.

## RED/GREEN trail

- **F1 RED:** `TestFenceDataEpochReevaluatesUpdateDuringBoundedHold` failed with `re-evaluated fence = "EPOCH_AHEAD", <nil>; want "allowed"` after the query write and an in-window real `HandleControl` update. **GREEN:** the same test and the existing unresolved `EPOCH_AHEAD` leg pass; the signal captures updates before, during, or after query emission and the hold remains bounded at 100ms.
- **F2 RED:** `TestCancelNeverClassifiesPreTransportWhenInvocationGateSucceeds` failed on iteration 0 with `successful invocation gate recorded as pre_transport`. **GREEN:** cancel-first classification passes that deterministic boundary interleaving plus the existing pre/post-invocation and already-cancelled gate tests.
- **F3 RED:** the first full battery failed only `TestRetryAfterIsRecordedAndNeverActedOn` at `{DialAttempts:1 ConnectionsEstablished:1 RequestWriteStarted:1 RequestWriteCompleted:0}`; isolated `-count=100` reproduced twice on unchanged transport bytes. The first allowed hypothesis—join `WroteRequest`—passed the isolated single-test loop but failed the full package stress because a successful early 429 response can accompany trace error `use of closed network connection`. Diagnostic evidence established that the handler returns without consuming the body after the connection wrapper has already completed a full write. The rejected join bytes and diagnostics were removed before commit. **GREEN:** the boundary-level `countingConn.Write` mechanism makes the same isolated test and the complete transport package pass at `-count=100` while the injected post-connect write failure remains `{1,1,1,0}`.

## Verification at final bytes

- `go test -mod=readonly -count=100 ./internal/connector/transport -run TestRetryAfterIsRecordedAndNeverActedOn` — exit 0.
- `go test -mod=readonly -count=100 ./internal/connector/transport` — exit 0, 31.339s; all dial/connect/write-start/write-complete cut vectors exact.
- `go test -mod=readonly -race -count=20 ./internal/connector/control ./internal/connector/attempt ./internal/connector/transport ./internal/connector/service` — exit 0 (`control` 3.911s, `attempt` 1.358s, `transport` 8.253s, `service` 2.111s).
- Pre-commit `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; fixtures 202.158s.
- Final-commit `go test -mod=readonly -p=1 -count=1 ./...` at `da491c1c` — exit 0; fixtures 200.605s, invariants 1.409s, replay/dogfood/zeroloss green.
- Final-commit `go vet ./...`, `go build ./...`, `go mod tidy -diff`, and `git diff --check` — each exit 0; the tidy tripwire emitted no diff and changed no module byte.

Acceptance criteria status: F1 update-inside-hold and unresolved-ahead legs pass; F2 successful invocation cannot coexist with a `pre_transport` fact; F3 counter truth is observed at the specified below-encoder boundary and all original four cut vectors remain exact. E2 target met.
Boundary contract proof: F1 changes only CTRL-C epoch synchronization and emits no attempt result/counter; F2 changes only cancellation ordering and retains existing event/outcome schemas; F3 retains `Counters` shape and transport retry/error behavior, touching neither carrier bytes nor E3 inputs. Repeated package and full-repository tests exercise each consumer boundary.
Evidence levels: E1 exact five-file commit diff; E2 targeted RED/GREEN, count-100 stress, race detector, two full batteries, vet, build, and tidy-diff. No E3/E4 claim.
Out-of-scope preserved: zero dependency/module, executor, observe, credentials, policy, request/translate/freeze, stream/outcome schema, fixture, generated artifact, merge, publication, or release byte changed. `git show --stat da491c1c` lists exactly five in-scope files.
Remaining risk: serialized restack, downstream battery rerun, substantive re-review, and operator MERGE-GATE remain outstanding. No merge, deployment, or live-verification claim is made.

FOLD_SCOPE:
- frank/internal/connector/control/control.go -> in
- frank/internal/connector/control/control_test.go -> in
- frank/internal/connector/attempt/attempt.go -> in
- frank/internal/connector/attempt/attempt_test.go -> in
- frank/internal/connector/transport/transport.go -> in
- frank/internal/connector/transport/transport_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: `s14-m8-connector@da491c1c458a9c22c322c2e08b1a4f81ba029b37`; commit `da491c1c` (`s14 REVIEW-FOLD: close F1-F3 end-review findings`) changes exactly five in-scope source/test files; governance actions are this report plus one append-only live-EOF s14 INDEX row; no push, PR, merge, publication, deployment, or release
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@da491c1c458a9c22c322c2e08b1a4f81ba029b37`
Next requested action: `s14.planner` re-verdicts F1+F2+F3 over exact commit `da491c1c`; on CLEAN, report upward for serialized restack, battery rerun, substantive re-review, and the separately authorized operator MERGE-GATE. No merge authority is requested or implied.
