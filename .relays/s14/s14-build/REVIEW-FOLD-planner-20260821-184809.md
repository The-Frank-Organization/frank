## REVIEW-FOLD — the SUCCESSOR fold relay (supersedes `REVIEW-FOLD-planner-20260821-182628.md` as the live fold grant; F1+F2 carried, **F3 ADDED**, the scope widened IN-FENCE): the full-battery blocker the fold surfaced (`SITREP-implementer-20260821-183627.md`) is DISPOSITIONED BY THIS SEAT as a THIRD END-REVIEW FINDING, not a master escalation — **`frank/internal/connector/transport/**` is inside s14's charter fence** (the prior OUT rows were out-of-FOLD-SCOPE only), the defect is pre-existing at head `316f97a6` (zero transport bytes differ; isolated count=100 reproduction; no F1/F2 dependency — the implementer's systematic isolation stands verified), and the fix moves NO frozen byte: the counter stays m-8-internal (addendum P2a — never a carrier or E3 input), and r12 §5.1's own mechanism words count writes "at the instrumented connection boundary BELOW the encoder" — the current `httptrace.WroteRequest` observation (transport.go:133) sits ABOVE it, so a boundary-faithful fix moves the implementation CLOSER to the frozen text. Master is CC'd and may countermand; if the fix were to require changing any r12-pinned semantic (the counter vocabulary, the no-replay pins, fixture-4's vectors), STOP — that becomes the escalation.

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
FILED_AT_LOCAL: 20260821-184809
IN_REPLY_TO: frank/.relays/s14/s14-build/SITREP-implementer-20260821-183627.md
FROM: s14.planner
TO: s14.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: successor fold — F1+F2 (GREEN, preserved) + F3 the transport write-completion truth race; scope widened in-fence; commit + battery + report

FOLD_SCOPE:
- frank/internal/connector/control/control.go -> in
- frank/internal/connector/control/control_test.go -> in
- frank/internal/connector/attempt/attempt.go -> in
- frank/internal/connector/attempt/attempt_test.go -> in
- frank/internal/connector/transport/transport.go -> in
- frank/internal/connector/transport/transport_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

Scope note: any discovered need outside these rows is a STOP-before-edit deviation escalation; any r12-pinned-semantic touch is a STOP regardless of path.

## §1 — The findings

**F1 + F2 — carried as folded (your RED/GREEN evidence on the trail at `183627` stands; the four edits preserved uncommitted).** Nothing re-opened.

**F3 — NEW — `transport.go:133/173-178`, the request-write-completion truth race:** `RequestWriteCompleted` increments only in the `httptrace.WroteRequest` callback, which the HTTP/1 write goroutine may run AFTER the response is exposed; `Counters()` joins nothing, so a successful response can be observed at `{1,1,1,0}` — a counter that can misreport its own truth window (your isolation: reproducible at count=100; pre-existing at `316f97a6`; no F1/F2 dependency). Fold inside these semantics: make write-completion observation truthful at the r12 §5.1 boundary — EITHER count completion at the instrumented connection wrapper's write return (below the encoder, the spec's own stated mechanism) OR join the write event before a response/counters are exposed — your mechanism choice; the counter VOCABULARY and every fixture-4 vector semantic stay byte-identical. RED first: your count=100 reproduction loop (cite the pre-fix red); GREEN: the same loop green + the fixture-4 per-cut vectors still exact + the full battery.

## §2 — The gate on your report

Commit shape your choice (the F1/F2 commit(s) + an F3 commit, or one bounded fold commit — each citing this fold relay); targeted tests after each finding; then the full E2 battery (`go test -mod=readonly -p=1 -count=1 ./...` + vet + tidy tripwire) green at the final commit; literal `git status --short` in the report. On the REVIEW-FOLD report I re-verdict; CLEAN ⇒ the slice reports upward for the serialized restack → rerun → re-review → the operator MERGE-GATE. No merge authority exists in this relay.

ACTIONS_GIT_REF: docs-workspace disk action — this successor fold relay + one s14 INDEX row (in-fence). No source byte by this seat, no commit, no lock, no shared path touched; the routing disposition above is a planner fence reading with master CC'd for countermand.
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/** + the implementer's preserved uncommitted F1/F2 edits on the code worktree)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer folds F3 (RED-first), commits the bounded fold set, runs the full battery green, and returns the REVIEW-FOLD report TO s14.planner for the re-verdict.
