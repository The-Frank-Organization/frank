## REVIEW-FOLD - s2-core review fold report for round 3 mint sanction

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-r3-fold-report
PARENT_DISPATCH_ID: s2-core-review-r3
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - fold report only; merge/S2-close and the Task-13.5 submit remain operator gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/REVIEW-FOLD-planner-20260704-143700.md
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: review fold round 3 report for branch s2-core-impl@7e5d1a3

Summary:
- Implemented the sanctioned conductor-internal `-mint <seat>` admin-time flag.
- Commit: `7e5d1a3bc9d63aff24dc4a92c5064a265054c5be` (`s2 IMPL: add admin mint flag`).
- No PR opened, no push performed, no merge authority claimed.
- The real Task-13.5 OI-S1-F11-SWEEP submit remains operator-executed and out of scope.

FOLD_SCOPE:
- cmd/frank/main.go -> in
- test/fixtures/main_assembly_test.go -> in
FOLD_SCOPE_RESULT: all-in

Fold status:
- Added `frank -root <store> -mint <seat> [-role <role>] [-operator]`.
- `-mint` runs before serve/recovery and before `-operator-submit`; it opens only the existing binding manager and calls `seat.Mint`, so the binding-table shape is unchanged.
- Admin-time guard: if the selected socket is live, `-mint` exits nonzero with `conductor is serving; -mint is admin-time only` before opening/changing bindings.
- Credential handling: successful mint prints exactly one `credential=<hex>` line to admin stdout; no success stderr; the credential is persisted only through the existing binding table.
- Seat-facing surface: no `mint` tool is added to `tools/list` or tool descriptions; the successful minted credential resolves through normal authenticated connect.

Fixture coverage:
- `TestFrankBinaryMintSeatAdminTimeCredential` covers `-mint system` rejected with `seat.ErrReservedSeatName` text and binding table unchanged.
- It covers duplicate mint rejected with `seat.ErrSeatAlreadyBound` text and binding table byte-identical.
- It covers successful operator mint, stdout-only credential line, credential not present in records/mailboxes/projections/outbox, authenticated connect with the minted credential, full-surface registry exactly `submit,project,read`, and no `mint` registry entry.
- It covers live-socket mint rejection and binding table unchanged.

Red-first evidence:
- `go test ./test/fixtures -run TestFrankBinaryMintSeatAdminTimeCredential -count=1` failed before implementation because `-mint` was not a defined flag.

Verification:
- `python3 <relay-lint tools>/relay-lint.py .relays/s2/s2-core-impl/REVIEW-FOLD-SCOPE-implementer-20260704-143930.md`: OK.
- `go test ./test/fixtures -run TestFrankBinaryMintSeatAdminTimeCredential -count=1`: PASS.
- `go test ./test/fixtures -run 'TestP1|TestS2ApplicabilityMapCoversEveryClassPointAndReport|TestFrankBinary(MintSeatAdminTimeCredential|ServesReadOnlyDiagnosticsOnDigestMismatch|AssemblesAuthenticatedSubmitProjectRead)' -count=1`: PASS.
- `go test ./test/fixtures -run TestSweep -count=1`: PASS.
- `go test ./cmd/frank -count=1`: PASS (`[no test files]`).
- `go test ./test/fixtures -count=1`: PASS.
- `go test -count=1 ./...`: PASS.
- `go test ./...`: PASS.
- `go test -race -count=1 ./internal/intake ./internal/engine ./internal/obligation -v`: PASS.
- `go vet ./...`: PASS.
- `golangci-lint run ./...`: PASS, `0 issues`.
- `git diff --check`: PASS.

Evidence level:
- E2 local tests/lint/vet/race evidence on branch `s2-core-impl`.
- Not claimed: E3/E4, real operator-owned OI submit execution, merge-readiness, S2-close, or live deployment proof.

Next requested action:
- Planner targeted re-verify of `s2-core-impl@7e5d1a3bc9d63aff24dc4a92c5064a265054c5be`.

ACTIONS_GIT_REF: branch `s2-core-impl@7e5d1a3bc9d63aff24dc4a92c5064a265054c5be`; commit `7e5d1a3bc9d63aff24dc4a92c5064a265054c5be`; FOLD_SCOPE artifact `.relays/s2/s2-core-impl/REVIEW-FOLD-SCOPE-implementer-20260704-143930.md` linted OK before source edits; report relay `.relays/s2/s2-core-impl/REVIEW-FOLD-implementer-20260704-144230.md` written after verification.
FINAL_GIT_STATUS_SHORT: implementation worktree clean after commit; main checkout tracked clean before report write, expected to remain tracked-clean because `.relays/` is gitignored.
