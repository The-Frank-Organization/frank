## REVIEW-FOLD - s2-core review fold report for RF-B/RF-M round 1

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-fold-report
PARENT_DISPATCH_ID: s2-core-review
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - fold report only; merge/S2-close remain operator gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/REVIEW-FOLD-planner-20260704-131809.md
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: review fold report for branch s2-core-impl@9e4829c

Summary:
- Folded the round-1 blocker/must-fix set on existing branch `s2-core-impl`.
- Commit: `9e4829c3ddfa7a3521de1b0b84952d18b5035b31` (`s2 IMPL: review-fold round 1 blockers and fixtures`).
- No PR opened, no push performed, no merge authority claimed.
- Operator-owned OI-S1-F11-SWEEP submit, exit gate, REVIEW-FOLD acceptance, and S2-close remain ahead.

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
- internal/obligation/owed.go -> in
- internal/obligation/owed_test.go -> in
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

Fold status:
- RF-B1/RF-B2/RF-B6: expanded child-process F11 coverage for S2 crashpoints and clean-completion class paths; report now says executed coverage rather than static map-only coverage.
- RF-B3: fixed GC marker resume/idempotence, added recovery GC pass, and covered marker/unlink convergence.
- RF-B4: main now serves only project/read on diagnostics, constructs submit loop/writer only after `result.Ready`, and channel `tools/list` reflects the active surface.
- RF-B5: moved obligation completion into the serialized loop turn; removed per-submit `gate.Complete` from connection goroutines.
- RF-M1: changed the F9 crash fixture to use the single Writer path at `post_intake_fsync:5`.
- RF-M2/RF-M3: added execution/report guardrails, live K2 channel fixture, non-operator owed item fixture, GC/recovery fixtures, and torn-tail journal/redo fixtures.
- RF-M4: implemented torn trailing JSON-line tolerance for intake and redo segment readers.

Optional dispositions:
- Kept `engine.TestReady` in production `ready.go` with the existing runtime `testing.Testing()` guard; the runtime path no longer self-mints Ready in `cmd/frank`.
- Kept the extra `-store` flag unchanged.
- Accepted `internal/engine/quarantine_test.go` as already within the Task-8/K2 intent recorded by the planner review fold.

Verification:
- `go test ./...`: PASS.
- `go test -race ./internal/intake ./internal/engine ./internal/obligation -v`: PASS.
- `go vet ./...`: PASS.
- `golangci-lint run ./...`: PASS, `0 issues`.
- `git diff --check`: PASS.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-impl/REVIEW-FOLD-implementer-20260704-133423.md`: OK.

Evidence level:
- E2 local tests/lint/vet/race evidence on branch `s2-core-impl`.
- Not claimed: E3/E4, operator-owned OI submit execution, merge-readiness, S2-close, or live deployment proof.

Remaining risk / planner attention:
- The F11 fixture now executes crash-expected S2 cells and one clean-completion path per class; it still does not perform a literal child process for every clean-completion class x crashpoint cell. The report wording avoids that overclaim.
- Ready construction remains exposed for recovery/test plumbing, but the shipped assembly path no longer constructs the externally served loop/writer before recovery returns Ready.
- The operator-owned OI-S1-F11-SWEEP submit remains intentionally unperformed by this fold.

Next requested action:
- Planner targeted re-verify of `s2-core-impl@9e4829c3ddfa7a3521de1b0b84952d18b5035b31`.

ACTIONS_GIT_REF: branch `s2-core-impl@9e4829c3ddfa7a3521de1b0b84952d18b5035b31`; commit `9e4829c3ddfa7a3521de1b0b84952d18b5035b31`; FOLD_SCOPE artifact `.relays/s2/s2-core-impl/REVIEW-FOLD-SCOPE-implementer-20260704-132203.md` linted OK before source edits; report relay `.relays/s2/s2-core-impl/REVIEW-FOLD-implementer-20260704-133423.md` written after verification.
FINAL_GIT_STATUS_SHORT: implementation worktree clean after commit; main checkout tracked clean before report write, expected to remain tracked-clean because `.relays/` is gitignored.
