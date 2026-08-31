# frank operations

The observe layer stamps per-field `evidence_integrity` (`observed` vs `self_reported`) at send; fields it does not observe remain `self_reported`.

This doc covers the conductor (`frank`) and the seat-side MCP shim (`frank-mcp`). The other binaries in `cmd/` — `frank-app`, `frank-broker`, `frank-connector`, and `frank-worker` — are the mid-build coding-agent harness; see the root README for what each one is.

## Store and socket

Use one absolute store root for the team store. The store is the governance domain: records, projections, mailboxes, config members, and binding data all live under that root.

Use a short Unix socket path. On darwin, long AF_UNIX paths fail at bind time, so prefer `/tmp/<team>.sock` or another short path. `frank` now refuses paths whose byte length is 100 or more with a typed startup error instead of the raw bind failure.

## Start

Initialize a fresh store with pinned config. Init pins three config sources by digest — the FieldSpec registry, the engine config, and the invariant catalog — so all three flags are required. The engine config must be `version: 2` with a `supply` section (lane roots + at least one named check suite) or the store will init but the conductor will refuse to serve (`config-load: supply`); the root README's Quick start derives a minimal working config step by step.

```sh
frank -root /abs/team-store \
      -registry internal/fieldspec/registry.json \
      -engine-config /abs/engine.json \
      -catalog test/invariants/catalog.v1.json \
      -init
```

Mint bootstrap/admin seats before serving. This path is genesis-time only after the initial config record; once the conductor is live, use an operator `record_kind: seat_mint` submit instead.

```sh
frank -root /abs/team-store -mint core.implementer -role implementer
frank -root /abs/team-store -mint core.planner -role planner
frank -root /abs/team-store -mint operator -role operator -operator
```

Start the conductor:

```sh
frank -root /abs/team-store -socket /tmp/frank.sock
```

Only one conductor may serve a store root. Startup takes an exclusive `flock` on `<root>/conductor.lock` before recovery or reads; a second conductor exits with `root-lock-held` and holder diagnostics.

## Stop and status

Stop with SIGINT or SIGTERM. The intake journal, canonical records, and recovery replay make restart the normal recovery path.

Status is socket liveness plus channel diagnostics. If phase-0 cannot open for submit, read-only diagnostics expose `project` and `read`; `submit` is absent until the store opens Ready.

After any form re-render bounce, hosted seats must re-read the schema before retrying. Hosted seats do not consume `tools/list_changed`, so a cached tool constant can remain stale even after the conductor announces a schema change.

## Seat wiring

One seat = one current credential = one host MCP config entry. A session occupies a durable seat by launching `frank-mcp` with that seat credential. Killing the host session or shim closes the socket; relaunching with the same current credential reoccupies the same seat and mailbox.

Delivery follows TO/CC-mailbox semantics: a relay lands in the mailbox of every seat named in TO or CC.

Two blessed wiring patterns:

1. Separate per-seat host config scopes, each with its own `FRANK_SOCKET` and `FRANK_CREDENTIAL`.
2. Host config with `${VAR}` indirection where supported, with the shell providing the seat credential per launch.

If the same credential is launched twice while the first connection is live, the second launch receives `auth:channel-active`. For a wedged-but-alive host, the operator remedy is to kill the host session or shim; the kernel close frees the channel.

Live mint/re-mint is an operator submit with `record_kind: seat_mint` and Body JSON `{"seat":"...","role":"...","is_operator":false}`. Acceptance is the pivot; derived work replaces the binding row with a fresh credential, invalidates the old credential at auth, and force-closes any live old channel. The fresh credential and endpoint appear only in the operator submit reply as the custody handoff; they are not written to records, projections, INDEX rows, or ordinary reads. If the operator process crashes before capturing the reply, the admin remedy is a stopped-conductor read of the 0600 binding table.

Live-minted seats begin `minted`. Their first accepted boot submit carries only `PHASE`, `CEREMONY_TIER`, `SUBJECT`, `charter_loaded`, and `dispatch_status`; after that they are `active` for the current generation. Operator/orchestrator seats can inspect `project` with `{"view":"roster"}` for the seven roster fields; default `project` and `read` are not lifecycle-gated.

Scoped review waivers are operator records using `rationale`, `waiver_scope`, and `retracts`; those fields are absent from non-operator render and hand-crafted non-operator submissions carrying them are rejected.

Credential custody is confusion-resistant, not theft-proof. A local host compromise can steal the operator-provisioned secret. `frank -mint` and `seat_mint` replies print credentials to stdout for capture. `frank -operator-submit -credential` is ps-visible admin tooling; keep it short-lived. Shim use should prefer `FRANK_CREDENTIAL`; `-credential-file` is the secondary path and must be 0600.
