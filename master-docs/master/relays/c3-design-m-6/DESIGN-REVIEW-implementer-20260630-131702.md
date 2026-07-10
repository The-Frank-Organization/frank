## DESIGN-REVIEW - m-6 implementer review of human-surface-scheduler design

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
DESIGN_REVIEW_VERDICT: must-revise
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-5.planner, m-5.implementer, operator
IN_REPLY_TO: c3-design-m-6/DESIGN-planner-20260630-131217.md
BUNDLE_ID: m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler)
SUBJECT: DESIGN-REVIEW must-revise - local surface model is sound, but the m-5 seam and egress-bounce bucket mapping are not lock-consistent

Verdict: **must-revise**.

The design is directionally strong: it keeps m-6 as a local-first projection over the locked relay/addressing substrate, promotes the ODB instead of inventing a second decision format, records the operator GRILL_LOCK, and correctly holds the away-mode inbound verdict-token cell behind the m-1 confirm-or-gap. I cannot approve design-lock yet because two boundary issues are still mechanically inconsistent in the artifacts.

## Findings

1. **Blocking - the m-5 seam is stale in the m-6 design doc and lacks the latest m-6 bind-confirm.**

   The m-6 design says the m-5 seam is resolved by COORD `122628` + `123022`, then locks `surface_intent {progress, review_checkpoint, advisory, result}` and treats `away_bridge_eligible` as an m-6-owned boolean; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:103-116`.

   That is not the current relay-tree state. The later m-5 final COORD relay adopts canonical `surface_intent {verdict, fyi, collaborate}`, declares `away_bridge_eligible` as a per-archetype capability ceiling, and says m-5 is awaiting m-6 binding confirm; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-125604.md:20-31` and `:46-47`. The folder listing has no later m-6 COORD after `125604`.

   Required revision: either file the missing m-6 bind-confirm to `125604` and fold the final three-value seam into the m-6 design, or explicitly reject/renegotiate the later m-5 final declaration in COORD before design-lock. The design doc must cite the actual final seam source and use one canonical `surface_intent` set and one `away_bridge_eligible` ownership model.

2. **Blocking - the bucket taxonomy double-assigns egress failure without a precedence rule.**

   Bucket A includes `egress_scan_result=blocked` as an operator-reaching verdict-required trigger; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:37`. Bucket D includes observe-veto/form-bounce outcomes as author-facing and non-operator-reaching; `:40-42`. Because the m-3 design treats failed egress scan as a veto, the same external-send failure can currently read as both A (operator decision queue) and D (author repair bounce).

   Required revision: add an explicit precedence/direction rule. For example, if the intended split is "away-mode external-send egress failure parks local and raises an A resummon, while ordinary observe-veto/form-bounce returns D to the author," state that exact distinction and the writer for each tag. If not, choose one route. The current text says "direction-explicit" but leaves this edge ambiguous at the first external-send chokepoint.

## Non-blocking Confirmations

- The ODB render/capture contract is acceptable as a Step-2/Step-4 split: bounded choices return through an operator-FROM verdict relay, the elaborate-more fork stays read-only, and the original gate remains parked until the fork returns; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:48-63`.
- The away-mode token bridge is not silently locked. OQ-1 is recorded as lock-blocking only for the away-mode binding cell, with the rest of m-6 allowed to lock after revision; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:75-80` and `:139-151`.
- The GRILL_LOCK is present and folded. The operator choices on bucket cut, away digest, meeting lane, resummon cadence, and write-capable fork rejection are recorded; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:128-135`.
- The local-first / egress fail-closed posture is acceptable once finding 2 is clarified; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:147-151`.

## Required Revision Checklist

- Fold or renegotiate the latest m-5 final COORD relay `COORD-planner-20260630-125604.md`.
- Update the m-6 design doc to cite the real final m-5/m-6 seam source.
- Resolve `surface_intent` to one canonical value-set across m-5 and m-6.
- Resolve `away_bridge_eligible` ownership/representation across m-5 and m-6.
- Add an explicit egress-failure precedence rule between Bucket A and Bucket D.
- Keep OQ-1 as the bounded m-1 confirm-or-gap for the away-mode inbound verdict-token cell only.
- Re-run relay-lint on the revised DESIGN relay and return a new DESIGN-REVIEW request.

ACTIONS_GIT_REF: wrote this read-only design-review relay only: `master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-131702.md` and appended a matching `master/relays/INDEX.md` row for the c3-design-m-6 design trail; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output before this relay write.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E2: `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-131702.md` returned `OK`.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-131702.md` returned `OK` for this relay and expected root noise for lint-exempt `INDEX.md`.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-design-m-6/DESIGN-planner-20260630-131217.md` returned `OK` for the incoming design-review request and expected root noise for lint-exempt `INDEX.md`.
- E1: `find master/relays/c3-design-m5-m6-coord -maxdepth 1 -type f -name '*.md' | sort` showed latest m-5/m-6 COORD file is `COORD-planner-20260630-125604.md`; no later m-6 bind-confirm exists in that dispatch folder.
- E1: `git -C pcode status --short` returned clean output.
- E1: `git status --short` at harness root exits 128 (`fatal: not a git repository...`), matching the expected structured unavailable form.
