## Team m-6 — c2 consumer-lens review (Human Surface & Scheduler)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-consumer-review-m-6
PARENT_DISPATCH_ID: c2-lock-prep
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only consumer-review; surface operator items in findings
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-6.planner, m-6.implementer
CC: master.orchestrator-reviewer, m-3.planner, m-4.planner, operator
BUNDLE_ID: m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler)

Phase scope — AUDIT / consumer-review (read-only). m-6 (already booted; you ran the c1 consumer-review) is the **warm consumer lens** for c2. Review the two design-complete c2 designs from your human-surface vantage and verify they expose what m-6 needs — a writer-with-no-reader / reader-with-no-writer check before the c2 lock. Not in scope: edits/branches/commits/PRs/code; no PLAN/IMPL.

Pair method: m-6.planner leads; m-6.implementer challenges independently; reconcile into one m-6 consumer-review relay.

**What to review (your consumer surfaces on m-3 + m-4):**
- **From m-3 (Observation & Evidence):** (a) the **observe-gate veto → gate→email bucket** projection — can you mechanically project an observe-bounce into your gate→email buckets? (b) the **egress / content-safety gate** for the **away-mode external bridge** (the first external send, §J1) — does m-3 expose the fail-closed scan + result you need? (c) the **Owner Decision Brief** evidence-summary content schema (m-3 promotes the agent-scripts ODB) — does it carry the fields your ODB surface needs (`completed_proof` as an evidence_ref, residual risk, recommendation, choices)?
- **From m-4 (Routing & Policy):** the **`routing` gate_category** (category-B / orchestrator-absorbed; escalates to **A** only on `human_decision_required` / `routing_unavailable`) — can you bucket it? And the **ODB content** m-4 exposes for the rare A-escalation (recommendation + enumerated model choices) — sufficient for your Owner Decision Brief projection?
- **The §J forward requirements you'll own:** the operator-configurable `gate_category` A/B map + protected-branch set; the away-mode external-inbox bridge (egress-gated). Confirm the m-3/m-4 designs don't foreclose these (they are *your* later work, but flag any contract gap now).

**Verdict shape:** for each surface — does the design expose what m-6 consumes (reader has a writer), or is there a gap? Flag gaps NOW (cheap pre-lock; expensive after). This is the same discipline as c1's consumer-review that caught the writer/reader mismatches.

Sources (E1 — cite file:line):
- m-3 design: `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md`.
- m-4 design: `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md`.
- The locked c1 contract `master/ARCHITECTURE.md` §J (your gate→email buckets, away-mode bridge) — do not reopen.

Deliverable: a reconciled m-6 consumer-review relay (planner + implementer reconciled) under `c2-consumer-review-m-6` — the per-surface reader-has-a-writer verdict; any gap m-3/m-4 must close before the lock; operator-judgment items or none; E1 cites. No source changes, no PR. Include ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT (read-only; cwd is not a git repo, structured unavailable form expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
