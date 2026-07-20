## RECONCILE — c1 m-1/m-2 design integration + consumer-review plan, for review

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c1-design-reconcile
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — orchestrator decomposition review; the joint co-foundational lock (later) is the operator/VP gate
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — both foundational designs are complete and pair-approved; I have deep-read both full artifacts (not just the SITREPs). Before I dispatch the consumer review (3 dispatches staged + lint-clean, held for your go), looping you in on the integration + the plan. Read the two design docs for the full text:
- master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md
- master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md

Integration assessment (my read; verify or refute):
- The seam is airtight. m-1's system-filled PARENT (parent_picker over a conductor-derived candidate set, §6/§12) is exactly what upgrades m-2's lineage engine from confusion-robust to forgery-robust by construction (m-2 §11.1). Both docs name the same certification-null-reserved / courier-as-TCB deferral as one joint crypto-hardening class (m-1 §8, m-2 §7). The envelope/system-field contract converged in the COORD thread (B1 parent_picker, B2 recipient_picker TO/CC, B3 ROLE-system).
- The VP guardrail held: the pairs PROVED the hypotheses rather than assuming them. m-1 §4 proves I1/I2 by enumerating the attack surface (payload-lie / victim-credential / unbound-connection / replay) and closing each under DI-1..DI-4, with the strength claim stated conditionally (degrades to operator-attested if DI-2 unmet). m-2 §10 classifies all 62 v2.8.8 checks (~33 prose-only dissolve / ~16 form-validation / ~13 lineage); m-2's rev1 fixed a real lineage-gating-order blocker (two-state submitted→accepted so authority still blocks before delivery).
- Operator grill outcomes locked: transport Option A (minted per-seat credential = by-construction); bespoke FieldSpec registry (not JSON-Schema+x-, with cited reasoning); strict form-only submit; X- overflow; schema_version + zero-migrator framework; Merkle/DKIM deferred (reserved seams).
- Residual (both designs name it honestly): Step-1 DI-2 (the lane/credential-isolation realization) is a build-time "fork-2 infra call"; "by construction" is a design proof obligation that becomes an E2 fact only at PLAN. Acceptable for a design-lock because the claim is conditional with the honest fallback label.

Consumer-review plan (the 3 staged dispatches): each AUDIT/read-only, TO the consumer pair (both seats), CC you + both foundational planners; cites the exact consumer-field sections (m-1 §7, m-2 §12); asks a sufficiency verdict + the domain's open question; carries a loop-in instruction (a found gap is relayed DIRECTLY to m-1.planner or m-2.planner, CC me, to fix before lock).
- m-3 (Observation & Evidence): the observe/evidence field slots, the submit() pre-send gate hook placement, the probe-from-outside isolation (DI-2), and the m-1.implementer "observer-only hook" constraint (m-1 §13.4).
- m-4 (Routing & Policy): the routing-record fields + forgery-robust dispatch authority; resolves Q-C (model-per-dispatch: header field vs separate routing relay).
- m-6 (Human Surface & Scheduler): the human-gate monotonic floor + gate→email buckets + the inbox-as-projection substrate; resolves Q-B (the Owner Decision Brief field list).

Joint-lock plan: the three consumer reports return → I reconcile them + any m-1/m-2 coordination they trigger → the joint m-1↔m-2 co-foundational lock, routed through you for the full review → that is the c1 design-lock (cycle terminal; no PLAN, audit+design only).

Requested review, assuming I may be confidently wrong:
First, is my integration read sound — any hypothesis the pairs (or I) treated as proven that the proof does not actually carry? Pressure-test the I1/I2 proof (m-1 §4) and the dissolve/survive completeness (m-2 §10) specifically.
Second, is the consumer-review decomposition right — the correct consumer fields routed to the correct domain, the open questions (Q-B/Q-C) correctly assigned to the owning consumer? Is there a consumer field or cross-domain concern my dispatches do NOT direct the lenses at that they should — for example an m-5 expansion-slot/archetype dependency (m-2 §15 Q-D, slot_in) that needs an owner before lock?
Third, the loop-in mechanism — a found gap relayed directly to the foundational planner (CC me) — sound, or should gaps route through the orchestrator instead?
Fourth, anything from your five c1-reconcile carry-forwards that the consumer review must explicitly verify before the lock?

Return one verdict: approve / revise / reroute / reject-or-defer / human-decision-required, with cited changes and named required edits.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: reviewer returns a verdict relay (FROM master.orchestrator-reviewer, TO master.orchestrator-planner, CC operator) on c1-design-reconcile; on approve I dispatch the 3 consumer reviews.
