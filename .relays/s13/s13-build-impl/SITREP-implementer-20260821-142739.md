## SITREP — s13 T4 commit gate held on the master-named s8 150→600 repair

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s13-build-impl
PARENT_DISPATCH_ID: s13-build-impl-2
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — master already assigned the repair to s14 in R8; s13 needs its exact governed carriage, not a new implementation choice
FILED_AT_LOCAL: 20260821-142739
IN_REPLY_TO: s13-build-impl/IMPL-planner-20260821-130456.md
PLAN_LOCK_ID: s13-build-plan @ sha256 00b1f60368d2e227f48a1bf22d8cb0e1ebeefb800bcc316aaab1bc2e5c9ad161
DESIGN_LOCK_ID: DS-s13-m10-module
FROM: s13.implementer
TO: master.orchestrator-planner
CC: s13.planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, s14.planner, s14.implementer
SUBJECT: T4 gate held per R8 — exact s8 timeout repair exists only as foreign uncommitted s14 bytes; return its governed checkpoint/carriage

## Status

T0–T3 are banked on `s13-m10-module` through
`9aee4e957bd4b23a01bb1d0d1afbbd75a9d84554` in four `s13 IMPL:` commits. T4's ruled SQLite
foundation is implemented but UNCOMMITTED: the active R6 byte is exactly `go 1.25.0`, the direct requirement is
exactly `modernc.org/sqlite v1.57.0`, no `toolchain` line exists, `go.sum` is tidy-generated, and the new store
package's focused, race, vet, and reduced-tag batteries are green.

The full required commit gate is HELD under master SITREP
`step3-t4-preflight/SITREP-orchestrator-planner-20260821-131315.md` R8. No s13 variant was authored and no foreign
process or byte was touched.

## Exact failure and isolation

The unchanged `go build ./... && go test ./... && go vet ./...` reached only the pre-existing fixture failure;
every changed package and all fast packages were green:

```text
--- FAIL: TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate (96.29s)
    s8_exit_gate_test.go:311: production suite state = "rejected", want accepted
FAIL github.com/jackli/frank/test/fixtures 144.629s
ok   github.com/jackli/frank/test/invariants 7.175s
FAIL
```

`go list -deps ./test/fixtures | rg 'github.com/jackli/frank/internal/(appipc|appctl)'` returned no output: no T1–T4
package participates in the fixture dependency graph. At the failure, one foreign whole-repository battery was live:

```text
PID 30587  go test -p=1 -count=1 ./...
cwd=/Users/jack/Programming/harness-s15-m9-worker/frank
```

The exact isolated leg then reproduced the master-named wall:

```text
go test ./test/fixtures -run '^TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate$' -count=1 -v
s8_exit_gate_test.go:345: production false suite submit: context deadline exceeded
--- FAIL: TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate (150.03s)
FAIL github.com/jackli/frank/test/fixtures 150.545s
```

The foreign s15 battery exited normally; s13 did not interrupt it.

## Exact repair state

The s14-owned file exists in its foreign worktree as an UNCOMMITTED modification:

```text
/Users/jack/Programming/harness-s14-m8-connector/frank/test/fixtures/s8_exit_gate_test.go
sha256 c339bc0f2b2467d29b05d1a273acd905cc1ec66b8922d741698ae05d03c1af8f
s14 branch HEAD 471532a1 (T9)
```

Its diff moves the context deadline `150*time.Second → 600*time.Second` and adds the capacity-boundary diagnostics
that name `ctx_err`, elapsed duration, and nested-suite observed color. This is outside s13's SCOPE_DIFF and is not
yet a Git object: s14 `git status --short` reports that fixture modified beside its current T10 work. The latest
fetched `origin/main@1a25d46822fc91f4b15b0fb3ded0a7cd56c4df71` still carries 150 seconds. Therefore s13 cannot lawfully
copy, cherry-pick, re-author, or infer a carriage for those bytes.

## T4 completed evidence below the held leg

- RED preceded `Open`, schema/version errors, and the store transaction API.
- Genesis/migration: `user_version=1`; v0 forward migration; higher-version and corrupt refusal.
- Floor: WAL, `synchronous=FULL`, 0700 runtime directory, 0600 database, immediate one-writer transactions.
- Exact amended 15-table census; `epoch_transitions`, `crossing_ops`, and deferred `events` absent.
- Closed state/counter/singleton negatives; transaction commit/rollback; no provider payload or payload-named column.
- `go test ./internal/appctl/store` green; `go test -race ./internal/appctl/store` green at 1.597s;
  `go vet ./internal/appctl/store` green.
- `go test -tags frank_test_reduced_limits ./internal/appipc/... ./internal/appctl/...` green:
  appipc 0.239s, store 0.479s, testutil 0.649s.
- Sequence-honest record:
  `.relays/s13/batteries/T4-sqlite-store-20260821.txt` in the isolated worktree.

Done: T0–T3 banked; T4 implementation and every in-scope focused/reduced leg green.

Not done: T4 commit; T5–T15; terminal IMPL report/review/restack/merge-gate sequence.

Blocked: only the required whole-repo commit gate, on the exact R8-owned s8 timing repair.

Scope drift risk: copying the visible s14 worktree diff would violate both ownership and s13's closed SCOPE_DIFF;
retrying the same 150-second wall would evade R8 rather than consume its named repair.

Next requested action: s14 completes and banks its owned repair; master returns a checkpoint/relay that binds the
exact repair commit/tree/file hash and the authorized carriage into s13's test gate (including any bounded scope
addition or base movement required). On that return, s13 consumes only those exact governed bytes, reruns the
unchanged full gate, banks T4, and continues straight-through.

ACTIONS_GIT_REF: docs-workspace disk action — this append-only blocker SITREP + one live-EOF s13 INDEX row; isolated source worktree remains s13-m10-module@9aee4e957bd4b23a01bb1d0d1afbbd75a9d84554 with uncommitted T4-only in-scope paths (go.mod, new go.sum, new internal/appctl/store tree, T4 battery); no merge, push, fixture edit, foreign-process action, or external side effect
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/.relays/s14/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-impl/SITREP-implementer-20260821-142739.md
?? frank/.relays/s14/s14-build/IMPL-planner-20260821-142737.md
?? frank/.relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-142139.md
?? frank/.relays/s14/s14-build/PLAN-planner-20260821-140427.md
