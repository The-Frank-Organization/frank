# frank

**A governed courier for multi-agent teams.** frank is a small trusted middle, written in Go, that
carries every agent-to-agent message through a single serialized commit path where **identity is
stamped by the channel, lineage is computed by the conductor** (frank's core process), **and
authority is enforced at commit** — never taken from the payload, never validated by convention. On top of that courier, a governed
coding-agent harness is now mid-build in the same tree.

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
That enforcement covers the governed surface; what it does and does not cover is spelled out plainly
in [What frank does *not* claim](#what-frank-does-not-claim-yet). The repository also ships its own
build record — the full relay trail and gate evidence of the governed multi-agent team that built it
(see [The build record](#the-build-record)).

## Features

### Transport + provenance (Step-1, closed)

- **Append-only record store** with crash-atomic multi-file commit (stage → fsync → rename;
  presence = committed) and recovery replay. A **single-threaded serialized commit loop** is the sole
  in-process writer for governance-surface mutations.
- **Channel-stamped FROM** — a **seat** (one agent's authenticated address on the courier) gets its
  identity from its authenticated channel (one active channel per credential, live re-mint with auth
  generations), never from what the seat writes.
- **Forms are the schema** — a declarative FieldSpec registry renders per-seat, per-phase forms, and
  the rendered form *is* the submit schema, pinned by a stable form digest. Config is pinned by a
  composite digest; changes go through an operator-authored `config_change` record, never silent edits.
- **Conductor-computed lineage** — the parent of a record is computed at commit from the sender's
  own read history intersected with the accepted graph; a seat's `parent_hint` is honored only when
  provable.
- **Typed terminal outcomes** — exactly `accepted`, `rejected`, `held`. Every bounce is a typed,
  recorded fact.
- **Human gates** — monotonic gate raising, gate categories with a fail-safe default, owed items and
  dispositions, scoped waivers with retraction, and a local-outbox-only rule for anything leaving the
  store, with an egress scanner at the drain.
- **A minimal wire** — Unix socket plus a stdio MCP shim (`frank-mcp`). A seat's entire tool surface
  is `submit` / `project` / `read`; no store, config, or socket path ever appears in a seat-visible
  surface.

### Observation + evidence (Step-2, closed)

- **Observe-as-send gate** (`internal/observe`) — a claim that names an observable is checked at the
  send boundary, not taken on faith; the check outcome is part of the record.
- **Evidence ladder E0–E4** with per-field `evidence_integrity` stamps: every field is marked
  `observed` (the conductor saw it happen) or `self_reported` (the sender said so). Rungs run from
  E0 — a bare assertion — upward, and a claim climbs only by conductor-side observation. The stamp
  is the honest boundary — the ladder upgrades claims by observation, never by assertion.
- **Executable claims** — a suite-class check executor runs pinned, lane-rooted, timeout-bounded
  check suites; "the tests pass" becomes something the conductor ran, not something a seat typed.
- **Ten executable laws** — the INV-CATALOG constitution ships as data
  (`test/invariants/catalog.v1.json`), pinned into every store at init and enforced by tests.

### The coding-agent harness (Step-3, mid-build)

Step-3 adds a barely-enough coding agent around the courier. Four new binaries ship in `cmd/` and
are honest about their state: **the end-to-end live agent loop is not yet assembled in the shipped
commands** — each piece runs and is tested at its own seam, and the wiring between them is being
built slice by slice.

```text
seats (any host) ── frank-mcp ──┐
                                ├─ unix socket ─► frank (the conductor) ─► append-only store
frank-worker ─── frank-broker ──┘      the governed relay plane: submit / project / read

frank-app ── supervises ─► frank-worker ── local tools (read / write / edit / apply_patch / bash)
    └─────── supervises ─► frank-connector ─► provider API      (bypasses the conductor)
```

- **`frank-app`** — the app control plane / supervisor: SQLite-backed run state, a frozen run
  manifest that fixes the eight tool names the worker may dispatch, one-shot dispatch tickets
  (design record F59 — an authorization is spent when used), and worker + connector supervision.
  The shipped binary's release binding is a **labeled development placeholder**
  (`sha256("frank-mvp-development")`), not a real release digest.
- **`frank-worker`** — the governed coding-agent runtime: `read` / `write` / `edit` / `apply_patch` /
  `bash` local tools behind a workspace-bounded backend, a uniform dispatch-authority path, the
  conductor as a native relay tool, and a JCS-canonical session journal with crash-tested recovery.
  Run standalone it exits 2 by design; it expects app-side wiring.
- **`frank-broker`** — holds the worker's conductor seat credential so the worker process never sees
  it: 0600/0700-hardened loads and a redacting log sink.
- **`frank-connector`** — the one process that holds provider credentials, with byte-exact
  `openai-responses.v1` request lowering. Provider traffic goes app-side through the connector and
  bypasses the conductor; the conductor stays a governed relay plane, not an app hub. (The courier's
  own trust boundary is unchanged by this split — see the D5 residual below.)

## Install

Requires Go 1.25+ on Linux or macOS (every wire is a Unix domain socket).

```sh
git clone https://github.com/The-Frank-Organization/frank
cd frank
go install ./cmd/frank ./cmd/frank-mcp
```

This installs the usable-standalone pair — `frank` (the conductor) and `frank-mcp` (the seat-side
MCP shim) — into your `GOBIN`. The Step-3 harness binaries build the same way but are mid-build (see
above).

## Quick start

Every command below was run end-to-end against this tree (2026-08-31) before landing in this README.
Run them from the cloned repo root, in order.

**1. Create a workspace and an engine config.** Serving requires a `version: 2` engine config with a
`supply` section: at least one **lane root** (an existing directory, given as an absolute,
symlink-free path — hence `pwd -P`, since `/tmp` is a symlink on macOS) and at least one named
**check suite** (an executable inside that lane). This is the Step-2 machinery: the conductor will
only serve if it knows where executable evidence comes from.

```sh
DEMO="$(mktemp -d)"; DEMO="$(cd "$DEMO" && pwd -P)"
mkdir -p "$DEMO/lane"
printf '#!/bin/sh\nexit 0\n' > "$DEMO/lane/check.sh"
chmod +x "$DEMO/lane/check.sh"

cat > "$DEMO/engine.json" <<JSON
{
  "version": 2,
  "gc_enabled": false,
  "segment_rotate_bytes": 4194304,
  "present_layers": {"observe": false},
  "supply": {
    "lane_roots": {"demo": "$DEMO/lane"},
    "schema_refs": {},
    "suites": {
      "demo-suite": {
        "lane": "demo",
        "command": "check.sh",
        "args": [],
        "timeout_class": "suite_bounded",
        "timeout_seconds": 60
      }
    }
  }
}
JSON
```

**2. Initialize a store.** The store root is the governance domain — records, projections,
mailboxes, and pinned config all live under it. Init pins three config sources by digest: the
FieldSpec registry, the engine config, and the invariant catalog (the ten executable laws ship
in-tree):

```sh
frank -root "$DEMO/store" \
      -registry internal/fieldspec/registry.json \
      -engine-config "$DEMO/engine.json" \
      -catalog test/invariants/catalog.v1.json \
      -init
```

**3. Mint seats.** CLI minting is genesis-time only; once the conductor is live, minting happens
through operator-authored `seat_mint` records. Each mint prints `credential=<64 hex>` exactly once —
capture it; credentials are never written to records or ordinary reads:

```sh
PLANNER_CRED="$(frank -root "$DEMO/store" -mint demo.planner -role planner | cut -d= -f2)"
OPERATOR_CRED="$(frank -root "$DEMO/store" -mint operator -role operator -operator | cut -d= -f2)"
```

**4. Start the conductor** — in a second terminal (same `DEMO`). One conductor per store root,
enforced with an exclusive lock. Keep the socket path short — long AF_UNIX paths fail at bind on
macOS, and frank refuses paths of 100 bytes or more with a typed error:

```sh
frank -root "$DEMO/store" -socket /tmp/frank-demo.sock
```

**5. Speak as a seat.** `frank-mcp` is a stdio MCP shim: one process per host session, one seat
credential, tool surface exactly `submit` / `project` / `read`. A tiny shell helper makes it
scriptable (each call opens a fresh shim session):

```sh
export FRANK_SOCKET=/tmp/frank-demo.sock
export FRANK_CREDENTIAL="$PLANNER_CRED"

mcp() {
  printf '%s\n' \
    '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"quickstart","version":"0"}}}' \
    '{"jsonrpc":"2.0","method":"notifications/initialized"}' \
    "{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"$1\",\"params\":$2}" \
  | frank-mcp | tail -n 1
}
```

The rendered form *is* the submit schema, so first fetch the current form digest from the tool
schema (re-fetch it before every submit — when the form refreshes, a stale digest gets a typed
re-render, not a silent accept):

```sh
DIGEST="$(mcp tools/list '{}' \
  | grep -o '"form_digest":{"const":"[0-9a-f]\{64\}"' | grep -o '[0-9a-f]\{64\}')"
```

Now send a relay. `PHASE` and `CEREMONY_TIER` are fields from the pinned form registry — the tier
sets how much a relay must prove, and a SITREP at tier `medium` must declare an evidence target.
(The `mcp` helper prints the full JSON-RPC envelope; the payload shown below each call is the
`result.content[0].text` value inside it.)

```sh
mcp tools/call "$(cat <<JSON
{"name":"submit","arguments":{"headers":{"PHASE":"SITREP","CEREMONY_TIER":"medium","EVIDENCE_TARGET":"E1","SUBJECT":"hello from the quick start"},"to":"operator","body":"First relay through the governed courier.","form_digest":"$DIGEST"}}
JSON
)"
```

```text
{"state":"accepted","decision_state":"accepted","post_commit_state":"complete","relay_id":"relay-…","intake_id":"intake-000001"}
```

**6. See a typed rejection.** Leave out `EVIDENCE_TARGET` and the relay is refused as a typed,
recorded fact — with a relay id, not an error string that vanishes:

```sh
mcp tools/call "$(cat <<JSON
{"name":"submit","arguments":{"headers":{"PHASE":"SITREP","CEREMONY_TIER":"medium","SUBJECT":"missing evidence target"},"to":"operator","form_digest":"$DIGEST"}}
JSON
)"
```

```text
{"state":"rejected","decision_state":"rejected","post_commit_state":"complete","relay_id":"relay-…","intake_id":"intake-000002","detail":"EVIDENCE_TARGET:required"}
```

**7. Read the other side.** The operator seat projects its mailbox and reads the accepted relay
(substitute the `relay_id` from your accepted submit):

```sh
FRANK_CREDENTIAL="$OPERATOR_CRED" mcp tools/call '{"name":"project","arguments":{}}'
FRANK_CREDENTIAL="$OPERATOR_CRED" mcp tools/call '{"name":"read","arguments":{"relay_id":"<relay_id from step 5>"}}'
```

The read returns the committed record: channel-stamped `from`, conductor-assigned lineage and
delivery state, and a checksum. Stop the conductor with Ctrl-C when done; on restart, recovery
replays the store cleanly.

To wire a real host session instead of a shell, register `frank-mcp` as an MCP server:

```json
{
  "mcpServers": {
    "frank": {
      "command": "frank-mcp",
      "env": {
        "FRANK_SOCKET": "/tmp/frank-demo.sock",
        "FRANK_CREDENTIAL": "<credential from mint>"
      }
    }
  }
}
```

Operational detail (restart and recovery, live re-mint, credential custody, wiring patterns) is in
[`docs/ops.md`](docs/ops.md).

## What frank does *not* claim (yet)

The courier records **transport and provenance facts only** about message content. Since Step-2, the
observe layer stamps what it actually observed — fields carrying `evidence_integrity: observed` were
checked by the conductor's own executors — but every field it did not observe remains
`self_reported`, and the record says which is which.

Identity is **confusion-resistant, not theft-proof**: a same-uid local process operating outside the
tool surface can still reach files and sockets (the recorded D5 residual).
Forgery-robust-*by-construction* is a deliberately deferred milestone, not a current claim.

There is **no sandbox in the MVP**. The worker's `bash` tool runs with the invoking user's ambient
authority — `internal/executor/executor.go` pins this residual in code: same-uid ambient filesystem,
network, and process access remains possible without an OS sandbox.

The whole system is **local, single-host, same-trust-domain**: every wire is a Unix domain socket
and nothing listens on TCP. It is **not hardened for external, untrusted, or multi-tenant use** —
that is a standing hard blocker, on the record, before any such deployment.

## Repository layout

```text
cmd/frank/           the conductor CLI: store init, seat mint, serve
cmd/frank-mcp/       stdio MCP shim — the seat-side wire (submit / project / read)
cmd/frank-app/       app control plane / supervisor: run state, run manifest, dispatch tickets (mid-build)
cmd/frank-worker/    governed coding-agent runtime: local tools + native relay tool (mid-build)
cmd/frank-broker/    seat-credential broker — the worker never holds its own conductor credential
cmd/frank-connector/ provider connector — holds provider credentials, speaks the provider wire
internal/            engine, store, fieldspec registry, lineage, gates, observe, executor, recovery, egress, ...
docs/ops.md          operations guide
docs/sprints/        per-slice build ledgers, plans, and gate evidence
test/                fixtures, invariant catalog, crash/replay harness, seam contract battery
.relays/             the build's full relay trail (slices s1 through s16a)
```

## The build record

frank was built by a governed multi-agent team running the same discipline the product enforces —
slice by slice, each closed by an adversarial merge gate (close points `s1-close` onward in the
ledgers), across Step-1, Step-2, and the Step-3 slices in flight. That working record ships with the
repository, as both provenance and the first end-to-end case study of the product's own discipline:

- `.relays/` — the full relay trail of the build: the slice teams' own traffic plus the governing
  team's dispatches and gates.
- `docs/sprints/` — the per-slice ledgers, plans, and gate evidence.

The coordination discipline itself — the role skills, the phase/relay protocol, and the relay linter
the team ran — is published separately as
[agentic-dev-team-skills](https://github.com/iwnlcern/agentic-dev-team-skills). It is exactly the
kind of convention-layer stack described in **Why** above, which is the point: frank is the trusted
middle a stack like that lacks.

## Development

```sh
go test ./...
```

The suite runs offline in about four minutes, dominated by `test/fixtures`: unit tests,
crash-atomicity and recovery replay fixtures (`test/replay`), real-conductor subprocess ceremonies,
invariant-catalog checks, and sweep tests that pin the claim-boundary language above to the code
that earns it.

## Status

Step-1 (transport + provenance) closed 2026-07-08 at `s6-close`. Step-2 (observation + evidence)
closed 2026-07-14 at `s11-close`. Step-3 — the MVP coding-agent harness — is mid-build: the courier
is live and governed, the harness binaries ship piecewise, and the end-to-end agent loop is still
being assembled.

## License

Apache-2.0 — see [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE).
