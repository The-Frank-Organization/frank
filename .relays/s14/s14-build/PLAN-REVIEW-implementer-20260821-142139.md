## PLAN-REVIEW — s14 build plan r7 bounded R9 executor-determinism fold: APPROVE; T9.4 carries all six master invariants, precedes the standalone T9.5 fixture commit, and binds the four-part GREEN ladder before preserved T10

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s14-build-plan-review-7
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this approves only the bounded R9 fold; implementation still requires a successor dispatch and merge remains operator-gated
GRILL_REQUIRED: no — master R9 already ruled the bounded owner repair; this review checks exact carriage and executability without redesign
FILED_AT_LOCAL: 20260821-142139
IN_REPLY_TO: frank/.relays/s14/s14-build/PLAN-planner-20260821-140427.md
PLAN_LOCK_ID: s14-build @ sha256 2f1ab1f2627eac98ecaa9c1ef107fb870f52dd75e8d61edd213b96dc397cfdba
PLAN_REVIEW_VERDICT: approve
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: approve plan r7 — exact R9 executor-only repair, repair-before-evidence order, four-part GREEN ladder, and unchanged carried plan

## Verdict

`PLAN_REVIEW_VERDICT: approve` for exact plan r7 SHA-256 `2f1ab1f2627eac98ecaa9c1ef107fb870f52dd75e8d61edd213b96dc397cfdba`.

The bounded fold matches master R9 SHA-256 `0f950fec07092f836458059778fb19abd2a162b64c350080fa87fd94a360d67f` on every operative invariant:

1. T9.4 makes nested module resolution offline by setting `GOPROXY=off` or a strict equivalent; missing content must fail immediately and be named.
2. The implementer chooses only one ruled local-source mechanism: pre-seed the staged cache from the go.sum-verified host module cache, or pass the host `GOMODCACHE` through. The fold must state the choice and isolation rationale; fresh per-run `GOCACHE` and `GOPATH` remain, and no host mutation beyond Go cache/lock discipline is admitted.
3. `finalizeRun` enriches only the existing string `FailingDetail` on the failure path with an explicitly capped captured-output tail. `suite-exit-mismatch` remains the prefix and outcome, predicate, rung, and timing semantics remain unchanged. Any record-schema or observe-contract change stops before edit and returns to master.
4. The T9.4 fence is exactly `internal/executor/**`, source and tests, for one commit. `scripts/dogfood-suite.sh` and every byte outside that package remain untouched; the package returns to consume-only after that commit.
5. GREEN is indivisible and ordered: offline zero-network module resolution; forced-inner-red detail surfacing with cap and token prefix; three consecutive isolated repaired-fixture greens; then two consecutive full batteries.
6. T9.4 is one liftable R9-citing commit, followed by the distinct R8-citing T9.5 fixture commit. Both repairs precede all GREEN evidence, and T10 remains preserved until that evidence is green.

The live source confirms the ruled seams exist without expanding the design surface. `internal/executor/executor.go` constructs the child environment in one list, currently assigns fresh per-run `GOCACHE`, `GOMODCACHE`, and `GOPATH` without `GOPROXY`, captures both stdout and stderr in `cappedCapture`, and drops that capture when `finalizeRun` emits the plain-string `suite-exit-mismatch`. New focused tests can live inside `internal/executor/**`; no record-schema or observe-package edit is required by the plan.

The executable starting point is branch `s14-m8-connector` at committed T9 HEAD `471532a151762d496186f2daee70b21fdd71062f`, with the R8 fixture repair present as an **unstaged** modification to `test/fixtures/s8_exit_gate_test.go` and preserved T10 as untracked `internal/connector/stream/**`. Thus r7's phrase “staged uncommitted” is read only as “present in the source worktree,” not as a claim about the Git index. That imprecision is non-operative because T9.4 and T9.5 each have disjoint, explicit one-commit fences; the successor must use path-exact staging and verify each commit tree.

T1a–T9, T10–T15, the R8 conditions, §2a, acceptance, escalation rails, the R6 module-directive law, end-review, restack, domain-owner routing, and operator MERGE-GATE carry byte-for-meaning from r6 and its predecessors. No implementation authority appears in this PLAN relay. No source, test, dependency, stage, commit, or branch mutation occurred during review. Implementation may resume only under a successor `PHASE: IMPL` relay containing the valid bare own-line token, addressed solely to `s14.implementer`, and parented by `s14-build-plan-review-7`.

## Boundary contract review

Writes: T9.4 writes only `internal/executor/**` for its one R9-citing commit; T9.5 then writes only `test/fixtures/s8_exit_gate_test.go` for its distinct R8-citing commit; afterward the carried plan writes only the existing connector and lane-owned evidence surfaces.
Reads: the host go.sum-verified module cache, staged suite tree, existing executor capture/verdict seams, the repaired production-dogfood fixture, and frozen connector contracts.
Target entity: deterministic offline execution of nested governed suites plus bounded failure diagnostics, followed by the unchanged fixture-capacity repair and connector stream work.
Downstream consumer: every s14 evidence run and, by byte-identical lift only, sibling slice batteries/restack; the repair diff also routes FYI to `m-3.planner` with the fold.
Contract: no live module proxy; local module source only; fresh run-local build/work caches; token-prefixed capped inner detail; verdict semantics byte-preserved; executor-only one-commit scope; R9 then R8 commits before GREEN evidence.
Proof: focused spawn-env offline test; forced-red capped-detail test; exact commit-tree checks; three consecutive isolated nested-suite greens; two consecutive full `go test -mod=readonly -p=1 -count=1 ./...` batteries; build/vet/diff/module tripwires.
No-consumer action: stop before edit or commit on any non-executor T9.4 byte, script edit, network-capable module resolution, non-host/go.sum-verified module source, unbounded detail, changed verdict mapping, schema/observe-contract change, combined repair commit, failed evidence leg, or any carried tripwire; route the exact deviation upward.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the carried connector slice retains the provider credential and authorization boundary
- migration/backfill/destructive-write/canonical-data-repair: no — T9.4/T9.5 and the carried connector work add no migration, destructive write, or canonical-data repair
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the governed-suite ladder and carried connector send boundary are trust-critical
- AI-or-automation-acts-downstream: yes — the executor runs governed checks and the connector ultimately sends model requests
- worker/scheduler/queue/retry/async-side-effect: yes — T9.4 governs a spawned nested process and the carried slice retains streaming, cancellation, transport, and one-shot effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the executor repair is shared conductor-tree machinery and may be lifted byte-identically across sibling slices
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is added by this fold
- test-runtime-role-mismatch: yes — R9 repairs the nested production-dogfood test runtime's network-dependent module source and missing inner diagnostics
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — R9 grants exactly one package for one commit and retains the schema-change stop, semantic mappings, evidence gates, and terminal merge gate
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade, failed-test acceptance, or risk waiver is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane-governance action only — this PLAN-REVIEW file plus one live-EOF s14 INDEX row; source worktree remains `s14-m8-connector@471532a151762d496186f2daee70b21fdd71062f` with inherited ` M test/fixtures/s8_exit_gate_test.go` and `?? internal/connector/stream/`; no source, test, dependency, branch, stage, commit, push, PR, merge, or publication action
FINAL_GIT_STATUS_SHORT:
 M .relays/s14/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-142139.md
?? .relays/s14/s14-build/PLAN-planner-20260821-140427.md
Next requested action: `s14.planner` issues the successor implementation dispatch parented by `s14-build-plan-review-7`; hand-relay that exact path to `s14.implementer`, which may then execute T9.4, commit T9.5 separately, prove the four-part GREEN ladder, and resume T10.
