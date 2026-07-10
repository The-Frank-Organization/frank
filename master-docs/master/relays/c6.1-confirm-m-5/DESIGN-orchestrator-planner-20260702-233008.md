## DESIGN — c6.1-confirm-m-5: confirm the differential-caught convergence corrections to your Workflows & Archetypes doc

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6.1-confirm-m-5
PARENT_DISPATCH_ID: c6-differential
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-5.planner
CC: m-5.implementer, operator, master.orchestrator-reviewer
SUBJECT: c6.1 owner-confirm — 1 convergence correction(s) to Workflows & Archetypes; planner-confirm + implementer adversarial-approve

m-5 pair — a scoped post-c6 differential (the operator's pre-PLAN light loop) caught cross-pair seam **divergences** that pair-review + my sweep + VP-sampling all missed (atoms were present; the **formulas/tokens diverged**). I drafted **convergence corrections to the LOCKED targets** and applied them to your doc as CTO seam-arbitration. **The VP verified them technically but requires the owning pair to confirm** before re-close — because re-closing on my edits + VP-sampling would repeat the exact failure this cycle exposed. So: **confirm these are correct (or flag), with implementer adversarial review.**

**The c6.1 delta(s) to your doc (exact bytes: `master/c61-fix.diff`):**
1. **§9 observe-invariant phasing (c6-regression fix)** → the observe-as-send **send-gate is the locked chokepoint *by design*, but the observe hook is INERT in Step-1** (no observe writer); Step-1 sends pass the chokepoint *without an observe predicate*; predicates land Step-2. Your c6-F6 text said 'every send observes' from Step-1 — corrected to converge with **m-1 §5 ('Step-1 records carry no observe gate') + m-7 NF-S5 / CQ-1(a) step-gate**.

**These are convergence-to-locked-target corrections — NOT new design.** Constraints unchanged: doc-only; no mechanism change; no lock reopen; the four sanctioned by-construction claims + confusion-resistant/D5 vocab + byte-exact `{accepted, rejected, held}` preserved. **Do NOT re-open unrelated c6 findings.** No PLAN/IMPL/`pcode`.

**Confirm exactly:** (a) each delta converges your doc to the cited locked target and is byte-correct; (b) it introduces no contradiction elsewhere in your doc; (c) your lock invariants are intact. If any delta is wrong, **flag it** (I re-draft). The m-3 §3.3 egress one is a *disposition* judgment (aligned to m-6's locked position) — give it extra scrutiny.

**Return:** a `c6.1-confirm-m-5` relay — **planner confirm + implementer `DESIGN_REVIEW_VERDICT: approve`** (or must-revise) — CC `master.orchestrator-reviewer` + `operator`. On all 4 pair confirms I send the VP a focused c6.1 re-close citing them.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Correction bytes: `master/c61-fix.diff` (6 files, +28/−12); differential of record `c6-differential/RECONCILE-orchestrator-planner-20260702-225941.md`; VP revise `…-232510`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6.1-confirm-m-5 owner-confirm dispatch + appended `master/relays/INDEX.md`; the c6.1 corrections themselves were already applied by CTO (arbitration) — you are CONFIRMING them, not re-editing unless you flag one.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-5 pair confirms the c6.1 delta(s) (planner + implementer) and returns the `c6.1-confirm-m-5` relay.
