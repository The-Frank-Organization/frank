# frank-worker

`frank-worker` is the governed model-worker process. The app control plane
(`frank-app`) owns lifecycle, authority, and broker assignment; the connector
(`frank-connector`) owns provider traffic. This command does not synthesize
either peer and does not accept provider credentials.

## Build, run, and test

Build the process from the `frank/` module root:

```sh
go build ./cmd/frank-worker
```

The standalone entry point is an integration guard, not a development server:

```sh
go run ./cmd/frank-worker
```

It exits with status 2 until the app-owned control and provider transports are
injected; live app-side wiring is mid-build. The executable runtime seam is
`internal/worker/runtime.Runner`; the in-process peers in `internal/worker/fake`
are test-only counterparts.

Run the command and governed-turn tests with:

```sh
go test ./cmd/frank-worker ./internal/worker/fake
```

Run the complete repository battery with `go test ./...`.

## Seam map

- The app control plane sends the typed hello/assignment, owns F59 authority
  decisions and turn terminalization, and receives attach and wake-forward
  reports.
- The broker validates the `{run_id, generation_id, turn_epoch}` attach tuple
  and returns only a connection-scoped capability.
- The connector receives opaque provider requests and returns normalized
  provider events and tool calls; the worker never receives provider
  credentials.
- The worker owns admission, context assembly, attempt sequencing, local/relay
  tool requests, journal durability, recovery, and the governed-turn state
  machine.
