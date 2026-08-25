## PLAN-REVIEW — s14 build plan r6 bounded R8 fixture-capacity fold: APPROVE; T9.5 carries exactly the one-file 600-second/discriminating-detail owner repair, its standalone pre-T10 commit and three-run GREEN obligation, with every fixture assertion and all other r5 scope unchanged

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s14-build-plan-review-6
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this approves only the bounded R8 fold; implementation still requires a successor dispatch and merge remains operator-gated
GRILL_REQUIRED: no — master R8 already ruled the bounded owner repair; this review checks exact carriage and executability without redesign
FILED_AT_LOCAL: 20260821-131510
IN_REPLY_TO: frank/.relays/s14/s14-build/PLAN-planner-20260821-131336.md
PLAN_LOCK_ID: s14-build @ sha256 0e44d5ea23d685429d624b4b2cc93b1a022a0866b84285960d47f718d6a4b5ac
PLAN_REVIEW_VERDICT: approve
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: approve plan r6 — exact R8 one-file fixture repair, standalone T9.5 commit, restored T10 gate, and unchanged remaining plan

## Verdict

`PLAN_REVIEW_VERDICT: approve` for exact plan r6 SHA-256 `0e44d5ea23d685429d624b4b2cc93b1a022a0866b84285960d47f718d6a4b5ac`.

The bounded fold matches master R8 SHA-256 `481cd51bbf0642141246edcd5ab311e58dccd5e5418297fbdd068afee146fe11` on every operative condition:

1. T9.5 changes the exact deadline byte in `test/fixtures/s8_exit_gate_test.go` from `150*time.Second` to `600*time.Second`.
2. T9.5 adds only the ruled discrimination: deadline exhaustion identifies `ctx.Err()` plus elapsed time, and the false-done mismatch identifies the nested suite's observed color plus elapsed time. No assertion is removed or weakened; every Accepted/Rejected expectation remains byte-unchanged.
3. The one-file grant is scoped to T9.5 only. The existing nested-scope guard and outer-reachability assertion remain byte-untouched, and the file returns to consume-only immediately after the repair commit.
4. The three failures already recorded in `SITREP-implementer-20260821-061012.md` are the RED evidence. GREEN requires two consecutive full batteries plus one isolated repaired-fixture run before T10 may commit.
5. T9.5 is its own cleanly liftable R8-citing commit, sequenced at committed T9 HEAD `471532a151762d496186f2daee70b21fdd71062f` before the preserved T10 bytes. T10 then commits only under the restored gate.

The live source worktree confirms the executable starting state: branch `s14-m8-connector` is at `471532a151762d496186f2daee70b21fdd71062f`; the target fixture still carries `context.WithTimeout(context.Background(), 150*time.Second)` at line 130 and the unchanged guard/reachability calls at lines 124/127; only the preserved in-fence `internal/connector/stream/**` T10 tree is untracked. Plan r6 adds no other source or governance surface. T1a–T9, T10–T15, §2a, acceptance, escalation rails, the R6 module-directive law, end-review, restack, and operator MERGE-GATE carry unchanged from r5/r4/r3/r2.

Run 2's R8 regrade is correctly informational: the fixture loudly detected that its nested suite was red, so no conductor defect is opened. The R-S12 recurrence tripwire remains named and not fired. Domain-owner objection timing and queue row 11 remain master's process, not a pair-local prerequisite or implementation claim.

No implementation authority appears in this PLAN relay. No source, test, dependency, stage, commit, or branch mutation occurred during review. Implementation may resume only under a successor `PHASE: IMPL` relay containing the valid bare own-line token, addressed solely to `s14.implementer`, and parented by `s14-build-plan-review-6`.

## Boundary contract review

Writes: T9.5 writes only `test/fixtures/s8_exit_gate_test.go` under the R8 one-commit grant; afterward the carried plan writes only the existing connector and lane-owned evidence surfaces.
Reads: the existing S8 production-dogfood fixture semantics, the `1ac14a0` nested-scope/reachability remedy, the recorded T10 RED evidence, and the frozen connector contracts.
Target entity: the production-dogfood fixture's hang backstop and failure diagnostics, followed by the unchanged T10 normalized provider-event implementation.
Downstream consumer: every per-commit serial full battery on s14 and, by byte-identical lift only, sibling slice batteries/restack.
Contract: 600 seconds is a hang backstop; capacity and semantic failures are distinguishable; all Accepted/Rejected expectations and nested-scope/reachability semantics remain unchanged; T9.5 is standalone and precedes T10.
Proof: exact one-file diff; isolated repaired fixture GREEN once; two consecutive `go test -mod=readonly -p=1 -count=1 ./...` batteries GREEN; build/vet/diff/module tripwires; commit ancestry T9.5 before T10.
No-consumer action: stop before commit on any extra file, changed assertion, changed guard/reachability byte, non-600-second deadline, missing discrimination field, combined T9.5/T10 commit, or failed required battery, and route the exact deviation upward.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the carried connector slice retains the provider credential and authorization boundary
- migration/backfill/destructive-write/canonical-data-repair: no — T9.5 and the carried connector work add no migration, destructive write, or canonical-data repair
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the carried connector slice retains trust-critical provider-send authority
- AI-or-automation-acts-downstream: yes — the connector ultimately sends model requests and normalizes their stream
- worker/scheduler/queue/retry/async-side-effect: yes — the carried slice retains streaming, cancellation, transport, and one-shot side effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the fixture is a shared conductor-tree gate and its exact repair may be lifted across sibling slices
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is added by this fold
- test-runtime-role-mismatch: yes — T9.5 specifically repairs an exhausted nested production-dogfood test-runtime capacity boundary
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — R8 grants exactly one file for one commit and preserves all semantic assertions and terminal gates
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade, failed-test acceptance, or risk waiver is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane-governance action only — this PLAN-REVIEW file plus one live-EOF s14 INDEX row; source worktree remains `s14-m8-connector@471532a151762d496186f2daee70b21fdd71062f` with the inherited preserved `?? internal/connector/stream/`; no source, test, dependency, branch, stage, commit, push, PR, merge, or publication action
FINAL_GIT_STATUS_SHORT:
 M .relays/s14/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-131510.md
Next requested action: `s14.planner` issues the successor implementation dispatch parented by `s14-build-plan-review-6`; hand-relay that exact path to `s14.implementer`, which may then execute T9.5 from T9 HEAD `471532a…`, prove the ruled three-run GREEN, commit T9.5, and resume T10.
