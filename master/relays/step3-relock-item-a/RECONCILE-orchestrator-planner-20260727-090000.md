## RECONCILE — item-A simplification amendment rev6 `7733e38b…`: folds VP-R8's single surgical blocker. §5.3 edge 1's source selector now enumerates ALL FOUR m-1 §4 parked halves — adds the omitted **m-9 D** (`:58`, the §2.2 writer gate / §2.3 create-open-verify+RED / §2.4 route-labeled sentinel legs / K6 exclusion) as a distinct source half with its own typed mapping (m-1 §D leg `023020` :42–45 "discharges parked half #2 m-9 D at-rest/redaction" + §D co-sign `123000` + close `160000`), kept distinct from the §D-redaction-co-sign half (shared targets, different source). No other rev5 change. Owners HELD until ratify, then RELEASED.

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
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-080000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review item-A simplification amendment rev6 `7733e38bd0c7b3f30b0158d40ef4560fcab5f2a5e911b28f619b13507cc3994e` — the r8 blocker folded (m-1 §4 edge now enumerates the fourth parked half m-9 D with its own typed mapping, distinct from the §D-co-sign half); every other rev5 correction preserved; on approve → operator ratification; owners held then released

## What changed vs rev5 `80318a91…` (the ONLY change)
Your r8 passed everything except one surgical gap. rev5 §5.3 edge 1 enumerated only three of m-1 §4's four parked halves. I verified at the bytes: m-1 §4 `:58` carries **both m-9 C and m-9 D** in one bullet (I missed m-9 D on the shared line), and the m-1 settlement leg `DESIGN-planner-20260723-023020.md:42–45` expressly discharges parked half **#2 (m-9 D at-rest/redaction)** as distinct from half **#4 (the §D redaction co-sign)**. Under the amendment's own "only listed edges govern" rule, that leg's presence as a target could not supersede a source half the edge never selected.

rev6 §5.3 edge 1 now enumerates **all four** halves and gives **m-9 D its own typed mapping** → m-1 §D leg `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md` (`d096a4b3…`) + §D co-sign `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md` (`2f3fb651…`) + close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md` (`fa2a634f…`). The §D-redaction-co-sign half stays a distinct source selector (shared targets, different source). No other byte moved.

## What I ask the VP to check
- Does edge 1 now select **all four** m-1 §4 halves, with m-9 D mapped distinctly from the §D-co-sign half at the exact settlement leg + co-sign + close?
- Is anything else in rev5 (which you passed) unintentionally moved? (Intent: nothing but edge 1.)

## Owners — HELD until ratification, then RELEASED
No owner marks/nominates anything; on ratification the item-A hold releases.

## Boundaries
Changes no byte until operator ratification. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No lock record authored yet (item A, post-ratification). Ratified rev12 `1125b0a0…` UNMOVED; the eight settled bases + frozen finals + governing amendments UNMOVED. **H-12 continues to hard-block external use.**

## Verification
Amendment rev6 `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` = SHA-256 `7733e38bd0c7b3f30b0158d40ef4560fcab5f2a5e911b28f619b13507cc3994e` (supersedes rev5 `80318a91…`). Verified at bytes: m-1 §4 `:58` carries m-9 C + m-9 D in one bullet, m-10 C `:59`, §D co-sign `:60`; m-1 settlement leg `:42–45` distinguishes half #2 (m-9 D) from half #4 (§D co-sign). Target hashes re-computed: m-1 §D leg `d096a4b3`, §D co-sign `2f3fb651`, close `fa2a634f`. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — edited `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` to rev6 (§5.3 edge 1 + status header only) + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no lock record authored, no owner action opened, no self-ratification.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews rev6 `7733e38b…` (edge 1 selects all four m-1 §4 halves incl. m-9 D with its distinct mapping); on approve → operator hash-bound ratification → master authors `STEP-3-INTERFACE-LOCK.md` (item A) → lane 4 → lane 5 (T4). Owners held until ratify, then released. H-12 stands.
