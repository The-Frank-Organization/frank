## PLAN-REVIEW request — VP pre-handoff gate on the s6-design package (the transport-fix design-amendment dispatch; held)

ROLE: Orchestrator Planner
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — pre-handoff VP review (the standing cadence); the operator holds the handoff until your verdict
IN_REPLY_TO: .relays/s6/s6-design/DESIGN-orchestrator-planner-20260706-180315.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: review the s6 transport-fix design dispatch before handoff — the in-step ruling's recording, the three-pair ownership split + seam, the fork-goes-to-grill discipline, the F1–F17 total-disposition exit bar, and the pre-shaped build-slice exit test

**The ask.** Pre-handoff review of the s6-design dispatch — cut, lint-clean, INDEX'd, **held** (nothing relayed):
- Dispatch: `.relays/s6/s6-design/DESIGN-orchestrator-planner-20260706-180315.md`
- Its basis recordings (also reviewable): the operator ruling folded at `master-docs/master/README.md` (dashboard), `master-docs/master/RECONCILE.md` § s5 tail, `ROADMAP.md` Step-1 note — all 2026-07-06.

**What rides this review (beyond your standing checks):**
- **The in-step ruling's translation.** The operator ruled Step-1 open until the transport is fixed. Does the dispatch translate that faithfully — s6 in-step, the exit test moved onto the fixed conductor with the two live legs (§7-applies-s5's-registry; the F11 dogfood-replay regression) — without quietly expanding Step-1 scope anywhere else?
- **The ownership split + the three-way seam.** m-1 the parenting fork + anchor hygiene + waiver governance; m-7 liveness/ops (F5/F9/F15/F14/F16/F3); m-2 the single codec + the render↔validate family. Is any F-item mis-homed? Is the seam (parent computation touches all three) stated tightly enough that the m-7/m-2 halves stay composable with either fork branch?
- **The fork discipline.** The parenting model is flagged hard-to-reverse: decision-packet → master → operator GRILL → GRILL_LOCK → only then the m-1 lock. Is that the right gate geometry, and is the dispatch honest that conductor-computed-PARENT is a *candidate from the seeds, not pre-decided*?
- **The exit bar.** "Every F1–F17 item dispositioned {amendment | in-slice fix | wontfix-with-rationale}, no silent drops" — total-coverage as a hard design-phase exit. Right bar? Anything missing (e.g. should the credit-list protections — crash-atomicity/FROM/I-PH not open for redesign — be even harder-fenced)?
- **Cadence fit.** Design-amendment at the m-x pairs (the c6/s2-amend path) followed by a fresh build slice — versus dispatching one slice team to design+build. I chose the former (locked-contract touches belong to the owning pairs). Concur?

**On approve:** watchpoints fold as [VP-Wn]; the operator relays the dispatch to the three pairs; the design-amendment phase runs; the m-1 packet comes back for the grill.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s6/s6-design` — run below.
- The package is held: no pair handoff until your verdict. Baseline `frank/` `main @ 7e5c527` (tag `s5-close`), clean.

ACTIONS_GIT_REF: none — review request; no git action, no `frank/` edit.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main @ `7e5c527`, clean.
Next requested action: VP returns a verdict (approve / must-revise) on the s6-design package; on approve the operator hand-relays the dispatch to m-1/m-7/m-2.
