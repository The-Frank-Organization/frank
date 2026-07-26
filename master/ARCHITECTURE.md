# master — ARCHITECTURE OF RECORD

CTO/VP integrated architecture-of-record for **`frank`** (the conductor courier). Grows over time; the per-domain design docs
under `domains/m-<n>-*/design/` are the authoritative detail, this file is the integration spine. Status of
each section noted inline.

---

## Cycle c1 — Step-1 foundations (Trust & Identity + Forms & Determinism)

**Status: Cycle c1 CLOSED / LOCKED** — VP close-confirm `c1-joint-lock` (20260629-180934); operator §J ratified.
**m-1 Trust & Identity + m-2 Forms & Determinism are jointly locked as the frank Step-1 design-of-record.** Both
pairs design-complete-rev2 + pair-approved; the shared contract mutually re-affirmed. (Scope was AUDIT + DESIGN
only — no PROCEED-TO-PLAN / implementation; consuming domains design against this contract in later cycles.) Sources: `domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md` (rev2),
`domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md` (rev2); the `c1-*` relay trail.

### 1. The integrated model — a governed-write stamping courier over a typed-envelope store
*(Claim boundary — this c1–c3 record predates the deployment-fork decision; per **§C4.3**, every "sole-writer /
forgery-robust / by-construction" claim below is scoped **confusion-resistant**: it holds against a confused agent
acting only through the tool surface, **not** against a same-uid code-executing seat — the D5 accepted-risk. Mechanisms
are unchanged; only the malicious-lane claims collapse, per `DESIGN-REVIEW-2026-07-01.md` §2B.)*
The conductor is a **mail system that is "not an open relay"**: authenticated-submission-only, it stamps the
sender from the channel, and it is the **sole writer through the governed `submit()` path** of an append-only relay store. The canonical object is a
**typed envelope** (m-2); on-disk markdown is a rendered projection only. The SMTP **security model** (not the
literal wire) is the idiom: MSA-authenticated `submit()`, envelope `FROM` set by the courier (SPF-analog),
`Received`-style courier-stamped lineage, **governed-write** Maildir mailbox projection, `project()` retrieval, and DKIM/DMARC
mapped to the deferred `certification` seam / the one-FROM rule.

### 2. m-1 — Trust & Identity (the TCB)
- **API:** `submit` (sole *governed* write path) / `project` (scoped inbox read) / `read` (lineage/migrator queries) /
  `mint_seat` (conductor-only; attestation folded in — no public `attest()`).
- **Identity is conductor-stamped (confusion-resistant):** `FROM`/`ROLE` are `lookup(connection)`, stamped from one certified channel —
  the seat has **no tool** to write the trusted field (a confused agent cannot forge it; **D5 accepted-risk**: a same-uid
  code-executing seat can write store files directly — out of scope, §C4.3). Closes the three surveyed self-asserted-identity gaps (stock-protocol
  agent-`FROM`, jcode `assign_role`/`from_session`, claude-code self-`from`) for a confused agent.
- **Invariants (proved §4):** **I1** governed-write / store-isolation (the sole *governed* write path is `submit()` —
  confusion-resistant, not malicious-lane containment), **I2** channel-isolation (**confusion-resistant**
  `FROM`), **I3/DI-5** observe-integrity (the conductor reads lane ground-truth from *outside* the lane, so
  observed evidence is conductor-observed, not lane-supplied). Deployment invariants **DI-1..DI-5** constrain
  the Step-1 realization.
- **Transport (operator grill):** Option A — minted per-seat credential over an isolated per-seat connection
  ⇒ **confusion-resistant** identity (a confused seat has no tool to forge `FROM`); **"forgery-robust by construction"
  is the shelved wrap-hardening milestone** (GRILL-LOCK D3), not the Step-1 claim — earned only if a later operator-gated
  spike meets DI-2/DI-5.
- **Operator/special address:** `operator` is a first-class minted-address-space entry — a delivery target +
  a special stamped `FROM` over the operator-relay channel (not a lane-minted credential; a **confused** lane has
  no tool to submit as `operator` — confusion-resistant; **D5 residual**: a same-uid seat writing `FROM:operator`
  to the store directly is out of scope, §C4.3).
- **Non-re-cut path:** Step-1 minted-token → standalone SO_PEERCRED/mTLS/SPIFFE swaps only the attestation
  backend; verbs, envelope, store shape, stamp semantics invariant.
- **Crypto-hardening (deferred):** courier-as-TCB for the initial release; Merkle/append-log hash-chain reserved (joint with
  m-2's `certification`).

### 3. m-2 — Forms & Determinism (the envelope)
- **Three-layer SMTP envelope:** certified system envelope / typed FieldSpec headers / free body + `X-`
  overflow. **One-FROM rule** (no lane-writable display-FROM — closes the DMARC two-FROM confusion for a confused agent).
- **Carrier:** a bespoke **FieldSpec registry** (ownership / seat-scope / consumer / lineage-role first-class;
  JSON-Schema vocabulary for the type/enum/required-when sublayer) — the single source the tool, courier, and
  linter read.
- **Field-ownership + fill-time authority:** `system` / `seat_scoped_enum` / `agent_enum_pick` / `free_text`;
  a forbidden option is *absent* from the seat's form (authority at FILL, not post-hoc). `parent_picker` /
  `recipient_picker` are the system-candidate-set hybrids. `required_when` is a **bounded** boolean over a
  closed atom set (never Turing-complete); `slot_in` is a **reserved opaque atom** (m-5 owns the tag-space later).
- **Linter refactor:** all 62 upstream checks classified — ~33 prose-only **dissolve**, ~16 **form-validation**,
  ~13 **cross-relay lineage engine**. Strict form-only submit (legacy markdown read-only).
- **Write path — reconciled seam (`readiness-fix-c1` — ✅ CLOSED 2026-06-30: CTO-arbitrated · pair-folded (m-1
  `submit()` + m-2 `send()`) · CTO re-verified byte-consistent · **VP closure co-sign 20260630-230335**):** one
  atomic accept, gates as **pre-append validation** — resolve/stamp →
  **form-validation** → **cross-relay lineage gate** (validates the candidate against the persisted `accepted`-graph
  with the candidate held **in-courier** — **no persisted `submitted` limbo**) → on pass, **one atomic `accepted`
  append** + deliver; on fail, a **terminal `rejected` evidenced record** + bounce. Authority-bearing records are
  non-consumable until the lineage gate passes — preserves the upstream "blocks before dispatch." **Step-1 = store + form
  + lineage** (named build boundary); the m-3 **observe-as-send** gate is a **reserved additive Step-2 hook**, not
  required for Step-1. *(Supersedes the prior two-state `submitted→accepted` path.)*
- **Versioning:** `schema_version` stamped immutably; read-time migrator registry shipped with **zero migrators**.

### 4. The shared contract (the seam — mutually re-affirmed: R1/R2/R3)
- **System-filled `PARENT` (m-1) strengthens the lineage engine (m-2) from confusion-robust → confusion-*resistant***
  (the seat cannot supply `PARENT` — a confused agent cannot forge lineage; D5 residual as §C4.3) — the load-bearing convergence.
- **R1 operator/special-address:** `operator` is a `recipient_picker` member with a defined operator-FROM path.
- **R2 routing decision = a separate seat-stamped routing relay** (m-2 FieldSpec + record-kind + accepted
  semantics; m-1 admits the accepted routing relay into the conductor-derived `parent_picker` candidate set for
  the dispatch it routes). The dispatch references it as **provenance/bookkeeping** — **model is never a gate
  input** (no `model_*` predicate enters the schema gate).
- **R3 observe-integrity:** observed evidence carries `evidence_integrity {observed | self_reported}`; the m-3
  hook is observer-only (a positive write-allowlist: it writes only the closed m-3 observed/computed set + a veto).
- **`certification` null-reserved** in every record (joint crypto-hardening deferral with m-1's Merkle).

### 5. identity ≠ authority — RATIFIED (closes m-1 open-Q #2)
m-1 owns **who** (confusion-resistant identity + governed-write store + `FROM`/`ROLE` stamp). **What a stamped seat may
do** is m-4 routing/policy + m-5 archetype-ceiling, keyed to the stamp (anti-confused-deputy). m-4 accepted the
boundary; the CTO/VP ratify it here.

### J. Operator-judgment items — RATIFIED (operator, c1 close)
**J1 — ODB `on_timeout` = `hold_and_resummon`, never auto-approve.** The conductor parks a gate and re-summons
on an escalating cadence; it never resolves a governance gate without the operator. Per-gate overrides are
**monotonic** (a gate may make itself more conservative, e.g. `block`; never opt an A-gate into
`take_recommended`). No hard `decision_deadline` by default (per-gate opt-in).
- *Forward (m-6 scheduler + m-3 egress):* **away-mode external bridge** — when the operator declares "out," the
  conductor mirrors A-bucket gates to the operator's real inbox (e.g. Gmail) and accepts the email reply as the
  verdict. **Opt-in, not the default** (default = local in-app inbox); it is the *first external send*, so it is
  gated by the **fail-closed egress scan** (secrets / PII / model-names — the decision-⑤ typed-ODB
  operator-facing model-name carve-out excepted, m-3 §7) before anything leaves.

**J2 — `gate_category` (default set; operator-CONFIGURABLE).** The membership, the A/B mapping, and the
protected-branch set are **operator-configurable policy** — a default ships, and the design accommodates
customization: `gate_category` is a **config-sourced enum** with the hardcoded `other`→A fail-safe; the A/B map
+ protected-branch set are m-6 config. *Forward (m-6 + config); the locked c1 `gate_category` slot does not
preclude it.*
- **A (reserved-to-human):** `merge_to_protected` · `irreversible_write` · `residual_risk_acceptance` ·
  `live_verify_skip` · `ceremony_downgrade` · `authz_security` · `product_semantics` · `scope_expansion` ·
  **`routing_escalation`** *(LANDED 2026-07-06 — the §C4 owed carry executed as prescribed: m-4 confirmed token +
  A-bucket ("strict precision improvement, no behavior change"), m-6 confirmed the surface (bucket A → ODB + park +
  J1, non-suppressible), m-2 supplied the exact `named_enums` mirror delta (append to `gate_category_A`; insert into
  `gate_category` before `other`; no B change; no FieldSpec-row change; MINOR/additive) — the three `s5-fidelity`
  answer relays 2026-07-06; the registry delta itself lands in s5-a's registry pass. Byte-distinct from
  `routing_unavailable` — the route_dispatch **outcome state** (§C2), never a §J2 member — and from
  `human_decision_required`, the `other`→A disposition label.)*
  *(Until the registry delta lands, routing-escalation force-A continues to flow via the `other`→A fail-safe below —
  correctness never depended on the member; m-4 §10:361-369.)*
- **B (orchestrator-absorbed):** `merge_feature_to_feature` · `routing` (model = payload) · `sequencing` ·
  `scope_within_bounds`.
- **Fail-safe:** `other` → A (`human_decision_required`) — unclassified ⇒ human-only, never auto-absorbed.
- **Merge split (operator refinement):** the merge bucket derives from the **target branch** + the
  operator-configured **protected-branch set** (default = prod branch(es) + the main working/dev/integration
  branch). Merge into a protected branch = **A** (affects shared dev / prod); a feature→feature merge, incl.
  tiered/stacked, = **B** (autonomy without affecting others or prod).
- **Direction — RAISE-ONLY (decision ③, RATIFIED c5 2026-07-02; owner-authored m-6 `c5-fold-decision-3`, m-6.impl approved).**
  An agent-pick of `gate_category` is **monotonic toward A** — it may escalate (more operator oversight), **never**
  de-classify an A-worthy decision down to B. A **known-A detector** (config-owned membership, operator-tunable per J2)
  forces ≥A regardless of agent pick, so an A-worthy decision mis-tagged B is **raised, not silently orchestrator-absorbed**
  (the most direct operator-not-surfaced vector, closed). Rides the existing HUMAN_GATE monotonic-raise (m-7 enforces at
  fill/submit) atop the J2 map + the CQ-3 A-floor — **no new gate class, no new mechanism**; a direction constraint on the pick.

### 6. PLAN carry-forwards (for the future build cycle — NOT this phase)
Recorded so the eventual build inherits them; none reopens design.
- **m-1:** DI-2 + DI-5 each tested as independent isolation/read-vantage properties (the "fork-2 infra call");
  `submit()` as one atomic conductor-owned commit (TOCTOU); seat-credential lifecycle (generation/rotation/
  revocation); m-3 hook observer-only resolving to the explicit m-2-declared field set; operator-relay channel
  unreachable by lanes; observed evidence = a point-in-time conductor snapshot on an immutable record.
- **m-2:** negative fixtures (monotonic-floor tamper rejection; closed-enum bucket routing; `X-` attempted
  gate-input rejected; a routing-ref case proving model is not a gate input); `completed_proof` readers respect
  `evidence_integrity` (self_reported ≠ observed proof).

**On VP approve + operator ratification of §J, Cycle c1 closes with these two foundations locked as the frank
Step-1 design-of-record.** Consuming domains (m-3/m-4/m-6) design against this locked contract; runtime/product
(m-7..m-12) are future cycles per `ROADMAP.md`.

---

## Cycle c2 — Step-1 runtime-intelligence (Observation & Evidence + Routing & Policy)

**Status: Cycle c2 CLOSED / LOCKED** — VP co-sign `c2-lock` (20260630-043859). **m-3 Observation & Evidence + m-4
Routing & Policy are jointly locked as the frank Step-1 runtime-intelligence layer** atop the locked c1 substrate.
Both design-complete + pair-approved (m-3 r1 `211003` + fold `040633`; m-4 r1 `203329` + fold `040641`); consumer
lenses m-5 (narrow) + m-6 cleared; the m-3↔m-4 seam reconciled both sides; all consumer folds (F1/F2/F3/M4-1)
bounded-additive, implementer-re-approved, **no m-2 micro-fold**. Operator-directed scope (GL-4 / m-5 narrow
engagement) is **operator-directed by current session context** (VP c2-lock-prep provenance note; no `FROM:operator`
relay required). Authoritative detail: the rev2 docs under `domains/m-3-observation-evidence/design/` +
`domains/m-4-routing-policy/design/`; the `c2-*` relay trail.

### C2.1 — m-3 Observation & Evidence (the observe-as-send gate)
**Observe-AS-send:** the conductor performs the phase-shaped done-predicate observation from *outside* the lane
(DI-5) inside m-1 `submit()`, after m-2 form-validation, before append/accept; **observer-only** against the
locked R3 write-allowlist (writes only the closed m-3 set + a veto). "No clean observation, no relay leaves" —
closes the universal self-reported-done gap (no surveyed system observes done from outside as a send-gate).
- `observe_gate()` → `observe_result{ predicate_result, veto, achieved_evidence (E-ladder), evidence_integrity
  {observed|self_reported}, egress_scan_result, degradation_notes }`; `EVIDENCE_TARGET` (intent) split from
  `achieved_evidence` (fact). E1/E2 passive Step-1; E4 = live-verify (operator-observed); executable-claims
  reserved (operator-gated registry descriptors, never arbitrary commands).
- **Egress / content-safety gate:** fail-closed at the **conductor-governed egress chokepoint** (the conductor's local outbox is the only egress the *governance system* offers — a governance-surface claim, **not** system-level sole-egress; **D5 residual**: a same-uid shell/curl network bypass is out of scope, §C4.3), dormant until
  the first external send (m-6 away-mode §J1); promotes the agent-scripts/claude-code rule sets; operator-config.
- **Archetype done-predicates** keyed on the opaque `slot_in` **work-archetype** atom (two-axis, C2.4) —
  conductor-classified at work-record acceptance, **non-lane-writable** (a confused seat has **no tool/verb** to set or
  re-tag it — confusion-resistant, closes re-tag-to-escape; **D5 residual**: same-uid direct store write out of scope, §C4.3).
- Record-level `evidence_integrity` rollup = m-3-internal (the per-field tag stays two-value; R3 unchanged).

### C2.2 — m-4 Routing & Policy (the governance-record layer)
**Routing is a first-class, recorded, justifiable governance decision** — a port of the SR 11-7→SR 26-2
override-register discipline into per-dispatch LLM routing. **Model = payload/bookkeeping, never a gate input (R2).**
- **`routing_decision` record** (m-2 FieldSpec consumer): `record_kind=routing` + `routing_record_kind`;
  `routing_assignments` (seat_scoped to planner/orch-planner; + per-assignment opaque `seat_archetype` /
  resolved `authority_ceiling`, C2.4 F2); `capability_prior_snapshot` (system, both layers, replay-complete);
  `justified_deviation` + `deviation_reason_code` (`required_when` declared_deviated); `outcome_feedback_ref`
  (null-reserved for a later release); `template_ref` (GL-4).
- **Two-layer capability prior (GL-1):** bucket membership (operator-owned; churns on model change) + bucket
  recommendation ((role,task_tag)→ranked buckets; later-release-tunable). Deviation computed **bucket-vs-bucket** (model
  identity NOT read). Snapshotted into the record.
- **`route_dispatch()` fail-closed:** no acceptable route → `human_decision_required` / `routing_unavailable`,
  never a silent default. **3-staged policy:** prior floor → justified deviation (initial release) → later-release outcome feedback
  (forward hook; sample drawn independent of `deviated`).
- **identity ≠ authority at the routing layer:** authority-ceiling-at-spawn caps routable authority; the router
  refuses to staff a forbidden seat (fail-closed). The ceiling keys on the opaque `seat_archetype` atom (C2.4).
- **GL-4 routing templates:** m-4 owns the template record mechanism (pre-filled `routing_decision` + `template_ref`;
  no-bypass invariant — an off-floor template assignment still records a deviation); **m-5** owns structures/lineup;
  **conductor-core** owns pane-spawn via **existing tmux/zellij/OS-terminal** (Step-1-consistent, not our TUI).
- **Novelty (F5):** the seat-stamped, persisted, auditable **deviation-against-a-declared-floor** artifact — not
  interpretable routing in general (Routesplain/Arch-Router), nor non-gradient adaptation in general (bandit routers).

### C2.3 — the m-3↔m-4 seam (reconciled both sides; R2-preserving by construction)
`routing_decision.deviated` = planner-**declared** boolean (the only model-touching atom the m-2 schema gate
reads, via `justified_deviation.required_when: declared_deviated==true`), cross-checked by a conductor-derived
`deviated_observed` (**bucket-vs-bucket; does not read model identity**). The four declared×observed cells deliver,
**EXCEPT silent deviation** (declared=false, observed=true) → m-3's **generic declared-vs-observed integrity veto**
(clean-tree class): an observe-layer bounce → honest re-declaration → m-2's existing justification `required_when`
fires. **Block the dishonesty, never the deviation** — no new gate class; **no model-derived predicate enters any
gate (R2 by construction)**; rides **snapshot-provenance** (not DI-5) so it holds for opaque lanes; an
internal-consistency check, not serving-model forensics (model = payload). The routing record is an evidenced
record (m-3 observer-only profile of the single R3 atom; m-4 declares observed fields, m-3 owns how;
`outcome_feedback_ref` binds the downstream observe-atom in a later release).

### C2.4 — the archetype tag-space: TWO orthogonal axes (consumer-lens folds, implementer-approved)
- **F3 — two axes:** `slot_in` = **work-archetype** (m-3, per-work-record) vs `seat_archetype` =
  **seat-archetype** (m-4, per-seat-at-spawn). A seat carries both; recorded as **distinct opaque tags**
  (m-4's "archetype vector" accommodates it).
- **F1 — provenance (split):** both conductor-owned / **non-lane-writable** (confusion-resistant — no lane tool writes it; D5 residual per §C4.3). `seat_archetype` =
  spawn-time; `slot_in`/work-archetype = conductor-classified at work-record **acceptance** (not spawn-fixed —
  long-lived seats may move bugfix→refactor→migration).
- **F2 — record home:** `seat_archetype` + resolved `authority_ceiling` recorded **per-assignment** on
  `routing_assignments` (replay-complete; both template and hand-authored paths).
- **BOUNDED:** both atoms stay **opaque**; concrete tag-space / invariant selection / ceiling semantics are
  **RESERVED to m-5 (c3)**. **De-locking note (operative):** any archetype values named anywhere in this section
  or the m-3/m-4/m-5 design docs (e.g. a `refactor` work-archetype, a `sensor` seat-archetype) are **non-locking
  candidate/example vocabulary, c3-owned** — the c2 lock binds only the two opaque atoms + the c3 reservation,
  nothing more. No concrete Step-1 values, no `required_when`/`visible_when` on tag values, no m-2 ownership
  change → **no m-2 micro-fold** (additive within the reserved opaque-atom space + m-4's own record).

### C2.5 — m-5 narrow engagement (surfaced; RESERVED to c3)
The c2 lock records m-5's bounded outputs as **proposals reserved to c3** — it does **NOT** lock the full m-5
archetype system. m-5 ownership of the **concrete tag-space, invariant selection, default per-archetype gate
composition, template structures/lineup, and authority-ceiling semantics is RESERVED to c3.**
- **initial-release template lineup (proposed):** T1 **Solo**, T2 **Adversarial Pair**, T3 **Sensor** (the orchestrator+N-pairs
  "conductor" template deferred to c3).
- **Sensor archetype** (the `/btw` side-question made spawnable): read-only ceiling, tool-blocked, single-turn,
  shared read-only fork, non-interrupting parallel, routed `fast_cheap`. **Integrity split:** answer *content* =
  `self_reported`/advisory + never gate-bearing; observable *metadata* (tool-blocked, no-source-actions, 1-turn)
  = `observed`.

### C2.6 — recorded ratifications + cleared gates
- **R2-boundary:** VP-ratified at lock-prep (`c2-lock-prep` 212213) — recorded here (narrow: bucket-vs-bucket;
  `chosen_model` readable by the observe layer only as payload; no model-derived predicate in any gate).
- **M4-1 (m-6 consumer):** the routing B→A escalation rides the **c1-locked monotonic HUMAN_GATE "routing-raise"**
  (`HUMAN_GATE_REQUIRED=raised` + routing `gate_category ∈ A-set` on the accepted record; ordinary routing stays
  category-B) — readable/stamped, no new gate class, R2-safe. **m-6 consumer-lens cleared** (reader-has-a-writer
  holds; m-3 built the egress gate + ODB schema for m-6's away-mode bridge; c1 Q-B/G3 continuity honored).
- **identity ≠ authority** extended to the routing layer (C2.2).

### C2.7 — interjection (forward requirement; ROADMAP cross-cutting)
First-class **steer / side-question / interrupt** (`ROADMAP.md` interjection rule + Step-3 exit test). steer+interrupt
baseline; the **side-question** = the m-5 read-only sensor archetype (C2.5) routed by m-4. Ownership: m-6 surface +
runtime mechanism + m-5/m-4 archetype/routing. Detail + cites in `jcode-ux-notes.md` (negative look) +
`codex-notes.md` (positive look) — local reference copies, not vendored. Ratified-when-designed; gates nothing now.

### C2.8 — PLAN carry-forwards (c2; for the build cycle — NOT this phase)
- **m-3:** observe-as-send as one atomic gate inside `submit()` (TOCTOU); the executable-claims conductor-EXECUTES
  mechanism (operator-gated registry descriptors); egress rule-set activation on the first external send; the
  opaque-lane `self_reported` floor; the `slot_in` classification-ordering (when/how the conductor classifies
  work-archetype at acceptance).
- **m-4:** router execution = Step-3 (provider adapters; the record becomes the spawn call); the benchmark/later-release
  outcome loop (sample independent of `deviated`); GL-4 pane-spawn via existing multiplexer infra; `seat_archetype`
  authority-bearing **only** within the routing mechanism (a watch-note — anything stronger routes as an
  m-2-adjacent fold).
- **m-5 (c3):** the full archetype system — concrete tag-space, invariant maps, gate composition, template
  schema/lineup, ceiling semantics, the conductor template.

**On VP c2 co-sign, Cycle c2 closes with m-3 + m-4 jointly locked as the frank Step-1 runtime-intelligence layer;
m-5's full archetype system + m-6's full human-surface design are c3.**

---

## Cycle c3 — Step-0 completion: the human surface (Workflows & Archetypes + Human Surface & Scheduler)

**Status: Cycle c3 CLOSED / LOCKED** — VP co-sign `c3-lock` (20260630-191315), co-signing both the lock (Q1) and the
**C3.6 integration capstone** (Q2). **m-5 Workflows & Archetypes + m-6 Human Surface & Scheduler are jointly locked**
as the final two Step-0 design domains — **completing the six-domain Step-0 design-of-record** (c1 foundations + c2
runtime-intelligence + c3 human-surface). Both design-complete + pair-approved (m-5 `133831`; m-6 `133839` r2 +
correction re-approve `185424`); the m-5↔m-6 seam reconciled both sides (`123022`⊕`131856`); Seam C (the m-1
away-token bridge) resolved **A** + folded (`182600`/`183008`/`183345`); both GRILL_LOCKs folded; **no c1..c2 contract
reopened** (m-1's certification-seam activation is additive — C3.4). Authoritative detail: `domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md`
+ `domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md`; the `c3-*` relay trail.

### C3.1 — m-5 Workflows & Archetypes (the governed expansion-slot)
**Net-new contribution (proved):** an **archetype** is one governed preset binding **{topology + gate-set +
authority-ceiling-at-spawn + observe-invariants + routing-prior}** as a single declared, seat-stamped, append-only,
auditable unit — riding m-1's append-only lineage with a per-assignment recorded, fail-closed-enforced ceiling (F2).
No surveyed system (internal or the 2023-2026 external landscape) binds these five into one auditable unit; the
parts are promoted (upstream panel pattern + AUTHORITY/tier enum; codex presets-as-data + graph-query-as-projection;
claude-code agent-type + `sideQuestion`), the **integration is built**.
- **Tag-space (LOCKED, `lower_snake_case`):** `slot_in` work-archetype `{feature_extension, refactor, cleanup,
  bugfix, migration, research_synthesis, qa_review, docs_chore}`; `seat_archetype` `{solo_worker, planner,
  implementer, reviewer, sensor, orchestrator_lead}`. **`actuator` = a DERIVED mutating-ceiling class, NOT a literal
  seat value** (GL-1); the literal `single_bounded_action` seat = Step-4/5 carry.
- **Composition + observer-selected invariants:** effective gate-set = `seat_archetype` (spawn-fixed) ⊗ `slot_in` (per-work-record,
  conductor-classified-at-acceptance). The two load-bearing invariants (`refactor` test-files-unchanged; `bugfix`
  red→green) hold **because the invariant is selected by the observer (conductor-side), not the observed lane** — a genuine
  observer-side control property (the observed lane cannot reach the selection), riding
  locked F1 (non-lane-writable classification — no lane tool writes it) + the locked observe-as-send gate; no new mechanism.
  **D5 residual**: a same-uid seat writing a false store record directly is out of scope (§C4.3).
- **Authority-ceiling vector (LOCKED, GL-2):** an **OPEN named-axis map** (not a fixed tuple), 3 initial-release axes —
  `write` (read_only<write_scratch<write_feature<merge_feature<merge_protected[A-gated]) ⊥ `dispatch` (none<route)
  ⊥ `tool` (none<read<exec) — a **partial order** (an `orchestrator_lead` routes but cannot write). A `seat_archetype`
  pins a MAX per axis at spawn (F2); templates tighten below, never loosen; `route_dispatch()` fail-closes above any
  axis. **`external_send` = a DEFERRED-but-modular additive 4th axis** (absent-default `none`/fail-closed; arrives
  with the literal actuator, Step-4/5; no re-cut).
- **Templates (LOCKED):** T1 `solo`, T2 `adversarial_pair` (planner↔implementer, diverse-model), T3 `sensor`; the
  **conductor/N-pair template DEFERRED to Step-5**. **Sensor (full):** read-only ceiling, tool-blocked, single-turn,
  non-interrupting fork, separate surface, `fast_cheap` (operator-overridable per-question escalation); integrity
  split (content self_reported/never-gate-bearing; metadata observed) — a **declared, seat-stamped** sensor.

### C3.2 — m-6 Human Surface & Scheduler (the promote-and-bind projection)
**A thin local-first projection over locked m-1..m-4 records** — not a new gate system, not a new schema, not a TUI
(Step-4). Three lanes (governance / collaboration / interjection).
- **Gate→email buckets (direction-explicit; no-bucket-without-a-writer):** **A** verdict-required (`gate_category∈A`
  ∪ `HUMAN_GATE=raised` ∪ `egress_scan_result=blocked`) → ODB+park+J1; **B** orchestrator-absorbed (live local
  digest, G3); **C** CC-FYI; **D** observe-bounce (author-facing, `delivery_state=rejected`+`failing_edge`; CQ-4 token —
  `bounced` retired). Only
  **A+C** reach the operator (alert vs notification). **Egress/D precedence** (F2 fix): egress is evaluated only at
  the external-send chokepoint on an already-A gate (→ `egress_blocked` → local resummon, never auto-send); D bounces
  at acceptance — different stage + `failing_edge`, mutually exclusive.
- **Owner Decision Brief:** **promote** the m-3-owned 7-field schema; m-6 designs **render + capture** — bounded
  choices as **buttons** (`agent_enum_pick`; pick only a legal verdict) → operator-FROM verdict relay; J1
  hold_and_resummon + refresh-before-resummon + never auto-approve. **"Elaborate more" = an operator-initiated
  read-only context-preserving fork** (governance→collaboration bridge): the gate stays parked; the fork's only
  output is the decision relayed back (write-capable fork DECLINED, G2); re-observe-on-resume.
- **Scheduler:** the durable **signal+await+timer** shape on the governed-write store as the checkpointer (a parked lane
  consumes nothing); a **7-state machine** (`active→parked_waiting_human→resummon_due→replied_pending_validation→
  resumed` + `bounced_repair` + `egress_blocked`); resummon **escalates the channel, never the verdict** (J1); no
  hard deadline by default.
- **Away-mode bridge (opt-in; the one new design):** A-gates mirror live + B a frequent digest (default sub-daily) +
  C local (G1); first external send **egress-gated fail-closed** (m-3 §7). The inbound verdict-token bridge = Seam C
  (C3.4).
- **Meeting-collaboration + interjection:** gate→lane routing off locked `{phase, GRILL_REQUIRED}` (design/grill
  gates → the meeting lane, **no-compress**, re-observe-on-resume; attach mechanism Step-3/4); the **interjection
  host** (steer/side-question/interrupt composer; the side-question answer on a separate non-lane surface — the
  explicit fix vs jcode's side-panel).

### C3.3 — the m-5↔m-6 seam-of-record (declare-before-bind; reconciled both sides)
`c3-design-m5-m6-coord`: **m-6 bind `123022` ⊕ m-5 confirm `131856`** (the `125604`/`131747` `{verdict,fyi,collaborate}`
excursion crossed the bind and was **retracted both sides**). m-5 DECLARES, m-6 BINDS. **Human-mode = two orthogonal
layers:**
- **posture** `{interactive, away, unattended}` — per-template/per-seat default; rides the F2 per-assignment home
  (no new m-2 field).
- **surface_intent** `{progress, review_checkpoint, advisory, result}` — **conductor-DERIVED, TOTAL** (`progress⇐
  otherwise`), over **non-gate records only** (like `record_integrity`; no new m-2 field — **C2.4 holds, no
  micro-fold**). **Gate-bearing records carry NO surface_intent** — delivery binds off the **locked**
  `gate_category`/HUMAN_GATE/J1 (so there is **no `verdict` value** — no duplicate-mechanism-as-vocabulary).
- **`away_bridge_eligible`** = an **m-6-owned per-gate boolean** (policy); m-5 declares only the `away` posture. A
  hard per-archetype never-bridge ceiling is a **reserved future m-5 hook**, not the initial release.
- **Interjection (Seam B):** m-6 surface + m-5 `accepts_interjection`-by-longevity + m-4 `fast_cheap` routing +
  runtime injection. (Non-blocking follow-up OQ-2: the elaborate-more fork's posture — a bounded m-5 COORD note.)

### C3.4 — Seam C: the m-1 away-mode inbound verdict-token bridge = A (additive; NO c1 reopen)
The away-mode bridge converts **untrusted SMTP/IMAP-inbound → a trusted operator-channel verdict record**. The
**first conditional-upstream-contract-check** (VP F3) resolved **(A) m-1 owns the mint/verify** (`182600`, accepted
`183008`, folded `183345`). **A is forced** by two locked m-1 invariants: **DI-1** (one-shot nonce-burn = an atomic
conductor-owned store append; confused lanes have no store-write tool — D5 residual: same-uid direct store write out of scope, §C4.3) + **DI-2** (the signing key is a TCB secret; m-6 custody = a
forbidden second identity authority). It is **additive** — the first activation of m-1's **already-reserved
`certification` (DKIM-analog) field** (present-but-null; email is the channel-stamp-unavailable case that seam was
reserved for): **no field re-owned, no fifth public verb, no schema/on-disk re-cut; c1 not reopened.**
- **m-1 (TCB) owns:** `mint(decision_id, seat, choice, expiry)` on egress; `verify(token)` on return
  (sig→audience→expiry→nonce-unused→seat-matches; nonce-burn = atomic conductor append; seat-match to m-1's minted
  address space) → on pass, stamp `FROM: operator` on the operator-relay channel.
- **m-6 (bridge, outside the TCB) owns + calls it:** egress trigger + token-bearing email render + POST-not-GET
  receipt + bucket routing + away UX; **supplies `expiry`** (validity-window = m-6/operator policy), which m-1
  enforces. Recorded as an **additive later-step build carry** (C3.7), not a c1 re-lock.

### C3.5 — GRILL_LOCKs (operator, 2026-06-30)
- **c3-grill-m-5** (GL-1..6): actuator-derived; 3-axis open-map ceiling + modular `external_send`; ship all 3
  read-only work-archetypes; the four-class non-gate `surface_intent`; `lower_snake_case`; sensor `fast_cheap`+escalation.
- **m-6 §9** (G1..G5): away-mode A-live + B-frequent-digest (default sub-daily); buttons + read-only elaborate-more
  fork (write-capable declined); Bucket-B live non-interrupting digest; escalate-channel-never-resolve; delegation
  deferred; §J config surface m-6-owned / values operator-owned.

### C3.6 — Integration-completeness capstone (the six-domain composition — the pre-close gate, VP F5)
**Certified: the six domains compose into a coherent, build-ready Step-0 design-of-record; no contradictions across
§1–§C3.**
- **The consume-graph is acyclic + writer-backed.** m-5 parameterizes m-2 (opaque atoms) + m-3 (observe-invariant via
  `slot_in`) + m-4 (ceiling+prior via `seat_archetype`, template record) + conductor-core (pane-spawn); m-6 projects
  m-1 (addressing/stamp/§J) + m-2 (gate fields/ODB slots) + m-3 (observe/egress/record_integrity/ODB schema) + m-4
  (routing escalation) + m-5 (the seam). All **reads** of locked contracts; **no writes to a locked domain; nothing
  reopened.** Every m-6 bucket/surface has a locked writer (no-consumer check passes both docs).
- **The three seams close.** A (human-mode two-layer, both docs identical) · B (interjection host, matched) · C
  (m-1 away-token = A, additive). The one cross-domain new item (m-1's certification-seam activation) is additive +
  recorded (C3.4/C3.7).
- **Locked-contract invariants preserved.** R2 (model = payload; no `surface_intent`/human-mode reads model identity)
  · R3 (no new m-2 field; `surface_intent` conductor-derived like `record_integrity`) · C2.4 (no m-2 micro-fold over
  tag values) · J1 (never auto-approve; escalate-channel-never-verdict) · DI-1/DI-2 (Seam C forced by them) ·
  identity≠authority (the ceiling vector).
- **Deferrals are recorded, not gaps.** Step-4/5 carries (literal actuator + `external_send` axis + standalone-runtime
  enforcement + the TUI + meeting-attach + the token-bridge build + digest scheduler) are build-cycle dependencies,
  not c3 blockers (C3.7).

### C3.7 — PLAN carry-forwards (c3; for the build cycle — NOT this phase)
- **m-5:** the archetype registry as config-data (templates + tag-space + ceiling maps); the
  conductor-classifies-`slot_in`-at-acceptance ordering; the open named-axis ceiling map + the additive `external_send`
  axis; the literal single-action actuator + standalone-runtime ceiling/topology enforcement (Step-4/5); negative
  fixtures (re-tag-escape blocked; fail-closed above-ceiling; no sensor in-place upgrade).
- **m-6:** the ODB render + buttons/elaborate-more capture; the park/wake durable-store transitions + m-3
  re-observe-on-wake; the away-bridge egress activation + digest scheduler; the fork primitive (runtime-owned) for
  elaborate-more + side-question; the meeting-attach (Step-3/4); the §J config surface; the Step-4 TUI (codex stack).
- **m-1 (additive, Seam C):** the inbound-token **mint/verify** surface — signing-key custody OS-isolated from all
  lanes (DI-2); nonce-burn = atomic conductor store append (`submit()` TOCTOU close); `certification`-seam activation
  scoped **inbound-verdict-only** (general DKIM/Merkle hardening stays the deferred courier seam); POST-not-GET +
  fail-closed verify. **Does not reopen c1.**

**Cycle c3 CLOSED (VP co-sign `c3-lock` 20260630-191315): m-5 + m-6 jointly locked — the six-domain Step-0
design-of-record is COMPLETE. The C3.7 build carries inherit to the future build cycle only; this lock grants no
PLAN/IMPL. Next is the operator's call per `ROADMAP.md`: the PLAN phase / Step-1 conductor-core build is a separate
operator-opened gate.**

---

## Cycle c4 — the CONDUCTOR-CORE substrate (the engine the six domains ride on)

**Why c4 exists.** The 2026-07-01 adversarial pre-build review returned **NO-GO**: the six domains designed the
*contracts*, but the **running program that executes them** — serialization, crash-atomicity, recovery, config
integrity, the interface guardrail — was nobody's domain and had no design. c4 stood up the 7th pair (**m-7
Conductor-Core**) to design that substrate (`DESIGN-REVIEW-2026-07-01.md` §2A). **LOCKED 2026-07-02** — VP co-sign
`c4-design-m-7/RECONCILE-orchestrator-reviewer-20260702-040327`; `DESIGN_LOCK_ID c4-design-m-7-lock`; full design-of-record
`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md`.

**The one-line boundary.** Conductor-core owns the **ENGINE** (how things run); the six domains own the **CONTRACTS**
(what is valid / required / gated); the engine **executes** those contracts — right order, right atomicity, behind the
right interface — and owns no policy.

### C4.1 The engine
- **One process, one commit loop.** N per-seat MCP channel handlers act as mail-carriers (authenticate the channel =
  m-1 identity binding, append to a durable intake journal, await a typed outcome); a **single-threaded serialized
  commit loop** runs every mutation to completion — `submit()` read-validate-append and `verify()` check-and-burn as
  one critical section. Reads are concurrent + lock-free over immutable committed records.
- **Crash-atomic commit — the named pivot.** The canonical typed-envelope record is staged → fsync → committed by one
  atomic `rename()` + dir-fsync (maildir linearization); **presence = committed**. INDEX.md / rendered `.md` / mailboxes
  are **derived projections** rebuilt from canonical records via a redo journal; canonical wins unconditionally. The
  m-1 `INDEX.md` on-disk layout is **unchanged** — only its crash-recovery *authority* becomes derived (CQ-8).
- **Durable FIFO + recovery.** Intake is append+fsync ahead of the in-memory queue; outcomes reference `intake_id`
  (atomic clear-on-pop, zero stale re-emission). Startup recovery = phases 0–4 (validate genesis → scan/quarantine →
  rebuild projections → restore runtime tables → open); **no authority consumption until recovery completes.**
- **Internal-fault disposition.** A trusted-side check that throws / times out / reads corrupt yields a typed outcome —
  authority record → **`HELD`** (a crash is not a yes; K8s `failurePolicy: Fail`), non-authority → bounce; the store
  never bricks (quarantine, not fail-stop). Terminal-state enum = byte-exact **`{accepted, rejected, held}`** (CQ-4).
- **Interface guardrail.** The conductor is an MCP server; each seat's trusted-side tool registry is exactly
  **`submit` / `project` / `read`**; `submit`'s input schema **IS** the rendered m-2 form (forbidden options absent);
  raw store/config/outbox/operator-channel paths + config *values* are absent from every seat surface. Delivery/wake =
  one `write()` onto the seat's pipe (per-runtime adapters). **Confusion-resistant, not adversarial isolation.**
- **Trusted config.** Per-domain sections (**m-2/m-3/m-4/m-5/m-6-authored** — the CQ-4b full author set, incl. m-5's archetype registry + m-2's declared section, x3-F2 c6) composed into one artifact under a **single top-level
  digest** + per-section stamps (CQ-4b), loaded once at trusted startup, digest-pinned to genesis, restart-only, absent
  from seat surfaces. Plus local-outbox-only external send (behind the m-3 egress gate), store genesis + GC, persisted
  seat-binding + decision-scoped `(decision_id, seat)` sibling-burn (CQ-6 base).
- **Conductor-internal provenance (s2-amend-m-1, 2026-07-04).** The engine's own **non-`submit`** records —
  genesis, gc_marker, recovery/incident, derived-outbox — are stamped `FROM = ROLE = "system"` (a reserved token,
  neither seat nor lane), `DeliveryState` within the byte-exact **`{accepted, rejected, held}`** (`held` for
  incident/quarantine); **`system` is never accepted from the public `submit` path** (extends I2 reject-unbound —
  confusion-resistant D4, D5 residual §C4.3). m-1 owns this *stamping* mechanism (authoritative: m-1 design-of-record
  §6); the record catalog/shapes are m-7/S2, `record_kind` is m-2. Review-driven fold of the S2 m-1-fidelity finding
  F-M1-1 (pair-approved; ratified into `RECONCILE.md`).

### C4.2 The engine × the six contracts (the seam matrix, §12 — 18 rows, biting negatives)
The engine **hosts + executes**, it does not re-own: m-1 store append + channel-stamped FROM (inside the loop); m-2
form render + fill-time authority + lineage gate; **phase-split required-set** (Step-1 gate never demands observe-owned
fields with no writer — CQ-1 (a) step-gate); m-3 observe hook + **decision-② class-conditional fail-closed** (CQ-2);
pure-judgment **A-floor** on the monotonic MAX (CQ-3); m-4 route_dispatch (R2-preserving — no model predicate in any
gate); m-5 slot_in classification at acceptance, post-gate/pre-observe/atomic-bind (CQ-5) + template/pane spawn; m-6
park/wake state machine + buckets/ODB. Every hosted row carries a biting negative fixture (§12) proving the guardrail.

### C4.3 Claim boundary (the line the NO-GO turned on)
Step-1 = attach + interface guardrail = **CONFUSION-RESISTANT**. The licensed "by construction" claims are exactly the
**four classes** the RECONCILE c5 kept-class records (not one): (1) the **serialized-loop kill** of the two-honest-seats
double-accept race (a control-flow property of the trusted engine, m-7 §2.4); (2) the **R2 no-model-predicate gate
grammar** — no model-derived predicate enters any gate (§C2.3); (3) the **observer-selected control properties** — the
F1 archetype invariants selected conductor-side, not by the observed lane (§C3.1); (4) **authority-ceilings-at-spawn**
(the §C3.1 ceiling vector). Everything else — attach + the interface guardrail — is confusion-resistant, not
by-construction. A same-uid, shell-bearing seat can write the store files directly — the documented **D5 accepted-risk**, out of scope;
adversarial containment / wrap / "sole-writer by construction" remain **shelved** (`GRILL-LOCK-deployment-fork-2026-07-01.md`).
The design's semantic claim-sweep (§16) is clean.

**Confusion-resistance is *tool-mediated* (external-review sharpening, 2026-07-03 — GPT-5.5-pro / Fable-5 read).** "Confusion-resistant" means a confused seat acting **through the `submit`/`project`/`read` tool surface** has no affordance for the wrong action — the guardrail removes **affordance, not access**. It does **NOT** contain confusion that routes *around* the tool surface: a coding seat has a shell, and a **confused** (not only malicious) agent executing a stale/hallucinated "write to the store at `<path>`" instruction via `bash` bypasses the guardrail exactly as a malicious one would (same-uid ⇒ the store is discoverable by `ls`). So **D5 covers shell-routed *confusion*, not just malice** — any seat reaching the store outside the tool surface, confused or adversarial, is the accepted residual. The honest claim is **tool-mediated confusion-resistance** + a *probabilistic* (affordance-following) defense against the rest — never a structural one. *(Same class as the c5/c6 sweep — an un-scrubbed overclaim, one layer down.)*

**Path-hygiene invariant (I-PH — Step-1-enforced).** No canonical **store / config / outbox** path may appear in **any seat-delivered surface**: bounce messages, error text, rendered projections, delivery payloads. Rationale: leaking a path *hands* a confused agent the very affordance the tool surface withholds — the conductor's own error message becomes the leak that converts confusion into a D5 bypass. Enforced by m-7 (guardrail + delivery), m-1 (store), m-2 (bounce/reason text); a **negative fixture** (Step-1 owed build item, §C4) proves no seat-facing output — every bounce/error included — contains a canonical store path.

**Cycle c4 CLOSED (VP co-sign `c4-design-m-7/…-040327`): the conductor-core substrate is LOCKED — re-baseline step (b)
COMPLETE.** Build-carries (non-locking): `re-mint-supersedes` (§2C away-bridge build step), CQ-7 row-parity, the
operator-gated runtime spikes (RUNTIME-RESEARCH §12). This lock grants no PLAN/IMPL/code/spike; **Step-1 PLAN** is the
operator-opened gate (re-baseline step (e)), after (c) global claim-sweep + (d) §2C-at-build-step.

**Integrated §2C build-carry ledger (recorded c5, mechanism deferred to step (d) — the one place a (d) builder inherits ALL §2C carries: away-token AND routing-lane):**
- **`re-mint-supersedes`** (m-7 CQ-6, m-1-confirmed-fit) — a resummon-mint burns the superseded prior-cycle `(decision_id, seat)` nonces.
- **Decision ④ away-token ROTATE + RE-OBSERVE** (RATIFIED c5, recorded m-1 `c5-fold-decision-4` + m-6 resummon edge, both impl-approved) — a refresh **rotates `decision_id` + burns prior nonces**; `verify` **re-observes current state and bounces the approval if it changed** since the operator last saw it (closes the stale-approval / TOCTOU window). **Dormant in Step-1** (no away-bridge/resummon exists yet); until (d), the base decision-scoped sibling-burn + m-6's never-auto-resolve-on-expiry FSM hold. Detailed design + fixture + adversarial review = a **step-(d) gate**, not design-locked here.
- **R2 `gate_referenceable` per-column FieldSpec** (m-4/m-2, recorded c5) — the **attribute is now declared first-class** (m-2 §4 `gate_referenceable: bool`, default false, `c6-fix-m-2`); the per-column **negative fixtures** over `chosen_model` and single-family bucket-valued proxies (R2's no-model-predicate grammar tested at the **live field grain**, not the ghost `selected_model`) remain a **step-(d) gate**.
- **Altitude-B per-row deviation grain** (m-4, recorded c5) — per-row `justified_deviation` / `deviation_reason_code` disposition so the silent-deviation veto is caught at **row grain, not record grain**. Value-set + fixtures = a **step-(d) gate**.
- **Away-mode trigger expressibility (m-5-F2, CTO ruling c6):** the m-5 posture model cannot yet express the away-mode trigger — **recorded a non-locking step-(d) build-carry**; the trigger's posture-model expression is designed when the **away-bridge** is built (it rides the deferred away-token/away-bridge mechanism above — not a Step-1 blocker).

**Owed Step-1-build fixtures (registration owed — recorded here so a (d)/Step-1 builder inherits them; flagged by the c6 pairs, receiving side CTO/m-7):**
- **③ known-A / RAISE-ONLY direction-invariant NF fixture** (m-7-side; flagged owed by m-2 `c6-fix-m-2` + m-6 §2, §J-ratified) — a B-pick / B-absorb over a known-A category ⇒ **raised to A + `gate_category` recorded**, never silently orchestrator-absorbed. Bound to Step-1 build; not yet in m-7's locked design-of-record (which locked before the ③ fold). *(SETTLED 2026-07-06, s5-escalations M-2, m-6+m-7 joint: **signal set** = S1 the CQ-3 A-floor over `(phase × record_kind)` [PRIMARY] + S2 the referenced-gate-record's own `gate_category` [mechanism, verdict-path store read] + S3 the §J2 merge-split content predicate, MAX-composed, `other`→A hardcoded; **mechanics** = the raise REWRITES the committed `gate_category` to the detector's named A member (else `other`) + stamps `gate_category_raised: "yes"`, atomic at the Q5 validate locus; lattice = the A/B MAP, never enum-index arithmetic; **Step-1 claim boundary** = exactly (S1)+(S2)+(S3)+fail-safe — NOT "every content mis-pick"; per-category content predicates beyond S1's grain are a Step-2+ carry.)*
- **⑤ ODB model-name egress fixture** (m-3-authored (a)/(b)/(c) set, `c6-fix-m-3` §13; receiving on m-7 / this §C4) — a model-name in the **exempt ODB `model_name` slot** (decision-⑤ carve-out) passes the egress scan; a model-name **outside** the carve-out is scanned/blocked. Bound to Step-1 build.
- **m-2 `GRILL_REQUIRED` FieldSpec row** (owed, m-6-F6) — m-2 declares the ported upstream `GRILL_REQUIRED` header's FieldSpec (owner/type/values) before the m-6 meeting-lane route binds it as a field; until then the route keys on the locked `phase` atom alone (m-6 §5). An owed m-2 build item.
- **§J2 explicit `routing_escalation` A-member** (owed cross-domain carry, m-4-F7 / x3-F1) — a **distinct** `routing_escalation` gate_category member for precise classification/telemetry of routing-escalation force-A. **Correctness already holds via the `other`→A fail-safe** (m-4 §7:365-367 — "no §J2 change required for correctness"), so this is **NOT a Step-1-PLAN blocker**; routed as a pre-wire cross-domain carry = CTO §J2 add + m-2 §J2-mirror + m-6 confirm. The member token is `routing_escalation`, distinct from the `routing_unavailable` route_dispatch outcome state (m-4 §7:363-364). *(c6.1a: an earlier c6.1 §J2 edit mis-named this member `routing_unavailable` — a differential-caught transcription slip, reverted; recorded here as the owed carry m-4 prescribed.)*
- **I-PH path-hygiene fixture** (external-review 2026-07-03) — no seat-delivered surface (bounce / error / projection / delivery payload) contains a canonical store/config/outbox path, so the conductor's own error text can't leak the affordance the tool surface withholds (§C4.3 I-PH). m-7-hosted; m-1/m-2 honor it in store + bounce/reason text.

**S6 TRANSPORT AMENDMENTS — DESIGN-OF-RECORD (VP co-signed 2026-07-06, `s6-design/RECONCILE-orchestrator-reviewer-20260706-220325`).** The integrated set **`master/S6-AMENDMENT-SET-2026-07-06.md` (r3)** amends this §C4's engine/store/form contracts per the F1–F17 transport findings (`TRANSPORT-FINDINGS-2026-07-06.md`) + the operator-requested B-1..B-3 boot stage: **branch-A conductor-computed PARENT** with fallback hints (`GRILL-LOCK-parenting-fork-2026-07-06`) · the **ONE canonical envelope codec** · stable-schema digest (A-1) · idempotent-replay intake (A-2) · **live mint as a loop mutation** (`seat_mint`, A-3 — supersedes admin-time-only) · the **I1-P store lock** (m-1 invariant / m-7 runtime) · `project()` **default-accepted** scope · **scoped waivers + `waiver_retraction`** · seat lifecycle **`minted→bound→active`** (derived-only activation, no persisted marker) + the typed boot form (`SITREP` + lifecycle-gating, no new vocabulary). The constituent domain amendment docs (m-1 + m-7 `2026-07-06-s6-transport-amendments.md`, m-2 `2026-07-06-s6-transport-codec-amendment.md`) are the authoritative deltas; where this §C4's older prose conflicts, **they win** — the full §C4 prose fold lands at the s6 build close. Registry impact: seven rows + two record classes, MINOR (no envelope migrator — the R-1 two-axes ruling).

**Step-3 routing-enforcement carries (registered 2026-07-06 — the m-4 (f)-confirm's two mandatory conditions, `s5-escalations`; receiving owner = the Step-3 router builder, with m-2 + m-4 co-signing at that gate; these must be explicit line-items, never "column validation later"):**
- **C1 — R2 at column grain.** The moment the FieldSpec grammar can address a `row_array` column, `chosen_model` and every model-identity column MUST carry `gate_referenceable:false` / be non-predicate — the by-construction R2 of the opaque Step-1 carrier is re-asserted explicitly when the grammar catches up. Load-bearing; not an afterthought of router work.
- **C2 — the `any_row` deviation-justification coupling.** `justified_deviation` + `deviation_reason_code`'s `required_when any(routing_assignments.declared_deviated == true)` (m-4 §5:210-211) — the readable half of the §2 silent-deviation veto — is inexpressible in the live grammar and rides this same carry; deferral is licensed ONLY because no Step-1 lane writes a routing record. It must be enforceable the instant the router writes one.

**INV-CATALOG — the named-invariant battery gate (registered 2026-07-07; seed = PROTOCOL-DEVIATIONS B9; receiving owner = CTO/m-7, with m-1/m-2/m-4 fidelity on their laws; lands as a follow-on at s6-close — explicitly NOT s6 scope):** consolidate the standing global laws — byte-exact `{accepted, rejected, held}` · the three-verb seat surface · R2 no-model-predicate · derived-only activation · I1-P sole-writer · I-PH · canonical-wins — from their scattered per-slice fixtures into **one `test/invariants` package**: each law NAMED, one executable check per law, the catalog file governed like `registry.json` (single-writer, owner fidelity on change). Effect: weakening a global law = a red battery **naming the law**, and the only path through is the amendment ritual — Cardinal rule 1 as a compile-time tripwire.

## CROSS-DOMAIN SCOPE AMENDMENT — the CONFUSION-FIREWALL DIRECTIVE (OPERATOR-RATIFIED 2026-07-11)

**Provenance + trail:** operator-originated (stated live in the m-2.planner session), scribed by m-2.planner (`master/relays/frank-threat-model-scope/SITREP-planner-20260711-162331.md` — surfaced, never enacted, at the correct authority grain), master-assessed with the four-effective-lock impact analysis (`…/SITREP-orchestrator-planner-20260711-162826.md`), **operator-ratified ("ratified", 2026-07-11) binding the exact text below.** Routings executed at adoption: the s8 build addendum (`s8-dispatch`) · the m-3 s9 adjudication design item (STEP-2-KICKOFF item 9) · the m-1 identity-seam rail (`frank-threat-model-scope-m1`).

> **(1)** The threat model is CONFUSION AND HALLUCINATION, not malicious adversarial agents: frank is a confusion/hallucination firewall for fallible agents, not a security product; design ceremony that re-imports the adversary is cut on sight. **(2)** The evidence mechanism is deterministic checks trusted-side + independent cheap-model adjudication for fuzzy claims + honest labeling; adjudication is medium-strength by design and labeled so. **(3)** The architecture builds extensible SEAMS (observe hook, evidence ladder, identity stamp) filled today with honest minimal mechanisms; no home-grown security primitive ships, ever; seams stay open by the per-surface rule — additive/open where ignoring loses only detail, closed/fail-closed where ignoring changes the meaning of acceptance (**Rail A**); mechanisms are cut by FUNCTION, not flavor — drift-detection, crash-safety, history-truth, and provenance machinery stays, claims worded to confusion grade (**Rail B**). **(4)** This REDUCES scope: no speculative frameworks, no doorways for doors nobody built. **(5)** Positioning leads with the confident-hallucination catch; robustness claims are never designed or pitched for.

**Standing interpretation:** Rails A/B govern every future open-vs-closed and keep-vs-cut decision; the c5/c6 claim-sweep discipline (confusion-grade wording beside every exclusivity-shaped claim) is the enforcement surface; this amendment CONFIRMS the standing claim boundary (§C4.3 tool-mediated confusion-resistance, D5, the D3 shelf) — it does not reopen it. At adoption, all four s8 design locks were verified surviving by function analysis: the composite digest = drift-detection + one-source-of-truth; adoption atomicity = crash-safety + history-truth; the A2 capability gate = the anti-"confidently-wrong-machinery" gate; the v1 executor ceiling and the shelved OS sandbox = this directive's logic applied before it existed.

---

# Step-3 architecture reframe — the conductor is ONE service in a larger app shell (RATIFIED 2026-07-15)

**Architecture-of-record:** `master/STEP-3-ARCH-AMENDMENT.md` (operator-ratified at SHA-256 `2d240eb6…`, VP-approved `step3-arch-packet/063000`). This section carries the reframe into the durable architecture record at full decision-bearing grain (the §1–§5 matrices are landed below). It **supersedes the conductor-hosted framing** of `STEP-3-KICKOFF.md` §§1–3 AND §§5–8 (old kickoff hash `983508fc…` preserved). Steps 1–2 are unchanged — no conductor byte changes.

## The topology
The app is a **modular monolith + supervised workers over local IPC — NOT networked microservices** (no API gateway; per-family stores/writers). **The conductor is one component, not the app hub.**
- **Conductor = the governed relay plane for stamped participants** (agent seats, orchestrator seats, the operator channel, reserved system-authored governance records) + its **own isolated store + sole governed writer**. It is NOT the app supervisor, run DB, provider client, turn engine, tool broker, terminal multiplexer, or general IPC bus.
- **App shell (greenfield, Division II):** **m-10 App Control Plane/Supervisor** (hosts+sequences; owns no policy) · **m-8** connector (holds provider credentials + does the provider wire, app-side; the last pre-wire enforcement host; `freeze→authorize→attach→send`) · **m-9** worker (the app-side turn runtime; the **only** app component that is a conductor seat) · human/terminal surfaces.

## Boundary invariants (the negative routes — NEVER transit the conductor)
provider request/response bytes · credentials/secrets · `LLMRequest` · tool-execution payloads · PTY/terminal streams · model-turn traffic · the run manifest + run/session/turn/attempt state · worker lifecycle. A violation is the category error the reframe corrects.

## Step-3 MVP — a barely-enough coding agent on the governed courier (OPERATIVE — the ratified MVP amendment)
**The normative spec is `master/STEP-3-MVP-AMENDMENT.md` r7, operator-ratified 2026-07-16 at SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`** (VP byte-bound approve `step3-arch-packet/…-035505`; ratification record `…-040405`; three operator grill decisions `…-023557`/`…-024350`/`…-025642`). It amends the ratified reframe packet at exactly four §1-named fragments (Sequence-A ceiling clause · first-stage order · the `:29` address-space phrase · the `:27` worker-as-principal phrase); the packet file stays byte-exact as the historical lock. This section is the architecture-of-record summary aligned to it — where wording differs, the amendment governs (H-3: one canonical representation).
The Step-3 MVP is a **minimal coding agent that collaborates with other agents through the governed conductor.** Scope is deliberately small; the permission/authority system, sandbox, and multi-model carousel are **Step-4**.

### Stage-6 RE-SCOPE ADDENDUM — operator-ratified 2026-07-21 (`master/STEP-3-STAGE6-AMENDMENT.md` rev12 `1125b0a0…`)
A **stage-5.1 third-party review** (`master/STAGE-5.1-EXTERNAL-REVIEW-2026-07-21.md` `b4e79f3b…`) found the frozen kernel to be real governance but not yet honestly a coding-agent MVP. The operator kept the "frank harness MVP" label and pulled scope up; the resulting **bounded re-scope amendment** was VP-APPROVED at decomposition grain (r12, twelve rounds) and **operator-ratified in-session 2026-07-21** (recorded agent-authored + operator-cited per §8b, ratify relay `step3-arch-packet/…-165500`). It is **additive** to the r7 MVP amendment above — it withdraws no approved mechanism and moves no bound byte of the nine frozen design finals + H-16 rev16 + the census. Where wording differs, the amendment's exact bytes govern (H-3). The addendum:
- **Sandbox FORGONE for the MVP** (D1): bash stays ambient (§4/`H-12`), the gap is documented, and **H-12 is promoted to a HARD pre-external-use blocker** — no untrusted / external / security-sensitive / multi-tenant use until a real sandbox lands (§10 deployment envelope + prohibited-use posture). The bash claim is **narrowed to invocation-context binding** (D2/D6: the effect descriptor is context-binding, `backend_id="ambient"`, no containment claim).
- **The exit test = six governance-property legs** (Governance · Durability · Crash-honesty · Injection-**visibility**-not-prevention · Handoff · Operability) **+ an objective overhead budget** (F59 authorize p95 ≤ 250 ms · relay ≤ 1 s · journal-commit ≤ 100 ms · per-turn added wall-clock p50 ≤ 20% PASS / 20–100% HOLD / > 100% FAIL), over a frozen `STEP-3-EXIT-FIXTURES.json` — **not** a benchmark score (D3). **Utility is DEMONSTRATED, not gated** (D4): public dogfood (CRM + bivpak, open auditable testaments) + honestly-labeled agent-as-operator SWE-bench, no threshold gate. **Decoupling invariant (D5): real-work / dogfood start ⊥ the exit gate.**
- **Durable session-state + resume is BUILT** (D7): a **worker-owned crash-safe session-content log** (field-standard; outcomes stay m-10-canonical — the frozen "no m-9-owned durable session store" invariant narrows to "no m-9-owned durable *outcome* store"). The trust model is **two time-scoped properties**: at settlement, `settled_with_content` is durable *evidence* the content + admitting marker/receipt fsync-linearized then; at resume, content is trusted **only** under matching positive evidence **AND** presence in the current recovered valid prefix — else `content_lost` → `DEGRADED`, never fabricated. The pre-admission manifest is a **producer-total m-10 three-class evidence union** (`settled_with_content` / `determinate_no_resume` / `uncertain`; completed-without-receipt ⇒ `uncertain`); `content_lost` is the **m-9 post-inspection reconciliation result**. An un-emittable oversized resume frame commits to **ONE terminal `FAILED`/`resume_frame_overflow`** state (no successor/lease/snapshot/revival, operator manual `resume_action`). Subsystem internals (record/segment/rotation grammar, the writer fence, exact frame/table) are delegated to the m-9/m-10 pairs under **F73**. **§D-settlement refinements (consolidated 2026-07-27):** the pre-admission disclosure set is restored **run-wide** (not per-turn), bounded by `MAX_PARKED_ROWS_PER_RUN=512` with a **run-terminal `parked_unknown_capacity_exceeded`** (no resume branch, no first-action entry); a zero-attempt **`turn_failed`** explicitly supersedes the prior framing; and the **D-4 Gate-2** claim is the consumer-side fail-closed validator + drift-detector over MVP-unreachable states (Gate 1 delivers the guarantee; the r21 §2.6 comparator bytes stay frozen).
- **Structural (items A/B/E):** the interface lock is re-cut with a `frozen_core_digest` join + `model_surface_digest` + typed E3 predicates; the pair order is an **acyclic DAG** (§6) with the **m-7 broker study resolving first**. *(Item A's original hashable Tier-HARD **bundle** mechanism — `master/STEP-3-INTERFACE-BUNDLE.json`, `bundle_sha256` over a canonical `lock_payload`, the extractor tool + `--verify`, HARD markers, the `bundle-soft-stability` fixture — is SUPERSEDED/WITHDRAWN by the item-A simplification amendment; see below.)*
- **The held joint interface-lock record `b7e1f0ef…` is SUPERSEDED** — replaced by a later, **shorter re-lock** over the interface-lock record's external SHA-256 + the whole-file-hard owner contracts. **The §11 sequence status (2026-07-27):** Lane 1 (m-7 broker study) **CLOSED** (2026-07-21, no H-24); Lane 2 (interface DAG legs) is **CLOSED** (9 settled owner bases + 5 joins) — owner producers pair-approved (m-1 `d34a7c47…`, m-2 `c3a8cd61…`, m-8 r5 `c0b7b488…` + 2a/2b discriminator r7 `734e44b7…`, m-9 §5-E recipes r12 `04422965…`, m-10 B/E rev3 `cd17db32…`) and m-3's sink/join leg closed **HONEST-PARTIAL** at r19 `92e08d09…` (T1–T8 live; N910/T9 a documented MVP limit; r7-mirror deferred v3). The **§D-settlement amendment** (`master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` rev4 `1fa71cb8…` + bound m-2 cell `5ec7a3d2…`) is **OPERATOR-RATIFIED (2026-07-25, recorded §8b at `step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-170000`)** — its four bounded mechanism corrections are now in force and the propagation matrix is open (m-9 + m-10 fold fresh successors; m-2 unchanged). **CONSOLIDATED into this architecture-of-record 2026-07-27** (lane-2 closed + §D join closed): the D7 resume additions [run-wide restore + `MAX_PARKED_ROWS_PER_RUN=512` + `parked_unknown_capacity_exceeded` terminal + `turn_failed` zero-attempt supersession] are folded into the D7 durability bullet above; the **`relay.submit` `canonical_resource` cell** (m-2, ratified `5ec7a3d2…`; §D-settlement Correction 3 / §5-C) is the settled form-determinism binding — `canonical_resource(relay.submit) = "relay.submit:" || SHA-256(JCS{ form_digest, dispatch_id?, to?, cc? | cc_unparsed? })`, where **`form_digest` is REQUIRED** (the cell is total — no MVP branch yields an unfillable value), the destination coordinates **`dispatch_id`/`to`/`cc` are omitted when absent**, and the CC member is a decoded string-array **or `cc_unparsed`** (mutually exclusive, distinct member names). Binding CC as an effect-delivery target **confers no relay authority** — the standing TO/CC authority protocol is untouched. **Still OPEN:** the fixture freeze (`STEP-3-EXIT-FIXTURES.json`) → the shorter stage-6 re-lock over the interface-lock record's external SHA-256 → T4 (behind the re-lock + H-16/H-26). *(The propagation matrix and the §D two-sided join CLOSED, closing Lane 2. Item A is now RATIFIED + AUTHORED — `master/STEP-3-INTERFACE-LOCK.md`, external SHA-256 `cbd1893c…` — and in VP + F73 review, per the item-A simplification amendment `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev7 SHA-256 `3443f73d…` (VP-approved, operator-ratified 2026-07-27), which WITHDRAWS the prior bundle mechanism and the extraction recipe `master/STEP-3-ITEM-A-RECIPE.md` r3 `06e6956e…` (not built).)* Ratification issued **no** DESIGN-lock / PLAN / T4 token / credential / provider call / release binding / live E3 / merge / deploy.

**Pinned MVP process topology (grill #1, amendment §2b):** the **conductor** = its own isolated service · the **app main process hosts the m-10 control plane as a MODULE** (not a separate daemon; its manifest/one-shot-authorization/epoch/IPC seams are designed as-if process-separated for a Step-4 split) · **m-9** = the supervised worker process (acts FOR the logical m-9 seat through an epoch-bound capability — F66) · **m-8** = a separate supervised connector process (the single egress chokepoint; F57-narrowed isolation). The seat broker lives OUTSIDE the replaceable worker generation; its placement = the m-7 channel/broker DESIGN under m-1's identity semantics.

**What the MVP builds:**
- **A coding-agent worker (m-9)** running a model turn (Codex-first) with **local tools: `read` · `write` · `edit` (str_replace) · `bash` · `apply_patch`** — app-side, never through the conductor. (str_replace is the model-agnostic base, apply_patch the Codex/GPT add-on — empirical: pi ships str_replace only, jcode/opencode ship both.)
- **The conductor as a NATIVE tool (not via MCP):** built-in `submit` · `project` · `read` (relay verbs) that speak to the conductor **service directly over its socket** (`internal/channel`), presented to the model as first-class tools for **agent-to-agent relay communication**. The conductor stays a **separate isolated relay plane** (reframe intact) — the native tool is a *client*, not a merge; governance (channel-stamped FROM, store, observe) is conductor-side, unchanged. (Distinguish the local-file `read` from the relay `read` in naming.)
- **The MCP server (`cmd/frank-mcp`) is RETAINED** as the adapter for **foreign harnesses** (agents on Codex / Claude Code / third-party joining the governed conductor — how the team's external seats connect). Not dropped; it is the interop face, not our internal path.
- **A shared conductor-client (a `frank/` refactor):** today the reusable client logic (`SubmitPayloadFromArguments`, `SchemaFromForm`, the re-render nudge, reconnect) is **trapped in the `cmd/frank-mcp` command** (package `main`; nothing else imports it — verified). The native-tool work **hoists it into a shared client** (`internal/channel` or a new `internal/conductorclient`) so BOTH the MCP frontend and the native tool are thin frontends over one client. Target layering: **engine (daemon) → `internal/channel` socket + shared conductor-client → { MCP frontend · native-tool frontend }.** (m-9↔m-7 seam — the client surface is m-7's interface-guardrail domain.)

**Authority in the MVP — the seam is built, shipped EMPTY:**
- **NO config-derived permission ceiling in the MVP.** The m-5 per-role, config-derived, fresh ceiling + `config_generation` freshness are **Step-4**. This **dissolves the seam-13 / freshness / m-5-amendment knot** for the MVP entirely — there is no ceiling to be fresh or stale about.
- The MVP's authorization is the **fixed tool-DISPATCH seam** (amendment §4): the **operator-ratified 8 canonical NAMES** — local `read`/`write`/`edit`/`apply_patch`/`bash` + conductor `relay.submit`/`relay.project`/`relay.read` — as the POLICY identity; **build identity** = per-tool vectors `{name, schema digest, catalog version, mapping version}` with named producers (m-9 local / m-2 relay verbs), bound at the first-stage interface-lock + a **post-build RELEASE-BINDING event** (F63). m-10 hosts the check (exact canonical-set equality over identity; fail-closed deny), owns no policy. **Authorized == executed via the F59 one-shot ticket** (grill #2): m-10 writes a durable ticket bound to `{run_id, turn_id, turn_epoch, tool_call_id, canonical_tool_name, canonical_args_digest}`; m-9 consumes exactly once (atomic) and executes exactly the digested call; duplicates/stale-epoch/mutations/crash-window replays rejected — park, never silent replay. It exercises the **enforcement seam** m-5's real ceiling plugs into at Step-4; no `config_generation`, no freshness claim. m-9's inert-until-authorized invariant is honored through the ticket.
- **`bash` = ambient host/external/destructive authority — operator-accepted residual** (no cwd-confinement/sandbox/irreversibility gate; audit is evidence, not prevention; a trusted executor is Step-4 hardening **H-12**).

**Wake-on-relay (stretch goal — push-based; the F61-honest contract, amendment §6):** `internal/channel` already carries a server→seat push (`PushTo`/`NextPush`, used by the park/wake prompts). Wake is event-driven — the conductor pushes on an inbound relay (no conductor protocol change), m-9 forwards the `relay_id`, m-10 schedules. **Guarantees stated honestly:** push = **best-effort + advisory** (no in-connection ack/retry claimed; a wake lost while connected is explicitly accepted); recovery = **durable rediscovery** (catch-up `project`/`read` on startup/reconnect — the record is truth); scheduling = **at-most-once** via durable `UNIQUE(relay_id)`. Non-gating; polling is the fallback.

**Deferred to Step-4 (was contested MVP scope):** the m-5 config-derived ceiling + `config_generation` freshness + per-role permissions; a hard sandbox + irreversibility gating (H-12); the per-model tool manifest / extensible registry / model carousel; live provider routing (m-4). **The reframe packet FILE stays byte-exact** — the ratified MVP amendment supersedes exactly its four §1-named fragments; no conductor byte/member change, no new conductor output, no direct edge (the native tool uses the existing `internal/channel` socket).

## Evidence honesty across the app boundary
An app-side provider send leaves the conductor's vantage, so a connector's send/deny/stream report is a **self-reported, worker-carried provider report at E0 — not a conductor observation** ("attestation" is avoided, amendment §3: the conductor establishes `FROM=m-9-seat` + body-claims-m-8-origin, not that m-8 authored the body), and it never promotes uncorroborated. Its **carrier** is the m-9 worker's existing `PHASE: SITREP` / `AUTHORITY: report-only` / `HUMAN_GATE_REQUIRED: no` relay with the m-3 app-event schema **in the body** (no new relay kind / FieldSpec row — no conductor change); top-level relay evidence describes carriage only, the body event carries its own `event_evidence=E0`/`event_integrity=self_reported` so a top-level observation cannot upgrade it. **Live E3 comes only from a separate integration harness / operator observation**, bound (F63/F65) to the exact **app/provider-vertical release** (app-main/m-10 + m-9 worker + m-8 connector digests or a covering `release_digest`) + manifest/catalog/policy digests — the **conductor service identity is a separate lifecycle, bound separately in the exit-test record** for the relay-exchange leg (whose evidence is the conductor's own observe-as-send E1/E2 records). The MVP ladder is **E0–E3**. **"Honest governed turn" = the app enforcing locked owner policies with correctly-labeled proof, NOT every app event becoming conductor-observed.**

## Ownership deltas (as amended by the ratified MVP amendment §7)
provider-send **mechanism** app-side, **policy stays m-3** · secrets m-1-governed, held at runtime in m-8 · **the m-5 ceiling STOOD DOWN for the MVP** (the enforcement seam = the operator-fixed 8-name dispatch constant, m-10-hosted; the untouched ceiling contract `643dd7c2…` is the Step-4 basis that plugs into the same seam) · m-7 retains conductor-host, loses the provider-credential contract (re-owned connector-side, m-8-authored) **+ authors the shared transport/client boundary AND the authenticated channel/BROKER contract incl. the F64 per-verb generation fence** · **m-1 authors the connector secret-boundary delta + the seat identity/credential-lifecycle semantics (F60: one broker-held credential per LOGICAL seat)** · **m-2 authors the form→tool-schema mapping + the relay-verb schema digests/mapping version (F58)** · **m-3 authors the egress policy + the E0 app-event schema + the E3 applicability evaluator (F62)** · **m-9 owns the catalog build + local-tool schema digests (F58) + the F59 executor half; m-10 owns the F59 one-shot authorization protocol + the durable app state + epoch fencing** · m-4/m-2 routing-record defers to Step-4 (Step-3 uses the app-side pinned run manifest — no routing execution).

## The operator direct route (a named non-governed route, ratified)
Authority-bearing **by construction** (confusion-not-malice: the live interactive channel authenticates the operator; forcing a governed relay to prove it is adversarial-shaped ceremony, cut). **Non-transitive:** a direct instruction authorizes only the addressed recipient within its ceiling; the citation is E0 evidence, not a transferable grant; a conductor-governed action needing a typed grant goes through the **landed grantor grammar** (a sanctioned grantor emits the typed grant under its own authority, citing the operator). No forged `FROM: operator`, no silent store mutation, no ceiling raise except via m-5's typed contract; live-runtime credentialed legs stay operator-performed. This is a Rail-A/Rail-B-consistent scope *reduction* (the confusion-firewall applied to the human channel).

## The matrices (landed from the ratified packet §1–§5, **as amended by the ratified MVP amendment r7 `2f75f2a1…`** — the durable architecture-of-record, not a summary)

### Boundary matrix (§1 — full decision-bearing grain, 10 columns)
| Component | Owner | Conductor SEAT? | Process boundary | API / IPC | Canonical state | Writer | Secrets | Authority / gates | Evidence |
|---|---|---|---|---|---|---|---|---|---|
| **Conductor — governed relay plane** | m-1..m-7 (m-7 hosts) | — (it IS the plane) | own isolated service | `submit`/`project`/`read`, stamped participants only | conductor relay store | m-7 sole governed writer | none | the locked relay gates (m-2 form/lint+lineage · m-3 observe-as-send · m-1 stamp · m-6 park/wake+ODB) | E0–E4 over **relay** traffic |
| **m-9 Model Runtime — the LOGICAL seat + its supervised worker generations** *(F66-amended)* | m-9 | **YES — the LOGICAL m-9 identity is the sole app seat**; the credential is **broker-held (m-7 broker, outside the replaceable worker generation)**; a worker generation gets an **epoch-bound revocable USE capability**, never the bytes | app-side worker process, supervised by m-10 | app IPC to m-10 + m-8; the seat's private authenticated channel to the conductor **via the broker** (every relay verb + push epoch-checked, F64) for genuine seat↔seat relays | m-9 turn/session/context state | m-9 | **no provider key; no other seat's secrets** | parses tool calls → INERT until the **m-10-hosted fixed tool-DISPATCH check issues the F59 one-shot ticket** (m-5's ceiling = Step-4); m-9 consumes-then-executes exactly the digested call | the worker seat is the **only submitter** of the E0-labeled worker-carried provider report |
| **m-10 App Control Plane / Supervisor** (NEW) | m-10 | **NO — trusted app component, no `submit` credential** | **a MODULE in the app main process** (§2b; seams designed as-if process-separated) | app IPC to workers + connector | app **run manifest** + supervisor + scheduler state + active-turn lease + **the durable app-state store incl. `tool_authorizations` (F59 tickets) + `leases`/`epochs`** | m-10 | orchestrates only **opaque credential references**, never secret bytes (incl. never the seat credential — F60) | hosts the **fixed tool-dispatch seam** (the operator-ratified 8-name constant; the m-5 ceiling plugs in at Step-4); authors the F59 ticket protocol; owns **no** policy; supervises + `turn_epoch`-fences without authoring as the worker | emits app events; not a conductor principal |
| **m-8 Provider Adapters / connector** | m-8 | **NO — trusted app component, not a seat** | **separate trusted process from the m-9 worker BEFORE the first E3** *(F57-amended: separation = accidental-disclosure reduction + non-injection into the enumerated surfaces; same-user inspection of peer-process state is an explicit unsandboxed MVP residual — no hard unreadability claim; the OS boundary is Step-4 H-12)* | receives `LLMRequest` over app IPC; provider HTTPS out via a frank-owned client | connector provider-attempt / stream telemetry | m-8 | **credential runtime holder + secret-store reader/writer under an m-1-authored boundary** | the **last pre-wire enforcement host**; the **`freeze → authorize → attach → send`** locus (one attempt per provider INVOCATION, no auto-retry — §2a) | the **provider-report source** (self-reported, worker-carried E0 — "attestation" avoided); returns app events over IPC, does not stamp a conductor record |
| **Human / terminal surface** | m-6 (governance semantics) + m-10 (PTY/TUI, Step-4+) | m-6 gates via conductor; terminal app-side | conductor-side gate→bucket/ODB; PTY app-side | m-6 surfaces via the relay plane; terminal via app IPC | m-6 HUMAN_GATE fields (relay store); PTY state app-side | m-7 / m-10 | none | m-6 HUMAN_GATE; the **operator direct route** (authority-bearing by construction, non-transitive) | conductor-side gate evidence |

### Traffic matrix (§2)
- **Governed relay traffic (transits the conductor, unchanged):** seat↔seat relays · the operator channel · reserved system governance records · the E0 governance/evidence summary · gate/park/wake/ODB. **No routing DECISION record, no lane-bearing FieldSpec row added in Step-3.**
- **Negative routes (NEVER the conductor):** provider request/response bytes · credentials/secrets · `LLMRequest` · tool-execution payloads · PTY/terminal streams · model-turn traffic · the run manifest + run/session/turn/attempt state · worker lifecycle. A violation is the category error the reframe corrects.
- **The operator direct route** — a named non-governed route (authority-bearing by construction, non-transitive, see the reframe section above).

### State-and-recovery matrix (§3)
| State family | Owner | Store | Writer | Crash/disagreement | Conductor authority |
|---|---|---|---|---|---|
| Conductor relay records + projections | m-1/m-7 | conductor store | m-7 sole governed writer | existing crash-atomic commit + recovery | **authoritative** |
| App run manifest + supervisor + active-turn lease + tool-authorization tickets | m-10 | m-10 store (durable; separate from the conductor store/writer) | m-10 | **fail-closed to interrupted/held; NEVER auto provider resend; a replacement starts only after the prior worker/attempt is proven terminal or on explicit operator disposition — and is fenced by a monotonic `turn_epoch` (stale generations actively rejected at m-8, the executor, and every broker relay verb/push; explicit UNKNOWN/PARTIAL states park, never silently replay)** | evidence, not authoritative |
| m-9 turn/session/context | m-9 | m-9 store | m-9 | m-9 one-active-turn invariant | evidence about a turn |
| m-8 provider-attempt/telemetry | m-8 | m-8 store | m-8 | attempt bound to `attempt_id`; no auto-resend (a user retry = a NEW `attempt_id`) | **provider-report source — self-reported E0 unless corroborated; live E3 = a separate observer** |
| Credential references + secret material | m-1 boundary / m-8 runtime holder | out-of-band; **never** conductor store | m-8 under m-1 boundary | rotation without genesis break | never in the conductor store |
| Terminal/PTY (Step-4+) | m-10 | app-side | m-10 | app-side | evidence surface only |

Stable IDs across the app stores: `run_id`/`turn_id`/`request_id`/`attempt_id` (+ the fencing `turn_epoch`); **no cross-store atomicity**. **Report-authority rule (amendment §3 — "attestation" avoided):** a conductor relay recording an app-side send is evidence — a self-reported worker-carried report at E0 — never authoritative app state, never the payload — carried in the m-9 worker's `SITREP` body (top-level relay evidence = carriage only; body event = `event_evidence=E0`).

### End-to-end sequences (§4)
- **Sequence A (app-side turn, MVP — the §1-superseded ceiling clause replaced per the ratified amendment):** m-10 writes the pinned run manifest (immutable lane ID + the locked tool-identity vector + digest — not a routing decision, not a gate input) → m-9 assembles the turn, tool call INERT until **the m-10-hosted fixed-set dispatch check passes (deny = not exactly the ratified 8 canonical IDs, fail-closed) and the F59 one-shot ticket is written; m-9 consumes the ticket exactly once and executes exactly the digested call** (the m-5 ceiling plugs into this same seam at Step-4) → m-9→m-8 `LLMRequest` (no keys in m-9) → m-8 `freeze→authorize→attach→send` (denial = zero send, no post-auth mutation of the frozen core, no secret exposed to m-3; **one attempt per provider INVOCATION, no auto-retry anywhere in the stack — a turn may hold multiple recorded attempts, §2a**) → m-8 returns stream + the provider report → the m-9 worker seat submits the single E0 worker-carried summary → **recovery/cancellation is app-side (§3, epoch-fenced, UNKNOWN/PARTIAL park-not-replay); a governed human gate is opened/read by the worker seat through the existing conductor verbs (no new m-10 conductor address or wake API)**. Zero conductor transit for provider bytes.
- **Sequence B (governed relay, UNCHANGED):** `submit()` → m-7 stamps FROM → m-2 form/lint + m-3 observe-as-send + lineage + single-writer commit → recipient `project()`/`read()`.
- **Honest governed turn** = the app enforcing locked owner policies with correctly-labeled proof (instrumented-test negatives, amendment §3/§10: policy-deny→zero provider-transport invocation · no post-authorize frozen-core mutation · secret resolver never invoked on the denied path · unknown/above-set tool→zero executor invocation · authorized==executed (the actual invocation equals the ticket) · stale-epoch rejected · no AUTOMATIC typed forwarding route into the conductor — agent-authored quoting stays legal; + one live E3 via a separate observer, bound to the exact app/provider-vertical release + manifest/catalog/policy digests, never laundered into the conductor summary — the conductor's own identity bound separately for the relay leg, F65).

### Scheduler split (§5)
Conductor governance-gate scheduler (m-6: park/wake exactly-once + ODB) — unchanged. App scheduler (m-10: worker scheduling, provider-await, cancellation, backpressure) — new, app-side. The bridge **reuses the worker seat's existing verbs** for any governed gate — NOT a new conductor event or m-10 principal.
