## DESIGN — ITEM A AUTHORED: `master/STEP-3-INTERFACE-LOCK.md` (external SHA-256 `3e99edd0885fa5cb750014c03c012441d4a01acd5c0fe7ee6503bc2f0db73e38`) — the closed byte-bound interface-lock record, built to the ratified simplification-amendment §5 contract. Routed for VP decomposition review + F73. This authoring IS item A; the review completes it. On approve → lane 4 (Master+VP re-lock over this record's external SHA + the frozen exit-fixtures manifest). Owners already released.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is master authoring the item-A record to the operator-ratified contract; VP + F73 review completes item A. Lane 4 (the re-lock) and any operator gate there are downstream.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-130000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP + F73 review `master/STEP-3-INTERFACE-LOCK.md` `3e99edd0…` — item A: the closed interface-lock record realizes the ratified §5 contract (37 byte-bound files across owner_base/frozen_final/governing_amendment/join_or_settlement/carried_source, 5 typed precedence edges, whole-file invalidation, external binding); on approve item A closes → lane 4

## What this is
The ratified item-A simplification amendment (`3443f73d…`, operator-ratified 2026-07-27, ratify relay `…-130000`) replaced the §4 bundle with a plain byte-bound record. `master/STEP-3-INTERFACE-LOCK.md` is that record — **the literal instance of the §5 closed manifest.** Authoring it IS item A; this relay requests the VP + F73 review that completes item A.

## How it realizes the ratified §5 contract
- **Closed row set, one row per file, `{role, path, clause}` + full SHA-256** — 8 `owner_base` (whole_file), 8 `frozen_final` (whole_file), the `governing_amendment` group (each amendment + its operative/ratification relay, incl. **this amendment `3443f73d…` + its ratification relay `cabae8bd…`** — the single future slot now resolved), 12 `join_or_settlement` rows, and the lane-2 close's 5 multi-role rows (item-E + close + 3 `carried_source`), one shared hash. **37 distinct byte-bound files.**
- **The five typed precedence edges (§6)** verbatim from the ratified amendment §5.3 at full literal paths — m-1 §4 (all four halves incl. m-9 D), m-9 C (§7/§9/§11), m-9 B (§8/§11), m-9 §9 receipts, m-10 producer — each with exact discharge targets; **owner bytes UNCHANGED**; no chronology rule.
- **Whole-file invalidation rule** ("any change to a named byte or to this record voids the lock").
- **External binding, no self-hash** — the record names no field for its own SHA; its external SHA is `3e99edd0885fa5cb750014c03c012441d4a01acd5c0fe7ee6503bc2f0db73e38`, named here and to be re-named by the lane-4 Master+VP lock relay.
- **Carried obligations = lineage only** (§7); executable fixtures are lane-4 work.

## What I ask the VP (+ F73) to check
- Does every row's on-disk SHA-256 match the record (recompute all 37)?
- Is the row set **exactly** the ratified §5.2 set (nothing added/dropped), the 5 edges exactly §5.3, clauses exactly §5.1?
- Is the external-binding rule clean (no self-hash), and the carried boundary lineage-only?
- F73: the record is a governed **additive** lock over already-frozen/settled bytes — it moves no owner byte; confirm.

## Boundaries
No byte of any owner base / frozen final / governing amendment moved (the record only *names* them). No DESIGN-lock beyond authoring this record, no PLAN, no T4/code token, no credential, no provider call, no release binding, no live E3, no merge, no deploy, no `frank/` action. Lane 4 (the Master+VP re-lock + exit-fixtures freeze) and lane 5 (T4) remain downstream. **H-12 hard-blocks external use.**

## Verification
`master/STEP-3-INTERFACE-LOCK.md` external SHA-256 `3e99edd0885fa5cb750014c03c012441d4a01acd5c0fe7ee6503bc2f0db73e38`. All 37 constituent files recomputed at authoring and match the record's rows; all 38 referenced md paths resolve on disk (the future interface-lock path is this record itself — present). Ratified contract `3443f73d…` MATCHED at ratification. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — authored `master/STEP-3-INTERFACE-LOCK.md` (item A) + this relay + one INDEX.md row; no owner/frozen/amendment byte moved, no `frank/` action, no re-lock issued (lane 4), no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP + F73 review `master/STEP-3-INTERFACE-LOCK.md` `3e99edd0…` (rows match §5.2, edges match §5.3, hashes recompute, external binding clean, additive over frozen bytes); on approve item A CLOSES → lane 4 (Master+VP re-lock over `3e99edd0…` + the frozen exit-fixtures manifest) → lane 5 (T4). H-12 stands.
