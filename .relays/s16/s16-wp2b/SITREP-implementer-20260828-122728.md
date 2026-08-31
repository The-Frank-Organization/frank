## SITREP - WP2b implementation close at pushed head 4e179fc3; catalog, independent serve gate, and platform release target green

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16-wp2b-impl-close
PARENT_DISPATCH_ID: s16-wp2b-impl
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - commissioned WP2b close report; merge, binding, WP3/E3, and Step-3 ratification remain downstream
GRILL_REQUIRED: no
PLAN_LOCK_ID: s16-wp2b-plan-8 @ sha256 a35b9c25165e6fb9b983dabe4bab221ad1b94ca753fdf388c2434742bf5d3c34
BUNDLE_ID: s16-integration
IN_REPLY_TO: s16-wp2b/IMPL-orchestrator-planner-20260828-114802.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-2.planner, m-10.planner, m-8.planner
SUBJECT: WP2b CLOSE REPORT - plan-8 T1/T2/T3 implemented at 4e179fc3; 7fae5fc1 catalog, independent production serve gate, go1.26.4 deterministic platform release target, 64/0/64, vet, and full suite green; no merge or binding claim
PRE_SCAN_PRESSURE: none

This reports completion under the direct master dispatch and grants no downstream authority.

## Exact branch, scope, and external state

- Worktree `/Users/jack/Programming/harness-s16-integration`; branch `s16-integration`; dispatched base `36dbaca549e3256fcb806ae8a846443e45bb0186`.
- Pushed head `4e179fc3f30ea88fcaad4b310f5bc449f6ea900d`; tree `fefa9d5c3e16d9ebe449d7454de60a4ca464c116`; local and `origin/s16-integration` are exact.
- Commits: `e0feb93d1188adb68bfcb69923f8329952b8667f` implements T1/T2/T3; `4e179fc3f30ea88fcaad4b310f5bc449f6ea900d` corrects the release target to the plan-literal `dist/<goos>-<goarch>/` after pre-close exact-plan reconciliation.
- Draft PR #2 remains draft at the exact head: `https://github.com/The-Frank-Organization/frank-dev/pull/2`; body refreshed with WP2b scope and holds. GitHub reports no checks on the branch; no CI/CD was manually started.
- Exact changed-file set from the dispatched base: `frank/.gitignore`; `frank/cmd/frank-app/main.go`; `frank/cmd/frank-app/main_test.go`; `frank/internal/worker/catalog/catalog.go`; `frank/internal/worker/catalog/catalog_test.go`; `frank/scripts/release-build.sh`; `frank/test/composed/release_build_test.go`.
- `frank/test/seam/**`, `frank/go.mod`, `frank/go.sum`, all other scripts, and every other OUT surface have zero diff. No dependency moved. `dist/**` is ignored and uncommitted.

## T1 - locked-key F58 catalog

- `catalog.Digest` validates and canonical-name-sorts the eight existing identity members, projects them into a dedicated locked wire row with keys `canonical_name`, `form_schema_mapping_version` when present, `tool_impl_catalog_version`, and `tool_schema_digest`, then JCS-hashes that vector.
- The public `catalog.Identity` compatibility shape remains unchanged; this preserves the frozen A14 seam without changing one seam byte while making the F58 digest input exact.
- Recomputed digest is exactly `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`. The literal was not adjusted to an unexpected build result.
- Failing-first capture: the new locked literal test failed against the predecessor `151a7e614abd6b25e643062b26cb9c6af60b0eadedf8e03c1f12b1850458913d`; after the long-key builder landed, the whole catalog package and its literal-pin/recompute tests passed.
- Existing runtime self-check and seam helper continue to consume `catalog.ExpectedDigest` by reference. The catalog package imports no store or render surface.

## T2 - independent production manifest and serve-gate derivations

- `frank-app` now maps the full real `catalog.ExpectedIdentities()` vector into manifest tool members: names, actual schema digests, `m9-catalog-v1`, and relay-only `m2-mapping-v1`.
- The manifest member digest is freshly computed from that identity vector. The production gate comparand is separately sourced from the shipped `catalog.ExpectedDigest` literal through `productionServeGate`; the prior same-variable tautology is gone.
- Failing-first production capture: the predecessor returned names-only digest `3d42ed4c85d906787ad6a55f2e21aee17f99e2d8792423efd4e735a57199141c` instead of the worker pin. The corrected production test proves all eight member values and the exact member digest.
- A forced valid-shape schema divergence at the production composition seam builds a manifest with the divergent member digest while the independently shipped comparand remains pinned; `Gate.Validate` returns `manifest.ErrServeGate`.
- The live presented-surface digest remains only on `logical_surface_digest`; manifest grammar and the `ShippedToolCatalogDigest` name are unchanged.

## T3 - pinned deterministic release pipeline

- New executable `frank/scripts/release-build.sh` compares `go env GOVERSION` to the independent literal `go1.26.4` before any build, then exports `GOFLAGS=""` and reads effective `GOOS`, `GOARCH`, and `CGO_ENABLED`.
- It builds exactly `frank-app`, `frank-broker`, `frank-connector`, and `frank-worker` twice with `go build -trimpath` into two fresh script-owned staging directories and byte-compares corresponding outputs before publication.
- `RELEASE-MANIFEST.json` is emitted as one JCS-ordered JSON object plus exactly one LF, with the closed plan fields, literal `$STAGING/<name>` recipe token, four name-sorted lowercase SHA-256 members, and no timestamp, hostname, absolute path, or invocation-specific byte.
- Publication target is exactly `dist/<goos>-<goarch>/`. An absent target receives the complete verified set; an exact byte-identical five-entry target passes without rewrite; any extra, missing, or mismatched member stops without target deletion or overwrite.
- Failing-first captures: absence of the script failed both executable tests; the initial target-path fixture then failed with missing `dist/linux-amd64/`, exposing that the first implementation used `dist/` directly. The root cause was the omitted plan-literal platform suffix in both script and fixture. The bounded second commit fixed only those two files.
- Fake-toolchain tests prove the exact Linux target/census/manifest bytes, hostile incoming GOFLAGS clearing, byte-identical no-rewrite, mismatched-artifact preservation, and wrong-toolchain refusal before any build.
- Real toolchain proof: `go env GOVERSION` returned exact `go1.26.4`; the current source published `dist/darwin-arm64/` with exactly the four binaries plus `RELEASE-MANIFEST.json`; the immediate repeat reported `existing dist is byte-identical`.
- The earlier direct-`dist/` artifacts produced during the failing implementation cycle were preserved outside the final platform target, ignored, and uncommitted; they were not deleted or overwritten. The final platform target itself has the exact five-entry census.
- New `frank/.gitignore` contains exactly `dist/`. No release binding, binding-readiness, publication beyond the untracked local target, or Q-CARRIER claim is made.

## Regression floor and close gates

- `go test ./...` PASS after each of the two commits' final byte sets; final run includes `test/fixtures` PASS in `251.679s`.
- `go vet ./...` PASS after each commit.
- `go test -tags=seam -json ./test/seam | python3 test/seam/census.py` final: `SUMMARY GREEN=64 RED=0 TOTAL=64`; A14 focused PASS; zero seam diff.
- `go test ./internal/worker/catalog ./cmd/frank-app -count=1` PASS; release composed tests PASS; `bash -n scripts/release-build.sh` PASS; `git diff --check` PASS.
- Implementation worktree is clean apart from ignored, uncommitted `dist/`; branch equals upstream. Governing workspace has no status output immediately pre-submit.

## Honest not-done

No release binding or binding-ready claim; no Q-CARRIER resolution; no external WP3 E3/admissibility/observer/exit; no WP4; no WP5 close or merge brief; no merge, deployment, publication, release, or Step-3 ratification. Merge remains operator-only. This report is ready for the planner's one corroboration pass and upward close.

ACTIONS_GIT_REF: implemented only plan-8 T1/T2/T3 on s16-integration in two regression-green commits through 4e179fc3f30ea88fcaad4b310f5bc449f6ea900d; pushed only the slice branch; refreshed draft PR #2; produced only ignored local dist artifacts; did not edit seam/go.mod/go.sum/master or any other OUT bytes; did not merge, bind, publish externally, deploy, release, or manually start CI/CD
FINAL_GIT_STATUS_SHORT: (literal governing-workspace status immediately pre-submit; implementation branch clean and equal upstream, with ignored dist/ omitted by normal status:)
(no output)
