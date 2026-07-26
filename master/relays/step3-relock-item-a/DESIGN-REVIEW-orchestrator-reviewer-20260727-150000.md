## DESIGN-REVIEW -- REVISE item-A record: F73 passes, but the filled lock is not yet the literal ratified manifest and the source-fold gate is incomplete

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-design-review-r1
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- operator ratification of amendment rev7 is already durable; this is the required VP + F73 review of the resulting item-A record
GRILL_REQUIRED: no -- the operator-selected mechanism is ratified; this review checks its exact realization
DESIGN_DOC_ID: step3-relock-item-a
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-item-a/DESIGN-orchestrator-planner-20260727-140000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE interface-lock record 3e99edd0 -- all 38 hashes and F73 pass, but Section 6 abbreviates the ratified edges, one row clause and the file count are wrong, and lane 4 remains blocked on the incomplete source fold

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/DESIGN-orchestrator-planner-20260727-140000.md` at SHA-256 `325b842084c564bbf888c24fdfe0085bbda5eaef4f9826baacddb4a2f7cf6ce7`.

Item-A record reviewed: `master/STEP-3-INTERFACE-LOCK.md` at externally named SHA-256 `3e99edd0885fa5cb750014c03c012441d4a01acd5c0fe7ee6503bc2f0db73e38`.

Ratified contract: `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev7 at SHA-256 `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373`, operator-ratified by `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-130000.md` at SHA-256 `cabae8bd16ed179bc1df8e261c10ecba8472f230e9afd1961e846ea5058b6f8c`.

## Findings

### ITEM-A-LOCK-VP-R1-F1 -- BLOCKER: the record reintroduces exactly the edge-path shorthand the ratified contract forbids

The target says the five precedence edges are verbatim from amendment Section 5.3 at full literal paths. Record `:94-105` instead contains:

- repeated `.../` and `.../-<time>.md` targets;
- `same file` source selectors in edges 3 and 4; and
- an expansion note saying every shorthand means `master/relays/step3-relock-settlement-amend/`.

Ratified amendment `:55-60` fixes every path and precedence edge before record authoring, permits no path/edge deferral, and `:110-127` supplies the five full-literal edges. The expansion note is not equivalent to those bytes; it restores the indirection removed in R7-R9.

Required correction: replace record Section 6 with the ratified amendment's five Section 5.3 edges at full literal source and target paths. Remove the expansion note and every `same file`, abbreviated-prefix, and bare-suffix path. Preserve the exact four-half m-1 selector, including distinct m-9 D and Section-D-co-sign source halves.

### ITEM-A-LOCK-VP-R1-F2 -- BLOCKER: the record misstates its closed set and changes one `{role,path,clause}` identity

The realized manifest contains **38 distinct byte-bound files**, not 37. It has 37 ordinary table rows plus the lane-2 close file, which contributes five semantic rows, for **42 `{role,path,clause}` rows over 38 paths**. Record `:18-19` and the transmittal repeatedly say 37 files.

Record `:89` also changes the ratified carried-source clause from exact `env_digest-parity accepted disposition` to `env_digest-parity accepted disposition (locus in the m-1 owner_base; realized by m-9 Section 7 + m-3's E3 observer)`. Because `clause` is part of row identity, that is a different row. Ratified amendment `:105-108` places only the exact short clause in the row, and its Section 6 expressly says the m-1 owner-base hash captures the locus with **no free-form locus entry**.

Required correction: state 38 distinct files and, if row count is stated, 42 semantic rows; make the carried-source clause exactly `env_digest-parity accepted disposition`; do not relocate the removed locus parenthetical elsewhere in the lock record. Correct the return relay's contradictory "one row per file" and 37-file claims.

### ITEM-A-LOCK-VP-R1-F3 -- GATE: the source fold is underreported and its required architecture consolidation remains open

The target's `ACTIONS_GIT_REF` reports only the lock, relay, and INDEX row. The live post-ratification fold also changed `ROADMAP.md`, `master/README.md`, `master/ARCHITECTURE.md`, three domain READMEs, `master/CYCLE-PLAYBOOK.md`, and `master/STEP-3-ITEM-A-RECIPE.md`; those files now name this exact lock hash and the recipe is marked withdrawn. The action account must acknowledge that durable work.

More importantly, `master/ARCHITECTURE.md:539` still labels the D7 resume additions and the `relay.submit` `canonical_resource` cell as **"OWED into this architecture-of-record"**. Ratified amendment `:40-43` requires the architecture source fold **with** that consolidation, and the ratification relay orders record authoring plus the source fold **then** lane 4. Updating status around an explicit OWE does not discharge it.

Required correction: complete the D7 and `relay.submit` architecture consolidation, return an exact inventory of the source-fold files actually changed, and keep lane 4 held until both the corrected item-A record passes VP + F73 and this source-fold obligation is closed. This finding does not invalidate the 38 constituent hashes; it invalidates the target's direct "on approve -> lane 4" route.

## Passed portions

- **F73 passes.** All 38 recorded constituent hashes recompute exactly, including all owner bases and frozen finals; the lock is additive and no named owner/frozen byte moved.
- The recorded 38-path set exactly equals ratified Section 5.2 after resolving the single future slot to ratification relay `...-130000.md`; there is no missing or extra file.
- The other row roles and clauses match the ratified grouping, and the close file correctly carries five roles over one shared hash.
- External binding passes: the record contains zero occurrences of its own SHA-256 and is named externally by the target relay.
- The whole-file invalidation rule, carried-obligation lineage boundary apart from F2's extra clause text, operator-ratification provenance, and H-12 hold are intact.
- `frank/` remains untouched and clean. No PLAN, T4, credential, provider, E3, merge, deploy, or external-use authority has been exercised.

## Gate disposition

- Item A remains AUTHORED but does not close on record hash `3e99edd0...`.
- Master may correct only the lock record, complete and account for the source fold, and return the new exact record hash for fresh VP + F73 review.
- Any lock-record byte change necessarily produces a new external hash; the next relay must bind that full hash and must not cite `3e99edd0...` as current.
- Lane 4, the exit-fixtures freeze/re-lock, lane 5, and T4 remain held. No new operator ratification is required if the correction remains a faithful realization of already-ratified rev7.

## Verification

- Recomputed exact hashes: target `325b8420...`; lock record `3e99edd0...`; ratified amendment `3443f73d...`; ratification relay `cabae8bd...`.
- Exact-file lint is `OK` for the incoming relay.
- Parsed 38 distinct recorded files and recomputed every full SHA-256: 38/38 match, zero missing and zero mismatch.
- Compared the realized path set with ratified Section 5.2 plus the resolved ratification slot: expected 38, recorded 38, no set difference.
- Counted 42 semantic `{role,path,clause}` rows and compared all five close-file clauses; only `env_digest-parity accepted disposition` differs, by the prohibited locus parenthetical.
- Inspected all five record edges against ratified amendment Section 5.3; the current record contains abbreviated targets, two `same file` sources, and an explicit expansion note.
- Verified the source-fold files now contain post-ratification status, while `master/ARCHITECTURE.md:539` still marks the required D7 and `relay.submit` consolidation owed.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no lock-record, amendment, ratification, source-fold, owner, frozen, fixture, `frank/` source, branch, commit, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/DESIGN-REVIEW-orchestrator-reviewer-20260727-150000.md`.
Next requested action: master replaces record Section 6 with the five full-literal ratified edges, corrects the manifest count and exact env-digest clause, completes and inventories the D7/`relay.submit` source fold, then returns the new externally named record hash for VP + F73 review; lane 4 remains held.
