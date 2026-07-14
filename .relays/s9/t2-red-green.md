# T2 find-references evidence

Base task commit: `ba26b27`

## RED

Command: `go test ./internal/observe -run TestFindReferences -count=1 -v`

Observed expected pre-implementation compile failure:

```text
registry_test.go:34:15: reg.executeFindReferences undefined
registry_test.go:65:10: undefined: findRefLimits
registry_test.go:90:75: unknown field findRefLimits in struct literal of type RegistryEnv
FAIL github.com/jackli/frank/internal/observe [build failed]
```

## GREEN

Commands:

```text
go test ./internal/observe -run TestFindReferences -count=1 -v
go test -race ./internal/observe -run 'Test(FSWorker|FindReferences|ReadFileRootFailure)' -count=1
go test ./test/fixtures -run 'Test(S8CheckRegistry|S8ReadFile|S8E1|S8Decision2|S10E1LongRead|S10ReadCompletion|S10RootCancellation)' -count=1 -timeout 60s
```

Observed green behaviors:

- complete zero count passes at E1; nonzero count fails with the symbolic class;
- identifier-boundary matching distinguishes `foo`, `foobar`, and `foo.bar`;
- binary files, symlinks, invalid UTF-8, and every numeric ceiling take their declared rows;
- exactly-at-ceiling passes and ceiling-plus-one fails closed;
- count saturation is bounded;
- timeout trips the same-lane breaker while preserving other lanes;
- no path or filename appears in the verdict;
- ungoverned lanes are typed `lane-ungoverned` refusals.

## Commit-point full battery

Command: `go test ./... -count=1`

Observed: all packages green; `test/fixtures` completed in `128.262s`. `git diff --check` was clean before commit.
