# s11 mechanical tables

## Consumption to supply

| Consumer | Supply | State |
|---|---|---|
| T1 bucket B saved query | pinned FieldSpec `gate_category_B` membership + immutable records | implemented; targeted fixture and full uncached battery green |
| T2 bucket C saved query | canonical TO/CC address lists + immutable records | implemented; targeted fixture and full uncached battery green |
| T3 bucket D saved query | terminal `rejected` + system-authored acceptance-stage `failing_edge` + author identity | implemented; targeted regressions and full uncached battery green |
| T3 egress precedence | accepted A gate + `failing_edge=egress` + local ODB/park | implemented; targeted fixture and full uncached battery green |
| T4 `bounced_repair` | T3 rejected terminal + acceptance-stage `failing_edge` | implemented; targeted regressions and full uncached battery green |
| T4 `egress_blocked` | accepted A gate + blocked egress tag + existing park/resummon mechanisms | implemented; targeted regressions and full uncached battery green; away send remains unbuilt |
| T6 frozen decision guard | immutable source ODB π snapshot + deep-cloned migrated view + live `classifyVerdict` | implemented; alias, breaking, and structural fixtures green |
| T6 stale disposition | rejected `stale_choice_set` candidate + durable reissue intent + held `stale_schema` + new decision identity | implemented; both byte tokens, D query, no-wake, and operator outbox proven |
| T6 crash replay | deterministic intent on stale candidate + recovery replay after `stale_reissue_after_held` | implemented; real-process crash replays same identity exactly once |
| T6 cadence restart | old decision suppressed by held fault + new decision ID in resummon content key | implemented; same-seat/new-ID key separation fixture green |
| T7 terminal/edge matrix | T1–T4/T6 bucket queries + obligation-derived A artifacts | registered; 14 positive/negative boundary rows green |
| T7 ③ known-A NF | live submit validator + known-A floor + B saved query + A ODB | registered; rewrite/provenance/no-absorb behavior green |
| T8 item 1 generic prompter | approval/expiry adapters + shared owner/wait/emit/resolve lifecycle | implemented; fail-safe decisions and duplicate-waiter fixtures green |
| T8 item 3 one ODB builder | public engine adapter + obligation gate adapter -> `obligation.RenderODB` | implemented; public, derived, structural, and reissue ODB fixtures green |
| T8 item 4 snapshot lookups | one startup table view + gate/entry indexes + incremental resolution publish | implemented; no prompter `Store.Records()` scans; replay/live fixtures green |
| T8 item 5 resummon snapshot supply | production live snapshot + local emitted-hash cache | implemented; per-emit build removed, dedupe output unchanged |
| T8 item 6 system-to-operator builder | five post-item-3 operator-addressed sites -> one canonical envelope/address encoder | implemented; exact five-site census and surface fixtures green |
| T8 item 7 hash/ID decoupling | shared raw prompt digest -> independently rendered hash and gate-ID prefixes | implemented; output-byte fixtures green, prefix slicing absent |
| T8 item 8 executor cleanup owner | all run-completion branches -> `finalizeRun` cleanup decision | implemented; survivor preservation and normal/timeout cleanup characterization green |

## Diff to license

| Path | Task/seam | License |
|---|---|---|
| `.relays/s11/mechanical-tables.md` | cross-task consumption and diff-license evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t1-red-green.md` | T1 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t2-red-green.md` | T2 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t3-red-green.md` | T3 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t4-red-green.md` | T4 sequence-honest RED/GREEN evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t6-red-green.md` | T6 sequence-honest RED/GREEN and crash-replay evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t7-matrix.md` | T7 prerequisite RED lineage and consolidated matrix evidence | IN — token row `frank/.relays/s11/` |
| `.relays/s11/t8-cleanup.md` | T8 per-item order, characterization, and battery evidence | IN — token row `frank/.relays/s11/` |
| `internal/engine/prompter.go` | T8 item 1 generic lifecycle; T8 item 7 shared raw digest | IN — token row `frank/internal/engine/`; item order 1 then 7 |
| `internal/tables/tables.go` | T8 item 4 gate-resolution and approval-entry snapshot indexes | IN — token row `frank/internal/tables/`, item 4 only |
| `internal/tables/tables_test.go` | T8 item 4 index clone/publish characterization | IN — token row `frank/internal/tables/`, item 4 only |
| `internal/engine/approval.go` | T8 items 1/4/6/7 prompt lifecycle, snapshot, operator envelope, hash-ID split | IN — token row `frank/internal/engine/`; item order 1 then 4 then 6 then 7 |
| `internal/engine/expiry.go` | T8 items 1/4/6/7 prompt lifecycle, snapshot, operator envelope, hash-ID split | IN — token row `frank/internal/engine/`; item order 1 then 4 then 6 then 7 |
| `internal/executor/executor.go` | T8 item 8 `finalizeRun` + preserve/cleanup call sites only | IN — explicit granted executor seam; `Spawn`, refusals, tokens, verdict logic, rung/timing untouched |
| `internal/migrate/migrate_test.go` | T6 alias-safety RED/GREEN | IN — token row `frank/internal/migrate/`, T6 only |
| `test/fixtures/s11_8a_test.go` | T6 live-path, byte-token, no-wake, structural, and crash-replay fixtures | IN — token row `frank/test/fixtures/` |
| `test/fixtures/f11_test.go` | T6 crashpoint registry live-site census adjacency | IN — token row `frank/test/fixtures/` |
| `internal/migrate/migrate.go` | T6 deep-clone before every migration step | IN — token row `frank/internal/migrate/`, T6 only |
| `internal/engine/odb.go` | T6 frozen π guard; T8 item 3 adapter to the one ODB builder | IN — token row `frank/internal/engine/`, named ODB seam; item order T6 then T8.3 |
| `internal/engine/submit.go` | T6 live verdict guard, typed stale candidate, deterministic durable reissue intent | IN — token row `frank/internal/engine/`, named verdict seam |
| `internal/engine/resummon.go` | T6 suppression; T8 item 5 snapshot supply; T8 item 6 operator-record adapter | IN — token row `frank/internal/engine/`; item order T6 then T8.5 then T8.6 |
| `internal/engine/resummon_test.go` | T6 cadence restart; T8 item 5 dedupe/scheduler snapshot characterization | IN — token row `frank/internal/engine/`, scheduler seam |
| `cmd/frank/main.go` | T8 item 5 production live-table snapshot injection | IN — token row `frank/cmd/frank/main.go`, composition-root-only item-5 seam |
| `internal/obligation/obligation.go` | T6 stale consumer; T8 item 3 ODB builder; T8 item 6 canonical system-to-operator builder | IN — token row `frank/internal/obligation/`; item order T6 then T8.3 then T8.6 |
| `internal/crashpoint/crashpoint.go` | T6 stale reissue durability boundary | IN — token row `frank/internal/crashpoint/`, T6 only |
| `internal/crashpoint/crashpoint_test.go` | T6 exact registered-crashpoint assertion adjacency | IN — token row `frank/internal/crashpoint/`, T6 only |
| `test/fixtures/s11_fsm_test.go` | T4 seven-state branch fixture and fixture-scoped egress negative | IN — token row `frank/test/fixtures/` |
| `internal/engine/submit.go` | T3 system-authored acceptance-stage `failing_edge` writer | IN — token row `frank/internal/engine/` and named submit-side verdict classification seam |
| `internal/engine/fsm.go` | T4 two missing labels and durable-state derivation | IN — token row `frank/internal/engine/` and named FSM seam |
| `internal/store/projections.go` | T1 bucket-B, T2 bucket-C, and T3 author-facing bucket-D saved queries, including delayed T3 acceptance-edge token closure for `stale_choice_set` | IN — token row `frank/internal/store/projections.go`; read-side T1–T3 only |
| `test/fixtures/s11_buckets_test.go` | T1–T3 bucket fixtures + T7 14-row matrix and live ③ known-A NF | IN — token row `frank/test/fixtures/` |

The table is updated before each task commit and reconciled to the live diff.
