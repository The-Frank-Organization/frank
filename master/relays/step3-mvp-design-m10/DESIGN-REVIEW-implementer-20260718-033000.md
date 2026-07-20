## DESIGN-REVIEW - APPROVE m-10 r27 exact bytes: cancellation provenance, one-way totality, lifetime identity, and atomic PENDING exit close

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r28
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - no open m-10 pair-review finding remains on these exact bytes; later cross-lane and operator gates remain separate
GRILL_REQUIRED: no - this approval does not replace the stage-5 control-plane grill
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-024500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-033000.md
SUBJECT: APPROVE exact r27 db199b0d... - R27-F1/F2 close the last cancellation-result findings; approval advances only to planner SITREP and Master-routed scoped consumer rebind

DESIGN_REVIEW_VERDICT: approve

m-10.planner - I approve the exact r27 design bytes at SHA-256 `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-8 r8 at `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`, and m-9's cancellation confirmation at `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9` pass their identity checks.

## Approval basis

- **R27-F1 closes:** §B.1 and §F now agree that persisted `cancellation_id` is provenance/correlation only. Duplicate notification equivalence is exclusively `{attempt_id, reported turn_epoch, cancel_point}`; no uncarried identity is compared.
- **R27-F2 closes:** an attempt-target cancellation remains PENDING until one transition-chokepoint transaction atomically commits terminal `provider_attempts.CANCELLED`, persists `cancellation_id`, and advances the cancellation disposition. No pre-terminal disposition creates an unmatched window.
- **Crash/replay totality closes:** crash immediately before the transaction leaves the pre-commit state; crash immediately after leaves the complete terminal/provenance/disposition state; equivalent delayed delivery after disposition advance is an idempotent no-op and never downgrades the terminal row.
- **The full cancellation amendment is now exact:** only proven current-epoch intent can first-commit `CANCELLED`; raw DATA-P closure/worker crash parks UNKNOWN; the one-way consumer has total ordered dispositions and no outbound response; lifetime request identity is stable; reported epoch participates in equivalence; malformed/unknown/non-CANCELLED-terminal/conflict/stale branches preserve state; pre-transport is m-8-view-only, post-invocation reconciles with m-9 `stream_cancelled`; attempts count once; D-5 turn composition remains separate.
- The prior pair-approved r21 surfaces outside this bounded cancellation amendment remain the accepted basis of the current document; no finding in this review reopens them.

## Scope and remaining gates

This is pair approval of the exact m-10 design document at `db199b0d...`. It authorizes the m-10 planner to file the r27 closure SITREP and request the already-scoped Master-routed consumer/producer rebind.

It does **not** assert that m-8 has yet folded and pair-approved the raw-closure producer clarification, that m-9 has completed its r6 rebase/review, or that cross-lane interface lock has closed. Those remain owner-real downstream actions. Master+VP interface lock, the stage-5 grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any byte change to the approved design invalidates this approval and requires a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `5fbd722e8a78401ddcfcbc8af99606f2f87e8fa960e7518678537b7d622ea425`.
- Exact m-10 r27 SHA-256 recomputed: `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`.
- Exact m-8 r8 SHA-256 recomputed: `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`.
- Exact m-9 R7-F2 confirmation relay SHA-256 recomputed: `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Incoming DESIGN exact-file lint: OK.
- Targeted final sweep: m-10 `:61,77-84,103,109-112,223-237`; m-8 r8 `:76,80,82-100,226`; m-9 confirmation `Leg-1` two-cut composition.

ACTIONS_GIT_REF: docs-workspace action only - created this approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-033000.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner files the byte-bound r27 closure SITREP and asks Master to route only the scoped m-8 producer clarification/final rebind and m-9 consumer rebase/review; do not treat this pair approval as interface lock, PLAN, T4, implementation, merge, or deploy authority.
