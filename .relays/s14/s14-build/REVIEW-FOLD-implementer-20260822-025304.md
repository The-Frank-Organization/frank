## REVIEW-FOLD — R14 F7 folded at `7f91f99b`; missing-module naming probe is synthetic and closure-invariant; exact-commit E2 green

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair delta re-verdict and master's retained s13 consumer leg are next; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260822-025304
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260822-024017.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer, m-3.planner, m-7.planner
SUBJECT: F7 R14 test-only successor folded at 7f91f99b — synthetic absent-module probe independent of consuming closure; delta re-verdict requested

Summary: R14 F7 is folded in one liftable one-test-file commit over `ebc9fbe4`. `TestSpawnNamesMissingOfflineModuleWithoutWaitingForNetwork` no longer copies the consuming root's `go.mod`/`go.sum` or imports `golang.org/x/text`. It authors its own synthetic module and sum, imports `example.invalid/frank-missing-module/probe`, sets an empty host module cache, and asserts the retained private diagnostic names the literal fabricated module. Its imports, graph, and expected error shape are now invariant across consuming roots. Accepted executor and fixture bytes are unchanged.

PR: none — this governed lane uses the existing `s14-m8-connector` branch; no push or PR action was authorized or performed.
R14 lock: master ruling `RECONCILE-orchestrator-planner-20260822-023206.md` at SHA-256 `3790ebc8167e1a853e24017684db68f08b2b47b51b05332b2d979654dbde0ef7`; planner fold grant `REVIEW-FOLD-planner-20260822-024017.md` at SHA-256 `0af0306299573f7772c40614e099d18e4cd83a10e1a9679519762d4ec6e7a656`; implementer-owned pre-edit scope `FOLD_SCOPE-implementer-20260822-024149.md` at SHA-256 `7e5bbc7ff8d11fb6847e60a19ff541f07b6da3cfe9c2d0e8a0a2bc12c034e7c2`.
Files changed: `internal/executor/executor_test.go` only; 11 insertions and 4 deletions.

## RED/GREEN trail

- Named mutation: restoring `copyRootModuleClosure` plus a root-owned module import makes the test's retained diagnostic depend on whichever dependency the consuming closure traverses first, rather than the module the probe intends to name.
- RED at unchanged F6 test bytes on the retained s13 consumer reproduction: `go test -mod=readonly -p=1 -count=1 ./internal/executor -run '^TestSpawnNamesMissingOfflineModuleWithoutWaitingForNetwork$' -v` exited 1 in 0.663s. The diagnostic named `github.com/dustin/go-humanize@v1.0.1: module lookup disabled by GOPROXY=off`, not `golang.org/x/text`; this is the exact R14 closure-dependent failure.
- GREEN on the synthetic successor: the same named test exited 0 in 0.563s on s14. The five-test owner cut then exited 0 in 8.615s before commit and 8.493s at the exact commit. The exact-commit nested zero-external-closure owner leg remained green inside that cut.
- Expected value provenance: `example.invalid/frank-missing-module` is a hand-written literal under the reserved non-resolving `.invalid` domain; it is not computed by executor code or copied from a consuming root.

## Per-test consuming-root closure statements at final bytes

- `TestSpawnPassesHostModuleCacheWithOfflineReadonlyFlags`: only the root's own closure, used closure-agnostically. It copies the root manifests solely to prove an already-available closure can execute from the passed host cache; its generated probe imports only the standard library and asserts no dependency identity or diagnostic shape.
- `TestSpawnNamesMissingOfflineModuleWithoutWaitingForNetwork`: nothing. It owns a synthetic `go.mod`, synthetic `go.sum`, fabricated import, empty cache, and literal expected module name; no consuming manifest byte is read.
- `TestGoModuleCachePathHonorsExplicitRunCache`: nothing. It exercises path canonicalization for an explicit temporary `GOMODCACHE` only.
- `TestSpawnNestedExecutorPackageFromZeroExternalClosure`: nothing. It owns a zero-require synthetic root and copies only the in-repository executor fixture/source closure; it does not consume the invoking root's module graph.
- `TestSpawnRetainsPrivateCappedTailOutsideBareFailureVerdict`: nothing. It executes a local shell fixture and checks retained-diagnostic cap, keying, mode, and public verdict shape.

For completeness, the two raw-string `TestOffline` declarations are generated probe payloads rather than package-level tests in `executor_test.go`: the first uses only the copied root closure closure-agnostically and imports standard library only; the second depends on nothing from the consuming root and is intentionally unreachable behind the fabricated missing import.

## Verification at final bytes

- Exact-commit five-test executor owner cut at `7f91f99b` — exit 0, 8.493s; synthetic naming, host-cache flags, explicit cache path, nested zero-external closure, and capped-tail retention all green.
- Exact-commit `go test -mod=readonly -p=1 -count=1 ./test/fixtures -run '^TestS8FXEXE1ExecutorProvidesOnlyRunScopedHandles$' -v` — exit 0, 0.599s.
- Exact-commit `go test -mod=readonly -race -p=1 -count=1 ./internal/executor` — exit 0, 8.854s, no detector finding.
- Pre-commit `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; executor 7.589s, fixtures 201.437s, invariants 1.477s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; executor 7.988s, fixtures 212.247s, invariants 1.369s, all replay/dogfood/zeroloss legs green.
- Exact-commit `gofmt -d internal/executor/executor_test.go`, `go vet ./...`, `go build ./...`, `go mod tidy -diff`, and `git diff HEAD^ --check` — each exit 0; formatter/tidy emitted no diff and module bytes remained unchanged.
- Exact-commit `git diff --name-only HEAD^` names only `frank/internal/executor/executor_test.go`; `git diff --name-only HEAD` and `git status --short` are empty.

Carriage v4 candidate bindings:
- `frank/internal/executor/executor.go` — SHA-256 `05b529fce363efbe6c22b0b164efb423870e1f32a8a1033cd24811a0387cda76` (unchanged).
- `frank/internal/executor/executor_test.go` — SHA-256 `858e28cfd1179bc215a3795d1eabbefaa9fa41a0a56f05e3a8a102a10a0b73cb`.
- `frank/test/fixtures/s8_executor_test.go` — SHA-256 `93f1e94b734b3b35d05cc4fcb856d3a25f1ab151621a2f1327ef569f599be6ab` (unchanged).
- `frank/test/fixtures/s8_exit_gate_test.go` — SHA-256 `c339bc0f2b2467d29b05d1a273acd905cc1ec66b8922d741698ae05d03c1af8f` (unchanged R8 fixture).
- `frank/go.mod` — SHA-256 `02a3d838d2c1f7429a6e0bdb0bd1eee05120a7162cec71f68b14c111df654ce7` (unchanged).
- `frank/go.sum` — SHA-256 `c5601e71a91165dc42d7f1bc40d8a3f7801202f881b5f0d229aca2f01d5ab035` (unchanged).

Acceptance criteria status: the named probe is fully synthetic and self-contained; its imports, module graph, empty cache, and expected module name are independent of the consuming closure; every executor test's root-closure dependency is stated above; one test-only commit cites R14 and the fold grant; the exact-commit E2 battery is green. E2 target met.
Boundary contract proof: `executor.go`, both accepted s8 fixtures, observe schema, retained diagnostic mechanism, public failure-token vocabulary, environment contract, and module manifests are byte-exact. Only the test setup and literal assertion for the R14-defective probe changed.
Evidence levels: E1 exact one-file commit diff and six carriage hashes; E2 retained-consumer RED, local GREEN, owner/nested/s8/race cuts, two full serialized batteries, vet, build, formatting, tidy-diff, and exact clean status. No E3/E4 claim.
Out-of-scope preserved: zero production, fixture, observe, schema, script, connector, dependency/module, generated artifact, merge, publication, deployment, or release byte changed.
Remaining risk: planner delta re-verdict, master's personally run retained s13 consumer leg, carriage v4, s13 exact-byte replacement/commit gate, serialized restack, substantive re-review, the m-3/m-7 windows, and operator MERGE-GATE remain outstanding. No merge, deployment, or live-verification claim is made.

FOLD_SCOPE:
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: `s14-m8-connector@7f91f99b8b8e97e52f1d8115df1e014c5affba9a`; commit `7f91f99b` (`s14 REVIEW-FOLD: isolate missing-module probe (R14 F7)`) changes exactly `frank/internal/executor/executor_test.go` over `ebc9fbe4`; governance actions are the pre-edit FOLD_SCOPE, this report, and their append-only live-EOF s14 INDEX rows; no push, PR, merge, publication, deployment, or release
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@7f91f99b8b8e97e52f1d8115df1e014c5affba9a`
Next requested action: `s14.planner` re-verdicts the F7 delta over exact commit `7f91f99b`; on CLEAN, master reruns the retained s13 consumer leg against these bound bytes and issues carriage v4 only after consumer green. No merge authority is requested or implied.
