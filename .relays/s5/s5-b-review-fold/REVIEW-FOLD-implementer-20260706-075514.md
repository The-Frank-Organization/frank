## REVIEW-FOLD scope - s5-b MF-1b remainder

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-b-review-fold
PARENT_DISPATCH_ID: s5-b-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-mechanisms
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-075026.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: FOLD_SCOPE before edits for MF-1b absorb-bounce remainder only
FOLD_SCOPE:
- internal/engine/submit.go -> in
- test/fixtures/s5_gate_raise_test.go -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: no code edits claimed for MF-1b yet; s5-b-mechanisms@a11e4bb; code worktree status clean before edit
FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/

Scope basis:
- MF-1b: absorb candidate with no seat-supplied `gate_category` plus detector hit plus bad `PARENT_DISPATCH_ID` must persist rejected without `gate_category`, `gate_category_raised`, or `gate_category_pick`.
- Implementation surface is the existing reject cleanup in `internal/engine/submit.go`.
- Fixture surface is the existing S5 gate raise fixture in `test/fixtures/s5_gate_raise_test.go`.

Not accepted for this fold:
- MF-2/MF-3 already-verified work.
- Optional items.
- T7, registered carries, integration ordering changes, merge, push, PR, or live activation.
