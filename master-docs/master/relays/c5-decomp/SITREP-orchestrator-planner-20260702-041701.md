## SITREP — re-baseline step (c) decomposition: the global claim-sweep + remaining decision-folds; VP review before dispatch

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: c5-decomp
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — scope/sequencing decomposition; VP-reviewed, operator on CC (operator authorized opening step (c))
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner — c4 is CLOSED (conductor-core LOCKED, your co-sign `c4-design-m-7/…-040327`). The operator authorized **re-baseline step (c)**. Per the decompose→review→dispatch spine, I bring you the step-(c) decomposition before any dispatch. This is a **cross-domain claim-hygiene + decision-fold pass over the LOCKED c1–c3 docs** — a doc-text pass (`DESIGN-REVIEW` §2B: *"a doc change, NOT mechanism removal — none of the 2A mechanisms collapse, only their malicious-lane claims"*), not a redesign.

**Step (c) = three parts (`DESIGN-REVIEW` §2B + §5c):**

**Part 1 — the global claim-sweep.** Relabel every "by-construction / structural / sole-writer / unbypassable / forgery-robust" adversarial-strength claim → the honest **D4 claim** (confusion-resistant; a malicious code-executing agent is explicitly out of scope) + record the **D5 accepted-risks** (config / store-write / operator-FROM under same-uid attach). Rule = §2B + the **semantic** methodology m-7 proved in its §16 sweep (scope exclusivity claims to the conductor-governed surface WITH the D5 residual beside them; keep genuine control-flow invariants like the serialized-loop kill). **Overclaim census** (candidate hits): **m-1 = 22** (the TCB — "forgery-robust by construction / sole-writer store I1", the epicenter of the WRAP-assumption the NO-GO inverted), **ARCHITECTURE = 15**, **m-2 = 11**, m-3 = 2, m-4 = 3, m-5 = 1, m-6 = 1. (Not all hits are overclaims — several are legit; the sweep is semantic, not find-replace.)

**Part 2 — fold the remaining operator decisions** (`READINESS-REGISTER` §Operator-decisions):
- **①** attach + confusion-resistant → **rides the m-1 claim-sweep** (record the confusion-resistant claim + wrap-upgrade path; runtime identity fields never accepted as FROM; conductor-owned channel/credential = sole stamp source). Owner: m-1.
- **② `self_reported` fail-closed → DONE in c4** (CQ-2 folded into m-3). Re-verify only.
- **③ RAISE-ONLY A/B** + known-A detector → owner **§J/CTO + m-6** (classification-direction invariant + detector).
- **④ away-token ROTATE + RE-OBSERVE** → owner **m-1 + m-6**. **Open Q:** fold the *decision* into the docs now (c), or defer entirely to (d) the away-bridge build step (where the mechanism builds, alongside the m-7 `re-mint-supersedes` carry)?
- **⑤ ODB model-name egress carve-out** (narrow, ODB→operator only; R2 untouched) → owner **m-3 + m-6 + m-4**.

**Part 3 — byte-consistency re-verify.** Confirm the c4 folds (② fail-closed, §2A.6 A-floor table via CQ-3, §2A.7 decision-② via CQ-2) are consistent across the docs and the sweep breaks nothing.

**Proposed approach (for your review):**
- **ARCHITECTURE.md** (CTO/VP-owned, 15 hits) → **CTO sweeps directly**, you review.
- **Domain docs** → the sweep propagates a *decided* operator claim (① confusion-resistant), so I lean **CTO-proposes-the-relabels + owning-planner-confirms-or-corrects** (bounded; lighter than the CQ produces) — the pair still owns its claims. The **decision-folds (③/⑤/④)** are genuine additions → **owner-authored** (not CTO-proposed).
- Whole sweep → you review for honest-line adherence + consistency; I run the byte-consistency re-verify.

**Where I want you to push (VP asks):**
- **Q1 — rigor/lift.** After the 6-pair CQ pass, is **CTO-proposes + planner-confirms** right for the *claim-sweep* (relabel hygiene), with **full-pair only for the decision-folds** (③/⑤/④, which add invariants)? Or hold the whole thing to full-pair?
- **Q2 — light-domain scope.** Sweep all six domain docs, or **CTO-spot-check the light ones** (m-3/m-4/m-5/m-6, 1–3 hits) and only re-engage **m-1 + m-2** (the heavy) for the sweep proper — with m-3/m-4/m-6 re-engaged anyway for their decision-folds?
- **Q3 — decision ④.** Fold the away-token rotate+re-observe *decision* now (c), or defer to (d) the away-bridge build step? I lean **defer to (d)** — it's a §2C away-bridge mechanism, dormant in Step-1, and pairs with the m-7 `re-mint-supersedes` carry already parked there.
- **Q4 — ownership.** Confirm the CTO may edit ARCHITECTURE directly + *propose* domain relabels (owner-confirms), without tripping the no-proxy-edit bar — since §2B is a directed hygiene pass propagating a locked operator decision, not a new design choice.

On your concurrence (+ any re-scope) I write the step-(c) dispatch(es). No doc edited yet.

Not authorized / not claimed: no doc edited by this SITREP, no pair re-engaged yet, no mechanism change, no PLAN/IMPL/`pcode`/spike, no Step-1 PLAN, no locked-contract *mechanism* reopen (claim-text hygiene only), no operator decision reopened (folding the recorded ones).

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c5-decomp/SITREP-orchestrator-planner-20260702-041701.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-decomp` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: scoped step (c) (read `DESIGN-REVIEW` §2B/§5c + `READINESS-REGISTER` operator-decisions + an overclaim census across the six locked docs + ARCHITECTURE); wrote this decomposition SITREP + appended `master/relays/INDEX.md`; no doc edit, no pair re-engaged, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP review of the step-(c) decomposition (Q1–Q4); on concurrence I dispatch the claim-sweep (CTO ARCHITECTURE + per-domain) + the decision-folds (③/⑤, ④ per Q3).
