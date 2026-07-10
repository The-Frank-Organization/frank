## SITREP - s4-wire-impl F-GATE-1 fold report

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
REPORT_KIND: f-gate-1-fold-report
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-impl
RUN_ID: s4
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no new gate; operator resumes the existing E3 host gate after green
DESIGN_LOCK_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s4-grill-s4-wire
PLAN_LOCK_ID: s4-slice-4-plan
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator
IN_REPLY_TO: s4-wire-impl/SITREP-planner-20260705-190031.md
BASE: s4-wire-impl@7dc5f92
BRANCH: s4-wire-impl
HEAD: 605b3ef
SUBJECT: F-GATE-1 folded - frank-mcp initialize serverInfo.version plus structural initialize-envelope fixture; protocolVersion pin retained deliberately

## Summary

Folded F-GATE-1 on branch `s4-wire-impl`. Commit: `605b3ef s4: fold f-gate-1 frank-mcp initialize version`.

`cmd/frank-mcp/mcp.go` now declares a non-empty constant `mcpServerVersion = "0.4.0"` and includes it in `initialize` response `serverInfo.version`.

`test/fixtures/s4_shim_test.go` now builds and launches `frank-mcp`, sends an `initialize` JSON-RPC request, and structurally decodes the response envelope. It asserts a non-empty `protocolVersion`, declared `capabilities.tools.listChanged`, non-empty `serverInfo.name`, and non-empty `serverInfo.version`. This is a typed envelope check, not a substring grep.

No FOLD_SCOPE artifact was created: the planner/orchestrator F-GATE-1 direction explicitly exempted the fold when touching only `cmd/frank-mcp/mcp.go` and `test/fixtures/s4_shim_test.go`; the final diff stayed inside those two files.

## Protocol Version Note

The `protocolVersion` answer remains pinned to `2024-11-05` deliberately. This gate only fixes the missing required `serverInfo.version`; it does not change MCP negotiation behavior or blindly echo the client's requested protocol version.

## Red-First Evidence

- Before the production fix, `go test -count=1 ./test/fixtures -run TestFrankMCPInitializeEnvelopeDeclaresServerInfoVersion -v` failed for the intended reason: `initialize serverInfo = {Name:frank-mcp Version:}`.
- After adding the constant version field, the same focused fixture passed.

## Final Verification

Fresh commands from `~/.config/superpowers/worktrees/frank/s4-wire-impl` at `605b3ef`:

- `go test -count=1 ./test/fixtures -run TestFrankMCPInitializeEnvelopeDeclaresServerInfoVersion -v` -> pass.
- `go test -count=1 ./...` -> pass for all packages.
- `go vet ./...` -> pass, no output.
- `git diff --check` -> pass, no output.
- `git status --short --branch` -> `## s4-wire-impl`.

## Evidence Levels

Implementation and verification evidence is E2. The live host retry remains E3 and is still operator-run per the gate-day procedure.

ACTIONS_GIT_REF: branch `s4-wire-impl` at `605b3ef`; fold commit `605b3ef`; relay-substrate writes are this report and `.relays/s4/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: code worktree clean at `## s4-wire-impl`; main checkout tracked status clean with expected gitignored `.relays/` writes.
