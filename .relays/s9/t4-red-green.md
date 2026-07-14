# T4 opaque-lane and lane_vcs evidence

Base task commit: `0f1aa42`

m-7 owner-byte source:
`master/relays/s9-build-fidelity-m7/SITREP-planner-20260713-184136.md`

## Anchor gate

Before editing, these four owner paths were byte-identical between
`39474d0` and `HEAD`: `internal/config/config.go`, `cmd/frank/main.go`,
`test/fixtures/s2setup_test.go`, and
`test/invariants/store_recovery_test.go`. E1-E10 therefore had no collision
with T1-T3 and were carried verbatim in commit `db9a166`.

## RED

The m-3 consumer tests initially failed to compile because
`RegistryEnv.LaneVCS` did not exist. After that seam landed, the FX-VCS tests
failed to compile because `config.Supply.LaneVCS` did not exist. These REDs
covered the two owner boundaries independently.

## GREEN

The final behavior proves:

- only a governed `lane_vcs: none` reaches the honestly labeled
  `opaque-lane-no-vantage` E0/self-reported branch;
- `git` and pre-v3 nil continue through the locked Section 13 observation
  floor and never opaque-accept;
- an ungoverned root is a typed machinery/config fault;
- root health runs in the detachable T1 worker, so a blocked filesystem stage
  returns a typed timeout without blocking the serialized caller;
- `RegistryEnv.LaneVCS`, `Pinned.Supply.LaneVCS`, and
  `Engine.Supply.LaneVCS` are cloned without aliasing, while v2 nil remains
  nil;
- v3 shape is total and closed to `git|none`; v2 smuggling, v3 omissions,
  wrong types/values, extra/missing lanes, downgrade/skip transitions, and v4
  reader-ceiling inputs fail closed;
- v3 dogfood and the dedicated v2 adoption/residency paths both execute.

Verification commands:

```text
go test -race ./internal/observe -run 'Test(DeclaredNone|UndeclaredOrGit|UngovernedRoot|RegistryEnvLaneVCS|AbsenceFloor)' -count=1
go test -race ./internal/config -run '^TestLaneVCS' -count=1
rg -n 'marker_absent|marker_present' internal/observe
rg -n 'EvalSymlinks|os\.Open|\.Stat\(|Readdirnames' internal/observe/checks_base.go
git diff --check
go test ./... -count=1
```

Observed: both race suites passed; both forbidden text sweeps returned no
matches; whitespace validation was clean; every package passed in the uncached
full battery, with `test/fixtures` completing in `133.725s`.
