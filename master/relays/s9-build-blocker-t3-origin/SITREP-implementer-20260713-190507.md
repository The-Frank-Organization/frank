# SITREP — T3 HOLD: the approved total origin table omits landed read-file refusal tuples

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s9-build-blocker-t3-origin
PARENT_DISPATCH_ID: s9-build-token
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — planner/guide correction is required before T3 resumes; merge remains operator-only
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan-rev11-s9.1
IN_REPLY_TO: master/relays/s9-build-token/IMPL-planner-20260713-183600.md
FROM: s9.implementer
TO: s9.planner
BUNDLE_ID: m-3-observation-evidence
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-7.planner, m-2.planner, m-1.planner
SUBJECT: T1 and T2 landed green; T3 stopped before commit because the plan's supposedly total origin-by-tuple table has no legal row for three extant non-machinery read-file refusals; choosing a row locally would either amend the plan or change terminals

## Summary

T1 landed at `ba26b27`; T2 landed at `5f6a7ec`. Both task commit points passed their focused/race batteries and full `go test ./... -count=1` batteries. T3's RED-first hostile matrix was then implemented far enough to pass its focused unit battery, but the first wider s8 regression exposed a plan-level omission. T3 is **HELD with no T3 commit**.

## The missing row

The plan calls its table at plan lines 157–163 the **TOTAL `(origin × tuple)` table**. Its `base-check` row permits only:

- `pass/pass/entry-rung`;
- `fail/fail/none`;
- `skipped/degraded/none`.

The landed read-file mechanism has a fourth, deliberate behavior family:

- `not-regular-file`;
- `read-size-exceeded`;
- `read-deadline-exceeded`.

Each is emitted as `readFileResultRefused` (`internal/observe/read_file_worker.go:94,113,116,123,127,136,139`) and converted by `runReadFile` to `refusedVerdictWithDetail` (`internal/observe/checks_base.go:120-131`): `unsafe/blocked/none`, **MachineryFault false**. The landed terminal contract is load-bearing: the s8 regression explicitly requires authority → `held` and non-authority → `accepted` through decision-② for `not-regular-file` (`test/fixtures/s8_decision2_test.go:124-168`).

No existing origin row accepts that tuple/class combination:

- `base-check` omits `unsafe/blocked/none`;
- `conductor-policy` is explicitly pre-spawn and its closed family omits these post-open read refusals;
- `conductor-machinery` would set `MachineryFault: true`, changing non-authority from accepted to rejected and changing the semantic class.

Therefore a build seat cannot make the table total without choosing new design bytes.

## E2 reproduction

Command:

```text
go test ./test/fixtures -run 'Test(S8VerdictBinding|S8ExecutableClaimAggregation|S8ProductionDogfood|S8AdversarialTruncation|S8Decision2|S8ReadFile|S10Suite|S10E1)' -count=1 -timeout 240s
```

Observed first dispositive failure:

```text
TestS8Decision2E1MachineryFaultsAndObservedAbsence/read_file_IO_error
want non-authority terminal accepted / observation-unavailable
got held/rejected machinery-edge behavior after the omitted tuple collapsed to
check-machinery-verdict-tuple-invalid
```

The production dogfood suite then rejected because its nested battery observed the same regression. This is not an isolated test expectation.

## Requested ruling

Please route one byte-exact correction through the m-3 guide and return it directly to `s9.implementer`:

1. Recommended: extend `base-check` with a closed `base-refusal` row `unsafe/blocked/none`, details exactly `{not-regular-file, read-size-exceeded, read-deadline-exceeded}`, `MachineryFault:false`, conductor-produced `under-timeout`, preserving decision-② terminals; or
2. explicitly place those details in another origin family with the same non-machinery semantics; or
3. if they are intentionally reclassified as machinery, issue the required locked-contract amendment and accept the terminal change explicitly.

The correction must also say whether the executor-origin validator may normalize empty timing in legacy/fake executors or whether the regression fakes must be updated to supply their required host timing. Production `executor.Host` already supplies the closed timing values; this second point is not the main blocker.

## Scope and state

No `executor.go`, config, main, fieldspec, T7/T8, or blocked-ledger byte was touched. The m-7 owner-byte return at `s9-build-fidelity-m7/SITREP-planner-20260713-184136.md` is addressed to `s9.planner`; this seat remains CC-only and has not carried it. T4 therefore also remains unopened here pending planner routing after T3.

ACTIONS_GIT_REF: `s9-evidence-thicken@5f6a7ec`; landed T1 `ba26b27` and T2 `5f6a7ec`; T3 has no commit; uncommitted T3 RED/partial implementation paths are `internal/observe/gate.go`, `internal/observe/registry.go`, `internal/observe/verdict_binding.go`, `internal/observe/verdict_binding_test.go`
FINAL_GIT_STATUS_SHORT:
 M internal/observe/gate.go
 M internal/observe/registry.go
?? internal/observe/verdict_binding.go
?? internal/observe/verdict_binding_test.go

Next requested action: s9.planner returns a directly addressed, guide-backed byte ruling for the missing read-refusal origin/tuple row (and the executor empty-timing compatibility point). T3 resumes from its uncommitted RED/partial state; no merge authority is requested or implied.
