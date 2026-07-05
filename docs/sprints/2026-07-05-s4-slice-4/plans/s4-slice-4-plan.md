# s4 Slice-4 Implementation Plan — the wire-up (shim · lifecycle · §7 record)

> **For the executing Implementer:** execute under Superpowers `executing-plans`, task-by-task, in plan order. Steps use checkbox (`- [ ]`) syntax. Every task ends with the battery green and a commit. Implementation authority arrives ONLY via the delegated dispatch relay under the standing conditions — not from this document.

**Goal:** a real host session files a relay through `submit()` over the per-seat MCP shim and a second real session receives it via `project()`/`read` — no human transport; plus lifecycle hardening (second-connect, per-recipient wake, frame bounds) and the §7 config-change record discharging `OI-S3-CONFIG-CHANGE`.

**Locked design:** `docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md` **r3 = main@2ef9437** (DESIGN_DOC_ID `s4-slice-4-design`; GRILL_LOCK `s4-grill-s4-wire`; approving review `s4-wire-design-r3-review-implementer` 20260705-040925). Section references below (§n) are design sections. **The design rules; this plan sequences it.**

**Architecture (one breath):** conductor changes are engine-internal hardening (frame bounds §6.4, active-channel index §4, per-recipient wake §5, the `config_change` mutation class §6); the shim (`cmd/frank-mcp`) is a NEW stdio binary translating real MCP ↔ the private socket dialect via the promoted `channel.Client` (§2–§3); docs are claim-honest deliverables (§7).

**Tech stack:** Go (existing module, zero new deps — MCP framing is newline JSON-RPC 2.0, hand-rolled like the existing dialect); existing crashpoint/child-SIGKILL harness; existing fixture patterns (`test/fixtures/main_assembly_test.go` = the socket-E2 donor).

## Global constraints (bind every task)

- Guardrail surface stays exactly `submit`/`project`/`read`; the shim adds affordance, never authority (§0).
- Terminal enum byte-exact `{accepted, rejected, held}`; no new value, anywhere.
- I-PH on every seat/host-visible byte incl. every shim-generated error (§3.5); ONE named exception: the `max_frame_bytes` value in frame refusals (§6.4 carve-out — assert exactly that one value leaks, nothing else).
- **m-1-PENDING labels:** the `config_change` record_kind token + provenance and the shim custody note are m-1-fidelity surface (packet in flight, `.relays/s4/s4-fidelity-m1/`). Tasks 4–5 carry them as **m-1-pending-confirm**; if m-1 prescribes different shapes, the plan folds VERBATIM or routes back (standing condition 2). **No dispatch before the m-1 approve is on record.**
- No binding-table shape change (hard stop + escalate); no in-band rotation/supersede; no socket-dialect rewrite; OUT list per §10.
- Every claim surface written by any task carries "transport/provenance only; done-state and `record_integrity` remain `self_reported` until Step-2 observe" + the D5 custody note where credentials appear ([VP-W2]).
- Run-FAIL-first on every negative fixture (S4-SC1, S4-NG4, S4-FR1/FR2, S4-IPH*, C7 refusal legs).
- Commit per task (`s4:` prefix); battery (`go test -count=1 ./...` + `go vet ./...`) green before every commit.

## File map (the SCOPE_DIFF universe)

| surface | files |
|---|---|
| engine config | `internal/config/config.go` (+`config_test.go`) — `MaxFrameBytes` member |
| channel layer | `internal/channel/server.go` (+ new `internal/channel/server_lifecycle_test.go`, `internal/channel/frame_test.go`) |
| assembly | `cmd/frank/main.go` (wake wiring, per-seat recovery nudge, socket-path pre-flight) |
| store | `internal/store/store.go` (`PendingDeliveryFor`), `internal/store/genesis.go` (chain walk), `internal/store/config_change.go` (new; + test) |
| engine loop | `internal/engine/loop.go` + `internal/engine/submit.go` (config_change acceptance + wake events) |
| registry | `internal/fieldspec/registry.json` (`config_change` token — m-1-pending-confirm) |
| crash harness | `test/fixtures/f11_test.go`, `test/fixtures/s4_config_change_test.go` (new) |
| shim | `cmd/frank-mcp/main.go`, `cmd/frank-mcp/mcp.go`, `cmd/frank-mcp/schema.go`, `cmd/frank-mcp/errors.go` (all new) + `cmd/frank-mcp/*_test.go` |
| shim/bridge fixtures | `test/fixtures/s4_shim_test.go`, `test/fixtures/s4_iph_test.go` (new) |
| docs | `docs/ops.md` (new), `cmd/frank-mcp/README.md` (new), `docs/usage-data.md` (new), **root `README.md` (ONE bounded delta — PENDING the orchestrator fence ruling; ASK filed with this plan)** |
| gate procedure | `docs/sprints/2026-07-05-s4-slice-4/results/e3-gate-procedure.md` (new) |

Everything else is OUT for SCOPE_DIFF purposes. `internal/seat/binding.go` is deliberately ABSENT (no shape change — the credential hash is computed channel-side).

---

### Task 1 — frame bound: config member + both-direction typed refusal (§6.4; S4-FR1/FR2)

**Files:** Modify `internal/config/config.go`, `internal/channel/server.go`; Create `internal/channel/frame_test.go`.
**Interfaces produced:** `config.EngineConfig.MaxFrameBytes int` (0 ⇒ default 1 MiB via `EngineConfig.FrameBytes() int`); server + client scanners sized from it; error classes `frame-too-large` (inbound) and outbound refusal payload `{"error_class":"frame-too-large","hint":"narrow the request: re-read by relay_id or paginate project"}` + the bound value.

- [ ] 1.1 Write failing tests in `frame_test.go`: `TestInboundOversizeFrameTypedRefusal` (dial a served socket, send one `len > FrameBytes()` line, assert a response frame with error containing `frame-too-large` and the bound value, then assert the SAME connection still answers `tools/list` — connection alive), `TestOutboundOversizeProjectRefusal` (serve a store whose `read` result exceeds a small configured `MaxFrameBytes` (set 4096 in the fixture store's engine.json), assert typed refusal + hint, not a dead socket), `TestFrameBytesDefault` (`EngineConfig{}.FrameBytes() == 1<<20`).
- [ ] 1.2 `go test ./internal/channel/ -run Frame -count=1` → FAIL (no member, no refusal path).
- [ ] 1.3 Implement: add `MaxFrameBytes int \`json:"max_frame_bytes,omitempty"\`` + `FrameBytes()`; thread the value into `Serve`/`ServeAuthenticated` and `Dial`/`DialAuthenticated` (new optional arg or a `FrameLimit` field set by the assembly — pick ONE, document in the code); replace both `bufio.NewScanner` defaults with `scanner.Buffer(make([]byte, 64*1024), limit)`; inbound: on `scanner.Err() == bufio.ErrTooLong` write one refusal frame then RESUME reading the next line if the transport allows, else close after writing; outbound: length-check every `write()` payload in `serverConn.write` and substitute the typed refusal JSON when over-bound.
- [ ] 1.4 Tests pass; whole battery green (`go test -count=1 ./...` — the omitempty member keeps every existing engine.json digest byte-identical: assert by running the s2/s3 fixture suites unmodified).
- [ ] 1.5 Commit `s4: frame bound — config-sourced 1MiB default, typed refusal both directions`.

### Task 2 — per-credential active-channel index (§4; S4-SC1/SC2/SC3)

**Files:** Modify `internal/channel/server.go`; Create `internal/channel/server_lifecycle_test.go`.
**Interfaces produced:** `Server.active map[[32]byte]*serverConn` under `s.mu`; `serverConn.credHash [32]byte`; auth reject error class `auth:channel-active`. NO seat/binding change.

- [ ] 2.1 Failing tests: `TestSecondConnectSameCredentialRejected` (two `DialAuthenticated` same credential: second errors `auth:channel-active`; FIRST connection still serves `tools/list`), `TestDistinctCredentialsBothConnect`, `TestProvenDeadRecovery` (first client `Close()`, brief poll for server-side reap, re-`DialAuthenticated` same credential succeeds), `TestKillHostEscapeHatch` = S4-SC3 (run a child process holding the connection, `SIGKILL` it, assert reconnect succeeds — kernel close-detection, the wedged-host escape hatch made mechanical).
- [ ] 2.2 Run → FAIL (today the second connect is accepted — the audit's live P4).
- [ ] 2.3 Implement: `sha256.Sum256(credential)` retained on the conn at auth; check-and-set in the auth branch under `s.mu`; delete in `run()`'s defer; never store or log the raw credential.
- [ ] 2.4 Tests pass; battery green; `-race` on `./internal/channel/`.
- [ ] 2.5 Commit `s4: one active channel per credential — typed reject + proven-dead recovery [VP-W1]`.

### Task 3 — per-recipient wake; retire broadcast + global pending (§5; S4-NG1..4)

**Files:** Modify `internal/channel/server.go`, `cmd/frank/main.go`, `internal/store/store.go`; extend `server_lifecycle_test.go` + `test/fixtures/s4_shim_test.go` assembly legs.
**Interfaces produced:** `serverConn.seat string` (from auth meta); `Server.PushTo(seat string, frame []byte) error` (writes to that seat's live conns only); `store.PendingDeliveryFor(seat string) (bool, error)`; wake frames `{"kind":"delivery-nudge","relay_id":"…"}` (recipients only) and `{"kind":"recovery-nudge"}` (own-seat, no seats list). `Push`/`broadcast`/`QueuePush`/`flushPending`/global `pending` REMOVED from the serving path (delete `QueuePush` + `pending` outright; keep nothing dead).

- [ ] 3.1 Failing tests: `TestNudgeRecipientOnly` (A submits TO seat-b: b's conn receives `delivery-nudge` with the relay_id; **a's conn receives nothing** — today FAILS by broadcast), `TestNudgeAllRecipients` (TO + CC seats each nudged once), `TestOfflineRecipientNudgedOnReconnect` (submit to offline b; b connects; b alone gets `recovery-nudge`), `TestNoCrossSeatMetadata` = S4-NG4 (capture EVERY frame b's connection ever receives across connect/auth/submit/recovery; assert no occurrence of seat-a's name, relay IDs not addressed to b, or any pending-state list), `TestPendingDeliveryFor` (store unit).
- [ ] 3.2 Run → FAIL (broadcast + all-seats recovery frame, live-confirmed in audit).
- [ ] 3.3 Implement: recipients of a committed record = its envelope TO/CC seats (the same set the mailbox intents used — derive in the assembly's post-commit hook where the committed `record.Record` is in hand: `process()` for recovery-path and `loop.AfterCommit`-adjacent wake in the serving path; wake emission = `PushTo(seat, frame)` fire-and-forget, write-error-continue); auth path: after `flushPending` removal, check `PendingDeliveryFor(meta.Name)` and send own-seat `recovery-nudge`; startup: delete the `PendingDeliverySeats` broadcast block in `main.go:199-212` (the on-auth check subsumes it); `PendingDeliverySeats` stays for any remaining internal use or is deleted if orphaned.
- [ ] 3.4 Tests pass (FAIL-first evidenced for NG4); battery + `-race` on channel/engine; grep `rg "QueuePush|broadcast\(" internal/ cmd/` → serving path clean.
- [ ] 3.5 Commit `s4: per-recipient wake (locked §8.3 grain); global pending queue retired; no cross-seat metadata [guide Q1]`.

### Task 4 — `config_change` acceptance surface + derived config materialization (§6.1–6.2; C7 refusal/commit legs) **[m-1-pending-confirm on the token]**

**Files:** Modify `internal/fieldspec/registry.json` (enum_set `record_kind` += `"config_change"` — **m-1-pending-confirm, fold verbatim on the verdict**), `internal/engine/submit.go`; Create `internal/store/config_change.go` + `internal/store/config_change_test.go`.
**Interfaces produced:** header contract `record_kind=config_change`, `member ∈ {engine, fieldspec}`, `new_digest` (headers), body = full new member content (string); `store.ConfigChangeIntents(rec record.Record) []Intent` (the derived member-file write via the existing atomic-write machinery — a projection-grade intent, replayed by recovery); engine check: non-operator submit carrying `record_kind=config_change` ⇒ typed reject naming the field; `new_digest` recomputed conductor-side over `{current members ⊕ body}` and must equal the header claim, else typed reject.

- [ ] 4.1 Failing tests (`config_change_test.go` + an engine leg): `TestConfigChangeNonOperatorRejected`, `TestConfigChangeDigestMismatchRejected`, `TestConfigChangeCommitMaterializesMember` (operator-submitted registry delta: exactly ONE canonical record appears (one rename — assert via the store's records dir before/after), `config/fieldspec/registry.json` bytes = the body, INDEX row present), `TestRunningConfigUnchangedUntilRestart` (post-commit, the SERVING registry still renders the old form — no hot reload).
- [ ] 4.2 Run → FAIL.
- [ ] 4.3 Implement per §6.2: validate → pivot → `applyIntents(ConfigChangeIntents(rec))` in the same loop iteration; the intent writer reuses `fsio.WriteFileAtomic`; the running `pinned` is untouched.
- [ ] 4.4 Tests pass; battery green (fresh stores: the shipped registry now contains the token — S1–S3 suites must stay green unmodified, the zero-regression floor).
- [ ] 4.5 Commit `s4: config_change mutation class — operator-only, embed-bytes, record-pivot-then-derived-materialization [guide Q2; m-1-pending-confirm token]`.

### Task 5 — phase-0 genesis→config-change chain walk (§6.2; C7 chain legs)

**Files:** Modify `internal/store/genesis.go`, `internal/recover/recover.go` (call-site only if signatures move); extend `config_change_test.go` + `internal/recover/recover_test.go`.
**Interfaces produced:** `(*Store) ExpectedConfigDigest() (string, error)` — genesis digest superseded in commit order by each accepted `config_change`'s `new_digest`, derived from committed records ONLY; `ValidateGenesis` becomes: loaded digest == ExpectedConfigDigest, with ONE re-materialization attempt (rewrite members from the latest chain record's body via `ConfigChangeIntents` replay) before the existing fail-closed-serving-reads diagnostics.

- [ ] 5.1 Failing tests: `TestChainWalkLatestWins` (genesis + two config_change records → expected = last), `TestPhase0RematerializesFromChain` (delete/corrupt `config/fieldspec/registry.json` after a chain exists → recovery restores bytes from the record body and opens Ready), `TestPhase0FreshStoreUnchanged` (no chain ⇒ byte-identical behavior to today — the S1–S3 zero-regression leg), `TestPhase0PersistentMismatchStillDiagnostics` (chain says X, files say Y, re-materialization forced to fail (read-only config dir) ⇒ Diag, never a brick, never Ready).
- [ ] 5.2 Run → FAIL. — [ ] 5.3 Implement. — [ ] 5.4 Battery + the restart re-render leg: `TestRestartWithNewRegistryBouncesStaleForm` (form digest rendered pre-change bounces `re-render` after restart; re-rendered form succeeds — [VP-W4] drift leg, E2 form).
- [ ] 5.5 Commit `s4: phase-0 walks the genesis→config-change chain; config/ re-materialized from records [guide Q2(iii)]`.

### Task 6 — crash harness gains the class (§6.2; F11 one-pivot; applicability row)

**Files:** Modify `test/fixtures/f11_test.go` (`f11Classes()` += `"config-change"`); Create `test/fixtures/s4_config_change_test.go` (child-SIGKILL legs, donor pattern = the existing per-class kills).
**Steps:** [ ] 6.1 add the class → the applicability-map assertion FAILS (missing row — the map is falsifiable, S2 property); [ ] 6.2 crash legs at every syscall boundary: pre-pivot (nothing happened; intake re-enqueue), between pivot and member-write (recovery re-materializes — record is truth), mid-member-write (atomic write discipline); assert exactly ONE canonical rename per change and store convergence each leg; [ ] 6.3 applicability row `crash-expected`/`clean-completion` per crashpoint name, trace-verified; [ ] 6.4 battery; commit `s4: crash matrix gains config-change (locked F11 :172 catches up) [ruling condition 4]`.

### Task 7 — the shim skeleton: real MCP over stdio ↔ `channel.Client` (§2, §3.1, §3.4, §3.5; S4-MCP1)

**Files:** Create `cmd/frank-mcp/main.go` (flags/env, lifecycle, reconnect backoff), `cmd/frank-mcp/mcp.go` (JSON-RPC 2.0 framing: `initialize` → `{protocolVersion, capabilities:{tools:{listChanged:true}}, serverInfo:{name:"frank-mcp"}}`, `notifications/initialized`, `ping`, `tools/list`, `tools/call`, graceful `unknown method` per JSON-RPC error object), `cmd/frank-mcp/errors.go` (THE scrub chokepoint: every MCP-bound error maps to `shim:conductor-unreachable | shim:auth-failed | shim:connection-lost | shim:frame-too-large | shim:protocol-error`; raw detail to stderr only, never the credential), + `mcp_test.go` goldens driving stdin/stdout pipes.
**Interfaces produced:** binary `frank-mcp`; env `FRANK_SOCKET`/`FRANK_CREDENTIAL` (+ `-socket`, `-credential-file`); socket-side ids allocated from 1 (never 0 — F-W5 pin); `project`/`read` static schemas; tool descriptions with the honesty line ("Files a governance relay (transport/provenance only …)").
**Steps:** [ ] 7.1 golden tests first (initialize handshake shape; tools/list = exactly three tools with the descriptions; unknown method → JSON-RPC error object; dial-failure at startup → initialize still coherent, tools/call returns `shim:conductor-unreachable`, and the golden asserts NO socket path substring in any stdout byte); [ ] 7.2 FAIL; [ ] 7.3 implement against a real served socket in-test (donor: main_assembly_test.go); [ ] 7.4 pass + battery; [ ] 7.5 commit `s4: frank-mcp shim skeleton — MCP stdio front, channel.Client back, scrubbed error chokepoint`.

### Task 8 — Form→JSON-Schema mapping + const digest (§3.2 closed rule; S4-SCH1/SCH2)

**Files:** Create `cmd/frank-mcp/schema.go` + `schema_test.go`.
**Interfaces produced:** `SchemaFromForm(form fieldspec.Form, digest string) map[string]any` implementing the §3.2 table VERBATIM: all properties string-typed under `headers` (+ envelope `to`/`cc`/`dispatch_id` + `body`); enum from rendered Options byte-exact; structured types described as canonical-JSON strings; `form_digest` required `const`; NO `from`/`role`/`relay_id`/`delivery_state` property, ever; zero reshaping (the value strings pass through byte-for-byte both ways).
**Steps:** [ ] 8.1 failing tests: `TestSchemaStringCarrierAllTypes` (a synthetic form covering text/enum/bool/id_ref/row_array/address_list maps per the table), `TestSchemaSystemFieldAbsence` (grep the marshaled schema for `"from"`/`"role"` etc. → absent), `TestSchemaConstDigest`, `TestSchemaEnumByteExact`, S4-SCH2 round-trip legs against a live socket (canonical string accepted; NON-canonical encoding → the conductor's typed field-named violation surfaces through the MCP result; schema description states the string carrier); [ ] 8.2 FAIL; [ ] 8.3 implement; [ ] 8.4 pass + battery; [ ] 8.5 commit `s4: rendered form IS the submit inputSchema — closed string-carrier rule, const digest [F2]`.

### Task 9 — drift loop: re-render → refresh + `tools/list_changed` (§3.3; S4-RR1)

**Files:** Modify `cmd/frank-mcp/mcp.go`; extend tests.
**Steps:** [ ] 9.1 failing test `TestPhaseSwitchDriftLoop`: host submits declaring `PHASE: PLAN` while holding the SITREP schema → result carries the structured non-fatal `re-render` violation with the hint text; shim emits `notifications/tools/list_changed`; next `tools/list` serves the PLAN-phase schema whose digest then succeeds; [ ] 9.2 FAIL; [ ] 9.3 implement (on violation class `re-render`: re-fetch Describe for the DECLARED phase/tier; never synthesize a digest; optional pre-fetch on visible phase mismatch); [ ] 9.4 pass + battery; [ ] 9.5 commit `s4: drift negotiation — re-render bounce refreshes schema via list_changed [VP-W4 drift leg]`.

### Task 10 — the I-PH bridge matrix (§3.5; S4-IPH1..7 + the carve-out leg)

**Files:** Create `test/fixtures/s4_iph_test.go`.
**Steps:** [ ] 10.1 one test per exit-gate surface class, each capturing REAL bytes and asserting zero occurrences of: store root path, `config/` paths, socket path, `binding`, `seats.json`, credential value — (1) tools/list descriptions, (2) input schemas across phases, (3) tool-call results incl. rejected/held outcomes, (4) notifications/poll hints, (5) reconnect errors (kill the socket mid-session; capture the shim's MCP-visible error), (6) credential-failure errors (bad credential), (7) shim diagnostics (dial failure golden from Task 7 re-asserted at matrix grain); plus `TestCarveOutExactlyOneValue`: the frame refusal MAY contain the bound value; assert the bound value appears in NO other surface class and no OTHER config value appears anywhere; [ ] 10.2 run-FAIL-first where instrumentable (plant a deliberate leak in a test double to prove the scanner bites), then green against the real surfaces; [ ] 10.3 battery; commit `s4: I-PH matrix across all seven bridge surface classes + ceiling carve-out pinned [VP-W3]`.

### Task 11 — ops surface + custody + usage docs + socket pre-flight (§7; honesty sweep)

**Files:** Create `docs/ops.md`, `cmd/frank-mcp/README.md`, `docs/usage-data.md`; Modify `cmd/frank/main.go` (socket-path length pre-flight: `len(path) >= 100` ⇒ typed startup error naming the darwin limit — operator-facing stderr, allowed); root `README.md` **ONLY IF the fence ruling grants it** (the one bounded delta: refresh the fresh-store sentence to name the now-landed §7 record; ruling relay cited in the SCOPE_DIFF row).
**Content contract (acceptance):** ops.md = start/stop/status + short-socket-path rule + team-store conventions + minting workflow end-to-end incl. **one seat = one credential = one config entry** and the two blessed wiring patterns (per-seat config scopes; `${VAR}` indirection) with per-host support pinned + the wedged-host remedy (guide Q5 condition 1, verbatim); shim README = custody posture (env-var default, file secondary, D5 note verbatim, "no in-band rotation — compromised credential = stop conductor + admin-time surgery"), poll-first posture, the transport-only line; usage-data.md = the store IS the record; read paths (`project`/`read`, records/, INDEX); "aggregation = s5". Sweep check: `rg` for the honesty line on every new claim surface; no "verifies/authorizes" anywhere in shim-visible text.
**Steps:** [ ] 11.1 socket pre-flight failing test (assembly test with a >104-byte path asserts the typed error, not `bind: invalid argument`); [ ] 11.2 implement + docs; [ ] 11.3 sweep greps green; battery; [ ] 11.4 commit `s4: ops/custody/usage surfaces — claim-honest; socket-path pre-flight [F-W3]`.

### Task 12 — E2 floors (gate lines, mechanical)

**Steps:** [ ] 12.1 `go test -count=1 ./...` (ALL packages incl. new) + `go vet` + `-race` on channel/seat/engine/intake/store; [ ] 12.2 enum grep floor (`rg -n '"bounced"' internal/ cmd/` → zero; the byte-exact triple asserted in the existing fixtures still green); [ ] 12.3 three-tool enumeration floor re-asserted THROUGH the shim (Task 7 golden covers it; re-run flagged); [ ] 12.4 zero-regression: S1/S2/S3 suites untouched-and-green, the real `$HOME/frank-s2-store` untouched (`ls` count unchanged — read-only check); [ ] 12.5 commit only if fixes were needed (`s4: floor fixes`), else record the runs in the task log.

### Task 13 — the E3 gate procedure (operator-run; documented, not automated)

**Files:** Create `docs/sprints/2026-07-05-s4-slice-4/results/e3-gate-procedure.md`.
**Content contract:** the exact operator-run sequence for every live exit-gate leg — store init (this store becomes the PERSISTENT team store: §7 backstop trips here); mint A/B (+ operator seat); wire host A (Claude Code) + host B (Codex) per ops.md; the live relay leg; the adversarial legs (no/bad credential, second-connect reject, forged-FROM submit, I-PH spot probes); the crash/liveness legs (kill frank mid-delivery → restart → wake re-issued exactly once; kill shim → reconnect → project catch-up; offline-seat nudge on reconnect); the §7 round-trip (operator submits a REAL registry delta as `config_change` on the now-EXISTING store → restart → phase-0 accepts via chain → stale form bounces re-render → re-rendered succeeds → crash legs already green at E2); **OI-S3-CONFIG-CHANGE disposition** (operator authors the `owed_disposition` through the operator channel citing the config_change relay; open owed set asserted EMPTY); every recorded claim carries the transport-only qualifier. Each leg names its evidence artifact (transcript/store paths) for the exit-gate report.
**Steps:** [ ] 13.1 write it; [ ] 13.2 dry-run every conductor-side step against a scratch store (everything except the two real hosts — those are gate-day, operator-designated); [ ] 13.3 commit `s4: E3 gate procedure of record`.

---

## Acceptance criteria (exit-gate mapping)

Every ROADMAP exit-gate line maps: live relay → T7-9+T13; adversarial → T2 (SC), T8 (no-FROM), T10 (I-PH matrix), T13 legs; crash/liveness → T6 + T13; §7 round-trip [VP-W4] → T4-6 + T13 (incl. OI discharge, open set empty); E2 floors → T12; honesty [VP-W2] → T11 + every artifact. Anti-half-fix guards: NG4/SC3/FR2/SCH2/carve-out are run-FAIL-first negatives; the applicability map stays falsifiable; no fixture may assert through the retired broadcast path.

## Out of scope (verbatim fence)

s5 consumer content · observe (Step-2) · routing (Step-3) · TUI (Step-4) · federation · external send · steer/interrupt beyond host-native · authority replacement · in-band rotation/supersede · socket-dialect rewrite · binding-table shape changes · the frozen s2-store upgrade (optional operator call, NOT a gate leg) · analytics over usage data.

## Operator-judgment items

Live-seat designation (gate-day); the §7 authorization itself (the mechanism); the README fence ruling (ASK filed alongside this plan); s4-close (merge is never implied by green fixtures).
