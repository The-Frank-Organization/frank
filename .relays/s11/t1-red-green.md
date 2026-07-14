# T1 bucket-B projection evidence

Base: `d91fcfb340b029c39c8493084ce2f227409aa546`

## RED

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11BucketBIsLiveNonInterruptingAndRaiseOnly$' -v
```

Observed expected failure before production code:

```text
test/fixtures/s11_buckets_test.go:26:17: st.ProjectBucketB undefined (type *store.Store has no field or method ProjectBucketB)
test/fixtures/s11_buckets_test.go:45:16: st.ProjectBucketB undefined (type *store.Store has no field or method ProjectBucketB)
FAIL github.com/jackli/frank/test/fixtures [build failed]
```

The RED proves the saved-query surface is absent. The fixture also requires a
B-set record to remain out of the operator ODB queue and a raised A record to
remain absent from B while producing an operator ODB.

## GREEN

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11BucketBIsLiveNonInterruptingAndRaiseOnly$' -v
```

Observed:

```text
=== RUN   TestS11BucketBIsLiveNonInterruptingAndRaiseOnly
--- PASS: TestS11BucketBIsLiveNonInterruptingAndRaiseOnly (0.16s)
PASS
ok github.com/jackli/frank/test/fixtures 0.914s
```

The saved query consumes the pinned registry's B membership. A B record is
returned without creating an ODB; a category rewritten to A is excluded from
the B view and follows the existing operator ODB path.

## Commit-point full battery

Command:

```text
go test -count=1 ./... && go vet ./... && git diff --check
```

Observed: exit 0. All packages passed; `test/fixtures` completed in
`124.631s`; `go vet ./...` and `git diff --check` were clean.
