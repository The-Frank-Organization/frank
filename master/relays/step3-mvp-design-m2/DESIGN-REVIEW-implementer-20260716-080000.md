## DESIGN-REVIEW - m-2 mapping rev4 must revise the stale document revision header

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-m2-mapping-review-r5
PARENT_DISPATCH_ID: step3-mvp-m2-mapping
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining finding is one m-2-owned metadata correction
GRILL_REQUIRED: no - unchanged for this stage-1 owner contract
DESIGN_DOC_ID: step3-mvp-design-m2-mapping
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-mvp-design-m2/DESIGN-planner-20260716-074500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: must revise rev4 - MR-10 and exact Appendix identity pass, but the live design header still identifies the document as rev2

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed the current design at exact SHA-256 `62efe9636d6e36b0e113c965cf670bc2a011d85cd46c92860053d1d2bc87cb94`, the directly addressed rev4 relay, review-r4, the unchanged amendment r7 at SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`, and the exact rev3-to-rev4 byte delta.

MR-10 is closed: sections 2.3.3 and 3.3 now both bind `V1-V8` and the 28-vector set. Reversing those two edits and deleting the new rev4 log line reproduces the prior approved-for-revision rev3 SHA `092cd10839923406d389d7434ee2bf6ad24591ee6f77bde3f26bc97584cdbf7a` exactly, proving the submitted delta is limited to the claimed three lines and Appendix A is byte-identical. Approval remains blocked by one stale live metadata locus.

This review grants no approved design hash, consumer confirmation, interface-lock readiness, PLAN, T4 token, `frank/` edit, merge, or runtime action.

## Finding

### MR-11 - the live document header still self-identifies as rev2

The design header says `DESIGN_DOC_ID: step3-mvp-design-m2-mapping - rev2` and summarizes only the rev0-to-rev2 history (`design:3`). The directly addressed relay calls the exact bytes rev4, and the normative revision log contains rev3 and rev4 entries (`design:281-297`). The same durable design record therefore has conflicting live revision identity at its primary metadata locus.

This is not historical narration: consumers opening the document read line 3 as its current revision identity. Pair-approving a hash described simultaneously as rev2 and rev4 would make later confirmation and interface-lock citations ambiguous even though the hash is exact.

Required revision: update the line-3 live revision marker to `rev4` and make its compact revision summary include or defer cleanly to the rev3/rev4 entries in section 9. Sweep live document metadata for any other stale current-revision marker. Do not alter Appendix A or mechanism text for this finding.

## Passed pressure checks

- MR-10 is exact: section 2.3.3 names V1-V8; section 3.3 names `S1, P1-P14, V1-V8, R1-R5` and 28 vectors.
- The exact inverse-delta reconstruction hashes to rev3 `092cd108...` byte-for-byte; only the two MR-10 loci and one rev4 history line moved.
- Appendix A still has 41 unique branch IDs and 28 ordered vectors; A.5 independently recomputes to `306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac`.
- The prior MR-6 through MR-9 closures, submit template digest, ownership boundaries, and downstream authority holds remain intact because no related byte moved.

## Revision acceptance bar

1. The live header identifies these bytes as rev4 and does not truncate the current revision history at rev2.
2. Appendix A remains byte-identical; any independently justified Appendix change requires explicit identity handling and fingerprint recomputation.
3. Return fresh design bytes/hash as the uniquely parented child of this review.

## Verification

Pre-write evidence:
- Exact relay routing verified: direct `TO: m-2.implementer`, `IN_REPLY_TO` review-r4, matching `DESIGN_DOC_ID`, review-only authority.
- Rev4 design and amendment hashes recomputed exactly as `62efe9636d6e36b0e113c965cf670bc2a011d85cd46c92860053d1d2bc87cb94` and `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Exact-file lint of the addressed rev4 relay exited 0.
- Mechanically reversed the two claimed normative edits and removed the rev4 log line; the result reproduced rev3 SHA `092cd10839923406d389d7434ee2bf6ad24591ee6f77bde3f26bc97584cdbf7a` exactly.
- Recounted 41 Appendix branch IDs and 28 A.4 vectors; independently reproduced A.6 from the 28 A.5 records exactly.
- Searched the complete current design for vector ranges, vector counts, revision markers, and MR-10 history.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no domain-doc, source, test, registry, conductor, PLAN, IMPL, branch, commit, merge, or runtime action
FINAL_GIT_STATUS_SHORT: `frank/` clean (`git status --short` produced no output); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-mvp-design-m2` exit 0
Next requested action: m-2.planner corrects MR-11 and returns fresh exact bytes/hash for review-r6; consumer confirmations remain blocked.
