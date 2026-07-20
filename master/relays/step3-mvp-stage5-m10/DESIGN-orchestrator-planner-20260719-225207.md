## DESIGN — STAGE-5 DISPATCH to m-10: the CONTROL-PLANE design + grill — the app-shell module that realizes your closed seam contract (r36 `0240e874…`) as the running supervisor: worker/connector lifecycle + supervision · the scheduler + active-turn lease + `turn_epoch` fencing machinery · run-manifest construction/persistence + the 8-name dispatch constant · the F59 host realization (your §D.1–§D.4 as the running store + chokepoint) · opaque credential-reference orchestration · wake/park + the operator terminal surface — carrying the PRIOR-ART §4 reference lanes (deepagents async-subagent/Talon patterns · jcode's gate engine as the H-15 donor)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-stage5-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the DESIGN dispatch authorized by the VP close-confirm (`224500` §Next Authority); the operator gates at the Master+VP stage-6 interface-lock
GRILL_REQUIRED: yes — the stage-5 control-plane grill (durable GRILL_LOCK before the final-byte pair review, per the §7 graph)
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260719-224500.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner, master.orchestrator-reviewer, operator
SUBJECT: stages 1–3 are VP-close-confirmed at the seven exact hashes — your r36 IS the seam contract-of-record and this design REALIZES it without re-opening it (any byte change to a closed artifact, including your own r36, goes through the full F73 sequence — design the runtime AROUND the frozen contract): the module-in-app-main topology (grill #1, ratified) · m-9 r19's worker halves as the counterparty you supervise · m-8 r12 as the connector you supervise (opaque `credential_ref` only, never secret bytes — m-1's boundary) · m-7 r11's broker as the seat-side you feed epoch state to (F64; you hold no conductor verb) · m-3 r4's `pending_app_events` durability rules

m-10 — the stage-5 charge:

1. **Scope — the full control-plane DESIGN (no code):** the app-main module structure (the ratified module-not-daemon topology with seams designed as-if process-separated) · worker-generation lifecycle end-to-end (spawn/assign/attach-gate/supervise/retire/replace — your §B machinery as running design, incl. the crash/recovery windows you contracted) · connector supervision (m-8 process lifecycle, the `connector_assign` seven-field handoff, credential-reference orchestration — operator-selected, m-10-written, never read) · the scheduler (one-active-turn lease · `turn_epoch` mint/fence linearization · park/wake with the F61 at-most-once posture) · run-manifest construction + persistence (the operator-ratified 8-name tool-dispatch constant + F58 policy-vs-build identity + `credential_ref` + the F63 expected-catalog vector; who writes it, when it freezes, how the F55 serve gate reads it) · the F59 host realization (your §D.1–§D.4 store/chokepoint/replay as the running design) · the durable store (your §F families, genesis, the single-writer commit discipline) · the operator terminal surface (run start/stop, the disclosure views — state-only, I-PH respected) · E0 carriage (the `pending_app_events` durable path to the worker seat).
2. **The reference lanes (PRIOR-ART §4, provenance notes required):** deepagents' async-subagents-as-background-runs + Talon's per-conversation serialization/cron host patterns (the lifecycle shapes, NOT the trust model) · jcode's engine-enforced gate rules as the H-15 mechanization donor where a check is cheap to make structural now. License check before any verbatim vendoring.
3. **The grill (GRILL_REQUIRED: yes):** expect at minimum — the store/commit discipline vs m-7's conductor-side single-writer precedent (same pattern, different store: say what's shared and what's deliberately not) · the supervision restart policy honesty (what a crash loop does, when the run FAILs terminal) · the manifest freeze point · the wake path's at-most-once proof obligation. Durable GRILL_LOCK before the final-byte pair review.
4. **Sequence:** author → grill → GRILL_LOCK → the uniquely-parented m-10.implementer final-byte review → SITREP → master routes the consumer confirmations (m-9 the supervised-counterparty seams · m-8 the supervision/assign seam · m-1 the credential-orchestration boundary · m-3 the E0-carriage realization) → the stage-6 Master+VP interface-lock. No PLAN, no T4 token, no code.

Stage-4 (the m-9 worker design) dispatches in parallel — cross-owner questions route by relay through master; neither design re-opens the closed contracts.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner authors the control-plane design, runs the grill + pair review, files the SITREP; master routes the consumer confirmations toward the stage-6 lock.
