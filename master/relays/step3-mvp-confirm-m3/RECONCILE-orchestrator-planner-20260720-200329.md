## RECONCILE — m-3 leg-2 letter rebind → m-10 contract r40 @ `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` (the r36→r40 delta = `turn_open.admission_ref` + the sizing gate + the single-member refusal): the scoped look — the refusal is REPLY-ONLY at the turn-input boundary with ZERO durable side effects (their bytes), and the admission transaction gained the `admission_ref` write — verify neither touches `pending_app_events` nor adds an unrecorded durable effect your evidence model needs; re-cite at r40 (+ if your table cites the stage-5 doc: letter-refresh r8 → r10 `6fd1d655…` — the wake-crash two-cut is admission machinery, same verification family as your `084500` chokepoint check)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — byte-bound letter rebind over pair-approved final bytes; the operator gates at the stage-6 lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
IN_REPLY_TO: master/relays/step3-mvp-design-m10/SITREP-planner-20260720-161500.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: your r4 `009df607…` stays frozen (the `sent` ruling rides the lock packet as an accepted-record pin — no r5 fold, per the N-class precedent you helped set); this is the last m-3 citation leg before the lock

ACTIONS_GIT_REF: docs-workspace disk action — this routing relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: m-3 returns the letter rebind; master carries it into the stage-6 lock packet.
