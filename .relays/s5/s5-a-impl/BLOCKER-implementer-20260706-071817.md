## IMPL blocker — full-suite acceptance now requires out-of-scope legacy test/fixture updates

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s5-a-impl-full-suite-blocker
PARENT_DISPATCH_ID: s5-a-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - scope/acceptance ruling required before commit
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer, operator
IN_REPLY_TO: .relays/s5/s5-a-impl/PLAN-planner-20260706-071046.md
SUBJECT: rev3 resolves_gate amendment folded and focused checks green; full `go test ./...` fails because EVIDENCE_TARGET and record_kind scope changes require edits outside the locked file map

Delta review: I approve the rev3 `resolves_gate` amendment as bounded and faithful to the stated operator-seat-scoped Step-1 semantics. I folded it into the WIP: `resolves_gate` now has `visible_when: {"any_of":[{"seat_is":["operator"]},{"role_in":["operator"]}]}`, and the fixture splits 37 strict-dormant names from the operator-only live affordance leg.

Focused verification now green:
- `go test ./internal/fieldspec -run 'TestRegistryS5MemberContainsRegistryPassRows|TestRegistryV2MemberParsesAndExposesLockedEnums'` -> PASS.
- `go test ./test/fixtures -run S5Registry` -> PASS.
- Payload check -> `version s5-fieldspec-v3`, `field_count 83`, `named_enum_count 24`, `resolves_gate_visible_when {'any_of': [{'seat_is': ['operator']}, {'role_in': ['operator']}]}`.
- `go vet ./...` -> PASS.
- `git diff --check` -> PASS.
- Raw annotation grep for m-2/m-1/m-6/s5-b honesty strings -> hits present.

Blocking verification:
- Command: `go test ./...`
- Result: FAIL.

Representative failures:
- `cmd/frank-mcp`: `TestSubmitArgumentsRoundTripStructuredStringCarrier` expected accepted submit, got rejected.
- `internal/engine`: submit/config-change tests reject with `EVIDENCE_TARGET:required` instead of their previous target assertions.
- `internal/obligation`: owed-item tests no longer project/open because the candidate is rejected; `TestOwedItemAcceptsNonOperatorSeat` also conflicts with the planned operator-only owed scope.
- `test/fixtures`: S4/F11/shim/config-change fixture paths reject before their former assertion points, commonly due to `EVIDENCE_TARGET:required`; some crash/applicability rows no longer hit their expected mutation point because the submit rejects earlier.

Root cause:
- The registry pass intentionally adds `EVIDENCE_TARGET.required_when = {"not":{"phase_in":[]}}`, and existing submit-path tests/fixtures still construct many records without `EVIDENCE_TARGET`.
- The registry pass intentionally narrows `record_kind` so `"*"` has only `diagnostics`, while existing owed/shim fixtures still expect non-operator owed affordance/acceptance.
- These are expected downstream fixture consequences of the approved data changes, not a `resolves_gate` regression.

Scope conflict:
- The locked plan still requires `go test ./...` green.
- The locked write surface still allows only `internal/fieldspec/registry.json`, `internal/fieldspec/registry_test.go`, new `test/fixtures/s5_registry_dormancy_test.go`, and optionally `internal/fieldspec/render_test.go` / `internal/fieldspec/validate_test.go` for registry-content assertions.
- The failing files include `cmd/frank-mcp/mcp_test.go`, `internal/engine/*_test.go`, `internal/obligation/owed_test.go`, and existing `test/fixtures/*`, which the dispatch/plan explicitly leave OUT.

Decision needed:
1. Expand scope and reissue the scope diff so I can update the affected legacy tests/fixtures to the S5 registry contract.
2. Or change the acceptance gate to exclude those legacy submit-path suites from this branch and leave their updates to integration/s5-b.
3. Or revise the registry contract itself, which would contradict the current S5 fixture and design acceptance, so I do not recommend it without a fresh design ruling.

Current worktree state:
- Branch/worktree: `~/frank-s5-team/s5-a`, branch `s5-a-registry`.
- No commit made.
- WIP files: `internal/fieldspec/registry.json`, `internal/fieldspec/registry_test.go`, new `test/fixtures/s5_registry_dormancy_test.go`.

ACTIONS_GIT_REF: WIP only; no commit; focused registry/fixture gates green, full-suite gate blocked by out-of-scope legacy tests as above
FINAL_GIT_STATUS_SHORT:
 M internal/fieldspec/registry.json
 M internal/fieldspec/registry_test.go
?? test/fixtures/s5_registry_dormancy_test.go

Next requested action: planner/orchestrator ruling on scope expansion vs acceptance narrowing. I will not patch the failing out-of-scope tests under the current dispatch.
