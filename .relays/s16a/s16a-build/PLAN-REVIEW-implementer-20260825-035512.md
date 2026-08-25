## PLAN-REVIEW — s16a WP2 PLAN r13 must revise cross-side commit sequencing only

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-8
PARENT_DISPATCH_ID: s16a-build-13
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — s16a.planner can issue the bounded sequencing successor; implementation and merge remain separately gated
PLAN_LOCK_ID: s16a-build-13
IN_REPLY_TO: s16a-build/PLAN-planner-20260825-035301.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: must-revise sequencing only — both required halves of every cross-side row must land in the same named green commit
PLAN_REVIEW_VERDICT: must-revise

`PLAN_REVIEW_VERDICT: must-revise` for exact PLAN r13 SHA-256 `eb9c19495c46a1d42c057690bbebad8640c79abaf76a8f9ab3e2cb6090b8dace` at `s16a-build/PLAN-planner-20260825-035301.md`.

R13 closes the two prior findings on authoritative scope and proof: §2 now carries the complete 44-owner-half matrix over 38 distinct rows, including B08-W and A16-CP; §3 now requires a valid moving tagged census rather than an impossible all-green battery. One residual execution contradiction remains in the suggested order.

## Blocking finding

- **S16A-WP2-PR-F3 — §4's cross-side sequencing can create a commit that violates §3.1.** Section 3 requires the named row/cluster CT to turn GREEN in its fix commit and requires the exact then-planned RED set at every commit. Section 4 nevertheless separates owner halves into W, CP, and CN waves, then says each cross-side row lands “first half then second half, order per wave.” A first-half commit can leave that row RED, directly violating §3.1. The order is also incomplete on its own terms: A16-CP is absent from the CP wave and B08-W is absent from the W-provider wave. The “mechanics yours” qualifier does not resolve which commit owns the row-green transition.

## Required bounded successor

Reissue the plan parented to this review with §4 corrected so each of the six cross-side rows lands all required halves in ONE named row commit: A13(W+CP), A14(W+CP), A16(W+CP), B08(W+CN), C01(CP+CN), and C08(CP+CN). Place each complete row commit once in the suggested order. If a multi-row cluster is genuinely required, name the cluster and require every included CT GREEN in that same commit. Do not authorize a committed first-half state whose named row remains RED.

Preserve every other substantive r13 byte unless a separately identified correction is necessary. In particular, do not reopen the complete W26/CP15/CN3 scope matrix, 38-row/59-5 arithmetic, moving-census law, non-conforming-side-only rule, coupled-test restriction, census modernization, A-2 carries, boundary contract, or WP3/WP4/out-of-scope fences.

Accepted and not reopened: authority lineage is exact (`s16a-build-13` parented to `s16a-build-plan-review-7`); engine-rendered bytes match the reviewed digest; r12 F1/F2 are substantively closed in §§2–3; all governing pins, the clean `s16a-conformance@f7040666…` worktree, and the accepted 21/43 coda remain unchanged.

Boundary-contract disposition: scope, target, consumers, and E2 proof are sufficient. Approval is withheld only because the ordered commit boundary is ambiguous and can contradict the per-commit acceptance law.

Remaining gates: no SCOPE_DIFF or implementation token may issue on this verdict. A sequencing-correct successor requires approval before the mechanical SCOPE_DIFF and separately addressed implementation relay. Merge remains operator-only.

Tests / verification: E1 full r13 read; exact r12→r13 diff; exact hash/lint/INDEX parent checks; direct reconciliation of §§2–4; clean implementation-worktree branch/head/status. No source tests or implementation actions were authorized.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — WP2 repairs authority and fail-closed seams; review-only here
- migration/backfill/destructive-write/canonical-data-repair: yes — C04 persistence remains planned; no review-time write
- money/inventory/orders/planning/accounting/trust-critical-state: yes — trust-critical seam
- AI-or-automation-acts-downstream: yes — worker and tool-authority behavior changes
- worker/scheduler/queue/retry/async-side-effect: yes — worker/control/connector lifecycles
- cross-repo/service-contract/generated-schema/shared-API-event: yes — registered cross-module contracts
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: no — r13 fixed the moving tagged-census rule
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — F3 is an execution/commit ambiguity and must be revised
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

No implementation action: no production/test byte, branch state, dependency, stage, commit, push, PR, merge, provider, credential, store, or runtime state changed.

ACTIONS_GIT_REF: engine-lane governance act only — drafted under `frank/.relays/s16a/.engine/drafts/s16a.implementer/` for submission; no source/test or implementation-worktree action claimed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/PLAN-REVIEW-implementer-20260824-225758.md
?? frank/.relays/s16a/s16a-build/PLAN-REVIEW-implementer-20260824-233824.md
?? frank/.relays/s16a/s16a-build/PLAN-REVIEW-implementer-20260825-013608.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260824-225355.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260824-232655.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260825-011947.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260825-035301.md
?? frank/.relays/s16a/s16a-impl/IMPL-planner-20260824-234004.md
?? frank/.relays/s16a/s16a-impl/SITREP-implementer-20260825-011602.md
