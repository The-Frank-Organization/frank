## REVIEW-FOLD - s2-core review fold report for round 2

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-r2-fold-report
PARENT_DISPATCH_ID: s2-core-review-r2
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - fold report only; merge/S2-close remain operator gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/REVIEW-FOLD-planner-20260704-134500.md
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: review fold round 2 report for branch s2-core-impl@16342e0

Summary:
- Folded RB2-1 and RB2-2 on existing branch `s2-core-impl`.
- Commit: `16342e0ce8fd28791fac21f261cc7404cab30d9b` (`s2 IMPL: review-fold round 2 trace and O3 fixture`).
- No PR opened, no push performed, no merge authority claimed.
- The real Task-13.5 OI-S1-F11-SWEEP submit remains operator-executed; this fold adds fixture semantics only.

FOLD_SCOPE:
- internal/crashpoint/crashpoint.go -> in
- internal/crashpoint/crashpoint_test.go -> in
- test/fixtures/applicability_map.go -> in
- test/fixtures/f11_test.go -> in
- test/fixtures/main_assembly_test.go -> in
- test/fixtures/s2_sweep_test.go -> in
- docs/sprints/2026-07-03-s2-slice-2/results/f11-sweep-report.md -> in
FOLD_SCOPE_RESULT: all-in

Fold status:
- RB2-1 folded with hit-trace row equality. `crashpoint.Hit` now appends to `FRANK_TEST_HIT_TRACE` before crashpoint target filtering. `TestS2ApplicabilityMapRowsMatchHitTrace` runs one child per class and asserts the trace-observed fired set equals the class row's `crash-expected` set exactly, both directions.
- RB2-1 map rows corrected to the true fired sets, including directory fsyncs, projection/delivery points where intents run, held outcome reply, recovery/genesis writes, quarantine/GC setup points, and owed-item/owed-disposition projection points.
- RB2-1 owed fixture bug corrected: F11 owed classes now submit `record_kind: owed_item` / `owed_disposition` instead of the class labels `owed-item` / `owed-disposition`.
- RB2-1 report updated to state the evidence mechanism: crash-expected child SIGKILLs plus hit-trace equality for row falsification, with clean-completion evidence tied to non-fired cells.
- RB2-2 folded with `TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition`: authenticated operator channel submit of the real OI-S1-F11-SWEEP typed owed payload, open surfacing via `project`, `read`, and `projections/owed/OPEN.md`, then an `owed_disposition` citing the F11 sweep report empties the open set.
- S1 assertions byte-identical by scope: no S1 source/doc assertion files changed.

Red-first evidence:
- `go test ./internal/crashpoint -run TestHitTraceAppendsNamesWithoutArmingCrashpoint -count=1` failed before implementation because no trace file was produced.
- `go test ./test/fixtures -run TestS2ApplicabilityMapRowsMatchHitTrace -count=1` failed before map correction with fired-but-clean row mismatches across the reviewed classes.

Verification:
- `go test ./internal/crashpoint -run TestHitTraceAppendsNamesWithoutArmingCrashpoint -count=1`: PASS.
- `go test ./test/fixtures -run 'TestS2ApplicabilityMap(CoversEveryClassPointAndReport|RowsMatchHitTrace)|TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition' -count=1`: PASS.
- `go test ./test/fixtures -count=1`: PASS.
- `go test -count=1 ./...`: PASS.
- `go test ./...`: PASS.
- `go test -race ./internal/intake ./internal/engine ./internal/obligation -v`: PASS.
- `go vet ./...`: PASS.
- `golangci-lint run ./...`: PASS, `0 issues`.
- `git diff --check`: PASS.
- `python3 <relay-lint tools>/relay-lint.py .relays/s2/s2-core-impl/REVIEW-FOLD-SCOPE-implementer-20260704-135015.md`: OK.
- `python3 <relay-lint tools>/relay-lint.py .relays/s2/s2-core-impl/REVIEW-FOLD-implementer-20260704-135628.md`: OK.

Evidence level:
- E2 local tests/lint/vet/race evidence on branch `s2-core-impl`.
- Not claimed: E3/E4, real operator-owned OI submit execution, merge-readiness, S2-close, or live deployment proof.

Next requested action:
- Planner targeted re-verify of `s2-core-impl@16342e0ce8fd28791fac21f261cc7404cab30d9b`.

ACTIONS_GIT_REF: branch `s2-core-impl@16342e0ce8fd28791fac21f261cc7404cab30d9b`; commit `16342e0ce8fd28791fac21f261cc7404cab30d9b`; FOLD_SCOPE artifact `.relays/s2/s2-core-impl/REVIEW-FOLD-SCOPE-implementer-20260704-135015.md` linted OK before source edits; report relay `.relays/s2/s2-core-impl/REVIEW-FOLD-implementer-20260704-135628.md` written after verification.
FINAL_GIT_STATUS_SHORT: implementation worktree clean after commit; main checkout tracked clean before report write, expected to remain tracked-clean because `.relays/` is gitignored.
