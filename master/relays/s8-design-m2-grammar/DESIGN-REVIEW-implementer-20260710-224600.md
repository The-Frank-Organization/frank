## DESIGN-REVIEW - s8 config/atom grammar rev2 must revise CQ-1 matrix census

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-m2-grammar-design-review-r3
PARENT_DISPATCH_ID: s8-m2-grammar-design-r3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - operator ratification of activation authorization remains lock-blocking after technical revision
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s8-m2-grammar-grill
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m2-grammar/DESIGN-planner-20260710-224500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-3.planner
SUBJECT: rev2 must revise - locked order and capability gate are corrected, but the CQ-1 ownership matrix has one inclusion and one omission

DESIGN_REVIEW_VERDICT: must-revise

### Finding

1. **BLOCKER - the lock-bearing CQ-1 matrix does not match the cited live registry.** The matrix says its Block-A rows were verified at `frank@2e1b4f0`, but class 2 includes `completed_proof` and class 3 omits `bucket_binding_observed` (`design:51-55`). At that exact baseline, `completed_proof` is `system` / `observed_value` with `required_when: null` and `visible_when: null`; it is not activated by CQ-1. Conversely, `bucket_binding_observed` is `system` / `computed_result` with both `required_when` and `visible_when` gated by `layer_present:observe`; it is a CQ-1 class-3 row. As written, the post-observe completeness contract and two-state fixtures would demand an unactivated field while failing to demand a real activated computed output.

   **Required revision:**
   - remove `completed_proof` from the CQ-1 class-2 row and state that it is outside this Block-A activation matrix;
   - add `bucket_binding_observed` to the CQ-1 class-3 row;
   - update every post-observe completeness, fixture, GRILL_LOCK, and summary reference to use the corrected sets;
   - make the two-state fixture compare an explicit expected class-2/class-3 ID set against the registry's `layer_present:observe` metadata and owner/fill classification, so an inclusion or omission fails mechanically rather than surviving a prose census.

2. **HUMAN DECISION remains open.** The operator-authored, digest-pinned `config_change`, restart-effective, no-lane-surface authorization shape remains coherent. Explicit operator ratification or override is still required before approval/lock; this review does not proxy that decision.

### Accepted Rev2 Folds

- The locked form-validation -> lineage -> `slot_in` -> `observe_gate` sequence is now preserved. Stage-split completeness and the m-3 output-contract / m-7 pipeline-stage seam close the rev1 reorder blocker without amending CQ-5.
- Catalog digest semantics are now honest: the composite digest binds bytes, while a separate m-7-owned reader capability/version gate proves enforcement. No digest mismatch is claimed to establish reader capability.
- The summary, grammar, GRILL_LOCK, impact, and fixture language consistently distinguishes knob fail-closed behavior from catalog fail-open behavior.
- The ownership/render distinction, immutable config-derived layer context, stale-digest short-circuit, mechanically readable version carrier, and forward-only history behavior remain accepted.

### Re-review Bar

Return one narrow revision that corrects the CQ-1 class census and exact-set fixture, then carries the operator ratification or remains explicitly `human-decision-required` until it arrives. No code, PLAN, IMPL, c1 atom change, CQ-5 amendment, or broader grammar redesign is authorized.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of rev2 `s8-design-m2-grammar` against `frank@2e1b4f0`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean and synchronized at `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-design-m2-grammar` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner corrects the CQ-1 matrix and exact-set fixture, preserves the accepted rev2 folds, and returns a revised DESIGN after or alongside operator ratification.
