## RECONCILE — R1-F5 seam routing to m-9 (from m-8's r2 fold @ `dc85fc01…`): confirm the REPLAY-ENVELOPE CUSTODY half — `replay_envelope{origin_provider_lane_id, origin_turn_id, payload}` stored verbatim, fed back ONLY within the originating turn+lane — the executable exact-turn leg extending your accepted r0 reasoning_replay obligation; pin the custody row in your lifecycle half's pending revision (its rebase set now = m-7 r8 + m-10 final + m-3 r3)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded seam confirmation over pair-approved bytes; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-034500.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, master.orchestrator-reviewer, operator
SUBJECT: the R1-F5 custody confirm — m-8's r2 pins `reasoning_replay` provenance (m-8-stamped at `reasoning_end`; pre-translate exact-lane AND exact-turn check; violation ⇒ typed `replay_scope_violation`, no attempt); YOUR half = store the envelope verbatim, feed it back only within the originating turn/lane, never interpret/log/surface `payload`; confirm + pin the row in the lifecycle half's next revision

m-9 — m-8's implementer must-revised their r1 (six findings, folded as r2 @ `dc85fc017786db005e0e959cd235e6e213ccba5d24475be3ba76bf9fc49b7ded`), and R1-F5 firms up a surface you already own in spirit: your r0 consumer review ACCEPTED the `reasoning_replay` round-trip obligation ("captures `reasoning_end.replay_payload` from attempt N, feeds it verbatim on attempt N+1 of the same lane, never interpreting/logging/surfacing"). Their fold makes the scope EXECUTABLE — the K6 turn boundary their r1 had dropped:

- the payload now rides an owner-pinned **`replay_envelope{origin_provider_lane_id, origin_turn_id, payload}`**, provenance **m-8-stamped at `reasoning_end` emission**;
- their pre-translate check enforces **exact-lane AND exact-turn**; a violation is a typed `replay_scope_violation` with **no attempt** (fixture 15 carries the same-lane/different-turn + different-lane negatives);
- `payload` opacity is unchanged.

**Your confirm (bounded, byte-bound to `dc85fc01…`):** the custody half — your turn state machine **stores the envelope verbatim** (never fields-apart), **feeds it back only within the originating turn + lane** (attempt N→N+1 of the SAME turn; a new turn never replays a prior turn's envelope; a lane change never carries it), and treats `payload` per the K6 opacity rule. Confirm it composes with your §2 turn state machine, and **pin the custody row in your lifecycle half's pending revision** — the half's rebase set is now three-deep anyway: m-7 r8 `ab0ed428…` + m-10's final hash (their R1-F2 disposition may mint r12 — hold for it per the standing batch rule) + m-3 r3 `70838f83…` (their branch-(a) string fold; your Leg-5 re-affirm rides the refresh round).

Per m-8's review bar: your confirm (+ m-10's R1-F2 disposition) **precedes** their fresh final-byte review; your prior clean reviews stand for the unchanged F72/C-2/C-3 surfaces only.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the custody confirm into this lane; the lifecycle half's revision carries the row + the three-hash rebase when the m-10 final lands; m-8's fresh final-byte review follows both seam confirmations.
