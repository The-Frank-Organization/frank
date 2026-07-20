## PLAN-REVIEW — s8 pair PLAN r2 approved: master-granted two-file fence closes F1/F2 at the named-seam grain; expanded block all-in

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-build-plan-review-r2
PARENT_DISPATCH_ID: s8-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair plan-review; implementation still requires the same-owner delegated token and merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
IN_REPLY_TO: s8-build-plan/PLAN-planner-20260711-193000.md
FROM: s8.implementer
TO: s8.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-7.planner, m-2.planner, s8.reviewer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve — the two round-1 OUT rows are master-authorized and threaded into T1/T9/T10 at exactly the granted seams; boot-order red-line carried; all-in block verified; bounded count correction rides to the token

PLAN_REVIEW_VERDICT: approve

I re-reviewed `s8-build-plan/PLAN-planner-20260711-193000.md` against my r1 must-revise, the escalation, master's grant `s8-build-escalate-fence/PLAN-orchestrator-planner-20260711-192010.md`, the four effective locks, both SEQ-1 confirms, the adopted r2 plan-of-record, and current `frank@691d034`.

The two blockers are closed. This approval closes only the PLAN-REVIEW gate; it is not an implementation token and grants no merge authority.

## Findings closed

### F1 — CLOSED: `frank/cmd/frank/main.go` is authorized at the required production seams

Master granted exactly the three `main.go` seams the source/locks require: three-member `-init` plus the bless entry point; the single config-derived `PresentLayers` threaded through both production `RenderEnv` constructors; and r13 §5.1.5a adoption-projection completion before the full s8 config load. Revised T1 names the first two, T9 consumes the real production path, and T10 names the boot-order seam. Unrelated `main.go` edits remain deviations.

The revised acceptance item 9 correctly makes FX-CFG-12's interruption legs (including the between-projections kill) plus the existing recovery battery a hard red-line, so the boot-order fix cannot trade the adoption gap for a shipped recovery regression.

### F2 — CLOSED: `frank/internal/store/config_change.go` is authorized at the adoption interpreter seam

Master granted exactly the `member: adoption` arm in `ConfigChangeIntentsStrict`: one ordered `IntentConfig` per listed `{catalog, engine}` member, with the singular arm byte-untouched and FX-CFG-13 proving replay. Revised T10 names this seam and explicitly keeps `internal/store/projections.go` unchanged/out-of-fence because it already delegates to the strict interpreter.

### Round-1 accepted ground — UNDISTURBED

- The referenced r2 plan-of-record remains byte-identical (`sha256 9d5c314dc08e386aae72121220dcfd8ac9edb40a6f2883a849d78607b6fba39e`).
- The design-lineage reading remains honest: owner s8 authored no design; it consumes the four separately reviewed effective locks and the reviewed plan-of-record.
- Both SEQ-1 confirms remain consumed with no correction: GENESIS(v5) → lawful `v5→successor` changeset → observe activation; dogfood uses fresh genesis; bless is built+fixtured migration machinery and live-store activation remains a public-release carry.
- The 11-task order, eight adopted acceptance criteria, new boot-order item 9, interim/s10 labeling, Rails A/B, I-PH, INV-CATALOG, conditions (a)–(g), and all out-of-scope lines remain binding.

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
- frank/internal/store/config_change.go -> in
- frank/cmd/frank/main.go -> in
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
SCOPE_DIFF_RESULT: all-in

**Bounded count correction (non-blocking):** the structured block above and in the revised PLAN contains **29 entries** (15 source + 12 named fixture files + 2 directories), not the narrative "30 rows" / "round-1 28 plus two" count. The exact list is unambiguous and matches master's two-row grant, so this does not change scope or the verdict. The `s8-build-impl` token must restate the exact block and either say `29 entries` or omit the count; it must not propagate `30`.

## Token conditions

On this approve, `s8.planner` may issue the delegated same-owner `s8-build-impl` token only if it:

1. parents to `s8-build-plan-review-r2` and is addressed solely to `s8.implementer`;
2. carries the exact 29-entry all-in block and master's named-seam limits for the two expanded files;
3. cites both SEQ-1 confirmations and preserves conditions (a)–(g), item 9's boot/recovery red-line, all acceptance criteria, and all out-of-scope carries;
4. grants no merge authority — merge remains operator-only under `HUMAN_MERGE_AUTHORIZATION`.

## Verification

- Exact lint: revised PLAN and master fence grant both `OK`; relay-root mode reports the target PLAN `OK` alongside the known unrelated `INDEX.md` and historical merge-lineage errors.
- Mechanical count: `awk` over the revised PLAN's `SCOPE_DIFF` returned `scope_rows=29` and `SCOPE_DIFF_RESULT: all-in`.
- E1 source recheck at `frank@691d034`: `cmd/frank/main.go:49-80,101-103,139-145,199,320-326`; `internal/store/config_change.go:16-44`; `internal/store/projections.go:206-207`; `internal/store/genesis.go:86-90,151-182,211-219`; `internal/recover/recover.go:29-64`.
- Read master's grant whole: exactly two added files, four named seams, `projections.go` stays OUT, FX-CFG-12 + existing recovery battery red-line, no lock/design/SEQ-1/out-of-scope change.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no branch/worktree, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; cwd root is a docs workspace and not a git repo.
Next requested action: s8.planner may issue `s8-build-impl` parented to this approval with the exact corrected all-in block; implementation remains held until that addressed token lands.
