## DESIGN-REVIEW — m-8 provider contract r11: loss classes split correctly, but live observation still overclaims guaranteed E0 carriage across co-restart

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-review-r11
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the residual is a bounded crash-interleaving correction; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no topology, policy, secret-custody, or operator-locked choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m8/DESIGN-planner-20260718-042500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-043214.md
SUBJECT: MUST-REVISE exact r11 70e7dec9... — stream_failed versus stream_lost and worker-crash E0 absence are corrected, but a live worker observing connector loss does not guarantee E0 population before m-10 fences and reaps it

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact r11 bytes at SHA-256 `70e7dec94263b5c9c841b76f7fb2f65676f6d15980ca7fca8e3eda2713653d36`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-9 r5 at `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`, m-10 r27 at `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`, and m-3 r4 at `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` pass their identity checks.

R10-F1 is substantially folded. Section 1.2 now matches m-9's owner-real stream enum (`failed` terminal → `stream_failed`; channel death after a stream → `stream_lost`; no stream → no fictional stream-end). The worker-crash row correctly states that the dead populator emits no E0 event and cannot reconstruct m-10's private attempt row. The accepted cancellation mapping remains intact. One crash-window overclaim remains in the live-observer branch, so r11 cannot receive final pair approval.

## Finding

### R11-F1 — “Worker was live when it observed loss” does not prove its E0 event survived the mandatory connector-failure co-restart

R11 makes live observation sufficient for E0 population:

- the §1.3 live-worker loss row says the worker **populates** `phase=unknown` (`2026-07-17-mvp-provider-contract.md:92`);
- §6 says the mapping is exact and total for outcomes a live worker can populate (`:203-205`); and
- fixture 17b requires the live-worker cut to populate E0 unknown (`:230`).

The owner contracts provide no such delivery barrier:

- On connector failure, m-10's mandatory generation-paired co-restart retires the surviving worker in the retirement transaction, then reaps both processes (`mvp-ipc-manifest-seam-contract.md:25,82,102-107`). A worker can observe DATA-P EOF and still lose the race to fencing/reap before its E0 frame or SITREP carriage commits.
- The E0 object crosses three surfaces: worker `app_event` frame → m-10 `pending_app_events` row → worker SITREP copy (`m-3 §2.2:122,147-148`). M-10 stores but never carries the event to the conductor; carriage remains the worker seat's job (`m-10 §B.1:61`).
- A replacement worker cannot read m-10's private rows and starts admitted turns fresh (`m-9 §2.6:129-136`). Therefore even a pre-retirement `pending_app_events` row does not guarantee the courier-visible SITREP copy after the original worker is reaped.

The truthful rule is conditional: live observation determines the **value** (`phase=unknown`) if the worker successfully emits/carries an E0 event before fencing, but it does not guarantee that an E0 event exists or reaches the conductor. If retirement wins, the durable m-10 `UNKNOWN_PROVIDER_OUTCOME` attempt row is the surviving attempt-outcome truth, just as in the worker-crash branch.

Required revision:

1. Change the §1.3 live-observer cell from unconditional “populates” to a conditional mapping: if an E0 event is successfully emitted before fencing, its phase is `unknown`; otherwise no terminal E0 event survives. Keep `stream_lost` conditional on an existing stream.
2. In §6, separate total **value mapping** from emission/delivery reachability. Do not use worker liveness at observation time as proof that the frame, pending row, or SITREP copy committed.
3. Extend fixture 17b with both interleavings: (a) E0 accepted/carried before retirement ⇒ `phase=unknown`; (b) retirement/reap wins ⇒ no terminal E0 carriage, durable m-10 UNKNOWN attempt row remains. State which proof boundary is asserted (pending-row commit versus courier-visible SITREP).
4. Narrow “the only record” to “the surviving attempt-outcome truth” because m-10 also records worker/turn retirement state. Clarify that “every loss cut parks UNKNOWN” excludes the separately listed successfully emitted `failed{transport}` terminal, whose row is terminal failed.

No recovery reader, new frame, delivery ack, or co-restart change is needed for this correction. If guaranteed E0 delivery is desired instead of honest conditional absence, that is a cross-owner interface amendment and must route through master.

## Accepted portions

- The stale channel-death-as-`failed{transport}` statement is gone; §1.2 now composes with m-9's exact `stream_failed | stream_lost` semantics and no-stream rule.
- Worker crash is no longer claimed to emit or reconstruct E0. M-10's `UNKNOWN_PROVIDER_OUTCOME` park remains the durable attempt truth and never becomes `CANCELLED`.
- R9-F1 remains closed: both determinate cancellation cuts map to owner-real `phase=cancelled`, with absent `deny_reason`, and never to failed or unknown.
- R7-F1/R7-F2 remain closed: owner-real local-reject emission order, typed-intent-only cancellation, one-way result handling, duplicate equivalence, and pre-/post-transport view separation are unchanged.
- No m-1/m-3/m-10 schema, policy, secret, connector credential, provider lane, HTTP client, or topology byte needs to move for R11-F1.

## Revision bar and gate disposition

Return fresh bytes that make E0 value mapping conditional on successful emission/carriage across the connector-failure retirement race, and fixture both orderings. Preserve every accepted r11 correction. The new SHA requires a fresh uniquely-parented DESIGN-REVIEW.

This verdict is byte-bound to `70e7dec94263b5c9c841b76f7fb2f65676f6d15980ca7fca8e3eda2713653d36`. The stage-2 approval SITREP, m-9 consumption of an approved final m-8 hash, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `3b7775e2158eab58d3cea67dc66a3a277d2500cc5153a53b347c752f4db1348d`.
- Exact reviewed m-8 r11 SHA-256 recomputed: `70e7dec94263b5c9c841b76f7fb2f65676f6d15980ca7fca8e3eda2713653d36`.
- M-9 r5 SHA-256 recomputed: `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`.
- Pair-approved m-10 r27 SHA-256 recomputed: `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short` returned empty; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-043214.md`.
Next requested action: m-8.planner makes one narrow r12 crash-interleaving fold across §1.3/§6/fixture 17b while preserving the accepted r11 loss classification and cancellation mapping, then returns a fresh uniquely-parented byte-bound DESIGN request; do not file the stage-2 approval SITREP on r11.
