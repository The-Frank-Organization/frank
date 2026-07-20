## IMPL - s7a F-S7-R2-COLGRAIN implemented red-first at d76c3ad; five-file fence and atomic singleton guard verified

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s7a-plan-m2-impl
PARENT_DISPATCH_ID: s7a-plan-m2-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - implementation report only; pair review, cross-domain fidelity, VP integration, and operator merge remain downstream
IN_REPLY_TO: master/relays/s7a-dispatch/DISPATCH-planner-20260710-140600.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-4.planner, m-4.implementer, m-7.planner, m-7.implementer
SUBJECT: F-S7-R2-COLGRAIN implemented on s7a-colgrain at d76c3ad; requesting pair review before master and m-4/m-7 fidelity

Summary: Implemented default-deny column-grain validation for `any_row` predicates. A row-array remains eligible only when its top-level FieldSpec is gate-referenceable, and the referenced row column must now also occur in that array's explicit `GateReferenceableColumns` allowlist.

PR: none - branch intentionally preserved for the relay-gated pair-review and integration path; no push, merge, or PR authority exercised.

Plan lock: `master/relays/s7a-dispatch/PLAN-planner-20260710-140000.md`, authorized by `s7a-plan-m2-dispatch`; no deviation.

Files changed:
- `internal/fieldspec/predicate.go`
- `internal/fieldspec/registry.go`
- `internal/fieldspec/registry.json`
- `internal/fieldspec/predicate_test.go`
- `internal/fieldspec/registry_test.go`

### Acceptance criteria status

- **AC1 PASS (E2):** two synthetic `routing_assignments.chosen_model` predicates, one `required_when` and one `visible_when`, reject during registry load. The error assertions pin the owner, `non gate-referenceable row field`, and `routing_assignments.chosen_model`; no path text is produced.
- **AC2 PASS (E2):** the shipped `routing_assignments.declared_deviated == "yes"` atoms load and compile, and the generic positive evaluation fixture now exercises that legal column.
- **AC3 PASS (E2):** a non-model `routing_assignments.seat` predicate is rejected by the same default-deny allowlist, proving this is not a `chosen_model` blocklist.
- **AC4 PASS (E2):** fresh uncached full suite, focused acceptance set, and `go vet ./...` all exit 0 at the committed tip.
- **AC5 PASS (E1):** `git diff --name-only 1d3e92c..d76c3ad` returns exactly the five authorized `internal/fieldspec/` files.
- **AC6 PASS (E1):** the shipped registry row adds exactly `routing_assignments.gate_referenceable_columns = ["declared_deviated"]`; the required additive-MINOR version marker advances from `s6-fieldspec-v4` to `s7a-fieldspec-v5`.
- **AC7 PASS (E2):** registry decoding asserts the singleton allowlist byte-for-byte as `[]string{"declared_deviated"}`.

### Red/green and atomicity proof

- Red commit `10ee3a2` (`test: expose any-row column gate gap`) adds the two `chosen_model` negatives plus the `seat` allowlist negative. Before production edits, `go test -count=1 ./internal/fieldspec -run '^TestRegistryLoadRejectsNonAllowlistedRowColumns$' -v` failed all three cases because registry load succeeded.
- Green commit `d76c3ad` (`fix: gate any-row predicates by column`) contains the production guard, `GateReferenceableColumns` member, singleton registry metadata, version treatment, positive retarget, and decoded singleton assertion together.
- Therefore no branch commit activates default-deny without the singleton present; the red fixture commit precedes both.

Boundary contract proof: byte-exact delivery states and the `submit`/`project`/`read` surface are untouched. No nested row-column schema, model-identity blocklist, routing column-set change, engine/store/seat-surface edit, or s7 branch edit was made.

Tests/verification:
- `go test -count=1 ./internal/fieldspec -run 'TestRegistryLoadRejectsNonAllowlistedRowColumns|TestPredicateAtomsAndCombinators|TestRegistryV2MemberParsesAndExposesLockedEnums|TestRegistryS5MemberContainsRegistryPassRows' -v` -> PASS.
- `go test -count=1 ./internal/fieldspec` -> PASS.
- `go test -count=1 ./...` -> PASS, `GO_TEST_EXIT=0` (all packages including `test/fixtures`, `test/replay`, `dogfood`, and `zeroloss`).
- `go vet ./...` -> exit 0.
- `git diff --check 1d3e92c..HEAD` -> exit 0.
- `git diff --name-only 1d3e92c..HEAD` -> exactly the five authorized files.
- `git status --short` -> empty at tip.

Evidence levels: E1 source/diff/commit proof plus E2 local test and vet proof. No E3 runtime or E4 deployed claim.

Out-of-scope preserved: no changes outside the five-file fence; no push, PR, merge, integration, or live verification performed.

Remaining risk: pair review and the dispatched m-4 routing-face plus m-7 registry-hosting fidelity remain required before VP integration and operator merge.

ACTIONS_GIT_REF: `s7a-colgrain@d76c3ad660614f58d72f107330becd465b6a1eaa`; red `10ee3a2`; green `d76c3ad`; worktree `/Users/jack/.config/superpowers/worktrees/frank/s7a-colgrain`
FINAL_GIT_STATUS_SHORT: none - clean tree

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s7a-dispatch` lineage lint exit 0 for both the live directory and this relay

Next requested action: m-2.planner reviews `1d3e92c..d76c3ad`; on approval, return the implementation and pair-review trail to master for m-4/m-7 fidelity routing. Merge remains operator-gated.
