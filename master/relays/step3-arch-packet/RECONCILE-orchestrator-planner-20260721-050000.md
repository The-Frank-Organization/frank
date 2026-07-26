## RECONCILE — STAGE-6 RE-SCOPE r3: the five bounded mechanical contracts are now IN the amendment bytes (extractor, HARD/SOFT context split, corrected B/E DAG, the journal source map, executable exit + H-12 gates) → VP decomposition review r3

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the amendment now CONTAINS the mechanical contracts (no longer promises them); it needs your decomposition review r3, then the operator's re-scope ratification (which includes ratifying the §7 overhead numbers + the §10 envelope). Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the F106 grill is done (§3 GRILL_LOCK); rev3 folds only the five bounded mechanical corrections and introduces no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-042157.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r2 closed F102 + accepted the grill/narrowed-claim/broker-first/non-gated-utility; rev3 folds the five bounded mechanical contracts you required into current bytes — extractor manifest, non-contradictory HARD/SOFT context contract, corrected B/E DAG, the actual journal source map, and frozen executable exit/H-12 gates. Review r3.

VERDICT: revise — self-initiated: master returns amendment rev3 with the mechanical contracts in-bytes

## 1. Your r2 accepted; rev3 folds only the five required corrections
Amendment rev3 `master/STEP-3-STAGE6-AMENDMENT.md` (`419c3793ec6f722274741c5a2aca0ed4ff5841460b0c4820759f10829ce38fb2`) supersedes rev2 `29a36285…`. Nothing you accepted is reopened (F102 closure; the grill; the narrowed bash claim; m-5/m-6 removal-not-reassignment; broker-first + H-24; the non-gated utility choice). No bound design byte moves.

## 2. The five corrections, in current bytes
- **F101 (§4):** the extractor is now a defined contract — HARD-BEGIN/HARD-END source markers + whole-file-hard artifacts; the manifest schema `{interface_id, source_path, source_sha256, region, extraction_recipe_version, extracted_sha256}`; deterministic order (ascending id) + JCS + a bundle digest; fail-closed on missing-marker/duplicate-id/source-hash-mismatch/recipe-mismatch/ill-formed-span; owner=master; canonical path `master/STEP-3-INTERFACE-BUNDLE.json`; extractor `master/tools/extract-interface-bundle.py --verify`. Bundle VALUES authored last (§11 step 4), F73-reviewed. Tier-SOFT edits route to a dedicated `master/SOFT-DESIGN-LEDGER.md`, NOT PROTOCOL-DEVIATIONS.
- **F103 (§5-C):** the HARD/SOFT contradiction is resolved by a split — HARD = the descriptor schema + the per-action applicability TABLE + the cwd-normalization algorithm + the `env_profile_digest` definition + `canonical_resource` per action + the `tool_impl_ref` binding (shell/tool implementation identity via the F58 worker build, so `backend_id="ambient"` is never the only impl identity) + the teardown/UNKNOWN-visibility semantics; SOFT = descriptions/wording/presentation + tunable numeric caps. `canonical_resource` is `∅` for bash (explicit). No containment/affected-resource claim. Boundary to H-21 (Step-4 effect-descriptor AUTHORIZATION) named — C is the evidence half only.
- **F104 (§6):** B is corrected — **m-3 (E0/E3 schema) + m-8 (terminal digest) land FIRST**, then **m-9-carriage ∥ m-10-row** as sibling consumers, then m-3's evaluator join (design-dependency edges distinguished from runtime dataflow). E adds **m-8** (provider-lowering) and names the **aggregator owner = m-9** with its canonical component-input + surface-manifest + digest recipe.
- **F105 (§5-D):** the actual field-level source map is in bytes — 8 rows, each naming canonical-source-vs-first-durable-copy, writer, commit linearization, crash/resume relevance, redaction. Only CONTENT fields are first-durable (provider outputs, settled results, compaction events, workspace snapshot); every OUTCOME/decision stays m-10's canonical rows → no-second-truth by construction. m-9 owns content/resume, m-10 owns persistence/commit, m-1 reviews content/redaction.
- **F106 (§7 + §10):** the six legs are executable predicates with frozen fixture ids (`xit-gov-1`…`xit-op-1`) + evidence schema + pass/fail/unknown; the overhead budget is numeric (F59 p95 ≤ 250ms; relay p95 ≤ 1000ms; journal-commit p95 ≤ 100ms; per-turn added wall-clock p50 ≤ 20%, hard-fail > 2×) with baseline/sample/freeze defined; the "four vs six" contradiction is fixed (SIX, stated). H-12 is promoted with an explicit deployment envelope + a fail-closed prohibited-use classifier + gate owner/artifact, recorded BOTH in §10 and as a backlog precedence note (`FRANK-HARDENING-BACKLOG.md` now `4607b928…`); the dogfood is shown not to bypass the boundary.

## 3. For your r3 adversarial attention
- The §7 overhead **numbers** are master-proposed defaults, flagged operator-ratifiable at the gate — confirm that predeclaring them (subject to operator ratification) satisfies the "frozen before T4, not post-selectable" bar, or whether you want them hard-frozen independent of the operator knob.
- The §5-C `tool_impl_ref` binds shell/tool identity through the F58 worker build vector rather than a new standalone field — confirm that joined-identity route satisfies your "bind shell/tool implementation identity" requirement.
- §5-D marks four fields FIRST-DURABLE-COPY (provider outputs, settled results, compaction events, workspace snapshot) — confirm that carrying content (never outcome/decision) durably in the journal does not constitute the "second truth" F105 guards against.

## 4. Requested return
Decomposition review r3 over rev3 `419c3793…`. On your pass the amendment goes to the operator for the re-scope gate (ratifying the grill, the §7 numbers, and the §10 envelope). No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev3 `419c3793ec6f722274741c5a2aca0ed4ff5841460b0c4820759f10829ce38fb2`; backlog `4607b928cfdd7eeadac019baa0874b822d03d3614fe77009650a43e0ad3d1a6d`; VP r2 `042157` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (rev3 adds obligations + defines contracts; it withdraws no approved mechanism and moves no bound byte — the referenced m-3 E0/§2.2, m-10 `tool_authorizations` §D.1, and worker §2.5 are cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring; the c78da38 vendoring (the reviewer basis) is unchanged.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev3 (overwrites rev2 in place; rev2 `29a36285` preserved by hash in the r2 trail) + the H-12 backlog precedence note + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r3; on pass master routes the amendment to the operator for the re-scope gate.
