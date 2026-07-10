## SITREP — c2 cycle decomposition proposal (loop-in-VP before dispatch)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: c2-decomp
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — operator authorized opening the next cycle; c2 scope is a category-B sequencing decision, VP-reviewed, operator on CC for visibility/redirect
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — c1 is CLOSED (m-1 + m-2 locked as the Step-1 design-of-record; close declaration `c1-joint-lock` 20260629-181713). The operator has authorized opening the next cycle. Per the playbook rule (loop you in before executing any authority-bearing orchestrator decision), I am bringing you the **c2 decomposition for adversarial review before I dispatch any audit.** This is a scope/sequencing proposal only — no PLAN/IMPL, still Step-0 AUDIT+DESIGN band.

**Proposed c2 scope — m-3 (Observation & Evidence) + m-4 (Routing & Policy), co-design.** The runtime-intelligence layer directly atop the locked c1 substrate.

Rationale:
1. **Dependency-correct.** Both consume *only* the now-locked m-1/m-2 contract (m-3: m-2 schema + m-1 isolation; m-4: m-2 schema + m-1 identity) → unblocked. Their consumers (m-5 = m-3 mechanism + m-4 routing; m-6 buckets = m-4 routing categories + m-3 gate) must design *after* them — same foundations-before-consumers discipline that ordered c1.
2. **m-4 is the headline.** The adaptive-routing pillar (altitude B, capability priors, 3-staged policy + justified deviation, the benchmark loop with feedback in a later release) is the conductor's thesis. Its *record contract* is already half-specified — R2 in `ARCHITECTURE.md` (routing = a separate seat-stamped relay; model never a gate input). c2 designs the full router/policy on top of R2.
3. **m-3 completes a c1 stub.** R3 observe-integrity + `evidence_integrity` were specified in c1 as the seam; c2 designs the full observe-as-send-gate, per-phase done-predicates, evidence ladder, executable claims, and the egress/content-safety gate.
4. **Low setup cost.** Both pairs are warm — booted (20260629-155401) and they ran the c1 consumer-review, so they hold the locked contract in context.

Roadmap mapping (`ROADMAP.md`): m-3 → Step 2 (governance hardening); m-4 → record contract designed in Step 1 (R2, done) + router execution at Step 3. c2 designs both now (Step-0 "designed-early, executed-later"). Phase band stays AUDIT+DESIGN; no build.

The m-3↔m-4 seam (for cross-pair COORD): the routing benchmark loop (feedback in a later release) *consumes m-3 observed evidence* (routing quality measured against observed outcomes), and a routing record may itself be an evidenced record. **Lighter than c1's PARENT-strengthens-lineage** — these are siblings on a shared foundation with one explicit seam, not a mutual load-bearing convergence. I propose co-design with a single COORD sub-thread on that seam, not the full c1-style mutual re-affirmation.

After c2: **c3 = m-5 + m-6** (the surfaces consuming m-3/m-4), then runtime/product/UI (m-7..m-12).

**Adversarial-review asks (where I want you to push):**
- Q1 — co-design m-3+m-4 together, or sequence m-3 first (it is marginally more foundational — m-4's benchmark loop consumes m-3 evidence)? I lean co-design; the seam is light enough.
- Q2 — c2 consumer lens: boot m-5 now to review as the primary downstream consumer of both, or use the warm m-6 as the lens and defer m-5 to c3? I lean **m-6 as lens + defer m-5** (m-5 is unbooted and its own design is c3; pulling it in only as a reviewer risks half-engaging it).
- Q3 — audit depth: full fresh domain audit (the agentic dev team protocol + jcode/claude-code/agent-scripts + the frank adaptive-routing-pillar export), or a compressed audit that leverages the c1 consumer-review the pairs already did? I lean full-but-focused (the consumer-review audited the *foundation*, not their own domain's prior art).
- Q4 — is m-3+m-4 the right c2 at all, or should m-6 (human surface — where the §J forward requirements concentrated: gate_category config, away-mode bridge) precede m-4? I argue no: m-6's buckets consume m-4's routing categories, so m-4 should lock first.

Requested: review/revise the decomposition. On your approve I will create the m-3/m-4 domain dirs + charters and dispatch the c2 audit. No PROCEED-TO-PLAN / implementation / merge authority sought.

ACTIONS_GIT_REF: wrote this relay; the c1 close artifacts (ARCHITECTURE.md status, README, RECONCILE ledger, CYCLE-PLAYBOOK.md, charter pointer) are committed to the docs workspace; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP returns review/revise on the c2 decomposition; on approve I scaffold m-3/m-4 and dispatch the c2 audit.
