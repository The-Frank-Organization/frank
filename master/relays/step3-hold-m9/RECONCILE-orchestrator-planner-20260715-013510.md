## HOLD (stop-work) — `step3-design-m-9` is STOPPED pending the Step-3 architecture reframe; preserve all bytes, return a bounded status handoff only

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-hold-m9
PARENT_DISPATCH_ID: step3-arch-reframe
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — containment ahead of an operator-ratified architecture re-cut; this relay stops work, it does not re-design
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-design-m-9-model-runtime
IN_REPLY_TO: master/relays/step3-arch-reframe/RECONCILE-orchestrator-reviewer-20260715-013000.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, master.orchestrator-reviewer, operator
SUBJECT: STOP-WORK on step3-design-m-9 — the "m-9 runs ON m-7's substrate" placement is invalidated; m-9 is an app-side runtime component, not conductor-internal; hold, preserve, hand back status

m-9 — **stop work on `step3-design-m-9` immediately.** The operator architecture-of-record correction (`step3-arch-reframe/…-011000`), VP-dispositioned `human-decision-required` (`…-013000`), invalidated a load-bearing assumption in your dispatch: **"m-9 runs ON m-7's substrate" is wrong** — the model-turn runtime is an **app-side component/worker**, not code inside the conductor's commit loop. m-9 uses the conductor only when a seat relays a governed message to another seat; its model-turn traffic, tool execution, and provider bytes do not transit the conductor. Your turn/session/context design must be re-cut against the app-shell topology before it can resume.

**This hold requires:**
1. **Stop** all semantic design edits, grills, paired review loops, consumer confirmations, and any move toward lock, PLAN, or implementation.
2. **Preserve every current artifact byte** — do not rewrite or delete work produced under the old framing; it becomes provisional audit input to the reframe.
3. **Return only a bounded status handoff** naming: your current artifact(s) + path(s); current verdict/lock state; unresolved findings; and the dependencies the reframe must disposition (esp. the m-9↔control-plane supervisor boundary, the tool-execution authorization path vs the conductor's `{submit,project,read}` guardrail, and the Q4 retry/idempotency seam vs the app-side connector's final-wire authorization).
4. Cross-lane note: the m-7 credential r3 lane + its three paired reviews are **provisional audit input only** — no r4, no lock.
5. Make **no** source, credential, provider-call, external-send, merge, deployment, or live-store action.

**Do not resume on the correction's prose alone.** Resumption requires the operator-ratified architecture-amendment packet + a refreshed consumer audit + a replacement single-author design dispatch. VP note: a new app-control-plane/supervisor domain (provisionally m-10) is recommended to own worker lifecycle/supervision — m-9 owns the turn runtime and must NOT also become its own supervisor.

ACTIONS_GIT_REF: none — a stop-work hold; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-013510.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the bounded status handoff TO master.orchestrator-planner (CC VP); no further design until the ratified reframe re-dispatches this lane.
