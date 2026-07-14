# s11 optional findings fold — RED/GREEN evidence

Reviewed base: `547ada9aa89b6edcb98769ba27f0418439236441`

Scope: optional finding 1 (bind the acceptance-bounce and Bucket-D edge relation) and optional finding 2 (assert the replacement decision's scheduler-assigned cadence slots). Findings 3–11 remain recorded no-change items.

## RED

After adding only the two tests:

```text
$ go test -count=1 ./internal/bounce ./internal/engine -run 'TestBucketDEdgesExtendAcceptanceBouncesOnlyWithStaleChoiceSet|TestS11StaleSchemaSuppressesOldCadenceAndReplacementRestartsIdentity'
# github.com/jackli/frank/internal/bounce_test [github.com/jackli/frank/internal/bounce.test]
internal/bounce/edges_test.go:17:14: undefined: bounce.AcceptanceBounceEdge
internal/bounce/edges_test.go:20:14: undefined: bounce.BucketDFailingEdge
...
# github.com/jackli/frank/internal/engine [github.com/jackli/frank/internal/engine.test]
internal/engine/resummon_test.go:200:23: undefined: g4ResummonInputs
FAIL
```

Both failures were the intended missing seams, not assertion or fixture errors.

## In-scope integration correction

The first classifier placement in `internal/bounce` exposed an import cycle because the existing formatter reaches `store` through `lineage -> tables -> intake`. The implementation was moved, without widening FOLD_SCOPE, to a single named classifier in `internal/store/projections.go`; `internal/engine/fsm.go` consumes that classifier and the external relation test proves the intended subset. No edge membership changed.

## Focused GREEN

```text
$ go test -count=1 ./internal/bounce ./internal/engine -run 'TestBucketDEdgesExtendAcceptanceBouncesOnlyWithStaleChoiceSet|TestS11StaleSchemaSuppressesOldCadenceAndReplacementRestartsIdentity'
ok  github.com/jackli/frank/internal/bounce  0.264s
ok  github.com/jackli/frank/internal/engine  0.511s
```

## Review-requested targeted GREEN

```text
$ go test -count=1 ./internal/bounce ./internal/store ./internal/engine
ok  github.com/jackli/frank/internal/bounce  0.204s
ok  github.com/jackli/frank/internal/store  1.292s
ok  github.com/jackli/frank/internal/engine  2.045s
$ go test -count=1 ./test/fixtures -run '^TestS11'
ok  github.com/jackli/frank/test/fixtures  3.662s
$ go vet ./internal/bounce ./internal/store ./internal/engine ./test/fixtures
$ git diff --check
```

All commands exited 0. This is E2 local proof; no runtime/live or merge claim is made.
