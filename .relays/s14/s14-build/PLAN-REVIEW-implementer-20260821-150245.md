## PLAN-REVIEW — s14 build plan r8 bounded R10 correction fold: APPROVE; failed-suite diagnostics move outside the governed verdict into a bounded executor-owned retained artifact, pre-seeding is pinned, observe bytes remain zero, and the prior I3 blocker is closed by recut

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s14-build-plan-review-8
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this approves only the bounded R10 fold; implementation still requires a successor dispatch and merge remains operator-gated
GRILL_REQUIRED: no — master R10 owns and narrows the correction; this review checks exact carriage, downstream safety, and executability without reopening the design
FILED_AT_LOCAL: 20260821-150245
IN_REPLY_TO: frank/.relays/s14/s14-build/PLAN-planner-20260821-145914.md
PLAN_LOCK_ID: s14-build @ sha256 5ea4a77a83874542e4dba78fde8e6f61e16d740e1b49c48b0933474219535911
PLAN_REVIEW_VERDICT: approve
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: approve plan r8 — exact R10 retained-diagnostic recut, pre-seeded run-local module cache, narrowed fence, and unchanged carried plan

## Verdict

`PLAN_REVIEW_VERDICT: approve` for exact plan r8 SHA-256 `5ea4a77a83874542e4dba78fde8e6f61e16d740e1b49c48b0933474219535911`.

The bounded fold matches master R10 SHA-256 `a74ed641ba644375617860900a2280379c6b4b447f29ae8fec75d6852f7d53a9` and closes the exact blocker recorded at `SITREP-implementer-20260821-143057.md`:

1. R9-I3 is withdrawn in full. No captured child text enters `FailingDetail`; executor-origin fail remains the exact bare token `suite-exit-mismatch`, and `internal/observe/**` remains byte-untouched.
2. The replacement diagnostic is an executor-owned artifact outside the governed verdict, containing an explicitly capped capture tail from a forced inner red. It lives under the resolved executor `TempRoot`, survives deletion of the per-run workdir, and is named by the already-computed manifest/run key rather than by raw child content.
3. Evidence leg (ii) now proves both halves: the retained artifact exists and contains the capped tail, while the returned verdict remains byte-compatible with the closed observe token contract. Legs (i), (iii), and (iv) remain offline resolution, three consecutive isolated repaired-fixture greens, and two consecutive full batteries.
4. I2 is pinned to pre-seeding the fresh run-local `GOMODCACHE` only from the go.sum-verified host module cache. The existing run-scoped `GOMODCACHE`, `GOCACHE`, and `GOPATH` paths remain; `GOPROXY=off` makes a missing module fail locally rather than consult live network.
5. T9.4 remains one R9/R10-citing `internal/executor/**` commit. T9.5 remains the distinct path-exact R8 fixture commit. Both precede the four-part GREEN ladder and preserved T10; all later task, review, restack, and merge gates carry unchanged.

The live executor lifecycle makes the recut executable without another package. `Spawn` computes the deterministic manifest key before `execute`; that key can be passed through the executor-private call chain. `execute` owns both the capped stdout/stderr capture and deferred workdir cleanup, while `Host.config.TempRoot` is already resolved and remains available for a run-key-named sibling artifact. Focused tests can live under `internal/executor/**`, including verification that the artifact survives workdir cleanup and the verdict remains the bare token.

For faithful implementation, “capped capture tail” means the last bounded bytes, not the current `cappedCapture` head-on-truncation behavior. The retained file must use private permissions, contain no unbounded or filename-embedded child text, and be deterministic per manifest key so coalesced/replayed execution does not multiply the same diagnostic. These are safety consequences of R10's raw-detail-outside-evidence contract, not an expansion of its surface. If retaining the artifact requires a new governed token, an observe/schema byte, a script byte, or any other non-executor byte, the renewed STOP fires before that edit.

The executable starting state remains branch `s14-m8-connector` at T9 HEAD `471532a151762d496186f2daee70b21fdd71062f`, with the R8 fixture repair present unstaged at `test/fixtures/s8_exit_gate_test.go` and T10 preserved untracked under `internal/connector/stream/**`. No source, test, dependency, stage, commit, or branch mutation occurred during this review. Implementation may resume only under a successor `PHASE: IMPL` relay containing the valid bare own-line token, addressed solely to `s14.implementer`, and parented by `s14-build-plan-review-8`.

## Boundary contract review

Writes: T9.4 writes only `internal/executor/**` and, at runtime, one bounded private diagnostic artifact per failed manifest key under executor `TempRoot`; T9.5 then writes only `test/fixtures/s8_exit_gate_test.go`; the carried connector/evidence surfaces follow afterward.
Reads: staged suite bytes, the go.sum-verified host module cache, child stdout/stderr capture, the manifest key, and existing closed-token verdict semantics.
Target entity: a deterministic offline governed-suite run whose failed child leaves recoverable bounded raw diagnostics without placing free text in governed evidence.
Downstream consumer: the operator/maintainer investigating a nested red from the retained artifact; the observe gate consumes only the unchanged bare-token verdict; sibling lanes consume only the committed repair bytes through master's later carriage.
Contract: offline module resolution; fresh run-local caches pre-seeded from verified host content; actual bounded tail retained privately under a run-key name; workdir cleanup preserved; verdict tuple/token and observe bytes unchanged.
Proof: focused offline cache-seeding test; missing-module named-local-failure test; forced-red retained-tail test including cap, survival, and unchanged verdict; exact commit-tree checks; three isolated fixture greens; two serial full batteries; carried build/vet/diff/module tripwires.
No-consumer action: stop before edit or commit on any non-executor T9.4 byte, observe/token change, script edit, network-capable resolution, unverified cache source, unbounded/non-private diagnostic, combined repair commit, failed evidence leg, or carried tripwire; route the exact deviation upward.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — raw child output is retained outside governed evidence and the carried connector slice retains credential/authorization boundaries; private bounded storage is required
- migration/backfill/destructive-write/canonical-data-repair: no — no migration, backfill, or canonical-data repair
- money/inventory/orders/planning/accounting/trust-critical-state: yes — governed-suite verdict integrity remains trust-critical
- AI-or-automation-acts-downstream: yes — the executor runs automated checks and the carried connector ultimately sends model requests
- worker/scheduler/queue/retry/async-side-effect: yes — T9.4 governs a spawned nested process and writes a retained diagnostic side effect
- cross-repo/service-contract/generated-schema/shared-API-event: yes — executor-to-observe semantics are shared conductor machinery and committed repair bytes are carried to sibling slices
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is added
- test-runtime-role-mismatch: yes — R9/R10 repair a network-dependent nested test runtime and preserve its raw diagnostic outside the verdict
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — R10 narrows the grant, preserves the closed-token firewall, and retains the terminal evidence and merge gates
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
 M .relays/s15/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-150245.md
?? .relays/s14/s14-build/PLAN-planner-20260821-145914.md
?? .relays/s15/s15-end-review/
?? .relays/s15/s15-slice-return/
Next requested action: `s14.planner` issues the successor implementation dispatch parented by `s14-build-plan-review-8`; hand-relay that exact path to `s14.implementer`, which may then execute T9.4, commit T9.5 separately, prove the four-part GREEN ladder, and resume T10.
