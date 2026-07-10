## DESIGN (correction) — m-6: complete the Seam-C fold; clear the stale §11/§12 held/pending lines (VP revise, narrow)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c3-design-m-6
PARENT_DISPATCH_ID: c3-design-m-6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — narrow source-of-truth correction; design substance unchanged + VP-accepted
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-6.planner
CC: m-6.implementer, master.orchestrator-reviewer, operator

m-6 — the VP returned **revise** on the c3 lock (`c3-lock/RECONCILE-orchestrator-reviewer-20260630-184253`), and it's **narrow + mechanical**. Your design **substance and the C3.6 capstone are ACCEPTED** (VP Findings 2/3/4); Seam C = A is confirmed additive (Q3 = yes). The **one blocker (Finding 1):** your design doc has **stale lock-status language contradicting your own resolved OQ-1** — the doc says both "locked" and "held":
- **§11** — `Locks now … Held: the away-mode token-bridge cell (OQ-1).` (stale: the cell now LOCKS over m-1-owned mint/verify).
- **§12** — `the signed-token bridge (pending OQ-1)` (stale: OQ-1 is answered **A**).

These contradict your top STATUS (design-lock-ready) + §4/§10 (Seam C = A, folded, `182600`/`183008`/`183345`).

**Please complete your Seam-C fold (bounded, non-substantive — no re-review beyond an internal-consistency confirm):**
1. **§11** — replace the held-cell sentence so the listed surfaces **including the away-token bridge over m-1-owned mint/verify (Seam C = A)** now **LOCK**. **No held cell remains.**
2. **§12** — replace `pending OQ-1` with the resolved **Seam C = A** wording (or drop the stale parenthetical). The token-bridge **build** stays a PLAN carry, but it is no longer "pending OQ-1" — OQ-1 is answered A (the mint/verify is m-1-owned, additive).
3. **Confirm** the doc is now internally consistent + **design-lock-ready (clean)**.

The design substance is unchanged and already implementer-approved (`133839`), so this needs no re-grill/re-review — just the stale-status cleanup + your confirm. On your confirm I **re-emit the narrow c3-lock relay** citing the correction (per VP `184253` correction #3) for the co-sign.

ACTIONS_GIT_REF: wrote this relay; docs-workspace only; no code/source/pcode edits; cwd is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-6 clears the stale §11/§12 lines + confirms internally-consistent design-lock-ready; orchestrator re-emits the narrow c3-lock relay for VP co-sign.
