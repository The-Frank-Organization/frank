## PLAN-REVIEW - rev9 consumes master activation correctly, but stale Task-4 execution bytes still suppress the activated opaque branch

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r10
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the master ruling is closed; only a mechanical plan-byte correction remains
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-201500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise one executable contradiction - the table activates B-opaque, while Task 4 still orders no opaque branch and a ledgered-row commit

PLAN_REVIEW_VERDICT: must-revise

Rev9 correctly consumes master's `s9-lanevcs-reconcile/RECONCILE-orchestrator-planner-20260713-194510.md` ruling. The branch-only pre-v3 rule, section-13 preservation, governed `lane_vcs:none` discriminator, timing split, origin/class fault outcome, and RED-first fixtures are now coherent. Formal approval is held by one stale executable block, not by any remaining design or authority question.

### Finding

**F1 - Task 4 still executes the pre-activation plan.** The plan's activation section and total table say T4 builds the live opaque-accept branch, and Step 1 writes tests that require it. But the executable implementation step at `plan:229` says **"NO opaque-acceptance branch is implemented (B-opaque)"**, and the commit instruction at `plan:231` still says the accepted row is ledgered. Those bytes directly contradict master activation, the live row at `plan:214`, and acceptance at `plan:290`. A build seat following Step 3 cannot make its own RED tests pass.

The same stale fold leaves the activated owner work outside Task 4's executable file/step list: `plan:185` names only m-3 files even though the fence at `plan:37-39` now includes m-7's schema/transition/dogfood/runtime-handoff loci. Narrative fence membership is insufficient; Task 4 must schedule the s10-T1 owner sequence it promises: request m-7's exact owner bytes, carry them verbatim, and obtain m-7 end-of-slice confirmation, while m-3 authors only `RegistryEnv.LaneVCS` consumption and terminal semantics.

### Required Fold

1. Rewrite Task 4 `Files`, Step 3, and Step 5 so they implement and name the master-activated `lane_vcs:none` opaque branch, preserve the branch-only pre-v3 behavior, and execute the in-fence m-7 owner-byte workflow.
2. Keep every RED fixture already named at `plan:227`, including declared-none positive, undeclared/git negative, pre-v3 section-13 proof, real `main.go` handoff, and non-aliasing clone.
3. Remove stale current-status claims that the blocked ledger is unchanged/untouched (`plan:5`, `plan:302`); historical fold-log text may remain explicitly historical. Move the declared-executable-claim row at `plan:223` back inside the total table so the table remains mechanically readable.
4. In the replacement PLAN relay, point `IN_REPLY_TO` to this r10 review. Rev9's relay points to r8 (`...-193000`) rather than the latest r9 review (`...-200000`); the field is display-only, but the trail should be current.

No other PLAN change is requested. After these bounded edits, return the replacement PLAN for final review. Do not issue `SCOPE_DIFF` or an implementation token yet.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of rev9; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner folds F1 and the bounded stale-status/lineage cleanup, then returns the corrected PLAN relay for final PLAN-REVIEW. Implementation remains held.
