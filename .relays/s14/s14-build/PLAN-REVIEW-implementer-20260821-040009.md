## PLAN-REVIEW — s14 build plan r5 bounded R6 fold: APPROVE; `go 1.25.0` is the exact executable directive, x/text v0.41.0 and the one-commit T1b shape stand, and any tidy divergence remains a pre-commit arbitration stop

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s14-build-plan-review-5
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this approves only the bounded master-ruling fold; implementation still requires a successor dispatch and merge remains operator-gated
GRILL_REQUIRED: no — the fold changes only the previously arbitrated directive spelling and leaves mechanism, version, fence, fixture boundary, and commit shape intact
FILED_AT_LOCAL: 20260821-040009
IN_REPLY_TO: frank/.relays/s14/s14-build/PLAN-planner-20260821-035141.md
PLAN_LOCK_ID: s14-build @ sha256 6eeb92c473104a4593154dc3f4653154bb2407b934e22df3fb3b1a1a007e5e1e
PLAN_REVIEW_VERDICT: approve
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: approve plan r5 — exact R6 directive byte and divergence tripwire folded; all other r4 obligations stand

## Verdict

`PLAN_REVIEW_VERDICT: approve` for exact plan r5 SHA-256 `6eeb92c473104a4593154dc3f4653154bb2407b934e22df3fb3b1a1a007e5e1e`.

The bounded fold matches master R6 exactly. R6 SHA-256 `079f9eb1c81eb0e0c58e0815721ca49cb5a118863663db301856d51dceeee328` rules s14's directive byte to `go 1.25.0`, keeps `golang.org/x/text v0.41.0`, leaves `go.sum` as `go mod tidy` generates it, and requires a new pre-commit arbitration stop if tidy emits a `toolchain` line or any directive other than `go 1.25.0`. Plan r5 carries each of those constraints without widening the fence or selecting a new mechanism.

The prior contradiction is therefore closed by authority rather than locally waived. T1b remains one honest commit containing the real `internal/connector/nfccheck` consumer, the ruled `go.mod` bytes, and tidy-generated `go.sum`; the dependency is neither retained by a shim nor downgraded. T1a remains committed at `3257ec35f1b82874c2b0fa382c17d5e71107307e`. T1c, T2–T15, §2a, the seventeen-task boundary, full-battery-at-every-commit, deterministic restack, end-review, and operator MERGE-GATE carry unchanged from r4/r3/r2.

The isolated code worktree was rechecked clean at `s14-m8-connector@3257ec35f1b82874c2b0fa382c17d5e71107307e`. No implementation authority is present in this PLAN relay, so no source, module, test, stage, or commit action resumed. Implementation may resume only under a successor `PHASE: IMPL` relay containing the valid bare own-line `DISPATCH IMPL`, addressed solely to `s14.implementer`, and parented by `s14-build-plan-review-5`.

## Boundary contract review

Writes: at T1b, `internal/connector/nfccheck/**`, exact ruled `go.mod`, tidy-generated `go.sum`, and pair-local evidence; later writes remain exactly within the carried plan fence.
Reads: x/text v0.41.0 plus the frozen m-8/m-3 contracts.
Target entity: the connector's check-only artifact-NFC predicate and reproducible module graph.
Downstream consumer: the T2 catalog validator and T3 policy validator.
Contract: `norm.NFC.IsNormalString` rejects non-NFC only at schema loci requiring NFC; it never rewrites bytes or becomes a global JCS rule.
Proof: RED-first predicate tests; tidy no-op; absence of a `toolchain` line; exact `go 1.25.0`; `go build ./...`; `go test -mod=readonly ./...`; `go vet ./...`; and the full existing battery on the same committed tree.
No-consumer action: stop before commit on any tidy divergence, retention shim, second x/text version, or expanded normalization locus, and route the exact bytes upward.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the wider connector slice retains provider credential and authorization boundaries
- migration/backfill/destructive-write/canonical-data-repair: no — no migration, destructive write, or canonical-data repair is added
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the wider connector slice retains trust-critical provider-send authority
- AI-or-automation-acts-downstream: yes — the connector ultimately sends model requests
- worker/scheduler/queue/retry/async-side-effect: yes — the wider slice retains streaming, cancellation, and one-shot effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the ruled shared Go directive and connector contracts are integration seams
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is added
- test-runtime-role-mismatch: no — the predicate and module graph are tested in their production Go package/runtime
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — the successor narrows one directive spelling and retains every gate
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk waiver is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane-governance action only — this PLAN-REVIEW file plus one live-EOF s14 INDEX row; code worktree remains clean at `s14-m8-connector@3257ec35f1b82874c2b0fa382c17d5e71107307e`; no source, test, dependency, branch, stage, commit, push, PR, merge, or publication action
FINAL_GIT_STATUS_SHORT:
 M .relays/s13/INDEX.md
 M .relays/s14/INDEX.md
 M docs/sprints/active/2026-08-20-s13-m10-module/plans/PL-s13-build-plan-20260820.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260821-034045.md
?? .relays/s13/s13-build-plan/PLAN-planner-20260821-035127.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-040009.md
?? .relays/s14/s14-build/PLAN-planner-20260821-035141.md
Next requested action: `s14.planner` issues the successor implementation dispatch parented by `s14-build-plan-review-5`; hand-relay that exact path to `s14.implementer`, which may then resume T1b from clean `3257ec35…`.
