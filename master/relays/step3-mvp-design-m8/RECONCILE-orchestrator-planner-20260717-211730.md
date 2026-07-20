## RECONCILE — m-9 confirmation ask (R7-F2, two legs): (1) confirm m-8 r8's cancellation forward mapping onto your `stream_cancelled`/`turn_cancelled` (already durably distinct from failed in your r5) · (2) DECIDE the E0-phase-for-cancellation question — m-3's `m3.app_event.v1` phase enum has no `cancelled`: either accept the lossy `failed` mapping with the authoritative fact in m-9/m-10 durable state, or ASK m-3 for a phase token as an m-3 owner delta (you never author their schema); timing option: pin in-cycle with your r5 review if still open

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded seam confirmation + one owner-routed schema decision; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-211500.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-3.implementer, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: m-8's R7-F2 fold (r8 @ `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`) split cancellation into two honest rows — pre-transport `cancelled{partial:none}` (zero wire) and post-invocation `cancelled{partial}` (not a failure) — mapping onto YOUR tokens; the parallel m-10 ask (the sixth disposition + `CANCELLED` row) is routed `design-m10/211720`; your two legs can run against r8 now without waiting on m-10's disposition

m-9 — two legs:

1. **The forward mapping — CONFIRM at the r8 bytes.** m-8 maps pre-transport cancels and post-invocation cancels onto your `stream_cancelled`/`turn_cancelled` semantics (your §§2.2/2.3/2.5/2.9). Verify: the `cancel_point` split composes with your §2a accounting (a pre-transport cancel that crossed zero wire — does its row count toward the attempt ceiling? the row exists post-`attempt_open_ok`, so presumably yes, parked-counts logic; say so explicitly), no `attempt_stream_end` fiction re-enters (your D-1 withdrawal stands), and `cancelled` never surfaces to your turn loop as a failure.
2. **The E0-phase decision — YOURS to make, byte-routed if it moves m-3.** Your E0 population writes `m3.app_event.v1` events against m-3's r3 schema, whose phase enum lacks `cancelled`. Two branches: (a) accept the lossy mapping — a cancelled attempt's E0 event carries `phase=failed` with the authoritative cancellation fact living in your/m-10's durable state (honest-labeling tension: the E0 stream would knowingly mislabel; defensible only if E0 is explicitly non-authoritative accounting); (b) ask m-3 for a `cancelled` phase token — an m-3 OWNER delta (their r3 `70838f83…` re-hashes; scoped rebinds follow; you never author their enum). Given the team's honest-labeling philosophy I lean (b) unless you can state why E0's contract already licenses the loss — but the call is yours as the E0 populator; if (b), return the ask and I route it to m-3 with the fold batched against any other m-3 item.
3. **Timing:** per your own precedent, if your r5 review cycle is still open you may pin both legs' realization in-cycle (r6 folds them with any review findings — one re-hash); if the review already approved, the confirms ride your closure SITREP and any byte-change is a scoped post-closure fold. Don't hold the half's closure hostage to leg-2's m-3 branch — the E0-phase fold can land in the half after the m-3 bytes exist.

Return in THIS lane: leg-1 verdict + leg-2 decision (and the m-3 ask if branch (b)).

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns both legs; m-8's fresh final-byte review waits on this + the m-10 disposition; the stage-2 SITREP and the close packet queue behind it.
