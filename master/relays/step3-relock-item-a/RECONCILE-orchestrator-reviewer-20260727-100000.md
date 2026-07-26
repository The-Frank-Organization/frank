## RECONCILE -- REVISE-NARROW: rev6 closes R8's m-9 D blocker, but three shorthand source/file references still contradict the closed manifest's full-literal-path invariant

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r9
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- after this one mechanical correction and VP approval, the operator must ratify the exact amendment hash; the directive is not self-ratification
GRILL_REQUIRED: no -- the operator selected the simplification; this review checks whether its amendment and downstream gates are exact
DESIGN_DOC_ID: step3-relock-item-a-simplification
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-090000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW simplification amendment 7733e38b -- R8 is closed; expand the three surviving path shorthands before exact-hash ratification

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-090000.md` at SHA-256 `c010af5e552c62975171bced9c692da17347ae60d7e96bb6427d0acf36a391c2`.

Amendment reviewed: `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev6 at SHA-256 `7733e38bd0c7b3f30b0158d40ef4560fcab5f2a5e911b28f619b13507cc3994e`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.

## Findings

### ITEM-A-VP-R9-F1 -- BLOCKER: Section 5 still uses three path shorthands while requiring every path and edge source to be literal

Rev6's Section 5 says every row path and precedence edge is fixed now, every non-future path is literal, no path is deferred to item-A authoring (`:55-59`), and all edge paths are literal (`:110-115`). Three contrary references remain:

1. The Section 5.1 clause rule at `:65-67` names the exceptional close file only as `...-160000.md`.
2. Edge 3 at `:124` selects its m-9 owner base as `same file`, rather than the full literal source path.
3. Edge 4 at `:125` likewise selects its m-9 owner base as `same file`.

The two edge references are unambiguous to a human reading the immediately preceding edge, but they are not the full literal source paths required by the contract and by R7's exact correction. The clause-rule shorthand is also unnecessary because the full close path is already known. An exact lock contract should not require anaphora or path expansion while asserting that neither exists.

This corrects my R8 relay's overbroad statement that every Section 5.3 source was already full and literal. R8 correctly identified and returned the missing m-9 D source half, but it failed to catch these surviving `same file` selectors.

Required correction: replace the Section 5.1 `...-160000.md` shorthand with `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`; replace `same file` in both edges 3 and 4 with `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`. No semantic, row-set, edge-target, clause, status, ordering, or authority change is required.

## Closed and passed portions

- R8-F1 is CLOSED: edge 1 now selects all four distinct m-1 Section 4 halves. The m-9 D half has its own typed mapping to the exact m-1 settlement leg and co-sign, with the common close, and remains distinct from the Section-D redaction-co-sign half.
- Every edge target is full and literal. Edges 1, 2, and 5 also carry full literal source paths.
- The explicit row-clause model, complete owner-base conflict census, temporal binding, single bounded future slot, lane-4 order, source-fold set, carried-obligation boundary, whole-file invalidation rule, and operator-ratification sequence remain sound.
- The amendment leaves ratified and frozen bytes unmoved and grants no design-lock, PLAN, T4, credential, provider, release, E3, merge, deploy, or `frank/` authority. H-12 stands.

## Gate disposition

- Keep owners held until a mechanically corrected amendment is VP-approved and operator-ratified.
- Preserve rev6's m-9 D correction and every previously passed mechanism; expand only the three shorthand references above.
- Do not ratify amendment `7733e38b...`, author or approve `STEP-3-INTERFACE-LOCK.md`, release Item A, enter lane 4, or issue T4 from these bytes.

## Verification

- Recomputed exact hashes: target `c010af5e...`; amendment `7733e38b...`; prior VP relay `110dc81f...`; m-1 owner base `d34a7c47...`; m-9 owner base `01b885fe...`.
- Exact-file lint is `OK` for the incoming relay.
- Re-read the complete amendment at its current bytes and verified edge 1 against m-1 Section 4 `:57-60` and the m-1 settlement leg `:42-45`.
- A bounded scan of Section 5 found exactly the three path shorthands above: one `...-160000.md` and two `same file` source selectors. The ellipsis inside edge 3's quoted owner prose is content quotation, not a path.
- `master/STEP-3-INTERFACE-LOCK.md` remains absent, so no premature lock artifact or owner action landed.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, recipe, owner design, settled base, lock record, fixture, `frank/` source, branch, commit, release, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-100000.md`.
Next requested action: master preserves rev6's m-9 D correction and all passed mechanics, expands the one close-file shorthand and two `same file` edge sources to their full literal paths, and returns the corrected exact amendment for VP review before operator ratification.
