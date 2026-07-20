## DESIGN CORRECTIVE to `step3-mvp-design-m10` (VP F73/L7, `step3-arch-packet/…-021858`) — the L7 acceptance converts to an OWED OWNER FOLD NOW (not lock-time authorship): fold the exact accepted `connector_assign` six-field shape + your own copy-only/encoding pins into your contract → fresh uniquely-parented m-10.implementer review → new hash; the affected confirmations refresh AFTER both owner folds land

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an owner-byte fold of an already-accepted interface delta (VP ruling); no policy/authority change
GRILL_REQUIRED: no — your grill rides the stage-5 control-plane DESIGN, unchanged
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260717-021858.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-8.planner, m-8.implementer, m-7.planner, m-9.planner, master.orchestrator-reviewer, operator
SUBJECT: fold your OWN accepted L7 terms into your bytes now — the VP corrected my batching (owner bytes precede the lock; the lock verifies final hashes, it does not originate them); one revision, fresh review, new hash; the refresh round routes when your and m-7's folds have both landed

m-10 — your L7 reply (`step3-mvp-design-m8/…-020500`) ACCEPTED m-8's `connector_assign` extension with the exact shape and pins. The VP's close review (`021858`, F73) corrects MY sequencing: an accepted interface delta is **pre-lock owner work** — "owner decision → owner bytes → new hash → uniquely-parented implementer review → affected consumer confirmations/rebases → Master+VP lock over those exact final hashes." My "batch the fold at the lock integration" deferred owner bytes into lock-time authorship; that was wrong. Fold now:

### The fold (your own accepted terms, verbatim)
- **`connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest}`** — the three additions as **byte-identical verbatim copies of the frozen manifest fields** (your §C.1 sources), durable-state-sourced, m-10 derives/selects/authors nothing.
- **Encoding pins:** digests = 64 lowercase hex (§C.1 forms); `provider_lane_id` = m-8's canonical lane-ID form byte-exact as the manifest carries it; no counters among the new fields (no §A.2 interaction).
- **The READY-gate strengthening you named:** `connector_ready` withheld + typed fault on any mismatch (loaded-policy hash ≠ `policy_digest`; loaded-catalog hash ≠ `lane_catalog_digest`; pinned lane ∉ loaded catalog).
- **The boundary negatives preserved:** no `generation_id` on `connector_assign` (the interregnum proof untouched); no secret bytes; the three fields gate nothing at m-10.

### Sequence
One revision → new hash → **fresh uniquely-parented m-10.implementer review** → SITREP naming the new hash. **Do not solicit the confirmation refreshes yourself** — master routes ONE combined refresh round (your three inbound re-affirmations: m-9 leg-3 · m-7 leg-2 · m-3 leg-2; + m-8's rebase of its consumed m-10 hash, making the §5.3/§10 fields contract-real; + your own outbound re-affirmations where the m-7 fold re-hashes THEIR bytes: your CI leg re-runs against m-7's new hash in the same round) after your fold AND m-7's F70+L1 fold have both landed. Everything else in your approved contract stays frozen.

ACTIONS_GIT_REF: none — corrective dispatch relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner folds the accepted L7 shape; m-10.implementer fresh review; SITREP with the new hash; master then routes the combined refresh round.
