## DESIGN-REVIEW — m-9 lifecycle half r5 full-byte review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — all four findings have bounded owner-correct or local contract-totality resolutions
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260717-210200.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-7.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260717-210652.md
SUBJECT: MUST-REVISE exact r5 452a352d... — the live m-8 cancellation seam is rejected, the promised r6 epoch/reason folds are incomplete, parked_unknown is asserted but not sequenced at the second gate, and D-5's equality-keyed facts lack closed field domains

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r5 bytes at SHA-256 `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`, not only the r5 fold loci. The directly-addressed relay, `DESIGN_DOC_ID`, lineage, and incoming exact-file lint pass. M-7 r11 at `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` and m-10 r21 at `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852` are pair-approved and match r5 at the attach and attempt-open seams.

The earlier F59 ordering, counter encoding, push-custody, replay-envelope custody, and EOF corrections survive. The D-4 §8b non-transitivity reasoning is sound in principle. A fresh whole-document pass nevertheless finds four blocking defects, including a live owner-seam rejection filed after this review request. These exact bytes cannot receive final pair approval.

## Blocking findings

### R5-F1 — The live m-8 owner seam rejects r5's cancellation mapping

R5 maps provider `cancelled` to `stream_cancelled` (`§2.2`), lands `cancelled{partial}` at `turn_cancelled` (`§2.5`), and emits the cancellation/terminal family as a distinct durable fact (`§2.9`). Those m-9 semantics are coherent, but they do not currently compose with the live m-8 owner bytes.

The exact m-8 r7 review filed at `step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-210409.md` is `MUST-REVISE`: m-8 r7 permits cancellation before transport invocation yet maps every cancellation to CTRL-C `transport_failed` and E0 `phase=failed`. That is false for the zero-wire branch and also mislabels an operator cancellation after invocation as a transport failure. The m-8 reviewer requires an honest cancellation disposition and says the selected representation must be reconciled with m-10's closed row enum and m-9's `stream_cancelled`/E0/turn mapping.

Required revision: hold r5 closure until m-8 returns pair-approved replacement bytes. Then consume the owner-real cancellation result, route every moved m-10 row/E0 seam through Master, and pin m-9's zero-wire and wire-crossed cancellation mappings to that final shape. Add both cancellation cuts. A letter-level closure rebind cannot cure a rejected semantic seam.

### R5-F2 — The claimed m-8 r6 triple-confirm fold is incomplete and the document carries mutually stale bases

R5's basis still names m-8 r4 `168c24b7...` with only three local-reject reasons (`:14`); §2.2 likewise says r4, adds the fourth `internal_integrity_fault`, then still calls the m-10 leg “to confirm,” labels it PROPOSED, and promises a fixture for “each of the three reason tokens” (`:94-98`). The §6 fixture still enumerates only those three reasons (`:225`). Elsewhere, the consumed-hash row and fold log say m-8 r6 `ab63f6eb...` (`:242,249`). This is not one byte-bound owner basis.

More importantly, m-9's own r6 triple-confirm at `step3-mvp-design-m8/RECONCILE-planner-20260717-190600.md` promised all three legs would fold into r5. Leg 3 pins m-8 DATA-P epoch-class replies: `STALE_EPOCH` is fencing/no local retry; `EPOCH_AHEAD` is a typed internal fault; neither produces `attempt_result`; parked-row versus no-row budget behavior is explicit. R5 contains no `EPOCH_AHEAD` consumer branch and uses `STALE_EPOCH` only for m-10 attempt-open or ticket/executor paths. It therefore did not fold the promised m-8 DATA-P leg.

Required revision: after m-8's final pair-approved bytes exist, rebase every normative m-8 citation to that one hash; make the local-reject enum and fixtures uniformly four-token; remove the stale PROPOSED/pending/three-token prose; and add explicit m-8 DATA-P `STALE_EPOCH`/`EPOCH_AHEAD` handling, including no-`attempt_result`, no-retry/new-attempt, and parked-row/no-row budget behavior. The replacement bytes require fresh review.

### R5-F3 — The second `parked_unknown` disclosure is not mechanically before successor work

The turn state machine enters ASSEMBLING before attempt-open and issues DATA-P after `attempt_open_ok` (`§2.2`). Section 2.6 then says the worker consumes `parked_unknown` from both `turn_open` and `attempt_open_ok` and surfaces it before successor work. That sentence does not define what happens if the second list contains an item that was absent from the first list. By then the request/context can already be assembled; r5 defines no comparison, delta, DATA-P hold, reassembly, or fail-closed branch. The claim that the new attempt arrives informed is therefore not reachable for a changed second disclosure.

Required revision: sequence both gates explicitly. The `turn_open` list must enter context before ASSEMBLING. On `attempt_open_ok`, compare its closed list against what has already been surfaced; any unsurfaced item must block DATA-P and be surfaced with the request reassembled, or take one explicitly chosen fail-closed disposition, before provider work. Add a fixture where the durable list changes between the two frames and prove no DATA-P or tool work occurs before the new item is surfaced.

### R5-F4 — D-5 calls its frame family exact while leaving the idempotency facts open

Section 2.9 defines `turn_terminal{..., attempts_summary_ref?}` without defining that optional member's type, namespace/target, canonical representation, or absence semantics. It defines `partial_disposition` only by prose — “partials-committed-labeled / in-flight-tool→unknown_effect” — rather than a closed enum or object. M-10 r21 §B.2 makes duplicate handling equivalence-keyed over exactly `{terminal, attempts_summary_ref?}` and `{partial_disposition}`. Without canonical m-9 field domains, m-10 cannot deterministically distinguish an equivalent resend from a conflicting report.

R5 also repeatedly locates D-2/D-5 at m-10 §B.1. `attempt_open_ok` is in §B.1, but m-10's `attach_result` admission gate, D-5 terminal/cancellation consumption, equivalence predicate, and D-4 admission disclosure are in §B.2. The citations at the r5 basis/§2.9/§5/§7 must identify the actual owner loci.

Required revision: either remove `attempts_summary_ref?` for this MVP or close its exact canonical domain and absence/equality semantics; define one exact `partial_disposition` domain and cancellation identity/matching semantics; then obtain m-10 owner confirmation for the comparator that consumes those facts. Correct the §B.1/§B.2 citations throughout.

## Accepted portions

- M-7 r11's three attach results, pinned order, transient-hold branch, and terminal tuple-mismatch branch match r5 §§1.2/1.6.
- M-10 r21's durable `attempt_open_ok`, no-row `attempt_open_reject`, budget rule, and D-2 attach gate match their principal r5 flow.
- D-4's choice of state-only disclosure and its rejection of a worker-forwarded operator citation as authority are correct; only the second-gate sequencing is incomplete.
- The prior F59 consume-before-execute, counter-string, push rediscovery, replay-custody, and EOF fail-closed corrections remain nonblockers.

## Revision bar and gate disposition

Return fresh bytes that:

1. consume a pair-approved m-8 cancellation disposition and one final m-8 hash;
2. fully fold the four-token local-reject and DATA-P epoch-class obligations from the m-8 confirmation;
3. make both `parked_unknown` gates executable before provider/tool work; and
4. close D-5's equality-keyed field domains and obtain the necessary m-10 owner return.

The next review must be fresh and uniquely parented to the replacement bytes. This verdict is byte-bound to `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`. The closure SITREP, m-10 reciprocal, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `22c03701ae8ec330059077f438a80315102fac1958f929aebf3f2a2982c25fba`.
- Exact reviewed m-9 r5 SHA-256 recomputed: `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Pair-approved m-10 r21 SHA-256 recomputed: `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852`.
- Current m-8 r7 SHA-256 recomputed: `b805edab019400d1bd6505dd17beddf1e9b092a05bd1a1b2fbe549cefc721083`; its exact review SHA-256 is `04549d253ff17d40ebc09c4b13a2b386e1b39de19578cecccda53219557b8928` with verdict `must-revise`.
- Incoming DESIGN and live m-8 DESIGN-REVIEW exact-file lint: OK.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260717-210652.md`.
Next requested action: m-9.planner holds r5 and the closure SITREP; after m-8 returns pair-approved cancellation/epoch bytes, fold R5-F1 through R5-F4, route the moved D-5/cancellation consumer seams to m-10 through Master, and return a fresh uniquely-parented byte-bound DESIGN request.
