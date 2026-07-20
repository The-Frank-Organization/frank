## Team m-4 — Routing & Policy: CONSUMER REVIEW of the m-1/m-2 foundations

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-consumer-review-m-4
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only consumer review; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-4.planner, m-4.implementer
CC: master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
BUNDLE_ID: c1-consumer-review-m-4
OWNER: m-4 (Routing & Policy), as consumer lens

Phase scope — AUDIT (read-only consumer review). Review whether the m-1/m-2 foundational interfaces express what the m-4 Routing & Policy domain needs — NOT full m-4 domain design (a later cycle), and not any edits or code. This consumer review is a hard prerequisite for the joint m-1↔m-2 co-foundational lock.

Pair roles & research method: m-4.planner leads + reconciles; m-4.implementer runs an INDEPENDENT pass, then reconcile. Deep external research optional (mostly internal: does the foundation serve m-4?). Independent paired review; the Planner does not spawn or direct the Implementer.

Context: m-1 (Trust & Identity) and m-2 (Forms & Determinism) are design-complete and pair-approved, HELD for the joint lock that needs your sign-off first. The routing decision is a seat-stamped on-disk record that rides submit() like any relay; this review confirms the foundation expresses the routing-record fields + gives forgery-robust dispatch authority.

Design docs to review:
- master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md — §5 (the API), §6 (system-field contract), §7 (consumer contract: the m-4 bullet — stamped store as routing-record write target).
- master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md — §3 (field-ownership), §4 (FieldSpec registry), §12 (consumer contract: the m-4 bullet), §15 Q-C (the open routing question for you to resolve).

Your consumer fields to validate (named in m-2 §12 + m-1 §7):
- role+model-per-dispatch (planner-emitted, altitude-B — the locked routing altitude).
- capability_prior (system/config — the GPT→math / Gemini→science / Opus→SWE floor).
- justified_deviation (free-text reason + a record_kind enum reusing the DESIGN_RECORD_KIND shape).
- benchmark / outcome-feedback handle (reserved for the v3.1 feedback loop).
- the m-1 stamped store as the write target for routing records; a routing record's FROM = the router seat's stamped identity, so dispatch authority is forgery-robust.

Review questions (answer each):
1. Do the routing-record field slots express altitude-B (role + model per dispatch) + capability priors + justified deviation + the v3.1 benchmark/feedback handle, all as typed FieldSpec fields?
2. RESOLVE Q-C (m-2 §15), your call as the routing owner: does model-per-dispatch belong in the routing-record HEADER (a per-relay field) or in a separate routing RELAY (a per-dispatch record)? This decides the FieldSpec shape — give the answer + the reasoning.
3. Does "the stamped store as write target + routing-record FROM = the router seat's stamped identity" give you forgery-robust dispatch authority (no lane can forge a routing decision)?
4. The record_kind reuses the DESIGN_RECORD_KIND enum shape — is that the right pattern for routing records, or does routing need its own record-kind enum?
5. Is the benchmark / outcome-feedback handle shaped so the v3.1 outcome-feedback loop (the open, non-gradient analog of Fugu's reward) bolts on WITHOUT re-cutting the record or the gate (the "shape record+policy for feedback day one" requirement)?

Loop-in instruction: if you find a GAP — a routing field, record shape, or store/stamp property your domain needs that the foundation does not express — relay the relevant foundational planner DIRECTLY (m-1.planner for store/stamp gaps; m-2.planner for schema field-slot/ownership gaps), CC me, to coordinate a fix before the lock. Loop them in; do not merely note it. Contract constraint (per VP review c1-design-reconcile): any gap that changes the joint m-1/m-2 contract must still return through orchestrator reconciliation before the lock — direct coordination must not become an unreviewed side-lock.

Deliverable: a file-relay consumer-review report (independent per seat, then reconciled): a verdict (sufficient / gaps-found / mis-owned), per-field findings, the resolved Q-C answer + reasoning, any coordination relays sent to m-1/m-2, and operator-judgment items or none. E1-cited. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
