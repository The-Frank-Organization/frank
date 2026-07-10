## REVIEW-FOLD - s4-wire-impl F-GATE-2 fold report

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-impl
RUN_ID: s4
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no new gate; orchestrator-routed hygiene items remain out of fold scope
DESIGN_LOCK_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s4-grill-s4-wire
PLAN_LOCK_ID: s4-slice-4-plan
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator, m-2.planner
IN_REPLY_TO: s4-wire-impl/SITREP-planner-20260705-223638.md
BASE: s4-wire-impl@605b3ef
BRANCH: s4-wire-impl
HEAD: 6a23cf0
SUBJECT: F-GATE-2 folded - five owed FieldSpec rows landed with record_kind_in predicates; rendered-form and validate fixtures green

FOLD_SCOPE:
- internal/fieldspec/registry.json -> in
- internal/fieldspec/registry_test.go -> in
- test/fixtures/s4_shim_test.go -> in
FOLD_SCOPE_RESULT: all-in

## Summary

Folded F-GATE-2 on branch `s4-wire-impl`. Commit: `6a23cf0 s4: fold f-gate-2 owed fieldspec rows`.

`internal/fieldspec/registry.json` now contains the five m-2-confirmed rows:
- `owner`, `source`, `target_surface`, and `disposition_path`: `free_text`/`text`, required only when `record_kind_in: ["owed_item"]`.
- `disposes_owed`: `free_text`/`id_ref`, required only when `record_kind_in: ["owed_disposition"]`.

All five rows have `layer: header`, `fill_constraints: free_text`, `lineage_role: none`, and `gate_referenceable: false`. The predicate form is the adopted `record_kind_in` atom.

The orchestrator-routed hygiene items in the planner relay were not touched: no record_kind seat-scope narrowing, no owner address-typing, and no owed-id picker work.

## Red-First Evidence

- Before the registry rows, `go test -count=1 ./internal/fieldspec -run 'TestRegistryV2MemberContainsOwedRecordRows|TestOwedRecordRequiredWhenPredicatesValidate' -v` failed for the intended missing surface: `missing owner row` and `missing violation owner/required in []`.
- Before the registry rows, `go test -count=1 ./test/fixtures -run TestOperatorDescribeToolsIncludesOwedRecordFields -v` failed for the intended rendered-form gap: `describe form missing owner`.
- After the registry rows, both focused commands passed.

## Fixtures Added

- `internal/fieldspec/registry_test.go`: structural row-shape assertions, structural `record_kind_in` predicate assertions, validate-required legs for each owed field, and a no-record_kind negative leg proving the five fields are not required without the predicate.
- `test/fixtures/s4_shim_test.go`: operator `DescribeTools` form assertion that `record_kind` can select owed kinds and the five owed fields are present at rendered-form grain.

## Owed Store Note

No live store write was performed by this fold. The existing scratch live-store owed open/dispose fixture, `TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition`, passed at the final tip; discharged owed records remain governed by the existing record/projection path, not by a store mutation in this fold.

## Final Verification

Fresh commands from `~/.config/superpowers/worktrees/frank/s4-wire-impl` at `6a23cf0`:

- `go test -count=1 ./internal/fieldspec -run 'TestRegistryV2MemberContainsOwedRecordRows|TestOwedRecordRequiredWhenPredicatesValidate' -v` -> pass.
- `go test -count=1 ./test/fixtures -run TestOperatorDescribeToolsIncludesOwedRecordFields -v` -> pass.
- `go test -count=1 ./test/fixtures -run 'TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition|TestOperatorDescribeToolsIncludesOwedRecordFields' -v` -> pass.
- `go test -count=1 ./...` -> pass for all packages.
- `go vet ./...` -> pass, no output.
- `git diff --check HEAD~1..HEAD` -> pass, no output.
- `git status --short --branch` -> `## s4-wire-impl`.

## Evidence Levels

Implementation and verification evidence is E2. The live host gate remains E3/operator-run.

ACTIONS_GIT_REF: branch `s4-wire-impl` at `6a23cf0`; fold commit `6a23cf0`; pre-edit scope artifact `s4-wire-impl/REVIEW-FOLD-implementer-20260705-224053.md`; relay-substrate writes are the scope artifact, this report, and `.relays/s4/INDEX.md` rows.
FINAL_GIT_STATUS_SHORT: code worktree clean at `## s4-wire-impl`; main checkout tracked status clean with expected gitignored `.relays/` writes.
