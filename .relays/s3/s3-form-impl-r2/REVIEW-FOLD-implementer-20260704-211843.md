## REVIEW-FOLD - s3-form-impl-r2 review blockers folded

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s3-form-impl-r2-implementer-fold-report
PARENT_DISPATCH_ID: s3-form-impl-r2-implementer-fold
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
IN_REPLY_TO: s3-form-impl-r2/FOLD_SCOPE-implementer-20260704-210508.md
FROM: s3-form.implementer
TO: s3-form.planner
CC: s3.orchestrator-planner, operator
SUBJECT: B1-B6 and M1-M7 folded on s3-form-impl@fe7308e; verification green
FOLD_SCOPE:
- cmd/frank/main.go -> in
- docs/sprints/2026-07-04-s3-slice-3/results/disposition-table.md -> in
- internal/engine/pipeline_test.go -> in
- internal/engine/submit.go -> in
- internal/fieldspec/registry.json -> in
- internal/fieldspec/render.go -> in
- internal/fieldspec/render_test.go -> in
- internal/fieldspec/validate.go -> in
- internal/fieldspec/validate_test.go -> in
- internal/lineage/lineage.go -> in
- internal/lineage/lineage_test.go -> in
- internal/tables/tables.go -> in
- internal/tables/tables_test.go -> in
- test/fixtures/main_assembly_test.go -> in
- test/fixtures/sweep_test.go -> in
- test/replay/dispositions.json -> in
- test/replay/harness.go -> in
- test/replay/replay_test.go -> in
FOLD_SCOPE_RESULT: all-in

Summary:
Fold commit: `s3-form-impl@fe7308ed830e7b11e3ed1b31e694a81d4ae07ad8` (`fe7308e s3 REVIEW-FOLD: close form lineage replay blockers`). No PR or merge action taken.

Per-finding disposition:
B1: folded. `test/replay/harness.go` now runs `ReplayAll` through real `Render`/`SubmitHandler`/`Validate`/`lineage.Engine` paths; accepted legs call the handler on a fresh store; fail legs require an actual rejection or lineage bounce; filename classifier removed.
B2: folded. `internal/lineage/lineage.go` adds L1 design-review chain, L2 pair-planner dispatch grant chain, L6 orchestrator-review visibility, and a narrowed `RealGrantState`; `registry.json` adds the pair-planner grant seat-scope path gated by that state.
B3: folded. `Validate` rejects `PARENT_DISPATCH_ID` outside the active candidate set; `cmd/frank` threads `turnContextForSeat` into Describe and Submit; stale-positive and outside-set submit tests landed.
B4: folded. `internal/tables.Live` publishes immutable snapshots; handler-side readers use snapshots; race coverage added for concurrent publish/read.
B5: folded. Main and loop consumers are repointed through the published table owner after commit and quarantine; binary fixture submits an owed disposition after live quarantine.
B6: folded. Served tables are built and published after recovery; corrupt-at-start binary fixture verifies the conductor serves and reads return the quarantine disposition instead of assembly failure.
M1: folded. Monotonic floors render `[floor, MAX]` and validate below-floor HGR/GRILL choices.
M2: folded. Disposition rows carry `context: live|reconstructed-observe`; reconstructed observe rows use the exact required label and stay out of `LiveCaughtCount`.
M3: folded. `dispositions.json` is generated code-first from 110 actual `relay-lint.py` assertion sites, with explicit line anchors including 844/846/848, 855/857/859/861/863/868, and 873; table generation is compared byte-for-byte.
M4: folded. Obsolete rows carry closed-vocabulary grounds only: `vanished-markdown-channel`, `strict-submit-api`, or `one-channel-role-from-stamp`.
M5: folded. Stale classifier wording swept; public test API renamed `ReplayAll`; no `ClassifyAll`, classifier, or filename wording remains in the replay/result surfaces.
M6: folded in-scope. The live owed projection path now builds from the supplied tables in `internal/engine/submit.go`; `internal/store/projections.go` was not touched because it is outside the r2 SCOPE_DIFF row set.
M7: folded. Parent outside-set and stale-positive render-to-submit legs were added; the main path no longer passes `lineage.TurnContext{}` directly to `ActiveLineageCandidates`.

Verification:
`go test -count=1 ./...` passed: all 20 test packages ok plus no-test packages.
`go vet ./...` passed with exit 0 and no output.
`go test -race -count=1 ./internal/engine ./internal/tables ./internal/fieldspec` passed for all three packages.
`git diff --check` passed with exit 0.

ACTIONS_GIT_REF: branch s3-form-impl@fe7308ed830e7b11e3ed1b31e694a81d4ae07ad8; commit fe7308e; fold-scope artifact `.relays/s3/s3-form-impl-r2/FOLD_SCOPE-implementer-20260704-210508.md` linted OK before source edits
FINAL_GIT_STATUS_SHORT: none - clean tree after commit
Remaining risk:
The M6 production rescan was removed from the live submit caller without editing `internal/store/projections.go`; that file remains as an out-of-scope compatibility helper for non-live callers/tests.
Next requested action:
Planner re-review of this fold report and branch `s3-form-impl@fe7308ed830e7b11e3ed1b31e694a81d4ae07ad8`.
