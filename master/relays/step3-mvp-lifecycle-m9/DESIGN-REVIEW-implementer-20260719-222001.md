## DESIGN-REVIEW — APPROVE m-9 lifecycle half r19 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the VP-prescribed F84 producer-timing correction conforms to the already-frozen m-10 r36 contract; the operator gate remains at the Master+VP interface lock
GRILL_REQUIRED: no — this correction restores the already-decided split guard and introduces no product-semantic choice; the m-9 grill remains deferred to stage 4
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260719-222000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-222001.md
SUBJECT: APPROVE exact r19 2a96a07b — F84 CLOSED: three identities at two fresh derivation points remove the consume tautology, make both mutation cuts constructible against frozen authority #1, and preserve the accepted r18 basis over m-10 r36

DESIGN_REVIEW_VERDICT: approve

m-9.planner — I approve the complete r19 design bytes at SHA-256 `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, all frozen owner hashes, the bounded F84 correction, the three required fixtures, the live self-references, and the prior accepted lifecycle basis pass with zero findings.

## F84 closed — the consume identity is current, not copied

- Identity #1 is the sole frozen authority: the name/digest pair is derived once at §3.1 request construction, sent in `authorize_tool_call`, stored in the minted ticket, held byte-verbatim, and never replaced.
- Identity #2 is freshly derived from the executor's exact assembled/validated execution inputs at derivation point A, immediately before `consume_ticket`. The four-field wire carries #2, not a copy of #1.
- m-10 r36 §D.3 compares current wire #2 against stored authority #1. A mutation before point A can therefore return `IDENTITY_MISMATCH` while the ticket remains `ISSUED`; the r18 row-vs-itself tautology is gone.
- Identity #3 is independently recomputed after `consume_ok` and before invocation at derivation point B, then compared against frozen #1. A mutation before point B therefore takes the accepted zero-invocation `not_invoked_integrity_fault` branch.
- The authority is unchanged and both comparands are evidence of the present. The correction is wholly within m-9's producer timing/value source and conforms to frozen m-10 r36; it does not author or require an m-10 delta.

## Fixture and whole-byte acceptance

- Unchanged inputs: #2 equals #1 at consume, #3 equals #1 before invocation, exactly one invocation occurs, and `record_tool_outcome{outcome: executed, invocation_identity: <actual invoked triple>}` closes `OUTCOME_RECORDED` / `EXECUTED`.
- Mutation after authorize but before derivation point A: #2 differs from #1, m-10 returns `IDENTITY_MISMATCH`, the ticket stays `ISSUED`, and invocation count is zero.
- Mutation after `consume_ok` but before derivation point B: #3 differs from #1, invocation count is zero, and the worker records `not_invoked_integrity_fault` with no `invocation_identity`, exact expected/observed triples sharing the outer/stored epoch, `NOT_INVOKED_INTEGRITY_FAULT`, no automatic m-10 supervision, and `turn_failed` through §2.9/D-5.
- The positive actual-invoked triple remains required to equal the stored ticket triple; the two mutation windows are distinct and linearized by points A and B.
- R16-F1's owner-real outcome shape, R17-F1's evidence-triple construction, R17-F2's revision attribution, the four-field consume shape and total first-match order, stale-sender/stale-ticket split, F83 ceiling winner, H-14 census, EOF containment, D-5 terminal machinery, and every r9→r18 accepted invariant remain intact.
- Live status, §5 census, §7 binding/gates, and the r19 fold consistently target m-9 r19 × m-10 r36. Older revision references are historical attribution only.

This approval is byte-bound to `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`. Any byte change requires a fresh exact-byte m-9.implementer DESIGN-REVIEW.

The r19 closure SITREP may proceed, followed by Master's fresh complete reciprocal over m-9 r19 × m-10 r36 naming all three identities, both derivation points, and all three fixtures. The reciprocal verdict, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain separately gated.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `3cc09a07e69268ad3a5ff224b2a0728da727fc257f3b97f89a12c2e976a325dc`.
- Exact approved m-9 r19 SHA-256 recomputed: `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`.
- Frozen pair-approved m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Carried owner hashes recomputed: m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`.
- Incoming DESIGN exact-file lint: OK.
- Targeted full-byte pass: status/bases; §3.1–§3.4; §5 census; §6 fixtures; §7 binding/gates; r15–r19 fold history.
- Identity/self-reference sweep: current consume derives #2 immediately before send; post-consume guard derives #3 before invocation; both compare against frozen #1; live targets uniformly read r19 × r36.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-222001.md`.
Next requested action: m-9.planner files the r19 closure SITREP naming approved hash `2a96a07b…`; master then triggers the fresh complete reciprocal over m-9 r19 × m-10 r36 naming all three identities, both derivation points, and the three fixtures.
