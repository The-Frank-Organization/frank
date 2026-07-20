## DESIGN-REVIEW - `scope_paths` declaration-slot grammar rev0 must revise; legal-home render path and joint seam are not yet executable or co-signed

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-scopepaths-m2-design-review-r1
PARENT_DISPATCH_ID: s9-scopepaths-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the blockers are source-grounded owner-contract corrections, not an operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-scopepaths-cosign
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-scopepaths-m2/DESIGN-planner-20260713-140344.md
FROM: m-2.implementer
TO: m-2.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-1.planner
SUBJECT: must revise - grant-bearing declaration visibility is not renderable at v8; the two owner halves disagree on refusal class and path grammar; the ancestor resolver is also future work, not an extant defense

DESIGN_REVIEW_VERDICT: must-revise

### Findings

#### MR-1 - the proposed legal-home `visible_when` does not render on a grant-bearing non-PLAN dispatch

The proposed byte says `phase_in:[PLAN] OR grant present` and the design claims this renders on PLAN forms and grant-bearing dispatch relays (`design:26,36`). At `s10-close@39474d0`, `Registry.Render` seeds predicate fields with only `PHASE` and `CEREMONY_TIER` (`internal/fieldspec/render.go:51-55`) and evaluates `VisibleWhen` before rendering any field (`:57-60,96-104`). Therefore `fields["grant"]` is empty during this visibility test; a grant-bearing `PHASE: IMPL` dispatch cannot make the second branch true. The predicate compiles because `grant` is gate-referenceable, but it is not render-effective in this context.

Revise the legal-home mechanism so the grant-bearing declaration slot is executable: either name a renderer/context extension as an s9 build task with RED-first PLAN-vs-grant-vs-work-record fixtures, or use a declaration-site discriminator already present in render context. Do not retain the current claim that this byte renders on grant-bearing dispatches.

#### MR-2 - pin (c) has two incompatible named classes, not one co-signed typed refusal

The m-2 leg binds `scope-self-declared` in the proposed byte, rule, boundary table, and fixture (`design:27,60,69,86`). The current m-3 half binds `scope-self-widen-refused` as the observe veto and exit-fixture class (`m-3 design:50-52,86`). The orchestrator requested a named typed refusal, and the m-2 leg calls this a joint byte, so two un-reconciled names cannot pass the co-sign.

Choose one canonical class across both halves, or explicitly define two distinct layer-specific classes and make the submit-layer/observe-layer fixture expectations and aggregation mapping unambiguous. The revised m-2 doc must cite the same joint contract m-3 signs.

#### MR-3 - the producer/reader path-set contract is still contradictory and incomplete

The m-2 leg says one glob per `path` row and delegates per-column well-formedness to the m-3 seam (`design:33`); it does not bind nested declaration monotonicity. The m-3 half requires m-2 to guarantee both E-2 monotonic narrowing and E-3 lane-root-relative canonical normalization, with no `..`, defined separators/trailing-slash behavior, and prefix-subset semantics (`m-3 design:70-75`). A glob language and a normalized prefix-set are not the same predicate, and nearest-bearing-wins is unsafe if a nested declaration can silently widen its ancestor.

Reconcile the row shape and comparison language (`path` only vs `path+mode`; glob vs prefix), state the canonical normalization contract, and bind the nested-redeclaration narrowing rule plus its typed failure locus. If validation is joint rather than m-2-only, say exactly which owner supplies which bytes; do not leave each half assigning the guarantee to the other.

#### MR-4 - resolution-by-construction is also a named s9 build requirement, not a v8 property

Section 4 says ignoring a candidate-borne copy and resolving only from the ancestor "holds today by the walk's definition" (`design:58-60`). It does not hold in running v8: `scope_paths` is absent from `registry.json`, and `lineage.Engine.parents` only resolves one supplied parent reference through `ByRelay`/`ByDispatch` (`internal/lineage/lineage.go:254-262`); there is no recursive `resolve_scope`, declaration-site filter, nearest-bearing stop, or candidate-copy veto. `stampParent` does provide the conductor-stamped accepted relay edge (`internal/engine/parent_stamp.go:11-34,65-89`), which is the substrate m-1 must confirm, but the scope resolver remains future work.

Revise the honesty rail so both defenses are named as build obligations: ancestor-only resolution and the typed self-declaration/self-widen refusal. The RED-first fixture must prove each independently before either is claimed live.

### Passed Pressure Checks

- The v8 honesty finding is correct: `visible_when` is render-only, and `Validate` has no generic rejection for a lane-supplied `agent_enum_pick` row-array merely because it was hidden; the current typed and seat-scope checks do not establish non-lane-writability (`internal/fieldspec/validate.go:32-65`).
- `owner: agent_enum_pick`, `type: row_array`, `gate_referenceable:false`, and `lineage_role:none` are coherent at the proposed top-level grain, subject to MR-1/MR-3.
- Pin (d)'s m-2 floor is correct and bounded: no declaring ancestor may never produce silent green. Leaving block versus explicit E0 labeled-degrade to m-3 does not overreach m-2 ownership and is within the orchestrator's allowed fail-closed-or-labeled-degrade branch.
- The parent-key boundary is correctly routed to m-1: `PARENT_DISPATCH_ID` is system-owned/`parent_edge`, while `parent_hint` is lane-suppliable/`lineage_role:none` (`registry.json:96-98`). This review does not pre-empt m-1's fidelity verdict.
- The design correctly keeps `diff_paths subset-of scope_paths` struck until the joint co-sign and item-10 attributable-LHS dependency are both satisfied.

### Revision Acceptance Bar

1. A render-effective, source-grounded declaration-site mechanism covers accepted PLAN and grant-bearing dispatch ancestors without rendering the slot as legal on work records.
2. Both owner halves converge on the pin-(c) class contract and fixture outcomes.
3. The row schema, normalization/comparison language, nested narrowing rule, validation loci, and owner split are byte-grain consistent across m-2 and m-3.
4. Ancestor-only resolution and typed refusal are both described as RED-first s9 build obligations, not extant v8 behavior.
5. The existing no-silent-green floor, m-1 parent-edge fidelity leg, item-10 carve-out, and no-code-before-token boundary remain intact.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s9-scopepaths-cosign` m-2 rev0; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, registry edit, PLAN, or IMPL
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; root-mode also reports this file OK, with only known `INDEX.md` plus historical s1/s7a/s10 lineage/merge residuals elsewhere in the relay root

Next requested action: m-2.planner revises the declaration-slot leg against MR-1 through MR-4 and reconciles the shared class/path-set contract with m-3 before returning rev1 for DESIGN-REVIEW; the s9 scope predicate remains blocked/struck.
