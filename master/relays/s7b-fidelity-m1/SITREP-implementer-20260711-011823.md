## SITREP — m-1 fidelity verdict on s7b close ordering at `e155aa6`: CONFIRM, `sync.Once` preserves channel and credential lifecycle semantics

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7b-fidelity-m1
PARENT_DISPATCH_ID: s7b-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — scoped m-1 fidelity return; VP integration and operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7b-fidelity-m1/SITREP-orchestrator-planner-20260711-010840.md
FROM: m-1.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-1.planner, m-7.planner, m-7.implementer
SUBJECT: confirm s7b close-once fix preserves B-3 bind/rebind, §8.5 re-attach, one-active-channel, and superseded-credential force-close semantics at s7b-close-once@e155aa6

**VERDICT: CONFIRM.** The `sync.Once` close ownership introduced by `a2a6966` removes duplicate channel-close ownership without changing the observable ordering that protects authentication, re-attach, or credential supersession.

### Contract fidelity

1. **Client close ordering is unchanged.** Before the fix, `Client.Close` closed `c.done` through select/default and then called `c.conn.Close`; after the fix it calls `c.closeDone()` and then `c.conn.Close()` (`internal/channel/server.go:521-528`). `readLoop` still signals `done` only from its defer after the read loop exits (`:557-589`). `sync.Once` changes duplicate ownership into idempotent ownership; it does not move the signal across the socket close or create another lifecycle transition.

2. **No authentication window opens.** Authentication and the one-active-channel guard remain server-owned and byte-identical: credential resolution, credential-hash lookup, `auth:channel-active`, and active-map insertion occur under `Server.mu` (`server.go:277-302`). A client-side `done` close cannot remove that entry. The entry is removed only by `serverConn.run`'s existing defer after the connection read loop exits (`:244-253`). Therefore the interval between client `done` and server cleanup stays fail-closed: a same-credential reconnect is rejected until the old socket is proven dead, then re-authentication can reoccupy the seat.

3. **B-3 bind/rebind semantics are preserved.** The locked m-1 contract says `bound` means current credential possession over a live authenticated channel, every reconnect re-binds, and reconnect never re-boots (`master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:95-103,108`). The fixed code changes no connect handler, lifecycle table, roster derivation, or boot gate. The focused station proves second live use of one credential remains rejected and proven-dead reconnect succeeds.

4. **§8.5 re-attach and supersession are preserved.** Re-attach still presents the same credential and resolves to the same seat without re-mint (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:127-128,141,151`). The remint path remains byte-identical: `MintOrReplace` is followed by `Server.ForceCloseSeat`, which closes the seat's authenticated socket; server cleanup removes its active credential entry (`cmd/frank/main.go:515-533`; `server.go:182-193,244-253`). The superseded-generation gate and binding manager are untouched.

5. **Server shutdown does not weaken admission.** `Server.Close` still signals server `done` before closing the listener and client sockets; `sync.Once` only makes that signal idempotent (`server.go:217-231`). During shutdown no new listener admission is opened, and restart begins with runtime `bound = empty` as already locked.

### Scope and remaining gates

- Fix commit `a2a6966` changes only `internal/channel/server.go` and `internal/channel/new_close_race_test.go`.
- `git diff --exit-code 2e1b4f0..e155aa6` over `cmd/frank/main.go`, `internal/seat`, `internal/engine/loop.go`, lifecycle tests, and the S6 lifecycle fixture is clean; the identity/auth-generation/remint paths did not change.
- This verdict lifts only the m-1 fidelity question. VP integration and the operator merge gate remain required; no broader s8 or live-channel claim is made here.

### Verification

- `go test -race -count=20 ./internal/channel -run '^(TestClientCloseIdempotentUnderRace|TestSecondConnectSameCredentialRejected|TestProvenDeadRecovery|TestForceCloseSeatDropsAuthenticatedChannel)$'` — PASS.
- `go test -count=1 ./test/fixtures -run '^(TestS6LiveMintBootActivationAndRoster|TestStaleNonSubmitRefusalIsNotLifecycleGating)$' -v` — PASS.
- `go test -count=1 ./test/invariants` — PASS.
- `go vet ./...` — PASS, exit 0 with no output.
- `git diff --check 2e1b4f0..e155aa6` — PASS.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s7b-fidelity-m1/SITREP-implementer-20260711-011823.md` — OK.

ACTIONS_GIT_REF: no `frank` code or git action by this report-only fidelity review; durable artifacts are `master/relays/s7b-fidelity-m1/SITREP-implementer-20260711-011823.md` and its append-only `master/relays/INDEX.md` row.
FINAL_GIT_STATUS_SHORT:
- workspace root: unavailable — `fatal: not a git repository (or any of the parent directories): .git`
- `frank/`: none — clean tree
- pinned worktree `s7b-close-once@e155aa6`:
  `?? .relays/s7b/INDEX.md`
  `?? .relays/s7b/s7b-close-once/IMPL-implementer-20260710-224959.md`
  `?? .relays/s7b/s7b-close-once/IMPL-implementer-20260711-003443.md`
  `?? .relays/s7b/s7b-close-once/RECONCILE-planner-20260711-004630.md`
Next requested action: accept the m-1 confirm into s7b integration and route the completed package to the VP. No merge authority is present here.
