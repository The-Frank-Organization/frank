## SITREP — exact R8/R9 carriage consumed; R9 is dependency-closure-sensitive and still blocks s13 T4

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s13-build-impl
PARENT_DISPATCH_ID: s13-build-impl-2
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this is an owner-repair/carriage defect, not a new product choice or merge request
FILED_AT_LOCAL: 20260821-223545
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-222814.md
PLAN_LOCK_ID: s13-build-plan @ sha256 00b1f60368d2e227f48a1bf22d8cb0e1ebeefb800bcc316aaab1bc2e5c9ad161
DESIGN_LOCK_ID: DS-s13-m10-module
FROM: s13.implementer
TO: master.orchestrator-planner
CC: s13.planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, s14.planner, s14.implementer, m-3.planner, m-7.planner
SUBJECT: R8/R9 exact carriage applied, but R9 inherits the executor's restricted GOMODCACHE and misses its test-only x/text module on s13

## Carriage consumed exactly

The addressed master carriage is SHA-256
`a17ad971f4b8e4e135c4338227f2d6ed6292d4ee2a358c3fe354d825fa1b3bb8`. I applied the two licensed Git objects with
`git cherry-pick -n` in the isolated `s13-m10-module` worktree:

```text
b9b2b3c2f99405055ef1f48fb0ee57adff1ed495
b86b8bc17ba56302d33b7aba4cf3f8af404d70ea
```

The staged census is exactly the three granted paths:

```text
M  frank/internal/executor/executor.go
A  frank/internal/executor/executor_test.go
M  frank/test/fixtures/s8_exit_gate_test.go
```

Their live SHA-256s exactly equal the carriage bindings:

```text
c339bc0f2b2467d29b05d1a273acd905cc1ec66b8922d741698ae05d03c1af8f  frank/test/fixtures/s8_exit_gate_test.go
aaa980d445f12e42adc804c2ba724574326bd7ccadee0264137eeacac5466806  frank/internal/executor/executor.go
02f244f56c632f0b9a89fcebb0c6abb636996f9ea7da5d1adef0cf66174119cd  frank/internal/executor/executor_test.go
```

No base moved, no carried byte diverged, and no scope was added. The prior T4 paths remain unstaged/untracked at
`HEAD 9aee4e957bd4b23a01bb1d0d1afbbd75a9d84554`.

## New deterministic gate failure

The carried fixture no longer reaches the 150-second capacity wall. Its exact isolated leg now fails in about one
second before the nested suite starts:

```text
go test ./test/fixtures -run '^TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate$' -count=1 -v
production suite state = "rejected", want accepted
degradation_notes: observe-machinery-fault
body: dogfood-battery-green:observe-machinery-fault
--- FAIL: TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate (1.07s)
FAIL github.com/jackli/frank/test/fixtures 1.637s
```

The executor returns before retaining a suite diagnostic: `seedGoModuleCache` maps the failure to
`executor-module-cache-miss`; the observe gate then correctly emits the public `observe-machinery-fault` class.

## Root cause proved

The R9 executor sets the nested suite's `GOMODCACHE` to its restricted run-local cache. The suite contains the new
`TestSpawnPreseedsRunLocalModuleCacheWithProxyOff`, and its `goModuleCachePath` deliberately honors inherited
`GOMODCACHE`. The inner executor therefore treats the outer restricted cache as the host source for its synthetic
`golang.org/x/text v0.41.0` dependency.

A controlled direct run of the exact carried test against an empty explicit inherited cache fails immediately:

```text
GOMODCACHE=<empty-absolute-dir> GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
  go test ./internal/executor -run '^TestSpawnPreseedsRunLocalModuleCacheWithProxyOff$' -count=1 -v
offline verdict = observe.CheckVerdict{Outcome:"unsafe", Predicate:"blocked", FailingDetail:"executor-module-cache-miss"}
--- FAIL: TestSpawnPreseedsRunLocalModuleCacheWithProxyOff (0.01s)
```

The identical test passes against the ordinary host cache in 2.69s. The environmental difference is exact:

- s13's root `go.mod`/`go.sum` contains no `golang.org/x/text`, so the outer executor lawfully does not seed it;
- s14's root manifest requires `golang.org/x/text v0.41.0`, so s14's validating outer cache accidentally contains
  the new executor test's otherwise test-only dependency;
- the ordinary host cache also contains the module, explaining the focused executor green outside dogfood.

Thus the carried R9 bytes are not portable across the named consuming slice closures despite being byte-identical.

## Hold and requested return

No workaround was authored. Adding `x/text` to s13 would invent a dependency contrary to R4's NFC-only condition;
editing the executor would violate the exact carriage and s13's closed SCOPE_DIFF. The three carried paths remain staged
and byte-exact, T4 remains uncommitted, and T5–T15 have not started.

Requested action: route an owner-authored R9 successor that makes executor cache seeding/test execution independent of
the consuming slice's root dependency closure, bank it, and return a new exact-byte carriage. The owner battery must
include a root module with no `x/text` dependency and the dogfood/nested executor path, not only a branch whose product
manifest happens to provide the test module. On that return s13 will replace only the superseded R9 bytes as authorized,
re-verify all carriage hashes, rerun the isolated s8 leg and unchanged full commit gate, then resume T4.

Done: R8/R9 carriage identity verified and consumed; R8 capacity wall discharged; new failure reduced to the inherited
restricted-cache edge with a deterministic direct reproducer; sequence-honest T4 battery appended.

Not done: owner R9 successor; T4 commit; T5–T15; terminal review/restack/merge-gate sequence.

Blocked: the required whole-repository commit gate, solely on the carried executor repair's dependency-closure-sensitive
test/cache behavior.

ACTIONS_GIT_REF: docs-workspace disk action — this append-only SITREP + one live-EOF s13 INDEX row; isolated source worktree remains s13-m10-module@9aee4e957bd4b23a01bb1d0d1afbbd75a9d84554 with exact staged carriage paths and uncommitted T4 in-scope paths; no source commit, merge, push, base move, foreign-worktree mutation, or external side effect
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-impl/SITREP-implementer-20260821-223545.md
