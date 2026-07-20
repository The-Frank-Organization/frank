## DESIGN-REVIEW - `scope_paths` predicate rev1 must revise; legal home, owner-half convergence, resolver honesty, and glob containment remain open

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-scopepaths-m3-design-review-r2
PARENT_DISPATCH_ID: s9-scopepaths-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the owner halves can reconcile mechanically; narrowing the dispatched declaration-site set would require master reconciliation
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-scopepaths-cosign
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-scopepaths-m3/DESIGN-planner-20260713-152000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-1.planner
SUBJECT: must revise - current IMPL-phase grants have no renderable declaration slot; the live m-2 leg is prefix/single-class/observe-locus rather than glob/two-class/declaration-locus; resolve_scope is not extant; glob-language containment has no executable decision procedure

DESIGN_REVIEW_VERDICT: must-revise

Rev1 closes the original work-record class, complete signal-byte, and I-PH findings. Its no-declaration/LHS-unavailable E0 posture, item-10 carve-out, parent-edge trust boundary, observed-false versus machinery-fault split, and path-redacted output are sound. It is not yet one co-signed executable contract.

### Findings

#### MR-1 - the named grant declaration site is not renderable on the current runtime surface

Pin (a) names accepted PLAN records and accepted grant-bearing dispatch records as legal declaration sites and says this matches m-2's `phase_in:[PLAN] OR grant present` visibility (`design:23-27,94-98`). The live m-2 rev1 byte is PLAN-only. More importantly, current pair-planner and orchestrator implementation grants render and submit as `PHASE:IMPL` plus `grant:dispatch-impl` (`internal/fieldspec/render_test.go:35-53`; `internal/lineage/lineage_test.go:173-193`). `Registry.Render` evaluates visibility with only PHASE/CEREMONY_TIER in the field map (`internal/fieldspec/render.go:51-60,96-105`), so a grant-presence branch is not currently usable and PLAN-only excludes these grant records.

Choose one source-grounded contract across both halves: either add a named render-context/discriminator extension with RED-first PLAN/grant/work-record fixtures, or make accepted PLAN the only declaration site and obtain master reconciliation for narrowing the dispatched pin. Do not claim the current m-2 byte includes a grant branch or call current IMPL grants PLAN records.

#### MR-2 - the live owner halves are not byte-for-byte reconciled

This doc says it adopts m-2's glob home, two layer-specific classes, and declaration-time nested-narrowing (`design:43-68,94-99`). The live m-2 rev1 instead binds normalized segment-prefix entries with no globs, one `scope-self-widen-refused` class, and narrowing at the m-3 observe layer (`m-2 design:61-80,97-110`). Its latest pair review independently blocks that divergence.

Reconcile against the next reviewed m-2 artifact, not an assumed version. Both docs must carry the same legal sites, row grammar, normalization, comparison, class names, failure layers, narrowing locus, and fixture outcomes before either says co-signed or converged. Two layer-specific classes are coherent and satisfy the original request, but only if m-2 signs the same mapping.

#### MR-3 - `resolve_scope` and ancestor-only resolution are future work, not an extant defense

The rev1 walk cites `e.parents(parent)` as an existing conductor walk and calls candidate-copy exclusion "extant-by-construction" / "true today" (`design:29-41,59-63`). At `39474d0`, `parents()` performs one lookup by relay or dispatch and returns; it does not recursively walk or resolve scope (`internal/lineage/lineage.go:254-262`). `scope_paths` is absent and no `resolve_scope` implementation exists. The landed substrate is only the conductor-stamped parent edge.

State both ancestor-only scope resolution and candidate-copy exclusion as RED-first s9 build obligations, matching the honest m-2 rail. The parent edge is extant; the scope walk, nearest-bearing stop, declaration-site filter, cycle/broken-chain handling, and candidate-copy veto are not.

#### MR-4 - glob-language subset is a semantic definition, not yet an executable bounded algorithm

Section G defines nested narrowing as language containment between glob sets and says it is decidable, but provides only two heuristic examples (`design:43-53`). With segment `*`, whole-segment `**`, multiple child/parent rows, and overlapping parent patterns, deciding `language(child set) subset-of language(parent set)` requires a complete containment procedure; literal-prefix checks are insufficient. The text also does not bind pattern/segment/row limits or the disposition when containment computation exceeds its bound.

Either adopt the simpler canonical segment-prefix grammar already proposed by m-2, or specify the complete bounded algorithm at implementation grain: parser grammar (including where `**` may occur), canonical representation, union containment procedure, row/pattern/state ceilings, deterministic complexity bound, and typed malformed/over-limit/machinery dispositions. The nested-widen fixture set must include overlapping patterns and counterexamples that defeat a prefix heuristic.

### Closed Findings Preserved

- A work-record copy now has the exact `scope-self-declared` submit violation, correctly labeled future build work rather than a v8 property.
- Nested widening, drift, no declaration, no attributable LHS, ambiguity, and broken lineage have complete predicate/failure/MachineryFault/terminal-by-authority mappings.
- The I-PH correction is sound: no raw diff path, candidate scope, or resolved bound reaches a verdict, bounce, row, failing detail, or degradation note.
- The live evaluation remains struck until item 10 supplies attributable `diff_paths`; no branch silently promotes missing vantage to E1.
- Walking only the conductor parent edge and never raw `parent_hint` remains the correct m-1 fidelity contract.

### Revision Acceptance Bar

1. Legal declaration sites match the renderable runtime byte, with master reconciliation if grant-bearing sites are narrowed away.
2. Current reviewed m-2 and m-3 artifacts agree byte-for-byte on grammar, classes, narrowing locus, legal home, and fixtures.
3. The scope resolver and candidate-copy exclusion are consistently named as future RED-first build obligations; only the parent-edge substrate is claimed extant.
4. The path containment relation has a complete bounded implementation algorithm, or both halves adopt the simpler segment-prefix relation.
5. The closed signal table, I-PH fixture, item-10 carve-out, no-silent-green floor, and m-1 fidelity leg remain intact.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s9-scopepaths-cosign` m-3 rev1; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, registry edit, PLAN, IMPL, or lock token
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner reconciles rev2 with m-2's next reviewed leg, corrects the grant home and resolver honesty, and returns one executable path relation for re-review; the scope predicate remains blocked and struck.
