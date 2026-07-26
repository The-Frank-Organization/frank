## DESIGN-REVIEW -- MUST-REVISE one source-fold lineage pointer only: lock cbd1893c and every substantive gate pass, but ROADMAP still routes the current record to the voided 14:00 transmittal

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-design-review-r3
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- operator ratification of amendment rev7 remains satisfied; no design decision or lock byte is reopened
GRILL_REQUIRED: no -- this return is one stale source-fold lineage pointer
DESIGN_DOC_ID: step3-relock-item-a
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-180000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: PRESERVE interface-lock cbd1893c exactly; replace ROADMAP's stale 14:00 transmittal pointer with the corrected review trail and remove its premature completion claim before Item A closes

VERDICT: revise

Review target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-180000.md` at SHA-256 `5319c7c2066067ea6979beedcafa4d0c05f56883ff7e462d4bc027cec3ed1716`.

Accepted item-A record: `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

Ratified contract remains `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev7 at SHA-256 `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373`.

## Finding

### ITEM-A-LOCK-VP-R3-F1 -- GATE: ROADMAP's current milestone still points to the voided-lock transmittal

`ROADMAP.md:247` now correctly names current record `cbd1893c...`, but `ROADMAP.md:249` says that record was routed through:

`step3-relock-item-a/DESIGN-orchestrator-planner-20260727-140000.md`

That 14:00 relay is not the current transmittal. Its title, subject, verification, and requested action bind voided record `3e99edd0885fa5cb750014c03c012441d4a01acd5c0fe7ee6503bc2f0db73e38`, state the superseded 37-file count, and route lane 4 over that old identity. The bounded hash-token scan in the incoming relay is literally clean because ROADMAP carries the stale identity through a relay path rather than by repeating the old hash. Semantically, the source fold still routes a reader from current hash `cbd1893c...` to a relay for void hash `3e99edd0...`.

Required correction:

1. Preserve `master/STEP-3-INTERFACE-LOCK.md` byte-for-byte at `cbd1893c...`.
2. In `ROADMAP.md:249-250`, remove the 14:00 pointer and route the milestone through the corrected/current trail: the 16:00 corrected-record transmittal, the 17:00 record/F73 acceptance with source-fold return, and the 18:00 source-fold correction.
3. Remove or qualify `which COMPLETES item A`; completion remains pending this final VP source-fold confirmation. Do not imply that the voided 14:00 review request completed the item.
4. Return a bounded scan over the eight fold files showing no current milestone points to the 14:00 transmittal, no current `3e99edd0...` binding, and no 37-file claim. Return the unchanged lock hash.

This is a one-file source-fold correction. No operator re-ratification is required.

## Passed scope

- **R1 record findings remain CLOSED:** Section 6 carries the five literal ratified edges; the record declares 38 distinct files / 42 semantic rows; all close-file clauses are exact, including bare `env_digest-parity accepted disposition`.
- **R2 source-fold content findings are CLOSED:** all eight folded files name `cbd1893c...`; none contains current `3e99edd0...` or a 37-file claim.
- **Architecture consolidation passes:** D7 carries the run-wide restore/capacity/terminal-disposition refinements, and `relay.submit` spells out `{form_digest, dispatch_id?, to?, cc? | cc_unparsed?}` with omission and mutual-exclusion rules.
- **F73 remains accepted:** the lock is unchanged at `cbd1893c...`; the prior 38/38 constituent-hash, 42-row, path-set, exact-clause, and no-self-hash checks stand.
- Incoming exact-file lint is `OK`; `frank/` is clean at local/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

## Gate disposition

- Item A remains authored but not closed on this relay.
- Lane 4, exit-fixtures freeze/re-lock, lane 5, and T4 remain held.
- Do not edit the accepted lock, owner/frozen bytes, ratified amendment, or settled formula.
- Return only the ROADMAP lineage correction and its bounded proof. The next clean pass can be the final VP confirmation.

## Verification

- Recomputed SHA-256: incoming `5319c7c...`; accepted lock `cbd1893c...`; corrected 16:00 transmittal `f93e6cb8...`; voided 14:00 transmittal `325b8420...`.
- Exact-file lint is `OK` for the incoming 18:00 relay.
- Bounded scan of the eight fold files found no old hash, no 37-file claim, and exactly one old-transmittal pointer: `ROADMAP.md:249`.
- Direct inspection of the 14:00 relay found repeated binding to voided hash `3e99edd0...`.
- `git -C frank status --short --branch` is `## main...origin/main`, with empty porcelain and HEAD/origin both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no lock record, source-fold file, amendment, owner/frozen artifact, fixture, `frank/` source, branch, commit, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/DESIGN-REVIEW-orchestrator-reviewer-20260727-190000.md`.
Next requested action: preserve lock `cbd1893c...`; correct only ROADMAP's stale 14:00 lineage pointer and premature completion wording; return zero-stale-lineage proof for final VP confirmation. Lane 4 remains held.
