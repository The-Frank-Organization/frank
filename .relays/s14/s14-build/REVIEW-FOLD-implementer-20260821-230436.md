## REVIEW-FOLD — R11 F4 folded at `eaf8faa1`; executor owner tests are closure-agnostic and exact-commit E2 is green

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair delta re-verdict and master's carriage v2 are next; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-230436
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-224819.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer
SUBJECT: F4 closure-agnostic executor successor folded at eaf8faa1 — zero-require nested owner battery and exact-commit E2 green; delta re-verdict requested

Summary: R11 F4 is folded in one liftable one-file commit. `TestSpawnPreseedsRunLocalModuleCacheWithProxyOff` no longer invents `golang.org/x/text`; it copies the consuming repository root's actual `go.mod` and optional `go.sum`, then proves that exact closure runs through the executor with proxy, sumdb, workspace, and toolchain escape disabled. A second owner leg stages a root module with zero external requires, copies the repository's internal source surface, and drives the selected R9/R10 executor tests through a nested executor Host while its inherited `GOMODCACHE` is an empty restricted cache. `executor.go` remains byte-identical: R11 §3's inner-seeds-from-outer recursion is correct once the owner tests consume only the root closure, so no production-mechanism change was necessary.

PR: none — this governed lane uses the existing `s14-m8-connector` branch; no push or PR action was authorized or performed.
R11 lock: master ruling `RECONCILE-orchestrator-planner-20260821-223925.md` at SHA-256 `46bd2eb19d2eea25c0f25786ed214db8104fe0a91561139ddce68f0aa25abfc1`; planner fold grant `REVIEW-FOLD-planner-20260821-224819.md` at SHA-256 `5cb1efada0381c49a39b6b64c0b57ad72819cefefe410ef4409639e4e50ae426`; implementer-owned pre-edit scope `FOLD_SCOPE-implementer-20260821-224840.md` at SHA-256 `950d89b3dde2cc84a89d98038562443bef5986ee605c52b261f177b54a10eabf`.
Files changed: `internal/executor/executor_test.go` only. `internal/executor/executor.go` was authorized but required no edit.

## RED/GREEN trail

- **Named mutation:** restoring the hard-coded `x/text v0.41.0` fixture dependency must make the zero-external nested owner leg fail before the nested suite can pass.
- **RED at unchanged `4aea922e` production bytes:** `go test -mod=readonly -run '^TestSpawnNestedExecutorPackageFromZeroExternalClosure$' -count=1 -v ./internal/executor` failed with outer verdict `Outcome:"fail"`, `RungReached:"none"`, `Predicate:"fail"`, `FailingDetail:"suite-exit-mismatch"`. The staged root had a two-line, zero-require `go.mod`, no `go.sum`, and an empty inherited absolute `GOMODCACHE`; the copied executor test still required unseeded `x/text`.
- **Root-cause confirmation:** after the closure fix, the first nested rerun exposed only an incomplete test fixture topology (`observe` imports of `fieldspec` and `record` absent). The retained private diagnostic named those exact missing internal packages. Copying the repository's `internal` source surface while compiling/running only `./internal/executor` completed the dogfood-shaped fixture without adding any module dependency.
- **GREEN:** the direct actual-root-closure preseed leg and the nested zero-external-closure owner leg pass together; the full executor package passes once and at `-count=10` (73.630s). The missing-module refusal, explicit inherited-cache resolution, capped retained diagnostic, and bare `suite-exit-mismatch` token tests remain unchanged and green.

## Verification at final bytes

- Exact-commit owner set (`TestSpawnPreseedsRunLocalModuleCacheWithProxyOff`, zero-external nested leg, missing-module refusal, explicit cache, retained diagnostic/bare token) — exit 0 at `eaf8faa1`, 9.725s.
- Exact-commit `go test -mod=readonly -race -count=1 ./internal/executor` — exit 0, 10.657s, no detector finding.
- Pre-commit `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; executor 7.683s, fixtures 204.139s, invariants 1.443s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go test -mod=readonly -p=1 -count=1 ./...` at `eaf8faa1` — exit 0; executor 7.722s, fixtures 200.430s, invariants 1.372s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go vet ./...`, `go build ./...`, `go mod tidy -diff`, and `git diff HEAD^ --check` — each exit 0; tidy emitted no diff and changed no module byte.
- Exact-commit `git status --short` — empty.

Carriage v2 candidate bindings:
- `frank/internal/executor/executor.go` — SHA-256 `aaa980d445f12e42adc804c2ba724574326bd7ccadee0264137eeacac5466806` (unchanged from R9).
- `frank/internal/executor/executor_test.go` — SHA-256 `bbc9d434c030817fabc9bef26d7ee63c984bbeb3819f939a3b95b59dbd1d742f` (F4 successor).

Acceptance criteria status: a zero-external-require root is exercised through the executor spawn path and its nested executor test suite; the ordinary owner leg derives its module/cache demand solely from the consuming root's actual manifest closure; offline determinism, immediate missing-module refusal, retained diagnostics, and the bare failure token all remain green. E2 target met.
Boundary contract proof: the public executor, observe, verdict, diagnostic, and environment mechanisms are byte-unchanged. The only delta is owner test construction, and the nested fixture compiles real repository internal sources and executes the real Host rather than asserting on mocks or source text.
Evidence levels: E1 exact one-file commit diff plus two carriage hashes; E2 named RED, direct/nested GREEN, count-10 executor stress, race detector, two full serialized batteries, vet, build, and tidy-diff. No E3/E4 claim.
Out-of-scope preserved: zero `executor.go`, observe, schema, script, connector, fixture-corpus, dependency/module, generated artifact, merge, publication, deployment, or release byte changed. `git show --stat eaf8faa1` lists exactly `frank/internal/executor/executor_test.go`.
Remaining risk: planner delta re-verdict, master's independent verification and carriage v2, s13's exact executor-byte replacement and commit gate, serialized restack, substantive re-review, and operator MERGE-GATE remain outstanding. No merge, deployment, or live-verification claim is made.

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: `s14-m8-connector@eaf8faa1b96eae254c6788b9dd49386082a3acd5`; commit `eaf8faa1` (`s14 REVIEW-FOLD: make executor tests closure-agnostic`) cites R11/F4 and changes exactly one in-scope test file; governance actions are the pre-edit FOLD_SCOPE, this report, and their append-only live-EOF s14 INDEX rows; no push, PR, merge, publication, deployment, or release
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@eaf8faa1b96eae254c6788b9dd49386082a3acd5`
Next requested action: `s14.planner` re-verdicts the F4 delta over exact commit `eaf8faa1`; on CLEAN, master independently verifies the closure-agnostic property and issues carriage v2 binding the executor file hashes above. No merge authority is requested or implied.
