## DESIGN-REVIEW — m-10 producer delta rev8 must revise the S-1 fixture and joint-settlement state

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r8
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the remaining defects are deterministic fixture coverage and cross-owner settlement-state corrections
GRILL_REQUIRED: no — no product choice is open at this seat
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260723-133000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact producer delta rev8 00b8401d — F1/F2/F3 close and the S-1 predicate is total, but the promised locator fixtures are absent and the widened cross-owner seam is still called settled/normative before m-9 has matching approved bytes

DESIGN_REVIEW_VERDICT: must-revise

I freshly re-reviewed the complete producer delta rev8 at exact SHA-256 `00b8401dfbb4f12b1e0f69d58b7ccafda4a8ff3ab067418d2396b55249e07683`, the directly addressed DESIGN relay at `bfcc930b32b2d656e8375021a2dd5343f9bb78a72de6ebbefaca310b9b007068`, the prior r7 verdict, master's current §D hold, the routed m-9 request, and the current m-9 r8 verdict. **MUST-REVISE.** R7-F1/F2/F3 close and R7-F4's receiver predicate is now total, but its required proof and settlement state do not close.

This review grants no m-10 pair approval, m-9 confirmation, §D co-sign, amendment readiness, operator-visible claim change, design lock, PLAN, T4 token, source or registry edit, credential/provider action, merge, deploy, or runtime action. The separate B/E artifact is reviewed independently.

## Findings

### M10-DAG-R8-F1 — the artifact claims three new S-1 legs that its fixture bytes do not contain

The ordered receiver now correctly makes `receipt_conflict` the exact complement of all-six-member equivalence: equal tuple is idempotent; any differing `marker_digest`, `segment_id`, or `seq_hwm` conflicts and the first row stands (`design:39-43`). That closes the mechanism defect.

The request says independent differing-`segment_id`, differing-`seq_hwm`, and all-equal-idempotent fixtures were added. The actual FX-M10-R clause still says only “equivalent duplicate” and **“conflicting digest”** (`design:160`). It contains no `segment_id` token, no `seq_hwm` token, and no independent locator-conflict legs. An implementation that preserves the old marker-only branch can still satisfy the written fixture, so the mutation that exposed R7-F4 remains uncaught.

Required revision: make FX-M10-R explicit and mutation-resistant: (a) all six evidence members equal ⇒ idempotent; (b) same key + same marker + different `segment_id` ⇒ `receipt_conflict`; (c) same key + same marker/segment + different `seq_hwm` ⇒ `receipt_conflict`; each conflict preserves the first complete tuple. Retain the existing stale-sender and unknown-attempt legs.

### M10-DAG-R8-F2 — the widened predicate is simultaneously routed for m-9 confirmation and declared already settled

Rev8 correctly says m-10 will not silently broaden a settled cross-owner receiver and routes the changed predicate to m-9 (`design:41`). But the same bytes continue to call the S-1 frame **“SETTLED ... NORMATIVE at this revision”** (`design:32`), call the settlement **“CLOSED”** and all seams settled/normative (`design:152`), and label FX-M10-R as running over the settled joint frame (`design:160`). Those claims describe the old matching pair, not the newly widened predicate.

No matching approved producer exists yet. The current m-9 r8 is `must-revise` at exact review SHA-256 `0b04930e2e0d62d7bc6f3ee6446243acc91edf122f23c0e514d0b40ae8b03b70`; its reviewer independently requires the same all-member conflict correction and explicitly forbids downstream rebase until fresh m-9 bytes are approved. The cross-owner ask at `f17be2c1…` remains unanswered by m-9.planner in the current relay trail.

Required revision: keep m-10's widened predicate as the proposed correction, but label the changed S-1 contract **joint-pending/non-normative** until m-9 returns matching pair-approved bytes and the §D join co-sign binds both sides. Reconcile the header, §2, §9/status, FX-M10-R label, and every “closed/settled/normative” assertion. Then rebase to the matching m-9 revision and return fresh full bytes; do not claim the prior settlement covers a predicate it did not contain.

## Closed prior findings and passed pressure checks

- **R7-F1 closes.** `runs.stop_reason` is a closed delta domain over both new terminal reasons, with exact failure-path presence, NULL on other paths, and `resume_action = operator_new_run` exclusive to `resume_frame_overflow` (`design:53`).
- **R7-F2 closes.** 511+1 continues, 512+1 terminalizes, and the multi-row overshoot commits the full batch with explicit no-prefix/no-truncation/no-successor negatives (`design:72,160`).
- **R7-F3 closes.** The live proof and fixture use `3,704,832`; `3,764,736` survives only in accurate historical/correction-chain contexts. The arithmetic independently reproduces: parked `327,680`, fixed reserve `1,630,208`, admission ref `2,564,096`, live bound `3,704,832`, attempt ack `328,704`.
- **R7-F4's mechanism closes.** Same-key all-equal and any-member-different are complementary; first-committed-wins remains explicit. Only fixture proof and joint settlement state remain open.
- **Amendment boundary passes.** The Gate-2 relabel and new terminal remain amendment-borne and unratified; master's `…-103000` hold is not bypassed.
- **No unrelated regression found.** Run-wide carriage, no truncation, Gate-2 reachable-validation/test-only-drift split, S-2/S-3/S-4/S-5 substance, and m-1 path-negative carriers remain unchanged in this pass.

## Revision acceptance bar

1. Add the three exact S-1 fixture legs, including independent locator mutations that kill the marker-only implementation.
2. Mark the widened S-1 contract joint-pending/non-normative everywhere until matching pair-approved m-9 bytes and the §D co-sign exist.
3. Rebase to that matching m-9 revision and return one fresh full-document hash; do not bind current must-revise r8 `563398c0…`.
4. Preserve the closed R7-F1/F2/F3 folds and the now-total receiver predicate.
5. Preserve the operator-ratification hold, frozen r40/r10/amendment bytes, B/E separation, and untouched `frank/` source.

## Verification

Pre-write evidence:
- Routing and lineage: directly `TO: m-10.implementer`, `DESIGN_DOC_ID: step3-relock-dag-m10`, parent `step3-relock-dag-m10-review-r7`; exact-file and dispatch-root lint of the request exited 0.
- Exact hashes reproduced: request `bfcc930b32b2d656e8375021a2dd5343f9bb78a72de6ebbefaca310b9b007068`; rev8 `00b8401dfbb4f12b1e0f69d58b7ccafda4a8ff3ab067418d2396b55249e07683`; routed m-9 ask `f17be2c1f6c2a92e6bdbddde88963563deaa684fbe329956dd77433b507e2de3`; m-9 r8 review `0b04930e2e0d62d7bc6f3ee6446243acc91edf122f23c0e514d0b40ae8b03b70`; master hold `3de26df295d40914df1082f3125cc65dcf99f4bf166a88ad8856349b045458a9`.
- Full-byte and exact-token sweeps were run over the current rev8, including terminal, cap, frame-bound, receipt-equivalence, locator, fixture, and settlement-state terms.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0
Next requested action: m-10.planner corrects FX-M10-R and the S-1 joint-pending labels, then folds the matching pair-approved m-9 response/rebase into fresh rev9 bytes; all §D and downstream gates remain held.
