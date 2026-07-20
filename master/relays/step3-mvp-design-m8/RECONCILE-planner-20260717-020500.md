## RECONCILE return — L7 bounded reply: **ACCEPT** m-8's `connector_assign` extension `{policy_digest, provider_lane_id, lane_catalog_digest}` with the exact shape + two copy-only pins below; NO byte fold now (`79fcf742…` frozen); fold batched at the lock integration per the routing

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded cross-interface reply; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-015435.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-8.planner, m-8.implementer
BUNDLE_ID: m-10-app-control-plane
SUBJECT: ACCEPT — the exact accepted shape: `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest}`, all three additions VERBATIM COPIES of the frozen-manifest fields, durable-state-sourced, no generation_id, no secret bytes; the scope-guard classification (freeze-time facts, not policy/authority) is my read too

master — m-8's §10 ask read at their exact bytes (`9d86c165…` recomputed, matches; their §5.3 load sequence + §10 row). **ACCEPT**, on these terms:

1. **Exact accepted shape:** `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest}` — the three additions are the run-frozen verification comparands their authorize/catalog path consumes (loaded-policy hash == `policy_digest`; loaded-catalog hash == `lane_catalog_digest`; pinned lane ∈ loaded catalog), with `connector_ready` withheld + typed fault on any mismatch. That STRENGTHENS my §B.1 READY gate (a connector that cannot prove it loaded the run's exact policy/catalog never becomes READY) — fail-closed in my direction, so I have no boundary objection.
2. **Copy-only pin (the same discipline as everything m-10 writes):** the three values are **byte-identical verbatim copies of the frozen manifest fields** (my §C.1: `policy_digest`, `provider_lane.lane_id`, `provider_lane.lane_catalog_digest`) — m-10 derives, selects, and authors nothing; sourced from durable frozen state only, like the rest of `connector_assign`.
3. **Encoding pin:** digests = 64 lowercase hex (the §C.1 forms); `provider_lane_id` = m-8's own canonical lane-ID form byte-exact as the manifest carries it. No counters among the new fields, so no §A.2 string-grammar interaction.
4. **Boundary check (the scope guard — concur):** freeze-time FACTS, not policy (m-3 authors the policy; m-10 copies its digest), not credential material (all three sit in m-1's NOT-secret census), not new authority (they gate nothing at m-10; they let m-8 verify its own loads). The interregnum proof is untouched — no `generation_id` rides `connector_assign`, before or after.
5. **No fold now:** `79fcf742…` stays frozen per the routing; the accepted fold lands at the Master+VP lock integration (batched with L1), takes a fresh uniquely-parented m-10.implementer review, and the affected confirms re-affirm in that round. m-8 may proceed on their unchanged r0 consumption meanwhile — their doc already treats this as an ask.

## Verification
- m-8 r0 hash recomputed: `9d86c165db6e48fe…` (exact); my doc `79fcf742…` unchanged (no edit made by this reply).
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-020500.md` — run at filing; result reported inline.

ACTIONS_GIT_REF: none — a bounded reply relay + one INDEX.md row timestamped 20260717-020500; no design-doc edit, no `frank/` edit, no code, no lock, no credential, no provider call.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: master carries this ACCEPT (+ the two pins) as lock-round input L7; the batched fold + fresh m-10 pair review + confirm re-affirmations land at the interface-lock integration.
