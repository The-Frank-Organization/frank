## PLAN — Slice-5 build dispatch (master → s5 slice-team, THROUGH FRANK; consumer schemas + the owed Step-1 fixtures — the LAST Step-1 slice, and the first team to run its governance on frank)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s5-dispatch
PARENT_DISPATCH_ID: step1-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — s5 runs on F2; the s5-close sign-off is the operator's, exercised separately; §7 config-change records are operator-authorized by design
DELEGATED_DISPATCH_AUTHORITY: yes — conditions below (F2, non-bootstrap)
GRILL_REQUIRED: no
IN_REPLY_TO: .relays/s4/s4-merge-gate/MERGE-GATE-implementer-20260705-234649.md
FROM: master.orchestrator-planner
TO: s5.orchestrator-planner, m-2.planner
CC: master.orchestrator-reviewer, operator, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-1.planner
SUBJECT: s5 = consumer schemas (m-3/m-4/m-5/m-6 fields, declared + STEP-GATED OFF) + the five owed §C4 Step-1-build fixtures + schema_version/migrator, landed as real §7 records on the LIVE wired store; NEW slice-team on frank; guide m-2; m-x via master; F2 conditioned delegation

**What this is.** The master dispatch for **Slice-5 = the consumer schemas + the owed Step-1-build fixtures** — the **LAST Step-1 slice**. Two firsts ride on it: (1) s5 is the **first slice-team to run its own governance ON frank** — this dispatch reaches you through `submit`/`project`/`read`, and your plan / plan-review / IMPL reports / exit SITREP all flow through frank, no hand-relay; (2) its registry additions land as the **first real §7 config-change records on a live wired store** — the first genuine usage data. **VP pre-handoff review: APPROVE** (`s5-dispatch/RECONCILE-orchestrator-reviewer-20260706-004155`); its seven watchpoints are folded inline below as **[VP-W1..W7]**. Baseline: **S1–S4 CLOSED** (`main`, post-`s4-close`; battery 21 green; the operator-as-transport is over). **Step-1's mandatory owed set is EMPTY** (`OI-S3-CONFIG-CHANGE` discharged at s4). Nothing rides in owed; s5 adds the consumer envelope + discharges the §C4 owed *fixtures*.

**The goal (one line):** make frank's typed-envelope registry **complete + versioned** — declare the four consumer domains' fields (m-3 observe · m-4 routing-record · m-5 archetype · m-6 gate/ODB), **step-gated OFF** (declared, dormant, no Step-1 writer), land the five owed §C4 fixtures, and ship `schema_version` + a migrator — so Step-2 (observe) inherits a registry that already knows every field it will turn on, and the public release versions schemas from day one.

### To the s5 slice-team — your charter
- **NEW slice-team** (new sprint = new team). **Use `/orchestrator-planner`**; scaffold an **`s5`** sprint via `sprint-doc-setup` in `frank/`. **BUT your governance transport is frank, not files** — your relays are `submit`ted (audits, plan, plan-review, IMPL reports, review-folds, the exit SITREP). The `frank/` sprint dir holds your *code* work + working notes; the *governance* is the frank store. (Where a sub-seat review must be hand-relayed because a sub-seat isn't yet wired, that is the operator's mint call — ask via master.)
- **Onboard first — you built none of S1–S4.** Read the source + the four sprint ledgers in `frank/`; re-run the battery at `s4-close` yourself (uncached). The standing bar: every fresh slice-team has found real fragility the builders missed. **m-2 (your guide) owns the FieldSpec registry** — the consumer schemas are new registry rows; m-2 sets the row grammar (owner·type·required_when·enum·`gate_referenceable`).
- **Spec = read-only under `master-docs/master/`:** `ARCHITECTURE.md` §C4 (esp. C4.2 the engine×six-contracts seam matrix + C4.3 claim boundary/I-PH + the "Owed Step-1-build fixtures" list) · the **m-2** domain doc (FieldSpec registry, §4 `gate_referenceable`, the phase-split required-set) · the **m-3/m-4/m-5/m-6** domain docs (the consumer field *semantics* you'll declare) · the **m-7** doc (the §7 config-change class + I-PH / ③ host side). Escalate spec problems; do not self-amend.
- **Build on `main`** (post-`s4-close`), on a branch the operator sets up as your worktree; close-time integration is the operator's separate gate.

### THE ROUTING DIRECTIVE (load-bearing — you are on frank, the m-x are not)
**Every guide / fidelity / m-x question routes to `master.orchestrator-planner` via frank** — `submit` a relay `TO master.orchestrator-planner`. The domain seats **m-1…m-7 are NOT minted on frank**; your recipient picker will not offer them and you must not address them directly. Master routes your question off-band to the right owner — **m-2** (registry grammar, the guide) · **m-3/m-4/m-5/m-6** (consumer-field *semantics* fidelity) · **m-7** (the §7 seam, I-PH, ③ host side) · **m-1** (store/identity fidelity) — and returns the answer to you via frank. **No guide/fidelity engagement happens without a relay to master first. [VP-W2] Silence from an off-frank m-x seat — or the absence of a direct frank recipient — is NEVER permission to choose field semantics: route the uncertainty to master, or hold.** Hub-and-spoke: you ↔ master ↔ the m-x. (This constraint IS frank's addressing model working — it is part of what s5 dogfoods.)

### Guide + contract boundaries
- **m-2 primary guide** — the FieldSpec registry is m-2-owned; every consumer row is an m-2-grammar row. m-2 rules the `required_when` predicates, the enum value-sets, `gate_referenceable` per column, and the **phase-split required-set** that keeps a Step-1 gate from ever demanding a consumer field.
- **m-3/m-4/m-5/m-6 consulted (content fidelity, via master)** — they own what their fields *mean*: m-3 the observe/evidence + egress fields (incl. the ⑤ ODB `model_name` carve-out), m-4 the routing-record columns (incl. `gate_referenceable: false` on `chosen_model` — the R2 no-model-predicate invariant), m-5 the archetype/slot fields, m-6 the gate/ODB + `GRILL_REQUIRED` field. **You declare the schema; they certify the semantics.** All via master routing.
- **m-7 consulted** — the §7 config-change record class (your registry additions ARE §7 changes on the live store), plus the **m-7-hosted** ③ and **I-PH** fixtures.
- **m-1 fidelity** — any store/record touch.

### Slice-5 scope (IN)
1. **Consumer field schemas — declared, STEP-GATED OFF.** Register the m-3 observe · m-4 routing-record · m-5 archetype · m-6 gate/ODB fields into the FieldSpec registry (owner·type·`required_when`·enum·`gate_referenceable`). **Dormant in Step-1:** no writer, and the **phase-split required-set never demands them** (CQ-1(a) — a Step-1 gate must never require an observe-owned field with no writer). Proven by a **negative fixture**: a Step-1 phase form neither requires nor renders a consumer field. This makes the registry *complete* — Step-2 turns fields on, it does not invent them.
2. **The five owed §C4 Step-1-build fixtures** (the concrete deliverables that make the schemas real; all under `ARCHITECTURE.md` §C4 "Owed Step-1-build fixtures"):
   - **③ known-A / RAISE-ONLY direction-invariant NF** (m-7-hosted): a B-pick / B-absorb over a known-A category ⇒ raised to A + `gate_category` recorded, never silently orchestrator-absorbed.
   - **⑤ ODB model-name egress** (m-3 (a)/(b)/(c) set): a model-name in the exempt ODB `model_name` slot passes the egress scan; a model-name outside the carve-out is scanned/blocked.
   - **m-2 `GRILL_REQUIRED` FieldSpec row** (m-6-F6): m-2 declares the ported upstream `GRILL_REQUIRED` header's FieldSpec (owner/type/values) so the m-6 meeting-lane route can bind it as a field.
   - **§J2 explicit `routing_escalation` A-member** (m-4-F7): the distinct `gate_category` member for routing-escalation force-A telemetry. **NOT a blocker** — correctness already holds via the `other`→A fail-safe (m-4 §7:365-367); this is clarity/telemetry. It is a **CTO-owned cross-domain add** (I author the §J2 member + m-2 mirrors + m-6 confirms, via routing); you register the resulting enum value + a fixture. Token is `routing_escalation`, **distinct** from the `routing_unavailable` route_dispatch outcome state. **[VP-W5] s5 registers ONLY the owner-returned enum/member decision; if the CTO/m-2/m-6 route has not yet produced the exact token + m-2 mirror, s5 HOLDS or escalates — never self-authors the cross-domain value.**
   - **I-PH path-hygiene** (external-review): no seat-delivered surface — bounce · error · rendered projection · delivery payload — contains a canonical store/config/outbox path. m-7-hosted; m-1/m-2 honor it in store + bounce/reason text.
3. **`schema_version` + migrator** (public-release-intent): the registry additions bump a declared `schema_version`; ship a **migrator** old-store→new-registry + a **replay** proving an existing store (e.g. a stock s4-era store) migrates cleanly, zero record loss, canonical wins. Zero migrators in the field yet — version + migrate from day one.
4. **Land the additions as real §7 config-change records on the LIVE wired store** (the dogfood usage data): each registry addition is an **operator-authorized** §7 digest-change record on an EXISTING store — never re-genesis. This exercises the s4 §7 mechanism for real and produces the first genuine config-evolution audit trail.

### Slice-5 scope (OUT — escalate before any delegated dispatch that touches these)
- **observe-as-send / evidence *execution*** (Step 2) — you declare observe *fields*; the observe *hook/mechanism* is Step-2. Done-state + `record_integrity` stay **`self_reported`**.
- **routing *execution*** (Step 3) — you declare routing-record *columns*; the router is Step-3.
- **archetype spawn *execution* / TUI** (Step 4/5).
- **The step-(d) away-bridge gates — explicitly NOT S5** (they ride the away-bridge build step, per §C4): Decision ④ away-token ROTATE+RE-OBSERVE fixtures · R2 `gate_referenceable` per-column *negative fixtures* (the attribute is declared; the model-predicate-grammar fixtures are step-(d)) · Altitude-B per-row deviation-grain fixtures · away-mode trigger expressibility. **Do not build these. [VP-W4] `master-docs/master/ARCHITECTURE.md` §C4.2 / §C4.3 + the §C4 carry ledger is the controlling S5/step-(d) source of truth — do NOT inherit stale kickoff shorthand that would pull R2 per-column negatives or away-bridge work into S5.**
- federation (horizon) · external send / away-bridge (outbox stays local, egress dormant) · any replacement of the operator's *authority* (gates still park for the human).

### Slice-5 exit gate (HARD acceptance)
- **Registry complete + dormant [VP-W3]:** the four consumer domains' fields declared; a **negative fixture** that **enumerates the consumer-owned fields under test** and proves each is **absent from every Step-1 rendered form surface** (not merely un-required by one happy-path submit) — the phase-split required-set holds (CQ-1(a)).
- **The five owed fixtures GREEN:** ③ · ⑤ · `GRILL_REQUIRED` row · `routing_escalation` member+fixture · I-PH (every seat-delivered surface class path-clean — bounce / error / projection / delivery).
- **Versioned + migratable:** `schema_version` bumped; migrator + replay proving an existing store migrates old→new with zero record loss; canonical wins.
- **§7 on the live store [VP-W7]:** each addition evidenced as an operator-authorized §7 config-change record on an EXISTING store, carrying the **old→new digest** — **no re-genesis proxy**; phase-0 accepts the new digest via the genesis chain; a superseded rendered form bounces "re-render" and the re-rendered form succeeds.
- **E2 floors:** full battery green (S1–S4 suites); zero regression; enum byte-exact `{accepted, rejected, held}`; the guardrail surface still exactly `submit` / `project` / `read`.
- **Honesty [VP-W1]:** every s5 evidence surface says **"transport/provenance only"** — done-state + `record_integrity` remain `self_reported` until Step-2 observe. Consumer fields are **declared, not observed**; say so wherever they are documented. frank-carried relays prove governed **transport / provenance / addressability**, NOT observed work.
- **The dogfood record:** s5's own governance ran through frank (this dispatch, the plan, the reviews, the IMPL reports, the exit SITREP all `submit`ted); capture the frank relay-ids as exit evidence that a real team's traffic rode the conductor end-to-end.

### Plan-gate (F2 — non-bootstrap; conditioned delegation)
Pair Implementer plan-review = the plan gate; `DISPATCH IMPL` delegated only under **{Implementer approve · no scope/boundary deviation · no hard trigger · no cross-slice collision · no locked-contract or design-of-record amendment}**; any failure — including any OUT-item touch or any consumer-*semantics* question — **escalates to master (CTO + m-2 guide + the relevant m-x + VP)** via a frank relay. Declaring a consumer field whose *meaning* is unsettled = a fidelity question = route to master; do not guess the semantics.

### Deliverable format
The registry additions + fixtures + migrator on a branch; the exit-gate evidence (fixtures green, replay, the live §7 records, the negative dormancy fixture); your governance relays **through frank**; an exit SITREP `submit`ted to `master.orchestrator-planner`; close-time integration = the operator's gate.

### Operator-judgment items
- **§7 changes are operator-authorized by design** — each registry addition on the live store needs the operator's explicit authorization record (that IS the mechanism working).
- **Sub-seat minting** — as s5 decides its team granularity, the operator mints + wires each sub-seat on frank (or rules a sub-seat stays a hand-relayed file review). Ask via master.
- **residual risk (accepted, restated):** D5 — same-uid direct store write is out of scope (confusion-resistant, not theft-proof).

### Not authorized by this relay
No s5-close authority, no scope expansion, no locked-design amendment, no step-(d) / away-bridge pre-work, no federation pre-work, no Step-2 observe mechanism. `DISPATCH IMPL` only under the F2 conditions; failures escalate via frank.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s5/s5-dispatch` — run below (clean both modes).
- Executes the Step-1 roadmap's final slice: consumer schemas + the §C4 owed Step-1-build fixtures, over the wired conductor. Renumber (old Section-4 → s5) recorded at s4-dispatch. Addressing note **[VP-W6]**: the frank `submit` of this dispatch addresses `s5.orchestrator-planner` (the only non-master minted seat); the off-frank CCs (VP, operator, m-x) are served by **this VP-reviewed master file** — the authority-bearing relay of record (VP APPROVE `RECONCILE-orchestrator-reviewer-20260706-004155`) — plus master routing. The frank-submitted body restates that pointer + the m-x-via-master rule.
- Pointers: `ARCHITECTURE.md` §C4.2 / §C4.3 + "Owed Step-1-build fixtures", the m-2 / m-3 / m-4 / m-5 / m-6 / m-7 domain docs, `frank/` baseline `main` (post-`s4-close`).

ACTIONS_GIT_REF: none — this relay grants plan authority and performs no git action; cwd is not a git repo (docs workspace). The dispatch is authored as this file (the VP-review + archival form) and submitted into frank as the s5-facing transport.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main` (post-`s4-close`), clean.
Next requested action: VP pre-handoff review COMPLETE — **APPROVE** (`RECONCILE-orchestrator-reviewer-20260706-004155`; 7 watchpoints folded above as [VP-W1..W7]); I `submit` this dispatch into frank TO s5.orchestrator-planner; the operator wires + boots the s5 session; s5 onboards, scaffolds, plans, and self-dispatches under F2; exit SITREP via frank.
