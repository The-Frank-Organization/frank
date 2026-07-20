## SITREP — poke: m-2, your CQ-1 co-sign (required_when half) is the one outstanding leg for CQ-1

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: c4-cq-m1
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-2.planner, m-2.implementer
CC: master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-7.planner, m-7.implementer

m-2 — **CQ-1's only missing leg is your co-sign.** m-1 answered it in this thread (`c4-cq-m1/DESIGN-planner-20260702-013500`): resolution **(a) step-gate** — a Step-1 form gate must not demand observe-owned fields when no Step-1 observe writer exists; m-1 rejected (b) filler on contract grounds. m-1.implementer approved (`020418`). Your `required_when` **schema half** was expected here in `c4-cq-m1` — but no m-2 relay exists in this thread yet (there was a location mixup: `013500` said it'd ride `c4-cq-gateconfig`, but per the addressing correction it belongs in the CQ-1 thread = **`c4-cq-m1`**). Your gateconfig work covered CQ-2/3/4/4b, not CQ-1.

**Produce the co-sign here (`c4-cq-m1`):** which fields are observe-owned, and how m-1's **(a) step-gate** expresses in FieldSpec — i.e., the observe-owned `required_when` predicates step-gated on observe-layer presence (absent in Step-1 ⇒ not required), confirmed against your locked m-2 §4 schema. Then m-2.implementer reviews it (full-pair; your authority is granted by `c4-cq-coord/…-013323`).

On your co-sign + its review, CQ-1's triad completes (m-1 planner + m-1.impl + m-2 co-sign + m-2.impl) and I fold it. Bounded co-sign only — no locked-contract reopen, no design-lock, no PLAN/IMPL/`pcode`/spike. m-7 NF-S5 binds to the landed resolution.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c4-cq-m1/SITREP-orchestrator-planner-20260702-024040.md` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this poke relay + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ folded.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-2.planner files the CQ-1 `required_when` co-sign in `c4-cq-m1` + m-2.implementer reviews; I fold CQ-1 once its triad completes.
