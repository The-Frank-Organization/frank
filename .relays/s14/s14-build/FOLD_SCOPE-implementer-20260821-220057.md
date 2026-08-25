## FOLD_SCOPE — successor s14 end-review fold, F3a complete-request write accounting only

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-fold-scope-3
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair-internal successor fold; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-220057
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-215400.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: successor pre-edit scope — F3a complete-request write accounting only, all-in with pinned-semantic stop intact

FOLD_SCOPE:
- frank/internal/connector/transport/transport.go -> in
- frank/internal/connector/transport/transport_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

The source worktree is clean at `s14-m8-connector@da491c1c458a9c22c322c2e08b1a4f81ba029b37`. The RED mutation is narrowly named: a body larger than one transport write whose connection fails on a later write must leave the literal counter vector `{DialAttempts:1, ConnectionsEstablished:1, RequestWriteStarted:1, RequestWriteCompleted:0}`. The existing implementation's first-full-Write completion increment must make that regression fail before any production edit.

The minimal fold will make completion depend on the complete encoded request crossing the instrumented connection, while preserving the four existing cut vectors, no-replay behavior, counter vocabulary, and all frozen carrier bytes. After targeted RED/GREEN, the exact existing vectors, count-100 stress, repeated race suites, full serialized battery, vet, and non-persistent tidy tripwire will run. Any need outside the three rows or any r12-pinned semantic change stops before edit and routes a new deviation.

ACTIONS_GIT_REF: governance-only successor pre-edit barrier — this FOLD_SCOPE relay plus one append-only live-EOF s14 INDEX row; source remains clean at `s14-m8-connector@da491c1c458a9c22c322c2e08b1a4f81ba029b37`; no source or test edit yet
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; source worktree clean)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer drives the named multi-write failure vector RED-first, folds the minimal F3a mechanism, proves the preserved vectors and E2 battery at one bounded commit, and returns the REVIEW-FOLD report. No merge authority is requested or implied.
