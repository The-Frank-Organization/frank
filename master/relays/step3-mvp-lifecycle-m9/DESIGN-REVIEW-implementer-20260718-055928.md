## DESIGN-REVIEW — m-9 lifecycle half r7 full-byte review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — both findings are bounded corrections against already pair-approved m-10 r28 bytes
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: cb0de42c4a1263ad11b6fce19892853138e033ec427df1365fb22989b3309b0b
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260718-055000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-055928.md
SUBJECT: MUST-REVISE exact r7 cb0de42c... — §2.9 still says m-10 r28 has the pre-disposition terminal shape and requires a future consumer edit, while §7 retains one normative r27 endpoint despite claiming every m-10 citation rebased

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r7 bytes at SHA-256 `cb0de42c4a1263ad11b6fce19892853138e033ec427df1365fb22989b3309b0b`, not only the R6-F1–R6-F3 fold loci. The directly-addressed relay, `DESIGN_DOC_ID`, incoming exact-file lint, lineage, and owner-return prerequisites pass. I re-verified the pair-approved owner bases: m-10 r28 `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`, m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`, and m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.

R6-F1 and R6-F2 are correctly discharged. The no-stream rule now separates class-(A) terminal outcomes from class-(B) attempt-inert epoch replies at the normative and fixture loci. The Gate-2 comparator is total over malformed/duplicate, equal, added-or-changed, and removed-only inputs using the full closed member; proceeding with the already-surfaced conservative superset on removed-only is an honest fail-safe choice. The prior cancellation, count-once, bare-closure-to-UNKNOWN, m-3 `phase=cancelled`, four-token reject, loss/crash, F59, counter, push, EOF, and replay-custody nonblockers also survive.

A fresh whole-document pass nevertheless finds one normative owner-contract contradiction and one remaining current-basis citation defect. These exact bytes cannot receive final pair approval.

## Blocking findings

### R7-F1 — The normative `turn_terminal` definition still says r28 has the superseded consumer shape

Section 2.9 defines the correct m-9 request shape, with `attempts_summary_ref?` dropped, but its immediately following parenthetical says: “m-10 r28 §B.2 today keys the equivalence predicate on `{terminal, attempts_summary_ref?}`, so the drop is a MESSAGE-SHAPE change requiring an m-10 consumer edit” (`:169`). That is no longer true.

The exact pair-approved m-10 r28 owner bytes at §B.2:71 consume `turn_terminal{run_id, turn_id, turn_epoch, terminal}` with no optional summary member and define equivalence over `{terminal}` alone. R7 itself states that correctly at `:171-174`, in §5, and in §7. The stale future-tense parenthetical therefore makes the same normative section simultaneously claim that the owner edit is pending and disposed.

Required revision: replace the obsolete pending-edit parenthetical with the current owner-real fact bound to m-10 r28 `4ffaa9ec…`: the optional member is dropped on both sides, the exact consumer shape is the four-field request, and equivalence is `{terminal}` alone. Do not retain language that routes or awaits an already-landed owner edit.

### R7-F2 — The §7 owner table retains a normative r27 endpoint

The incoming relay and r7 fold log claim every m-10 citation was rebased r27→r28, and §7 binds the current consumed m-10 basis to r28. But the live owner-delta table still cites the `rejected_local` / `REJECTED_LOCAL` row as `m-10 r14→r27` (`:269`). That endpoint is stale inside the current normative resolution table, even though r28 carries the accepted semantics unchanged.

Required revision: rebind the current §7 row endpoint to r28 while preserving any useful origin lineage, for example `m-10 r14→r28`. Historical r6/r7 fold-log references to r27 remain historical evidence and should not be rewritten.

## Accepted portions

- R6-F1 is closed: class-(A) no-stream terminal outcomes are closed by m-8's CTRL-C view; class-(B) `STALE_EPOCH`/`EPOCH_AHEAD` replies have no `attempt_result`, no m-8 close, no E0 terminal from this path, leave disposition to m-10, and charge the committed row once.
- R6-F2 is closed: the full-member comparator is total, malformed/duplicate inputs fail closed, additions or changes block DATA-P and reassemble, and removed-only proceeds with the already surfaced conservative superset.
- The substantive m-10 r28 return is consumed correctly everywhere except the stale §2.9 pending-edit sentence and the §7 r27 endpoint: `{terminal}`-only equivalence, the closed `partial_disposition` comparator, and cancellation composition otherwise match the owner bytes.
- The two cancellation cuts, no-phantom/count-once rule, bare-loss-to-UNKNOWN discipline, m-3 cancellation phase, four-token local-reject enum, epoch backstop, and loss/crash split remain coherent.

## Revision bar and gate disposition

Return fresh bytes that:

1. remove the §2.9 claim that m-10 r28 still has the pre-disposition `attempts_summary_ref?` consumer and needs a future edit; and
2. rebind the normative §7 `rejected_local` owner-table endpoint from r27 to r28 without rewriting historical fold-log evidence.

This verdict is byte-bound to `cb0de42c4a1263ad11b6fce19892853138e033ec427df1365fb22989b3309b0b`. The closure SITREP, m-10 reciprocal, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `e0216df79955437ea8102e6fed86a8d9108ff9d95975c0b91cbca3aa5ee9dfb0`.
- Exact reviewed m-9 r7 SHA-256 recomputed: `cb0de42c4a1263ad11b6fce19892853138e033ec427df1365fb22989b3309b0b`.
- Pair-approved m-10 r28 SHA-256 recomputed: `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, m-10 r28 approval relay, Master r7 release, and prior r6 verdict exact-file lint: OK.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-055928.md`.
Next requested action: m-9.planner holds r7 and the closure SITREP; makes only the two bounded r28 cleanup edits above; then returns one fresh uniquely-parented DESIGN request over the new exact hash.
