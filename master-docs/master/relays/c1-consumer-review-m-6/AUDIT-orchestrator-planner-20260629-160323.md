## Team m-6 — Human Surface & Scheduler: CONSUMER REVIEW of the m-1/m-2 foundations

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-consumer-review-m-6
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only consumer review; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-6.planner, m-6.implementer
CC: master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
BUNDLE_ID: c1-consumer-review-m-6
OWNER: m-6 (Human Surface & Scheduler), as consumer lens

Phase scope — AUDIT (read-only consumer review). Review whether the m-1/m-2 foundational interfaces express what the m-6 Human Surface & Scheduler domain needs — NOT full m-6 domain design (a later cycle), and not any edits or code. This consumer review is a hard prerequisite for the joint m-1↔m-2 co-foundational lock.

Pair roles & research method: m-6.planner leads + reconciles; m-6.implementer runs an INDEPENDENT pass, then reconcile. Deep external research optional (mostly internal: does the foundation serve m-6?), though the Owner Decision Brief prior art (references/agent-scripts) is worth a look for Q-B. Independent paired review; the Planner does not spawn or direct the Implementer.

Context: m-1 (Trust & Identity) and m-2 (Forms & Determinism) are design-complete and pair-approved, HELD for the joint lock that needs your sign-off first. The human inbox = a projection of the addressing graph; gates project to operator email buckets. This review confirms the foundation expresses the human-gate fields + gives you the addressing/mailbox substrate to build the two-surface split (email governance vs meeting collaboration) + the scheduler on later.

Design docs to review:
- master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md — §5 (the API; project() + mailboxes), §6 (system-field contract; recipient_picker TO/CC + the system-owned address space), §7 (consumer contract: the m-6 bullet — addressing graph + inbox = projection).
- master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md — §3 (field-ownership; the HUMAN_GATE monotonic floor), §12 (consumer contract: the m-6 bullet), §15 Q-B (the open Owner-Decision-Brief question for you to resolve).

Your consumer fields to validate (named in m-2 §12 + m-1 §7):
- HUMAN_GATE_REQUIRED — owner:system, fill_constraints: monotonic (the agent may only RAISE the floor, never lower it); human_gate_reason.
- gate→email bucket — system-derived from TO/CC + verdict (the A reserved-to-human / B orchestrator-absorbed / C CC-only-FYI / D lint-bounce policy projection).
- TO/CC operator; human-decision-required (auto-escalate-TO-operator).
- the Owner Decision Brief 7-field sub-schema (agent-pick recommendation + enumerated choices; port from agent-scripts).
- park/wake/summon-urgency (scheduler fields).
- the m-1 addressing graph (TO/CC + the system-owned minted seat-address space); inbox = projection of the addressing graph; the mailbox / project() substrate backs your operator inbox + gate→email buckets.

Review questions (answer each):
1. Do the human-gate field slots express the monotonic floor (agent may only RAISE), human-decision-required auto-escalation, and the gate→email bucket projection?
2. RESOLVE Q-B (m-2 §15), your call as the human-surface owner: the Owner Decision Brief sub-schema — port the agent-scripts 7 fields verbatim (plain-language change / why-now / completed-proof / tradeoffs / recommendation / exact-choices / …), or trim/extend? Give the field list + reasoning.
3. Does the addressing graph + inbox-as-projection + the mailbox / project() substrate (m-1 §5/§7) give you what you need to build the operator inbox + the two-surface split (email governance lane vs meeting/attach-to-session collaboration lane)?
4. The gate→email bucket is "system-derived from TO/CC + verdict" — is that derivation sufficient for your A/B/C/D bucket policy, or do you need an explicit bucket field / additional inputs?
5. park/wake/summon-urgency — are these expressible as fields/hooks now, or a clean reserved seam for the m-6 scheduler design (later cycle)?

Loop-in instruction: if you find a GAP — a human-gate field, bucket input, addressing/mailbox property, or scheduler hook your domain needs that the foundation does not express — relay the relevant foundational planner DIRECTLY (m-1.planner for addressing/mailbox/project gaps; m-2.planner for schema field-slot/ownership gaps), CC me, to coordinate a fix before the lock. Loop them in; do not merely note it. Contract constraint (per VP review c1-design-reconcile): any gap that changes the joint m-1/m-2 contract must still return through orchestrator reconciliation before the lock — direct coordination must not become an unreviewed side-lock.

Deliverable: a file-relay consumer-review report (independent per seat, then reconciled): a verdict (sufficient / gaps-found / mis-owned), per-field findings, the resolved Q-B answer (the ODB field list) + reasoning, any coordination relays sent to m-1/m-2, and operator-judgment items or none. E1-cited. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
