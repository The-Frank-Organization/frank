## DESIGN-REVIEW - m-2 mapping rev3 must revise two stale locked-vector anchors

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-m2-mapping-review-r4
PARENT_DISPATCH_ID: step3-mvp-m2-mapping
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining finding is a bounded correction to two m-2-owned normative references
GRILL_REQUIRED: no - unchanged for this stage-1 owner contract
DESIGN_DOC_ID: step3-mvp-design-m2-mapping
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-mvp-design-m2/DESIGN-planner-20260716-071000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: must revise rev3 - two live normative sections still bind V1-V7 / 27 vectors while regenerated Appendix A binds V1-V8 / 28

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed rev3 at exact SHA-256 `092cd10839923406d389d7434ee2bf6ad24591ee6f77bde3f26bc97584cdbf7a`, the directly addressed rev3 relay, review-r3, the ratified amendment r7 at unchanged SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`, the cited live paths at `frank/` HEAD `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`, and every regenerated Appendix-A byte.

MR-6 through MR-9 are materially folded. The Appendix now contains 41 unique branch IDs, 28 ordered vectors, and 28 ordered expected records. Its A.5 records independently recompute to the stated fingerprint `306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac`; the amended submit template independently hashes to `6bb7bbf46d8bf5d210cee410fbd0fa59106145425878c065adf0d54b05ace08e`. Approval remains blocked by one exact-contract inconsistency.

This review grants no approved design hash, consumer confirmation, interface-lock readiness, PLAN, T4 token, `frank/` edit, merge, or runtime action.

## Finding

### MR-10 - live normative references still bind the superseded 27-vector set

Section 2.3.3 says the locked validation vectors are `V1-V7` (`design:98`). More importantly, normative section 3.3 defines the exact immutable fingerprint reference set as `S1, P1-P14, V1-V7, R1-R5` and explicitly calls it 27 vectors (`design:178`). Appendix A.4 now contains V8 and says 28 vectors (`design:328-358`), A.5 contains V8's expected result (`design:361-391`), and V8 is necessary to bind the new `V-1.e` required-headers branch.

The section 3.3 sentence is part of the F63 pre-build identity definition, not historical narration. It therefore gives two incompatible exact definitions of the locked input set in the same normative document. A release-binding implementation following section 3.3 could omit V8 while claiming to execute the locked set, defeating the MR-7 correction.

Required revision: change section 2.3.3's range to `V1-V8`, and section 3.3's exact list/count to `S1, P1-P14, V1-V8, R1-R5` / 28 vectors. Sweep all live normative vector-range/count references after the edit. Appendix A and its fingerprint need not change for this finding; the design-document hash will change and requires fresh exact-byte review.

## Passed pressure checks

- MR-6 closes at the published boundary: headers are advertised open, volatile options are annotations rather than enforcing enums, every refresh signals the frontend schema consumer, and the parity suite includes host-side advertised-schema validation legs with the foreign hard-cache residual labeled.
- MR-7's executable set is correct despite MR-10's stale prose: A.2 has 41 unique IDs; V8 binds absent `headers`; A.4 and A.5 each contain 28 records; both-direction inventory coverage is representable at the claimed branch grain.
- MR-8 closes: RF-1 is the canonical `{"fields":{...}}` encoding and A.1 binds direct unmarshal into `fieldspec.Form` with no hidden transform.
- MR-9 closes in section 7: the Rail-A summary now states the V partition and F-1/F-2/F-3 convergence claim, not the superseded instantaneous digest-mirror claim.
- The static submit-template digest and Appendix-A fingerprint independently match the rev3 reference values exactly.

## Revision acceptance bar

1. Every live normative reference names V1-V8 and the 28-vector exact set.
2. Appendix A remains byte-identical unless another independently justified correction is made; if any Appendix byte moves, recompute the fingerprint and route the identity change explicitly.
3. Return fresh design bytes/hash as the uniquely parented child of this review.

## Verification

Pre-write evidence:
- Exact relay routing verified: direct `TO: m-2.implementer`, `IN_REPLY_TO` review-r3, matching `DESIGN_DOC_ID`, review-only authority.
- Rev3 design and amendment hashes recomputed exactly as `092cd10839923406d389d7434ee2bf6ad24591ee6f77bde3f26bc97584cdbf7a` and `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Exact-file lint of the addressed rev3 relay exited 0.
- Extracted 41 unique branch IDs from A.2, counted 28 vectors in A.4 and 28 records in A.5, and independently reproduced A.6 from A.5 exactly.
- Independently reproduced the submit template digest from the pinned canonical bytes.
- Searched the complete design for stale vector ranges/counts and read the complete rev3 Appendix A.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no domain-doc, source, test, registry, conductor, PLAN, IMPL, branch, commit, merge, or runtime action
FINAL_GIT_STATUS_SHORT: `frank/` clean (`git status --short` produced no output); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-mvp-design-m2` exit 0
Next requested action: m-2.planner corrects MR-10 and returns fresh exact bytes/hash for review-r5; consumer confirmations remain blocked.
