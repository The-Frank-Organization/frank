## DESIGN — c6.1-confirm-m-3: confirm the differential-caught convergence corrections to your Observation & Evidence doc

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6.1-confirm-m-3
PARENT_DISPATCH_ID: c6-differential
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, operator, master.orchestrator-reviewer
SUBJECT: c6.1 owner-confirm — 2 convergence correction(s) to Observation & Evidence; planner-confirm + implementer adversarial-approve

m-3 pair — a scoped post-c6 differential (the operator's pre-PLAN light loop) caught cross-pair seam **divergences** that pair-review + my sweep + VP-sampling all missed (atoms were present; the **formulas/tokens diverged**). I drafted **convergence corrections to the LOCKED targets** and applied them to your doc as CTO seam-arbitration. **The VP verified them technically but requires the owning pair to confirm** before re-close — because re-closing on my edits + VP-sampling would repeat the exact failure this cycle exposed. So: **confirm these are correct (or flag), with implementer adversarial review.**

**The c6.1 delta(s) to your doc (exact bytes: `master/c61-fix.diff`):**
1. **§9 `deviated_observed` formula (c6-regression fix)** → **`declared_bucket ≠ rank-1(recommended)`** (deviation-from-floor), with **`bucket_binding_observed := chosen_model ∈ members(declared_bucket)`** as a *separate boolean* binding check. Your c6 fold wrote `chosen_bucket ≠ declared_bucket` (the binding check, not the deviation) — now **byte-identical to m-4 §2/§9 (GL-1) + m-2 §17.3**. Fold-log §241 records the correction.
2. **§3.3 egress row (c6-regression fix)** → an outbound `egress_scan_result=blocked` is the **non-terminal `egress_blocked` park + resummon** (m-6 §4 / m-7 NF-S9), **not** terminal `held`. The acceptance-time content-safety egress veto stays terminal `rejected` (your own §3.2). Converges to **m-6 §2:46's locked 'egress-block ≠ terminal token'** + your §3.2.

**These are convergence-to-locked-target corrections — NOT new design.** Constraints unchanged: doc-only; no mechanism change; no lock reopen; the four sanctioned by-construction claims + confusion-resistant/D5 vocab + byte-exact `{accepted, rejected, held}` preserved. **Do NOT re-open unrelated c6 findings.** No PLAN/IMPL/`pcode`.

**Confirm exactly:** (a) each delta converges your doc to the cited locked target and is byte-correct; (b) it introduces no contradiction elsewhere in your doc; (c) your lock invariants are intact. If any delta is wrong, **flag it** (I re-draft). The m-3 §3.3 egress one is a *disposition* judgment (aligned to m-6's locked position) — give it extra scrutiny.

**Return:** a `c6.1-confirm-m-3` relay — **planner confirm + implementer `DESIGN_REVIEW_VERDICT: approve`** (or must-revise) — CC `master.orchestrator-reviewer` + `operator`. On all 4 pair confirms I send the VP a focused c6.1 re-close citing them.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Correction bytes: `master/c61-fix.diff` (6 files, +28/−12); differential of record `c6-differential/RECONCILE-orchestrator-planner-20260702-225941.md`; VP revise `…-232510`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6.1-confirm-m-3 owner-confirm dispatch + appended `master/relays/INDEX.md`; the c6.1 corrections themselves were already applied by CTO (arbitration) — you are CONFIRMING them, not re-editing unless you flag one.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-3 pair confirms the c6.1 delta(s) (planner + implementer) and returns the `c6.1-confirm-m-3` relay.
