# s7 INV-CATALOG verification

Candidate implementation commit: `eaaf5f0`

## Contracted names

```text
TestLawIntakeOutcomeOneToOne
TestLawDerivedOnlyActivation
TestLawSoleGovernedWriter
TestLawPathHygiene
TestLawCanonicalWins
TestLawOnePivotPerMutation
TestLawRebuildBeforeOpen
TestLawTerminalEnumByteExact
TestLawThreeVerbSurface
TestLawR2NoModelPredicate
```

Exactly ten top-level `TestLaw*` entries are present.

## Invariant package

`go test -count=1 ./test/invariants`:

```text
ok      github.com/jackli/frank/test/invariants    1.105s
```

`go test -race -count=1 ./test/invariants`:

```text
ok      github.com/jackli/frank/test/invariants    3.862s
```

`go test -count=3 ./test/invariants`:

```text
ok      github.com/jackli/frank/test/invariants    3.446s
```

## Full repository first pass

`go test -count=1 ./...` exited 0:

```text
?       github.com/jackli/frank/cmd/frank                   [no test files]
ok      github.com/jackli/frank/cmd/frank-mcp               0.343s
ok      github.com/jackli/frank/internal/bounce              0.731s
ok      github.com/jackli/frank/internal/channel             0.853s
ok      github.com/jackli/frank/internal/config              0.943s
ok      github.com/jackli/frank/internal/crashpoint          1.163s
ok      github.com/jackli/frank/internal/egress              1.316s
ok      github.com/jackli/frank/internal/engine              3.313s
ok      github.com/jackli/frank/internal/fieldspec           1.498s
ok      github.com/jackli/frank/internal/fsio                1.900s
ok      github.com/jackli/frank/internal/gate                2.736s
ok      github.com/jackli/frank/internal/gc                  2.959s
ok      github.com/jackli/frank/internal/intake              6.030s
ok      github.com/jackli/frank/internal/lineage             2.224s
ok      github.com/jackli/frank/internal/migrate             2.315s
ok      github.com/jackli/frank/internal/obligation          2.785s
ok      github.com/jackli/frank/internal/record              2.382s
ok      github.com/jackli/frank/internal/recover             4.203s
ok      github.com/jackli/frank/internal/seat                2.485s
ok      github.com/jackli/frank/internal/store               4.278s
ok      github.com/jackli/frank/internal/tables              1.814s
ok      github.com/jackli/frank/test/fixtures                27.939s
ok      github.com/jackli/frank/test/invariants              3.557s
ok      github.com/jackli/frank/test/replay                  1.775s
ok      github.com/jackli/frank/test/replay/dogfood          1.761s
ok      github.com/jackli/frank/test/replay/zeroloss         2.645s
?       github.com/jackli/frank/test/seatproc                [no test files]
```

Count: 25 tested packages, 2 no-test-file packages, zero failures.

`go vet ./...` exited 0 with no output.

## Scope

Candidate diff paths are confined to:

- `test/invariants/**`
- `docs/sprints/2026-07-10-s7-inv-catalog/**`
- `.relays/s7/**`

No production source, registry, or record-kind path changed.

The final-tip repeat is recorded in the IMPL relay after this evidence file is committed.
