## SITREP - s4-wire-impl implementation report

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
REPORT_KIND: implementation-report
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-plan-review-implementer-r3
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - live-seat designation, section-7 authorization itself, E3 gate-day execution, s4-close, and merge remain operator gates
DESIGN_LOCK_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s4-grill-s4-wire
PLAN_LOCK_ID: s4-slice-4-plan
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator
IN_REPLY_TO: s4-wire-impl/IMPL-planner-20260705-124727.md
BASE: main@28dfa33
BRANCH: s4-wire-impl
HEAD: 796b7be
SUBJECT: implementation report - Tasks 1-13 complete on branch s4-wire-impl; final E2 floors green; merge not authorized

## Summary

Implemented the s4 wire-up plan r3 in order on branch `s4-wire-impl` from `main@28dfa33`. The final branch tip is `796b7be` (`s4: E3 gate procedure of record`).

Merge is not authorized by this report. The live-seat designation, section-7 gate-day authorization, E3 live host run, s4-close, and merge remain separate operator gates.

## Commit Map

| task | commit | evidence surface |
|---|---|---|
| 1 | `0ef8cf6` | frame bound from engine config, default 1 MiB, typed inbound/outbound frame-too-large refusals |
| 2 | `0ffbe8d` | one live channel per credential, typed second-connect rejection, dead-connection recovery |
| 3 | `1a80ac1` | per-recipient wake, retired global pending queue, no cross-seat metadata |
| 4 | `e1ed9ba` | operator-scoped `config_change`, member bytes in canonical record, conductor recomputed digest |
| 4b | `7d64f33` | redacted non-operator read view for `config_change`, full operator/admin read path |
| 5 | `3ce66c9` | phase-0 genesis-to-config-change chain, config materialized from records |
| 6 | `a1f998d` | crash matrix gains config-change applicability class |
| 7 | `49c60e9` | `frank-mcp` stdio shim skeleton and scrubbed error chokepoint |
| 8 | `b6b233d` | rendered form is the MCP submit schema, closed string-carrier rule, digest const |
| 9 | `ca7bc90` | drift negotiation via re-render bounce and `notifications/tools/list_changed` |
| 10 | `49c5f1f` | I-PH bridge matrix across seven surface classes and ceiling carve-out |
| 11 | `c8bd239` | ops/custody/usage docs, socket-path preflight, README ruled-in delta |
| 12 | no commit | floor verification only; no fixes needed |
| 13 | `796b7be` | E3 gate procedure of record |

## FAIL-First / Negative Fixture Evidence

- Task 11 preserved red-first output before the socket preflight fix: `go test -count=1 ./test/fixtures -run TestFrankBinarySocketPathPreflight` failed with the raw Darwin bind failure (`bind: invalid argument`). After `validateSocketPath`, the same fixture passed and asserts typed `socket path too long` output containing `darwin`, with no raw bind leak.
- Task 12 floor checks re-ran the shim enumeration, enum grep, race set, and full suite from the final branch tip; no floor-fix commit was required.
- Task 13 dry-run commands from `docs/sprints/2026-07-05-s4-slice-4/results/e3-gate-procedure.md` ran against scratch stores and passed. They cover conductor-side store init/mint/submit/project/read, operator owed disposition, second-connect/dead-recovery, offline wake, config-change restart/re-render, config-change redaction, and I-PH spot matrix surfaces. The two real host sessions remain gate-day operator-designated, not satisfied here.

## Final Verification

Fresh final-tip commands from `~/.config/superpowers/worktrees/frank/s4-wire-impl`:

- `go test -count=1 ./...` -> pass for all packages.
- `go vet ./...` -> pass, no output.
- `go test -race -count=1 ./internal/channel ./internal/seat ./internal/engine ./internal/intake ./internal/store` -> pass.
- `rg -n '"bounced"' internal/ cmd/` -> zero matches (rg exit 1).
- `go test -count=1 ./cmd/frank-mcp -run 'TestToolsListGolden|TestToolsListUsesRenderedSubmitSchemaWhenReachable'` -> pass.
- `go test -count=1 ./test/fixtures -run 'TestFrankBinaryAssemblesAuthenticatedSubmitProjectRead|TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition|TestFrankBinaryReissuesRecoveryWakeForExistingMailbox|TestOfflineRecipientNudgedOnReconnect|TestRestartWithNewRegistryBouncesStaleForm|TestConfigChangeReadRedactedForNonOperator|TestConfigChangeReadFullForOperator' && go test -count=1 ./internal/channel -run 'TestSecondConnectSameCredentialRejected|TestProvenDeadRecovery|TestKillHostEscapeHatch' && go test -count=1 ./test/fixtures -run 'TestS4IPHBridgeSurfaceMatrix|TestCarveOutExactlyOneValue'` -> pass.
- `git diff --check` -> pass, no output.
- `$HOME/frank-s2-store` read-only count before final gate: `31`; after final gate: `31`.
- `git status --short --branch` -> `## s4-wire-impl`.

## Notes

- `docs/ops.md` uses an explicit generated `engine.json` source for `frank -init`; the repo does not ship `internal/config/default-engine.json`.
- `.relays/` is gitignored. Relay proof of this report is this file plus the appended `.relays/s4/INDEX.md` row and exact-file relay lint.

ACTIONS_GIT_REF: branch `s4-wire-impl` at `796b7be`; implementation commits are local in worktree `~/.config/superpowers/worktrees/frank/s4-wire-impl`; relay-substrate writes are this file plus `.relays/s4/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: code worktree clean at `## s4-wire-impl`; main checkout tracked status not modified by code work; `.relays/` has the expected gitignored report write.
