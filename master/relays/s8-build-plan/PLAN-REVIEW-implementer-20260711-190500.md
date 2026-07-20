## PLAN-REVIEW — s8 pair PLAN must revise: the adopted r2 scope fence omits two production seams required by locked T1/T10 behavior

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-build-plan-review-r1
PARENT_DISPATCH_ID: s8-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this is a mechanical scope correction; master must reauthorize the fence before approval, while slice merge remains separately operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2
IN_REPLY_TO: s8-build-plan/PLAN-planner-20260711-184500.md
FROM: s8.implementer
TO: s8.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-7.planner, m-2.planner, s8.reviewer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must-revise — adoption, lineage, locks, SEQ-1, acceptance, and execution mapping are sound, but the mechanical scope check finds cmd/frank/main.go and internal/store/config_change.go OUT and required by the locked config/bless contract

PLAN_REVIEW_VERDICT: must-revise

The fresh-pair PLAN is correctly addressed and correctly grants no implementation authority. Its byte-reference adoption of `master/domains/m-3-observation-evidence/plan/2026-07-11-s8-observe-spine-plan.md` resolves to the reviewed r2 artifact (`sha256 9d5c314dc08e386aae72121220dcfd8ac9edb40a6f2883a849d78607b6fba39e`); the four locks are effective; both SEQ-1 confirms are consumed without correction; the eight acceptance criteria and out-of-scope carries survive; and the fresh-pair sole-writer/fidelity mapping is consistent with the later operator staffing ruling.

The design-lineage reading in PLAN §2 is also accepted. Owner `s8` authored no design doc: this PLAN consumes effective, separately reviewed locks plus the already-reviewed r2 plan-of-record. Adding `DESIGN_RECORD_KIND: design-doc` would falsely imply a same-owner s8 DESIGN → DESIGN-REVIEW chain; `audit-record` would be false. The existing `PLAN_LOCK_ID` plus explicit effective-lock references are the honest shape.

Approval is blocked by the binding do-not-exceed scope fence. Current source proves two files outside that fence are required to implement the locked behavior; this is `SCOPE_DIFF_RESULT: deviation-present`, so the pair cannot self-authorize them and `s8.planner` must not issue `DISPATCH IMPL`.

## Blocking findings

### F1 — `frank/cmd/frank/main.go` is OUT but required by T1/T9/T10

The adopted file map assigns three-member genesis, config-derived `PresentLayers`, restart activation, and offline bless to in-fence library files, but the production composition root lives in `cmd/frank/main.go`:

- `cmd/frank/main.go:49-80` owns CLI flags and `-init`; it supplies only `{fieldspec, engine}` to `store.Init`, so neither the required catalog source/member nor a `bless-s8` offline mode can be reached without changing this file.
- `cmd/frank/main.go:139-145` and `:320-326` construct the two production `fieldspec.RenderEnv` values. T1's one immutable config-derived `PresentLayers` cannot reach render, validate, and grant-digest on the running surface unless these production constructors consume it.
- `cmd/frank/main.go:101-103` performs full `config.Load(StoreRootConfigPaths(...))` before `frankrecover.RunWithProcessor` at `:199`. The r13 §5.1 adoption contract requires all adoption member projections to recover before the full s8 load; the current order is the opposite. The boot sequence must change at this composition root (or another separately authorized production seam), not only in `genesis.go`.

This is not optional CLI polish: without the file, FX-CFG-1/7/12 cannot prove the locked production path they claim.

### F2 — `frank/internal/store/config_change.go` is OUT but required by the adoption record contract

`internal/store/config_change.go:16-44` is the canonical `ConfigChangeIntentsStrict` interpreter used by canonical projection rebuild. It resolves exactly one `configTarget` and emits exactly one `IntentConfig` containing the whole record body. The locked adoption variant instead requires a closed `{catalog, engine}` body to emit exactly two ordered `IntentConfig` projections while preserving the unchanged singular arm. `internal/store/projections.go:206-207` routes every `config_change` rebuild through this interpreter.

Changing only in-fence `genesis.go` and `submit.go` cannot make accepted adoption records replay/recover through the canonical interpreter. `internal/store/config_change.go` must be in the authorized source fence (no change to `projections.go` is presently required because it already delegates to the strict interpreter).

## Mechanical scope check

SCOPE_DIFF:
- frank/internal/observe/gate.go -> in
- frank/internal/observe/registry.go -> in
- frank/internal/observe/checks_base.go -> in
- frank/internal/executor/executor.go -> in
- frank/internal/engine/submit.go -> in
- frank/internal/engine/loop.go -> in
- frank/internal/engine/completeness.go -> in
- frank/internal/config/config.go -> in
- frank/internal/fieldspec/predicate.go -> in
- frank/internal/fieldspec/validate.go -> in
- frank/internal/fieldspec/render.go -> in
- frank/internal/fieldspec/registry.json -> in
- frank/internal/store/genesis.go -> in
- frank/test/fixtures/s8_config_activation_test.go -> in
- frank/test/fixtures/s8_registry_changeset_test.go -> in
- frank/test/fixtures/s8_observe_spine_test.go -> in
- frank/test/fixtures/s8_check_registry_e1_test.go -> in
- frank/test/fixtures/s8_executor_test.go -> in
- frank/test/fixtures/s8_verdict_binding_test.go -> in
- frank/test/fixtures/s8_pipeline_stages_test.go -> in
- frank/test/fixtures/s8_decision2_test.go -> in
- frank/test/fixtures/s8_suppliability_guard_test.go -> in
- frank/test/fixtures/s8_exit_gate_test.go -> in
- frank/test/fixtures/s8_adoption_test.go -> in
- frank/test/fixtures/s8_adversarial_test.go -> in
- frank/test/invariants/ -> in
- frank/.relays/s8/ -> in
- frank/cmd/frank/main.go -> OUT
- frank/internal/store/config_change.go -> OUT
SCOPE_DIFF_RESULT: deviation-present

The master-visible `master/relays/s8-build-*` files are governance transport artifacts, not implementation-source expansion; they do not cure the two OUT source rows.

## Required revision / routing

1. Do not issue `s8-build-impl`.
2. Relay this deviation to `master.orchestrator-planner` under condition (b)/(c)'s stop-and-escalate discipline and request an expanded implementation fence containing at least `frank/cmd/frank/main.go` and `frank/internal/store/config_change.go`.
3. After master reauthorization, revise the execution/file map so T1 names the production `RenderEnv`/three-member-init seam and T10 names the CLI/pre-load-recovery + multi-member-intent seam; reissue the PLAN for review with the expanded mechanical block.
4. Preserve every existing lock, SEQ-1 order, acceptance criterion, carry, and out-of-scope line. This finding asks for no design change.

## Verification

- Exact PLAN lint with relay-root context: target file reported `OK`; root mode also reported pre-existing unrelated `INDEX.md` and historical merge-lineage errors.
- Read the adopted r2 plan in full; file map `:55-68`, T1 `:74-82`, T10 `:172-179`, acceptance `:191-200`, out-of-scope `:202-203`.
- Read all four effective locks, both SEQ-1 confirms, the orchestrator dispatch/addendum, master reconcile declaration, prior r2 plan review, and prior gate-clearance token.
- E1 source inspection at `frank@691d034`: `cmd/frank/main.go:49-80,101-103,139-145,199,320-326`; `internal/store/config_change.go:16-44`; `internal/store/projections.go:206-207`; `internal/store/genesis.go:86-90,151-182,211-219`; `internal/recover/recover.go:29-64`.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no branch/worktree, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; cwd root is a docs workspace and not a git repo.
Next requested action: s8.planner routes the deviation to master for fence expansion, then reissues the PLAN; implementation remains held.
