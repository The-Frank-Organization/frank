## REVIEW-FOLD SCOPE - s2-core fold scope before edits

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-fold-scope
PARENT_DISPATCH_ID: s2-core-review
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - fold scope only; merge/S2-close remain operator gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/REVIEW-FOLD-planner-20260704-131809.md
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: FOLD_SCOPE before editing s2-core-impl branch for RF-B1..RF-B6 and RF-M1..RF-M4

FOLD_SCOPE:
- cmd/frank/main.go -> in
- docs/sprints/2026-07-03-s2-slice-2/results/f11-sweep-report.md -> in
- internal/channel/capability_test.go -> in
- internal/channel/server.go -> in
- internal/crashpoint/crashpoint.go -> in
- internal/crashpoint/crashpoint_test.go -> in
- internal/engine/loop.go -> in
- internal/engine/loop_test.go -> in
- internal/engine/ready.go -> in
- internal/engine/ready_test.go -> in
- internal/gc/gc.go -> in
- internal/gc/gc_test.go -> in
- internal/intake/journal.go -> in
- internal/intake/journal_test.go -> in
- internal/intake/writer.go -> in
- internal/intake/writer_test.go -> in
- internal/obligation/obligation.go -> in
- internal/obligation/obligation_test.go -> in
- internal/recover/recover.go -> in
- internal/recover/recover_test.go -> in
- internal/store/projections.go -> in
- internal/store/store.go -> in
- internal/store/store_test.go -> in
- test/fixtures/applicability_map.go -> in
- test/fixtures/f10_test.go -> in
- test/fixtures/f11_test.go -> in
- test/fixtures/f9_c6_test.go -> in
- test/fixtures/iph_test.go -> in
- test/fixtures/main_assembly_test.go -> in
- test/fixtures/s2_sweep_test.go -> in
- test/fixtures/s2setup_test.go -> in
- test/seatproc/testseat.go -> in
FOLD_SCOPE_RESULT: all-in

Scope basis:
- RF-B1/RF-B2/RF-B6: executed class-by-crashpoint, recovery-boundary, quarantine/rotation, and operator-channel fixtures under `test/fixtures`, plus report regeneration.
- RF-B3: GC idempotence, marker completion, recovery obligation, and runtime hook touch `internal/gc`, `internal/obligation`, `internal/recover`, and `cmd/frank`.
- RF-B4: Ready/Diagnostics and channel surface fixes touch `internal/engine`, `internal/intake`, `internal/channel`, `cmd/frank`, and channel/assembly fixtures.
- RF-B5: obligation turn moves from per-connection submit goroutine into the serialized engine loop.
- RF-M1..RF-M4: writer mid-stream, claim/path guardrails, small obligation/GC/owed/quarantine legs, and torn-tail tolerance touch the scoped fixture/journal/store files.

Rules:
- No edit outside the FOLD_SCOPE rows above.
- Any newly required file outside the rows above is a deviation and stops the fold before edit.
- Merge/S2-close/operator-owned OI submit remain out of scope.

ACTIONS_GIT_REF: no source edits yet; this FOLD_SCOPE relay written before implementation worktree edits; implementation branch remains `s2-core-impl@dc8354474e56ade4d0e015e2b7c6a253215c01a7` before fold edits.
FINAL_GIT_STATUS_SHORT: main checkout tracked clean before FOLD_SCOPE write; implementation worktree clean at `s2-core-impl@dc8354474e56ade4d0e015e2b7c6a253215c01a7` before fold edits.
