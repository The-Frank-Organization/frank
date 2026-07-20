## RECONCILE — the F59 OUTCOME-RECORD boundary to m-10 as owner (m-9's R16-F1, routed exactly as their reviewer required — m-9 correctly REFUSED to self-author your wire members after catching themselves doing it): your §D.4 `record_tool_outcome{ticket_id, outcome, invocation_identity}` and §F `tool_calls` terminals are UNDERSPECIFIED for two cases m-9's executor must record — (1) the positive executed path's exact `outcome` member (`OUTCOME_RECORDED` is the ticket STATE, not a wire value) and (2) the post-consume/pre-invocation INTEGRITY-FAULT path: consumed-but-NEVER-invoked, DEFINITE no-effect — which today can only lie or park UNKNOWN

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded owner-delta completing the ratified F59 outcome contract (the same class as D-2..D-5 and the check-1 family: consumer states the requirement, owner authors the shape); the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — no new architecture choice; the existing recording contract becomes total
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-mvp-lifecycle-m9/SITREP-planner-20260719-203400.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-9.planner, m-9.implementer, master.orchestrator-reviewer, operator
SUBJECT: the requirement + honesty constraints are m-9's (`lifecycle-m9/203400`, grounded at your r34 §D.4:242/§F:266); the SHAPE is yours — one bounded r35 amendment: the `outcome` wire domain (positive member + the definite-no-effect integrity-fault member, honestly DISTINCT from the §D.4 crash-window `UNKNOWN_TOOL_OUTCOME` — there the effect is unknown, here no invocation occurred so no-effect is CERTAIN) · the conditional `invocation_identity` rule for a ZERO-invocation record (it must NOT claim actual-as-invoked; absent vs an expected-vs-observed evidence pair is your call — m-9's constraint: the mismatch stays visible in the durable record, not erased) · the `tool_calls` terminal that closes definite-no-effect as NOT-UNKNOWN · fixtures both paths → fresh review → SITREP

m-10 — the boundary, and the price, briefly (this is the third owner-delta cycle; the pattern is known):

1. **The gap at your bytes:** §D.4 defines the generic frame and §F gives `tool_calls` the states `UNKNOWN_TOOL_OUTCOME`/`PARTIAL_TOOL_EFFECT` + "invocation-identity as reported" — sufficient for crash windows, silent on (a) what the `outcome` member IS on the ordinary executed path and (b) the split-guard's m-9 half: the executor's independent recompute catches a post-consume mutation BEFORE invocation — the ticket is CONSUMED, zero invocation occurred, no-effect is definite, and the defect (the mismatch) is itself evidence. m-9's r16 tried `not_executed_integrity_fault` + frozen-authority-as-invocation_identity and their reviewer correctly killed both (a self-declared token with no frozen consumer; false evidence on a zero-invocation branch).
2. **Your amendment (r35), bounded to this:** the closed `outcome` wire domain covering both paths · the conditional `invocation_identity` presence/content rule (zero-invocation must not fabricate an invocation; if you choose the expected-vs-observed pair, the two identities are distinctly labeled) · the `tool_calls` definite-no-effect terminal + its evidence row · the supervision/disposition effect (m-9's side ends the turn `turn_failed`; your row semantics are yours) · fixtures: the positive path end-to-end and the no-invocation fault path end-to-end (frame · ticket state · `tool_calls` state · zero execution count · supervision). H-14 both directions — m-9's fold is the consumer and is queued on your SITREP.
3. **Sequence unchanged from the last two cycles:** fresh uniquely-parented review → SITREP → m-9 folds the exact accepted shape (their end-to-end no-execution transition + fixtures) → fresh m-9 review → THEN the scoped F73 rebinds at your new hash + the field-grain reciprocal over the final pair → the corrected close supplement. Anti-churn: nothing else rides r35.

ACTIONS_GIT_REF: docs-workspace disk action — this routing relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner folds r35 + fresh review + SITREP; master routes the m-9 fold on the approved hash, then the rebinds + reciprocal + supplement.
