## RECONCILE — revise accepted: m-4 added to COORD-1 (6 pairs), CQ-2 scoped to decision-②, CQ-status mapping required

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c4-cq-coord
PARENT_DISPATCH_ID: c4-design-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-4.planner, m-4.implementer, m-5.planner, m-5.implementer, m-6.planner, m-6.implementer

Partner — revise **accepted**; both findings are correct catches. Folding all three required edits before I write the dispatches:

**Finding 1 — m-4 added (I take your preferred path, not the scope-out).** CQ-4b's config artifact composes m-4-authored capability-prior / routing-policy config (m-7 design-of-record §7:106-109). Under the operator's full-pair rigor, that surface cannot close with m-4 absent. I will **not** claim a clean read-only path: the composition decision (one unified artifact + shared digest vs per-domain files) can change how m-4's capability-prior config is loaded, integrity-checked, and versioned — that touches m-4's deployment surface, not just its container. So:
- **COORD-1 owner set = m-2 + m-3 + m-4 + m-6 full pairs + CTO.** Decomposition updated: **6 stood-down pairs / ~12 sessions** (was 5 / ~10).
- **m-4's role is scoped to CQ-4b only** (it has no other open CQ): confirm the composition + load contract preserves its locked capability-prior config's assumptions — a bounded confirm, not a re-author. m-4 stands back down after.

**Finding 2 — CQ-2 wording corrected.** COORD-1 closes the **decision-② subset** of re-baseline step (c), **NOT** the full "fold the five decisions." The other four decisions' folds are separate step-(c) work outside COORD-1's scope; COORD-1 neither scopes nor verifies all five. The dispatch will say exactly this.

**Q1 fold — lead-plus-co-sign made explicit.** Each joint CQ's closure artifact carries the co-signer's confirmation: **m-2 co-signs CQ-1** (its `required_when` half), **m-6 co-signs CQ-6** (the away-token edge), **m-3 co-signs CQ-5** (the observe-pipeline-ordering half). No joint CQ closes on a single seat's say-so.

**Required-edit 3 — CQ-status mapping baked into every COORD's acceptance criteria.** Each closure artifact MUST map every named CQ to an exact status: **{closed · corrected-by-artifact · still-open / non-locking-carry}** — no CQ left implicit.

**Q4 — scope-hold confirmed.** Each COORD is confirm-or-produce for its named CQ rows only: no c1/c2/c3 cycle reopen, no locked-contract reopen, no PLAN, no implementation, no `pcode/`, no spike, no design-lock by implication. Re-engaged pairs answer the bounded question and stand back down (conditional-upstream-contract-check at scale, not a reopened cycle).

**Corrected decomposition (what I will now write):**
- **COORD-1 `c4-cq-gateconfig`** — CQ-2 (decision-② subset of step (c)), CQ-3 (A-floor table), CQ-4 (terminal tokens), CQ-4b (config composition) → **m-2 + m-3 + m-4 + m-6** full pairs + CTO (CQ-4b ruling draft supplied by me).
- **COORD-2 `c4-cq-m1`** — CQ-1 (m-2 co-sign), CQ-6 (m-6 co-sign), CQ-8 → **m-1** full pair.
- **COORD-3 `c4-cq-slotin`** — CQ-5 (m-3 co-sign) → **m-5** full pair.

No re-review needed — I'm executing your prescribed edits, not re-clustering; you're CC'd on the three dispatches so the fixes are visible as they land. If I've mis-taken any fix, flag it and I hold.

Not authorized / not claimed: no CQ resolved, no pair re-engaged yet (dispatches follow), no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-012056.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read the VP revise `c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-011714`; wrote this fold-acceptance relay + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved, no pair re-engaged.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: I write the 3 COORD dispatches (COORD-1 now m-2/m-3/m-4/m-6 + CTO) + the CQ-4b config-composition ruling draft, with these folds baked in; the operator re-engages the 6 pairs.
