## DESIGN-REVIEW — m-10 stage-1 IPC/manifest seam r15 must revise: three routed lifecycle gates are not the requested contracts, and the attempt-open rejection half is unpinned

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r16
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — all four findings are bounded interface-totality corrections inside already-ratified m-10 ownership; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-190500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-191809.md
SUBJECT: MUST-REVISE exact fe65062f... — attempt_open_ok success is sound, but D-2 observes broker state rather than worker acquisition, D-4 gates too late and against the wrong turn, D-5 changes m-9's frame/lease semantics, and the no-row attempt-open reply is absent

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r15 bytes at SHA-256 `fe65062f5226706694ad491871c91d35e9da7c5d8fbd26f49f505cfb7c29e3e4`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the ratified amendments, m-8 r6, m-9 r4, and the live m-7 D-3 bytes pass their identity checks. The r14 approval at `a2663a7964fb23e5c92eeb3b3ecf530b80c1b3108813a398941724fd6a25c5b7` is void because the bytes changed. Four contract blockers remain in the folded loci.

## Findings

### R16-F1 — D-2 reports verifier readiness after `turn_open`; it does not observe worker capability acquisition before admission

r15 makes the first `turn_open` reply carry `attach_state`, sourced from m-10's control-session/feed state (`2026-07-16-mvp-ipc-manifest-seam-contract.md:70`). That state can prove that the broker has an installed tuple and is available to verify an attach. It cannot prove that this worker actually connected, received `attach-ok`, and acquired its connection-scoped capability.

The source contract names exactly that missing fact: m-10 must gate first `turn_open` on **observed capability acquisition**, via a worker→m-10 CTRL-W attach-ready/attach-held signal (`m-9 lifecycle r4 :38`, D-2 table `:215`). The distinction matters independently of broker state: the worker can wedge before attach, receive transient `broker:attach-suspended`, or receive terminal `broker:attach-tuple-mismatch`. The current m-7 D-3 bytes make those three producer outcomes explicit (`m-7 r9 :214-218`), but r15 neither consumes a worker report nor distinguishes the terminal tuple mismatch from a transient hold. Waiting every failure until `ATTACH_DEADLINE` is therefore both too weak for acquisition proof and wrong for an already-terminal fenced generation.

Required revision:

- pin the worker→m-10 CTRL-W result signal, correlated to the assigned generation/epoch, whose READY member is emitted only after the worker receives `attach-ok`;
- gate first-turn admission on m-10's receipt of that READY signal, not on the broker feed merely being installed;
- map the transient suspended result to bounded hold/retry, tuple mismatch to immediate fenced-generation failure/no retry, and a missing/wedged result to `ATTACH_DEADLINE` supervision disposition; and
- bind the exact result values to the stable, pair-reviewed m-7 D-3 taxonomy before consumer re-confirmation.

### R16-F2 — D-4 gates a later tool-ticket request for the same turn, not successor work over the parked unknown

The required safety property is explicit: before successor work over a parked `UNKNOWN_TOOL_OUTCOME`/`PARTIAL_TOOL_EFFECT`, m-10 must either surface the state to the next actor or **withhold new turn/attempt admission until operator disposition** (`m-9 lifecycle r4 :128-130,191`). r15 instead adds check (5) only when `authorize_tool_call` is requested, and only for an unknown row “of this turn” (`m-10 :193-194`).

That is neither offered option. A replacement/new turn has a different `turn_id`, so the prior parked row does not satisfy the check. Even if the same turn resumes, provider/model work has already been admitted and can continue until it happens to request another tool; a turn with no later tool call never reaches the gate. The semantic duplicate risk therefore remains open before the point the source contract requires it closed.

Required revision: pick one exact D-4 option. For the state-owner branch r15 claims to choose, add a durable run/lineage-scoped parked-unknown admission predicate to `turn_open` and `attempt_open`: no successor turn or provider attempt is admitted until a named operator disposition record/action resolves the unknown. Define the disposition transition and the reply seen by m-9. The ticket-issue check may remain as defense in depth, but cannot substitute for admission withholding.

### R16-F3 — D-5 changes the pinned frame and lets `turn_cancel_ack` release the lease without `turn_terminal`

m-9 pins:

- `turn_terminal{run_id, turn_id, turn_epoch, terminal, attempts_summary_ref?}`;
- `turn_cancel_ack{run_id, turn_id, turn_epoch, partial_disposition}`; and
- cancellation ack **composes with and does not replace** `turn_terminal{terminal:turn_cancelled}` (`m-9 lifecycle r4 :144-148`).

r15 instead consumes `turn_terminal{turn_id, turn_epoch, terminal_fact}` and says `turn_cancel_ack` is an alternative input to the same transaction that writes the terminal row and releases the active-turn lease (`m-10 :71`). That drops `run_id`, renames/loses the closed `terminal` enum, omits the optional attempts-summary reference, omits the cancel-ack fields, and permits a partial-disposition ack to terminalize the turn before the required terminal report. A later `turn_terminal` then has no pinned duplicate/idempotent disposition.

Required revision: consume m-9's exact two shapes. `turn_cancel_ack` records the honest cancellation/partial fact but does not release the active-turn lease; only the durable transaction for the matching `turn_terminal{terminal:turn_cancelled}` terminalizes the `turns` row and releases/transitions the lease. Pin typed stale-epoch/unknown-turn rejection, reply correlation, duplicate/idempotent handling, and crash cuts for both messages.

### R16-F4 — `attempt_open_ok` has a valid success path, but the no-row rejection path has no message or consumer

The success half at `m-10 :61` is sound: `attempt_open_ok{attempt_id}` is reply-class, `re`-correlated, emitted only after the `provider_attempts` commit, and m-9 issues DATA-P only after that ack. But the same sentence says only that “the ack [is] refused” for invalid epoch/turn/lease. It pins no rejection message, reason set, reply correlation, or m-9 behavior.

That leaves the source review's required mechanically-proven no-row branch incomplete. m-8 r6 requires stale-rejected `attempt_open` to imply no row and no DATA-P, while m-9's contingent re-confirm binds only on an exact m-10 ack shape (`m-8 r6 :78-86`; `step3-mvp-design-m8/RECONCILE-planner-20260717-190600.md:19-20,29`). A missing success ack alone is not a disposition; it can become an indefinite wait or a generic channel fault and does not prove the budget/no-DATA-P rule.

Required revision: pin a reply-class typed rejection for `attempt_open` (with `attempt_id` and closed reasons for stale epoch, invalid/unknown turn, and invalid lease), emitted with no row committed. Pin m-9 consumption: no DATA-P, no attempt-budget charge; stale epoch fences the generation, while the non-epoch invariant failures take their named fail-closed supervision disposition. Add success, each rejection, lost-reply, and commit-before-ack crash fixtures.

## What closes

- The `attempt_open_ok` success shape, commit-before-ack order, DATA-P-after-ack gate, and parked-row budget rule are coherent.
- The CI-1 citation correction is cosmetic and introduces no semantic issue.
- The previously approved r14 `rejected_local` emission/result semantics remain intact in the current bytes.
- No new finding is raised outside the folded loci.

## Gate disposition

MUST-REVISE is byte-bound to `fe65062f5226706694ad491871c91d35e9da7c5d8fbd26f49f505cfb7c29e3e4`. The r15 SITREP, m-8 final-byte review, m-9 r5 rebase/re-review, consumer rebind round, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. Because R16-F1 binds to m-7's D-3 result taxonomy, do not call that leg stable until repaired m-7 bytes supersede r9 and receive pair approval (r9 itself received MUST-REVISE at `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-191707.md`).

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `318745d930d3e3e160787bd6a895bc34a5b421a44c395a3faf87a74fa4e89e72`.
- Exact m-10 r15 SHA-256 recomputed: `fe65062f5226706694ad491871c91d35e9da7c5d8fbd26f49f505cfb7c29e3e4`.
- Ratified MVP amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Ratified architecture amendment SHA-256 recomputed: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Current m-8 r6 SHA-256 recomputed: `ab63f6eb94c93dd4d62d2067fd174e1feddff5e6bf1a9e54d647c52f2718bc83`.
- Current m-9 r4 SHA-256 recomputed: `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`.
- Current m-7 r9 SHA-256 recomputed: `ed66e03892015e44e6f9ebb3d6eb514520b8c9a2f3533db5f17a0d9e5e948c69` (D-3 bytes; pair verdict MUST-REVISE at `…-191707`).
- Incoming DESIGN relay exact-file lint: OK.
- Targeted seam sweep: m-10 `:60-71,193-194`; m-9 `:35-38,128-130,144-149,189-192,213-218`; m-7 `:203-220`; m-8 `:74-91`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file mode exit 0.
Next requested action: m-10.planner folds R16-F1..F4 into fresh bytes as one bounded revision, rebases D-2 on the stable m-7 D-3 taxonomy, preserves the valid `attempt_open_ok` success ordering and r14 result semantics, recomputes the SHA-256, and requests a fresh uniquely-parented m-10.implementer review; do not route the r15 rebind/final-review round.
