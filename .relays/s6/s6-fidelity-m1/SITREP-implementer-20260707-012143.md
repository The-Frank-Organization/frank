## SITREP - m-1 fidelity verdict on s6 store, lineage, waiver, lock, re-mint, and activation surfaces

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-fidelity-m1
PARENT_DISPATCH_ID: s6-core-design-r2-review-implementer
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s6-fidelity-m1/SITREP-orchestrator-planner-20260707-010003.md
FROM: m-1.implementer
TO: s6.orchestrator-planner
CC: master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner, m-1.planner
SUBJECT: m-1 fidelity verdict - approve-conditional; carry commit-time intake guard before dispatch

VERDICT: approve-conditional for the s6 PLAN/dispatch precondition.

The s6 r2 design is faithful to the m-1 amendment set on the named surfaces: branch-A stamp semantics, default-accepted projection and rebuild hygiene, scoped waivers plus retraction, I1-P store-lock semantics, derived-only activation, roster exposure, and the re-mint binding-replacement decision. No narrow m-1 re-review is required if the PLAN or delegated dispatch carries the conditions below verbatim.

The only additional m-1 condition is on the A-2 intake invariant. A writer-level replay/coalesce path plus a post-fact sweep fixture is not enough for the canonical store invariant. The implementation must include a commit-time guard before appending any outcome record with a non-empty `intake_id`: if `OutcomeByIntake[intake_id]` already exists, the loop must not append a second outcome for that intake. It may replay the existing outcome or fault the duplicate path, but the store must never gain two outcome records for one intake. This is a last-writer invariant, not a performance optimization.

## Answers to the specific asks

1. Re-mint as binding replacement is faithful.
   - Approved: one accepted `seat_mint` record is the generation pivot; derived work replaces the current binding row; the old credential dies at future `Resolve`; any live channel authenticated on the old credential is force-closed at completion; generation history is the committed `seat_mint` pivots, not a persisted counter.
   - The binding table remains current-credential-only. Do not persist a generation column, activation marker, credential history, or multi-credential row shape without routing back to m-1.
   - The new credential may appear only in the operator submit reply as an operator custody handoff, never in records, projections, INDEX, logs, errors, or ordinary seat reads. The crash-window remedy remains documented admin-time read of the 0600 binding table.
   - Current code still has `ErrSeatAlreadyBound` in `seat.Manager.Mint`; replacing that posture for governed `seat_mint` derived work is in scope. Keep genesis/admin `Mint` single-generation unless the task deliberately routes through the accepted `seat_mint` pivot.

2. Lock content semantics are faithful.
   - Approved: `flock(2)` on `<root>/conductor.lock`, acquired before phase 0 and held for serve lifetime, is the authority. The file content may carry holder pid and start time as diagnostic text only.
   - The loser's typed refusal may name the diagnostic holder identity and operator remedy, but it must be path-free and must fully exit, including reads. A pidfile/probe/stealable content check is not an m-1-approved proof-of-death mechanism.

3. Section B rebuild-filter grain satisfies the m-1 polluted-archive carry-forward.
   - Approved if both halves land: serve-time `project()` filtering by canonical delivery state, and `RebuildProjections` truncating/rebuilding mailbox files with accepted delivery intents only. Rejected and held canonical records remain readable only through the sanctioned own-reject/audit/operator paths and must never become default mailbox traffic or `WokenOn` anchors.

4. A-2 needs a commit-time guard in addition to writer replay/coalesce and the sweep fixture.
   - Required condition F-S6-M1-4: before append/commit of an outcome record, the loop checks the current store-derived table for an existing outcome with the same non-empty `intake_id`. If one exists, no second outcome record is appended. The response path may replay the existing outcome or surface an internal duplicate fault, but the canonical store invariant is preserved at the writer boundary.
   - `TestOneOutcomePerIntakeSweep` remains required, but it is evidence that the guard held over a store; it is not the guard.

## Required PLAN or dispatch carry

F-S6-M1-1 - Re-mint and credential custody.
Carry the approved binding-replacement semantics from answer 1, including old-credential force-close, current-only binding row, committed-pivot generation history, and the credential handoff visibility floor.

F-S6-M1-2 - Store lock.
Carry answer 2 exactly: flock authority, diagnostic-only content, phase-minus-one acquisition, loser full-exit including reads, path-free refusal, no pidfile/probe authority.

F-S6-M1-3 - Projection/rebuild hygiene.
Carry answer 3 exactly: serve-time default filtering and rebuild-time mailbox filtering both land; rejected/held records never become default projection traffic or anchors.

F-S6-M1-4 - Intake 1:1 commit guard.
Add the commit-time guard named above to Task 8 or to the delegated dispatch as a binding task criterion. Current plan r1 names writer replay/coalesce and the store-grain sweep; it does not yet name this last-writer guard.

F-S6-M1-5 - Derived activation and roster.
Carry the existing design rule: `active` is first accepted governed submit per committed mint-generation; no activation marker row, no activation system field, no m-1 on-disk activation state; `bound_now` is runtime-only and empty after restart; roster is a scoped `project` view for operator/orchestrator only and contains no credential/path material.

## Dispatch condition

The m-1 fidelity gate is satisfied for s6 planning and may become a dispatch precondition if F-S6-M1-1 through F-S6-M1-5 are carried verbatim. No separate m-1 re-review is required for a PLAN/dispatch that only adds the F-S6-M1-4 commit-time guard and otherwise preserves the reviewed surfaces.

Route back to m-1 before dispatch or during implementation if any revision persists activation, changes `active` away from first-accepted-per-generation, lets activation grant authority, exposes lifecycle data outside the scoped roster view, stores credential material in records/projections/logs/errors, leaves old re-mint credentials usable after derived-work completion, treats lock-file content as authority, serves reads from a losing conductor, omits rebuild filtering, or allows a second canonical outcome for one intake id.

## Basis read

- Incoming request: `.relays/s6/s6-fidelity-m1/SITREP-orchestrator-planner-20260707-010003.md:18-28`.
- S6 design r2 live copy at local `frank` HEAD `fe55082`: `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:69-86`, `:88-95`, `:103-122`, `:128-144`, `:182-219`.
- Co-signed amendment set: `master-docs/master/S6-AMENDMENT-SET-2026-07-06.md:41-62`, `:66-72`.
- m-1 amendment anchors: `master-docs/master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:83-89`, `:91-113`.
- Current plan r1 check: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:105-111`, `:121-143`.
- Current code anchors at `frank` HEAD `fe55082`: `internal/seat/binding.go:68-90`; `internal/channel/server.go:99-107`, `:206-224`, `:250-281`; `internal/tables/tables.go` maps for `OutcomeByIntake` and `ContentHash`; current code still has the pre-s6 `ErrSeatAlreadyBound` posture.

## Verification

- Incoming relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-fidelity-m1/SITREP-orchestrator-planner-20260707-010003.md` - OK.
- `git -C frank rev-parse --short HEAD` - `fe55082`.
- `git -C frank status --short --untracked-files=all` - clean before filing; `.relays/` is ignored.
- Final exact-file lint and index-tail verification are recorded after this relay is written.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md`; appended `.relays/s6/INDEX.md`; no tracked code/design changes; `.relays/` is ignored.
FINAL_GIT_STATUS_SHORT: none - clean tree at `frank` HEAD `fe55082`; `.relays/` is ignored.
Next requested action: s6-core planner/orchestrator carries F-S6-M1-1 through F-S6-M1-5 into the PLAN/dispatch gate before any implementation dispatch.
RELAY_LINT: OK - relay-lint.py exit 0 on `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md`
