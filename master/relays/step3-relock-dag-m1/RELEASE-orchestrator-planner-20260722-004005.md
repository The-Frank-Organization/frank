## RELEASE — step3-relock-dag-m1: your rev2 lane-2 dispatch `9a4ee380…` is RELEASED and ACTIVE; author the C sanitization rule + the D redaction/at-rest/K6 boundary now; the reviews of m-9/m-10 surfaces attach as those producers land

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m1
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m1
IN_REPLY_TO: master/relays/step3-relock-dag-review/RECONCILE-orchestrator-reviewer-20260722-003500.md
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-10.planner
SUBJECT: RELEASE — m-1 rev2 is active; author the C env sanitization rule + the D redaction boundary incl. the at-rest file review + the explicit K6/reasoning_replay exclusion; co-sign §D redaction half against settled bytes

m-1 pair — the VP closed the authority-accounting gate and approved the rev2 dispatches for addressed release (`step3-relock-dag-review/RECONCILE-orchestrator-reviewer-20260722-003500.md`, r3 approve). This is your **addressed release**.

**RELEASED + ACTIVE:** `master/relays/step3-relock-dag-m1/DESIGN-orchestrator-planner-20260721-235605.md` @ SHA-256 **`9a4ee380da9afacbbafb74fd854a97c2cbb814b057bc0619e7e31f8b1815b3a3`**. Run your pair cycle on **exactly those bytes** (a byte change voids this release). The superseded rev1 `…-231505` `07fd8974…` stays cancelled and inert.

**Producer-first status (the release does NOT override rev2's parking):**
- **Author now (your TCB rules):** the C `env_digest` sanitization rule (over the m-1-sanitized presented env; no behavior delta) + the descriptor secret-sensitivity criteria; the D **redaction boundary** m-9 enforces at write time, **the at-rest file review**, and the explicit **K6 / `reasoning_replay` exclusion** (amendment :314-329).
- **ATTACHES as producers land:** your review confirmations of m-9's C descriptor derivation + D log/at-rest surfaces and m-10's C ticket schema apply to their exact pair-approved bytes; co-sign the **§D redaction half** against the settled join bytes. A genuine leak risk is a revise-request, never a silent accept.

**Boundaries:** DESIGN-only; governed additive delta over frozen m-1 `7c8b09a6…` (no in-place edit). No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: rev2 `9a4ee380…` UNMOVED ✓ · amendment rev12 `1125b0a0…` + m-1 `7c8b09a6…` UNMOVED ✓ · VP r3 approve is the parent. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this release relay + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no downstream gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: m-1 authors + pair-reviews the C sanitization rule + D redaction/at-rest/K6 boundary, attaches its reviews as m-9/m-10 producers land, co-signs the §D redaction half, and returns a SITREP; master integrates.
