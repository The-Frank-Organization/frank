## FOLD_SCOPE — successor s14 end-review fold, F1+F2 preserved and F3 transport counter race added in-fence

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-fold-scope-2
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair-internal successor fold; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-184953
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-184809.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: successor pre-edit scope — F1+F2 four paths preserved, F3 transport mechanism/test added, all-in with pinned-semantic stop intact

FOLD_SCOPE:
- frank/internal/connector/control/control.go -> in
- frank/internal/connector/control/control_test.go -> in
- frank/internal/connector/attempt/attempt.go -> in
- frank/internal/connector/attempt/attempt_test.go -> in
- frank/internal/connector/transport/transport.go -> in
- frank/internal/connector/transport/transport_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

F1 and F2 remain at their already-observed RED/GREEN bytes, uncommitted. F3's pre-fix RED is the unchanged `TestRetryAfterIsRecordedAndNeverActedOn` mismatch reproduced both in the full battery and twice under isolated `-count=100`. The minimal fold will join the existing `httptrace.WroteRequest` event before a successful response becomes observable; it will not rename counters, alter no-replay behavior, change fixture-4 vector semantics, or move any frozen carrier byte.

After the F3 targeted GREEN and count-100 stress, the exact fixture-4 cuts, repeated race suites, full serialized battery, vet, and a non-persistent tidy tripwire will run before commit. Any need outside the seven rows or any r12-pinned semantic change stops before edit and routes a new deviation.

ACTIONS_GIT_REF: governance-only successor pre-edit barrier — this FOLD_SCOPE relay plus one append-only live-EOF s14 INDEX row; source retains only the four prior in-scope uncommitted F1/F2 files at `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`; no transport edit yet
FINAL_GIT_STATUS_SHORT:
 M internal/connector/attempt/attempt.go
 M internal/connector/attempt/attempt_test.go
 M internal/connector/control/control.go
 M internal/connector/control/control_test.go
Next requested action: s14.implementer applies the minimal F3 write-event join, proves targeted and count-100 GREEN plus fixture-4 invariants, runs the full E2 gate, commits the bounded seven-row fold, and returns the REVIEW-FOLD report. No merge authority is requested or implied.
