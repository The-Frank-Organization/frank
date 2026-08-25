## RESTACK RETURN — s14 is restacked at `d70729a7a528b85b8086bb6a4eb9b256cda7e4cf`; locked gates and carriage-v4 are green; two reviewed manifest bytesets require planner re-review

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s14-build-restack-1
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — planner re-review is next; merge remains separately operator-gated
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260822-161509
IN_REPLY_TO: frank/.relays/s14/s14-build/IMPL-planner-20260822-160513.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s15.planner
SUBJECT: s14 restack complete at d70729a7 — exact conflict census, main-superset manifest review delta, locked E2 gates green, planner re-review requested

Summary:
- Master restack word SHA-256: `c36a9a79fd2fe36d997e4e516a8d4832eb1466e1dcfde88bbeed28fccff6f6ca`; pair dispatch SHA-256: `f8414278e9c9f24ca4110248d57d44f76339f36434748ff3b0b21a9b45a91cd5`.
- Pre-restack branch tip: `7f91f99b8b8e97e52f1d8115df1e014c5affba9a`.
- Execution-time target-main pin: `60fa348d01d5c66d93c5c42428f53c0013e0f9b2`, the s13 merge receipt commit named by the word.
- Restacked branch tip: `d70729a7a528b85b8086bb6a4eb9b256cda7e4cf` on `s14-m8-connector`; the target pin is its merge-base and ancestor.
- Twenty commits remain above target-main. Six of the former 26 replays became empty because main already carried the ruled final repair bytes.

## Conflict census

- Replay 2/26, `REBASE_HEAD` `7e7d4f2ff078a71c033b09877b38d11473deee15`: exactly `frank/go.mod` (content) and `frank/go.sum` (add/add). For `go.mod`, stage 2/main SHA-256 was `29d34d18f193269fc49155db630ce7f6391a1b155c928b6e22b2ef34d01d5b50`; stage 3/s14 was `02a3d838d2c1f7429a6e0bdb0bd1eee05120a7162cec71f68b14c111df654ce7`. For `go.sum`, stage 2/main was `5b5228f6fabd5ccf55f2d77b4c9e2aed8e6f3e12801bb9c35f29e668cf6042c6`; stage 3/s14 was `c5601e71a91165dc42d7f1bc40d8a3f7801202f881b5f0d229aca2f01d5ab035`. Both resolved to exact target-main bytes as §2b requires. The resulting T1b replay kept its two NFC source/test files and inherited the main manifest superset.
- Replay 12/26, `REBASE_HEAD` `b86b8bc17ba56302d33b7aba4cf3f8af404d70ea`: `executor.go` and `executor_test.go`. Stage 2 was exact carriage-v4 and exact pre-restack-tip identity (`05b529fc…`, `858e28cf…`); stage 3 was the superseded R9/R10 intermediate (`aaa980d4…`, `02f244f5…`). Both resolved to the identical stage-2/main blobs; the replay became empty. Replay 13/26 R8 was then automatically dropped as already upstream.
- Replay 23/26, `REBASE_HEAD` `eaf8faa1b96eae254c6788b9dd49386082a3acd5`: `executor_test.go` only. Stage 2 was exact carriage-v4/pre-restack identity `858e28cf…`; stage 3 was the superseded F4 intermediate `bbc9d434…`. It resolved to the identical stage-2/main blob and the replay became empty.
- Replay 24/26, `REBASE_HEAD` `df26d6122f4eaf93e0ea66d04753c24e69b0b29e`: `executor.go` and `executor_test.go`. Stage 2 was exact carriage-v4/pre-restack identity (`05b529fc…`, `858e28cf…`); stage 3 was the superseded R12 intermediate (`8e5751ff…`, `9a56769f…`). Both resolved to exact stage-2/main blobs and the replay became empty. Replays 25/26 R13 and 26/26 R14 F7 were automatically dropped as already upstream.
- Replays 1, 3–11, 14–22 were conflict-free. No other source conflict occurred. `frank/.relays/s14/**` inherited the target-main banked bytes with no branch-side committed divergence; `frank/.relays/s13/**` and the s13 sprint tree inherited main silently; no s14 sprint-doc conflict existed, so the R15 class did not fire.

The shared-repair pauses were intermediate-history conflicts rather than final-byte non-identity: at every pause, stage 2 equaled both target-main and the reviewed pre-restack s14 tip byte-for-byte. No source content was synthesized or semantically merged during resolution.

## Reviewed-byte delta and required re-review

`git diff --exit-code 7f91f99b..d70729a7 -- cmd/frank-connector internal/connector internal/executor test/fixtures/s8_executor_test.go test/fixtures/s8_exit_gate_test.go` exited 0. Every reviewed connector byte and all four shared repair files are unchanged.

Exactly two s14-reviewed paths changed relative to the pre-restack tip: `frank/go.mod` and `frank/go.sum`. They are target-main's explicitly ruled superset: x/text remains pinned at v0.41.0, `go 1.25.0` remains exact, and main adds s13's `modernc.org/sqlite v1.57.0` closure. The delta is 66 insertions and one replacement across the two manifests. These are expected §2b resolution bytes, but they are still changed reviewed bytes; planner substantive re-review is requested rather than claiming “none.”

## Appipc-adoption evaluation

No adoption is warranted in this restack. Both packages use the same four-byte big-endian length prefix and 4 MiB ceiling, but `internal/appipc` exposes only raw `ReadFrame`/`WriteFrame` payload I/O and its own error vocabulary. `internal/connector/frame` additionally owns canonical JSON envelope encoding, connector channel/type validation, counter grammar, reply requirements, and per-channel monotonic sequence state, and its `Encode` API returns governed wire bytes rather than writing them. Importing the raw appipc primitive would not replace the connector codec; it would add a cross-slice package dependency and require reviewed error/API changes for no observed correctness gain. The full composed tree is green with both implementations. Any later common-primitive extraction remains serialized seam arbitration/s16 work, not a silent s14 restack edit.

## Gates at `d70729a7a528b85b8086bb6a4eb9b256cda7e4cf`

- `go test -mod=readonly -count=100 ./internal/connector/transport` — exit 0, 32.507s.
- Executor five-test owner cut — exit 0, 9.069s; host-cache flags, synthetic missing-module naming, explicit cache path, nested zero-external closure, and retained capped tail all green.
- `go test -mod=readonly -race -p=1 -count=1 ./internal/executor` — exit 0, 9.137s, no detector finding.
- `go test -mod=readonly -race -count=20 ./internal/connector/control ./internal/connector/attempt ./internal/connector/transport ./internal/connector/service` — exit 0; control 3.622s, attempt 1.468s, transport 8.563s, service 1.740s; no detector finding.
- `go test -mod=readonly -p=1 -count=1 ./...` — exit 0; executor 7.842s, fixtures 244.216s, invariants 1.733s, all appctl/appipc/connector/replay/dogfood/zeroloss legs green.
- Chained `go vet ./...` — exit 0, silent.
- Fresh final `go mod verify` — `all modules verified`; `go mod tidy -diff` — exit 0, silent.
- Fresh final carriage hash check and literal `git status --short` — hashes exact; status empty.

Carriage-v4 at the restacked tip:
- `internal/executor/executor.go` = `05b529fce363efbe6c22b0b164efb423870e1f32a8a1033cd24811a0387cda76`.
- `internal/executor/executor_test.go` = `858e28cfd1179bc215a3795d1eabbefaa9fa41a0a56f05e3a8a102a10a0b73cb`.
- `test/fixtures/s8_exit_gate_test.go` = `c339bc0f2b2467d29b05d1a273acd905cc1ec66b8922d741698ae05d03c1af8f`.
- `test/fixtures/s8_executor_test.go` = `93f1e94b734b3b35d05cc4fcb856d3a25f1ab151621a2f1327ef569f599be6ab`.

Evidence levels: E1 exact staged-side hashes, conflict census, graph/ancestor proof, reviewed-surface diff, and carriage hashes; E2 full uncached battery, vet, count-100 transport, executor owner/race, connector race-20, module verification, and tidy stability. No E3/E4 claim.

Out-of-scope preserved: no new source/test byte was authored during the restack; no appipc adoption, shared-seam edit, dependency edit, merge, push, PR, deployment, release, or live verification occurred. Per the finishing workflow and explicit no-merge boundary, branch `s14-m8-connector` and its linked worktree remain preserved.

ACTIONS_GIT_REF: `s14-m8-connector` restacked from `7f91f99b8b8e97e52f1d8115df1e014c5affba9a` onto execution-time target `60fa348d01d5c66d93c5c42428f53c0013e0f9b2`; returned tip `d70729a7a528b85b8086bb6a4eb9b256cda7e4cf`; source worktree clean; no merge or push
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `d70729a7a528b85b8086bb6a4eb9b256cda7e4cf`

Next requested action: `s14.planner` performs the required restack re-review, explicitly including the exact main-superset `go.mod`/`go.sum` delta and the unchanged connector/carriage proof, then returns the reviewed tip to `master.orchestrator-planner`. Master, not this seat, authors the operator MERGE-GATE brief.
