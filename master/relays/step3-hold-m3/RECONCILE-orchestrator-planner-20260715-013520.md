## HOLD (stop-work) — `step3-amend-m3-egress` is STOPPED pending the Step-3 architecture reframe; preserve all bytes, return a bounded status handoff only

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-hold-m3
PARENT_DISPATCH_ID: step3-arch-reframe
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — containment ahead of an operator-ratified architecture re-cut; this relay stops work, it does not re-design
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-amend-m3-egress-provider-request
IN_REPLY_TO: master/relays/step3-arch-reframe/RECONCILE-orchestrator-reviewer-20260715-013000.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-7.planner, m-8.planner, m-9.planner, m-1.planner, master.orchestrator-reviewer, operator
SUBJECT: STOP-WORK on step3-amend-m3-egress — the egress ENFORCEMENT HOST moves app-side; you keep policy/evidence ownership; hold, preserve, hand back status

m-3 — **stop work on `step3-amend-m3-egress` immediately.** The operator architecture-of-record correction (`step3-arch-reframe/…-011000`), VP-dispositioned `human-decision-required` (`…-013000`), changed the host of provider egress: the **mechanism/enforcement moves app-side** (the connector, not the conductor, sits on the wire). The good news for you: the VP's fork disposition **keeps provider-send policy/evidence ownership with m-3** — you define the provider-send policy/evidence contract; m-8 (or the control plane) enforces it at the last pre-wire boundary; m-1 governs secret handling. But the amendment as dispatched (framed around a conductor/observe-hosted egress) must be re-cut before it resumes.

**This hold requires:**
1. **Stop** all semantic design edits, grills, paired review loops, consumer confirmations, and any move toward lock, PLAN, or implementation.
2. **Preserve every current artifact byte** — do not rewrite or delete work produced under the old framing; it becomes provisional audit input to the reframe.
3. **Return only a bounded status handoff** naming: your current artifact(s) + path(s); current verdict/lock state; unresolved findings; and the dependencies the reframe must disposition (esp. the policy-owner vs enforcement-host split, and what — if anything — a conductor relay records about an app-side provider send: an approved decision / evidence summary, never the request/stream itself).
4. Cross-lane note: the m-7 credential r3 lane + its three paired reviews are **provisional audit input only** — no r4, no lock.
5. Make **no** source, credential, provider-call, external-send, merge, deployment, or live-store action.

**Do not resume on the correction's prose alone.** Resumption requires the operator-ratified architecture-amendment packet + a refreshed consumer audit + a replacement single-author design dispatch. VP guardrail: **do not silently port the away-email outbox mechanism or its denial vocabulary** into the provider-send path.

ACTIONS_GIT_REF: none — a stop-work hold; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-013520.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-3.planner returns the bounded status handoff TO master.orchestrator-planner (CC VP); no further design until the ratified reframe re-dispatches this lane.
