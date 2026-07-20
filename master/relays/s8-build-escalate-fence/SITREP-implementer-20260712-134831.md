## SITREP — Row 3 reverted after adversarial review: the quoted absence-floor sentence lacks the phase table and observation inputs needed to stamp an honest rung; owner precision required, Rows 4–5 remain independent

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-row3-precision
PARENT_DISPATCH_ID: s8-build-row2-go
RUN_ID: s8
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — m-3 owner precision and master fence ruling are required; live adoption and merge remain operator-only
GRILL_REQUIRED: no
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: Row 3 (five-row re-lift)
IN_REPLY_TO: master/relays/s8-build-t10-close/SITREP-planner-20260712-073000.md
FROM: s8.implementer
TO: s8.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: Row 3's first implementation was rejected by its fresh review and is byte-clean reverted at `e12a87b`; the current Candidate/schema cannot evaluate the locked per-phase base predicates without inventing inputs or policy, so please route the smallest precision question to m-3 while independent Rows 4–5 continue

Row 2 is landed and task-review approved at `b50d822`. Row 3 then followed the granted RED-first path: `fe8d12a` wired the absent-declaration evaluator through the production submit handler and proved a real-socket E0→E1 transition. The fresh task review correctly returned `CHANGES_REQUIRED`: that evaluator never consumed `PHASE`, treated any resolvable `@ref` plus a matching status as `phase-done`, rejected honest `none ...` forms, and corrupted Git porcelain's meaningful leading status column with `TrimSpace`. Its fixture mirrored the same normalization and therefore could not catch the defect.

The feature commit is now neutralized by the normal history-preserving revert `e12a87b`. `git diff b50d822..e12a87b -- internal/observe/registry.go internal/engine/submit.go test/fixtures/s8_exit_gate_test.go` is empty; the focused present-declaration typed-reject + aggregation controls pass; Row 2's mechanism is untouched. The branch carries only the already-existing `.relays/s8/` bookkeeping dirt outside committed source.

**Why this needs owner precision rather than another inferred fix:** the quoted sentence requires the base **phase-shaped done predicate** plus Git-provenance observation. The locked m-3 source names materially different sets: read-only artifact/no-unauthorized-action; IMPL expected diff + artifacts + checks + conditional scope subset; MERGE/LIVE target merge/deploy/live proof. The production evaluator receives `Candidate{Record, Seat, Phase, Authority}`. The current governed surface supplies neither an allowed-artifact predicate, required-check results when declarations are absent, diff/scope facts, nor deploy/live proof; it also supplies no closed `SITREP` mapping, although SITREP is the live production fixture phase. `PHASE + ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT` cannot prove those predicates. Repeating the generic-ref implementation would stamp E1 for a predicate the conductor never observed.

**Smallest m-3 question:** for the current s8 `Candidate` and registry surfaces, provide the exact closed phase→base-predicate table for every live phase including `SITREP`; state which unavailable inputs degrade versus veto; define the canonical byte grammar mapping no-action/clean-tree declarations to raw Git observations; and, where IMPL/MERGE/LIVE require facts not present today, name the smallest sanctioned input/fence expansion or explicitly defer those phase rows. The implementation will not infer semantics from body prose or from an arbitrary existing commit.

Rows 4 and 5 share neither bytes nor semantics with this blocker and remain authorized by `…-055510`/`…-060500`; I will continue them while this answer routes. The full-set exit gates remain held: no exact-head battery, fidelity closure, slice-exit package, merge, push, PR, tag, or live adoption is claimed.

## Verification

- Row 3 review: `.superpowers/sdd/four-row-3-review.md` — `CHANGES_REQUIRED`, with the missing phase dispatch and porcelain-byte findings.
- Blocker analysis: `.superpowers/sdd/four-row-3-blocker.md` — locked input inventory, live grammar examples, and smallest owner question.
- `git show --stat e12a87b` — targeted revert of exactly the three Row 3 paths.
- `git diff b50d822..e12a87b -- internal/observe/registry.go internal/engine/submit.go test/fixtures/s8_exit_gate_test.go` — empty.
- Focused present-declaration validation/aggregation controls — PASS after revert.

ACTIONS_GIT_REF: s8-observe-spine@e12a87b (Row 3 `fe8d12a` neutralized by targeted revert; effective new feature head remains approved Row 2 at `b50d822`)
FINAL_GIT_STATUS_SHORT: worktree source clean at `e12a87b`; only pre-existing `.relays/s8/INDEX.md` modification and `.relays/s8/s8-build-t9-review/` untracked bookkeeping remain
Next requested action: s8.planner routes the smallest precision question above to m-3/master; s8.implementer proceeds on independent Rows 4–5 and returns their RED→GREEN reviews; Row 3 resumes only on an owner-precise re-lift.
