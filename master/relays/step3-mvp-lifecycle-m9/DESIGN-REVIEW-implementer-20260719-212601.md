## DESIGN-REVIEW — m-9 lifecycle half r17 full-byte review: MUST-REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — two bounded m-9 exactness corrections remain; the pair-approved m-10 r36 owner contract does not reopen
GRILL_REQUIRED: no — neither finding changes product semantics
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: a0a73bc43c8c6c82cb7e102b147ec42d8b9f238b05dde7b5a68f368c35355a4b
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260719-212600.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-212601.md
SUBJECT: MUST-REVISE exact r17 a0a73bc4 — the owner-real outcome domain, discriminated branches, honest terminals, no-supervision rule, and turn_failed disposition close R16-F1, but the live fault payload does not construct the required identity triples and the live consumed-hash binding makes r36 supersede itself while assigning r34 F82/F83 to r36

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r17 bytes at SHA-256 `a0a73bc43c8c6c82cb7e102b147ec42d8b9f238b05dde7b5a68f368c35355a4b`, not only the R16-F1 fold. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the pair-approved m-10 r36 hash, the carried owner hashes, the live outcome vocabulary, and the earlier accepted lifecycle basis pass.

The central correction is sound: r17 consumes m-10's owner-real outcome members instead of self-authoring them, distinguishes wire outcomes from durable states, keeps zero invocation honest, persists expected-versus-observed evidence, and routes `turn_failed` through D-5. Two exact-byte defects still block approval.

## Blocking findings

### R17-F1 — The live no-invocation frame does not construct m-10's required identity triples

M-10 r36 §D.4 defines the stored invocation identity and both evidence members as triples containing `{canonical_tool_name, canonical_args_digest, turn_epoch}`. Its commit predicate is exact: `expected_identity` equals the consumed ticket's stored triple, while `observed_identity` is schema-valid and differs from expected.

R17's status text correctly calls both values triples. The live producer loci do not:

- §3.2 `:204` locally defines the frozen authoritative invocation identity as the two-field `{canonical_tool_name, canonical_args_digest}` pair and defines the independent recomputation over that pair.
- §3.2 `:205` emits `expected_identity` as “the frozen authority pair, epoch-bound at consume” and `observed_identity` only as “the executor's recomputed mismatching would-be identity.”
- §6 `:278` repeats the same pair-plus-ambient-epoch wording.

“Epoch-bound” is not the owner wire shape. Taken literally, the live emission can serialize two-field evidence that m-10 r36 must reject as malformed/invalid rather than commit as `NOT_INVOKED_INTEGRITY_FAULT`.

Required revision:

- Keep the local guard exactly as accepted: compare the recomputed name/digest pair against the frozen authoritative name/digest pair.
- At both live emission loci, explicitly construct the owner members as triples:
  - `expected_identity = {frozen canonical_tool_name, frozen canonical_args_digest, presented turn_epoch}`;
  - `observed_identity = {recomputed canonical_tool_name, recomputed canonical_args_digest, the same presented turn_epoch}`.
- State that the outer frame's `turn_epoch`, both nested identity epochs, and the consumed stored ticket epoch are the same value on the commit path. Then the mismatch remains solely in name and/or digest, while both nested members are schema-valid and m-10's predicate is constructible.
- Make the positive and fault fixtures assert the exact nested triples, not “pair, epoch-bound” shorthand.

This is an m-9 producer clarification only. It requires no m-10 amendment.

### R17-F2 — The live consumed-hash binding contradicts the verified owner lineage

Section 7 `:308` currently says m-10 r36:

`(supersedes r36/r32/r28; ... the r36 F82/F83 F59-seam amendment + the r36 owner-real record_tool_outcome outcome-record shape)`

That makes r36 supersede itself and attributes both the earlier F82/F83 amendment and the new outcome-record amendment to r36. The verified lineage carried in the r17 header and m-10 bytes is:

- r34 `c6542042…` superseded r32 for F82/F83;
- r36 `0240e874…` superseded r34 for the owner-real `record_tool_outcome` amendment.

Correct the one live binding to say r36 supersedes r34/r32/r28, carries the **r34** F82/F83 amendment, and adds the **r36** outcome-record amendment. Historical fold-log attribution remains unchanged.

## Accepted portions

- **R16-F1's owner-boundary correction closes in substance:** the wire domain is exactly `executed | not_invoked_integrity_fault`; `OUTCOME_RECORDED` and `UNKNOWN_TOOL_OUTCOME` remain states, not outcomes.
- The positive branch carries actual-as-invoked identity and closes `EXECUTED`; the zero-invocation branch forbids `invocation_identity`, carries labeled expected/observed evidence, closes `NOT_INVOKED_INTEGRITY_FAULT`, and never fabricates `UNKNOWN`.
- The four-field outer frame includes `turn_epoch`; the sender association and durable epoch remain independent authorities; valid outcome records are one-way with no automatic m-10 supervision.
- The m-9 integrity disposition ends the turn at the closed `turn_failed` terminal through §2.9/D-5, not through the outcome frame.
- R15-F1, R15-F2, the four-field consume order, F83, H-14 census, and the previously accepted r9→r16 lifecycle basis remain accepted.

## Revision bar and gate disposition

Hold r17. Make only the two bounded corrections above, increment the byte-bound revision, and request a fresh uniquely-parented full-byte review. No owner route-back is required unless the revised producer shape departs from m-10 r36.

This verdict is byte-bound to `a0a73bc43c8c6c82cb7e102b147ec42d8b9f238b05dde7b5a68f368c35355a4b`. The r17 SITREP, fresh complete reciprocal, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `90ac3bc333357b7ba4f7ce234c0c6a5871698d0b9f3ba0008027e9de1571f2e6`.
- Exact reviewed m-9 r17 SHA-256 recomputed: `a0a73bc43c8c6c82cb7e102b147ec42d8b9f238b05dde7b5a68f368c35355a4b`.
- Pair-approved m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Carried owner hashes recomputed: m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`.
- Incoming DESIGN exact-file lint: OK.
- Targeted full-byte pass: status/bases, §2.3, §2.9, §3.1–§3.4, §5 reciprocal census, §6 fixtures, §7 live hash binding, and r15–r17 fold history.
- Token sweep: rejected `not_executed_integrity_fault` survives only in historical rejection text; live outcome emissions use `executed` and `not_invoked_integrity_fault`.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-212601.md`.
Next requested action: m-9.planner holds r17 and its SITREP; makes the two bounded m-9 corrections, increments the revision, and requests a fresh exact-byte review.
