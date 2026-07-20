## IMPL — s11 runnable slice built, verified, pushed, and opened as draft PR #1; T5/T10 correctly report-and-hold with acceptance OPEN; request the one end-of-slice adversarial review

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s11-build-report
PARENT_DISPATCH_ID: s11-build-token
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — merge remains operator-only; T5 and T10 also remain acceptance-open behind g2 and dc respectively
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-token/IMPL-planner-20260714-145200.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the authorized runnable s11 set is on `s11-comms-thicken@547ada9`, pushed to the private frank-dev remote with draft PR #1; T1-T4/T6-T9/T11 are E2-green, T5/T10 remain explicit FINDING-4 report-and-hold, and the eight-retained-item T8 wording is preserved; request the single whole-slice pair review, no merge

## Summary

The runnable s11 slice is implemented on the exact dispatched base and is ready
for the one B11 end-of-slice adversarial review. The branch adds the B/C/D saved
queries, completes the seven-state local FSM surface, hardens parked decisions
across schema changes, registers the terminal/edge and known-A negatives,
lands the eight retained cleanup items in card order, and moves both G4 timers
onto startup-pinned operator config without any auto-approve policy.

T5 and T10 are not claimed complete. The live inbox was re-audited immediately
before handoff: its newest row remains the 145200 build token. g2 still has only
the m-5 planner proposal (no implementer review/master completion), and no dc
design-cell return exists. Both tasks therefore exit report-and-hold with their
acceptance OPEN, exactly as FINDING-4 requires; no stub or assumption was added.

## PR and plan lock

- PR: private draft [frank-dev #1](https://github.com/iwnlcern/frank-dev/pull/1), `s11-comms-thicken` → `main`.
- Base: `main@d91fcfb340b029c39c8493084ce2f227409aa546`.
- Head: `s11-comms-thicken@547ada9aa89b6edcb98769ba27f0418439236441`.
- Plan: `s11-comms-thicken-plan-s11.3`, consuming the SHA-pinned guide plan and
  the r3 plan/review/token chain.
- Branch shape: 16 commits, 35 paths, 2431 insertions, 485 deletions.

## Task and acceptance status

| Task | Commit(s) | Status |
|---|---|---|
| T1 bucket B | `a30df05` | live, non-interrupting, raise-only |
| T2 bucket C | `1db30c5` | live operator-CC FYI, no decision obligation |
| T3 bucket D | `1b1ed68` | rejected acceptance-edge author return; egress stays A |
| T4 seven-state FSM | `219e8db` | `bounced_repair` + fixture-scoped local `egress_blocked` live |
| T5 elaborate-more fork | none | REPORT-AND-HOLD; g2 incomplete; acceptance OPEN |
| T6 8a hardening | `99e945c` | stale reject + held signal + new-identity crash-safe reissue live |
| T7 matrix/known-A NF | `e076bc4` | 14 boundary rows + raise-only negative live |
| T8 cleanup | `e28282d`..`d27fa5a` | eight retained items, one commit/full battery each |
| T9 G4 config | `6e25f20` | both timers operator-configured; no auto-approve encoding |
| T10 re-prompt/claimless-held | none | REPORT-AND-HOLD; dc absent; acceptance OPEN |
| T11 exit evidence | `547ada9` | tables, labels, I-PH, scope, catch ledger reconciled |

T8: eight of nine — item 2 rescoped by master
(`s11-build-escalate-fence/RECONCILE-…-143010`), carried post-Step-2 to
m-7+m-3.

## Files changed and boundary-contract proof

The 35 changed paths reconcile one-for-one to
`frank/.relays/s11/mechanical-tables.md`. The diff-to-license table preserves
same-file task order for submit, resummon, ODB, FSM, and main. The three
explicit forbidden families are absent from the branch diff:

- no `internal/observe/` change;
- no `internal/fieldspec/registry.json` change;
- no store write-path change (`store.go`, `lock.go`, `quarantine.go`).

The branch also leaves the T5 executor fork seam untouched, does not build the
away mirror, does not activate external send, and does not change locked
terminal vocabulary. The only executor change is T8's master-granted
`finalizeRun`/cleanup ownership seam; `Spawn`, refusal rules, fault tokens,
verdict logic, rung values, and timing values are byte-untouched.

## Tests and verification

Fresh committed-head command:

```text
go test -count=1 ./... && go vet ./...
```

Exit 0 at `547ada9`: `test/fixtures` 130.930s; `test/invariants` 4.157s;
`test/replay` 1.644s; dogfood 1.465s; zeroloss 2.131s; vet emitted no error.

Focused exit proof also passed:

- all `TestS11*` fixtures, including the real-process T6 crash replay and two
  production-binary G4 timer cases;
- `TestLawPathHygiene`, `TestLawTerminalEnumByteExact`, and
  `TestLawR2NoModelPredicate`;
- 35/35 diff paths found in the mechanical table;
- forbidden-seam scan empty; `git diff --check d91fcfb...HEAD` clean;
- final `git status --short` empty.

## Evidence levels, remaining risk, and next action

Evidence is E2: local compiled production binaries, durable-store/crash
fixtures, invariant suites, and replay suites. It is not deployed E3/E4 proof.
The durable evidence package is `frank/.relays/s11/`, especially
`t6-red-green.md`, `t7-matrix.md`, `t8-cleanup.md`, `t9-red-green.md`,
`mechanical-tables.md`, and `t11-exit.md`.

Remaining work is deliberately authority/gate-bound: s11.planner's one
whole-slice review; any resulting REVIEW-FOLD; g2/dc completion and separately
authorized T5/T10 work; then an operator-addressed merge decision. Draft PR #1
does not grant merge authority.

ACTIONS_GIT_REF: `s11-comms-thicken@547ada9aa89b6edcb98769ba27f0418439236441` over exact base `main@d91fcfb340b029c39c8493084ce2f227409aa546`; pushed to private `origin/s11-comms-thicken`; draft PR #1 opened; no merge, tag, or worktree cleanup.
FINAL_GIT_STATUS_SHORT: none — clean tree
Acceptance criteria status: T1-T4/T6-T9/T11 satisfied at E2; T5/T10 correctly report-and-hold with acceptance OPEN; T8 honestly eight-of-nine under the master rescope.
Boundary contract proof: 35/35 paths licensed; forbidden observe/registry/store-write families absent; exact terminal enum and R2 invariants green.
Out-of-scope preserved: T5 fork, T10 edges, away send, external bridge, E3/E4, registry data, store write path, merge.
Remaining risk: end-of-slice review has not yet run; g2/dc remain open; local E2 is not deployed proof.
Next requested action: s11.planner performs the single end-of-slice adversarial review against `d91fcfb...547ada9` and draft PR #1, then returns bounded REVIEW-FOLD findings or routes the still-open gate state and merge decision upward; only the operator may authorize merge.
