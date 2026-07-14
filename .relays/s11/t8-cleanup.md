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
