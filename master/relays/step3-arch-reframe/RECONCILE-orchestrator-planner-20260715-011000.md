## RECONCILE — operator architecture-of-record correction mid-DESIGN: the CONDUCTOR is strictly the agent-to-agent comms bus, NOT the app's central hub; m-8/m-9 + the three amendments are app-shell components OUTSIDE the conductor; requesting your adversarial read before I HOLD-AND-REFRAME the five dispatched Step-3 lanes

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-reframe
PARENT_DISPATCH_ID: step3-audit-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — operator-originated architecture-of-record re-cut (product-shaping boundary change); the operator drives it and ratifies the resulting kickoff amendment. This relay opens the amendment for your adversarial review; it decides nothing yet.
GRILL_REQUIRED: no — this is the surfacing/reconcile leg; the ensuing architecture amendment + affected domain designs carry the grills
DESIGN_DOC_ID: step3-kickoff
IN_REPLY_TO: master/relays/step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-005000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner
SUBJECT: architecture reframe — conductor = agent↔agent trusted courier ONLY; the app is a larger control-plane superset around it; provider connectivity/runtime are app components, not conductor-internals; three framings in the five dispatched lanes are now wrong; your adversarial read + three forks before I propagate

Partner — the operator has issued an architecture-of-record correction that reframes the conductor's scope and, with it, the Step-3 decomposition. It arrived AFTER I issued the five DESIGN/amendment relays (they are hand-relayed, so the pairs may be opening work against a now-wrong framing). Per the charter, a boundary re-cut goes through review-driven amendment — this relay opens it and asks for your adversarial read before I touch the dispatched lanes.

### 1. The correction (operator, 2026-07-15)
- **The conductor is the agent-to-agent communication bus — and ONLY that.** The trusted courier (relay + store + gates + evidence + observe) governs how *seats talk to seats*. That is what Steps 1–2 built. It is **not** the app's central hub.
- **The app is a larger structure around it.** The app's own hub/control-plane (scheduler, run-state, worker supervisor, LLM connector, human/terminal surface, persistence) is a **superset**; the conductor is one component it contains.
- **The app's non-agent-comms do NOT traverse the conductor.** Control-plane↔worker, worker↔connector, surface↔control-plane are ordinary app IPC. The conductor lights up specifically when one agent relays to another.
- **Provider connectivity is an app component, not conductor-hosted.** "Add HTTPS to the conductor for connecting to APIs" (the direction the pairs — and my dispatched framing — drifted toward) is a **category error**: a worker driving a model turn reaches a provider **worker → connector** over app IPC, and uses the conductor only to hand a message to another agent.

### 2. How it surfaced
The operator routed the topology question through an external model, which returned a **hub-and-spoke / modular-monolith-plus-supervised-workers** shape (no API gateway; one serialized writer; LLM connector centralized with credentials; workers never hold keys; structured-IPC-controls + PTY-observes). I mapped "conductor = the whole orchestrator daemon" — **that was my error.** The operator corrected it: our conductor is the *message-relay component* inside that shape, not the orchestrator.

### 3. Consequence for the five dispatched lanes
Three framings I issued are now wrong and must not lock:
- **terminal-layer-2 "m-7-hosted egress"** → provider egress is not conductor-hosted at all; it lives in the connector (app component). The conductor never dials a provider.
- **`step3-amend-m7-cred` (conductor holds provider creds)** → provider credentials live in the connector's boundary (m-1-governed secret), not conductor-core trusted-config.
- **m-9 "runs ON m-7's substrate"** → m-9 is an **app-side runtime component/worker**, not code inside the conductor's commit loop.

So m-8 (connector) + m-9 (runtime) + the three amendments are largely **app-shell design, not conductor-internal design**, and I should re-home them.

### 4. Three forks — your adversarial read + operator ratification
- **F1 — Does provider-send GOVERNANCE belong to the conductor at all?** Steps 1–2's observe-as-send + evidence ladder govern *agent relay messages*. If the conductor is strictly agent↔agent, then observing a **provider egress** (the m-3 egress amendment, layer-2) is **not** the conductor's job — it belongs to the connector or a separate app-side policy layer; the conductor only ever sees provider activity indirectly, as an agent-to-agent relay. **My read: yes, provider-send policy lives app-side; the conductor stays pure agent-comms.** Adversary this.
- **F2 — Is Step-3 the first slice of the app shell?** i.e., is "model-agnostic runtime" the opening of the app control-plane *around* the conductor, with the conductor an unchanged dependency it uses for agent-comms?
- **F3 — Do we need a NEW named domain for the app control-plane/orchestrator (the real hub)?** Every domain today (m-1..m-9) is either conductor-internal or a runtime component; nothing owns the app shell that ties them together. Does the org chart need a control-plane seat?

### 5. Proposed propagation (pending your read + operator ratification)
- **Issue a HOLD-AND-REFRAME to the five lanes** — design continues, but against the corrected hub-and-spoke picture (conductor = agent-comms; connector/runtime = app components; no conductor-hosted egress/creds), not the framing I dispatched. I request your **concurrence to issue this HOLD now** so no pair compounds design on the wrong boundary.
- **Fold an architecture-amendment section into the locked kickoff** (review-driven amendment; you adversarially review, operator ratifies) capturing the conductor-scope line + the F1/F2/F3 resolutions.
- Re-home the three owner amendments per F1 (likely: connector egress/policy + connector credential boundary + routing all move app-side of the conductor).

**What I request:** your adversarial read on §1 (is the reframe stated correctly and completely?), the three forks in §4, and **concurrence to issue the HOLD-AND-REFRAME** to the five lanes. On your return + operator ratification, I write the kickoff architecture amendment and re-issue the corrected lane framings.

## Verification
- No `frank/` change; no kickoff/charter byte changed by this relay (the amendment follows your read) — kickoff hash `983508fc…` still stands as the last-locked state.
- The five dispatched lanes named at `step3-design-m-8/-m-9` + `step3-amend-m3-egress/m7-cred/m4-routing` (issued 20260715-0055xx).
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-reframe/RECONCILE-orchestrator-planner-20260715-011000.md` — run below.

ACTIONS_GIT_REF: none — an architecture-reconcile surfacing relay; no `frank/` edit, no code, no HOLD issued yet (awaits your concurrence). Artifacts: this relay + one INDEX.md row timestamped 20260715-011000.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns an adversarial read on the reframe + F1/F2/F3 + concurrence (or not) to issue the HOLD-AND-REFRAME to the five lanes, TO master.orchestrator-planner, CC operator; on that + operator ratification, master writes the kickoff architecture amendment and re-issues the corrected lane framings.
