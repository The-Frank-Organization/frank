## IMPL blocker - r2 exposed non-inventoried main_assembly owed posture change

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s5-a-impl-r2-main-assembly-blocker
PARENT_DISPATCH_ID: s5-a-impl-r2
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - scope/class ruling required before continuing
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
BRANCH: s5-a-registry
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer, operator
IN_REPLY_TO: .relays/s5/s5-a-impl-r2/IMPL-planner-20260706-083119.md
SUBJECT: stopped under r2 fence; class-(a)/(c)/(d) edits applied, but `main_assembly_test.go` now needs a class-(b) owed/record_kind expectation change although the addendum tagged that file as class (a) only

Current state:
- r2 dispatch linted clean and carried the live own-line implementation token at line 50.
- Worktree: `~/frank-s5-team/s5-a`, branch `s5-a-registry`, still uncommitted.
- I applied the class-(a) `EVIDENCE_TARGET` additions, the F11 mutation-helper required-field move, and the named owed non-operator inversion/rename in the implementation worktree.
- I did not commit.

Focused evidence before stop:
- `go test -count=1 ./cmd/frank-mcp` -> PASS.
- `go test -count=1 ./internal/engine` -> PASS.
- `go test -count=1 ./internal/obligation` -> PASS.
- `go test -count=1 ./test/fixtures -run 'TestF11|TestS2|TestFrankBinaryAssembles|TestFrankBinaryOperatorChannelO3|TestFrankBinaryReadCorruption|TestRunningConfig|TestConfigChange|TestRestartWithNewRegistry|TestNudge|TestCCRecipient|TestOfflineRecipient'` -> FAIL only at `TestFrankBinaryReadCorruptionQueuesLiveQuarantine`.

Remaining failure:
- `test/fixtures/main_assembly_test.go:622`: `post-quarantine disposition = {State:rejected RelayID:...}, want accepted`.

Why I stopped:
- The remaining test constructs the post-quarantine owed leg with `meta := seat.SeatMeta{Name: "seat-a", Role: "implementer"}` at `test/fixtures/main_assembly_test.go:555`.
- It then submits `record_kind: owed_item` and `record_kind: owed_disposition` through that same non-operator client at `test/fixtures/main_assembly_test.go:598-619`.
- The S5 registry WIP now defines `record_kind` scope as operator-only for `owed_item` and `owed_disposition`, while `"*"` has only `diagnostics` (`internal/fieldspec/registry.json:125`).
- Updating this test to use an operator channel or to assert the non-operator bounce is a class-(b) owed/record_kind posture update, not merely class (a).
- The r2 dispatch says the eight inventoried files admit only their tagged grant classes. The addendum tags `test/fixtures/main_assembly_test.go` as class (a), not class (b). I therefore stopped instead of broadening the edit.

Decision needed:
1. Amend the inventory/addendum/r2 fence to admit class (b) for `test/fixtures/main_assembly_test.go`, then I can update the quarantine owed leg to the settled operator-only posture.
2. Or give a narrower direction for this test that remains class (a) only.
3. Or rule that this is a fresh escalation outside r2 and should be handled by a later integration/fold lane.

ACTIONS_GIT_REF: WIP only in `~/frank-s5-team/s5-a`; no commit; primary checkout source drift from my earlier patch-target mistake was reversed before this blocker; current implementation worktree diffstat is 10 tracked files changed plus new `test/fixtures/s5_registry_dormancy_test.go`
FINAL_GIT_STATUS_SHORT:
 M cmd/frank-mcp/mcp_test.go
 M internal/engine/config_change_test.go
 M internal/engine/pipeline_test.go
 M internal/fieldspec/registry.json
 M internal/fieldspec/registry_test.go
 M internal/obligation/owed_test.go
 M test/fixtures/f11_test.go
 M test/fixtures/main_assembly_test.go
 M test/fixtures/s4_config_change_test.go
 M test/fixtures/s4_shim_test.go
?? test/fixtures/s5_registry_dormancy_test.go
