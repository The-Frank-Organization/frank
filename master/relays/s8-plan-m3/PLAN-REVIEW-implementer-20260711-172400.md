## PLAN-REVIEW - s8 observe-spine plan r2 approved; F1 folded and addendum consumed

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-plan-m3-review-r2
PARENT_DISPATCH_ID: s8-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - pair plan-review; implementation token remains delegated to m-3.planner and merge remains operator-only
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2
IN_REPLY_TO: s8-plan-m3/PLAN-planner-20260711-171500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve - r2 folds F1 through Task 6.5, consumes the s8-dispatch addendum, and keeps the all-in scope fence with completeness and pipeline fixture rows

PLAN_REVIEW_VERDICT: approve

I reviewed `s8-plan-m3/PLAN-planner-20260711-171500.md` and the revised r2 plan doc `master/domains/m-3-observation-evidence/plan/2026-07-11-s8-observe-spine-plan.md` against my prior must-revise `s8-plan-m3/PLAN-REVIEW-implementer-20260711-155230.md`, the `s8-dispatch` addendum `SITREP-orchestrator-planner-20260711-163144.md`, both SEQ-1 owner confirmations, and current `frank@691d034`.

The r2 plan is approved for the next delegated implementation-token step. This review grants no implementation token by itself; it closes the pair review gate so m-3.planner may issue that token under the standing dispatch conditions.

## Findings Closed

### F1 - CLOSED: the missing staged pipeline is now an executable task with named owners and fixtures

Task 6.5 now supplies the lock-bearing stage that r1 lacked:
- step-3 computes `authority_class` in `frank/internal/engine/submit.go` after `slot_in`/gate-category classify and before observe/disposition;
- step-4.5 creates the m-7-hosted profile-aware completeness stage in `frank/internal/engine/completeness.go`;
- Option-B `surface_intent` is derived only for non-gate-bearing records, defaults to `progress`, and is absent on gate-bearing records;
- fixtures prove `authority_class` is computed, `surface_intent` derivation/absence is correct, producer-manifest violations reject, and m-3 cannot write either field.

Evidence:
- The file map names `submit.go`, `loop.go`, and new `completeness.go` for T6.5 (`master/domains/m-3-observation-evidence/plan/2026-07-11-s8-observe-spine-plan.md:61-63`).
- Task 6.5 names the owner split: m-2 owns the formula and registry amendment; m-5/m-6 own the already-locked surface-intent producer/profile contract; m-7 hosts step-4.5; m-3 consumes (`...:131-140`).
- T7 now explicitly consumes T6.5's computed `authority_class` instead of assuming it exists (`...:147`).
- Acceptance criterion 7 and self-review lock coverage now include the staged pipeline and oracle (`...:199`, `...:223`).

### Addendum - CLOSED: SEQ-1 steer, Rails A/B, and the s9 seam are folded without reopening locks

The `s8-dispatch/…-163144` addendum is consumed correctly:
- fresh-genesis-first is the operator-ratified SEQ-1 default; T10 still builds and fixtures bless/adoption, while activation against the actual shipped live store is a named public-release-migration carry (`...:51`, `...:203`);
- Rails A/B are global build-time review criteria for new surfaces, with T6.5 stating Rail A/B inline (`...:32`, `...:135`, `...:200`);
- s8 remains deterministic E1/E2 only, and the s9 fuzzy-claim adjudication rung is out-of-scope with the check-registry seam kept additive/open (`...:33`, `...:203`).

## SEQ-1 Gate State

The r2 plan text still says the m-2 half is pending because the m-2 confirm landed after the r2 doc was cut. That is stale-but-conservative, not a blocker: the plan's hard condition remains "no delegated token until both confirms are in hand; correction means reissue affected tasks."

Both confirms are now present and lint-clean:
- `s8-plan-m3-seq1/SITREP-planner-20260711-151621.md` - m-7 confirms genesis pins v5, FX-CFG-7 stays valid, the changeset is post-genesis `v5 -> successor`, and bless bypasses only the seat fill-gate while trusted validation still runs.
- `s8-plan-m3-seq1/SITREP-planner-20260711-153500.md` - m-2 confirms both `adoption` and `catalog` tokens land in the one step-2 changeset, `v5 -> successor` is the lawful forward fieldspec transition, and the changeset-before-observe order is grammar-required. No correction to the plan order.

Required token-note carry: when m-3.planner issues the delegated implementation token, cite both SEQ-1 confirmations as consumed so the stale "m-2 pending" status is not propagated into build execution.

## Mechanical Scope Diff

This is the all-in fence approved for the delegated token. The r2 additions from F1 are included explicitly.

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
SCOPE_DIFF_RESULT: all-in

## Gates and Residuals

- This review does not authorize implementation directly. m-3.planner's delegated token must parent to this approval and preserve dispatch conditions (a)-(g).
- Merge remains operator-only with `HUMAN_MERGE_AUTHORIZATION`.
- Any registry/schema change beyond the scheduled T2 changeset, locked-contract change, or cross-domain m-1/m-6 shape change stops the build and routes to master.
- The s9 fuzzy-claim adjudication rung and activation of bless/adoption against the actual shipped live store remain out-of-scope carries, not s8 build work.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-plan-m3/PLAN-planner-20260711-171500.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-dispatch/SITREP-orchestrator-planner-20260711-163144.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-plan-m3-seq1/SITREP-planner-20260711-151621.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-plan-m3-seq1/SITREP-planner-20260711-153500.md` - OK.
- Read `master/domains/m-3-observation-evidence/plan/2026-07-11-s8-observe-spine-plan.md` - checked r2 revision line, global Rails/s9 constraints, SEQ-1 order/status, file map, T6.5, T7 consumes line, acceptance criteria, out-of-scope, and self-review.
- Read `master/relays/s8-plan-m3/PLAN-planner-20260711-171500.md` - confirmed `FROM: m-3.planner`, `TO: m-3.implementer`, `AUTHORITY: plan-only`, `PLAN_LOCK_ID: s8-observe-spine-plan-r2`, and re-review request.
- Read both SEQ-1 confirms; m-7 and m-2 confirm no correction to the order.
- `git -C frank rev-parse --short HEAD` - `691d034`; `git -C frank status --short` - clean.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no branch/worktree, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: cwd root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git`; `git -C frank status --short` returned clean at `691d034`.
Next requested action: m-3.planner may issue the delegated implementation token, parented to this approval, after citing both SEQ-1 confirmations as consumed and preserving the all-in scope fence above.
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-plan-m3/PLAN-REVIEW-implementer-20260711-172400.md`.
