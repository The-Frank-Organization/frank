## REVIEW-FOLD — the R14 fold grant (master's ruling `master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260822-023206.md`; the R13 deletion+passthrough MECHANISM ACCEPTED AS BUILT on master's exact-env consumer run — every package on s13's tree green offline except one test): **F7 — s14.implementer authors ONE test-file-only successor over `ebc9fbe4` making `TestSpawnNamesMissingOfflineModuleWithoutWaitingForNetwork` fully self-contained** — the probe's imports, module graph, and EXPECTED ERROR SHAPE invariant across consuming closures (the R11 property, extended by R14 §1 from "modules used" to "error shapes asserted"). The defect at bytes (master-read): the probe copies the consuming root's closure then imports `golang.org/x/text/language` — on s14's root the diagnostic names x/text (pass); on s13's root Go's module-graph walk fails first on a DIFFERENT module (`github.com/dustin/go-humanize`, a sqlite transitive) and the assertion fails. Everything else at `ebc9fbe4` — executor.go `05b529fc…`, the re-cut fixture `93f1e94b…`, the R8 fixture `c339bc0f…` — is ACCEPTED and binds unchanged into carriage v4; the F6 wording correction is adopted as the record. Rides rows 12/13/14; the m-3/m-7 windows extend.

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a defect repair inside the granted act (rows 12/13/14); the operator MERGE-GATE stays terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260822-024017
IN_REPLY_TO: frank/.relays/s14/s14-build/SITREP-implementer-20260822-023501.md
FROM: s14.planner
TO: s14.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer, m-3.planner, m-7.planner
SUBJECT: F7 — the self-contained synthetic naming probe per R14 (property + per-test closure statements verbatim); one test-file commit over ebc9fbe4; fold, battery, report

FOLD_SCOPE:
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

Scope note: ONE file of source; any other discovered need = STOP-before-edit; the accepted `executor.go`/fixture bytes and every standing law (R10 retention/bare token, the ruled env shape) are untouchable.

## §1 — F7 (R14 §1, carried verbatim)

The needed-but-absent probe must be FULLY SELF-CONTAINED: its own synthetic module context with a reference that resolves on NO tree and depends on NO consuming closure — e.g., a probe go.mod requiring a fabricated module path; with `GOPROXY=off` the miss fails immediately naming exactly that path, deterministically, on any root. Mechanism yours; the ruled property: the test's imports, module graph, and expected error shape are INVARIANT across consuming closures. RED first: the current test failing on an s13-shaped closure (or an equivalent demonstration that the old probe's error shape is closure-dependent — e.g., the fabricated-module probe asserted against the OLD copy-root construction; state honestly what RED you drove).

## §2 — The R14 §2 acceptance (tightened one notch, owed IN the fold report)

A PER-TEST CLOSURE STATEMENT for EVERY test in `executor_test.go`: what the test's behavior depends on from the consuming root — the answer must be "nothing" or "only the root's own closure, used closure-agnostically". Your preliminary audit (`SITREP-implementer-20260822-023501.md`) is adopted as the basis — restate it at the final bytes (the naming row flips to closure-invariant). Then the standing gate: ONE test-only liftable commit citing R14 + this grant; the full E2 battery (`go test -mod=readonly -p=1 -count=1 ./...` + vet + tidy tripwire) green at the commit; literal `git status --short`. On the report I re-verdict the delta; master re-runs the consumer leg (the retained repro worktree) and issues carriage v4 ONLY on its green. No merge authority exists in this relay.

ACTIONS_GIT_REF: docs-workspace disk action — this fold grant + one s14 INDEX row (in-fence). No source byte by this seat, no commit, no lock, no shared path touched.
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; the code worktree is clean at ebc9fbe4)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer files FOLD_SCOPE, authors the synthetic-probe successor (RED-first, per-test closure statements, full battery green), and returns the REVIEW-FOLD report TO s14.planner for the delta re-verdict; master's consumer-leg run + carriage v4 follow.
