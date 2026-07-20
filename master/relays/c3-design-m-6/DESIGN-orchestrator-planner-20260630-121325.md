## Team m-6 — Human Surface & Scheduler: PROCEED TO DESIGN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c3-design-m-6
PARENT_DISPATCH_ID: c3-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — design surfaces operator-judgment items (away-mode policy, cadence, reply-grammar, B/C visibility, meeting boundary, delegation, §J config) + the m-1 confirm-or-gap; grill them
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-6.planner
CC: m-6.implementer, m-5.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
BUNDLE_ID: m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags product-semantics decisions. Not in scope: source/test edits, branches, commits, PRs, scaffolding, prototype code. Design-lock is the terminal — no implementation / PLAN.

Basis: your reconciled `c3-audit-m-6` (`054107`) is APPROVED by the VP (`c3-reconcile` verdict: approve). Design the **five surfaces** your audit produced (**promote-and-bind**, a thin local-first projection over locked m-1..m-4) — but treat the surfaced designs as **HYPOTHESES to PROVE and LOCK**, not proven facts.

Co-design with m-5 (the seam): coordinate in the shared COORD thread **`c3-design-m5-m6-coord`** (seeded — read it first; cite its current state). **Declare-before-bind (VP F2): do NOT bind surface/scheduler behavior to the human-mode vocabulary until m-5 declares it in the COORD.** You MAY lock your **neutral substrate** independently (the bucket-taxonomy mechanism, the ODB render, the park/wake state machine); the archetype-specific bindings (bucket-intensity / channels / resummon-class / interjection-affordances per human-mode) wait on m-5's declaration.

Design questions / grill agenda (your DQ-1..6 + OJ-1..7 — grill them; they may block LOCK, not dispatch):
1. **The m-1 confirm-or-gap [LOCK-BLOCKING — VP Finding 4].** The inbound away-mode verdict-token bridge. Draft the precise bounded question; I route it to m-1.planner as the **first conditional-upstream-contract-check** (in the COORD). **Do NOT design-lock the away-mode bridge** until m-1 answers (m-6-owns-bridge-over-m-1-crypto / m-1-owns-mint) **or** it is explicitly recorded as a blocker / human decision.
2. **Bucket cut (DQ-3)** — A/B/C/D + direction-explicit (only A+C reach the operator; B = orchestrator's; D = author's). Grill: does D belong outside the operator-surface taxonomy entirely?
3. **Meeting-surface depth (DQ-4)** — design the **route-to-meeting + re-observe-on-resume contract** now; **defer the attach mechanism to Step-3/4 runtime** (my recommendation — confirm at grill). Protect the collaboration lane as deliberately as the governance one (the export split).
4. **Resummon cadence + away defaults (DQ-5)** — the two-timer model; escalate the channel, never the verdict; **confirm none auto-resolves (J1).**
5. **OPERATOR-judgment items (OJ-1..7)** — away-mode destination policy (allowed inboxes, Gmail-first, may B/C ever mirror); cadence classes + max-resummon-before-block; reply grammar (exact-choice-ID vs alias vs structured); Bucket-B visibility (hidden/digest/live); meeting-vs-governance boundary; delegation (recommend DEFER, solo-operator); the §J `gate_category` map / protected-branch / away-opt-in / egress-whitelist **config surface** (m-6 designs the surface, operator owns the values). Grill these.

Hard proof requirements (prove, don't assert):
- PROVE **promote-and-bind**: each surface either **binds a locked mechanism** (m-1 addressing/stamp, m-2 `gate_category`/HUMAN_GATE, m-3 egress/observe/`record_integrity`, m-4 routing escalation) or designs the **away-mode bridge** the export left undesigned. No rebuild — the ODB content schema is **promoted from m-3** (render the surface + capture, not a new schema); the TUI is **deferred to Step-4** (design the contract).
- The **A/B/C/D bucket taxonomy** bound to locked mechanism (no-bucket-without-a-writer; direction-explicit); the **ODB** surface + bounded `agent_enum_pick` capture → operator-FROM verdict relay (J1 `hold_and_resummon`, refresh-before-resummon, never auto-approve); the **7-state park/wake machine** on the durable append-only store (`active → parked_waiting_human → resummon_due → replied_pending_validation → resumed` + `bounced_repair` + `egress_blocked`); the **opt-in, A-only, egress-gated away bridge** + the signed-token bridge (its LOCK gated on the m-1 confirm-or-gap); the **governance-vs-collaboration split** + re-observe-on-resume; the **interjection host** (Claude-Code three-mechanism, not jcode's side-panel).
- **Local-first / egress fail-closed**: every Step-2 mechanism is local-only over the relay store; the sole external send (the away bridge) is egress-gated by m-3 §7 (fail-closed). Mechanism builds Step 2; full email-client UX builds Step 4 — design the contracts now.

Guardrails (VP):
- **No surface value lock outside this DESIGN grill.**
- **Do NOT reopen locked m-1..m-4** — design against them. The **only** sanctioned upstream touch is the **bounded m-1 confirm-or-gap** (DESIGN-phase, orchestrator-routed).
- **Declare-before-bind** — the human-mode vocabulary from m-5 (the COORD).
- **The away-mode-bridge LOCK is gated on the m-1 confirm-or-gap** (do not lock it until resolved or recorded as a blocker/human decision).

Boundary contract — name the consumed fields before lock:
- Consumes m-5 (the human-mode two-layer vocabulary + the interjection surface / sensor archetype — the COORD); m-1 addressing/stamp + §J; m-2 `gate_category`/HUMAN_GATE; m-3 egress/observe/`record_integrity`; m-4 routing.
- No downstream consumer (you are the terminal design domain; m-7..m-12 runtime/product are a later step). Flag any binding point with no locked mechanism behind it.

Out of scope: m-5 archetype internals (sibling `c3-design-m-5`); the locked m-1..m-4; the full email-client UX **build** (Step-4 — design the contract, don't build); any code.

Relay hygiene: keep the pair-thread DISPATCH_ID `c3-design-m-6`; address the design-review request TO m-6.implementer (not the orchestrator).

Deliverable: a design doc (Superpowers brainstorming + design-grill), recorded as DESIGN_DOC_ID `c3-design-m-6-human-surface-scheduler` under `master/domains/m-6-human-surface-scheduler/design/`, containing — the five proven surface designs (bucket taxonomy / ODB / park-wake+away-bridge / email+meeting / interjection host); the COORD-thread seam resolution (the human-mode binding + the m-1 confirm-or-gap result) cited; operator decisions/defaults folded into a GRILL_LOCK; the named consumer boundary contract; open questions (incl. any unresolved m-1 confirm-or-gap recorded as a blocker/human decision, NOT silently locked). Then send the design-review request TO m-6.implementer, and report design-complete to the orchestrator for the c3 lock (after the COORD reconcile + the m-1 confirm-or-gap resolution). Do not self-advance to PLAN.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
