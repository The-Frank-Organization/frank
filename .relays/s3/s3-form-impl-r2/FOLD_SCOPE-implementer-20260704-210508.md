## REVIEW-FOLD FOLD_SCOPE - s3-form-impl-r2 adversarial review fold

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s3-form-impl-r2-implementer-fold
PARENT_DISPATCH_ID: s3-form-impl-r2-review
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s3-grill-s3-form
DESIGN_DOC_ID: s3-slice-3-design
DESIGN_LOCK_ID: s3-slice-3-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s3-slice-3-plan
IN_REPLY_TO: s3-form-impl-r2/SITREP-planner-20260704-210019.md
FROM: s3-form.implementer
TO: s3-form.planner
CC: s3.orchestrator-planner, operator
SUBJECT: fold scope pre-filed before review-fold edits for B1-B6 and M1-M7
FOLD_SCOPE:
- README.md -> in
- cmd/frank/main.go -> in
- docs/sprints/2026-07-04-s3-slice-3/results/disposition-table.md -> in
- internal/channel/server.go -> in
- internal/engine/loop.go -> in
- internal/engine/pipeline_test.go -> in
- internal/engine/submit.go -> in
- internal/fieldspec/fieldspec.go -> in
- internal/fieldspec/registry.json -> in
- internal/fieldspec/registry_test.go -> in
- internal/fieldspec/render.go -> in
- internal/fieldspec/render_test.go -> in
- internal/fieldspec/validate.go -> in
- internal/fieldspec/validate_test.go -> in
- internal/lineage/lineage.go -> in
- internal/lineage/lineage_test.go -> in
- internal/obligation/owed_test.go -> in
- internal/recover/recover.go -> in
- internal/tables/tables.go -> in
- internal/tables/tables_test.go -> in
- test/fixtures/main_assembly_test.go -> in
- test/fixtures/sweep_test.go -> in
- test/replay/dispositions.json -> in
- test/replay/harness.go -> in
- test/replay/replay_test.go -> in
FOLD_SCOPE_RESULT: all-in

Scope note: M6 names `internal/store/projections.go`, which is not present in the r2 dispatch SCOPE_DIFF row set. I am not touching that file under this fold scope. The intended M6 fold is to remove the live submit path's rescan by building owed projections from the in-memory tables in `internal/engine/submit.go` and to pin that with an in-scope probe. If that cannot satisfy the finding, I will stop and send a deviation relay before editing any out-of-scope file.

ACTIONS_GIT_REF: branch s3-form-impl@269fbf8 before fold edits; this artifact and one INDEX row only in the gitignored relay tree
Next requested action: implement the fold on the existing `s3-form-impl` branch and report per-finding disposition.
