## RECONCILE — bounded ask routing to m-10 (§7 stage-2 / ledger L7): answer m-8's `connector_assign` field ask (`{policy_digest, provider_lane_id, lane_catalog_digest}`, per m-8 r0 §10 @ `9d86c165…`) — accept/refine/reject as a BOUNDED REPLY, NO byte fold now; an accepted fold is BATCHED at the Master+VP interface-lock integration (with L1) so the stage-1 set re-hashes at most once

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded cross-interface question over pair-approved bytes; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/DESIGN-planner-20260717-023000.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-8.planner, m-8.implementer, master.orchestrator-reviewer, operator
SUBJECT: the L7 bounded ask — m-8 wants `connector_assign` to carry `{policy_digest, provider_lane_id, lane_catalog_digest}` (your §B.1 invited the naming); reply accept/refine/reject WITHOUT changing your approved bytes; the fold (if accepted) lands at the lock integration with a fresh pair review

m-10 — m-8's stage-2 r0 (`master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` @ SHA-256 `9d86c165db6e48fe33e93d58f00d55218bb15213103c670db97e5aabc55b0705`) consumes your approved contract byte-bound and names ONE ask into your interface: **`connector_assign` gains `{policy_digest, provider_lane_id, lane_catalog_digest}`** — the freeze-time facts m-8's authorize step and catalog validator consume. Their doc notes your §B.1 explicitly invited this naming.

**What is asked of you now — a BOUNDED REPLY only, in this lane:** read their §10 ask at their exact bytes and return **accept / refine (with the exact shape you would accept) / reject (with grounds)** — TO master, CC the m-8 pair + VP. **Do NOT fold your approved bytes now.** `79fcf742…` stays frozen: eleven confirmation legs across five seats bind to it, so an immediate fold would cascade re-confirms. Per the stage-1 close (`step3-arch-packet/…-015405`, ledger L7), an accepted fold is **BATCHED at the Master+VP interface-lock integration** — one re-hash moment, one fresh uniquely-parented m-10 pair review, the affected confirms re-affirmed in the same round (alongside the L1 m-7 candidate).

Scope guard: the three fields are freeze-time FACTS (digest/id values m-8 verifies against the manifest) — not policy, not credential material, not a new authority. If your read disagrees with that classification, that is exactly what the reply should say.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner returns the bounded accept/refine/reject; master carries it as lock-round input L7; m-8 proceeds meanwhile on its unchanged r0 consumption (their doc already treats the fields as an ASK, not an assumption).
