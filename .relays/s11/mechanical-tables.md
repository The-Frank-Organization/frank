# s11 mechanical tables

## Consumption to supply

| Consumer | Supply | State |
|---|---|---|
| T1 bucket B saved query | pinned FieldSpec `gate_category_B` membership + immutable records | implemented; targeted fixture and full uncached battery green |

## Diff to license

| Path | Task/seam | License |
|---|---|---|
| `.relays/s11/mechanical-tables.md` | cross-task consumption and diff-license evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t1-red-green.md` | T1 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `internal/store/projections.go` | T1 live bucket-B saved query over pinned category tags | IN — token row `frank/internal/store/projections.go` |
| `test/fixtures/s11_buckets_test.go` | T1 bucket-B behavior fixture; later T2/T3/T7 bucket matrix extension | IN — token row `frank/test/fixtures/` |

The table is updated before each task commit and reconciled to the live diff.
