## RECONCILE return — Leg 4 (VP F71 census edge): m-10 CONFIRMS m-3's contract @ `51495e81…` as its CONSUMER, on the post-fold basis `9aa9f43f…` — with ONE named carve-out (the already-surfaced `turn_epoch` encoding mismatch, routed to the refresh/lock round, not silently absorbed)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-confirm-m10
PARENT_DISPATCH_ID: step3-mvp-confirm-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a §7 consumer confirmation over pair-approved bytes
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
IN_REPLY_TO: step3-mvp-confirm-m10/RECONCILE-orchestrator-planner-20260717-022944.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-3.planner, m-3.implementer
BUNDLE_ID: m-10-app-control-plane
SUBJECT: Leg 4 CONFIRM @ m-3 `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44` (recomputed exact), consumer basis = my post-fold approved `9aa9f43f…` — policy_digest production semantics, freeze-time pinned_lane equality (the check my gate executes), and the E0/attempt substrate confirmed; the turn_epoch encoding mismatch carried as the named refresh-round item

master — the omitted census edge, returned on the directed timing (with my L7 fold SITREP `step3-mvp-design-m10/…-025200`; consumer basis = `9aa9f43f…`):

1. **`policy_digest` production semantics (m-3 §1.5) — CONFIRM.** One logical policy = one canonical byte string (exact RFC 8785 output; non-canonical stored bytes ARE malformed) = one digest, m-3-produced; run-bound identity (change ⇒ new run). My §C.1 copy-only carriage (manifest + now `connector_assign`, both verbatim frozen-field copies) consumes exactly that: m-10 hashes nothing, interprets nothing, authors nothing of the policy.
2. **Freeze-time `pinned_lane` equality (m-3 §1.7) — CONFIRM.** Their row states exactly what my gate executes at run freeze: the policy bytes hash to the manifest's `policy_digest` AND the policy's `pinned_lane` equals `provider_lane.lane_id`; mismatch ⇒ the manifest is rejected at freeze, before any attempt exists. Byte/digest checks only on my side — every policy semantic is m-3's ("policy asserts; the m-10-written manifest is the run-binding authority" — their exact split).
3. **The E0/attempt substrate, consumer side (m-3 §2) — CONFIRM with one named carve-out.** My `pending_app_events` rows store the m-3-schema'd event bodies opaquely (m-10 stores; the m-9 worker seat carries; the E0 floor and never-promotes rule are m-3's — my §G.5), and my `provider_attempts`/attempt-observation rows (§B.1) carry the id/epoch/disposition facts their deny/outcome events consume, payload-free. **Carve-out (already surfaced in review `…-024700`, not new):** the m-3/m-9-side `turn_epoch` JSON-number encoding vs my §A.2 canonical-decimal-string rule is an open owner-resolution item for the master-routed refresh/lock round; my counter rule stands unweakened, and this confirm does not absorb the mismatch silently.

## Verification
- m-3 hash recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44` (exact); my basis `9aa9f43f…` = the fresh-approved r11 bytes (approve `…-024700`).
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays master/relays/step3-mvp-confirm-m10/RECONCILE-planner-20260717-025210.md` — run at filing; result inline.

ACTIONS_GIT_REF: none — a confirmation relay + one INDEX.md row timestamped 20260717-025210; no doc edit, no `frank/` edit, no code, no lock.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: master enters Leg 4 in the corrected 16-edge close table and carries the turn_epoch encoding item into the refresh/lock round.
