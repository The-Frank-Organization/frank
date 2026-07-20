## DESIGN-REVIEW - m-2.implementer - Forms & Determinism

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c1-design-m-2-review
PARENT_DISPATCH_ID: c1-design-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - design revision only; no operator decision required
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
BUNDLE_ID: m-2-forms-determinism
OWNER: m-2 (Forms & Determinism)
ACTIONS_GIT_REF: created this DESIGN-REVIEW relay at master/relays/c1-design-m-2/DESIGN-REVIEW-Implementer-20260628-173310.md; no code/source changes; cwd git status unavailable because workspace root is not a git repository
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; git status reports fatal not-a-git-repository at /mnt mount boundary

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

This is a narrow must-revise, not a rejection. The design is directionally correct and preserves the m-2 audit conclusions: bespoke FieldSpec registry, fill-time authority, strict form-only submission, bounded predicates, X- overflow as non-gating, and separate form-validation plus lineage are the right architecture. Two contract ambiguities must be fixed before design lock because they affect v2.8.8 authority preservation and the m-1/m-2 boundary.

## Blocking revisions

1. Clarify whether `ROLE` is envelope/system-owned or header/form-owned.

Evidence:
- The layer table puts `ROLE` in the typed header/form layer with other header fields. Source: [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:23), [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:24).
- The m-1 boundary contract later classifies `ROLE` as an envelope/system field stamped by m-1, alongside `FROM`, `relay_id`, `DISPATCH_ID`, `PARENT_DISPATCH_ID`, `timestamp`, and `schema_version`. Source: [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:167).
- The coordination relay also asks m-1 to stamp `ROLE` from the same channel as `FROM`, so the v2.8.8 ROLE/FROM tripwire dissolves by construction. Source: [COORD-planner-20260628-172553.md](/mnt/c/Users/lijia/programming/harness/master/relays/c1-design-m-2/COORD-planner-20260628-172553.md:1).

Why this blocks lock: the design's claim that `role_from_consistency_error` dissolves depends on a single source of truth for `ROLE`. If `ROLE` remains a header pick while `FROM` is courier-stamped, the ROLE/FROM mismatch class is not structurally eliminated. If the intended model is that `ROLE` is system-owned and may be rendered as a read-only header view, say that explicitly in the layer table and FieldSpec examples.

Required revision: make `ROLE` unambiguously `owner: system`. Either move it to the envelope layer or state that header rendering is a read-only projection of the system-owned envelope value. Then update §10b and §12 so they use the same layer/owner model.

2. Clarify lineage-engine gating order: invalid authority lineage must block delivery/consumption, not merely be detected after append.

Evidence:
- The render/validate flow says send runs form-validation and observe-as-send, then the relay is stamped and emitted atomically, and the cross-relay lineage engine runs after that over the append-only store. Source: [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:64), [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:68).
- The design correctly classifies design-review lineage, pair-Planner dispatch lineage, non-addressee IMPL trap, merge lineage, scope-flip drift, and orchestrator-review visibility as surviving cross-relay lineage gates. Source: [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:145), [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:148), [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:151), [2026-06-28-v3-form-schema-design.md](/mnt/c/Users/lijia/programming/harness/master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:153).
- v2.8.8 protocol expects relay-lint before handoff/delegated dispatch and treats lint errors as blockers unless waived. Source: [protocol.md](/mnt/c/Users/lijia/programming/harness/extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/agent-pair-implementer/protocol.md:397), [protocol.md](/mnt/c/Users/lijia/programming/harness/extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/agent-pair-implementer/protocol.md:407).

Why this blocks lock: if an invalid design-lock PLAN, delegated dispatch, merge claim, or non-addressee IMPL report can be emitted/delivered before lineage validation, v3 weakens a v2.8.8 authority gate. An append-only store can record failed attempts, but consumers must not act on authority-bearing records until lineage passes.

Required revision: define a two-state write path such as `attempted/submitted` -> lineage validation -> `accepted/deliverable`, or state that "emitted" means "stored as an attempted record but not deliverable/consumable until lineage passes." Also name which phases/records are lineage-blocking before delivery: at least design-doc PLAN locks, delegated dispatch/merge grants, substantive IMPL reports, merge claims, and orchestrator authority relays.

## Review asks

1. Dissolve/survive table: mostly sound. I do not see a dropped v2.8.8 check except for the `ROLE` ambiguity above: the `ROLE/FROM` tripwire only dissolves if both values are system-owned from one channel. The authority content of dispatch/merge tokens is preserved as typed grant fields; only lexical token detection dissolves.

2. Bounded required-when vocabulary: sufficient for intra-relay form shape as written. Cross-relay predicates such as "parent is an earlier approving DESIGN-REVIEW" correctly belong to lineage, not `required_when`. Do not add arbitrary function predicates.

3. Unpreservable gates beyond GATE-1: no additional gate appears unpreservable, provided revision 2 makes lineage blocking before delivery/consumption. Typed per-relay form alone cannot prove lineage, and the design correctly routes that to the lineage engine.

4. One-FROM / certification-reserve split: consistent with m-1's locked identity model, provided revision 1 resolves `ROLE` as system-owned alongside `FROM`. Certification reserve/defer is acceptable because `certification` is non-consumed until a verifier exists.

## Non-blocking notes

- The strict legacy rule is correct. A transitional markdown submission path would force retaining the prose-only checks that the design intends to delete.
- `X-` overflow is safe only if `consumers: []` and `lineage_role: none` are enforced by schema tests; acceptance criterion 4 should include a negative fixture for attempted `X-` gate input.
- The m-4 open question about routing-record header vs separate routing relay is non-blocking for m-2 as long as FieldSpec can model either a per-relay field or a linked per-dispatch record.

## Acceptance criteria review

- DESIGN_DOC_ID present and stable: pass.
- Boundary contract present: pass with blocking revision for `ROLE` layer consistency.
- Target entity present: pass; v3 typed relay envelope and FieldSpec registry.
- Acceptance criteria draft present: pass with blocking revision for lineage-blocking delivery semantics.
- Operator decisions/defaults present: pass; GRILL_LOCK records carrier, certification, strict legacy, versioning, overflow, and bounded required-when decisions.
- Rejected alternatives present: pass.
- Open questions present and non-blocking: pass.

Next requested action: revise the design doc narrowly on `ROLE` ownership/layering and lineage-engine delivery gating, then resend for DESIGN-REVIEW on the same DESIGN_DOC_ID.
