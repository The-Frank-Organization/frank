## RECONCILE -- REVISE the stage-1 close: the confirmation census is incomplete and the m-7/m-10 gating edge is not closed in the approved owner bytes

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage1-close-review-r1
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- these are owner-byte, confirmation-routing, and sequencing corrections inside the ratified architecture
GRILL_REQUIRED: no -- the m-7 placement grill is valid and remains closed; this review introduces no new topology choice
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260717-015405.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- stage 1 is CLOSE-PENDING: repair the m-7 snapshot-absent contract branch, complete the 16-edge pre-m-8 census, pin m-8's owner schema, and sequence L1/L5/L7 before the lock

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260717-015405.md` at SHA-256 `45b3db1f80cbd6b872d830bc9d001a6e2853258c1a5a8ccbfed025a46342747f`, including the post-target m-8 consumer review, m-10 L7 response, and m-9 lifecycle return now visible in the append-only tail.

## Findings

### F70 -- BLOCKER: m-7/m-10 is not cross-resolved in the approved owner bytes

The two approved contracts disagree at the bootstrap/adoption cut:

- m-10's exact `79fcf742...` contract explicitly separates control-session establishment from snapshot installation and, for non-terminal or `ABORTED` transitions, supplies **no install-eligible snapshot** while the broker remains suspended (`master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:87-92`).
- m-7's exact `f072bd99...` contract gives only the generic no-installed-state suspended floor at `master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:114-121`, while its actual broker-start and adoption procedures still require receiving/presenting and installing a snapshot at `:199-202`.

The later m-7 confirmation relay's statement that step 5 is conditional on a supplied snapshot is an interpretation, not an owner-byte amendment. A consumer confirmation cannot add an absent branch to the pair-approved producer contract. Therefore target rows 1 and 5, and the closure claim at target `:41`, do not pass exact-byte review.

Required correction: m-7 folds the explicit snapshot-absent bootstrap/adoption branch into its owned contract, takes a fresh uniquely-parented m-7 implementer review, and every byte-bound confirmation touching the changed m-7 contract is rerun on the new hash. The reciprocal m-7/m-10 checks must prove the same transition-ID behavior in both directions.

### F71 -- BLOCKER: the promised full confirmation table has 13 rows but the pre-m-8 census has 16

The dispatched final consumer sets, excluding m-8's stage-2 confirmations until that lane closes, are:

| Producer | Required stage-1 consumers |
|---|---|
| m-1 | m-7, m-9, m-10 |
| m-2 | m-9, m-7, m-10 |
| m-3 | m-9, m-10 |
| m-7 | m-9, m-10, m-1, m-2, m-3 |
| m-10 | m-9, m-7, m-3 |

That is 16 producer-to-consumer edges. The source dispatches establish these sets at:

- m-2 original `step3-mvp-design-m2/...-041620:28` plus supplement `...-043520:23,30`;
- m-3 original `step3-mvp-design-m3/...-041700:29`, preserved by supplement `...-043510:25`;
- m-7 original `step3-mvp-design-m7/...-041630:31` plus supplement `...-043459:29,35`;
- m-10 original `step3-mvp-design-m10/...-041640:30` plus supplement `...-043530:23,30`.

The target table omits exactly:

1. m-2 -> m-7;
2. m-3 -> m-9;
3. m-3 -> m-10.

The live routing proves the omission: `step3-mvp-confirm-m7/...-010010` routes only m-1 and m-10; `step3-mvp-confirm-m9/...-010020` routes m-2/m-7/m-10/m-1 but not m-3; and `step3-mvp-confirm-m10/...-010000` routes m-7/m-2/m-1 but not m-3. Route all three missing, byte-bound. m-8's m-1/m-3/m-10 confirmations remain stage-2 close items and are not added to this 16-edge pre-m-8 table.

### F72 -- BLOCKER on the m-8 final-fold path: `tool_result.content` has no owner-defined type

m-8 declares a closed `input_item` enum but leaves `tool_result{tool_call_id, content}` without a type or encoding at `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md:44`. The m-9 consumer review correctly notices this at `master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-012600.md:34-40`, but cannot pin the provider/request owner's schema "on its side." The amendment assigns `LLMRequest` bytes to m-8 (`master/STEP-3-MVP-AMENDMENT.md:83`).

Required correction: m-8 pins the MVP type and encoding in its contract, rehashes, obtains an m-9 consumer rereview on the revised bytes, and only then routes the fresh m-8 implementer final-byte review. `string` is the already-proposed bounded MVP branch; a different branch needs explicit owner rationale.

### F73 -- BLOCKER before stage 6: L1/L5/L7 are pre-lock owner work, not lock-time authorship

- **L1 is not optional polish.** m-7 declares full-domain `config_generation <uint64>` as a JSON number in trust-bearing serve-stamp and relay-evidence objects (`m-7 ...transport-broker.md:260,268,291`). m-10 proves why trust-bearing full-domain counters cannot safely cross JCS JSON as numbers and requires canonical decimal strings (`m-10 ...seam-contract.md:32-39`). Resolve this in owner bytes by string encoding or an explicit narrower numeric domain, then review and reconfirm. Since m-7 already owes F70 bytes, fold L1 in the same owner revision.
- **L7 is now accepted, so it is no longer merely a candidate.** `step3-mvp-design-m8/RECONCILE-planner-20260717-020500.md:21-27` accepts the exact six-field `connector_assign` shape. m-10 must fold it, rehash, take a fresh pair review, and rerun affected confirmations before stage 6.
- **L5 belongs in m-9's stage-4 owner bytes.** The worker design must pin whether the shared client is inside `m9_worker_build_digest` or separately built and transitively covered by `release_digest`. The interface lock verifies and binds that decision; it does not originate it.

The valid sequence for every L1/L7 change is: owner decision -> owner bytes -> new hash -> uniquely-parented implementer review -> affected consumer confirmations/rebases -> Master+VP lock over those exact final hashes.

### F74 -- REQUIRED record corrections

- Target `IN_REPLY_TO` points to `...-043205`; the latest VP prerequisite approval is `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-044033.md`. The corrected close must reply to this review and retain `043205` only as historical lineage.
- Target `:61` says "All 13 confirmation relays." Its table has 13 edges represented by 9 response files, and the required census is 16 edges. Count edges and files separately.

## Disposition

The five owner hashes and their final pair approvals are real at the reviewed snapshot:

- m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`
- m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`
- m-3 `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`
- m-7 `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`
- m-10 `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453`

The m-7 `own-supervised-process` grill outcome is inside the ratified option set and remains valid. Eleven of the target's thirteen listed edges are clean on the cited bytes; rows 1 and 5 fail F70, and three required edges are absent under F71. Stage 1 is therefore **CLOSE-PENDING**, not closed.

Preserve the already-issued m-8 stage-2 and m-9 stage-3 DESIGN-only routing. They may continue nonlocking authoring/review, but:

- hold m-8 implementer final-byte review until F72 is folded and m-9 rereviews;
- hold m-9 stage-3 final closure and reciprocal confirmation until the repaired m-7 hash is rebased;
- issue no stage-6 interface lock, PLAN, T4 code token, credential, provider call, release binding, merge, or deployment authority.

## Required Return

Return one corrected close packet after:

1. m-7 folds F70 + L1, receives fresh final-byte pair approval, and all affected m-7-touching confirmations are refreshed;
2. the missing m-2 -> m-7, m-3 -> m-9, and m-3 -> m-10 confirmations land;
3. m-8 folds F72, m-9 rereviews the revised request schema, and m-8 receives fresh implementer final-byte review;
4. accepted L7 is folded and reviewed in m-10 owner bytes before lock, while m-9's later stage-4 bytes pin L5;
5. the close packet carries the corrected 16-edge pre-m-8 table on current hashes, distinguishes edges from relay files, and threads from this VP review.

No operator decision is required unless an owner proposes changing the ratified topology or claim boundary.

## Verification

- Exact-file lint of the target relay passes.
- Target and all five stage-1 owner hashes were recomputed in this review; they match the values above.
- Operative MVP amendment remains exact `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- `frank/` is clean on `main@502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`, tracking `origin/main` at `+0/-0`.
- This reviewer relay exact-file lint ends `OK`; its INDEX row is present exactly once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260717-021858.md` and appended its `master/relays/INDEX.md` row; no governing source, historical relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean on `main@502e06cc07b5` at `+0/-0` versus `origin/main`.
Next requested action: route the bounded owner-byte and missing-confirmation corrections above, then return the current exact hashes and complete 16-edge table for a fresh VP close review.
