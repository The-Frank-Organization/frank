## DISPATCH — s7a F-S7-R2-COLGRAIN implementation authorized on the clean chain: unique DISPATCH_ID, parenting to review-r1; all four conditions + AC6 met; the atomicity condition carried

ROLE: Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7a-plan-m2-dispatch
PARENT_DISPATCH_ID: s7a-plan-m2-review-r1
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
SCOPE_DIFF_RESULT: all-in
HUMAN_GATE_REQUIRED: no — delegated dispatch under `PLAN-orchestrator-planner-20260710-125439` (DELEGATED_DISPATCH_AUTHORITY: yes) + master option-1 ruling `135739`; the s7a merge stays operator-gated
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-dispatch/PLAN-REVIEW-implementer-20260710-140100.md
FROM: m-2.planner
TO: m-2.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-4.planner, m-4.implementer, m-7.planner, m-7.implementer
SUBJECT: implementation authorized (bare token below, own line) — build the approved s7a plan `PLAN-planner-20260710-140000.md` (DISPATCH_ID s7a-plan-m2); clean chain (token → review-r1 → s7a-plan-m2 PLAN); the four conditions + AC6 met with the atomicity condition carried

**All delegated-dispatch conditions are satisfied on the clean chain — implementation is authorized.**

1. **approve:** m-2.implementer `PLAN-REVIEW-implementer-20260710-140100.md` (`DISPATCH_ID: s7a-plan-m2-review-r1`) — `approve`, a real re-assertion parenting to the unique `s7a-plan-m2` PLAN.
2. **SCOPE_DIFF = all-in:** the five-file `internal/fieldspec/` fence + the registry singleton fully cover F-S7-R2-COLGRAIN; no scope-out row, no extra path (the mechanical block below).
3. **no hard trigger:** bounded single-domain fieldspec fix; no cross-domain authority change, no must-revise outstanding, no deviation.
4. **AC6 cleared:** master `SITREP-orchestrator-planner-20260710-133700.md` GRANTED the singleton `routing_assignments.gate_referenceable_columns = ["declared_deviated"]`; carries unchanged.

The mechanical scope block (all-in; five-file fence; no OUT row):

SCOPE_DIFF:
- internal/fieldspec/predicate.go -> in
- internal/fieldspec/registry.go -> in
- internal/fieldspec/registry.json -> in
- internal/fieldspec/predicate_test.go -> in
- internal/fieldspec/registry_test.go -> in

Build the plan of record — `master/relays/s7a-dispatch/PLAN-planner-20260710-140000.md` (mechanism, `GateReferenceableColumns` contract, AC1–AC7, red-first sequence, fence, branch). Execution defers to your side (Superpowers `executing-plans`); I stand by for inline questions/blockers and the IMPL report. The bare implementation token follows on its own line:

DISPATCH IMPL

**Carried build conditions (from the plan + master's clearance — do not trade):**
- **Atomicity (master's AC6 condition):** the `gate_referenceable_columns` singleton and the default-deny guard land such that **no commit on `s7a-colgrain` has default-deny active without the singleton present** — sequence the registry delta with or before the guard commit, never after. Red-first fixtures still precede the guard.
- **§9 registry-version:** the new `GateReferenceableColumns` member is additive — apply the MINOR registry-version treatment inside the five-file fence (the guard enforces the always-locked §5 `:88` with zero shipped-registry breakage per master's census). If any shipped predicate would break (none per census), stop and re-engage master.
- **Fence (AC5):** the five `internal/fieldspec/` files only; no engine/store/seat-surface; no s7-branch touch. Any needed path outside re-engages master.
- **AC1 reject shape:** a registry-load Go error with pinned substrings (`owner` + `non gate-referenceable row field` + `routing_assignments.chosen_model`), no path text — not a runtime bounce.
- **AC6/AC7:** the shipped registry delta is exactly the singleton; `registry_test.go` asserts the decoded allowlist byte-exactly `["declared_deviated"]`.

**Branch:** `s7a-colgrain` off BASE `frank/main@1d3e92c`.

**Return path:** your IMPL report (parenting to this dispatch, `s7a-plan-m2-dispatch`) → my pair review → master → m-4 fidelity (R2 routing face) + m-7 fidelity (registry-load/hosting surface) → VP integration of s7a → operator merge → the s7 pair rebases + folds the row-3 negatives.

ACTIONS_GIT_REF: wrote this dispatch relay (bare token to m-2.implementer, unique DISPATCH_ID, parenting to review-r1) + appended `master/relays/INDEX.md`; no `frank/` edit, no branch cut, no code — implementation authority passes to m-2.implementer.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; `frank/` main clean @ 1d3e92c, read-only).
Next requested action: m-2.implementer cuts `s7a-colgrain` off `1d3e92c` and executes the approved plan (red-first; atomicity held), returning the IMPL report parented to `s7a-plan-m2-dispatch` for m-2.planner pair review.
RELAY_LINT: OK — relay-lint.py exit 0 (file + relay-root lineage walk both clean)
