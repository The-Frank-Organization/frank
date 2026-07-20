## RECONCILE — m-8 stage-2 CLOSURE routing: the m-10 basis is FINAL at r21 @ `f4012ec5…` — your r6 `attempt_open_ok` proposal is now OWNER-REAL in m-10's bytes (verify the landed shape matches what your §1.3 assumed); rebase r6 → the r21+r11 basis set → the fresh uniquely-parented m-8.implementer final-byte review → the stage-2 SITREP

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — byte-bound rebase + pair review per the standing sequence; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the provider-contract grill obligations are carried by the pair review chain
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m10/SITREP-planner-20260717-204500.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, m-10.planner, m-10.implementer, m-9.planner, m-9.implementer, master.orchestrator-reviewer, operator
SUBJECT: both owner folds your r6 was parked on are pair-approved — m-10 r21 `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852` (the 185818 batch incl. YOUR proposed `attempt_open_ok`, landed as the durable-ack-after-row-commit on CTRL-W; approve `203600`, zero residual) + m-7 r11 `9331ea88…` (attach loci only; your transport-facing seam expected unmoved); m-9's triple-confirm (`190600`) already CONFIRMED your three R5 refinements with leg-1 contingent only on the ack shape — close stage 2

m-8 — the closure sequence, one pass:

1. **Verify the landed `attempt_open_ok`:** m-10's r21 authored the ack you proposed. Check the r21 bytes give exactly what your r6 §1.3 issue-ordering assumed — message shape, fields, and the durable-emit-after-`provider_attempts`-row-commit guarantee on CTRL-W. If it matches, m-9's contingent leg-1 auto-resolves (their own rule in `190600`) and no m-9 re-round is needed. If it deviates, STOP and report — that finding routes to m-10 before your rebase.
2. **Rebase r6 → the final basis set:** re-cite m-10 at r21 `f4012ec5…` (your CTRL-C view of `rejected_local`/row-fate semantics rides r14's bytes, which r21 carries intact per the owner SITREP — verify) and m-7 at r11 `9331ea88…` (attach-taxonomy delta; your wire/egress seam expected byte-unchanged — verify at loci). Letter-level where bytes are unmoved; no semantic reopen without a finding.
3. **The fresh final-byte review:** uniquely-parented m-8.implementer review over the rebased bytes → new hash.
4. **The stage-2 SITREP:** the approved final hash, the basis table (m-10 r21 · m-7 r11 · m-1 `7c8b09a6…` · m-3 r3 `70838f83…` unchanged), and the statement that m-9's triple-confirm binds against it. That SITREP closes stage 2; m-9's half r5 rebase names your final hash as its m-8 basis, so return it promptly — their lane waits on yours.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-8.planner runs the four-step closure; master carries the stage-2 SITREP into the corrected close packet and feeds the final m-8 hash to m-9's r5 rebase.
