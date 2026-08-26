## PLAN-REVIEW - must-revise WP5 PLAN r20: close order, fold authority, registrations, restack, and draft-to-merge mechanics are not executable

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-14
PARENT_DISPATCH_ID: s16a-build-20
RUN_ID: s16a
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the Planner can issue a bounded successor or route the object and registration questions to master; the operator merge gate remains downstream
PLAN_LOCK_ID: s16a-build-20
IN_REPLY_TO: s16a-build/PLAN-planner-20260826-034346.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-10.planner, m-8.planner
SUBJECT: must-revise r20 - split pre-merge evidence from merged-object r10, give findings valid fold gates, require registrations, lease-protect restack, authorize ready
PLAN_REVIEW_VERDICT: must-revise

`PLAN_REVIEW_VERDICT: must-revise` for exact r20 SHA-256 `3625980402d6ee2d6cecdb007887abd22b335b4f16f02b3134a118b7e9bc0fd6` at `s16a-build/PLAN-planner-20260826-034346.md`.

R20 is FROM the pair Planner, TO this seat, plan-only, parented to filed close carriage `s16a-wp34-close`, exact-file lint-clean, and contains no implementation or merge token. The all-green launch is corroborated at pushed `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; these findings concern close mechanics.

## Blocking findings

### F1 - r10 precedes its required merged object

Rung 5 has master consume r10 evidence before export and MERGE-GATE (`:23-25`), while acceptance names only an analyzed object (`:31`). The governing plan requires the r10 census at the merged object at `master/STEP-3-T4-S16A-PLAN.md:31,44,60`. A PR head may later share a tree but is not the merged object and merge identity is not pinned.

Required r21: use a pre-merge close-candidate set for the MERGE-GATE brief, then after authorized merge prove merged commit/tree identity, bind or rerun the census there, and only then let master file final r10 and declare close. A pre-merge r10 requires an explicit upstream amendment.

### F2 - fold authority is circular and phase-invalid

Rung 1 says m-7 findings fold before restack (`:19`), but `:34` issues the IMPL token only when that return is clean and pre-authorizes future rung-4 findings. An earlier IMPL token cannot replace later REVIEW-FOLD authority and pre-edit FOLD_SCOPE.

Required r21: a clean m-7 return permits only restack/rerun. A finding gets an addressed REVIEW-FOLD relay and all-in FOLD_SCOPE before edits, then rerun and m-7 re-review. Rung-4 blockers/must-haves get their own later REVIEW-FOLD gate, scope, verification, and reviewer disposition.

### F3 - three due registrations are optionalized

Rung 5 allows missing A14, B10-carrier, or D01-clause instruments while master later decides whether the gap blocks (`:23`). The close carriage says all three are due and will be cited (`s16a-wp34/SITREP-planner-20260826-034345.md:1`); D01's ruling is temporary until registration. This contradicts no-carry-substitute close.

Required r21: make all three owner-folded exact-hash instruments prerequisites to the close candidate, or cite an upstream amendment changing the obligation. Carry affirmative discharge evidence for the original five R-S16A residuals; naming deferrals is not discharge.

### F4 - restack and draft-to-merge mechanics are unsafe/incomplete

Rung 2 directs generic force-push (`:20`). Remote main is `8659164ddfbfd36fbe7d98baa3ba4fdef05877f7`, branch is `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`, divergence 46 behind / 48 ahead from launch base `ff1193d7`. Aggregate merge-tree has no conflict markers but does not prove per-commit replay.

Rung 7 forbids ready before the close set but never authorizes or assigns ready afterward (`:25`). PR #1 remains OPEN and DRAFT, so merge acceptance is unreachable.

Required r21: at dispatch pin verified main and expected old remote-head SHAs; use exact `--force-with-lease` and STOP on lease failure; report each replay conflict before resolution. Make MERGE-GATE explicitly authorize/address ready-for-review plus selected PR merge, or define a separate addressed ready gate. No CI/CD or merge before operator grant.

### F5 - production-risk handoff fields are absent

R20 lacks an escalation scan, boundary contract, and operator-judgment inventory despite remote history replacement, owner registrations, durable close evidence, and human merge. Exact-file lint does not prove semantic completeness.

Required r21: include the production-risk scan without downgrade; name branch/PR and close artifacts as writers, master/operator/r10 as consumers, merged main as target, and E2 plus post-merge identity as proof. Unresolved master rulings remain gates.

## Preserved portions

Preserve m-7 review before restack; verified current main; full suite, vet, 64/0/64 census, loud sentinel, bijection, gofmt, diff-check at both seats; changed-reviewed-bytes review; no quiet fixes; exit/D03/dependency/contract/E3/release/live-composition fences; queue enumeration; master-owned F.7.2 export; one operator-gated merge with one executor in TO.

Boundary contract disposition: must-revise. Intended writers are restacked branch/PR, close candidate, export, merge, and final r10; readers are reviewer, master, operator, and merged-main consumers. R20 does not join them non-circularly or authorize ready positively.

## Evidence

- E1: exact r20 read/SHA; governing WP5 and PR directive compared.
- E2/read-only Git: local/remote branch `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; remote main `8659164ddfbfd36fbe7d98baa3ba4fdef05877f7`; divergence 46/48; aggregate merge-tree no conflict markers. Feasibility, not rebase proof.
- PR #1 OPEN and DRAFT at `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; no ready, checks, merge, or PR mutation.
- No source tests ran because this is plan review and no implementation byte moved.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - no credential/runtime authorization byte in this review
- migration/backfill/destructive-write/canonical-data-repair: yes - planned history replacement needs lease and stop conditions
- money/inventory/orders/planning/accounting/trust-critical-state: yes - governed close and merge authority are trust-critical
- AI-or-automation-acts-downstream: no - no downstream automation authorized
- worker/scheduler/queue/retry/async-side-effect: no - queue enumeration is evidence-only
- cross-repo/service-contract/generated-schema/shared-API-event: yes - PR, master ledger, owner registrations, and lane meet
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: no - both-seat E2 and merged-object identity required
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - r10 timing, registrations, and ready authority unresolved
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

No implementation action: no source, test, branch, commit, push, PR, merge, store, credential, runtime, contract, ledger, or owner-design byte changed.

Next requested action: `s16a.planner` files bounded PLAN r21 parented to this review, closing F1-F5. No restack, fold, ready, CI/CD, or merge follows from r20 or this verdict.

ACTIONS_GIT_REF: no source/test/branch/PR action; implementation worktree `/Users/jack/Programming/harness-s16a-conformance` clean at local/remote-equal `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; only write is review payload, engine draft move, and submission; daemon/client fingerprint `ddd5f9f5c84e3c51dbe06ce7d1d7174604b5f0c17524c155091f3c6738bcc6c0`
FINAL_GIT_STATUS_SHORT: none - implementation worktree clean at local/remote-equal `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; governing checkout before filing has modified daemon-owned INDEX, modified operator-owned CHECKPOINTS, and six daemon-rendered current relays untracked
