## SITREP — s14 T9.4 stopped before edit: R9-I3's token-prefixed `FailingDetail` requires an out-of-fence observe-contract change because the live executor-origin validator accepts only the exact bare token

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s14-build-t94-observe-contract-blocker-1
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — master must first rule the observe-contract scope conflict; implementation remains stopped and merge remains operator-gated
GRILL_REQUIRED: no — this reports the exact R9-I3 stop condition at live bytes and requests bounded scope disposition, not a redesign by the pair
FILED_AT_LOCAL: 20260821-143057
IN_REPLY_TO: frank/.relays/s14/s14-build/IMPL-planner-20260821-142737.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-3.planner, m-7.planner
SUBJECT: T9.4 blocked before edit — enriched executor failure detail is rejected by the exact-match observe contract outside R9's executor-only fence

Phase: IMPL, stopped before the first T9.4 source/test edit under R9-I3 and the successor dispatch's explicit schema/observe-change tripwire.
Current artifact: exact r7 dispatch SHA-256 `fe6ec6a53beb2db961a1e77f8bd2e71ebbff0e9a97474f4dc4b46aba2e19b884`; source branch `s14-m8-connector` remains at T9 HEAD `471532a151762d496186f2daee70b21fdd71062f`.

Claims:
- The requested executor change is mechanically possible inside `internal/executor/executor.go:288-294`: the mismatch branch currently assigns the exact bare string `suite-exit-mismatch`, and the existing capture is available to `finalizeRun`. — evidence E1 — source bytes
- The enriched value required by R9-I3 cannot survive the existing downstream observe contract. `internal/observe/verdict_binding.go:135-142` accepts an executor-origin failed verdict only when `FailingDetail == "suite-exit-mismatch"` exactly. A value such as `suite-exit-mismatch: <bounded tail>` therefore fails `validOriginClass`. — evidence E1 — source bytes
- The failure is semantic, not merely diagnostic: `validateBoundVerdict` converts that rejected detail to `check-machinery-verdict-origin-class-mismatch` at `internal/observe/verdict_binding.go:64-67`, replacing the intended observed fail with a machinery fault. That violates R9-I3's required outcome/predicate/rung preservation. — evidence E1 — source bytes
- Repairing the contract requires at least `internal/observe/verdict_binding.go` plus its focused tests. Those bytes are outside the one-commit `internal/executor/**` grant, and R9-I3 expressly says any observe-contract change stops and returns to master before edit. — evidence E1 — R9 + r7 scope diff + source bytes
- I2 remains implementable without widening the fence by pre-seeding the fresh run-local `GOMODCACHE` from the go.sum-verified host module cache. Passing the host cache through was rejected locally because the existing out-of-fence S8 isolation fixture asserts the run-scoped cache path exactly; no mechanism byte was written. — evidence E1 — `test/fixtures/s8_executor_test.go:15-26` + executor environment bytes

Actions:
- Verified the exact dispatch token, sole-addressee routing, approved parent, plan hash, linked worktree, and source status. — ACTIONS_GIT_REF source read-only at `s14-m8-connector@471532a151762d496186f2daee70b21fdd71062f`
- Stopped before T9.4 test or production edit, staging, commit, or evidence run. The inherited R8 fixture and T10 worktree bytes remain untouched and unstaged. — ACTIONS_GIT_REF source `git status --short` remains ` M test/fixtures/s8_exit_gate_test.go` plus `?? internal/connector/stream/`
- Filed only this lane-owned SITREP and one live-EOF s14 INDEX row in the governance worktree. — ACTIONS_GIT_REF disk paths `frank/.relays/s14/s14-build/SITREP-implementer-20260821-143057.md` and `frank/.relays/s14/INDEX.md`

Done: authority and plan-lock verification; isolation verification; R9 root-cause seam inspection; observe-contract collision identified at exact bytes; allowed I2 mechanism selected but not implemented.
Not done: no T9.4 edit/test/commit; no T9.5 commit; no GREEN-ladder run; no T10 commit; no T11-T15 work; no push, PR, merge, publication, or release.
Blocked: R9-I3 requires enriched `FailingDetail`, while the current observe validator admits only the bare token. The required contract change is outside the granted `internal/executor/**` package.
Scope drift risk: editing `internal/observe/**` under the current dispatch would violate R9-I4 and the successor SCOPE_DIFF; omitting the change would make the enriched verdict fail closed as a machinery fault.
Tests / verification: no source test was run because the plan's stop condition fired before edit. Static contract proof is at `internal/observe/verdict_binding.go:64-67,135-142` and `internal/executor/executor.go:288-294`; the incoming dispatch linted OK before action.

Writes: none to source; this report writes only the s14 relay lane and its INDEX row.
Reads: executor failure verdict and captured output; observe executor-origin tuple/class validation.
Target entity: an observed failed suite verdict whose `FailingDetail` retains the `suite-exit-mismatch` prefix and carries a bounded inner-output tail.
Downstream consumer: `observe.validateBoundVerdict`, which currently rejects the target value before it reaches executable-claim evidence.
Contract: R9-I3 requires detail-only enrichment with outcome/predicate/rung mapping preserved; current exact-match validation makes those requirements mutually incompatible inside the executor-only fence.
Proof: exact comparison and fault-conversion source paths cited above; a successor scope ruling must bind the validator and focused regression tests before implementation resumes.
No-consumer action: stop and route to master; do not emit enriched detail that the validator converts to a machinery fault.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no — the stopped T9.4 edit has not touched credential or authorization bytes
- migration/backfill/destructive-write/canonical-data-repair: no — no data or migration action
- money/inventory/orders/planning/accounting/trust-critical-state: yes — executable-claim verdict integrity is trust-critical state
- AI-or-automation-acts-downstream: yes — this validator governs automated suite evidence consumed by the conductor
- worker/scheduler/queue/retry/async-side-effect: yes — T9.4 governs a spawned nested process
- cross-repo/service-contract/generated-schema/shared-API-event: yes — executor-to-observe verdict validation is a shared internal contract and the repair may be lifted across sibling slices
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is involved
- test-runtime-role-mismatch: yes — the blocker is exactly a mismatch between the requested executor diagnostic and its runtime validator
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — continuing requires master to expand or recut the executor-only scope
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or risk acceptance is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: governance-only report action — this SITREP plus one live-EOF s14 INDEX row; source stays `s14-m8-connector@471532a151762d496186f2daee70b21fdd71062f` with inherited unstaged T9.5 and untracked T10 bytes; no source/test/dependency/stage/commit/branch/push/PR/merge/publication action
FINAL_GIT_STATUS_SHORT:
 M .relays/s13/INDEX.md
 M .relays/s14/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260821-142739.md
?? .relays/s14/s14-build/IMPL-planner-20260821-142737.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-142139.md
?? .relays/s14/s14-build/PLAN-planner-20260821-140427.md
?? .relays/s14/s14-build/SITREP-implementer-20260821-143057.md
Next requested action: `s14.planner` routes this exact collision to master; master either grants the minimal `internal/observe/**` contract-and-test companion with owner review, supplies a different already-admitted detail carrier, or revises I3. A fresh successor plan review and dispatch are required before s14 resumes T9.4.
