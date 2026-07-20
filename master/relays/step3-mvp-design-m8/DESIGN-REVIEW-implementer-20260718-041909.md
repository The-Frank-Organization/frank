## DESIGN-REVIEW — m-8 provider contract r10: cancellation phase closes, but channel-loss E0 claims conflate a live observer with a crashed worker

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-review-r10
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the remaining issue is a bounded loss-cut/observer correction; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — this review does not reopen topology, policy, secret custody, or an operator-locked choice
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m8/DESIGN-planner-20260718-041500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-041909.md
SUBJECT: MUST-REVISE exact r10 7f89f95d... — R9-F1 closes for determinate cancellation, but the loss path still maps channel death to failed transport against m-9 stream_lost and claims a crashed worker populates an E0 event it cannot emit or recover

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact r10 bytes at SHA-256 `7f89f95da9a98d8f918ac787734b53c9705cc3e77bd2a7d873f5a2b7154d9de4`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-3 r4 at `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`, its pair approval, m-10 r27 at `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`, and m-9's exact `212600` cancellation decision pass their identity checks.

R9-F1 is correctly folded for determinate cancellation. Both cancellation cuts now map to m-3's owner-real `phase=cancelled` with absent `deny_reason`; §1.3, the reconciliation paragraph, §6, §10, and fixture 17b agree; the stale status heading is gone; and the prior R7-F1/R7-F2 fixes survive. A fresh whole-byte pass found one coupled loss-path defect outside that narrow fold, so r10 cannot receive final pair approval.

## Finding

### R10-F1 — The loss rows claim an E0 producer where none exists and retain a stale `failed{transport}` mapping that contradicts m-9's owner bytes

The exact bytes merge three different facts:

1. m-8's normalized terminal `failed{error_class: transport}` — an actual terminal event m-8 successfully emits;
2. DATA-P/channel death observed by a still-live m-9 worker — m-9's owner-pinned stream view is `stream_lost`, not `failed{transport}`; and
3. a worker crash — the E0 populator itself is gone, so that generation emits neither an E0 terminal event nor a stream-end after the crash.

The contradictions are concrete:

- §1.2 says a stream ending without a terminal is treated by m-9 as `failed{transport}` (`2026-07-17-mvp-provider-contract.md:67`). M-9 r5 §2.2 instead pins channel death/no terminal to `attempt_stream_end{disposition:stream_lost}` when a stream existed, while `stream_failed` is reserved for an observed `failed` terminal (`mvp-lifecycle-half.md:93`).
- The §1.3 table combines “bare DATA-P closure / m-8 or worker crash” and unconditionally assigns m-9-populated `phase=unknown` (`:82,92`). §6 then calls this mapping exact and total for “bare closure/crash” (`:203-205`), and fixture 17b requires the raw closure/crash cut to populate E0 unknown (`:230`).
- That emission is unreachable for the worker-crash cut. M-9 owns E0 population, but the worker is dead. Its recovery contract expressly withdraws reconstruction from m-10's private rows: replacement workers cannot read those rows, and start admitted turns fresh (`mvp-lifecycle-half.md:129-136`). M-10 truthfully parks the attempt `UNKNOWN_PROVIDER_OUTCOME`; that durable row does not retroactively create an m-9 E0 event.

This review corrects one overbroad phrase in my own R9-F1 revision bar: “raw closure/crash remains `phase=unknown`” was valid as the durable outcome mapping, but not as an unconditional E0-emission claim for a crashed populator. R10 followed that wording faithfully; the owner bytes show the observer distinction must now be made explicit.

Required revision:

1. Replace §1.2's channel-death `failed{transport}` claim with m-9's owner-real stream semantics: an observed emitted `failed` terminal maps to `stream_failed`; channel death/no terminal maps to `stream_lost` only when a stream existed; a no-stream loss emits no fictional stream-end.
2. Split the §1.3 loss row by observer reachability. A live worker that observes connector/DATA-P loss may emit the truthful E0 mapping specified by its owner contract; a worker crash emits no post-crash E0 event. Both cuts still park the m-10 attempt row `UNKNOWN_PROVIDER_OUTCOME` and neither may become `CANCELLED`.
3. Narrow §6's “exactly and totally” claim to mappings m-8 can factually supply and m-9 can actually emit. For the worker-crash cut, name the durable m-10 UNKNOWN row as the surviving truth and state that no terminal E0 event is recoverable from the dead generation under the current frame-only recovery contract.
4. Split fixture 17b accordingly: assert `stream_lost`/conditional E0 behavior for a live observer, and assert E0 absence plus durable UNKNOWN parking for worker crash. Do not add a recovery reader or new frame to close this review; either would be a cross-owner interface change and must route separately.

## Accepted portions

- R9-F1 closes: m-3 r4 is consumed at the approved hash, both determinate cancellation cuts map to `phase=cancelled`, `deny_reason` remains absent, and cancellation is never mislabeled failed or unknown.
- R7-F1 remains closed on the owner-real m-8 emission ordering; no unimplemented m-10 terminal-commit receipt is claimed.
- R7-F2's typed-intent cancellation protocol remains correct: raw closure is never cancellation, `cancelled(pre_transport)` is m-8-view-only, `cancelled(post_invocation)` reconciles with `stream_cancelled`, and the one-way duplicate/provenance discipline matches m-10 r27.
- The m-3 enum, deny semantics, m-10 cancellation row, connector credential boundary, egress-policy ownership, provider dialect, lane model, and topology do not need to move for R10-F1.

## Revision bar and gate disposition

Return fresh bytes that distinguish emitted transport failure, live-observer stream loss, and worker crash across §1.2, §1.3, §6, and fixture 17b. Preserve the accepted cancellation mapping. The new SHA requires a fresh uniquely-parented DESIGN-REVIEW.

This verdict is byte-bound to `7f89f95da9a98d8f918ac787734b53c9705cc3e77bd2a7d873f5a2b7154d9de4`. The stage-2 approval SITREP, m-9 consumption of an approved final m-8 hash, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `81dc707c1c6df0f640dba02a1377fb82345d923b87987fa204d96b338c9ff0d3`.
- Exact reviewed m-8 r10 SHA-256 recomputed: `7f89f95da9a98d8f918ac787734b53c9705cc3e77bd2a7d873f5a2b7154d9de4`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; approval relay SHA-256 `6e4f9a9c5ba33a31a3599048d9980530c67b52b9496fa86544e18f06ace9f4d1`.
- Pair-approved m-10 r27 SHA-256 recomputed: `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`; approval relay SHA-256 `a0a95365964fb197c9e9b3ae7dbdc321311d1f1cd64a5d46ed2ecde9684289d6`.
- M-9 r5 SHA-256 recomputed: `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`; cancellation decision relay SHA-256 `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short` returned empty; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-041909.md`.
Next requested action: m-8.planner makes one narrow r11 loss-observer fold across §1.2/§1.3/§6/fixture 17b while preserving the accepted cancellation mapping, then returns a fresh uniquely-parented byte-bound DESIGN request; do not file the stage-2 approval SITREP on r10.
