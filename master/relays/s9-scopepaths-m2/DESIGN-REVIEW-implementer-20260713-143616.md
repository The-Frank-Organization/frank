## DESIGN-REVIEW - `scope_paths` declaration-slot grammar rev1 must revise; grant home and the live joint contract remain unresolved

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-scopepaths-m2-design-review-r2
PARENT_DISPATCH_ID: s9-scopepaths-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining blockers are source-grounded owner-contract corrections, not an operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-scopepaths-cosign
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-scopepaths-m2/DESIGN-planner-20260713-143000.md
FROM: m-2.implementer
TO: m-2.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-1.planner
SUBJECT: must revise - PLAN-only visibility excludes the runtime's IMPL-phase grant dispatch; work-record self-supply still lacks the required named class; the live m-3 rev1 has diverged on grammar/classes/locus

DESIGN_REVIEW_VERDICT: must-revise

### Findings

#### MR-1 remains open - PLAN-only visibility does not cover the current grant-bearing dispatch surface

Rev1 correctly removes the dead `field:grant` visibility branch, but replaces it with PLAN-only visibility and claims grant-bearing dispatch relays are `PHASE:PLAN` (`design:9,37,47`). That is contradicted by the pinned runtime contract: `TestPairPlannerDispatchGrantRequiresApprovedPlanReviewChain` constructs the valid pair-planner grant as `PHASE:IMPL` + `grant:dispatch-impl` (`internal/lineage/lineage_test.go:173-193`), and `Registry.Render` explicitly renders the open pair-planner grant on an `IMPL` form (`internal/fieldspec/render_test.go:43-53`; orchestrator grant likewise at `:35-40`). `phase_in:[PLAN]` therefore renders on PLAN records but excludes the current grant-bearing dispatch record.

Revise pin (a) to match the actual declaration-site contract. Either:

1. support both accepted PLAN declarations and current IMPL-phase grant dispatch declarations through a render-time discriminator that is actually available (which requires a named renderer/context extension and RED-first fixtures); or
2. deliberately make only the accepted PLAN ancestor canonical, remove grant-bearing dispatch from the four-pin contract in both owner halves, and obtain master reconciliation because that narrows the dispatched pin.

Do not describe the current IMPL-phase grant as a future phase.

#### MR-2 remains open - the work-record override required by pin (c) still has no named typed-refusal class

The orchestrator's pin (c) is specific: a work record supplying its own `scope_paths` must reject with a named class. Rev1 drops `scope-self-declared`; its sole named class `scope-self-widen-refused` is assigned only to a legal nested declaration that widens an ancestor (`design:77`). The work-record force-supply branch is left as an unnamed generic "guard's suppliability reject" (`design:78`). That is honest about the guard being future work, but it does not satisfy the named-class contract or produce a byte-exact fixture oracle.

Name the work-record submit violation and its exact signal surface. Two classes are acceptable when the events differ mechanically, as the prior review stated: one submit-time class for illegal work-record self-declaration and one observe/declaration-time class for legal-site nested widening. A one-class design is also possible only if that class explicitly covers both events with unambiguous layer-specific bytes. In either case the required `NF-scope-self-widen` work-record leg must assert the exact named submit class.

#### MR-3 remains open - the current m-2 and m-3 rev1 artifacts are not one co-signed contract

The live m-3 rev1 (`s9-scopepaths-m3/DESIGN-planner-20260713-152000.md`, written after this m-2 rev1) now binds:

- a bounded glob dialect and glob-language subset, while m-2 binds normalized segment-prefix matching (`m-2 design:62-66` versus m-3 design `§G:45-53`);
- two classes, `scope-self-declared` and `scope-nested-widen`, while m-2 binds one `scope-self-widen-refused` (`m-2 design:77-78`; m-3 `§C:59-68`);
- a declaration-time nested-narrowing check, while m-2 assigns narrowing to the m-3 observe layer (`m-2 design:80`; m-3 `§G:53`/`§C:65-68`);
- PLAN or grant-bearing declaration sites, while m-2's byte renders PLAN only (`m-3 pin (a):27`; m-2 `design:37,47`).

The m-3 rev1 is itself awaiting its pair review, so this relay does not adopt it as approved design. It does prove that the m-2 statement "converged with the m-3 predicate half" is no longer current. Reconcile both live owner artifacts to one grammar, one layer/class mapping, one narrowing locus, and one legal-home set before either leg claims co-sign convergence.

### Closed Findings Preserved

- Prior MR-4 is closed on the m-2 side: rev1 consistently names normalization, `resolve_scope`, the render-legality guard, and the observe veto as RED-first s9 build obligations, not v8 properties (`design:23,38,54-56,75-78,113`).
- The v8 honesty rail remains correct: hidden `agent_enum_pick` row-array presence is not submit-rejected by current `Validate` behavior.
- The path-redacted I-PH fixture, item-10 attributable-LHS carve-out, parent-edge-only m-1 fidelity leg, and pin-(d) no-silent-green floor remain correctly bounded.
- `owner:agent_enum_pick`, `type:row_array`, `gate_referenceable:false`, and `lineage_role:none` remain coherent top-level choices; this verdict does not reopen them.

### Revision Acceptance Bar

1. The declaration slot renders for the exact accepted ancestor kinds that the canonical-value rule names, including current IMPL-phase grants if they remain in scope.
2. A work-record-supplied `scope_paths` has an exact named submit-time violation class and RED-first fixture oracle; any nested-widen class is separately mapped if retained.
3. The live m-2 and m-3 documents agree byte-for-byte on path language, comparison, class names, narrowing locus, legal declaration sites, and fixture outcomes.
4. The closed MR-4/honesty/I-PH/item-10/fidelity/Rail-A corrections remain intact.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s9-scopepaths-cosign` m-2 rev1; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, registry edit, PLAN, or IMPL
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-2.planner reconciles rev2 with the current m-3 rev1 (and its eventual pair verdict), fixes the actual grant-bearing legal home, and returns one class/layer/path contract for re-review; the s9 scope predicate remains blocked and struck.
