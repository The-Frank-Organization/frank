## RECONCILE — item-A simplification amendment rev5 `80318a91…`: folds VP-R7's single blocker. F1(1) every §5.3 edge source+target now at FULL literal path (no `…/`, no deferral). F1(2) every §5.2 row carries an explicit clause (`whole_file` default; the close file's rows carry distinct clauses). F1(3) owner-base census REDONE over operative sections → the 3-edge set expands to 5: adds the m-1 §4 "PARKED producer-attaching halves" edge (m-9-C/m-10-C/§D-redaction discharges incl. the newly-added m-10-C confirmation row `774cd380`) and splits m-9 into C (§7/§9/§11) + B (§8/§11) + receipts (§9 items 4/5) edges. Owners HELD until ratify, then RELEASED.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this AMENDS ratified §4/§6/§11/§12; on VP approve it goes to operator hash-bound ratification (§8b). Master does not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-060000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review item-A simplification amendment rev5 `80318a91aa665df4dcbcf877e32637f2a0418d43540adae33671be4e0470f6df` — the single r7 blocker folded (full literal edge paths; a clause on every row; complete owner-base census incl. the omitted m-1 §4 + m-9 C/B loci); on approve → operator ratification; owners held then released

## What changed vs rev4 `abd0d723…`
Your r7 passed everything except one blocker with three exact defects. All three folded:

- **R7-F1(1) — edge paths abbreviated.** rev5 §5.3 writes **every source and target at its full literal path** (`master/relays/step3-relock-settlement-amend/…` and `master/relays/step3-relock-dag-m10/…`); the `…/` deferral note is deleted. No expansion is left to the item-A record.

- **R7-F1(2) — clause missing from most rows.** rev5 §5.1 adds the **clause rule** and §5.2 carries it: every row is `clause whole_file` (the lock binds the whole file) **except** the close file `…-160000.md`, whose rows carry the distinct clauses stated inline (item-E inventory / lane-2 close / N910 / r7-mirror / env-parity). §5.1 also clarifies the edge's section selector is not the owner-base row's clause (which stays `whole_file`).

- **R7-F1(3) — census incomplete.** rev5 §5.3 **redoes the census over operative sections**, expanding 3 edges → **5**:
  1. **m-1 §4 "The PARKED producer-attaching halves"** (`:57–60`; m-9-C, m-10-C, §D-redaction-co-sign) → discharged by item-C rows (m-9-C), the **newly-added m-10-C confirmation row** `master/relays/step3-relock-dag-m10/SITREP-planner-20260722-015123.md` (`774cd380…`, m-10-C), and the m-1 §D leg `…023020` + §D co-sign `…123000` (§D-redaction), ratified in the close `…160000`. My rev4 wrongly said m-1 carried no status.
  2. **m-9 C-consumption** at §7 `:476`, §9 `:499`/`:509`, §11 `:559` (not just the §9 ledger row) → item-C rows + close.
  3. **m-9 B-consumability** at §8 `:494`, §11 `:559` → §B sink (`…131500` + `…133000`) + close. (Previously ungoverned.)
  4. **m-9 §9 items 4/5** JOINT-PENDING receipts → §D co-sign + close (carried from rev4).
  5. **m-10 producer** §10/`assign`/S-1..S-5 JOINT-PENDING → §D co-sign + §B sink + B-carriage + close (carried from rev4).
  m-2/m-3/m-8 confirmed no operative cross-seam status; revision-history statements preserved as history; owner bytes unchanged in every edge.

## What I ask the VP to check
- Are all §5.3 edge paths now **fully literal**?
- Does **every** §5.2 row carry a clause (`whole_file` or the close file's distinct clause)?
- Is the census now **complete at operative sections** — m-1 §4 + all m-9 C/B/receipt loci + m-10, with the exact discharge targets and owner bytes unchanged?

## Owners — HELD until ratification, then RELEASED
No owner marks/nominates anything; on ratification the item-A hold releases.

## Boundaries
Changes no byte until operator ratification. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No lock record authored yet (item A, post-ratification). Ratified rev12 `1125b0a0…` UNMOVED; the eight settled bases + frozen finals + governing amendments UNMOVED. **H-12 continues to hard-block external use.**

## Verification
Amendment rev5 `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` = SHA-256 `80318a91aa665df4dcbcf877e32637f2a0418d43540adae33671be4e0470f6df` (supersedes rev4 `abd0d723…`). Verified at their bytes: m-1 §4 title `:57–60`; m-9 loci §7 `:476`, §8 `:494`, §9 `:499`/`:509`, §11 `:559`. Discharge relays re-hashed on disk: m-10-C confirm `774cd380`, m-1 §D leg `d096a4b3`, m-1 C-confirm `3070f34c`, m-9 B `95e8c6aa`, m-3 B `185acf9e`. Grep confirms no `…/`-abbreviated relay path remains in §5. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — edited `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` to rev5 + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no lock record authored, no owner action opened, no self-ratification.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews rev5 `80318a91…` (edge paths literal; every row clause present; census complete at operative sections); on approve → operator hash-bound ratification → master authors `STEP-3-INTERFACE-LOCK.md` (item A) → lane 4 → lane 5 (T4). Owners held until ratify, then released. H-12 stands.
