## FOLD_SCOPE — R12 F5 exact-build-need executor successor only

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-fold-scope-5
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded defect repair under the existing rows-12/13 grant; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-233721
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260821-233412.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer
SUBJECT: pre-edit scope — R12 F5 narrows offline seeding to available artifacts while preserving named needed misses, all-in

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

The source worktree is clean at `s14-m8-connector@eaf8faa1b96eae254c6788b9dd49386082a3acd5`. The RED mutation is a real spawn from a consuming root whose valid go.sum contains an additional graph-only module absent from an explicitly restricted fixture host cache: the current all-go.sum batch must refuse it before the suite runs.

The minimal fold will seed and hash-verify checksum-listed module artifacts that actually exist in the host cache, treat unavailable graph-only artifacts as a non-event, and preserve `GOPROXY=off` for the staged suite's own build/run. A required absent module will be exercised through an actual import and must still fail immediately, remain named in the retained private diagnostic, and keep the public bare `suite-exit-mismatch` token. Neither mechanism nor tests will warm a cache.

After the two real-spawn RED/GREEN legs, the standing F4 closure-agnostic, offline-determinism, retained-diagnostic, and bare-token cuts run with the full serialized E2 battery, race, vet, build, and non-persistent tidy tripwire. Any need outside the three rows or any observe/schema/script byte stops before edit and routes a deviation.

ACTIONS_GIT_REF: governance-only pre-edit barrier — this FOLD_SCOPE relay plus one append-only live-EOF s14 INDEX row; source remains clean at `s14-m8-connector@eaf8faa1b96eae254c6788b9dd49386082a3acd5`; no source or test edit yet
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; source worktree clean)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer drives the R12 extra-absent graph-only reproduction RED-first, preserves the needed-absent named negative, folds the minimal two-file successor, proves the full E2 battery at one liftable commit, and returns REVIEW-FOLD for delta re-verdict. No merge authority is requested or implied.
