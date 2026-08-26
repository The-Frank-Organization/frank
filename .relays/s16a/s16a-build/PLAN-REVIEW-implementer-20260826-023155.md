## PLAN-REVIEW — s16a WP3-coda + WP4 PLAN r19 approved: the joined D1 ruling closes F1/F2 and the explicit CT-D01/CT-D02 recuts close F3

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-13
PARENT_DISPATCH_ID: s16a-build-19
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — r19 binds the routed owner ruling and the executable seam-predicate recuts; implementation still requires the separate pair-Planner scope diff and addressed token, and the operator's next gate remains WP5 MERGE-GATE
PLAN_LOCK_ID: s16a-build-19
IN_REPLY_TO: s16a-build/PLAN-planner-20260826-022447.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer, m-10.planner, m-9.planner, m-8.planner
SUBJECT: approve PLAN r19 — conversion-before-probe fcntl coda, narrowed handover recording residual, typed in-memory replacement outcomes, and executable D01/D02 recuts are bounded and testable
PLAN_REVIEW_VERDICT: approve

`PLAN_REVIEW_VERDICT: approve` for exact PLAN r19 SHA-256 `a9b630f9f7d8c67749093c0f704dfd5ca9e8000a0421f5589a8a79ef4706dbab` at `s16a-build/PLAN-planner-20260826-022447.md`.

Findings: none.

## Review disposition

R19 is the correct successor to the r18 must-revise review: FROM `s16a.planner`, TO this seat, plan-only, parented to `s16a-build-plan-review-12`, exact-file lint clean under `--no-freshness`, and containing no live implementation or merge token. The joined D1 ruling at `s16a-wp34/RECONCILE-orchestrator-planner-20260826-021752.md` re-hashes to `f5072ff072e3a07f9611e5dc03f9c7fbad96311747bf8d95935b8b9f6f32d435` and is engine-rendered.

- F1 is closed: controller conversion to the `fcntl F_SETLK` family lands before the broker probe, with stable never-unlinked inode, one-descriptor/no-same-process-close lifetime, Linux/Darwin holder-PID floor, bounded/blocking acquisition, and probe-after-accept ordering all explicit. The named `session.go` and broker handshake packages can contain the required platform-tagged helpers without a dependency or `go.mod` change.
- F2 is closed by the owner-routed narrowing: the coda lands the probe, verified live-session replacement, old-session close, and exactly four typed in-memory outcomes; durable `control_handover` recording plus `adopted` recording defer in full on already-filed residual `R-S16A-CTRL-HANDOVER-REC`. The plan forbids relabeling a log/test hook/in-memory slice as durable recording and authorizes no m-10 byte.
- F3 is closed: CT-D01 is explicitly replaced with executable arbitrary-precision integer and refusal predicates across applicable encode/digest paths; CT-D02 is explicitly replaced with an exhaustive shared vector corpus executed against all three consumers. Both predicates remain anti-vacuous and RED until the extraction satisfies them.

The accepted r18 portions remain bounded: D2 is 32 OS-CSPRNG bytes rendered as 64 lowercase hex; D3 is disclosure-only; the shared canonicalizer removes rather than shadows the three duplicate implementations; D03, `fieldspec.CanonicalMarshal`, contracts/ledger/master bytes, dependencies, unrelated conductor internals, merge/ready, and E3/exit remain out of scope. The final acceptance partition is exactly `64 GREEN / 0 RED / 64`, with per-commit plain suite, vet, tagged census, sentinel, and bijection proof.

Boundary contract: approved as written. The controller writes the OS lock and handshake; the broker reads peer PID/lock holder plus token/generation and surfaces one typed result to the controller/supervisor. The shared canonicalizer writes the single canonical/refusal behavior consumed by appipc, connector, and worker; the executable shared corpus reads all three consumer surfaces. Durable handover recording is not a reader-less partial build because every recording surface is excluded now and registered as one deferred residual.

Tests / verification:
- E2 focused baseline passed: `go test -count=1 ./internal/broker ./internal/appctl/brokerclient ./internal/appipc ./internal/connector/jcs ./internal/worker/jcs`.
- E2 Darwin floor cross-build passed: `GOOS=darwin GOARCH=arm64 go test ./internal/broker ./internal/appctl/brokerclient`.
- Expected pre-implementation RED was reproduced: `go test -tags seam -count=1 ./test/seam -run 'TestCT_D0[12]'` fails only CT-D01 on the current float-policy disagreement and CT-D02 on the absent shared implementation/vector execution.
- Implementation worktree is clean and local/remote-equal at `3566d37aafed8a51a2cd9effa4abbf658c9224fa`.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — CI-1 controls broker controller authority and token custody; the owner ruling is bound verbatim
- migration/backfill/destructive-write/canonical-data-repair: no
- money/inventory/orders/planning/accounting/trust-critical-state: yes — controller handover and canonical evidence are trust-critical, with exact fail-closed acceptance
- AI-or-automation-acts-downstream: yes — the broker authorizes governed tool operations; scope stays behind the existing authority fence
- worker/scheduler/queue/retry/async-side-effect: yes — concurrent accept and live-session replacement are explicitly scoped and tested
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-7/m-10 broker boundary and the three-module canonicalizer are owner-routed and bounded
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: no — r19 explicitly replaces both contradictory/vacuous seam predicates with executable contract predicates
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — the only deferred recording scope is named on the filed residual; no live-verify claim is made
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

Remaining gates: this approval authorizes no implementation. `s16a.planner` must mechanically emit `SCOPE_DIFF_RESULT: all-in` over r19's controller/broker/token/canonicalizer/test surfaces before a separately addressed IMPL relay may carry the live token. Any new m-10 recording byte, dependency change, D03/fieldspec byte, or other named exclusion is a deviation and returns UP. Merge, ready-for-review, CI/CD, E3, exit, deployment, and release remain unauthorized.

No implementation action: no source, test, branch, commit, push, PR, merge, store, credential, provider, or runtime byte was changed by this review.

Next requested action: `s16a.planner` performs the mechanical r19 scope diff; if and only if all rows are in, submit the separate IMPL relay carrying the live token and parent it to `s16a-build-plan-review-13`.

ACTIONS_GIT_REF: no source/test/branch action; implementation worktree `/Users/jack/Programming/harness-s16a-conformance` clean at local/remote-equal `3566d37aafed8a51a2cd9effa4abbf658c9224fa`; this seat's only write is this review payload, the engine draft move, and relay submission.
FINAL_GIT_STATUS_SHORT: none — implementation worktree clean at `3566d37aafed8a51a2cd9effa4abbf658c9224fa`
