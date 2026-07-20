## FOLD_SCOPE — s11 optional findings 1 and 2 only

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s11-build-slice-review-fold-scope
PARENT_DISPATCH_ID: s11-build-slice-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair-internal optional fold only; merge remains operator-only and T5/T10 remain acceptance-OPEN behind g2/dc
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-slice-review/REVIEW-FOLD-planner-20260714-163830.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: pre-edit scope for optional findings 1 and 2 — centralize the acceptance-bounce edge relation and make replacement cadence-slot restart independently asserted; findings 3 through 11 remain recorded with no code change

FOLD_SCOPE:
- frank/internal/bounce/edges.go -> in (finding 1 shared named edge classification)
- frank/internal/bounce/edges_test.go -> in (finding 1 relation assertion)
- frank/internal/engine/fsm.go -> in (finding 1 shared acceptance-bounce consumer)
- frank/internal/store/projections.go -> in (finding 1 shared Bucket-D consumer)
- frank/internal/engine/resummon.go -> in (finding 2 testable G4 input construction)
- frank/internal/engine/resummon_test.go -> in (finding 2 replacement cadence-slot assertion)
- frank/.relays/s11/fold-optionals-red-green.md -> in (fold E2 evidence)
- frank/.relays/s11/mechanical-tables.md -> in (path-to-finding license reconciliation)
FOLD_SCOPE_RESULT: all-in

Finding 1 will preserve the intentional relation `Bucket D = acceptance-bounce subset + stale_choice_set`; no edge token or projection membership changes. Finding 2 will preserve both shipped G4 cadence-slot literals while proving the replacement decision is assigned `g4-no-response-1` by the scheduler input constructor. Tests are written and observed RED before production edits, then the touched packages plus s11 fixtures run GREEN as directed by the review relay.

Findings 3–11 are not edit licenses in this fold: 3–7 are unreachable-path/refactor observations with correct current fallback behavior; 8–9 would change host-facing error semantics; 10 is historical commit evidence and cannot be repaired by a new byte; 11 is already backstopped by the positive projection rows. Their disposition will be restated in the fold report.

ACTIONS_GIT_REF: none — pre-edit FOLD_SCOPE artifact; no frank edit yet; s11-comms-thicken clean at 547ada9aa89b6edcb98769ba27f0418439236441
FINAL_GIT_STATUS_SHORT: none — clean s11-comms-thicken worktree at 547ada9aa89b6edcb98769ba27f0418439236441
Next requested action: s11.implementer applies exactly the eight-row fold above with RED-first evidence, commits and pushes the existing PR branch, then returns a REVIEW-FOLD report. No merge authority is requested or implied.
