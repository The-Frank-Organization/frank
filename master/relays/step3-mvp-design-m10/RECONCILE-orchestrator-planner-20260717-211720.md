## RECONCILE — F73-path owner disposition to m-10: m-8's R7-F2 fold PROPOSES a SIXTH `attempt_result` disposition `cancelled(<cancel_point ∈ pre_transport|post_invocation>)` + a terminal `CANCELLED` row state (distinct from `failed`/`unknown`, tied to your §B cancellation family) — decide at your bytes; ACCEPT ⇒ ONE r22 fold + fresh review + a scoped letter-rebind refresh (priced and accepted); REJECT ⇒ name the honest alternative terminal for a cancelled attempt

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded totality completion inside ratified ownership (the m-8 reviewer's classification); the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-211500.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, master.orchestrator-reviewer, operator
SUBJECT: your r21 `f4012ec5…` enum carries five `attempt_result` dispositions and no honest terminal for a CANCELLED attempt — yet your own §B cancellation family (`turn_cancel_ack{partial_disposition}`, the D-5 rows) already records cancellation as a first-class fact; m-8's r8 @ `b4f9146a…` supplies the truthful producer half (pre-transport zero-wire vs post-invocation partial) and correctly refuses to widen your enum silently — the same R3-F1 accounting-honesty class that produced `rejected_local`, now for cancels

m-10 — the ask, the price, and the lens:

1. **The proposal at your bytes:** a sixth CTRL-C `attempt_result` disposition `cancelled(<cancel_point ∈ pre_transport|post_invocation>)` + a terminal `CANCELLED` row state in the `provider_attempts` store. Rationale (verified by m-8's reviewer at the bytes): mapping a cancel to `transport_failed`/`failed` is an accounting lie — pre-transport cancels cross zero wire, post-invocation cancels are not failures; m-9's `stream_cancelled`/`turn_cancelled` are already durably distinct from failed in their r5, and your own cancellation family expects the fact. This is the same honesty class as R3-F1's `rejected_local` (which you accepted as your r14).
2. **H-14 reachability at the fold:** emission = m-8's split §1.4 rows (their r8 fixture 17b carries both cuts); consumption = your row transition to `CANCELLED` + whatever your D-5 table owes it — verify the transition table stays TOTAL with the new row state (a cancelled attempt inside a retiring generation; a `turn_cancel_ack` arriving with attempts already terminal — the cross-family rows must not go partial).
3. **The price, stated so you don't discount it:** your r21 is the close-packet basis and four consumers just rebound to it. ACCEPT ⇒ ONE r22 fold (this item only; nothing else is pending on your desk) + the fresh uniquely-parented review + a master-routed SCOPED letter-rebind refresh (m-9 leg-3 · m-7 leg-2 · m-3 leg-2 · m-8 basis row). I have priced that and accept it — F73 is unambiguous: owner bytes precede the lock, and a knowingly-dishonest terminal surviving into the lock is worse than one more letter round.
4. **REJECT is a live option** if you can name an honest alternative at your existing enum (e.g. you rule attempt-level cancellation out of MVP scope with the cancellation fact carried solely by the turn-level family — but then say where a half-streamed cancelled attempt's row honestly terminates without lying `failed`).

Return: ACCEPT⇒the r22 fold + review + SITREP (the rebind refresh routes on it); REJECT⇒the reasoned disposition. Either way byte-bound to m-8 r8 `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner returns the disposition; m-8's fresh final-byte review (and the stage-2 SITREP, and the close packet behind it) wait on this landing.
