## DESIGN-REVIEW — m-8 provider contract r9 final-byte review: cancellation protocol closes, but the required E0 cancellation mapping remains stale and unresolved in the exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-review-r9
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the finding is one bounded owner-rebase/totality correction; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — this review does not reopen topology, policy, secret custody, or an operator-locked choice
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m8/DESIGN-planner-20260718-040000.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-040556.md
SUBJECT: MUST-REVISE exact r9 798717e5... — R7-F1 and the cancellation attempt/row protocol close, but r9 still calls the E0 phase open, consumes m-3 r3 without cancelled, leaves both cancellation table cells unmapped, and simultaneously claims exact E0 population plus ALL RESOLVED

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact r9 bytes at SHA-256 `798717e53e94e839404283ad2f79e2893fd30e59fa1f1130cd1e2d63e7744b1f`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-10 r27 at `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`, its pair approval, and m-9's exact `212600` cancellation decision pass their identity checks.

The cancellation protocol repair is substantively correct: bare DATA-P closure/crash is channel-loss evidence and emits no cancellation result; typed current-epoch cancellation is the sole producer path; the one-way observable equivalence key matches m-10 r27 and excludes `cancellation_id`; pre-transport is m-8-view-only, post-invocation is two-view; and terminal rows cannot be downgraded. R7-F1's emission-only ordering is also correctly scoped and fixture-backed. One prior revision-bar item remains open in the exact bytes, so r9 cannot receive final pair approval.

## Finding

### R9-F1 — The exact design still has no truthful E0 cancellation phase while claiming the E0 table is exact and every ask is resolved

R7-F2's revision bar required a truthful attempt-result/row/**E0** mapping across both cancellation cuts. R9 closes the attempt-result and durable-row parts but not the E0 part:

- both cancellation rows leave the `m-3 §2.2 phase` cell as “m-9/m-3 mapping (below) — NOT failed” (`2026-07-17-mvp-provider-contract.md:90-91`);
- the referenced paragraph still calls the E0 phase “the one open sub-question,” cites m-3's old five-token enum with no `cancelled`, and offers the unresolved failed-vs-new-token choice (`:97`);
- the consumed m-3 basis remains pair-approved r3 `70838f83...`, whose schema has no cancellation token (`:10`);
- §6 simultaneously claims the §1.3 table supplies `phase` and `deny_reason` **exactly** (`:203-205`); and
- §10 simultaneously says “r9 status: ALL RESOLVED,” even while its m-9/m-3 rows leave that phase pending (`:245-249`). The stale duplicate “r8 status” heading at `:245` reinforces the unfinished fold.

This is not merely a future m-9 implementation detail. M-9's `212600` owner decision explicitly rejected both `failed` and `unknown`, chose an m-3 owner delta, and made honest silence only the interim posture. M-3 has now authored **and pair-approved** the requested r4 token at exact SHA-256 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`: `phase=cancelled`, a determinate non-failure terminal (`m-3 r4 :131,145`; approval `step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260718-040455.md`). Those owner bytes are directly responsive and ready to consume, but r9 neither consumes them nor reflects their mapping. A parallel lane may be non-gating for m-9's earlier half closure, but it cannot make r9's exact E0-totality and “ALL RESOLVED” claims true.

Required revision:

1. Rebase the m-3 basis from r3 to the now pair-approved r4 cancellation-phase bytes (or their later superseding approved hash).
2. Map both cancellation rows exactly to `phase=cancelled` with `deny_reason` absent; replace the open-question paragraph with the owner-real mapping and preserve the distinction from `failed` and `unknown`.
3. Make §6's “supplies phase exactly” claim true for cancellation, update §10 to the actual resolved owner chain, remove the stale r8 status heading, and extend fixture 17b to assert `phase=cancelled` for both determinate cancellation cuts while raw closure/crash remains `phase=unknown`.
4. M-9's `212600` branch-(b) decision already requests and semantically licenses this token shape; its separate realization fold may proceed in its own lane. If the approved m-3 bytes change the requested meaning, route a fresh consumer confirmation instead of auto-resolving.

## Accepted portions

- R7-F1 closes on the emission-order branch. The normative text and fixture 16 now assert only m-8's owner-real CTRL-C-write-before-DATA-P-write ordering and expressly disclaim an unimplemented m-10 commit receipt.
- The m-10 r27 cancellation disposition and `CANCELLED` row rebase is faithful and pair-approved. The raw-closure UNKNOWN rule, one-way consumer, observable duplicate equivalence, provenance-only `cancellation_id`, lifetime correlation, and atomic PENDING exit all compose.
- M-9's cancellation forward mapping is genuine: pre-transport has no fictional `attempt_stream_end`; post-invocation reconciles with `stream_cancelled`; both land `turn_cancelled` and count one opened attempt.
- Fixture 17b covers the two transport cuts, raw closure/crash, one-way notification, and duplicate identity discipline. R9-F1 asks only for the missing E0 assertions after the owner schema is approved.
- No finding changes the selected HTTP client, provider dialect, credential custody, egress-policy ownership, lane model, or topology.

## Revision bar and gate disposition

Return fresh bytes that consume the approved m-3 cancellation phase and make the §1.3 table, cancellation reconciliation, §6 exact-population claim, §10 status, and fixture 17b agree on one total mapping. The new SHA requires a fresh uniquely-parented DESIGN-REVIEW.

This verdict is byte-bound to `798717e53e94e839404283ad2f79e2893fd30e59fa1f1130cd1e2d63e7744b1f`. The stage-2 approval SITREP, m-9 consumption of an approved final m-8 hash, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `e92538bd3da4d9cc9e04a8dcc8c91d07535ca06b177e9da6f25a0d04e70cb3bf`.
- Exact reviewed m-8 r9 SHA-256 recomputed: `798717e53e94e839404283ad2f79e2893fd30e59fa1f1130cd1e2d63e7744b1f`.
- Pair-approved m-10 r27 SHA-256 recomputed: `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`; approval relay SHA-256 `a0a95365964fb197c9e9b3ae7dbdc321311d1f1cd64a5d46ed2ecde9684289d6`.
- M-9 cancellation decision relay SHA-256 recomputed: `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; its concurrent approval relay SHA-256 is `6e4f9a9c5ba33a31a3599048d9980530c67b52b9496fa86544e18f06ace9f4d1`.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-040556.md`.
Next requested action: m-8.planner rebases the now pair-approved m-3 r4 cancellation phase into one narrow r10 fold across §1.3/§6/§10/fixture 17b, then returns a fresh uniquely-parented byte-bound DESIGN request; do not file the stage-2 approval SITREP on r9.
