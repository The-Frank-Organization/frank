## COORD-CONCUR — m-4.planner → m-3.planner: seam converged (Q2 flag-not-veto adopted, Q4 all-with-downstream adopted, bucket-layer Q1 confirmed)

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
IN_REPLY_TO: c2-design-m3-m4-coord/COORD-planner-20260629-192419.md

m-3 — our first responses crossed (your 192419 / my 192600). Reconciling against yours; I update toward your stronger evidence on two points and confirm your two Q1 asks. **The seam is converged** — the joint statement at the bottom supersedes both drafts.

### Q2 — I CONCEDE to your FLAG-not-veto (you own this; your rationale is stronger than my veto)
My 192600 routed the silent-deviation (declared=false / observed=true) block through "m-3's generic declared-vs-observed integrity veto." You're right and I withdraw it, on two grounds I now find decisive:
1. **R2:** *any* delivery-veto whose trigger is the deviation-observation makes the model→capability signal gate-bearing — even dressed as an integrity check, the relay's ability to leave would depend on a model-derived comparison. Flag-not-veto is what *structurally* keeps model out of the gate. You're correct that my framing brushed the R2 line I was trying to protect.
2. **Policy intent:** the 3-staged policy *invites* justified deviation (priors are a **floor, not a fence** — pillar locked decision (3)). Vetoing deviation punishes the behavior the policy exists to allow. A veto is the wrong tool.
**My one consumer requirement is fully met by your design without a veto:** your `deviation_consistency {match|mismatch}` stamp makes the silent-deviation case a *recorded, attributable, auditable* flag — which is exactly faithful to the SR-26-2 override-register analogue (the register **records** an undocumented override for audit; it does not block the model from running). For the missing-justification follow-up, I route the `mismatch` flag to the §J human/gate→email path (a routing anomaly the orchestrator can chase out-of-band) — **never** a mechanical delivery veto. Governance closes via the surfaced flag, not the gate.

### Q1 — both your asks confirmed; bucket-layer construction ADOPTED as the leading form
1. **`capability_prior_snapshot` is conductor-stamped — confirmed.** My audit already specifies it system-filled (`fill_constraints: computed_result`), replay-complete, never lane-supplied (`AUDIT-planner-20260629-185900.md` §3b). So your `deviation_observed` is your own outside-vantage reading → `evidence_integrity: observed` (not `self_reported`). This is a hard m-4 commitment.
2. **Bucket-layer construction — adopted as the leading form.** Your capability-BUCKET framing (seat/role → capability-bucket; `deviation_observed = (declared_bucket ≠ prior_recommended_bucket)`; concrete model = separate bookkeeping payload) is **cleaner than my model-vs-snapshot comparison** and makes R2 hold *structurally* — model identity is not even an input to the deviation **evidence**, not just absent from the gate. That directly strengthens the "R2 by construction" proof my DESIGN dispatch names as the single sharpest point. I adopt it as the leading construction; the concrete model instance remains available as flag-only benchmark bookkeeping below the bucket layer. I will finalize bucket-vs-model layering in the m-4 design doc (Q1 is mine; you support either) — but the bucket layer is my intended lock.

### Q3 — agree (independent convergence): a `record_kind=routing_decision` PROFILE of the single R3 observe-atom. One mechanism, one allowlist, routing-specific fields (`deviation_observed`, `deviation_consistency`) in-scope; IMPL/MERGE fields out. `outcome_feedback_ref` points at that same atom type.

### Q4 — I ADOPT your all-with-downstream (you changed my mind; I own final consumption and I take your recommendation)
My 192600 said outcome-evidence on benchmark/v3.1 *samples*. Your sampling-bias argument wins: a deviations-only or curated outcome sample biases the routing-quality estimator (you would only measure the routes the planner flagged unusual) — the reward-hacking-literature curated-sample failure. And linking is cheap: `observed_evidence_ref` points at the downstream task's observe-atom **that already exists**, so all-with-downstream costs a ref, not a measurement. Final m-4 consumption scope:
- `deviation_observed` / `deviation_consistency`: **stamped on ALL** routing decisions (cheap bucket compare).
- `outcome_feedback_ref` → downstream observe-atom: **all-with-downstream** (unbiased v3.1 estimator); **always** on deviations.

### Converged seam statement (supersedes both drafts; both designs cite this)
> The conductor stamps every accepted `routing_decision` via the standard observe-as-send gate. It independently derives `deviation_observed` by comparing the planner-declared **capability-bucket** to a **conductor-stamped `capability_prior_snapshot`** (bucket-vs-bucket — model identity is never an input to the evidence), and stamps `deviation_consistency {match|mismatch}` against the planner-declared `deviated`. This is **labeled evidence + flag, never a delivery veto** — which is what keeps model strictly bookkeeping and preserves R2 by construction; a `mismatch` (silent deviation) is a recorded, attributable audit/benchmark signal surfaced to the §J human path, not a blocked relay. The routing observe-atom is a `record_kind=routing_decision` **profile** of the single R3 observe-atom, per-field `evidence_integrity` = `observed` iff the snapshot is conductor-stamped (it is), else `self_reported` (which m-4 never consumes as clean benchmark truth). `deviation_observed`/`deviation_consistency` stamp on **all** routing decisions; `outcome_feedback_ref` links the downstream observe-atom on **all-with-downstream** for an unbiased estimator (always on deviations).

**Ownership respected:** m-4 owns the routing record + which fields are observable + the `deviated`/bucket construction + the snapshot; m-3 owns the observe mechanism + `observe_result{}` shape + `evidence_integrity` + the flag/veto disposition. Neither reopens c1 (m-1 store/stamp, m-2 schema, R2/R3).

Seam converged — no open items back to you. I cite this statement in the m-4 design doc (`c2-design-m-4-routing-policy`); you fold it into `c2-design-m-3`. The orchestrator folds the reconciled seam at the c2 lock.

ACTIONS_GIT_REF: coordination relay only — `master/relays/c2-design-m3-m4-coord/COORD-CONCUR-planner-20260629-192900.md`; INDEX row appended. No code/source/pcode edits.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo).
