## PLAN-REVIEW request — VP pre-handoff gate on the s6 BUILD package (dispatch + boot held; the slice Step-1 closes on)

ROLE: Orchestrator Planner
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-dispatch
PARENT_DISPATCH_ID: s6-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — pre-handoff VP review (the standing cadence); the operator holds the handoff until your verdict AND the m-7 revert window passing quiet
IN_REPLY_TO: .relays/s6/s6-dispatch/PLAN-orchestrator-planner-20260706-221040.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: review the s6 build dispatch + boot before handoff — your co-sign's [VP-W1..W3] folded (hold-on-revert · FX-B1g explicit · seven rows no marker); the exit gate carries the step-exit test whole (incl. the operator §7 leg + the F11 live replay + the live boot); the no-perf-work fence; point-not-restate discipline

**The ask.** Pre-handoff review of the s6 build package — cut, lint-clean, INDEX'd, **held**:
- Dispatch: `.relays/s6/s6-dispatch/PLAN-orchestrator-planner-20260706-221040.md`
- Boot: `.relays/s6/boot/s6-boot-orchestrator-planner/SITREP-orchestrator-planner-20260706-221040.md`

**What rides this review (beyond your standing checks):**
- **Your three co-sign watchpoints, folded:** [VP-W1] the dispatch self-holds until the m-7 revert window passes quiet (stated in its plan-gate section — the operator relays only after); [VP-W2] FX-B1g appears explicitly in the exit gate with its scope; [VP-W3] the registry pass is fenced to EXACTLY the seven named rows, no-marker stated in-scope. Confirm the foldings carry your intent.
- **Point-not-restate:** the dispatch treats the r3 set + the four constituent docs as the spec and points at them. Check I haven't paraphrased any contract in a way that could drift from the co-signed text (the s5 lesson: restated prose is where conflicts hide).
- **The step-exit test as the exit gate:** the ROADMAP:83-85 legs + the operator §7 leg + the F11 live replay + the live boot (mint→wire→boot→derived-active→roster) — is the test complete against the operator's in-step ruling, and are the operator-authored legs clearly marked as the operator's (never the team's to simulate — the s4 honesty line)?
- **The OUT fence:** no perf work (the latency exoneration), no Step-2 pre-work, no dogfood-in-slice, amendment-path-only for locked touches. Anything missing?
- **The onboarding bar:** read-the-ledger-before-the-spec ordering in the boot (the WHY before the WHAT) — concur it's load-bearing, not decoration?

**On approve** (+ the revert window quiet): the operator relays the boot then the dispatch to a fresh session; s6 onboards, plans under F2, builds; master's independent battery at the close tip; the operator's integration gate; **the Step-1 close fold.**

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s6/s6-dispatch` — run below.
- Package held; baseline `frank/` `main @ 7e5c527` (tag `s5-close`), clean; the §C4 pointer + dashboard landed this pass.

ACTIONS_GIT_REF: none — review request; no git action, no `frank/` edit.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main @ `7e5c527`, clean.
Next requested action: VP returns approve / must-revise; on approve + the quiet revert window the operator boots the s6 team.
