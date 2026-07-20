## DESIGN-REVIEW — APPROVE m-10 r28 exact bytes: `{terminal}`-only D-5 equivalence remains total; partial/cancellation consumer seams confirmed

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r29
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — no open m-10 pair-review finding remains on these exact bytes; later cross-lane and operator gates remain separate
GRILL_REQUIRED: no — this approval does not replace the stage-5 control-plane grill
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-053000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-053100.md
SUBJECT: APPROVE exact r28 4ffaa9ec... — the batched m-9 comparator/cancellation return and m-3 r4 letter rebind are clean; approval advances only to the Master-routed sweep and m-9 r7 closure path

DESIGN_REVIEW_VERDICT: approve

m-10.planner — I approve the exact r28 design bytes at SHA-256 `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-9 r6 at `1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb`, and m-3 r4 at `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` pass their identity checks.

## Approval basis

- **Item (1) closes:** m-10 now consumes m-9's exact `turn_terminal{run_id, turn_id, turn_epoch, terminal}` request shape. Removing undefined `attempts_summary_ref?` and narrowing same-type equivalence to `{terminal}` does not under-distinguish any D-5 transition-table row.
- **H-14 totality closes:** same committed `terminal` is the equivalent-resend branch; a different committed `terminal` is the same-type conflict branch; cancellation composition uses the terminal value (`turn_cancelled`) plus the separately committed cancel-ack fact. Malformed, unknown-turn, stale-epoch, and fresh-legality branches do not consume the removed member. The ordered table therefore remains total across both request families.
- **Item (2) closes:** `turn_cancel_ack` equivalence consumes exactly `{partial_disposition}` over the closed turn-output enum `{none, partials_committed_labeled}`. Tool-effect state remains on `tool_calls` plus the D-4 disclosure, while `cancel_point` remains attempt-row provenance; neither axis is silently collapsed into the comparator.
- **Item (3) closes:** m-9's two cancellation cuts, count-once rule, bare-closure-to-UNKNOWN rule, and loss-not-cancel rule compose with m-10 §B.1. The worker neither carries nor compares the m-10-owned durable `cancellation_id`; inbound duplicate equivalence remains limited to observable frame facts.
- **The m-3 letter rebind is exact:** m-3 r4 recomputes to `009df607...`; its `phase=cancelled` and canonical string `turn_epoch` surfaces remain compatible with the r28 consumer basis. The remaining m-7/m-3/m-8 sweep is still Master-routed work, not silently closed here.
- The pair-approved r27 surfaces outside this bounded two-locus amendment remain the accepted basis of the current document; no finding in this review reopens them.

## Scope and remaining gates

This is pair approval of the exact m-10 r28 design document at `4ffaa9ec...`. It authorizes the m-10 planner to return the byte-bound approval to Master for the remaining owner sweep and the m-9 closure sequence.

m-9 r6 is presently MUST-REVISE for its own R6-F1/R6-F2 local findings and the now-resolved R6-F3 owner dependency. m-9 must fold one r7 over the returned m-10 r28 basis and receive a fresh m-9.implementer approval before its closure SITREP and m-10 reciprocal can close stage 3. This approval does not waive or pre-answer that review.

Master+VP interface lock, the remaining sweep, the stage-5 grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any byte change to the approved design invalidates this approval and requires a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `8c616311aed4657973ddc635123ad745f32577bfad5fe9f8afbfc12d1a2c0ebc`.
- Exact m-10 r28 SHA-256 recomputed: `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`.
- Exact m-9 r6 SHA-256 recomputed: `1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb`.
- Exact m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Incoming DESIGN exact-file lint: OK.
- Targeted final sweep: m-10 §B.1 cancellation-result consumer and §B.2 D-5 request/equivalence/transition table; m-9 r6 §2.5 and §2.9; m-3 r4 §2.2 cancellation and counter encoding.

ACTIONS_GIT_REF: docs-workspace action only — created this approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-053100.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner returns this exact-byte approval to Master; Master routes the remaining m-7/m-3/m-8 sweep and m-9 folds/reviews r7 before any stage-3 closure claim. Do not treat this approval as interface lock, PLAN, T4, implementation, merge, or deploy authority.
