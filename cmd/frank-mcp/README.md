# frank-mcp

transport/provenance only; done-state and `record_integrity` remain `self_reported` until Step-2 observe.

`frank-mcp` is a stdio MCP shim for one host session and one frank seat credential. It translates MCP `initialize`, `tools/list`, and `tools/call` to the private frank socket dialect and keeps the seat-facing tool set at exactly `submit`, `project`, and `read`.

## Configuration

Default configuration uses environment variables:

```sh
FRANK_SOCKET=/tmp/frank-s4.sock
FRANK_CREDENTIAL=<seat credential>
```

`-socket` may override `FRANK_SOCKET`. `-credential-file <path>` is the secondary credential source for hosts that cannot provide per-server environment variables; the file must be mode 0600. There is intentionally no bare credential CLI flag.

A local host compromise can steal the operator-provisioned secret. There is no in-band rotation, supersede, or revoke path in this slice; a compromised credential means stopping the conductor and doing admin-time store surgery.

## Runtime posture

The shim is poll-first. Hosts should call `project` at turn start; nudge notifications and `notifications/tools/list_changed` only accelerate convergence. If a host hides custom notifications, mailbox truth still converges through `project` and `read`.

Every shim-generated MCP-visible error is scrubbed to one fixed class: `shim:conductor-unreachable`, `shim:auth-failed`, `shim:connection-lost`, `shim:frame-too-large`, or `shim:protocol-error`. Socket paths and credential bytes stay out of MCP stdout.
