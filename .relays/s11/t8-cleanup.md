# T8 cleanup-card evidence

Base: `e076bc4`

T8 scope: eight of nine — item 2 rescoped by master
(`s11-build-escalate-fence/RECONCILE-…-143010`), carried post-Step-2 to
m-7+m-3.

Each retained item lands in card order with its own commit and a green battery
before the next item begins.

## Item 1 — generic prompter

Characterization baseline: the T7-boundary uncached full suite was green.
The approval and expiry duplicate-waiter and operator-only replay/live fixtures
were rerun after extracting their shared pending-owner, emit, await, fail-safe,
dedupe, and resolve lifecycle into `genericPrompter[D]`; all passed. Approval
still defaults deny, expiry still defaults kill, the verbs/content hashes/gate
IDs are unchanged, and the domain-specific stored-decision lookups remain in
their adapters.

Targeted command:

```text
go test -count=1 ./internal/engine
go test -count=1 ./test/fixtures -run '^TestS10(ApprovalPrompterSharesDuplicateGateWaiters|ExpiryPrompterSharesDuplicateGateWaiters|ExpiryPrompterRejectsNonOperatorExtendAcrossReplayAndLivePaths)$' -v
```

Observed: GREEN.

Between-item battery: `go test -count=1 ./... && go vet ./...` GREEN.

## Item 6 — shared system-to-operator record builder

Fresh post-item-3 census before editing found five operator-addressed builders:
approval, expiry, resummon, the consolidated ODB builder, and T6's held
`stale_schema` signal. The old separate `engine/odb.go` builder site had folded
into obligation at item 3; the T6 held signal is also operator-addressed, so the
verified post-item-3 count remained five without absorbing any non-operator
record. All five now call `obligation.SystemOperatorRecord`, the sole encoder
of the canonical TO list and `{From:system, To:operator, Role:system}` envelope.

Targeted command:

```text
go test -count=1 ./internal/engine ./internal/obligation
go test -count=1 ./test/fixtures -run '^TestS(9OwnerDecisionBrief|10(Approval|Expiry|ProductionScheduler|ODB|Park)|11StaleChoice)' -v
test "$(rg -n 'SystemOperatorRecord\(SystemOperatorInput|obligation\.SystemOperatorRecord\(obligation\.SystemOperatorInput' internal/engine internal/obligation | wc -l | tr -d ' ')" = 5
```

Observed: GREEN; one address-list encoder and exactly five governed call sites.

Between-item battery: `go test -count=1 ./... && go vet ./...` GREEN.

## Item 5 — drop per-emit tables.Build

The scheduler constructor now requires a table-snapshot supplier; production
injects the already-published `liveTables.Snapshot`. `Emit` resolves content
hashes from that snapshot plus a scheduler-local post-submit cache, so crash
refire dedupe remains byte-identical without rebuilding tables from disk for
each emit. The two other `tables.Build` sites in `resummon.go` are unchanged:
`ArmParked` and the post-timer `due` recheck are the plan-classified
evidence-only sites.

Targeted command:

```text
go test -count=1 ./internal/engine -run '^TestS10(ResummonTimerCrashRefireDedupesBySeatDecisionAndCadenceSlot|ProductionSchedulerArmsParkedGateAndEmitsExactlyOneResummon)$' -v
test "$(rg -n 'tables\.Build' internal/engine/resummon.go | wc -l | tr -d ' ')" = 2
```

Observed: GREEN; both remaining build loci are `ArmParked`/`due`, and
`outcomeForContentHash` contains no build.

Between-item battery: `go test -count=1 ./... && go vet ./...` GREEN.

## Item 4 — prompter lookups from table snapshots

Approval and expiry prompters now build one startup `tables.Live` view and use
the `VerdictsByGate` / `ApprovalGates` indexes for replay and entry-scope
lookups. Live operator resolutions increment that view after the committed gate
is verified. The prior per-lookup `Store.Records()` disk scans are gone; gate
identity verification still uses the single immutable `Store.Read(gateID)`.

Targeted command:

```text
go test -count=1 ./internal/tables ./internal/engine
go test -count=1 ./test/fixtures -run '^TestS10(ApprovalPrompter|ExpiryPrompter)' -v
```

Observed: GREEN; the snapshot clone test also proves both new indexes survive
publish without mutating an older snapshot.

Between-item battery: `go test -count=1 ./... && go vet ./...` GREEN.

## Item 3 — one ODB builder

`obligation.RenderODB` is now the single record builder. The public
`engine.RenderODB` adapter preserves its existing input API, while
`completeODB` supplies its gate-derived phase/schema/choice-row values to the
same builder. The adapter flags preserve the pre-refactor byte distinctions:
the public render includes its dispatch header; derived gate ODBs retain
`PHASE:SITREP`, the gate schema version, and representational choice columns.

Targeted command:

```text
go test -count=1 ./internal/engine ./internal/obligation
go test -count=1 ./test/fixtures -run '^TestS(9OwnerDecisionBriefIsExactAndModelIsPayloadOnly|10ODB|10Park|11StaleChoice|11StructuralChoice)' -v
```

Observed: GREEN, including the T6 structural-column and crash-reissue ODB paths.

Between-item battery: `go test -count=1 ./... && go vet ./...` GREEN.
