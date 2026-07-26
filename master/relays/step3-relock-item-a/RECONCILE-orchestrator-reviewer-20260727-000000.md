## RECONCILE -- REVISE-NARROW: the plain record-lock direction is viable, but the amendment surface, fixture ordering, and lock manifest are incomplete

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r4
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- after the bounded corrections and VP approval, the operator must ratify the exact amendment hash; the directive is not self-ratification
GRILL_REQUIRED: no -- the operator selected the simplification; this review checks whether its amendment and downstream gates are complete
DESIGN_DOC_ID: step3-relock-item-a-simplification
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-230000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW simplification amendment 680e6fcb -- keep the plain byte-lock, but supersede every ratified bundle dependency, resolve fixture digests before lane 4, include the governing amendments/precedence in the lock manifest, and choose one carried-obligation boundary

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-230000.md` at SHA-256 `a284a3f8477612334e20d4ba7f06f7e3133c8b4a6b2bbd2647fd73b5f5831683`.

Amendment reviewed: `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` at SHA-256 `680e6fcb930a1fc0f2f6c04dd02a0dd8c76d98710927c4d3ef4ee27b2b8c9476`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.

## Findings

### ITEM-A-VP-R4-F1 -- BLOCKER: replacing Section 4 alone leaves three ratified bundle dependencies live

Amendment `:4` says it amends only Section 4. The controlling bytes also require the old mechanism in:

- Section 6 `:359`: Item A is the extraction recipe + bundle authored over B-E.
- Section 11 `:424-427`: author the recipe/bundle and stability fixture, then re-lock over `bundle_sha256`.
- Section 12 `:432-435`: VP must specifically validate `bundle_sha256` provenance exclusion and soft-edit stability.

Those clauses remain operative if only Section 4 is replaced. The amendment's prose sequencing at `:31-32` does not identify them as superseded, so lane 4 would simultaneously require a plain record and the withdrawn bundle hash. The live `ROADMAP.md`, `master/ARCHITECTURE.md`, and `master/README.md` also still define Item A and lane 4 in terms of `bundle_sha256`; several domain dashboards route their next action to that same mechanism.

Required correction: make the amendment's supersession surface explicit and exact: replace Section 4, Section 6's Item-A edge, Section 11 steps 4-5, and the obsolete Section 12 bundle-specific VP criterion. State the replacement lane-4 predicate over `STEP-3-INTERFACE-LOCK.md`. Include a post-ratification source-fold manifest for `ROADMAP.md`, `master/README.md`, `master/ARCHITECTURE.md`, affected domain dashboards, and the withdrawn recipe status so the architecture-of-record cannot keep routing two mechanisms.

### ITEM-A-VP-R4-F2 -- BLOCKER: un-fusing the fixture manifest from Item A does not resolve its before-T4 digest contradiction

Amendment `:19` says the fixture-input digests do not exist until T4, while `:31-32` still freezes and hashes `STEP-3-EXIT-FIXTURES.json` in lane 4 before lane 5/T4. Ratified Section 7 `:377-387` requires concrete per-fixture `input_artifact_sha256` values and concrete baseline digests in that frozen manifest. Moving the manifest out of the interface lock changes which digest moves; it does not make those future values exist before the re-lock.

Required correction: choose one coherent ordering. The narrow path is to author and content-address the immutable fixture-input and baseline artifacts before lane 4, freeze the manifest with final non-placeholder digests, then let T4 build the executable fixtures against them. If those artifacts truly cannot exist before T4, this amendment must also change Section 7's manifest contract and/or Section 11's lane order through the same VP/operator gate. Do not describe the current un-fusing as dissolving the circularity.

### ITEM-A-VP-R4-F3 -- BLOCKER: the proposed simple lock omits governing bytes and has no conflict-precedence rule

Amendment `:21-27` lists eight current owner files, five joins, frozen finals, and carried statements. The lane-2 close says the settled interface is also underpinned by three governing amendment packets: stage-6 rev12 `1125b0a0...`; the m-3 schema amendment `9e874df8...` plus bound contract `6e2abe40...`; and the Section-D settlement amendment `1fa71cb8...` plus the m-2 cell (`...-160000.md:37`). The proposed lock omits those amendment artifacts even though the target itself says they remain unmoved.

Whole-file owner bases also contain authorship-time status later superseded by the join/settlement trail, notably m-9 r17's PARKED C-ticket line. A plain hash list can safely lock that history only if the record states which exact settlement/join record governs conflicting status. Semantic names such as "item-C" or "Section-D join" are not enough to reproduce the lock.

Required correction: define the lock record as a closed manifest of literal path/relay + full SHA-256 + role (`owner_base`, `frozen_final`, `governing_amendment`, `join_or_settlement`, `carried_source`) and include all three governing packets. Enumerate every join/co-sign/confirmation relay by exact path and hash. Add explicit precedence edges saying later governing amendments and settlement/join records supersede named historical status clauses without changing the owner-base bytes. Bind the final record itself at its review hash.

### ITEM-A-VP-R4-F4 -- BLOCKER: carried obligations are simultaneously inside and outside the interface lock

New Section 4 at amendment `:9` says the lock record byte-binds the two carried limits and env-digest parity locus. Amendment `:18-19` then says the carried obligations were never part of the interface lock and belong to the lane-4 exit-fixture manifest. Both cannot define the single source of truth.

Required correction: choose one boundary. Recommended: the interface lock records exact source-relay hashes for the accepted N910/r7/env-parity design dispositions as governing lineage, while lane 4 alone owns their executable fixture records and expected rows. State that split explicitly; do not duplicate free-form obligation text in both artifacts.

## Passed portions

- The operator-directed plain byte-bound record is a viable MVP mechanism if ratified. It matches the project's existing exact-hash approval discipline and may deliberately withdraw soft-edit stability because the named design artifacts are frozen.
- The amendment correctly uses VP review followed by operator hash-bound ratification; master does not self-ratify.
- Removing marker insertion, the extractor, and per-interface extraction artifacts eliminates the r3 source-boundary and post-marker provenance defects.
- No owner action is needed for the simple lock, and the addressed hold remains effective until the amendment is ratified.
- Produce/lock/T4 authority remains separated; no code, provider, merge, deploy, or external-use authority is present. H-12 stands.

## Gate disposition

- Keep owners held until the corrected amendment is VP-approved and operator-ratified.
- Preserve the plain record-lock direction; re-cut only the amendment completeness, lane ordering, and lock manifest contract.
- Do not author or approve `STEP-3-INTERFACE-LOCK.md`, release Item A, or enter lane 4 from amendment `680e6fcb...`.

## Verification

- Recomputed exact hashes: target `a284a3f...`; amendment `680e6fcb...`; ratified rev12 `1125b0a0...`; m-3 schema amendment `9e874df8...`; settlement amendment `1fa71cb8...`; lane-2 close `fa2a634f...`.
- Exact-file lint is `OK` for the incoming relay.
- Scanned current non-relay governance sources for `bundle_sha256`, the bundle path, extractor, stability fixture, lock path, and exit-fixture path. Confirmed live bundle dependencies in the ratified amendment, roadmap, architecture, dashboard, and domain status.
- Re-read the lane-2 close's exact owner, join, amendment, carried, and sequencing inventory.
- The Item-A tree contained no owner marker/nomination action and no simple lock record before this review.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, recipe, owner design, settled base, lock record, `frank/` source, branch, commit, release, fixture, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-000000.md`.
Next requested action: master preserves the operator-selected plain lock, expands the exact supersession/source-fold surface, resolves fixture digests before lane 4, completes the lock manifest and precedence graph, and returns the corrected amendment for VP review before operator ratification.
