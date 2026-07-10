## COORD (seed) — c3 m-5↔m-6 seam thread: human-mode declare-before-bind · interjection host · m-1 confirm-or-gap

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c3-design-m5-m6-coord
PARENT_DISPATCH_ID: c3-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — coordination thread; resolves at DESIGN, before any c3 lock
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-5.planner, m-6.planner
CC: m-5.implementer, m-6.implementer, master.orchestrator-reviewer, operator

m-5, m-6 — this is the **single shared COORD thread** for c3 (VP-approved, `c3-reconcile/RECONCILE-orchestrator-reviewer-20260630-121138` approved-action). Opened early so **both c3 design docs cite its current state** as they develop. Coordinate the seam here; do not side-lock it inside either design doc. **Reconcile this thread before any c3 lock.** The m-5↔m-6 ordering is **declare-before-bind (VP F2):** m-5 declares vocabulary/semantics; m-6 binds surface behavior only after.

**Agenda (resolve here; cite current state in both designs):**

**(a) Seam A — human-mode declare-before-bind [the load-bearing item].** m-5's reconcile converged the structure: a **two-layer** vocabulary — `human_mode` **posture** {interactive / away / unattended} ⊥ `surface_intent` **delivery-class**. m-5 **DECLARES** both value-sets; m-6 **BINDS** surface/scheduler behavior (bucket-intensity / channels / resummon-class / interjection-affordances / meeting-enabled) to the (posture × surface_intent) pair — **only after m-5 declares.** Open deltas to resolve here:
- (i) the exact `surface_intent` value-set;
- (ii) whether `operator_gate` / `hold_and_resummon` are m-5-declared `surface_intent` values OR references to the **locked** A-bucket `gate_category` / J1 mechanisms (m-5's reconcile flagged this — do not duplicate locked mechanism as vocabulary);
- (iii) `away_bridge_eligible` as a boolean capability flag (m-5.implementer) vs a mode value;
- (iv) whether `human_mode` + `surface_intent` are one m-2 field or two (no m-2 micro-fold over concrete values — C2.4).
m-6 may lock its **neutral substrate** (bucket-taxonomy mechanism, ODB render, park/wake state machine) independently; the archetype-specific bindings wait on m-5's declaration.

**(b) Seam B — interjection host.** m-6 owns the **surface** (the steer / side-question / interrupt composer + choice); m-5 owns the **side-question SENSOR archetype** (full design); m-4 **routes** it (`fast-cheap`, locked); the **runtime** owns boundary-injection / soft-cancel / fork (later-step build). Resolve: the spawn+answer contract (m-5 declares the sensor ceiling + "answer lands on a separate, non-lane surface, lane not interrupted"; m-6 binds the rendering + the three-way choice), and **which archetypes accept interjection** (a 1-turn sensor is not steered; a long-lived implementer/orchestrator is).

**(c) m-1 confirm-or-gap — the inbound away-mode verdict-token bridge [LOCK-BLOCKING — VP Finding 4].** m-6's away-mode bridge converts **untrusted SMTP-inbound → a trusted operator-channel verdict record.** Reconciled position: anchor on m-1's **existing** operator-channel stamp **+ an m-6-owned signed one-time per-`(decision,seat,choice)` token bridge** (POST-not-GET; replay-nonce ≠ validity-window). The **bounded question to m-1:** does m-1's `mint_seat`/channel-stamp model **extend** to minting/verifying the inbound one-time verdict-token (m-1 owns the mint), **or** is it an m-6 construct over m-1-owned crypto + the existing channel (m-6 owns the bridge)?
- This is the **first conditional-upstream-contract-check** (VP F3). It fires as a **bounded confirm-or-gap relay to m-1.planner**, orchestrator-routed, when m-6's design reaches the away-mode binding (m-6 drafts the precise question; I route it + re-engage m-1). **The away-mode bridge MUST NOT design-lock** until m-1 answers (confirm m-6-owns-bridge / m-1-owns-mint) **or** it is explicitly recorded as a blocker / human decision. Nothing m-1 must change to *open* DESIGN.

**Ownership boundary (do not cross):** m-5 owns the archetype + human-mode vocabulary + sensor semantics; m-6 owns the surface/scheduler binding + the away-mode bridge; m-4 routing is locked; m-1 owns the TCB/identity (the confirm-or-gap). Neither reopens locked c1/c2 except via the bounded m-1 confirm-or-gap.

Exchange COORD relays in this DISPATCH_ID (`c3-design-m5-m6-coord`). Reach a reconciled seam statement (all three items) before either pair reports design-complete; the orchestrator folds it at the c3 lock.

ACTIONS_GIT_REF: none — coordination seed relay; docs workspace, no code/source/pcode edits; cwd is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
