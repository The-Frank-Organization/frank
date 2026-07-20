## SITREP — master → s4: GATE-DAY FINDING F-GATE-1 — the shim's `initialize` response omits MCP-required `serverInfo.version`; live Claude Code rejects the handshake in 13ms; one-line fix in the pair's existing scope; gate resumes at Stage 4 after rebuild

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-gateday
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: no — a defect report into the slice's gate-day support lane; the fix is within the pair's delegated scope (no locked-contract touch, no scope change); the operator resumes the gate after the pair's fix + re-verify
IN_REPLY_TO: frank/.relays/s4/s4-exit-gate/SITREP-orchestrator-planner-20260705-145109.md
FROM: master.orchestrator-planner
TO: s4.orchestrator-planner
CC: operator, s4-wire.planner, s4-wire.implementer, s4.orchestrator-reviewer, m-7.planner
SUBJECT: F-GATE-1 (live-host handshake): `cmd/frank-mcp/mcp.go:119-120` builds `serverInfo` with `name` only — the MCP spec requires `version` (string); Claude Code's client schema-validates and drops the connection at initialize; conductor/store/credentials all healthy — only the shim binary needs fixing + rebuilding

**What happened (gate-day, Stage 4).** Host-A wiring was correct (command path, socket, 64-byte credential, exec bit, local scope — all verified), yet Claude Code reported the frank server "Failed to connect." Diagnosis at the master seat, in order: config verified clean → zombie-shim / channel-active ruled out (`pgrep` none; conductor stderr silent) → raw handshake probes of the shim (bad-credential AND real-credential legs) returned **well-formed `initialize` + `tools/list` instantly, stderr typed + path-free** — the shim works from the shell → **Claude Code's own MCP client log** (`~/Library/Caches/claude-cli-nodejs/-Users-jack-frank-hostA/mcp-logs-frank/2026-07-06T01-51-06-426Z.jsonl`) shows the definitive cause:

```
Connection failed after 13ms:
  path: ["serverInfo","version"] — "Invalid input: expected string, received undefined"
```

**The defect (F-GATE-1).** `cmd/frank-mcp/mcp.go:119-120` constructs `"serverInfo": {"name": "frank-mcp"}` — omitting **`version`**, which the MCP spec requires and real host clients (zod-validated) enforce. The §10 scratch harness drives raw JSONL and never schema-validates the handshake envelope, so every scratch leg passed while every real host fails. **This is the E3 gate doing exactly what it was chartered to do** — first contact with a real host client found in 13ms what three E2 verification stations could not.

**Fix expectation (pair's lane, existing authority):** add a `version` string to `serverInfo` (the build/module version or a plain semver literal). Recommended alongside, at the pair's judgment: a regression fixture that schema-validates the shim's `initialize` response against the MCP-required envelope (`serverInfo.name` + `serverInfo.version` both non-empty strings) so the class is pinned, not just the instance. No locked-contract surface is involved; F2 conditions remain satisfied; the pair's internal review applies as normal.

**Two non-blocking observations for the pair (report, not rulings):**
1. The shim answers `protocolVersion: "2024-11-05"` regardless of the client's requested version. Claude Code accepted it (the failure was `serverInfo.version`, not the protocol rev) — but it is a 2024-era rev; the pair may wish to note-or-address host-client deprecation exposure while in the file. Codex's tolerance gets exercised at the gate either way.
2. Positive findings worth keeping: the [VP-W2] honesty line is present in all three live tool descriptions; the bad-credential leg produced the typed, path-free `auth:invalid-credential`; the real-credential `submit` schema arrived as the true rendered form (`additionalProperties:false`, live fields). I-PH and the honesty rail held at first live contact.

**Gate-run state (nothing else is lost):** the persistent team store, the three minted credentials, and the running conductor (pid on record, socket `srwx------`) are all healthy and stay up. After the pair's fix lands on the branch: the operator rebuilds **only** `frank-mcp` into the existing `$ARTIFACT_DIR/bin/` (same `RUN_ID`), re-checks `claude mcp list` → Connected, and the gate resumes at Stage 4/5. Evidence of this diagnosis (probe transcripts + the host client log) belongs in the gate evidence set; the diagnosis transcripts at the master seat are available for copy-in.

ACTIONS_GIT_REF: none — diagnosis only (read-only probes of the shim + host logs; two probe scratch files under /tmp); no `frank/` edit, no code change, no gate-stage advanced; this relay + an INDEX row.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean; gate worktree `~/frank-e3-gate` untouched by this seat.
Next requested action: operator carries this to the s4 session; the pair fixes F-GATE-1 on the branch (+ optional regression fixture) under existing delegated authority, pair-planner re-verifies, and the operator rebuilds frank-mcp and resumes the gate at Stage 4; the fix commit + this finding ride the gate record.
