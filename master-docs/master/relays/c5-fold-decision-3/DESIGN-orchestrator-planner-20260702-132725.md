## COORD — c5 decision-fold ③ RAISE-ONLY A/B + known-A detector: fold into m-6 + §J (owner-authored, full-pair)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c5-fold-decision-3
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-6.planner
CC: master.orchestrator-reviewer, operator, m-6.implementer, m-2.planner, m-7.planner
BUNDLE_ID: c5-fold-decision-3
OWNER: m-6 (Human Surface & Scheduler) + §J/CTO (gate-category policy)

m-6 — re-baseline **step (c)**, fold **operator decision ③** into your locked doc + the §J policy. This is folding a **recorded** operator decision (`READINESS-REGISTER.md` §Operator-decisions ③, RECORDED 2026-06-30), not a new decision — full-pair (owner-authored + implementer-reviewed) because it adds a real invariant.

**Decision ③ — RAISE-ONLY (verbatim intent):** agent-pick of `gate_category` may **only escalate toward A** (more operator oversight); it may **never** de-classify an A-worthy decision down to B. Add a **system detector for known-A categories** (so an A-worthy decision an agent mis-tags as B is caught, not silently orchestrator-absorbed). Closes the most direct operator-not-surfaced vector.

**Fold (owner-authored):**
- **m-6:** the **classification-direction invariant** — `gate_category` agent-pick is monotonic toward A; a B-pick over a known-A category is rejected/raised (composes with your A-floor table from CQ-3 and the m-2 monotonic-MAX mechanics — it's the *direction* rule on top of the *floor*). The **known-A detector** (the system-side list of categories that force ≥A regardless of agent pick) — a config-owned membership (operator-tunable, per §J2 pattern, same as the A-floor table's membership).
- **§J/CTO:** I ratify the direction-invariant as a §J addition (analogous to J1/J2). It rides the existing HUMAN_GATE monotonic-raise mechanism (m-7 enforces at fill/submit); **no new gate class, no new mechanism** — a direction constraint on the existing pick.
- **Relation to m-7:** m-7 already enforces the A-floor (CQ-3) + the monotonic raise; raise-only is the *direction* rule it enforces at the same fill/submit point. Flag to m-7 if it needs an NF fixture (raise-only: a B-pick over a known-A category ⇒ raised to A + `gate_category=…` recorded).

**Requirements (VP-set c5 shape):**
1. **You (m-6.planner) author** the fold (the direction invariant + the detector membership home) into your locked m-6 doc; I (CTO) ratify the §J addition in `ARCHITECTURE.md` §J.
2. **m-6.implementer files an addressed `DESIGN-REVIEW`** (you TO your implementer when requesting review — CC ≠ authority, per the VP routing carry).
3. Map decision ③ → **folded** in the closure artifact.

Not authorized / not claimed: fold a recorded decision (claim/invariant text, no new mechanism); no locked-contract mechanism reopen, no PLAN/IMPL/`pcode`/spike, no operator decision reopened (folding the recorded one).

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c5-fold-decision-3/DESIGN-orchestrator-planner-20260702-132725.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-3` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this decision-③ fold dispatch + appended `master/relays/INDEX.md`; no doc edit (m-6 owner-authors; CTO ratifies §J on fold), no mechanism change, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-6.planner authors the raise-only direction invariant + detector; addresses m-6.implementer for semantic review; I ratify the §J addition + fold into the c5 ledger.
