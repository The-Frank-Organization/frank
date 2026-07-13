# s10 T11 mechanical reconciliation

Base: `main@8941889`

Branch: `s10-comms-spine`

Plan lock: `s10-comms-spine-plan-r2-s10.2`

## Consumption to supply

| Task | Consumes | Supplies | Proving surface |
|---|---|---|---|
| T1 | v7 registry and m-2/m-7 owner bytes | governed v8 reader marker, `odb`, `resummon_command` | `TestS10V7ReaderRefusesV8MarkerBeforeContent` |
| T2 | T1 v8 bytes and existing init machinery | store born directly at v8, no pre-v8 record | `TestS10FreshGenesisIsBornAtV8WithoutPreV8Records` |
| T3 | T1 `odb` kind | minimal seven-field ODB and frozen choices | `TestS10ODBRejectsChoiceOutsideFrozenSet` |
| T4 | T3 ODB and accepted A classification | ODB + `parked_waiting_human`; non-A stays live | `TestS10OnlyAcceptedAGateEmitsODBAndParksWaitingForHuman` |
| T5 | T3 frozen choices and T4 park | operator-FROM, in-set, one-shot gate resolution | `TestS10OperatorReplyValidatesFrozenChoiceAndResolvesOnlyOnce` |
| T6 | T5 validated reply and m-3 observe hook | newest-authoritative local re-observe before wake | `TestS10WakeReobservesParkedGateBeforeResuming` |
| T7 | T4-T6 durable records | derived crash-safe `resumed`, exactly one wake | `TestS10CrashSafeParkAndWakeConvergeExactlyOnce` |
| T8 | T7 liveness and T1 `resummon_command` | timer reasons and A-2 `(seat, decision, cadence-slot)` dedupe | `TestS10ResummonTimerCrashRefireDedupesBySeatDecisionAndCadenceSlot` |
| T9 | T3-T5 prompt path and m-7 expiry amendment | E1/E2 soft-expiry ODB `{kill, extend}` plus absolute hard block ceiling | `TestS10LongRunningCheckParksAlertsAndAppliesOperatorExtend` |
| T10 | T3-T5 live verdict path and m-3 named entry | typed default deny and live approval gate-lift, zero spawn | `TestS10LiveOperatorVerdictLiftsApprovalGateThroughODBWithoutExecution` |
| T11 | T1-T10 | fresh-v8 EXIT LEG 3 and whole-slice evidence | `TestS10ExitLeg3FreshV8GateWakesExactlyOnceAfterLocalReobserve` |

All produced interfaces have a named consumer or are terminal slice outputs. PARK-ACROSS-V8 is avoided by T1/T2 ordering; the s11 freeze/re-issue branch is not implemented.

## Diff to license

| Task / commit | Changed production seam | License reconciliation |
|---|---|---|
| T1 `94f8ab1` | `fieldspec/registry.json`; `config.go` capability/successor sites; `registry_test.go`; v8 version carrier in `store/genesis.go` | T1 governed transition plus the exact master-granted config/registry seams; fixture pin updates are semantic-invariance companions |
| T2 `b250bd1` | `store/genesis.go` fresh-init derivation | T2 sole production seam; fixture companions under `test/fixtures/` |
| T3 `bb9a96e` | new `engine/odb.go` | T3 new-unit license under `engine/` |
| T4 `1d73839` | `obligation/obligation.go` `completePark` path | T4 named obligation seam |
| T5 `daab2c1` | `engine/submit.go` gate-resolution validation; ODB/obligation choice support | T5 gate-resolution acceptance arm and its frozen-choice support |
| T6 `0ae553e` | `engine/submit.go` wake re-observe arm | T6 named wake-delivery arm; no registry entry added |
| T7 `f6fbccd` | new `engine/fsm.go` | T7 derived recovery/wake unit under `engine/` |
| T8 `82cdd19` | new `engine/resummon.go`; `fsm.go` resummon state | T8 scheduler/A-2 units under `engine/`; production handler is composed at the licensed main block in T9 |
| T9 `5b28d10` | executor expiry disposition; E1 read disposition; loop nested verdict arm; obligation custom choices; licensed main composition blocks | T9 named seams. `Spawn` admission/class/timeout checks are byte-untouched; absolute ceiling only kills/blocks |
| T10 `e7a97a0` | additive `observe/registry.go` entry + pre-spawn gate; loop approval arm; licensed main composition blocks | T10 gate-lift license. `internal/executor/executor.go` has no T10 diff and nothing spawns |
| T11 | `test/fixtures/s10_exit_test.go` and `.relays/s10/` evidence only | T11 fixture/evidence license; no production source seam |

Every changed path is inside the expanded mechanical block. No config/schema change exists after T1. No merge action is taken.

## Label to mechanism and summary-line check

| Label / claim | Runtime mechanism | Executable check |
|---|---|---|
| `s10-fieldspec-v8` | marker-first exact capability and forward successor | old reader refuses before content; fresh reader loads |
| `parked_waiting_human` | derived ODB then park records from accepted A gate | non-A negative plus A positive |
| `resumed` / exactly-once wake | `GateState` derives accepted resolution dominance | crash park/wake legs and EXIT LEG 3 count one |
| `resummon_due` | internal serialized `resummon_command` with A-2 content hash | crash-refire produces one command |
| “silent auto-kill is gone” | descriptor/read timeout is a soft prompt; absolute hard ceiling is kill-only | park+alert+extend, operator kill, and no-response ceiling legs |
| “static-only gate is gone” | `run-suite-unbounded` default-deny consults a live operator verdict | unapproved typed refusal; live gate-lift; zero executor calls |

Summary line verified: the slice implements one fresh-v8 A-gate path from accepted send through ODB/park, validated reply, local re-observe, and exactly-one wake; both interim approval-channel defaults are sunset without widening the executor boundary.
