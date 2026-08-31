## SITREP - WP2c implementation close at pushed head aea5f064; production release-binding carrier realized and post-realization candidate rebuilt

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16-wp2c-impl-2-close
PARENT_DISPATCH_ID: s16-wp2c-impl-2
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - commissioned WP2c close report; the successor Master+VP binding act, WP3, merge, and Step-3 ratification remain downstream
GRILL_REQUIRED: no
PLAN_LOCK_ID: s16-wp2c-plan-2 @ sha256 7b7018e3b8863955565bb10e8ab0135f6f9fe725ac2a927365bd74524833940c
BUNDLE_ID: s16-integration
IN_REPLY_TO: s16-wp2c/IMPL-orchestrator-planner-20260829-051740.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner
SUBJECT: WP2c CLOSE REPORT - carrier §B production seam and complete §B.7 table realized at aea5f064; authorized recoverable disposal executed once; unchanged go1.26.4 pipeline rebuilt exact-five candidate with covering 96e2cba9; 64/0/64, vet, and full suite green; candidate claimed by no one
PRE_SCAN_PRESSURE: none

This reports completion under the direct master dispatch and grants no downstream authority.

## Exact branch, commit, scope, and review surface

- Worktree `/Users/jack/Programming/harness-s16-integration`; product root `/Users/jack/Programming/harness-s16-integration/frank`; branch `s16-integration`; dispatched base `4e179fc3f30ea88fcaad4b310f5bc449f6ea900d`.
- One realization commit: `aea5f064c7c0d80ac7c81649809ac66c0011382d` (`feat(app): consume release binding manifest`); tree `df495deac8d99a9785670535c731ca9b2699c23f`. Local head and `origin/s16-integration` are exact; normal Git status is empty (the ruled `dist/**` output is ignored).
- Exact changed-file set from the dispatched base: `frank/cmd/frank-app/main.go`; `frank/cmd/frank-app/main_test.go`; `frank/cmd/frank-app/release_binding.go`; `frank/cmd/frank-app/release_binding_test.go`; `frank/test/composed/turn_test.go`.
- Fence proof: `frank/internal/appctl/manifest/**`, `frank/scripts/**`, `frank/test/seam/**`, every other `frank/internal/**` and cmd surface, `frank/go.mod`, `frank/go.sum`, and `master/**` have zero source diff. No dependency moved.
- Draft PR #2 remains DRAFT and OPEN at the exact head: `https://github.com/The-Frank-Organization/frank-dev/pull/2`. Its body is refreshed through WP2c with scope, floor, candidate digest, and the explicit holds. GitHub reports no checks; no CI/CD was manually started.

## T1 - the locked §B consumption seam

- Global `--release-manifest` and `--bound-at-ref` options are presence-tracked separately from their values. Present-together or absent-together is enforced before opening the state store; either half configuration reports typed `half_config` and cannot reach starter construction or admission.
- Both absent call the isolated development population with exact prior values: `BoundAtRef: "working-tree"`, `BuildDigests: nil`, and `ReleaseDigest: 2b14f5b5a76f04d60316f9edca375a5f50c17c17dd2a0f6220314f90e331cc21` (SHA-256 of the unchanged labeled bytes `frank-mvp-development`).
- Both present validate `bound_at_ref` before filesystem access in the total order `empty -> over_length -> embedded_nul -> invalid_utf8 -> dev_sentinel`; the 129-byte overlapping NUL plus invalid-UTF-8 vector reaches exactly `over_length`.
- The production chain is ordered without look-ahead: file read (`manifest_unreadable`); clean symlink resolution and final `dist/<goos>-<goarch>/RELEASE-MANIFEST.json` tuple capture (`path_form_violation`); the one `canonicaljson.Canonicalize` primitive plus exactly one trailing LF (`non_canonical`); the complete eight-member closed schema (`structural_violation` with a distinguishable clause); validated tuple binding (`path_form_violation`); exact-byte SHA-256 and `ReleaseBinding` population.
- The schema fixes the build-command literal, empty `goflags`, `go1.26.4` pin, version equality, nonempty platform strings, exact sorted F63 member set and lowercase 64-hex digests; `cgo_enabled` is string-typed and intentionally value-unvalidated. JSON `null` is explicitly not accepted as a string.
- `productionStarter` receives the parser-produced binding and supplies that exact value to `manifest.Build`; no hand-built bound-path binding is used. The pre-existing `validateRelease` XOR grammar is untouched.

## T2 - complete §B.7 production-path acceptance table

- Failing-first evidence: the paired-option regression initially failed because both flags were unknown and the predecessor emitted the generic state-dir diagnostic. After only the presence gate, both half directions passed. The complete production-path table then failed to compile against the predecessor three-argument starter factory, proving the binding had no population route; the realized four-argument route plus validator made it green.
- Positive published-form fixture admits through `execute`, preserves the valid ref verbatim, and compares `ReleaseDigest` to an independently recomputed SHA-256 over the exact manifest bytes. A symlink alias resolving to the same published-form path also admits.
- Negative path fixtures cover wrong final shape and a `dist/linux-amd64/` directory carrying a darwin/arm64 manifest. Canonical fixtures cover byte perturbation, a top-level duplicate, and a nested member-object duplicate; both duplicates reach only `non_canonical`.
- Closed-schema fixtures independently cover unknown member, missing member, numeric and null wrong types, fixed-literal violation, version/pin equality violation, member-set/order violation, and malformed digest. Ref fixtures cover all five causes plus the overlapping precedence vector. A correctly shaped absent file reaches `manifest_unreadable`.
- Every negative asserts nonzero execution, no starter construction, a zero-value captured binding (no development fallback), and no state-directory creation. The composed whole-app E3 fixture passes both production flags and verifies the persisted frozen manifest carries the exact `bound_at_ref` and freshly recomputed `release_digest`.
- Focused fresh results: `go test ./cmd/frank-app -count=1` PASS; `go test ./test/composed -count=1` PASS; the bound whole-app case PASS.

## T3 - authorized disposal, unchanged rebuild, and candidate bytes

- The required standing label is carried verbatim: the prior release at covering digest `9052ec0f...` and head `4e179fc3...` remains a VALID, REPRODUCIBLE **PRE-REALIZATION RELEASE WITNESS - explicitly NON-BINDING**; it was never promoted and is not the terminal F63 identity.
- Only after T1/T2 were committed at a clean new head, the one granted disposal moved exactly `frank/dist/darwin-arm64/` to the named recoverable directory `/Users/jack/.Trash/frank-s16-wp2c-20260829-054705`. The pre-move target contained exactly `RELEASE-MANIFEST.json`, `frank-app`, `frank-broker`, `frank-connector`, and `frank-worker`; nothing else existed under `dist/`. The post-move `dist/` census was empty. No second disposal occurred.
- `scripts/release-build.sh` was unchanged and returned zero at clean head `aea5f064c7c0d80ac7c81649809ac66c0011382d`, reporting `published verified artifacts to .../dist/darwin-arm64`. Its independent `go env GOVERSION` comparison accepted the locked `go1.26.4` toolchain.
- Final five-entry census, exact: `RELEASE-MANIFEST.json`; `frank-app`; `frank-broker`; `frank-connector`; `frank-worker`.
- Fresh on-disk member SHA-256 values, matching the canonical manifest members: `frank-app` = `708a631e71a240241112d5f048104c65d3529a595939a39d3cc94f891d5b3dbb`; `frank-broker` = `40bc6c8576aecbd2e329f8c8d2963c067344d7b9aa26313e5f06ac7babcf80c0`; `frank-connector` = `1f40bd17cd295c41457fa9bf43ed1e9c0686fc59be64d457c7db2dd0296355fa`; `frank-worker` = `975c19d6028881eb582957c7f54e5e8269e427aea99f29c0014a70b80a7688fb`.
- `RELEASE-MANIFEST.json` is 615 bytes, one canonical JSON line plus exactly one LF. Fresh covering SHA-256 over those exact bytes = `96e2cba98f4847c3c6d7e0ae13192bf69d19e63523ddc3543c6ef8e282772a96`.
- Fresh production-consumer proof: `go run ./cmd/frank-app --state-dir <fresh-temp>/state --release-manifest <absolute rebuilt dist path> --bound-at-ref aea5f064c7c0d80ac7c81649809ac66c0011382d status` returned zero. That drives the one canonical round-trip, closed schema, tuple binding, exact-byte digest, and verbatim final-ref population on the actual rebuilt candidate.
- These are the NEW CANDIDATE BINDING BYTES, CLAIMED BY NO ONE. No s16 act fixes, binds, promotes, or publishes them as a release; the successor Master+VP act owns that decision.

## Regression floor and close gates

- Before the sole commit, one aggregate command retained each return code and returned zero: `go test -tags=seam -json ./test/seam | python3 test/seam/census.py` -> `SUMMARY GREEN=64 RED=0 TOTAL=64`; `go vet ./...` -> PASS/no findings; `go test -count=1 ./...` -> PASS, including `test/composed` in `8.431s` and `test/fixtures` in `255.062s`.
- `git diff --check` PASS. The commit contains no generated `dist/**` byte. The one commit therefore satisfies the floor-at-every-commit rail.
- Source branch is pushed and exact with upstream; PR #2 is refreshed and remains draft; no merge path was invoked.
- Findings: none. The implementation proved no need for a manifest-grammar or other fence amendment.

## Honest not-done and hold

No release binding, binding-ready verdict, Master+VP co-sign, promotion, WP3 plan-2 or implementation, WP4, WP5, merge, deployment, external publication, release, Step-3 exit, or Step-3 ratification. Merge remains operator-only. This close is ready for `s16.planner`'s independent corroboration and upward routing; the pair holds all further product and dist mutation pending fresh authority.

ACTIONS_GIT_REF: realized only plan-2 T1/T2 in one regression-green commit aea5f064c7c0d80ac7c81649809ac66c0011382d on s16-integration; pushed only that slice branch; refreshed draft PR #2; under the renewed one-time grant moved exactly the superseded dist/darwin-arm64 target to named recoverable Trash, then ran the unchanged pinned release pipeline once and derived candidate-only hashes; did not edit manifest/scripts/seam/go.mod/go.sum/master or any other OUT bytes; did not merge, bind, promote, deploy, release, or manually start CI/CD
FINAL_GIT_STATUS_SHORT: (literal governing-workspace status immediately pre-submit; implementation branch clean and equal upstream, with ignored dist/ omitted by normal status:)
(no output)
