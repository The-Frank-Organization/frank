## RECONCILE -- REVISE: rev2 preserves the right product choices, but the extractor and journal source map are still promises, two DAG legs remain wrong, and the exit gate is not predeclared

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r2
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the bounded operator grill is complete, but operator re-scope ratification remains held until the amendment contains the mechanical contracts and pre-T4 gate definition it currently promises for later
GRILL_REQUIRED: no -- GRILL_LOCK_ID `step3-stage6-rescope-grill-1` durably settles sandbox exclusion, ambient bash with the narrowed claim, non-gated utility demonstrations, and the property-plus-overhead exit direction
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-041500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- accept the grill, F102 closure, narrowed bash claim, broker-first placement, and non-gated utility proof; finish the hash extractor, hard invocation-context contract, B/E dependency graph, journal truth map, and executable pre-T4 exit/H-12 gates

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-041500.md` at SHA-256 `f9aabf6feef32159cce0c8af7c10106f0adbae243bf22292ccc66e1ad71927d1`.

Proposed amendment rev2: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `29a36285b9c2c5a49deff828490e00c4b99b8a249ed2f6941f2f8821683d1bc6`.

## Findings

### F101 -- BLOCKER, STILL OPEN: the amendment names a future extractor but does not define a mechanically reproducible lock boundary

Amendment `:119-139` says a versioned extraction recipe will pull hard material into a canonical bundle, and `:223-229` schedules authoring that recipe only after the pair deltas. There is still no artifact path or schema, source-marker or dedicated-interface-file mechanism, canonical ordering/encoding, source-SHA binding, duplicate/missing-source failure rule, generator owner, output identity, or exact verification command. The current mixed realization docs contain no machine-identifiable Tier-HARD regions. A future extractor is a sound direction; it is not yet the hashable mechanism that rev2 and target section 3 claim exists.

Required correction: define the extraction contract in the amendment now while allowing the final bundle values to be populated after B-E settle. At minimum bind `{source_path, source_sha256, interface_id, extraction_recipe_version, extracted_sha256}`, deterministic order and encoding, stable explicit source markers or dedicated interface artifacts, fail-closed behavior for missing/duplicate IDs and source-hash mismatch, generator/verifier owner, canonical output path, recipe digest, and exact generation/verification commands. The final bundle may still be authored last and reviewed under F73.

Amendment `:133` also routes every Tier-SOFT edit into `master/PROTOCOL-DEVIATIONS.md`. That register is constitutionally scoped to deviations from the stock agentic-dev-team protocol (`CLAUDE.md:117-125`; `PROTOCOL-DEVIATIONS.md:1-9`), not ordinary product-design evolution. Name a dedicated soft-design ledger or keep those changes in the owning design histories; do not overload the framework-deviation register.

### F103 -- BLOCKER, PARTIALLY CLOSED: the broad bash effect claim is correctly removed, but the exact invocation-context fields are simultaneously HARD and SOFT

The grill and amendment `:71-96,147-155` correctly remove containment and per-effect destructive/protected-hold claims. That closes the central F103 objection without reopening a sandbox decision.

The remaining contract contradicts itself. Amendment `:125-136` makes effect identity, canonicalization, process teardown, and UNKNOWN visibility Tier-HARD, while `:191-195` makes path resolution, `..`, symlink behavior, shell identity, fixed cwd, environment allow-list, resource limits, and background-process behavior Tier-SOFT. Those are the semantics of the claimed exact invocation context and honest process outcome, not merely UX. The descriptor also requires `canonical_resource` for an action set that includes arbitrary `bash`, without defining field applicability, and `backend_id="ambient"` does not identify the shell/tool implementation that interpreted the invocation.

Required correction: split section 7. Keep presentation and ergonomics soft, but place the context-identity and teardown semantics needed by the descriptor in the hard bundle. Define per-action required/absent/null rules for every descriptor field; define what `canonical_resource` means for each local action and whether it is inapplicable for `bash`; define cwd normalization and `env_profile_digest` snapshot/application semantics; and bind shell/tool implementation identity through the descriptor or an explicitly joined release/build identity. This does not require containment, affected-resource inference, or a per-effect hold.

### F104 -- BLOCKER, STILL OPEN: the replacement graph is not a correct producer/consumer DAG for B or E

Amendment `:179-189` removes the old pair-total cycle, and C, D, the m-7-first placement, and conditional H-24 are sound. Two legs remain incorrect:

- B orders `m-8 -> m-9 carrier -> m-10 row -> m-3 E0/E3`. The frozen m-3 contract owns the `m3.app_event.v1` E0 schema and the E3 schema/evaluator; m-9 cannot implement the new E0 carriage before that schema delta exists. The m-9 E0 carrier and m-10 attempt row are sibling consumers of m-8 output, not a serial `m-9 -> m-10` dependency. The graph conflates runtime dataflow with design/schema dependency.
- E includes provider-specific tool lowering in `model_surface_digest` (`:169-177`) but names only m-9 and m-2 as producers. m-8 owns provider wire translation and therefore owns or must confirm the provider-lowering contribution. The unnamed `canonical aggregator` has no owner, input schema, output artifact, or digest recipe.

Required correction: express B as explicit schema/producer/consumer joins, with m-3's E0/E3 schema available before m-9 emission and m-8 feeding the independent m-9-carriage and m-10-row consumers before m-3's final evaluator join. Add m-8 to E, name the aggregator owner, and define its canonical component-input and output/digest contract. Distinguish design dependency edges from runtime event flow.

### F105 -- BLOCKER, STILL OPEN: section 5-D requires a source map but does not contain one

Amendment `:157-167` says the amendment "MUST ship a field-level source map" and lists topics that map must define. No field-level map follows. The text does not name the journal records or fields, identify which existing event/blob is canonical for each, declare which values become first-class durable records, or give writer/reader and commit ordering at field/record grain. Calling m-9 the source/content owner and m-10 the persistence owner leaves the same ambiguity F105 rejected.

Required correction: include the actual architectural source map in current amendment bytes. Cover at least attempt input references, provider-visible output items, parsed tool calls, settled tool results, compaction events/checkpoints, objective/constraint references, workspace snapshot identity, and UNKNOWN markers. For each, name the canonical source or first-durable-copy record, writer, readers, content/blob ID and digest, transaction/linearization relative to canonical attempt/tool rows, replay/idempotency and crash behavior, retention/GC, size bound, access/redaction rule, and resume relevance. Physical implementation detail may remain with the pair designs; source-of-truth and commit ownership may not.

### F106 -- BLOCKER, PARTIALLY CLOSED: the grill settles the product direction, but its gated acceptance contract and H-12 promotion are not yet executable

The embedded grill is valid and closes the former backend/bash choices. The operator's decision that CRM/bivpak dogfood and SWE-bench are demonstrations rather than exit thresholds is explicit and accepted; their deferred corpus details are therefore non-blocking.

The gated half is still post-selectable. D3 (`:82-85`) requires six governance-property legs plus a predeclared objective overhead budget, but section 7 (`:196-206`) gives labels rather than executable predicates and replaces a budget with a "loose hard ceiling against catastrophe." It provides no numeric ceiling, baseline/normalization, fixture or task IDs, sample count, evidence schema, pass/unknown rule, or freeze point before T4. Section 3's design-lock summary also says "four property legs" at `:116`, contradicting the six named at `:196-202`.

H-12 is called a HARD pre-external-use blocker at `:72-77`, but the source backlog still records it as Step-4 hardening for an accepted MVP residual (`FRANK-HARDENING-BACKLOG.md:27`), and rev2 names no gate artifact, enforcing authority, or classification test for `untrusted`, `external`, `security-sensitive`, or `multi-tenant` use. Public dogfood must not accidentally bypass that boundary merely because it is non-gating for Step-3 close.

Required correction: before operator re-scope ratification, define the six legs as executable acceptance predicates with fixed fixture/evidence identities and pass/fail/unknown behavior; specify the numeric overhead ceilings, measurement baseline and method, repetitions/sample rule, and a pre-T4 freeze gate. Reconcile four versus six. Record H-12's promotion in its source ledger or state explicit amendment precedence, define the allowed trusted-operator deployment envelope and prohibited-use classifier, and name the release/deploy/public-use gate owner and artifact. No additional product grill is required unless those corrections introduce a new choice outside the existing lock.

## Accepted closures and direction preserved

- **F102 is CLOSED.** Dropping the fixed policy removes the attempted ownership transfer. m-10 remains the F59 authorization host, m-9 remains executor/consumer, and keeping m-5 stood down plus m-6 out is removal-not-reassignment. Their CC routing is sufficient context.
- **F103's claim narrowing is accepted.** Ambient bash is explicit; there is no sandbox, containment, exact-affected-resource, or destructive/protected per-effect claim in the MVP.
- **The operator grill is durable.** Sandbox exclusion, ambient bash, property-plus-overhead exit direction, non-gated public dogfood/SWE-bench, exit/dogfood decoupling, and descriptor retention are not reopened by this review.
- **Broker sequencing is corrected.** The m-7 study resolves before affected m-9/m-10 finals and re-lock; H-24 runs before re-lock if cross-epoch completion survives. The separate secret-holding process and epoch fence remain.
- **C and D's two-sided ownership shape is directionally correct.** C keeps authorization in m-10 and execution in m-9. D correctly identifies m-9/m-10 as a coordinated seam, subject to F105's missing source/commit map.
- The held joint-lock proposal remains correctly superseded. No previously approved design byte moved: all nine design finals, H-16 rev16, and H-17 census v3 recompute to the r4-approved hashes.

## Gate disposition

- Proposed stage-6 amendment rev2 `29a36285...`: REVISE; not ready for operator re-scope ratification.
- F102: closed. F103 and F106: direction closed, mechanical contracts still open as bounded above. F101, F104, and F105: still open.
- Pending all-artifact joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- Stage-6 re-lock remains unrequestable until the corrected interface deltas, exact bundle recipe/artifact, reviews, and consumer joins complete.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, public/external use outside the explicitly allowed envelope, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r3 over new amendment bytes containing: (1) a deterministic extractor/manifest contract; (2) a non-contradictory hard invocation-context and teardown contract with per-action field applicability; (3) corrected B and E DAGs with m-8 and a named aggregator; (4) the actual journal field/source/commit map; and (5) frozen executable property/overhead and H-12 use-boundary gates. Preserve the grill, F102 closure, narrowed bash claim, m-5/m-6 disposition, broker-first/H-24 placement, non-gated utility choice, and all unchanged artifact hashes.

## Verification

- Target SHA-256: `f9aabf6feef32159cce0c8af7c10106f0adbae243bf22292ccc66e1ad71927d1`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1874.
- Amendment rev2 SHA-256: `29a36285b9c2c5a49deff828490e00c4b99b8a249ed2f6941f2f8821683d1bc6`.
- External-review SHA-256: `b4e79f3bfdc4ca0b2f5023b441da1e0006c7064a5d4b6b06bb4e9e2a7b5bf7d1`; r1 reviewer parent SHA-256: `b72d3eaf3bd6bc70e01e26665822c78dd969646e48ffb69fb61a874c32b3e4f6`.
- All nine prior design finals, H-16 rev16 `a349a329...`, and H-17 census v3 `959b1928...` recompute exactly to the r4 manifest.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken. The target's older `6e4d657 local / c78da38 origin` status line is superseded by this fresh live-state check and is not a design finding.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master folds only the five bounded mechanical corrections above and returns amendment rev3 for decomposition review r3; operator re-scope ratification remains held.
