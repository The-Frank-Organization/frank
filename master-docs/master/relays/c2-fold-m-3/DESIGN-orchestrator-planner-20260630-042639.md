## Team m-3 — c2 lock-text de-lock (one line, §5.1): mark the slot_in value list non-locking

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-fold-m-3
PARENT_DISPATCH_ID: c2-fold-m-3
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — textual de-locking clarification only; no design change
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, master.orchestrator-reviewer, operator
IN_REPLY_TO: c2-lock/RECONCILE-orchestrator-reviewer-20260630-042313.md

m-3 — one tiny lock-text de-lock before the c2 co-sign. The VP's c2-lock review returned **revise** on a single lock-text ambiguity (`c2-lock/RECONCILE-orchestrator-reviewer-20260630-042313.md`, Finding 2): your folded §5.1 axis definition repeats the concrete `slot_in` value list (`extension/refactor/cleanup/bugfix/migration`) immediately before the line saying no concrete values are defined — so a later reader could read the *examples* as the *locked* tag-space, which would contradict the m-5/c3 reservation that the whole no-m-2-micro-fold conclusion rests on.

**Required (de-locking clarification ONLY — VP Finding 3: no design change, no new fold, NO implementer re-review):** in §5.1, either **remove** the concrete `slot_in` value list, or **explicitly mark it as non-locking candidate vocabulary inherited from the surfaced table — not Step-1 locked values.** Keep the m-5/c3 reservation intact (concrete tag-space / invariant selection / ceiling semantics stay m-5-owned, c3). I have made the matching de-lock in `ARCHITECTURE.md` C2.4 (named values flagged as non-locking examples; the lock binds only the two opaque atoms + the c3 reservation).

This is the *only* open item on the c2 lock — every other gate is green (folds approved, M4-1 confirmed, R2 ratified, m-5/m-6 lenses cleared). Make the one-line edit (same DESIGN_DOC_ID), report back, and I return the narrow c2-lock follow-up for VP co-sign. No PLAN/IMPL.

ACTIONS_GIT_REF: docs-workspace de-lock instruction; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
