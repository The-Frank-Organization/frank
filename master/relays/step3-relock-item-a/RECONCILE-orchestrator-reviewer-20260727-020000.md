## RECONCILE -- REVISE-NARROW: rev2 closes the prior four directions, but the proposed lock cannot self-hash, its operative lane-4 replacement reverses fixture order, and its manifest/source-fold sets remain open

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r5
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- after the bounded corrections and VP approval, the operator must ratify the exact amendment hash; the directive is not self-ratification
GRILL_REQUIRED: no -- the operator selected the simplification; this review checks whether its amendment and downstream gates are complete
DESIGN_DOC_ID: step3-relock-item-a-simplification
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-010000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW simplification amendment c99cd78e -- keep the plain byte-lock and corrected fixture model, but bind the record externally, make lane-4 order single, close the literal manifest, and complete the source-fold set

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-010000.md` at SHA-256 `82c8397bfd9023d9ab05a852a3e5a07f98e3967f3c0525a34cd59289c62c3e50`.

Amendment reviewed: `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev2 at SHA-256 `c99cd78e806aa60a0fc80f5e78786d96c6de95766ee100aa3a9e762ec69dd35c`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.

## Findings

### ITEM-A-VP-R5-F1 -- BLOCKER: a final file cannot practically contain its own SHA-256

Amendment `:47-48`, `:68`, `:111-113`, and `:126-128` require `STEP-3-INTERFACE-LOCK.md` to state and be ratified over its own final hash. Writing hash X into the file changes the bytes whose SHA-256 was X; recomputing and replacing X changes them again. This is not the exact-hash discipline used by prior approvals, where a separate review or ratification relay names the finalized artifact hash.

Required correction: remove the literal self-hash field and every "self-binding hash" dependency. The finalized record states the generic invalidation rule ("any change to this record or a named artifact voids the lock"), while the VP/interface-lock relay names the record's full SHA-256 externally. Any later ratification record cites that same external hash. No fixed-point construction or placeholder is permitted.

### ITEM-A-VP-R5-F2 -- BLOCKER: the operative Section 11 step-5 replacement reverses the corrected fixture order

Amendment `:47` replaces ratified Section 11 step 5 with interface-lock ratification over the record and **then** fixture-manifest freeze. Amendment `:63-69` and `:126-128` correctly require the opposite: author/content-address fixture inputs and baselines, freeze the final manifest, then re-lock over both the record and manifest. Because `:47` is the text that supersedes the controlling clause, the amendment currently installs the stale order while its rationale describes the repaired one.

Required correction: make every sequencing locus identical: lane 4 authors and content-addresses immutable fixture inputs/baselines -> freezes `STEP-3-EXIT-FIXTURES.json` with final digests -> Master+VP re-locks over the externally hashed interface-lock record plus that frozen manifest -> lane 5/T4 builds executable fixtures and fills no bound slot.

### ITEM-A-VP-R5-F3 -- BLOCKER: the claimed closed manifest still depends on unlisted and ambiguous bytes

The record's included rev12 bytes still prescribe the bundle. The amendment that supersedes those bytes, plus the operator relay that makes that amendment operative, are absent from the required manifest at `:86-102`. Named bytes alone therefore do not reproduce the new governing mechanism or prove that the authored amendment was ratified.

The same section also describes five conceptual joins using ellipses and grouped legs, leaves item-E's binding relay to be "named at authoring," does not name the three `carried_source` relays, and gives the lane-2 close/correction no explicit role. That is not yet the literal path-plus-full-hash enumeration required by `:75-84`; a conceptual seam can contain several independently hashed relay files. Finally, `:108-109` says a "later record" generally governs without defining later by an explicit edge, and can let an unrelated later join override any owner status.

Required correction: add the ratified simplification packet itself (amendment plus its authority-bearing ratification relay) to the role model; enumerate one manifest row per actual artifact/relay, including every join leg, close/correction record, and carried-source relay, at exact path and full hash; and replace chronology-based precedence with explicit `(source path/hash + named clause) -> (superseding path/hash)` edges. Manifest order and filename time must carry no authority.

### ITEM-A-VP-R5-F4 -- BLOCKER: the post-ratification source-fold manifest omits three live old-mechanism routes

Amendment `:50-56` includes the roadmap, dashboard, architecture, m-3 dashboard, and withdrawn recipe. The current source scan also finds operative old-sequence text in:

- `master/CYCLE-PLAYBOOK.md:408` -- recipe/bundle authoring and shorter re-lock;
- `master/domains/m-1-trust-identity/README.md:111` -- next engagement is the extraction bundle;
- `master/domains/m-2-forms-determinism/README.md:59` -- re-engage at the extraction bundle.

Required correction: add those three files to the post-ratification fold manifest. Historical relay/ledger entries remain append-only history and need no rewrite.

## Passed portions

- Rev2 now explicitly supersedes every operative rev12 bundle locus: Section 4, Section 6's item-A edge, Section 11 steps 4-5, and the bundle-specific Section 12 VP criterion.
- The substantive inputs-before-freeze fixture model is coherent and fixes R4-F2 once the contradictory step-5 replacement is corrected. T4 fills no hash-bound slot.
- The carried-obligation boundary is now single: the interface record carries exact disposition lineage; lane 4 alone owns executable fixture records and expected rows.
- The role-tagged whole-file lock direction, explicit historical-status edges, and exact-byte invalidation rule are sound once the literal set and external record binding are closed.
- The amendment still preserves VP review -> operator hash-bound ratification, leaves ratified/frozen owner bytes unmoved, and grants no design-lock, PLAN, T4, credential, provider, release, E3, merge, deploy, or `frank/` authority. H-12 stands.

## Gate disposition

- Keep owners held until a corrected amendment is VP-approved and operator-ratified.
- Preserve the operator-selected plain record-lock direction and the corrected fixture-input model.
- Do not ratify amendment `c99cd78e...`, author or approve `STEP-3-INTERFACE-LOCK.md`, release Item A, enter lane 4, or issue T4 from these bytes.

## Verification

- Recomputed exact hashes: target `82c8397b...`; amendment `c99cd78e...`; prior VP relay `8c3a85b...`; ratified rev12 `1125b0a0...`; m-3 schema amendment `9e874df8...`; settlement amendment `1fa71cb8...`; lane-2 close `fa2a634f...`; env-locus correction `84c6d6ab...`.
- Exact-file lint is `OK` for the incoming relay.
- Re-read the amended supersession, fixture-ordering, manifest, precedence, carried-boundary, sequencing, and authority clauses at the exact bytes.
- Scanned the live roadmap, architecture, dashboard, cycle playbook, and domain dashboards for the withdrawn bundle mechanism; confirmed the three omitted source-fold files above.
- `master/STEP-3-INTERFACE-LOCK.md` remains absent, so no premature lock artifact or owner action landed.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, recipe, owner design, settled base, lock record, fixture, `frank/` source, branch, commit, release, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-020000.md`.
Next requested action: master preserves the plain record-lock, removes in-file self-hashing in favor of external relay binding, makes the lane-4 order identical at every locus, closes the literal artifact/relay and precedence sets (including this amendment's ratification packet), completes the source-fold manifest, and returns the corrected exact amendment for VP review before operator ratification.
