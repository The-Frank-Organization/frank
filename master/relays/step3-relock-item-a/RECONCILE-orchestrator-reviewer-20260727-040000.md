## RECONCILE -- REVISE-NARROW: rev3 fixes self-hashing, lane order, and the source fold, but its claimed literal row set remains open and the amendment-ratification relay cannot cite a future record hash

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r6
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- after the two bounded corrections and VP approval, the operator must ratify the exact amendment hash; the directive is not self-ratification
GRILL_REQUIRED: no -- the operator selected the simplification; this review checks whether its amendment and downstream gates are complete
DESIGN_DOC_ID: step3-relock-item-a-simplification
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-030000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW simplification amendment 512e9c52 -- retain external binding and the corrected lane order, but make the artifact/role/edge set actually literal and remove the impossible pre-authoring operator citation

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-030000.md` at SHA-256 `7b145fe00d6f578d7b8241274264584aee33b8b06749b046af7449ed9ca22d92`.

Amendment reviewed: `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev3 at SHA-256 `512e9c52efd517044ef144168408cb17659a70aa112e7f2d5d8e48e097e096f0`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.

## Findings

### ITEM-A-VP-R6-F1 -- BLOCKER: Section 5 still does not contain the closed literal row and precedence sets it claims

Amendment `:63-67` says this amendment fixes every row's role and literal path, with no row deferred to authoring. The actual inventory still:

- says the m-3 and settlement ratification relays are resolved later (`:96-97`), although their exact operative records already exist at `step3-stage6-m3-schema-amend/RECONCILE-orchestrator-planner-20260722-234500.md` and `step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-170000.md`;
- uses `.../` path abbreviations throughout the join, correction, carried-source, and precedence rows (`:101-111`, `:120`);
- permits the author to refine the item-E/carried paths later (`:113-114`), directly contradicting the fixed-row-set claim;
- permits the author to add further precedence edges later (`:122-123`), so the ratified edge set is not closed.

The row model is internally inconsistent as well. Section 5.1 requires one row per actual file with one scalar `role`, but the same lane-2-close file is listed as item-E, the close record, and each of three `carried_source` rows (`:105-111`). Amendment/contract/ratification files are also grouped inside bullets despite the one-file-row rule. Those entries cannot be represented faithfully by the declared shape without duplicate file rows or lost roles/clauses.

Required correction: expand every pre-existing path literally, including the two known ratification records; define only the future simplification-ratification record as a single post-ratification slot that the final record must resolve to one literal path/hash; remove all author-time path refinement. Choose one coherent representation: either a unique artifact table with closed `roles[]` and exact clause selectors, or semantic-binding rows whose identity explicitly includes `{role, path, clause}` and may repeat a file. Perform the owner-base conflict census now and make the typed precedence-edge list exhaustive; do not leave additional edge discovery to item-A authoring.

### ITEM-A-VP-R6-F2 -- BLOCKER: the current operator-ratification record cannot cite the later interface-lock hash

Section 5.4 `:125-129` correctly binds the final record externally, but says "the operator ratification record cites that same external hash." Under the amendment's own sequencing (`:138-146`), the operator first ratifies this amendment and only then does master author `STEP-3-INTERFACE-LOCK.md`. That ratification record cannot cite the hash of bytes that do not yet exist.

Required correction: the present operator-ratification record binds only this amendment's exact hash. The later item-A VP review relay and lane-4 Master+VP lock relay bind the finalized interface-lock record hash. Remove the operator citation from Section 5.4 unless a distinct post-item-A human gate is deliberately added to the sequencing, authority fields, and replacement clauses.

## Passed portions

- R5-F1's core repair passes: no in-file self-hash or fixed-point placeholder remains; the generic invalidation rule plus external relay binding is sound.
- R5-F2 passes: Section 11 step 5, Section 4, and Section 7 now state one order -- content-address fixture inputs/baselines, freeze final manifest, re-lock over record plus manifest, then T4.
- R5-F4 passes: the source-fold manifest now covers every live old-mechanism route found in the roadmap, dashboard, architecture, cycle playbook, and m-1/m-2/m-3 dashboards. Historical relay/ledger text remains append-only.
- All eight `owner_base` and eight `frozen_final` literal paths exist and independently hash to the listed digest prefixes.
- The single carried-obligation boundary and the plain whole-file invalidation mechanism remain sound.
- The amendment preserves VP review -> operator hash-bound amendment ratification, leaves ratified/frozen bytes unmoved, and grants no design-lock, PLAN, T4, credential, provider, release, E3, merge, deploy, or `frank/` authority. H-12 stands.

## Gate disposition

- Keep owners held until a corrected amendment is VP-approved and operator-ratified.
- Preserve the operator-selected plain record lock, external final-record binding, corrected fixture order, and complete source-fold set.
- Do not ratify amendment `512e9c52...`, author or approve `STEP-3-INTERFACE-LOCK.md`, release Item A, enter lane 4, or issue T4 from these bytes.

## Verification

- Recomputed exact hashes: target `7b145fe0...`; amendment `512e9c52...`; prior VP relay `1aaa8d09...`; ratified rev12 `1125b0a0...`.
- Verified the known ratification records on disk: stage-6 rev12 record `7c367c7f...`; m-3 schema record `49c811fd...`; settlement record `984071fb...`.
- Exact-file lint is `OK` for the incoming relay.
- Re-read every external-binding, ordering, row-set, role, precedence, carried-source, source-fold, sequencing, and authority clause at the exact amendment bytes.
- Independently hashed all eight owner bases and eight frozen finals; every path and listed digest prefix matches.
- Scanned current non-relay status sources for the withdrawn bundle mechanism; the seven files now named by the fold manifest are the complete live set found.
- `master/STEP-3-INTERFACE-LOCK.md` remains absent, so no premature lock artifact or owner action landed.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, recipe, owner design, settled base, lock record, fixture, `frank/` source, branch, commit, release, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-040000.md`.
Next requested action: master preserves rev3's corrected mechanism/order/fold, makes the artifact/role/clause and precedence sets literally closed at amendment time (with one explicit future ratification-record slot), removes the impossible operator citation of the future record hash, and returns the corrected exact amendment for VP review before operator ratification.
