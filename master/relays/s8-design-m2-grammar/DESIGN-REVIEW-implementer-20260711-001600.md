## DESIGN-REVIEW - s8 config/atom grammar rev5 must revise Option-B stage and fixture echoes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-m2-grammar-design-review-r6
PARENT_DISPATCH_ID: s8-m2-grammar-design-r6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - the two operator legs remain required at the reconciled lock and are not proxied by this review
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s8-m2-grammar-grill
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m2-grammar/DESIGN-planner-20260711-001500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-5.planner, m-6.planner, m-7.planner
SUBJECT: rev5 must revise - owner confirms and Option B are accepted, but the numbered stage rule and exact-set fixture still encode the superseded global-predicate model

DESIGN_REVIEW_VERDICT: must-revise

### Findings

1. **BLOCKER - the executable stage rule still contradicts the confirmed partition and Option B.** Section 8 accurately records m-7's `authority_class` step-3 placement, m-3's ten-field step-4 manifest, m-5/m-6 Option B, and step 4.5. But the primary numbered rule still says the post-observe check validates outputs “the observe hook produced” for required class 2/3, and that all system/computed rows retain CQ-1 `required_when` (`design:65`). That is false after the confirmed reconciliation: `authority_class` is a step-3 m-2 value, `surface_intent` is not an m-3 output, and Option B explicitly removes `surface_intent`'s global `required_when`.

   Residual lock-bearing echoes remain nearby: the matrix still says the `surface_intent` amendment is “pending” (`design:57`); the sequence says only “m-7-confirmed pre-disposition” rather than the confirmed **step 3**, and does not name derivation/completeness as **step 4.5** (`design:64-65`); the seam calls Option B a `surface_intent` “predicate fix” (`design:71`); §4 calls it a “registry-predicate amendment” (`design:88`), although Option B removes static applicability predicates.

   **Required revision:** make the primary sequence byte-identical to the owner records:
   1. form validation handles every seat-fillable activated predicate;
   2. lineage, then step-3 classification computes `authority_class` under the s5 ③ raise-rewrites tripwire;
   3. m-3 step 4 produces its confirmed manifest plus the separately noted git-ref/control outputs;
   4. step 4.5 applies the confirmed producer/profile manifests, derives Option-B `surface_intent` only for non-gate records (`progress` otherwise), and validates all applicable outputs before commit.

   Sweep the matrix, numbered rule, seam, §4, GRILL_LOCK, and design-lock impact. Use “registry-row amendment removing the static predicates,” not “predicate amendment/fix,” and mark all four confirms complete.

2. **BLOCKER - the independent fixture equality is impossible under Option B as currently written.** The fixture defines **expected** solely as fields gated by `layer_present:observe` in the registry, then requires that set to equal the union of producer outputs (`design:72`). Option B removes `surface_intent`'s global `required_when`/static applicability predicate, while the confirmed m-5/m-6 profile manifest still derives `surface_intent` on every non-gate record. Therefore the producer union contains `surface_intent` and the registry-derived expected set does not; a correct Option-B implementation fails the stated no-extra equality.

   **Required revision:** define independent expectations at their actual homes:
   - registry expectation: activated FieldSpec requirements, partitioned by seat/system ownership;
   - producer/profile expectation: confirmed m-3 manifest, m-2 step-3 classification, and m-5/m-6 Option-B applicability/derivation;
   - actual: assembled record plus producer-write provenance at step 4.5.

   Assert each expectation against actual independently, then assert committed-record behavior for gate and non-gate profiles. Specifically prove: gate-bearing record has no `surface_intent`; non-gate record gets exactly one valid value and defaults to `progress`; m-3 cannot write `authority_class`/`surface_intent`; no required registry field is absent; no producer writes outside its manifest. Do not force the predicate-free profile obligation into equality with the registry predicate set.

3. **HUMAN DECISIONS remain open at the reconciled lock.** This review does not proxy the operator's activation-authorization ratification or m-3's three grill defaults. Both remain required on record per the master close-out.

### Accepted Rev5 Reconciliation

- The four owner returns and master close-out are authoritative and mutually consistent.
- m-3's ten-field Block-A manifest, two exclusions, git-ref/veto precision, and Step-2 defaults are accepted.
- m-7's `authority_class` step-3 placement, step-4.5 amendment vehicle, and four hosting constraints are accepted.
- m-5/m-6 Option B is accepted: no `gate_bearing` atom, no static `surface_intent` requiredness, total conductor derivation for non-gate records, absence on gate-bearing records.
- The prior form-validation, census, pipeline-order, reader-capability, layer-context, stale-digest, version, and forward-history folds remain accepted.

### Re-review Bar

Return one consistency-only revision that rewrites the primary sequence and fixture to the confirmed step-3/step-4/step-4.5 Option-B contract. No new owner consultation is required unless the fold changes an owner-confirmed decision. Keep both operator legs explicitly pending until recorded. No code, PLAN, IMPL, registry edit, c1 edit, or silent §3 amendment is authorized.

ACTIONS_GIT_REF: none - read-only final DESIGN-REVIEW of rev5 `s8-design-m2-grammar` against the four owner confirms and `frank@2e1b4f0`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean and synchronized at `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-design-m2-grammar` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner performs the bounded Option-B sequence/fixture consistency sweep and returns the final DESIGN while the two operator legs proceed in parallel.
