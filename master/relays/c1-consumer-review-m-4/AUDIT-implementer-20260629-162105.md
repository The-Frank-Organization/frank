## Team m-4 Implementer - independent consumer review of m-1/m-2 foundations

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-consumer-review-m-4
PARENT_DISPATCH_ID: c1-consumer-review-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
BUNDLE_ID: c1-consumer-review-m-4
OWNER: m-4 (Routing & Policy), as consumer lens

CONSUMER_REVIEW_VERDICT: sufficient

PRIMARY_BUCKET: already-closed
still-open: no blocking m-4 foundation gap found; m-4 still owes its own later routing-domain design.
already-closed: the m-1/m-2 foundation already provides the typed field slots and stamped-store write target needed by m-4 once Q-C is resolved as a dedicated routing relay.
product-overlapped: none.
recommended-next: m-4.planner should reconcile the paired reviews; orchestrator can then fold the Q-C answer into the m-1/m-2 joint lock without a separate m-1 or m-2 gap relay.

## Evidence Basis

- m-1 `submit()` is the sole write path, stamps `FROM`/`ROLE` from the bound seat connection, runs gates before append, appends the record/INDEX row, and projects to mailboxes: `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:75-82`.
- m-1's on-disk contract makes markdown a rendered projection of the canonical typed envelope, with immutable append-only records and system-owned `FROM`, `ROLE`, `relay_id`, `DISPATCH_ID`, `timestamp`, `schema_version`, and `certification`: `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:84-91`.
- m-1 explicitly names m-4's consumer contract: stamped store as routing-record write target; routing record `FROM` equals the router seat's stamped identity; routing records ride `submit()` like any relay: `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:93-101`.
- m-2's FieldSpec registry has first-class owner, type, enum, seat-scope, required/visible predicates, consumers, and lineage-role slots: `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:31-41` and `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:43-64`.
- m-2 explicitly names the m-4 routing-record fields: `role+model-per-dispatch`, `capability_prior`, `justified_deviation`, `record_kind`, and benchmark/outcome-feedback handle: `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:168-175`.
- The prior v3 routing pillar locks altitude B and the staged policy requirement: role+model per dispatch, static capability priors, recorded deviation justification, and v3.1 feedback handle shaped on day one: `extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md:14-18` and `extracted/agentic-dev-team-skills-v3-export/v3-design/agentic-dev-team-skills-v3-roadmap.md:44-60`.

## Review Questions

1. Do the routing-record field slots express altitude-B plus priors, deviation, and feedback?

Yes, at the foundation level. The typed FieldSpec carrier can express the required fields as typed, owner-scoped slots, and m-2 names m-4 as the consumer. The later m-4 domain design should refine the exact field IDs, but the foundation does not need a new primitive.

Required shape for the later m-4 design:
- `routing_assignments`: object or row_array, planner-emitted, carrying role plus selected model for each dispatched seat.
- `capability_prior_ref` or `capability_prior_snapshot`: system/config-owned reference or object naming the static prior used.
- `justified_deviation`: free_text required when selected model differs from the prior floor.
- `routing_record_kind`: routing-owned enum using the DESIGN_RECORD_KIND pattern, not design-domain literal values.
- `benchmark_case_ref` / `outcome_feedback_ref`: nullable id_ref fields reserved now for v3.1.

2. Resolve Q-C: header field or separate routing relay?

Answer: separate routing RELAY, with the routing fields typed on that routing relay's header/envelope, not copied as ordinary per-relay fields on every work relay.

Reasoning:
- The routing decision is authority-bearing. It needs its own accepted, stamped, lineage-visible record whose `FROM` is the router seat, matching m-1's "routing records ride submit()" contract.
- Altitude B is per dispatch: one routing decision can assign role+model for every seat in that dispatch. Repeating the model choice on every downstream work relay creates drift risk and makes the model look like a trust-bearing field.
- The v3 identity pillar says the seat is trust-bearing while the model behind a seat is payload/routing/benchmark bookkeeping, not a gate input: `extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md:33-35`.
- The downstream dispatch relay should consume or parent/reference the accepted routing relay. It may render a read-only projection for convenience, but the canonical routing decision remains the routing relay.

3. Does stamped store + router-seat `FROM` give forgery-robust dispatch authority?

Yes for seat identity and routing-decision authorship. m-1 stamps `FROM` from the seat connection and rejects unbound submit, while lane-authored payload `FROM` is ignored; the store is sole-writer and append-only. A lane cannot directly write the routing record or submit as another router seat under the stated m-1 invariants.

Boundary note: identity is not the whole authority model. m-1 owns who wrote the record; m-4/m-5 still own what that stamped seat may authorize. m-1 names this boundary explicitly at `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:101`.

4. Should routing reuse DESIGN_RECORD_KIND or define its own enum?

Use the DESIGN_RECORD_KIND pattern, not the literal design-domain enum values. Routing needs its own named enum set, for example `routing_decision`, `routing_deviation`, and later feedback/update kinds as m-4 designs them. The important reusable part is the FieldSpec/lineage pattern: a closed enum that distinguishes record kind for required-when, lineage, and authority gates.

5. Is the feedback handle shaped for v3.1 without recutting the record or gate?

Sufficient at the foundation level, with one carry-forward requirement for m-4 design: the dedicated routing relay must reserve nullable id_ref/object slots for policy/prior version and benchmark/outcome linkage now. m-2's FieldSpec has `id_ref`, `object`, `consumers`, and versioned schema machinery, and m-2 already names the benchmark/outcome-feedback handle. That is enough to avoid a foundation re-cut if m-4 makes the handle concrete during its own design.

## Per-Field Findings

- `role+model-per-dispatch`: sufficient. It belongs on the dedicated routing relay as assignments for the dispatch, then downstream work relays reference or project it read-only.
- `capability_prior`: sufficient. Treat as system/config-owned prior reference or snapshot, consumed by the router and later benchmark loop.
- `justified_deviation`: sufficient. Free-text reason should be required when the selection deviates from the prior floor.
- `record_kind`: sufficient if interpreted as a routing-owned enum using the DESIGN_RECORD_KIND shape; mis-owned if literal design enum values are reused.
- `benchmark/outcome-feedback handle`: sufficient as a reserved FieldSpec slot; m-4 must name exact IDs in its own design so v3.1 does not need a schema/gate re-cut.
- stamped store/write target: sufficient. m-1 gives routing records the same `submit()` path and stamped `FROM` identity as other records.

## Coordination

Coordination relays sent: none.

Reason: no blocking m-1 store/stamp gap and no blocking m-2 field-slot/ownership gap found. Q-C is resolved here for m-4.planner reconciliation; the expected follow-up is a reconciliation/fold into the joint lock, not a direct contract repair relay.

## Operator Judgment

Operator-judgment items: none for the foundation lock.

Residual risk: m-4's later domain design must make the routing enum values and feedback handle field IDs concrete. This is not a foundation blocker because m-2's FieldSpec and m-1's stamped store already express the needed shapes.

ACTIONS_GIT_REF: workspace file writes in non-git docs workspace: `master/relays/c1-consumer-review-m-4/AUDIT-implementer-20260629-162105.md`; INDEX row appended in `master/relays/INDEX.md`.
FINAL_GIT_STATUS_SHORT: unavailable - fatal: not a git repository (or any of the parent directories): .git
