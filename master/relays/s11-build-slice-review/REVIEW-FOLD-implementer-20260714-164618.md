## REVIEW-FOLD — optional findings 1 and 2 folded, targeted E2 green, existing draft PR pushed at `e86644d`; findings 3–11 explicitly recorded no-change; no merge

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s11-build-slice-review
PARENT_DISPATCH_ID: s11-build-slice-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — merge remains operator-only (`HUMAN_MERGE_AUTHORIZATION` at grant time); T5/T10 remain acceptance-OPEN behind g2/dc
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-slice-review/REVIEW-FOLD-planner-20260714-163830.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the two review-recommended optionals are folded in one attributable commit and pushed to draft PR #1; the edge relation is single-source and relation-tested, replacement cadence slots are independently asserted from the scheduler constructor, targeted packages plus every TestS11 fixture are green; the other nine optionals are intentionally no-change; request planner merge-decision relay TO operator, not merge

## Summary

Optional finding 1 is folded without changing edge membership: `internal/store/projections.go` now owns one named acceptance-bounce subset, Bucket D extends that exact subset only with `stale_choice_set`, and `internal/engine/fsm.go` consumes the shared classifier. `internal/bounce/edges_test.go` asserts the relation and the deliberate exclusions (`egress`, `stale_schema`, empty).

Optional finding 2 is folded without changing cadence values: `ArmParked` now consumes one `g4ResummonInputs` constructor, and the stale-schema replacement test independently asserts `g4-no-response-1` and `g4-answered-stalled-1` before checking the replacement content hash. The prior hash-only assertion is no longer the sole restart proof.

PR: draft PR #1 — https://github.com/iwnlcern/frank-dev/pull/1 — API-confirmed head `e86644ddf10ca9bbdc4c098f443ad3eab73c4e20`

Plan lock: `s11-comms-thicken-plan-s11.3`; no locked-design, task, terminal-enum, bucket-membership, cadence-token, T5/T10, or merge-state change.

## Findings disposition

1. **Folded:** shared acceptance-bounce subset plus relation assertion; Bucket D remains the subset plus `stale_choice_set`.
2. **Folded:** replacement scheduler inputs independently assert both G4 cadence-slot values and drive the content-hash check.
3. **Recorded no-change:** the first `violation` assignment is ineffectual but the sentinel/fail-closed behavior is correct; tidying it adds no behavioral proof.
4. **Recorded no-change:** marshal/owner order differs only on the unreachable marshal-failure path and converges to the same deny/kill result.
5. **Recorded no-change:** the in-memory snapshot makes the former `tables.Build` error path unreachable; reintroducing it would undo the reviewed T8 snapshot refactor.
6. **Recorded no-change:** multiple accepted resolutions for one gate are prevented by the terminal gate surface; no order semantics are invented for an unreachable state.
7. **Recorded no-change:** the schema-version default is unreachable behind the `TargetSchemaVersion < 1` rejection guard.
8. **Recorded no-change:** symbolizing the host-side obligation error would change host-facing diagnostics; the seat-facing record remains path-free.
9. **Recorded no-change:** same host/composition-root diagnostic class as finding 8 and a pre-existing pattern, not a bounded behavior-preserving fold.
10. **Recorded no-change:** historical RED provenance in commit `99e945c` cannot be repaired by a later source byte; the relay ledger remains the honest evidence.
11. **Recorded no-change:** the two all-false rows remain backstopped by positive-membership rows on the same projections; no false-green gap exists in the complete matrix.

## Scope and files

FOLD_SCOPE:
- frank/internal/bounce/edges.go -> in
- frank/internal/bounce/edges_test.go -> in
- frank/internal/engine/fsm.go -> in
- frank/internal/store/projections.go -> in
- frank/internal/engine/resummon.go -> in
- frank/internal/engine/resummon_test.go -> in
- frank/.relays/s11/fold-optionals-red-green.md -> in
- frank/.relays/s11/mechanical-tables.md -> in
FOLD_SCOPE_RESULT: all-in

The required pre-edit artifact is `master/relays/s11-build-slice-review/FOLD_SCOPE-implementer-20260714-164316.md`; it was indexed and exact-file lint-clean before any frank edit. Final diff touches seven of its eight allowed rows and no path outside it. `.relays/s11/mechanical-tables.md` remains path-total against `d91fcfb..e86644d` (empty `comm -3` reconciliation).

Files changed: `.relays/s11/fold-optionals-red-green.md`, `.relays/s11/mechanical-tables.md`, `internal/bounce/edges_test.go`, `internal/engine/fsm.go`, `internal/engine/resummon.go`, `internal/engine/resummon_test.go`, `internal/store/projections.go`.

## Acceptance and boundary proof

- T1–T4/T6–T9/T11 remain E2-complete; the fold strengthens T4/T6 maintenance proof without changing the observable contract.
- T5/T10 remain acceptance-OPEN per FINDING-4 because g2 implementer/master closure and dc return are still absent.
- The T8 master rescope remains eight-of-nine; item 2 remains the named post-Step-2 m-7+m-3 carry.
- Boundary contract is unchanged: rejected acceptance-stage edges feed the author-facing D projection; only the four shared edges derive `bounced_repair`, while `stale_choice_set` remains D-only. G4 input remains local summon only and carries the same two stable cadence tokens.
- Forbidden families remain untouched: no `internal/observe/`, `registry.json`, store write path, raw seat surface, merge, or deployment edit.

## RED/GREEN and verification

`.relays/s11/fold-optionals-red-green.md` records genuine pre-production REDs: undefined `AcceptanceBounceEdge` / `BucketDFailingEdge` and undefined `g4ResummonInputs`. The first classifier placement exposed an import cycle through the existing bounce formatter dependency graph; it was corrected inside FOLD_SCOPE by making store the single classifier owner.

Fresh post-commit targeted verification at `e86644d`:

- `go test -count=1 ./internal/bounce ./internal/store ./internal/engine` — exit 0 (`bounce` 0.614s, `store` 1.259s, `engine` 2.290s).
- `go test -count=1 ./test/fixtures -run '^TestS11'` — exit 0 (3.339s).
- `go vet ./internal/bounce ./internal/store ./internal/engine ./test/fixtures` — exit 0.
- `git diff --check HEAD^` — exit 0.
- `git status --short` — no output.
- `git ls-remote origin refs/heads/s11-comms-thicken` and the PR API both resolved `e86644ddf10ca9bbdc4c098f443ad3eab73c4e20` after push.

Evidence level: E2 local targeted tests/vet plus E1 diff/scope/remote-ref proof. No E3/E4 or merge claim.

ACTIONS_GIT_REF: `s11-comms-thicken@e86644ddf10ca9bbdc4c098f443ad3eab73c4e20`; one review-fold commit `e86644d` (`refactor: fold s11 review optionals`) over reviewed head `547ada9`; pushed to `origin/s11-comms-thicken`; draft PR #1 API-confirmed at the same head; no merge
FINAL_GIT_STATUS_SHORT: none — clean s11-comms-thicken worktree at `e86644ddf10ca9bbdc4c098f443ad3eab73c4e20`, equal to upstream

Remaining risk: findings 3–11 remain explicit ledger observations, not silent omissions. T5/T10 and merge remain gated exactly as before.

Next requested action: s11.planner re-verifies this bounded fold and authors the merge-decision relay TO the operator. Merge remains operator-only; this report grants none.
