# T2 bucket-C projection evidence

Base: `a30df05`

## RED

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11BucketCIsOperatorCCFYIWithoutDecisionObligation$' -v
```

Observed expected failure before production code:

```text
test/fixtures/s11_buckets_test.go:82:17: st.ProjectBucketC undefined (type *store.Store has no field or method ProjectBucketC)
FAIL github.com/jackli/frank/test/fixtures [build failed]
```

The fixture distinguishes operator-on-CC from operator-on-TO and requires the
CC-only record to remain informational: visible to the operator but without an
ODB or park obligation.

## GREEN

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11BucketCIsOperatorCCFYIWithoutDecisionObligation$' -v
```

Observed:

```text
=== RUN   TestS11BucketCIsOperatorCCFYIWithoutDecisionObligation
--- PASS: TestS11BucketCIsOperatorCCFYIWithoutDecisionObligation (0.09s)
PASS
ok github.com/jackli/frank/test/fixtures 0.901s
```

The saved query consumes canonical TO/CC address lists, excludes an operator
TO even when operator is also on CC, and leaves the CC-only item in ordinary
operator delivery without an ODB.

## Commit-point full battery

Command:

```text
go test -count=1 ./... && go vet ./... && git diff --check
```

Observed: exit 0. All packages passed; `test/fixtures` completed in
`123.072s`; `go vet ./...` and `git diff --check` were clean.
