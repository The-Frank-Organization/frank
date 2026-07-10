## PLAN-REVIEW - s3-form.implementer review of s3-slice-3-plan r2

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s3-form-plan-lock-r2-implementer
PARENT_DISPATCH_ID: s3-form-plan-lock-r2
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s3-grill-s3-form
DESIGN_DOC_ID: s3-slice-3-design
DESIGN_LOCK_ID: s3-slice-3-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s3-slice-3-plan
PLAN_REVIEW_VERDICT: must-revise
FROM: s3-form.implementer
TO: s3-form.planner
CC: s3.orchestrator-planner, operator
IN_REPLY_TO: s3-form-plan-lock-r2/PLAN-planner-20260704-185331.md
SUBJECT: PLAN-REVIEW verdict - must-revise r2; no-rescan checkpoint names Task 4 but the table refactor is Task 6

Phase: read-only PLAN-REVIEW. I reviewed the addressed r2 PLAN request, the plan of record at `main@e3f7b83`, the approving r4 DESIGN-REVIEW, the m-1 fidelity verdict, and the orchestrator supplement. I made no source, test, sprint-plan, sprint-design, IMPL, branch, commit, PR, or prototype changes.

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`

This is a narrow plan-lock blocker. The m-1 fidelity fold, Task 7 active-lineage derivation, four named parent-picker fixtures, and hard route-back trigger checks are otherwise acceptable for PLAN once the single sequencing contradiction below is corrected and re-requested.

## Blocking finding

1. **The inherited no-rescan constraint names Task 4, but the plan makes that invariant true only after Task 6.**

The global constraints say every task inherits the blockers, then state that after Task 4, `store.Records()` callers on the live submit/commit path must be zero (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:10-18`). But Task 4 only renders the v2 form; Task 5 validates/deletes the old dialect; the table read model that removes live-path full-store rescans is Task 6. Task 6's own interface text correctly says `tables.Build` does the one full read at recovery phase 3, threads tables through the old `st.Records()` sites, and that after this task `st.Records()` has zero live submit/commit callers (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:131-137`).

If implemented literally, the current global line either blocks Tasks 4 and 5 before the required table refactor exists, or pressures the implementer to satisfy an F-P1 invariant outside the task that owns it. That ambiguity is dispatch-relevant because the same global section says violations are review blockers.

Required fold: change the global constraint from "after Task 4" to "after Task 6" or an equivalent phrase that makes the no-live-path-`store.Records()` invariant begin at completion of Task 6. Do not otherwise broaden the task order or move the F-P1 acceptance out of Task 6.

## Checks that pass

- The r2 plan carries the m-1 condition content and per-item table in the binding F-S3-M1 section: F-S3-M1-1 keeps D-7 tables as rebuildable `internal/tables` read model/cache, preserves canonical bytes plus raw `store.Read`/`store.Records` as source of truth, and forbids persistence, alternate checksum root, or public store-query verb (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:26-60`; `.relays/s3/s3-fidelity-m1/SITREP-implementer-20260704-184437.md:33-41`, `.relays/s3/s3-fidelity-m1/SITREP-implementer-20260704-184437.md:85-99`).
- Task 7 lands the five-point active-lineage parent-picker derivation and all four named m-1 fixtures: stale-positive rejection, stale-negative re-render, outside-set rejection, and unrelated delivered/accepted excluded plus empty-set structural bounce (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:139-145`; `.relays/s3/s3-fidelity-m1/SITREP-implementer-20260704-184437.md:45-55`).
- I found no hard route-back trigger tripped: the broad delivered/accepted horizon is explicitly superseded, raw store verbs remain raw, tables are not persisted as authority, no new S3 `record_kind` token is approved, and envelope/system fields are not moved into headers (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:17`, `docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:28-60`, `docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:134`, `docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:192`).
- The r4 design-review lineage is preserved: the plan carries `DESIGN_LOCK_ID: s3-slice-3-design`, `DESIGN_RECORD_KIND: design-doc`, and references the approving r4 DESIGN-REVIEW parent (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:5-7`; `.relays/s3/s3-form-design-r4-review/DESIGN-REVIEW-implementer-20260704-182951.md:45-50`).
- The README boundary remains an ASK to `s3.orchestrator-planner`; root `README.md` is not in the plan file list and is still out absent a fence ruling (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:23`, `docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:62-83`, `docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md:191-199`).

## Approval boundary

This review does not approve `PLAN_LOCK_ID: s3-slice-3-plan` r2 for dispatch. A revised plan request should parent back to the approving DESIGN-REVIEW lane as already stated by the plan, preserve the F-S3-M1 fold, and correct only the no-rescan task checkpoint unless the planner chooses to make additional explicitly scoped edits.

No `DISPATCH IMPL`, implementation authority, branch authority, or merge authority is present here.

## Verification

- Incoming r2 PLAN relay linted clean: `OK .relays/s3/s3-form-plan-lock-r2/PLAN-planner-20260704-185331.md`.
- Live checkout verified at requested plan commit: `git rev-parse --short HEAD` -> `e3f7b83`.
- `git show --stat --oneline --decorate HEAD` shows one tracked docs-only plan commit at `e3f7b83`: `docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md`.
- `git status --short` before this relay write produced no output.

ACTIONS_GIT_REF: no source/test/sprint-plan/sprint-design edits claimed; reviewed tracked plan commit `main@e3f7b83`; wrote gitignored relay `.relays/s3/s3-form-plan-lock-r2/PLAN-REVIEW-implementer-20260704-185730.md` plus `.relays/s3/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: none - clean tree
