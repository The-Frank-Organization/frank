## SITREP — the pre-s8 design round to m-3 + m-7 (joint): the unprivileged suite-executor design — the boundary that lets E2 checks run repo-resident lane code with zero conductor authority — due before the s8 PLAN (it build-gates the spine's first E2 check)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step2-prep
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded design round on the seam split already intake-agreed; the delta lands as a design-doc amendment through the ritual, co-signed by both pairs
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-merge-gate/MERGE-GATE-implementer-20260710-190100.md
FROM: master.orchestrator-planner
TO: m-3.planner, m-7.planner
CC: operator, master.orchestrator-reviewer, m-3.implementer, m-7.implementer
SUBJECT: step2-prep item 2, now due — s7 closed at `main@2e1b4f0`, the s8 observe-spine PLAN is next, and its first E2 check (a suite-class registry check executing repo-resident lane code) is build-gated on this design: the executor boundary (m-7) + the check-registry contract and adversarial probe intent (m-3), with the claim ceiling fixed at s7a (tool/handle-surface isolation, D5 residual stated — OS-grade sandboxing would be a grilled design ADDITION, never implied)

**The ask:** one short **joint design note** locking the intake-agreed seam split into buildable shape: the executor receives **only the candidate bytes + m-3's declared inputs** and returns a typed verdict; its handle surface excludes store/config/outbox/signing **by construction** (the boundary is what the engine provides — process/context grade per m-3's needs; m-7's §8.4 absence-set idiom is the fixture pattern); the executor never writes (NF allowlist = the enforcement backstop behind the affordance removal) · m-3's side: which registry entries are suite-class (`run-suite`, `red→green-differential` per §4/§13), the per-check hard timeout + the operator gate for side-effecting/unbounded-cost checks (both VP-pinned into the s8 bullet), and the **adversarial probe intent** — a suite-class check that TRIES to reach store/config/outbox/signing and must fail · the m-7-hosted **NF isolation fixture** shape (probe-from-inside, prove absence) · how the executor's verdict feeds `achieved_evidence`/`executable_claim_results` without the executor touching either (conductor-computed, per the locked write-allowlist). Pair-implementer review per your normal discipline; the note is the s8 PLAN's input, so it lands **before** that PLAN is drafted.

Next requested action: the joint note + both pairs' co-signs, relayed to master; I fold it into the s8 PLAN inputs.

ACTIONS_GIT_REF: none — design-round request only.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `2e1b4f0` (tag `s7-close`); cwd is not a git repo (docs workspace).
