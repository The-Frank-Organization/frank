# Routing & Policy — design (the governance-record layer)

**DESIGN_DOC_ID:** c2-design-m-4-routing-policy
**Cycle/phase:** c2 / DESIGN (terminal = design-lock; no PLAN/IMPL this phase)
**Owner:** m-4 (Routing & Policy) — design-lead m-4.planner; adversarial review m-4.implementer
**Status:** LOCKED-at-c2 (design-of-record) — m-4.implementer DESIGN-REVIEW **approve** on rev1
(`c2-design-m-4/DESIGN-REVIEW-implementer-20260629-203329.md`); rev1 folded the three must-revise findings
(F1 precise R2 invariant + locked bucket-vs-bucket form + bucket-binding atom §2; F2 `deviation_reason_code`
`required_when` §5; F3 template no-bypass invariant §7) and the non-blocking carry-forward (leftover membership
shorthand in §0/§9 normalized). **rev2 — c2 fold** (`c2-fold-m-4` 20260630): folded the VP-approved m-5/m-6
consumer-lens findings (F2 per-assignment `seat_archetype`+`authority_ceiling`; F3 `seat_archetype` distinct
opaque tag; **M4-1 CONFIRMED** — B→A escalation via the existing c1 monotonic HUMAN_GATE, no new gate class) —
bounded-additive, see §15 fold-log. **LOCKED at the c2 lock** — the R2-boundary is **RATIFIED**
(readiness-fix-c1/c4 + the design-review reconciles); the m-5 concrete-semantics reservation is **carried to c3**
(m-5-owned). **c5 folds:** claim-sweep (§16) + decision-⑤ R2-guard (§17). **c6 re-review cleanup** (`c6-fix-m-4`;
`master/DESIGN-REREVIEW-2026-07-02.md` CONDITIONAL-GO): §5/§7/§13 build-carry + template-authoring + reason-code
sourcing + status/§J reconcile (see §18 c6 fold-log). Builds on the VP-approved
`c2-audit-m-4` reconcile and the
locked c1 contract (`master/ARCHITECTURE.md` R2/§5/§J). The m-3↔m-4 seam is RECONCILED both sides
(`c2-design-m3-m4-coord`); the one R2-boundary item **was escalated to orchestrator/VP and RATIFIED at the c2 lock**
(readiness-fix-c1/c4 + the design-review reconciles).
**Sources of record:** `c2-audit-m-4/RECONCILE-planner-20260629-190200.md` (the audit), the reconciled seam
statement (`c2-design-m3-m4-coord/COORD-planner-20260629-192916.md`), this cycle's operator design-grill.

> **The thesis.** Routing is a first-class, recorded, justifiable **governance decision** — not implicit config.
> Every surveyed system selects a model with no recorded, justifiable, auditable decision; frank ports the
> SR 11-7→SR 26-2 model-risk **override-register** discipline (floor + documented override + reason + approval)
> into per-dispatch LLM routing. The whole design holds one invariant: **the model is payload/bookkeeping, never
> a gate input (R2).**

---

## §0. GRILL_LOCK — operator decisions (durable)

Folded from the c2 design-grill with the operator (final authority). These are locked design inputs.

- **GL-1 — capability-prior shape = two-layer bucket model.** The prior is split into **Layer 1: bucket
  membership** (`bucket → {models}`, churns on model change) and **Layer 2: bucket recommendation**
  (`(role, task_tag) → ranked bucket list`, churns on judgment change). Deviation is computed at the **bucket**
  layer (`deviated_observed := declared_bucket ≠ rank-1 recommended bucket` — a bucket-vs-bucket comparison
  that does **not** read model identity; an auxiliary binding check `chosen_model ∈ members(declared_bucket)`
  reads the model as payload only, never as a gate input — see §2 for the precise R2 invariant). Rationale:
  makes R2 structural at the deviation comparison; model-churn = a one-line membership edit (not N recommendation rewrites); benchmark signal is
  durable across model churn. The direct `(role,task_tag)→model` form is the documented fallback (still
  R2-safe, but flag-only rather than structural).
- **GL-2 — later-release tuning scope = recommendations only.** The later-release outcome-feedback loop may retune **Layer 2
  (recommendations)** only; **Layer 1 (membership) stays operator/config-owned.** A benchmark may route tasks
  among operator-defined capability classes; it may **not** silently reclassify what a class *means*. Both
  layers ship default-seeded then operator-configurable (the §J pattern). This is a governance bound on the
  loop's write-scope, not a record-shape constraint (the snapshot captures both layers regardless).
- **GL-3 — the initial-release record scope = record-now / execute-Step-3 (option B).** The initial release records the routing decision
  while **riding existing runtimes**; the conductor does not drive its own provider adapters (that is Step-3 per
  `ROADMAP.md`). The initial release enforces **declaration honesty** (the deviation check, §2), not **execution fidelity**
  (which model actually served) — the latter is a documented Step-1→Step-3 boundary (§12). The
  orchestrator-planner seat authors the record; recording is **always-on but near-zero-friction on-floor**
  (auto-filled when the planner takes the prior with no deviation).
- **GL-4 — the initial release ships routing TEMPLATES (operator-directed scope addition).** The initial release ships 1–3 selectable team
  templates; selecting one spawns the team (panes/sessions opened, named per seat, started with preset models,
  awaiting input) — **option B (ride existing runtimes; no own-runtime).** m-4 owns the **routing-template
  mechanism** (§7); **m-5 owns the template structures + the shipped lineup**; conductor-core owns the
  pane-spawn. The cross-domain + roadmap parts are flagged to the orchestrator
  (`c2-design-m-4/SITREP-planner-20260629-200500.md`).

---

## §1. The gap this closes (recap; full detail in the audit)

No surveyed system records a *justifiable governance decision* for model selection — at most observability
(*which* model ran), never a governed *why* a human can audit and override before it takes effect. The stock protocol: no
routing field at all (out-of-band lane hand-assignment). jcode: runtime provider switch, no decision record.
claude-code: per-task model override recorded as role-only telemetry, no "why". External predictive routers:
a learned scalar over a threshold. The routing record closes this by making the decision a **first-class,
seat-stamped, lineage-gated relay**: attributable (rides m-1 `submit()`, **confusion-resistant** stamped `FROM` — m-1-owned; **D5 residual**: a malicious same-uid seat forging a store record is out of scope),
auditable (append-only, lineage-walkable), overridable (a §J category-B gate that escalates to A only on
`human_decision_required`). **Novelty is located precisely in §11** (not interpretable routing in general; not
non-gradient adaptation in general).

---

## §2. The `routing_decision` record — R2 by construction

The single sharpest proof obligation: **no `model_*` predicate ever enters the schema gate.** The construction:

1. **Declared (gate-side).** Per assignment, the planner fills a plain boolean `declared_deviated`
   (`agent_enum_pick`). The **only** model-touching atom the m-2 schema gate reads is
   `justified_deviation.required_when: declared_deviated == true`. The atom is a plain boolean — **no model
   comparison enters the gate.**
2. **Observed (observe-side, m-3).** Per GL-1 the deviation is computed at the **bucket layer**: the planner
   declares a `declared_bucket` per assignment, and the conductor derives
   `deviated_observed := declared_bucket ≠ rank-1(recommended bucket for (role, task_tag))` against the
   conductor-stamped `capability_prior_snapshot`. **This comparison is bucket-vs-bucket — it does not read
   model identity.** Separately, the conductor derives an auxiliary **bucket-binding** observation
   `bucket_binding_observed := chosen_model ∈ members(declared_bucket)` — this *does* read `chosen_model`
   (payload) to confirm the declared bucket is honest about the concrete model. The observed atom set is
   `{deviated_observed, bucket_binding_observed, declared_bucket, rank1_recommended_bucket, chosen_model}`
   with `evidence_integrity: observed`.
3. **`deviated` is not a freestanding truth bit** — it is a *declared* bit cross-checked by an *observed* bit,
   the same declare/observe geometry as the monotonic `HUMAN_GATE` and the clean-tree integrity check.

> **The precise R2 invariant (this supersedes any stronger phrasing elsewhere).** The observe layer **may read
> `chosen_model` (payload)** to derive an integrity observation (the bucket-binding check). What R2 forbids is
> that **no model-derived predicate enters the m-2 schema gate, the authority gate, the lineage gate, or the
> work-dispatch header.** The *deviation* comparison is bucket-vs-bucket (model identity is not an input to
> *it*); the only model-touching atom any *gate* reads is the plain `declared_deviated` boolean. "Model =
> payload" means: readable by the observe layer to produce *labeled evidence*, never a gate/authority input.
> A **bucket-binding mismatch** (`chosen_model ∉ members(declared_bucket)` — the planner declared a bucket its
> own chosen model isn't in) is handled by the **same generic declared-vs-observed integrity veto** as the
> silent-deviation case below (observe-layer bounce, never a gate predicate).

**Disposition of the four declared×observed cells (m-3-owned; reconciled):**

| declared | observed | disposition |
|---|---|---|
| true | true | flag + deliver (`observed`, consistent) |
| **false** | **true** | **block** — m-3's generic declared-vs-observed integrity veto (clean-tree class) |
| true | false | deliver (over-declaration, harmless) |
| false | false | deliver (on-floor) |

The block on **(false, true)** — a *silent/undeclared* deviation — is the **veto on the dishonesty, not on the
deviation**. Vetoing a deviation per se would make the model→capability signal gate-bearing (R2 violation) and
would punish the justified deviation the policy exists to *invite*; vetoing the *undeclared-ness* is the
generic clean-tree-class integrity veto and introduces **no new gate class** and **no model-derived schema
predicate**. Mechanism: the observe-layer bounce prompts honest re-declaration → m-2's own
`justified_deviation.required_when` then fires → the justification is **captured** → resubmit delivers
`(true,true)` with the recorded reason. **m-3 enforces honest declaration; m-2 enforces justification content;
neither uses a model-derived predicate.**

**Snapshot-provenance, not DI-5.** `deviated_observed = f(declared chosen_model on the record,
conductor-stamped snapshot)` — both inputs are conductor-available **without reading inside the lane**. So the
honesty check holds even for **fully-opaque remote lanes** (a strict strengthening over DI-5-dependent
observations like the clean-tree check); it is `observed` whenever the snapshot is conductor-stamped (by
construction, since the floor is the conductor's own policy). It degrades to `self_reported` only if the
snapshot were ever lane-supplied — foreclosed by the m-4 commitment that the snapshot is system-filled.

**Scope (avoids over-claim).** This is an **internal-consistency / honest-declaration** check (declared
`deviated` vs declared `chosen_model` vs conductor snapshot). It is **not** a forensic verification that the
recorded model *actually served* the seat — that would need outside model-attribution (m-1 channel territory),
which under "model = payload, not trust-bearing" the initial release explicitly does **not** require. The honesty of the
*declaration* is enforced for all lanes; the identity of the *serving* model stays untrusted bookkeeping by
design (§12).

> **R2-boundary item — RATIFIED at the c2 lock (both pairs aligned):** routing the silent-deviation block
> through m-3's generic integrity-veto (not a model-derived `required_when`) is the cross-domain R2-preservation
> choice; **ratified** (readiness-fix-c1/c4 + the design-review reconciles). Neither m-3 nor m-4 accepts a
> model-derived schema predicate.

---

## §3. The `route_dispatch()` API — fail-closed

`route_dispatch(target_dispatch_id, assignments[], policy_context) -> accepted routing_relay_id`
(+ a read-only projection for the dispatch header/body).

- **Inputs:** target dispatch id; the seats/roles to staff; the archetype tag vector (opaque, m-5-owned —
  §8/§10); task-domain hints / `task_tag`; `evidence_target`; optional constraints (budget/latency/privacy —
  reserved/forward, §6); the available runtime routes; an optional prior routing relay for reroute.
- **Internals:** (policy engine: bucket-recommendation lookup at fill time) → (planner-stamped
  `routing_decision` record) → (conductor form+lineage validation, stamp `FROM`, accept, admit to m-1's
  `parent_picker` candidate set). The dispatch references the accepted record as `routing_ref`
  (provenance/bookkeeping only).
- **Fail-closed (load-bearing):** if no acceptable route exists, emit `human_decision_required` /
  `routing_unavailable` — it **MUST NOT silently fall back to a default model.** A silent default would be an
  unrecorded routing decision, re-opening the gap this design closes.
- **Step-1 vs Step-3:** in the initial release the record is *written* and the human (or the GL-4 template-spawn) launches the
  lane; at Step-3 the same record is *executed* by the conductor's own runtime (§12). The API surface and the
  record do not change between steps.

---

## §4. The capability-prior table — declared, versioned, snapshotted (two-layer)

Per GL-1:
```
Layer 1 — bucket membership (operator/config-owned; churns on model change):
   top-tier-coding      → { Opus-4.x, … }
   strong-general-code  → { Sonnet-4.x, … }
   strong-math          → { GPT-5.x, … }
   strong-science       → { Gemini-3.x, … }
   fast-cheap           → { Haiku-4.x, … }

Layer 2 — bucket recommendation ((role, task_tag) → ranked buckets; later-release-tunable per GL-2):
   (implementer, swe-debug) → [ top-tier-coding, strong-general-code ]
   (researcher, science)    → [ strong-science ]
   …
```
- **Declared, not learned** — seeded from Fugu's published priors (GPT→math, Gemini→science/GPQA, Opus→SWE/
  debug; pillar §4.2 confirms Fugu's learned peaks match these), model cards, and ours. Each recommendation
  row carries a **`basis`** (benchmark name / model-card claim / [later release] observed-outcome) — what makes it a
  *justifiable* floor, not a magic constant.
- **Snapshotted into the record** (`capability_prior_snapshot`, system-filled `computed_result`,
  replay-complete — captures **both** layers at decision time). Not a live config ref: an immutable record
  cannot be reconstructed after the fact, and a later release must attribute each outcome to the prior version in force.
- **Config-sourced + operator-configurable** (§J pattern); a default ships, the operator edits.

---

## §5. The `routing_decision` FieldSpec (consumer of the locked m-2 registry)

Top-level discriminator `record_kind = routing` + a routing-internal `routing_record_kind` enum (its **own**
value set — `routing_decision`; reserved `routing_deviation` / `routing_update` for a later release — reusing the
`DESIGN_RECORD_KIND` *shape*, not its values; reusing `design-doc|audit-record|direct-override` would mis-own
the field, confirmed at `tools/relay-lint.py:42`). Fields (each a FieldSpec; the locked m-2 registry already
reserves `router` in `consumers` and `routing_ref` in `lineage_role`, form-schema §4:60-63):

| field | owner / fill | type | notes |
|---|---|---|---|
| `routing_record_kind` | agent_enum_pick (routing set) | enum | record-kind discriminator (own values) |
| `routing_assignments` | seat_scoped_enum → planner / orch-planner **(+ `operator` on `template_ref`-bearing template-spawn records only — §7, m-4-F4)** | row_array | altitude B; per row: `{ seat, role, task_tag, declared_bucket, chosen_model (pinned\|resolved), declared_deviated:bool, pin_mode:{pinned\|slot}, seat_archetype (opaque; m-5-valued, c3), authority_ceiling (resolved-at-spawn, recorded) }`. **F2 fold:** `seat_archetype` + resolved `authority_ceiling` are recorded **per assignment** for replay/audit on **both** the template *and* hand-authored (non-template) staffing paths — the ceiling that denied write authority is replay-complete *on the record*, not only enforced (§8). Both are **opaque** (concrete values m-5-owned, c3). |
| `capability_prior_snapshot` | system (`computed_result`) | object | both layers, replay-complete (§4) |
| `justified_deviation` | free_text | text | **required_when `any(routing_assignments.declared_deviated == true)`** (§2) |
| `deviation_reason_code` | agent_enum_pick | enum | machine-readable reason (SR-26-2 "reason code" + the `justified_deviation` narrative); **required_when `any(routing_assignments.declared_deviated == true)`** — same grain/treatment as `justified_deviation`, so the reason code can never be omitted where the policy says it is load-bearing. **Value-set = config-sourced enum (m-4-F5, c6): default-seeded + operator-configurable + hardcoded `other` fail-safe** (the §J2 `gate_category` pattern). **Shipped default vocabulary:** `capability_gap` · `cost_budget` · `latency_budget` · `bucket_unavailable` · `operator_directive` · `experiment` (later-release-reserved) · `other` (fail-safe ⇒ the `justified_deviation` free-text carries the reason). m-2 §17.3 mirrors the value-set (flagged). |
| `constraints` | agent_enum_pick / free_text | object | budget/latency/privacy — **reserved/forward** (§6) |
| `outcome_feedback_ref` | system / id_ref | id_ref | **null-reserved** (a later release); links the downstream observe-atom |
| `template_ref` | system / id_ref | id_ref | set when spawned from a template (§7); null otherwise |

`consumers: [ router, lineage_engine, human_surface, observe_gate ]`. `lineage_role: routing_ref` on the
referencing dispatch. **Model values appear only inside `routing_assignments` as payload** — never in a
`required_when`/`visible_when` predicate, never in an authority/lineage gate.

---

## §6. The 3-staged policy + the later-release forward hook

- **Stage 1 — capability-prior floor.** The policy engine suggests the rank-1 recommended bucket per
  `(role, task_tag)`; the seat-form renders it as the default (a slot resolves to a member of that bucket).
- **Stage 2 — justified deviation (ships in the initial release).** The planner may pick off-floor → declares
  `declared_deviated = true` → the form requires `justified_deviation` + `deviation_reason_code` (the reason code
  drawn from the config-sourced value-set seeded in §5 — default + operator-configurable + `other` fail-safe).
  Recorded, attributable, auditable, human-overridable.
- **Stage 3 — outcome feedback (later-release forward hook only).** The benchmark consumes m-3 observed outcomes (via
  `outcome_feedback_ref`) to retune **Layer 2 recommendations only** (GL-2). Non-gradient (bandit/tally) — but
  per §11(b) the **differentiator is the persisted, auditable decision+update artifact, not the non-gradient
  mechanism.** Shaped day-one via the snapshot + ref fields so it bolts on without re-cutting the record or
  gate. **later-release sampling caveat (held):** draw the outcome sample **independent of `declared_deviated`** — a
  deviations-only/curated sample biases the routing-quality estimator (reward-hacking curated-sample failure).
- **`constraints` (budget/latency/privacy)** are **reserved/forward** policy inputs (the external cost-quality
  literature lives here); capability-prior is the initial-release floor. Their weight in the policy is an open question
  for a later cycle (§13), not an initial-release closure.

---

## §7. Routing templates (GL-4 — the operator-directed extension)

A **routing template** is a saved, named, reusable set of per-seat model assignments — the **model-assignment
side** of a team template (m-5 owns the structure). Each assignment is:
- **pinned** — an exact model (`lead = Gemini-3-pro`); always that model; or
- **slot** — a capability bucket (`lead = strong-science`); the router resolves it to a current member of that
  bucket at spawn (this is where GL-1's bucket layer pays off — templates are churn-stable);
- mixable per seat (`pin_mode` on each `routing_assignments` row).

**On template-spawn:** m-4 emits a `routing_decision` record for the team — pinned rows filled directly, slot
rows resolved via the Stage-1 prior — with `template_ref` set (provenance: which template).

**No-bypass invariant (mechanical — a template is a pre-fill, never an escape hatch).** A template-spawn
`routing_decision` MUST NOT bypass Stage 2. For **every off-floor assignment** (pinned *or* slot), the
**emitted** `routing_decision` sets `declared_deviated = true` for that row and **snapshots**
`justified_deviation` + `deviation_reason_code` *into the emitted record* (the template author's justification
is **copied in** — snapshot, not live-ref, consistent with `capability_prior_snapshot`). So the emitted record
is self-contained and replay-complete; `template_ref` records *which* template but the justification is never
chased as informal template prose. An off-floor template assignment therefore produces a recorded deviation in
the emitted record **exactly as a hand-authored one would** — and the §2 deviation-honesty check + the m-2
`required_when` on `justified_deviation`/`deviation_reason_code` apply unchanged. A template is thus a
**reusable, pre-filled routing decision** that changes **nothing** in the locked contract or R2 (it is a
record, not a gate input). **Operator-customizable** (save/edit), config-sourced like the priors.

**Template-spawn authoring model (m-4-F4, c6 — decided before the GL-4 build step).** A template-spawn
`routing_decision` is authored **`FROM = operator`** on the operator-relay channel: the operator selecting the
template is the authorizing act, and at spawn there may be **no orchestrator-planner seat yet** (the template
*creates* the team, its orchestrator seats included). The `routing_assignments` **seat-scope is widened to admit
`operator` as an authoring seat on `template_ref`-bearing records only** — ordinary hand-authored routing stays
planner / orch-planner scoped. **`declared_deviated` stays a *declared* (agent/operator-owned) bit — it is NOT
re-typed system/computed on the template path;** re-typing it would reintroduce the readiness-fix-c4
declared-vs-observed conflation the m-3 silent-deviation seam depends on (the declared bit must remain an
independent agent assertion, cross-checked by the conductor's observed bit). The pre-declared template
justification is snapshotted into the emitted record (the no-bypass invariant above). **Propose-vs-stamp holds
(seam):** the template/operator *proposes* role/archetype from a fill-time-pruned candidate set; the conductor
*stamps* the resolved `seat_archetype` + `authority_ceiling` per-column. **Cross-domain mirror required (flagged
to the orchestrator — m-4 records the decision, does not write those docs):** m-2 §17.3 (widen the
`routing_assignments` seat-scope to admit `operator` on `template_ref`-bearing records) + m-7 S11 (template-spawn
author = operator-relay).

**Ownership boundary:** m-4 = the template's model-assignment mechanism + record emission on spawn. **m-5** =
the template structures (seats/panes/gate-set/read-only-ness) + the shipped 1–3 lineup. **Conductor-core** =
opening/naming the panes + boot-relay delivery. Roadmap sequencing of the initial-release scope addition is the
orchestrator's (SITREP `c2-design-m-4/SITREP-planner-20260629-200500.md`).

---

## §8. identity ≠ authority at the routing layer — the authority-ceiling cap

m-1 owns **who** (the stamp); m-4 owns **what a stamped seat may route** (§5 ratified). At the routing layer
this realizes as: the **archetype authority-ceiling-at-spawn caps what the router may assign.** The router
**refuses to staff a seat the archetype ceiling forbids** — e.g. a read-only sensor team's seat can never be
routed into write/merge authority; `route_dispatch()` returns `routing_unavailable` (fail-closed, §3) rather
than down-grading silently. This is anti-confused-deputy: authority is keyed to the stamped seat within the
archetype ceiling, never to the model behind it.

**F3 fold — the tag is `seat_archetype`, a distinct opaque tag in the archetype vector (per-seat-at-spawn).**
The archetype tag-space has two orthogonal axes: m-3's **`slot_in`** = *work-archetype* (per-record,
conductor-classified at work-record acceptance) vs m-4's **`seat_archetype`** = *seat-archetype*
(per-seat-at-spawn) — `seat_archetype` is the tag the prior table keys on for the authority-ceiling (this §8)
and the capability recommendation (§10). The two are recorded as **distinct opaque tags** (model-free; concrete
tag-space + ceiling semantics stay **m-5-owned, c3**).

**F2 fold — record the resolved ceiling per assignment (replay-complete for the non-template path too).** For
template-spawned teams `template_ref` carries the per-seat archetype; for **hand-authored** staffing the
applied `seat_archetype` + resolved `authority_ceiling` are written **per assignment** (§5) so the ceiling that
denied write authority is auditable/replayable on the record itself — not only enforced at spawn and then lost.

---

## §9. The m-3↔m-4 seam — RECONCILED (cite)

The reconciled seam statement (both pairs; `c2-design-m3-m4-coord/COORD-planner-20260629-192916.md`, m-4
reconcile-final `…COORD-RECONCILE-planner-20260629-193400.md`):

> `routing_decision.deviated` is a planner-declared boolean (gate-side; the only model-touching atom the m-2
> schema gate reads, via `justified_deviation.required_when: declared_deviated==true`), cross-checked by a
> conductor-derived `deviated_observed := chosen_model ∉ rank-1(capability_prior_snapshot)` for
> `(role, task_tag)` [locked canonical = bucket-vs-bucket on `declared_bucket` per GL-1/§2; this membership
> phrasing is the equivalent fallback] (observe-side; an m-3 `record_kind=routing_decision`
> **profile** of the single R3 observe-atom). The check is observed on **all** routing decisions and rides on
> **snapshot-provenance, not DI-5** (conductor-stamped snapshot ⇒ `evidence_integrity: observed`, including for
> opaque lanes). The four declared×observed cells deliver, EXCEPT silent deviation (declared=false,
> observed=true), which m-3 blocks via its generic declared-vs-observed integrity veto — an observe-layer
> bounce that prompts honest re-declaration, which then triggers m-2's existing justification requirement. No
> model-derived predicate enters the m-2 schema gate (R2 preserved) — escalated to orchestrator/VP for
> ratification at the c2 lock. Task-outcome evidence (`outcome_feedback_ref`) is null-reserved in the initial release and
> benchmark-sampled in a later release (sampled independent of `deviated`).

m-4 declares the closed routing-observed field set for m-3's R3 allowlist:
`{ deviated_observed, bucket_binding_observed, declared_bucket, rank1_recommended_bucket, chosen_model }`
(+ reserved task-outcome). m-3 owns the observe mechanism, the veto/flag disposition, and the
`evidence_integrity` result shape (a profile of the single observe-atom). **Canonical deviation form locked =
bucket-vs-bucket** (§2) — m-3's preferred construction (`c2-design-m3-m4-coord/COORD-planner-20260629-192419.md`),
with the `chosen_model ∉ rank-1` membership phrasing in the reconciled statement above as its equivalent
fallback. This refinement (the bucket-binding atom + the locked bucket-vs-bucket form) is **additive to the
m-4-declared observed set** — m-4 owns *which* fields are observed, m-3 owns *how* — so it needs no seam
re-negotiation; flagged to m-3 at design-complete. Seam status: **reconciled both sides**; the R2-boundary
ratification is **complete** (c2 lock; readiness-fix-c1/c4 + the design-review reconciles).

---

## §10. Consumer boundary contract (named before lock)

- **→ m-3:** the `routing_decision` record is a possible **evidenced record**; the benchmark/later-release loop consumes
  m-3 observed evidence via `outcome_feedback_ref`. m-4 declares which fields are observed (§9); m-3 owns how it
  observes (observer-only, positive R3 write-allowlist). **Reconciled (§9).**
- **→ m-5 (lock prerequisite, NOT optional):** the opaque **`seat_archetype`** tag (F3, §8) parameterizes the
  Layer-2 recommendation lookup; the authority-ceiling caps routable authority (§8). m-4 carries `seat_archetype`
  + the resolved `authority_ceiling` as **opaque** per-assignment fields (§5); the **concrete `seat_archetype`
  tag-space + ceiling semantics + the routing-template structures/lineup (GL-4) are m-5's (c3).** Consumer-lens
  round complete (m-5 surfaced F2/F3, folded here, bounded-additive); the c2 m-4 lock preserves m-5 ownership of
  the concrete semantics. **Surfaced, not closed.**
- **→ m-6 (M4-1 — CONFIRMED, rides the existing c1 gate; no new gate class):** ordinary `routing` stays
  **category-B (orchestrator-absorbed)** — no raise. On **`human_decision_required` / `routing_unavailable`**,
  m-4 **raises the c1-locked monotonic `HUMAN_GATE_REQUIRED`** on the routing record — this is exactly the
  **"m-4 routing-raise"** the locked m-2 contract already names as a sanctioned monotonic raiser
  (`m-2 form-schema §3`: floor = MAX across {system, agent, m-5 archetype-ceiling, **m-4 routing-raise**}) — and
  stamps a **force-A `gate_category`** on the record.
  - **The atom (readable + stamped on a *consumable* record):** `HUMAN_GATE_REQUIRED = raised` + a **force-A
    `gate_category`**. The routing-outcome states that **trigger** the raise are `human_decision_required` /
    `routing_unavailable` — these are **routing-outcome states, NOT §J2 `gate_category` members** (`human_decision_required`
    is §J2's `other`→A disposition label; `routing_unavailable` is not in §J2). **Force-A resolution (m-4-F7 /
    x3-seam-byte-integrity-F1, c6):** the escalation's category-A is **guaranteed today** by the hardcoded
    **`other` → A fail-safe** in ARCHITECTURE §J2 (an unclassified routing-outcome escalation ⇒ human-only) — **no
    §J2 change is required for correctness.** A **distinct explicit `routing_escalation` A-member** in the §J2
    default map is *recommended* for precise classification/telemetry and is routed as a **c6 cross-domain carry**
    (CTO §J2 edit + m-2/m-6 confirm) **before Step-1 PLAN** — a clarity improvement, not a design-lock dependency.
    m-6 reads `HUMAN_GATE_REQUIRED = raised` + the force-A category off the **accepted, consumable** routing record
    to distinguish an **escalated** routing gate from an **absorbed** (category-B, no raise) one — it does **not**
    depend on the `route_dispatch()` return value.
  - **The stamp path:** the routing record rides m-1 `submit()` → stamped `FROM` + lineage-accepted → the raised
    `HUMAN_GATE_REQUIRED` + A-set reason are stamped on the consumable record; m-6 projects A-set to the
    human-surface (Owner Decision Brief: recommendation + enumerated model choices). The A/B category membership
    + email projection are m-6/§J config (operator-configurable per §J J2) — m-4 names the *requirement* (a
    readable force-A atom on the record), not a new enum.
  - **R2-safe:** the raise trigger is a **routing-outcome state** (`human_decision_required` / `routing_unavailable`),
    **not** a model value — so no model-derived predicate enters the gate (R2 preserved). No new gate class is
    introduced; the escalation reuses the c1 monotonic `HUMAN_GATE` + §J `gate_category` mechanisms verbatim.
  Warm lens at consumer-review (now complete).

No orphaned primitive — every field has a named consumer.

---

## §11. Novelty statement + two honest qualifications (F5 — no overclaim)

**The contribution, located precisely:** the **seat-stamped, persisted, auditable routing/deviation artifact** —
a per-dispatch routing decision recorded as a **confusion-resistant**-stamped (m-1 stamp; **D5 residual**), lineage-gated relay carrying a justified
deviation against a declared capability-prior floor — i.e. the **port of the SR 11-7→SR 26-2 model-risk
override-register discipline into per-dispatch LLM routing**, where it has never been applied.

Two priors the design **concedes** (the contribution is NOT either of these):
- **(a) Interpretable routing already exists** — Routesplain (2511.09373), Arch-Router (2506.16655) make
  routing rationale interpretable/intervenable. Differentiator: the *persisted seat-stamped
  **deviation-against-a-declared-floor** audit artifact*, not interpretability per se.
- **(b) Non-gradient adaptation already exists** — bandit routers (C2MAB-V 2405.16587; PILOT/LinUCB
  2508.21141) update routing online without backprop. Differentiator: the *auditable, persisted decision+update
  artifact*, not the non-gradient mechanism. The "non-gradient analog of Fugu's reward" framing is hedged
  accordingly.

---

## §12. Step-1 / Step-3 split + the documented execution-fidelity boundary

`ROADMAP.md`: the routing-record contract is **Step-1**; router **execution** is **Step-3** (model-agnostic
runtime + provider adapters + the router executing the Step-1 record + the benchmark). The initial release (Step-1) **rides
existing runtimes** and does not drive its own agents.

**Documented boundary (the way the upstream protocol §10 documented deferred forgery):** the initial release enforces **declaration honesty**
(§2, all lanes, snapshot-provenance) but **not execution fidelity** (whether the launched lane actually ran the
recorded model). **This execution-fidelity gap is a documented initial-release residual (D5-class): a lane may serve a
different model than the record declares, and the initial release does not verify it** — verifying the serving model would need
outside model-attribution (m-1 channel territory), explicitly out of scope under "model = payload." **At Step-3
(future — standalone-runtime hardening, NOT an initial-release/Step-1 claim)** the record *becomes* the spawn call, so the
conductor spawns exactly the recorded model by its own control-flow — closing the gap **structurally at Step-3
only**, with no change to the record shape. The
GL-4 template-spawn already narrows the gap in practice for the initial release (each session opens with its preset model).
**Execution-dependency routed to the orchestrator as a Step-3 item** (do not design the runtime now).

---

## §13. Open questions

1. **`constraints` (budget/latency/privacy) weight in the policy** — reserved for a later cycle; the initial release is
   capability-prior-led. Confirm deferral.
2. **`prior_default_set` width** — rank-1 bucket (current default) vs a top-k accepted set before a pick counts
   as a deviation. m-4 DESIGN-internal tunable; the seam mechanism is identical either way (§9).
2a. **bucket-binding-mismatch grain** — `bucket_binding_observed` is computed per assignment; confirm the
   disposition (observe-layer bounce, same integrity class as silent-deviation, §2) is the intended grain vs a
   record-level rollup. m-4 DESIGN-internal; does not affect the gate (no model predicate either way).
3. **m-5 seam disposition** — the concrete archetype-tag + ceiling semantics + template structures/lineup
   (lock prerequisite, §10). **RESOLVED/carried** — m-5 booted and confirmed through c4/c5; the concrete
   semantics are **reserved to c3 (m-5-owned)**, and the c2 m-4 lock preserves that ownership.
4. **Record-level `evidence_integrity` rollup** — m-3 DESIGN-internal (their reconcile flagged it); does not
   affect the m-4 per-field contract.
5. **R2-boundary ratification** — **RESOLVED** — ratified at the c2 lock (readiness-fix-c1/c4 + the design-review
   reconciles; §2; both pairs aligned); no model-derived schema predicate.
6. **§2C build-carry (routing lane) — non-locking, gated to Step-1 build (step-(d)) (m-4-F1 / x1-fatal-resolution-F1, c6).**
   Two routing-lane carries are owed at **Step-1 PLAN**, not at the design lock, and are recorded in the §2C
   build-carry ledger (ARCHITECTURE §C4 + RECONCILE step-(d) + README): **(i) R2 `gate_referenceable` per-column** —
   the FieldSpec `gate_referenceable: bool` set **false** on `chosen_model` (and any model-identity column), with
   **negative fixtures** over `chosen_model` *and* single-family-bucket proxies (a bucket whose only member is one
   model must not become a model-name gate by proxy); **(ii) altitude-B per-row deviation grain** — the
   `any(routing_assignments.declared_deviated == true)` aggregate fixture across a multi-seat fan-out (no deviating
   row escapes justification). Both are **build/fixture** obligations, not record-shape or lock changes; the m-2
   `gate_referenceable` flag that backs (i) is already landed (m-4-F3, m-2 §5/§12).

---

## §14. Evidence (E1)

Locked contract — `master/ARCHITECTURE.md:66-69` (R2), `:74-77` (§5 identity≠authority), `:89-102` (§J); m-2
FieldSpec `router`/`routing_ref` — `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:60-63`;
pillar altitude-B/3-staged/Fugu-priors — the pre-build design-state export (not vendored);
roadmap Step-1/Step-3 split — `ROADMAP.md:42-65,91-92`; the implicit-routing gap + external prior-art (RouteLLM
2406.18665, SR 26-2, Routesplain 2511.09373, Arch-Router 2506.16655, C2MAB-V 2405.16587, PILOT 2508.21141) —
`c2-audit-m-4/RECONCILE-planner-20260629-190200.md`; the reconciled m-3 seam —
`c2-design-m3-m4-coord/COORD-planner-20260629-192916.md` + `…COORD-RECONCILE-planner-20260629-193400.md`;
operator grill + scope directive — this cycle + `c2-design-m-4/SITREP-planner-20260629-200500.md`.

---

## §15. c2 fold-log

Bounded, additive fold of the VP-approved m-5/m-6 consumer-lens findings (`c2-fold-m-4` dispatch
`20260630-035412`; basis `c2-consumer-reconcile/RECONCILE-orchestrator-reviewer-20260630-034321.md`). No PLAN/
IMPL; R2 and the reconciled m-3↔m-4 seam unchanged; no m-2 ownership/enum change (bounded-shape guardrail held).
- **F2 (m-5)** — added per-assignment opaque `seat_archetype` + resolved `authority_ceiling` to
  `routing_assignments` (§5), recorded for replay/audit on **both** template *and* hand-authored paths (§8). The
  per-assignment field is the replay-complete option (both m-5 seats preferred it).
- **F3 (m-5)** — named the prior-keying tag `seat_archetype`: a **distinct opaque tag** in the archetype vector,
  per-seat-at-spawn, orthogonal to m-3's `slot_in` work-archetype (per-record, conductor-classified at
  acceptance) (§8/§10). Concrete tag-space + values stay m-5-owned (c3).
- **M4-1 (m-6) — CONFIRMED.** The routing B→A escalation reaches the human-surface as a readable force-A atom on
  a consumable record via the **existing c1 monotonic `HUMAN_GATE` "m-4 routing-raise"** + a **force-A
  `gate_category`** — **no new gate class**, R2-safe (trigger = routing-outcome state, not a model value). The
  force-A category is guaranteed by §J2's hardcoded **`other`→A fail-safe** today; an explicit `routing_escalation`
  §J2 A-member is a **c6 cross-domain carry** (CTO §J2 + m-2/m-6 confirm, §10), not a lock dependency. Atom + stamp
  path named in §10 → m-6. (This was the named lock-blocker-if-unconfirmable; it is confirmed through the existing
  shape.)

## §16. c5 claim-sweep fold-log (claim-text hygiene — no mechanism change)

Applied the VP-ratified checklist (`c5-claim-sweep-architecture` `20260702-131320`) via `c5-claim-sweep-light`
(`20260702-132139`). Claim-text only; no mechanism, no locked-contract reopen. Test applied: *claims a malicious
seat is stopped* → RELABEL to **confusion-resistant + D5 residual**; *holds by the engine's own control-flow/grammar
or an observer-side selection the lane can't reach* → KEEP.
- **RELABEL (2)** — **§2.1 :69** and **§11 :360** "forgery-robust-stamped" → **confusion-resistant-stamped
  (m-1-owned stamp) + D5 residual** (a malicious same-uid seat forging a store record is out of scope). The `FROM`
  stamp rides m-1's `submit()`; m-4 mirrors m-1's relabel — m-4 does not own the stamp mechanism.
- **SCOPE (1)** — **§12 :381–:387** the execution-fidelity gap is now named a **documented initial-release residual (D5-class)**;
  "Step-3 closes the gap **structurally**" is explicitly scoped to **Step-3 only (future standalone-runtime
  hardening, NOT an initial-release/Step-1 claim)**.
- **KEEP (justified)** — **R2 "by construction / structural"** (§2 :76, §0 :39/:41, §2 :78 "proof obligation"): the
  **R2 gate-grammar invariant** — model-identity is non-gate-referenceable by the m-2 schema grammar; holds by the
  trusted engine's own control-flow, the licensed class (same as ARCHITECTURE R2 KEEP). **§4 :187** "immutable
  record cannot be reconstructed" = replay-completeness fact, not a malice claim. **Authority-ceiling** statements
  (routes-but-no-write-tool, §8) = confusion-resistant enforcement, kept.
- **KEEP — late-classified (m-4-F8; missed by the original static grep, token line-wrapped across :127–:128).**
  **§2 snapshot-provenance :127–:129** "it is `observed` whenever the snapshot is conductor-stamped (by
  construction, since the floor is the conductor's own policy)" = **observer-side selection the lane cannot reach**
  (the conductor stamps its own snapshot from its own policy ⇒ the deviation check is `observed` without lane
  vantage) — trusted-engine control-flow / observer-selected-control class, NOT a malice/write-exclusion claim, so
  KEEP (wording unchanged). Re-run the c5 sweep greps in **whitespace-tolerant** form (tolerant of line-wrapped
  `(by\nconstruction`) so line-split tokens are no longer missed.

## §17. c5 decision-⑤ fold — ODB model-name egress carve-out: R2-guard (m-4 half)

Operator decision ⑤ (`READINESS-REGISTER.md` §Operator-decisions ⑤, RECORDED 2026-06-30), folded joint with m-3
(egress scan) + m-6 (ODB). **m-4 half = the R2-guard confirmation: the carve-out does NOT touch R2.** No mechanism
change; R2 unchanged.

The carve-out (m-3/m-6-owned) exempts **only** the model-name field inside a conductor-generated operator-facing ODB
from the **confidentiality** egress scan (not the safety/content class). m-4 confirms this is R2-safe on **two
orthogonal axes**:
- **Gate-referenceability (R2's axis) — untouched.** R2 forbids a `model_*` predicate entering the schema / authority
  / lineage / work-dispatch gates. The carve-out adds no such predicate; the m-2 gate-referenceable-field allowlist
  still structurally excludes model-identity fields (readiness-fix-c4). The model-name stays **payload/bookkeeping**.
- **Egress-confidentiality (the carve-out's axis) — orthogonal.** Whether a payload value may *leave the boundary* in
  an operator-facing ODB is a different question from whether it may *gate a decision*. ⑤ relaxes a confidentiality
  scan on one display field; it does not make the model gate-referenceable.

**Human-surface ≠ machine gate (the load-bearing nuance).** The ODB renders the model-name for the **operator** to
read when exercising human judgment — an Owner Decision Brief exists to *inform* operator decisions. R2 constrains
**automated** gate predicates, not what a human is shown; surfacing model-name to the operator is transparency
consistent with the identity≠authority / human-override design, and is categorically distinct from a `model_*`
predicate in a machine gate. **Peer-bias protection intact** — no gate gains a model input. m-4 co-confirm complete;
carve-out closes with m-3 (scan) + m-6 (ODB) co-confirms.

## §18. c6 re-review cleanup fold-log (doc-only — no mechanism change, lock invariants unchanged)

Applied the c6 re-review findings routed to m-4 (`c6-fix-m-4`; `master/DESIGN-REREVIEW-2026-07-02.md`
CONDITIONAL-GO, VP-concurred). **Doc-only consistency folds into the locked doc**; Routing & Policy lock invariants
(R2, the deviation gate, the record shape, the four sanctioned by-construction claims, the confusion-resistant/
D5-residual vocabulary) are **unchanged**. Per-finding:
- **m-4-F1 / x1-fatal-resolution-F1 (B) — §2C build-carry deferral marker.** Added §13 item 6: the two routing-lane
  build-carries owed at Step-1 PLAN — (i) R2 `gate_referenceable` per-column FieldSpec bool on `chosen_model` +
  single-family-bucket negative fixtures; (ii) altitude-B per-row deviation grain. Non-locking, recorded in the §2C
  ledger; the m-2 flag backing (i) is already landed (m-4-F3).
- **m-4-F4 (B, ◆) — template-spawn authoring model.** §7: `FROM = operator` on the operator-relay channel;
  `routing_assignments` seat-scope widened to admit `operator` on `template_ref`-bearing records **only**;
  `declared_deviated` **stays declared** (not re-typed system/computed — the rejected option would reintroduce the
  readiness-fix-c4 declared-vs-observed conflation). Cross-domain mirror (m-2 §17.3 + m-7 S11) flagged to the
  orchestrator.
- **m-4-F5 (B, ◆) — `deviation_reason_code` value-set.** §5/§6: declared **config-sourced enum** (default-seeded +
  operator-configurable + `other` fail-safe, §J2 pattern); seeded the shipped default vocabulary
  (`capability_gap` · `cost_budget` · `latency_budget` · `bucket_unavailable` · `operator_directive` ·
  `experiment` · `other`). m-2 §17.3 value-set mirror flagged.
- **m-4-F6 (H) — status hygiene.** Header → **LOCKED-at-c2** (+ c5/c6 fold notes); §2 box + §9 R2-boundary marked
  **RATIFIED**; §13.3 (m-5 seam) + §13.5 (R2-boundary) marked **RESOLVED** with refs.
- **m-4-F7 / x3-seam-byte-integrity-F1 (H/M) — §J gate_category byte-reconcile.** §10/§15: the escalation stamps a
  **force-A `gate_category`** guaranteed by §J2's `other`→A fail-safe (no false claim of a non-existent A-member);
  `human_decision_required`/`routing_unavailable` re-stated as routing-outcome **triggers**, not §J2 members; an
  explicit `routing_escalation` A-member routed as a c6 cross-domain carry (CTO §J2 + m-2/m-6 confirm) before Step-1.
- **m-4-F9 (M) — ARCHITECTURE §J1 ⑤ carve-out:** already satisfied by the CTO gov-surface apply (same line as
  m-6-F5; `ARCHITECTURE §J1:99-102` already annotates the decision-⑤ typed-ODB model-name carve-out) — **verified,
  not redone.**
- **m-4-F8 (M) + m-4-F3 (B) — CTO-applied, verified not redone:** F8 late-classified KEEP (§16 :500–506);
  F3 R2 `chosen_model` re-anchor landed in m-2 §5/§12 (`gate_referenceable` flag).

**Cross-domain mirrors flagged to the orchestrator (m-4 records the decision; does not write those docs):** m-2
§17.3 (template `operator` seat-scope + `deviation_reason_code` value-set) + m-7 S11 (template author) + the CTO
§J2 explicit `routing_escalation` A-member (optional clarity carry).

**rev1 (implementer must-revise fold):** (a) **F4 §5 primary FieldSpec row** — the `routing_assignments` owner cell
now carries the same conditional `operator`-on-`template_ref`-records exception as §7 (the two seat-scope sources no
longer compete); (b) **F6 header** — the live "one R2-boundary item is escalated … for ratification at the c2 lock"
sentence past-tensed to "was escalated and RATIFIED at the c2 lock" (no resolved item described as currently
escalated).
