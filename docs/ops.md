# frank operations

transport/provenance only; done-state and `record_integrity` remain `self_reported` until Step-2 observe.

## Store and socket

Use one absolute store root for the team store. The store is the governance domain: records, projections, mailboxes, config members, and binding data all live under that root.

Use a short Unix socket path. On darwin, long AF_UNIX paths fail at bind time, so prefer `/tmp/<team>.sock` or another short path. `frank` now refuses paths whose byte length is 100 or more with a typed startup error instead of the raw bind failure.

## Start

Initialize a fresh store with pinned config:

```sh
cat > /tmp/frank-engine.json <<'JSON'
{"gc_enabled":false,"segment_rotate_bytes":4194304}
JSON
frank -root /abs/team-store -registry internal/fieldspec/registry.json -engine-config /tmp/frank-engine.json -init
```

Mint seats before serving:

```sh
frank -root /abs/team-store -mint s4-wire.implementer -role implementer
frank -root /abs/team-store -mint s4-wire.planner -role planner
frank -root /abs/team-store -mint operator -role operator -operator
```

Start the conductor:

```sh
frank -root /abs/team-store -socket /tmp/frank-s4.sock
```

## Stop and status

Stop with SIGINT or SIGTERM. The intake journal, canonical records, and recovery replay make restart the normal recovery path.

Status is socket liveness plus channel diagnostics. If phase-0 cannot open for submit, read-only diagnostics expose `project` and `read`; `submit` is absent until the store opens Ready.

## Seat wiring

One seat = one credential = one host MCP config entry. A session occupies a durable seat by launching `frank-mcp` with that seat credential. Killing the host session or shim closes the socket; relaunching with the same credential reoccupies the same seat and mailbox.

Two blessed wiring patterns:

1. Separate per-seat host config scopes, each with its own `FRANK_SOCKET` and `FRANK_CREDENTIAL`.
2. Host config with `${VAR}` indirection where supported, with the shell providing the seat credential per launch.

If the same credential is launched twice while the first connection is live, the second launch receives `auth:channel-active`. For a wedged-but-alive host, the operator remedy is to kill the host session or shim; the kernel close frees the channel. There is no heartbeat, supersede, remint, revoke, or in-band rotation in this slice.

Credential custody is confusion-resistant, not theft-proof. A local host compromise can steal the operator-provisioned secret. `frank -mint` prints the credential to stdout for admin capture. `frank -operator-submit -credential` is ps-visible admin tooling; keep it short-lived. Shim use should prefer `FRANK_CREDENTIAL`; `-credential-file` is the secondary path and must be 0600.
