## RECONCILE -- REVISE: r10 has the right states, but `content_lost` is on the wrong side of `turn_open` and the overflow branch has no single durable outcome

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev10 is not ready until the manifest producer timeline and overflow disposition are coherent
GRILL_REQUIRED: no -- the semantic classes, one-carrier decision, and fail-closed posture are settled; this return is limited to producer timing, one durable overflow outcome, and discriminating proof
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-161500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- preserve r10's determinate-terminal mapping and full-frame sizing, but separate the pre-admission manifest from post-inspection content loss and close overflow to one durable fail-closed terminal

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-161500.md` at SHA-256 `844359d7354fe34ef675de72af8e2adb3e56f38d9a968db369d133174a6702f5`.

Proposed amendment rev10: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `1efb9e571cfaec69ab5f4eac12b4ac70a7eb3f0b2a3efa7a4516f31bf4a92d22`.

## r10 closures -- ACCEPTED

- The determinate terminal/no-resume class is now truthful: genuine zero-content states are separated from failed/post-invocation-cancel partial bytes that are intentionally discarded, and `NOT_INVOKED_INTEGRITY_FAULT` maps to definite no-effect plus `turn_failed` (`STEP-3-STAGE6-AMENDMENT.md:205-223`, `:252-260`).
- The first-action table no longer lets denied/rejected/failed/cancelled rows leak into clean-positive continuation. Any retry remains an explicit fresh user-requested act.
- Complete-frame pre-sizing names every required member and preserves the one-carrier/no-chunking decision (`:241-248`).
- `xit-dur-3/4/5` are correctly added beneath the existing Durability leg; the gate remains six legs (`:324-336`).

These decisions stand. The remaining defects are in when a class can be known and where the overflow result durably lives.

## Findings

### F105-D2-R10 -- BLOCKER: the immutable pre-admission manifest includes a class that only post-`turn_open` m-9 can discover, while `completed` without a receipt is still unmapped

The owner timeline is fixed by the same amendment:

1. m-10 produces the settlement manifest from its canonical rows (`:184-186`).
2. m-10 persists the immutable manifest bytes in the continuation-admission transaction and sends them on `turn_open` (`:231-240`).
3. Only after receiving `turn_open` can m-9 inspect its private log and report `RESUMABLE` or `DEGRADED` (`:261-269`).

Against that order, `content_lost` cannot be a class in the manifest m-10 commits before step 3. Rev10 defines it as "an exhibited settled-without-content cut" (`:219`), but m-10 cannot inspect the log and no pre-admission owner-real report exhibits that cut. It can emit the positive evidence it owns; only m-9 can later discover that the referenced content is absent/corrupt and derive `content_lost -> DEGRADED`. The current four-class **manifest** union therefore conflates a pre-admission evidence carrier with a post-inspection reconciliation result.

One canonical branch remains non-deterministic for the same reason. A provider row terminal `completed` without a committed content-ready receipt is not `settled_with_content`, is not one of the enumerated `determinate_no_resume` terminals, and is not listed under `uncertain`. The earlier sentence maps missing receipt to "`uncertain` or `content_lost`" (`:197-204`), but m-10 must choose exactly one class before m-9 inspection. As written, the claimed exhaustive mapping still has an `or` where a producer-decidable rule is required.

There is also a stale schema contradiction: `:184-188` still says **every** entry carries `args_digest`, while the owner-real split at `:221-223` correctly forbids requiring it on provider entries.

Required correction:

1. Separate the two closed unions/stages. The immutable m-10-produced **manifest evidence** must map every canonical row plus receipt-presence predicate exactly once. A safe mapping is: completed+receipt or executed-after-marker -> positive settled evidence; determinate terminal/no-resume -> determinate; completed-without-receipt and UNKNOWN/PARTIAL -> uncertain. Exact names may differ, but no producer-time `or` or omission may remain.
2. Keep `content_lost` as an m-9 **reconciliation result** after log inspection: a manifest entry carrying positive settled evidence but missing from the valid prefix yields `content_lost`, `DEGRADED`, and the existing receipt-gated report. It is not an initial manifest member unless a new owner-real pre-admission producer/record is explicitly introduced.
3. Replace the stale generic entry sentence with the already-correct split schemas: tool entries carry `args_digest`; provider entries do not.

This preserves r10's semantic distinctions; it only makes their producer and transition order executable.

### F105-D3-R10 -- BLOCKER: the oversized-frame branch simultaneously admits, does not commit, and lets two different outcomes remain open

For `> FRAME_MAX`, rev10 says a typed operator-visible `DEGRADED` branch occurs, then says "the continuation admits in re-derive/degraded mode, **or** abandons to the operator," while also promising that no continuation/lease/snapshot is committed (`:241-248`). Those claims cannot all hold:

- An admitted continuation needs a durable row/lease and a legal `turn_open`; the required manifest cannot fit and one-carrier forbids chunking or an alternate transfer.
- If no continuation row/snapshot is committed, its `resume_disposition` column cannot hold durable `DEGRADED` and no m-9 exists on the other side of `turn_open` to report it.
- The next paragraph says m-10 cannot truthfully choose `RESUMABLE` or `DEGRADED` at admission because only post-`turn_open` m-9 can inspect the log (`:261-265`), directly contradicting the pre-`turn_open` m-10-selected `DEGRADED` branch.
- Automatic re-derive without the manifest is unsafe: the worker would not receive the determinate/uncertain rows whose disclosure prevents replay.

The branch is therefore neither closed nor durably operator-visible. Because one-carrier is settled and the frame cannot be emitted, the existing fail-closed rule at `:267-269` supplies the safe outcome: **no successor turn, no active-turn lease, no snapshot, and one durable m-10-owned run-terminal/parked record with a closed overflow reason and operator `stop_reason`/`resume_action` projection**. Exact table/message encoding remains pair DESIGN, but master must name one outcome, its durable home, and the no-successor invariant. "Admit or abandon" cannot be delegated as an implementation choice.

### F106-R10 -- BLOCKER: the new fixtures still admit the exact omission and permanent-hold mutants

- `xit-dur-3` says terminal-first/receipt-absent is merely "not settled" (`:332`). A mutant that omits the canonical completed row from the manifest passes. Require **exactly one `uncertain` evidence entry** for completed-without-receipt, and separately exercise positive-evidence-then-missing-prefix -> m-9 `content_lost`/durable `DEGRADED`.
- `xit-dur-4` proves zero work before the disposition receipt but has no positive post-receipt assertion. A worker that remains blocked forever passes. Require the selected first action to become reachable exactly once after the same durable receipt while preserving all pre-receipt zero-work cuts.
- `xit-dur-5` targets a "typed `DEGRADED` branch" whose durable row/reason and successor absence are undefined. Bind it to the single overflow outcome selected above, asserting exact durable rows, no successor/lease/snapshot, and the operator projection.

The six-leg structure remains accepted; only these sub-fixture predicates need to discriminate the corrected transitions.

## Gate disposition

- Proposed stage-6 amendment rev10 `1efb9e57...`: REVISE on the two producer/state contradictions and their proof predicates; not ready for operator re-scope ratification.
- Rev10's truthful determinate-terminal semantics, source-specific final schemas, no-auto-advance first action, complete-frame sizing inputs, one-carrier rule, and added Durability sub-fixture structure are accepted.
- Every r8/r9 closure remains accepted; the decomposition-versus-pair-design grain remains settled with no operator arbitration request.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r11 over new amendment bytes that: (1) make the m-10 pre-admission manifest producer-total and move `content_lost` to the m-9 post-inspection reconciliation result; (2) replace the oversized-frame admit-or-abandon branch with one durable m-10-owned fail-closed outcome and no successor; (3) remove the stale provider-`args_digest` sentence; and (4) make `xit-dur-3/4/5` reject omission, permanent hold, and undefined overflow-state mutants. Preserve all accepted closures and keep exact frame/table/record internals delegated under F73.

## Verification

- Target SHA-256: `844359d7354fe34ef675de72af8e2adb3e56f38d9a968db369d133174a6702f5`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1890.
- Amendment rev10 SHA-256: `1efb9e571cfaec69ab5f4eac12b4ac70a7eb3f0b2a3efa7a4516f31bf4a92d22`; VP r9 parent SHA-256: `0c9235de5298a7fbd9ee38c8db789c3c1ea9ede47c74a28b8c89faef3db706be`.
- Relevant frozen bases remain unchanged: m-8 provider `4b670a79...`; m-9 lifecycle `4d3bd14e...`; m-9 worker `cb7ff970...`; m-10 seam `d2ce9831...`; m-10 control plane `6fd1d655...`; MVP amendment `2f75f2a1...`.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update.
Next requested action: master folds only the manifest-producer/result split, single durable overflow terminal, stale schema sentence, and discriminating sub-fixture predicates, then returns amendment rev11 for decomposition review r11; operator re-scope ratification remains held.
