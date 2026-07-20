## Team m-2 — Forms & Determinism: DESIGN REFINEMENT (rev2 — fold consumer-review findings)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c1-refine-m-2
PARENT_DISPATCH_ID: c1-consumer-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — design refinement; operator-judgment items (ODB on_timeout default; gate_category enum membership) surface in the doc
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-2.implementer, m-1.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
BUNDLE_ID: m-2-forms-determinism
OWNER: m-2 (Forms & Determinism)

Phase scope — DESIGN refinement (rev2). Fold the consumer-review findings below into your design doc (rev2, same DESIGN_DOC_ID); re-run your pair design-review; report design-complete-rev2 for the joint lock. Still AUDIT + DESIGN-only — no PLAN/IMPL. Not in scope: code, branches, m-1 store/identity internals (sibling), the consumer domains' mechanisms (m-2 declares the SLOTS; the sibling owns the filling/consuming mechanism).

Basis: the c1 consumer review (m-3/m-4/m-6) is reconciled and VP-approved (c1-consumer-reconcile, approve on the fold). Your design's SHAPE is endorsed — these are field-schema refinements, not a redesign. Sources: c1-consumer-review-m-6 (the planner pass G1/G2 + the Q-B ODB resolution), c1-consumer-review-m-4 (the RECONCILE-planner Sharpenings A-D + the Q-C separate-routing-relay resolution), c1-consumer-reconcile (the orchestrator adjudication).

Fold these into rev2:

1. G1 — HUMAN_GATE_REQUIRED render-affordance (m-6 gap). The field is owner:system + monotonic, but §3's render keys affordance on owner and owner:system renders "not shown" — so the agent cannot RAISE it as written. Fix mirrors the pickers: make HUMAN_GATE_REQUIRED a HYBRID — the system sets the floor; the agent selects within [floor, MAX]; the courier validates the pick is monotonically at or above the floor. Add monotonic to §3's hybrid-render clause; stop tagging it bare owner:system. The floor composes as the MAX across all raisers (system baseline, agent, m-5 archetype-ceiling, m-4 routing-raise) — a lower write never wins.

2. G2 — gate→email bucket inputs (m-6 gap, answers Q4). "system-derived from TO/CC + verdict" cannot mechanically split the human-only bucket A from the orchestrator-absorbed bucket B. Make human_gate_reason a CLOSED ENUM keyed to the protocol operator-judgment categories (the human-only category A covers product-semantics, irreversible writes, and merge decisions; category B covers routing/sequencing/scope) — canonical-iff-consumed: the bucket projection is a mechanical consumer, so its driver must be enum, not free text. Also expose the two-state delivery_state (accepted or bounced) + the failing_edge as readable computed slots, so the lint-bounce (bucket D) email can state why. m-2 declares these slots; m-6 owns the bucket projection mechanism.

3. Q-B — Owner Decision Brief sub-schema (m-6 resolution; m-6 is the human-surface owner, you encode the slots). Encode the ~10-field ODB sub-schema per the m-6.planner resolution: port the agent-scripts 7 (plain_language_change, why_now, completed_proof, tradeoffs_risks, recommendation, choices, subject_ref), retype three to the v3 substrate (subject_ref = id_ref into the relay store; completed_proof = evidence_ref pulled from m-3 observed evidence, never agent free-text; choices = row_array the operator reply selects within); extend with gate_category (= the same closed enum as the G2 human_gate_reason fix — one category set serves bucket-routing and the brief); reserve two scheduler fields null (decision_deadline, on_timeout). Provenance (FROM/PARENT/DISPATCH_ID) auto-projects from the certified envelope — not a manual ODB field.

4. Routing-relay FieldSpec — Sharpenings A/B/C + D-schema-half (m-4 routing; Q-C resolved = a SEPARATE seat-stamped routing relay, not a dispatch header). Encode the routing-relay record fields: routing_assignments (type row_array, owner seat_scoped_enum to planner/orch-planner only — altitude-B fill-time authority, each row = target seat/role + selected model from a closed model enum); capability_prior_snapshot (owner system, fill_constraints computed_result/observed_value — the snapshot of the prior in effect at decision time, replay-complete); justified_deviation (owner free_text, required_when selected_model is off the prior floor); routing_record_kind (own named-enum, DESIGN_RECORD_KIND SHAPE not values); outcome_feedback_ref (owner system, type id_ref, null-reserved for v3.1, lineage_role none). D-schema-half: the routing relay carries accepted/deliverable lineage semantics and its record-kind identity; m-1 owns the parent_picker candidate-set half (the dispatch references the accepted routing relay) — coordinate the join with m-1.planner. Keep model OUT of the work relay's authority header (pillar :33 — model is payload, never a gate input).

5. slot_in reservation (already dispositioned, VP-approved). In the §5 required-when predicate vocabulary, keep the slot_in atom SHAPE but define NO concrete archetype/slot enum values for Step 1, and no required-when predicate may branch on a concrete slot; m-5 (Workflows & Archetypes) owns the tag-space + concrete slot semantics in a later cycle. Note it as a reserved atom (same pattern as certification/Merkle).

Co-foundational re-affirm (hard boundary): G2's address interaction, the ODB completed_proof = m-3 evidence_ref dependency, and the routing-relay parent_picker join all touch the shared contract. Coordinate with m-1.planner (the COORD-thread mechanism) and RE-AFFIRM the joint envelope/system-field contract in rev2. Neither domain re-locks in isolation; the orchestrator runs the joint co-foundational lock after both rev2s + the re-affirm, under the VP's full review.

Deliverable: the rev2 design doc (same DESIGN_DOC_ID; a rev2 fold-log + updated body — incl. the updated FieldSpec entries + the routing-relay record shape); the m-1↔m-2 re-affirm (COORD relay); your pair design-review (Template I to m-2.implementer); and a design-complete-rev2 SITREP to me for the joint lock. Design Q&A inline; the rev2 doc + re-affirm are file-first.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
