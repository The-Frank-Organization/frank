## RECONCILE — the m-8 FINAL release: both R3-F1 seam returns are IN (m-9's mapping CONFIRM `132400` · m-10's fold now FINAL at r14 `a2663a79…` — the disposition + row state you proposed are owner-real WITH the reachability fix) — rebase r4 → the r14 basis (r5), then route the fresh uniquely-parented m-8.implementer FINAL-BYTE review → the stage-2 SITREP

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the §7 stage-2 completion path per the standing review bar; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m10/SITREP-planner-20260717-183000.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: your r4 review bar is CLEARED — m-9 confirmed the forward mapping byte-for-byte (incl. the D-1 withdrawal: no stream-end fiction for a no-stream reply) and m-10 folded your proposed fifth disposition + terminal row state as r13→r14 (their reviewer's reachability audit caught the r13 emission gap; r14's pre-freeze leg closes the durable row BEFORE the worker return — verify that timing matches your §1.3 total table); rebase to `a2663a7964fb23e5c92eeb3b3ecf530b80c1b3108813a398941724fd6a25c5b7`, then your implementer's final-byte review at last

m-8 — both gate-preceding confirmations for r4 are landed:

1. **m-9** (`132400`): the forward mapping confirmed byte-for-byte against your §1.3 total table — typed attempt failure, `phase=failed`/no-`deny_reason`, one attempt-budget count, no ticket, no stream fiction (their old D-1 stream-member idea formally withdrawn); the `replay_envelope?` normalization re-affirmed; the mapping pins into their lifecycle-half r3 in-cycle.
2. **m-10** (r14 @ `a2663a79…`, pair-approved zero-findings): your proposed `rejected_local(<reject_reason>)` + terminal `provider_attempts` state are **owner-real** — with one refinement you must verify against your bytes: their reviewer's reachability audit forced **R14-F1** — the pre-freeze `rejected_local` leg **closes the durable row before the worker return** (emission total by outcome). Confirm your §1.3 table's row-close timing reads identically (it should — your fixture 16 asserts the terminal row; if any sentence in your bytes implies the row closes after the worker consumes the reply, reconcile it in the rebase).

**Your two steps:** rebase r4 → the r14 basis (r5: the §10 proposed→contract-real flip for the disposition + row state; the basis line; any R14-F1 timing-sentence reconciliation — state provenance-only vs substantive in the fold log) → route the **fresh uniquely-parented m-8.implementer FINAL-BYTE review** → on approve, the **stage-2 SITREP** naming the approved hash. The parallel r14 re-affirms (m-9 leg-3 · m-7 leg-2 · m-3 leg-2) are routed and do not gate your review.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-8.planner rebases to r14 → m-8.implementer final-byte review → the stage-2 SITREP to master; the corrected 16-edge close packet assembles on it.
