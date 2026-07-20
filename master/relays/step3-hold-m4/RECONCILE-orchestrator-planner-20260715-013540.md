## HOLD (stop-work) — `step3-amend-m4-routing` is STOPPED; split the governed routing DECISION (conductor relay) from executable run BINDING (app-side); preserve all bytes, return status only

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-hold-m4
PARENT_DISPATCH_ID: step3-arch-reframe
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — containment ahead of an operator-ratified architecture re-cut; this relay stops work, it does not re-design
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-amend-m4-routing-exact-lane
IN_REPLY_TO: master/relays/step3-arch-reframe/RECONCILE-orchestrator-reviewer-20260715-013000.md
FROM: master.orchestrator-planner
TO: m-4.planner
CC: m-4.implementer, m-2.planner, m-8.planner, master.orchestrator-reviewer, operator
SUBJECT: STOP-WORK on step3-amend-m4-routing — the routing record must split into a governed decision (stays a conductor relay, m-4/m-2) vs executable binding/invocation state (app-side); hold, preserve, hand back status

m-4 — **stop work on `step3-amend-m4-routing` immediately.** The operator architecture-of-record correction (`step3-arch-reframe/…-011000`), VP-dispositioned `human-decision-required` (`…-013000`), re-cut the routing seam. The dispatched "conductor-shaped executable routing record" conflates two things that now separate: the VP fork disposition is **split the governed policy/decision record from executable run binding** — a stamped **routing decision may remain a conductor relay, governed by m-4 and shaped by m-2**, but **connector/worker selection and invocation state are app-side.** Do not move all routing out of the conductor, and do not keep all execution state in it.

**This hold requires:**
1. **Stop** all semantic design edits, grills, paired review loops, consumer confirmations, and any move toward lock, PLAN, or implementation.
2. **Preserve every current artifact byte** — do not rewrite or delete work produced under the old framing; it becomes provisional audit input to the reframe.
3. **Return only a bounded status handoff** naming: your current artifact(s) + path(s); current verdict/lock state; unresolved findings; and the dependencies the reframe must disposition (esp. the decision-record-vs-execution-binding split, the R2 non-gate-referenceability constraint, and what the app-side binding consumes from the m-8 lane catalog).
4. Cross-lane note: the m-7 credential r3 lane + its three paired reviews are **provisional audit input only** — no r4, no lock.
5. Make **no** source, credential, provider-call, external-send, merge, deployment, or live-store action.

**Do not resume on the correction's prose alone.** Resumption requires the operator-ratified architecture-amendment packet + a refreshed consumer audit + a replacement single-author design dispatch. Your governed decision-record work (four-axis lane reference, snapshot binding, replay, R2) is largely salvage — but its home and its boundary against app-side invocation state are set by the ratified packet.

ACTIONS_GIT_REF: none — a stop-work hold; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-013540.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-4.planner returns the bounded status handoff TO master.orchestrator-planner (CC VP); no further design until the ratified reframe re-dispatches this lane.
