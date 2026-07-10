## Team m-4 — reconciled AUDIT view (routing & policy primitive)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: read-only
DISPATCH_ID: c2-audit-m-4
PARENT_DISPATCH_ID: c2-audit-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-4.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, m-4.implementer, operator
BUNDLE_ID: m-4-routing-policy
OWNER: m-4 (Routing & Policy)

Reconciles the two independent passes — `AUDIT-planner-20260629-185900.md` + `AUDIT-implementer-20260629-185224.md` (each carries full E1 detail; this view is the ledger + merged design-of-record-candidate + single verdict). Resolution discipline: toward evidence, not seniority.

---

## Reconcile ledger
| Item | Status | Resolution |
|---|---|---|
| 4-bucket verdict | **agree** | identical: still-open = the governance-record primitive; already-closed = carrier (locked m-1/m-2) + selection runtime (jcode/CC) + published priors; product-overlapped = jcode runtime (Step-3); recommended-next = thin governance layer, write R2 record now, execute Step-3. |
| Implicit-routing gap | **agree** | identical: surveyed systems select via config/flags/runtime/coordinator/learned-scalar; none records a separate seat-stamped, justified routing decision consumed as provenance. |
| R2 preserved (model never a gate input) + identity≠authority + category-B | **agree** | both build on the locked contract; neither reopens it. |
| Step-1 ride / Step-3 execute split | **agree** | record writable on existing runtimes now; execution = Step-3; no execution dependency blocks the DESIGN lock. |
| 3-staged policy (prior → justified deviation → outcome feedback in a later release) | **agree** | identical staging; benchmark loop = forward hook, not Step-0. |
| m-3 seam (evidenced record + benchmark consumes observed evidence; m-4 ≠ m-3 internals) | **agree** | route into the `c2-*-coord` COORD thread at DESIGN. |
| m-5 seam (archetype tags parameterize priors; ceiling caps routable authority) | **agree** | **cannot lock m-4 without an m-5 seam disposition at lock** — surface, don't close. |
| capability-prior = versioned table w/ reason codes + snapshot the used version into the record | **agree** | independent convergence (corroborates the c1 snapshot sharpening). |
| Router API shape | **different-coverage → merged** | implementer's `route_dispatch()` single verb + **fail-closed failure mode** ⊕ planner's policy-engine/record/dispatch-ref decomposition. Adopt the verb + fail-closed; keep the decomposition as its internals. |
| Justified-deviation encoding | **different-coverage → merged** | implementer's **machine-readable reason code** + planner's **free-text narrative** = the SR-26-2 "reason code + narrative" override-register pattern. Adopt both. |
| `required_when` and R2 | **different-coverage → adopt planner** | the predicate atom must be the agent-declared `deviated` boolean, never a model-value comparison (no `model_*` predicate enters the gate). Implementer preserved R2 at consumption; planner adds the fill-time construction. |
| External prior-art positioning | **different-coverage → adopt planner** | implementer relied on local sources; planner's web sweep supplies the external survey + **two honest qualifications** (below). The reconciled view carries them so DESIGN does not overclaim. |
| record-kind naming (`record_kind: routing` vs `routing_record_kind` own enum) | **resolved (not a conflict)** | two levels: a top-level `record_kind = routing` discriminator **and** a routing-internal `routing_record_kind` sub-enum (`routing_decision` / reserved `routing_deviation`,`routing_update`). Reuse the DESIGN_RECORD_KIND *shape*, not its values. |
| Cost/latency/privacy `constraints` | **different-coverage → fold as optional/forward** | capability-prior is the initial floor; constraints are optional policy inputs / reserved (the external cost-quality literature lives here). Settle weight at DESIGN. |
| genuine disagreements | **none** | — |
| operator-decision-needed | **none new** | the three flagged items (§ Operator) are config/scope confirmations, not blockers. |

---

## Reconciled pair verdict — PRIMARY_BUCKET: still-open
The frank routing/policy **primitive** is still-open: a routing decision that is *first-class, recorded, attributable, justifiable, and human-overridable*. The **carrier** (record + forgery-robust stamp + lineage + `parent_picker` admission) is already provided by locked c1 m-1/m-2 (FieldSpec reserves `router` consumer + `routing_ref` lineage_role, form-schema §4:60-63); the **selection runtime** (jcode `MultiProvider`/failover; CC per-task `model` override) already exists as Step-3 wire-not-rebuild assets; the **priors** are published. The build = the **governance-record layer** on top: the routing record + declared capability-prior table + 3-staged policy + the thin fail-closed router API. **Recommend PROCEED-TO-DESIGN** with the two seam conditions below.

## Reconciled implicit-routing gap (+ honest qualifications)
Every surveyed system selects a model/role with **no recorded, justifiable governance decision** — at most observability (*which/what* ran), never a governed *why* a human can audit and override before it takes effect. The stock protocol: no routing field at all (out-of-band lane hand-assignment). jcode: runtime provider switch, no decision record. claude-code: per-task model override recorded as role-only telemetry, no "why". agent-scripts: prose logs, not a per-decision artifact. External: nine predictive routers = learned scalar over a threshold; four commercial products = dashboards/telemetry; bandits = throwaway in-memory state.

**The routing record closes it by porting a mature GRC primitive** — the "default-from-floor + justified override + recorded reason + human approval" discipline of **SR 11-7 → SR 26-2** model-risk override registers (ISO 27001 / NIST POA&M / ADR / MLOps approval-gate family) — into per-dispatch LLM routing, where it has never been applied. **Two qualifications DESIGN must concede (not overclaim):**
- **(a)** Routesplain (2511.09373) + Arch-Router (2506.16655) already make routing rationale interpretable/intervenable — differentiate on the *seat-stamped, persisted **deviation-against-a-declared-floor** audit artifact*, not on "we make routing interpretable."
- **(b)** Non-gradient bandit outcome-feedback already exists (C2MAB-V 2405.16587; PILOT/LinUCB 2508.21141) — so the later-release loop's novelty is the *auditable, persisted decision+update artifact*, **not** the non-gradient mechanism. The "Fugu-reward-analog" framing must be hedged accordingly.

## Reconciled design recommendation (merged)
- **Router API:** `route_dispatch(target_dispatch_id, assignments[], policy_context) -> accepted routing_relay_id` (+ read-only projection for the dispatch). Internals = (policy engine: prior lookup at fill time) → (planner-stamped `routing_decision` record) → (conductor admits it to m-1 `parent_picker`; dispatch carries `routing_ref` provenance only). **Fail-closed:** no acceptable route emits `human_decision_required` / `routing_unavailable` — it MUST NOT silently fall back to a default model.
- **Capability-prior:** declared, versioned `map<(role, task_tag, evidence_target, [archetype_tag]) → ranked routes + reason codes + basis>`; **snapshotted** into the record (replay-complete, not a live ref); config-sourced + operator-configurable (§J pattern).
- **Routing-record schema (m-2 FieldSpec consumer):** top-level `record_kind=routing` + internal `routing_record_kind` (own enum, DESIGN_RECORD_KIND *shape* not values); `routing_assignments` row_array `{seat, role, model_or_class, runtime_route_key, provider_family, capability_prior_id, policy_stage, deviated:bool}` seat_scoped to planner/orch-planner; `capability_prior_snapshot` (system computed_result); `justified_deviation` free_text **required_when `any(routing_assignments.deviated==true)`** (the atom is the **agent-declared boolean**, never a `model_*` comparison — R2 by construction; conductor observes declared-vs-snapshot as m-3 evidence) **+ machine-readable `deviation_reason_code`**; `constraints` (budget/latency/privacy — optional/forward); `outcome_feedback_ref` id_ref null-reserved (a later release). `consumers:[router, lineage_engine, human_surface, observe_gate]`.
- **3-staged policy:** (1) prior floor → (2) justified deviation w/ reason-code+narrative (ships at Step-1) → (3) outcome feedback in a later release consuming m-3 evidence via `outcome_feedback_ref` (forward hook; differentiator = persistence/auditability).

## Reconciled boundary contract + seams
- **→ m-3:** routing record = evidence-addressable record (m-3 observer-only, positive write-allowlist R3); benchmark consumes m-3 observed evidence via `outcome_feedback_ref`; m-4 declares which fields are observed, m-3 owns how. **Co-design in `c2-*-coord`.**
- **→ m-5:** archetype tags parameterize priors; authority-ceiling-at-spawn caps routable authority (router refuses to staff a seat the ceiling forbids — identity≠authority at the routing layer). **m-4 cannot lock without an m-5 seam disposition at lock.**
- **→ m-6:** `routing` = category-B; routing decision normally orchestrator-absorbed; `human_decision_required` / `routing_unavailable` / unclassified `other` surface to the human path; m-6 owns bucketing + projection (ODB content for A-escalation). m-6 = warm consumer lens.
- No orphaned primitive (every field has a named consumer).

## Operator-judgment items (none blocking)
1. Capability-prior **seed values** = operator-configurable (ship a default seed: Fugu priors + ours; operator confirms/customizes).
2. **Step-1 record scope** — recommend recording the decision for human-launched lanes now (Step-1 "automated operator-relay"), execution deferred to Step-3. Confirm at PROCEED-TO-DESIGN.
3. m-5 seam disposition is a **DESIGN coordination requirement**, not an operator item yet.

## Coordination relays sent
None. No blocking m-1/m-2 gap; R2 is preserved, not reopened. The m-3↔m-4 + m-5 seams are routed through DESIGN coordination (the COORD thread + lock-time disposition), per the dispatch's contract constraint — no unreviewed side-locks.

---

ACTIONS_GIT_REF: workspace file writes in non-git docs workspace — `master/relays/c2-audit-m-4/AUDIT-planner-20260629-185900.md`, `master/relays/c2-audit-m-4/RECONCILE-planner-20260629-190200.md`; INDEX rows appended in `master/relays/INDEX.md`. Five read-only research lenses (transcripts under `tasks/`); no source changes.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo).
