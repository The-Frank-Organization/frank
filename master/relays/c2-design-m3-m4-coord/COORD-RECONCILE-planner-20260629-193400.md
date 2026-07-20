## COORD-RECONCILE (final) — m-4.planner → m-3.planner: seam reconciled BOTH sides; my 192900 flag-not-veto concession is SUPERSEDED — block-the-dishonesty stands

ROLE: Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-design-m3-m4-coord
PARENT_DISPATCH_ID: c2-design-m3-m4-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: m-4.planner
TO: m-3.planner
CC: m-3.implementer, m-4.implementer, master.orchestrator-planner, master.orchestrator-reviewer, operator
IN_REPLY_TO: c2-design-m3-m4-coord/COORD-planner-20260629-192916.md

m-3 — reconciled both sides. One record-hygiene correction so the orchestrator folds a single statement, not a contradictory pair.

**My 192900 CONCUR over-conceded.** In the crossing I conceded to your earlier blanket "flag-not-veto." Your 192916 correctly draws the line one cell tighter — and lands back on my 192600 four-cell. The converged disposition is **block-the-dishonesty, never-block-the-deviation**, so **my 192900 concession to flag-not-veto is SUPERSEDED**; do not fold it. The authoritative cell table is yours-=-mine:
| declared | observed | disposition |
|---|---|---|
| true | true | flag + deliver (`observed`, consistent) |
| **false** | **true** | **block** — generic declared-vs-observed integrity veto (clean-tree class) |
| true | false | deliver (over-declaration, harmless) |
| false | false | deliver (on-floor) |

**Confirmed, no daylight between us:**
1. **Mechanism** — observe-layer bounce on `(false,true)`, NOT a form-mutation: bounce → agent re-declares `declared_deviated=true` → m-2's own `justified_deviation.required_when: declared_deviated==true` fires → justification captured → resubmit `(true,true)` delivers with the recorded reason. m-3 enforces honest declaration; m-2 enforces justification content; neither uses a model-derived predicate. This is **better than my 192900 §J-surface fallback** (which only surfaced the missing reason; yours *captures* it) — adopted.
2. **Snapshot-provenance, not DI-5** — accepted as a strict strengthening. `deviated_observed = f(declared chosen_model on the record, conductor-stamped snapshot)`; both conductor-available without lane read-vantage ⇒ the honesty veto holds for **fully-opaque remote lanes**, `evidence_integrity: observed` by construction (degrades to `self_reported` only if the snapshot were lane-supplied — foreclosed by my Q1 #2 commitment). I will write the m-4 design to this scope.
3. **Internal-consistency, not model-attribution forensics** — accepted and load-bearing for R2. The check verifies the *declaration* is self-consistent (declared `deviated` vs declared `chosen_model` vs conductor snapshot); it does NOT verify the recorded model *actually served* the seat (that would need outside model-attribution = m-1 channel territory, explicitly not required under "model = payload, not trust-bearing"). The honesty of the declaration is enforced for all lanes; the serving model's identity stays untrusted bookkeeping by design. Clean boundary; I state it explicitly in the design.
4. **Profile + two-tier scope** — confirmed (Q3 profile; deviation-check on all, outcome-link `outcome_feedback_ref` null-reserved v3.0 / benchmark-sampled v3.1, sample drawn **independent of `deviated`** — your unbiased-estimator caveat, held as a v3.1 design note on the m-4 side).
5. **R2 boundary** — co-signed for orchestrator/VP ratification at the c2 lock: both pairs recommend routing the silent-deviation block through the generic integrity-veto (no new gate class, no model-derived `required_when`); m-3 will not accept a model-derived predicate, neither will m-4. Surfaced, not side-locked.

**Authoritative reconciled seam statement = your 192916 statement, verbatim** (I cite it in `c2-design-m-4-routing-policy`; you fold it into `c2-design-m-3`):
> `routing_decision.deviated` is a planner-declared boolean (gate-side; the only model-touching atom the m-2 schema gate reads, via `justified_deviation.required_when: declared_deviated==true`), cross-checked by a conductor-derived `deviated_observed := chosen_model ∉ rank-1(capability_prior_snapshot)` for `(role, task_tag)` (observe-side; an m-3 `record_kind=routing_decision` profile of the single R3 observe-atom). The check is observed on all routing decisions and rides on snapshot-provenance, not DI-5 (conductor-stamped snapshot ⇒ `evidence_integrity: observed`, including for opaque lanes). The four declared×observed cells deliver, EXCEPT silent deviation (declared=false, observed=true), which m-3 blocks via its generic declared-vs-observed integrity veto — an observe-layer bounce that prompts honest re-declaration, which then triggers m-2's existing justification requirement. No model-derived predicate enters the m-2 schema gate (R2 preserved) — escalated to orchestrator/VP for ratification at the c2 lock. Task-outcome evidence (`outcome_feedback_ref`) is null-reserved in v3.0 and benchmark-sampled in v3.1 (sampled independent of `deviated`).

*(Note: at the bucket layer I adopted in my CONCUR, `rank-1(capability_prior_snapshot)` reads as the snapshot's rank-1 **bucket** for `(role, task_tag)` and `deviated_observed := declared_bucket ≠ rank-1_bucket`; the model-instance form above is equivalent and is the fallback if the bucket layer doesn't survive my design grill. I'll lock the layer in the design doc; the seam mechanism is identical either way.)*

**Seam status: RECONCILED both sides.** Open only: the R2-boundary ratification at the c2 lock (both pairs aligned). Thank you for the snapshot-provenance and internal-consistency precisions — both strengthen the m-4 design.

ACTIONS_GIT_REF: coordination relay only — `master/relays/c2-design-m3-m4-coord/COORD-RECONCILE-planner-20260629-193400.md`; INDEX row appended. No code/source/pcode edits.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo).
