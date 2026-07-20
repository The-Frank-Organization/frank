## SITREP — s10 T1 RED proven; implementation holds for the owner-authored v8 registry and capability bytes

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s10-build-t1-review
PARENT_DISPATCH_ID: s10-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this hold needs the already-assigned m-2/m-7 owner returns; merge remains separately operator-gated
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
IN_REPLY_TO: s10-build-impl/IMPL-planner-20260712-232010.md
FROM: s10.implementer
TO: s10.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: T1 RED is E2-proven on the authorized branch; T1 GREEN is blocked because the live relay trail contains neither m-2's final `odb`/`resummon_command` registry rows + required-field/seat-scope calls + v7-to-v8 successor bytes nor m-7's v8 capability-marker byte/confirm; condition (f) hold, no owner bytes adapted

Summary: Began T1 under the valid delegated token. Added only the RED-first old-reader phase-0 fixture required by the adopted T1 plan. The fixture fails for the intended missing-feature reason: the shipped registry marker remains `s8-fieldspec-v7`, while the test requires `s10-fieldspec-v8`. No production or registry bytes were written.

Plan lock: `s10-comms-spine-plan-r2-s10.2`; T1 only. T9 and T10 remain closed. T2 through T8 and T11 preparation have not started.

SCOPE_DIFF:
- frank/test/fixtures/s10_fieldspec_test.go -> in
SCOPE_DIFF_RESULT: all-in

Files changed:
- `frank/test/fixtures/s10_fieldspec_test.go` — T1 old-reader marker-first refusal fixture, within the authorized `frank/test/fixtures/` row and the adopted T1 RED requirement.
- `frank/.relays/s10/t1-red-20260712-225520.txt` — file-captured RED output, within the authorized slice-relay row.

ACTIONS_GIT_REF: `s10-comms-spine@8941889a432a0a02a4183a16fb0015c230c9561b`; uncommitted T1 RED fixture; `git status --short` = `?? test/fixtures/s10_fieldspec_test.go`; no task commit because T1 is incomplete

Acceptance criteria status:
- T1 Step 1 RED: met at E2; failure is the missing v8 marker, not a compile/setup error.
- T1 Steps 2-5: blocked on owner bytes/confirm.
- T1/T2 owner-fidelity-before-T4 gate: still closed; no park fixture opened.

Boundary contract proof:
- Writes so far: test-only expectation for the governed v8 marker/refusal behavior.
- Reads: shipped embedded registry and `config.ValidateFieldspecReaderMarker`.
- Downstream consumer: the T1 transition and fresh-v8 genesis remain unmodified.
- No-consumer action: hold rather than invent m-2/m-7-owned bytes.

Tests/verification:
- Baseline before edits: `go test ./...` — green.
- RED: `go test ./test/fixtures -run '^TestS10V7ReaderRefusesV8MarkerBeforeContent$' -count=1` — expected FAIL: registry version `s8-fieldspec-v7`, want `s10-fieldspec-v8`.
- Capture: `frank/.relays/s10/t1-red-20260712-225520.txt`.
- `git diff --check` — clean.

Evidence levels: baseline and RED command output are E2; missing owner relays and current source/version marker are E1.

Out-of-scope preserved: no production code, registry, successor map, capability set, genesis, park/wake, executor, composition, T9, T10, merge, push, or PR action. No owner-owned byte was inferred or adapted.

Remaining risk: T1 cannot be completed faithfully until the exact owner returns arrive. The plan explicitly leaves the final per-token bytes and required-field call to m-2 and the capability marker to m-7; choosing them here would violate the owner-fidelity gate and condition (c).

Blocked: yes — owner fidelity input absent after refreshing `master/relays/INDEX.md` and the live s10 relay tree.

Next requested action: route a T1 fidelity request to m-2.planner for the exact v8 registry transition rows, per-kind required fields, conductor seat-scope, and v7-to-v8 successor byte; route the v8 capability marker to m-7.planner. Return their owner-authored bytes/confirm to this seat. On receipt I will complete T1 GREEN, run the package/fixture/full batteries, commit exactly T1, and request task review. T4 remains hard-held until both owner confirms are in hand.
