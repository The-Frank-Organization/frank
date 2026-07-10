## Team m-1 — Trust & Identity: DESIGN REFINEMENT (rev2 — fold consumer-review findings)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c1-refine-m-1
PARENT_DISPATCH_ID: c1-consumer-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — design refinement; operator-judgment items (the DI-5 Step-1 honest-fallback) surface in the doc
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, m-2.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-1-trust-identity
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)

Phase scope — DESIGN refinement (rev2). Fold the consumer-review findings below into your design doc (rev2, same DESIGN_DOC_ID); re-run your pair design-review; report design-complete-rev2 for the joint lock. Still AUDIT + DESIGN-only — no PLAN/IMPL. Not in scope: code, branches, m-2 schema internals (sibling), the consumer domains' mechanisms.

Basis: the c1 consumer review (m-3/m-4/m-6) is reconciled and VP-approved (c1-consumer-reconcile, approve on the fold). Your design's SHAPE is endorsed — these are targeted refinements, not a redesign. Sources: c1-consumer-review-m-3 (the planner pass + COORD-planner-20260629-162500 DI-5 argument), c1-consumer-review-m-6 (planner pass G3), c1-consumer-reconcile (the orchestrator adjudication; I adjudicated both pairs toward their planners on the evidence).

Fold these three contract items + two non-blockers into rev2:

1. DI-5 — observe-integrity invariant (m-3 gap; VP-confirmed real and distinct from DI-2). Name an explicit deployment invariant alongside DI-1..DI-4: the conductor can read the lane's relevant ground-truth (workspace / git / process state) directly, from OUTSIDE the lane's trust boundary, so an observed-value field is the conductor's own reading and the lane cannot substitute or intercept it — observed evidence is conductor-observed, not lane-supplied, by construction. DI-2 (credential confidentiality between siblings) is necessary but orthogonal: it does not establish the read-vantage. Apply your §4 honest-fallback discipline — where a Step-1 realization cannot meet DI-5, the degraded label is "self-reported evidence / confusion-resistant," NOT observe-by-construction. DI-5 constrains the deferred fork-2 infra realization at PLAN exactly as DI-2 does.

2. Operator/human address-space (m-6 gap G3; VP-confirmed contract-level). Define the special operator/human address as a first-class entry in the system-owned minted address space: (i) it has a mailbox/projection (the human reads via project() / the operator-relay channel); (ii) it is a recipient_picker member (agents may address TO operator / CC operator); (iii) define the operator-FROM stamping path — the human authors via the operator-relay channel, NOT a minted lane credential, so operator is a delivery target plus a special stamped FROM, not a forgeable submit-lane. Reconcile with the charter, which lists operator as a named address.

3. Sharpening-D, m-1 half (m-4 routing; VP-corrected split). Confirm/add that an accepted routing relay (m-2's separate seat-stamped routing record) may appear in the conductor-derived parent_picker / reference candidate set for the dispatch relay it routes — AND that this reference does NOT make model-choice a trust-bearing gate input (the dispatch references the routing relay as bookkeeping; the model behind a seat stays payload, never a gate field — pillar :33). Make m-1's parent_picker candidate-set contract explicit on this point.

4. Non-blocking — restate §13.4 #4 ("observer-only") as a POSITIVE write-allowlist: the m-3 observe hook may write ONLY the closed m-3 observed/computed field set and emit a pass/fail veto; it may not write any system_only/identity/envelope field and has no delivery effect. Makes "observer-only" mechanically enforceable and squares §13.4 with m-2 §12 (where the hook writes its own observed fields = observation, not identity authorship).

5. Non-blocking — align your §5 step-numbering with m-2 §4's submitted then accepted ordering, so observe-before-append reads identically across both docs (an implementer must not read m-2 step-4 "append as submitted" as preceding observe).

Co-foundational re-affirm (hard boundary): DI-5, G3, and Sharpening-D touch the shared contract — G3's operator-address interacts with m-2's recipient_picker / address space; the routing-relay parent_picker candidate interacts with m-2's routing FieldSpec + lineage semantics. Coordinate with m-2.planner (the COORD-thread mechanism) and RE-AFFIRM the joint envelope/system-field contract in rev2. Neither domain re-locks in isolation; the orchestrator runs the joint co-foundational lock after both rev2s + the re-affirm, under the VP's full review.

Deliverable: the rev2 design doc (same DESIGN_DOC_ID; a rev2 fold-log + updated body); the m-1↔m-2 re-affirm (COORD relay); your pair design-review (Template I to m-1.implementer); and a design-complete-rev2 SITREP to me for the joint lock. Design Q&A inline; the rev2 doc + re-affirm are file-first.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
