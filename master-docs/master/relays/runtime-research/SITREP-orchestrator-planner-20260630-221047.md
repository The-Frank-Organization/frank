## SITREP — operator decisions ②–⑤ RECORDED: all five readiness decisions now closed

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: runtime-research
PARENT_DISPATCH_ID: runtime-research
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — records decisions the operator already made; does not itself gate
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — the operator has made the remaining four readiness decisions. **All five (Cluster 2 / 5a / 6a / 6b / 6c) are
now RECORDED** in `master/READINESS-REGISTER.md` → *Operator decisions — recorded log*. Each went to the stricter/safer
posture:

- **② 5a — FAIL-CLOSED for merge/authority (A-gate) classes.** An unobservable authority-class record is held/escalated,
  not delivered on `self_reported`. → bounded **m-3** fix.
- **③ 6a — RAISE-ONLY.** Agent-pick may only escalate `gate_category` toward A, never de-classify down to B; add a
  known-A detector. → bounded **CTO/§J + m-6** fix.
- **④ 6b — ROTATE + RE-OBSERVE.** Away-token refresh rotates `decision_id`/burns prior nonces; `verify` re-observes and
  bounces on changed state. → bounded **m-1 + m-6** fix.
- **⑤ 6c — YES, narrow ODB→operator exemption.** The model-name field in a conductor-generated operator-facing ODB is
  exempt from the *confidentiality* egress scan only; **R2 (model ≠ gate input) untouched**; general egress rule intact;
  away-bridge transport still gated by the away opt-in. → bounded **m-3 + m-6 + m-4** fix.

**This clears item (3) of your sequencing gate** ("operator decisions recorded for Cluster 2, 5a, 6a, 6b, 6c"). The
gate's two remaining MUST items stand: **Cluster 1** (m-1/m-2 `submit()` write-path) and **Cluster 4a/4b** (m-2 stale
schema + R2 hole) — both unreconciled. **No Step-1 PLAN opens** until those are reconciled.

**Nothing routed, nothing built.** The five decisions each imply a bounded owner-fix; I have **not** opened any of them —
they route only as bounded VP-gated fix relays on the operator's go. No pair re-engaged, no contract changed, no code,
no spike. Docs-only, cwd.

Next: the operator directs whether to route the two MUST-reconcile fixes (Cluster 1 → CTO+VP+m-1/m-2; Cluster 4a/4b →
m-2+m-4) and the bounded owner-fixes, or to pause. I route nothing without the operator's go.

ACTIONS_GIT_REF: recorded decisions ②–⑤ in `master/READINESS-REGISTER.md`; wrote this relay + appended `master/relays/INDEX.md`; no code/source/pcode edits, no design-doc changes, no PLAN, no spike, no pair re-engaged.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: operator directs the next step (route Cluster 1 + 4a/4b MUST-reconcile fixes + the bounded owner-fixes · or pause); VP has visibility on all five recorded decisions.
