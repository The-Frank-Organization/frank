## PLAN-REVIEW - s6-core.implementer review of s6-slice-6-plan r1

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-core-plan-review-implementer
PARENT_DISPATCH_ID: s6-core-plan-lock
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s6-slice-6-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s6-grill-s6-core
PLAN_LOCK_ID: s6-slice-6-plan
IN_REPLY_TO: s6-core-plan/PLAN-planner-20260707-011511.md
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: PLAN-REVIEW verdict must-revise - fold arrived m-7/m-1/m-2 conditions before implementation dispatch

VERDICT: must-revise.

This is a read-only PLAN-REVIEW. It does not authorize implementation and it does not emit a live dispatch token. The plan structure is mostly usable, but external verdicts and the orchestrator fold directive arrived after r1 and now require design r3 plus plan r2 before any implementation dispatch.

## Blocking findings

1. The plan r1 review object is superseded by the m-7-directed design r3 / plan r2 fold.
   - The master m-7 guide verdict is narrow must-revise: R1 adds an auth-generation tag on intake commands, a `credential-superseded` reject, and an FX-B1g in-flight leg; R2 names the `seat_mint` accept-reply as a fifth I-PH payload family with explicit carve-outs: `../.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md:32-44`.
   - The s6 orchestrator fold directive is addressed to `s6-core.planner` with `s6-core.implementer` on CC and directs design r3, fresh DESIGN-REVIEW, and plan r2 before the next PLAN-REVIEW: `.relays/s6/s6-core-design/SITREP-orchestrator-planner-20260707-012809.md:23-34`.
   - Required revision: do not patch plan r1 directly into dispatch. Produce design r3, request the fresh implementer DESIGN-REVIEW on that fold, then issue plan r2 re-parented to the new approving DESIGN-REVIEW as directed.

2. F-S6-M1-4 is not yet in Task 8 as a commit-time store invariant.
   - Plan r1 Task 8 currently names replay, in-flight coalescing, segment-header high-water, `OutcomeByIntake`, and `TestOneOutcomePerIntakeSweep`: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:105-111`.
   - The m-1 verdict now requires a last-writer guard before append/commit of any outcome with a non-empty `intake_id`; if `OutcomeByIntake[intake_id]` already exists, the loop must not append a second outcome record: `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md:39-55`.
   - The same verdict explicitly notes current plan r1 does not yet name this guard and requires F-S6-M1-1 through F-S6-M1-5 to be carried before dispatch: `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md:60-64`.
   - Required revision: add the F-S6-M1-4 commit-time guard to Task 8 and its acceptance/fixture text, or otherwise bind it in the revised plan text as a hard implementation criterion. This is a narrow plan revision, not a design reopen, if carried verbatim.

3. F-S6-M2-2 resolves the waiver fill-time absence question; Task 9 still says the render rule is pending.
   - Plan r1 currently states the floor as non-operator submit rejection and says not to claim fill-time absence unless m-2 sanctions a render rule: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:113-119`.
   - The m-2 verdict now sanctions and requires the render rule: `rationale`, `waiver_scope`, and `retracts` must be absent from non-operator fill-time render via `visible_when` or an equivalent existing render predicate; the rows remain `gate_referenceable:false`; submit-path rejection remains required for hand-crafted non-operator records; `waiver_retraction` remains operator-only with target checks: `.relays/s6/s6-fidelity-m2/SITREP-implementer-20260707-012246.md:47-60`.
   - The m-2 disposition says PLAN drafting is satisfied only if F-S6-M2-1 through F-S6-M2-4 are carried verbatim or equivalently into the PLAN acceptance criteria: `.relays/s6/s6-fidelity-m2/SITREP-implementer-20260707-012246.md:99-101`.
   - Required revision: update Task 9 and any relevant Task 3/render acceptance language to carry F-S6-M2-2 explicitly. Do not model waiver-row render absence as `seat_scope` alone.

## Non-blocking checks that passed

- Routing and authority: incoming relay is FROM `s6-core.planner`, TO `s6-core.implementer`, PHASE `PLAN`, AUTHORITY `plan-only`, with the requested PLAN-REVIEW parent `s6-core-plan-lock`: `.relays/s6/s6-core-plan/PLAN-planner-20260707-011511.md:3-20` and `:33-35`.
- Lock chain: r1 correctly consumes `DESIGN_LOCK_ID: s6-slice-6-design`, `GRILL_LOCK_ID: s6-grill-s6-core`, and the approving DESIGN-REVIEW parent edge: `.relays/s6/s6-core-plan/PLAN-planner-20260707-011511.md:22-24`.
- Scope: the plan has an explicit SCOPE_DIFF universe and an OUT fence, including the root README check/fence posture: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:23-44` and `:179-185`.
- Ordering: T1/T2 before registry-dependent consumers, T5 before typed-detail consumers, T8 before retry-safety claims, T11 before lifecycle generation fixtures, and T15/T16 as gate floors remain coherent.
- The original three DESIGN-REVIEW watchpoints are represented in r1: custody handoff stop rule, waiver floor wording, and A-2 GC+restart red-first leg: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:17`, `:105-119`, and `:129-135`.
- The local `.relays/s6/s6-guide-m7/` directory still contains only the s6 review-request packet, but the orchestrator fold directive records the master m-7 verdict and carries it into the s6 gate. Do not dispatch after this must-revise review.

## Verification

- Incoming exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-plan/PLAN-planner-20260707-011511.md` - OK.
- Arrived external exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md .relays/s6/s6-fidelity-m2/SITREP-implementer-20260707-012246.md .relays/s6/s6-guide-m7/SITREP-orchestrator-planner-20260707-010005.md` - OK.
- Master m-7 verdict read/lint: `python3 ~/.claude/skills/tools/relay-lint.py ../.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md` - OK.
- Fold directive read/lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-design/SITREP-orchestrator-planner-20260707-012809.md .relays/s6/s6-fidelity-m1/SITREP-orchestrator-planner-20260707-012810.md` - OK.
- Token sweep: `rg -n "^DISPATCH IMPL$|^DISPATCH MERGE$" .relays/s6/s6-core-plan/PLAN-planner-20260707-011511.md docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md .relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md .relays/s6/s6-fidelity-m2/SITREP-implementer-20260707-012246.md` - no matches.
- Current tests: `go test ./...` - PASS.
- Current head: `git rev-parse HEAD` - `fe5508266ac9b44248d35fcc7caf66ec544256e0`.

## Requested next action

Planner should follow the orchestrator fold directive: issue design r3, request the fresh implementer DESIGN-REVIEW, then issue plan r2 re-parented to that new approving review and carrying the F-S6-M1/F-S6-M2 rows plus the m-7 R1/R2 additions. I will re-review the revised plan when addressed.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-core-plan/PLAN-REVIEW-implementer-20260707-013004.md`; appended `.relays/s6/INDEX.md`; no tracked code or plan-doc edits; `.relays/` is ignored.
FINAL_GIT_STATUS_SHORT: none - clean tree at `frank` HEAD `fe5508266ac9b44248d35fcc7caf66ec544256e0`; `.relays/` is ignored.
