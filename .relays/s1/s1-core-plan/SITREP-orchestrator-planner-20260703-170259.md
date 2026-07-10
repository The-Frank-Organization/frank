## COORD — ASK-1 RULING: root README.md is IN-FENCE (fence amended explicitly); external-gate packets issued; hold until all four approves land

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-core-plan
PARENT_DISPATCH_ID: s1-core-design
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the four external approves must land as relays in .relays/s1/ before SCOPE_DIFF-then-dispatch
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: s1-core.implementer, s1.orchestrator-reviewer, operator
IN_REPLY_TO: s1-core-plan/SITREP-planner-20260703-163256.md
SUBJECT: README fence ruling (in-fence — the fence wording is hereby amended to include it); plan-gate + both fidelity packets written and handed to the operator; your hold posture is correct

**ASK-1 RULING — root `README.md` is IN-FENCE.** The PROCEED-TO-PLAN fence wording ("new Go
source/test/fixture/build files") did not anticipate the honesty surface; the file serves
ratification condition 2 (claim honesty via Task 12 README spec + SWEEP) and the m-7 §16
sweep applies to it — it is exactly the kind of seat-/user-facing surface the slice must
control. The scope list in `s1-core-plan/PLAN-orchestrator-planner-20260703-153356.md` is
amended by this relay to read, additionally:
- repo-root `README.md` (the S1 honesty/claim surface; Task 12; SWEEP-covered)
Whether an out-of-fence file is necessary is the orchestrator's call — this one is, and it is
now in-fence. The planned fallback (move under docs/) is NOT needed; keep the authoritative
file-structure block as committed at plan r2 (main@a24bf57). Cite this relay as the fence
evidence in your SCOPE_DIFF row for README.md.

**ASK-2 — external-gate packets issued** (all three written this session, operator-carried):
- Plan gate → m-7.planner + master.orchestrator-reviewer:
  .relays/s1/s1-plan-gate/SITREP-orchestrator-planner-20260703-170259.md
- m-1 fidelity → m-1.implementer:
  .relays/s1/s1-fidelity-m1/SITREP-orchestrator-planner-20260703-170259.md
- m-2 fidelity → m-2.implementer:
  .relays/s1/s1-fidelity-m2/SITREP-orchestrator-planner-20260703-170259.md

Your hold posture is correct and unchanged: SCOPE_DIFF runs only after all four approves are
relays in .relays/s1/; delegated dispatch only per the full condition set in the
PROCEED-TO-PLAN (:24-25). No further orchestrator input is owed before the gates return.

ACTIONS_GIT_REF: wrote this relay + the three gate-packet relays + INDEX rows under .relays/s1/ (gitignored substrate); reconciliation entry committed on main (see git log; commit follows this batch)
FINAL_GIT_STATUS_SHORT: none — clean tree (verified before this batch; the only tracked change this turn is the RECONCILE.md entry, committed separately)
Next requested action: hold for the four external approves; on their arrival run SCOPE_DIFF (README.md row cites this ruling) and proceed per the delegation conditions.
