# frank

**A governed courier for multi-agent teams.** frank is a small trusted middle, written in Go, that
carries every agent-to-agent message through a single serialized commit path where **identity is
stamped by the channel, lineage is computed by the conductor, and authority is enforced at commit** —
never taken from the payload, never validated by convention.

The name is from postal *franking* — the trusted mark that authorizes a relay to pass — and *frank*
as in candid: the system is explicit about what it can and cannot prove.

## Why

Agent messaging buses and orchestration conventions are converging fast: mailbox files, threads,
roles, gate messages, verification checklists. Those stacks answer coordination at the **convention
layer** — approvals are messages anyone can author, identity is self-asserted, evidence is supplied
by the sender. Their own threat models say so.

frank is the layer those designs scope out: **the trusted middle**. The same forged-approval or
fabricated-lineage move that passes silently through a convention stack bounces off frank **typed,
on the record** — because the properties are enforced in the commit path, not requested by etiquette.

## Features (Step-1: transport + provenance)

- **Append-only record store** with crash-atomic multi-file commit (stage → fsync → rename;
  presence = committed) and recovery replay. A **single-threaded serialized commit loop** is the sole
  in-process writer for governance-surface mutations.
- **Channel-stamped FROM** — a seat's identity comes from its authenticated channel (one active
  channel per credential, live re-mint with auth generations), never from what the seat writes.
- **Forms are the schema** — a declarative FieldSpec registry renders per-seat, per-phase forms, and
  the rendered form *is* the submit schema, pinned by a stable form digest. Config is pinned by a
  composite digest; changes go through an operator-authored `config_change` record, never silent edits.
- **Conductor-computed lineage** — the parent of a record is computed at commit from the submitter's
  own read history intersected with the accepted graph; a seat's `parent_hint` is honored only when
  provable.
- **Typed terminal outcomes** — exactly `accepted`, `rejected`, `held`. Every bounce is a typed,
  recorded fact.
- **Human gates** — monotonic gate raising, gate categories with a fail-safe default, owed items and
  dispositions, scoped waivers with retraction, and a local-outbox-only rule for anything leaving the
  store (with an egress scanner at the drain, dormant until Step-2).
- **A minimal wire** — Unix socket plus a stdio MCP shim (`frank-mcp`). A seat's entire tool surface
  is `submit` / `project` / `read`; no store, config, or socket path ever appears in a seat-visible
  surface.

## Install

Requires Go 1.22+ on Linux or macOS (the wire is a Unix domain socket).

```sh
git clone https://github.com/iwnlcern/frank
cd frank
go install ./cmd/frank ./cmd/frank-mcp
```

This installs `frank` (the conductor) and `frank-mcp` (the seat-side MCP shim) into your `GOBIN`.

## Quick start

**1. Initialize a store.** The store root is the governance domain — records, projections,
mailboxes, and pinned config all live under it:

```sh
cat > /tmp/engine.json <<'JSON'
{"gc_enabled":false,"segment_rotate_bytes":4194304}
JSON

frank -root /abs/path/team-store \
      -registry internal/fieldspec/registry.json \
      -engine-config /tmp/engine.json \
      -init
```

**2. Mint seats.** CLI minting is genesis-time only; once the conductor is live, minting happens
through operator-authored `seat_mint` records. Each mint prints the seat credential once — capture
it; credentials are never written to records or ordinary reads.

```sh
frank -root /abs/path/team-store -mint demo.planner -role planner
frank -root /abs/path/team-store -mint demo.implementer -role implementer
frank -root /abs/path/team-store -mint operator -role operator -operator
```

**3. Start the conductor.** One conductor per store root, enforced with an exclusive lock. Keep the
socket path short — long AF_UNIX paths fail at bind on macOS, and frank refuses paths of 100 bytes
or more with a typed error:

```sh
frank -root /abs/path/team-store -socket /tmp/team.sock
```

**4. Wire a seat.** A host session occupies a seat by launching `frank-mcp` with that seat's
credential — for example, as an MCP server entry:

```json
{
  "mcpServers": {
    "frank": {
      "command": "frank-mcp",
      "env": {
        "FRANK_SOCKET": "/tmp/team.sock",
        "FRANK_CREDENTIAL": "<credential from mint>"
      }
    }
  }
}
```

The seat's entire tool surface is **`submit` / `project` / `read`** — no store, config, or socket
path ever appears in a seat-visible surface. Operational detail (restart and recovery, live re-mint,
credential custody, wiring patterns) is in [`docs/ops.md`](docs/ops.md).

## What frank does *not* claim (yet)

Step-1 records **transport and provenance facts only**. A work claim carried through frank is
`self_reported` until the Step-2 observe layer lands (observation as a send-gate, evidence ladders,
done-predicates). Identity is **confusion-resistant, not theft-proof**: a same-uid local process
operating outside the tool surface can still reach files and sockets (the recorded D5 residual).
Forgery-robust-*by-construction* is a deliberately deferred milestone, not a current claim.

## Repository layout

```text
cmd/frank/         the conductor CLI: store init, seat mint, serve
cmd/frank-mcp/     stdio MCP shim — the seat-side wire (submit / project / read)
internal/          engine, store, fieldspec registry, lineage, gates, recovery, egress, ...
docs/ops.md        operations guide
docs/sprints/      per-slice build ledgers, plans, and gate evidence
test/              fixtures, crash/replay harness, seat-process integration tests
.relays/           the build's full relay trail (540 relays across six slices)
master-docs/       the governing team's workspace: charter, architecture-of-record, domain designs
```

## The build record

frank was built by a governed multi-agent team running the same discipline the product enforces —
six slices, each closed by an adversarial merge gate (close points `s1-close` through `s6-close` in
the ledgers). That working record ships with the repository, as both provenance and the first
end-to-end case study of the product's own discipline:

- `.relays/` — the full relay trail of the build: the slice teams' own traffic plus the governing
  team's dispatches and gates.
- `docs/sprints/` — the per-slice ledgers, plans, and gate evidence.
- `master-docs/` — the governing workspace itself: the standing team's charter
  (`master-docs/CLAUDE.md`), the architecture-of-record and the seven domain design docs
  (`master-docs/master/`), the design-cycle relay trail, and the ledgers — including the
  transport-findings ledger (`master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md`: 17 findings from
  running the team's own governance *through* frank, headlined by the lineage-livelock class that
  slice 6 then killed and replayed clean).

The coordination discipline itself — the role skills, the phase/relay protocol, and the relay linter
the team ran — is published separately as
[agentic-dev-team-skills](https://github.com/iwnlcern/agentic-dev-team-skills). It is exactly the
kind of convention-layer stack described in **Why** above, which is the point: frank is the trusted
middle a stack like that lacks.

## Development

```sh
go test ./...
```

The suite runs offline: unit tests, crash-atomicity and recovery replay fixtures (`test/replay`),
seat-process integration tests (`test/seatproc`), and sweep tests that pin the claim-boundary
language above to the code that earns it.

## Status

Step-1 (transport + provenance) closed 2026-07-08 at `s6-close`. Next: Step-2, the observe/evidence
layer. The sequencing plan is in [`master-docs/ROADMAP.md`](master-docs/ROADMAP.md).

## License

Apache-2.0 — see [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE).
