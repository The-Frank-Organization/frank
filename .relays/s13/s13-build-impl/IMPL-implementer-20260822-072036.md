## IMPL — s13 m-10 supervisor-module implementation report: **T0–T15 COMPLETE at component E2** on clean `s13-m10-module@f090868fb28347de93464fd56df4514928f479cd`; exact final-head ordinary + reduced + race + compile-matrix gates green; no push, PR, restack, merge, E3, deployment, or release claimed

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s13-build-implementer-report
PARENT_DISPATCH_ID: s13-build-impl-2
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the Planner's one adversarial end-review is next; serialized restack/rerun/re-review and the operator's terminal MERGE-GATE remain separate later acts, with s13 preceding s14 then s15
FILED_AT_LOCAL: 20260822-072036
IN_REPLY_TO: frank/.relays/s13/s13-build-impl/IMPL-planner-20260821-130456.md
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer
SUBJECT: terminal implementation report — s13 m-10 module T0–T15 complete at E2 on clean f090868f; planner end-review requested; merge remains gated

Summary:
- Executed the locked sixteen-task plan T0–T15 straight-through on isolated branch `s13-m10-module`, launched from exact `main@b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9`, now clean at `f090868fb28347de93464fd56df4514928f479cd`.
- Built the app IPC codec/type/limit layer, private SQLite store, serialized applier, frozen manifest gate, F59 authority host, lifecycle/epoch retirement, broker control client, scheduler, settlement/evidence carriage, E0 persistence, bounded terminal surface, and boot recovery/negative sweep.
- Carried the four R8–R14 executor/fixture files at master's exact carriage-v4 hashes before T4 and retained those hashes through the terminal head.
- Every source/test commit ran `go build ./... && go test ./... && go vet ./...`; all T1–T15 battery records are under `frank/.relays/s13/batteries/`.
- No GitHub CI/CD was run. No PR was opened and nothing was pushed, restacked, merged, deployed, published, or released.

PR: none — local concurrent-slice branch only; no publication or merge authority inferred.

Plan and authority locks:
- Design artifact: `frank/docs/sprints/active/2026-08-20-s13-m10-module/designs/DS-s13-m10-module-20260820.md`, SHA-256 `3e74c4c1fe76dbaf07f3ca9547ffeb88a958cf23d0f6180afa9d49cd71a81361`.
- Plan artifact: `frank/docs/sprints/active/2026-08-20-s13-m10-module/plans/PL-s13-build-plan-20260820.md`, SHA-256 `00b1f60368d2e227f48a1bf22d8cb0e1ebeefb800bcc316aaab1bc2e5c9ad161`.
- Approving plan review: `s13-build-plan-review-6`, SHA-256 `6e3bb7e34ace4e9fa8b1cd5b3335264c752918d583d26e7bba3d9a73ce12e80c`.
- Live dispatch: this report's `IN_REPLY_TO`, SHA-256 `8c6fca228fe36fbb59a8fd574bd475f00c36dfd75a14a3c8f199dd79e5f33429`.
- Final carriage-v4 authority: `master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260822-041652.md`, SHA-256 `b89a96e9c5e4320847e167c71ee17e75ece153721e21294f333f2e7a99f9f3e7`.

Branch and commits, in order:
- Base: `b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9`.
- `6a5930da` governed input bank (T0) · `f2342906` appipc core (T1) · `e0be7420` message families/fakes (T2) · `9aee4e95` limits/compile matrix (T3) · `690193ce` exact R8–R14 carriage · `3402276d` SQLite store (T4) · `a19bfe74` applier (T5) · `95e0b370` manifest (T6) · `f488fb77` F59 (T7) · `67a3e60d` F59 correction · `b1d3ac59` supervision/retirement (T8) · `98eccf77` broker client (T9) · `7e44bfd6` scheduler (T10) · `5c6aa2f3` settlement/carriage (T11) · `4b4aa034` E0 (T12) · `d5983356` terminal (T13) · `3bc9b5ff` recovery/negatives (T14) · `f090868f` completion evidence (T15).
- Head: `f090868fb28347de93464fd56df4514928f479cd`.

Implemented contract by task family:
- **T1–T3 appipc:** 4-byte big-endian bounded frames; canonical decimal full-`uint64` counters; strict JCS subset; typed CTRL-W/CTRL-C/broker families; real socketpair fakes; production/reduced complementary limit selectors; executable compile-negative P2 witness.
- **T4–T5 durability:** SQLite v1 schema with WAL + `synchronous=FULL`, 0700/0600 privacy, integrity/version refusal, forward-only genesis migration, one connection; one serialized write/query loop; commit-before-emission/reply; one transaction and one `state_seq` advance per mutation; no other `database/sql` import.
- **T6 manifest:** immutable canonical bytes + digest, exact operator-ratified 8-name tool set, policy/lane/digest equality, workspace-root identity, release binding, frozen-row serve gate and vector-drift refusal.
- **T7 F59:** ordered authority checks, canonical effect descriptors, row-first one-shot tickets, consume identity match, stale/future fencing, outcome totality, replay/idempotence, state-sensitive expiry/parking, exact 64-bit ceiling behavior.
- **T8 supervision:** sanitized child environment, close-on-exec socketpair/death-pipe construction, process-group spawn, health/deadline/termination ladder, lease retirement with E+1 fencing and full parking, same-epoch pre-lease washout, distinct parked-cap and tenth-failure terminals, durable backoff/reset/cancel.
- **T9 broker:** exclusive fcntl control lock, durable `control_generation` advance before Unix dial, authenticated handshake, detached spawn command, total state-proposal fold, two-form assign gate, durable broker-event ack/dedup, framed proposal/install/deadline re-proposal.
- **T10 scheduler:** atomic admission predicate, turn/lease/wake/disclosure commit, actual encoded-size bounds, rowless initial overflow, continuation terminal triple, attempt-open row-before-ack, typed cancellation, terminal and duplicate-wake handling.
- **T11 settlement/carriage:** canonical unknown/partial producer rows with ancestry, telemetry-never-input cut evidence, reduced S-1 receipt tuple, continuation snapshot/disposition, cap behavior, B/E three-digest first-commit-wins carriage with `m10_row_state`.
- **T12 E0:** closed v1/v2 m-3 app-event validation, duplicate-key rejection, exact raw-byte persistence, exact duplicate idempotence, identity conflict refusal, no silent drop; `workspace_root_path` and `session_log_path` rejected.
- **T13 terminal:** registry-exact mutations `run start`/`run stop`/`run cancel [--hard]`; read-only committed snapshots for `status`/`attempts`/`tickets`/`parked`/`wakes`; ids/states/digests only; forbidden mutations refuse without `state_seq` motion; persistent FAILED banner and non-zero scripted exits; recovery hook always precedes command dispatch.
- **T14 recovery/negatives:** durable matrix (a) leased retirement/one mint, (b) committed retirement/no re-mint, (c) pre-lease washout/same epoch, (d) initial E=1/no mint; terminal no-restart; control establishment before reconstruction; always-propose current durable tuple; bounded retry and fail-closed mismatch; no conductor authority dependency, provider-payload type, credential-byte column, or sensitive validation echo.

Carriage-v4 exact-byte proof at terminal head:
- `internal/executor/executor.go` = `05b529fce363efbe6c22b0b164efb423870e1f32a8a1033cd24811a0387cda76`.
- `internal/executor/executor_test.go` = `858e28cfd1179bc215a3795d1eabbefaa9fa41a0a56f05e3a8a102a10a0b73cb`.
- `test/fixtures/s8_exit_gate_test.go` = `c339bc0f2b2467d29b05d1a273acd905cc1ec66b8922d741698ae05d03c1af8f`.
- `test/fixtures/s8_executor_test.go` = `93f1e94b734b3b35d05cc4fcb856d3a25f1ab151621a2f1327ef569f599be6ab`.

Dependency and scope proof:
- `go.mod` carries exact `go 1.25.0`, no `toolchain` directive, direct `modernc.org/sqlite v1.57.0` and `golang.org/x/text v0.41.0`; `go mod verify` = `all modules verified`; `go mod tidy -diff` = no output/no drift.
- Path-classification over every `git diff --name-only b7f406b2..f090868f` entry = PASS: zero path outside `frank/cmd/frank-app/**`, `frank/internal/appctl/**`, `frank/internal/appipc/**`, ruled `go.mod`/`go.sum`, s13 sprint/relay evidence, and the four carriage-v4 exceptions.
- Total branch delta: 122 files, 15,702 insertions, 22 deletions (governed input/relay/battery files included).

Verification at terminal bytes:
- `go test -run '^TestLimitsCompileMatrix$' -count=1 ./internal/appipc` — PASS (`0.468s`); ordinary build + reduced-selector positive + named compile-negative failure all asserted.
- `go test -race -count=3 ./internal/appctl/... ./internal/appipc/... ./cmd/frank-app` — PASS; selected: recovery `3.761s`, scheduler `4.614s`, supervisor `5.205s`, appipc `5.173s`, frank-app `4.763s`.
- `go test -tags frank_test_reduced_limits -count=1 ./internal/appipc/... ./internal/appctl/...` — PASS; appipc `0.233s`, recovery `2.380s`, settlement `1.743s`, terminal `2.534s`.
- Pre-T15 terminal ordinary `go build ./... && go test ./... && go vet ./...` — PASS; fixtures `194.080s`; fresh `go test -count=1 ./test/invariants` — PASS `1.261s`.
- Exact final-head ordinary command on `f090868f`: first attempt encountered one intermittent foreign governed-fixture red in `TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate` when its SECOND nested whole-suite run returned red and therefore satisfied the deliberately false `expect_green=false` claim. No byte changed. Exact isolated rerun of that test — PASS `154.646s`; subsequent exact full `go build ./... && go test ./... && go vet ./...` — PASS, fixtures `204.090s`, vet clean. This history is disclosed, not erased; the final observed gate is green.
- `git diff --check` — PASS. Literal source `git status --short` immediately before this report — no output (clean).

Evidence levels:
- E2: all implementation, durability, process/socketpair, fake-counterpart, real local Unix broker-control, race, compile-negative, reduced-limit, boundary, and repository-battery claims above.
- E3/E4: none claimed. No live credentialed provider call or composed real m-8/m-9 turn ran in this slice.

Composition boundary and remaining risk, stated without promotion:
- The exported terminal `Runner` invokes an injected `Starter`; this standalone s13 head deliberately passes no composed starter, so `run start` fails loud until the serialized app composition supplies the real s14 connector + s15 worker launcher. The registry and runtime refusal/mutation properties are E2-proven; a real composed start is not claimed here.
- Recovery E2 proves durable selection, fencing/washout, control establishment, tuple install/re-proposal, and assign-gate opening. Fresh real CTRL-W/CTRL-C/DATA-P creation, connector bootstrap, lease-bind, worker assign, and first admission remain the serialized sibling-composition check after restack; local fake/component greens are not misreported as that composed execution.
- The exact-head fixture's one intermittent nested red is a retained review datum despite the unchanged isolated and final full greens. The four carriage files remain byte-exact and were not locally repaired.
- Restack onto the then-current serialized predecessor, rerun ordinary/reduced/compile batteries, and perform substantive re-review before any merge grant. H-12, queue rows, E3, release binding, and same-UID residuals remain with their named owners/gates.

Out-of-scope preserved:
- No store submit/poll/export act (the store is dormant under `020247`), GitHub CI/CD, live credentialed provider call, push, PR, restack, merge, deployment, publication, release, operator decision, or sibling-lane closure.
- No policy ownership, conductor seat/credential/verb, provider credentials, provider payload hot-path, routing, Step-4 ceiling policy, or parked-row clearing path was added.

ACTIONS_GIT_REF: source branch `s13-m10-module@f090868fb28347de93464fd56df4514928f479cd`; base `b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9`; final diff `b7f406b2..f090868f`; implementation worktree `/Users/jack/Programming/harness-s13-m10-module/frank`; this report + one append-only s13 INDEX row in the shared hand-relay workspace; no store/push/PR/restack/merge act.
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s13-m10-module@f090868fb28347de93464fd56df4514928f479cd` immediately before filing this report.
Next requested action: `s13.planner` performs the one adversarial end-review over exact head `f090868fb28347de93464fd56df4514928f479cd`, including the two composition-boundary statements and the disclosed intermittent fixture result. On a clean/folded verdict, return the merge-ready SITREP upward for serialized restack/rerun/re-review and the operator's terminal MERGE-GATE. No merge is authorized by this report.
