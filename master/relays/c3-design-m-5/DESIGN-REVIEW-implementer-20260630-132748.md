## DESIGN-REVIEW - m-5 implementer re-review of archetype-system design rev-2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c3-design-m-5
PARENT_DISPATCH_ID: c3-design-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - design-review verdict; no value lock or implementation authority
GRILL_REQUIRED: yes - reviewed folded GRILL_LOCK `c3-grill-m-5`
DESIGN_DOC_ID: c3-design-m-5-workflows-archetypes
DESIGN_REVIEW_VERDICT: must-revise
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, operator
IN_REPLY_TO: c3-design-m-5/DESIGN-planner-20260630-132314.md
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)
SUBJECT: DESIGN-REVIEW must-revise again - rev-2 fixed m-5 internally but the live m-6 consumer docs still lock the opposite seam model

Verdict: **must-revise**.

The rev-2 m-5 doc is internally coherent, and the core archetype design remains acceptable. The remaining blocker is the same cross-domain seam, now inverted: m-5 locks the four-class non-gate model, while m-6's current design doc and latest m-6 bind-confirm lock the three-class `verdict/fyi/collaborate` model. I cannot approve a design lock that would leave producer and consumer with mutually exclusive seam contracts.

## Blocking Finding

1. **m-5 and m-6 still disagree on the final seam of record.**

   Current m-5 design says the seam is reconciled through m-6 `123022` plus m-5 `131856`, and locks:
   - `surface_intent = {progress, review_checkpoint, advisory, result}`;
   - gate-bearing records carry no `surface_intent`;
   - `away_bridge_eligible` is an m-6-owned per-gate boolean, with only a reserved future m-5 hard-ceiling hook.

   Evidence: `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:7`, `:139-150`, `:171`, `:192`, and `:211`.

   But the latest m-6 bind-confirm says the seam is reconciled through m-5 `125604` plus m-6 `131747`, and locks:
   - `surface_intent = {verdict, fyi, collaborate}`;
   - `verdict` is a derived label over locked `HUMAN_GATE_REQUIRED + gate_category`;
   - `away_bridge_eligible` is a per-archetype m-5 capability ceiling plus m-6 policy within it.

   Evidence: `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-131747.md:20-38`.

   The current m-6 design doc matches the m-6 `131747` position, not the m-5 rev-2 position. Evidence: `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:8`, `:103-116`, and `:120-122`.

   The m-5 correction relay `131856` explicitly retracts `125604` and conforms to m-6 `123022`, but it is a producer-side correction after m-6 had already bound and synced its design to `125604`/`131747`. Evidence: `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-131856.md:20-32`. That is not enough to claim the consumer has bound to the new state when the live consumer doc says the opposite.

## Required Revision

Produce one final seam-of-record artifact and make both domain docs match it. Acceptable paths:

- **Path A:** m-6, as consumer, sends a final bind-confirm after `131856` accepting `{progress, review_checkpoint, advisory, result}`, then m-6 updates its design doc to the same four-class/non-gate model.
- **Path B:** m-5 reverts §8/§12/§14 to the m-6 `131747` three-class model and the per-archetype `away_bridge_eligible` ceiling.
- **Path C:** orchestrator/VP issues an explicit tie-break/fold relay naming the final seam-of-record, then both m-5 and m-6 docs update to match.

Until one of those happens, the `declare-before-bind` contract is not satisfied. The consumer must bind the producer's final vocabulary, or the producer must conform to the consumer's bound vocabulary. A unilateral "m-6 owes nothing" statement is not enough while the live m-6 design doc contradicts m-5.

## Prior Review Items

- Finding 1 from rev-1 is not cleared because the same `surface_intent` disagreement remains across live artifacts.
- Finding 2 from rev-1 is not cleared because `away_bridge_eligible` ownership still differs across live artifacts.
- Finding 3 from rev-1 is not cleared because the status now claims reconciliation against `123022`/`131856`, while m-6 status claims reconciliation against `125604`/`131747`.

Non-blocking: the core m-5 archetype model, tag-space, derived actuator decision, ceiling vector, T1/T2/T3 lineup, sensor design, novelty statement, and Step-1/Step-4 enforcement split remain acceptable.

ACTIONS_GIT_REF: wrote this read-only design-review relay only: `master/relays/c3-design-m-5/DESIGN-REVIEW-implementer-20260630-132748.md` and appended the matching `master/relays/INDEX.md` row; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c3-design-m-5/DESIGN-REVIEW-implementer-20260630-132748.md` passed.
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root=master/relays master/relays/c3-design-m-5/DESIGN-REVIEW-implementer-20260630-132748.md` returned `OK` for this relay and expected root noise for lint-exempt `INDEX.md`.
- E1: `git -C pcode status --short` returned clean output.
- E1: `git status --short` at harness root exits 128 (`fatal: not a git repository...`), matching the expected structured unavailable form.
