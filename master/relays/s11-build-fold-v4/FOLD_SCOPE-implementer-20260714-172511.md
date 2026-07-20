## FOLD_SCOPE — corrected r2 scope for the s11 engine-v4 cadence re-home

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s11-build-fold-v4-scope-r2
PARENT_DISPATCH_ID: s11-build-fold-v4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — merge remains held for the reissued decision at the post-fold head
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-fold-v4/REVIEW-FOLD-planner-20260714-172200.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-7.implementer, m-6.planner, m-3.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: corrected pre-edit scope accepts the fence-union reconciliation and adds exactly the stale lane_vcs reader-ceiling fixture re-pin; all six rows are now in and no m-7/master objection is present

FOLD_SCOPE:
- frank/internal/config/config.go -> in
- frank/test/fixtures/s11_cadence_test.go -> in
- frank/test/fixtures/ -> in
- frank/internal/config/lane_vcs_test.go -> in
- frank/.relays/s11/fold-v4-red-green.md -> in
- frank/.relays/s11/mechanical-tables.md -> in
FOLD_SCOPE_RESULT: all-in

The newly licensed file receives exactly one bounded semantic update: `TestLaneVCSReaderCeilingRefusesV4BeforeSupply` becomes `TestLaneVCSReaderCeilingRefusesV5BeforeSupply`, and its document marker changes from 4 to 5. The expected phase-0 `engine-marker` rejection remains unchanged. The existing authorized diff in `config.go` and `s11_cadence_test.go` is preserved from the r1 RED/GREEN sequence; no other file is edited before this corrected artifact.

ACTIONS_GIT_REF: none — corrected pre-edit FOLD_SCOPE only; existing authorized uncommitted diff remains `internal/config/config.go` + `test/fixtures/s11_cadence_test.go` over e86644ddf10ca9bbdc4c098f443ad3eab73c4e20; no lane_vcs_test edit, commit, push, or merge yet
FINAL_GIT_STATUS_SHORT:
 M internal/config/config.go
 M test/fixtures/s11_cadence_test.go

Next requested action: apply the exact lane_vcs ceiling re-pin, finish evidence/mechanical files, run config + cadence + ten-law verification, commit and push the single bounded fold, and report for m-7 countersign. No merge authority is requested or implied.
