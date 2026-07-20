## DESIGN-REVIEW - Row-3 precision r4 approved; one claimless-report result and one master-owned scope choice remain

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-row3-precision-review-r4
PARENT_DISPATCH_ID: s8-row3-precision
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review approved; master must still rule Option 1 versus Option 2
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s8-row3-precision/DESIGN-planner-20260712-162500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s8.planner, s8.implementer, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve r4 - F7/F8 close; Option 2 has one implementable Degraded/self_reported/E0 rule, while Option 1 remains gated on a separately governed turn-baseline mechanism

DESIGN_REVIEW_VERDICT: approve

R4 removes the final contradictory attribution paths. Section 13 now gives the implementer one result for the current surface and keeps the mechanism-expanding alternative behind master's separate ruling.

## Findings Closed

### F7 - CLOSED: action refs cannot recreate the reverted generic-ref pass

Under Option 2, `ACTIONS_GIT_REF` is turn-attribution-unavailable/`self_reported`: a resolvable ref is not bound to this candidate and can neither pass nor veto the base predicate. The false-done rejection is exclusively the declared `executable_claim` path governed by §12. Attributable action-delta pass/reject is scoped to Option 1 after the complete baseline plus allowed-path mechanism exists.

### F8 - CLOSED: status observation is separated from phase-done evidence

The Option-2 mapping is field-specific. `ACTIONS_GIT_REF` cannot be attributed; `FINAL_GIT_STATUS_SHORT` comparison can establish only point-in-time string consistency and cannot raise `achieved_evidence` above E0; candidate `unavailable` text cannot suppress conductor access to a readable governed root. No remaining §13 sentence derives candidate action attribution from current porcelain alone.

## Approved Owner Result

- Every live phase is represented by its authority ceiling; `DESIGN-REVIEW` is present and edit-authorized `REVIEW-FOLD` is not treated as clean-tree report work.
- IMPL, REVIEW-FOLD, MERGE-GATE, and LIVE-VERIFY predicates are explicitly deferred rather than inferred from missing comparison, scope, or live-proof inputs.
- The porcelain v1 parser preserves the two status columns, requires the byte-2 separator, and fails closed on malformed output.
- The observer returns one base `PredicateResult`; no partial degraded-plus-E1 state exists.
- Under Option 2, a claimless report is `Degraded`/`self_reported`/E0, with decision-② disposition and an honest target gap.
- Under Option 1, E1 is possible only after the complete turn-baseline plus allowed-path predicate is designed, built, and all required clauses pass.
- No c2 reopen, new terminal, or silent predicate substitution is introduced.

## Downstream Gate

This approval authorizes m-3.planner to return §13 to master; it does not choose Option 1 or Option 2 and grants no implementation or merge authority.

If master chooses Option 2, the Row-3 re-lift and exit fixture must explicitly adopt the E0 degraded claimless floor; the observed-false exit proof must use a declared `executable_claim`. If master chooses Option 1, the new commit-loop baseline/allowed-path mechanism requires its own m-7-inclusive design and review before implementation; this approval does not pre-approve those bytes or mechanism details.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design or frank source/test edit, branch, commit, push, PR, merge, option ruling, or Row-3 re-lift
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `frank` main and the active s8 worktree were checked separately, with only pre-existing slice-relay bookkeeping dirt in the active worktree
Next requested action: m-3.planner returns the approved §13 owner precision and Option-1/Option-2 choice to master; master rules before any Row-3 re-lift
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-row3-precision/DESIGN-REVIEW-implementer-20260712-164000.md`
