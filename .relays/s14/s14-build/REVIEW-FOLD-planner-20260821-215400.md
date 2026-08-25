## REVIEW-FOLD — the re-verdict over fold commit `da491c1c458a9c22c322c2e08b1a4f81ba029b37`: **F1 + F2 + F3 ACCEPTED as folded** (verified at final bytes by this seat: F1's bounded `EpochQueryHoldDeadline` hold + signal wake + ONE total re-evaluation [below⇒STALE / equal⇒allowed / still-ahead⇒EPOCH_AHEAD rejected-retriable]; F2's cancel-before-classify on disk; the commit scope-exact at five files; my OWN independent full battery + vet GREEN at `da491c1c`) — **with ONE remaining narrow finding, F3a, in the same transport scope; this successor grant carries it:** the landed completion mechanism (`countingConn.Write`, `completeOnce` on the FIRST fully-flushed Write call) marks the WHOLE request complete when only its first chunk crossed — for a multi-Write request (an `LLMRequest` body routinely exceeds one write chunk) failing MID-BODY, the counter reads `{1,1,1,1}` where the addendum §6.2 partial-write vector REQUIRES `{…,0}`. Honestly scoped: the evidence chain cannot be corrupted by this (a false P2a=1 still dies at m-3's P2b complete-capture requirement — no wrong E3 pass is manufacturable), and every EXISTING vector stays exact (their fixtures are single-chunk); but the addendum vector is a locked semantic the §2a matrix claims LOCAL, and a claimed-green leg may not carry a known-false edge.

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a pair-internal fold inside the charter fence; the operator MERGE-GATE stays terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-215400
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-implementer-20260821-190508.md
FROM: s14.planner
TO: s14.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: F1/F2/F3 accepted at da491c1c; one remaining finding F3a — completion must mean the COMPLETE encoded request crossed; fold, battery, report

FOLD_SCOPE:
- frank/internal/connector/transport/transport.go -> in
- frank/internal/connector/transport/transport_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

Scope note: any discovered need outside these rows is a STOP-before-edit deviation escalation; any r12-pinned-semantic touch (counter vocabulary, no-replay pins, the four existing cut vectors' meanings) is a STOP regardless of path.

## §1 — F3a (the one finding)

**`transport.go` (`countingConn.Write` + counter exposure):** `request_write_completed` must be true IFF the COMPLETE encoded request crossed the instrumented boundary (the addendum P2a semantic: "a complete request was written"; §6.2: partial-write ⇒ completed 0). The landed `completeOnce`-on-first-successful-Write marks completion at the first chunk. Fold inside these semantics — mechanism yours; non-binding candidates: (i) join the err-free `httptrace.WroteRequest` signal (the true full-write fact) before counters/response become observable, deciding the early-response/closed-conn case from the conn-level no-failed-write fact your diagnostics already mapped; (ii) any other shape that keeps the four existing vectors exact and makes the mid-body cut read 0. RED first: a multi-chunk request (body > one write chunk) whose connection fails mid-body ⇒ `{DialAttempts:1, ConnectionsEstablished:1, RequestWriteStarted:1, RequestWriteCompleted:0}`; GREEN: that vector + the four existing vectors + the count-100 loops + the full battery.

## §2 — The gate on your report

One bounded commit citing this grant; targeted RED/GREEN + the count-100 stress; then the full E2 battery (`go test -mod=readonly -p=1 -count=1 ./...` + vet + tidy tripwire) green at the commit; literal `git status --short`. On the report I re-verdict; CLEAN ⇒ the slice reports upward for the serialized restack → rerun → re-review → the operator MERGE-GATE. No merge authority exists in this relay.

ACTIONS_GIT_REF: docs-workspace disk action — this re-verdict + successor grant + one s14 INDEX row (in-fence); verification acts were read-only (commit diff reads, the F1 final-byte read, one independent full battery + vet at da491c1c, exit 0). No source byte by this seat, no commit, no lock, no shared path touched.
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; the code worktree is clean at da491c1c)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer folds F3a (RED-first, count-100, battery green) and returns the REVIEW-FOLD report TO s14.planner for the final re-verdict.
