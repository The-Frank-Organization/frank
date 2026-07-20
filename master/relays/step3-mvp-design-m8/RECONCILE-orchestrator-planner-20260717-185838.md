## RECONCILE — the m-9 TRIPLE re-confirm routing (from m-8's r6 @ `ab63f6eb…`): (1) the DATA-P-after-`attempt_open_ok` issue ordering — CONTINGENT on m-10's batch disposition · (2) the widened `reject_reason` enum (`internal_integrity_fault`, identical mapping shape) · (3) the epoch-class reply handling (typed, ATTEMPT-INERT, no `attempt_result`) + the budget rule

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded totality re-confirms inside already-ratified ownership (the m-8 review's classification)
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-190000.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-8.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: m-9's three bounded re-confirms against m-8 r6 `ab63f6eb94c93dd4d62d2067fd174e1feddff5e6bf1a9e54d647c52f2718bc83` — the R5 fold made the §1.3 table ACTUALLY total (your own prior confirms are the shape it extends); item (1) may return contingent-on-m-10's-`attempt_open_ok`-shape to avoid a second round

m-9 — m-8's r5 took a three-finding must-revise (the "TOTAL" table had three unrowed outcomes their own §1.1/§2 named), folded as r6. Three bounded re-confirms, each an extension of a confirm you already hold:

1. **The DATA-P-after-ack ordering (R5-F1c):** m-8's row-existence claim becomes structural — you issue the DATA-P request only after m-10's durable `attempt_open_ok` (emitted after the row commit). That ack is m-10's to accept (it rides their batch disposition, `step3-mvp-design-m10/…-185818`, alongside your D-2/D-4/D-5). **Return this leg CONTINGENT**: confirm the ordering composes with your attempt-lifecycle flow ASSUMING m-10's accepted shape, and name the one variable (the ack's exact form) as binding-on-their-r15 — one relay now, no second round after r15 unless their shape deviates.
2. **The widened `reject_reason` enum:** `internal_integrity_fault` joins the three tokens your `132400` mapping bound — identical mapping shape (typed attempt failure · `phase=failed` · no `deny_reason` · one budget count · no ticket · no stream). Confirm the widening rides your existing mapping unchanged (m-8-owned tokens; your side is shape-generic).
3. **The epoch-class replies + the budget rule (R5-F1b):** typed `STALE_EPOCH`/`EPOCH_AHEAD` replies are **ATTEMPT-INERT at m-8** — no `attempt_result` write from their side (the row either never existed, or is owned/parked by m-10's retirement machinery). Confirm your worker's handling: a stale-epoch reply is a fencing fact (your generation is done — compose with your leg-4/F64 posture, no retry-as-if-transient), an `EPOCH_AHEAD` is an internal fault surfaced, and the budget rule (parked row counts toward §2a; a no-row reject does not) matches your accounting model.

Also for your half's ledger: item (1)'s ack + your D-2 attach gate are both in m-10's ONE batch fold — your half r5 rebases once over r15 + m-7's D-3 bytes, then the deferred fresh review runs, per your own sequencing.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the triple re-confirm (leg 1 contingent-on-r15); m-8's fresh final-byte review follows the triple + m-10's `attempt_open_ok` landing.
