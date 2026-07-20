## DESIGN-REVIEW - s8 config/atom grammar rev1 must revise pipeline order and old-reader gate

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-m2-grammar-design-review-r2
PARENT_DISPATCH_ID: s8-m2-grammar-design-r2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - operator ratification of activation authorization remains lock-blocking after technical revision
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s8-m2-grammar-grill
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m2-grammar/DESIGN-planner-20260710-223000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-3.planner
SUBJECT: rev1 must revise - ownership split is corrected, but observe enrichment reorders locked CQ-5, catalog digest mismatch does not exclude old readers, and stale rev0 claims remain

DESIGN_REVIEW_VERDICT: must-revise

### Findings

1. **BLOCKER - the new enrichment-before-requiredness rule contradicts the locked conductor pipeline.** Rev1 says m-3 `observe_gate` fills/computes classes 2/3 before the observe-owned required-set is checked (`design:57,74,100,120`). The locked m-7 pipeline says the opposite order: step 2(a) is m-2 form validation including the phase-split required-set, step 2(b) lineage; step 3 classifies `slot_in`; only step 4 runs `observe_gate` (`v3-conductor-core-design.md:68-72,154-155`). m-3's CQ-5 co-sign likewise pins `slot_in` post-form/lineage and `observe_gate` after that (`v3-observe-evidence-design.md:114,217,238`). The document cannot claim “no c1 reopen” while silently moving m-3 ahead of m-2/lineage.

   **Required revision:** preserve the locked order and split completeness by stage:
   - pre-observe form validation checks stale digest, seat ownership/type/scope, and the seat-owned activated requirements;
   - lineage and `slot_in` remain before observe;
   - `observe_gate` fills conductor-observed/computed classes;
   - a post-observe completeness check validates the hook's allowed outputs before append.

   System/computed rows may retain CQ-1 `required_when` as their schema declaration, but they are not satisfied “before the required-set fires” by moving the hook. Name the post-observe completeness validator and its owner seam. If the intended design truly reorders form/lineage/observe, route a review-driven amendment to the locked CQ-5 pipeline instead of changing it here.

2. **BLOCKER - the catalog old-reader exclusion mechanism is false against current `config.Load`.** Rev1 says a reader that cannot enforce the catalog omits it from the config digest and is stopped by a pinned-digest mismatch (`design:38,104`). Current loading does not work that way: `config.Load` reads every supplied member, including unknown names, into `loaded` (`config.go:43-50`) and hashes all entries (`:66-70,74-82`), while interpreting only `engine` and optional `fieldspec` (`:52-63`). A pre-s8 reader can therefore receive the catalog member, compute the same composite digest, and still perform no catalog enforcement. Digest equality proves bytes, not reader capability.

   **Required revision:** make governed-runtime participation depend on a mechanically checked reader capability/version, not digest mismatch. Examples include a required config-format/engine version whose unsupported value fails load, or a required-member handler registry that rejects a pinned member lacking an enforcement handler. m-7 owns the mechanism, but the m-2 grammar must state the actual contract: digest binds bytes; a separate capability gate proves this reader enforces the catalog.

3. **BLOCKER - rev0 contradictions survive the fold and would mislead PLAN.** Full-document consistency must match the corrected matrix:
   - Section 1 still says the same Block-A fields become “required/rendered” knob-on (`design:19`), contradicting the three-class matrix at `:47-58`.
   - “One grammar, two members” still labels both unknown defaults `ignore-unknown-safe-default` (`design:43`), contradicting the catalog fail-open correction at `:38`.
   - GRILL_LOCK resolved decision `design:97` still calls `no-catalog-enforcement on old readers` a safe default, contradicting `:104`.

   **Required revision:** byte-consistency sweep all summary, GRILL_LOCK, design-impact, and fixture language. Use class-specific activation wording and distinguish activation fail-closed from catalog fail-open everywhere.

4. **HUMAN DECISION remains correctly surfaced, not yet closed.** The folded operator-authored, digest-pinned `config_change`, restart-effective, no-lane-surface position is coherent (`design:112-120`) and matches the m-7 dispatch. Operator ratification or override is still required before approval/lock. This review does not proxy that answer.

### Accepted Rev1 Folds

- The ownership/stage matrix correctly distinguishes seat-rendered, conductor-observed, and conductor-computed rows.
- One immutable config-derived layer context across render, validation, and grant-digest is the correct invariant.
- Stale digest short-circuiting required-set evaluation, followed by fresh-form validation, closes the double-bounce class.
- The knob requires a mechanically readable version carrier; knob value changes remain config-digest acts, not record schema bumps.
- Catalog legacy behavior is correctly identified as fail-open in the detailed section; only its proposed exclusion mechanism and residual summaries need correction.
- Forward-only read/project history remains sound and does not re-run current fieldspec requiredness.

### Re-review Bar

Return a new unique pair-design revision that:

- preserves the locked form/lineage/slot/observe order and names the post-observe output-completeness check;
- replaces digest-mismatch capability claims with an enforceable reader-version/handler gate owned at the m-7 seam;
- removes all three residual contradictory statements;
- carries the operator ratification or remains explicitly `human-decision-required` until it arrives.

No code, PLAN, IMPL, c1 atom change, or silent CQ-5 amendment is authorized.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of rev1 `s8-design-m2-grammar` against `frank@2e1b4f0`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean and synchronized at `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-design-m2-grammar` lineage lint exit 0 for both the live directory and this relay

Next requested action: m-2.planner folds the locked-order, reader-capability, and consistency findings, coordinates the m-3/m-7 seams, and returns a revised DESIGN after or alongside the operator ratification.
