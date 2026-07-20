## Team m-3 — Observation & Evidence: CONSUMER REVIEW of the m-1/m-2 foundations

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-consumer-review-m-3
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only consumer review; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-3.planner, m-3.implementer
CC: master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
BUNDLE_ID: c1-consumer-review-m-3
OWNER: m-3 (Observation & Evidence), as consumer lens

Phase scope — AUDIT (read-only consumer review). Review whether the m-1/m-2 foundational interfaces express what the m-3 Observation & Evidence domain needs to build on later — NOT full m-3 domain design (a later cycle), and not any edits or code. This consumer review is a hard prerequisite for the joint m-1↔m-2 co-foundational lock (the VP's lock condition).

Pair roles & research method: m-3.planner leads the consumer review + reconciles; m-3.implementer runs an INDEPENDENT pass, then reconcile. Deep external research is optional here (the task is mostly internal: does the foundation serve m-3?); use it only if a domain prior-art question arises. Independent paired review; the Planner does not spawn or direct the Implementer.

Context: m-1 (Trust & Identity) and m-2 (Forms & Determinism) are design-complete and pair-approved, HELD for the joint lock that needs your sign-off first. Transport is locked to Option A (minted per-seat credential = forgery-robust by construction). The observe-as-send gate is YOUR domain; this review confirms the foundation gives you the hooks + field slots to build it without churn.

Design docs to review:
- master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md — §5 (the submit/project/read/mint_seat API), §6 (system-field contract), §7 (consumer boundary contract: the m-3 bullet), §13.4 (the m-3-hook constraint).
- master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md — §3 (field-ownership), §4 (FieldSpec registry + the render/validate flow, esp. step 4b observe-as-send), §12 (consumer contract: the m-3 bullet).

Your consumer fields to validate (named in m-2 §12 + m-1 §7):
- ACTIONS_GIT_REF, FINAL_GIT_STATUS_SHORT — owner:system, fill_constraints: observed_value.
- EVIDENCE_TARGET — agent_enum_pick (intent); achieved_evidence — owner:system, observed; the target>achieved auto-flag as a system-computed *_RESULT.
- the per-phase done-predicate hook — a required_when / observe-gate input run at submit().
- executable-claim ref; egress/content-safety result (egress is post-v3.0 per m-2 §15 — confirm it is a clean reserved seam).
- the m-1 store-isolation boundary (DI-2) reused as your probe-from-outside-the-lane primitive; the submit() pre-send gate hook where your done-predicate runs BEFORE append ("no clean observation, no relay leaves").

Review questions (answer each; this is the substance):
1. Are the observe/evidence field slots (owner/type/fill_constraints) in the FieldSpec right for the observe-as-send gate + the E1–E4 evidence ladder?
2. Is the submit() pre-send gate hook sufficient and correctly placed (m-1 §5 step 3 / m-2 §4 step 4b) for your per-phase done-predicate to run before append — i.e. can a false "done" be rejected pre-send?
3. Does the m-1 store-isolation boundary (DI-2, shared with the stamper) give you the probe-from-outside-the-lane isolation you need (the conductor observes the lane workspace from outside it)?
4. m-1.implementer flagged (m-1 §13.4) that the m-3 hook must be observer/validator only — it cannot author or mutate FROM/ROLE/PARENT/relay_id or perform mailbox-delivery effects. Confirm that constraint is workable for your observe gate (it should be — observe reads ground truth + fills observed fields; it does not author identity).
5. Executable claims (the conductor runs the agent's check at the observe step) — are they expressible as a field/hook now, or a reserved seam?

Loop-in instruction: if you find a GAP — a field slot, hook, ownership, or isolation property your domain needs that the foundation does not express — relay the relevant foundational planner DIRECTLY (m-1.planner for store/stamp/isolation/hook gaps; m-2.planner for schema field-slot/ownership gaps), CC me, to coordinate a fix before the lock. Do not merely note it; loop them in so it resolves before lock. Contract constraint (per VP review c1-design-reconcile): any gap that changes the joint m-1/m-2 contract must still return through orchestrator reconciliation before the lock — direct coordination must not become an unreviewed side-lock.

Deliverable: a file-relay consumer-review report (independent per seat, then reconciled): a verdict (sufficient / gaps-found / mis-owned), per-field findings (confirmed-expressible / gap / mis-ownership), any coordination relays you sent to m-1/m-2, and operator-judgment items or none. E1-cited. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT (read-only; cwd is not a git repo).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
