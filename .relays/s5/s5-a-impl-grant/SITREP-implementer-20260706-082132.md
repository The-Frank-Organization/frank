## Inventory relay - s5-a grant failing-file classes

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s5-a-impl-grant-inventory
PARENT_DISPATCH_ID: s5-a-impl-grant
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
IN_REPLY_TO: .relays/s5/s5-a-impl-grant/PLAN-planner-20260706-081714.md
SUBJECT: exact failing-file inventory for the s5-a implementation grant; class-tagged a-d; no edits started

Inventory command:
- `go test -count=1 -json ./... > /tmp/s5-a-full-go-test.json`
- Result: FAIL.
- Failed packages: 4.
- Raw evidence: `/tmp/s5-a-full-go-test.json`.

Scope-grant class legend:
- (a) add `EVIDENCE_TARGET` to legacy candidate constructions, assertion-preserving, value per the test context default E1/E2.
- (b) owed/record_kind seat-scope expectations updated to the 053113 settled posture.
- (c) crash/applicability fixtures moved past the new required field with original mutation-point assertion intent preserved.
- (d) exactly one: `TestOwedItemAcceptsNonOperatorSeat` inverts and renames.

Inventory:

| path | class | failing tests / edit note | s5-b surface collision |
| --- | --- | --- | --- |
| `cmd/frank-mcp/mcp_test.go` | (a) | `TestSubmitArgumentsRoundTripStructuredStringCarrier` now rejects before the SCOPE_DIFF canonicality assertion; add `EVIDENCE_TARGET` to accepted and bad submit headers while preserving the canonical/non-canonical assertion target. | none; s5-b marked `cmd/*` OUT |
| `internal/engine/config_change_test.go` | (a) | `TestConfigChangeDigestMismatchRejected` rejects on `EVIDENCE_TARGET:required` before the intended `new_digest` rejection; add `EVIDENCE_TARGET` to `configChangeRecord(...)` constructions. | none |
| `internal/engine/pipeline_test.go` | (a) | `TestSubmitHandlerStampsAndAcceptsValidCandidate`, `TestSubmitHandlerAssignsRelayIDAndProjectionIntents`, `TestOperatorVerdictOneShotRunsThroughSubmitHandler`, `TestSubmitHandlerBuildsOwedProjectionFromProvidedTable`, and `TestSubmitHandlerRejectsStalePositiveParentAfterRender` reject before their identity/projection/verdict/parent assertions; add context-appropriate `EVIDENCE_TARGET` to the affected submit records. | none |
| `internal/obligation/owed_test.go` | (a), (d) | `TestOwedItemSubmitProjectsOpenAndDispositionClosesIt` needs `EVIDENCE_TARGET` on operator owed/disposition constructions; `TestOwedItemAcceptsNonOperatorSeat` is the named class (d) inversion/rename and should reject under the settled operator-only owed posture after required-field coverage is present. | none |
| `test/fixtures/f11_test.go` | (c) | Config-change crash matrix rows, S2 clean-completion config-change, and applicability rows for config-change/operator-verdict/owed-item/owed-disposition fail before their prior mutation points; add/move `EVIDENCE_TARGET` in the mutation helpers so original crash/applicability assertions remain the target. | none |
| `test/fixtures/main_assembly_test.go` | (a) | `TestFrankBinaryAssemblesAuthenticatedSubmitProjectRead`, `TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition`, and `TestFrankBinaryReadCorruptionQueuesLiveQuarantine` reject submit/disposition records before accepted/project/read/quarantine assertions; add `EVIDENCE_TARGET` to the binary, owed, and quarantine submit records. | none |
| `test/fixtures/s4_config_change_test.go` | (a) | `TestRunningConfigUnchangedUntilRestart`, `TestConfigChangeReadRedactedForNonOperator`, `TestConfigChangeReadFullForOperator`, `TestConfigChangeProjectionsCarryNoMemberBytes`, and `TestRestartWithNewRegistryBouncesStaleForm` reject config-change records before config-change/read/redaction/stale-digest assertions; add `EVIDENCE_TARGET` to `configChangeRelay(...)` and the stale candidate where needed. | none |
| `test/fixtures/s4_shim_test.go` | (a) | `TestNudgeRecipientOnly`, `TestNudgeAllRecipients`, `TestCCRecipientMailboxedNudgedOnceAndPathClean`, and `TestOfflineRecipientNudgedOnReconnect` reject shim submit records before nudge/CC/offline assertions; add `EVIDENCE_TARGET` to `s4Relay(...)`. | none |

No `ESC` classes observed in the fresh run.

No standalone class (b) failing file was observed. The record_kind posture currently appears as the named class (d) inversion in `internal/obligation/owed_test.go`; F11 owed applicability failures are classified as class (c) because they must reach the prior mutation points after required-field coverage is present.

No s5-b collision was observed. The failing files above are tests/fixtures outside s5-b's granted implementation surface; the only potentially broad directory, `cmd/*`, is explicitly OUT for s5-b.

ACTIONS_GIT_REF: inventory-only relay; no source edits in this step; full battery run generated `/tmp/s5-a-full-go-test.json`; standing WIP remains uncommitted in `~/frank-s5-team/s5-a`
FINAL_GIT_STATUS_SHORT:
 M internal/fieldspec/registry.json
 M internal/fieldspec/registry_test.go
?? test/fixtures/s5_registry_dormancy_test.go
