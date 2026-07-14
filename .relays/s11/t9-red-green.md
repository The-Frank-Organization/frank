# T9 G4 operator-config cadence evidence

## RED — production config rejects the G4 surface

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11G4Cadence' -v
```

Observed before implementation: FAIL. Both independent timer cases stop at
store initialization with `config-load: engine`, proving that the current
operator config schema cannot carry either G4 cadence. The negative that adds
`auto_approve` is already rejected, pinning the no-verdict-policy boundary.

## GREEN

The engine member now accepts one exact `resummon_cadence` object containing
only `no_response_seconds` and `answered_but_stalled_seconds`. Omission preserves
the shipped one-hour defaults; explicit zero remains a valid immediate timer.
Production derives one immutable cadence from the startup-pinned config and
uses it both at startup and after every accepted commit.

Command:

```text
go test -count=1 ./internal/config ./internal/engine
go test -count=1 ./test/fixtures -run '^TestS11G4Cadence' -v
! rg -n 'DefaultResummonCadence' internal/engine cmd/frank
```

Observed: GREEN. The production binary independently fires each configured
timer at zero while the other remains at one hour. Both emissions are
resummon commands only; neither creates an accepted gate resolution. A config
object containing `auto_approve` is rejected as `config-load`.

Full battery: `go test -count=1 ./... && go vet ./...` GREEN.
