## HOLD (stop-work) — `step3-design-m-8` is STOPPED pending the Step-3 architecture reframe; preserve all bytes, return a bounded status handoff only

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-hold-m8
PARENT_DISPATCH_ID: step3-arch-reframe
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — containment ahead of an operator-ratified architecture re-cut; this relay stops work, it does not re-design
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-design-m-8-provider-adapters
IN_REPLY_TO: master/relays/step3-arch-reframe/RECONCILE-orchestrator-reviewer-20260715-013000.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, m-9.planner, m-1.planner, m-3.planner, m-4.planner, m-7.planner, master.orchestrator-reviewer, operator
SUBJECT: STOP-WORK on step3-design-m-8 — an operator architecture-of-record correction invalidated the dispatched framing (conductor is NOT the app hub / provider connectivity is app-side); hold, preserve, hand back status

m-8 — **stop work on `step3-design-m-8` immediately.** An operator architecture-of-record correction (`step3-arch-reframe/…-011000`), VP-dispositioned `human-decision-required` (`…-013000`), invalidated the framing your dispatch carried: the conductor is the governed **relay plane for stamped participants**, NOT the app's central hub; the LLM **connector and provider egress are app-side components**, not conductor-hosted; "add HTTPS to the conductor" was a category error. Your provider-adapter/connector design must be re-cut against the app-shell topology before it can resume — not continued on the dispatched picture.

**This hold requires:**
1. **Stop** all semantic design edits, grills, paired review loops, consumer confirmations, and any move toward lock, PLAN, or implementation.
2. **Preserve every current artifact byte** — do not rewrite or delete work produced under the old framing; it becomes provisional audit input to the reframe.
3. **Return only a bounded status handoff** naming: your current artifact(s) + path(s); current verdict/lock state; unresolved findings; and the dependencies the reframe must disposition (esp. where the connector process boundary, credential attach, and last-pre-wire enforcement land relative to the conductor).
4. Cross-lane note: the m-7 credential r3 lane + its three paired reviews are **provisional audit input only** — no r4, no lock.
5. Make **no** source, credential, provider-call, external-send, merge, deployment, or live-store action.

**Do not resume on the correction's prose alone.** Resumption requires the operator-ratified architecture-amendment packet + a refreshed consumer audit + a replacement single-author design dispatch. VP fork disposition relevant to you: provider request bytes/network send bypass the conductor; **m-8 is the likely connector actuator / last pre-wire enforcement host**; m-3 keeps provider-send policy/evidence ownership; m-1 governs secrets.

ACTIONS_GIT_REF: none — a stop-work hold; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-013500.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-8.planner returns the bounded status handoff TO master.orchestrator-planner (CC VP); no further design until the ratified reframe re-dispatches this lane.
