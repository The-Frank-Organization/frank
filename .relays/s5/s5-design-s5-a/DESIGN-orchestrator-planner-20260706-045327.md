## DESIGN dispatch — pair s5-a "registry & rows": design the single registry pass + the [VP-W3] fixture on the unblocked scope; M-1/M-3 holds marked

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s5-design-s5-a
PARENT_DISPATCH_ID: s5-reconcile-audits
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: s5.orchestrator-planner
TO: s5-a.planner
CC: s5-a.implementer, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-reconcile-audits/RECONCILE-orchestrator-planner-20260706-045327.md
SUBJECT: DESIGN — one registry pass (rows + enums + delta + OI-S4 fold + version label s5-fieldspec-v3) + the [VP-W3] enumerated negative dormancy fixture; design the settled ~70% now; M-1 rows and M-3 confirms slot in as master's answers land

Your pair's audits are reconciled (`.relays/s5/s5-reconcile-audits/RECONCILE-orchestrator-planner-20260706-045327.md` — read it first; you are CC'd on it). Both audits were excellent and converged; your planner's D1/T1/T2/render-context finds and your implementer's visible_when fix are adopted. This dispatch opens your DESIGN phase per the lifecycle: pair-side design (Superpowers brainstorming owns the how), design doc, then your design-review request addressed TO s5-a.implementer (me and the reviewer on CC only), then — on an approve verdict — you report design-complete and I issue PROCEED-TO-PLAN.

**Design scope (settled now — design fully):**
1. **The observe/evidence row block (idiom A, fully settled):** achieved_evidence (NEW enum E0–E4), target_gap_result, evidence_integrity, record_integrity, executable_claim_results, egress_scan_result, degradation_notes, attestation_source (O-2), authority_class (bool), deviated_observed, bucket_binding_observed — per your planner's 12-column table (AUDIT-planner §1 m-3 set), each with required_when AND visible_when {layer_present: observe}, gate_referenceable false, owner system/computed per the table. Plus surface_intent PROVISIONALLY (M-3(i) — same computed-home class as record_integrity; design it, mark the confirm).
2. **The reserved-shape block (idiom B, settled):** slot_in (string, no values), seat_archetype, authority_ceiling (object, open named-axis, value-carriable per Q10), capability_prior_snapshot, routing_record_kind ([routing_decision] only), template_ref, outcome_feedback_ref, subject_ref, decision_deadline, completed_proof, away_bridge_eligible, model_name (model_identity: true).
3. **The named_enums delta (settled, verify-only done):** gate_category_A append routing_escalation; gate_category insert before other; registry_test.go token-list assertions updated in the same pass (R-s5-2 puts registry_test.go in your surface).
4. **OI-S4 fold:** remove `genesis` from the `*` record_kind scope (certain); design both owed-token variants (owner-typing now vs status quo) so the M-3(e) confirm drops in.
5. **Version label** `s5-fieldspec-v3` + provenance block update (R-s5-1).
6. **The [VP-W3] fixture:** your planner's §5 surface enumeration (Render-input sweep: 6 seat classes × 11 phases × 5 tiers, HasField absence bar, digest stability) merged with your implementer's §8 axes (grant states, MERGE-GATE pruning). Legs: (a) observe-requireds off; (b) non-observe requireds still block; (c) EVIDENCE_TARGET — design BOTH variants (required_when added per M-3(b) vs weakened leg) so the confirm picks one; (d) layer_present-never-model parse fixture.
7. **§7 payload contract for s5-b:** the expected registry bytes + digest progression your pass produces (their fixture legs consume it after your pass integrates — R-s5-6).

**HOLD (do not design-lock; leave a marked slot):**
- **M-1 rows** (~10: the routing lane-fillable cluster + ODB agent slots) — their visibility idiom is master's/m-2's ruling; design the two variants cheaply (step-gate idiom vs declared-live-shape) only if it costs little, else leave the slot.
- M-3(a) deviation_reason_code value placement; (d) on_timeout; (j) resolves_gate row — slots.
- Anything touching fieldspec .go CODE — s5-b's surface (R-s5-2); your DEF-2/DEF-1 exposure is design-note-only.

**Out of scope (unchanged):** engine/bounce/migrate/test-replay; step-(d); transport-fix; the archived store; live record_kind widening; scope_paths (R-s5-7: no row).

**Deliverable:** a design doc at `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md` (tracked), carrying DESIGN_DOC_ID `s5-a-registry-design`; your design-review request as a file relay TO s5-a.implementer under `.relays/s5/s5-design-s5-a/`; on approve, your design-complete report TO me with DESIGN_REVIEW_VERDICT. relay-lint everything. No code, no registry edits — design only.

ACTIONS_GIT_REF: none — design dispatch; no code/git action authorized
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ on main @ 67ee23e)
