## PLAN-REVIEW — s16a WP1 PLAN r5 / narrow lineage-literal correction required

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-1
PARENT_DISPATCH_ID: s16a-build-5
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — s16a.planner can issue the bounded successor; the operator's next substantive gate remains MERGE-GATE
PLAN_LOCK_ID: s16a-build-5
IN_REPLY_TO: s16a-build-5/PLAN-planner-20260823-155140.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: formal PLAN-REVIEW — must-revise one stale authority-lineage literal in §5(a); every other r5 surface passes
PLAN_REVIEW_VERDICT: must-revise

`PLAN_REVIEW_VERDICT: must-revise` for exact tracked PLAN r5 SHA-256 `346f7fdf736a8fd8c3d2ac91611417eb4f1d8ad5b980620fc7d182c4d247b40e` at checkpoint `a25f4ababe7df4adc1fb58e8a1d1f1f7368ba378`.

The filed-parent gate is closed. `git ls-files` tracks the plan, the HEAD blob and working-tree bytes hash identically to `346f7fdf736a8fd8c3d2ac91611417eb4f1d8ad5b980620fc7d182c4d247b40e`, and committed `HEAD:frank/.relays/s16a/INDEX.md` contains the r5 plan row. The bank receipt `master/relays/t4-s16a-commission/SITREP-orchestrator-planner-20260823-193747.md` is consistent with those independent checks.

Blocking finding:
- **S16A-PR1-F1 — §5(a) names the superseded plan as the token precondition.** At line 63, r5 says the build token requires “your formal PLAN-REVIEW approve of THIS r4, parented to its dispatch.” That conflicts with the title and action record, which say r5 supersedes r4; with the header `DISPATCH_ID: s16a-build-5`; and with §9, which requires the verdict to parent to this relay's dispatch. Carried literally, the token precondition can be read as approving or parenting to the superseded r4 rather than the tracked r5 reviewed here. Exact authority lineage cannot rely on reading around that contradiction.

Required successor delta:
- Issue PLAN r6 as the bounded successor to r5.
- In §5(a), replace the stale “THIS r4” reference with “THIS r5” and make the parent target explicit as `s16a-build-5`.
- Preserve every other substantive r5 byte unless a separately identified correction is necessary. Re-bank the exact successor and its INDEX row before re-presenting it.

Closed and not reopened:
- Locked bases and hashes, including plan r7, A-1 with E-1, and the twice-corrected s16a charter.
- Exact WP1 fence: `frank/test/seam/**` plus lane relays only; production/source, dependencies, `master/**`, and all later work packages remain out.
- The 64-row battery: 20 GREEN pins including G20, 44 expected-RED gate rows, and D03 excluded.
- Compile-clean RED-by-assertion law, no skips, anti-vacuity stop rules, row-binding comments, one-test-per-row rule, exact pre-existing-test exclusion command, separate seam command, and script-derived census.
- PM guidance, m-10 item-3 withdrawal, boundary contract, E2 acceptance criteria, WP1-close fidelity hold, and no WP2 byte before that round.
- No additional design, scope, acceptance, boundary-contract, test-command, count, PM-guidance, or banking finding.

Remaining gates: no SCOPE_DIFF or implementation token on this verdict. After a tracked corrected successor receives `PLAN_REVIEW_VERDICT: approve`, the planner must mechanically produce `SCOPE_DIFF_RESULT: all-in` before any separately addressed implementation relay may carry the live token. Merge remains separately operator-gated.

No implementation action: no branch, worktree, source/test byte, dependency, stage, commit, token, merge, store, credential, or provider action was created or changed by this review.

Tests / verification: E1 exact HEAD/worktree hash equality, `git ls-files`, committed INDEX-row lookup, receipt comparison, full plan line review, and relay-lint. No source test was authorized.

Next requested action: `s16a.planner` returns a checkpoint-banked PLAN r6 closing only S16A-PR1-F1, then re-presents it for the next formal PLAN-REVIEW.

ACTIONS_GIT_REF: lane-governance disk action only — this PLAN-REVIEW plus one append-only `frank/.relays/s16a/INDEX.md` row; no source/test or git-state action claimed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build-5/PLAN-REVIEW-implementer-20260823-214313.md
?? frank/.relays/s16a/s16a-build-5/SITREP-planner-20260823-213200.md
