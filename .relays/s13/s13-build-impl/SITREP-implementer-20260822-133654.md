## FOLD_SCOPE — s13 F1/F2 pre-edit scope census: **deviation-present; one required cmd test seed lies outside the carried source fence**

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s13-build-impl-fold-scope-1
PARENT_DISPATCH_ID: s13-build-impl
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — master can correct the one-file fence without an operator act; the bounded s13 design amendment and pair review remain next before source edits
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260822-133146.md
FROM: s13.implementer
TO: master.orchestrator-planner
CC: s13.planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner
SUBJECT: F1/F2 fold scope cannot yet be all-in — cmd/frank-app/main_test.go contains a literal RUNNING schema seed and must move with the ruled lowercase run_phase domain

FOLD_SCOPE:
- frank/internal/appctl/** -> in
- frank/internal/appipc/** -> in
- frank/.relays/s13/batteries/FOLD-F1-F2-20260822.txt -> in
- frank/.relays/s13/s13-build-impl/FOLD_SCOPE-implementer-20260822-133654.md -> in
- frank/.relays/s13/s13-build-impl/REVIEW-FOLD-implementer-*.md -> in
- frank/.relays/s13/INDEX.md -> in
- frank/cmd/frank-app/main_test.go -> OUT
FOLD_SCOPE_RESULT: deviation-present

No source, test, schema, battery, or design byte was edited. The OUT row is required by a mechanical literal census, not by a requested feature expansion:

- The carried Q-A ruling binds the v1 schema literally to `runs.run_phase IN ('created','create_authorized','established')`; terminality moves solely through `runs.state`.
- Exact head `f090868fb28347de93464fd56df4514928f479cd` contains a raw SQL run seed in `cmd/frank-app/main_test.go:66` with `run_phase='RUNNING'`.
- That test seed will violate the corrected v1 CHECK as soon as `internal/appctl/store/schema.go` carries the ruled domain. It must change to the applicable lowercase phase for the required full repository gate to execute.
- Retaining uppercase aliases in the schema merely to avoid the OUT edit would violate the owner-final literal domain. Skipping `./cmd/frank-app` would violate the required full gate. No in-fence production change can rewrite a raw SQL literal compiled into this out-of-fence test.

The production/test source fence otherwise remains exactly the carried `internal/appctl/**` + `internal/appipc/**`. The battery and relay rows above are evidence/handoff substrate already licensed by the locked plan; they add no product surface. The planner-authored bounded design-doc amendment is a separate actor-owned documentation step and is not claimed in this Implementer's edit scope.

ACTIONS_GIT_REF: none — no fold/source/test/schema edit; this report-only scope-census relay plus one s13 INDEX row are the only s13.implementer docs actions
FINAL_GIT_STATUS_SHORT:
 M .relays/s13/INDEX.md
 M ../master/relays/CHECKPOINTS.md
 M docs/sprints/active/2026-08-20-s13-m10-module/designs/DS-s13-m10-module-20260820.md
?? .relays/s13/s13-build-design/DESIGN-planner-20260822-133706.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260822-133654.md

Next requested action: `master.orchestrator-planner` extends the fold's source fence by exactly `frank/cmd/frank-app/main_test.go` for the literal phase-seed correction, or provides a source-backed contradiction. Then `s13.implementer` banks a corrected all-in successor scope; `s13.planner` authors the bounded design-doc amendment binding Q-A/Q-B for pair review; only after approval do RED-first F1/F2 source edits begin. No merge authority is requested or implied.

