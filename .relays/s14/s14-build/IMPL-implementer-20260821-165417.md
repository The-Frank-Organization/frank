## IMPL — s14 m-8 connector implementation report + end-of-slice adversarial review: **APPROVE at E2** on `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`; T1a–T15 complete, the review-found live-stream/cancellation/bootstrap/opacity cuts closed in `316f97a6`, every §2a LOCAL + PRODUCER leg green, source worktree clean; no push, PR, merge, restack, E3, or sibling-lane closure claimed

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s14-build-implementer-report
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the serialized operator MERGE-GATE remains terminal; s13 precedes s14, then s15; restack/rerun/re-review remain owed before any merge grant
FILED_AT_LOCAL: 20260821-165417
IN_REPLY_TO: frank/.relays/s14/s14-build/IMPL-planner-20260821-150509.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: implementation report + pair end-review APPROVE at E2 — s14 m-8 connector complete on clean local branch at 316f97a6; planner end-review requested, merge remains gated

Summary:
- Implemented T1a–T15 on the isolated local branch `s14-m8-connector`, launched from exact `main@b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9`, now at `316f97a6025111cd3aa65e2841b4465010694bb8`.
- T9.4 and T9.5 are distinct liftable commits under the two bounded grants. T9.4 changed only `internal/executor/**`, preserved the bare governed verdict token and every `internal/observe/**` byte, forced offline resolution, preseeded a fresh run-local cache from go.sum-verified host content, and retained a private capped LAST-byte diagnostic by run key. T9.5 changed only the ruled fixture path and exact capacity/discrimination bytes.
- T10–T15 implement the normalized provider stream, CTRL-C lifecycle and epoch fence, total v2 outcomes, typed cancellation, the real connector process, and the package/wire/fake-counterpart close battery.
- My end-of-slice adversarial review found substantive live-composition defects at the initial T15 head `ec151eb9`; I fixed them before verdict in commit `316f97a6`: batch-only stream release → emit-as-parse callback; cancellation/event terminal races → single-terminal arbitration with immediate context abort; invocation truth moved to the last pre-dial gate; policy bytes now load before `hello` and are reused without a second read; redirect + Retry-After failure truth fixed; opaque replay validation no longer trims bytes; the missing lane-id opacity structural guard added. The corrected head was re-reviewed and is **APPROVE at E2, zero remaining findings**.
- No PR was opened and nothing was pushed or merged. This is the local implementation handoff required before the Planner's end-review and the serialized restack/rerun/re-review sequence.

PR: none — local concurrent-slice branch only; no publication authority inferred.

Plan lock:
- Plan of record: `frank/.relays/s14/s14-build/PLAN-planner-20260821-145914.md` r8, SHA-256 `5ea4a77a83874542e4dba78fde8e6f61e16d740e1b49c48b0933474219535911`.
- Approving review: `s14-build-plan-review-8`, SHA-256 `02959ce7f61c7faf45c8287723b97afc958f53cc3f7ef5b4b6eea8c1a02e8a73`.
- Live dispatch: this report's `IN_REPLY_TO`, exact SHA-256 `9a1c1cc30e169bf67a0b34b35e3cc16be153c101d4614dce495c2efb60ec7756`.

Branch and commits:
- Base: `b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9`.
- Head: `316f97a6025111cd3aa65e2841b4465010694bb8`.
- Task commits, in order: `3257ec35` T1a · `7e7d4f2f` T1b · `d192c211` T1c · `7d8e1600` T2 · `eacc835a` T3 · `53242e16` T4 · `12f018a4` T5 · `9f12a9be` T6 · `4dc769d4` T7 · `afb134be` T8 · `471532a1` T9 · `b86b8bc1` T9.4 · `b9b2b3c2` T9.5 · `93405aec` T10 · `afabbd10` T11 · `1b27584e` T12 · `b8df2b30` T13 · `c9ca7d05` T14 · `ec151eb9` T15 · `316f97a6` end-review corrections.

Files changed:
- Production + tests: `cmd/frank-connector/**`, `internal/connector/**`.
- Ruled dependency bytes: `go.mod`, `go.sum` (`go 1.25.0`, pinned `golang.org/x/text v0.41.0`; later tidy tripwires did not move either byte set).
- One-commit R9/R10 grant: `internal/executor/executor.go`, `internal/executor/executor_test.go`.
- One-commit R8 grant: `test/fixtures/s8_exit_gate_test.go`.
- Zero `internal/observe/**`, conductor protocol, frozen-oracle, script, appipc, seatclient, or sibling-slice bytes changed.

Acceptance criteria status — §2a LOCAL + PRODUCER matrix (E2; cross-lane closure expressly not claimed):

| fixture | s14 evidence |
|---|---|
| r12-1 denied sends nothing | LOCAL green: all nine ordered deny reasons; zero resolver/transport counters. |
| r12-2 attach after authorize | LOCAL green: resolver unreachable on every deny and reference-without-authority negative. |
| r12-3 frozen immutable | LOCAL green: copy-on-read core/body, digest/body mutation guard. |
| r12-4 no retry | LOCAL green: fresh-dial/nothing-written/headers/midstream vectors, pool absence, nil GetBody, redirect refusal, h2 never offered, Retry-After recorded only. |
| r12-5 stale/ahead epoch | LOCAL green; PRODUCER fake m-10 query/answer path green; real m-10 remains s16. |
| r12-6 CTRL EOF | LOCAL fail-closed/zero further path green; PRODUCER fake EOF green; real supervision remains s16. |
| r12-7 endpoint grammar | LOCAL full frozen vector consumer green. |
| r12-8 catalog | LOCAL closed-schema/semantic negatives + AST opacity guard green. |
| r12-9 policy | LOCAL exact-JCS/digest/membership negatives; PRODUCER assignment comparand green; real manifest remains s16/m-10. |
| r12-10 stream dialect | LOCAL corpora, terminal, timeout, overflow, inert fragments, opaque replay green; real credentialed re-record remains live-E3/s16. |
| r12-11 determinism | LOCAL repeated translate-byte property green. |
| r12-12 sentinel secret | LOCAL process output + CTRL-C/DATA-P/events/typed errors/core/digests sweep green; PRODUCER sentinel/process contributed; m-9/m-10/conductor legs remain sibling. |
| r12-13a object route | LOCAL package-boundary/no-conductor-import-or-serializer AST guard green; 13b remains sibling with sentinel type set contributed. |
| r12-14 wire capture | LOCAL exact frozen request + sole auth attach + host/content-length/connection census, no uncensused header, green. |
| r12-15 replay scope | LOCAL wrong-lane/wrong-turn/verbatim/legacy-field legs green. |
| r12-16 reject totality | LOCAL all reasons, zero stream/resolver/transport, CTRL-before-DATA barrier + reversed mutation green; PRODUCER fake CTRL receiver green; real terminal row remains s16/m-10. |
| r12-17 epoch inertness | LOCAL typed replies, zero result/counters; PRODUCER fake query/no-row assertion; real rows remain s16/m-10. |
| r12-17b cancellation | LOCAL pre/post disposition truth, zero-wire pre-dial gate, partial truth, duplicate idempotence, raw-loss non-cancellation, never-failed supply green; PRODUCER fake intent/row assertions green; two-view reconciliation remains s16/m-9/m-10. |
| add §6.1–§6.11 | LOCAL presence/absence, P2a counters, B identity, E determinism/zero-tools/field-set/order/strict/independence, v2 decodability, payload/secret absence, refusal-stage totality all green; named observer/classifier halves are PRODUCER self-checks only, real m-3/m-9/m-10 consumers remain sibling/restack. |

Boundary contract proof:
- The production call chain is `request.Parse → translate.Translate → freeze.Freeze → authorize.Evaluate → credentials.Attach → transport.SendGated`, with secret resolution lexically absent before the allowed verdict and the mutation guard before the last pre-dial invocation gate.
- The app-internal AST boundary test rejects connector→conductor imports, conductor→connector imports, and lane-id semantic parsing outside the catalog owner.
- `cmd/frank-connector` accepts only clean absolute artifact paths inside the private runtime directory, re-marks both inherited FDs close-on-exec, disables core dumps before loading credentials, clears the ordinary environment, performs one credential→catalog→policy→hello/assign/ready incarnation, and treats controlled shutdown as clean without solo restart.
- No provider payload type is serialized into a conductor payload; no credential secret enters argv, frames, events, errors, logs, frozen core, B, or E.

Tests/verification (fresh corrected-head E2 unless explicitly historical):
- `go test -mod=readonly -race -count=20 ./internal/connector/service ./internal/connector/stream ./internal/connector/attempt ./internal/connector/transport ./internal/connector/control` — PASS.
- `go test -mod=readonly -race -count=1 ./internal/connector/... ./cmd/frank-connector` — PASS.
- `go vet ./cmd/frank-connector ./internal/connector/...` — PASS.
- `go build -o /tmp/frank-s14-build/frank-connector ./cmd/frank-connector` — PASS.
- `go mod tidy` diff tripwire — PASS, no `go.mod`/`go.sum` movement.
- Corrected-head `go test -mod=readonly -p=1 -count=1 ./...` — PASS; `test/fixtures` 185.017s.
- Earlier T15 full battery — PASS; `test/fixtures` 195.000s. T14 full battery — PASS; `test/fixtures` 189.366s.
- R9/R10 ladder: offline missing-cache and retained-LAST-tail legs PASS; repaired fixture isolated THREE consecutive PASS (145.68s, 146.79s, 141.55s); full battery TWO consecutive PASS before T10 (fixture legs 179.935s, 182.763s).

Evidence levels:
- E2: all implementation, boundary, fake-counterpart, fake-provider, process-spawn, transport-capture, race, and repository-battery claims above.
- E3/E4: none claimed. The credentialed live dialect re-record and composed real-sibling app remain later gates.

Out-of-scope preserved:
- No merge, push, PR, deployment, release, GitHub CI/CD, live credentialed provider call, real m-9/m-10/m-3 consumer claim, or operator decision performed.
- Framing-vs-live-appipc remains a mandatory serialized restack check; this branch used the in-fence byte-conformant connector codec exactly as planned.
- H-12 and the same-UID MVP residual remain unchanged.
- Queue rows 7/9/11/12/13 remain for the operator merge brief; this report does not discharge them.

Remaining risk:
- The concurrent-branch contract is still live: restack onto the then-current serialized predecessor, rerun the full batteries, and perform substantive re-review before any s14 merge grant.
- Cross-lane fixture columns remain owned by s16 and the named sibling domains. Local fakes are producer assertions, not evidence that the real consumers have landed.

ACTIONS_GIT_REF: source branch `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`; base `b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9`; final diff `b7f406b2..316f97a6`; implementation worktree `/Users/jack/Programming/harness-s14-m8-connector/frank`; this report + one append-only s14 INDEX row in the governance workspace; no PR/push/merge.
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `316f97a6025111cd3aa65e2841b4465010694bb8` immediately before filing this report.
Next requested action: s14.planner performs the one end-review over exact head `316f97a6025111cd3aa65e2841b4465010694bb8`; on a clean verdict, report upward for the serialized restack/rerun/re-review and operator MERGE-GATE. No merge is authorized by this report.
