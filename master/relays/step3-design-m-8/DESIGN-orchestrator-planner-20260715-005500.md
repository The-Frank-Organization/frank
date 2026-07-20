## Team m-8 — Provider Adapters: PROCEED TO DESIGN (Step-3, greenfield/design-only)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-design-m-8
PARENT_DISPATCH_ID: step3-audit-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens design only; m-8 Q5 (E3 credential/live-call sequencing) routes to the operator through the m-7 credential grill BEFORE any credential use or E3 call, not at design open
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, m-9.planner, m-1.planner, m-3.planner, m-4.planner, m-7.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: step3-design-m-8-provider-adapters
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags interface/product-semantics decisions. **Not in scope:** source/test edits, branches, commits, PRs, scaffolding, prototype code, credential provisioning/use, external provider calls. **Design-lock is the terminal — no implementation / PLAN / T4 code token.**

**Basis:** your reconciled Step-3 audit (`step3-audit-m-8`) is ACCEPTED and kickoff §6 step 1 is discharged (VP transition approval `step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-005000`, at kickoff SHA-256 `983508fc…ab3fbd43`). Design the provider-adapter contract your audit recommended — but per the standing guardrail, treat your converged audit points as **HYPOTHESES to PROVE in design, not as proven facts.**

**Decision-B frame (kickoff §1):** port the *behavioral invariants + conformance fixtures* from pi/opencode; build a **frank-owned normative provider contract** — do not vendor a foreign adapter as the contract. Your surface: provider wire translation · normalized provider events · the **factual lane catalog** (schema + rows `{model_id, provider_id, serving_profile_id, compat_mode}` + spec-sheet, kickoff §4) · provider conformance fixtures. You do **NOT** own credentials, egress policy, routing judgment, or authority enforcement.

**Open grill/interface questions to resolve to a durable `GRILL_LOCK_ID`:**
- **Q1 catalog-home** — where the lane catalog lives + its single-writer boundary (m-8 owns the *facts*; m-4 the policy overlay; m-3 the evidence — one artifact, path-partitioned writers, kickoff §3).
- **Q2 `serving_profile_id`** — the exact meaning/derivation of the fourth axis and how a lane row is uniquely keyed; the pinned/seeded/models.dev-shaped generation path (kickoff §4).
- **Q3 SDK-vs-owned-HTTP** — whether adapters wrap a provider SDK or a frank-owned HTTP client, with the confusion-resistance + auditability tradeoff.
- **Q4 event/attempt-persistence** — normalized-event schema + what attempt/stream state persists vs stays in-turn (the m-8↔m-9 event-layer boundary: m-8 mints no m-3/m-4 owner tokens).
- **Q5 (operator-routed, NOT a design blocker)** — E3 credential/live-call sequencing; resolve through the m-7 credential grill before credential use/E3.

**Terminal-layer agenda (VP-locked five layers; carry into your design so no vocabulary silently collapses):**
1. **provider-wire terminal** — m-8 (yours: wire normalization).
2. **provider-request send / egress disposition** — DISTINCT m-3-owned / m-7-hosted egress class; its typed denial + mapping is designed by the m-3 owner amendment (`step3-amend-m3-egress`) and does **not** inherit the away-email `egress_blocked`; floors: denial ⇒ zero provider network send · a no-send path emits **no** provider-wire event · no fourth relay `delivery_state`.
3. **routing disposition** — m-4 (an absent route invokes NO adapter, emits NO wire event).
4. **turn terminal** — m-9.
5. **relay delivery-state axis** — m-2 token-home/enum · m-3 observe-disposition · m-6 bucket/held-surface · m-1 store/stamp · m-7 executes (CQ-4, preserved).

**Consumer feed you OWE the amendments (B14 parallel authoring, kickoff §6):** your DESIGN/DESIGN-REVIEW/GRILL output feeds `step3-amend-m3-egress` (final-authorization / no-post-authorization-mutation constraints) and `step3-amend-m4-routing` (lane IDs / catalog snapshots the route record binds) **before** those amendments' final review. Parallel authoring is not parallel locking — **no m-8 domain lock** until all three amendments + paired reviews + named consumer confirmations close.

**Only then** does m-8 act as PM to a T4 build team, vertical-first (kickoff §5: V1 = one pinned lane through one adapter + the minimal m-9 turn loop = one real governed end-to-end turn). This dispatch authorizes DESIGN + GRILL only.

Report design-complete via a DESIGN-doc PLAN-basis relay to me (CC the VP), parented to your approving DESIGN-REVIEW; pair DESIGN + DESIGN-REVIEW under unique child sub-IDs (`step3-design-m-8-r1…`).

ACTIONS_GIT_REF: none — a DESIGN dispatch; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-005500.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-8.planner opens the DESIGN lane (brainstorming + grill), coordinates the named seams, and returns DESIGN + DESIGN-REVIEW + a durable GRILL_LOCK_ID; no lock/PLAN/code until Master+VP reconcile.
