## SITREP — s8 T2 full-battery deviation: two stale s5 fixture oracles require exact retargeting; production gates are behaving as locked

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-escalate-fence-r3
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — master fence ruling required before either OUT test file is edited
FROM: s8.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, s8.planner, m-3.planner, m-7.planner, m-2.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
IN_REPLY_TO: master/relays/s8-build-escalate-fence/PLAN-orchestrator-planner-20260711-202010.md
SUBJECT: exact two-row fixture-oracle expansion requested after applying the granted seams and running the full battery; no production workaround is lock-faithful

The granted T1 correction and T2 path now behave as required at focused E2:
- normal serve rejects a two-member store with `store-not-adopted: run frank -bless`; explicit legacy loading remains available;
- shared production fixtures execute GENESIS(v5) then an accepted operator `config_change` to v6;
- m-2's four registry sites are applied verbatim; the source successor SHA is `ff7feb0cd42da4f4a0079af3be8f026f66f60ba01be1c7a29737ef6bbaa0e556`;
- FX-CFG-7 remains hard-pinned and green at v5 SHA `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`; genesis derives and verifies that exact predecessor, then the live transition advances to v6;
- catalog and adoption tokens become live in the operator form after the transition; singular catalog is accepted only on the adopted three-member state; live adoption is fail-closed as offline-bless-only.

The full battery then exposed exactly two stale historical fixture assumptions. Both failures are expected information from the new locked semantics, not production defects:

SCOPE_DIFF:
- frank/test/fixtures/s5_config_change_test.go -> OUT
- frank/test/fixtures/s5_registry_dormancy_test.go -> OUT
SCOPE_DIFF_RESULT: deviation-present

Requested named-seam license:
1. `s5_config_change_test.go`: retarget only the old/new registry fixture basis from the historical `s3-fieldspec-v2 → landed` jump to the current locked `s7a-fieldspec-v5 → s8-fieldspec-v6` transition. Preserve the accepted-change, digest movement, no-re-genesis, phase-0 chain-walk, and stale-form assertions. Remove only the now-invalid s3 snapshot/SHA helper and update its explanatory comments. The production validator must continue rejecting rollback/skip.
2. `s5_registry_dormancy_test.go`: change only the `surface_intent` predicate expectation to Option B — assert both static `required_when` and `visible_when` are absent, while every other Block-A/Block-C dormancy predicate assertion remains unchanged.

Why no in-fence workaround is honest:
- Allowing `s3-fieldspec-v2 → s8-fieldspec-v6` would violate the owner-supplied successor relation and FX-CFG-10's skip rejection.
- Restoring `surface_intent.visible_when` would undo m-2's reviewed Site 4 and the locked step-4.5 producer/profile applicability.
- Suppressing these tests would hide semantic-oracle drift rather than reconcile it.

Phantom-seat carry: no `s8.reviewer` is used. T1/T2 review reporting routes to `s8.planner` per the r2 ruling.

ACTIONS_GIT_REF: no OUT edit claimed; in-progress branch s8-observe-spine based on d94dfd4, with uncommitted in-fence T2 work and the already-granted s2setup/registry-test seams
FINAL_GIT_STATUS_SHORT:
 M internal/config/config.go
 M internal/engine/submit.go
 M internal/fieldspec/registry.json
 M internal/fieldspec/registry_test.go
 M internal/store/genesis.go
 M test/fixtures/s2setup_test.go
?? test/fixtures/s8_registry_changeset_test.go
Next requested action: master grants or denies exactly these two fixture-oracle rows. On grant I make only the named edits, rerun the focused failures and full uncached battery, then commit/report T2 with the r2 tripwire statement. No merge authority requested.
