## RECONCILE — item-A simplification amendment rev7 `3443f73d…`: folds VP-R9's single mechanical blocker. Expanded the three surviving path shorthands to full literal paths — §5.1 clause rule (`…-160000.md` → `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`) and edges 3 & 4's `same file` source selector → `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`. R8's m-9 D correction CLOSED; no semantic/row-set/edge-target/clause/ordering/authority change. Owners HELD until ratify, then RELEASED.

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
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-100000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review item-A simplification amendment rev7 `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373` — the r9 mechanical blocker folded (three path shorthands → full literal paths); R8 m-9 D correction preserved; every other passed mechanism unmoved; on approve → operator ratification; owners held then released

## What changed vs rev6 `7733e38b…` (the ONLY change — mechanical)
Your r9 CLOSED R8 and left one mechanical blocker: three path shorthands surviving under the full-literal-path invariant. All three expanded, no semantic change:
1. **§5.1 clause rule** — `…-160000.md` → `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
2. **§5.3 edge 3 (m-9 B)** source — `same file` → `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`.
3. **§5.3 edge 4 (m-9 §9 receipts)** source — `same file` → `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`.

Nothing else moved: the row set, all five edges' targets, edge 1's four m-1 §4 halves incl. the m-9 D mapping, clauses, ordering, external binding, carried boundary, and fold set are byte-unchanged from rev6.

## What I ask the VP to check
- Are the three references now full literal paths, with no residual `same file` / bare-filename shorthand in §5?
- Is rev6's passed content otherwise byte-identical (only the three expansions + the status header changed)?

## Owners — HELD until ratification, then RELEASED
No owner marks/nominates anything; on ratification the item-A hold releases.

## Boundaries
Changes no byte until operator ratification. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No lock record authored yet (item A, post-ratification). Ratified rev12 `1125b0a0…` UNMOVED; the eight settled bases + frozen finals + governing amendments UNMOVED. **H-12 continues to hard-block external use.**

## Verification
Amendment rev7 `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` = SHA-256 `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373` (supersedes rev6 `7733e38b…`). Grep of §5 finds no `same file` selector and no bare `…-160000.md` path (the sole remaining "same file" string is the status header describing the fix). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — edited `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` to rev7 (three path expansions + status header) + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no lock record authored, no owner action opened, no self-ratification.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews rev7 `3443f73d…` (three shorthands now literal; rest byte-unchanged); on approve → operator hash-bound ratification → master authors `STEP-3-INTERFACE-LOCK.md` (item A) → lane 4 → lane 5 (T4). Owners held until ratify, then released. H-12 stands.
