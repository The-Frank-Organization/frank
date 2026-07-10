## PLAN-REVIEW -- s2-slice-2-plan r5 narrow state-label review

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s2-core-plan-review-implementer-r5
PARENT_DISPATCH_ID: s2-core-plan-lock-r5
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
PLAN_REVIEW_VERDICT: must-revise
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: PLAN-REVIEW verdict -- must-revise; stale spec-of-record pointer is closed, but two plan task labels still overstate m-1 approval state

Reviewed:
- Parent PLAN request: `.relays/s2/s2-core-plan/PLAN-planner-20260704-041121.md`.
- Prior narrow review: `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-040358.md`.
- Root-lint disposition: `.relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-034500.md`.
- Folded docs at `main@845a7d1`: `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md` and `docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md`.

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`

This is a narrow residual revision. F1 from my prior review is closed: the plan now says the current in-tree design is the spec of record, distinguishes r3 `main@da25aec` substantive folds from r4 state-label normalization, and preserves `main@6e3b67f` only as the r2 lineage parent (`s2-slice-2-plan.md:7`). Task 13.2 now routes the remaining m-1 narrow re-review to the post-fold review object, with m-1 approve required before dispatch (`s2-slice-2-plan.md:159`).

F2 is mostly closed in the design, but not closed in the plan task text. The prior required revision asked to normalize design and plan wording so folded shapes are m-1-prescribed/required and pending m-1 narrow re-review, not m-1-approved. Two plan task bodies still use the over-finalized labels:

- `docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md:67` says `the exact m-1-fixed shape (r3 F-M1-1)`.
- `docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md:133` says `section 4.6 m-1-approved shape`.

Those labels conflict with the now-correct gate state elsewhere: plan line 13 still says m-1 shapes are proposals pending fidelity, line 159 requires m-1 approve before dispatch, and design line 126 says the section-4 items become m-1-APPROVED only when m-1's narrow re-review confirm is on record.

Required revision:
- Replace the two remaining task-body labels with the same state used in the design, for example `m-1-PRESCRIBED` / `m-1-prescribed-pending-confirm`.
- Keep the existing line 7 lineage distinction and Task 13.2 m-1 re-review routing; those parts now satisfy F1.

## Non-blocking checks

- The incoming r5 relay exact-file lints clean.
- The old superseded r2 root-mode lint error is dispositioned by `.relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-034500.md`; I am not treating that residue as this review's blocker.
- The r4 label-fix commit is doc-only: `git show --name-only --format=oneline 845a7d1 --` lists only the S2 design and plan docs.
- No `DISPATCH IMPL` token is present in the parent relay; this remains review-only.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/PLAN-planner-20260704-041121.md` -> `OK .relays/s2/s2-core-plan/PLAN-planner-20260704-041121.md`.
- `rg -n "m-1-APPROVED|m-1-approved|m-1-FIXED|m-1-fixed|main@6e3b67f|main@da25aec|main@845a7d1|spec of record|review object|NARROW|PRESCRIBED|pending" docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md` -> residual `m-1-fixed` at plan line 67 and residual `m-1-approved` at plan line 133; lineage/current-review wording present at plan lines 7 and 159.
- `go test -count=1 ./...` -> pass across all packages (`cmd/frank` and `test/seatproc` report no test files).
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-041446.md` -> `OK .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-041446.md`.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s2 .relays/s2/s2-core-plan/PLAN-planner-20260704-041121.md .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-041446.md` -> known `INDEX.md` header noise plus old superseded r2 lineage error; both r5 target files OK. The old r2 error is covered by `.relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-034500.md`.

ACTIONS_GIT_REF: no source/test/sprint-spec edits; reviewed tracked fold commit `main@845a7d1`; wrote gitignored relay `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-041446.md` plus `.relays/s2/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: none - clean tree
