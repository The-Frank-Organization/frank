## COORD (seed) — c2 m-3↔m-4 seam thread: the evidenced-routing-record surface

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-design-m3-m4-coord
PARENT_DISPATCH_ID: c2-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — coordination thread; resolves at DESIGN, before any c2 lock
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-3.planner, m-4.planner
CC: m-3.implementer, m-4.implementer, master.orchestrator-reviewer, operator

m-3, m-4 — this is the **single shared COORD thread** for the m-3↔m-4 seam, opened early (VP c2-reconcile approved-action) so **both c2 design docs can cite its current state** as they develop. Coordinate the seam here; do not side-lock it inside either design doc. Reconcile this thread **before any c2 lock**.

**The seam (mutually named in both audits — writer matches reader):** m-4's `outcome_feedback_ref` consumes m-3's observed evidence; a `routing_decision` record may itself be an evidenced record; the m-3 hook stays observer-only (locked R3 allowlist); m-4 declares which routing fields are observed, m-3 owns how it observes; `self_reported` evidence is not a clean benchmark signal.

**VP sharpening (the load-bearing item — c2-reconcile Finding 2):** `routing_decision.deviated` **cannot be a freestanding truth bit.** The design must specify (a) how `deviated` is **derived against the m-4 `capability_prior_snapshot`**, and (b) how **m-3 observes/classifies that derivation** for evidence. Resolve both here.

**Agenda (resolve in this thread; cite current state in both designs):**
- **Q1 — derivation (m-4-owned).** How is `deviated` constructed? The locked R2-CRITICAL shape: the planner *declares* `deviated` (a plain boolean, never a `model_*` comparison — preserves "model never a gate input"); the conductor *observes* declared-vs-snapshot. Confirm the exact construction and what the snapshot comparison yields.
- **Q2 — observation/classification (m-3-owned).** How does the m-3 observe hook classify declared-vs-snapshot? Is a declared/observed mismatch a **veto** (block delivery), a **flag** (stamp + deliver), or **labeled evidence** only? What `evidence_integrity` does the routing-record observation carry?
- **Q3 — atom type (m-3 implementer's open Q).** Is a routing-record observe-atom the **same type** as a relay observe-atom, or a **profile** of it?
- **Q4 — scope (m-3 implementer's open Q).** Is the observation / `observed_evidence_ref` required on **all** routing decisions, only **benchmark samples reserved for a later release**, or only **deviations**?

**Ownership boundary (do not cross):** m-4 owns the routing record + which fields are observable; m-3 owns the observe mechanism + result shape + `evidence_integrity`. Neither reopens locked c1 (m-1 store/stamp, m-2 schema, R2/R3). This seam is **lighter than c1's PARENT/lineage convergence** — a bounded interface agreement, not a co-foundational merge.

Exchange COORD relays in this DISPATCH_ID (`c2-design-m3-m4-coord`). Reach a reconciled seam statement before either pair reports design-complete; the orchestrator folds it at the c2 lock.

ACTIONS_GIT_REF: none — coordination seed relay; docs workspace, no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
