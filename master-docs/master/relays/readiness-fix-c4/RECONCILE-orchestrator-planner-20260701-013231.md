## RECONCILE — CTO re-verify: Cluster 4 fully reconciled (4a+4b+4c, all four owner-confirms in) → recommend CLOSED (VP closure co-sign requested); two pre-PLAN items tracked

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: readiness-fix-c4
PARENT_DISPATCH_ID: readiness-fix-c4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — CTO re-verify; recommends closure for VP co-sign; operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner — Cluster 4 is fully reconciled. All four owner-confirms are in, m-2 applied the posture ruling, and I
re-verified the folded schema **line-by-line against source**. Recommending Cluster 4 **CLOSED**, pending your closure
co-sign. This is the **last MUST item**.

**Owner-confirms (all CONFIRM):**
- **m-3** (`000930`) — `record_integrity` authorship-faithful; m-1 "m-2-declared set" met.
- **m-4** (`013000`) — record `:203` un-extended; mirror now follows it; re-verified m-2's §17.3 matches `:203`. No open finding.
- **m-5** (`012910`) — **option 1, no c3-lock amendment**; owns its `235944` imprecision; effective posture derives from
  `template_ref`-else-`seat_archetype` (both on the record).
- **m-6** (`012906`) — `(posture × surface_intent)` reads cleanly via `seat_archetype`; standalone-field removal correct.

**CTO re-verify (E1 — I read `m-2 …forms-determinism…§17.3:292` + `§17.6`):**
- **4b mirror faithful to `m-4 …:200-210`:** per-row `routing_assignments` = `{seat, role, task_tag, declared_bucket,
  chosen_model, pin_mode, declared_deviated, seat_archetype, authority_ceiling}` — **matches `:203` exactly**;
  `deviation_reason_code` (`:297`, `required_when any(routing_assignments.declared_deviated==true)`) + reserved-shape
  `constraints`/`template_ref` present. **No standalone posture field** — `:292` states posture "rides `seat_archetype`
  (m-5 c3 lock `:142`); m-6 resolves posture by reading the recorded `seat_archetype`." ✓
- **4c homes declared:** `record_integrity`/`surface_intent` as computed slots (§17.6); posture rides `seat_archetype`;
  all **non-gate-referenceable** (audit/delivery outputs, R2-consistent). ✓
- **4a intact:** the R2 trigger + grammar-enforcement + `declared_deviated`/`deviated_observed` declare/observe split
  are unchanged. ✓
- **Invariants held:** R2 (model never a gate input; allowlist excludes model-identity fields); m-4 record-contract
  unchanged; no `ARCHITECTURE` change; `seat_archetype` opaque; no c2/c3 reopen (this *applies* the c3 lock).

**Recommendation: Cluster 4 → CLOSED.** All of 4a/4b/4c reconciled; mirror faithful; four owners confirm; CTO
re-verified. Requesting your **closure co-sign** (I do not self-declare a MUST-gate closure).

**Two pre-PLAN items tracked — NEITHER a Cluster-4 blocker (both owners explicit):**
1. **Row-level parity (§15 Q-F)** — formal §17.3 rows for the 7 remaining m-3 observe fields (`achieved_evidence`,
   `evidence_integrity`, the `*_RESULT` set, `executable_claim_results`, `egress_scan_result`, `degradation_notes`,
   `routing_decision` observed profile); m-3 (owner) + m-6 (consumer) want it before Step-1 PLAN to make the §13 fixtures
   machine-checkable. Bounded, additive, non-gate.
2. **Runtime-`away` resolution (m-6 `012906`, NEW — a good catch):** m-6, as the away-bridge consumer, correctly notes
   `away` is an **operator-global RUNTIME availability toggle** (flips A-gate delivery across all lanes, changes
   mid-flight), **not a per-seat spawn-default** like `interactive`/`unattended`. So it does **not** ride the
   per-assignment-frozen `seat_archetype`, and has no clean recorded home under the ruling. The posture vocab
   `{interactive|away|unattended}` arguably conflates *per-seat run-mode* with *operator availability*. This is the
   **c3-lock-amendment lane** m-5/m-6 named — a design item for m-5 (vocab) + m-6 (away-bridge) + m-1 (runtime-state
   home), to resolve **before the away-bridge PLAN (Step-2)**. Tracked; not opened here.

**Gate implication:** on your co-sign, the **full MUST-before-Step-1 gate is satisfied** — five operator decisions ✅ +
Cluster 1 ✅ + Cluster 4 ✅. The only remaining item is the **Step-1 PLAN phase transition**, which stays the
**operator's** to authorize (crosses the charter AUDIT+DESIGN boundary). The two pre-PLAN items resolve before their
relevant PLAN, not before the gate.

Not authorized / not claimed: no unilateral Cluster-4 closure (awaiting your co-sign); no Step-1 PLAN; no code/pcode/
spike; no c3-lock amendment opened; no other c2 change.

ACTIONS_GIT_REF: read-only re-verify of m-2's folded §17.3/§17.6 vs m-4 `:203`/`:200-210` + the four owner-confirms; wrote this relay + appended `master/relays/INDEX.md`; no design-doc edits, no code/source/pcode, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP closure co-sign for Cluster 4 → full MUST gate satisfied; then the Step-1 PLAN phase transition is the operator's to authorize. Pre-PLAN items (row-parity Q-F · runtime-`away` resolution) route on the operator's go, before their relevant PLAN.
