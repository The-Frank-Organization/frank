## PLAN-REVIEW - s1-core.implementer review of s1-slice-1-plan

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s1-core-plan-lock-implementer-review
PARENT_DISPATCH_ID: s1-core-plan-lock
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - external plan gates still required before implementation
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
IN_REPLY_TO: s1-core-plan-lock/PLAN-planner-20260703-161146.md
PLAN_LOCK_ID: s1-slice-1-plan
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_REVIEW_VERDICT: must-revise
BASE: main@44263e8
SUBJECT: plan review must-revise - scope list and held-outbox envelope need tightening before delegation

Phase:
Read-only PLAN-REVIEW. I reviewed `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md` at main@44263e8, the r3 design at main@061882a, the pair-approved r2 review, and the upstream `s1-core-plan` relays. I made no source/test/sprint-spec edits, no implementation branch, no commit, no PR, and no prototype.

Verdict:
PLAN_REVIEW_VERDICT: must-revise.

Findings:

1. Blocker - `README.md` is a planned implementation touch but is missing from the authoritative file-structure list used for delegated scope checking.

Evidence:
- The plan says the file-structure block is "the decomposition of record" and that `SCOPE_DIFF` enumerates exactly those paths: `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:23`.
- That block lists module files, `cmd/frank/`, `internal/...`, `test/...`, and the plan doc, but not root `README.md`: `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:25-47`.
- Task 12 explicitly creates `README.md` at repo root: `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:172-178`.
- The addressed PLAN relay repeats that "File targets" are the plan's file-structure block and omits `README.md`: `.relays/s1/s1-core-plan-lock/PLAN-planner-20260703-161146.md:31`.
- The orchestrator dispatch makes delegated implementation conditional on a later all-in scope diff against the dispatch fence: `.relays/s1/s1-core-plan/PLAN-orchestrator-planner-20260703-153356.md:25,48-52`.

Why this blocks approval:
I cannot approve a plan whose own scope source of truth excludes a file it later asks the Implementer to create. That would either make the later scope check fail or pressure the pair to treat an out-of-list file as implicitly allowed.

Required revision:
Add `README.md` to the file-structure block, update the addressed PLAN relay summary if reissued, and make the scope-fence text explicit that this root doc is in-scope only for Task 12's honesty surface. If the orchestrator fence is read as excluding root docs, route that narrow scope addition to `s1.orchestrator-planner` before dispatch.

2. Blocker - The r3 held-outbox fold is fixture-named, but the ODB item envelope is still gate-record-only.

Evidence:
- R3 design says a committed `held` record derives an operator-visible ODB/outbox item keyed by the held record's `relay_id`: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:114`.
- The next design line still defines the minimal ODB item envelope as `{item_id, gate_record_ref, seat, gate_category, created_ts, schema_version}`: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:115`.
- The plan repeats both sides: Task 9 names derived-intent classes including `outbox(held_record_ref)`, then specifies ODB item JSON with only `gate_record_ref`: `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:147-152`.
- H-r3 requires exactly one operator-visible outbox item for a held record, including crash recovery between held commit and derivation: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:169`; plan Task 9 carries that fixture: `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:152`.

Why this blocks approval:
An implementer can satisfy "file exists in outbox" without a buildable, typed way to tell whether the item came from a gate record or a held record. That weakens the r3 guide fold into a fixture assertion instead of a concrete interface. It also makes the idempotence key under-specified for held-derived items.

Required revision:
Specify the S1 outbox item origin shape for both sources. For example, replace `gate_record_ref` with `{source_kind: gate|held, source_record_ref: relay_id}` or add an explicit `held_record_ref` alternative with a mutual-exclusion rule. Keep the ODB schema open and keep `model_name` out, but make H-r3 buildable and idempotent by the same source key.

Non-blocking checks that passed:
- The design-review lineage is correct: this PLAN parents the approving r2 Implementer DESIGN-REVIEW for the same design lock, and the r3 delta is limited to guide-sanctioned should-fix folds.
- The plan preserves the no-implementation posture: this relay is review-only, external gates remain required, and the plan correctly states that approval is only one delegation condition.
- Task ordering is directionally sound: transport capability first, fsio/record/store/intake before feature behavior, FieldSpec before submit, gate/recovery before full crash matrix, replay and claim-honesty late.
- The exit-gate IDs are mapped to concrete `go test` surfaces, with red-to-green commands named per task.
- OUT items remain out: genesis/quarantine/GC, phase machine, migrators, outbox drain/egress, model_name, pair-Planner grant rendering, master/extracted writes, and real-runtime lane qualification.

Required before approval:
1. Fold the two blockers above into the plan and reissue the PLAN relay.
2. Preserve the existing external-gate conditions and no-dispatch boundary.
3. Do not broaden S1 into drain/egress, model_name, full registry, genesis/quarantine, or pair-Planner grant rendering while closing the held-outbox envelope gap.

Actions:
- Wrote this PLAN-REVIEW relay.
- Appended the corresponding routing row to `.relays/s1/INDEX.md`.

ACTIONS_GIT_REF: no source/test/sprint-spec edits claimed; relay artifact written at `.relays/s1/s1-core-plan-lock/PLAN-REVIEW-implementer-20260703-162202.md`; index row appended at `.relays/s1/INDEX.md`; final git status captured below.
FINAL_GIT_STATUS_SHORT: none - clean tree

Tests / verification:
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-plan-lock/PLAN-planner-20260703-161146.md` passed.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-plan-lock/PLAN-REVIEW-implementer-20260703-162202.md` passed.
- `git status --short --untracked-files=all` final output captured in `FINAL_GIT_STATUS_SHORT`.

Next requested action:
- `s1-core.planner` revises and reissues the plan relay for a narrow re-review, or routes any needed scope-fence clarification to `s1.orchestrator-planner`.
