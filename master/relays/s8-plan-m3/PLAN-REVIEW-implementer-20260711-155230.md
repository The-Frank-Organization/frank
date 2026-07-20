## PLAN-REVIEW - s8 observe-spine plan must revise missing authority_class and step-4.5 completeness work

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-plan-m3-review-r1
PARENT_DISPATCH_ID: s8-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - pair plan-review; no new operator decision surfaced
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r1
IN_REPLY_TO: s8-plan-m3/PLAN-planner-20260711-154500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must-revise - plan maps the m-3 spine and SEQ-1, but omits the executable owner/task for m-2 authority_class and m-7 step-4.5 profile-aware completeness before decision-2 can run

PLAN_REVIEW_VERDICT: must-revise

I reviewed `s8-plan-m3/PLAN-planner-20260711-154500.md` and `master/domains/m-3-observation-evidence/plan/2026-07-11-s8-observe-spine-plan.md` against the orchestrator dispatch `s8-dispatch/PLAN-orchestrator-planner-20260711-145129.md`, the effective m-3 registry lock r1/r1a/r1b, the m-2 grammar lock, the m-7 executor/config locks, and current `frank@691d034`.

The plan is close and the core observe-spine decomposition is sound, but I cannot approve the delegated implementation chain yet. One lock-bearing build stage is missing from the task graph. If the planner folds that stage explicitly, the rest of this review gives the all-in scope fence and residual pre-token gates to preserve.

## Finding

### F1 - BLOCKER: decision-2 consumes `authority_class`, and Option-B consumes step-4.5 completeness, but the plan never assigns their executable build task

The m-2 grammar lock does not merely say "remove the `surface_intent` registry predicates." It locks a staged pipeline:
- step 3 computes `authority_class` after `slot_in` classify and before observe, so m-3's disposition reads a computed value, not a lane-supplied or absent header.
- step 4 is m-3's ten-field observe manifest.
- step 4.5 is an m-7-hosted profile-aware completeness stage that validates producer manifests and derives Option-B `surface_intent` for non-gate-bearing records only, before append.

Evidence:
- `master/domains/m-2-forms-determinism/design/2026-07-10-s8-config-atom-grammar.md:64-66` locks the step-3 `authority_class` compute and step-4.5 profile-aware completeness/`surface_intent` derivation.
- `...s8-config-atom-grammar.md:67-77` states that `authority_class` is not an observe output, `surface_intent` is not an observe output, and the fixture must prove m-3 cannot write either field.
- `...s8-config-atom-grammar.md:91-94` names the same seam in the m-7 config boundary; `...:138-140` makes it part of design-lock impact.
- Current `frank@691d034` has only registry declarations for `authority_class`/`surface_intent`, not an implementation of the compute/completeness stage: `rg -n "authority_class|surface_intent|step-4\\.5|profile-aware|producer/profile" frank/internal frank/test --glob '*.go' --glob '*.json'` returns `registry.json` rows and dormancy tests, no executable stage.

The plan currently loses that locked stage:
- The file map assigns `surface_intent` only to the registry changeset row, and assigns no file/task for computing `authority_class` or running step-4.5 completeness (`master/domains/m-3-observation-evidence/plan/2026-07-11-s8-observe-spine-plan.md:49-60`).
- T2 removes `surface_intent`'s static predicates and adds tokens, but does not implement the profile-aware derivation/completeness stage (`...:76-83`).
- T3 says `observe.Gate` reads, never writes, `slot_in`/`authority_class` (`...:87-94`), and T7 keys decision-2 on computed `authority_class` (`...:125-132`), but no earlier task produces that computed value.
- The plan's self-review maps `s8-design-m2-grammar` only to T1/T2 (`...:201-205`), omitting the lock's step-3/step-4.5 obligations.

Required revision:
- Add an explicit build task, or split into existing tasks with explicit owner gates, for:
  1. computing `authority_class` at the locked step-3 point, after `slot_in`/gate-category classification and before m-3 disposition reads it;
  2. running the m-7-hosted step-4.5 profile-aware completeness stage before append;
  3. deriving Option-B `surface_intent` only for non-gate-bearing records, defaulting to `progress`, absent on gate-bearing records;
  4. proving m-3 cannot write `authority_class` or `surface_intent`, no producer writes outside its manifest, and the committed behavior matches the m-2 fixture oracle.
- Name the owner split. m-2 owns the `authority_class` formula and the registry-row amendment; m-5/m-6 own the `surface_intent` profile contract; m-7 hosts the step-4.5 pipeline stage; m-3 consumes the values for disposition and observe output completeness.
- Thread that task into the file map, milestone dependencies, acceptance criteria, and scope diff. Without it, decision-2 can only read an absent/stale/lane-supplied `authority_class`, and Option-B `surface_intent` is merely removed from FieldSpec requiredness rather than produced anywhere.

## Accepted Shape, Pending F1 Fold

1. SEQ-1 is correctly treated as a pre-token gate, not a silent assumption. `s8-plan-m3-seq1/SITREP-planner-20260711-153000.md` exists and asks m-7/m-2 to confirm the genesis-first, `v5 -> successor`, fresh-genesis dogfood, bless-as-migration order. A concurrent m-7 response, `s8-plan-m3-seq1/SITREP-planner-20260711-151621.md`, is now present and confirms m-7's half; I found no m-2 half in the `s8-plan-m3-seq1` relay directory at this review point. The revised plan may still approve if it keeps "no delegated token until both confirms are in hand; correction means reissue affected tasks."

2. The vertical slice order T3-T6 is otherwise coherent. `SubmitHandler` returns a record/intents at `frank/internal/engine/submit.go:24-105`, and `engine.Loop.process` commits that result immediately after handler return at `frank/internal/engine/loop.go:142-159`, so a handler-side observe hook can be pre-append if it runs before returning accepted intents.

3. The timeout/no-vantage correction is properly encoded in T7 as a fixture pair: killed check -> authority `held` / non-authority `rejected`+fault-edge, separate no-vantage -> `accepted`+`self_reported`.

4. The executor boundary in T5 matches the effective m-7 lock at the plan level: closed `CheckVerdict` only, RunResult host-internal, per-run provided writable surface, exit-confirmed cleanup, survivor machinery fault with preserved workdir, no OS-sandbox claim.

5. The §6.1 I-PH contract is present in the global constraints and in T4/T5/T6/T11. Keep it conductor-side: param refusal before spawn, verdict redaction after return, no raw store/config/outbox/socket path or effective config values in seat-visible surfaces.

## Mechanical Scope Diff

This is the all-in fence I would approve after F1 is folded. If the revised task adds a new file for step-3/step-4.5, include it explicitly; if it edits a file outside this set, escalate before token.

SCOPE_DIFF:
- frank/internal/observe/gate.go -> in
- frank/internal/observe/registry.go -> in
- frank/internal/observe/checks_base.go -> in
- frank/internal/executor/executor.go -> in
- frank/internal/engine/submit.go -> in
- frank/internal/engine/loop.go -> in
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
- frank/test/fixtures/s8_decision2_test.go -> in
- frank/test/fixtures/s8_suppliability_guard_test.go -> in
- frank/test/fixtures/s8_exit_gate_test.go -> in
- frank/test/fixtures/s8_adoption_test.go -> in
- frank/test/fixtures/s8_adversarial_test.go -> in
- frank/test/invariants/ -> in
- frank/.relays/s8/ -> in
SCOPE_DIFF_RESULT: all-in

## Gates and Residuals

- No implementation token is authorized from this review. A delegated dispatch must wait for a revised PLAN and an approving PLAN-REVIEW.
- SEQ-1 m-7/m-2 confirmations remain a pre-token gate unless folded into the revised plan as already received.
- No operator merge authority is implied; merge remains operator-only with `HUMAN_MERGE_AUTHORIZATION`.
- If implementation discovers a need to change any locked design contract, add schema/registry changes beyond the scheduled T2 changeset, or touch cross-domain m-1/m-6 shapes, stop and route to master.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-plan-m3/PLAN-planner-20260711-154500.md` - OK.
- Read `master/relays/s8-plan-m3/PLAN-planner-20260711-154500.md` - confirmed `FROM: m-3.planner`, `TO: m-3.implementer`, `AUTHORITY: plan-only`, `DESIGN_LOCK_ID: s8-design-m3-registry`, and PLAN-REVIEW request.
- Read `master/domains/m-3-observation-evidence/plan/2026-07-11-s8-observe-spine-plan.md` - checked file map, tasks T1-T11, acceptance criteria, out-of-scope lines, boundary contract, delegation conditions, and self-review.
- Read `s8-dispatch/PLAN-orchestrator-planner-20260711-145129.md` - confirmed delegated dispatch authority conditions and scope-of-record.
- Read m-3 effective design r1/r1a/r1b and prior m-3 implementer approval; checked timeout/I-PH/provided-surface folds.
- Read m-2 grammar lock lines `:64-77`, `:91-94`, `:138-140`; checked current `frank` with `rg` for `authority_class`/`surface_intent` executable support.
- Read current `frank/internal/engine/submit.go` and `frank/internal/engine/loop.go`; confirmed handler-side observe can be pre-append if it runs before handler return.
- Read `s8-plan-m3-seq1/SITREP-planner-20260711-151621.md` after its concurrent index row appeared; exact lint OK and it confirms only the m-7 half of SEQ-1.
- `git -C frank rev-parse --short HEAD` - `691d034`; `git -C frank status --short` - clean.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no branch/worktree, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: cwd root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git`; `git -C frank status --short` returned clean.
Next requested action: m-3.planner folds F1 into the s8 observe-spine plan, preserves the SEQ-1 pre-token confirm gate, and returns a revised PLAN for another PLAN-REVIEW.
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-plan-m3/PLAN-REVIEW-implementer-20260711-155230.md`; narrow dispatch-root lint is not the proof of record here because `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s8-plan-m3` reports the incoming PLAN's design-lineage context is outside this narrow root: `PLAN-planner-20260711-154500.md: DESIGN_LOCK_ID 's8-design-m3-registry' has no earlier same-owner DESIGN relay carrying matching DESIGN_DOC_ID`.
