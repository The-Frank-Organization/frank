## DESIGN-REVIEW — r4 MUST-REVISE: the r3 findings close, but the terminal-seal guard runs after four ordinary prefix-termination branches and therefore does not cover the physical bytes it claims

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r4
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the one remaining finding is an evaluation-order defect inside the selected terminal-seal mechanism; no operator-ratified product choice is reopened
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260722-120000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-1.planner, m-2.planner, m-3.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260722-130000.md
SUBJECT: r4 `26dc372c…` materially closes both r3 blockers, but MUST-REVISE remains on one exact ordering defect — §1.4's terminal-seal branch is evaluated after self-consistency, conflicting-duplicate, gap/regression, and chain branches, so four classes of non-collapsed physical post-seal bytes return an ordinary truncated prefix instead of the promised whole-run fail-close

## Verdict

**MUST-REVISE** on exact design SHA-256 `26dc372c5b4a91b0dca75fd77b12df14aca73d44b98a54f55679fc2db024cf82`.

R4 closes both r3 findings in substance. `boundary_record` is total, including the `segment_open`-only case; the false blanket late-write claim is honestly split into 11a/11b/11c; semantic identity and edge progression are recovery checks; append/create selection is bound to current `assign.generation_id`; and the cross-generation forward-link exception correctly preserves the rotation-seal-then-replacement history without inventing a foreign `generation_id` ordering rule.

One evaluation-order defect prevents approval. The design says **every** non-collapsed physical record after an accepted seal fails closed for the whole run, but the ordered algorithm reaches that rule only after four earlier branches that return ordinary prefix termination.

## M9-DAG-R4-F1 — BLOCKER: branch 6 is too late to enforce terminal-seal semantics

The §1.4 algorithm is explicitly first-match-wins (§1.4:73-81). Its current order is:

1. self-consistency / parse failure ⇒ prefix ends before R;
2. adjacent byte-identical duplicate ⇒ collapse;
3. conflicting duplicate ⇒ prefix ends before R;
4. regression or gap ⇒ prefix ends before R;
5. chain mismatch ⇒ prefix ends before R;
6. last accepted kind is `segment_seal` ⇒ whole-run fail-closed.

Therefore only a well-formed, next-sequence, correctly chained record reaches branch 6. Four non-collapsed post-seal cases contradict the normative “any further physical record R” claim (§1.4:80-83):

- a torn or unparseable line after the seal matches branch 1;
- a same-`seq` different-byte line matches branch 3;
- a regressed or gapped `seq` matches branch 4;
- a next-`seq` record with the wrong `prev_digest` matches branch 5.

Each case reports a clean prefix ending at the seal instead of the whole-run fail-close required by row 11a and property (ii). The existing post-seal complete-round fixture exercises only the one shape that passes branches 1–5, so it cannot prove the universal terminal-seal claim.

**Required correction:** make the terminal state the first state-dependent check after reading the next physical line. If the last accepted record is a seal, compare the raw next line to the exact accepted seal bytes: an immediately adjacent byte-identical repeat may take the existing collapse exception; **every other physical byte sequence**, including torn/unparseable content, conflicting sequence, gap/regression, or wrong chain, must fail closed for the whole run before generic prefix-termination logic. An equivalent file-offset/finality rule is acceptable if it has the same total result. Extend the terminal-seal battery with the four cases above and assert they all take the seal-specific whole-run fault, never branches 1/3/4/5.

## Accepted on these bytes

- **R3-F1 mechanism:** terminal seal plus fail-closed disposition is the right choice; no legal producer path writes after it. `boundary_record` is one unambiguous real-record definition: last honoured marker, otherwise `segment_open` at seq `"0"`. The sealed/unsealed equation coincides, and 11b/11c correctly distinguish a moved trusted boundary from an untrusted suffix. R4-F1 changes only guard priority so the terminal rule matches its stated domain.
- **R3-F2 mechanism:** composition identity, envelope homogeneity, exact same-generation rotation increment, cross-generation rotation reset/non-reuse, and the current-assignment append table close the legal-path gap. Not ordering opaque `generation_id` locally is correct; any stronger order must come from m-10. Restricting seal `next_segment_id` equality to same-generation edges is sound because backward identity and boundary equality still bind the cross-generation replacement edge.
- **Earlier repairs remain accepted:** the per-run `flock`, disposed-but-live fail-closed cut, seal-before-successor order, complete topology predicate, round-key membership, per-kind provenance, and duplicate algorithm all stand.
- **Unchanged surfaces:** §1.5, §1.6, D2's temporal trust split, `content_lost`, the disposition-receipt no-work gate, five first-action branches, §7.1 narrowing, E/C/B ownership, and the five parked consumers remain sound. The m-9/m-10 receipt frame remains correctly joint and parked.
- **Additivity/authority:** worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, amendment rev12 `1125b0a0…`, and m-10 r40 `d2ce9831…` reproduce exactly. No frozen design byte moved; this review takes no PLAN, T4/code, lock, release-binding, E3, merge, or deploy authority.

## Re-review gate

Return one revised additive delta changing the terminal-seal evaluation order and its fixtures, with a fresh full-document SHA-256. Preserve every accepted r4 mechanism. Re-run the existing battery plus post-seal torn/unparseable, conflicting-same-seq, gap/regression, and wrong-`prev_digest` inputs; every non-identical physical suffix after a seal must reach the same whole-run fail-close. Do not file the lane SITREP, F73 confirmations, or §D join record from r4.

## Verification

- Reviewed design SHA-256: `26dc372c5b4a91b0dca75fd77b12df14aca73d44b98a54f55679fc2db024cf82`.
- Incoming DESIGN relay SHA-256: `1fa85daf542d090d4edb56b357772dc6f4f3b2e30e0ba271201d37cca1768226`; exact-file lint: `OK`.
- Released rev2 dispatch SHA-256: `254950dd5e164d151739aff827efd1f8ea67887832ca9b969334370453003f25`.
- Frozen bases reproduced: worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; lifecycle r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260722-130000.md`.
Next requested action: m-9.planner folds M9-DAG-R4-F1 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; all downstream lane-2 gates remain held.
