## REVIEW-FOLD scope - s5-b panel must-fixes

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
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-073232.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: FOLD_SCOPE before edits for MF-1/MF-2/MF-3 only
FOLD_SCOPE:
- internal/engine/submit.go -> in
- test/fixtures/s5_gate_raise_test.go -> in
- test/replay/zeroloss/zeroloss_test.go -> in
- test/fixtures/s5_iph_test.go -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: no code edits claimed; s5-b-mechanisms@81e0551; worktree status clean
FINAL_GIT_STATUS_SHORT: none - clean tree

Scope basis:
- MF-1: strip raise-derived headers on all reject paths in `internal/engine/submit.go`; add the rejected-record persistence fixture in `test/fixtures/s5_gate_raise_test.go`.
- MF-2: strengthen the canonical-wins replay leg in `test/replay/zeroloss/zeroloss_test.go` so the corrupted projection is asserted present while the replay view remains canonical.
- MF-3: add egress drain sentinel strings to the S5 I-PH sweep in `test/fixtures/s5_iph_test.go`.

Not accepted for this fold:
- OPT-1/OPT-2/OPT-3 optional items.
- T7 and all registered carries.
- Merge, push, PR, and any live activation work.
