# s9 mechanical tables

## Consumption to supply

| Consumer | Supply | State |
|---|---|---|
| T1 read-file | governed `RegistryEnv.Lanes[lane_ref]`, timeout router | landed in T1 shared worker |
| T2 find-references | T1 descriptor-rooted worker | landed at `5f6a7ec` |
| T3 verdict binding | conductor `Selection`, registry `CheckEntry`, executor verdict | implemented; full commit-point battery green |
| T4 absence floor | T1 root health, T3 rows, m-7 `lane_vcs` owner bytes | implemented; m-7 bytes at `db9a166`; full battery green |
| T5 attestation negative | system-owned FieldSpec admission | verified covered; test-only fixture |
| T6 mixed rollup | T3 row output | verified: `authority_mixed` reaches decision-② `held` |
| T7 bounce class | m-2 registry co-sign | closed |
| T8 envelope key hygiene | m-1 and m-2 co-signs | closed |
| T9 exit fixtures | T1 through T6 plus opened owner-gated tasks | implemented; consolidated runnable set + decision-⑤ pair green |

## Diff to license

| Path | Task/seam | License |
|---|---|---|
| `.relays/s9/mechanical-tables.md` | cross-task consumption and diff-license evidence, including MF-1 fold | IN — token row `frank/.relays/s9/` |
| `.relays/s9/t1-red-green.md` | T1 RED/GREEN evidence | IN — token row `frank/.relays/s9/` |
| `.relays/s9/t2-red-green.md` | T2 RED/GREEN evidence | IN — token row `frank/.relays/s9/` |
| `.relays/s9/t3-red-green.md` | T3 RED/GREEN and rev13 guide-byte evidence | IN — token row `frank/.relays/s9/` |
| `.relays/s9/t4-red-green.md` | T4 owner-byte and RED/GREEN evidence | IN — token row `frank/.relays/s9/` |
| `.relays/s9/t5-verification.md` | T5 verification-first evidence | IN — token row `frank/.relays/s9/` |
| `.relays/s9/t9-red-green.md` | T9 consolidated runnable-set and decision-⑤ evidence | IN — token row `frank/.relays/s9/` |
| `cmd/frank/main.go` | T4 m-7 E9 runtime handoff carried verbatim in `db9a166` | IN — named `main.go` seam |
| `internal/config/config.go` | T4 m-7 E1-E8 `lane_vcs` owner bytes carried verbatim in `db9a166` | IN — named `config.go` seam |
| `internal/config/lane_vcs_test.go` | T4 FX-VCS load/transition/ceiling fixture matrix in `1b87261` | **OUT — escalated to master under `s9-build-escalate-fence`** |
| `internal/observe/checks_base.go` | T4 total input table, governed-none branch, worker root-health consumption | IN — token row `frank/internal/observe/` |
| `internal/observe/checks_base_test.go` | T4 opaque/nil/git/config/serialization fixtures | IN — token row `frank/internal/observe/` |
| `internal/observe/fs_worker.go` | T1 shared detachable worker plus T2 descriptor-rooted scan operation | IN — token row `frank/internal/observe/` |
| `internal/observe/fs_worker_test.go` | T1 detach/breaker/component-swap fixtures | IN — token row `frank/internal/observe/` |
| `internal/observe/gate.go` | T3 thickened system-owned claim rows | IN — token row `frank/internal/observe/` |
| `internal/observe/read_file_worker.go` | T1 extraction with read-file semantics preserved | IN — token row `frank/internal/observe/` |
| `internal/observe/registry.go` | T1 worker hook, T2 registry/dispatch, T3 binding pass, T4 cloned `LaneVCS` consumer | IN — token row `frank/internal/observe/` |
| `internal/observe/registry_test.go` | T2 complete and fail-closed scan matrix | IN — token row `frank/internal/observe/` |
| `internal/observe/verdict_binding.go` | T3 conductor-origin envelope, total tuple validator, redaction, signal derivation | IN — token row `frank/internal/observe/` |
| `internal/observe/verdict_binding_test.go` | T3 adversarial matrix and rev13 base-refusal boundary | IN — token row `frank/internal/observe/` |
| `test/fixtures/s2setup_test.go` | T4 m-7 E10 v3 dogfood descriptor plus derived v2 residency helper | IN — token row `frank/test/fixtures/` |
| `test/fixtures/s8_adoption_test.go` | T4 dedicated v2 residency/adoption path | IN — token row `frank/test/fixtures/` |
| `test/fixtures/s8_adversarial_test.go` | T3 six-column row and truncated-pass regression | IN — token row `frank/test/fixtures/` |
| `test/fixtures/s8_config_activation_test.go` | T4 v3 dogfood genesis assertion | IN — token row `frank/test/fixtures/` |
| `test/fixtures/s8_exit_gate_test.go` | T3 valid-origin aggregation and dogfood row regression | IN — token row `frank/test/fixtures/` |
| `test/fixtures/s8_suppliability_guard_test.go` | T5 lane-forgery negative and conductor default | IN — token row `frank/test/fixtures/` |
| `test/fixtures/s8_supply_test.go` | T4 v3 duplicate-marker fixture | IN — token row `frank/test/fixtures/` |
| `test/fixtures/s9_exit_test.go` | T9 fixture-scoped decision-⑤ egress pair | IN — token row `frank/test/fixtures/` |
| `test/invariants/store_recovery_test.go` | T4 m-7 E10 invariant-test descriptor lines carried in `db9a166` | **OUT — escalated to master under `s9-build-escalate-fence`** |

The table above is path-total for `git diff --name-only 39474d0..76179ec`.
The two OUT rows await master's named-row ruling; their substantive owner-byte
fidelity is separately confirmed and does not convert them to IN. No
`internal/executor/executor.go` diff. No T7/T8 or blocked-ledger code.
