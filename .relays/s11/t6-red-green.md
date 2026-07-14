# T6 integrated 8a hardening evidence

Base: `219e8db`

## RED

Commands:

```text
go test -count=1 ./internal/migrate -run '^TestApplyDeepClonesHeadersAndXFieldsBeforeInPlaceMigrator$' -v
go test -count=1 ./test/fixtures -run '^TestS11(StaleChoice|StructuralChoice)' -v
```

Observed expected failures before production code:

```text
--- FAIL: TestApplyDeepClonesHeadersAndXFieldsBeforeInPlaceMigrator
    migrate_test.go:107: migrated view = ... Headers:map[choices:[{"label":"Ship","value":"approve"}]] ... XFields:map[source:mutated]
FAIL github.com/jackli/frank/internal/migrate

test/fixtures/s11_8a_test.go:69:32: undefined: engine.SubmitHandlerWithMigration
test/fixtures/s11_8a_test.go:166:35: undefined: engine.SubmitHandlerWithMigration
FAIL github.com/jackli/frank/test/fixtures [build failed]
```

The first RED proves an in-place migrator aliases both source maps. The E2E RED
proves the live verdict path has no migration registry/guard seam. The pending
fixture binds the two byte-distinct reason tokens, no wake, deterministic
replacement identity, real process crash after the held signal, recovery to
the same replacement, and the reorder/representational-column GREEN case.

## GREEN

Commands:

```text
go test -count=1 ./internal/migrate -run '^TestApplyDeepClonesHeadersAndXFieldsBeforeInPlaceMigrator$' -v
go test -count=1 ./internal/engine -run '^TestS11StaleSchemaSuppressesOldCadenceAndReplacementRestartsIdentity$' -v
go test -count=1 ./test/fixtures -run '^TestS11(StaleChoice|StructuralChoice)' -v
go test -count=1 ./internal/engine ./internal/obligation ./internal/migrate ./internal/store ./internal/recover
```

Observed: all commands passed. The breaking-migration fixture proves the stale
candidate is `rejected` with `failing_edge: stale_choice_set`, appears in the
operator's bucket-D query, carries no projection intents, and does not resolve
the old gate. `CompleteAuto` consumes that committed durable intent, commits an
operator-visible `held` record with `failing_edge: stale_schema`, and only then
commits the deterministic new-identity gate/ODB/park. A real SIGKILL at
`stale_reissue_after_held` leaves the held record durable; normal recovery
replays the same replacement identity exactly once. The old decision is
resummon-ineligible, while the replacement's new decision ID creates fresh
cadence keys for the same seat and slot series.

The structural fixture migrates reordered rows with an added display-only
column, preserves π (`value -> label`), and resolves against the migrated view.
The source ODB bytes and in-memory table maps remain unchanged in both cases.
The final uncached `go test -count=1 ./... && go vet ./...` battery also passed.
