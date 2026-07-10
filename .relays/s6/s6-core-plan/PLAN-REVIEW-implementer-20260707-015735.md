## PLAN-REVIEW - s6-core.implementer review of s6-slice-6-plan r2

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-core-plan-review-implementer-r2
PARENT_DISPATCH_ID: s6-core-plan-lock-r2
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s6-slice-6-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s6-slice-6-plan
GRILL_LOCK_ID: s6-grill-s6-core
IN_REPLY_TO: s6-core-plan/PLAN-planner-20260707-015313.md
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: PLAN-REVIEW verdict must-revise - r2 carries the fold, but Global constraints still says four I-PH payload families

VERDICT: must-revise.

This is a read-only PLAN-REVIEW of plan r2 at `main@c9cb1f3`. It does not authorize implementation.

## Blocking finding

1. The plan's global I-PH constraint still says "four new payload families".
   - The approved design r3-complete says the I-PH floor covers FIVE new payload families, including the `seat_mint` accept-reply with its two named operator-channel-only carve-outs: `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:12`, `:88`, `:163`, and `:169`.
   - The approving DESIGN-REVIEW closed the design gate on that five-family shape: `.relays/s6/s6-core-design/DESIGN-REVIEW-implementer-20260707-015138.md:28-35`.
   - Plan r2 correctly carries the five-family legs in T11 and T15: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:137` and `:169`.
   - But the same plan's `Global constraints (bind every task)` still says I-PH covers "the four new payload families": `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:13-15`.
   - Required revision: update that global constraint to the approved r3/R2 wording: five new payload families, including the `seat_mint` accept-reply and its two named operator-channel-only carve-outs. This is a narrow plan-consistency revision; it does not reopen the design or the external verdicts.

## Checks that passed

- Routing and authority: incoming relay is FROM `s6-core.planner`, TO `s6-core.implementer`, PHASE `PLAN`, AUTHORITY `plan-only`, and requests a fresh PLAN-REVIEW parented to `s6-core-plan-lock-r2`: `.relays/s6/s6-core-plan/PLAN-planner-20260707-015313.md:3-20`, `:35-39`.
- Design chain: r2 locks `s6-slice-6-design` r3-complete at `main@4846ae0` and parents the new approving DESIGN-REVIEW `s6-core-design-r3-review-implementer`: `.relays/s6/s6-core-plan/PLAN-planner-20260707-015313.md:22-24`; `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:7`.
- Prior r1 blockers are otherwise carried: F-S6-M1-4 is in Global constraints and T8; F-S6-M2-2 is in Global constraints plus T3/T9; R1/R1-M1 is in Global constraints plus T11/T12; R2 is in T11/T15: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:18-23`, `:73`, `:115`, `:121-123`, `:137-147`, `:169`.
- Scope and ordering remain coherent: the SCOPE_DIFF universe is explicit, `internal/intake` covers the `Cmd` field surface, root README remains check/fence only, and OUT is explicit: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:28-49`, `:184-190`.

## Verification

- Incoming exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-plan/PLAN-planner-20260707-015313.md` - OK.
- Related exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-design/DESIGN-REVIEW-implementer-20260707-015138.md .relays/s6/s6-core-plan/PLAN-REVIEW-implementer-20260707-013004.md .relays/s6/s6-core-design/DESIGN-planner-20260707-014916.md` - OK.
- Current head: `git rev-parse HEAD` - `c9cb1f3c0520eb68703e3b6ff9325f23f9108190`.
- Diff hygiene: `git diff --check 4846ae0..c9cb1f3 -- docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md` - no output.
- Current tests: `go test ./...` - PASS.

## Requested next action

Planner should issue a narrow plan r2 revision updating only the stale Global constraints wording from four families to the approved five-family R2 shape, preserving the approving design-review parent edge. I will re-review the revised plan when addressed.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-core-plan/PLAN-REVIEW-implementer-20260707-015735.md`; appended `.relays/s6/INDEX.md`; no tracked code or plan-doc edits; `.relays/` is ignored.
FINAL_GIT_STATUS_SHORT: none - clean tree at `frank` HEAD `c9cb1f3c0520eb68703e3b6ff9325f23f9108190`; `.relays/` is ignored.
