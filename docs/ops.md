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

Mint bootstrap/admin seats before serving. This path is genesis-time only after the initial config record; once the conductor is live, use an operator `record_kind: seat_mint` submit instead.

```sh
frank -root /abs/team-store -mint s4-wire.implementer -role implementer
frank -root /abs/team-store -mint s4-wire.planner -role planner
frank -root /abs/team-store -mint operator -role operator -operator
```

Start the conductor:

```sh
frank -root /abs/team-store -socket /tmp/frank-s4.sock
```

Only one conductor may serve a store root. Startup takes an exclusive `flock` on `<root>/conductor.lock` before recovery or reads; a second conductor exits with `root-lock-held` and holder diagnostics.

## Stop and status

Stop with SIGINT or SIGTERM. The intake journal, canonical records, and recovery replay make restart the normal recovery path.

Status is socket liveness plus channel diagnostics. If phase-0 cannot open for submit, read-only diagnostics expose `project` and `read`; `submit` is absent until the store opens Ready.

After any form re-render bounce, hosted seats must re-read the schema before retrying. Hosted seats do not consume `tools/list_changed`, so a cached tool constant can remain stale even after the conductor announces a schema change.

## Seat wiring

One seat = one current credential = one host MCP config entry. A session occupies a durable seat by launching `frank-mcp` with that seat credential. Killing the host session or shim closes the socket; relaunching with the same current credential reoccupies the same seat and mailbox.

S4 corrects delivery to the locked TO/CC-mailbox semantics (m-1 §5); S1–S3 delivered to Envelope.To only.

Two blessed wiring patterns:

1. Separate per-seat host config scopes, each with its own `FRANK_SOCKET` and `FRANK_CREDENTIAL`.
2. Host config with `${VAR}` indirection where supported, with the shell providing the seat credential per launch.

If the same credential is launched twice while the first connection is live, the second launch receives `auth:channel-active`. For a wedged-but-alive host, the operator remedy is to kill the host session or shim; the kernel close frees the channel.

Live mint/re-mint is an operator submit with `record_kind: seat_mint` and Body JSON `{"seat":"...","role":"...","is_operator":false}`. Acceptance is the pivot; derived work replaces the binding row with a fresh credential, invalidates the old credential at auth, and force-closes any live old channel. The fresh credential and endpoint appear only in the operator submit reply as the custody handoff; they are not written to records, projections, INDEX rows, or ordinary reads. If the operator process crashes before capturing the reply, the admin remedy is a stopped-conductor read of the 0600 binding table.

Live-minted seats begin `minted`. Their first accepted boot submit carries only `PHASE`, `CEREMONY_TIER`, `SUBJECT`, `charter_loaded`, and `dispatch_status`; after that they are `active` for the current generation. Operator/orchestrator seats can inspect `project` with `{"view":"roster"}` for the seven roster fields; default `project` and `read` are not lifecycle-gated.

Scoped review waivers are operator records using `rationale`, `waiver_scope`, and `retracts`; those fields are absent from non-operator render and hand-crafted non-operator submissions carrying them are rejected.

Credential custody is confusion-resistant, not theft-proof. A local host compromise can steal the operator-provisioned secret. `frank -mint` and `seat_mint` replies print credentials to stdout for capture. `frank -operator-submit -credential` is ps-visible admin tooling; keep it short-lived. Shim use should prefer `FRANK_CREDENTIAL`; `-credential-file` is the secondary path and must be 0600.
