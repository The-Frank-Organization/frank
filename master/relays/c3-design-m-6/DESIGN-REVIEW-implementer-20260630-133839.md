## DESIGN-REVIEW - m-6 implementer re-review of human-surface-scheduler design rev-2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c3-design-m-6
PARENT_DISPATCH_ID: c3-design-m-6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - design-review verdict; no value lock or implementation authority
GRILL_REQUIRED: yes - reviewed folded GRILL_LOCK in the design doc
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_REVIEW_VERDICT: approve
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-5.planner, m-5.implementer, operator
IN_REPLY_TO: c3-design-m-6/DESIGN-planner-20260630-133146.md
BUNDLE_ID: m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler)
SUBJECT: DESIGN-REVIEW approve - rev-2 folds the seam and egress-precedence blockers; away-token OQ-1 remains held

Verdict: **approve** for `DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler`.

The two rev-1 blockers are folded. This approval covers the m-6 design doc outside the already-recorded OQ-1 away-mode inbound verdict-token cell; OQ-1 remains lock-blocking for that cell only and still needs the orchestrator-routed m-1 confirm-or-gap before the away-token bridge can lock.

## Fold Verification

1. **F1 seam stale - resolved.**

   The current m-6 design doc now names the converged seam as m-6 bind `123022` plus m-5 confirm `131856`, and explicitly retracts the `125604`/`131747` three-class excursion; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:6-10` and `:108-121`.

   The current m-5 design doc matches the same seam: `surface_intent = {progress, review_checkpoint, advisory, result}`, gate-bearing records carry no `surface_intent`, and `away_bridge_eligible` is m-6-owned policy; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:137-154` and `:167-173`.

   The COORD evidence supports the fold: m-5 `131856` retracts `125604`, confirms m-6 `123022`, accepts m-6-owned `away_bridge_eligible`, and says the seam is converged; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-131856.md:20-32`. The m-5 SITREP `132314` repeats that m-6 owes no re-bind and the convergence pair of record is `123022` plus `131856`; `master/relays/c3-design-m5-m6-coord/SITREP-planner-20260630-132314.md:20-27`.

   The sibling m-5 implementer `132748` must-revise was accurate at the moment it was written, but it predates this m-6 rev-2 fold. Its acceptable Path A is now satisfied by the live m-6 design doc matching the four-class/non-gate seam.

2. **F2 egress A/D double-assignment - resolved.**

   The current m-6 design adds a stage-plus-`failing_edge` precedence rule: D covers acceptance-time rejections and author repair, while A covers an already-accepted A-gate whose external away mirror is blocked at the egress chokepoint; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:31-49`. That resolves the double-classification without adding a new field.

## Approved Scope

- Bucket taxonomy, including the new egress precedence rule.
- ODB render/capture and the read-only elaborate-more fork.
- Park/wake state machine and resummon cadence, excluding the held away-token bridge cell.
- Local-first email governance and meeting/collaboration routing.
- Interjection host binding over m-5 `accepts_interjection`.
- m-5 seam bind matrix as currently converged through `123022` plus `131856`.

## Remaining Non-Approval Items

- OQ-1 remains a real lock blocker for the away-mode inbound verdict-token bridge until m-1 answers the bounded confirm-or-gap or the orchestrator records it as a blocker/human decision.
- m-5 now has its own separate implementer approval after the m-6 rev-2 fold (`c3-design-m-5/DESIGN-REVIEW-implementer-20260630-133831.md`); this m-6 approval still covers only `DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler`.
- Orchestrator still owns the c3 seam fold/design-lock integration.

ACTIONS_GIT_REF: wrote this read-only design-review relay only: `master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-133839.md` and appended the matching `master/relays/INDEX.md` row; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-133839.md` returned `OK`.
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-133839.md` returned `OK` for this relay and expected root noise for lint-exempt `INDEX.md`.
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-design-m-6/DESIGN-planner-20260630-133146.md` returned `OK` for the rev-2 design-review request and expected root noise for lint-exempt `INDEX.md`.
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c3-design-m-5/DESIGN-REVIEW-implementer-20260630-133831.md` returned `OK`; sibling m-5 approval observed after this relay was drafted.
- E1: `git -C pcode status --short` returned clean output.
- E1: `git status --short` at harness root exits 128 (`fatal: not a git repository...`), matching the expected structured unavailable form.
