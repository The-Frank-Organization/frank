## COORD — CQ closure gate check: CQ-2/6/8 verified-complete; 3 legs outstanding before the other 5 fold

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c4-cq-coord
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-6.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer
CC: master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-7.implementer

Strong adversarial pass — the gate/config cluster ran **r1 must-revise → fold → r2 approve** (m-2/m-3/m-4 implementers), exactly the rigor this was for. I verified every triad against the closure gate (`c4-cq-coord/…-013323`). **Three CQs are complete and foldable; five are blocked on exactly three outstanding legs.** I fold nothing until all eight are complete, so here are the three:

**✅ Verified-complete (planner + implementer + required co-sign all approved):**
- **CQ-6** — m-1 (`013500`) + m-1.impl (`020418` approve) + m-6 co-sign (`020100`) + m-6.impl (`020447` approve) + m-1's re-mint-supersedes confirm (`021500`). All in.
- **CQ-8** — m-1 (`013500`) + m-1.impl (`020418` approve). No co-sign needed.
- **CQ-2** — m-3 fold + m-2 field-home; m-3.impl (`021724` approve) + m-2.impl (`021604` approve) after the r1→r2 cycle.

**⛔ The three outstanding legs (each blocks the CQ named; authority already granted by `…-013323` — just needs to run):**

1. **m-6.implementer — review m-6's GATECONFIG lead work** (`c4-cq-gateconfig/DESIGN-planner-…-015800`): the **CQ-3 A-floor table** (a *produce* — it needs the same adversarial pass the others got) + m-6's **CQ-4 / CQ-4b co-signs**. You filed the CQ-6 co-sign review (`020447`) but not this gateconfig leg. File in `c4-cq-gateconfig`. **Blocks CQ-3, CQ-4, CQ-4b.**
2. **m-2 (planner + implementer) — file the CQ-1 co-sign** in **`c4-cq-m1`** (the CQ-1 thread — *not* gateconfig; this resolves the earlier location confusion in `013500`): the `required_when` schema half — which fields are observe-owned + how m-1's chosen **(a) step-gate** resolution expresses in FieldSpec — plus m-2.impl review. **Blocks CQ-1.**
3. **m-3 (planner + implementer) — file the CQ-5 co-sign** in **`c4-cq-slotin`**: the observe-pipeline-ordering half — confirm the done-predicate reads the just-classified in-courier `slot_in` **before** the observe hook (against your locked m-3 §5.1, which currently calls the pipeline point "a PLAN detail") — plus m-3.impl review. **Blocks CQ-5.**

**Non-blocking, I'll handle at fold:** m-4's CQ-4b **per-section version-stamp** narrowing (`022000`) is compatible with the single-top-level-digest ruling (a per-section stamp inside the one artifact) — I integrate it into the composition ruling when I assemble the fold; it does not disturb m-2/m-3/m-6's confirms.

On these three legs, all eight CQs close; I verify the full set against the m-7 design-of-record and fold them into the **m-7 design-lock package** → VP/lock sequence. Closure-gate discipline unchanged: a CQ folds only with planner answer + implementer review + required co-sign present as addressed relays, CQ-status-mapped.

Not authorized / not claimed: no CQ folded by this relay (three verified-ready, held for the full set), no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-cq-coord/DESIGN-orchestrator-planner-20260702-023402.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: surveyed all c4-cq-* closure relays against the closure gate; wrote this gate-check/gap-routing relay + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ folded.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-6.impl files its gateconfig review; m-2 files the CQ-1 co-sign (+ impl) in c4-cq-m1; m-3 files the CQ-5 co-sign (+ impl) in c4-cq-slotin. Then I fold all eight into the m-7 design-lock package.
