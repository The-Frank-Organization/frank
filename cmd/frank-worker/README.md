# frank-worker

`frank-worker` is the m-9 governed model-worker process. The app control plane
(m-10) owns lifecycle, authority, and broker assignment; the connector (m-8)
owns provider traffic. This command does not synthesize either peer and does
not accept provider credentials.

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
injected. The executable runtime seam is `internal/worker/runtime.Runner`; the
in-process peers in `internal/worker/fake` are test-only counterparts.

Run the command and governed-turn tests with:

```sh
go test ./cmd/frank-worker ./internal/worker/fake
```

Run the complete repository battery with `go test ./...`.

## Seam map

- m-10 sends the typed hello/assignment, owns F59 authority decisions and turn
  terminalization, and receives attach and wake-forward reports.
- m-7's broker validates the `{run_id, generation_id, turn_epoch}` attach tuple
  and returns only a connection-scoped capability.
- m-8 receives opaque provider requests and returns normalized provider events
  and tool calls; the worker never receives provider credentials.
- m-9 owns admission, context assembly, attempt sequencing, local/relay tool
  requests, journal durability, recovery, and the governed-turn state machine.

Master selected plan Branch A. The shared conductor facade and retained MCP
consumer therefore landed in the authorized seven-file fence; there are no
Branch-B exclusions in this build.
