## RECONCILE — m-9 consumer confirmation, Leg 2 of 4: the m-7 transport/broker contract — CONFIRM (byte-bound @ `f072bd99…`)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m9
PARENT_DISPATCH_ID: step3-mvp-confirm-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded byte-bound consumer confirmation; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-7.planner, m-7.implementer, master.orchestrator-reviewer
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-confirm-m9/RECONCILE-orchestrator-planner-20260717-010020
SUBJECT: CONFIRM — m-7's stage-1 contract (`master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md`, r6, SHA-256 verified `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`) — the caller seam + capability/attach surface are sufficient for both m-9 consumers (native tool + worker); no capability m-9 needs is absent; worker-side error-disposition obligations accepted; no findings

CONFIRMATION (the §4 m-9 obligations, against the exact bytes):
1. **The two-operation caller seam — CONFIRM.** `Call` over exactly the three canonical relay verbs + typed `Describe` (metadata, never a fourth verb, never in the 8-name dispatch set) is the complete transport surface both m-9 consumers need: the native tool (submit/project/read + the m-2 re-render refresh path via `Describe`) and the worker's rediscovery loop (catch-up `project`/`read` + `Describe` through the broker). Placement-agnostic consumption against the interface (§1.2), with `ManagedClient` and the worker-side capability client as the two implementations, fits the m-9 design exactly. **No capability m-9 needs is absent** — checked against the coding-agent loop: relay verbs ✓, form rediscovery ✓, push receipt ✓, nothing else is on the m-9 relay surface by design (the manifest-fixed 8-name set needs no runtime tool listing).
2. **The closed retry contract — CONFIRM.** Connection-loss-class-only, at-most-one-retry, single-flight replacement with push-reader continuity, and the broker-path retry gate re-entering the epoch fence (§1.5) are consumable and honest; the m-9 loop treats a served application rejection as terminal (never re-sent by transport) and `broker:unknown-outcome` on a fenced submit retry as a rediscovery trigger (the record is truth) — both fit the worker's turn state machine.
3. **`relay.read`'s conditionally-mutating classification — CONFIRM.** The §1.5.4 disposition contract (any later read returns the authoritative current disposition: `checksum-mismatch` pending or `record-quarantined` with incident identity; duplication impossible per the four landed mechanisms) fits the worker's rediscovery loop: a lost first response is healed by re-read, no worker-side dedup or repair bookkeeping is needed.
4. **The USE capability + attach surface — CONFIRM.** The §2.3 capability (opaque, per-generation, connection-scoped, epoch-bound, broker-memory-only) + the §2.10 attach protocol (present `{run_id, generation_id, turn_epoch}` from m-10's `assign`; capability returned on that connection) + the §2.8 closed worker surface (three relay calls + typed `Describe` + push receipt — nothing else) are exactly the surface the m-9 worker consumes: **the worker path needs no credential bytes, no direct dial, and nothing beyond that closed surface — confirmed.** The attach tuple matches m-10's CI-2 `assign` shape byte-for-byte (cross-checked at `79fcf742…` §B.1).

ACCEPTED OBLIGATIONS (m-9 lanes, named): worker-side typed handling of the broker classes — `broker:stale-epoch` (generation fenced: fail closed, no local retry; supervision replaces), `broker:suspended`/`preparing` (transient: hold, re-invoke read-only ops), `broker:record-unavailable` (payload withheld: re-invoke/rediscover; for submit, rediscover the committed truth), `broker:unknown-outcome` (rediscover) — plus reattach-after-broker-restart with FRESH capability material at unchanged epoch (§2.7 row 2, m-1 MR2 aligned), and rediscovery-not-push as the only delivery guarantee consumed (§2.6).

NOTES (non-blocking): the §3 conductor-identity producer contract and §3.3 relay-leg evidence are outside the m-9 consumer surface (observer/m-3/Master-VP edges) — read, no objection, not confirmed here as they are not mine to confirm.

Duplicate/already-built gate: not applicable — a bounded confirmation over named bytes.
Boundary contract: not applicable — no artifact beyond this confirmation; consumer = master's stage-1 confirmation table.

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no frank/ edit; no code
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: master holds this leg for the stage-1 confirmation table; no m-7 action owed.
