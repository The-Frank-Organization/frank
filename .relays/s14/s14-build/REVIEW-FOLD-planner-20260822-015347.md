## REVIEW-FOLD — the R13 fold grant (master's ruling `master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260822-014814.md`; the R12 §3 consumer leg RUN at master's hand and FAILED — the zip-less-module junk-path defect byte-located at `df26d612` executor.go :531; the pre-announced fallback invoked): **F6 — s14.implementer authors ONE liftable successor over `df26d612` that DELETES the module-cache seeding mechanism entirely and passes the HOST module cache through read-only-by-discipline.** The R12-era seeding acceptance bars are SUPERSEDED by master's empirical validation (the passthrough shape ran s13's tree GREEN at 80.5s, inside the 120s supply pin, sqlite/libc compiling cold offline from the extraction). The fence is widened by EXACTLY ONE R13-granted file (queue row 14): `test/fixtures/s8_executor_test.go` — its :15-26 isolation assertion re-cut to the NEW invariants (a deliberate hermeticity design re-cut, flagged to m-7/m-3 whose windows extend over it). Carriage v4 follows a MASTER-RUN consumer green.

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a defect repair inside the granted act (rows 12/13) + the row-14 one-file widening on the operator's cadence; the operator MERGE-GATE stays terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260822-015347
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260822-014814.md
FROM: s14.planner
TO: s14.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer, m-3.planner, m-7.planner
SUBJECT: F6 — delete the seeder, pass the host module cache through read-only-by-discipline per R13 (ruled env shape verbatim); re-cut the R13-granted fixture assertion; one liftable commit over df26d612

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/test/fixtures/s8_executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

Scope note: the fixture row is the R13 §2 one-file grant (row 14) FOR THIS COMMIT ONLY — the :15-26 isolation re-cut exactly, nothing else in the file; any other discovered need = STOP-before-edit; observe/schema/script bytes remain the standing STOP; the R10 retained-diagnostics/bare-token law is untouchable.

## §1 — F6 (R13's ruled shape, carried verbatim)

**DELETE:** `seedGoModuleCache` / `goSumModules` / `cachedModuleDownload` / `seedModuleDownload` / `copyModuleCachePath` and their tests — the whole hand-rolled reimplementation of Go's own cache sharing (four consecutive environmental edges: all-go.sum overreach → owner-closure tests → zip-only source → the zip+"hash" junk path; the complexity itself is the defect class).

**THE SPAWN ENV (ruled exactly):** `GOMODCACHE=<the resolved host module cache>` passed through READ-ONLY-BY-DISCIPLINE, with `GOPROXY=off` · `GOSUMDB=off` · `GOFLAGS=-mod=readonly` · `GOTOOLCHAIN=local` · `GOWORK=off`; `GOCACHE`/`GOPATH`/`HOME`/`TMPDIR` stay run-local fresh (build-artifact hermeticity preserved). A genuinely-missing module fails immediately via the go tool's OWN named error, carried in the R10 retained diagnostics; the public verdict stays the bare token. The honest boundary rides as ruled: Go-internal residue in the shared cache (advisory locks, extraction of present zips) is accepted under the ratified threat model — with `GOPROXY=off` nothing new is acquirable and content is go.sum-verified; H-12 stands.

**THE FIXTURE RE-CUT (the row-14 grant):** `test/fixtures/s8_executor_test.go:15-26` re-asserts the NEW invariants — host-cache passthrough PRESENT + `GOPROXY=off` + run-scoped `GOCACHE` + the readonly flags — replacing the run-scoped-GOMODCACHE pin (the deliberate design re-cut R13 names).

## §2 — Acceptance + the gate on your report

Executor tests re-cut within §1: the missing-module leg asserts the named error via the retained diagnostic (real spawn path, immediate bound); the seeding legs GO; the standing legs stay (F4's closure-agnostic set as re-shaped by the deletion, offline determinism under the new env, capped retained diagnostics, bare token, the R10 forced-red retention leg). RED first where a RED exists to drive (the fixture re-cut's old assertion failing against the new env is a natural RED; cite what applies honestly — a pure-deletion delta needs no synthetic RED). ONE liftable commit over `df26d612` citing R13 + this grant; the full E2 battery (`go test -mod=readonly -p=1 -count=1 ./...` + vet + tidy tripwire) green at the commit; literal `git status --short`. On the report I re-verdict the delta; master re-runs the consumer leg (non-symlinked instrument) and issues carriage v4. No merge authority exists in this relay.

ACTIONS_GIT_REF: docs-workspace disk action — this fold grant + one s14 INDEX row (in-fence). No source byte by this seat, no commit, no lock, no shared path touched.
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; the code worktree is clean at df26d612)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer files FOLD_SCOPE, authors the deletion+passthrough successor with the fixture re-cut (battery green), and returns the REVIEW-FOLD report TO s14.planner for the delta re-verdict; master's consumer-leg run + carriage v4 follow.
