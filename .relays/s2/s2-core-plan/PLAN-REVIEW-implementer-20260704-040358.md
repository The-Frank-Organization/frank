## PLAN-REVIEW -- s2-slice-2-plan r4 narrow m-1 fold review

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s2-core-plan-review-implementer-r4
PARENT_DISPATCH_ID: s2-core-plan-lock-r4
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
SUBJECT: PLAN-REVIEW verdict -- must-revise; m-1 fold mechanics map through, but plan/design state text still points at stale or over-finalized m-1 state

Reviewed:
- Parent PLAN request: `.relays/s2/s2-core-plan/PLAN-planner-20260704-040121.md`.
- m-1 fidelity verdict: `.relays/s2/s2-fidelity-m1/SITREP-implementer-20260704-034158.md`.
- Prior r3 approval: `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-033128.md`.
- Root-lint disposition: `.relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-034500.md`.
- Folded docs at `main@da25aec`: `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md` and `docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md`.

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`

This is a narrow revision request. The three m-1 mechanics are present in the plan:

- F-M1-1 is mapped to Task 2 and related surfaces: genesis envelope/header homes, `system/system` provenance, and `seat.Mint("system")` reservation appear in `s2-slice-2-plan.md:66-69`; GC-marker provenance appears at `:132-135`.
- F-M1-2 is mapped to Task 3 and Task 11: `ErrChecksum` vs `ErrQuarantined`, post-eviction read behavior, and channel frames appear at `:74-77` and `:140-143`.
- F-M1-3 is mapped to Task 1, Task 2, and Task 11: store-root config members, Init-time source materialization, and serve-time store-root loading appear at `:57-69` and `:140-143`.

The blocker is that the documents now mix three different states: pre-fold design r2 (`main@6e3b67f`), folded design/plan r3 (`main@da25aec`), and m-1-final approval that is still explicitly pending. That makes the next m-1 packet and eventual dispatch parent ambiguous.

## Blocking findings

### F1 -- The plan still names the pre-fold design commit as the spec of record and m-1 review object

The r4 relay says this review object is "plan r3 + design r3 at main@da25aec" and that m-1 narrow re-review must confirm the folded lines before dispatch (`PLAN-planner-20260704-040121.md:20-31`). The current plan header says the plan cites the design doc at `main@6e3b67f` as the spec of record (`s2-slice-2-plan.md:7`), and Task 13 still describes the m-1 review object as design section 4 at `main@6e3b67f` (`s2-slice-2-plan.md:159`).

That is not just stale prose: `main@6e3b67f` is the design r2 commit before the m-1 fold, while `main@da25aec` is the commit that changed both design and plan for F-M1-1..3. A downstream m-1 narrow packet or delegated dispatch using the plan as written can cite the wrong design version.

Required revision:
- Update the plan's spec-of-record reference to the folded design/plan state at `main@da25aec`, or otherwise state precisely that r2 design review remains the lineage parent while r3 folded lines at `main@da25aec` are the current review surface.
- Update Task 13's m-1 gate to route the post-fold narrow m-1 re-review against design/plan r3, not the stale `main@6e3b67f` review object.

### F2 -- The design overstates m-1 approval before the required m-1 narrow re-review exists

The design status line correctly says r3 is "pending pair narrow re-review + m-1 narrow re-review" (`s2-slice-2-design.md:5`), and the r4 relay repeats that m-1 must confirm the folded lines before dispatch (`PLAN-planner-20260704-040121.md:31`). But design section 4 says the checked items are "m-1-APPROVED shapes (post-fold), no longer open proposals" (`s2-slice-2-design.md:126`), and the r3 fold-log says several items are "marked m-1-approved-post-fold" (`:200`).

Those statements outrun the current gate state. The m-1 verdict was `must-revise` and required a narrow re-review confirming the fold before implementation dispatch. Until that follow-up exists on record, the folded shapes can be described as m-1-prescribed or m-1-required post-fold shapes, but not as finally m-1-approved.

Required revision:
- Normalize the design and plan wording so it says the r3 shapes are m-1-prescribed/required and pending m-1 narrow re-review, or produce the m-1 approval first and cite it.
- Avoid mixed labels such as "APPROVED" in section 4 while other nearby lines still say m-1 approval is pending.

## Non-blocking checks

- The r4 relay exact-file lints clean: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/PLAN-planner-20260704-040121.md` -> OK.
- The old superseded r2 root-mode lint error is dispositioned by `.relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-034500.md`; I am not treating that residue as this review's blocker.
- The m-1 fold commit is doc-only: `git show --stat --oneline --decorate HEAD` reports `da25aec` touching only the S2 design and plan docs.
- `go test -count=1 ./...` passes across all packages.

## Verification

- `git show --stat --oneline --decorate HEAD` -> `da25aec (HEAD -> main) s2 DESIGN+PLAN: r3 folds ...`; two doc files changed.
- `git diff --stat c16f261..da25aec -- docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md` -> two files, 35 insertions, 25 deletions.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/PLAN-planner-20260704-040121.md` -> `OK .relays/s2/s2-core-plan/PLAN-planner-20260704-040121.md`.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-034500.md` -> `OK .relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-034500.md`.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s2 .relays/s2/s2-core-plan/PLAN-planner-20260704-040121.md` -> known INDEX noise plus old superseded r2 error; target r4 PLAN itself OK. The old r2 error is covered by the scoped operator waiver in `SITREP-orchestrator-planner-20260704-034500.md`.
- `go test -count=1 ./...` -> pass across all packages (`cmd/frank` and `test/seatproc` have no test files; all test packages OK).
- `git status --short` before relay write -> no output.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-040358.md` -> `OK .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-040358.md`.

ACTIONS_GIT_REF: no source/test/sprint-spec edits; reviewed tracked fold commit `main@da25aec`; wrote gitignored relay `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-040358.md` plus `.relays/s2/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: none - clean tree
