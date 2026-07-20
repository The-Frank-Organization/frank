## SITREP — the s7 endgame order to the s7 pair: bring the merged guard into your branch (merge `main@54420dbc` INTO `s7-inv-catalog` — no rebase, the trail's SHAs stay valid), fold the row-3 `any_row` negatives into the named invariant, pair-review, report; the operator countersigned the s7a trail record, so nothing upstream blocks you

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator countersigned the s7a corrective merge record ("ratified", 2026-07-10, against `s7a-merge-gate/MERGE-GATE-…-160704`), closing the executor blocker; this fold runs under your standing s7 dispatch (REVIEW-FOLD loop); the s7 merge stays operator-gated at the slice's end
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/SITREP-orchestrator-planner-20260710-113841.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-2.implementer, m-4.implementer, m-2.planner, m-4.planner
SUBJECT: s7 row-3 unblocked — F-S7-R2-COLGRAIN is closed in `main@54420dbc` (the s7a default-deny column guard + the truthfully-attributed `s7a-fieldspec-v5`); execute the held row-3 fold under your standing authority: integrate main, add the `any_row:routing_assignments.chosen_model` negatives to `TestLawR2NoModelPredicate`, pair-review, report

**Context:** the guard your row 3 was held on now exists in `main` — merged by the operator's named executor, VP-final-approved, master-verified (`54420dbc`, parents `1d3e92c`+`2bc0763`, serialized battery green). The operator's countersign closed the s7a trail incident; the two s8 obligations (the genesis condition; `OI-S7A-CLOSE-ONCE-RACE`) stand and do not gate this fold.

**The bounded fold (under your standing s7 dispatch; test-only fence unchanged):**
1. **Integrate main into your branch:** in the s7 worktree, `git merge main` (bringing `54420dbc` into `s7-inv-catalog` at `81dce49`) — **merge, not rebase**: every SHA in the slice trail (`eaaf5f0`/`bc88fe8`/`35aabb9`/`81dce49`) stays valid. Expect no conflicts (your diff is `test/invariants/` + slice docs; s7a's is `internal/fieldspec/`).
2. **Fold row 3:** extend `TestLawR2NoModelPredicate` with the two synthetic negatives the m-2/m-4 fidelity verdicts required — `required_when` AND `visible_when` predicates `any_row:routing_assignments.chosen_model` — asserting **registry-load rejection** (they run green against the merged guard; the mechanism-level red was already proven in s7a's red commit `10ee3a2` on the pre-guard parser, so no scratch-red is owed here). Keep the catalog row-3 claim text within the bounded wording m-4 pinned ("R2 enforced at column grain" = the live `any_row` grammar + the shipped registry — no broader claim). If the fold needs any touch beyond `test/invariants/`, stop and report.
3. **Verify:** focused `go test -count=1 ./test/invariants` (all ten rows + the new negatives) · the full uncached battery **serialized** (`-p=1`, per the standing honest-sequence rule while `OI-S7A-CLOSE-ONCE-RACE` is open) · `go vet` · the diff audits (fold commit test-only; branch-to-main diff = `test/invariants/` + slice docs only).
4. **Pair review** (m-7.planner, diff-grain, per your R-1/row-6 pattern) → your report TO master with the trail. I then route **row 3 to m-2 + m-4 for the final re-confirms**, assemble the full s7 integration package for the VP, and the slice goes to the operator's merge gate.

Next requested action — m-7.implementer: the three-step fold + pair review + report. Nothing else is authorized (no s7 merge, no tag, no catalog governance change — the §7-pinning carry stays s8's).

ACTIONS_GIT_REF: none — no git action by this relay (the fold order; the merge-into-branch and fold commits are the executor's under the standing dispatch).
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` clean at `54420db`; the s7 worktree at `81dce49` awaiting the integrate+fold; cwd is not a git repo (docs workspace).
