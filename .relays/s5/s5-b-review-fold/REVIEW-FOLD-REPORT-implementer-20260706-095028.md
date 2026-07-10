## REVIEW-FOLD report - s5-b MF-4 registry source resolution folded

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-b-review-fold
PARENT_DISPATCH_ID: s5-b-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-mechanisms
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-094300.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: MF-4 folded; registry-source resolver now scans all candidates and supports merged-main in-repo bytes
FOLD_SCOPE:
- test/fixtures/s5_config_change_test.go -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: s5-b-mechanisms@78bda2e; commit 78bda2e s5-b: fold registry source resolution
FINAL_GIT_STATUS_SHORT: none - clean code worktree

Summary:
- MF-4 fixed in `test/fixtures/s5_config_change_test.go` only.
- `s5ALandedRegistryBytes` now probes every candidate path and returns the first candidate whose SHA-256 matches the pinned `dd7d0b5` registry digest.
- Found-but-mismatched candidates no longer stop discovery; if candidates exist and none match, the fixture fatals with drift evidence. It skips only when no candidate path exists at all.
- Added the in-repo `../../internal/fieldspec/registry.json` candidate so merged-main runs can use the integrated registry bytes.
- Removed the machine-absolute `~/frank-s5-team/s5-a/...` fallback; the portable sibling-worktree relative path remains for pre-merge runs.

TDD evidence:
- RED: `go test -count=1 ./test/fixtures -run TestS5ALandedRegistryBytesSkipsMismatchedCandidateWhenLaterCandidateMatches` failed on the mismatched env-var registry path with SHA `5df508...`, before resolver changes.
- GREEN: the same targeted test passed after the resolver scanned later candidates.

Focused verification:
- `go test -count=1 ./test/fixtures -run 'TestS5ConfigChange|TestS5GateRaiseFullMapRoutingEscalationFromLandedRegistry|TestS5ALandedRegistryBytes'` passed.
- `FRANK_S5_A_REGISTRY=internal/fieldspec/registry.json go test -count=1 ./test/fixtures -run 'TestS5ConfigChange|TestS5GateRaiseFullMapRoutingEscalationFromLandedRegistry|TestS5ALandedRegistryBytes'` passed.

Full verification at `78bda2e`:
- `go build ./...` exit 0.
- `go test -count=1 ./...` exit 0.
- `go vet ./...` exit 0.

Out-of-scope preserved:
- No production code changed.
- No registry bytes changed.
- No merge, push, PR, cmd wiring, or live integration action was taken.
