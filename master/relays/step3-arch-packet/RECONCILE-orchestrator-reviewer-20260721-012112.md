## RECONCILE -- REVISE: H-16 and the m-9 pins pass, but F91 and F96 remain open; the stage-6 lock stays held

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-interface-lock-review-r2
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator's stage-6 gate remains held because the census is not yet the canonical machine-readable full-row artifact and the packet is not yet the deterministic exact-path/full-hash manifest required by F96
GRILL_REQUIRED: no -- the remaining defects are master-owned census serialization/splitting and evidence-manifest corrections; the accepted owner designs, H-16 join, and P4/P5 product disposition do not need another grill
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-005128.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- preserve F92-F95, F97-F100, the nine design hashes, and H-16 rev16; rebuild H-17 as one canonical exact-field table and replace every F96 shorthand with an exact path plus full SHA-256 before returning lock-review r3

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-005128.md` at SHA-256 `9f4be88ac7e4ca94844021809f49a3b6623af7837717787c66c6a8a54701114e`.

## Findings

### F91 -- BLOCKER REMAINS: the rebuilt census is 39 unique headers, but it is not the required canonical machine-readable full-row table

Several F91 corrections are real and preserved: current `master/H17-CENSUS.md` hashes to `54208535e50723924cc8b61bc254757b5750574d7caf95db26f24adeead114d7`; a mechanical header scan finds 39 rows and 39 unique IDs; the duplicate `m8-provider-send` row is gone; merge/release/deploy are outside the effect set; and `conductor-relay-accept` now carries the H-16 rev16 monotonic outcome split.

The artifact still fails the contract in three independent ways:

1. It is not one table with one canonical row grammar. `master/H17-CENSUS.md:7` expressly substitutes two special mappings for the m-9 rows: the row header is treated as `effect_id`, and one combined `policy_owner/policy_artifact` field is split by convention. A mechanical schema-key scan confirms that all 16 m-9 rows lack a named `effect_id` field and lack separate `policy_owner` and `policy_artifact` fields. The prose mapping above the rows does not make each row full-field under `master/H17-CENSUS-SCHEMA.md:9-32`, and it leaves a parser dependent on section-specific grammar.
2. Required `policy_artifact` cells are not exact documents/digests. The schema requires the exact document/digest at `master/H17-CENSUS-SCHEMA.md:18`. Counterexamples include the abbreviated H-16 digest and `frank@6e4d657` composite at census `:22`, the unbound `s11-close` label at `:24`, and the truncated r40/amendment digests at `:36`. Section-level provenance does not repair row-local fields that cite several different policy artifacts.
3. The prior non-append split is still incomplete. Census `:24` combines conductor-side `project`, `read`, and `Describe` into `conductor-seat-serve`, despite F91's requirement to split every non-append verb. Those operations have different disclosure targets and serve semantics. Census `:30` similarly combines genesis/migration and GC while supplying only the genesis/migration linearization fields; its own replay cell distinguishes once-per-store genesis from idempotent GC. Those are not one effect row at the required grain.

Required correction: emit one canonical machine-readable table whose columns are the schema-v1 fields, one row per effect, with each cell exact or a schema-legal null token. Normalize the owner material into that table without moving the approved owner-design bytes; carry full exact paths/digests in every `policy_artifact` cell; split conductor `project`, `read`, and `Describe`; and split genesis/migration from GC unless one exact shared linearization/record contract can honestly cover both. Recompute the census hash and rerun the row-count, uniqueness, required-field, and full-digest checks.

### F96 -- BLOCKER REMAINS: the packet repeats the shorthand locators the deterministic-manifest correction prohibited

Target `:1,18` says every supporting record is enumerated by exact path plus SHA-256 and that this packet is the lock's sole locator. The body does not meet that claim:

- The nine-design table gives full hashes for designs but no hash for most final-review relays, and its review paths omit the `master/relays/` root (`:24-32`). The m-8 basis verdict is itself only a truncated digest (`:29`).
- The two m-9 records use truncated hashes followed by "full: recompute at file" (`:61-62`). A deterministic manifest must carry the digest; it cannot delegate reconstruction to its consumer.
- The six grills are timestamp fragments or in-document labels with no exact path/full hash (`:68`). The amendment itself is only `2f75f2a1...` with no exact path/full digest (`:53`).
- N1-N4, P1-P3, the L-ledger, and H-26 are labels or abbreviated references rather than exact records (`:71`).
- The confirmation set still uses 8-character prefixes and two bare `090000` locators (`:74`). There are in fact two files named with that timestamp in different lanes, exactly the ambiguity F96 called out.
- The inherited amendment/gate records required by F96 are not enumerated as exact-path/full-hash rows.

Required correction: replace the narrative manifest with one complete record table. Every design, final approval, grill, operator decision, pin/erratum source, consumer confirmation, reciprocal, m-9 certification, H-16 join leg, schema, census, ratified amendment, and inherited gate must have its exact `master/...` path and current 64-hex SHA-256 in the packet itself. Timestamp shorthand and hash prefixes may be display aliases only after the exact locator is present.

## Accepted evidence preserved

- All nine design artifacts recompute to the exact hashes in target section 1; no accepted design byte moved.
- F92 is closed at semantic grain by the actual m-9 Implementer certification `master/relays/step3-mvp-stage4-m9/RECONCILE-implementer-20260720-223000.md`, current SHA-256 `c59d4e6b5dfdfe59b21690bd4ed3d299e2d95e44d99d20501b08081f8da759e8`, and exact-file lint is OK.
- F94/P4 is closed: stage 6 binds only the interface identity contract plus expected catalog vector; actual build and release digests bind postbuild before E3.
- F95/P5 is closed at product-decision grain by the operator-cited option-1 record `master/relays/step3-mvp-stage4-m9/RECONCILE-orchestrator-planner-20260720-214811.md`, current SHA-256 `8e4eebe64c72835f67ca4f707bab23cadeb4ef0a285745f784514d1943c7308a`. The runnable worker digest includes linked bytes; the separate shared-client component digest supplies attribution; both enter release binding; the dependency-insensitive `iff` claim is withdrawn.
- F93 and F97-F100 remain closed at H-16 design-contract grain. Rev16 still recomputes to `a349a3297c608915510e7bf6de2e304b824897aae4184f7c71db928149ae316d`; the pair approval, both owner confirmations, master half, and VP half all recompute to target section 2's exact hashes. H-26 remains a separately scoped open code defect.
- The canonical H-17 schema still recomputes to `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`; this verdict rejects the census serialization/content against that schema, not the schema bytes.

## Gate disposition

- Stage-6 joint interface lock: HELD.
- Operator ratification of that lock: NOT REQUESTABLE from this packet.
- T4 PM/PLAN/code token, H-16 implementation, credentials, provider calls, release binding, live E3, merge, and deploy: remain separately HELD under the existing sequence.
- Correct only F91 and F96 at master grain. Do not move the nine approved design artifacts, H-16 rev16, owner confirmations, or the P4/P5 records.
- Step 2 remains closed.

## Required return

Return lock-review r3 only after the canonical H-17 table and the self-contained exact-path/full-SHA manifest exist on current bytes. Bind the returned packet to the new census hash and preserve the accepted closure set above.

## Verification

- Target path is directly addressed to this seat, indexed, exact-file lint-clean, and hashes to `9f4be88ac7e4ca94844021809f49a3b6623af7837717787c66c6a8a54701114e`.
- The nine design hashes, H-16 rev16, five H-16 approval/join records, H-17 schema, H-17 census, m-9 certification, and operator P5 record were recomputed from current disk bytes.
- Census mechanical scan: 39 row headers, 39 unique extracted IDs, zero duplicate IDs; 16/16 m-9 rows lack the canonical named `effect_id`, separate `policy_owner`, and separate `policy_artifact` fields.
- The two bare `090000` files were resolved live in distinct lanes; their ambiguity is current, not historical.
- `frank/` remained read-only at `6e4d657913229027fc94a1e2a8c2348b05c09a75` for this review.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master rebuilds H-17 at the canonical one-row-per-effect table grain, emits the actual F96 exact locator table, and returns only those corrected current bytes for lock-review r3.
