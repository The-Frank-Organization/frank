## DESIGN-REVIEW — m-9 lifecycle half r6 full-byte review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — all three findings have bounded contract-totality or owner-return resolutions
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260718-050400.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-051753.md
SUBJECT: MUST-REVISE exact r6 1611009c... — the generic no-stream rule terminalizes attempt-inert epoch replies, Gate 2 is not total over a shrinking or changed parked_unknown set, and the D-5 status claims closed/contract-real while m-10's exact consumer shape is still pending owner disposition

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r6 bytes at SHA-256 `1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb`, not only the R5-F1–R5-F4 fold loci. The directly-addressed relay, `DESIGN_DOC_ID`, incoming exact-file lint, and lineage pass. I re-verified the pair-approved owner bases: m-10 r27 `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`, m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`, and m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.

The two cancellation cuts, count-once rule, bare-closure-to-UNKNOWN rule, m-3 `phase=cancelled`, four-token local-reject enum, loss/crash split, corrected §B.1/§B.2 citations, and prior F59/counter/push/EOF/replay-custody corrections survive. A fresh whole-document pass nevertheless finds three blockers. These exact bytes cannot receive final pair approval.

## Blocking findings

### R6-F1 — The generic no-stream rule contradicts the owner-real attempt-inert epoch branch

Section 2.2 first includes the DATA-P epoch backstop in the no-stream set and says “each such attempt row is closed by m-8's CTRL-C view alone” (`:93`). Its generic forward-mapping row again includes “DATA-P epoch backstop” and says the worker records a “typed attempt terminal,” with the typed DATA-P/CTRL-C reply as the terminal answer (`:99`).

That is the opposite of the exact owner contract and r6's own adjacent paragraph. M-8 r12 §1.3 says `STALE_EPOCH` and `EPOCH_AHEAD` are attempt-INERT: no `attempt_result`, no stream, no m-8 close, no E0 view from this path; m-10's retirement/epoch machinery owns the committed row's fate. R6 states that correctly at `:94` and in fixture text at `:245`, but the generic rules at `:93/:99` re-terminalize the same branch and invent a CTRL-C close that does not exist.

Required revision: split no-stream **terminal outcomes** (`denied`, `rejected_local`, `cancelled(pre_transport)`) from no-stream **attempt-inert epoch replies**. The epoch branch emits no `attempt_stream_end`, but it also records no m-8-view terminal, expects no `attempt_result`, emits no E0 terminal from this path, and leaves row disposition to m-10 while charging the already-committed row once. Make the main rule and fixture assert those facts without calling the typed epoch reply a terminal close.

### R6-F2 — Gate 2 is not a total comparison over the closed `parked_unknown` lists

The new two-gate sequence correctly handles the safety-critical growth case: equal lists proceed; any Gate-2 identity absent from Gate 1 blocks DATA-P, surfaces the new item, and reassembles (`§2.6:143-147`). But those are not the only possible relations between the two snapshots. M-10 r27 §B.2 says parked rows reach terminal through ordinary owner machinery, so the set can shrink between frames. A member can also retain `{turn_id, tool_call_id, ticket_id}` while another closed field differs. R6 defines neither branch: a removal-only list is unequal but contains no new item, and the identity-only membership rule ignores a changed `state`, `canonical_tool_name`, or `canonical_args_digest`.

Required revision: make Gate 2 total over a precisely defined set comparator. Reject duplicate/malformed identities; compare the full closed member under the identity key; define equal, added/changed, and removed-only relations. Added or changed unsurfaced facts must block DATA-P and reassemble. For removed-only facts, choose and state one honest behavior — reassemble from the Gate-2 snapshot, or proceed with the already-surfaced conservative superset — so every received pair of valid lists has a deterministic branch. Extend the fixture beyond growth to removal and same-identity changed-member cuts.

### R6-F3 — D-5 is simultaneously called closed/contract-real and acknowledged not to match the current m-10 bytes

R6's status says every owner delta is resolved and route-back is closed at m-10 (`:6`), and §2.9 says the terminal/cancellation carriage is “contract-real on both sides” (`:168`). Yet the exact m-9 `turn_terminal` shape has dropped `attempts_summary_ref?`, while the current pair-approved m-10 r27 shape and equivalence predicate still include it (`m-10 r27 §B.2:71`). R6 itself correctly acknowledges that mismatch at `:163`, `:215`, `:219-225`, and `:254-266`.

The live routing advanced during this review: Master has now routed the three-item owner request directly to m-10 in `step3-mvp-design-m10/RECONCILE-orchestrator-planner-20260718-051544.md`. That discharges “awaiting master routing,” but it does not supply m-10's disposition, updated bytes, comparator confirmation, or pair review. The current exact contracts therefore remain unmatched.

Required revision: describe the actual state consistently as **routed, awaiting m-10 owner disposition**; do not call the moved frame shape/comparator contract-real on both sides. Wait for m-10 to dispose the optional-member semantics, confirm or amend the `partial_disposition` comparator and cancellation consumption, and pair-review any new bytes. Then rebind r6 to the final m-10 hash and return one fresh review. If m-10 rejects either offered domain, fold its owner-real alternative rather than preserving the current claim.

## Accepted portions

- R5-F1's cancellation repair is correct: pre-transport is m-8-view-only with no stream end; post-invocation is two-view with `stream_cancelled`; both count one; neither is failure; raw loss never becomes cancellation.
- R5-F2's single m-8 r12 basis, four-token local-reject enum, and explicit `STALE_EPOCH`/`EPOCH_AHEAD` local dispositions are present. The remaining defect is the contradictory generic no-stream aggregation, not the dedicated epoch paragraph.
- R5-F3's growth-path safety action is now mechanical: a newly parked item at Gate 2 blocks DATA-P and forces surfacing/reassembly. The remaining defect is comparison totality.
- R5-F4's m-9-owned shape is materially improved: `attempts_summary_ref?` is absent and `partial_disposition ∈ {none, partials_committed_labeled}` is a closed, deterministic output-axis domain. Its owner consumer has now been correctly routed but has not yet returned.
- M-3 r4 cancellation-phase and m-8 r12 loss/crash folds are consumed faithfully. The prior F59 ordering, canonical counter strings, rediscovery/push boundary, EOF containment, and replay-envelope custody remain nonblockers.

## Revision bar and gate disposition

Return fresh bytes that:

1. separate attempt-inert epoch replies from m-8-view terminal no-stream outcomes at every normative and fixture locus;
2. totalize the Gate-2 `parked_unknown` comparator over equal, added/changed, removed-only, duplicate, and malformed inputs; and
3. consume m-10's final owner disposition at its pair-approved hash, with status/route-back text matching the live state.

This verdict is byte-bound to `1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb`. The closure SITREP, m-10 reciprocal, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `b80afa85ca2e92af0460708123ae9c64103e0d9605665fca79658eb19d2dc99c`.
- Exact reviewed m-9 r6 SHA-256 recomputed: `1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb`.
- Pair-approved m-10 r27 SHA-256 recomputed: `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, three owner approval relays, r6 seam relay, Master r6 route, and Master m-10 owner route exact-file lint: OK.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-051753.md`.
Next requested action: m-9.planner holds r6 and the closure SITREP; folds R6-F1/R6-F2 after m-10 returns its owner disposition and any pair-approved bytes; then rebases the status/contract hashes and returns one fresh uniquely-parented DESIGN request.
