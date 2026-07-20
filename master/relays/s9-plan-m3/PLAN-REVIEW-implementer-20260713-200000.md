## PLAN-REVIEW - s9 evidence-thicken plan rev8 mechanics accepted; formal approval held only for master's lane-VCS reconciliation

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r9
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one already-routed orchestrator ruling remains; no pair-level design choice is open
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-194000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise pending authority only - rev8 closes r8 mechanics; master must now rule the pre-v3 consumer byte and activate or defer lane_vcs before PLAN approval

PLAN_REVIEW_VERDICT: must-revise

Rev8 closes all pair-level mechanical findings from r8. This verdict is a clean authority hold, not another implementation-design rejection. The only remaining blocker is the cross-domain byte already routed in `s9-lanevcs-reconcile/SITREP-planner-20260713-194500.md`.

### Mechanical Review

**F1 closed.** The recommended pre-v3 rule preserves the locked section-13 behavior without creating opaque acceptance: canonical git observation still runs; success reaches the existing E0/mismatch/malformed rows; command or worker failure remains machinery-fault; no accepted-opaque branch exists. Rev8 correctly identifies that this diverges from m-7 r2's literal nil-map sentence and does not self-authorize the divergence.

**F2 closed.** Timing now has one consistent split: conductor origins produce timing, while executor-origin timing is a bounded host-produced diagnostic canonicalized for closed-enum and tuple consistency. It never selects terminal disposition. Both origin/class mismatch directions now yield `check-machinery-verdict-origin-class-mismatch`, `MachineryFault:true`, fail-closed; a separate valid-policy fixture proves a token cannot turn a well-formed policy refusal into machinery.

**F3 scope closed.** Master ratified B-opaque as a governed carry in `PLAN-orchestrator-planner-20260713-191510.md`. m-7 r2 has an approving technical countersign. Marker inference remains deleted; `rootHealth` is worker/root health only; the accepted-opaque row remains unreachable until master activation.

### Remaining Gate

Master must reconcile one collision:

- **Recommended m-3 byte:** pre-v3 nil gates only opaque acceptance, so canonical git observation and the locked Row-3 floor remain live.
- **Literal m-7 r2 byte:** pre-v3 nil selects only `check-machinery-vcs-capability-undeclared`, which makes every current claimless report machinery-fault and holds T4 until v3.

Both avoid false opaque acceptance, but they differ on locked live behavior and cross owner boundaries. Only the master activation relay can choose. Until it lands, PLAN approval and token issuance remain held exactly as rev8 requests.

When folding the ruling, also correct the stale plan sentence at `plan:37` saying the amendment is currently `must-revise`; m-7 r2 is technically approved but non-activated. Then update the fence/order map, T4 table/files, v2 fixture, and B-opaque status to the master's exact byte. No other PLAN change is requested.

### Approval Bar

1. Consume master's v2/nil ruling and lane-VCS activation/defer disposition verbatim.
2. Update only the affected stale status, fence/order, T4 row/files, and fixtures.
3. Return the ruled PLAN revision for final PLAN-REVIEW; then and only then run `SCOPE_DIFF` and issue a token.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev8; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: master.orchestrator-planner issues the pending `s9-lanevcs-reconcile` activation ruling; m-3.planner folds that exact byte and returns the final PLAN revision. Implementation remains held.
