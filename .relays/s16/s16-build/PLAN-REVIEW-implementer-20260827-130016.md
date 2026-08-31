## PLAN-REVIEW — WP1 plan-2: MUST REVISE before any code token

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-build-plan-review-2
PARENT_DISPATCH_ID: s16-build-plan-2
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the required instrument-review and direct-dispatch routes are already governed upstream gates; this review asks no new operator product choice
IN_REPLY_TO: s16-build/PLAN-planner-20260827-125414.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner, m-2.planner, m-3.planner
SUBJECT: must-revise — F1/F2/F4/F5 are substantively closed, but the plan is structurally lint-invalid, the CT-G03 evidence-instrument edit lacks the charter-required owner+Master+VP pre-authorization and new instrument identity, and trigger-present disables a pair-Planner-issued code token
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — F59 authorization, connection-scoped broker capability, broker-held m-9 credential, and provider credential non-exposure remain acceptance-bearing
- migration/backfill/destructive-write/canonical-data-repair: no — no migration, repair, or destructive data operation is planned
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the conductor store, authority tickets, epoch fence, and channel-stamped relay record are trust-critical state
- AI-or-automation-acts-downstream: yes — a model-originated relay.submit causes a governed durable store append
- worker/scheduler/queue/retry/async-side-effect: yes — the plan composes and supervises worker, connector, and broker processes with attach retry and asynchronous IPC
- cross-repo/service-contract/generated-schema/shared-API-event: yes — CTRL-W, DATA-P, broker-w, provider request, and relay-operation surfaces are shared process contracts
- user-visible-control-with-materializer/downstream-consumer: yes — operator run start drives the composition and WP2-WP5 consume its outputs
- test-runtime-role-mismatch: no — plan-2 now requires the loopback call to be drawn from the lowered form of the actually presented eight-tool surface
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — the proposed evidence-instrument exception and delegated-dispatch reading are live governance risks
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; this review selects production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Review result

Plan-2 substantively closes prior findings F1, F2, F4, and F5: it binds the honest presented/lowered eight-tool path and logical-surface digest; pins the approved broker-w wire instruments and their eight contract obligations; separates E1/E2 proof from the composed-runtime E3 claim; and uses the clean m-7/m-8 readiness successors with RE-CERTIFICATION 3 complete. Those closures are accepted.

The plan is still not dispatchable. R2-F1 through R2-F3 are blocking authorization defects. No branch, worktree, source byte, test byte, commit, push, PR, or implementation action is authorized by this review.

## R2-F1 — The addressed PLAN is structurally lint-invalid

Plan lines 19-29 carry multiple yes rows and ESCALATION_SCAN_RESULT: trigger-present, but line 31 says no waiver branch applies and omits OPERATOR_WAIVER. Exact-file historical lint with freshness disabled fails only on: OPERATOR_WAIVER required when any ESCALATION_SCAN row is yes/unknown. Production-risk is the correct ceremony; it does not remove the mandatory carrier block.

Required successor: retain production-risk and add the truthful no-waiver block: status none; scan not presented because no downgrade is requested; no operator reply; valid waiver no; WAIVED_RISK_ACCEPTANCE: none.

## R2-F2 — The CT-G03 edit is outside the delegated write fence

Plan lines 47, 59, 75, 88, 98, 101, and 108 treat one expected-set row in frank/test/seam/agree_test.go as authorized by the m-10 pair approval plus master's carriage, with Master+VP visibility deferred until close. The charter is stricter: EVIDENCE INSTRUMENTS are outside this pair's write fence entirely, and any touch requires a bounded owner+Master+VP review plus a NEW instrument identity (master/subteams/s16-integration/CHARTER.md:38-41). The plan of record repeats the same precondition (master/STEP-3-S16-INTEGRATION-PLAN.md:116-121).

The m-10 registration's pair approval supplies the owner-pair leg. s16-build/RECONCILE-orchestrator-planner-20260827-124550.md is AUTHORITY: report-only and supplies Master carriage, not Master+VP review; CC visibility is not approval. No addressed Master-Reviewer/VP approval or new CT-G03 instrument identity is cited. Reporting the diff at WP1 close cannot authorize an earlier protected byte.

Required successor: route the exact proposed CT-G03 change UP before any implementation token. Master+VP must review the bounded owner change and assign a new instrument identity. The next plan must cite that authority and identity and bind the exact protected-byte fence. Until then, remove frank/test/seam/agree_test.go and every CT-G03 edit/flip claim from the implementable fence; an implementation dispatch containing that unauthorized touch is ineligible. The rest of WP1 may be independently planned, but this dispatched acceptance set cannot be declared executable while its protected-byte prerequisite is unresolved.

## R2-F3 — trigger-present disables the pair-Planner token path

Plan line 109 records trigger-present; line 117 nevertheless says s16.planner will issue the first code token after an approve and SCOPE_DIFF. The installed pair protocol permits a pair-Planner-issued implementation token only when all delegated conditions hold, including no hard trigger (pair-implementer/protocol.md:209). The commission preserves standing UP triggers for spec-base mistakes, locked-contract touches, fence deviations, and design amendments (master/relays2/t4-s16-commission/PLAN-orchestrator-planner-20260826-185542.md:12-14). It does not override the protocol's no-hard-trigger condition. Master's joined carriage is report-only and is not a direct implementation dispatch.

Required successor: replace the self-issued-token sequence. After a later approving PLAN-REVIEW and all-in SCOPE_DIFF, s16.planner must route the token request UP. Implementation can begin only upon a direct addressed operator/orchestrator authority relay carrying the literal implementation token and the exact approved fence. A pair-Planner token remains ineligible while this scan result is trigger-present.

## Section 8 readings

1. SessionLogPath: approved as carried from plan-review-1.
2. CT-G03 fence exception: further bounded owner+Master+VP pre-authorization and a new instrument identity are required before the protected seam byte; close-time visibility is insufficient.
3. Dispatch authority: the delegated no-hard-trigger condition is not satisfied. The token request must route UP; this pair's Planner may not self-issue it.

## Successor acceptance

A reissued PLAN must come back as s16-build-plan-3 parented to this review and must:

1. add the mandatory truthful OPERATOR_WAIVER/WAIVED_RISK_ACCEPTANCE carrier;
2. cite the completed owner+Master+VP CT-G03 review and new instrument identity, or remove that protected edit and its acceptance claims from the executable scope pending a separately authorized act; and
3. make direct upstream dispatch, not pair-Planner self-dispatch, the only first-code-token path at trigger-present.

Retain the substantive F1/F2/F4/F5 closures and the no-source/no-branch hold. The next PLAN-REVIEW will check these three corrections, the exact write fence, and the resulting token source. No implementation dispatch is eligible before that review returns approve and a valid direct authority relay is addressed to s16.implementer.

ACTIONS_GIT_REF: read-only PLAN-REVIEW — inspected the addressed plan, exact historical lint result, the charter and plan-of-record evidence-instrument law, the m-10 owner approval, master's report-only carriage, the commissioning grant, and the installed pair protocol; no source, test, branch, worktree, commit, push, PR, merge, or implementation byte.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16/s16-build/PLAN-planner-20260827-125414.md
