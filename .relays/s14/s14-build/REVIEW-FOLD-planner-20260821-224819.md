## REVIEW-FOLD — the R11 fold grant (master's successor ruling `master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-223925.md` over the R9/R10 executor repair's PORTABILITY DEFECT, surfaced by s13's carriage consumption and byte-verified by master): **F4 — s14.implementer authors ONE liftable CLOSURE-AGNOSTIC successor commit over `4aea922e`, executor package only** (the existing rows-12/13 grant; a defect repair INSIDE the granted act; no new authority). The end-review's CLEAN is superseded ONLY at the changed bytes; my bounded re-verdict covers the delta; master then re-verifies and issues carriage v2; s13 swaps ONLY the executor bytes (the R8 fixture byte stays valid and consumed). s14's HOLD at the (new) head for the post-s13-merge restack is otherwise unchanged.

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a defect repair inside the existing grant, riding queue rows 12/13; the operator MERGE-GATE stays terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-224819
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-223925.md
FROM: s14.planner
TO: s14.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer
SUBJECT: F4 — the closure-agnostic executor successor per R11 (ruled property + owner battery verbatim); one liftable commit over 4aea922e; fold, battery, report

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

Scope note: any discovered need outside these rows is a STOP-before-edit deviation escalation (R11 §3: if the inherited-GOMODCACHE recursion must change anyway, stay in-fence and STATE WHY in the fold report); any observe/schema/script byte remains the standing STOP.

## §1 — F4 (R11's ruled property + owner battery, carried verbatim)

**The defect (master-verified at `b86b8bc1`):** (a) `goModuleCachePath()` (executor.go :457-460) honors an inherited `GOMODCACHE` first — inside a nested dogfood run the inner executor treats the outer restricted run-local cache as its host source; (b) `TestSpawnPreseedsRunLocalModuleCacheWithProxyOff`'s synthetic inner module hard-requires `golang.org/x/text v0.41.0` — present in s14's root closure BY ACCIDENT of its own manifest, absent from s13's. Byte-identical code, closure-dependent behavior — the class the convergent-overlap law exists to exclude.

**The ruled property (R11 §1 — author the mechanism inside it):** the executor package — the seeding mechanism AND its tests — behaves correctly for ANY consuming root closure, INCLUDING one with zero external requires. No executor test may depend on a module outside the consuming root's own go.sum closure: a synthetic fixture module must be fully self-provisioned by the test (e.g., a local replace/vendored fixture needing no cache seed), or the assertion restated over the root's ACTUAL closure ("whatever the closure is, the nested run works offline").

**The owner battery (R11 §2 — the ruled evidence bar):** a fixture ROOT MODULE WITH NO x/text (ideally zero external requires) exercised through the executor spawn path, PLUS the nested-dogfood-shaped leg — portability proven at the PROPERTY grain, never by the accident of the owner branch's manifest.

**R11 §3 carried:** inner-seeds-from-outer is ACCEPTABLE once the tests are closure-agnostic (inner ⊆ outer ⊆ host, each containing the root closure); change it only if you must, in-fence, with the why stated.

## §2 — The gate on your report

ONE liftable commit over `4aea922e` citing R11 + this grant; RED first (the s13-shaped reproduction: the fixture-root-without-x/text leg failing at the current bytes); GREEN = the owner battery + the existing R9/R10 legs (offline determinism, retained diagnostics, bare token) + the full E2 battery (`go test -mod=readonly -p=1 -count=1 ./...` + vet + tidy tripwire) at the commit; literal `git status --short`. On the report I re-verdict the delta; master then re-verifies for carriage v2. No merge authority exists in this relay.

ACTIONS_GIT_REF: docs-workspace disk action — this fold grant + one s14 INDEX row (in-fence). No source byte by this seat, no commit, no lock, no shared path touched.
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; the code worktree is clean at 4aea922e)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer files FOLD_SCOPE, authors the closure-agnostic successor (RED-first, owner battery, full battery green), and returns the REVIEW-FOLD report TO s14.planner for the delta re-verdict; master's carriage v2 follows.
