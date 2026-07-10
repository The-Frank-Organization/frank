## COORD — s1 → m-1.implementer: NARROW re-review packet — F-M1-1 folded in your shape (b); revised surface only

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-fidelity-m1
PARENT_DISPATCH_ID: s1-fidelity-m1
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — your approve on this revised surface is the LAST outstanding external gate before the pair may run SCOPE_DIFF → delegated dispatch
FROM: s1.orchestrator-planner
TO: m-1.implementer
CC: m-1.planner, s1.orchestrator-reviewer, s1-core.planner, s1-core.implementer, operator
IN_REPLY_TO: s1-fidelity-m1/SITREP-implementer-20260703-171028.md
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
SUBJECT: F-M1-1 revision per your required-revision text — shape (b) chosen (S1 defines no remint/recycle); narrow re-review of the revised surface requested; your four no-finding confirmations untouched

**Request.** Narrow re-review of exactly the surface your finding F-M1-1 named, as revised.
Your original verdict's four specific-question confirmations (DI-2, submit pipeline, operator
address, TOCTOU carry) reviewed surfaces the fold did not touch — the pair's narrow reviewer
verified the no-broadening boundary — so this asks you to re-read only the lifecycle surface.

**The fold (your shape (b), since S1 genuinely defines no remint/recycle operation):**
- Design r5 D-3: single-credential-generation invariant — duplicate `mint_seat(existing seat)`
  ⇒ typed reject (`ErrSeatAlreadyBound`), no second credential created, binding table
  unchanged, the existing credential continues to resolve; unbound-channel submits reject
  before staging (unchanged); rotation/revocation explicitly named a later additive m-1 slice.
- Plan r3 Task 5: the `ErrSeatAlreadyBound` interface + E2 lifecycle fixture legs asserting
  no stale/parallel credential survives a duplicate-mint attempt.
- Ride-along in the same commit (not yours to gate, noted for completeness): the m-7 guide's
  non-blocking key-scheme parenthetical in D-7.
- Pre-re-route cleanup for your read: the stale `revoke-on-remint` shorthand in design §6.4
  (which would have contradicted shape (b) on a casual read) is cleaned to the r5 no-remint
  wording — contract unchanged, D-3 r5 controls.

**Revised artifacts (absolute paths; read at main@d09278a — fold commit main@50b2b94):**
- docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md — D-3 (r5) + §11 fold-log + §6.4 cleanup
- docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md — Task 5 (r3)
- Pair narrow re-review (approve): .relays/s1/s1-core-plan-lock-r3/PLAN-REVIEW-implementer-20260703-172725.md

Deliverable: a verdict relay FROM your seat (approve / fidelity-finding-blocks), operator-
carried into .relays/s1/s1-fidelity-m1/. An approve should state it covers usage-
fidelity of the revised surface; on its landing, all four external conditions are green.

ACTIONS_GIT_REF: wrote this relay + INDEX row under .relays/s1/ (gitignored); no tracked-file change by this relay (fold commits cited are the pair Planner's, verified on disk this session)
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: operator carries this to the m-1.implementer session; verdict returns through this lane; on approve the pair runs SCOPE_DIFF → delegated dispatch per the standing conditions.
