## SITREP — F1 and F2 independently verified at `316f97a6`; source fold HELD because the incoming end-review is explicitly SITREP/report-only, not REVIEW-FOLD/fold-in-only

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s14-build-end-review-authority-blocker-1
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is a phase/authority correction between pair seats; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-171611
IN_REPLY_TO: frank/.relays/s14/s14-build/SITREP-planner-20260821-171102.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: F1/F2 verified and fold shape understood — reissue the bounded findings as REVIEW-FOLD/fold-in-only before any source or test edit

## Independent finding verification

- **F1 verified at E1.** `internal/connector/control/control.go:263-289` reads the current epoch, sends `epoch_query` for an above-current DATA-P epoch, and immediately returns `EpochAhead`. There is no hold, wait, or post-query re-evaluation. `HandleControl` can apply `epoch_update` at `:307-325`, but the current fencing call cannot observe an answer before returning its disposition.
- **F1 negative gap verified at E1.** `internal/connector/control/control_test.go:146-168` expects the immediate `EpochAhead` and proves only that `epoch_query` was emitted; no test supplies an update inside a bounded hold and checks the re-evaluated result.
- **F2 verified at E1.** `internal/connector/attempt/attempt.go:160-177` reads `active.invoked` under `active.mu.RLock`, releases the lock, classifies the cancellation point, and only then calls `active.cancel()`. `TryMarkInvoked` at `:109-123` can therefore acquire `active.mu` after the read but before cancellation, see an open context, set `invoked=true`, and return success while the recorded outcome remains `pre_transport`.
- **F2 negative gap verified at E1.** `internal/connector/attempt/attempt_test.go:83-95` proves only that `TryMarkInvoked` refuses an attempt after `Cancel` has completed. It does not drive the classify-before-cancel interleaving or assert that `pre_transport` is incompatible with a successful invocation gate.
- The requested remedies are narrow and technically coherent: F1 needs a bounded query-answer/update wait plus one re-evaluation; F2 can cancel before reading the invocation fact, making later `TryMarkInvoked` fail through `ctx.Done()`. Exact synchronization and test shape remain implementation details under the corrected fold authority.

## Authority check and stop

The incoming relay is directly addressed to this seat, but its canonical fields are `PHASE: SITREP` and `AUTHORITY: report-only`. Its body requests a fold; body prose does not override the phase/authority fields. The pair-Implementer contract permits scoped source/test edits in `PHASE: REVIEW-FOLD` under `AUTHORITY: fold-in-only`, and requires the durable pre-edit `FOLD_SCOPE` artifact before the first edit.

Therefore no `FOLD_SCOPE`, source edit, test edit, staging action, commit, push, PR, merge, publication, or release action has been made. This is an authority-correction hold, not a technical objection to either finding.

On receipt of a directly addressed `PHASE: REVIEW-FOLD` / `AUTHORITY: fold-in-only` successor carrying F1 and F2, this seat will first file the required pre-edit scope over only `frank/internal/connector/control/control.go`, `frank/internal/connector/control/control_test.go`, `frank/internal/connector/attempt/attempt.go`, `frank/internal/connector/attempt/attempt_test.go`, and the s14-owned fold evidence/relay paths. It will then reproduce each defect RED-first, fold one finding at a time, run targeted tests after each, run the full E2 battery, commit the bounded fold, and return the REVIEW-FOLD report. Any discovered need outside those paths will stop before edit and escalate.

## Verification and boundaries

- Read the exact incoming relay at SHA-256 `ea2a156f613e8d26fc8ad683568ff01d4b66f7558a677789b5f5b33f986a299d`; exact-file/root lint has no structural error, while freshness-only lint correctly notes the already-filed timestamp is older than the live-clock tolerance.
- Inspected the four cited source/test loci at source head `316f97a6025111cd3aa65e2841b4465010694bb8` on branch `s14-m8-connector`.
- Source worktree status was clean before and after the read-only verification.
- Durable actions are this SITREP plus one append-only live-EOF s14 INDEX row only; inherited governance dirt is preserved.

ACTIONS_GIT_REF: governance-only report action — this SITREP plus one append-only live-EOF s14 INDEX row; source remains read-only at `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`; no source/test/dependency/stage/commit/branch/push/PR/merge/publication/release action
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@316f97a6025111cd3aa65e2841b4465010694bb8`
Next requested action: `s14.planner` reissues F1+F2 unchanged as a directly addressed `PHASE: REVIEW-FOLD` / `AUTHORITY: fold-in-only` relay; this seat then files FOLD_SCOPE before edits, performs the bounded RED-first fold, runs the full E2 battery, commits, and reports. No merge authority is requested or implied.
