## REVIEW-FOLD - VP F1 provenance corrected red-first at 2bc0763; exact two-file delta ready for pair review

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s7a-plan-m2-fold-f1
PARENT_DISPATCH_ID: s7a-plan-m2-fold-f1-scope
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - fold report and pair-review request only; m-7 final-bytes re-confirm, VP re-approval, and operator merge remain downstream
IN_REPLY_TO: master/relays/s7a-dispatch/SITREP-orchestrator-planner-20260710-145439.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.implementer, m-4.implementer, m-7.planner, m-4.planner
SUBJECT: VP F1 folded at s7a-colgrain@2bc0763; exact provenance tuple pinned and final two-file delta ready for pair review

Summary: Replaced the weak provenance-owner nonempty check with a byte-exact four-value tuple assertion, demonstrated it failing against the stale s6 attribution, then changed only those four registry provenance values to truthful s7a attribution.

FOLD_SCOPE:
- internal/fieldspec/registry_test.go -> in
- internal/fieldspec/registry.json -> in
- master/relays/s7a-dispatch/FOLD_SCOPE-implementer-20260710-145847.md -> in
- master/relays/s7a-dispatch/REVIEW-FOLD-implementer-20260710-150027.md -> in
- master/relays/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

### Exact provenance tuple

- `owner`: `m-2`
- `design_doc_id`: `F-S7-R2-COLGRAIN`
- `plan_lock_id`: `s7a-plan-m2`
- `note`: `s7a m-2 pair build under the operator B10 second-application ruling; restores c1 section 5 column-grain fidelity`

These values identify the owning pair, the governing audit finding, the clean-chain plan of record, the operator staffing ruling, and the c1 fidelity restoration. They no longer route a v5 reader to the superseded s6 pass.

### Red/green proof

- Red `37ac1dc` (`test: pin s7a registry provenance`): exact map assertion added first. `go test -count=1 ./internal/fieldspec -run '^TestRegistryV2MemberParsesAndExposesLockedEnums$' -v` failed with the complete observed s6 tuple versus the expected s7a tuple.
- Green `2bc0763` (`fix: attribute fieldspec v5 to s7a`): changed only the four provenance values in `registry.json`; the same focused test then passed.

### AC6 wording amendment

AC6 now reads: **one semantic row delta (the singleton `routing_assignments.gate_referenceable_columns = ["declared_deviated"]`) plus the required provenance-attribution metadata update**. The provenance fold does not add or alter any semantic field row.

### Non-movement proof

- `git diff --name-only d76c3ad..2bc0763` returns exactly `internal/fieldspec/registry.json` and `internal/fieldspec/registry_test.go`.
- The JSON fold changes only `provenance.owner`, `provenance.design_doc_id`, `provenance.plan_lock_id`, and `provenance.note`.
- The guard, Go member, singleton allowlist, `s7a-fieldspec-v5` version, named enums, field rows, and every other registry byte are unchanged.
- `git diff --name-only 1d3e92c..2bc0763` remains exactly the original five-file s7a fence.

Tests/verification:
- Focused four-test fieldspec command -> PASS, including exact provenance, legal `declared_deviated`, and all three default-deny negatives.
- `go test -count=1 ./internal/fieldspec` -> PASS.
- `go test -count=1 ./...` -> PASS, `GO_TEST_EXIT=0`, including `test/fixtures`, replay, dogfood, and zeroloss.
- `go vet ./...` -> exit 0.
- `git diff --check d76c3ad..HEAD` -> exit 0.
- `git status --short` -> empty at `2bc0763`.

Evidence levels: E1 exact source/diff/commit proof plus E2 red/green, package, full-suite, and vet proof. No E3/E4 claim.

Out-of-scope preserved: no predicate behavior, registry shape, singleton, version, engine/store/seat surface, m-4 semantic surface, push, PR, merge, or live-verification change.

ACTIONS_GIT_REF: `s7a-colgrain@2bc0763`; red `37ac1dc`; green `2bc0763`; fold base `d76c3ad`; worktree `/Users/jack/.config/superpowers/worktrees/frank/s7a-colgrain`
FINAL_GIT_STATUS_SHORT: none - clean tree

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s7a-dispatch` lineage lint exit 0 for both the live directory and this relay

Next requested action: m-2.planner reviews the exact `d76c3ad..2bc0763` delta under a unique pair sub-ID, then returns it to master for m-7 final-bytes composite-digest/s8-genesis re-confirm. Merge remains separately operator-gated.
