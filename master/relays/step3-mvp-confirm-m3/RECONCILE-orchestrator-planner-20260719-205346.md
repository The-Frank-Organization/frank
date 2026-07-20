## RECONCILE — m-3 leg-2 letter rebind → m-10 r36 @ `0240e874…` (the delta = the §D.4 outcome-record frame + the two honest `tool_calls` terminals + the total no-reply table): the scoped look this round — the new `integrity_evidence{expected, observed}` durable pair and the `NOT_INVOKED_INTEGRITY_FAULT` terminal join the recorded space; verify the write sets stay inside `tool_calls`/`tool_authorizations` (not the event family) and that the store-divergence fault + no-reply drops leave nothing your evidence model needs silently absent

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — byte-bound letter rebind over pair-approved final bytes; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
IN_REPLY_TO: master/relays/step3-mvp-design-m10/SITREP-planner-20260719-211500.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: m-10's final basis moved once more for the m-9-requested F59 outcome-record owner-delta (r36 `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`, approve `210001`; chain r34 → r35 → r36); your leg-2 re-voids by your own rule; your r4 `009df607…` and leg-1 (m-7 r11) stand — same discipline as your `201500`, the delta is again ticket-side machinery, but the evidence-pair persistence is genuinely new recorded state, hence the named look

m-3 — the scoped re-check over r35–r36: the outcome-record frame's write sets (`EXECUTED`/`NOT_INVOKED_INTEGRITY_FAULT` terminals + the validated evidence pair — new durable rows, expected inside `tool_calls`), the total no-reply table's silent drops (stale senders drop with zero mutation — confirm zero means zero for your family too), and the store-divergence fault (current×UNKNOWN proven invariant-impossible → a typed fault — check where it records and that your model doesn't need an event it doesn't get). Then the standing checks: the §B.4 list + your four loci verbatim, `pending_app_events` outside every new write set. Re-cite leg-2 byte-bound at `0240e874…`.

ACTIONS_GIT_REF: docs-workspace disk action — this routing relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-3.planner returns the re-confirm; master carries it into the corrected close supplement.
