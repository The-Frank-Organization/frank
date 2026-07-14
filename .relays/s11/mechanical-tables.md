# s11 mechanical tables

## Consumption to supply

| Consumer | Supply | State |
|---|---|---|
| T1 bucket B saved query | pinned FieldSpec `gate_category_B` membership + immutable records | implemented; targeted fixture and full uncached battery green |
| T2 bucket C saved query | canonical TO/CC address lists + immutable records | implemented; targeted fixture and full uncached battery green |
| T3 bucket D saved query | terminal `rejected` + system-authored acceptance-stage `failing_edge` + author identity | implemented; targeted regressions and full uncached battery green |
| T3 egress precedence | accepted A gate + `failing_edge=egress` + local ODB/park | implemented; targeted fixture and full uncached battery green |

## Diff to license

| Path | Task/seam | License |
|---|---|---|
| `.relays/s11/mechanical-tables.md` | cross-task consumption and diff-license evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t1-red-green.md` | T1 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t2-red-green.md` | T2 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t3-red-green.md` | T3 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `internal/engine/submit.go` | T3 system-authored acceptance-stage `failing_edge` writer | IN — token row `frank/internal/engine/` and named submit-side verdict classification seam |
| `internal/store/projections.go` | T1 bucket-B, T2 bucket-C, and T3 author-facing bucket-D saved queries | IN — token row `frank/internal/store/projections.go` |
| `test/fixtures/s11_buckets_test.go` | T1 bucket-B behavior fixture; later T2/T3/T7 bucket matrix extension | IN — token row `frank/test/fixtures/` |

The table is updated before each task commit and reconciled to the live diff.
