## PLAN-REVIEW - s7b close-once mechanism sound; fixture disposition must name lifecycle, actual timeout owner, and mechanical fallback

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s7b-plan-review
PARENT_DISPATCH_ID: s7b-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - pair Planner revision next; implementation, fidelity, VP integration, and merge remain downstream
DESIGN_RECORD_KIND: audit-record
PLAN_LOCK_ID: s7b-close-once-plan
IN_REPLY_TO: master/relays/s7b-dispatch/PLAN-planner-20260710-215733.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: verdict must-revise - approve the close-once mechanism in substance, but repair three fixture-plan gaps before delegated dispatch; no implementation token accepted or requested

PLAN_REVIEW_VERDICT: must-revise

### Findings

1. **BLOCKER - the fallback pin is not mechanically specified, so the plan cannot produce the required all-in scope diff.** The VP condition in `RECONCILE-orchestrator-reviewer-20260710-201732.md` requires a checked-in canonical executable target/script/CI entry that forces the selected parallelism and is then exercised. The plan says only "fall back to the pin WITH the recorded rationale" and names no artifact, command, or consumer. A rationale-only fallback is explicitly insufficient. Revise the plan to either remove the fallback and commit to fixture hardening, or name the exact checked-in pin file, invocation, assertion that consumers use it, and scope disposition. If that file is outside the orchestrator's two-fixture/shared-helper fence, route the deviation to master before any delegated token.

2. **BLOCKER - package-level `sync.Once` caching lacks a process-lifetime path and cleanup owner.** Current `buildFrank` builds to `filepath.Join(t.TempDir(), "frank")` (`test/fixtures/main_assembly_test.go:702-709`). Caching that path at package scope while retaining `t.TempDir` makes the cache invalid as soon as the first owning test cleans up. Revise with the exact lifecycle contract: a package-process-owned temporary directory, one build result/error, and cleanup after the package run (for example an explicit `TestMain` owner), with no first-test `t.Cleanup` ownership of the cached binary. Include failure behavior so a failed first build cannot yield an empty/stale path to later tests.

3. **BLOCKER - the declared S4 target does not own the timeout the plan says to raise.** `TestConfigChangeProjectionsCarryNoMemberBytes` calls `newS4ShimHarness`, but its 8-second context and `buildFrank` call live in `test/fixtures/s4_shim_test.go:225-250`; `s4_config_change_test.go` has neither. To meet the plan's own context `>=30s` requirement, add `test/fixtures/s4_shim_test.go` to the exact file targets and say whether `s4_config_change_test.go` changes at all. This shared-helper path is within the orchestrator fence, so it is a plan correction, not a scope expansion.

4. **MUST CLARIFY - carry the VP's gate-lift boundary literally.** A successful s7b merge lifts only `OI-S7A-CLOSE-ONCE-RACE` as the live-channel blocker. It does not declare s8 or dogfood globally ready; design, genesis/config, implementation, and operator gates remain. The current plan gestures at this but its return-path sentence says the live-channel gate lifts without the explicit non-global boundary.

### Accepted plan portions

- The production defect is still open and exactly located: `Client.Close` and `Client.readLoop` independently select/default-close `c.done`; `Server.Close` has the same latent double-call class.
- `sync.Once` plus one unexported close owner per `done` channel is coherent, minimal, and preserves reader/reconnect semantics.
- The server leg is in-fence and same-class; retain it unless the revised plan gives a technical reason to sever it.
- The named channel regression, focused reconnect race repetition, zero-remaining-idiom grep, serialized uncached full battery, vet, file-captured evidence, and sequence-honest reporting are appropriate.
- Work-item separation is correct: fixture startup remains distinct unless captured failure text proves unification.

### Fresh evidence at `main@2e1b4f0`

- `go test -race -count=20 ./cmd/frank-mcp -run '^TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss$'` -> FAIL immediately with `panic: close of closed channel` through `internal/channel.(*Client).readLoop.func1`, `server.go:560`. This confirms the red-first base condition remains reproducible; no source mutation was needed.
- `buildFrank` is shared by `main_assembly_test.go`, `s4_shim_test.go`, `s5_wire3_test.go`, `s6_iph_test.go`, `s6_lifecycle_test.go`, `s6_lock_test.go`, `s6_mint_test.go`, and `sweep_test.go`; changing it is a package-wide fixture-runtime contract, even though the accepted flaky surfaces remain the two named tests.
- `waitForSocket` is likewise shared across those files and currently has a fixed 4-second deadline.
- No existing `closeOnce`, package-level Frank binary cache, or `TestMain` closes these items already.

### Mechanical scope diff

SCOPE_DIFF:
- internal/channel/server.go -> in
- internal/channel/server_lifecycle_test.go or a new internal/channel close-race test file -> in
- test/fixtures/main_assembly_test.go -> in
- test/fixtures/s6_iph_test.go -> in
- test/fixtures/s4_config_change_test.go -> in
- test/fixtures/s4_shim_test.go -> in against the orchestrator's shared-helper fence, but absent from the submitted plan target list
- canonical executable parallelism-pin target -> OUT/undiffable because no file path is specified; if selected, name it and re-engage master when outside the fixture/shared-helper fence
SCOPE_DIFF_RESULT: deviation-present

No hard trigger requires operator judgment for the close mechanism itself. The blocker is plan precision and mechanical scope, not product semantics.

Duplicate/already-built gate: still-open. No idempotent close primitive exists on `Client.done` or `Server.done`; no package-lifetime Frank build cache exists; the two fixture flake surfaces remain registered and undispositioned.

Not authorized / not done: no branch/worktree creation, source or test edit, prototype, commit, push, implementation dispatch, merge, or live verification. The bare implementation token remains required after an approving revised PLAN-REVIEW with `SCOPE_DIFF_RESULT: all-in`.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; frank source remained read-only at main@2e1b4f036c3cfb66902655dcbc3c6702eb1c4023
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main` clean at `2e1b4f0`
Next requested action: m-7.planner revises the fixture leg and exact targets, then returns a superseding PLAN for another PLAN-REVIEW. Do not issue the delegated implementation token from this must-revise verdict.
