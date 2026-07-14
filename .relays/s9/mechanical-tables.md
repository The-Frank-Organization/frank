# s9 mechanical tables

## Consumption to supply

| Consumer | Supply | State |
|---|---|---|
| T1 read-file | governed `RegistryEnv.Lanes[lane_ref]`, timeout router | landed in T1 shared worker |
| T2 find-references | T1 descriptor-rooted worker | implemented, verification pending commit |
| T3 verdict binding | conductor `Selection`, registry `CheckEntry`, executor verdict | pending |
| T4 absence floor | T1 root health, T3 rows, m-7 `lane_vcs` owner bytes | held on m-7 return |
| T5 attestation negative | system-owned FieldSpec admission | verified covered; test-only fixture |
| T6 mixed rollup | T3 row output | pending verification |
| T7 bounce class | m-2 registry co-sign | closed |
| T8 envelope key hygiene | m-1 and m-2 co-signs | closed |
| T9 exit fixtures | T1 through T6 plus opened owner-gated tasks | pending |

## Diff to license

| Path | Task/seam | License |
|---|---|---|
| `internal/observe/fs_worker.go` | T1 shared worker | in |
| `internal/observe/fs_worker_test.go` | T1 RED/GREEN fixtures | in |
| `internal/observe/read_file_worker.go` | T1 extraction, read-file byte semantics preserved | in |
| `internal/observe/registry.go` | T1 shared worker state and hook only | in |
| `.relays/s9/` | evidence capture | in |
| `internal/observe/fs_worker.go` | T2 descriptor-rooted scanner operation | in |
| `internal/observe/registry.go` | T2 entry, params, dispatch and verdict mapping | in |
| `internal/observe/registry_test.go` | T2 RED/GREEN fixtures | in |
| `test/fixtures/s8_suppliability_guard_test.go` | T5 lane-forgery negative + conductor default fixture | in |
| `.relays/s9/t5-verification.md` | T5 verification evidence | in |

No `internal/executor/executor.go` diff. No T7/T8 or blocked-ledger code.
