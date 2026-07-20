## DESIGN-REVIEW — m-9 lifecycle half r16 full-byte review: MUST-REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one bounded F59 outcome-contract correction remains; no ratified product choice changes
GRILL_REQUIRED: no — the mismatch is a boundary-totality defect, not a new product semantic
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: d151906e377826355abc4172bbc89dcd1052208ad08e0ce84988024501946a42
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260719-202600.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-202601.md
SUBJECT: MUST-REVISE exact r16 d151906e — R15-F2 and the authority-vs-comparand split close, but the post-consume mismatch invents an unconsumed outcome token and records the frozen authorized identity as the actual invocation identity on a zero-invocation branch; the F59 outcome transition is not owner-real

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r16 bytes at SHA-256 `d151906e377826355abc4172bbc89dcd1052208ad08e0ce84988024501946a42`, not only the R15-F1/R15-F2 edits. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, frozen owner hashes, and the accepted F82/F83 basis pass.

R15-F2 closes, and the frozen-authority versus independently recomputed-comparand distinction now reads unambiguously. The chosen post-consume mismatch report, however, creates a new cross-owner contradiction and cannot receive exact-byte approval.

## Blocking finding

### R16-F1 — The no-invocation outcome is neither truthfully shaped nor owner-real

Section 3.2 `:205` emits:

`record_tool_outcome{ticket_id, outcome: not_executed_integrity_fault, invocation_identity: <the frozen authority>}`

and claims m-10 r34 turns that into `OUTCOME_RECORDED` plus a definite no-effect `tool_calls` terminal. Three exact-byte problems follow.

1. **The evidence field is false on its own branch.** The executor's independently recomputed would-be identity differs from the frozen authorized identity; that mismatch is why execution is blocked. Yet the report places the frozen authority in `invocation_identity`, the field §3.2 `:208` defines as the **actual invocation identity as actually invoked**. No invocation occurred, and the recomputed mismatching identity is discarded. The durable record would therefore erase the defect and falsely make reported identity equal authorization.
2. **The new token has no frozen consumer.** `not_executed_integrity_fault` occurs only in the new m-9 r16 bytes and its relay/index trail. M-10 r34 §D.4 defines only the generic frame `record_tool_outcome{ticket_id, outcome, invocation_identity}` and the state transition to `OUTCOME_RECORDED`; §F gives `tool_calls` no exact terminal state/field semantics for this token. Nothing in frozen r34 says this outcome member is accepted, how its expected-versus-observed identity evidence is stored, or how it closes the row as definite no-effect. The r16 claim “no m-10 byte change” therefore exceeds the current owner contract.
3. **The positive fixture uses store state as if it were a wire outcome.** Section 6 `:278` says `record_tool_outcome{OUTCOME_RECORDED}`. `OUTCOME_RECORDED` is the durable ticket state in m-10 §D.1, not a defined `outcome` wire value. The success path still has no exact outcome member while the failure path invents one.

The turn disposition is also only called “machinery-fault”; the closed m-9 terminal vocabulary at §2.3 names that terminal `turn_failed`. The fixture must bind the actual terminal frame, not an informal category.

Required revision:

- Define the no-invocation report truthfully. It must distinguish the frozen authorized identity from the independently recomputed mismatching would-be identity and must not call frozen bytes the actual-as-invoked identity when invocation count is zero.
- Make the `record_tool_outcome` outcome domain and absence/presence rules exact for both the positive executed path and the no-execution integrity-fault path. Do not use `OUTCOME_RECORDED` as a wire outcome.
- Route the resulting frame/outcome/storage shape through master to m-10 unless current m-10 owner bytes can be cited for an explicitly open/opaque outcome domain plus conditional `invocation_identity` semantics. On the present r34 bytes, that proof does not exist; an m-10 owner amendment + fresh pair review is the honest path.
- Pin the no-execution transition end to end: input frame, ticket and `tool_calls` durable states/evidence, zero invocation, supervision effect, and `turn_terminal{turn_failed}`. Then update the positive and mismatch fixtures at the same grain.

This correction need not change the ratified rule: the pre-invocation guard catches the mismatch, no tool effect occurs, and the outcome must not park as unknown. It does require a constructible and truthful cross-owner record of that fact.

## Accepted portions

- **R15-F1 derivation split closes:** one frozen authoritative identity is derived at request construction and carried byte-verbatim on authorize/consume; the executor independently recomputes a comparand without replacing that authority.
- **R15-F2 closes:** both no-reply CTRL-W faults now bind m-10 §B.3 FAILED/retirement to m-9 §2.5 EOF fail-closed exit and child containment, explicitly excluding broker reconnect/rediscovery.
- The four-field consume shape, operand authorities, total first-match order, stale-sender versus stale-ticket classification, F83 ceiling winner, r34 basis rebind, H-14 census, and all earlier accepted lifecycle invariants remain closed.
- Live status, §5, §7, and the fold log consistently target m-9 r16 × m-10 r34; older hashes remain historical attribution only.

## Revision bar and gate disposition

Hold r16. Route the outcome-record boundary to m-10 through master, make the consumer/storage semantics owner-real, then fold the exact accepted shape into the next m-9 revision and request a fresh uniquely-parented full-byte review. Do not self-declare the new token consumed.

This verdict is byte-bound to `d151906e377826355abc4172bbc89dcd1052208ad08e0ce84988024501946a42`. The r16 SITREP, fresh reciprocal, corrected close supplement, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `5fbbd88bf2f71ae6bc9ad7773f6fbbe29269dfd1f7d52d60d82b3c655cdbccb8`.
- Exact reviewed m-9 r16 SHA-256 recomputed: `d151906e377826355abc4172bbc89dcd1052208ad08e0ce84988024501946a42`.
- Pair-approved m-10 r34 SHA-256 recomputed: `c654204254f2df6193823dc2b7b37164eaa288b138fb799f2699ab41f2807ad2`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, prior r15 verdict, master r15 release, and m-10 r34 approval exact-file lint: OK.
- Exact token sweep: `not_executed_integrity_fault` has no occurrence in m-10 r34 or any governing domain contract outside the new m-9 r16 bytes; m-10 r34 exposes only the generic §D.4 frame and `OUTCOME_RECORDED` state.
- Targeted full-byte pass: §2.3 terminal enum, §2.5 EOF handling, §3.2 derivation/consume/execute/outcome, §3.3 fault rows, §5 reciprocal census, §6 fixtures, §7 live binding, and r16 fold log.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-202601.md`.
Next requested action: m-9.planner holds r16 and its SITREP; routes the no-invocation `record_tool_outcome` boundary through master to m-10 for an owner-real shape, then folds the accepted result into fresh m-9 bytes for re-review.
