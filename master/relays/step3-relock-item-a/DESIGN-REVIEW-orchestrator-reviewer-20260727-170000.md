## DESIGN-REVIEW -- MUST-REVISE source-fold gate only: corrected lock cbd1893c passes VP + F73, but all eight folded files still bind the voided hash and the architecture leaves relay.submit's target object implicit

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-design-review-r2
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- operator ratification of amendment rev7 remains satisfied; this is the VP + F73 review of its item-A realization
GRILL_REQUIRED: no -- no design decision is reopened
DESIGN_DOC_ID: step3-relock-item-a
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-item-a/DESIGN-orchestrator-planner-20260727-160000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: PRESERVE interface-lock cbd1893c exactly; repoint the eight source-fold files from voided 3e99edd0, correct ROADMAP's 37-file count, and spell out the settled relay.submit target object before Item A closes

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/DESIGN-orchestrator-planner-20260727-160000.md` at SHA-256 `f93e6cb8a29b6797c9842428f96e468d09cdf8e7f2ec2d87392686768509ceb7`.

Corrected item-A record reviewed: `master/STEP-3-INTERFACE-LOCK.md` at externally named SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

Ratified contract remains `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev7 at SHA-256 `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373`.

## Finding

### ITEM-A-LOCK-VP-R2-F1 -- GATE: the source fold still names the voided record and does not fully state the settled `relay.submit` target

The corrected lock itself passes. The source-fold claim does not:

1. **All eight inventoried source-fold files still cite voided hash `3e99edd0...`, and none names current hash `cbd1893c...`:** `ROADMAP.md`; `master/README.md`; `master/ARCHITECTURE.md`; the m-1, m-2, and m-3 domain READMEs; `master/CYCLE-PLAYBOOK.md`; and `master/STEP-3-ITEM-A-RECIPE.md`. Several route lane 4 over the old hash. A whole-record edit voids the old lock by definition, so these are operative stale bindings, not harmless history.
2. **`ROADMAP.md:247` still describes a 37-file manifest.** The corrected and verified manifest is 38 distinct files across 42 semantic rows.
3. **The architecture consolidation leaves `target` undefined.** `master/ARCHITECTURE.md:539` says `"relay.submit:" || SHA-256(JCS(target))`, but the settled cell at `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md:71-85` and bound m-2 source `5ec7a3d2...:18-43,89-94` fixes the target object as `{form_digest, dispatch_id?, to?, cc? | cc_unparsed?}`, with absent optionals omitted and decoded `cc` mutually exclusive with `cc_unparsed`. The CC member was the pair-review correction that made the effect-target projection complete; replacing it with an undefined metavariable does not consolidate that decision-bearing grain.

Required correction: preserve `master/STEP-3-INTERFACE-LOCK.md` byte-for-byte at `cbd1893c...`; repoint all eight source-fold files to the corrected hash; change ROADMAP's manifest count to 38 files / 42 rows where both are useful; and replace the architecture's undefined `JCS(target)` shorthand with the settled exact target-member formula plus the omission and CC-branch rule. Return a bounded scan proving no source-fold file still names `3e99edd0...` as current and no 37-file claim remains.

## Closed findings and passed scope

- **R1-F1 CLOSED:** record Section 6 has full literal source and target paths, no `same file`, no prefix/suffix expansion, and the same five semantic edges as ratified amendment Section 5.3. The only textual difference is omission of the non-semantic `(r8-F1)` review-history tag.
- **R1-F2 CLOSED:** the record states 38 distinct files across 42 semantic rows; all five close-file clauses match exactly, including bare `env_digest-parity accepted disposition`, with no relocated locus prose.
- **R1-F3's D7 half CLOSED:** `master/ARCHITECTURE.md:537` now carries run-wide restore, `MAX_PARKED_ROWS_PER_RUN=512`, run-terminal `parked_unknown_capacity_exceeded`, zero-attempt `turn_failed` supersession, and the D-4 Gate-2 clarification. The action inventory is now disclosed.
- **F73 passes again:** all 38 constituent hashes recompute with zero mismatch, all owner/frozen bytes remain unchanged, the expected and realized path sets match, and the record contains no self-hash.
- External binding, invalidation, carried-source lineage, operator-ratification provenance, and H-12 remain correct. `frank/` remains untouched and clean.

## Gate disposition

- Do not edit the corrected lock record. Its exact hash `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636` is accepted at record/F73 grain and should survive the source-fold correction unchanged.
- Item A does not close on this relay because the ratified post-ratification fold still routes readers and lane 4 to a voided lock identity.
- Master may make only the source-fold corrections above and return the unchanged lock hash for final VP confirmation. No operator re-ratification is required.
- Lane 4, exit-fixtures freeze/re-lock, lane 5, and T4 remain held.

## Verification

- Recomputed current hashes: target `f93e6cb8...`; corrected lock `cbd1893c...`; ratified amendment `3443f73d...`; prior VP review `e6c1ada9...`.
- Exact-file lint is `OK` for the incoming relay.
- Parsed and re-hashed all 38 record paths: 38/38 match; 42 semantic rows declared; zero self-hash occurrence.
- Compared record Section 6 to amendment Section 5.3; the five edge bodies differ only by removal of the `(r8-F1)` review-history label.
- A bounded scan over the eight source-fold files returned eight files containing `3e99edd0...`, zero containing `cbd1893c...`, and one live 37-file claim in `ROADMAP.md`.
- Re-read the settled `relay.submit` formula from amendment Correction 3 and the pair-approved m-2 cell, including the decoded-CC versus `cc_unparsed` branch.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no lock record, source-fold file, amendment, owner/frozen artifact, fixture, `frank/` source, branch, commit, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/DESIGN-REVIEW-orchestrator-reviewer-20260727-170000.md`.
Next requested action: preserve lock `cbd1893c...`, repoint all eight source-fold files and the ROADMAP count, expand the architecture's settled `relay.submit` target formula, and return bounded zero-stale-hash proof for final VP confirmation; lane 4 remains held.
