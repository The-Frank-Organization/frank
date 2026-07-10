## IMPL blocker — s5-a registry pass contract conflict at resolves_gate render dormancy

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s5-a-impl-blocker
PARENT_DISPATCH_ID: s5-a-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - planner contract ruling required before further implementation
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer, operator
IN_REPLY_TO: .relays/s5/s5-a-impl/IMPL-planner-20260706-065734.md
SUBJECT: blocked during TDD green pass: locked resolves_gate row shape renders on Step-1 forms, but the locked 38-name VP-W3 sweep requires it absent from every rendered form

Status: blocked before commit. I did not self-author a behavior change because the implementation dispatch says "Any OUT-file need, collision, or contract question: STOP and relay to me."

Work completed before the blocker:
- Validated the IMPL dispatch and baseline worktree.
- Ran baseline `go test ./...` green in `~/frank-s5-team/s5-a`.
- Added the planned red assertions in `internal/fieldspec/registry_test.go` and new `test/fixtures/s5_registry_dormancy_test.go`.
- Confirmed the tests fail red against the unedited registry.
- Applied the atomic registry content pass far enough that `go test ./internal/fieldspec` is green.

Failing evidence:
- Command: `go test ./test/fixtures -run S5Registry`
- Result: FAIL.
- Reproducible first failure class: `TestS5RegistryDormancySweepAndRecordKindScope/...: resolves_gate rendered unexpectedly`.

Root-cause evidence:
- Locked design §6 says `resolves_gate` is `id_ref`, owner `free_text`, fill `free_text`, with only `required_when: {"all_of":[{"record_kind_in":["gate_resolution"]}]}` and no `visible_when`.
- The live renderer renders free-text fields unless they are system/computed/system_only/computed_result or their `visible_when` is false (`internal/fieldspec/render.go:50-57`).
- The locked design §7 and locked plan step 1b require the 38-name VP-W3 sweep to assert `form.HasField(f) == false` for all 36 new rows plus ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT. The design explicitly counts Block D as `resolves_gate + gate_category_pick`.
- Therefore the locked row shape and locked fixture acceptance are incompatible for `resolves_gate`: with the specified row data, `Render` correctly exposes `resolves_gate` on every Step-1 form.

Decision needed:
1. Keep `resolves_gate` lane-fillable per design §6. If so, update the 38-name sweep contract to exclude `resolves_gate` and acknowledge that Block D is not fully render-dormant.
2. Keep the 38-name sweep strict. If so, issue an explicit design/plan clarification adding a `visible_when`/step-gate to `resolves_gate` and state how a gate_resolution author is expected to supply the referenced gate id without production Go changes.

Current worktree state:
- Branch/worktree: `~/frank-s5-team/s5-a`, branch `s5-a-registry`.
- No commit made.
- WIP diffstat from tracked files: `internal/fieldspec/registry.json | 74 ++++++++++++++++++++++----`; `internal/fieldspec/registry_test.go | 103 ++++++++++++++++++++++++++++++++++--`.
- New untracked file: `test/fixtures/s5_registry_dormancy_test.go`.

Verification already run:
- `go test ./...` at baseline before edits: PASS.
- `go test ./internal/fieldspec` after registry pass: PASS.
- `go test ./test/fixtures -run S5Registry` after registry pass: FAIL as above.

ACTIONS_GIT_REF: WIP only; no commit; worktree dirty at `~/frank-s5-team/s5-a` with tracked registry/test edits plus new fixture file
FINAL_GIT_STATUS_SHORT:
 M internal/fieldspec/registry.json
 M internal/fieldspec/registry_test.go
?? test/fixtures/s5_registry_dormancy_test.go

Next requested action: planner ruling on the `resolves_gate` contract, then I can either adjust the fixture expectation or fold the clarified row gating and continue the existing WIP through full verification and commit.
