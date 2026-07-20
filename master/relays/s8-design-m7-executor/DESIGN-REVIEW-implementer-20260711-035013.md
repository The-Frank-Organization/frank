## DESIGN-REVIEW - s8 executor host r4 technically approved; operator and reconcile gates remain

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-executor-review-r4
PARENT_DISPATCH_ID: s8-design-m7-executor
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - this is technical approval only; the operator must ratify the v1 residual or elect the sandbox round before lock effectiveness
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-executor
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-executor-grill-r1
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s8-design-m7-executor/DESIGN-planner-20260711-034701.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
SUBJECT: approve technical r4 - all prior findings are resolved; lock remains conditional on operator sandbox election and master reconciliation with m-3's provided-surface scoping flag

DESIGN_REVIEW_VERDICT: approve

R4 is technically sound for the m-7 executor-host surface. The two r3 findings are closed and no new host-contract blocker remains.

## Review

- Network semantics now consistently reject every absence claim without a sandbox. Suite-class policy marks network unsanctioned; no network handle or credential is provided; ambient access remains possible and untested.
- The effective lock identifier is `s8-design-m7-executor-r4`, matching the reviewed revision and `s8-executor-grill-r1`.
- The host accepts only suite-class runs while preserving m-3's existing side-effecting taxonomy and policy ownership.
- Manifest identity binds pre-spawn bytes without claiming staged-copy immutability. All cache content is run-scoped under the provided workdir.
- Execution is honestly at-least-once under crash; the conductor guarantee is one committed verdict/one pivot, with ambient same-uid repetition risk explicit.
- Timeout cleanup requires direct-child reap and bounded process-group death confirmation; survivor expiry becomes a machinery fault with workdir preservation.
- `RunResult` remains host-internal. Only m-3's closed `CheckVerdict` crosses the boundary, followed by m-3's conductor-side redaction.
- Fixtures FX-EXE-1..6 cover the claim boundary, kill/reap state machine, committed-verdict idempotency, ceiling marking, two-axis fault disposition, and closed verdict shape.

## Remaining Gates

- This approval does not ratify the OS-sandbox election. The reconciled lock becomes effective only after an operator record selects the v1 residual or a sandbox design round.
- Master must reconcile m-3's "writes are confined there" sentence to the provided-surface plus stated-residual scope; this review does not edit or approve m-3-owned wording.
- No c4 amendment, PLAN, or IMPL authority follows from this relay alone.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no effective lock, and no sandbox choice inferred.

ACTIONS_GIT_REF: wrote this approving r4 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner returns the technically approved executor r4 artifact to master for m-3 reconciliation and the separate operator sandbox decision; no further pair revision is requested unless reconciliation changes the m-7 contract.
