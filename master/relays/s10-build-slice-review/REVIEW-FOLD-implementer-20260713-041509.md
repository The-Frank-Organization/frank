## REVIEW-FOLD — second pre-merge fold complete: MF-2 through MF-8 and AO-2/AO-3 are RED→GREEN in one bounded commit; full uncached and race batteries green; clean head; no merge

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s10-build-slice-review
PARENT_DISPATCH_ID: s10-build-slice-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the pair-internal fold is complete; the operator's merge-decision go-ahead still requires this fold's planner re-verification, and merge remains operator-authorized only
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
IN_REPLY_TO: s10-build-slice-review/REVIEW-FOLD-planner-20260713-034634.md
FROM: s10.implementer
TO: s10.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the operator-ordered second-review fold is complete in one attributable commit `39474d0` over `f481042`; all seven must-fixes and both accepted-optional hardenings have behavioral RED evidence, focused GREEN evidence, race coverage for the concurrency seams, and two final full uncached green runs (the second post-commit); the initial full run correctly found a path-hygiene sink-census integration issue, repaired in-fence without changing the catalog; request planner re-verification, not merge

FOLD_SCOPE:
- frank/internal/store/genesis.go -> in
- frank/internal/engine/loop.go -> in
- frank/internal/engine/expiry.go -> in
- frank/internal/engine/approval.go -> in
- frank/internal/engine/resummon.go -> in
- frank/internal/engine/submit.go -> in
- frank/internal/observe/read_file_worker.go -> in
- frank/internal/observe/registry.go -> in
- frank/internal/executor/executor.go -> in
- frank/internal/obligation/obligation.go -> in
- frank/cmd/frank/main.go -> in
- frank/test/fixtures/ -> in
- frank/internal/engine/ -> in
- frank/.relays/s10/ -> in
FOLD_SCOPE_RESULT: all-in

The required pre-edit artifact is `.relays/s10/fold2-scope-20260713-034905.md`. It existed before every second-fold source/test edit and contains the exact block above. No OUT discovery occurred.

## Findings folded

- **MF-2 — v8 genesis integrity:** v8 now peels count-guarded through exactly its marker and `record_kind` delta into v7, then through the existing v7→v6→pinned-v5 chain. The pinned-v5 SHA remains byte-unchanged. A locked-byte corruption fails genesis; pristine v8 still materializes as v8 after predecessor validation.
- **MF-3 — commit-time snapshot guards:** the serialized loop rebuilds the durable table immediately before relevant commits and rechecks `resolves_gate`, `owed_disposition`, `waiver_retraction`, pinned config digest/transition, and `seat_mint` validation. The nested-service race now produces exactly one accepted gate resolution and one typed `resolves_gate:already-resolved` rejection. Seat remint is explicitly tested as the existing latest-wins rule, not invented uniqueness.
- **MF-4 — completed read precedence:** both late-kill and hard-ceiling arms drain an already-ready worker result, call `finishReadFile`, return the real pass, and leave the lane breaker closed.
- **MF-5 — completed suite precedence:** executor deliberation watches `waited`; a finished suite returns its actual result even when a late kill or ceiling event is also selectable. Genuinely running kill/ceiling behavior is unchanged.
- **MF-6 — live resummon hosting:** the production composition root constructs `ResummonScheduler`, arms after startup and each completed turn, and hosts the two G4 classes over durable parked gates. No-response emits local; an answered-but-stalled attempt emits louder-local; accepted resolution cancels by durable-state check; both use stable cadence slots and A-2 content-hash dedupe. The current explicit default is one hour for each class. Production timers are `time.NewTimer` instances stopped on context cancellation; the injectable channel seam remains test-only.
- **MF-7 — scoped registry fault:** the registry loads once per `CompleteAuto` pass. Explicit HUMAN/egress gates still complete when the file is corrupt; ambiguous category-only gates fail closed. The discovering pass records one typed `held-gate-registry-unavailable` fault and returns `ErrGateRegistry`; later turns do not wedge or duplicate the fault, and repaired config can classify the skipped gates on a future pass.
- **MF-8 — root cancellation:** read expiry derives from `RegistryEnv.Context`; executor config now carries the production root context. Cancel during a pending prompt resolves promptly through conservative timeout/kill behavior rather than waiting to the hard ceiling.
- **AO-2 — duplicate prompters:** approval and expiry prompters now attach same-gate callers to one broadcast pending result. One durable prompt submission wakes every attached waiter; owner failure/cancel fails the shared result closed.
- **AO-3 — nested reply abandonment:** nested `ReplyCh` delivery now has the same bounded timeout escape as the top-level loop, so an abandoned receiver cannot deadlock the serialized writer.

## RED, integration repair, and GREEN evidence

- `.relays/s10/fold2-red-20260713-040702.txt` records the behavioral REDs for every MF/AO. MF-4/MF-5 and the read half of MF-8 were additionally rerun against a disposable detached `f481042` worktree with only their tests applied; that worktree was removed after capture.
- `.relays/s10/fold2-targeted-green-20260713-040702.txt` records all focused regressions, `git diff --check`, and race runs. Post-commit race reruns also passed: engine 1.740s; fixtures 1.958s.
- `.relays/s10/fold2-full-red-repair-20260713-040702.txt` records the first integrated battery's useful RED: the two new direct `bounce.Format` sites made `TestLawPathHygiene` report 21 sinks vs 19 pins, which caused the production dogfood E2 suite to reject. Both sites now reuse the existing typed formatter; the invariant catalog and every OUT file stayed untouched. The invariant and production-dogfood tests then passed independently.
- `.relays/s10/fold2-full-green-20260713-040702.txt` records the final pre-commit `go test ./... -count=1` success: fixtures 141.168s, invariants 3.959s, replay 0.881s, dogfood 0.929s, zeroloss 1.919s.
- Fresh post-commit `go test ./... -count=1` at `39474d0` also exited 0: fixtures 137.941s, invariants 4.205s, replay 1.142s, dogfood 1.171s, zeroloss 2.126s. `git diff --check HEAD^ HEAD` exited 0 and the final worktree status is clean.

## Boundary, owner bytes, and named carries

`internal/executor/executor.go`'s `Spawn` admission/class/timeout block remains byte-untouched; only config context and the licensed expiry deliberation seam changed. `cmd/frank/main.go` changes are confined to the executor-host and loop-construction composition blocks. The pinned-v5 SHA, FX-CFG-7, the v8 owner bytes, terminal enum, locks, and merge state are untouched.

The dispatch's carries remain open and unchanged: non-Accepted-resolution/blocked-prompter re-prompt semantics plus the claimless authority-floor Held edge route to the s11 m-3/m-6 design cell; pinned run-verdict cache invalidation remains m-7's A-2 ledger item; and the s11 cleanup card remains the generic prompter, shared expiry arbiter, one ODB builder, table-snapshot prompter lookup, per-emit `tables.Build` removal, shared system→operator envelope builder, ContentHash/GateID decoupling, `finalizeRun` cleanup, and genesis reverse-ladder cleanup.

No merge, push, tag, or branch cleanup was performed. This seat still has no merge authority.

ACTIONS_GIT_REF: `s10-comms-spine@39474d06ef1ea90b539efec4de1f194bb9c8a261`; one second-review fold commit `39474d0` (`fix: harden s10 pre-merge concurrency`) over reviewed head `f48104261841809eb21f2ec5f4b73631ae4ce341`; exact base remains `8941889a432a0a02a4183a16fb0015c230c9561b`
FINAL_GIT_STATUS_SHORT: none - clean tree

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s10-build-slice-review/REVIEW-FOLD-implementer-20260713-041509.md` — run before handoff.
- `git -C /Users/jack/.config/superpowers/worktrees/frank/s10-comms-spine status --short` — no output after the post-commit full battery.

Next requested action: s10.planner adversarially re-verifies this bounded second fold, then updates/executes the already-held merge-decision go-ahead. Merge itself remains blocked until the separate valid HUMAN_MERGE_AUTHORIZATION reaches the authorized merge seat.
