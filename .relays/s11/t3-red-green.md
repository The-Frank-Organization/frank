# T3 bucket-D projection and precedence evidence

Base: `1db30c5`

## RED

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11BucketDIsAuthorFacingAndEgressBlockedStaysA$' -v
```

Observed expected failure before production code:

```text
test/fixtures/s11_buckets_test.go:130:17: st.ProjectBucketD undefined (type *store.Store has no field or method ProjectBucketD)
test/fixtures/s11_buckets_test.go:137:20: st.ProjectBucketD undefined (type *store.Store has no field or method ProjectBucketD)
FAIL github.com/jackli/frank/test/fixtures [build failed]
```

The fixture requires rejected acceptance-stage records to return only through
the author-facing D query, requires form-validation rejections to carry the
system-authored `failing_edge`, and keeps an accepted egress-blocked A gate on
the local operator ODB path rather than D.

## GREEN

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11BucketDIsAuthorFacingAndEgressBlockedStaysA$' -v
```

Observed:

```text
=== RUN   TestS11BucketDIsAuthorFacingAndEgressBlockedStaysA
--- PASS: TestS11BucketDIsAuthorFacingAndEgressBlockedStaysA (0.19s)
PASS
ok github.com/jackli/frank/test/fixtures 0.984s
```

The handler now overwrites acceptance-stage rejection provenance with a
system-authored edge. The D query accepts only the four locked acceptance
edges for the authoring seat. `egress` is excluded; an accepted blocked A gate
still produces its local operator ODB.

## Commit-point full battery

Commands:

```text
go test -count=1 ./internal/engine -run 'TestSubmitHandler'
go test -count=1 ./test/fixtures -run 'Test(S5SubmitGuard|S10OperatorReply|S11BucketD)'
go test -count=1 ./... && go vet ./... && git diff --check
```

Observed: all commands exited 0. The full run passed all packages;
`test/fixtures` completed in `123.229s`; vet and diff checks were clean.
