## REVIEW-FOLD - s4-wire-impl fold round 1 report

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-plan-review-implementer-r3
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no new gate; merge/s4-close and operator veto path remain separate
DESIGN_LOCK_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s4-grill-s4-wire
PLAN_LOCK_ID: s4-slice-4-plan
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator
IN_REPLY_TO: s4-wire-impl/SITREP-planner-20260705-141524.md
BASE: s4-wire-impl@796b7be
BRANCH: s4-wire-impl
HEAD: 7dc5f92
SUBJECT: fold round 1 complete - MF-1..MF-4 addressed on s4-wire-impl; final E2/race gates green; no merge authority

FOLD_SCOPE:
- internal/channel/server.go -> in
- internal/channel/server_push_test.go -> in
- internal/recover/recover_test.go -> in
- internal/store/genesis.go -> in
- internal/store/projections.go -> in
- internal/store/project_test.go -> in
- test/fixtures/s4_shim_test.go -> in
- docs/ops.md -> in
- docs/sprints/2026-07-05-s4-slice-4/results/e3-gate-procedure.md -> in
FOLD_SCOPE_RESULT: all-in

## Summary

Folded the four must-fix findings from `SITREP-planner-20260705-141524.md` on the existing `s4-wire-impl` branch. Commit: `7dc5f92 s4: fold r1 must-fixes -- push deadline, config self-heal, canonical TO/CC [MF-1..4]`.

No merge authority is claimed. Operator veto path, s4-close, and merge remain separate gates.

## Finding Disposition

- MF-1: addressed. `writePushes` now uses a push-only write path with a 100ms write deadline; timed-out pushes are dropped best-effort and mailbox truth remains authoritative. Red-first fixture: `TestPushWriteDropsStalledRecipient`.
- MF-2: addressed. Phase-0 rematerializes the latest accepted `config_change` per approved member across the chain, using config intents only. Red-first fixture: `TestPhase0RematerializesLatestConfigChangePerMember`; existing fail-closed mismatch leg still passes.
- MF-3: addressed. Delivery recipients now parse TO/CC through canonical `fieldspec.ParseTyped` address-list parsing; any non-canonical header carriage falls back to `Envelope.To` only. Red-first fixture: `TestDeliveryRecipientsRequireCanonicalAddressLists`. `internal/store/projections.go` and `internal/store/project_test.go` are in via absorption ruling `s4-wire-impl/RECONCILE-orchestrator-planner-20260705-140849.md`.
- MF-4: addressed. Added CC recipient mailbox/nudge/dedupe/path-clean assertion surface and the stated-not-silent delivery semantics line to `docs/ops.md` and `docs/sprints/2026-07-05-s4-slice-4/results/e3-gate-procedure.md`.

Optional findings were not folded except where the must-fix implementation naturally avoided side effects: MF-2 replays config intents only, so it does not re-emit INDEX/mailbox projection intents during rematerialization.

## Red-First Evidence

- `go test -count=1 ./internal/channel -run TestPushWriteDropsStalledRecipient` failed before MF-1 with `writePushes blocked on a stalled recipient`; passed after the push deadline.
- `go test -count=1 ./internal/recover -run TestPhase0RematerializesLatestConfigChangePerMember` failed before MF-2 with diagnostics instead of Ready; passed after per-member rematerialization.
- `go test -count=1 ./internal/store -run TestDeliveryRecipientsRequireCanonicalAddressLists` failed before MF-3 with recipients widened from non-canonical CC; passed after canonical parsing.
- MF-4 is assertion/doc coverage over the already-ratified TO/CC behavior; focused fixture and statement grep passed.

## Final Verification

Fresh final-tip commands from `~/.config/superpowers/worktrees/frank/s4-wire-impl` at `7dc5f92`:

- `go test -count=1 ./...` -> pass.
- `go vet ./...` -> pass, no output.
- `go test -race -count=1 ./internal/channel ./internal/engine ./internal/store` -> pass.
- `go test -count=1 ./internal/channel -run 'TestPushWriteDropsStalledRecipient|TestSecondConnectSameCredentialRejected|TestPushToAuthenticatedSeatOnly' && go test -count=1 ./internal/recover -run 'TestPhase0RematerializesLatestConfigChangePerMember|TestPhase0PersistentMismatchStillDiagnostics|TestPhase0RematerializesFromChain' && go test -count=1 ./internal/store -run 'TestDeliveryRecipientsRequireCanonicalAddressLists|TestPendingDeliveryForUsesDurableRecipientMailboxes' && go test -count=1 ./test/fixtures -run 'TestCCRecipientMailboxedNudgedOnceAndPathClean|TestNudgeAllRecipients|TestNoCrossSeatMetadata'` -> pass.
- `git diff --check` -> pass, no output.
- `rg -n 'S4 corrects delivery to the locked TO/CC-mailbox semantics \(m-1 §5\); S1–S3 delivered to Envelope\.To only\.' docs/ops.md docs/sprints/2026-07-05-s4-slice-4/results/e3-gate-procedure.md` -> two hits, one per required surface.
- `git status --short --branch` -> `## s4-wire-impl`.

## Evidence Levels

Implementation and verification evidence is E2. The E3 live two-host gate remains operator-run per the existing gate procedure.

ACTIONS_GIT_REF: branch `s4-wire-impl` at `7dc5f92`; fold commit `7dc5f92`; relay-substrate writes are `REVIEW-FOLD-implementer-20260705-142253.md`, this report, and `.relays/s4/INDEX.md` rows.
FINAL_GIT_STATUS_SHORT: code worktree clean at `## s4-wire-impl`; main checkout tracked status clean, with expected gitignored `.relays/` writes.
