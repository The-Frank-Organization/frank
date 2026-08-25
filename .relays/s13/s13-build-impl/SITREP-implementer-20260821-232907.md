## SITREP — carriage v2 consumed exactly; mandatory s8 acceptance remains RED on checksum-list cache overreach

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s13-build-impl
PARENT_DISPATCH_ID: s13-build-impl-2
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this is a failed named acceptance leg requiring an owner successor, not a product or merge decision
FILED_AT_LOCAL: 20260821-232907
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-232450.md
PLAN_LOCK_ID: s13-build-plan @ sha256 00b1f60368d2e227f48a1bf22d8cb0e1ebeefb800bcc316aaab1bc2e5c9ad161
DESIGN_LOCK_ID: DS-s13-m10-module
FROM: s13.implementer
TO: master.orchestrator-planner
CC: s13.planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, s14.planner, s14.implementer, m-3.planner, m-7.planner
SUBJECT: v2 one-file successor is byte-exact but s13's required isolated s8 leg and normal-root preseed leg still fail

## Exact v2 consumption

The addressed carriage is SHA-256 `82eac3276f7434eb0bf49b818e495074daacf35971d69e0c4f86a71a867123da`.
Commit `eaf8faa1b96eae254c6788b9dd49386082a3acd5` exists with parent `4aea922e458dba50de2c21edc10bf989694b7d8c`
and changes only `frank/internal/executor/executor_test.go`. I applied it with `git cherry-pick -n`.

The staged census remains exactly the three licensed carriage paths, with every v2 binding satisfied:

```text
c339bc0f2b2467d29b05d1a273acd905cc1ec66b8922d741698ae05d03c1af8f  frank/test/fixtures/s8_exit_gate_test.go
aaa980d445f12e42adc804c2ba724574326bd7ccadee0264137eeacac5466806  frank/internal/executor/executor.go
bbc9d434c030817fabc9bef26d7ee63c984bbeb3819f939a3b95b59dbd1d742f  frank/internal/executor/executor_test.go
```

No base, executor implementation byte, fixture byte, or T4 source byte moved.

## Named acceptance leg failed

I ran acceptance check (b) exactly as directed, before the full gate:

```text
go test ./test/fixtures -run '^TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate$' -count=1 -v
production suite state = "rejected", want accepted
degradation_notes: observe-machinery-fault
body: dogfood-battery-green:observe-machinery-fault
--- FAIL: TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate (1.11s)
FAIL github.com/jackli/frank/test/fixtures 1.758s
```

Because mandatory precursor (b) failed, acceptance check (c), the unchanged full commit gate, did not run.

## Root cause reduced in the normal s13 environment

The successor's focused group proves R11's zero-external-closure test is green but its consuming-root preseed test is
still red without the old empty-cache reproducer and without any adversarial environment:

```text
TestSpawnPreseedsRunLocalModuleCacheWithProxyOff: FAIL executor-module-cache-miss
TestSpawnNamesMissingOfflineModuleWithoutWaitingForNetwork: PASS
TestSpawnNestedExecutorPackageFromZeroExternalClosure: PASS
TestSpawnRetainsPrivateCappedTailOutsideBareFailureVerdict: PASS
FAIL github.com/jackli/frank/internal/executor 5.268s
```

`copyRootModuleClosure` now correctly copies s13's own `go.mod` and `go.sum`. The remaining failure is in the unchanged
R9 implementation: `goSumModules` treats every non-`/go.mod` checksum as a mandatory offline download target. s13 has
25 such entries, while 15 lack complete artifacts in the ordinary host cache. The missing entries include
`github.com/google/pprof`, `github.com/hashicorp/golang-lru/v2`, `golang.org/x/{mod,sync,tools}`, and ten modernc
compiler/build modules. Those checksums are present from the module graph, but their packages are not imported by the
normal s13 build; `go test ./internal/appctl/store -count=1` remains green at 0.311s in the same environment.

Therefore the relay's stated premise that normal builds populate the required cache is false for this legitimate
consumer: normal Go build/test populates needed imported-package artifacts, not every historical or build-only module
zip named by `go.sum`. A deliberate online `go mod download` cache-warm would hide the ambient-cache dependency and make
the acceptance non-portable, so none was performed.

## Hold and requested return

The v2 carriage remains staged and byte-exact. T4 remains uncommitted; T5–T15 remain not started.

Requested action: route an owner successor whose acceptance battery uses a non-empty consuming root with extra valid
`go.sum` entries absent from the host artifact cache, then return new hash-bound bytes. The implementation must seed the
modules actually needed for the governed suite without requiring unrelated checksum-listed artifacts or network warming.
On return s13 will consume only the successor bytes, rerun the isolated s8 leg, then the unchanged full gate.

Done: v2 identity verified and consumed; named acceptance (b) executed; failure reduced to 15 absent artifacts out of
25 checksum targets; real s13 SQLite package independently reconfirmed green; T4 battery appended.

Not done: acceptance (c); T4 commit; T5–T15; review/restack/merge-gate sequence.

Blocked: required s8 acceptance, on R9's all-`go.sum` ambient-cache precondition.

ACTIONS_GIT_REF: docs-workspace disk action — this append-only SITREP + one live-EOF s13 INDEX row; isolated s13 worktree remains at 9aee4e957bd4b23a01bb1d0d1afbbd75a9d84554 with three exact staged carriage paths and uncommitted T4 paths; no source commit, merge, push, base move, network cache-warm, foreign-worktree mutation, or external side effect
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-impl/SITREP-implementer-20260821-232907.md
