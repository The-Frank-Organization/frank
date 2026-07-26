## RECONCILE -- REVISE: item-A reopens the ratified F101 defects, has no unique owned interface inventory, and does not freeze the exit-fixture contract

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r1
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- owners remain under the addressed hold; master must re-cut and return the recipe before any release
GRILL_REQUIRED: no -- the operator-ratified Item-A and exit-gate choices are fixed; this return restores their mechanical contracts
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-171500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE item-A recipe 44bb27fa -- restore the ratified bundle schema/hash domain, replace semantic owner nomination with the versioned extractor+source inventory, give joins/carried records unique owners and interface IDs, and freeze the full exit-fixture manifest; owners remain held

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260726-171500.md` at SHA-256 `32b87ead902f779c9ce2db95aaff0fafa22ee59b1c45f302e85583d70691ebbb`.

Recipe reviewed: `master/relays/step3-relock-item-a/DESIGN-orchestrator-planner-20260726-170000.md` at SHA-256 `44bb27fa7420a00bebd4013ce2bb87fcfc66919c3865c293ebe7657d986d931a`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, especially Sections 4, 7, and 11.

## Findings

### ITEM-A-VP-R1-F1 -- BLOCKER: the recipe replaces the ratified bundle schema and hashes provenance it says is excluded

Recipe `:31-41` defines a new object `{bundle_version, owners, base_hashes, joins, carried}` and then sets `bundle_sha256 = SHA-256(JCS(STEP-3-INTERFACE-BUNDLE.json))`. Ratified Section 4 `:83-93` fixes a different, exact contract:

```text
top level = {schema, recipe_version, recipe_sha256, bundle_sha256, lock_payload, provenance}
lock_payload = {recipe_version, recipe_sha256,
                interfaces:[{interface_id, extracted_sha256} sorted by interface_id]}
bundle_sha256 = SHA-256(JCS(lock_payload))
provenance = {sources:[{interface_id, source_path, source_sha256,
                       region:"marker"|"whole_file"}]}
```

Under the proposed whole-document formula, every `base_hashes` value is inside the constitutional digest even though `:35` and `:41` call those hashes provenance-only. A Tier-SOFT edit changes the mixed source's full-file hash, changes the bundle object, and moves `bundle_sha256`. That is the exact F101 failure the re-scope replaced. The proposed object also omits the ratified `schema`, recipe identity, top-level `bundle_sha256`, canonical `lock_payload`, and separated `provenance` fields.

Required correction: use the ratified top-level schema and hash **only** JCS(`lock_payload`). Keep mixed-source full-file SHA values exclusively in `provenance`; only a `whole_file` source's identical `extracted_sha256` legitimately enters `lock_payload`. Joins and carried obligations may affect the digest only by becoming declared HARD interfaces with unique `interface_id` + `extracted_sha256`, never as an unversioned second top-level hash domain.

### ITEM-A-VP-R1-F2 -- BLOCKER: Tier-HARD/Tier-SOFT is a semantic vote, not a decidable extraction boundary

Recipe `:28-29` defines HARD as "normative, semantically load-bearing" and SOFT as wording that "does not change a normative element"; `:48-57` then asks each owner to nominate a canonical object and mark what it deems SOFT. An extractor cannot decide either category from those words. No source path, interface ID, region kind, literal marker span, recipe version, recipe digest, extraction command, or failure behavior is specified. Current domain sources contain zero `HARD-BEGIN`/`HARD-END` markers, and `master/tools/extract-interface-bundle.py` does not yet exist; those absences are expected before Item A, but the release recipe must say exactly what will be authored rather than substitute manual classification.

Ratified Section 4 `:94-105` already fixes the mechanism: literal versioned HARD markers or a declared `whole_file` source; verbatim region extraction; master-owned `master/tools/extract-interface-bundle.py`; exact `--verify`; sorted interface IDs; and fail-closed rejection of absent/ill-formed spans, duplicate or undeclared IDs, source-SHA mismatch, and recipe-version mismatch.

Required correction: before owner release, choose and enumerate the actual source strategy for every interface (marker span in an exact source path or a dedicated whole-file-hard artifact), assign its stable `interface_id` and recipe version, and bind the exact generator/verify commands plus every ratified fail-closed check. Owner review decides whether the selected bytes are the right contract; the extractor mechanically decides which bytes are hashed. Also specify the shipped soft-stability fixture's path, mutation input, command, and expected pair of results: a SOFT mutation changes `provenance.source_sha256` but not `bundle_sha256`; a HARD-region mutation moves `extracted_sha256` and `bundle_sha256`.

### ITEM-A-VP-R1-F3 -- BLOCKER: the six topic lists do not establish single ownership, and joins/carried content has no source or reviewer

Recipe `:34-37` places six raw owner payloads plus `joins` and `carried` content into the proposed bundle. The owner scopes at `:48-55` are semantic topic lists, not a complete interface inventory. They do not prevent two owners from restating the same formula, a consumer from hashing a producer's bytes, or a load-bearing join/carry from being omitted. The `joins` and `carried` objects have no unique interface IDs, source paths, HARD regions, sole author, pair/VP review chain, or extracted digests. Therefore the asserted parallel independence at `:62-63` and no-foreign-byte-hashing property are not decidable.

Required correction: provide a complete manifest table before release with, for every HARD interface, `{interface_id, sole_owner, source_path, region, recipe_version, settled_basis}` and one non-overlapping extraction source. Producer formulas appear once under their owner; consumer contracts bind/reference them rather than copy them. Cross-domain join records and the carried-obligation record need explicit master-owned integration sources reviewed by the VP (or another already-settled sole owner), and enter `lock_payload` only through their own extracted digests. The nine settled hashes remain lineage/provenance evidence, not raw hash inputs. With that table, demonstrate that every load-bearing item is present exactly once and every owner's work is independent until master's assembly barrier.

### ITEM-A-VP-R1-F4 -- BLOCKER: `STEP-3-EXIT-FIXTURES.json` is not specified, and a v3-deferred caveat is not a current passing fixture leg

Recipe `:22-26` promises the six-property fixture manifest, but `:43-46` defines only three carried items. It omits the ratified Section 7 `:363-393` contract: the six gate IDs/predicates, required durability sub-fixtures, per-fixture `{fixture_id,input_artifact_sha256,fault_injection_point,expected_canonical_rows,observer_id,evidence_locator}`, crash counter expectation, two-record handoff expectation, resume/degraded expectations, sample weights totaling exactly 30 governed turns + 100 tool calls, frozen baseline digests, overhead thresholds, and evidence schema.

The three carried items also need distinct dispositions:

- N910 can be a frozen expected cut: no `m3.b_sink.v1` record plus m-10 `UNKNOWN_PROVIDER_OUTCOME -> uncertain`, mapped to a named current fixture/predicate.
- `env_digest` parity can be a frozen current fixture with canonical logical input/expected preimage bytes plus duplicate-name and non-UTF-8 reject vectors, binding m-1 Section 5 `:63`, m-9 Section 7, and the future m-3 E3 observer.
- r7-mirror is explicitly **v3-deferred** with a re-open predicate. It cannot be represented as a current Step-3 pass-required fixture unless the missing independent m-10 resolution is first designed. Record it as a non-gating carried limit/re-open condition in a closed, typed section, not as evidence that a current fixture passes.

Required correction: freeze the full ratified manifest schema and all six existing gate legs, then define a closed additive representation for these three carried records with explicit fixture mapping, gating/non-gating status, expected fields, owners, and source loci. Preserve the honest exit claim `T1-T8 live; N910 documented MVP limit; r7-mirror deferred-v3`; never turn the carried metadata into "complete lane-2 coverage."

## Passed portions

- The addressed hold is correct and effective. The owners are in `TO`, the prior nominate-now authority is explicitly revoked, and no owner nomination exists in the item-A trail.
- The nine settled anchors and three amendment anchors are the right lineage basis; the corrected `env_digest` recipe locus is m-1 Section 5 `:63`, not m-9 Section 10.
- Item A correctly **produces but does not lock** the bundle. The assembled-bundle VP review, lane-4 lock, lane-5/T4 sequencing, H-12 boundary, and all no-code/no-provider/no-merge holds are preserved.
- A six-owner parallel extraction can be valid after F2/F3 supply a mechanical, non-overlapping source inventory. Master's assembly remains the barrier.

## Gate disposition

- Keep the addressed hold on `...-170000`; do not release any owner nomination authority.
- Recipe `44bb27fa...` is not a valid extraction contract and must not seed owner payloads.
- Master re-cuts the recipe against the exact ratified Section 4/Section 7 schemas, including the complete interface source inventory and carried-record dispositions, then returns it for VP decomposition re-review.
- No bundle, fixture freeze, interface lock, DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, E3, merge, deploy, or external use is authorized.

## Verification

- Recomputed exact hashes: review-routing target `32b87ead...`; recipe `44bb27fa...`; ratified stage-6 amendment `1125b0a0...`; settlement amendment `1fa71cb8...`.
- Exact-file lint is `OK` for both incoming relays.
- Read ratified Sections 4, 7, 11 and the prior F101 closure trail; scanned current domain sources for HARD markers (zero files), checked the canonical extractor/bundle/fixture paths (not yet authored), and confirmed the item-A tree contains only `...-170000` and `...-171500` -- no nomination leak.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain; HEAD and `origin/main` both equal `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no recipe, amendment, owner design, settled base, `frank/` source, branch, commit, release, bundle, fixture, lock, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260726-180000.md`.
Next requested action: master re-cuts Item A to the ratified bundle/extractor and full exit-fixture contracts, keeps all six owners held, and returns the exact replacement recipe for VP re-review.
