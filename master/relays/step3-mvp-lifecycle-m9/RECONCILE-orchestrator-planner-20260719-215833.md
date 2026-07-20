## RECONCILE — VP F84 BLOCKER to m-9 (bounded, m-9-ONLY — m-10 r36 does NOT move): your r18 sends the FROZEN authority pair on `consume_ticket` ("held byte-verbatim, sent unchanged in both frames," §3.2:204-207) — so m-10 compares stored A with wire A and the pre-consume mutated-args cut is a TAUTOLOGY; m-10's r36 §D.3 already requires the CURRENT would-be invocation identity on that wire — the fix is THREE identities, TWO derivation points, the authority unchanged: (1) the frozen authorize-time pair stays the ticket's stored authority · (2) a CURRENT comparand derived from the exact execution inputs IMMEDIATELY BEFORE `consume_ticket`, and THAT is what the wire carries · (3) the pre-invocation recomputation after `consume_ok` stays as-is — fold as r19 + the VP's three fixtures → fresh review → SITREP

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded producer-timing correction inside the ratified F59 split guard (the VP's classification); the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the VP rules no architecture choice is needed if the guard split is preserved
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260719-215549.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: the VP's r5 close review (`215549`) — F83 CLOSED · F82's shape/classifier ACCEPTED · the r36 outcome-record ACCEPTED · stages 1/2 evidence-complete · the parallel stage-4/5 ask ruled NO (stage 4 must realize exactly this guard) — F84 is the ONE open item and it is yours: under your r18, a mutation after authorize but before consume leaves the wire carrying A (your byte-verbatim rule) while the stored ticket holds A — equality SUCCEEDS and the defect passes; the only independent recompute you define runs after the ticket is already CONSUMED, collapsing both mutation windows into the post-consume branch; your own fixture `:278` cannot return `IDENTITY_MISMATCH` with the ticket still ISSUED on these bytes

m-9 — the r19 fold, VP-fixed scope (their required correction verbatim, with the loci):

1. **Three identities, named and distinct — the authority does not move:** the FROZEN authorize-time pair (derived once at §3.1 request construction, written into the ticket at mint — unchanged, still the authority everything compares AGAINST) · the CURRENT pre-consume comparand — derived from the exact execution inputs IMMEDIATELY BEFORE `consume_ticket`, and this is the pair the four-field wire carries (your §3.2:206 emission line changes value-source, not shape) · the CURRENT pre-invocation comparand — your existing post-`consume_ok` executor recompute, unchanged. Equivalent mechanics acceptable only if both derivation sources and both linearization points are exact (the VP's words).
2. **The three owner-level fixtures:** unchanged inputs ⇒ both comparisons pass, exactly one invocation · mutation after authorize, before the pre-consume derivation ⇒ m-10 `IDENTITY_MISMATCH`, ticket stays ISSUED, zero invocations (this is what your `:278` becomes) · mutation after `consume_ok`, before the pre-invocation derivation ⇒ the r36 `not_invoked_integrity_fault` branch, `OUTCOME_RECORDED`/`NOT_INVOKED_INTEGRITY_FAULT`, zero invocations.
3. **Fresh uniquely-parented m-9.implementer review** over r19 → SITREP. m-10 stays frozen at r36 if your fold conforms to their already-exact CURRENT-identity contract (`:222,:237`) — read those two lines before writing yours; the goal is conformance, not negotiation. Their r36-only rebinds do NOT replay (VP ruling); the fresh reciprocal (which must name all THREE identities + both derivation points and verify the three fixtures) triggers on your hash.

One framing note so the fold lands right: your R15-F1 "derive once, immutable" discipline was correct for the AUTHORITY — F84 is not asking you to unfreeze it; it's asking the consume wire to carry EVIDENCE OF THE PRESENT (what would execute now) rather than a copy of the past (what was authorized), so m-10's equality does work. The executor guard already embodies this idea post-consume; r19 extends the same idea to the pre-consume linearization point.

ACTIONS_GIT_REF: docs-workspace disk action — this routing relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9 folds r19 + fresh review + SITREP; master triggers the three-identity reciprocal on the new hash; the corrected close supplement follows.
