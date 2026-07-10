## PLAN-REVIEW - s1-core.implementer narrow re-review of s1-slice-1-plan r3

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s1-core-plan-lock-r3-implementer-review
PARENT_DISPATCH_ID: s1-core-plan-lock-r3
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - m-1 external approval, SCOPE_DIFF, delegated dispatch, and merge gate remain required
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
IN_REPLY_TO: s1-core-plan-lock-r3/PLAN-planner-20260703-172222.md
PLAN_LOCK_ID: s1-slice-1-plan
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_REVIEW_VERDICT: approve
BASE: main@50b2b94
SUBJECT: plan r3 narrow re-review approve - F-M1-1 fold and guide key-scheme nit are buildable

## Phase Discipline

This is a read-only PLAN-REVIEW response to
`.relays/s1/s1-core-plan-lock-r3/PLAN-planner-20260703-172222.md`.

No source files, fixtures, docs, or sprint specs were edited. I did not start
implementation because this relay is plan-review only and carries no valid
implementation dispatch token.

## Verdict

Approve for the narrow r3/r5 fold scope requested by the addressed relay.
I found no blocking issue in the F-M1-1 fold or the guide key-scheme
parenthetical.

## Review Scope

- Incoming relay: `.relays/s1/s1-core-plan-lock-r3/PLAN-planner-20260703-172222.md`.
- Requested scope: narrow read-only review of plan r3 plus design r5 delta only.
- Prior pair approval outside this fold remains the r2 PLAN-REVIEW at
  `.relays/s1/s1-core-plan-lock-r2/PLAN-REVIEW-implementer-20260703-162820.md`.
- Commit reviewed: `main@50b2b94`.

## Fold Verification

1. F-M1-1 is folded in the m-1-approved shape (b).

   Evidence:
   - m-1's blocking finding allowed shape (b) if S1 has no remint/recycle
     operation: duplicate mint must return existing binding or reject without
     creating a second credential, with an E2 no-stale/parallel-credential test
     (`.relays/s1/s1-fidelity-m1/SITREP-implementer-20260703-171028.md:32-35`).
   - The orchestrator bounded the revision to one lifecycle invariant plus E2
     fixture in Task 5, with the guide nit riding along
     (`.relays/s1/s1-core-plan/SITREP-orchestrator-planner-20260703-171643.md:26-56`).
   - Design r5 now states S1 has no remint/recycle/rotation operation, exactly
     one credential generation per seat, typed rejection on already-bound
     `mint_seat`, no second credential, unchanged binding table, existing
     credential still resolving, and unbound credentials rejecting before
     staging (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:45`).
   - Plan r3 Task 5 now exposes the buildable interface:
     `ErrSeatAlreadyBound`, no second credential, byte-identical binding table,
     and exactly one credential generation per seat
     (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:118-123`).
   - The Task 5 E2 leg checks double mint, binding bytes before/after, original
     credential resolution, exactly-one credential scan, and absent-credential
     rejection before staging
     (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:123`).

   Result: this is buildable and idempotence/testable enough for implementation.
   The shape avoids pretending rotation/revocation is live in S1 while preserving
   the m-1 carry as a later additive slice.

2. The guide key-scheme nit is folded without changing the outbox envelope.

   Evidence:
   - The guide nit asked for a parenthetical that `gate_record_ref` is
     `source_record_ref` with `source_kind=gate`
     (`.relays/s1/s1-core-plan/SITREP-orchestrator-planner-20260703-171643.md:40-43`).
   - Design r5 adds exactly that parenthetical in D-7
     (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:112`).
   - The r4 two-source ODB envelope remains unchanged and open, with no
     `model_name` slot or drain/egress semantics added
     (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:115`).

   Result: the fold clarifies one key scheme; it does not re-open the r4 outbox
   envelope approval.

## Carry-Forward Checks

- The addressed relay keeps this review narrow and says no dispatch token exists
  in the run (`.relays/s1/s1-core-plan-lock-r3/PLAN-planner-20260703-172222.md:33-35`).
- F11 still asserts exactly one pivot per mutation class
  (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:13,158-164,167`).
- P1/SWEEP still captures all fixture output surfaces, including push frames and
  tool descriptions, not only bounce/error text
  (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:175-181`).
- Out-of-scope lines remain intact: no genesis/quarantine/GC, migrators,
  outbox drain/egress, `model_name`, pair-Planner grant rendering, master or
  extracted writes, or real-runtime lane qualification
  (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:203-205`).

Non-blocking note: design §6.4 still contains the old shorthand
"revoke-on-remint" in the open-items list
(`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:185`).
I do not treat it as a blocker because the controlling r5 D-3 line, r5 fold log,
and Task 5 plan interface now define S1 as no-remint/no-recycle with typed
duplicate-mint rejection. If the planner wants to reduce m-1 reviewer friction,
that shorthand can be cleaned before the m-1 re-route without changing the
approved implementation contract.

## Approval Boundary

This approval is limited to the r3 plan and r5 design delta described above. It
does not:

- dispatch implementation,
- clear m-1's external approval,
- clear SCOPE_DIFF,
- clear delegated-dispatch lineage,
- clear merge or live-verification gates,
- approve any scope outside the current s1-core plan surface.

## Actions

- Wrote this PLAN-REVIEW relay.
- Appended the corresponding routing row to `.relays/s1/INDEX.md`.

ACTIONS_GIT_REF: no source/test/sprint-spec edits claimed; relay artifact written at `.relays/s1/s1-core-plan-lock-r3/PLAN-REVIEW-implementer-20260703-172725.md`; index row appended at `.relays/s1/INDEX.md`; final git status captured below.
FINAL_GIT_STATUS_SHORT: none - clean tree

## Verification

- Incoming relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-plan-lock-r3/PLAN-planner-20260703-172222.md` -> OK.
- New review relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-plan-lock-r3/PLAN-REVIEW-implementer-20260703-172725.md` -> OK.
- Final git status: `git status --short --untracked-files=all` -> empty output.

## Next Requested Action

`s1-core.planner` may report this pair approval to `s1.orchestrator-planner` for
the m-1 narrow re-route, then hold for the final external approval and
delegated-dispatch conditions.
