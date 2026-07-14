# T7 bucket matrix and known-A negative evidence

Base: `99e945c`

## Sequence record

T7 is a fixture-registration task consuming the already-landed T1–T4 and T6
mechanisms; it has no licensed production-code step. Its consolidated matrix
therefore passed on its first run. The behavior REDs preceded their supplies:

- T1: `ProjectBucketB` was undefined before `a30df05`.
- T2: `ProjectBucketC` was undefined before `1db30c5`.
- T3: `ProjectBucketD` was undefined before `1b1ed68`.
- T4: `bounced_repair` and `egress_blocked` were absent before `219e8db`.
- T6: `stale_choice_set` was absent from the D query before `99e945c`.

The T7 table itself contains a negative for every bucket boundary: rejected A
and B/C-shaped records return through D; accepted or held records never fall
through to D; held remains A; egress is A only at the accepted-A stage; operator
TO excludes C; and non-B terminals exclude the B digest.

## GREEN

Command:

```text
go test -count=1 ./test/fixtures -run '^TestS11(BucketTerminalAndFailingEdgeMatrix|KnownABPickIsRaisedRecordedAndNeverAbsorbed)$' -v
```

Observed: all 14 matrix rows and the ③ known-A fixture passed. The known-A
fixture traverses the live submit validator: an agent `routing` B-pick over an
`authz_security` A floor is accepted only after rewriting `gate_category` to
`authz_security`, recording `gate_category_raised=yes` and
`gate_category_pick=routing`; the record is absent from bucket B and produces
an operator ODB in bucket A.

Rail A is closed at every acceptance-meaning boundary; Rail B reuses the
existing FieldSpec, obligation, and saved-query mechanisms. No runtime or
schema bytes changed in T7.

The uncached `go test -count=1 ./... && go vet ./...` battery passed before
the T7 commit boundary.
