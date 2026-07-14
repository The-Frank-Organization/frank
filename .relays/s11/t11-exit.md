# T11 s11 exit evidence

Implementation base: `main@d91fcfb340b029c39c8493084ce2f227409aa546`.
Implementation tip before this evidence-only exit commit:
`6e25f2059fdaa71e0ba5798f3c5334161508208f`.

## Acceptance readback

- T1–T4: B/C/D projections and the full seven-state FSM are live. Bucket B is
  non-interrupting and raise-only; C is operator-CC FYI without obligation; D
  is rejected acceptance-bounce author return; accepted egress block stays A.
- T6: frozen decision π, migrate-then-validate, rejected
  `stale_choice_set`, held/operator-visible `stale_schema`, no stale wake, and
  crash-replayed replacement under one new decision identity are live.
- T7: the 14-row terminal/edge matrix and the known-A B-pick negative are live.
- T8: eight of nine — item 2 rescoped by master
  (`s11-build-escalate-fence/RECONCILE-…-143010`), carried post-Step-2 to
  m-7+m-3. The eight retained items landed in card order, one commit and one
  full green battery each.
- T9: both G4 timer classes come from startup-pinned operator config. The
  schema exposes durations only, rejects `auto_approve`, and production emits
  resummon commands rather than resolutions.
- T5 remains report-and-hold with acceptance OPEN: g2 has an m-5 planner
  proposal but no m-5 implementer review and no completion to master.
- T10 remains report-and-hold with acceptance OPEN: no dc design-cell return
  exists. No stub, assumption, or silent scope reduction was introduced.

## Label to mechanism sweep

| Label | Mechanism | Executable proof |
|---|---|---|
| bucket B | `Store.ProjectBucketB` over pinned classification | live B/raise-only fixture + matrix |
| bucket C | `Store.ProjectBucketC` over canonical TO/CC | live CC-only/no-obligation fixture + matrix |
| bucket D | `Store.ProjectBucketD` over rejected acceptance edges | D/egress precedence fixture + matrix |
| `bounced_repair` | `GateState` + `acceptanceBounceEdge` | seven-state fixture |
| `egress_blocked` | accepted egress-block branch in `GateState` | local park/resummon fixture; away send explicitly unbuilt |
| `stale_choice_set` | `classifyVerdict` guard → rejected candidate + durable reissue intent | live stale/crash fixture |
| `stale_schema` | obligation recovery → held system/operator record | live held/outbox/crash fixture |
| G4 configured cadence | `EngineConfig.ResummonCadenceDelays` → production `ArmParked` | two production-binary timer cases |

Every seat-visible label above has a writer and a consuming projection/state
derivation. The held T5/T10 labels have no pretend mechanism.

## I-PH, terminal enum, R2, and scope

Commands:

```text
go test -count=1 ./test/invariants -run '^TestLaw(PathHygiene|TerminalEnumByteExact|R2NoModelPredicate)$' -v
go test -count=1 ./test/fixtures -run '^TestS11' -v
test -z "$(git diff --name-only d91fcfb...HEAD | rg '^(internal/observe/|internal/store/(store|lock|quarantine)\.go$|internal/fieldspec/registry\.json$)' || true)"
git diff --check d91fcfb...HEAD
```

Observed: GREEN. Canonical store/config/socket/outbox paths remain rejected at
the registered seat-visible sinks; terminal values remain byte-exact
`{accepted,rejected,held}`; model data remains payload/render-only and never a
predicate; `internal/observe/`, `registry.json`, and the store write path are
absent from the branch diff.

## Mechanical reconciliation and catch ledger

- Consumption→supply and diff→license rows reconcile to every branch path.
- Same-file task order is explicit for submit, resummon, ODB, FSM, and main.
- Catch #1 (pre-code): executor license contradiction, pair-found and
  master-ruled before implementation.
- Catch #2 (pre-code): item-5/item-6 locus table truth errors, pair-found and
  corrected before implementation.
- Catch #3 (in-build, pre-item-7 commit): incremental evidence insertion left
  T8 prose sections out of card order; the code commits were already ordered,
  and the ledger was reconciled to 1,3,4,5,6,7,8,9 before the next item.

The fresh full uncached suite, vet result, final head, and clean status are
reported in the file-first implementation relay after this evidence commit.
