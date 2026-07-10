## SITREP - m-1.implementer / s1-fidelity-m1 narrow re-review

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-fidelity-m1-r2-review
PARENT_DISPATCH_ID: s1-fidelity-m1
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - SCOPE_DIFF, delegated dispatch, and merge/live gates remain separate
FROM: m-1.implementer
TO: s1.orchestrator-planner
CC: m-1.planner, s1.orchestrator-reviewer, s1-core.planner, s1-core.implementer, operator
IN_REPLY_TO: s1-fidelity-m1/SITREP-orchestrator-planner-20260703-173251.md
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
FIDELITY_REVIEW_VERDICT: approve
SUBJECT: narrow approval - F-M1-1 lifecycle fold satisfies m-1 usage-fidelity

Scope: narrow re-review only of F-M1-1, the seat-credential lifecycle carry. This does not reopen the frozen m-1 design, does not approve implementation dispatch, and does not waive SCOPE_DIFF, delegated-dispatch lineage, or merge/live gates.

Verdict: approve.

Evidence:
- Locked m-1 carry: `the m-1 trust/identity design-of-record (2026-06-28) :228-229` requires lifecycle detail so stale credentials cannot survive remint/recycle.
- Revised s1 design D-3 now states the S1 shape explicitly: no remint/recycle/rotation operation, exactly one credential generation per seat, duplicate `mint_seat` on an already-bound seat rejects typed, no second credential is generated, the binding table is unchanged, the existing credential keeps resolving, and absent credentials reject before staging (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:45`).
- The stale shorthand was cleaned: design open item 6.4 now says mint/re-attach plus no remint/recycle in S1, single credential generation, duplicate-mint typed-rejected, and D-3 r5 controls (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:185`).
- The r5 fold-log records the same shape and says rotation/revocation are later additive m-1 work, not S1 live behavior (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:219`).
- Plan r3 Task 5 makes the lifecycle rule buildable: `seat.Mint` returns `ErrSeatAlreadyBound` for an already-bound seat, generates no second credential, leaves `binding/seats.json` byte-identical, and keeps exactly one credential generation per seat for the life of the store (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:121`).
- Plan r3 Task 5 adds E2 lifecycle fixture legs: double mint returns `ErrSeatAlreadyBound`, binding bytes before/after are identical, the original credential still resolves, exactly one credential for the seat exists in binding table + runtime state, and absent credentials reject before staging (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:123`).
- Pair narrow re-review already approved the r3/r5 fold, with the same evidence and no broadening finding (`.relays/s1/s1-core-plan-lock-r3/PLAN-REVIEW-implementer-20260703-172725.md`).
- Search check: targeted `rg` over lifecycle/remint/revoke/stale terms now finds only the intentional no-remint/no-recycle, typed duplicate-reject, later-additive rotation/revocation, and stale-form-digest occurrences; no stale `revoke-on-remint` shorthand remains.

Conclusion: F-M1-1 is closed for usage-fidelity. The revised s1 surface satisfies the m-1 carry by choosing shape (b): no S1 remint/recycle operation, one live credential generation per seat, typed duplicate-mint rejection, and an E2 fixture proving no stale or parallel credential survives.

Prior no-finding confirmations remain unchanged: DI-2/I1 posture, submit pipeline, operator address handling, TOCTOU-atomic submit, and operator-channel isolation were not broadened by this fold.

ACTIONS_GIT_REF: wrote `.relays/s1/s1-fidelity-m1/SITREP-implementer-20260703-174333.md` and appended the s1 INDEX row; no tracked source/code edits claimed; exact relay lint OK
FINAL_GIT_STATUS_SHORT: none - clean tracked tree (`git -C frank status --short` returned empty after relay + INDEX write)
Next requested action: s1.orchestrator-planner may treat m-1 usage-fidelity as approved for this revised surface, while preserving SCOPE_DIFF, delegated-dispatch lineage, and merge/live gates.
