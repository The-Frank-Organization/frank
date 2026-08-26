## RECONCILE — OPERATOR DIRECTIVE carried down (effective immediately): the s16a branch goes up as a DRAFT PULL REQUEST on origin; the no-merge fence is unchanged

*(Successor to `s16a-pr/RECONCILE-orchestrator-planner-20260825-072029.md`, post-submit lint: non-canonical CEREMONY_TIER; content otherwise identical. This relay is the directive of record.)*

**The directive (operator, 2026-08-25; recorded as PROTOCOL-DEVIATIONS V30 + charter CORRIGENDUM 3):** slice teams make their work available as pull requests for ease of review. For s16a, the implementer's act now:

1. **Push the branch:** `git push -u origin s16a-conformance` from the implementation worktree (origin = `The-Frank-Organization/frank-dev`). The prior "no push" fence is AMENDED to "no MERGE" — branch pushes are now directed; after each banked WP commit, push again so the PR tracks the lane.
2. **Open a DRAFT PR** against `main`: title `s16a: seam conformance battery + fixes (T4)`; description = the slice scope (charter + plan r7 pointers), the current census (state it as of HEAD at open), the WP ladder with status, and a line stating the governance record is the relay trail under `frank/.relays/s16a/` — the PR is the REVIEW SURFACE, not an authority channel. Refresh the description at each WP close.
3. **The merge fence is UNCHANGED:** nothing merges without the operator's WP5 MERGE-GATE grant (executor in TO, `HUMAN_MERGE_AUTHORIZATION` verbatim); the grant now EXECUTES as the PR merge. Do not mark the PR ready-for-review until the WP5 close set is filed.
4. Report the PR URL + the pushed tip sha in your next SITREP; master verifies both at its own invocation.

For the record and not the pair's act: master now also syncs the standalone `frank` repo (`The-Frank-Organization/frank`) from workspace `main` after every frank-touching merge (PROTOCOL-DEVIATIONS V31; catch-up `975bd925` → `45869f93` landed this act). All other standing state — the r14 token, the B10 pins from `s16a-wp2/RECONCILE-orchestrator-planner-20260825-064318.md`, B10-lands-last, the WP ladder — is untouched by this directive.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s16a-pr-directive
PARENT_DISPATCH_ID: t4-s16a-commission
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — carrying an operator directive down verbatim; the operator's next gate remains the terminal WP5 MERGE-GATE (now executed as the PR merge)
IN_REPLY_TO: s16a-wp2/SITREP-planner-20260825-062725.md
FROM: master.orchestrator-planner
TO: s16a.planner, s16a.implementer
CC: master.orchestrator-reviewer, operator
SUBJECT: operator directive — push s16a-conformance to origin + open a DRAFT PR to main (review surface; relay trail stays the record); no-merge fence unchanged, MERGE-GATE executes as the PR merge; report PR URL + tip sha next SITREP

ACTIONS_GIT_REF: engine-lane governance act — this directive drafted at .engine/drafts/master.orchestrator-planner/ and submitted through relay submit; this act also appended PROTOCOL-DEVIATIONS V30/V31 + charter CORRIGENDUM 3 and pushed the standalone frank sync commit 45869f93 (git commit-tree snapshot, fast-forward 975bd925 -> 45869f93, .engine/key absence verified in the synced tree); no slice source, test, branch, or commit byte from this seat.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M master/relays/CHECKPOINTS.md
