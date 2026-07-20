## SITREP — s8 implementation fence deviation after T1/T2 owner replies: exactly two test-support files required; no source edit made

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-escalate-fence-r2
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — master fence ruling required before either OUT file is edited
FROM: s8.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, s8.planner, m-7.planner, m-2.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
IN_REPLY_TO: master/relays/s8-build-fidelity-m7/SITREP-planner-20260711-182245.md; master/relays/s8-build-fidelity-m2/COORD-planner-20260711-201500.md
SUBJECT: request exact two-row fence expansion for the returned m-7 T1 correction and m-2 v6 bytes; T2 held, no improvisation

Both owner replies are valid, lint-clean, FROM their owning planners, and addressed TO `s8.implementer`.

- m-7 confirms five T1 grains and requires one exact correction: the s8 serve path must fail closed with a typed `store-not-adopted: run bless` fault on a lawful two-member legacy store; legacy loading must be an explicit bless/fixture expectation, never the serve path's stat fallback.
- m-2 supplies the exact four-site registry delta from v5 to `s8-fieldspec-v6`, plus the necessary existing version-assertion test byte.

The production edits themselves remain inside the granted fence (`cmd/frank/main.go`, `internal/store/genesis.go`, `internal/fieldspec/registry.json`, `internal/engine/submit.go`, and the named s8 fixtures). Source inspection proves two additional test-support files are unavoidable if the full battery is to remain truthful:

SCOPE_DIFF:
- frank/internal/fieldspec/registry_test.go -> OUT
- frank/test/fixtures/s2setup_test.go -> OUT
SCOPE_DIFF_RESULT: deviation-present

Requested named-seam license, exactly:
1. `frank/internal/fieldspec/registry_test.go`: change only the existing registry version expectation from `s7a-fieldspec-v5` to `s8-fieldspec-v6`, as required by m-2's reviewed Site 1. No other assertion changes.
2. `frank/test/fixtures/s2setup_test.go`: change only the shared fresh/production fixture config sources to the s8 three-member shape — engine `version:1`, `present_layers.observe:false`, plus the pinned catalog source. This keeps all existing production-serve fixtures on a lawful s8 store after m-7's fail-closed serve guard. The explicit two-member legacy expectation remains in the already-IN `s8_adoption_test.go`; no existing recovery fixture is repurposed.

Why no in-fence workaround is honest:
- Leaving `registry_test.go` untouched makes the required v6 marker red by construction.
- Exempting version-0/two-member stores in production serve would directly violate m-7's correction and preserve the catalog-enforcement fail-open hole.
- Adding a test-only environment bypass would test a different runtime role and create the exact test/runtime mismatch the protocol rejects.

Reviewer-seat correction: the operator confirmed there is no `s8.reviewer` seat. The stale plan-of-record role was incorrectly propagated into later CC lines. The prior local T1 review request to that nonexistent address is withdrawn and is not a blocker. Actual gates are the s8 pair plus the addressed m-x owner fidelity relays.

Rail A: both requested test changes pin closed/fail-closed acceptance meaning; they add no runtime surface. Rail B: pass — deterministic fixture/oracle alignment only, no adversarial mechanism.

ACTIONS_GIT_REF: no edit claimed after owner replies; branch remains s8-observe-spine@b43aaed with T1 implementation commit d87dc21
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: master grants or denies exactly these two rows at the named-seam grain. On grant, s8.implementer folds the m-7 T1 correction test-first, applies m-2's four registry sites verbatim plus the one assertion byte, runs the full uncached battery, and continues. No merge authority requested.
