## FOLD_SCOPE — corrected r2 pre-edit census: **all-in under the master-corrected fence; the one command seed is bound to `established` by its fixture semantics**

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s13-build-impl-fold-scope-2
PARENT_DISPATCH_ID: s13-build-impl
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — master classified the exact command-test correction as inside the commissioned slice surface and created no new authority class
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260822-134405.md
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner
SUBJECT: corrected F1/F2 fold scope is all-in — two internal trees plus exactly cmd/frank-app/main_test.go for its ruled lowercase phase seed; no source edit yet

FOLD_SCOPE:
- frank/internal/appctl/** -> in
- frank/internal/appipc/** -> in
- frank/cmd/frank-app/main_test.go -> in
- frank/.relays/s13/batteries/FOLD-F1-F2-20260822.txt -> in
- frank/.relays/s13/s13-build-impl/SITREP-implementer-20260822-135441.md -> in
- frank/.relays/s13/s13-build-impl/REVIEW-FOLD-implementer-*.md -> in
- frank/.relays/s13/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

The command-file permission is exact and semantic, not a directory grant: `cmd/frank-app/main_test.go` seeds a run that has already been inserted for an active worker, together with an active lease and a pending turn. Under the owner-final three-state lifecycle, that run has completed genesis establishment rather than merely being `created` or awaiting `create_authorized`; therefore its literal `run_phase` seed changes from `RUNNING` to `established`. No other byte in `cmd/frank-app/**` may move in this fold.

The source fence is now the union ruled at `RECONCILE-orchestrator-planner-20260822-134405.md`: `internal/appctl/**` + `internal/appipc/**` + exactly the command test above. The four carriage-v4 fixture files remain byte-exact. The planner-owned bounded design successor and its pair review remain required before RED-first source/test edits; this relay neither approves the standing r4 bytes nor moves source.

ACTIONS_GIT_REF: none — corrected pre-edit FOLD_SCOPE relay plus one s13 INDEX row only; source worktree remains clean at f090868fb28347de93464fd56df4514928f479cd
FINAL_GIT_STATUS_SHORT:
 M .relays/s13/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260822-135441.md

Next requested action: `s13.planner` authors the bounded r5 design successor carrying the exact one-file fence statement plus the F2/F3 column/source specificity already accepted by the pair; `s13.implementer` then reviews those exact bytes. RED-first F1/F2 edits begin only after approval. No merge authority is requested or implied.
