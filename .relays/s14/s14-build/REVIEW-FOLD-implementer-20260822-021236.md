## REVIEW-FOLD — R13 F6 folded at `ebc9fbe4`; seeder deleted, host cache passed through, exact-commit E2 green

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair delta re-verdict and master's independent non-symlinked s13 consumer leg are next; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260822-021236
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260822-015347.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer, m-3.planner, m-7.planner
SUBJECT: F6 R13 successor folded at ebc9fbe4 — seeding surface deleted, resolved host cache passed through under offline readonly discipline; delta re-verdict requested

Summary: R13 F6 is folded in one liftable three-file commit over `df26d612`. The executor resolves the host module cache once per run and passes it directly as `GOMODCACHE`; `GOPROXY=off`, `GOSUMDB=off`, `GOFLAGS=-mod=readonly`, `GOTOOLCHAIN=local`, and `GOWORK=off` remain fixed, while `HOME`, `TMPDIR`, `GOCACHE`, and `GOPATH` remain run-local. The entire hand-rolled seeding surface is deleted: the module-download type; sum reader/selector; discovery, seed, path-copy, and file-copy helpers; the run-local module-cache directory; and the two seeding-specific owner legs. The replacement real-spawn owner test and the R13-granted s8 fixture cut assert the new environment contract. The genuinely-needed missing-module test still fails locally through Go's own named error in the retained private diagnostic while the public verdict remains bare `suite-exit-mismatch`.

PR: none — this governed lane uses the existing `s14-m8-connector` branch; no push or PR action was authorized or performed.
R13 lock: master ruling `RECONCILE-orchestrator-planner-20260822-014814.md` at SHA-256 `657cc72d2f7b28977ba657ef6cf7c1415e37112c61a323b1b85c2487afedf651`; planner fold grant `REVIEW-FOLD-planner-20260822-015347.md` at SHA-256 `8a5f9d355da7e6e1ab28ba75c86748b4076ca3d2bb0a4821d0467c903ba480b7`; implementer-owned pre-edit scope `FOLD_SCOPE-implementer-20260822-015608.md` at SHA-256 `37f3670e5645f359b82556d339ee95c332cfd0bbfdbcbdf1ffc3ecd59d112b1a`.
Files changed: `internal/executor/executor.go`, `internal/executor/executor_test.go`, and the one-file R13 widening `test/fixtures/s8_executor_test.go` only.

## RED/GREEN trail

- **Named mutation:** restoring the old `$PWD/.cache/go-mod` environment must make both the real-spawn owner leg and s8 isolation fixture fail their host-cache passthrough assertion.
- **RED at unchanged `df26d612` production bytes:** after re-cutting tests first, `go test -mod=readonly -p=1 -count=1 ./internal/executor ./test/fixtures -run '^(TestSpawnPassesHostModuleCacheWithOfflineReadonlyFlags|TestS8FXEXE1ExecutorProvidesOnlyRunScopedHandles)$'` exited 1. The owner leg failed at 0.20s and s8 fixture at 0.21s; both observed `Outcome:"fail"`, `Predicate:"fail"`, bare `FailingDetail:"suite-exit-mismatch"` because the current executor supplied its run-local module cache instead of the expected host cache.
- **GREEN fixture normalization:** the first compiled production run exposed that `goModuleCachePath` intentionally canonicalizes symlinks while macOS temp paths may be presented through `/var`. Both test expectations were corrected with `filepath.EvalSymlinks`; this does not weaken the RED because the old implementation still supplies `$PWD/.cache/go-mod`.
- **GREEN:** the same two-test command exited 0 after host-cache passthrough, 2.491s for executor and 0.667s for fixtures. Exact-commit reruns remain green.
- **Pure deletion:** no synthetic test was added for deleted private helpers. The commit-level symbol tripwire proves all named seeding types/functions and `.cache/go-mod` references are absent from the three governed files.
- **Standing cuts:** the full executor package includes the needed-module immediate/name negative, F4 nested closure leg re-shaped around the passthrough owner test, explicit cache-path resolution, capped retained diagnostic, replay stability, R10 forced-red retention, and bare-token assertions; all remain green.

## Exact defect-source correction

The R13/F6 simplicity ruling and deletion remedy stand, but their exact `df26d612` failure-source wording is not byte-accurate. At those bytes, `cachedModuleDownload` overwrites the absent JSON `Zip` field with `base + ".zip"` before returning. The value reaching the copy list is therefore non-empty, so an extraction-only module fails first at `os.Stat` of the missing synthesized `.zip` path; execution does not reach a literal junk path `"hash"`. This correction was recorded before source edit in the FOLD_SCOPE artifact. It does not alter the empirically validated passthrough shape or its acceptance bar.

## Verification at final bytes

- Exact-commit `go test -mod=readonly -p=1 -count=1 ./internal/executor` at `ebc9fbe4` — exit 0, 10.348s.
- Exact-commit `go test -race -mod=readonly -p=1 -count=1 ./internal/executor` — exit 0, 11.293s, no detector finding.
- Exact-commit `go test -mod=readonly -p=1 -count=1 ./test/fixtures -run '^TestS8FXEXE1ExecutorProvidesOnlyRunScopedHandles$'` — exit 0, 0.706s.
- Pre-commit `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; executor 7.536s, fixtures 203.282s, invariants 1.434s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go test -mod=readonly -p=1 -count=1 ./...` at `ebc9fbe4` — exit 0; executor 8.446s, fixtures 199.936s, invariants 1.508s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go vet ./...`, `go build ./...`, `go mod tidy -diff`, and `git diff HEAD^ --check` — each exit 0; tidy emitted no diff and changed no module byte.
- Exact-commit deletion tripwire over the three governed files for `seedGoModuleCache`, `readGoSums`, `goSumModules`, `cachedModuleDownload`, `seedModuleDownload`, `copyModuleCachePath`, `copyModuleCacheFile`, `moduleDownload`, and `.cache/go-mod` — no match, exit 0 under the inverted grep gate.
- Exact-commit `git status --short` — empty.

Carriage v4 candidate bindings:
- `frank/internal/executor/executor.go` — SHA-256 `05b529fce363efbe6c22b0b164efb423870e1f32a8a1033cd24811a0387cda76`.
- `frank/internal/executor/executor_test.go` — SHA-256 `05ab2025e302d6c54b40e6e0443f4d8555bdd42d3d9102a75d0f9d04aeaf9048`.
- `frank/test/fixtures/s8_executor_test.go` — SHA-256 `93f1e94b734b3b35d05cc4fcb856d3a25f1ab151621a2f1327ef569f599be6ab`.
- `frank/go.mod` — SHA-256 `02a3d838d2c1f7429a6e0bdb0bd1eee05120a7162cec71f68b14c111df654ce7` (unchanged).
- `frank/go.sum` — SHA-256 `c5601e71a91165dc42d7f1bc40d8a3f7801202f881b5f0d229aca2f01d5ab035` (unchanged).

Acceptance criteria status: seeding and run-local module-cache bytes are deleted; the resolved host cache is passed through with all five offline/readonly flags; build/cache/temp/home paths remain run-local; missing required modules fail immediately and named through retained diagnostics; the s8 row-14 assertion pins the new contract; standing executor and full E2 batteries remain green. E2 target met.
Boundary contract proof: module source is now shared read-only-by-discipline, not OS-enforced read-only. With proxy disabled, the suite cannot acquire a new module, but Go may write advisory locks or extract an already-present zip in its own global content-addressed cache; that honest residue is the R13-accepted boundary. Build artifacts and process-private paths remain run-local. Observe verdict schema, retained diagnostic cap/mode/keying, and public failure-token vocabulary are unchanged.
Evidence levels: E1 exact three-file commit diff, deletion tripwire, and five carriage hashes; E2 named RED/GREEN, executor package/race, s8 cut, two full serialized batteries, vet, build, tidy-diff, and exact clean status. No E3/E4 claim.
Out-of-scope preserved: zero observe, schema, script, connector, other fixture-corpus, dependency/module, generated artifact, merge, publication, deployment, or release byte changed. `git diff-tree --no-commit-id --name-only -r ebc9fbe4` lists exactly the three granted paths.
Remaining risk: planner delta re-verdict, master's personally run non-symlinked s13 isolated s8 consumer leg, carriage v4, s13 exact-byte replacement/commit gate, serialized restack, substantive re-review, the m-3/m-7 objection windows over the deliberate hermeticity re-cut, and operator MERGE-GATE remain outstanding. No merge, deployment, or live-verification claim is made.

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/test/fixtures/s8_executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: `s14-m8-connector@ebc9fbe4d8eadfb81d8d8e9c9fa7ee0b58c30d86`; commit `ebc9fbe4` (`s14 REVIEW-FOLD: pass through host module cache (R13)`) changes exactly the three in-scope files over `df26d612`; governance actions are the pre-edit FOLD_SCOPE, this report, and their append-only live-EOF s14 INDEX rows; no push, PR, merge, publication, deployment, or release
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@ebc9fbe4d8eadfb81d8d8e9c9fa7ee0b58c30d86`
Next requested action: `s14.planner` re-verdicts the F6 delta over exact commit `ebc9fbe4`; on CLEAN, master personally runs the non-symlinked s13 isolated s8 consumer leg against these bound bytes and issues carriage v4 only after consumer green. No merge authority is requested or implied.
