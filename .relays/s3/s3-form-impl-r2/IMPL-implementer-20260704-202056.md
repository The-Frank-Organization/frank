## IMPL - s3-form implementation r2 partial/blocker report

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s3-form-impl-r2-implementer
PARENT_DISPATCH_ID: s3-form-impl-r2
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - merge stays the operator's separate S3-close gate; this report does not request merge
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s3-grill-s3-form
DESIGN_DOC_ID: s3-slice-3-design
DESIGN_LOCK_ID: s3-slice-3-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s3-slice-3-plan
IN_REPLY_TO: s3-form-impl-r2/IMPL-planner-20260704-193915.md
FROM: s3-form.implementer
TO: s3-form.planner
CC: s3.orchestrator-planner, operator
SUBJECT: implementation branch built and verified at s3-form-impl@b4a1764, but Task 11 full-replay acceptance is not met by the current replay master-docs/report; no S3-close claim

Summary:
- I resumed from the superseding r2 dispatch, reset the implementation branch onto `main@354718b`, and implemented the S3 form/registry line on branch `s3-form-impl`.
- The branch tip is `s3-form-impl@b4a1764`.
- The final branch diff is inside the r2 SCOPE_DIFF file set. During the final scope audit I found a transient Task 6 touch to `internal/obligation/owed.go`; commit `b4a1764` restores that helper to the main-branch shape, so it is absent from the final `main..HEAD` diff.
- Full local verification is green at `b4a1764`.
- I am blocking S3-close/readiness because the locked Task 11 acceptance says FULL 243-fixture replay, both legs, zero uncovered fail-side rows; the current replay master-docs/report still permits and records `uncovered-S3` rows.

## Branch and commits

Branch:
- `s3-form-impl`

Base:
- `main@354718b`

Tip:
- `b4a1764`

Commits on branch:
- `6bb0804` s3 IMPL: registry v2 data model + load-time checks + member content (S3-G1..G5)
- `5e00ef7` s3 IMPL: bounded predicate evaluator + CQ-1 layer wrap (D-3)
- `27592b2` s3 IMPL: canonical JSON-in-string typed carrier (D-2)
- `772eb5e` s3 IMPL: registry-driven render + full-context digest with candidate-set split (D-4)
- `0b72357` s3 IMPL: registry-driven validate, live digest, full enum sweep; 6-enum dialect deleted (S3-A2..A4, G6/G7, RR1/RR1b)
- `c36f927` s3 IMPL: runtime tables; live path off full-store rescans (F-P1, S3-P1)
- `2ed9a1a` s3 IMPL: the section 10c lineage engine over tables; grant-narrowing carry landed (S3-L1..L7)
- `b59f6b3` s3 IMPL: serve rendered form schema and live digest via descriptions (S3-E2)
- `e1f26c1` s3 IMPL: migrator registry + read-time apply + refusal legs, zero production migrators (S3-V1..V5)
- `e995a22` s3 IMPL: assemble pinned v2 registry fresh-store path (S3-E1, V6, RR2)
- `e751d69` s3 IMPL: disposition artifacts and replay coverage
- `d9ad2d5` s3 IMPL: regression floor and claim sweeps
- `b4a1764` s3 IMPL: restore owed helper outside dispatch scope

## Acceptance status

Implemented and verified locally:
- Registry v2 load model, member data, load-time predicate checks, GRILL rows, and row-array carrier.
- Bounded predicate evaluator and observe-layer wrapping.
- Canonical JSON-in-string typed carrier.
- Registry-driven render, full-context digest, live digest validation, seat scope, enum sweep, and old 6-enum dialect removal.
- Runtime tables and live submit/complete paths using table state.
- Lineage engine over tables, including non-addressee IMPL, merge-grant prerequisite, scope-flip drift, active candidate derivation, and grant state.
- Served form schema/digest through `tools/descriptions`.
- Migration registry/read facade/refusal legs and fresh-store pinned registry assembly.
- README fresh-store qualifier delta under the r5 fence ruling.
- Task 12 claim/I-PH sweeps, including code-only deletion of `isAGateCategory` and README qualifier coverage.

Blocked/not claimable:
- Task 11 as locked is not fully satisfied.
- Evidence: `test/replay/classmap.go` walks the external 243-file relay-lint corpus, but `test/replay/report.md` still has an `## uncovered-S3` section with many rows.
- Evidence: `test/replay/replay_test.go` still asserts `scope-diff-row-array` and `fold/FD1-fold-edit-no-foldscope.md` as `uncovered-S3`, while the locked Task 11 requires all FAIL-side rows caught-or-obsolete and zero uncovered.
- Therefore I am not claiming "FULL replay, both legs green" or S3 exit readiness.

## Scope proof

Final `git diff --name-only main..HEAD` at `b4a1764`:
- `README.md`
- `cmd/frank/main.go`
- `docs/sprints/2026-07-04-s3-slice-3/results/disposition-table.md`
- `internal/channel/server.go`
- `internal/engine/loop.go`
- `internal/engine/pipeline_test.go`
- `internal/engine/submit.go`
- `internal/fieldspec/canonical.go`
- `internal/fieldspec/canonical_test.go`
- `internal/fieldspec/fieldspec.go`
- `internal/fieldspec/fieldspec_test.go`
- `internal/fieldspec/predicate.go`
- `internal/fieldspec/predicate_test.go`
- `internal/fieldspec/registry.go`
- `internal/fieldspec/registry.json`
- `internal/fieldspec/registry_test.go`
- `internal/fieldspec/render.go`
- `internal/fieldspec/render_test.go`
- `internal/fieldspec/validate.go`
- `internal/fieldspec/validate_test.go`
- `internal/lineage/engine_test.go`
- `internal/lineage/lineage.go`
- `internal/lineage/lineage_test.go`
- `internal/migrate/migrate.go`
- `internal/migrate/migrate_test.go`
- `internal/obligation/obligation.go`
- `internal/obligation/owed_test.go`
- `internal/tables/tables.go`
- `internal/tables/tables_test.go`
- `test/fixtures/f11_test.go`
- `test/fixtures/iph_test.go`
- `test/fixtures/main_assembly_test.go`
- `test/fixtures/s2setup_test.go`
- `test/fixtures/sweep_test.go`
- `test/replay/classmap.go`
- `test/replay/dispositions.json`
- `test/replay/replay_test.go`

No `../master` or `../extracted` writes were made by the implementation branch. The replay harness reads the external corpus under `../extracted`, but does not write it.

## Verification evidence

Fresh final runs from `~/.config/superpowers/worktrees/frank/s3-form-impl` at `b4a1764`:
- `go test -count=1 ./...` -> PASS for all packages.
- `go vet ./...` -> PASS, no output.
- `go test -race -count=1 ./internal/engine ./internal/tables ./internal/fieldspec` -> PASS for all three packages.
- `rg -n 'func isAGateCategory|isAGateCategory\(' --glob '*.go'` -> no matches.
- `rg -n 'the S3 registry rides \`store\.Init\`|registry evolution on an existing store awaits the §7 config-change record' README.md test/fixtures/sweep_test.go` -> matches `README.md` and `test/fixtures/sweep_test.go`.
- Implementation worktree status after verification: clean.
- Main relay checkout status before writing this relay: clean.

## Requested planner action

Please adjudicate Task 11:
- If the intended scope is the current representative/corpus-walk disposition artifact, issue an explicit fold narrowing the FULL replay acceptance.
- If the locked FULL 243-fixture replay remains required, dispatch a focused follow-up with the required pass/fail classification source and expected per-fixture mapping, or direct me to continue Task 11 under the existing branch.

PR: none.
Merge: none.
S3-close claim: none.

ACTIONS_GIT_REF: branch `s3-form-impl` at `b4a1764`; source/test/docs commits listed above; gitignored relay report `.relays/s3/s3-form-impl-r2/IMPL-implementer-20260704-202056.md` plus `.relays/s3/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: implementation worktree clean at `s3-form-impl@b4a1764`; main relay checkout clean before this gitignored relay/index write.
