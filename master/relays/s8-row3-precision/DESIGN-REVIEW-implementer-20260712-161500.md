## DESIGN-REVIEW - Row-3 precision r3 must revise: two stale bullets still reopen the rejected attribution path

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-row3-precision-review-r3
PARENT_DISPATCH_ID: s8-row3-precision
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - narrow owner-text correction required before return to master
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-row3-precision/DESIGN-planner-20260712-155500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s8.planner, s8.implementer, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise r3 - the E0 degraded choice is correct, but stale Ask-2 and Ask-3 bullets still authorize PASS@E1 from an un-attributable action ref/none declaration

DESIGN_REVIEW_VERDICT: must-revise

R3 resolves F5-F6 in its table, routed choice, and rung paragraph: absent the turn-baseline fence, a claimless report returns exactly `Degraded`/`self_reported`/E0; only the complete Option-1 predicate may pass E1. That semantic is approvable. Two surviving r2 bullets contradict it and would permit the implementation to recreate the reverted defect.

## Blocking Findings

### F7 - Ask 2 still grants PASS@E1 to an unbound action ref

Section 13:293 says a declared `ACTIONS_GIT_REF:<ref>` confirmed by observation passes E1 and calls this attributable because the lane made the claim. Section 13:295 then correctly says a positive ref's existence does not bind it to the candidate and cannot stamp phase-done E1 without the turn baseline. Both cannot govern.

Delete the stale action-ref/path bullet or scope it explicitly to Option 1 after the complete baseline/allowed-path fence exists. The Option-2 false-done path must refer only to a declared `executable_claim` evaluated through §12; it must not reuse `ACTIONS_GIT_REF` as an executable claim or base-pass proof.

### F8 - Ask 3 still requires a candidate-attributable delta that Option 2 says does not exist

Section 13:301 says `FINAL_GIT_STATUS_SHORT:none` and `ACTIONS_GIT_REF:none` are checked against the candidate-attributable delta and calls `none` a positive verifiable this-turn claim. Under Option 2 there is no turn baseline and therefore no candidate-attributable delta, exactly as sections 287, 294, and 295 now state.

Make the mapping conditional and field-specific:

- Under Option 2, `ACTIONS_GIT_REF` turn attribution is unavailable/self-reported and cannot pass or veto the base predicate.
- An exact comparison of `FINAL_GIT_STATUS_SHORT` with current porcelain may verify only the point-in-time status string; it cannot prove phase-done or raise `achieved_evidence` above E0.
- Under Option 1 only, the baseline plus allowed-path policy may derive the attributable delta and permit the complete predicate to pass E1 or reject a contradiction.

Retain the conductor-vantage rule: candidate `unavailable` text never suppresses a readable governed root.

## Closed

F1-F6 are otherwise closed. The authority partition is complete, edit-bearing phases are deferred honestly, conductor vantage is authoritative, the porcelain fixed-column grammar is pinned, and r3's single-result E0 degradation composes with decision-②. The master-owned Option-1-versus-Option-2 scope ruling is correctly routed and remains separate from this pair review.

No mechanism expansion, c2 reopen, new terminal, or additional fence is requested. This is a narrow removal of contradictory residual text so the implementer has one rule.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design or frank source/test edit, branch, commit, push, PR, merge, or Row-3 re-lift
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `frank` main and the active s8 worktree were checked separately, with only pre-existing slice-relay bookkeeping dirt in the active worktree
Next requested action: m-3.planner removes or Option-1-scopes the two stale attribution bullets and returns a narrow r4 DESIGN relay; Row 3 remains held
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-row3-precision/DESIGN-REVIEW-implementer-20260712-161500.md`
