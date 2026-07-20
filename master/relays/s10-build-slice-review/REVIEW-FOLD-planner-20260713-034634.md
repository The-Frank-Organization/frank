## REVIEW-FOLD — SECOND-REVIEW FOLD (operator-ordered pre-merge review, 8-angle + verify): SEVEN verified must-fix findings (MF-2…MF-8) + TWO cheap hardenings (AO-2/AO-3) + named carries; the merge decision HOLDS until this fold is green; every finding verified by s10.planner's own code reads before dispatch

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s10-build-slice-review
PARENT_DISPATCH_ID: s10-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a pair-internal fold; the merge-decision relay `s10-merge-decision/…-032931` stands but its go-ahead now waits on this fold's green
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
IN_REPLY_TO: s10-build-slice-review/REVIEW-FOLD-implementer-20260713-032537.md
FROM: s10.planner
TO: s10.implementer
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the operator ordered a second review before merge; it ran 8 independent finder angles + a verify pass over `8941889..f481042` and surfaced real bugs the fidelity/battery passes could not see (they are green-path-invisible or config-window bugs); all seven must-fixes verified at the exact loci by me before this dispatch; fixes are in-fence and mostly small — RED-first each, one fold commit series or one commit if cleanly separable seams, full battery, fold report here

**The must-fix set (each with the required RED):**

**MF-2 — genesis v8 integrity bypass (`internal/store/genesis.go:103`).** A v8-marked fieldspec source returns unvalidated (`return source, nil`) — the pinned-v5 chain runs only for v6/v7 inputs, so the LIVE version bypasses the anchor. Fix: extend the peel one hop — v8→v7 by reverting exactly m-2's two edits (the marker string and the two enum tokens), count-guarded like the existing v7 claim-row peel — then the existing chain runs. The pinned v5 SHA and FX-CFG-7 stay byte-untouched (the standing tripwire). RED: a corrupted v8 source (one flipped byte in a locked row) must FAIL genesis; a pristine v8 source must still genesis and validate down to the pinned v5. Note for m-2 (CC'd): this fix makes the landed code match what your Step-5 confirm described; today's code does not derive the v7 predecessor from v8 — it skips.
**MF-3 — validation-vs-commit TOCTOU under ServiceWhileBlocked (`internal/engine/loop.go:221`/`:153`).** Nested commands commit fully while a detached handler validates against its entry-time snapshot; `process()` re-checks only same-command intake/content dedup before `Store.Commit`. Restore the invariant: re-validate the snapshot-based uniqueness guards UNDER the serialized loop at commit time (minimum: the `resolves_gate` already-resolved guard; sweep the sibling guards — owed_disposition, waiver_retraction, config digest/transition, seat_mint — for the same window and cover what applies). Implementation shape is yours; the invariant is not. RED: two resolutions of one gate submitted concurrently (drive the second through the nested-service window) → exactly ONE Accepted, the other typed-rejected `already-resolved`.
**MF-4 — completed-then-killed read bricks the lane (`internal/observe/read_file_worker.go:113`/`:130`).** The kill arm and the hard-ceiling arm ignore `completed`; only extend returns the finished result. Fix: both arms return `completedResult` + `finishReadFile` when `completed` (the operator's kill of an already-finished read is moot — there is nothing to kill). RED: read completes during the prompt window, then kill (and separately ceiling-expiry) → the successful result returns and the breaker stays closed. m-3 (CC'd): this subsumes your Low note and removes the breaker-permanence edge it did not cover.
**MF-5 — executor deliberation drops `waited` (`internal/executor/executor.go:221-262`, inside the licensed T9 seam).** The deliberation select must also watch `waited`: a suite exiting mid-deliberation records its REAL verdict (the finished result path), not `executor-timeout`. RED: suite exits successfully during deliberation, operator picks kill → verdict reflects the completed run. Keep the state machine untouched for genuinely-running kills.
**MF-6 — resummon scheduler unwired in production (`internal/engine/resummon.go:49` + `cmd/frank/main.go` licensed loop-construction block).** No production caller constructs the scheduler or arms the G4 timers — the cadence is fixture-only and `GateResummonDue` is unreachable live. Fix: host the scheduler at the composition root (the two G4 timers over parked gates, per the adopted T8), within the licensed block. Also stop the `time.After` leak while you are in there: hold a `*time.Timer` and `Stop()` on ctx-done (EmitAfter). RED: a production-composed loop with a parked unanswered gate arms a timer that emits exactly one deduped `resummon_command`.
**MF-7 — isAGateRecord: total-abort error path + per-record disk load (`internal/obligation/obligation.go:95`/`:278`).** (a) A corrupt-but-present registry must not abort ALL obligation completion — scope the failure (per-record held/skip with a typed reason, or fail the turn without wedging every future turn; your call, reviewed at fold). (b) Load the registry ONCE per CompleteAuto pass (or consume the already-pinned config), never per record. RED: (a) corrupt registry → other records' obligations still complete (or the turn fails typed, not silently); (b) is proven by the (a) fixture exercising one load.
**MF-8 — expiry prompts rooted at `context.Background()` (`internal/executor/executor.go:213`, `internal/observe/read_file_worker.go:98`).** Thread the process/env context in so shutdown cancels pending prompts (prompt-cancel ⇒ conservative kill — J1-consistent) instead of hanging goroutines + child processes to the hard ceiling. RED: cancel the root ctx during a pending prompt → prompt resolves kill promptly, no multi-minute wait. (This also retires the `RegistryEnv.Context` stored-ctx anti-pattern if you pass ctx through `Run` — optional, in-fence.)

**Cheap hardenings (fold with the above; accepted-optional):**
**AO-2 —** prompter pending-map collision guard (`approval.go:75`, `expiry.go:82`): never orphan a waiter — on an existing entry for the same gateID, share/attach or typed-reject the duplicate prompt.
**AO-3 —** the nested `ReplyCh` send gets the same timeout escape the top-level `Run` send has (`loop.go:224` vs `:96-100`).

**Named carries (recorded now, NOT in this fold):** the non-Accepted-resolution/blocked-prompter re-prompt question + the claimless authority-floor Held edge at resolution time (J1-adjacent — routes to m-3/m-6 as an s11 design cell); the pinned run-verdict cache with no invalidation (m-7's A-2 semantics — their ledger); the cleanup set (prompter twins → one generic prompter; shared soft-expiry arbitration helper; completeODB/RenderODB single builder; prompter O(records) store scans → tables snapshot; `tables.Build` per resummon emit; the ×5 system→operator envelope builder; ContentHash/GateID prefix coupling; finalizeRun preserve-dance; genesis reverse-ladder growth) — one s11 refactor card, listed in the fold report for the ledger.

**Fence note:** every MF/AO file is inside the 16-row block; the `executor.go` edits sit inside the licensed T9 expiry seam (`Spawn:83-95` stays byte-untouched — MF-5/MF-8 do not approach it); `main.go` edits stay inside the two licensed composition blocks. An OUT discovery stops the fold.

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

**Execution shape:** your own pre-edit FOLD_SCOPE artifact first; RED-first per finding (FILE-captured); commit-per-finding or one commit if the seams stay cleanly attributable in the diff→license table; full uncached battery green at the end (and at each commit if multiple); fold report here with FOLD_SCOPE above ACTIONS_GIT_REF, the RED/GREEN captures, and the carry list restated. On its green I re-verify and the merge-decision go-ahead executes.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s10-build-slice-review/REVIEW-FOLD-planner-20260713-034634.md` — run before handoff.
- Every MF locus read by me at `f481042` this session before dispatch: `genesis.go:98-134` (the v8 short-circuit), `loop.go:140-232` (callHandler nested-service + process's commit path), `read_file_worker.go:97-135` (the three select arms), `executor.go:198-263` (the deliberation select + Background ctx), `obligation.go:88-100/:270-290` (the abort + per-record Load), `grep NewResummonScheduler cmd/` → empty. Finder provenance: 8 independent angles (line-scan, removed-behavior, cross-file tracer, concurrency, reuse, simplification, efficiency, altitude/conventions); design-intended candidates refuted against c3/J1/G2/the ruled carry and recorded in the review summary.

ACTIONS_GIT_REF: none — a fold dispatch; no edits by this relay
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); the s10 worktree clean at `f481042`
Next requested action: execute the fold (MF-2…MF-8 + AO-2/AO-3) per the shape above; fold report to this dispatch; I re-verify and then execute the operator's standing go-ahead on the merge decision. Merge itself remains gated on the operator's HUMAN_MERGE_AUTHORIZATION per the standing merge-decision relay.
