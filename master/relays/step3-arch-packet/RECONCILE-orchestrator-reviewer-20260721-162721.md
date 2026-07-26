## RECONCILE -- REVISE: rev11 closes the producer split, but its resume-time invariant contradicts `content_lost` and overflow still has two lifecycle meanings

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r11
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev11 needs one final consistency fold before its exact bytes are ratifiable
GRILL_REQUIRED: no -- the manifest stages, one-carrier decision, and fail-closed direction are settled; the return is limited to temporal truth, exact overflow lifecycle state, and exhaustive fixture predicates
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-162500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- producer-total manifest and post-inspection content_lost are accepted; make the trust invariant time-correct, choose run-terminal versus parked exactly once, and close the last omission/permanent-hold fixture branches

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-162500.md` at SHA-256 `2a8036cd18ce69447fdffde55924ebeb3baf0a0f5590732af9adf724f45ac378`.

Proposed amendment rev11: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `61fe014c0fe66c3096a750d9da3ca08c3ae6030f3c4a891b62749a0ee20da0dd`.

## r11 closures -- ACCEPTED

- The immutable pre-admission manifest is now a producer-total m-10 evidence union: every canonical tool/provider terminal or park maps once to `settled_with_content`, `determinate_no_resume`, or `uncertain`; provider completed-without-receipt maps deterministically to `uncertain` (`STEP-3-STAGE6-AMENDMENT.md:193-231`).
- `content_lost` is correctly moved to the m-9 post-`turn_open` reconciliation result, with `DEGRADED` reported through the receipt-gated disposition path (`:233-235`).
- Tool/provider entry schemas are source-real; the stale universal provider `args_digest` requirement is gone (`:193-198`, `:237-239`).
- Oversized resume frames now have no successor, lease, snapshot, alternate carrier, or automatic re-derive (`:257-267`).
- `xit-dur-3/4/5` now contain the requested no-omission, post-receipt progress, and overflow-state structure under the same six-leg gate (`:344-356`).

Those r10 findings close. The residuals below are exact-byte contradictions in the newly folded wording and proof predicates, not a request to reopen the architecture.

## Findings

### F105-D2-R11 -- BLOCKER: the stated construction invariant makes the accepted `content_lost` result unreachable

Rev11 correctly calls `settled_with_content` **positive evidence** that content should be in the prefix and says m-9 verifies it at stage 2 (`:215-235`). But `:241-245` still declares `settled_with_content => content in the durable valid prefix, by construction`. Read at continuation/resume time, that implication forbids the exact state `:233-235` and `xit-dur-3` require: a `settled_with_content` entry whose referenced content is now missing/corrupt, producing `content_lost`.

The two true invariants are temporal and conjunctive:

1. **At settlement/content-ready commit:** positive evidence proves the content and admitting marker had durably linearized at that earlier point.
2. **At resume trust:** content is trusted only when the manifest carries matching positive evidence **AND** m-9 finds the matching content in the current recovered valid prefix; evidence-without-current-content yields `content_lost`, never trusted content.

Required correction: replace the timeless implication with those two time-scoped properties. The class name may remain; the bytes must not simultaneously guarantee current presence and specify its absence branch.

### F105-D3-R11 -- BLOCKER: `run-terminal/parked` still leaves two lifecycle outcomes under the claimed single overflow terminal

Rev11 removes the unsafe continuation, but `:259-265` says m-10 commits one **"run-terminal/parked record."** Terminal and parked are behaviorally different: terminal controls run GC/restart and forbids same-run continuation; parked preserves a nonterminal run awaiting an operator transition. The slash leaves the product state and allowed future transitions to pair implementation even though the heading claims one closed outcome.

The amendment already supplies the safe decision at `:286-288`: inability to establish the degraded path is **fail-closed run-terminal**. Apply it here explicitly. Pin one lifecycle result, for example the existing m-10 run terminal `FAILED` carrying closed reason `resume_frame_overflow`, with no same-run successor and an operator-visible manual next action. The exact table column/message encoding remains pair DESIGN; terminal-versus-parked and whether same-run revival exists do not.

Required correction: remove `run-terminal/parked`; name the one lifecycle state, its terminality, the no-same-run-successor rule, and the operator projection. Keep no continuation/lease/snapshot.

### F106-R11 -- BLOCKER: one omission mutant and one permanent-hold mutant still pass the written predicates

- `xit-dur-3` now requires exactly one `uncertain` entry for terminal-first/receipt-absent, but receipt-first/terminal-absent remains only "not settled" (`:352`). After retirement the extant provider row is canonically UNKNOWN/PARTIAL and must also appear **exactly once as `uncertain`**; omission still passes the current branch.
- `xit-dur-4` says the selected first action becomes "reachable exactly once" after receipt. Reachability is not an observed machine outcome; an implementation can expose a runnable transition and never execute it. Bind the fixture to one selected branch and require its corresponding durable/wire action to be **observed exactly once after receipt**, with zero such observations before receipt.
- `xit-dur-5` should assert the exact terminal state selected above, not the unresolved `run-terminal/parked` family.

The sub-fixture structure and six-leg count remain accepted. These are bounded predicate corrections so the fixtures prove the architecture they name.

## Gate disposition

- Proposed stage-6 amendment rev11 `61fe014c...`: REVISE on the temporal invariant, overflow lifecycle token, and three exact predicate clauses; not ready for operator re-scope ratification.
- The three-class producer union, completed-without-receipt mapping, post-inspection `content_lost`, source-specific schemas, no-successor overflow direction, and all earlier r8-r10 closures are accepted.
- The decomposition-versus-pair-design grain remains settled; no operator arbitration is requested.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r12 over new amendment bytes that: (1) replace `settled_with_content => current content` with settlement-time evidence plus resume-time evidence-and-presence trust; (2) choose one terminal overflow lifecycle state with no same-run successor, removing `run-terminal/parked`; and (3) make `xit-dur-3/4/5` require exact uncertain carriage on both missing-half orders, an observed exactly-once post-receipt action, and the exact overflow terminal. Preserve every accepted rev11 mechanism and keep record/frame internals delegated under F73.

## Verification

- Target SHA-256: `2a8036cd18ce69447fdffde55924ebeb3baf0a0f5590732af9adf724f45ac378`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1892.
- Amendment rev11 SHA-256: `61fe014c0fe66c3096a750d9da3ca08c3ae6030f3c4a891b62749a0ee20da0dd`; VP r10 parent SHA-256: `def00664988ff03a9ef77ab160036e2a2f9fd9555cba86e37785e000f71e6317`.
- Relevant frozen bases remain unchanged: m-8 provider `4b670a79...`; m-9 lifecycle `4d3bd14e...`; m-9 worker `cb7ff970...`; m-10 seam `d2ce9831...`; m-10 control plane `6fd1d655...`; MVP amendment `2f75f2a1...`.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update.
Next requested action: master folds only the time-scoped trust invariant, exact overflow terminal state, and discriminating `xit-dur-3/4/5` predicates, then returns amendment rev12 for decomposition review r12; operator re-scope ratification remains held.
