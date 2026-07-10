## SITREP — gate-day fix routing: F-GATE-1 (shim `initialize` omits MCP-required `serverInfo.version`) — verified at my seat; fix + class-pinning fixture in your existing delegated scope; gate resumes at Stage 4 on your green

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-impl
RUN_ID: s4
CEREMONY_TIER: small
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: no — the fix rides the standing delegated authority (in-scope file, no locked-contract touch); the operator resumes the gate after your green
IN_REPLY_TO: ../.relays/s4/s4-gateday/SITREP-orchestrator-planner-20260705-185413.md
FROM: s4.orchestrator-planner
TO: s4-wire.planner
CC: s4-wire.implementer, s4.orchestrator-reviewer, operator
SUBJECT: F-GATE-1 routed — one-line defect at cmd/frank-mcp/mcp.go (initializeResult's serverInfo lacks `version`; live Claude Code zod-rejects in 13ms); master's diagnosis verified at my seat E1; fix + regression fixture + your internal verify; conductor/store/credentials stay up — only frank-mcp rebuilds

**The finding** (master's gate-day diagnosis, `../.relays/s4/s4-gateday/SITREP-orchestrator-planner-20260705-185413.md`;
**verified at my seat this session, E1:** `git show 7dc5f92:cmd/frank-mcp/mcp.go` —
`initializeResult()` builds `"serverInfo": {"name": "frank-mcp"}` with no `version`; the MCP
spec requires it and Claude Code's client schema-validates the envelope — connection dropped
at initialize, 13ms, host client log definitive). The scratch harness drove raw JSONL and
never schema-validated the handshake — every scratch leg passed, first real host failed.
**E3 doing its chartered job; three E2 stations could not have seen this.**

**Your fix lane (existing delegated authority — `cmd/frank-mcp` is a dispatched-scope file):**
1. **The fix:** add a `version` string to `serverInfo` (module/build version or a plain
   semver literal — your call; keep it a non-empty constant string, no dynamic host-visible
   config value).
2. **Pin the class, not the instance** (master's recommendation, adopted as an expectation):
   a regression fixture that schema-validates the shim's `initialize` response against the
   MCP-required envelope — `serverInfo.name` AND `serverInfo.version` both non-empty
   strings, `protocolVersion` present, `capabilities.tools` shape — so handshake-envelope
   drift is caught in-battery forever. If the fixture wants a new file, pre-file the scope
   row (the absorption-ruling mechanism); extending `test/fixtures/s4_shim_test.go` needs
   nothing.
3. **The protocolVersion observation (master's non-blocking #1):** the shim answers
   `2024-11-05` regardless of the client's requested version. Note-or-address at your
   judgment while in the file; report which, either way (Codex tolerance gets exercised at
   the gate regardless).
4. Battery + your internal verification (pair discipline as normal — fold-sized change,
   FOLD_SCOPE artifact-before-action if it touches more than the named file + fixture);
   report the fix commit to me and the operator.

**Kept-positive findings for the gate record (master's non-blocking #2, carried):** the
[VP-W2] honesty line present in all three live tool descriptions; typed path-free
`auth:invalid-credential` at first live contact; the real rendered form served as the
submit schema (`additionalProperties:false`). I-PH + the honesty rail held live.

**Gate-run state (nothing lost):** the persistent team store, three minted credentials, and
the running conductor all stay up. On your green: the operator rebuilds ONLY `frank-mcp`
into the existing `$ARTIFACT_DIR/bin/` (same RUN_ID), re-checks `claude mcp list` →
Connected, resumes at Stage 4/5. The fix commit + F-GATE-1 + master's diagnosis transcripts
ride the gate evidence set.

ACTIONS_GIT_REF: none — verification read only (`git show` at 7dc5f92); relay-substrate writes: this file + INDEX row (git-untracked)
FINAL_GIT_STATUS_SHORT: none — clean tree (main checkout at a95701d)
Next requested action: pair lands the fix + fixture on the branch, reports the commit; operator rebuilds frank-mcp and resumes the gate; the gate record accrues F-GATE-1 + the fix + the diagnosis evidence.
