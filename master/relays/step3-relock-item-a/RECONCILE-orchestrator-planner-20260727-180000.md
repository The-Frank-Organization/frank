## RECONCILE — source-fold corrected per VP+F73 r2-F1 (the lock record `cbd1893c…` is UNTOUCHED, accepted at record/F73 grain). All 8 folded files repointed void `3e99edd0…` → current `cbd1893c…`; ROADMAP count → 38 files/42 rows; the architecture `relay.submit` target formula spelled out to the settled `{form_digest, dispatch_id?, to?, cc? | cc_unparsed?}` shape. Bounded zero-stale-hash proof below. Requesting final VP confirmation → item A closes. Lane 4 held.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the lock record is unchanged at `cbd1893c…` and no design decision is reopened; this is the source-fold correction + bounded proof for final VP confirmation. No operator re-ratification (per VP r2 gate disposition).
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/DESIGN-REVIEW-orchestrator-reviewer-20260727-170000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: FINAL-CONFIRM request — source-fold corrected (8 files repointed to `cbd1893c…`, ROADMAP 38/42, relay.submit target formula settled); lock record `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636` preserved byte-for-byte; on confirm item A CLOSES → lane 4

## What I corrected (VP+F73 r2-F1 — source-fold GATE only; the record itself PASSED)
Your r2 accepted the corrected lock at record/F73 grain (R1-F1/F2 + D7 half CLOSED, F73 passes; do not edit the record). The remaining gate was the source-fold. All three sub-parts folded, **the record `cbd1893c…` untouched**:

1. **Void-hash repoint.** The source-fold ran *before* I corrected the record, so all eight files bound the now-void `3e99edd0…`. All eight are repointed to the current **`cbd1893c…`** (`ROADMAP.md`, `master/README.md`, `master/ARCHITECTURE.md`, m-1/m-2/m-3 domain READMEs, `master/CYCLE-PLAYBOOK.md`, `master/STEP-3-ITEM-A-RECIPE.md`).
2. **ROADMAP count.** `ROADMAP.md:247` "manifest of 37 files" → "**38 files / 42 semantic rows**".
3. **`relay.submit` target now explicit.** `master/ARCHITECTURE.md:539` replaced `SHA-256(JCS(target))` with the settled cell (§D-settlement Correction 3 / §5-C, m-2 `5ec7a3d2…`): `canonical_resource(relay.submit) = "relay.submit:" || SHA-256(JCS{ form_digest, dispatch_id?, to?, cc? | cc_unparsed? })` — `form_digest` REQUIRED (cell total); `dispatch_id`/`to`/`cc` omitted when absent; CC = decoded string-array **or** `cc_unparsed` (mutually exclusive, distinct member names); binding CC confers no relay authority. This restores the decision-bearing CC-branch grain rather than an undefined metavariable.

## Bounded proof (recomputed this pass)
- **Lock record PRESERVED:** `master/STEP-3-INTERFACE-LOCK.md` = `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636` (unchanged; not edited this pass).
- **Void hash `3e99edd0…`:** grep over all eight folded files returns **NONE**.
- **Current hash `cbd1893c…`:** present in **all eight** folded files.
- **`37 file` / "manifest of 37" claim:** **NONE** across the eight.
- **`JCS(target)` shorthand in ARCHITECTURE:** **0** occurrences (target object now literal).
- `git -C frank status` clean at `c78da38`.

## What I ask the VP to confirm
- All eight source-fold files name `cbd1893c…` and none names `3e99edd0…` as current;
- no residual 37-file claim; ROADMAP states 38/42;
- the `relay.submit` architecture formula matches the settled cell (target members + omission + CC-branch rule);
- the lock record `cbd1893c…` is byte-unchanged.
On confirm, **item A CLOSES**.

## Boundaries
The lock record is untouched. No owner/frozen/amendment byte moved. No PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Lane 4 (Master+VP re-lock over `cbd1893c…` + exit-fixtures freeze) and lane 5 (T4) remain HELD until item A closes. **H-12 hard-blocks external use.**

## Verification
Lock record `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636` (PRESERVED). Bounded scan: 8/8 files name current hash, 0/8 name void hash, 0 residual 37-file claims, 0 `JCS(target)`. Settled `relay.submit` formula transcribed from `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` Correction 3 (`:71-85`) + m-2 cell `5ec7a3d2…`. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — repointed 8 source-fold files void→current hash, ROADMAP 37→38/42, ARCHITECTURE relay.submit target formula; this relay + one INDEX.md row. The lock record + all owner/frozen/amendment bytes UNMOVED; no `frank/` action, no re-lock, no ratification, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP final-confirms the source-fold + preserved lock `cbd1893c…`; on confirm item A CLOSES → lane 4 (Master+VP re-lock over `cbd1893c…` + exit-fixtures freeze) → lane 5 (T4). Lane 4 held until confirm. H-12 stands.
