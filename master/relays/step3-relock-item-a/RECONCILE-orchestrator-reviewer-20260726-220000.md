## RECONCILE -- REVISE: r3 fixes the missing join and weights, but freezes unresolved hashes and hashes superseded/foreign-owned source bytes

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r3
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- master can correct the extraction and freeze recipe within the ratified contract; owners remain held
GRILL_REQUIRED: no -- the operator-ratified Section 4/Section 7 choices are fixed
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-210000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE recipe 06e6956e -- do not hash placeholder fixture digests before T4, distinguish settled-basis from post-marker source hashes, replace whole-document owner spans that include foreign/stale bytes, and close the actual JSON unions; owners remain held

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-210000.md` at SHA-256 `c64d969359c3fe994c5e8fe69a5427a217e4cfc7aea0200d0434dd6ba33536de`.

Recipe reviewed: `master/STEP-3-ITEM-A-RECIPE.md` r3 at SHA-256 `06e6956e1c2c591d6cf0a322971ca250a66957c82c0ad09b16da927591033419`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, especially Sections 4, 7, and 11.

## Findings

### ITEM-A-VP-R3-F1 -- BLOCKER: the whole-file exit-fixtures interface is frozen and locked before its required digests exist

Recipe `:56-59` makes `STEP-3-EXIT-FIXTURES.json` a `whole_file` HARD interface, so every byte in that file contributes to `if.exit-fixtures.extracted_sha256` and therefore `bundle_sha256`. Recipe `:65-67` freezes that file before bundle assembly, and `:111-112` then re-locks the resulting bundle before T4.

But `:96` says `input_artifact_sha256`, `baseline_artifact_digest`, and `baseline_config_digest` are unresolved required-field slots that T4 will fill, and `:112` again places that fill after lane 4. Filling any slot changes the whole-file SHA, changes `bundle_sha256`, and invalidates the just-completed re-lock. Leaving placeholders instead violates ratified Section 7 `:377-387`, which defines those digests as fields of the frozen manifest, and contradicts r3's own rule that assembled review rejects every unresolved digest/placeholder. The claimed no-circularity ordering is therefore impossible as written.

Required correction: produce and content-address every fixture input plus the baseline artifact/config before freezing `STEP-3-EXIT-FIXTURES.json`, bundle assembly, and lane 4. T4 may build the executable fixture implementation against those already-frozen inputs; it may not fill a hash-bound manifest. If those artifacts genuinely cannot exist before T4, route an explicit sequence/schema amendment through VP and operator rather than silently creating mutable slots.

### ITEM-A-VP-R3-F2 -- BLOCKER: the owner rows still are not literal verifier inputs, and their hashes become stale when markers are inserted

Recipe `:35-44` calls its paths literal, but every row starts `domains/...`; from the governing workspace cwd all eight paths are absent, while the files live under `master/domains/...`. No alternate path base is defined. The purported exact heading anchors also contain abbreviated `...` text rather than the full literal on-disk headings, and `:32` permits owners to place markers merely "within" the range instead of at one fixed boundary.

The `full source_sha256` values do match the eight current unmarked files, but marker insertion necessarily changes each full-file hash. Ratified provenance stores the **post-marker on-disk** `source_sha256`, and `--verify` must reject a mismatch. R3 neither distinguishes the listed pre-marker settled-basis digest from the final provenance digest nor proves that marker placement changed only the two marker lines. Keeping the listed SHA makes every released source fail verification; replacing it after marker placement silently loses the mechanical link to the settled bytes.

Required correction: use literal workspace-relative `master/domains/...` paths (or define and verify one explicit path base), full exact heading strings, and fixed insertion boundaries. Rename the current values `settled_basis_sha256`; compute the final post-marker `source_sha256` separately; and fail closed unless removing exactly the declared marker lines reproduces the settled-basis bytes/hash. A dedicated derived extraction artifact is also valid if its derivation from the settled basis is exact and reviewed.

### ITEM-A-VP-R3-F3 -- BLOCKER: one interface per source file does not establish semantic single ownership and includes known superseded state

Recipe `:32` infers "no consumer hashes a producer" from "one interface per source." File ownership does not imply ownership of every formula copied into that file. The proposed m-9 span (`:42`, Sections 1 through 9) includes m-9's own ownership table stating that the component recipe and three relay schemas are **m-2-owned** (`m-9 delta :434-444`), then reproduces the m-2-authored `relay.submit` formula (`:475-485`). Those same producer semantics are also hashed in the two m-2 interfaces. The producer formula therefore does not appear exactly once under its owner.

The same m-9 span includes Section 9's live text that the m-10 C ticket is `PARKED` (`:496-510`), while the lane-2 close explicitly rules that item discharged by the settlement trail and says the parked line is superseded (`step3-relock-settlement-amend/...-160000.md:20-23`). The bundle would hash the stale PARKED statement and the current master `if.join.item-c` binding together without any precedence rule. Conversely, the m-3 span starts at Section 1 and omits Section 0a, even though Section 0a calls itself the single source of producer truth and the selected live sections cite it rather than restate it (`m-3 delta :15-25`).

Required correction: do not use broad mixed-document spans as a proxy for semantic ownership. Return to exact non-overlapping current-contract regions, or author dedicated owner interface artifacts containing only that owner's current formulas plus typed references to producer interface IDs. Exclude superseded status prose and make current settlement/join records the sole home of trail-only state; ensure every selected reference resolves inside the bundle.

### ITEM-A-VP-R3-F4 -- BLOCKER: the claimed closed fixture schemas are internally open

Recipe `:78-79` gives a field list plus optional "typed extensions" but no discriminator specifying which fields are required or forbidden for each fixture, and it leaves `fault_injection_point`, `expected_canonical_rows`, `observer_id`, and locator shapes unconstrained. Recipe `:101-105` calls `carried_records` closed with base fields `{carried_id,gating,disposition,source_locus,fixture_binding}`, then adds `reopen_predicate` for `r7_mirror`, a field absent from that closed shape. `fixture_binding` alternates between a two-fixture composite, a sub-vector, and `none`, while `disposition` is prose for two records and an enum token for the third.

Required correction: publish the actual closed schema or canonical discriminated unions used to author the JSON: exact field types, enums, required/forbidden members per fixture/carried ID, extension applicability, and `additionalProperties=false` semantics. The final manifest may supply concrete values later in Item A, but its author must not invent the data model after this release gate.

## Passed portions

- The Section 4 top-level/hash-domain repair remains correct.
- The fifth normative join, `if.join.b-carriage`, is now present, and a whole-file interface is a valid way to digest carried records once the manifest is actually final.
- The durability predicate references now bind the exact ratified rev12 bytes and preserve the substantive missing-half, receipt, and overflow cuts.
- The concrete sample-weight table sums to exactly 30 governed turns and 100 tool calls.
- The overhead budget is correctly recorded as already operator-ratified and immutable for T4.
- The addressed hold remains effective and leak-free; no owner marker or nomination relay exists.
- Item A remains produce-not-lock and grants no code, provider, merge, deploy, or external-use authority.

## Gate disposition

- Keep all six owners held. Do not issue the addressed RELEASE from recipe `06e6956e...`.
- Master re-cuts the freeze ordering and source strategy before another VP review.
- The later assembled-artifact review still must see concrete, non-placeholder fixture/baseline digests, final post-marker provenance, the exact extractor/negative fixture, and the closed manifest instance before lane 4.

## Verification

- Recomputed exact hashes: target `c64d9693...`; recipe `06e6956e...`; ratified rev12 `1125b0a0...`.
- Exact-file lint is `OK` for the incoming relay.
- Recomputed all eight owner-source hashes; each r3 full digest matches the current unmarked file. Confirmed zero owner-source HARD markers and confirmed all eight recipe paths are absent relative to the workspace cwd while their `master/domains/...` paths exist.
- Read every on-disk `##` heading in the eight sources; r3 uses abbreviated, non-literal anchor text. Inspected the selected m-3 and m-9 ranges against the lane-2 close and ownership declarations.
- The Item-A tree contained no owner nomination/marker relay before this review.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no recipe, amendment, owner design, settled base, `frank/` source, branch, commit, release, marker, bundle, extractor, fixture, lock, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260726-220000.md`.
Next requested action: master re-cuts Item A so every hashed input is final before lane 4, every source cut is literal and current, semantic ownership is unique, and the JSON unions are actually closed; owners remain held pending the next VP approve plus an addressed release.
