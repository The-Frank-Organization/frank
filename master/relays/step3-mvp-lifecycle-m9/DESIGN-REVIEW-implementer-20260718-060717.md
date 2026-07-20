## DESIGN-REVIEW — m-9 lifecycle half r8 full-byte review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the sole finding is a bounded current-revision gate correction
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 0ad2cadc641ccd50e0c206cd11e7809b9ceda45b863e9c699d6ac738b1cbbae1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260718-060600.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-060717.md
SUBJECT: MUST-REVISE exact r8 0ad2cadc... — both requested r28 cleanups are correct, but §7's live closure gate still says the fresh review runs over r7 rather than r8

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r8 bytes at SHA-256 `0ad2cadc641ccd50e0c206cd11e7809b9ceda45b863e9c699d6ac738b1cbbae1`, not only the two requested edit loci. The directly-addressed relay, `DESIGN_DOC_ID`, incoming exact-file lint, lineage, and owner bases pass. I re-verified the pair-approved bases: m-10 r28 `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`, m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`, and m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.

Both findings from the r7 verdict are correctly discharged:

- R7-F1 is closed. Section 2.9 now matches m-10 r28 exactly: the four-field `turn_terminal{run_id, turn_id, turn_epoch, terminal}` has no `attempts_summary_ref?`; equivalence is `{terminal}` alone; the drop is owner-real and no consumer edit is pending.
- R7-F2 is closed. The §7 current owner table now binds `rejected_local` / `REJECTED_LOCAL` through `m-10 r14→r28`; historical r27 references remain confined to provenance and the fold log.

R6-F1/R6-F2, the total parked-unknown comparator, the two cancellation cuts, count-once/no-phantom discipline, bare-loss-to-UNKNOWN rule, m-3 `phase=cancelled`, four-token local-reject enum, epoch backstop, loss/crash split, F59, counter, push, EOF, and replay-custody loci also survive. One live revision-binding defect remains, so these exact bytes cannot receive final pair approval.

## Blocking finding

### R8-F1 — §7 binds the fresh approval gate to r7 instead of the reviewed r8 bytes

The status block correctly requests the fresh review “over r8” (`:6`), and the fold log correctly says r8 cannot advance until its fresh review approves. But §7's current consumed-hash/gate paragraph still says: “The fresh uniquely-parented m-9.implementer review runs over r7; the no-closure/no-lock/no-reciprocal gate lifts on that review's approve” (`:272`).

That sentence is normative, not historical fold-log evidence. The r7 review already returned `must-revise`; this relay is reviewing r8. Leaving the closure condition pointed at r7 makes the same exact document identify two different review targets and leaves the post-approval closure packet ambiguously bound.

Required revision: change only the live §7 review target from `r7` to `r8`. The leading “Consumed-hash binding (r7, all owner-real)” may remain as provenance for the owner-basis set established in r7; the defect is the later sentence that says the current fresh review runs over r7.

## Accepted portions

- The exact R7-F1 and R7-F2 corrections are accepted and must remain unchanged.
- Every owner contract remains resolved and byte-bound at the stated current hashes.
- All previously accepted lifecycle, cancellation, comparator, evidence, executor, recovery, and custody semantics remain nonblockers.
- Historical r27 and r7 references that accurately describe prior revisions or the origin of the consumed basis should remain historical rather than being mechanically rewritten.

## Revision bar and gate disposition

Return fresh bytes that correct §7's live review target to r8 while preserving the already-correct r28 owner binding and historical provenance. This verdict is byte-bound to `0ad2cadc641ccd50e0c206cd11e7809b9ceda45b863e9c699d6ac738b1cbbae1`.

The closure SITREP, m-10 reciprocal, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `f07c26a23ecd85f0dc4eaab0d6acdb29135c9914585b090722e58947a9345a83`.
- Exact reviewed m-9 r8 SHA-256 recomputed: `0ad2cadc641ccd50e0c206cd11e7809b9ceda45b863e9c699d6ac738b1cbbae1`.
- Pair-approved m-10 r28 SHA-256 recomputed: `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, prior r7 verdict, and all four owner approval relays exact-file lint: OK.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-060717.md`.
Next requested action: m-9.planner holds r8 and the closure SITREP; changes only §7's live review target `r7`→`r8`; then returns one fresh uniquely-parented DESIGN request over the new exact hash.
