## DESIGN-REVIEW - s8 executor host r3 must revise two final lock-wording echoes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-executor-review-r3
PARENT_DISPATCH_ID: s8-design-m7-executor
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - technical approval remains separate from the operator-owned OS-sandbox election
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-executor
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-executor-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-executor/DESIGN-planner-20260711-033342.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
SUBJECT: r3 must revise - the rejected-alternative row retains obsolete claim-level network wording and the lock-impact line still names superseded r2; all run-contract folds otherwise approve

DESIGN_REVIEW_VERDICT: must-revise

R3 resolves the three r2 findings in the normative run contract: cache content is run-scoped, replay safety is provided-surface scoped, and network use is unsanctioned rather than absent. Two final echoes remain.

## Findings

### F1 - The rejected-alternative record still preserves the obsolete network claim

Section 2.6 correctly says network use is unsanctioned and never claimed absent (`2026-07-11-s8-executor-host.md:24`). The GRILL_LOCK rejected alternative still ends with "claim-level only, honestly worded" (`design:74`), retaining the exact model r3 replaced. Rejected alternatives are lock content, not inert history; this wording can reintroduce an absence claim downstream.

Required fold: rewrite that row to reject **all network-absence claims without a sandbox** and preserve only the suite-policy statement: network use is unsanctioned; no network handle/credential is provided; ambient access remains possible and untested.

### F2 - The design-lock impact still names r2

The artifact status and incoming relay are r3, but the GRILL_LOCK names `s8-design-m7-executor-r2` as the reconciled lock (`design:83`). That points at the superseded revision.

Required fold: change the effective lock ID to `s8-design-m7-executor-r3` and sweep all lock-impact/reconcile text for stale revision identifiers.

## Confirmed

- All v1 cache content is run-scoped and follows the full group-death/cleanup state machine.
- At-least-once execution is now bounded to no sanctioned/provided-surface side effect for conforming checks, with ambient repetition risk explicit.
- Network is policy-unsanctioned, not enforcement-absent.
- The r1 taxonomy, exactly-once-commit, descendant-death, and closed-verdict folds remain intact.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no lock, and no sandbox choice inferred.

ACTIONS_GIT_REF: wrote this r3 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F2 into executor-host r4 and returns a final DESIGN relay for re-review.
