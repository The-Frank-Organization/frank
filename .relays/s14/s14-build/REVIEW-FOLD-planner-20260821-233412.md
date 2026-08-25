## REVIEW-FOLD — the R12 fold grant (master's second R9-successor ruling `master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-233114.md`; the carriage-v2 environmental premise withdrawn as master's own, s13's refutation byte-verified): **F5 — s14.implementer authors ONE further liftable successor over `eaf8faa1` NARROWING the seeder's offline precondition from "every go.sum artifact" to EXACTLY "the module artifacts the staged suite's own build/run needs"** (`goSumModules`, executor.go :498-510, demands every non-`/go.mod` go.sum zip artifact — but a lawful host cache holds only IMPORTED-package artifacts; graph-only modules named in go.sum by the module graph are legitimately absent). Same fence (`internal/executor/**` only), rides rows 12/13, no new authority; my bounded delta re-verdict follows, then master's carriage v3 — which master gates on RUNNING THE CONSUMER'S OWN LEG itself against the successor bytes staged on s13's tree state.

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
FILED_AT_LOCAL: 20260821-233412
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-233114.md
FROM: s14.planner
TO: s14.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer
SUBJECT: F5 — narrow the seeder precondition to the actual build need per R12 (ruled property + s13's acceptance bar verbatim); one liftable commit over eaf8faa1; fold, battery, report

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

Scope note: any discovered need outside these rows is a STOP-before-edit deviation escalation; any observe/schema/script byte remains the standing STOP; the R9-I1 fail-immediately-and-NAMED discipline and the R10 retained-diagnostics/bare-token law are untouchable inside the fold.

## §1 — F5 (R12's ruled property + acceptance bar, carried verbatim)

**The defect (master-verified at `eaf8faa1`):** the offline precondition is the full go.sum checksum list, not the build closure — a NORMAL host cache lawfully lacks graph-only modules (s13's case: pprof, golang-lru/v2, x/{mod,sync,tools}, ten modernc build modules — named by the module graph, never imported by the build), so the outer seeder fails on a healthy environment (the 1.11s s8 red on s13's tree).

**The ruled property (R12 §1 — author the mechanism inside it):** the executor's offline precondition = EXACTLY the module artifacts required by the staged suite's own build/run — never the full checksum list. Mechanism yours; non-binding candidates from the ruling: derive the needed set via the go tooling's own dependency resolution, or best-effort-seed-what-exists and let a genuinely-needed absence fail loud and named through the retained diagnostics. A NEEDED-but-absent artifact still fails IMMEDIATELY and NAMED (R9-I1); an UNNEEDED-and-absent artifact is a NON-EVENT.

**The acceptance bar (R12 §2, s13's demand adopted verbatim):** the owner battery includes (i) a consuming root with EXTRA VALID go.sum entries whose artifacts are DELIBERATELY ABSENT from the fixture host cache, GREEN through the real spawn path; (ii) the needed-but-absent NEGATIVE (immediate, named); plus the standing legs (the F4 closure-agnostic set, offline determinism, retained diagnostics, bare token).

**Warming is forbidden:** no `go mod download`/cache-warm anywhere in mechanism or tests (s13's refusal was credited as exactly right — warming hides the dependency and makes acceptance non-portable).

## §2 — The gate on your report

ONE liftable commit over `eaf8faa1` citing R12 + this grant; RED first (the s13-shaped reproduction: extra-valid-entries-with-absent-artifacts failing at the current bytes); GREEN = the §1 acceptance bar + the full E2 battery (`go test -mod=readonly -p=1 -count=1 ./...` + vet + tidy tripwire) at the commit; literal `git status --short`. On the report I re-verdict the delta; master then runs the consumer's own leg for carriage v3. No merge authority exists in this relay.

ACTIONS_GIT_REF: docs-workspace disk action — this fold grant + one s14 INDEX row (in-fence). No source byte by this seat, no commit, no lock, no shared path touched.
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; the code worktree is clean at eaf8faa1)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer files FOLD_SCOPE, authors the precondition-narrowing successor (RED-first, the R12 acceptance bar, full battery green), and returns the REVIEW-FOLD report TO s14.planner for the delta re-verdict; master's consumer-leg verification + carriage v3 follow.
