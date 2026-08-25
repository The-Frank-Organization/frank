# seatclient

`seatclient` is the shared, transport-neutral seat-facing conductor boundary.
It contains no raw store paths and never receives, stores, logs, or transports
seat credential bytes.

## Packages and seams

- `formschema` maps the live m-2 form to strict native/MCP tool schemas,
  validates the exact submit/project/read argument shapes, and provides stable
  F58 schema identities. Its form digest and volatile fields are refreshed at
  the defined F-1/F-2 boundaries.
- `conduct` exposes only `relay.submit`, `relay.project`, and `relay.read` over
  an already-authenticated transport. It also supplies bounded attach,
  reconnect classification, mailbox rediscovery, and advisory wake forwarding.

Authentication stays with the caller. In the retained MCP frontend,
`ensureClient` remains the sole credential acquirer and performs the preserved
close, re-authenticate, one-retry choreography. Native worker and MCP requests
then traverse the same `conduct` facade bytes.

## Build and test

From the `frank/` module root:

```sh
go test ./internal/seatclient/...
go test ./cmd/frank-mcp
go test ./...
go vet ./...
```

There is no standalone seatclient executable. Callers inject an authenticated
`conduct.Transport`; worker-side lifecycle code injects its app/broker peers.

Master selected plan Branch A. The authorized MCP consumer refactor is present,
including shared-schema, reconnect, H-16 reachability, and native/MCP parity
tests. There are no Branch-B exclusions in this build.
