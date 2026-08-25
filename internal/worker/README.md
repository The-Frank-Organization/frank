# m-9 worker internals

This tree implements one bounded governed coding-agent turn. It does not own
provider wire translation, credentials, tool-authority policy, conductor
storage, app scheduling, or slice-close evidence export.

## Package map

- `catalog`: the exact eight-name F58 tool catalog and pinned identities.
- `wire`: closed frame envelopes and monotonic decimal counters.
- `journal`: the one-file durable session log, writer fence, crash recovery,
  round membership, and replay boundary.
- `executor`: F59 prepare/authorize/consume/record choreography.
- `tools`: the five local tools and their workspace/process boundaries.
- `relaytool`: native `relay.submit`, `relay.project`, and `relay.read` mapping.
- `turn`: admission, epoch fencing, assembly, attempt, tool-round, cancellation,
  and terminal state transitions.
- `provider`: one normalized provider attempt with E0 evidence derivation.
- `contextmgr`: pinned context tiers, compaction, and logical-surface digest.
- `resume`: settlement-manifest decoding and the five resume first actions.
- `runtime`: typed control/broker/provider assembly for one honest turn.
- `fake`: in-process m-10 and m-8 counterparts used only by E2 tests.

## Build, run, and test

From the `frank/` module root:

```sh
go build ./cmd/frank-worker
go test ./internal/worker/...
go test ./...
go vet ./...
```

The app owns live process startup. For local execution, construct a
`runtime.Runner` with explicit `Control`, `Broker`, `Provider`, tool `Backend`,
and wake-relay `ObjectiveResolver` dependencies. The end-to-end example is
`cmd/frank-worker/worker_test.go`; it proves durable replay but performs no real
provider egress.

## Trust seams

The worker treats parsed model tool calls as inert until the control peer
authorizes and consumes an F59 ticket. Tool outcomes are journaled through a
complete round marker before they are reported. A resume decision trusts only
the frozen full-key prefix and the two bounded trust windows. Push notifications
are advisory: durable project/read rediscovery remains authoritative.

The full requirement-to-test map, including partial and owner-external
counterpart gates, is in `FIXTURES.md`. Master selected Branch A, so this build
has no Branch-B exclusions. Live E3, m-10 receipt persistence, m-8 wire egress,
m-3 independent reconstruction, store export, merge, and slice close remain
outside this E2 implementation claim.
