## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r1
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the findings are bounded corrections required by the existing dispatch; the operator merge gate remains separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-044605.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-045404.md
SUBJECT: must-revise - the monotonic split is right, but volatile pending state, caller-coupled AfterAccepted sequencing, and legacy fail-closed compatibility are not yet contract-real

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed the H-16 design note at exact SHA-256 `eb7a4bd3973991666360b701429c94a13cfb34f0ded98ada4abb8a3552014969` against the lane dispatch and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.

The core decision is correct: a committed decision must never be relabeled by post-commit work, and `AfterAccepted` must not be blindly replayed after its caller-coupled credential delivery is lost. Three gaps prevent the proposed mechanism from satisfying the dispatch's durable crash/retry and fail-closed compatibility bars.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, or deploy.

## Findings

### R1-F1 - the first retryable fault is memory-only, so restart silently converts `pending` to `complete`

The dispatch requires every post-commit failure to become a **durable** derived-work fault and explicitly requires the cut `failure -> restart -> retry` to return the committed decision plus the pending fault (`DESIGN-orchestrator-planner-20260720-043321.md:22-24`).

The note instead stashes the first retryable failure only in one in-memory slot. A durable park record is written only when a second distinct failure displaces the first (`2026-07-20-h16-outcome-split.md:37-45`). Replay derives `pending` only from that live slot and derives `failed` from a parked record; absent either, it returns `complete` (`:47-50`).

After a crash between the first `pending` reply and any second failure, the new Loop has neither slot nor parked record. Rebuilt tables therefore classify the committed record as `complete`, lose the remaining-hook cursor, and cannot retry the work. T2 is impossible under the specified mechanism.

Required revision:

1. Durably represent **every** post-commit failure before returning `pending` or `failed`, including the source relay ID, exact remaining-hook cursor, retryable/parked status, safe reason, and a stable unique identity.
2. Define the exact append-only record/obligation and resolution shape by which heal or park terminalizes that durable fault, plus the table derivation that reconstructs it before admission/replay after restart.
3. Treat any single in-memory slot as a scheduling cache only, never the source of truth. If the implementation stays single-active-slot, startup must hydrate it from the durable unresolved fault and a second failure must not erase either fault.
4. Rewrite T2/T5 to prove restart reconstruction and two-fault durability, not merely same-process slot behavior.

### R1-F2 - retrying an earlier hook cannot safely complete the still-unrun `AfterAccepted` delivery

The live order is `completeTurn -> AfterGateResolution -> AfterApprovalResolution -> AfterAccepted` (`internal/engine/loop.go:168-187`). If any of the first three fails, `AfterAccepted` has not run. The note retries the earlier remaining hooks at the top of **any** later `process`, then clears the slot on success, while declaring `AfterAccepted` non-retriable because its mint extras are meaningful only in the matching direct reply (`2026-07-20-h16-outcome-split.md:39-44`).

That leaves no correct branch:

- clearing the slot after the earlier hooks heal skips `AfterAccepted` but reports `complete`;
- continuing into `AfterAccepted` at the top of an unrelated command can mint a credential whose extras go to no caller, the exact action the note forbids; or
- waiting for same-intake replay contradicts the stated top-of-every-process retry rule and still needs a durable cursor distinguishing not-yet-run from ran-but-reply-lost.

This is reachable for an accepted `seat_mint`: `AfterAccepted` is the sole live reply path through `completeSeatMintBinding(..., includeReply=true)` and its credential/endpoint are intentionally unpersisted (`cmd/frank/main.go:301-303,603-621`).

Required revision:

1. Specify the exact post-commit hook state machine, including not-started/running-or-unknown/completed for the caller-coupled delivery hook and the ordering relationship to retryable hooks.
2. Prove that no unrelated command can execute a mint-producing `AfterAccepted`.
3. For an accepted mint whose earlier hook failed before delivery, choose an honest disposition: either a same-intake, caller-present execution path with a durable cursor and exact response rule, or immediate durable `failed`/operator re-mint. Do not report `complete` without delivery and do not mint into a discarded result.
4. Add the earlier-hook-fails-before-seat-mint-delivery cut and the unrelated-command negative to the battery.

### R1-F3 - an additive ignored field does not make old state-only callers fail closed

The dispatch explicitly requires callers that understand only the old shape to fail closed rather than misread the split (`DESIGN-orchestrator-planner-20260720-043321.md:23`). Under the proposed additive shape, an old decoder ignores `post_commit_state` and reads `state: accepted` for both `pending` and `failed` (`2026-07-20-h16-outcome-split.md:17-35,52-56`). That is not fail-closed; the old caller cannot distinguish complete acceptance from incomplete derived work.

T7 proves only that old JSON decodes. It does not prove the required disposition. The current host also has a state-only accepted consumer in the delivery-nudge callback (`cmd/frank/main.go:337-341`), confirming that consumer behavior must be swept rather than inferred from transport pass-through.

Required revision:

1. Choose an outcome/version or negotiated surface that makes a legacy consumer mechanically unable to treat a non-complete post-commit result as complete acceptance.
2. Enumerate every current native/MCP/host consumer and pin its pre-commit, complete, pending, and failed behavior. Update the nudge rule as part of this seam if it must wait for `complete`.
3. Replace T7 with a true old-consumer fail-closed test in addition to additive/new-consumer decode tests. If additive encoding remains, provide the concrete compatibility mechanism that closes ignored-field behavior; absence of an in-tree authority action is not a wire-compatibility proof.

## Accepted direction

- `State`/decision truth must remain byte-exact `{accepted, rejected, held}` and immutable after commit.
- A separate post-commit dimension with `{complete, pending, failed}` is the correct semantic model.
- `supersededCredentialOutcome` must preserve `credential-superseded` detail while reporting derived-work state separately.
- `AfterAccepted` must never be automatically rerun after a possibly realized mint whose extras no matching caller can receive.
- The correction remains bounded to Outcome, the five failure sites, replay coherence, exact consumers needed for compatibility, and tests; INV-CATALOG and Step-2 closure remain untouched.
- No operator product decision is required. These findings enforce the already-routed durability, caller-coupling, and fail-closed requirements.

## Revision bar

Return fresh design bytes and a fresh hash that:

1. Make pending and failed derived-work states durable and restart-reconstructible before reply.
2. Define a total hook cursor/state machine that never loses or callerlessly executes `AfterAccepted`.
3. Make old-shape callers mechanically fail closed on non-complete work, with an exact consumer census.
4. Add red-first tests for the corrected crash, two-fault, pre-mint earlier-failure, unrelated-command, and legacy-consumer cuts.
5. Preserve the accepted monotonic split, bounded scope, branch-only sequence, and operator merge gate.

The VP parallel-narrowing lane had no additional relay in `h16-outcome-split` at this review sweep. Check it again before IMPL as the dispatch requires.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-044605.md` lints OK and directly addresses `m-7.implementer`.
- Design note hash recomputed: `eb7a4bd3973991666360b701429c94a13cfb34f0ded98ada4abb8a3552014969`.
- Re-read the dispatch, full design note, `internal/engine/loop.go:24-54,130-199,259-282,299-374`, and `cmd/frank/main.go:190-217,274-324,337-341,603-651`.
- Lane sweep found only the orchestrator dispatch and this incoming planner DESIGN before this review; no VP narrowing was present.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0.
Next requested action: m-7.planner folds R1-F1..F3 into fresh design bytes, checks the lane again for VP narrowing, recomputes the hash, and requests a fresh uniquely-parented review before any IMPL branch work.
