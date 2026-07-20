## SITREP — build-readiness review decomposition (loop-in-VP before running the lenses)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: readiness-decomp
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — operator chose a design-validation pass before any build; scope/sequencing proposal, VP-reviewed, operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — c3 is CLOSED; the **six-domain Step-0 design-of-record is complete** (`c3-lock` co-sign 191315 / close 191525). The operator chose to **validate the design-of-record before committing to Step-1** (not to open the build yet). Per loop-in-VP, here's the decomposition for adversarial review **before I run it.**

**Proposed: a bounded, read-only `readiness-review` — a design-of-record BUILD-READINESS review.** NOT a new design cycle, NOT a build. It stays in the **AUDIT band** (a read-only audit of the locked design), locks nothing, opens no PLAN/IMPL. Scope reviewed: `ARCHITECTURE.md` §1–§C3 + the six per-domain design docs + the relay trail.

**Method.** I run **fresh read-only adversarial reviewer lenses** (the sanctioned subagent spawn — fresh eyes that did not design it), each a distinct dimension, then reconcile into a **build-readiness register** + a go/no-go verdict. A lens finding a real **design gap** routes to the owning pair as a **bounded fix** (conditional re-engage, VP-gated — the c3 upstream-check pattern); the default output is "build-ready + here is the risk register for the build."

**The lenses (dimensions):**
1. **Composition & seams (end-to-end trace).** Trace a real relay through the whole pipeline (submit → channel-stamp FROM → form-validate → observe-as-send → egress → route/record → archetype gate-set → human-surface bucket/ODB). Any unstated data-flow gap; any field a consumer reads that no writer provides; any seam closed-on-paper-but-holed.
2. **Build-risk ranking.** The hardest/riskiest mechanisms to implement: store-append **atomicity**; observe-as-send **TOCTOU** (the check→edit→send window); the inbound-token **nonce-burn**; the fail-closed **egress chokepoint**; **ceiling enforcement** (best-effort host vs conductor-uniform); durable **park/wake**. Rank by risk; flag spike candidates.
3. **Unstated assumptions / unknowns.** What the design assumes but doesn't validate — especially "the conductor holds an **outside-the-lane read handle**" on Step-1 ride-existing-runtimes (the DI-5 honesty floor); anything hand-waved.
4. **Step-1 slice readiness.** Can the *specific* Step-1 deliverable (sole-writer store + channel-stamped FROM + inline lint/form gate + a local governance outbox, riding existing runtimes) be built from what's specified? What's under-specified for THAT slice.
5. **Adversary / subversion.** It's a governance/TCB system — where can identity / authority / observe be subverted; is "forgery-robust by construction" actually airtight in the design (not just asserted)?

**Deliverable.** A **build-readiness register** — each item: severity · domain(s) · **design-gap vs build-note** · recommended action {fix-in-design / spike / accept-as-build-task} — plus a **go/no-go verdict** and a recommendation (proceed to Step-1 / fix specific gaps first / spike specific pieces). The spike decision itself (a spike = the first code, a scope-step) is **surfaced to the operator, not taken here.**

**Bound.** Read-only; no new design lock; the charter's design-only scope is preserved (this is AUDIT-band review, not build). A real design gap is a bounded owning-pair fix (VP-gated), never a reopen. No m-1..m-4 (or m-5/m-6) contract change except via such a bounded, reviewed fix.

**Adversarial-review asks (where I want you to push):**
- **Q1 — the lens set.** Is the 5-lens set right, or is a dimension missing/redundant? Candidates I considered: a dedicated **operator-surface / human-in-the-loop** lens (m-6-heavy), and a **versioning/migration** lens (given the day-one public-release + zero-migrator intent). I lean folding those into lenses 1/4 rather than adding lenses — push if you disagree.
- **Q2 — fresh lenses vs pair re-engage.** Run fresh red-team subagent lenses (orchestrator-run), or re-engage the six pairs to stress their own domains? I lean **fresh lenses** (they didn't design it → less blind) + **conditional pair re-engage only on a found gap**. Concur?
- **Q3 — the bound.** Agree this stays **AUDIT-band read-only** (no design lock, no code), and a **spike is a separate operator-opened step** if a lens flags one?

Requested: review/revise the readiness-review decomposition. On your approve I run the lenses, reconcile the register, and bring it to you (+ the operator) with the go/no-go. No PLAN / build / spike authority sought here.

ACTIONS_GIT_REF: wrote this relay; docs-workspace only; no code/source/pcode edits; cwd is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP returns review/revise on the readiness-review decomposition; on approve I run the read-only lenses + reconcile the build-readiness register + bring the go/no-go to the VP + operator.
