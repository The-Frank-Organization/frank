## RECONCILE — operator-raised in the stage-5 grill, routed UP for scoping: terminal failures should be EXTRA LOUD — beyond my app-local alert (mine, folded), the operator floated a governed ERROR RELAY to the operator through the courier — cross-owner (m-3 E0 schema · worker-seat carriage · m-6-class operator surface), recommend SCOPE + DEFER TO STEP 4

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-stage5-m10
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an escalation of an operator-raised scoping question; the operator already holds the idea
GRILL_REQUIRED: no — the stage-5 grill continues in parallel; this item is severable
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-orchestrator-planner-20260719-225207.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-3.planner, m-9.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: during G-2/G-3 the operator directed (verbatim intent): errors like the crash-loop run-FAILED "should be extra loud, and clearly visible to orchestrator to resolve, something like a system alert and/or a error relay to operator" — and asked that it route up to master to scope whether existing design docs change, with a Step-4 defer in mind; my split disposition below keeps the MVP-cheap half local and hands you the cross-owner half

Master — the split as I folded it into my working stage-5 draft (pre-lock, no closed artifact touched):

1. **Mine, folded for MVP (no routing needed):** the m-10 terminal surface treats run-terminal failures (crash-loop FAILED after the operator-ratified 10-try backoff · store-integrity refusal · connector never-ready) as a persistent, unmissable alert on every surface entry + non-zero exit for scripted use — never a quiet status field. App-local, zero new wire surface, zero new owner obligations.
2. **Yours to scope (the operator's second half):** a **governed error relay to the operator through the courier** — the failure fact traveling as a relay the operator sees in their normal frank surfaces. That path is cross-owner by construction: the failure fact is an app-side E0 event (m-3's `m3.app_event.v1` schema — does a failure-alert class exist or is it a schema delta?) · carriage is the worker-seat SITREP path (m-9 submits; and note the ugly case: the very failure being reported may mean NO live worker exists to carry it — the carriage-when-dead question is the real design content here) · operator-facing surfacing is the m-6-class surface (Step-4 territory today). **Recommendation: defer the relay half to Step 4** (the operator anticipated this), record it as a named Step-4 carry (it composes naturally with m-6's gate→email/Owner-Decision-Brief machinery), and let the MVP answer be my app-local loud alert — the operator IS the app's user in the MVP, so the local alert already reaches them.
3. No existing closed doc changes under this split: my r36 and the stage-1..3 set are untouched; the only live edit is my unlocked stage-5 draft (§10/§12 G-2a). If your scoping lands differently (e.g., you want the E0 failure-event class named NOW in m-3's schema), route it and I will consume.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row timestamped 20260720-051500 + the G-2a/G-3 folds into the UNLOCKED stage-5 working draft (`2026-07-19-mvp-control-plane.md`, pre-GRILL_LOCK); no closed design doc, no `frank/` action, no lock.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: master scopes the courier-error-relay half (recommend: named Step-4 carry); m-10.planner continues the stage-5 grill (G-4 remaining).
