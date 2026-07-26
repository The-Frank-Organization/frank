## DESIGN-REVIEW — rev6 MUST REVISE: PREPARING replay intent closes, but the claimed total disposition table omits bootstrap and has overlapping reject branches

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r6
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-184737.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: byte-exact re-review of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev6 at SHA-256 `b17fa35dfb4e37be48cd912586ed3e1580f78320527e7a8ec43c1095d0de4d9a` — the two requested folds are present, but the replacement table is not yet total or mutually exclusive

## Verdict

**MUST REVISE.** R5-F2 closes. R5-F1's intended live-transition behavior also closes in substance: a target-equal retransmission joins without resetting the deadline or duplicating the cut, and a conflicting proposal cannot replace the active target. The fixtures now cover both correlation forms and conflict.

The exact six-branch table nevertheless omits the fresh-broker/no-installed-state case already required by the study's broker-death convergence path, and its reject predicates overlap while PREPARING. One small ordered or mutually-exclusive table correction is required before m-10 can confirm the wire contract.

## Closed From R5

- **R5-F1 CLOSED at mechanism level:** target-equal replay joins the live transition under either correlation; the original deadline, target, and cut identities remain unchanged; `transition-started` is never install proof.
- **R5-F2 CLOSED:** Q2 and section C now state the identical two-form logical gate. No event-only assign-gate phrase remains.

## Finding

### R6-F1 — BLOCKER — the “TOTAL” proposal table omits no-installed-state bootstrap and contains overlapping active-transition reject predicates

**Omitted branch:** the table covers no-active equality, same-epoch advancement, newer epoch, active-target equality/difference, stale/regressing, and malformed input. A fresh broker has **no installed state** and no active transition, so a valid first `state_proposal` matches none of those branches. That state is reachable and normative: section Q3.3.1's broker-death convergence leg says the fresh instance installs the proposed durable tuple over an empty in-flight set and returns proof under the new correlation.

**Overlapping branches:** while PREPARING, a stale/regressing proposal is both “DIFFERENT from the active target” and “stale/regressing tuple (any state),” permitting either `rejected-transition-active` or `rejected-stale`. A malformed proposal similarly cannot be safely target-compared but is not given explicit precedence over the active-different branch. The table is exhaustive in labels but not deterministic in predicates.

**Required correction:** make the table ordered or make every predicate disjoint, and pin these cases exactly:

1. No installed state + no active transition + valid proposal installs synchronously over the empty in-flight set and returns one exact disposition/proof shape.
2. Malformed input has one explicit precedence and disposition in every broker state.
3. A well-formed stale/regressing proposal during PREPARING has one exact disposition; regardless of the chosen token, it leaves the active target and original deadline untouched and is never proof.
4. Add one fixture for each case: fresh-instance first proposal, malformed-during-PREPARING, and stale/regressing-during-PREPARING. Assert one response token, no ambiguity, no target/deadline mutation, and no extra cut.

This does not reopen the new enum, require durable transition state, or alter the clean-cut/no-H-24 determination.

## Re-review Gate

Return one byte-exact rev7 preserving the closed r1-r5 mechanisms and replacing only the disposition table/fixtures needed to cover bootstrap and remove predicate overlap. Consumer confirmations, join record, SITREP, and re-lock remain held until those bytes pair-approve.

## Verification

- Reviewed rev6 at exact SHA-256 `b17fa35dfb4e37be48cd912586ed3e1580f78320527e7a8ec43c1095d0de4d9a`; incoming relay at exact SHA-256 `6fbb187b9144525469122915ffd5c480e41edafe0317e2c86d7109f3ac714932`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Incoming exact-file relay lint reports `OK`.
- No product design or code bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner folds R6-F1 into one exact rev7 and relays it for byte-exact adversarial re-review; m-9/m-10 confirmations and re-lock remain held on rev6.
