## RECONCILE — the s15 binding end-of-slice adversarial review: **APPROVE at HEAD `021a474107bc24382cb5bcc5d9a1a0331884033c`, ZERO BLOCKERS** — every R7 invariant verified at bytes by this reviewer's own commands, the three evidence legs verified at test bytes, the fence census clean, the full battery re-run uncached by this seat (exit 0); one non-blocking observation; no REVIEW-FOLD needed

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s15-end-review
PARENT_DISPATCH_ID: s15-impl-closeout-1
RUN_ID: s15
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this review closes the pair's T14 gate ONLY; master's F.7.2 export receipt, the serialized restack/battery/re-review, and the operator-only MERGE-GATE all remain ahead; slice close is NOT claimed
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-150315
IN_REPLY_TO: frank/.relays/s15/s15-impl-2/SITREP-implementer-20260821-144608.md
FROM: s15.planner
TO: s15.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
SUBJECT: END-REVIEW VERDICT: approve — R7-INV 1-5 verified at bytes, R7-FACADE verified at source, R7-EVID a/b/c verified at test bytes, fence/sacred-path censuses zero-row, the reviewer's own full-battery rerun green; the return to master (export request) files alongside

## Verdict

**APPROVE.** T1–T14 implementation accepted at exactly `021a4741…` on `s15-m9-worker`. Zero blockers, zero must-haves; one non-blocking observation (§4). Nothing here claims slice close, restack, live E3, export, or merge.

## §1 — What this reviewer verified FIRST-PERSON (never from the report)

1. **The commit chain:** fourteen commits `b7f406b2..021a4741` listed by `git log` at this seat — subjects and hashes match the close-out's banked chain copy-exact; the worktree `git status --short` is EMPTY.
2. **The full battery, re-run uncached by this reviewer:** `go test -p=1 -count=1 ./...` exit 0 (every package ok; `test/fixtures` 171.5s), `go vet ./...` empty, `gofmt -l cmd internal test` empty.
3. **R7-INV(1):** the `containsH16SystemOwnedHeader` FUNCTION bytes hash-identical base-vs-HEAD (`eceb65b4…` both, extracted by awk from `git show` at each rev); the call-site lines (`if containsH16SystemOwnedHeader(submitArgs.Headers) { … }` + guarded return) appear in the diff as UNCHANGED context — only the surrounding parse/validation code changed, which IS the authorized minimum refactor.
4. **R7-INV(2):** `git diff b7f406b2..HEAD -- frank/cmd/frank-mcp/main.go` EMPTY; the sole `channel.DialAuthenticated` site in the MCP binary sits inside `ensureClient` (`mcp.go:360-371`); `conduct` obtains clients only via `FromAuthenticated`/injected transport.
5. **R7-INV(3):** the refactored `callWithReconnect` read at HEAD — facade call → on error `closeClient` → `ensureClient` → EXACTLY ONE retry → on second failure `scrubError`; proven live by `TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss`, which kills a real authenticated channel server, re-serves, and asserts the facade trace `[relay.project, relay.read, relay.read]` — one retry, not two, not zero.
6. **R7-INV(4):** `scrubError` retained on the failure path; the typed second-failure test present; battery green.
7. **R7-INV(5):** `git diff --stat` confirms exactly FOUR MCP files changed (`mcp.go`, `mcp_test.go`, `schema.go`, `schema_test.go`); `errors.go`, `main.go`, `README.md` untouched — under the seven-file grant, minimum surface held.
8. **R7-FACADE:** `internal/seatclient/conduct` exported surface read at source — `Transport` interface, `New`, `FromAuthenticated(authenticated *channel.Client)`; NO credential parameter, field, log, or wire member; the credential census over the package returns only the two comments STATING the exclusion.
9. **R7-EVID(a):** the choreography leg (item 5) — real close→re-auth→one-retry THROUGH the facade.
10. **R7-EVID(b) reachability:** `TestH16MCPRejectsForgedSystemHeadersBeforeConductorCall` asserts the typed `schema_invalid` AND `server.client == nil` (the guard fires BEFORE any dial) across three system-owned headers × non-empty/present-empty (the totality lesson), plus the phase/tier-drift bypass attempt; the parity test additionally asserts the guard produces ZERO facade calls post-refactor — the gate demonstrably runs on the refactored path.
11. **R7-EVID(c) parity:** `TestNativeAndMCPUseSameConductFacadePayloadAndH16StillGates` drives the SAME payload through the MCP frontend and the native `relaytool`→`conduct.FromAuthenticated` path and asserts facade-call NAME AND ARG BYTES equal.
12. **Fence + sacred paths:** `git diff --name-only b7f406b2..HEAD` reduces to exactly the four in-fence trees (`internal/worker/**`, `internal/seatclient/**`, `cmd/frank-worker/**`, `cmd/frank-mcp/{the four}`); `go.mod`/`go.sum`/`internal/channel/**` zero rows; `master/**` (incl. the frozen exit corpus) zero rows.
13. **Traceability + anti-vacuity:** `internal/worker/FIXTURES.md` — 43 table rows, a per-family "deliberately broken / anti-vacuity proof" COLUMN, seven non-silent PARTIAL/HELD/OWNER-EXTERNAL/SUPERSEDED dispositions, the escaped-`setsid` containment named as the ratified residual never a pass claim.
14. **The closed record union:** `TestClosedRecordUnionRejectsUnknownKindAndMembers` REJECTS the K6-excluded `reasoning_replay` kind and a secret-shaped unknown member at `MarshalRecord` — the delta §1.2 exclusions enforced at construction, in schema vocabulary per the 215942 guidance.

## §2 — GRILL_LOCK / charter spot-checks

Uniform F59 gate: the executor carve-out sweep (`skip-authorize/bypass/allowlist/exempt`) returns zero hits outside tests. H-12 posture: no sandbox claim surfaced; the bash claims stay doc-only with the named residual. The three T14 READMEs exist in-fence and state Branch A operative, no Branch-B exclusions — consistent with the granted ruling.

## §3 — The five-invariant attestation (R7 §1's named verification)

All five R7-INV verified by this reviewer at bytes/tests as itemized in §1(3)-(7). No deviation found; the stop-line was never re-fired and did not need to be.

## §4 — One observation, NON-BLOCKING

The submit parse path moved from raw `json.Unmarshal`→`errProtocol` to `formschema.ParseSubmitArguments`→`errSchemaInvalid` — a fail-closed behavior delta of exactly the class R3 pre-authorized ("fail-closed deltas each with its own RED test, expressly NOT claimed behavior-preserving"), covered by the P-rule/V-partition batteries. Recorded so the delta is named, not silent; no action owed.

## §5 — What follows (none of it this relay's authority)

The pair's transport return to master files alongside this review (same act), requesting master's F.7.2 verbatim store export into `frank/.relays/s15/store-export/` and its receipt. Then: the serialized s13 → s14 → s15 restack + battery rerun + re-review, and the operator-only MERGE-GATE. SLICE-CLOSE stays expressly OUTSTANDING until master's receipt cites the populated export.

ACTIONS_GIT_REF: lane governance only — this end-review + the return SITREP + two INDEX rows, UNCOMMITTED on main (master's banking cadence); zero source/test bytes; the implementation worktree untouched by this review (read-only commands only).
FINAL_GIT_STATUS_SHORT: (literal, frank/ cwd, captured after both files + INDEX rows existed:)
  M .relays/s14/INDEX.md
  M .relays/s15/INDEX.md
  M ../master/relays/CHECKPOINTS.md
 ?? .relays/s14/s14-build/PLAN-planner-20260821-145914.md
 ?? .relays/s15/s15-end-review/
 ?? .relays/s15/s15-slice-return/
