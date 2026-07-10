## IMPL report - s5-a registry pass r2

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s5-a-impl-r2
PARENT_DISPATCH_ID: s5-a-impl-r2
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
BRANCH: s5-a-registry
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: implementation report for s5-a-registry-plan r2; 084234 main-assembly delta folded
ACTIONS_GIT_REF: s5-a-registry@dd8189d
FINAL_GIT_STATUS_SHORT: none - clean tree

Summary:
- Implemented the s5-a registry pass under the r2 implementation dispatch.
- Folded the `20260706-084234` planner ruling: `test/fixtures/main_assembly_test.go` is treated as approved `(a)+(b)` scope, with the quarantine owed-item/disposition leg submitted through the operator channel.
- No merge, PR, push, or cleanup action was taken.

PR:
- none; merge/PR was not authorized by this relay.

Plan lock:
- DESIGN_LOCK_ID: s5-a-registry-design
- PLAN_LOCK_ID: s5-a-registry-plan
- Dispatch relay: `.relays/s5/s5-a-impl-r2/IMPL-planner-20260706-083119.md`
- Delta ruling relay: `.relays/s5/s5-a-impl-r2/PLAN-planner-20260706-084234.md`

Files changed:
- `cmd/frank-mcp/mcp_test.go`
- `internal/engine/config_change_test.go`
- `internal/engine/pipeline_test.go`
- `internal/fieldspec/registry.json`
- `internal/fieldspec/registry_test.go`
- `internal/obligation/owed_test.go`
- `test/fixtures/f11_test.go`
- `test/fixtures/main_assembly_test.go`
- `test/fixtures/s4_config_change_test.go`
- `test/fixtures/s4_shim_test.go`
- `test/fixtures/s5_registry_dormancy_test.go`

Commit:
- `dd8189d` s5-a: land registry pass
- Diffstat: 11 files changed, 574 insertions, 85 deletions.

Acceptance criteria status:
- Registry payload: version is `s5-fieldspec-v3`; field count is 83; named enum count is 24.
- `gate_category_A` includes `routing_escalation`; combined `gate_category` keeps A/B order and appends `routing_escalation`, `other`.
- `EVIDENCE_TARGET` is required for every phase using the empty `phase_in` negation shape.
- Observe action rows `ACTIONS_GIT_REF` and `FINAL_GIT_STATUS_SHORT` are observe-visible and observe-required.
- `record_kind` is seat-scoped to operator for `owed_item`, `owed_disposition`, `gate_resolution`, `disposition`, `diagnostics`, and `config_change`; all other seats keep `diagnostics`.
- `resolves_gate` is visible only to operator seat/role and required for `gate_resolution`.
- Added s5 registry dormancy fixture coverage for observe-only rows, OI-S4 record-kind posture, operator-positive/non-operator-negative `resolves_gate`, EVIDENCE_TARGET requiredness, model-layer rejection, registry mutation rejection, predicate controls, and required raw annotations.

Grant fold status:
- Class `(a)` settled-contract updates were folded across MCP/config/pipeline/owed/F11/main-assembly/s4 fixtures by supplying `EVIDENCE_TARGET`.
- Class `(b)` was folded only for the `main_assembly` post-quarantine owed leg per the 084234 ruling, by switching owed-item/disposition submissions to the operator channel while preserving the downstream assertions.
- Class `(c)` F11 crash/applicability helper moves were folded by supplying the new required field in `submitPayloadForRegistry`.
- Class `(d)` OI-S4 posture was folded by renaming and inverting the non-operator owed-item test to assert rejection.

Boundary contract proof:
- Production mechanism changes are limited to `internal/fieldspec/registry.json`.
- All Go source edits outside the registry payload are tests or fixtures needed to keep existing batteries aligned with the new registry contract.
- No s5-b surface, validator mechanism, engine mechanism, migration, lineage, store, or command behavior implementation was changed.
- The implementation remains on the local `s5-a-registry` branch; no merge to protected or feature branches was attempted.

Tests/verification:
- Focused post-delta quarantine check before commit: `go test -count=1 ./test/fixtures -run TestFrankBinaryReadCorruptionQueuesLiveQuarantine -v` exit 0.
- Focused fixture subset before commit: `go test -count=1 ./test/fixtures -run 'TestF11|TestS2|TestFrankBinaryAssembles|TestFrankBinaryOperatorChannelO3|TestFrankBinaryReadCorruption|TestRunningConfig|TestConfigChange|TestRestartWithNewRegistry|TestNudge|TestCCRecipient|TestOfflineRecipient'` exit 0.
- Final fresh battery at `dd8189d`: `go test -count=1 ./...` exit 0.
- Final static check at `dd8189d`: `go vet ./...` exit 0.
- Whitespace check at `dd8189d`: `git diff --check` exit 0.
- Payload probe at `dd8189d` confirmed `version s5-fieldspec-v3`, `field_count 83`, `named_enum_count 24`, expected gate category arrays, operator record-kind seat scope, EVIDENCE_TARGET requiredness, `resolves_gate` operator visibility, and observe action/status visibility.

Evidence levels:
- E1: file diff, commit, and registry payload inventory above.
- E2: full local test/vet battery plus focused fixture and registry dormancy coverage.
- E3: not claimed.
- E4: not claimed.

Honesty / workspace notes:
- A patch-target slip briefly touched the primary checkout earlier in the run; it was reversed before the worktree commit. Primary checkout tracked source status is unchanged; it still shows only untracked sprint docs under `docs/sprints/2026-07-06-s5-slice-5/`.
- One supplemental payload probe initially used the wrong JSON field key and failed locally. The checked-in tests had already passed; the probe was corrected against the actual `named_enums`/`id` schema and rerun successfully before this report.

Out-of-scope preserved:
- No implementation of s5-b validator, egress, detector, replay, or away-mode mechanisms.
- No production Go mechanism changes outside the registry payload.
- No live/deploy verification and no E3/E4 claim.
- No PR, push, merge, or cleanup action.

Remaining risk:
- Integration remains dependent on the planner/orchestrator sequence that lets s5-b consume the s5-a registry payload.
- Dormancy assertions remain bounded to local E2 evidence; submit-path non-lane-writability remains the s5-b guard surface.

Next requested action:
- Planner review of this implementation report and downstream orchestration of the s5-a/s5-b integration gate.
