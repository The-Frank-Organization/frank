## Team m-9 — Model Runtime: PROCEED TO DESIGN (Step-3, greenfield/design-only)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-design-m-9
PARENT_DISPATCH_ID: step3-audit-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens design only; no credential/external-call/authority-ceiling decision is made here
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: step3-design-m-9-model-runtime
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags interface/product-semantics decisions. **Not in scope:** source/test edits, branches, commits, PRs, scaffolding, prototype code, credential handling, external provider calls, authority-ceiling changes. **Design-lock is the terminal — no implementation / PLAN / T4 code token.**

**Basis:** your reconciled Step-3 audit (`c7-audit-m-9`) is ACCEPTED and kickoff §6 step 1 is discharged (VP transition approval `step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-005000`, at kickoff SHA-256 `983508fc…ab3fbd43`). Design the model-turn/session/context runtime your audit recommended — treat your converged audit points as **HYPOTHESES to PROVE in design, not proven facts.**

**Scope frame (kickoff §2):** the turn/session/context state machine that **drives a model through m-8's adapters, governed** — the "frank runs a real turn" layer. **Governed tool-execution requests:** you parse tool calls and *request* execution; a parsed tool call stays **INERT** until the existing trusted authority/tool-exposure path authorizes it. You do **NOT** re-own m-7's substrate (process/commit loop/recovery/trusted-config/seat-guardrail), m-5's authority ceiling, or credentials/endpoints/egress/routing.

**Open grill/interface questions to resolve to a durable `GRILL_LOCK_ID`:**
- **Q1 process-placement** — where the runtime loop sits *on* m-7's process/commit loop without reimplementing it (the m-9↔m-7 host boundary).
- **Q2 within-substrate schema** — the session/context/turn state shape and what persists across turns vs stays in-turn.
- **Q3 observation-granularity** — observe-as-send over the model's streamed output: unit/cadence of observation and how it composes with the gate.
- **Q4 retry/idempotency relative to final-wire authorization** — turn-terminal semantics vs the m-3 provider-request egress authorization; **you OWE this into the `step3-amend-m3-egress` consumer packet** (retry/idempotency + no-post-authorization-mutation).
- **Q6 m-7/m-6 seam** — where a turn's human-surface interaction (m-6) meets the runtime on m-7; m-6.planner is CC'd to receive this source packet.

**Terminal-layer agenda (VP-locked five layers; carry into your design):**
1. provider-wire terminal — m-8.
2. provider-request send / egress disposition — DISTINCT m-3-owned / m-7-hosted egress class; typed denial + mapping owner-designed (`step3-amend-m3-egress`), does not inherit away-email `egress_blocked`; floors: zero send · no wire event on no-send · no fourth `delivery_state`.
3. routing disposition — m-4 (absent route invokes NO adapter).
4. **turn terminal** — m-9 (yours: turn semantics).
5. relay delivery-state axis — m-2 token-home/enum · m-3 observe-disposition · m-6 bucket/held-surface · m-1 store/stamp · m-7 executes (CQ-4, preserved).

**Consumer reviews you OWE (B14 parallel authoring, kickoff §6):** consumer-review the named seams — the authority path against m-5/m-7, and the provider-request egress point against m-3/m-7. Your DESIGN/DESIGN-REVIEW/GRILL output feeds `step3-amend-m3-egress` before its final review. Parallel authoring is not parallel locking — **no m-9 domain lock** until all three amendments + paired reviews + named consumer confirmations close.

**Only then** does m-9 act as PM to a T4 build team, vertical-first (kickoff §5: V1 = one pinned lane through one adapter + the minimal m-9 turn loop = one real governed end-to-end turn). This dispatch authorizes DESIGN + GRILL only.

Report design-complete via a DESIGN-doc PLAN-basis relay to me (CC the VP), parented to your approving DESIGN-REVIEW; pair DESIGN + DESIGN-REVIEW under unique child sub-IDs (`step3-design-m-9-r1…`).

ACTIONS_GIT_REF: none — a DESIGN dispatch; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-005510.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner opens the DESIGN lane (brainstorming + grill), coordinates the m-7/m-6/m-5/m-3/m-8 seams, and returns DESIGN + DESIGN-REVIEW + a durable GRILL_LOCK_ID; no lock/PLAN/code until Master+VP reconcile.
