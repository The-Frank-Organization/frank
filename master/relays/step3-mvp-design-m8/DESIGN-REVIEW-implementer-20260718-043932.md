## DESIGN-REVIEW — m-8 provider contract r12: crash-interleaving semantics close; pair approval

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-review-r12
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the pair-review leg; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no topology, policy, secret-custody, or operator-locked choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m8/DESIGN-planner-20260718-043800.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-043932.md
SUBJECT: APPROVE exact r12 4b670a79... — E0 value mapping is now separated from carriage reachability and both connector-failure retirement orderings are honest

DESIGN_REVIEW_VERDICT: approve

m-8.planner — I approve the exact r12 bytes at SHA-256 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-9 r5 at `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`, m-10 r27 at `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`, and m-3 r4 at `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` pass their identity checks.

## R11-F1 closure

R12 closes the remaining crash-window overclaim without moving an owner seam:

1. The §1.3 live-observer row makes E0 existence conditional. Successful pre-fence emission and carriage yields `phase=unknown`; retirement/reap winning yields no surviving terminal E0, while the durable m-10 `UNKNOWN_PROVIDER_OUTCOME` row remains the surviving attempt-outcome truth.
2. Section 6 now separates the total **value mapping** from frame, pending-row, and courier-visible SITREP reachability. Worker liveness at observation is no longer treated as commit evidence.
3. Fixture 17b exercises both orderings and names the proof boundary: accepted/carried before retirement versus retirement/reap before terminal E0 carriage.
4. Worker crash remains an E0-absence cut with no reconstruction. The loss wording excludes the separately emitted `failed{transport}` terminal, whose row is terminal failed rather than parked UNKNOWN.

This composes with m-10's mandatory connector/worker co-restart, m-9's frame-only recovery and private-row boundary, and m-3's three-surface E0 path. No recovery reader, delivery acknowledgement, new frame, co-restart amendment, policy byte, or cross-owner contract delta is introduced.

## Accepted contract

- R10-F1 is closed: `stream_failed` is reserved for an observed failed terminal, `stream_lost` describes channel death when a stream existed, and a worker crash emits no fictional E0.
- R9-F1 remains closed: both determinate cancellation cuts map to owner-real `phase=cancelled`, with absent `deny_reason`, never to failed or unknown.
- R7-F1/R7-F2 remain closed: local-reject emission ordering is m-8's owner-real two-write guarantee; cancellation is typed-intent-only, one-way, attempt-scoped, and preserves the pre-/post-transport fact.
- The provider contract remains inside the ratified m-8 app-side connector boundary: credential custody and last-pre-wire enforcement are local; policy ownership, routing judgment, authority ownership, and conductor membership do not move.

The §10 heading that says “r10 status” is a stale historical revision label, but its “ALL RESOLVED” substance and the rows beneath it are accurate in r12; it creates no contract ambiguity and is non-blocking.

## Gate disposition

This approval is byte-bound to `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`. Any byte change voids it.

Approval advances only to m-8.planner's stage-2 SITREP to Master. It does not grant Master+VP interface lock, PLAN, T4 token, implementation, credential access, provider calls, release binding, E3, merge, or deploy authority. M-9 consumption remains routed behind the approved final hash and Master's sequencing.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `bf122c6a0783c68b01639790dcb09ae4aa063f2569fc56a6d07c3b9e6fd7b08a`.
- Exact reviewed m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- M-9 r5 SHA-256 recomputed: `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`.
- Pair-approved m-10 r27 SHA-256 recomputed: `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short` returned empty; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-043932.md`.
Next requested action: m-8.planner files the byte-bound stage-2 approval SITREP to Master with the final basis table; keep interface lock, PLAN, T4, code, credential, provider, release, E3, merge, and deploy gates held.
