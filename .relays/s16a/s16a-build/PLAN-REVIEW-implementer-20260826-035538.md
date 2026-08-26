## PLAN-REVIEW successor - must-revise remains after master WP5 ruling; prior findings narrowed and corrected forward

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-15
PARENT_DISPATCH_ID: s16a-build-20
RUN_ID: s16a
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - r21 can bind the master ruling and close the remaining mechanics; operator merge authority remains downstream
PLAN_LOCK_ID: s16a-build-20
IN_REPLY_TO: s16a-wp5/RECONCILE-orchestrator-planner-20260826-035119.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-10.planner, m-8.planner
SUBJECT: successor review after master ruling - bind corrected registrations/export/ready facts; still revise merged-object r10, fold authority, lease safety, and production-risk handoff
PLAN_REVIEW_VERDICT: must-revise

This is the corrective-forward successor to `s16a-build/PLAN-REVIEW-implementer-20260826-035308.md`, prompted by master ruling `s16a-wp5/RECONCILE-orchestrator-planner-20260826-035119.md` at SHA-256 `1336560045530e6dd30d508054809ad0fd86bc6f07c75e42dc41456b5729e508`. The reviewed plan remains exact r20 SHA-256 `3625980402d6ee2d6cecdb007887abd22b335b4f16f02b3134a118b7e9bc0fd6`.

## Corrections to the prior review

- A14 (`67acb760...`) and the B10 worker-r12 carrier (`63f5c49d...` section 7.1a) are already discharged. Only the D01 clause remains due, and master routed it; it must land before the brief. Prior F3 is narrowed accordingly.
- F.7.2 is satisfied by construction: the engine root `frank/.relays/s16a` plus the closing workspace commit is the export of record; no duplicate export directory is owed.
- Master assigns the PR ready-for-review flip to the MERGE-GATE brief. The prior assertion that no ready step was assigned is withdrawn; r21 must carry this exact positive assignment and executor routing.
- Master confirms the high-level sequence. This review does not reopen the ruled registration/export/ready decisions.

## Remaining blocking findings

### F1 - r20 does not bind the new ruling

R20 predates the ruling and still says all three registration gaps may await a later blocking decision, calls the export form unresolved, and states only a negative ready condition. Those bytes are now stale. Required r21: bind the exact two discharged registration hashes, D01-before-brief gate, satisfied-by-construction export root plus closing hash, and ready-at-brief rule.

### F2 - merged-object r10 identity remains unstated

Master confirms r10 before the MERGE-GATE, while the ratified plan requires the r10 census at the merged object (`master/STEP-3-T4-S16A-PLAN.md:31,44,60`). The ruling does not state how a pre-merge record satisfies that literal. Required r21: distinguish pre-merge close-candidate evidence from final post-merge r10/close, or explicitly bind a master interpretation that proves reviewed-head tree identity to the merged object and defers the final close claim until that identity exists. Do not silently call a PR head merged.

### F3 - review findings still lack phase-correct fold authority

R20 says m-7 findings fold before restack but issues the IMPL token only on a clean return, and pre-authorizes future end-review findings. Required r21: each non-clean return gets a later addressed REVIEW-FOLD relay and pre-edit all-in FOLD_SCOPE, followed by rerun and reviewer disposition. A clean return permits only restack/rerun. No earlier IMPL token authorizes unknown future findings.

### F4 - restack rewrite still needs an exact lease

Master expects a clean restack because the 46 behind commits are governance-only, but per-commit replay remains unproven and r20 still says generic force-push. Required r21/dispatch: pin verified main and old remote branch SHAs; use exact `--force-with-lease`; STOP on lease failure; report every replay conflict before resolution. Preserve the draft and no-CI/no-merge fences.

### F5 - production-risk handoff completeness remains absent

Required r21: add the completed escalation scan, boundary contract, operator-judgment inventory, writers/readers/target, both-seat E2, and post-merge identity proof. The master ruling closes decisions; it does not retroactively add missing r20 fields.

## Evidence and incident disclosure

- Current branch and remote remain clean/equal at `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; PR #1 remains OPEN and DRAFT.
- Exact-file lint and engine projection of review 14 passed; the inherited INDEX divergence remains separate.
- During review-14 drafting, the relay write hook rejected three workspace-root temporary payload paths after writing them. A concurrent master checkpoint captured those untracked files. Cleanup then appeared as tracked deletions, so this seat immediately restored their checkpointed bytes and leaves their disposition to master: `PLAN-REVIEW-draft-s16a-20260826-034754.md`, `PLAN-REVIEW-implementer-20260826-035002.md`, `review_payload.txt`. No foreign history was rewritten.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - no credential/runtime authorization byte in this review
- migration/backfill/destructive-write/canonical-data-repair: yes - planned remote history replacement requires lease protection
- money/inventory/orders/planning/accounting/trust-critical-state: yes - close and merge authority are trust-critical
- AI-or-automation-acts-downstream: no - no downstream automation authorized
- worker/scheduler/queue/retry/async-side-effect: no - queue enumeration is evidence-only
- cross-repo/service-contract/generated-schema/shared-API-event: yes - PR, engine root, master ledger, owner registrations, and lane meet
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: no - both-seat E2 and merged-object identity required
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - merged-object identity and fold/lease mechanics remain open
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

No implementation authority or action: no source, test, branch, commit, push, PR, merge, CI/CD, store, credential, runtime, contract, ledger, or owner-design byte changed.

Next requested action: file r21 parented to this successor review and bind the master ruling plus F1-F5. No restack or fold precedes its approval and a separate addressed token.

ACTIONS_GIT_REF: no source/test/branch/PR action; implementation worktree clean at local/remote-equal `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; review artifact submission only; daemon/client fingerprint `ddd5f9f5c84e3c51dbe06ce7d1d7174604b5f0c17524c155091f3c6738bcc6c0`
FINAL_GIT_STATUS_SHORT: none - implementation worktree clean at local/remote-equal `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; governing checkout before filing has only daemon INDEX, operator CHECKPOINTS, and daemon-rendered current relay dirt
