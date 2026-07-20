## RECONCILE — revise accepted: owner-pair semantic review on the claim-sweep; decision-④ split (record now / mechanism at (d)); 5-lane shape confirmed

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c5-decomp
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner — revise **accepted** in full; both corrections are right and mine to fold.

1. **Owner-pair adversarial confirmation on the claim-sweep.** You're right that the sweep *is* the NO-GO's failed claim class and can't ride a planner-confirm alone. Folded: CTO prepares the candidate relabels + a **semantic sweep checklist** (the §2B rule + m-7's §16 exclusivity-class method); the **owning planner folds/confirms**; the **owning implementer files a review-only `DESIGN-REVIEW` semantic approve** of the relabels. Light ceremony + batching for the low-hit domains, but the pair reviewer is never bypassed.
2. **Decision-④ split.** Folded: **record the decision now (c)** as a **non-locking §2C build-carry** in m-1/m-6 + the integrated ledger — rotate `decision_id`, burn prior nonces, re-observe current state, bounce/hold if changed; **defer the mechanism / fixture / adversarial proof to (d)** (before park/wake or away-bridge ships). Not a full defer, not a full fold.

**Q1–Q4 accepted as answered:** rigor = CTO-candidate + planner-fold + implementer-semantic-approve for the sweep, **full-pair mandatory for ③/⑤ and the ④ build-carry record**, no grill absent a new mechanism decision; scope = **all six domain docs + ARCHITECTURE**, no out-of-scope spot-checks (m-1/m-2 heavier, light domains batched-but-owner-confirmed); ④ = record-now/mechanism-at-(d); ownership = CTO edits ARCHITECTURE directly + *proposes* domain relabel patches, **never the final owner-confirming act** on a locked domain doc (owner folds or explicitly accepts + implementer reviews).

**Confirmed dispatch set (your 5-lane shape):**
1. **`c5-claim-sweep-architecture`** — CTO edits `ARCHITECTURE.md` (establishes the semantic checklist the domain lanes inherit); VP reviews.
2. **`c5-claim-sweep-m-1`** + **`c5-claim-sweep-m-2`** — CTO candidate sweep → owning planner fold/confirm → implementer semantic review (the heavy trust/identity + schema overclaim load).
3. **`c5-claim-sweep-light`** — one batched relay to m-3/m-4/m-5/m-6 (1–3 hits each); each domain owner-pair confirms its own claim text.
4. **Decision-folds** — **③** RAISE-ONLY A/B + known-A detector (§J/CTO + m-6, owner-authored + implementer-reviewed); **⑤** ODB model-name carve-out (m-3 + m-6 + m-4, owner-authored + implementer-reviewed); **④** away-token record-now build-carry (m-1 + m-6 + ledger; mechanism → (d)).
5. **Byte-consistency re-verify** — CTO cross-doc check after the folds; VP reviews **before step (c) is marked closed.**

**Sequencing:** I fire `c5-claim-sweep-architecture` first (it sets the checklist), then the domain sweep + decision-fold lanes in parallel; the re-verify is the closing gate. I'll carry a single **c5 status ledger** (which lane closed, owner-confirmed + implementer-approved) so nothing folds into "step (c) closed" without its owner-pair confirmation — the same closure discipline that held the CQ gate.

No re-review needed on this fold (I'm executing your prescribed shape); you're CC'd on each lane + the closing re-verify comes to you. If I've mis-taken any item, flag it and I hold.

Not authorized / not claimed: no doc edited by this relay, no pair re-engaged yet (the lanes follow), no mechanism change, no PLAN/IMPL/`pcode`/spike, no Step-1 PLAN, no operator decision reopened.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c5-decomp/RECONCILE-orchestrator-planner-20260702-042317.md` — OK
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-decomp` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read the VP revise `c5-decomp/RECONCILE-orchestrator-reviewer-20260702-042018`; wrote this fold-acceptance relay + appended `master/relays/INDEX.md`; no doc edit, no pair re-engaged, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: I write lane 1 (`c5-claim-sweep-architecture`, CTO-edited) + the candidate sweeps for the domain lanes; the operator re-engages the pairs for the fold/confirm + semantic review.
