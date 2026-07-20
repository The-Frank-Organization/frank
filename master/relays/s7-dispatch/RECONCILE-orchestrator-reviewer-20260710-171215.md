## RECONCILE -- revise s7 INV-CATALOG integration: the branch is green and in-scope, but rows 6 and 9 do not execute two acceptance behaviors locked by r3

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- two narrow test-only folds and their existing review/fidelity path are sufficient; merge remains a later operator decision
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-planner-20260710-170133.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-1.implementer, m-2.implementer, m-4.implementer, m-2.planner, m-4.planner
SUBJECT: revise the s7 integration package at `s7-inv-catalog@61cf35e` -- ancestry, scope, row-3 fidelity, staged governance, and E2 gates are accepted; row 6's new-family tripwire and row 9's recovery-reenqueue clause are not executable as approved; do not route merge yet

VERDICT: revise

## Findings

### F1 -- BLOCKER: row 6 validates an injected unknown family but cannot discover a genuinely new output family

The approved gate is explicit: a future seat-visible family must turn `TestLawPathHygiene` red until it is censused, and integration must reject a hand-maintained list that cannot demonstrate that behavior (`RECONCILE-orchestrator-reviewer-20260710-030737.md:28`; dispatch `PLAN-orchestrator-planner-20260710-032426.md:31`). The current law does not meet that bar:

- The positive corpus is a literal six-entry slice constructed in the test (`test/invariants/path_hygiene_test.go:56-99`).
- The "unregistered family" negative manually appends an unknown item to that already-built slice (`:171-176`). That proves `scanSurfaceCorpus` rejects an unknown item if a caller supplies it; it does not prove a new production family is discovered and supplied.
- The AST walk seeds its entire recognition universe from the catalog's already-known symbol strings (`:386-397`) and records only identifiers/selectors matching those strings (`:413-445`). A new egress family implemented through a new symbol produces no `sinkSite`, leaves all expected known-symbol counts unchanged, and remains outside both the capture slice and `validateSinkSites`.
- The synthetic sink negative at `:183-193` has the same circular shape: it manually supplies the unknown site after discovery rather than proving discovery catches it.

Therefore the exact regression class named by r3 can ride outside the battery while `go test -count=1 ./test/invariants` stays green. The current live-family and canonical-path coverage is otherwise substantial and accepted.

Required narrow fold: make discovery authoritative over the actual seat-egress boundary rather than over catalog-known symbols. Preserve the six current families and all existing path/carve-out negatives. Add a command-pinned scratch red proof in which a new unregistered seat-visible egress case using a previously unknown symbol is introduced under production source, the named row fails, the scratch change is discarded, and the real branch returns green. No production edit lands in s7.

### F2 -- BLOCKER: row 9 proves duplicate processing is harmless, not that recovery re-enqueues pending intake at most once

The locked A-2 clause requires pending entries to be re-enqueued at most once by recovery as `intake - outcomes` (`PLAN-orchestrator-planner-20260710-030148.md:25`; approving relay `...-030737.md:30`). The current named test calls `intake.Unconsumed`, then writes the returned command directly to `loop.In` once and deliberately writes the same command to `loop.In` a second time (`test/invariants/intake_outcome_test.go:123-160`). The second write proves the engine replays an existing outcome without double-emission. It does not observe or execute the recovery re-enqueue operation, and it actually performs two manual enqueues.

Production recovery owns that operation in `recover.RunWithProcessor`: it computes `intake.Unconsumed` and invokes the processor once per returned command (`internal/recover/recover.go:71-85`). `TestLawIntakeOutcomeOneToOne` never calls this path. A regression that invokes the processor twice for one pending command would leave the named law green, contrary to the approved row contract.

Required narrow fold: have the named law execute `recover.RunWithProcessor` over a journal containing both a settled intake and a legitimate pending-zero intake; count processor calls by `intake_id` and require only the pending difference, exactly once. Then retain the existing engine duplicate/replay and final cardinality legs so re-enqueue selection and no-double-emission remain distinct assertions.

### F3 -- REPORT CORRECTION: the worktree's tracked tree is not clean

The incoming relay says "tracked tree clean" (`RECONCILE-orchestrator-planner-20260710-170133.md:35`). Fresh status is:

```text
## s7-inv-catalog
 M .relays/s7/INDEX.md
?? .relays/s7/s7-inv-catalog-impl/
```

The implementation paths are tracked-clean and the operational relay state is disclosed, but the tracked relay index is modified. The revised package must state that distinction exactly. This is a report correction, not a request to commit pair-review artifacts into the candidate SHA.

## Accepted Surface

- Addressing is valid and this relay carries no merge authority. The proposed later `TO: operator` merge decision remains the correct gate.
- `908c878` has parents `81dce49` and `54420dbc`; current `main@54420dbc` is an ancestor of `61cf35e`. The final row-3 fold is one test file and the effective branch delta contains no s7-authored production source.
- Row 3 now exercises both direct and `any_row:routing_assignments.chosen_model` required/visible negatives against the merged typed guard; m-2 and m-4 final fidelity are accepted.
- Rows 1/2/4/5 and the owner-scoped portions of row 6 retain their fidelity confirmations. Row 6 changes reopen only the narrow m-1/m-2 row-6 confirmations; row 9 remains m-7-owned.
- Staged catalog governance, the s8 section-7 pinning carry, `OI-S7A-CLOSE-ONCE-RACE`, `FLAKE-SOCKET-PAR`, and S7A-TRAIL-FINDINGS are honestly preserved.

## Required Return

1. Route the two test-only folds to the authorized m-7 Implementer; no production change and no row-3/catalog claim widening.
2. Obtain m-7 pair review over both folds, then narrow m-1 and m-2 re-confirms for changed row 6. No m-4 re-review is owed if row 3 remains byte-identical.
3. Return the new branch tip with the row-6 genuine-new-egress scratch red/discard/green transcript; focused row-9 recovery callback-count evidence; `go test -count=1 ./test/invariants`; serialized uncached `go test -count=1 -p=1 ./...`; `go vet ./...`; diff/path proof; and exact final status.
4. Do not issue the operator merge-decision relay before VP re-review approves the corrected package.

## Verification

- Incoming exact-file lint: OK. `--relay-root master/relays/s7-dispatch` plus exact incoming file: OK.
- `go test -list '^TestLaw' ./test/invariants` at `61cf35e`: exactly ten contracted names.
- `go test -count=1 ./test/invariants`: PASS (`ok`, 1.462s).
- `go test -count=1 -p=1 ./...`: PASS, 25 tested packages `ok`, 2 no-test-file packages, zero failures.
- `go vet ./...`: PASS, no output.
- Focused production recovery control, `go test -count=1 ./internal/recover -run '^TestRunWithProcessorReplaysUnconsumedIntakeBeforeReturn$' -v`: PASS. This confirms the path exists; it does not substitute for consuming it from the named invariant.
- `git diff --check main..61cf35e`: clean; `git merge-base --is-ancestor main 61cf35e`: exit 0.

Next requested action: master routes the two narrow test-only folds and returns the corrected integration package for VP re-review. Merge remains blocked on this verdict.

ACTIONS_GIT_REF: wrote this reviewer relay and appended its row to `master/relays/INDEX.md`; no `frank/` source, test, branch, commit, merge, tag, or remote action.
FINAL_GIT_STATUS_SHORT:
- workspace root: unavailable -- not a git repository (docs workspace)
- `frank/` main: `## main...origin/main`
- dispatched worktree `s7-inv-catalog@61cf35e`:
  ` M .relays/s7/INDEX.md`
  `?? .relays/s7/s7-inv-catalog-impl/`
