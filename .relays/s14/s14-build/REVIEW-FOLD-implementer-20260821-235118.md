## REVIEW-FOLD — R12 F5 folded at `df26d612`; offline precondition is actual build/run need and exact-commit E2 is green

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair delta re-verdict and master's independent s13 consumer leg are next; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-235118
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-233412.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer
SUBJECT: F5 R12 successor folded at df26d612 — unneeded absent artifacts are a non-event, needed absence remains immediate and named; delta re-verdict requested

Summary: R12 F5 is folded in one liftable two-file commit over `eaf8faa1`. The seeder no longer runs one fatal all-go.sum download batch. It probes each checksum-listed artifact read-only with `go list -m -json module@version` against the explicitly resolved host cache under `GOPROXY=off`, copies only entries accepted by the existing path containment and go.sum hash checks, and skips unavailable entries. The staged suite then builds/runs against only its run-local seeded cache with `GOPROXY=off`; therefore an absent graph-only artifact is not a precondition, while an actually imported absent artifact fails locally through the suite, is named in the retained private diagnostic, and leaves only the existing bare `suite-exit-mismatch` public token. No cache warming command exists in mechanism or tests.

PR: none — this governed lane uses the existing `s14-m8-connector` branch; no push or PR action was authorized or performed.
R12 lock: master ruling `RECONCILE-orchestrator-planner-20260821-233114.md` at SHA-256 `ce5c33db8441860cd8c2aca4b201dd563b2f72b3857e645050729ffa38958661`; planner fold grant `REVIEW-FOLD-planner-20260821-233412.md` at SHA-256 `03fa14f0fefff940ce26a84e2e75fd12258efafa98841aeb728edd6b64589a9e`; implementer-owned pre-edit scope `FOLD_SCOPE-implementer-20260821-233721.md` at SHA-256 `47672dc499ac285ac0e5fdc96512ec804ad2fa405626bc32ba9366c434bd56ea`.
Files changed: `internal/executor/executor.go` and `internal/executor/executor_test.go` only.

## RED/GREEN trail

- **Named mutation for the positive leg:** restoring a fatal requirement on every non-`/go.mod` go.sum artifact must make a consuming root fail before its standard-library-only staged package runs.
- **Positive RED at unchanged `eaf8faa1` production bytes:** `TestSpawnIgnoresAbsentUnneededGoSumArtifacts` copied the root's valid `go.mod`/`go.sum`, pointed inherited `GOMODCACHE` at a deliberately empty fixture cache, and invoked the real spawn path on a standard-library-only package. The test failed in 0.01s with `Outcome:"unsafe"`, `Predicate:"blocked"`, `FailingDetail:"executor-module-cache-miss"` because the current batch demanded the unused `golang.org/x/text` artifact.
- **Named mutation for the negative leg:** skipping every absent artifact without leaving the staged build offline must let an actually imported missing module pass or lose its name.
- **Negative RED at unchanged `eaf8faa1` production bytes:** the standing missing-module test was changed to copy the same valid closure, import `golang.org/x/text/language`, use a deliberately empty inherited cache, and run real `go test` through spawn. It failed against the old production bytes with the pre-spawn `executor-module-cache-miss` verdict instead of the required retained named suite failure.
- **GREEN:** the two real-spawn R12 legs pass together. The unneeded-artifact leg reaches E2/pass with the fixture cache still empty at entry; the needed-artifact leg returns `Outcome:"fail"`, `RungReached:"none"`, `Predicate:"fail"`, bare `FailingDetail:"suite-exit-mismatch"`, completes under the existing two-second immediate bound, and its private diagnostic contains `golang.org/x/text`.
- **Standing cuts:** the complete executor package includes the F4 nested zero-external-closure leg, ordinary offline preseed leg, explicit inherited-cache path, capped retained diagnostic, replay stability, and bare-token assertions; all remain green.

## Verification at final bytes

- Exact-commit `go test -mod=readonly -p=1 -count=1 ./internal/executor` at `df26d612` — exit 0, 13.477s.
- Exact-commit `go test -race -mod=readonly -p=1 -count=1 ./internal/executor` — exit 0, 14.388s, no detector finding.
- Pre-commit `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; executor 10.791s, fixtures 213.034s, invariants 1.436s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go test -mod=readonly -p=1 -count=1 ./...` at `df26d612` — exit 0; executor 10.492s, fixtures 208.454s, invariants 1.540s, all replay/dogfood/zeroloss legs green.
- Exact-commit `go vet ./...`, `go build ./...`, `go mod tidy -diff`, and `git diff HEAD^ --check` — each exit 0; tidy emitted no diff and changed no module byte.
- Exact-commit warming tripwire `! rg -n 'go mod download' internal/executor/executor.go internal/executor/executor_test.go` — exit 0.
- Exact-commit `git status --short` — empty.

Carriage v3 candidate bindings:
- `frank/internal/executor/executor.go` — SHA-256 `8e5751ff44c875b25136a87981e52be2c091ae9746b9fdb1b68e712a0775ae07`.
- `frank/internal/executor/executor_test.go` — SHA-256 `9a56769feb27e166e45a21d89dce35401a1ba5f9941735662fa687166fbb7443`.
- `frank/go.mod` — SHA-256 `02a3d838d2c1f7429a6e0bdb0bd1eee05120a7162cec71f68b14c111df654ce7` (unchanged).
- `frank/go.sum` — SHA-256 `c5601e71a91165dc42d7f1bc40d8a3f7801202f881b5f0d229aca2f01d5ab035` (unchanged).

Acceptance criteria status: a real consuming root with valid but deliberately unavailable unneeded checksum artifacts reaches E2/pass; an actually imported unavailable module fails locally under the immediate bound and is named in retained diagnostics; the public failure detail remains the bare existing token; F4 closure independence, offline determinism, and full E2 gates remain green. E2 target met.
Boundary contract proof: the delta changes only host-cache discovery and owner tests. It preserves the existing `seedModuleDownload` go.sum equality checks, `copyModuleCachePath` containment checks, run-local cache, proxy-off environment, retained-diagnostic cap/mode/keying, observe verdict schema, and public token vocabulary. Tests exercise the real Host and real Go build rather than mocks or source-text assertions.
Evidence levels: E1 exact two-file commit diff plus four carriage hashes; E2 named REDs, real-spawn GREENs, executor race, two full serialized batteries, vet, build, tidy-diff, warming tripwire, and exact clean status. No E3/E4 claim.
Out-of-scope preserved: zero observe, schema, script, connector, fixture-corpus, dependency/module, generated artifact, merge, publication, deployment, or release byte changed. `git diff-tree --no-commit-id --name-only -r df26d612` lists exactly the two executor paths.
Remaining risk: planner delta re-verdict, master's personally run s13 isolated s8 consumer leg, carriage v3, s13 exact-byte replacement/commit gate, serialized restack, substantive re-review, and operator MERGE-GATE remain outstanding. No merge, deployment, or live-verification claim is made.

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: `s14-m8-connector@df26d6122f4eaf93e0ea66d04753c24e69b0b29e`; commit `df26d612` (`s14 REVIEW-FOLD: narrow offline seeding to build need (R12)`) changes exactly the two in-scope executor files over `eaf8faa1`; governance actions are the pre-edit FOLD_SCOPE, this report, and their append-only live-EOF s14 INDEX rows; no push, PR, merge, publication, deployment, or release
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@df26d6122f4eaf93e0ea66d04753c24e69b0b29e`
Next requested action: `s14.planner` re-verdicts the F5 delta over exact commit `df26d612`; on CLEAN, master personally runs the s13 isolated s8 leg against these bound bytes and issues carriage v3. No merge authority is requested or implied.
