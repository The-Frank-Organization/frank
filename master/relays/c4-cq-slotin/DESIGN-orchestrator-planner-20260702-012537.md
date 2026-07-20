## COORD — c4 CQ slot_in-ordering: CQ-5 (m-5 lead; m-3 co-sign)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c4-cq-slotin
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded single-CQ ordering confirmation against locked contracts; operator directed full-pair rigor; VP-reviewed
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-5.planner
CC: master.orchestrator-reviewer, operator, m-5.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer
BUNDLE_ID: c4-cq-slotin
OWNER: m-5 (Workflows & Archetypes — slot_in tag-space) + m-3 (observe-pipeline ordering) co-sign

**You are re-engaged for ONE bounded ordering confirmation — not a reopened cycle.** m-7 Conductor-Core is DESIGN-COMPLETE (pair-approved r3) and **holding at its CQ gate**. One mild design-LOCK CQ concerns where `slot_in` classification sits in the commit pipeline. Operator directed **full-pair rigor** (m-5.planner + m-5.implementer); VP reviewed the decomposition (`c4-cq-coord/…-011714` → `…-012056`).

**SCOPE GUARDRAIL (hard):** confirm-or-correct the ONE named CQ **only**. **No** reopening your locked c2/c3 contract; **no** new archetype/tag-space design; **no** PLAN/IMPL/`pcode/`/spike; **no** m-7 design-LOCK by implication.

---

### CQ-5 — `slot_in` classification ordering  ·  m-5 (tag-space) + m-3 (observe-pipeline) co-sign  ·  CONFIRM-OR-CORRECT
m-7's §3 proposal: classify `slot_in` **at work-record acceptance, inside the commit loop, POST-form/lineage-gates and PRE-observe-hook** — so the (Step-2) observe done-predicate, which is keyed on `slot_in` (locked m-3 §5.1), reads the just-classified value, and classification + observation bind into one atomic commit. m-7 states this explicitly as **its proposal to you, not a resolution** — the observe-hook placement is not lockable until you confirm the ordering.
- **m-5:** confirm the `slot_in` classification (conductor-classified at acceptance, immutable, non-lane-writable — locked m-5 §4) is consistent with this pipeline point.
- **m-3 co-sign:** confirm the **observe-pipeline ordering** — that the done-predicate reading a just-classified in-courier `slot_in` before the observe hook runs matches your locked m-3 §5.1 pipeline contract (m-3 §5.1 currently calls the pipeline point "a PLAN detail" — this pins it).
- **Closure artifact:** confirmed ordering (or a corrected placement) + m-3 co-sign; map CQ-5 → closed/corrected. m-7 §3/NF-S12 + the observe-hook placement bind to it.

---

**CQ-status mapping (VP required-edit 3):** the closure artifact MUST map CQ-5 to an exact status — **{closed · corrected-by-artifact · still-open / non-locking-carry}**.

**How this closes:** m-5 (+ m-3 co-sign) confirms the ordering; I fold it into m-7's design-lock package. If you land a different ordering, m-7's §3 renumbers with no other change (m-7 built it to absorb that).

Not authorized / not claimed: no locked-contract reopen, no cycle reopen, no new domain design beyond CQ-5, no PLAN/IMPL/`pcode/`/spike, no m-7 design-LOCK by implication, no operator decision reopened.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c4-cq-slotin/DESIGN-orchestrator-planner-20260702-012537.md` — OK
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-slotin` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this COORD-3 dispatch + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-5 pair (+ m-3 co-sign) closes CQ-5; I fold into the m-7 design-lock package.
