## DESIGN-REVIEW — m-9 lane-2 r8 must revise: S-1 is not total over its own evidence tuple; §9 denies the S-1/S-2 folds; the new path-negative proof claims impossible zero occurrences

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r8
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — all findings are determinate contract-totality and internal-consistency corrections; Master's §2.6 ruling is obeyed and no operator-ratified claim is adjudicated here
GRILL_REQUIRED: no — no product choice or unresolved ownership boundary is decided by this review
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
DESIGN_DOC_SHA256: 563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260723-101500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-2.planner, m-3.planner, m-3.implementer, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260723-140000.md
SUBJECT: MUST-REVISE exact m-9 r8 563398c0 — Master's §2.6 hold and §6/§8 byte identity pass, but S-1 leaves same-key/same-marker/different-locator receipts unclassified; §9/§11 still call the newly folded S-1/S-2 inputs parked and unconsumed; both m-1 negative guards falsely claim their own named strings occur nowhere

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r8 artifact at exact SHA-256 `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c`, with your `RECONCILE-planner-20260723-111500` gate correction and Master's authoritative `RECONCILE-orchestrator-planner-20260723-103000` ruling applied. **MUST-REVISE.** Three findings survive.

## M9-DAG-R8-F1 — S-1 is not total over its declared six-member equivalence tuple

Section 2 correctly defines duplicate equivalence over `{run_id, turn_id, attempt_id, marker_digest, segment_id, seq_hwm}` (`:285`), but its conflict rule fires only for a same-key **different `marker_digest`** (`:287`). The input

`same {run_id, turn_id, attempt_id}` + same `marker_digest` + different `segment_id` or different `seq_hwm`.

is therefore not equivalent and is not marker-conflicting. On the receiver it reaches the already-occupied unique key without any selected disposition. The ordered contract is not total over its own stated domain.

This is independently visible in the m-9 bytes. The later m-10 r7 review found the symmetric receiver defect and current proposed m-10 rev8 `00b8401d…` changes conflict to the exact complement of equivalence; that cross-lane result corroborates rather than creates this finding. The m-10 `…-134500` request is addressed to you, not to me; I do not act on its CC routing.

**Required correction:** define `receipt_conflict` for **any same-key non-equivalent evidence tuple**, with the first committed tuple standing. Preserve equivalent-duplicate before stale-sender. Strengthen the “not producible by correct m-9” statement from marker-only reasoning to the whole evidence tuple: a correct replay of one receipt re-emits all six evidence members byte-identically. Carry an all-equal idempotent leg plus independent different-`segment_id` and different-`seq_hwm` conflict legs; align Status and the r8 fold record.

## M9-DAG-R8-F2 — §9 and §11 deny the S-1/S-2 folds r8 claims

R8 says the exact S-1 frame is jointly settled and folded into §2 (`:275`, `:290`) and calls the S-2 frames settled and folded into §3 (`:301`). But §9 still says:

- item 4 is parked “before I fold it into §3's no-work gate”;
- item 5 is parked until the joint frame settles; and
- “Nothing in this delta consumes any of the four remaining parked inputs (2–5) as though it were settled” (`:405-407`).

Section 11 then still owes F73 confirmations on all four items and separately owes the joint content-ready frame (`:438`). These live statements cannot all describe the same r8 state. This is the same self-denying parking class as R6-F2, now across two newly folded items rather than one.

**Required correction:** make the staging ledger truthful and preserve the join gate. Items 2 and 3 remain parked. Items 4 and 5 are now carried as exact joint-contract folds but are **not normative** until both pairs approve matching bytes and co-sign the §D join. Update §9's closing sentence, §11's owed list, Status, and the fold log consistently. If instead you intend items 4/5 to remain genuinely parked, retract the §2/§3 “settled/folded” claims; do not describe the same input as both folded and unconsumed.

The current m-10 producer rev8 is still proposed and not pair-approved, so this correction must not convert the joint-pending state into a premature producer approval. Rebase to its eventual pair-approved exact hash before the join.

## M9-DAG-R8-F3 — both new path-negative evidence claims are self-falsifying

The new m-1 guards are directionally correct, but their proof language is false on the bytes that carry it:

- §1.12 says `session_log_path` is “Verified absent from ... this delta” while naming that exact token; the r8 fold record names it again.
- §7 says `workspace_root_path` “appears nowhere in ... this delta” while naming that exact token; the r8 fold record names it again.

Fresh exact-token census returns **two occurrences of each token in this delta**. Worker r7 and lifecycle r21 have zero, but r8 cannot use a lexical-zero claim after adding the explicit negative guard. The real invariant is carrier-scoped, not string-scoped.

**Required correction:** replace both zero-occurrence claims with the mechanically true statement: outside the explicit normative-negative and lineage/history clauses, no conductor record, projection, INDEX row, typed error, or E0 body carries the path. Verify the actual carrier/member census; do not use raw token absence as proof once the contract necessarily names the forbidden member.

## Passed pressure checks

- **Master's §2.6 ruling conforms.** R8 does not fold, imply, or pre-commit the Gate-2 relabel. The relabel and the r7 `turn_failed` clarification remain for the single §D-settlement amendment's accept-additive-carry versus explicit-r21-amendment decision. No revision may fold §2.6 before that amendment ratifies.
- **§4 carrier claim passes.** The run-wide parked set rides m-10-owned `turn_open` and `attempt_open_ok`; continuation ancestry also carries `uncertain` manifest entries. M-9 correctly authors no durable re-surfacing copy.
- **S-2 semantics pass apart from the staging contradiction.** The reportable domain, present-IFF rule, committed-pair receipt/conflict carriage, adopt-and-proceed behavior, and no-work receipt gate agree with the settled pair exchange.
- **Five S-3 action families pass.** The local encodings agree with the pair settlement and `relay.*` remains held for m-2's pair-approved form-derived target shape.
- **§6 and §8 are byte-identical to approved r7.** I extracted the exact r7 section bytes captured during the `…-081500` review and compared them directly with r8: §6 `cmp` exit 0 at section SHA-256 `aa1bb710b27db15948a3f317f0bb11cd4ffea51d8d47ebebcbf2529950fb77dc`; §8 `cmp` exit 0 at `13ba55064d8fd12bfbcd9d499579b1fc2f8f98df014c3fd4906ad068b07ee853`. That identity fact passes, but **r8 remains unapproved**, so m-3/m-10 must not bind `563398c0…` as a producer basis.
- Frozen lifecycle r21 `4d3bd14e…` and worker r7 `cb7ff970…` remain byte-exact and unedited.

## Revision acceptance bar

1. Make the S-1 same-key split total: all six evidence members equal ⇒ equivalent duplicate; any member differs ⇒ `receipt_conflict`, first tuple stands; add the three named legs.
2. Reconcile §2/§3 with §9/§11: items 2/3 parked, items 4/5 exact-folded but joint-pending, with no normativity before matching pair approvals plus the §D co-sign.
3. Replace the impossible m-1 lexical-zero assertions with a carrier-scoped negative census and sweep Status/fold-history echoes.
4. Preserve Master's §2.6 hold, the consolidated §D-amendment adjudication, the `relay.*` hold, §4's m-10-only carriers, and the exact §6/§8 content.
5. Return one fresh full-document hash for byte-bound re-review. Until then, m-3/m-10 do not rebase to r8 and no §D join, amendment, integrated re-lock, PLAN, T4/code, provider, E3, merge, or deploy gate advances from this verdict.

## Exact evidence

- R8 DESIGN relay SHA-256: `c6665fa17564100705a8300d98f62a832b565b6c8d366bf443de55c8f74499fb`; exact-file lint OK; directly `TO: m-9.implementer`.
- Review-gate correction SHA-256: `e4cd7f20ff2afcd13788d1dc3dc446fcb711d6f9b9b46d2af64f71fe36b9704e`; exact-file lint OK; directly `TO: m-9.implementer`.
- Current m-9 delta SHA-256: `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c`.
- Master's ruling SHA-256: `3de26df295d40914df1082f3125cc65dcf99f4bf166a88ad8856349b045458a9`.
- Current proposed m-10 producer rev8 SHA-256: `00b8401dfbb4f12b1e0f69d58b7ccafda4a8ff3ab067418d2396b55249e07683` — contextual, not treated as pair-approved authority.
- Frozen lifecycle r21 SHA-256: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; frozen worker r7 SHA-256: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.

## Authority boundary

This is a byte-bound m-9 pair review only. It does not adjudicate the amendment-borne r21 claims, approve m-10 rev8, answer the m-10 planner's directly addressed carrier request on m-9.planner's behalf, co-sign the §D join, or authorize any downstream rebase or action gate.

ACTIONS_GIT_REF: docs-workspace action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design document, frozen artifact, `frank/` source, branch, commit, joint record, amendment, lock, PLAN, T4 token, credential, provider call, release binding, merge, or deploy action
FINAL_GIT_STATUS_SHORT: unavailable — harness root is not a git repository; `frank/` clean at `c78da3815a34480590071295c1e09bb7d53c10b6`
RELAY_LINT: OK — exact-file and dispatch-root lint exit 0
Next requested action: m-9.planner folds M9-DAG-R8-F1..F3 together with the directly routed m-10 S-1/carrier response, returns one fresh r9 hash for full-byte review, and otherwise holds the final post-ratification §7-`relay.*` + §2.6 batch.
