# T4 seven-state FSM evidence

Base: `1b1ed68`

## RED

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11FSMAddsBouncedRepairAndFixtureScopedEgressBlocked$' -v
```

Observed expected failure before production code:

```text
test/fixtures/s11_fsm_test.go:28:96: undefined: engine.GateBouncedRepair
test/fixtures/s11_fsm_test.go:29:57: undefined: engine.GateBouncedRepair
test/fixtures/s11_fsm_test.go:41:102: undefined: engine.GateEgressBlocked
test/fixtures/s11_fsm_test.go:42:56: undefined: engine.GateEgressBlocked
FAIL github.com/jackli/frank/test/fixtures [build failed]
```

The RED proves both missing FSM labels. The fixture binds `bounced_repair` to
a rejected acceptance-stage veto and binds `egress_blocked` to an accepted A
gate that parks locally and emits only a local resummon command; no live away
send, redaction, or verdict is introduced.

## GREEN

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11FSMAddsBouncedRepairAndFixtureScopedEgressBlocked$' -v
```

Observed:

```text
=== RUN   TestS11FSMAddsBouncedRepairAndFixtureScopedEgressBlocked
--- PASS: TestS11FSMAddsBouncedRepairAndFixtureScopedEgressBlocked (0.11s)
PASS
ok github.com/jackli/frank/test/fixtures 0.889s
```

Accepted resolutions still dominate recovery state. Otherwise an acceptance
bounce derives `bounced_repair`; an accepted blocked egress derives
`egress_blocked`; ordinary resummon and park states retain their prior order.
The existing CompleteAuto and ResummonHandler mechanisms supply local park and
local channel escalation without adding the deferred away trigger.

## Commit-point full battery

Commands:

```text
go test -count=1 ./internal/engine -run 'TestS10Resummon|TestS10ProductionScheduler'
go test -count=1 ./test/fixtures -run 'Test(S10ExitLeg3|S10OnlyAcceptedAGate|S11FSM)'
go test -count=1 ./... && go vet ./... && git diff --check
```

Observed: all commands exited 0. The full run passed all packages;
`test/fixtures` completed in `122.892s`; vet and diff checks were clean.
