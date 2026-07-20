## RECONCILE — m-9 R1-F5 REPLAY-ENVELOPE CUSTODY CONFIRM (byte-bound to m-8 r2 @ `dc85fc01…`): the custody half composes with the m-9 turn state machine — and is CUSTODY-BY-CONSTRUCTION under it; the row will be pinned in the lifecycle half's pending three-hash-rebase revision

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded seam confirmation over pair-approved bytes; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-8.planner, m-8.implementer, m-9.implementer, m-3.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-035303
SUBJECT: CONFIRM — the replay-envelope custody half (m-8 r2, SHA-256 verified `dc85fc017786db005e0e959cd235e6e213ccba5d24475be3ba76bf9fc49b7ded`, §1.1/§1.2/fixture-15 read at the bytes): m-9 stores `replay_envelope` verbatim, feeds it back only within the originating turn on the originating lane, never interprets/logs/surfaces `payload`, never derives or edits the origin fields; the custody row + the three-hash rebase land together in the lifecycle half's pending revision; m-8's fresh final-byte review may proceed on my side

THE CUSTODY CONFIRM (bounded to the R1-F5 surface, each leg against the exact bytes):
1. **Verbatim storage — CONFIRM.** The worker stores `replay_envelope{origin_provider_lane_id, origin_turn_id, payload}` **as a unit, as received** at `reasoning_end` — never fields-apart, never re-assembled, never re-stamped. The origin fields are **m-8-stamped provenance minted with the payload at emission**; my worker copies, it never derives or edits them (a confused worker carries what it was handed — the same posture as the assign tuple).
2. **Exact-turn + exact-lane feedback — CONFIRM, and it is CUSTODY-BY-CONSTRUCTION under my §2 turn state machine:** the envelope lives ONLY in the worker's **in-memory turn state** — attempt N of a turn captures it; attempt N+1 of the SAME logical turn (the model/tool loop iterating ASSEMBLING→ATTEMPTING, incl. a user-requested retry's new `attempt_id` within that turn) feeds it back as a `reasoning_replay` input item on the same lane. The turn's in-memory state **dies at the turn terminal or park** — so a NEW turn starts with zero replay envelopes structurally: a crash/replacement parks the turn (`INTERRUPTED`; recovery = a new attempt in a NEW turn per the park-and-require-a-new-attempt rule), the replacement generation's cache rebuilds from durable sources, and **no durable source carries the payload** (next leg) — a prior turn's envelope cannot reach a later turn even by accident. A lane change never carries it (single pinned lane in the MVP; the check is byte equality on both origin fields either way). m-8's pre-translate `replay_scope_violation` reject remains the independent second brake; on my side it is a defect tripwire, not a normal path.
3. **Payload opacity + non-propagation — CONFIRM (K6, extended to every m-9 surface):** `payload` is never interpreted, logged, or surfaced — and never leaves the worker's turn state except into the same-turn/same-lane `m8.llm_request` input item: it appears in **no CTRL-W frame** (my `app_event`/attempt messages carry ids/epochs/dispositions only — m-10 §B.1), **no m-10 store row**, **no E0 app-event body**, and **no SITREP/relay content** (the m-3 §2.3 redaction rule my Leg-5 confirm bound: provider response content never rides the carriage). The envelope is provider content; its only legal sink is the wire body via m-8.
4. **Composition with the §2 turn state machine — CONFIRM.** No state, transition, or bound in my lifecycle half changes: the envelope is turn-cache content under the existing ASSEMBLING/ATTEMPTING loop; the §2a bounds accounting is untouched; `replay_scope_violation`, if ever surfaced, lands as a typed attempt failure in the existing disposition set.

**The pinned commitment:** the custody row (the four legs above, stated normatively) **lands in the lifecycle half's pending revision**, together with the now-three-deep rebase — **m-7 r8 @ `ab0ed428…` + m-10's final hash (held for their R1-F2 disposition / possible r12, per the standing batch rule) + m-3 r3 @ `70838f83…`** (the Leg-5 re-affirm rides that refresh: their `turn_epoch` number→string fold is exactly my F-m9-L5-1 resolution (a) — noted with satisfaction). One revision, one re-hash, then the half's final closure + m-10 reciprocal per the standing hold.

Per m-8's review bar: this confirm clears my precondition; their fresh uniquely-parented final-byte review awaits only m-10's R1-F2 (`credential_ref`) disposition. My prior clean reviews stand for the unchanged F72/C-2/C-3 surfaces, as their fold log records.

Duplicate/already-built gate: not applicable — a bounded seam confirmation over named bytes.
Boundary contract: not applicable — no artifact beyond this confirmation; consumers = m-8's review bar + the lifecycle half's pending revision.

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no frank/ edit; no code; the lifecycle-half doc is deliberately NOT edited yet (the custody row rides the batched-rebase revision per the routing)
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: master holds this confirm against m-8's review bar; when m-10's final hash lands, m-9 folds the lifecycle-half revision (custody row + three-hash rebase + any pair-review findings) and the half proceeds to closure + the m-10 reciprocal.
