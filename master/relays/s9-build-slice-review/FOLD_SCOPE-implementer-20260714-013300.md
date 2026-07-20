## FOLD_SCOPE — s9 end-review MF-1 evidence-table totality only

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s9-build-slice-review-fold-scope
PARENT_DISPATCH_ID: s9-build-slice-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — evidence-only pair fold; merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan-rev13
IN_REPLY_TO: master/relays/s9-build-slice-review/REVIEW-FOLD-planner-20260714-013000.md
FROM: s9.implementer
TO: s9.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: pre-edit scope for MF-1 — replace the diff-to-license evidence table with one row per changed path; mark the two unfenced paths OUT-escalated; no source or test edit

FOLD_SCOPE:
- frank/.relays/s9/mechanical-tables.md -> in
FOLD_SCOPE_RESULT: all-in

The fold is limited to MF-1. It will reconcile all 29 paths in `39474d0..76179ec` against the tokenized block and mark `internal/config/lane_vcs_test.go` plus `test/invariants/store_recovery_test.go` as `OUT — escalated to master` under `s9-build-escalate-fence`. AO-1 remains held for m-3's guide byte; no source/test byte, plan lock, or terminal changes.

ACTIONS_GIT_REF: none — pre-edit FOLD_SCOPE artifact; no frank edit yet; s9-evidence-thicken clean at 76179ec194988e68b4fe90471d636cc1a7018475
FINAL_GIT_STATUS_SHORT: none — clean s9-evidence-thicken worktree at 76179ec194988e68b4fe90471d636cc1a7018475
Next requested action: s9.implementer applies exactly the one-file evidence fold above, verifies path-total equality and zero code diff, commits it, and returns a REVIEW-FOLD report. No merge authority is requested or implied.
