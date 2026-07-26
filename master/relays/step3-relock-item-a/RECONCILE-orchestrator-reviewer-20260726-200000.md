## RECONCILE -- REVISE: the constitutional hash is repaired, but the source inventory and exit-fixture freeze remain non-executable

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r2
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- master can make the bounded mechanical corrections below; owners remain held
GRILL_REQUIRED: no -- the operator-ratified Section 4/Section 7 choices are fixed and must not be reopened
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-190000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE recipe a98e85a1 -- Section 4 hash-domain repair passes, but replace shorthand sources with exact extraction inputs, restore the missing B-carriage and carried interfaces, and make the Section 7 fixture freeze executable; owners remain held

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-190000.md` at SHA-256 `2b41a9c84e01f7c9e602a757f047e568dd05669abf7924f404761691343cce5d`.

Recipe reviewed: `master/STEP-3-ITEM-A-RECIPE.md` at SHA-256 `a98e85a196099b2588d3a2da88c2df1ef1ca37e57c99480984e826b4e600edeb`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, especially Sections 4, 7, and 11.

## Findings

### ITEM-A-VP-R2-F1 -- BLOCKER: the declared source inventory is still not an exact extractor input

Recipe `:52-84` says master fixes `source_path` and establishes a complete, non-overlapping inventory, but the rows contain non-path placeholders such as `domains/m-1-...`, ditto marks, and section annotations inside the path column. `settled_basis` values are prefixes rather than full digests. The eight intended basenames do resolve uniquely on disk and their current hashes match the displayed prefixes, so the ownership direction is plausible; the table itself is nevertheless not data that `extract-interface-bundle.py` can consume.

The same section says exact span granularity is still "refinable at owner review" (`:84`). Several rows group broad or discontiguous material in one `region=marker` declaration, while other rows draw interfaces from the same sections. Until every interface has one exact canonical source and an unambiguous contiguous marker cut (or a dedicated `whole_file` artifact), the claimed non-overlap and no-foreign-byte properties are assertions, not decidable properties. Approval here would release owners to choose the hash boundary after the release gate.

The negative fixture is likewise not yet concrete. Recipe `:40-44` names a directory and the production `--verify` command, but it does not name a fixture manifest/input, mutation script and arguments, isolated output, or a command that proves both expected states without mutating a settled source. A directory plus "scripted mutation" is not the frozen mutation input required by r1 F2.

Required correction: expand every row to one literal workspace-relative path and full settled-basis SHA-256; identify the exact heading/byte anchors for each one-span marker or route a dedicated whole-file-hard extraction artifact; split discontiguous or overlapping selections. Specify the isolated soft-stability fixture files and exact commands for the baseline, SOFT mutation, HARD mutation, and expected digest comparisons. Marker placement still receives fresh owner/pair review, but the release recipe must bind where each marker is allowed to go.

### ITEM-A-VP-R2-F2 -- BLOCKER: the 27-row inventory omits a normative join and all carried-obligation bundle interfaces

The lane-2 close names five normative joins at `step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md:37`: Section D, Section B sink, item E, **B-carriage (m-8 -> m-10 -> m-3)**, and item C. Recipe `:79-82` authors only four master join rows; there is no B-carriage join. Owner-local `if.m8.frozen-core-digest`, `if.m9.b-carriage`, and `if.m10.be-carriage-row` do not replace the missing master integration record that binds the producer/carriers/evaluator.

The prior correction also required the three carried obligations to have owned, reviewed HARD sources if they affect the bundle. Recipe Part 2 has no `if.carried.*` interface and Part 3 places only prose outside `lock_payload`; sequencing `:113` assembles the bundle before it establishes any digest-bearing carried source. This does not satisfy the original Item-A scope ("bundle over ... the carried obligations") or r1 F1/F3's rule that such content enters the constitutional digest only through declared `{interface_id, extracted_sha256}` entries.

Required correction: add the master-owned B-carriage join source and extracted interface. Give N910, env-digest parity, and r7-mirror each a unique owned/reviewed HARD source/interface (or one closed carried-obligations source with separately decidable IDs), then sequence that source before bundle assembly. If `STEP-3-EXIT-FIXTURES.json` is itself the source, freeze it before the bundle and declare its exact interface regions; do not leave a circular "bundle first, fixtures second" dependency.

### ITEM-A-VP-R2-F3 -- BLOCKER: Part 3 summarizes the exit gate but does not freeze an executable fixture manifest

Recipe `:92-100` labels its predicates "verbatim Section 7", but durability at `:95` is a topic list. It omits the structured conditions fixed at ratified `:371`, including the exact valid-prefix/closed-manifest rule, both missing-half orders and omission mutants, duplicate/idempotence and `content_lost` cuts, the three pre-receipt crash cuts plus zero-work/observed-once rule, and the full frame-overflow terminal/no-successor/no-lease/no-snapshot/no-revival rule.

Recipe `:101` lists required field names but supplies neither a closed JSON schema nor the actual per-fixture `sample_weight` assignments. Saying future values will sum to 30 governed turns and 100 tool calls does not freeze those values. It also names no producers or frozen paths for the `input_artifact_sha256`, baseline artifact, or baseline config. Those artifacts do not currently exist, and sequencing `:113` has no step that creates them before their digests must be written. The prose list at `:103-106` similarly is not a closed typed `carried_records` schema: it has no record shape, enum closure, source/owner fields, or exact fixture-vector IDs/expected canonical rows.

Required correction: bind the full ratified predicates without lossy restatement; define the closed manifest and `carried_records` JSON shapes; assign each fixture's concrete `{governed_turns, tool_calls}` weight; route and order every content-addressed input/baseline artifact before manifest freeze; and define exact N910/env-digest vector IDs plus expected rows. The assembled-manifest review must reject any unresolved digest, placeholder, unowned observer, or arithmetic-only weight promise.

### ITEM-A-VP-R2-F4 -- BLOCKER: the recipe reopens an already-discharged operator decision on overhead

Recipe `:101` calls the overhead numbers "operator-ratifiable at the gate" and says they become immutable only "once ratified." The operator already ratified rev12, including the Section 7 numbers, in `step3-arch-packet/RECONCILE-orchestrator-planner-20260721-165500.md:18-27`. The current wording makes fixed thresholds appear adjustable at a future gate and conflicts with this relay's own `HUMAN_GATE_REQUIRED: no`.

Required correction: state that the `250 ms / 1000 ms / 50 ms` p95 ceilings and total-wall-clock bands are already operator-ratified and immutable for T4 at rev12. Any adjustment is a fresh addressed amendment/ratification, not part of Item A.

## Passed portions

- The r1 F101 defect is repaired. Recipe `:11-25` uses the exact ratified top-level schema and computes `bundle_sha256` only from JCS(`lock_payload`); mixed-source full-file hashes stay in provenance.
- Recipe identity, literal marker syntax, the master extractor path, `--verify`, and all ratified fail-closed classes are present at `:27-38`.
- The N910 gating-cut direction, env-digest parity direction and corrected m-1 Section 5 locus, and non-gating `deferred_v3` r7-mirror disposition are semantically aligned. They still need the exact typed/source binding required above.
- The addressed hold remains effective and leak-free. No owner nomination or marker relay exists in the Item-A tree.
- Item A still produces but does not lock; no DESIGN-lock, PLAN, T4 token, provider call, `frank/` action, merge, deploy, or out-of-envelope authority is present.

## Gate disposition

- Keep all six owners held. Do not issue the addressed RELEASE from recipe `a98e85a1...`.
- Master makes the four bounded corrections and returns one exact replacement recipe for VP re-review.
- The next approval remains decomposition/release approval only. The assembled bundle, extractor, fixture artifacts, and frozen manifest still require their own F73/VP evidence before lane 4.

## Verification

- Recomputed exact hashes: target `2b41a9c...`; recipe `a98e85a1...`; ratified rev12 `1125b0a0...`; lane-2 close `fa2a634f...`; ratification record `7c367c7f...`.
- Exact-file lint is `OK` for the incoming relay.
- Resolved all eight intended source basenames and recomputed their current full SHA-256 values; every displayed prefix matches, but the recipe rows remain abbreviated/non-literal.
- Scanned the Item-A tree: only `...-170000`, `...-171500`, reviewer `...-180000`, and target `...-190000` existed before this review; no owner nomination or marker placement landed.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain; HEAD and `origin/main` both equal `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no recipe, amendment, owner design, settled base, `frank/` source, branch, commit, release, marker, bundle, extractor, fixture, lock, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260726-200000.md`.
Next requested action: master re-cuts Item A to exact extractor inputs, a complete join/carried interface inventory, and an executable already-ratified exit-fixture freeze; owners remain held pending the next VP approve plus an addressed release.
