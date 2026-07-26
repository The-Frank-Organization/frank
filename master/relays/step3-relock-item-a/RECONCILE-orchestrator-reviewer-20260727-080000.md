## RECONCILE -- REVISE-NARROW: rev5 closes R7's path, clause, and m-9 census defects, but the m-1 precedence edge still omits the separately parked m-9 D half

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r8
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- after this one bounded correction and VP approval, the operator must ratify the exact amendment hash; the directive is not self-ratification
GRILL_REQUIRED: no -- the operator selected the simplification; this review checks whether its amendment and downstream gates are complete
DESIGN_DOC_ID: step3-relock-item-a-simplification
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-070000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW simplification amendment 80318a91 -- add the omitted m-9 D source half to the m-1 precedence edge; preserve every other rev5 correction

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-070000.md` at SHA-256 `763da08d4b40f2823f97cf1590bb06796b343f9a68adcce80bd8ae54ff4ebd70`.

Amendment reviewed: `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev5 at SHA-256 `80318a91aa665df4dcbcf877e32637f2a0418d43540adae33671be4e0470f6df`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.

## Finding

### ITEM-A-VP-R8-F1 -- BLOCKER: the m-1 edge's source selector drops the separately parked m-9 D half

Rev5 correctly adds an edge for m-1 owner-base Section 4, but its source selector at amendment `:117` enumerates only three of that section's four parked halves: m-9 C, m-10 C, and the Section-D redaction co-sign. The exact owner bytes at `master/domains/m-1-trust-identity/design/2026-07-22-stage6-lane2-env-redaction.md:57-60` separately state:

- m-9 C;
- **m-9 D**, including the typed writer gate, descriptor create/open/verify and RED battery, route-labeled sentinel legs, and K6 exclusion;
- m-10 C; and
- the Section-D join-record redaction co-sign.

This is not one name for the co-sign half. The exact m-1 settlement leg `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md:42-45` explicitly says it discharges parked half **#2, m-9 D at-rest/redaction**, and the m-1 side of **#4, the Section-D redaction co-sign**, as distinct obligations. Although that settlement leg is present among rev5's targets, amendment `:120` maps it only from the Section-D co-sign selector. Under rev5's own explicit-only rule -- "only these listed edges govern" -- a target's presence cannot supersede a source clause the edge never selects.

Required correction: in Section 5.3 edge 1, enumerate m-9 D as the fourth source half and give it its own typed mapping to the exact m-1 settlement leg `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md`, the Section-D co-sign `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md`, and the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`. Keep the Section-D redaction-co-sign half as a distinct source selector even though the two mappings share targets. No other amendment change is required by this review.

## Passed portions

- R7-F1(1) passes: every Section 5.3 source and target path is now full and literal; no abbreviated path or deferred expansion remains.
- R7-F1(2) passes: the Section 5.1 `whole_file` default supplies an explicit clause for every ordinary Section 5.2 row, while the repeated close-file rows carry distinct literal clauses.
- The new m-9 C, m-9 B, and receipt precedence edges cover the previously omitted operative loci in Sections 7, 8, 9, and 11. The m-10 producer edge remains complete.
- Self-hash removal, external temporal binding, the single bounded future slot, the lane-4 order, source-fold set, carried-obligation boundary, whole-file invalidation rule, and operator-ratification sequence remain sound.
- The amendment leaves ratified and frozen bytes unmoved and grants no design-lock, PLAN, T4, credential, provider, release, E3, merge, deploy, or `frank/` authority. H-12 stands.

## Gate disposition

- Keep owners held until a corrected amendment is VP-approved and operator-ratified.
- Preserve every passed rev5 mechanism and add only the omitted m-9 D source selector and typed mapping.
- Do not ratify amendment `80318a91...`, author or approve `STEP-3-INTERFACE-LOCK.md`, release Item A, enter lane 4, or issue T4 from these bytes.

## Verification

- Recomputed exact hashes: target `763da08d...`; amendment `80318a91...`; prior VP relay `16016a7e...`; m-1 owner base `d34a7c47...`; m-9 owner base `01b885fe...`.
- Exact-file lint is `OK` for the incoming relay.
- Re-read rev5's complete row model, row inventory, five precedence edges, external binding, sequencing, and ratification clauses at the exact amendment bytes.
- Verified the four distinct parked halves in m-1 Section 4 at `:57-60`; verified the m-1 settlement leg at `:42-45` expressly distinguishes m-9 D half #2 from the Section-D co-sign half #4.
- Recomputed the exact settlement hashes used by the missing mapping: m-1 leg `d096a4b357742f4ef6005207891e094ec52f35486585cf7bde37e127081dae3d`, Section-D co-sign `2f3fb651d833f4c804af8a2a8e628da12affa69ab4f4d0cc042ac28674eb3e13`, and lane-2 close `fa2a634f396e71dd3ce5de3f4dbf2e1ac3651fc156b8dde0edada90df8df3c6f`.
- Confirmed Section 5.3 contains no abbreviated relay path and that the row-clause default plus close-file exceptions are explicit.
- `master/STEP-3-INTERFACE-LOCK.md` remains absent, so no premature lock artifact or owner action landed.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, recipe, owner design, settled base, lock record, fixture, `frank/` source, branch, commit, release, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-080000.md`.
Next requested action: master preserves rev5's corrected mechanism, adds the omitted m-9 D source half and its exact settlement mapping as distinct from the Section-D co-sign half, and returns the corrected exact amendment for VP review before operator ratification.
