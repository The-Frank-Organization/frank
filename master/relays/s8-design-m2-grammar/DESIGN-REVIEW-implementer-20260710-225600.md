## DESIGN-REVIEW - s8 config/atom grammar rev3 must revise activation-stage partition

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-m2-grammar-design-review-r4
PARENT_DISPATCH_ID: s8-m2-grammar-design-r4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - operator ratification of activation authorization remains lock-blocking after technical revision
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s8-m2-grammar-grill
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m2-grammar/DESIGN-planner-20260710-225500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-3.planner, m-5.planner, m-6.planner
SUBJECT: rev3 must revise - evidence census is corrected, but seat-proxy requirements and sibling-computed outputs remain assigned to the wrong activation stages

DESIGN_REVIEW_VERDICT: must-revise

### Findings

1. **BLOCKER - pre-observe form validation still excludes real observe-gated seat requirements.** Rev3 correctly acknowledges the ten routing/ODB rows using `layer_present:observe` as the s5-M-1 dormancy proxy (`design:57`), but the executable stage rule still says form validation checks **only** the two class-1 evidence refs (`design:60`) and the seam summary still says the CQ-1 guardrail touches only observe-owned rows (`design:64`). At `frank@2e1b4f0`, `justified_deviation` and `deviation_reason_code` are seat-fillable and each has `required_when = layer_present:observe AND any routing_assignments.declared_deviated == yes`. Once activation is on, those are real conditional form requirements. The other eight proxy rows become visible but are not unconditionally required. A PLAN following the current stage rule can omit both required deviation fields or incorrectly treat the whole proxy cluster as merely optional prose.

   **Required revision:** partition the pre-observe seat surface explicitly:
   - evidence refs: `ACTIONS_GIT_REF`, `FINAL_GIT_STATUS_SHORT`, required and rendered when observe is active;
   - proxy-gated consumer rows: all ten become render-eligible when observe is active, while only `justified_deviation` and `deviation_reason_code` become conditionally required when any declared-deviation row says yes;
   - form validation evaluates every seat-fillable activated predicate after the stale-digest short-circuit, not “only class 1.”

   Extend the two-state fixture with declared-deviation false/true cases proving the two conditional requirements, while the remaining proxy rows are visible without becoming blanket-required. Replace the false “CQ-1 touches only observe-owned rows” echo with the exact distinction: CQ-1 proper is observe-owned; s5-M-1 deliberately reuses the atom as a non-observe-owned dormancy proxy.

2. **BLOCKER - the computed class is not one `observe_gate` producer set, and global completeness contradicts the locked `surface_intent` profile.** The matrix groups `target_gap_result`, `record_integrity`, `deviated_observed`, `bucket_binding_observed`, `surface_intent`, and `authority_class` as if `observe_gate` computes all six (`design:55,62-65`). The locked homes disagree:
   - m-3 produces the evidence outputs and routing-profile observed set;
   - m-2 computes `authority_class`, which m-3 consumes as the CQ-2 disposition key (`v3-form-schema-design.md:345`; m-3 `:63`), so it cannot first appear as an undifferentiated post-observe output;
   - `surface_intent` is a conductor-derived m-5 vocabulary / m-6 binding and is present on **non-gate-bearing records only**; gate-bearing records carry none (`v3-form-schema-design.md:346`; m-6 `:121-138`). Yet the live registry currently gates `surface_intent` only on `layer_present:observe`, and rev3's global class-3 completeness rule would require it on every activated record.

   **Required revision:** split computed rows by actual producer, consumer, record profile, and pipeline stage. Name where `authority_class` is computed before m-3 uses it. Reconcile `surface_intent` activation with its locked non-gate-only contract: either correct the registry predicate through the required review-driven amendment or define an equivalent profile-aware completeness rule accepted by m-5/m-6/m-7; do not silently require it on gate-bearing records. Post-observe completeness must be profile-aware and must validate each owner's allowed/required outputs, not treat every computed CQ-1 row as an m-3 write.

3. **BLOCKER - the proposed “registry-derived exact-set” fixture lacks an independent actual side.** Deriving an expected set from the registry and then asserting that derived set against itself cannot prove that m-3/m-2/m-5/m-6/m-7 actually produce the required outputs. It also cannot make a stray prose inclusion fail unless the implementation has a separate output manifest to compare. The fixture becomes meaningful when the registry-derived expected set is compared against independent producer allowlists/profile manifests and exercised through accepted records. For every applicable record profile, assert: activated registry requirements = union of the authorized producer outputs, with no missing, extra, or wrong-owner write; separately assert the committed-record completeness result. This also forces the producer/stage split in finding 2 to be executable.

4. **HUMAN DECISION remains open.** The operator-authored, digest-pinned `config_change`, restart-effective, no-lane-surface authorization position remains coherent. Explicit operator ratification or override is still required before approval/lock; this review does not proxy that decision.

### Accepted Rev3 Fold

- The evidence-row census is now correct: `completed_proof` is outside CQ-1; `bucket_binding_observed` belongs in the computed evidence/routing set. No reversion of that correction is requested.
- The distinction between genuine observe-owned rows and the s5-M-1 dormancy proxy is directionally correct; findings 1-2 require carrying that distinction into the executable stage and producer contracts.
- All previously accepted rev2 folds remain accepted: locked pipeline order, stage-split completeness concept, reader-capability gate, knob fail-closed/catalog fail-open consistency, immutable layer context, stale-digest short-circuit, version carrier, and forward-only history.

### Re-review Bar

Return one revision that (a) includes the proxy-gated seat predicates in form validation, (b) partitions computed outputs by producer/profile/stage including the non-gate-only `surface_intent` rule and pre-disposition `authority_class`, and (c) gives the exact-set fixture an independent producer-manifest and behavioral-record oracle. Carry operator ratification or remain explicitly `human-decision-required` until it arrives. No code, PLAN, IMPL, silent c1/CQ-5 amendment, or unreviewed sibling-domain semantic change is authorized.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of rev3 `s8-design-m2-grammar` against `frank@2e1b4f0`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean and synchronized at `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-design-m2-grammar` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner corrects the activation-stage/producer partition with m-3/m-5/m-6/m-7 confirms, preserves the accepted census and rev2 folds, and returns a revised DESIGN after or alongside operator ratification.
