## RECONCILE — the FINAL rebind + closure routing to m-9 (owner set final: m-7 r11 `9331ea88…` · m-10 r21 `f4012ec5…`): leg-2 → r11 (your D-3, answered total) · leg-3 → r21 (your D-2/D-4/D-5 + the `attempt_open_ok` shape — resolve your contingent leg-1) · m-10's named precondition (the disclosure fields + frame directions confirm) · then the half r5 rebase over r11+r21+r6 → the DEFERRED fresh review → the closure SITREP

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — byte-bound rebinds + the batched half closure per the standing sequence; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the m-9 grill rides stage-4
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-mvp-design-m10/SITREP-planner-20260717-204500.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-7.planner, m-7.implementer, m-10.planner, m-10.implementer, m-8.planner, master.orchestrator-reviewer, operator
SUBJECT: everything your half was waiting for is owner-real — m-7 r11 (the three typed attach results in pinned order, suspended-first total over §2.4∪PREPARING, mismatch terminal-for-generation, truthful outcome/reason recording, FX-TB-19) + m-10 r21 (the 185818 batch through seven hardening rounds: acquisition-gated first admission ON the r11 tokens · D-4 state-only disclosure on both admission frames · D-5 total cross-family transition table · attempt_open_ok); rebind, confirm the disclosure shapes, rebase ONCE, run the deferred review, close the half

m-9 — the two owner folds your r4 marked AWAITING are pair-approved:

1. **Leg-2 rebind → m-7 r11 @ `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.** Your D-3 answered in full: `broker:attach-suspended` FIRST (total over the three §2.4 causes + PREPARING; nothing minted; tuple never evaluated; bounded hold-and-retry under supervision — your typed-transient-hold disposition licensed exactly) · `attach-ok`/`attach-tuple-mismatch` both requiring live-control+no-barrier+installed-state, mismatch TERMINAL for the presenting generation (your fenced-⇒-exit disposition licensed; no-hammering fixture in FX-TB-19). Verify your §1.6/§1.7 consumption rows against the pinned order + the outcome/reason recording.
2. **Leg-3 rebind → m-10 r21 @ `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852`** — AND resolve your contingent triple-leg-1: verify r21's **`attempt_open_ok`** is the durable-ack-after-row-commit CTRL-W shape you bound on (if conforming, your contingency auto-resolves per your own rule; if it deviates, that one leg re-verifies). Also **m-10's named precondition**: confirm the **exact D-2/D-4 disclosure fields + frame directions** (the acquisition-gated first-admission frames carrying the state-only disclosure) and the **D-5 transition-table half** your §2.9 family reciprocates.
3. **The half r5 rebase — ONCE, then close:** fold the r11 + r21 + m-8 bases — NOTE m-8 is rebasing r6 onto r21/r11 in parallel (routed this round) and will re-hash; cite m-8's FINAL post-rebase hash from their stage-2 SITREP if it lands before your fold, else cite r6 `ab63f6eb…` and take the letter-level refresh in the closure SITREP (r6's semantics are what you triple-confirmed; the rebase moves bases, not your seam) (+ the after-ack ordering sentence, the four-token reason list, and your D-2/D-3/D-4/D-5 consumption rows against the owner-real bytes) → the **deferred fresh uniquely-parented m-9.implementer review** → the **closure SITREP** naming the approved half hash. **The m-10 reciprocal confirmation routes on that SITREP** — it is the stage-3 close.

Return: the rebinds/confirms in THIS lane (one relay, legs separately dispositioned); the r5 fold + review + closure in the half's lane.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the rebind/confirm set; the half closes through its lane; master routes the m-10 reciprocal on the closure SITREP.
