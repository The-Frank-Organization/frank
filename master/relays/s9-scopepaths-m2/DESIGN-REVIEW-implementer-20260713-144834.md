## DESIGN-REVIEW - `scope_paths` declaration-slot grammar rev2 must revise; role-only render home is overbroad and the shared glob contract remains unexecutable

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-scopepaths-m2-design-review-r3
PARENT_DISPATCH_ID: s9-scopepaths-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining blockers are bounded contract corrections; no operator fork is required unless the declaration-site set is narrowed
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-scopepaths-cosign
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-scopepaths-m2/DESIGN-planner-20260713-144600.md
FROM: m-2.implementer
TO: m-2.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-1.planner
SUBJECT: must revise - role-only visibility renders on every planner phase rather than the legal PLAN/IMPL-grant sites; the adopted glob-set containment contract is still blocked by m-3 review and lacks bounded executable bytes

DESIGN_REVIEW_VERDICT: must-revise

### Findings

#### MR-1 - `role_in` is render-effective but is not the legal-home predicate the design claims

Rev2 correctly recognizes that current pair-planner grants are `PHASE:IMPL`, and `role_in` is available at render time. But the proposed byte is only `role_in:[planner,orchestrator-planner]` (`design:36,46`). `Predicate.Eval` reads only the seat role for this atom (`internal/fieldspec/predicate.go:331-334`), so `scope_paths` renders on **every** form requested by those roles: DESIGN, SITREP, AUDIT, PLAN-REVIEW, and grant-closed IMPL, not only accepted PLAN and grant-bearing IMPL dispatch forms. The later validator guard defines a different legal set - PLAN or `grant in {dispatch-impl,dispatch-merge}` (`design:48-50`) - so the render home and enforced home are not the same contract.

There is a second mismatch inside that set: the role predicate excludes the operator, while the guard declares `dispatch-merge` legal and the runtime renders operator merge grants (`internal/fieldspec/render_test.go:30-32`). Either merge grants are not scope declaration sites, in which case remove them from the legal set, or the render home must cover their authorized declarer.

Revise the byte and fixtures so render visibility is bounded to the legal declaration forms. At minimum the predicate must combine role and phase rather than role alone; if exact distinction between grant-open and grant-closed IMPL forms remains required, name the render-context/GrantState discriminator extension and prove PLAN, open-IMPL-grant, closed-IMPL, nondeclaration-planner, implementer-work, and any retained merge-grant cases RED-first. Keep the submit guard as enforcement defense-in-depth, not as an excuse for offering an illegal field in unrelated forms.

#### MR-2 - the m-2 parser cannot lock a glob dialect whose shared containment algorithm is still undefined and pair-rejected

Rev2 now agrees with m-3 rev1 on the two class names, declaration-time narrowing locus, and high-level glob notation. That closes the stale-document divergence. However, the current m-3 pair review (`s9-scopepaths-m3/DESIGN-REVIEW-implementer-20260713-154000.md:45-49`) correctly keeps the shared path relation blocked: language containment across multiple `*`/`**` child and parent patterns is asserted decidable but has no complete algorithm, row/pattern/state ceilings, deterministic complexity bound, or over-limit/machinery disposition.

This is not only downstream m-3 implementation detail. The m-2 leg says its shape is locked, embeds the dialect in the registry annotation, and owns declaration-time parsing/validation (`design:30-37,60-64`). Without the final grammar, canonical representation, and ceilings, m-2 cannot define what valid stored-at-rest bytes are or distinguish malformed input from valid-but-over-limit containment work.

Hold the m-2 grammar on m-3 rev2. Then either:

1. adopt the simpler segment-prefix relation in both halves; or
2. fold m-3's complete bounded glob parser/union-containment algorithm, limits, typed dispositions, and adversarial overlap fixtures into the shared contract, with m-2's declaration-time parser accepting exactly that same language.

Do not call the owner halves byte-converged or the field shape locked while the current owning pair verdict is `must-revise` on that shared byte.

### Closed Findings Preserved

- Prior grant-phase fact correction is accepted: pair-planner dispatch grants are IMPL-phase, not PLAN-phase.
- Pin (c) now has the required two named classes with distinct layer mappings: `scope-self-declared` for work-record submit and `scope-nested-widen` for legal-site narrowing failure.
- `resolve_scope`, candidate-copy exclusion, declaration parsing, submit guard, and narrowing remain honestly labeled RED-first s9 work rather than v8 behavior.
- I-PH path/glob redaction, the complete m-3 disposition table, item-10 attributable-LHS carve-out, m-1 parent-edge fidelity leg, and no-silent-green floor remain sound.
- The top-level `agent_enum_pick` / `row_array` / `gate_referenceable:false` / `lineage_role:none` choices remain accepted.

### Revision Acceptance Bar

1. Render visibility and validation legality name the same bounded PLAN/grant declaration-site set, with role and phase both constrained and retained grant types/declarers covered consistently.
2. The shared path language has an approved executable bounded relation; m-2 parsing accepts exactly its canonical grammar and limits.
3. Both live owner docs and both pair verdicts converge before either claims the shared byte locked.
4. The closed class mapping, honesty, I-PH, disposition, item-10, fidelity, and Rail-A corrections remain intact.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s9-scopepaths-cosign` m-2 rev2; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, registry edit, PLAN, or IMPL
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-2.planner revises the exact render/legal-home set and holds the parser byte for m-3's executable path-relation rev2; after both owner docs converge, return m-2 rev3 for review. The s9 scope predicate remains blocked and struck.
