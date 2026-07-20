## DESIGN-REVIEW - c6-fix-m-7 pair slice

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-7
PARENT_DISPATCH_ID: c6-fix-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - planner can fold narrow doc-only fixes
GRILL_REQUIRED: no
DESIGN_DOC_ID: c4-design-m-7-conductor-core
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c6-fix-m-7/DESIGN-planner-20260702-205730.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

The c6 fold lands most of the requested fixes, but two c6-scoped misses remain. Both are doc-only and narrow.

## Findings

1. **Blocker - F11 is defined, but the eventual PLAN acceptance gate still only requires F1-F10.**

The c6 dispatch requires the m-7-F1 fold to add a crash-between-canonical-renames fixture / one-pivot-per-mutation fixture, and the planner fold added F11 with the right mutation-class enumeration (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:172`, `:314`).

But §19 acceptance criterion 1 still says `F1-F10 + NF-S1..S18 pass as conductor-registry checks` (`...v3-conductor-core-design.md:245`). That leaves the new F11 outside the eventual Step-1 PLAN gate even though F11 is the biting fixture for the VP-ratified m-7-F1 resolution. Required revision: update §19 AC1 to require F1-F11. If §22's fixture summary is intended to summarize the current fixture set rather than the original c4 lock package, update that summary too; if §22 intentionally stays unchanged as c4 lock history, leave it and make the distinction explicit in the c6 fold-log.

2. **Blocker - the README line called out by c6 still says decision-2 fail-closed is `self_reported` only.**

The c6 rerereview source explicitly scopes m-7-F7 to README lines 47-49 and x2-claim-honesty-F8 to README line 3 plus 48-49 (`master/DESIGN-REREVIEW-2026-07-02.md:163`, `:186`). The design doc correctly widens decision-2 to `record_integrity in {self_reported, mixed}` in S7, CQ-2, and the fold-log (`...v3-conductor-core-design.md:157`, `:218`, `:277`).

The README host/execute line remains stale: `decision-2 fail-closed on authority-class self_reported` (`master/domains/m-7-conductor-core/README.md:48-49`). Required revision: align that README summary to the c6-widened contract, e.g. `record_integrity in {self_reported, mixed}` for authority-class records, while preserving the seam-matrix pointer.

## Checks Passed

- m-7-F1 main mechanism text is in the right shape: HELD embeds the candidate in one compound canonical record; `verify()` commits one compound operator-verdict record whose presence implies the burn; Phase 3 derives burn state from verdict records (`...v3-conductor-core-design.md:76`, `:93`, `:100`, `:152`).
- m-7-F5 intake single-writer discipline is explicit in the diagram, §2.1, and §2.2 (`...v3-conductor-core-design.md:28-31`, `:52`, `:57`).
- x3-F2 author set now includes m-2/m-3/m-4/m-5/m-6 in §7 and S15 (`...v3-conductor-core-design.md:109`, `:165`).
- m-7-F6 NF-S9 is regrained to non-seat destinations and no longer conflicts with seat-pipe delivery (`...v3-conductor-core-design.md:159`).
- m-7-F9 adds the writability/reachability token family to the F8 sweep, and the prior `non-lane-writable` instance is scoped to the tool surface with the D5 residual (`...v3-conductor-core-design.md:71`, `:233`).
- The byte-exact terminal enum `{accepted, rejected, held}` remains intact, and the D5/confusion-resistant claim boundary is not strengthened by the c6 folds (`...v3-conductor-core-design.md:16`, `:104`, `:166`, `:231-233`).

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-7/DESIGN-planner-20260702-205730.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-7/DESIGN-orchestrator-planner-20260702-204518.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-7` - OK before this relay write
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '1,130p'` - reviewed status, diagram, §2, §3, §5, §6, §7.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '130,240p'` - reviewed §9, §12, §13, §15, §16.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '240,340p'` - reviewed §19, §20, §21, §22 and c6 fold-log.
- `nl -ba master/domains/m-7-conductor-core/README.md | sed -n '1,120p'` - reviewed status and hosted-contract summary.
- `rg -n 'F1–F10|F1-F10|F11|self_reported|record_integrity ∈ \{self_reported, mixed\}' master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md master/domains/m-7-conductor-core/README.md` - found F11 present but AC still `F1-F10`; README m-3 line still `self_reported` only.
- `rg -n "candidate \\+ a held|separate burn record|only socket-writing|m-6/m-3/m-4-authored|BOOTING" master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md master/domains/m-7-conductor-core/README.md` - no stale-form hits for the planner's listed grep set.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-7/DESIGN-REVIEW-implementer-20260702-210010.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-7` - OK after this relay write
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C pcode status --short` - clean

ACTIONS_GIT_REF: read `master/relays/c6-fix-m-7/DESIGN-planner-20260702-205730.md`, parent dispatch `master/relays/c6-fix-m-7/DESIGN-orchestrator-planner-20260702-204518.md`, c6 source/review artifacts, the m-7 design doc, and m-7 README; wrote `master/relays/c6-fix-m-7/DESIGN-REVIEW-implementer-20260702-210010.md`; appended `master/relays/INDEX.md`; no design doc/README/source/code/`pcode`, no PLAN, no IMPL, no design-lock reopen, no spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: m-7.planner folds §19 AC1 to include F11 and aligns the README m-3 decision-2 line to `{self_reported, mixed}`, then returns r2 for review.
