## REVIEW-FOLD report - s6-core implementer - s6-core-impl-fold-remint

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s6-core-impl-fold-remint
PARENT_DISPATCH_ID: s6-core-impl-remint-ruling
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: REVIEW-FOLD-planner-20260707-103642.md
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: fold report for released re-mint crash-window option A fix

FOLD_SCOPE:
- internal/seat/binding.go -> in
- internal/seat/binding_test.go -> in
- cmd/frank/main.go -> in
- test/fixtures/s6_mint_test.go -> in
- test/fixtures/s6_iph_test.go -> in
- docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md -> in
- .relays/s6/s6-core-impl/FOLD_SCOPE-implementer-20260707-103947.md -> in
- .relays/s6/s6-core-impl/REVIEW-FOLD-implementer-20260707-232510.md -> in
- .relays/s6/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF:
branch `s6-transport-impl` at `1f6cd08` (`s6: fold remint crash-window provenance recovery`).

FINAL_GIT_STATUS_SHORT:
none - clean tree in `~/frank-s6-impl` after commit `1f6cd08`.

Fold summary:
- `internal/seat/binding.go`: binding rows now carry optional `realized_mint_ref`, populated only by derived `seat_mint` completion; credential replacement and the provenance ref are persisted by the same binding-table write.
- `cmd/frank/main.go`: live `seat_mint` completion passes the accepted pivot relay id, and startup recovery now scans latest accepted `seat_mint` pivots by store commit order before opening the authenticated socket. A stale or missing binding row is replaced; an already realized pivot is a no-op.
- `test/fixtures/s6_mint_test.go`: added the SIGKILL fixture for crash after accepted re-mint pivot record rename and before binding replacement; restart must repair before serve, reject the superseded credential on first post-restart auth, and realize the latest pivot.
- `test/fixtures/s6_iph_test.go`: extended the I-PH custody sweep so the new binding marker stays out of read/project surfaces and every non-binding file.
- `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md`: added the directed one-line Section 7 option-A addendum citing `s6-fidelity-m1/SITREP-implementer-20260707-102208.md`.

Red-first evidence:
- RED command: `go test ./test/fixtures -run TestS6RemintCrashBeforeBindingReplacementRecoversBeforeServe -count=1`
- RED result before production fix: failed after restart because the binding row had empty `realized_mint_ref` for the latest accepted re-mint pivot. Credential material from the failure row is intentionally elided from this relay.
- GREEN command: `go test ./test/fixtures -run TestS6RemintCrashBeforeBindingReplacementRecoversBeforeServe -count=1`
- GREEN result after fix: passed.

Redline compliance:
- R1 ok: the only new binding-row field is optional `realized_mint_ref`, storing the accepted `seat_mint` pivot relay id only; no timestamp, counter, credential material, hash, socket/session id, or role text was added.
- R2 ok: `MintOrReplace` writes the replacement credential and `realized_mint_ref` into one binding row before the single atomic binding-table persist.
- R3 ok: recovery computes latest accepted `seat_mint` pivot per seat from store commit order, replaces missing/stale rows, no-ops realized rows, and leaves legacy/genesis rows with no accepted pivot untouched.
- R4 ok: the scan runs before `channel.ServeAuthenticated`; the crash fixture removes the stale Unix socket and proves the superseded credential fails first post-restart auth.
- R5 ok: accepted records are canonical authority; the binding row is only the derived completion artifact.
- R6 ok: the marker remains in 0600 binding custody and is covered by read/project/non-binding-file negative sweeps.
- R7 ok: no activation derivation, roster rule, or R1/intake-generation rule changed.
- R8 ok: the SIGKILL fixture kills after the accepted pivot record rename, restarts, proves pre-serve repair, checks latest-pivot realization, and checks old credential custody cleanup.

Verification:
- Focused checks passed:
  `go test ./test/fixtures -run TestS6RemintCrashBeforeBindingReplacementRecoversBeforeServe -count=1`
  `go test ./internal/seat -count=1`
  `go test ./test/fixtures -run 'TestS6(RemintCrashBeforeBindingReplacementRecoversBeforeServe|IPHSeatMintReplyCarveOutsScoped|StartupCompletesCommittedInitialSeatMintBinding)' -count=1`
- Full battery passed on final scoped tree:
  `go test -count=1 ./...`
  `go vet ./...`
  `go test -race -count=1 ./internal/seat ./internal/engine ./internal/store ./internal/channel`
  `git diff --check`

Scope hygiene:
- No files outside the pre-filed `FOLD_SCOPE` were committed.
- `internal/gc/gc_test.go` was not touched; the absorption-ruling citation condition is not triggered.
- No merge or PR action is claimed by this relay.

Next requested action:
s6-core.planner reviews `s6-transport-impl@1f6cd08`; if accepted, proceed to the exit-gate pass described by the planner/orchestrator trail.
