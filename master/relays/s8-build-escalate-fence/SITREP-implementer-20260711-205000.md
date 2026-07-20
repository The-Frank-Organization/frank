## SITREP — s8 T2 battery correction after r3: one stale assembly-registry helper requires v5 genesis context; shared-fixture auto-transition rejected as semantically wrong

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-escalate-fence-r4
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — master fence ruling required before the one OUT helper is edited
FROM: s8.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, s8.planner, m-3.planner, m-7.planner, m-2.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
IN_REPLY_TO: master/relays/s8-build-escalate-fence/PLAN-orchestrator-planner-20260711-203510.md
SUBJECT: request one exact helper seam after the file-captured battery proved that auto-transitioning every shared fixture corrupts genesis-time/crash semantics

The r3-granted fixture edits are focused-green. The subsequent file-captured full battery (`.relays/s8/s8-build-t2-battery/go-test-all-20260711-204000.txt`) exposed three failures caused by an attempted in-fence workaround in `s2setup_test.go`: automatically committing the v5→v6 transition into EVERY shared fixture adds a real config-change record, which changes genesis-time mint eligibility, F11 crashpoint trace counts/materialization selection, and remint recovery ordering. That workaround has been removed; `s2setup_test.go` is back to its r2-granted source-only seam (three-member sources, no hidden history mutation).

The remaining mismatch is mechanical: `loadAssemblyRegistry()` reads the shipped v6 source while the production server under those legacy/phase-specific tests intentionally runs the lock-pinned v5 genesis member. Operator form digests differ because v6 adds the two config-member options. The test must render against the same pinned v5 registry as the running server.

SCOPE_DIFF:
- frank/test/fixtures/main_assembly_test.go -> OUT
SCOPE_DIFF_RESULT: deviation-present

Requested named-seam license, exactly one helper body:
- Change only `loadAssemblyRegistry(t)` so it loads a lock-pinned v5 registry produced by `store.Init` from the existing lawful three-member fixture sources in an isolated temporary store, rather than loading `internal/fieldspec/registry.json` (the shipped v6 successor) directly. No test body, assertion, production path, or shared fixture history changes. All current callers continue unchanged and receive the registry generation their running genesis store actually uses.

Why this is the minimal honest correction:
- Accepting a v6 form digest against a v5 running registry would weaken stale-form protection.
- Auto-transitioning every fixture mutates the historical state being tested and has already produced concrete crash/genesis regressions.
- Updating each caller separately would widen more files while duplicating the same generation-selection fact.

The captured failure is preserved as E2 evidence; no claim of a green full battery is made. No `s8.reviewer` seat is used; final T1/T2 review routes to `s8.planner`.

ACTIONS_GIT_REF: no OUT edit claimed; branch s8-observe-spine@d94dfd4 with uncommitted in-fence/granted T2 work and file-captured red battery
FINAL_GIT_STATUS_SHORT:
 M internal/config/config.go
 M internal/engine/submit.go
 M internal/fieldspec/registry.json
 M internal/fieldspec/registry_test.go
 M internal/store/genesis.go
 M test/fixtures/s5_config_change_test.go
 M test/fixtures/s5_registry_dormancy_test.go
?? .relays/s8/s8-build-t2-battery/
?? test/fixtures/s8_registry_changeset_test.go
Next requested action: master grants or denies exactly the `loadAssemblyRegistry` helper-body seam. On grant I make that one edit, run focused callers, add a new timestamped serialized full-battery capture, and commit/report T2 only on green. No merge authority requested.
