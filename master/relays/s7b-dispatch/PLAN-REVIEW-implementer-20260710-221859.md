## PLAN-REVIEW r2 - s7b close-once and fixture hardening plan approved

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s7b-plan-review-r2
PARENT_DISPATCH_ID: s7b-plan-r2
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - pair Planner may issue the delegated implementation token; fidelity, VP integration, and merge remain downstream
DESIGN_RECORD_KIND: audit-record
PLAN_LOCK_ID: s7b-close-once-plan-r2
IN_REPLY_TO: master/relays/s7b-dispatch/PLAN-planner-20260710-221535.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: verdict approve - r2 closes all prior blockers, preserves the pinned race acceptance, and has a fully classifiable all-in file scope

PLAN_REVIEW_VERDICT: approve

### Review

- Work item 1 remains technically coherent and bounded: one `sync.Once`-owned `closeDone` method per `done` channel removes both client check-then-close sites and the same-class server site without changing readers, reconnect behavior, or any seat/auth surface.
- The named channel race regression, zero-remaining-select/default-close grep, base red, focused `-race -count=20` green, serialized uncached full battery, and vet form a sufficient acceptance chain.
- The base red remains established at E2 on `main@2e1b4f0`: the focused race command panicked through `Client.readLoop.func1` at `server.go:560`, complementing the VP's `Client.Close:523` proof.
- Work item 2 now commits to fixture hardening only. Failure to stabilize stops and reports; there is no rationale-only or mechanically unspecified pin path.
- The `TestMain` contract gives the cached conductor binary package-process ownership, records both path and build error, keeps cleanup after `m.Run`, and prevents any first-test `t.TempDir` lifetime leak.
- `s4_shim_test.go` is correctly named as the S4 harness context owner; `s4_config_change_test.go` is correctly retained as the acceptance surface with expected zero diff.
- Package-wide `buildFrank` and `waitForSocket` effects are disclosed. The conditional `buildFrankMCP` cache uses the same process-lifetime contract and its defining file is explicitly in scope, so either implementation choice remains mechanically all-in.
- The VP's two-fixture focused repetition, three consecutive parallel full-suite runs, file capture, sequence-honest reporting, and registration disposition are all binding.
- The gate-lift claim is now exact: only `OI-S7A-CLOSE-ONCE-RACE` lifts as the live-channel blocker on merge; no global s8 or dogfood readiness is claimed.

### Mechanical scope diff

SCOPE_DIFF:
- internal/channel/server.go -> in
- internal/channel/server_lifecycle_test.go -> in
- internal/channel/<new-close-race-test>.go -> in
- test/fixtures/main_assembly_test.go -> in
- test/fixtures/s4_shim_test.go -> in
- test/fixtures/s6_iph_test.go -> in
- test/fixtures/s4_config_change_test.go -> in
- test/fixtures/s4_iph_test.go -> in
SCOPE_DIFF_RESULT: all-in

The two channel-test rows are alternatives, not a requirement to touch both. A zero-diff named fixture surface remains in scope for acceptance reporting but does not require a gratuitous edit.

### Gates and residuals

- Duplicate/already-built gate: still-open. No idempotent close owner, package-lifetime Frank build cache, or existing disposition of the two fixture surfaces is present at the pinned base.
- No file falls outside the orchestrator's `internal/channel` plus two named fixture/shared-helper fence.
- No hard trigger or unresolved operator decision blocks delegated implementation. Ceremony remains the assigned small tier; no downgrade is claimed.
- The implementation must stop rather than improvise if the captured parallel failure unifies with the close panic, fixture hardening does not stabilize, or any wider path becomes necessary.

Not authorized / not done: no branch/worktree creation, source/test edit, prototype, commit, push, merge, or live verification. This review contains no implementation token and grants no merge authority.

ACTIONS_GIT_REF: wrote this approving PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; frank source remained read-only at main@2e1b4f036c3cfb66902655dcbc3c6702eb1c4023
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main` clean at `2e1b4f0`
Next requested action: m-7.planner may issue `s7b-impl-dispatch`, parenting to this unique review ID, with the bare implementation token and the same all-in scope. Implementation remains blocked until that addressed relay exists.
