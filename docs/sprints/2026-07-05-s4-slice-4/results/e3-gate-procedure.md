# s4 E3 gate procedure

transport/provenance only; done-state and `record_integrity` remain `self_reported` until Step-2 observe.

This is the operator-run procedure of record for the s4 wire-up exit gate. The gate-day run uses two real host sessions: host A is Claude Code, host B is Codex. Local scratch dry-runs may use `frank-mcp` directly, but they do not satisfy the live-host designation gate.

## 0. Evidence layout

Use one evidence directory per gate attempt:

```sh
export RUN_ID="$(date -u +%Y%m%dT%H%M%SZ)"
export ARTIFACT_DIR="$PWD/docs/sprints/2026-07-05-s4-slice-4/results/e3-$RUN_ID"
mkdir -p "$ARTIFACT_DIR"/{bin,transcripts,secrets}
chmod 700 "$ARTIFACT_DIR/secrets"
export FRANK="$ARTIFACT_DIR/bin/frank"
export FRANK_MCP="$ARTIFACT_DIR/bin/frank-mcp"
go build -o "$FRANK" ./cmd/frank
go build -o "$FRANK_MCP" ./cmd/frank-mcp
```

Attach transcripts and redacted host config to the exit-gate report. Do not attach credential values or credential files.

## 1. Persistent store init

This store becomes the persistent team store; the section 7 backstop trips here.

```sh
export TEAM_STORE="/abs/frank-s4-team-store"
export FRANK_SOCKET="/tmp/frank-s4.sock"
cat > "$ARTIFACT_DIR/engine.json" <<'JSON'
{"gc_enabled":false,"segment_rotate_bytes":4194304}
JSON
"$FRANK" -root "$TEAM_STORE" -registry internal/fieldspec/registry.json -engine-config "$ARTIFACT_DIR/engine.json" -init \
  >"$ARTIFACT_DIR/transcripts/01-init.stdout" \
  2>"$ARTIFACT_DIR/transcripts/01-init.stderr"
```

Evidence: `transcripts/01-init.*`, the persistent store path, and the store-root `records/genesis.json`.

## 2. Mint seats

Mint one credential per durable seat and save secrets out of the public report:

```sh
HOST_A="s4-wire.host-a"
HOST_B="s4-wire.host-b"
OPERATOR="operator"
A_CRED="$("$FRANK" -root "$TEAM_STORE" -mint "$HOST_A" -role implementer | sed 's/^credential=//')"
B_CRED="$("$FRANK" -root "$TEAM_STORE" -mint "$HOST_B" -role planner | sed 's/^credential=//')"
OP_CRED="$("$FRANK" -root "$TEAM_STORE" -mint "$OPERATOR" -role operator -operator | sed 's/^credential=//')"
umask 077
{
  printf 'A_CRED=%s\n' "$A_CRED"
  printf 'B_CRED=%s\n' "$B_CRED"
  printf 'OP_CRED=%s\n' "$OP_CRED"
} > "$ARTIFACT_DIR/secrets/credentials.env"
```

Evidence: `transcripts/02-mint-redacted.txt` with only seat names and roles, not credential values. The expected invariant is one seat = one credential = one host MCP config entry.

## 3. Start conductor

```sh
"$FRANK" -root "$TEAM_STORE" -socket "$FRANK_SOCKET" \
  >"$ARTIFACT_DIR/transcripts/03-frank.stdout" \
  2>"$ARTIFACT_DIR/transcripts/03-frank.stderr" &
FRANK_PID=$!
printf '%s\n' "$FRANK_PID" > "$ARTIFACT_DIR/frank.pid"
```

Evidence: process id, socket path, and startup transcript. Keep the socket path short.

## 4. Wire the two real hosts

Host A (Claude Code) MCP server configuration:

```sh
FRANK_SOCKET=/tmp/frank-s4.sock
FRANK_CREDENTIAL=<A_CRED from secrets/credentials.env>
command: /abs/path/to/frank-mcp
```

Host B (Codex) MCP server configuration:

```sh
FRANK_SOCKET=/tmp/frank-s4.sock
FRANK_CREDENTIAL=<B_CRED from secrets/credentials.env>
command: /abs/path/to/frank-mcp
```

Evidence: redacted host config screenshots or text dumps showing socket path, command path, and that credential values are supplied by environment or a 0600 file. Credential bytes are not report artifacts.

## 5. Live relay leg

Host A calls `tools/list`, captures the `submit` input schema, and uses the schema's `form_digest` const in a `submit` call:

```json
{
  "headers": {
    "PHASE": "SITREP",
    "AUTHORITY": "report-only",
    "CEREMONY_TIER": "medium",
    "SUBJECT": "E3 live relay"
  },
  "to": "s4-wire.host-b",
  "dispatch_id": "e3-live-relay",
  "body": "transport/provenance only; done-state and record_integrity remain self_reported until Step-2 observe.",
  "form_digest": "<schema const from tools/list>"
}
```

Host B then calls `project` and `read` for the returned relay id.

Evidence: `transcripts/10-host-a-tools-list.jsonl`, `transcripts/11-host-a-submit.jsonl`, `transcripts/12-host-b-project.jsonl`, and `transcripts/13-host-b-read.jsonl`. The exit report records the relay id, not a content-verification claim.

S4 corrects delivery to the locked TO/CC-mailbox semantics (m-1 §5); S1–S3 delivered to Envelope.To only.

## 6. Adversarial legs

No credential and bad credential:

```sh
printf '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}\n' |
  FRANK_SOCKET="$FRANK_SOCKET" FRANK_CREDENTIAL= "$FRANK_MCP" \
  >"$ARTIFACT_DIR/transcripts/20-no-credential.jsonl" 2>&1 || true
printf '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}\n' |
  FRANK_SOCKET="$FRANK_SOCKET" FRANK_CREDENTIAL=bad-credential "$FRANK_MCP" \
  >"$ARTIFACT_DIR/transcripts/21-bad-credential.jsonl" 2>&1 || true
```

Expected MCP-visible class: `shim:auth-failed`.

Second-connect reject:

```sh
{ printf '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}\n'; sleep 30; } |
  FRANK_SOCKET="$FRANK_SOCKET" FRANK_CREDENTIAL="$A_CRED" "$FRANK_MCP" \
  >"$ARTIFACT_DIR/transcripts/22-first-live.jsonl" 2>&1 &
FIRST_SHIM_PID=$!
sleep 1
printf '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}\n' |
  FRANK_SOCKET="$FRANK_SOCKET" FRANK_CREDENTIAL="$A_CRED" "$FRANK_MCP" \
  >"$ARTIFACT_DIR/transcripts/23-second-connect.jsonl" 2>&1 || true
kill "$FIRST_SHIM_PID" 2>/dev/null || true
```

Expected MCP-visible class: `shim:auth-failed`; the conductor-side rejection class is `auth:channel-active`.

Forged-FROM submit:

Submit a raw payload through `frank -operator-submit` while authenticated as host A, deliberately setting `envelope.from` and `envelope.role` to other values. The committed record must read back as `from = s4-wire.host-a` and `role = implementer` because the conductor stamps identity from the authenticated channel.

I-PH spot probes:

```sh
rg -n "$TEAM_STORE|$FRANK_SOCKET|binding|seats.json" "$ARTIFACT_DIR/transcripts" && exit 1 || true
rg -n "$A_CRED|$B_CRED|$OP_CRED" "$ARTIFACT_DIR/transcripts" && exit 1 || true
```

Evidence: transcripts for each adversarial leg plus the zero-hit I-PH spot-probe output.

## 7. Crash and liveness legs

Run these with the persistent store and real host wiring:

1. Kill `frank` after host A submits to offline host B; restart `frank`; host B reconnects and receives one recovery nudge, then `project` and `read` catch up.
2. Kill host A's shim; relaunch host A with the same credential; `project` catches up without reminting.
3. Keep host B offline while host A submits; host B reconnects and receives the offline-seat recovery nudge.

Evidence: `transcripts/30-kill-frank-before-restart.*`, `31-host-b-recovery-nudge.jsonl`, `32-host-a-shim-reconnect.jsonl`, and `33-offline-seat-reconnect.jsonl`.

## 8. Section 7 config-change round trip

Generate a real registry delta from the existing store-root registry:

```sh
export REGISTRY_DELTA="$ARTIFACT_DIR/registry-delta.json"
python3 - <<'PY'
import json, os, pathlib
root = pathlib.Path(os.environ["TEAM_STORE"])
out = pathlib.Path(os.environ["REGISTRY_DELTA"])
doc = json.loads((root / "config" / "fieldspec" / "registry.json").read_text())
doc.setdefault("provenance", {})["e3_gate_marker"] = os.environ["RUN_ID"]
out.write_text(json.dumps(doc, separators=(",", ":")) + "\n")
PY
export NEW_DIGEST="$(python3 - <<'PY'
import hashlib, os, pathlib
root = pathlib.Path(os.environ["TEAM_STORE"])
members = {
    "engine": (root / "config" / "engine.json").read_bytes(),
    "fieldspec": pathlib.Path(os.environ["REGISTRY_DELTA"]).read_bytes(),
}
manifest = b""
for name in sorted(members):
    manifest += name.encode() + b"\0" + hashlib.sha256(members[name]).hexdigest().encode() + b"\n"
print(hashlib.sha256(manifest).hexdigest())
PY
)"
```

Operator host calls `tools/list`, captures the current `form_digest`, then submits:

```json
{
  "headers": {
    "PHASE": "SITREP",
    "AUTHORITY": "report-only",
    "CEREMONY_TIER": "medium",
    "SUBJECT": "E3 registry config_change",
    "record_kind": "config_change",
    "member": "fieldspec",
    "new_digest": "<NEW_DIGEST>"
  },
  "dispatch_id": "e3-config-change",
  "body": "<contents of registry-delta.json>",
  "form_digest": "<operator tools/list schema const>"
}
```

Restart `frank`. Phase-0 must accept the genesis to config-change chain. Submit one stale-form relay using the pre-restart digest and record the re-render rejection. Then call `tools/list` again, submit the same relay with the refreshed digest, and record acceptance.

Evidence: `transcripts/40-config-change-submit.jsonl`, `41-frank-restart.*`, `42-stale-form-rejected.jsonl`, `43-refreshed-form-accepted.jsonl`, and the config-change relay id.

## 9. OI-S3-CONFIG-CHANGE disposition

Find the open owed row for `OI-S3-CONFIG-CHANGE` in the persistent store:

```sh
export OI_RELAY_ID="$(awk -F'|' '/OI-S3-CONFIG-CHANGE/ {gsub(/ /, "", $2); print $2}' "$TEAM_STORE/projections/owed/OPEN.md")"
test -n "$OI_RELAY_ID"
```

Through the operator channel, submit an `owed_disposition` with `disposes_owed = $OI_RELAY_ID`. The body must cite the accepted config-change relay id and this E3 evidence directory.

```json
{
  "headers": {
    "PHASE": "SITREP",
    "AUTHORITY": "report-only",
    "CEREMONY_TIER": "medium",
    "SUBJECT": "Dispose OI-S3-CONFIG-CHANGE",
    "record_kind": "owed_disposition",
    "disposes_owed": "<OI_RELAY_ID>"
  },
  "dispatch_id": "e3-oi-s3-config-change",
  "body": "Disposes OI-S3-CONFIG-CHANGE via config_change <relay_id> and evidence <ARTIFACT_DIR>. transport/provenance only; done-state and record_integrity remain self_reported until Step-2 observe.",
  "form_digest": "<operator tools/list schema const>"
}
```

Assert the open set is empty:

```sh
! rg -n 'relay-' "$TEAM_STORE/projections/owed/OPEN.md"
```

Evidence: `transcripts/50-owed-disposition-submit.jsonl` and `51-open-owed-empty.txt`.

## 10. Local scratch dry-run

These commands exercise every conductor-side class above against temporary stores, excluding only the two real host sessions:

```sh
go test -count=1 ./test/fixtures -run 'TestFrankBinaryAssemblesAuthenticatedSubmitProjectRead|TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition|TestFrankBinaryReissuesRecoveryWakeForExistingMailbox|TestOfflineRecipientNudgedOnReconnect|TestRestartWithNewRegistryBouncesStaleForm|TestConfigChangeReadRedactedForNonOperator|TestConfigChangeReadFullForOperator'
go test -count=1 ./internal/channel -run 'TestSecondConnectSameCredentialRejected|TestProvenDeadRecovery|TestKillHostEscapeHatch'
go test -count=1 ./test/fixtures -run 'TestS4IPHBridgeSurfaceMatrix|TestCarveOutExactlyOneValue'
```

Record the command output in `transcripts/90-local-dry-run.txt`. Green scratch fixtures are not a substitute for the gate-day host transcripts.
