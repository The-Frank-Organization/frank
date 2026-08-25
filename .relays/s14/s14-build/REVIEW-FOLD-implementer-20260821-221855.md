## REVIEW-FOLD — F3a folded at `4aea922e`; complete-request accounting, preserved response-first vector, and exact-commit E2 battery green

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
FILED_AT_LOCAL: 20260821-221855
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-215400.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: F3a complete-request write accounting folded at 4aea922e — RED/GREEN, count-100, race, and exact-commit E2 green; final re-verdict requested

Summary: the one authorized F3a finding is folded in one bounded two-file commit. `countingConn` now observes only bytes actually accepted by the plaintext HTTP/1.1 connection, parses the emitted request header to obtain `Content-Length`, and increments `request_write_completed` exactly once only after the complete body length has crossed. A request-level `WroteRequest` join plus first-response-byte ordering preserves the already-pinned response-first close vector only when no write failed before the provider response began. The counter vocabulary, no-replay pins, four existing cut-vector meanings, frozen carrier bytes, and all prior F1/F2 bytes are unchanged.

PR: none — this governed lane uses the existing `s14-m8-connector` branch; no push or PR action was authorized or performed.
Plan lock: successor fold grant `REVIEW-FOLD-planner-20260821-215400.md` at SHA-256 `da10e2c01da9a8aa5561300298bc8c25e8735ed332043eab78915305e6a71f5a`; implementer-owned pre-edit scope `FOLD_SCOPE-implementer-20260821-220057.md` at SHA-256 `1626c1dbcafbdbfd53b1a8f37e5f5d6e764ab077e3c39c66f75bf132c673a87d`.
Files changed: `internal/connector/transport/transport.go`, `internal/connector/transport/transport_test.go`.

## RED/GREEN trail

- **Named mutation:** restoring completion-on-first-successful-`Write` must make the new regression report completed 1 instead of the hand-derived literal 0.
- **RED at unchanged `da491c1c` production bytes:** `go test -mod=readonly -run '^TestMidBodyWriteFailureDoesNotCompleteRequest$' -count=1 ./internal/connector/transport` failed only at `Counters() = {DialAttempts:1 ConnectionsEstablished:1 RequestWriteStarted:1 RequestWriteCompleted:1}, want {DialAttempts:1 ConnectionsEstablished:1 RequestWriteStarted:1 RequestWriteCompleted:0}`. The fixture sends a 128-KiB body, allows the first connection write, fails the next write while the server consumes the body, and observes no retry.
- **Rejected request-trace-only hypothesis:** joining `httptrace.WroteRequest` made the new vector green but failed the preserved response-first `Retry-After` vector under count-100. Go's transport source established the root cause: `WroteRequest` fires before the write loop's outer buffered flush, so it cannot alone prove the final bytes crossed.
- **Final GREEN mechanism:** connection-level accepted-byte accounting supplies the normal full-write fact; `GotFirstResponseByte` ordering plus the request-write join isolates the already-pinned response-first close case from a pre-response or mid-body write failure. The new vector, all existing transport vectors, and the complete transport package then remained green through count-100.

## Verification at final bytes

- `go test -mod=readonly -count=100 ./internal/connector/transport` at `4aea922e` — exit 0, 31.639s; the new mid-body vector and all existing dial/connect/write-start/write-complete vectors remained exact.
- `go test -mod=readonly -race -count=20 ./internal/connector/control ./internal/connector/attempt ./internal/connector/transport ./internal/connector/service` at `4aea922e` — exit 0 (`control` 3.613s, `attempt` 1.469s, `transport` 8.212s, `service` 1.983s), no detector finding.
- Pre-commit `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; fixtures 190.965s, invariants 1.342s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go test -mod=readonly -p=1 -count=1 ./...` at `4aea922e` — exit 0; fixtures 196.969s, invariants 1.519s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go vet ./...`, `go build ./...`, `go mod tidy -diff`, and `git diff HEAD^ --check` — each exit 0; tidy emitted no diff and changed no module byte.
- Exact-commit `git status --short` — empty.

Acceptance criteria status: the multi-write mid-body cut is exactly `{1,1,1,0}`; the fresh-dial failure, first-write failure, response cuts, redirect/no-pool, response-first `Retry-After`, TTFB, and caller-context behaviors retain their established results; no retry was introduced. E2 target met.
Boundary contract proof: the fold changes only below-encoder HTTP/1.1 write accounting and its real TLS integration fixture. It retains the public `Counters` shape, transport error behavior, wire request bytes, provider credential path, and every E3 input. The full transport count-100 and repository batteries exercise the direct consumers.
Evidence levels: E1 exact two-file commit diff; E2 targeted RED/GREEN, count-100 package stress, race detector, two full serialized batteries, vet, build, and tidy-diff. No E3/E4 claim.
Out-of-scope preserved: zero control, attempt, executor, observe, credentials, policy, request/translate/freeze, stream/outcome schema, fixture corpus, dependency/module, generated artifact, merge, publication, deployment, or release byte changed. `git show --stat 4aea922e` lists exactly the two authorized transport files.
Remaining risk: planner re-verdict, serialized restack, downstream battery rerun, substantive re-review, and operator MERGE-GATE remain outstanding. No merge, deployment, or live-verification claim is made.

FOLD_SCOPE:
- frank/internal/connector/transport/transport.go -> in
- frank/internal/connector/transport/transport_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: `s14-m8-connector@4aea922e458dba50de2c21edc10bf989694b7d8c`; commit `4aea922e` (`s14 REVIEW-FOLD: count complete request writes`) cites the successor grant and changes exactly two in-scope source/test files; governance actions are the pre-edit FOLD_SCOPE, this report, and their append-only live-EOF s14 INDEX rows; no push, PR, merge, publication, deployment, or release
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@4aea922e458dba50de2c21edc10bf989694b7d8c`
Next requested action: `s14.planner` re-verdicts the accepted F1/F2/F3 plus F3a over exact commit `4aea922e`; on CLEAN, report upward for serialized restack, battery rerun, substantive re-review, and the separately authorized operator MERGE-GATE. No merge authority is requested or implied.
