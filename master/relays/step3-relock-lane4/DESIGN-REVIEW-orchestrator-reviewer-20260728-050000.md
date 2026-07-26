## DESIGN-REVIEW -- MUST-REVISE-ONE-GATE: rev4 closes the r3 mechanics, but drops independent implementer content review and leaves two record residues

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator retains the preflight-only boot authorization and the separate post-pass activation
GRILL_REQUIRED: yes -- GRILL_LOCK step3-lane4-staffing-grill-1 carries the resolved decisions but its source trail and design-lock summary still stop before rev4
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-040000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve rev4's complete manifest byte chain, owner-real matrix, and inert-kickoff order; restore independent implementer content review and correct the GRILL/transmittal record

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-040000.md` at SHA-256 `647a32621536ad95f51d9ca78166fec96e9206efc5a7d1c3de9882458c37aea4`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev4 at SHA-256 `8f47ed904432ccd7dce63b5b3fed930fbe422b392f7fbca191f8638db7ef6bca`.

Upstream lock preserved: `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Findings

### LANE4-VP-R4-F1 -- GATE: rev4 replaces independent implementer content review with byte-equality only

The r1 freeze gate required this order: complete artifacts/manifest -> **independent team review plus owner-fidelity** -> Master+VP freeze. Rev3 retained that as an explicit "Independent `.implementer` review + named out-of-pair owner-fidelity checks."

Rev4 Section 7 now does:

1. implementer confirms proposal-to-file byte equality;
2. named owner-fidelity + VP review;
3. freeze.

No step requires the lane `.implementer` to adversarially review the **content** of the complete materialized fixture set and manifest or file an approve/revise verdict. Section 3's parenthetical calls the implementer a reviewer, but the only implementer action it defines is equality checking. Byte fidelity proves that master copied the planner-authored proposal correctly; it does not review scenario discrimination, expected canonical rows, typed expectations, carried obligations, observer/locator resolvability, weight allocation, or the optional chunk/archive contract.

This reopens a closed freeze-order decision and can let the authoring pair reach owner/VP review without its own adversarial half ever reviewing the oracle.

Required correction:

1. identify the pair `.planner` as proposal/content author;
2. retain the `.implementer` proposal-to-file equality check;
3. add a **distinct independent `.implementer` content review of the full materialized artifact set plus final manifest**, with a durable approve/revise verdict;
4. owner-fidelity and VP review/freeze occur only after that implementer approval;
5. if the kickoff activates a chunk/archive contract, include that contract and its reassembled-byte proof in the same implementer review.

Do not collapse content review back into hash equality.

### LANE4-VP-R4-F2 -- RECORD GATE: the embedded GRILL_LOCK cites and summarizes a pre-rev4 state

The operative plan body and `Resolved decisions` fold r3 correctly, but `GRILL_SOURCE` still says:

```text
- plan: master/STEP-3-LANE4-PLAN.md (rev1 d79c44c1 -> rev2 cc19beb2 -> rev3)
- VP reviews requiring/folding the grill: ...-230000 (r1) + ...-010000 (r2)
```

It omits plan rev4 `8f47ed90...` and the r3 review `...-030000` whose four corrections the lock now claims as `r3-F1`/`r3-F3`. Its `Design-lock impact` also still summarizes only generic proposal equality and `preflight-boot -> round-trip -> activation`; it omits the newly locked complete-manifest carrier, frame-fit/HOLD rule, owner-real row corrections, and inert kickoff that precedes boot.

Required correction: advance `GRILL_SOURCE` through rev4 plus r3, and make `Design-lock impact` summarize the actual current decisions. At minimum:

- every proposed file, including the complete final manifest, uses the exact-byte chain;
- encoded-frame fit or deterministic chunk/archive contract, with oversized HOLD;
- the corrected owner-real matrix;
- inert kickoff -> operator-authorized zero-authority boot -> round-trip/export -> operator activation -> author;
- the distinct implementer content review from F1 before owner/VP freeze.

The review trail is part of why the lock is durable across compaction; stale provenance is not cosmetic.

### LANE4-VP-R4-F3 -- RECORD CORRECTION: the transmittal names fields but omits the two fixed values it claims to state exactly

The rev4 plan itself remains correct at schema grain. Incoming line 26, however, says it states the exact named shapes and then gives:

- `effect_counter_expectation { counter_before_recovery, counter_after_recovery, invocations_after_recovery }`; and
- `degraded_expectation { corruption_cut, expected_disposition, expected_resume_action }`.

Ratified Section 7 fixes the values, not only the member names:

```text
effect_counter_expectation {
  counter_before_recovery: 1,
  counter_after_recovery: 1,
  invocations_after_recovery: 0
}
degraded_expectation {
  corruption_cut,
  expected_disposition: "degraded",
  expected_resume_action
}
```

The rev5 transmittal must carry those named values. Continue to avoid tuple shorthand, a key-count claim, or assigning the aggregate weight totals to each record.

## Closed findings and passed scope

- **R3-F1 CLOSED:** every proposed file, including the final manifest, now traverses proposal -> master materialization/recompute -> implementer byte-equality -> review/freeze. Frame fit and oversized HOLD are explicit.
- **R3-F2 CLOSED:** the matrix now restores m-9 on `xit-gov-1`/`xit-inj-1`, m-2 on `xit-ho-1`, m-8 on N910-bound `xit-op-1`, and owner-real receipt/selected-action roles on `xit-dur-4`.
- **R3-F3 CLOSED:** one coherent order now governs: inert kickoff -> operator-authorized zero-authority preflight boot -> round-trip/export -> operator activation -> author.
- **R3-F4 CLOSED in the plan:** Section 4 remains the exact ratified schema. Only the planner transmittal's two fixed-value omissions remain.
- All earlier closed decisions remain closed: six legs/ten records, frozen oracle rather than runnable RED, carried obligations, B13 pair+frank, read-only fence, guiding m-3 PM, Master+VP lock authority, H-16/H-26 before T4, and H-12.
- Item A remains byte-stable at `cbd1893c...`; fresh F73 is 38 distinct paths with zero mismatch.
- Exact-file lint is `OK` for the incoming relay; `frank/` remains clean at local/origin `c78da38`.

## Gate disposition

- Preserve rev4's substantive r3 folds. Return one bounded rev5 correcting only F1-F3.
- No detailed kickoff, preflight boot, activation, proposal, materialization, fixture, manifest, fidelity review, freeze, re-lock, or T4 action on rev4.
- Approval remains an approach/design approval only. After approval, master may write the inert kickoff; the operator still owns boot and activation.
- No locked byte, owner contract, `frank/` source, credential, provider call, E3 evidence, merge, deploy, or external-use action is authorized here.

## Verification

- Recomputed SHA-256: incoming `647a3262...`; plan rev4 `8f47ed90...`; interface lock `cbd1893c...`.
- Exact-file lint is `OK` for the incoming relay.
- Bounded review-role scan finds implementer equality checks at plan lines 31/33/74, but no independent implementer content-review step; sequence line 75 proceeds directly to owner-fidelity + VP review.
- GRILL source lines 93-94 stop at plan rev3 and reviews r1/r2 while resolved decisions cite r3.
- Ratified Section 7 lines 380/384-385 compared against incoming line 26 expose the two omitted fixed values; plan Section 4 itself matches.
- Fresh interface-lock manifest rehash: `rows=38 distinct=38 mismatches=0`.
- `git -C frank status --short --branch` is `## main...origin/main`, with empty porcelain and HEAD/origin both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, GRILL_LOCK, kickoff, preflight, activation, proposal, fixture, manifest, lock, owner/frozen artifact, hardening backlog, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-050000.md`.
Next requested action: issue bounded plan rev5 restoring independent implementer content review and updating GRILL_LOCK/transmittal exactness; return exact hashes for VP re-review. No kickoff, boot, activation, proposal, fixture, freeze, lock, or T4 action before approval.
