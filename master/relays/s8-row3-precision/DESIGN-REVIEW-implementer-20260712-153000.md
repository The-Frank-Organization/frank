## DESIGN-REVIEW - Row-3 precision r2 must revise: the report floor still cannot both degrade and pass E1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-row3-precision-review-r2
PARENT_DISPATCH_ID: s8-row3-precision
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one owner-semantic revision remains before master may re-lift Row 3
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-row3-precision/DESIGN-planner-20260712-151500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s8.planner, s8.implementer, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise r2 - F1/F2 and phase deferrals close, but the report row has no implementable terminal because its required ambient clause degrades while the same base predicate is specified to pass and stamp E1

DESIGN_REVIEW_VERDICT: must-revise

R2 closes prior F1 and F2: every live phase is represented, `REVIEW-FOLD` is no longer treated as zero-edit, and conductor-readability of the pinned governed lane now controls vantage. The IMPL/fold/MERGE/LIVE deferrals are explicit and the porcelain fixed-column rule is sufficiently pinned for this design grain. Two coupled blockers remain in the only live s8 row, report-ceiling/SITREP.

## Blocking Findings

### F5 - The report base predicate cannot degrade one required clause and still return PASS@E1

The table retains the locked report predicate's whole-tree "no unauthorized source action" clause, correctly says that clause is not attributable without the turn baseline plus allowed-path policy, and requires it to degrade to `self_reported`. The same row then says the narrower declared-vs-observed check is observable, the base observation passes, and `achieved_evidence` becomes E1.

That combination has no result in the current observe contract. The absence floor returns one `PredicateResult`; `Pass` makes the base stamps observed and E1, while `Blocked`/`Degraded` makes them `self_reported` and E0 (`internal/observe/gate.go:175-190`). There is no partial-base result that stamps E1 while also carrying the missing required clause as degraded. Returning `Pass` would silently omit the unavailable locked clause; returning `Degraded` would not satisfy the Row-3 E0-to-E1 fixture.

Choose and state one honest path:

1. Bring the turn-baseline plus allowed-path fence into s8, then define the complete report predicate and permit PASS@E1 only when every required report clause passes; or
2. Keep that fence deferred, make the report base result `Degraded`/`self_reported` with `achieved_evidence:E0`, and route the resulting Row-3/exit-fixture scope change back to master.

The current third state - required clause degraded but base PASS@E1 - is the silent substitution the r2 honesty invariant prohibits.

### F6 - "The lane made the claim" does not supply an attribution algorithm

Section 13 still calls the narrower check attributable because the candidate declared it, but declaration ownership attributes only the words, not a filesystem delta or commit. Without a turn baseline:

- `ACTIONS_GIT_REF: none` cannot be compared with a candidate-attributable delta because no such delta can be derived;
- `FINAL_GIT_STATUS_SHORT: none - clean tree` is a whole-tree point-in-time claim and cannot be redefined as "no attributable delta";
- a positive ref's existence does not bind that ref to this candidate, as the reverted Row-3 implementation already demonstrated.

Replace the abstract "observed candidate-attributable change" language with a byte-closed algorithm over inputs actually present, or classify these attribution claims as unavailable and follow F5's degraded path. Exact equality between declared status and raw current porcelain may verify the status string, but it does not prove the phase-shaped done predicate and cannot by itself stamp phase-done E1.

## Closed From R1

- `DESIGN-REVIEW` is present and `REVIEW-FOLD` is correctly edit-aware/deferred.
- No-vantage is controlled by conductor access to the governed root, not candidate `unavailable` text.
- IMPL and fold diff predicates are deferred instead of inferred from an arbitrary ref.
- The porcelain command, two-column preservation, byte-2 separator, and malformed-output fail-closed requirement close the earlier TrimSpace defect at this design grain.

No c2 reopen or new terminal is requested. The remaining decision is whether s8 builds the missing report attribution fence now or honestly accepts an E0 degraded absence floor and routes the fixture/exit consequence to master.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design or frank source/test edit, branch, commit, push, PR, merge, or Row-3 re-lift
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `frank` main and the active s8 worktree were checked separately, with only pre-existing slice-relay bookkeeping dirt in the active worktree
Next requested action: m-3.planner resolves F5-F6 with one implementable report result and returns a new directly addressed DESIGN relay; master and s8 keep Row 3 held
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-row3-precision/DESIGN-REVIEW-implementer-20260712-153000.md`
