# Workflows & Archetypes — DESIGN (the archetype system: the governed expansion-slot)

DESIGN_DOC_ID: c3-design-m-5-workflows-archetypes
RUN_ID: master
CYCLE: c3
OWNER: m-5 (Workflows & Archetypes) — design-lead m-5.planner
STATUS: **design-complete** — m-5.implementer DESIGN-REVIEW **approve** (`c3-design-m-5/DESIGN-REVIEW-implementer-20260630-133831.md`); m-5↔m-6 COORD seam **RECONCILED both sides** (m-6 `123022` ⊕ m-5 `131856`; both domain docs matched on the four-class non-gate seam after m-6 withdrew the crossed `131747` path). **LOCKED per c3-lock (VP co-sign 20260630-191315).** Builds on the LOCKED c1+c2 contract (`master/ARCHITECTURE.md` §1–§C2); does NOT reopen m-1..m-4.
GRILL_LOCK_ID: c3-grill-m-5 (folded in §12; operator decisions of 20260630)
DESIGN_REVIEW_VERDICT: approve — m-5.implementer `c3-design-m-5/DESIGN-REVIEW-implementer-20260630-133831.md` (rev-1 `131617` + rev-2 `132748` must-revise folded; cross-domain seam cleared — both docs on the four-class non-gate model)

---

## 1. Frame, scope, and what is already locked

**The one job.** Bind the two c2-reserved opaque atoms (`slot_in`, `seat_archetype`) into a concrete, m-3/m-4-consumable, m-6-bindable **archetype system** — and prove the net-new contribution: an **archetype is a single governed expansion-slot preset that binds {topology + gate-set + authority-ceiling-at-spawn + observe-invariants + routing-prior} as ONE declared, seat-stamped, append-only, auditable unit.** No surveyed system (internal — the stock protocol, codex, claude-code — or the external 2023-2026 framework landscape) binds these five into one auditable unit (§2, §11).

**Locked c1+c2 contract this design parameterizes (NOT reopened):**
- **Two opaque atoms reserved to m-5** (ARCHITECTURE.md:55, :180-194): `slot_in` = work-archetype (m-3, per-work-record); `seat_archetype` = seat-archetype (m-4, per-seat-at-spawn). c3 binds their concrete vocabulary; **no m-2 micro-fold** (no `required_when`/`visible_when` over concrete values — C2.4).
- **F1 provenance** (ARCHITECTURE.md:184-186): both atoms conductor-owned / **non-lane-writable** (confusion-resistant — no lane tool writes/re-tags them; D5 residual, §4). `seat_archetype` spawn-fixed; `slot_in` **conductor-classified at work-record acceptance** (not spawn-fixed — long-lived seats move bugfix→refactor→migration).
- **F2 record home** (ARCHITECTURE.md:187-188): `seat_archetype` + resolved `authority_ceiling` recorded per-assignment on `routing_assignments` (replay-complete; template + hand-authored paths).
- **m-3** owns the observe-gate + classifies `slot_in` → selects the done-predicate/invariant add-on (C2.1). **m-4** owns routing + the `seat_archetype` capability-prior key + the GL-4 routing-template **record** mechanism (pre-filled `routing_decision` + `template_ref`, no-bypass) + the authority-ceiling fail-closed staffing refusal (C2.2). **conductor-core** owns pane-spawn via existing tmux/zellij/OS-terminal. **M4-1**: routing B→A escalation rides the c1 monotonic HUMAN_GATE routing-raise.

**Non-goals this cycle:** no PLAN/IMPL/merge; no reopening m-1..m-4; the conductor/N-pair template stays Step-5; the literal single-bounded-action actuator seat + standalone-runtime ceiling enforcement are Step-4/5 carries (§9, §13).

**Evidence base.** The c3 audit + F4 pair-reconcile (`c3-audit-m-5/AUDIT-planner-20260630-053308.md`, `…/AUDIT-implementer-20260630-053116.md`, `…/RECONCILE-planner-20260630-120326.md`); the prior-art sweep (3 source agents + a 95-subagent deep-research workflow); the m-5↔m-6 COORD seam, reconciled both sides (m-6 bind `c3-design-m5-m6-coord/COORD-planner-20260630-123022.md` ⊕ m-5 confirm `…/COORD-planner-20260630-131856.md`).

---

## 2. The unifying binding (the net-new contribution — PROVED, not asserted)

**Claim.** An **archetype** is a governed preset that binds five dimensions into one declared, seat-stamped, append-only, auditable unit:

```text
archetype  ⇒  { topology, gate-set, authority-ceiling-at-spawn, observe-invariants, routing-prior }
```

**Proof it is net-new (promote the parts, build the binding):** each dimension exists in prior art, but as a *separate, non-uniform, non-auditable* mechanism. PROMOTE, do not rebuild:
- topology preset → the upstream review-panel preset+selection+justification PATTERN (`review-panels.md:5-41`); codex `agent-graph-store` query-surface (`store.rs:17-55`) **as an append-only projection** over m-1 lineage (NOT codex's mutate-in-place);
- gate-set / authority-ceiling → the upstream `AUTHORITY`/`CEREMONY_TIER` enums (`protocol.md:16,248-251`); claude-code agent-type `tools[]`/`disallowedTools[]` **enforced** ceiling (`agentToolUtils.ts:122-225`);
- preset-as-data → codex `collaboration-mode-templates` mask (`config_types.rs:721-727`) + claude-code frontmatter agents + CrewAI JSONC/YAML (data, but per-agent runtime caps only, no governed bundle);
- observe-invariants-by-work-kind → m-3 §5 candidate predicates; external RefAgent/VeriMAP (work-type→check-type, but ungoverned);
- routing-prior → claude-code per-agent-type `model`; m-4 `seat_archetype` capability-prior key (locked).

**The integration is the contribution.** The external 2026 topology survey (doi 10.3390/fi18060326) independently separates **template ⊥ realized-graph ⊥ append-only-trace** — validating our preset/instance/lineage split — but audit-first immutable lineage appears externally *only* in a security-provenance line (Merkle/CT identity chains), never bound to a topology/archetype. **Our archetype unit, riding m-1's append-only lineage with a seat-stamped FROM + a per-assignment recorded ceiling (F2), is the cross-over no surveyed system has.** (E1: the prior-art sweep, `c3-audit-m-5/AUDIT-planner-20260630-053308.md` §2/§3.1.)

---

## 3. The two-axis tag-space (LOCKED vocabulary — `lower_snake_case`)

Two orthogonal opaque atoms; the values are m-5-owned (C2.4 de-locking note); locked here at the c3 DESIGN grill.

### 3.1 `slot_in` — work-archetype (per-work-record; conductor-classified-at-acceptance; selects the m-3 observe-invariant)
`{ feature_extension, refactor, cleanup, bugfix, migration, research_synthesis, qa_review, docs_chore }`
Mutating/code work: feature_extension, refactor, cleanup, bugfix, migration. Read-only investigation: research_synthesis, qa_review. Low-mutation hygiene: docs_chore.

### 3.2 `seat_archetype` — seat-archetype (per-seat-at-spawn; m-4-keyed; selects ceiling + routing-prior + tool-set + behavioral-mode + default human-mode)
`{ solo_worker, planner, implementer, reviewer, sensor, orchestrator_lead }`
**`actuator` is NOT a literal value** — it is a **derived mutating-ceiling class** (§7.2, GL-1).

---

## 4. The composition rule + the confusion-resistance proof (observer-selected control property)

**Composition.** A running seat's effective gate-set = **`seat_archetype` (spawn-fixed) ⊗ `slot_in` (per-work-record, classified-at-acceptance).** The `seat_archetype` fixes {ceiling, tool-set, routing-prior, behavioral-mode, default human-mode}; each work-record's `slot_in` adds its observe-invariant family on top of the phase-shaped base predicate, evaluated **within** the seat's ceiling. The conductor records both; the lane rewrites neither.

**Work-invariant families (the m-3 observe add-on per `slot_in`):**
| slot_in | observe invariant (m-3 selects) |
|---|---|
| feature_extension | additive-only diff (where feasible) + suite-green + **no test edits** unless test-addition declared |
| refactor | suite-green **AND test-files-unchanged** |
| cleanup | find-references = 0 (or intentional-redirect evidence) + no new public behavior |
| bugfix | repro **red-on-parent → green-on-fix** (where executable) + suite-green |
| migration | reversibility/rollback note + compatibility evidence + broader smoke; human-gate tends A for irreversible/data |
| research_synthesis | cited artifact present + **no source actions** + uncertainty preserved |
| qa_review | finding/no-finding evidence + dedup/novelty + **no source actions** unless escalated |
| docs_chore | artifact present + source-actions confined to doc/config paths + minimal gate |

**PROOF (confusion-resistance via observer-selected control) — a confused lane has no tool to re-tag its own `slot_in`.** The two load-bearing invariants (refactor `test-files-unchanged`; bugfix `red→green`) defend against the reward-hacking failures m-3 cites (edit-the-tests-to-stay-green; hardcode-the-artifact). They hold **iff the lane cannot choose its own `slot_in`.** By locked **F1**, `slot_in` is **conductor-classified at work-record acceptance** (confusion-resistant — no lane tool/verb writes or re-tags it). So a *confused* lane has no verb to re-tag `refactor`→`feature_extension` to dodge `test-files-unchanged`: the conductor, observing from outside the lane (DI-5), classifies the work-archetype from the observable change, and the invariant for *that* classification is the gate on the door (observe-as-send). **The invariant is selected by the observer, not the observed** — this observer-selection is a genuine control property (the selection lives conductor-side, unreachable by the observed lane); it rides the locked F1 provenance + the locked observe-as-send gate, requiring no new mechanism. **D5 residual:** classification alone does not contain a *genuinely malicious* seat — one that could write a false store record or reach a shell outside its tool surface (e.g. under attach) is out of scope for the initial release; that residual is bounded by m-1 store isolation + the tool-ceiling enforcement carried to Step-4/5 (§9, §13), not by this proof. (E1: ARCHITECTURE.md:142-143,184-186, §C3:277; m-3 design §2/§5.)

---

## 5. The authority-ceiling vector (LOCKED — 3 axes, open named-axis map, GL-2)

**Representation = an OPEN named-axis map**, not a fixed tuple — so axes are additive without a re-cut (operator modularity requirement, GL-2):
```text
authority_ceiling: { write: <lvl>, dispatch: <lvl>, tool: <lvl> }   # initial release
```
Each axis is its own monotonic lattice; consumers (m-4 `route_dispatch`, the F2 per-assignment record) **iterate present axes**; an axis **absent** on an older record falls to that axis's **declared absent-default** (read-time-migrator-friendly; consistent with the locked append-only store + zero-migrator versioning). **Every axis MUST declare an absent-default, and it is that axis's most-restrictive floor (fail-closed): `write → read_only`, `dispatch → none`, `tool → none`** (and the deferred `external_send → none`, below). A record-absence therefore reads as the **minimum grant on that axis, never an escalation** — monotonic-tightening-consistent, so an old or partial record can only ever be *under*-authorized, never silently widened.

**The three initial-release axes (partial order — NOT a single ladder):**
1. **write:** `read_only < write_scratch < write_feature < merge_feature < merge_protected`(A-gated). Merge folds in as the top tiers.
2. **dispatch:** `none < route` (orthogonal — an `orchestrator_lead` routes but cannot write).
3. **tool:** `none < read < exec` (orthogonal — `sensor` = none; `reviewer` = read; `implementer` = exec).

**Semantics:** a `seat_archetype` pins a **MAX** point per axis at spawn (recorded per-assignment, F2). A template/expansion-slot may **TIGHTEN below** the archetype max on any axis (monotonic, like c1 J1 gate-override monotonicity); it may **never loosen above**. `route_dispatch()` **fail-closes** (`routing_unavailable`) if asked to staff above any axis (locked C2.2).

**DEFERRED but modular: `external_send`.** Not an initial-release seat axis (no initial-release seat initiates an external send — the conductor is the only external sender the **governance surface** offers, i.e. conductor-governed egress, mirroring the m-3 chokepoint / ARCHITECTURE §C2 egress relabel; **D5 residual**: a same-uid shell/curl send is out of governance scope; egress is governed by the locked m-3 chokepoint + `away_bridge_eligible`, §8). Reserved as an **additive 4th axis** (`none < internal < external`) arriving with the literal external-send actuator (Step-4/5); its absent-default `= none` (fail-closed AND accurate — pre-axis seats never had send authority). Adding it = one key + one consumer + its absent-default; **no re-cut** (GL-2).

---

## 6. Per-archetype invariant composition (LOCKED candidate composition)

| seat_archetype | write | dispatch | tool | routing-prior (m-4 bucket) | behavioral mode | default human-mode (posture) |
|---|---|---|---|---|---|---|
| sensor | read_only | none | none | fast_cheap | answer-once | away |
| solo_worker | write_feature | none | exec | coding | execute | interactive |
| implementer | write_feature | none | exec | top_tier_coding | execute | interactive |
| planner | read_only | route | read | strong_reasoning | plan/design | interactive |
| reviewer | read_only | none | read | strong_general | adversarial | interactive |
| orchestrator_lead | read_only | route | read | strong_reasoning | decompose/route | interactive |

(`actuator` = a derived class: `write_feature`+`exec` + heavy gates, §7.2.) Ceilings are MAX points; templates tighten. The routing-prior is the m-4 capability-prior bucket keyed on `seat_archetype` (m-4-owned; m-5 supplies the key). Dispatch cells carry only lattice tokens (`none < route`, §5); the `planner`'s `route` is scoped to design-author dispatch by the T2 template (§7), a template tightening expressed outside the axis value, not a distinct dispatch point.

---

## 7. Templates (T1/T2/T3, LOCKED structures) + sensor/actuator

**Template schema (m-5-owned, data — promotes codex/claude-code/CrewAI presets-as-data):**
`{ template_id, seats[], topology, gate_set, human_mode, pane_layout }`; each seat = `{ seat_name, role, seat_archetype, pane_intent, model_slot }`. m-4 fills `model_slot` (pinned|slot+bucket) + emits `routing_decision` + `template_ref`; no-bypass unchanged. conductor-core spawns panes.

| # | template | seats / topology | panes | gate-set | read-only-ness | model slots |
|---|---|---|---|---|---|---|
| **T1 solo** | 1 `solo_worker` (or role-specific `implementer`); no peer | 1 | by `slot_in`; observe-as-send gate always-on (the `slot_in` invariant *families* bite from Step-2, §9) | mutating, bounded by ceiling | 1 (coding bucket) |
| **T2 adversarial_pair** | `planner` ↔ `implementer`; review/fold lineage | 2 | + design/plan/review/fold by phase; `slot_in` work-invariant | planner read/design-leaning | planner=strong_reasoning, implementer=top_tier_coding (diverse-model adversarial) |
| **T3 sensor** | 1 `sensor`; non-interrupting side-branch, NOT in delivery lineage | 1 side-surface | none (read-only ⇒ no delivery/merge gate) | strict read-only, tool-blocked, 1-turn | fast_cheap |

The **conductor/N-pair template** is DEFERRED to Step-5 (no initial-release consumer; nested/recursion/scheduler pressure not needed to bind the atoms).

### 7.1 Sensor (full)
Read-only ceiling `(read_only, none, none)`; single-turn; shared read-only context fork; **non-interrupting** (lane continues); answer → separate non-lane surface (never the lane inbox). **Integrity split (locked C2.5):** answer *content* = `self_reported`/advisory/**never gate-bearing**; observable *metadata* (tool-blocked, no-source-actions, 1-turn) = `observed`. **Routing-prior = `fast_cheap`** (locked C2.5) + a recorded caveat (claude-code `sideQuestion` inherits the parent model for a cache-hit, `sideQuestion.ts:82-84`) → an **operator-overridable per-question escalation** for a domain-heavy side-question (m-4 boundary). Direct prior art: `runSideQuestion`/`runForkedAgent` (`sideQuestion.ts:53-102`); our net-new = the **declared, seat-stamped** sensor (not an anonymous self-fork).

### 7.2 Actuator (DERIVED class for the initial release; literal seat = Step-4/5 carry, GL-1)
`actuator` is a **derived mutating-ceiling class** — a label over a write-bearing ceiling (`write_feature`+`exec`) with heavy gates — **not** a literal `seat_archetype` value in the initial release. The **read-only→write boundary is hard human-gated**: a `sensor` **emits work into a separately-spawned actuator** (decision-brief / candidate-work-item into a human/conductor queue); it **never upgrades in place** (E1: the pre-build design-state export's adaptive-routing-pillar doc, lines 56-58). The **literal `single_bounded_action` actuator seat** (tight 1-action/1-tool effector, gate-bearing actuation) is a **Step-4/5 carry** — its single-action ceiling is not enforceable on a ride-existing-runtime (§9); it arrives with the deferred `external_send` axis (§5).

---

## 8. The m-6 boundary contract (seam resolution — cite the COORD; declare-before-bind)

m-5 DECLARES; m-6 BINDS surface/scheduler behavior after. **RECONCILED both sides** in `c3-design-m5-m6-coord`: m-6 binding-confirm `123022` (bound the four non-gate `surface_intent` classes + gate-bearing-carries-none + m-6-owned `away_bridge_eligible`) ⊕ m-5 confirm `131856` (conformed to m-6's bound set; retracted the crossed `125604` `{verdict, fyi, collaborate}`). The earlier `122628`/`122616` relays were crossed drafts, superseded.

**Human-mode = two orthogonal layers (m-5-declared):**
- **posture** `{ interactive, away, unattended }` — a per-template/per-seat default; rides the F2 per-assignment record home (no new m-2 field).
- **surface_intent** `{ progress, review_checkpoint, advisory, result }` — a **conductor-DERIVED**, **TOTAL** (exactly one per record) delivery-class over **non-gate records only** (no new m-2 field, no micro-fold; derived like m-3's `record_integrity` rollup). Each is binding-distinct (m-6-validated, `c3-design-m5-m6-coord/COORD-planner-20260630-123022.md`):
  - `progress` ⇐ ambient/ephemeral checkpoint (no park; batchable digest);
  - `review_checkpoint` ⇐ pair/design-review context → the collaboration/meeting lane (do not compress; re-observe-on-resume);
  - `advisory` ⇐ the sensor side-question answer surface (Seam B; never gate-bearing);
  - `result` ⇐ terminal deliverable (retained + navigable);
  - `progress ⇐ otherwise` makes the derivation total.
  **Gate-bearing records carry NO surface_intent** — delivery is fully determined by the **locked** `HUMAN_GATE_REQUIRED + gate_category` (A/B, J2) + J1; m-6 binds A-bucket → ODB off those locked fields (no duplicate-mechanism-as-vocabulary; this is why there is **no `verdict` value**). **m-6-owned gate→lane routing (no new field):** a gate-bearing record with `phase ∈ {DESIGN, DESIGN-REVIEW}` ∨ `GRILL_REQUIRED=yes` routes to the collaboration/meeting lane (NOT compressed to the ODB), derived off locked `{gate_category, phase, GRILL_REQUIRED}`. m-6 binds surface behavior to the **(posture × surface_intent)** pair (two independently-readable atoms) + the locked gate fields. The **bounce (author-facing)** is an m-3 observe-veto outcome, **not** a surface_intent.
- **`away_bridge_eligible`** = an **m-6-owned per-gate boolean** (m-6/operator policy: opt-in, A-only, egress-gated, locked J1). In the initial release **m-5 declares only the `away` posture**, not this flag. A **hard never-bridge capability ceiling** (for an archetype whose gate's very *existence* is sensitive, beyond content-egress scanning) is a **RESERVED future m-5 hook, NOT an initial-release ask** (m-6 binding-confirm `123022`). NOT a posture/surface_intent value, NOT a write-axis (§5).

**Seam B — interjection.** m-6 owns the surface (steer/side-question/interrupt composer + the separate non-lane answer rendering); m-5 owns the `sensor` archetype + declares **`accepts_interjection { steer, interrupt, side_question_target }` per archetype** (1-turn sensor: no steer/interrupt, IS the side-question target; long-lived implementer/orchestrator: all three; reviewer: interrupt; actuator: interrupt-before-commit = §7.2 grill cell); m-4 routes the side-question `fast_cheap`; the runtime owns injection/cancel/fork (Step-3+).

**Seam C — m-1 inbound away-mode verdict-token bridge.** m-6/m-1-owned; m-5's only touchpoint is the `away` posture (the bridge's trigger). LOCK-BLOCKING for the away-mode binding only (VP F4); m-6 drafts the bounded m-1 question, orchestrator routes it; **gates nothing in m-5's domain.**

---

## 9. Step-1-rideable + tiered ceiling enforcement (route the dependency, not a blocker)

The archetype system DESIGN is Step-1-rideable as a **recorded governance contract**; ceiling ENFORCEMENT is **tiered**:
- **Step 1 (ride existing runtimes):** archetype + ceiling DECLARED + recorded per-assignment (F2) + enforced **best-effort by host config** — real on claude-code's agent-type `tools[]`/`disallowedTools[]` allowlist (`agentToolUtils.ts:122-225`); partial/host-dependent elsewhere. The governance VALUE (recorded, auditable, fail-closed routing-refusal) holds regardless.
- **Steps 4-5 (standalone runtime):** the conductor enforces ceilings + topology uniformly (ROADMAP Step-5 "a human picks a workflow at spawn that sets the authority ceiling"). The literal single-action actuator (§7.2) + the `external_send` axis (§5) + uniform tool-blocking specifically need this.
- **Observe-invariant phasing (mirrors the ceiling tiering — resolves the §7-T1-vs-§9 read):** the observe-as-send **send-gate is the locked m-3 chokepoint *by design***, but in **Step-1 the observe hook is inert** — no observe writer (m-1 §5: "Step-1 records carry **no observe gate**"; m-7 NF-S5 / the CQ-1(a) step-gate), so Step-1 sends pass the chokepoint **without an observe predicate** (the observe pipeline + done-predicates land at Step-2). The per-`slot_in` observe-invariant **families** (§4/§7 — refactor `test-files-unchanged`, bugfix `red→green`, …) **bite from Step-2**, when m-3's observe pipeline + done-predicates land. Step-1 records the `slot_in` classification (F1) and rides the always-on gate; the invariant families enforce as m-3 ships them — the same **declared-now / enforced-as-the-tier-lands** shape as the ceiling above. So T1's "observe-as-send always-on" (§7) names the *gate mechanism* (Step-1), not a claim that every `slot_in` family predicate is enforced pre-Step-2.
- **No re-cut:** the ceiling is already recorded per-assignment (F2) as an open named-axis map (§5); standalone enforcement reads the same fields. **Dependency routed to the orchestrator** as a later-step build concern.

---

## 10. Consumer boundary contract (no orphan)
- **Writes:** the archetype system (tag-space + composition + ceiling vector + templates + sensor/actuator), as config-sourced data.
- **→ m-3:** `slot_in` selects the observe-invariant family (§4); conductor-classified-at-acceptance (F1).
- **→ m-4:** `seat_archetype` keys the capability-prior + the ceiling vector (§5/§6); templates feed the GL-4 record mechanism (`template_ref`, no-bypass).
- **→ m-6:** the two-layer human-mode vocabulary (posture × surface_intent) + the sensor/interjection contract (§8); `away_bridge_eligible` is **m-6-owned policy** (m-5 declares only the `away` posture; hard never-bridge ceiling = reserved future hook) — reconciled in the COORD.
- **→ conductor-core:** pane-spawn via existing multiplexer (Step-1).
- **No-consumer check:** every archetype/template/value has a named consumer + an m-3/m-4 mechanism. The one flagged item — the literal single-action actuator's enforcement — has no Step-1 mechanism (recorded contract only), routed to orchestrator (§9); its consumer (m-3 observe of the actuation outcome) exists.

---

## 11. Novelty statement (precise; no overclaim)
**Promoted, not built:** the preset+selection+justification pattern (upstream); presets-as-data (codex/claude-code/CrewAI); the AUTHORITY/tier + tool-allowlist ceiling primitive; the topology query-surface; the agent-type shape + the `sideQuestion` sensor-invariant; the per-work-kind gate analog (RefAgent/VeriMAP). **Genuinely new:** the **archetype as ONE governed unit binding {topology + gate-set + authority-ceiling-at-spawn + observe-invariants + routing-prior}**, riding m-1's append-only lineage with a seat-stamped FROM + a per-assignment recorded, fail-closed-enforced ceiling — the template ⊥ instance ⊥ append-only-trace separation bound to audit-first immutable lineage, which the external survey separates conceptually but never binds. We do NOT claim novelty in topology, presets, or per-task gating individually (each is prior art); the contribution is the **governed, auditable integration.**

---

## 12. GRILL_LOCK — c3-grill-m-5

GRILL_LOCK_ID: c3-grill-m-5
GRILL_REQUIRED: yes
GRILL_SOURCE: c3 audit + F4 reconcile; m-3/m-4 locked designs; prior-art sweep (3 source agents + 95-subagent deep-research); m-5↔m-6 COORD; operator design-grill 20260630.

Resolved decisions (operator, final authority; 20260630):
- **GL-1 — actuator = DERIVED ceiling-class for the initial release** (NOT a literal `seat_archetype` value); literal `single_bounded_action` seat = Step-4/5 runtime-enforcement carry. Read-only→write hard human-gated; sensors emit into a separately-spawned actuator. — source: operator + orchestrator steer.
- **GL-2 — ceiling vector = 3 axes (write · dispatch · tool), partial order, OPEN named-axis map** with per-axis monotonic tightening + per-axis absent-default; `external_send` DEFERRED but **modular** (additive 4th axis, absent-default `none`/fail-closed, arrives with the literal actuator; no re-cut). — source: operator (explicit "build modularly so adding it isn't a PITA").
- **GL-3 — read-only/low-mutation work-archetype ship-set = ship all three** (`research_synthesis`, `qa_review`, `docs_chore`); each has a distinct m-3 done-predicate. — source: operator.
- **GL-4 — surface_intent = `{progress, review_checkpoint, advisory, result}`** (non-gate delivery-classes; m-6-validated binding-distinct, `c3-design-m5-m6-coord/COORD-planner-20260630-123022.md`), all **conductor-DERIVED + TOTAL** (`progress ⇐ otherwise`); **gate-bearing records carry NO surface_intent** — they bind off the locked `gate_category`/`HUMAN_GATE`/J1 (so there is **no `verdict` value** — resolves no-duplicate-locked-mechanism); m-6 routes design/grill gates to the collaboration lane off locked `{phase, GRILL_REQUIRED}`. — source: operator (20260630; corrected after reading m-6's crossed binding-confirm `123022` — the earlier `{verdict, fyi, collaborate}` was m-6's superseded `122616` proposal).
- **GL-5 — naming = `lower_snake_case` roster** (§3). — source: operator + orchestrator steer.
- **GL-6 — sensor routing-prior = `fast_cheap`** (C2.5-locked) + operator-overridable per-question escalation for domain-heavy side-questions; cache-vs-cost caveat recorded (m-4 boundary). — source: operator.

Rejected / deferred:
- literal single-action actuator seat in the initial release — deferred (not ride-runtime-enforceable; Step-4/5).
- `external_send` as an initial-release seat axis — deferred (no initial-release consumer; conductor = the only external sender the *governance surface* offers, D5 residual); reserved modular.
- conductor/N-pair template in the initial release — deferred to Step-5 (no initial-release consumer).
- any m-2 `required_when`/`visible_when` over concrete tag values — rejected (C2.4 no micro-fold).

---

## 13. PLAN carry-forwards (future build cycle — NOT this phase)
- the archetype registry as config-sourced data (templates + tag-space + ceiling maps); the conductor-classifies-`slot_in`-at-acceptance ordering (m-3 PLAN carry).
- the open named-axis ceiling map + per-axis absent-default + the additive `external_send` axis path; the literal single-action actuator + standalone-runtime ceiling/topology enforcement (Step-4/5).
- negative fixtures: a lane re-tag-to-escape attempt blocked by conductor-classification (refactor test-files-unchanged; bugfix red→green); a `route_dispatch` fail-closed above-ceiling refusal; a sensor that cannot upgrade-in-place; a template off-floor slot recording a deviation (no-bypass).

## 14. Lock prerequisites (open at the c3 lock)
1. m-5.implementer DESIGN-REVIEW approve — **DONE** (`133831`; rev-1 `131617` + rev-2 `132748` must-revise folded).
2. The m-5↔m-6 COORD seam reconcile — **DONE** (reconciled both sides: m-6 `123022` ⊕ m-5 `131856`; surface_intent four-class non-gate, gate-bearing off locked fields, away_bridge_eligible m-6-owned, interjection composer bound).
3. The Seam-C m-1 confirm-or-gap (m-6-led) — lock-blocking for the away-mode binding only; gates nothing in m-5.
4. Orchestrator records the Step-4/5 carries (literal actuator + external_send axis + standalone enforcement) as later-step build dependencies, not c3 blockers.

---

## 15. c5 claim-sweep fold (claim-text hygiene — NO mechanism change, NO locked-contract reopen)
Applied the ratified confusion-resistant checklist (`c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320`, VP-approved) to this doc under `c5-claim-sweep-light`. **Claim-text only; every mechanism, lock (GL-1..GL-6), and F1/F2 invariant is byte-unchanged in substance.** Full-net classified survivor list:

**RELABELED → confusion-resistant + D5 residual (5):**
- **§4 heading (`:62`)** "tamper-resistance proof" → "confusion-resistance proof (observer-selected control property)" — keeps the observer-selection substance; drops the malicious-containment word.
- **§4 proof (`:78`)** "tamper-resistant invariants cannot be escaped by lane re-tagging" / "the whole tamper-resistance" → "a *confused* lane has no tool to re-tag its `slot_in`"; **KEPT** the observer-selected-control framing (the licensed KEEP class, ARCHITECTURE §C3:277) + added the explicit **D5 residual** (a genuinely malicious seat writing a false store record / reaching a shell under attach is out of scope for the initial release, bounded by m-1 isolation + Step-4/5 tool-ceiling).
- **§1 F1 ref (`:19`)** "non-lane-writable" → scoped to confusion-resistant (no lane tool writes/re-tags) + D5 residual pointer; mirrors ARCHITECTURE F1 relabel.
- **§5 external_send (`:97`)** "the conductor is the **sole external sender**" → "the only external sender the *governance surface* offers" + D5 residual (same-uid shell/curl out of scope); mirrors the m-3 chokepoint / ARCHITECTURE §C2 egress relabel. *(This hit was not in the orchestrator's m-5 census; caught by the required own full-net grep.)*
- **§12 Rejected/deferred (`:198`)** the hyphenated shorthand "conductor **sole-sender**" → "conductor = the only external sender the *governance surface* offers, D5 residual"; same claim class as §5 `:97`, matched by extending the sweep net to the hyphenated forms ("sole-sender", "sole sender"). *(Surfaced in the c6 seam-byte-integrity pass; the c5 regex net did not catch the hyphenated shorthand.)*

**KEEP (justified — no edit):**
- **§5 ceiling (`:92`)** "an `orchestrator_lead` routes but cannot write" — **authority-ceiling** (the dispatch axis grants no write *tool*); the checklist's named KEEP exemplar (`131320:32`).
- **§5 (`:95`,`:161`,`:178`,`:190`)** `route_dispatch` **fail-closes** above ceiling / "never loosen above" / "fail-closed-enforced ceiling" — authority-ceiling + monotonic-lattice structural properties. KEEP.
- **§1/§7/§9 no-bypass refs (`:21`,`:119`,`:170`,`:207`)** GL-4 routing-template record grammar "no-bypass" — **m-4-owned** record-emitting grammar (no lane verb skips the record; confusion-resistant); mirrors m-4's own c5 sweep (`c4/c5` m-4 `:76`/`:360`); D5 residual carried at m-4. No local edit (not m-5's claim to rewrite).
- **§1/§2 (`:15`,`:31`,`:44`,`:178`)** "append-only / immutable lineage" — **append-only** is the genuine structural store property (KEEP); "immutable lineage" is a consumed **m-1** store ref → mirrors m-1's governed-write relabel; residual carried at m-1.
- **§6/§7 (`:109`,`:124`)** "adversarial" (reviewer / T2 adversarial_pair) — diverse-model **review posture**, not a security-containment claim. KEEP.
- **§9 negative-fixture names (`:207`)** "re-tag-to-escape attempt **blocked** by conductor-classification" / "sensor that **cannot upgrade-in-place**" — negative **fixture** names (cf. ARCHITECTURE `:125` KEEP): they name what the fixture asserts (the observer-selected classification property + the sensor authority-ceiling); inherit the same confusion-resistant scoping + D5 residual as the underlying claims they test. KEEP.

Invariants: R2 held (no m-5 field is gate/model-referenceable); C2.4 held (no tag-value micro-fold); F1/F2 substance unchanged; no m-4 record / `ARCHITECTURE` change; no c2/c3 reopen. m-5.implementer semantic-approve **recorded** — verdict approve (`c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134809`).

---

## 16. c6-fix-m-5 fold (re-review cleanup — DOC-ONLY; NO mechanism change, NO design-lock reopen)
Applied the `c6-fix-m-5` dispatch (`c6-fix-m-5/DESIGN-orchestrator-planner-20260702-204512`; source `master/DESIGN-REREVIEW-2026-07-02.md` §5 + appendix, VP-concurred CONDITIONAL-GO). **Consistency folds INTO the locked doc; lock invariants (GL-1..GL-6, F1/F2, R2, C2.4) unchanged.** Preserves the four sanctioned by-construction claims, the swept confusion-resistant / D5-residual vocabulary, and the byte-exact `{accepted, rejected, held}` enum.

**Applied (m-5-owned, my domain judgment):**
- **m-5-F6 (M, cross-doc-contradiction — §7 T1 vs §9):** T1's "observe-as-send always on" read as contradicting §9's tiered/best-effort Step-1 framing. **Fix:** annotated §7 T1 cell + added a §9 phasing bullet — the observe-as-send *gate mechanism* is always-on from Step-1; the per-`slot_in` observe-invariant *families* bite from **Step-2** (when m-3's observe pipeline lands), mirroring the ceiling's declared-Step-1 / enforced-later tiering. No new mechanism; a phase annotation only.
- **m-5-F7 (C/OBS, gap — §5):** §5 mandated per-axis absent-defaults but declared one only for the deferred `external_send` axis. **Fix:** folded one sentence into §5 — every axis MUST declare an absent-default = its most-restrictive floor (fail-closed): `write→read_only`, `dispatch→none`, `tool→none` (`external_send→none`); a record-absence reads as the minimum grant, never an escalation. Makes explicit what GL-2's monotonic-tightening already implied; no lattice change.

**Coordinated (seam, addressed COORD not a self-edit):**
- **m-7-F2 / x3-seam-byte-integrity-F2 (B, CTO:seam — trusted-config author set omits m-5):** the trusted-config artifact (CQ-4b: per-domain sections, each stamped, one top-level digest) named the m-5-owned archetype registry but omitted m-5 from its author set. **m-5 CQ-4b confirm** filed to m-7 (`c6-fix-m-5/COORD-planner-20260702-…`): the archetype registry (§3 tag-space + §6 ceiling composition + §7 templates + §5 ceiling maps) is an **m-5-authored, m-5-stamped section** of the artifact, loaded at genesis (already my §13 config-sourced-data carry) — no mechanism change, consistent with the c3 lock. m-7 amends its §7/S15 + ARCHITECTURE C4.1 (their doc, their dispatch `c6-fix-m-7`).

**Verified CTO-already-applied (did NOT redo):** m-5-F4 (status header → LOCKED per c3-lock, `:7`); m-5-F5 (§6 token normalization — tokens lower_snake_case/consistent); x2-claim-honesty-F5 + x3-seam-byte-integrity-F6 (§12 `:198` hyphenated "sole-sender" relabel + §15 survivor-list entry, `:225`). All present and correct.

**Flagged out-of-scope (NOT applied — scope confirm requested from CTO):** **m-5-F2** (posture-vocab away-trigger) is named in the dispatch's canonical-resolution guidance but `owner: CTO:seam`, absent from the dispatch's "Your findings" table + the "3 findings" framing, and absent from the CTO-already-applied list. posture is a **locked c3 vocabulary** and the away-mode trigger touches the m-6-owned `away_bridge`/Seam-C boundary — so I did not re-record it unbidden. If the m-5 posture-vocab slice is mine, route it as a bounded seam COORD and I will apply within the locked lattice + coordinate with m-6.
