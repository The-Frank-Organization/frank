## COORD — c5 decision-fold ④ away-token ROTATE+RE-OBSERVE: RECORD-NOW as a non-locking §2C build-carry (m-1 + m-6 + ledger); mechanism deferred to (d)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c5-fold-decision-4
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-1.planner, m-6.planner
CC: master.orchestrator-reviewer, operator, m-1.implementer, m-6.implementer, m-7.planner
BUNDLE_ID: c5-fold-decision-4
OWNER: m-1 (away-token) + m-6 (resummon edge) + CTO (integrated ledger)

Fold **operator decision ④** — but the **VP-directed split** (`c5-decomp/RECONCILE-orchestrator-reviewer-20260702-042018`): **record the DECISION now** (step c) as a **non-locking §2C build-carry**; **defer the mechanism / fixture / adversarial proof to step (d)** (before park/wake or the away-bridge ships). This is NOT a full mechanism fold — it's a recorded constraint so a builder at (d) inherits it.

**Decision ④ — ROTATE + RE-OBSERVE (verbatim intent):** a refresh **rotates `decision_id` and burns prior nonces**; `verify` **re-observes** current state and **bounces** the approval if it changed since the operator last saw it. Closes the stale-approval / TOCTOU window.

**Record-now (owner-authored, as a non-locking build-carry):**
- **m-1:** record in your locked doc's build-carries / §2C section — the away-token refresh boundary: **rotate `decision_id`, burn prior-cycle nonces** on refresh; `verify` **re-observes and bounces on change**. Note it as **build-step (d)** work, m-1-confirmed-fit (pairs with the m-7 `re-mint-supersedes` carry already parked at §2C), **not** design-locked here (detailed design + fixture + adversarial review owed at (d)).
- **m-6:** record the **resummon edge** — a resummon triggers the rotate + re-observe (the m-6 side of the boundary); the 7-state park/wake FSM's refresh transition carries it as a build-carry.
- **CTO (me):** add decision ④ to the **integrated build-carry ledger** (`ARCHITECTURE.md` §C4 build-carries + `RECONCILE.md`) alongside the m-7 `re-mint-supersedes` carry — one place a (d) builder sees the full away-token freshness constraint.

**Residual (state it, per the m-7 pattern):** dormant in Step-1 (no away-bridge/resummon exists yet); until (d) builds it, the base decision-scoped sibling-burn + m-6's never-auto-resolve-on-expiry FSM hold. The full-pair adversarial review of the rotate/re-observe mechanism is a **(d) gate**, not a (c) one.

**Requirements (VP-set c5 shape):**
1. m-1 + m-6 **record** the decision as a non-locking build-carry in their locked docs (recorded constraint, **no mechanism design, no fixture** — that's (d)); I add it to the integrated ledger.
2. Each **addresses its implementer** for a review-only semantic `DESIGN-REVIEW` approve **of the record** (that it's correctly scoped record-now / mechanism-at-(d), not smuggled in as design-locked).
3. Map decision ④ → **recorded as §2C build-carry** (not "folded/locked") in the closure artifact.

Not authorized / not claimed: **record-now only** — no away-token mechanism design, no fixture, no design-lock of ④; no locked-contract mechanism reopen, no PLAN/IMPL/`pcode`/spike, no operator decision reopened. Mechanism/proof deferred to step (d).

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c5-fold-decision-4/DESIGN-orchestrator-planner-20260702-132818.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-4` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this decision-④ record-now dispatch + appended `master/relays/INDEX.md`; no doc edit (owners record; CTO ledgers on fold), no mechanism design, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-1 + m-6 record ④ as a non-locking §2C build-carry + address implementers for review; I add it to the integrated ledger + the c5 status ledger.
