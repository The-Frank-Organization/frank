# T1 shared FS worker evidence

Base: `39474d0`

## RED

Command: `go test ./internal/observe -run TestFSWorker -v`

Observed expected failure before production code:

```text
internal/observe/fs_worker_test.go:12:14: undefined: FSStage
internal/observe/fs_worker_test.go:30:5: unknown field FSStageHook in struct literal of type RegistryEnv
FAIL github.com/jackli/frank/internal/observe [build failed]
```

Read-file compatibility RED:

```text
=== RUN   TestReadFileRootFailureKeepsLandedMachineryClass
root failure detail = "check-machinery-fs-root-open", want landed read-file machinery class
--- FAIL: TestReadFileRootFailureKeepsLandedMachineryClass
```

## GREEN

Commands:

```text
go test ./internal/observe -run 'Test(FSWorker|ReadFileRootFailure)' -count=1 -v
go test ./test/fixtures -run 'Test(S8ReadFile|S8E1|S8Decision2|S10E1LongRead|S10ReadCompletion|S10RootCancellation)' -count=1 -timeout 45s
go test -race ./internal/observe -run 'Test(FSWorker|ReadFileRootFailure)' -count=1
go test -race ./test/fixtures -run 'Test(S8ReadFileBreaker|S8ReadFileConfinement|S8ReadFileDetaches)' -count=1 -timeout 45s
```

Observed:

```text
PASS: TestFSWorkerDetachesEveryStage (root-open, directory-read, file-open, metadata, read, close)
PASS: TestFSWorkerBreakerBoundsWorkers
PASS: TestFSWorkerComponentSwapNoOutsideBytes
PASS: TestReadFileRootFailureKeepsLandedMachineryClass
ok github.com/jackli/frank/test/fixtures
ok github.com/jackli/frank/internal/observe (race)
ok github.com/jackli/frank/test/fixtures (race)
```

The initially hanging legacy confinement/detachment probes were reproduced with timed stack dumps. Root cause was lost legacy hook sequencing during extraction; the composite worker hook now preserves `before -> root open -> component opens -> metadata -> read` while the new shared hook additionally covers directory-read and close.

## Commit-point full battery

Command: `go test ./... -count=1`

Observed: all packages green; `test/fixtures` completed its nested dogfood batteries in `123.751s`. `git diff --check` was clean before commit.
