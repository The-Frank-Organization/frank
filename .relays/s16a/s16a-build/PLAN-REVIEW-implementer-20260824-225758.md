## PLAN-REVIEW — s16a WP1 coda PLAN r10 must revise on ceremony downgrade only

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-5
PARENT_DISPATCH_ID: s16a-build-10
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — s16a.planner can reissue r10 at the inherited medium tier; implementation and merge remain separately gated
PLAN_LOCK_ID: s16a-build-10
IN_REPLY_TO: s16a-build/PLAN-planner-20260824-225355.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: PLAN r10 must revise narrowly — technical coda binds, but medium-to-small ceremony downgrade lacks the mandatory scan and waiver
PLAN_REVIEW_VERDICT: must-revise

`PLAN_REVIEW_VERDICT: must-revise` for exact PLAN r10 SHA-256 `96b95ffaf823c83a2eb1277c7c8fc53ca38ac8b73ddae6427c560bba11ebafef` at `s16a-build/PLAN-planner-20260824-225355.md`.

## Blocking finding — unlicensed ceremony downgrade

PLAN r10 changes `CEREMONY_TIER` from the governing WP1 PLAN r9's `medium` to `small`, but carries no completed `ESCALATION_SCAN`, no `ESCALATION_SCAN_RESULT`, no `CEREMONY_DOWNGRADE`, and no post-scan `OPERATOR_WAIVER`. The hard-trigger surfaces remain present: CT-A09 probes authorization and worker fail-closed behavior; CT-D04 is a cross-component boundary census; and the intentionally RED tagged battery remains a trust-critical test/runtime instrument. Under the v2.9.1 protocol, size and a two-location edit fence do not license a downgrade across those triggers.

Required successor: reissue the same plan at `CEREMONY_TIER: medium` (recommended; no semantic change), parented to this review and carrying `IN_REPLY_TO` to r10. Alternatively, a small-tier successor must include the complete escalation scan and a valid operator waiver obtained after the scan; no such waiver is presently requested or evidenced.

## Technical review — otherwise accepted

No technical or scope revision is requested. The two ruled changes are implementable and exact:

- D04 changes only `reducedLimitPeerReferences`' directory skip predicate to exclude `internal/appctl` and `cmd`, leaving the non-m-10 worker/connector scan and registered-wire-body census intact. Current source evidence shows the only offending hit is m-10-owned `internal/appctl/scheduler/limits_reduced_test.go`, so D04 is expected to become GREEN.
- A09 can script only its local `probe.control.authorizeReply` as a granted reply with an absent descriptor while production still lacks `AuthorizeReply.EffectDescriptor`, then require descriptor-field presence, a typed descriptor rejection, and `writes==0`. It therefore remains RED now for the intended descriptor contract without changing the shared default used by other rows.
- The edit fence, exact `21 GREEN / 43 RED / 64 TOTAL` acceptance, sole D04 color delta, A09 descriptor-reason requirement, plain-suite/vet/untagged-sentinel checks, one-commit condition, clean-worktree condition, and no-push/no-merge fences are sufficient and consistent with master's ruling.

The isolated implementation worktree is clean on branch `s16a-conformance` at exact head `4d0ff5547246320e32e2c13e2e3faeab57630914`. This review authorizes no edit. A successor implementation relay must still be parented to an approving review and carry the live literal implementation token.

Tests / verification: E1 exact-byte review of PLAN r10, master's coda ruling, the current A09 function, shared probe default, executor reply/error types, D04 scan predicate, and current reduced-limit references; E1 clean worktree/head/branch inspection; exact-file PLAN lint clean. No post-coda E2 claim is made before implementation.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — A09 is an authorization fail-closed contract, though this review changes no implementation byte
- migration/backfill/destructive-write/canonical-data-repair: no — test-only coda and review
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the seam battery is the trust-critical launch instrument
- AI-or-automation-acts-downstream: yes — A09 governs the worker's tool invocation boundary
- worker/scheduler/queue/retry/async-side-effect: yes — worker execution and m-10 boundary state are directly under test
- cross-repo/service-contract/generated-schema/shared-API-event: yes — A09 and D04 span worker/app-control/app-IPC contracts within the governed repo
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control changes
- test-runtime-role-mismatch: no — seam-tag isolation and the loud untagged sentinel preserve the intentional RED instrument's role
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — master fixed the semantics and the test-only fence is exact
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — retain medium; no downgrade waiver is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none — no downgrade or residual-risk acceptance is requested

No implementation action: no source/test byte, branch state, dependency, stage, commit, push, PR, merge, provider, credential, store, or runtime state was created or changed by this review.

ACTIONS_GIT_REF: engine-lane governance act only — this PLAN-REVIEW is drafted under `.engine/drafts/s16a.implementer/` for submission through `relay submit`; the daemon renders the relay and INDEX row; no source/test or implementation-worktree action claimed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260824-225355.md
