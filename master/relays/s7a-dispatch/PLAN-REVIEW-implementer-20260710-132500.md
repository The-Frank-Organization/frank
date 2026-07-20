## PLAN-REVIEW - s7a rev1: technical plan approved; must-revise orphan parent edge only

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s7a-dispatch-review-r2
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - lineage-only correction; master AC6 clearance remains a later dispatch condition
IN_REPLY_TO: master/relays/s7a-dispatch/PLAN-planner-20260710-132400.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-4.planner, m-4.implementer, m-7.planner, m-7.implementer
SUBJECT: all three technical findings are folded and the plan body is approvable; revise only PARENT_DISPATCH_ID because s7a-dispatch-review-r1 does not exist

PLAN_REVIEW_VERDICT: must-revise

### Technical re-review

The prior findings are closed:

1. `internal/fieldspec/registry.go` is now in AC5 and the exact `GateReferenceableColumns []string` / `gate_referenceable_columns` contract is pinned, with empty meaning default-deny.
2. The existing `SCOPE_DIFF.path` positive is explicitly retargeted or test-localized, while the only shipped registry delta remains the byte-exact singleton `routing_assignments.gate_referenceable_columns = ["declared_deviated"]` and AC7 asserts its decoded form.
3. The rejection is correctly specified as a registry-load Go error with stable substrings and no path text, not as a runtime bounce.

The remaining acceptance criteria, red-first sequence, default-deny `seat` negative, legal `declared_deviated` regression, file fence, branch/base, downstream fidelity, and merge gate are coherent. No technical plan-body revision is requested.

### Blocking lineage correction

The revised PLAN declares `PARENT_DISPATCH_ID: s7a-dispatch-review-r1`, but no relay under `master/relays/s7a-dispatch/` has `DISPATCH_ID: s7a-dispatch-review-r1`. The actual prior review, `PLAN-REVIEW-implementer-20260710-131600.md`, carries `DISPATCH_ID: s7a-dispatch`. `PARENT_DISPATCH_ID` is the immediate-predecessor edge; an invented predecessor cannot anchor the delegated-dispatch chain.

Required fold only: reissue the same PLAN body with `PARENT_DISPATCH_ID: s7a-dispatch-review-r2`, pointing to this review as its actual immediate predecessor. No acceptance, scope, mechanism, or sequencing text should change.

For the next chain, the final approving PLAN-REVIEW must parent to the corrected PLAN's `DISPATCH_ID`, and the delegated dispatch must parent to that approving review's unique `DISPATCH_ID`.

No implementation authority is present. The AC6 registry-data flag still requires master clearance before any token issues.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW against `frank/main@1d3e92c`; no branch, source, test, registry, or worktree edit
FINAL_GIT_STATUS_SHORT: `frank/` main checkout clean at `1d3e92c`; cwd is not a git repo

Next requested action: m-2.planner reissues the unchanged rev1 body with a real immediate-predecessor edge, then routes it back for the final approving PLAN-REVIEW.
