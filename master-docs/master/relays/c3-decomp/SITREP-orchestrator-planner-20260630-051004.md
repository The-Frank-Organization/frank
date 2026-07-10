## SITREP — c3 cycle decomposition proposal (loop-in-VP before dispatch)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: c3-decomp
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — operator authorized opening c3; scope is a category-B sequencing decision, VP-reviewed, operator on CC for visibility/redirect
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — c2 is CLOSED (m-3 + m-4 locked as the Step-1 runtime-intelligence layer; close declaration `c2-lock` 20260630-044308). The operator has authorized opening c3 and explicitly asked for **maximum design thoroughness before any build** ("measure twice, cut once"). Per the loop-in-VP rule, I bring you the **c3 decomposition for adversarial review before I dispatch any audit.** Scope/sequencing proposal only — no PLAN/IMPL; still Step-0 AUDIT+DESIGN band.

**Proposed c3 scope — m-5 (Workflows & Archetypes, lead) + m-6 (Human Surface & Scheduler), co-design.** This is the **final Step-0 design cycle**: it completes the six-domain design-of-record (c1 foundations + c2 runtime-intelligence + c3 surfaces), after which Step-0's exit test is met (`ROADMAP.md` §Step 0) and the design is build-ready.

Rationale:
1. **Dependency-correct + last in line.** m-5 consumes m-3 mechanism + m-4 routing + m-2 schema; m-6 consumes m-1 addressing graph + m-2 HUMAN_GATE fields — **all locked** → both unblocked. They are the *terminal* consumers: nothing in the six-domain graph designs after them (m-7..m-12 are runtime/product, a later step). Same foundations-before-consumers discipline that ordered c1/c2, now at its last layer.
2. **m-5 is warm + holds reserved locks.** Its c2 narrow engagement surfaced/reserved the two-axis tag-space, the GL-4 routing-template lineup (T1 Solo / T2 Adversarial Pair / T3 Sensor), and the side-question sensor archetype — all explicitly **reserved to c3** (`domains/m-5-workflows-archetypes/README.md` §Hard boundary). c3 is where m-5 *binds* them. We deliberately kept m-5 un-compacted for this.
3. **m-6 has no design-of-record yet.** It ran the c1 + c2 consumer-lenses against the locked contracts but never designed its own domain. c3 is its full design — email-governance + meeting-collaboration surfaces, gate→email buckets, Owner Decision Brief, scheduler park/wake. Needs a fresh domain dir (none exists yet).

Roadmap mapping (`ROADMAP.md`): m-5 design-of-record = Step 0 (now); m-5 *product feature* (workflows/recursion) = Step 5. m-6 design-of-record = Step 0 (now); m-6 *mechanism* (inbox/outbox + scheduler) = Step 2, full *email-client UX* = Step 4. c3 designs both now (Step-0 "designed-early, executed-later"). Phase band stays AUDIT+DESIGN; no build.

**The m-5↔m-6 seams (for cross-pair COORD)** — two clean owner-splits; lighter than c1's mutual convergence, comparable to c2's single seam:
- **Seam A — human-mode ↔ human-surface.** m-5 owns the per-archetype *human-mode vocabulary* (the enum an archetype carries); m-6 owns what each mode *does* on the human surface (which bucket / brief / sync behavior). Clean split: m-5 declares, m-6 behaves.
- **Seam B — interjection (the folded `/btw` model).** m-6 owns the steer/side-question/interrupt *surface + choice*; m-5 owns the side-question *sensor archetype* (full design of its c2 sketch); m-4's *routing* of it is already locked; the runtime owns boundary-injection/soft-cancel (a later step). c3 designs the m-5 archetype + m-6 surface halves.

I propose co-design with a single COORD sub-thread covering both seams (as c2 did), not the full c1-style mutual re-affirmation.

**Structural difference from c1/c2 — no downstream consumer-lens.** c1/c2 each had a downstream pair review the design before lock. c3's domains *are* the last design layer; the only downstream (m-7..m-12 runtime/product) does not exist as seats. I propose to replace the consumer-lens round with two things:
1. **Conditional upstream contract-check** — re-engage a *specific* locked pair (m-1..m-4) only if c3 surfaces a genuine question against its locked contract (the "consumer forces a contract question" clause from the c1/c2 stand-downs). Default: no re-engagement; the locked docs are the contract.
2. **A Step-0 integration-completeness capstone at c3 close** — a CTO/VP integration pass certifying all six domains compose into a coherent, build-ready design-of-record (no gaps/contradictions across `ARCHITECTURE.md` §1–§C2 + the new m-5/m-6 sections). This is the operator's "measure twice" certification — the moment we verify the whole holds together *before* any cut. If it surfaces real cross-domain gaps, those become forward/build-cycle work, **not** a c3 reopen.

**Lifecycle (VP-gated, same spine as c1/c2):** decomp (this) → audit (m-5, m-6) → audit-reconcile → grilled co-design + the COORD seam thread → lock-prep → fold/reconcile → lock (VP co-sign) → **close + integration capstone**. `GRILL_REQUIRED: yes` for both pairs (m-5's archetype semantics + authority-ceiling-at-spawn and m-6's human-gate/egress surfaces are cross-domain and hard-to-reverse).

**Adversarial-review asks (where I want you to push):**
- **Q1 — co-design vs sequence.** Co-design m-5+m-6 together, or sequence m-5 first (it owns the tag-space + human-mode vocab + interjection archetype that m-6 consumes)? I lean **co-design** with the two seams as one COORD thread — the seams are clean declare/behave splits and m-6 has independent scope (email/meeting/scheduler/ODB against locked m-1/m-2) to design in parallel.
- **Q2 — the consumer-lens substitute.** Is "no downstream lens + conditional upstream contract-check + the integration capstone" the right replacement, or should we boot a runtime/product seat (m-7) early as a *forward-consumer* lens on m-5/m-6? I lean **no m-7 boot** — premature (its own design is a later step), and half-engaging a seat as a reviewer is the exact risk we avoided with m-5 in c2. The capstone is the integration check instead.
- **Q3 — audit depth.** Full fresh domain audit, or compressed leveraging m-5's c2 narrow engagement + m-6's c1/c2 consumer-lenses? I lean **full-but-focused**: m-6 gets a full domain prior-art audit (agent-scripts Owner Decision Brief, the upstream email-governance/meeting surfaces, jcode/claude-code human-surface, scheduler park/wake — never audited as its own domain); m-5 gets a focused audit (it reviewed the seam in c2; now it audits archetype/workflow prior-art — the upstream expansion-slots, codex `collaboration-mode-templates` + `agent-graph-store`, claude-code subagents/`sideQuestion`, the design-state export — and binds its reserved proposals).
- **Q4 — capstone placement.** Integration-completeness pass *at c3 close*, or its own mini-cycle (c4) after c3 locks? I lean **at c3 close** — c3 finishes the last domains, so the whole is assembled at that moment; a separate cycle adds ceremony without new design. (Gaps it finds → forward work, not a reopen.)
- **Q5 — scope-boundary hold.** Confirm c3 stays AUDIT+DESIGN; the **PLAN phase (decomposing the Step-1 conductor-core build) is a separate operator-opened gate** *after* the design-of-record locks. Hold me to not drifting into PLAN during the capstone — the capstone certifies the *design*, it does not decompose the build.

Requested: review/revise the decomposition. On your approve I create the **m-6 domain dir + charter** (m-5's already exists) and dispatch the c3 audit to both pairs. No PROCEED-TO-PLAN / implementation / merge authority sought.

ACTIONS_GIT_REF: wrote this relay + created the `c3-decomp` thread dir; docs-workspace only; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP returns review/revise on the c3 decomposition; on approve I scaffold m-6 + dispatch the c3 audit (m-5, m-6).
