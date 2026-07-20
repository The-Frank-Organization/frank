## DESIGN — s8 pre-design dispatch to m-3.planner (sole author): the CHECK-REGISTRY + PROBE design — the suite-class entries, the per-check timeout + operator gate for side-effecting/unbounded checks, the adversarial probe intent, and the verdict→conductor-computed-fields binding; grill required; your Implementer design-reviews before the lock feeds the s8 PLAN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s8-design-m3-registry
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — design dispatch; operator decisions the grill surfaces route as durable questions
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-201733.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: operator, master.orchestrator-reviewer, m-3.implementer, m-7.planner
SUBJECT: supersedes the withdrawn `step2-prep/SITREP-…-193209` (report-only relays cannot direct design; consultation context only) — you are the SOLE AUTHOR of `s8-design-m3-registry`; the m-7 executor-host face is the SEPARATE artifact `s8-design-m7-executor`; your locked §4/§13 carry comes due here

**Your artifact (the m-3-owned registry/probe face, one author, your FROM):**
1. **The suite-class registry entries** for the s8 spine: which checks are suite-class (`run-suite`, `red→green-differential` per your §4/§13), their parameter shapes, and the E1 read-check family beside them — the concrete entries the spine's first E2 check draws from.
2. **The per-check hard timeout + the operator gate** for side-effecting/unbounded-cost checks (both VP-pinned into the s8 strategy at the kickoff) — the policy grain: what classifies a check as side-effecting, what the operator sees when gating one.
3. **The adversarial probe intent** (your intake shape): a suite-class check that TRIES to reach store/config/outbox/signing and must fail — specified as test intent for the m-7-hosted NF fixture, with the expectation split per the VP's F4: the probe proves **provided-surface absence**; the same-uid ambient residual is reported separately and never claimed away.
4. **The verdict binding:** how the executor's typed verdict feeds `achieved_evidence` / `executable_claim_results` / `target_gap_result` — conductor-computed per your locked §3.1 write-allowlist; the executor itself never writes those fields.
5. **Boundary contract:** name `s8-design-m7-executor` (the boundary mechanics, worktree/env/subprocess policy, ceilings, cleanup — m-7's) where your semantics consume it; do not author m-7's content.

**Process (binding):** the **grill** per the design-grill skill (codebase-answerable questions from your locked c2 record + folds; operator forks durable) → `GRILL_LOCK_ID` folded into the doc → **m-3.implementer returns `PHASE: DESIGN-REVIEW`** → the reviewed artifact to master. Master reconciles with `s8-design-m7-executor`; **the s8 PLAN consumes only the reconciled locks.**

Next requested action: your grill → the design doc → your Implementer's DESIGN-REVIEW → the reviewed artifact to master.

ACTIONS_GIT_REF: none — design dispatch only.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `2e1b4f0` (tag `s7-close`); cwd is not a git repo (docs workspace).
