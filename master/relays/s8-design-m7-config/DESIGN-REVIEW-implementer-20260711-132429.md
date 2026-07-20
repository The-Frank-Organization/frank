## DESIGN-REVIEW - s8 config host r7 must revise the version-transition predicate

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r7
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the operator selected F5 fork (a) in `SITREP-orchestrator-planner-20260711-130825.md`; only technical findings remain
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-132018.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r7 must revise - F2's capability table is now concrete, but F1's key-path predicate neither recognizes the full schema surface nor defines forward transitions for string markers

DESIGN_REVIEW_VERDICT: must-revise

F2's per-member capability table is accepted: the marker forms, initial values, supported ceiling/sets, marker-first load order, and typed pre-interpretation failure are concrete enough for PLAN. F1 is correctly moved into the trusted acceptance path and recovery correctly replays accepted history only. Three defects remain in the transition predicate itself.

## Findings

### F1 - An untyped JSON key-path set does not detect the schema changes section 1.1 requires a bump for

Section 1.1 explicitly includes a type change in the engine schema surface. Section 2.4 compares only deep JSON key-path sets. `{"x":1}` and `{"x":"one"}` have the same key-path set, as do scalar-to-container changes at an existing path under common leaf-path definitions. They therefore pass unchanged-version validation even though the design says they require a bump. Arrays make the rule additionally ambiguous: the design does not say whether indices participate, how element shapes are normalized, or how ordinary catalog row/value changes avoid becoming false schema drift.

Required fold: replace "key-path set" with a mechanically complete, canonical schema-surface predicate. It must distinguish JSON node kinds and nested object/array shape while remaining insensitive to ordinary values and list indices. A handler-declared strict schema validator or version-to-schema descriptor is the simpler robust shape; a structural signature is also viable only if its exact object/array normalization is specified. Add executable unchanged-marker rejects for a same-path scalar type change and a container/element-shape change, plus an unchanged-version accept for a genuine value-only change.

### F2 - The biconditional contradicts the stated lawful semantics-only bump

Section 2.4 and the GRILL_LOCK say "version unchanged is lawful iff the key-path set is unchanged." That biconditional makes unchanged shape imply unchanged version. The next bullet and FX-CFG-10(f) instead say a version bump with unchanged shape is lawful. Both cannot define the acceptance validator.

Required fold: state the one-way invariant directly: an unchanged marker requires an unchanged canonical schema surface; a changed schema surface requires the member's lawful forward version transition. A marker bump with unchanged schema may remain lawful as the explicitly named semantics-only residual. Sweep the relay summary, GRILL_LOCK, and fixture wording to the same implication-based rule.

### F3 - "Changed marker" is not a successor relation for catalog and fieldspec versions

For catalog/fieldspec, section 2.4 accepts any changed marker when shape changes and delegates the successor value to an owning-domain grammar, but the capability table declares only exact supported sets and no allowed-transition relation. Once a future handler supports both old and new markers, a rollback from the new marker to the old supported marker is merely "changed" and passes both checks. A skip across supported markers has the same problem. The current singleton sets hide the defect but do not define the evolution mechanism the design claims.

Required fold: make each string-marker handler declare an allowed forward-transition relation (or an ordered successor map supplied by the owner contract) in addition to its supported set. Acceptance requires that relation for shape-changing transitions; rollback and skip between otherwise supported markers are typed `config-version-transition` rejects. Add those two executable legs for a simulated multi-marker handler. This consumes the owner-defined vocabulary; it does not let m-7 invent m-2's marker sequence.

## Confirmed

- The acceptance gate is correctly located in the serialized loop before pivot, and its reject class is path-free.
- Recovery replay of already-accepted transitions without re-validation preserves store-is-truth.
- The closed F2 capability table resolves the r6 parser/comparison/initial-value ambiguity.
- The corrected two-member-reader versus three-member-store digest argument and FX-CFG-9 remain valid.
- The F5 operator gate is satisfied; no further operator decision is required for these findings.
- All previously accepted descriptor, digest, canonical-census, step-4.5, and activation folds remain approved.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, and no effective reconciled lock.

ACTIONS_GIT_REF: wrote this r7 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F3 into config-host r8 and returns a DESIGN relay for re-review; master holds reconcile-A completion meanwhile.
