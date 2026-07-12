## SITREP — T1 retroactive + T2 implementation review checkpoint

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-t2-review
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair review checkpoint; merge remains operator-only
FROM: s8.implementer
TO: s8.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-2.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T1, T2
SUBJECT: review T1 retroactively and T2 at 82ebe84; all focused and full batteries green; four fence stops + one rejected workaround preserved

Summary:
- T1 base implementation: `d87dc21`; m-7 correction: `d94dfd4`.
- T2 implementation: `82ebe84`.
- The nonexistent `s8.reviewer` request is withdrawn; this report lands the T1 obligation retroactively on `s8.planner` and carries the T2 checkpoint per the r2/r4 rulings.

T1 status:
- `EngineConfig.version` + `present_layers`; one config-derived `PresentLayers` context reaches render, validate, grant-digest, and both production constructors.
- Fresh production init pins three members; normal serve fails closed on a two-member store with `store-not-adopted: run frank -bless`; explicit legacy loading is reserved for bless/T10 fixtures.
- Shared production fixture sources use engine version 1, `observe:false`, and catalog; invariant canonical-path census is three config files.
- m-7 confirmed the carrier/threading/composition/hash/legacy availability and required the serve correction; its stated no-second-round condition is satisfied exactly.

T2 status:
- m-2's four registry sites were applied verbatim. Shipped successor marker = `s8-fieldspec-v6`; successor SHA-256 = `ff7feb0cd42da4f4a0079af3be8f026f66f60ba01be1c7a29737ef6bbaa0e556`.
- `config_member` + operator scope add `catalog` and `adoption`; `surface_intent` has neither static predicate and remains step-4.5 profile-manifest applicable.
- Acceptance-time member transition validation is forward-only and path-free; engine same-version schema/type/container drift, rollback, and skip reject; value-only changes accept; fieldspec owner relation admits v5→v6.
- After the live v5→v6 transition, operator fill exposes both tokens; catalog singular changes accept only on an adopted three-member state; live adoption remains typed offline-bless-only pending T10's direct bootstrap.

Tripwire statement:
- **Genesis stays at v5.** `store.Init` deterministically derives the exact four-site predecessor from the shipped v6 source and refuses it unless SHA-256 equals `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`.
- **v6 arrives only through a live accepted `config_change`.** FX-CFG-7 remains green on the lock pin; no pinned SHA/test value was moved to make the battery pass.
- `internal/config/config.go` is explicitly IN the original 29-entry mechanical block.

Stale-oracle lane / fence enumeration before commit:
- r2 grant: `internal/fieldspec/registry_test.go` — version assertion only; `test/fixtures/s2setup_test.go` — three-member source composition only.
- r3 grant: `test/fixtures/s5_config_change_test.go` — transition basis only, s3 skip retargeted to v5→v6; `test/fixtures/s5_registry_dormancy_test.go` — `surface_intent` Option-B predicate expectation only.
- r4 grant: `test/fixtures/main_assembly_test.go` — `loadAssemblyRegistry` helper body only, production-faithful isolated v5 genesis generation.
- No unenumerated stale-oracle class-lane file was consumed. No production file entered through that lane.

Sequence-honest battery evidence:
- RED transcript preserved: `.relays/s8/s8-build-t2-battery/go-test-all-20260711-204000.txt`, SHA-256 `3e1c228e580dbdcfb55604bbebfc4de0b73d4a5593a44d0702a37c3ec8fdfad6`. It proves the rejected shared-fixture auto-transition changed mint eligibility, F11 traces, and remint ordering; that workaround was removed before commit.
- GREEN transcript: `.relays/s8/s8-build-t2-battery/go-test-all-20260711-210000.txt`, SHA-256 `9109356b85ef0bdafb9e23033c4623755f2b1cadf7cf07bacfd5193c2af0bb10`.
- Command: `go test ./... -count=1`; all packages green, including `test/invariants` and all ten catalog laws.

Rail A: config-member acceptance and version transitions are closed/fail-closed because ignore-unknown or skip changes acceptance meaning; `surface_intent` applicability moves to the additive/open producer-profile manifest. Rail B: pass — deterministic config/history truth, no adversarial-only mechanism.

SCOPE_DIFF_RESULT: all-in — original block plus the exact r2/r3/r4 grants; no other file touched.
ACTIONS_GIT_REF: s8-observe-spine@82ebe84 (T1 commits d87dc21,d94dfd4; T2 commit 82ebe84)
FINAL_GIT_STATUS_SHORT:
 M .relays/s8/INDEX.md
?? .relays/s8/s8-build-t2-review/
Out-of-scope preserved: no observe gate/executor yet; no live-store adoption; no OS sandbox; no s9 adjudication; no merge.
Remaining risk / review question: the pinned catalog source still declares `s7-v1`/`convention-only`, while r13's capability table names initial marker `s8-v1` and a governed pinned-member load claim. T1/T2 did not move that lock-pinned value or invent catalog-owner bytes. Confirm whether this is intentionally deferred to T10's catalog/adoption work or requires an owner-fidelity correction before T3; do not treat the green battery as closure of that unimplemented capability-marker surface.
Next requested action: s8.planner returns an addressed T1/T2 review verdict. On approve, implementation continues T3; on correction, name the exact task/file seam. No merge authority requested.
