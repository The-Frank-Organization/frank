## DESIGN-REVIEW - m-6.implementer review of c6-fix-m-6

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-6-review-r1
PARENT_DISPATCH_ID: c6-fix-m-6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded review of doc-only consistency cleanup
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c6-fix-m-6/DESIGN-planner-20260702-210000.md
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-7.planner
BUNDLE_ID: c6-fix-m-6
OWNER: m-6 Human Surface and Scheduler c6 cleanup review

DESIGN_REVIEW_VERDICT: must-revise

I reviewed `c6-fix-m-6/DESIGN-planner-20260702-210000.md` against the CTO c6 dispatch, the source rereview inventory, the live m-6 design doc, and the m-6 README.

Most of the c6 fold is acceptable. One narrow README status drift remains, so this cannot be approved yet.

## Blocking Revision

1. m-6-F7 is only partially closed: the README still contains stale active pre-lock status text.

Evidence:
- `master/relays/c6-fix-m-6/DESIGN-orchestrator-planner-20260702-204515.md` routes F7 as a status-tail cleanup so the domain front door no longer reads as pre-lock or awaiting review.
- `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:6` is now current: `c3 LOCKED`, with the c4/c5/c6 fold trail and no PLAN/IMPL.
- `master/domains/m-6-human-surface-scheduler/README.md:37` is also current at the top status bullet.
- But `master/domains/m-6-human-surface-scheduler/README.md:38` still says `c3 DESIGN doc r2 - both must-revise blockers folded; re-review requested`, followed by `Held:` and `Design-complete gates on: implementer rev-2 approve + m-1 answer`.

Why this blocks: line 38 is still inside the live `## Status` section, immediately under the corrected locked status. A reader can still understand the domain as held, awaiting rev-2 review, and waiting on design-complete gates. That is the exact stale status drift c6 F7 was meant to clear before Step-1 PLAN work reads the README as the front door.

Required revision: remove the stale line, move it to an explicitly historical/past-tense area, or rewrite it so the current `## Status` section cannot be read as pre-lock, held, or awaiting rev-2 design-complete gates.

## Accepted Checks

- F3 is folded: decision 5 no longer claims the m-7 egress fixture exists; it says the m-7 fixture registration remains owed until m-7 lands it (`m-6 design:61`).
- F6 is folded: meeting-lane routing keys on locked `phase in {DESIGN, DESIGN-REVIEW}` while `GRILL_REQUIRED=yes` is only corroborating until m-2 has an actual v3 registry row (`m-6 design:105`).
- F8 is folded: no numeric `ARCHITECTURE.md:<line>` anchors remain in the m-6 design doc; the relevant references now use stable section anchors.
- F9 is folded: Seam-C build carries now distinguish the four `COORD-182600` carries from the c5 decision-4 carry (`m-6 design:96`, `:157`, `:177-182`).
- The CTO pre-applied F1/F2 slices remain visible: the terminal-token map preserves byte-exact `{accepted, rejected, held}`, narrows `held` to fault/fail-closed disposition only, and ordinary A-gate parking stays `accepted` (`m-6 design:44`, `:46`, `:84`).
- No mechanism change, design-lock reopen, PLAN, IMPL, spike, or `pcode/` edit is licensed by this review.

## Verification

Pre-write review evidence:
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-6/DESIGN-planner-20260702-210000.md` - OK.
- `sed -n '1,220p' master/relays/c6-fix-m-6/DESIGN-planner-20260702-210000.md` - read addressed planner relay.
- `sed -n '1,220p' master/relays/c6-fix-m-6/DESIGN-orchestrator-planner-20260702-204515.md` - read CTO c6 dispatch and constraints.
- `sed -n '1,220p' master/DESIGN-REREVIEW-2026-07-02.md` and targeted line reads - reviewed c6 source inventory and m-6 F3/F6/F7/F8/F9 assignments.
- `sed -n '1,220p' master/relays/c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md` - confirmed VP approval for CTO apply plus pair dispatch; no PLAN/IMPL.
- `nl -ba master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md | sed -n '1,210p'` - reviewed the live c6 fold.
- `sed -n '34,42p' master/domains/m-6-human-surface-scheduler/README.md` - confirmed the remaining stale README status line.
- `grep -nE 'ARCHITECTURE\.md:[0-9]' master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md` - no matches.
- `grep -nE 'DESIGN-LOCK-READY|Awaiting only' master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md master/domains/m-6-human-surface-scheduler/README.md` - no matches.
- `rg -n "m-6-F[123456789]|Awaiting only|DESIGN-LOCK-READY|re-review requested|Held:|Design-complete gates|fixture registration owed|GRILL_REQUIRED|ARCHITECTURE\.md:[0-9]|c3 LOCKED|accepted, rejected, held|confusion-resistant|governed-write" master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md master/domains/m-6-human-surface-scheduler/README.md` - confirmed the accepted c6 folds and the remaining README line 38 blocker.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

Post-write verification:
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-6/DESIGN-REVIEW-implementer-20260702-211201.md` - OK (`OK master/relays/c6-fix-m-6/DESIGN-REVIEW-implementer-20260702-211201.md`).
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-6` - OK (`OK master/relays/c6-fix-m-6`).
- `perl -ne 'print "$ARGV:$.:$_" if /[^\x20-\x7E\n]/' master/relays/c6-fix-m-6/DESIGN-REVIEW-implementer-20260702-211201.md` - no output.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no README edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128); `git -C pcode status --short` clean.
Next requested action: m-6.planner patches the README status section so it no longer presents stale pre-lock/re-review/held gates as current, then re-relays the bounded c6-fix-m-6 revision for re-review.
