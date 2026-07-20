## DESIGN-REVIEW - Row-3 precision must revise: phase partition and observable predicates are not yet closed

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-row3-precision-review
PARENT_DISPATCH_ID: s8-row3-precision
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - owner design revision is required before master may re-lift Row 3
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-row3-precision/DESIGN-planner-20260712-143000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s8.planner, s8.implementer, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - the table omits DESIGN-REVIEW, misclassifies edit-authorized REVIEW-FOLD, conflates author cwd with the pinned lane, and leaves IMPL matching-diff without a comparison base

DESIGN_REVIEW_VERDICT: must-revise

The Row-3 re-lift remains held. Section 13 correctly preserves porcelain's leading status column, distinguishes observable contradiction from no-vantage, labels the `scope_paths` and MERGE/LIVE deferrals, and binds a passing base observation to E1. Four blocking defects remain in the claimed closed table.

## Blocking Findings

### F1 - The phase partition is neither complete nor authority-correct

The table claims every live phase but omits `DESIGN-REVIEW`, which is a live protocol/FieldSpec phase. It also places `REVIEW-FOLD` in the zero-source-edit report group, although `REVIEW-FOLD` with `AUTHORITY: fold-in-only` explicitly authorizes scoped edits after a pre-action `FOLD_SCOPE`. A conforming fold with an expected diff would therefore be rejected by the proposed clean-tree predicate.

Revise the table to enumerate every live phase and use the already-supplied `Candidate.Authority` wherever phase alone does not determine the action ceiling. At minimum, `DESIGN-REVIEW` needs an explicit row/group and `REVIEW-FOLD` needs an edit-aware predicate (or an explicit s8 deferral); it cannot remain in Group A.

### F2 - Candidate cwd vocabulary cannot determine conductor vantage

Section 13 says the conductor observes the pinned governed-lane root, but the byte mapping treats `FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo` as sufficient to enter no-vantage degradation. Those are different facts. The docs-authoring cwd can be non-git while the pinned `frank` lane is fully observable; conversely, a candidate cannot make an observable lane opaque by writing `unavailable`.

Make conductor vantage authoritative: no-vantage is entered only when the conductor cannot resolve/read the governed lane root. When that root is observable, the candidate's structured absence is compared with the actual observation and cannot suppress it. State which target the two action-report fields describe when author cwd and governed lane differ.

### F3 - Dirty status does not prove this candidate performed an unauthorized action on the stated surface

The table strikes the locked allowed-artifact/expected-tree refinement because the surface lacks it, then equates "no unauthorized source action" with an entirely clean tree. That is not an operationalization of the same predicate. A point-in-time dirty status cannot attribute pre-existing dirt or permitted relay bookkeeping to this candidate, and `ACTIONS_GIT_REF: none` reports this turn's action, not a claim that all lane bytes are clean. The active s8 worktree itself demonstrates the distinction: source is clean while tracked `.relays` bookkeeping is dirty.

Either provide the smallest sanctioned baseline/allowed-path input needed to classify the observed delta, or label the unavailable attribution as degradation/deferral. Do not call dirty-vs-`none` an observed contradiction unless the conductor has enough pinned facts to establish that the dirty bytes belong to the candidate and are unauthorized.

### F4 - The IMPL predicate still contains an undefined implementation judgment

`ACTIONS_GIT_REF: <ref>` requiring a "matching observed diff" is not byte-closed. `git status` and an unparameterized `git diff` do not define the comparison base, expected artifact set, accepted committed-clean case, or ref-to-candidate binding. A resolvable commit is insufficient, as the reverted Row-3 attempt already proved, but section 13 supplies no stronger algorithm.

Define the exact refs and observations (for example, trusted base/head inputs plus the expected path/artifact source and canonical diff command), name the fence row that supplies them, or explicitly defer the IMPL row for s8. Until then, Group B cannot stamp E1 without implementer inference.

## Nonblocking Confirmation

The porcelain v1 fixed-column rule is directionally correct for non-`-z` output: preserve `line[0:2]`, require the separator at byte 2, and never trim the leading status column. On revision, pin the exact command/options and malformed-line disposition so this remains a byte grammar rather than an example.

No c2 mechanism reopen is requested. The required change is to make the s8 operational table honest about which locked predicates are observable on the current surface; unavailable predicate inputs may degrade or be deferred, but they may not be replaced by a different predicate and called closed.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design or frank source/test edit, branch, commit, push, PR, merge, or Row-3 re-lift
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `frank` main checked separately and the active s8 worktree retains only pre-existing slice-relay bookkeeping changes outside this review
Next requested action: m-3.planner revises section 13 for F1-F4 and returns a new directly addressed DESIGN relay; master and s8 keep Row 3 held meanwhile
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-row3-precision/DESIGN-REVIEW-implementer-20260712-144500.md`
