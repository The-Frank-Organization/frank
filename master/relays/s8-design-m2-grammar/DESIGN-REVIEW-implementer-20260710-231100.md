## DESIGN-REVIEW - s8 config/atom grammar rev4 must revise stale global producer claims and resolve surface profile

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-m2-grammar-design-review-r5
PARENT_DISPATCH_ID: s8-m2-grammar-design-r5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - operator ratification of activation authorization remains lock-blocking after technical and sibling resolution
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s8-m2-grammar-grill
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m2-grammar/DESIGN-planner-20260710-231000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-3.planner, m-5.planner, m-6.planner
SUBJECT: rev4 must revise - new producer partition is directionally correct, but stale global observe claims remain and surface_intent has no executable profile predicate yet

DESIGN_REVIEW_VERDICT: must-revise

### Findings

1. **BLOCKER - the old global `observe_gate` contract still survives in lock-bearing sections and contradicts rev4's producer partition.** Rev4 correctly states at `design:64-70` that the computed class is not one m-3 producer set. But builders encounter the opposite rules first and again at the boundary/lock:
   - the matrix says all six computed rows, including `authority_class` and `surface_intent`, are computed “from the observed set” (`design:55`);
   - pipeline steps 3-4 say `observe_gate` computes class 3 and post-observe completeness checks outputs the observe hook produced (`design:62-63`);
   - the §4 seam still says class 2/3 are filled by `observe_gate`, with only m-3/m-7 owners (`design:86`);
   - GRILL resolved decision still says conductor-observed/computed are filled by `observe_gate` (`design:112`).

   These are not historical fold-log quotations; they are current design rules. They would authorize exactly the wrong-owner implementation rev4 says the independent-manifest fixture must reject.

   **Required revision:** replace the matrix's single computed row with producer/profile rows, and propagate the per-producer contract through the numbered pipeline, §4 seam, GRILL_LOCK, design-lock impact, and fixture summaries. Use one consistent sequence: all activated seat predicates at form validation; m-2 `authority_class` at an m-7-confirmed pre-disposition point; m-3 observe outputs at step 4; sibling-confirmed `surface_intent` only on its applicable profile; then per-profile completeness before append.

2. **BLOCKER - `surface_intent` is correctly surfaced as a contradiction but not yet resolved to an executable registry contract.** “Add the not-gate-bearing gate” (`design:68`) is not a predicate specification. The current grammar has no `gate_bearing` atom; `authority_class` is deliberately non-gate-referenceable; and the authority/gate profile is derived from the locked authority record-kind set **or** config-sourced gate-category class A. Duplicating that formula into a static FieldSpec predicate risks drift from the authoritative classifier and its config. The offered alternative, profile-aware completeness alone, still conflicts with the registry's unconditional `required_when: layer_present:observe`, and therefore cannot satisfy rev4's fixture equality between registry-derived requirements and producer outputs for gate-bearing records.

   **Required revision:** after m-5/m-6/m-7 confirmation, choose and specify one executable source of truth:
   - a concrete, grammar-valid registry predicate whose inputs are authoritative and non-circular; or
   - an explicit reviewed amendment that removes/changes the global registry requiredness and assigns profile applicability to a named producer/profile manifest consumed by completeness.

   Name the exact predicate/manifest shape, its owner, the stage at which the profile is known, and fixtures for both gate-bearing and non-gate-bearing records. Do not predicate on `authority_class` unless the locked non-gate-referenceable rule is itself amended through review, and do not hardcode the config-sourced A map into a second drifting list.

3. **DEPENDENCY - sibling contracts are routed but unconfirmed.** Rev4 appropriately stops claiming unilateral authority. m-3 must confirm the observe-output manifest; m-5/m-6 must confirm `surface_intent` semantics and the selected amendment; m-7 must confirm the `authority_class` compute point and profile-aware completeness stage. Until those returns are reconciled into this artifact, the independent-manifest fixture has no locked actual side and this design cannot approve.

4. **HUMAN DECISION remains open.** The operator-authored, digest-pinned `config_change`, restart-effective, no-lane-surface authorization position remains coherent. Explicit operator ratification or override is still required before approval/lock; this review does not proxy that decision.

### Accepted Rev4 Folds

- Form validation now correctly covers all seat-fillable activated predicates, including conditional `justified_deviation` and `deviation_reason_code` requirements.
- The CQ-1-proper versus s5-M-1 dormancy-proxy distinction is correct.
- The producer/profile partition at `design:64-70` is the correct direction: m-3 does not own `authority_class` or `surface_intent`, and completeness needs an independent producer-manifest plus behavioral-record oracle.
- Surfacing the live `surface_intent` registry conflict as a review-driven amendment, rather than silently editing it, is correct governance.
- All prior accepted census, pipeline-order, capability-gate, layer-context, stale-digest, version, and forward-history folds remain accepted.

### Re-review Bar

Return a byte-consistent design after the sibling confirms select the exact `surface_intent` profile mechanism and m-7 places `authority_class`/completeness in the pipeline. Carry operator ratification or remain explicitly `human-decision-required`. No code, PLAN, IMPL, silent registry change, c1/CQ-5 amendment, or proxy sibling decision is authorized.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of rev4 `s8-design-m2-grammar` against `frank@2e1b4f0`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean and synchronized at `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-design-m2-grammar` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner gathers m-3/m-5/m-6/m-7 confirms, selects an executable `surface_intent` profile contract, sweeps every lock-bearing echo to the per-producer sequence, and returns a revised DESIGN after or alongside operator ratification.
