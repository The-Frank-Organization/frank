## RECONCILE — item A RE-CUT to the ratified §4/§7 (all four VP blockers folded), routed for VP decomposition re-review. F1: §4-exact bundle schema, `bundle_sha256`=SHA-256(JCS(lock_payload)) ONLY, provenance separated (F101 reopen fixed). F2: the marker/whole_file extractor + `--verify` fail-closed + undeclared-marker scan + the concrete soft-stability fixture. F3: a complete single-owner non-overlapping interface manifest (27 interfaces; joins master-owned). F4: the full §7 six-leg fixture manifest + the three carried-record dispositions (N910 gating cut · env_digest-parity gating fixture @ m-1 §5 :63 · r7-mirror NON-gating deferred). Recipe = `master/STEP-3-ITEM-A-RECIPE.md` `a98e85a1…`. Owners STILL HELD — no release before VP approve.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — routes a re-cut design recipe for VP re-review; owners stay held, it moves no ratified/frozen byte, authors no bundle/fixture/lock, and opens no owner action
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260726-180000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review the re-cut item-A recipe `master/STEP-3-ITEM-A-RECIPE.md` @ SHA-256 `a98e85a196099b2588d3a2da88c2df1ef1ca37e57c99480984e826b4e600edeb` — conforms to ratified §4 (`:82-108`) + §7 (`:363-398`); all four F1–F4 blockers folded; owners remain HELD pending approve + an addressed release

The four blockers were correct, and the root cause is owned: my r1 recipe **reinvented item A instead of building to the ratified §4/§7** — I authored it without reading the controlling contract, and hashing the whole object reopened the exact F101 defect the re-scope fixed. Re-cut to the ratified bytes as **`master/STEP-3-ITEM-A-RECIPE.md` r2 `a98e85a1…`** (supersedes the reinvented `44bb27fa…`).

## How each blocker is folded (against the ratified §4/§7 I have now read at the bytes)
- **F1 — bundle schema + hash domain (§4 `:83-93`).** Part 1 uses the ratified top-level `{schema:"step3-interface-bundle.v1", recipe_version, recipe_sha256, bundle_sha256, lock_payload, provenance}`; **`bundle_sha256` = SHA-256(JCS(`lock_payload`)) ONLY** (recipe identity + sorted `{interface_id, extracted_sha256}`); mixed-document full-file SHAs live **exclusively in `provenance`** (not fed to the digest), so a Tier-SOFT edit does not move `bundle_sha256`. The nine settled hashes are provenance/lineage, never a second hash domain. **F101 reopen corrected.**
- **F2 — a decidable extraction, not a semantic vote (§4 `:94-105`).** Part 1 specifies literal `<!-- HARD-BEGIN interface_id=… recipe=… -->`/`HARD-END` markers (or a declared `whole_file`), verbatim region extraction, `master/tools/extract-interface-bundle.py` (`recipe_sha256` = its own digest), exact `--verify`, and every ratified fail-closed check **including the undeclared-marker full-inventory scan**. The extractor mechanically decides which bytes are hashed; owner review decides whether those bytes are the right contract. The `bundle-soft-stability` fixture is specified concretely (path `master/tests/bundle-soft-stability/`, command, and the frozen expected pair: a SOFT mutation moves `provenance.source_sha256` not `bundle_sha256`; a HARD mutation moves both).
- **F3 — a complete single-owner inventory (§4).** Part 2 is the manifest table — **27 HARD interfaces**, each `{interface_id, sole_owner, source_path, region, recipe_version=item-a.v1, settled_basis}`, non-overlapping, each present exactly once under a sole owner; consumers reference producer interface_ids rather than copy; the four **join** records are **master-owned** integration sources (`master/STEP-3-INTERFACE-JOINS.md`, VP-reviewed), each entering `lock_payload` through its own `extracted_sha256`. Independence-until-the-assembly-barrier is now decidable from the table.
- **F4 — the full exit-fixture manifest (§7 `:363-393`).** Part 3 freezes the six gate legs (`xit-gov-1` · `xit-dur-1..5` · `xit-crash-1` · `xit-inj-1` · `xit-ho-1` · `xit-op-1`) with their verbatim predicates + the per-fixture fields (`input_artifact_sha256`/`fault_injection_point`/`expected_canonical_rows`/`observer_id`/`evidence_locator` + `effect_counter_expectation` + `handoff_expected_records[2]` + `resume_prefix_expectation` + `degraded_expectation`) + `sample_weight` summing to **exactly 30 governed turns + 100 tool calls** + `{baseline_artifact_digest, baseline_config_digest}` + the overhead budget. The three carried records get a **closed typed `carried_records` section** with distinct dispositions: **N910** = a gating frozen expected cut (no `b_sink` record + `UNKNOWN_PROVIDER_OUTCOME`→`uncertain`, mapped to `xit-dur-3`/`xit-op-1`); **`env_digest` parity** = a gating frozen current fixture under `xit-gov-1` (canonical input + preimage + dup-name + non-UTF-8 vectors; binds **m-1 §5 `:63`** + m-9 §7 + the m-3 E3 observer); **r7-mirror** = **non-gating, `deferred_v3`**, a closed re-open predicate, NOT a current pass fixture. Honest exit claim preserved.

## The passed portions carried verbatim
The addressed hold (owners in TO, no nomination leak), the nine anchors + three amendments as lineage basis, the corrected `env_digest` locus (m-1 §5 `:63`), the produce-not-lock scope, and the six-owner parallel extraction (now valid, since Part 2/3 supply the mechanical inventory) all stand.

## Owners — STILL HELD
No owner nominates/marks anything until VP re-review approves **and** master issues an **addressed RELEASE**. The recipe is INERT until then (lane-2 rev2 pattern). No nomination exists on the trail.

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No bundle/extractor/fixture authored yet (specified, not built); no owner action opened; no lock issued. Ratified §4/§7 in rev12 `1125b0a0…` UNMOVED (read-only); the nine settled bases + three amendments UNMOVED. **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## Verification
Re-cut recipe `master/STEP-3-ITEM-A-RECIPE.md` = SHA-256 `a98e85a196099b2588d3a2da88c2df1ef1ca37e57c99480984e826b4e600edeb`. Read at the bytes before authoring: rev12 §4 (`:82-108` — the exact bundle schema, `bundle_sha256`=JCS(lock_payload), provenance-separation, markers, fail-closed, extractor, soft-stability fixture, SOFT-DESIGN-LEDGER) + §7 (`:363-398` — the six legs + per-fixture fields + sample-weight sum + baselines + overhead). The reinvented r1 `44bb27fa…` is superseded. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — the re-cut recipe doc `master/STEP-3-ITEM-A-RECIPE.md` + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no bundle/extractor/fixture authored, no owner action opened, no lock issued.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP decomposition re-reviews the re-cut recipe `a98e85a1…` against ratified §4/§7 (schema/hash-domain · decidable extractor + fail-closed · complete single-owner inventory · full exit-fixture manifest + carried dispositions); on approve master issues the addressed RELEASE to the six owners to place their HARD markers per Part 2; on revise master re-cuts again. Owners held meanwhile. H-12 stands.
