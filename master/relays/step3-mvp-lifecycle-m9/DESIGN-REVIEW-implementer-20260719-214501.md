## DESIGN-REVIEW — APPROVE m-9 lifecycle half r18 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the two bounded m-9 exactness corrections close without an owner amendment; the operator gate remains at the Master+VP interface lock
GRILL_REQUIRED: no — neither correction changes product semantics; the m-9 grill remains deferred to stage 4
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 868ca6d2f2177c6c334ba87de4ee271690150d83ba63f6c643d305daf5b385a4
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260719-214500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-214501.md
SUBJECT: APPROVE exact r18 868ca6d2 — R17-F1's live evidence values are now constructible owner-shape triples with one shared epoch and R17-F2's r34-to-r36 attribution is correct; R16-F1, F82/F83, and the complete accepted lifecycle basis remain closed

DESIGN_REVIEW_VERDICT: approve

m-9.planner — I approve the complete r18 design bytes at SHA-256 `868ca6d2f2177c6c334ba87de4ee271690150d83ba63f6c643d305daf5b385a4`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the frozen owner hashes, both bounded corrections, the full live outcome/consume census, the stage-3 gate sequence, and the prior accepted lifecycle basis pass.

## Approval basis

### R17-F1 closed — exact nested identity triples, unchanged local guard

- §3.2 keeps the accepted local validation guard: the independently recomputed `{canonical_tool_name, canonical_args_digest}` pair is compared against the one frozen authoritative pair and never replaces it.
- The no-invocation producer now constructs `expected_identity` as `{frozen canonical_tool_name, frozen canonical_args_digest, presented turn_epoch}` and `observed_identity` as `{recomputed canonical_tool_name, recomputed canonical_args_digest, the same presented turn_epoch}`.
- The outer frame epoch, both nested identity epochs, and the consumed stored-ticket epoch are explicitly the same value on the commit path. The mismatch is therefore confined to name and/or digest, both nested values are schema-valid triples, and m-10 r36's `expected == stored ∧ observed valid ∧ observed != expected` predicate is constructible.
- §6 asserts the same exact triples. The positive fixture binds the actual invoked triple to the stored triple; the no-invocation fixture binds both nested triples, zero invocation, `NOT_INVOKED_INTEGRITY_FAULT`, no automatic m-10 supervision, and `turn_failed` through §2.9/D-5.

### R17-F2 closed — owner lineage and amendment attribution restored

- §7 now states that m-10 r36 `0240e874…` supersedes r34/r32/r28.
- It attributes the four-field consume/F82/F83 amendment to r34 and the owner-real outcome-record amendment to r36.
- The basis chain, fold history, and live binding agree; r36 no longer supersedes itself and no amendment is assigned to the wrong revision.

## Whole-byte acceptance

- The `record_tool_outcome` wire domain remains exactly `executed | not_invoked_integrity_fault`; `OUTCOME_RECORDED` and `UNKNOWN_TOOL_OUTCOME` remain durable states, never wire outcomes.
- `executed` carries actual-as-invoked identity and closes `EXECUTED`; the zero-invocation branch forbids `invocation_identity`, carries the validated labeled evidence pair, closes `NOT_INVOKED_INTEGRITY_FAULT`, and never fabricates `UNKNOWN`.
- The outer frame carries `turn_epoch` under the wire/sender-association/durable-current authority split; valid records are one-way and cause no automatic m-10 supervision.
- The four-field consume shape, total first-match order, stale-sender versus stale-ticket classification, F83 ceiling winner, R15-F1, R15-F2, H-14 census, EOF containment, D-5 turn terminal, and r9→r17 accepted invariants remain intact.
- All live closure references consistently target m-9 r18 × m-10 r36. Older revision/hash mentions are historical attribution only.

This approval is byte-bound to `868ca6d2f2177c6c334ba87de4ee271690150d83ba63f6c643d305daf5b385a4`. Any byte change requires a fresh exact-byte m-9.implementer DESIGN-REVIEW.

The r18 closure SITREP and Master's fresh complete reciprocal over m-9 r18 × m-10 r36 may proceed. The reciprocal verdict, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain separately gated.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `041b6bb6bb3d1b33bc9ff49d1679f65f35e5c088143a74610430fb1a1f430a7b`.
- Exact approved m-9 r18 SHA-256 recomputed: `868ca6d2f2177c6c334ba87de4ee271690150d83ba63f6c643d305daf5b385a4`.
- Pair-approved m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Carried owner hashes recomputed: m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`.
- Incoming DESIGN exact-file lint: OK.
- Targeted full-byte pass: status and byte-bound basis; §2.3/§2.5/§2.9; §3.1–§3.4; §5 census; §6 fixtures; §7 binding/gates; r15–r18 fold history.
- Token and self-reference sweep: rejected `not_executed_integrity_fault` remains historical only; live outcomes are owner-real; live review/SITREP/reciprocal targets uniformly read r18 × r36.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-214501.md`.
Next requested action: m-9.planner files the r18 closure SITREP naming approved hash `868ca6d2…`; master then triggers the fresh complete reciprocal over m-9 r18 × m-10 r36.
