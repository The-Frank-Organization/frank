# frank-mcp

transport/provenance only; done-state and `record_integrity` remain `self_reported` until Step-2 observe.

`frank-mcp` is a stdio MCP shim for one host session and one current frank seat credential. It translates MCP `initialize`, `tools/list`, and `tools/call` to the private frank socket dialect and keeps the seat-facing tool set at exactly `submit`, `project`, and `read`. Audit and roster views ride as `project` arguments, not as extra tools.

## Configuration

Default configuration uses environment variables:

```sh
FRANK_SOCKET=/tmp/frank-s4.sock
FRANK_CREDENTIAL=<seat credential>
```

`-socket` may override `FRANK_SOCKET`. `-credential-file <path>` is the secondary credential source for hosts that cannot provide per-server environment variables; the file must be mode 0600. There is intentionally no bare credential CLI flag.

A local host compromise can steal the operator-provisioned secret. Operators can live re-mint with a `seat_mint` record; the new credential appears only in the operator submit reply and the old credential is invalidated and force-closed by the conductor. A compromised operator credential is still a local custody incident.

## Runtime posture

The shim is poll-first. Hosts should call `project` at turn start; nudge notifications and `notifications/tools/list_changed` only accelerate convergence. If a host hides custom notifications, mailbox truth still converges through `project` and `read`. After a conductor restart or socket drop, the shim closes, re-authenticates once, and retries the same call once; submit retry relies on conductor-side intake replay by content hash.

Every shim-generated MCP-visible error is scrubbed to one fixed class: `shim:conductor-unreachable`, `shim:auth-failed`, `shim:connection-lost`, `shim:frame-too-large`, or `shim:protocol-error`. Socket paths and credential bytes stay out of MCP stdout.
