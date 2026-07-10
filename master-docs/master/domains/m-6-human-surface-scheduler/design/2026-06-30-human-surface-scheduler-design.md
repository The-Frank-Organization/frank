# m-6 — Human Surface & Scheduler (DESIGN-OF-RECORD)

DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler) — pair: `m-6.planner` (lead) + `m-6.implementer` (adversarial design-reviewer)
PHASE: DESIGN (c3) — design-lock is terminal; no PLAN/IMPL this cycle
STATUS: **c3 LOCKED** (`c3-lock` VP co-sign `RECONCILE-orchestrator-reviewer-20260630-191315`; orchestrator lock `190627`/`191525`) — pair-approved `133839` (rev-2); seam converged (`123022`⊕`131856`); **Seam C RESOLVED (A)** — away-token cell folded over m-1-owned mint/verify + locked (`182600`/`183008`); GRILL_LOCK §9 folded; all cells closed. **c4-cq-gateconfig:** CQ-3 (pure-judgment A-floor table, canonical home `c4-cq-gateconfig/DESIGN-planner-20260702-015800`), CQ-4, CQ-4b — m-6 legs **pair-approved** (`024620`); the CQ-4 terminal-token alignment (`delivery_state {accepted,rejected,held}`; `bounced` retired → `rejected`; `held`→bucket A) is **folded into §2/§4** here so the c3 lock consumes clean tokens. **c5:** claim-sweep relabels folded (§2/§4/§11, claim-text only, `c5-claim-sweep-light`); **operator decision ③ (RAISE-ONLY + known-A detector) folded into §2** (`c5-fold-decision-3`); **decision ④ (ROTATE+RE-OBSERVE) recorded as a non-locking §2C/§12 build-carry** (mechanism at step d, `c5-fold-decision-4`); **decision ⑤ (ODB model-name egress carve-out) folded into §3** (joint m-3/m-4, `c5-fold-decision-5`). Post-lock review-driven consistency folds — c4 CQ (`c4-cq-*`), c5 sweep+decisions (`c5-*`), c6 cleanup (`c6-fix-m-6`) — recorded in the §13 fold-log; **lock invariants unchanged**. No PLAN/IMPL (c3 terminal = design-lock).
BASIS: VP-approved reconciled audit `c3-audit-m-6` (`RECONCILE-planner-20260630-054107`); dispatch `c3-design-m-6/DESIGN-orchestrator-planner-20260630-121325`
SEAM: m-5↔m-6 resolved in `c3-design-m5-m6-coord` — **m-6 bind `123022` ⊕ m-5 confirm `131856`** (the converged FOUR-class model below; the interim `125604`/`131747` `{verdict,fyi,collaborate}` excursion crossed `123022` and was **retracted both sides** per m-5 SITREP `132314` + m-5.implementer must-revise `131617`)
GRILL_LOCK: §9 (operator decisions, 2026-06-30) — referenced from this DESIGN_LOCK
CONSUMES (locked, not reopened): m-1 addressing/stamp + §J · m-2 `gate_category`/HUMAN_GATE/`delivery_state` · m-3 egress/observe/`record_integrity`/ODB-schema · m-4 routing escalation · m-5 human-mode vocabulary + sensor archetype (the seam)

---

## 1. Thesis — promote-and-bind (proven, not asserted)

m-6 is a **thin local-first projection over locked m-1..m-4 records** — not a new gate system, not a new schema, not a TUI (Step-4). Every surface either **binds a locked mechanism** or designs the **one thing the export left undesigned — the away-mode bridge**. Proof obligations discharged per surface below; the no-rebuild lines:
- **ODB content schema** is **promoted from m-3** (`m-3 …design.md:146-151`, the 7-field bundle from agent-scripts `maintainer-orchestrator/SKILL.md:53-69`) — m-6 designs the **render + capture**, not the schema.
- **Egress** is **m-3's** (`m-3 …design.md:116-122`) — m-6 **consumes** the fail-closed scan, never builds it.
- **Identity/stamp + addressing** are **m-1's** — m-6 **projects** the addressing graph into an inbox; the verdict is an **operator-FROM** record (R1).
- **The TUI** is **Step-4** — this doc locks **contracts**; mechanism builds Step-2, full email-client UX Step-4 ("designed early, executed later").

**Roadmap mapping:** design-of-record = Step-0 (now); inbox/outbox + scheduler **mechanism** = Step-2; full email-client **UX** = Step-4 (`ROADMAP.md`).

**The organizing frame — three human-surface lanes** (the governance-vs-collaboration split from the pre-build design-state export, not vendored):
1. **GOVERNANCE** (lane→operator; verdict; minimize friction): buckets + ODB + park/wake + away bridge (§2–§4).
2. **COLLABORATION** (operator↔lane; shaping; maximize richness): the meeting/attach + the **"elaborate more" fork** (§3, §5).
3. **INTERJECTION** (operator→running-lane; redirect): steer/side-question/interrupt (§6).

---

## 2. Gate→email bucket taxonomy (binds locked mechanism; direction-explicit)

The inbox is a **projection of the locked addressing graph** (per the pre-build design-state export, not vendored), not a new store. Four buckets, each with a **locked writer** (no-bucket-without-a-writer); **direction-explicit** — only **A + C** reach the operator.

| Bucket | Trigger (LOCKED writer) | Reaches operator? | Surface behavior | Capture/return |
|---|---|---|---|---|
| **A — verdict-required** | `gate_category ∈ §J A-set` (`ARCHITECTURE.md §J2 A-set`) ∪ `HUMAN_GATE_REQUIRED=raised` (incl. **m-4 routing-raise**, `m-4 …design.md:339-344`) ∪ `egress_scan_result=blocked` (`m-3 …design.md:116-122`) | **YES** — primary decision queue | **ODB** (§3); park; J1 resummon; non-suppressible (monotonic floor) | bounded-choice (buttons) → operator-FROM verdict relay |
| **B — orchestrator-absorbed** | `gate_category ∈ §J B-set` (`ARCHITECTURE.md §J2 B-set`) | **No** (→ A only if a monotonic raiser promotes it) | **live local digest, non-interrupting** (GRILL_LOCK G3) | none (orchestrator resolves) |
| **C — CC-FYI** | operator on **CC** not TO (`README.md:49-59`; export `:37`) | **Informational** | low-priority FYI feed; batchable | none (CC creates no obligation) |
| **D — observe-bounce** | observe-veto: predicate-fail / integrity-veto / form-bounce → `delivery_state=rejected` + `failing_edge` (`m-2 …design.md:256-258`; `m-3 …design.md:61-63`) | **No — author-facing** | returns to the **authoring seat** with the named failing edge | author repairs + resubmits |

**Direction (DQ-3 resolved at grill):** D **stays in the taxonomy** (it is a delivery outcome with a locked writer) but is **tagged author-facing** and **never enters the operator decision queue** (unless the failed edge is itself an A-gate). Only A (alert) and C (notification) reach the operator — the alert-vs-notification split is a property of the **bucket grammar** (it keys on the locked writer + TO/CC direction, not on seat assertion).

**Classification-direction invariant (operator decision ③, RAISE-ONLY — folded c5, `READINESS-REGISTER.md` §Operator-decisions ③).** The `gate_category` agent-pick is **monotonic toward A**: an agent may escalate a B-category to A (more operator oversight), but may **never de-classify an A-worthy decision down to B**. This composes with the CQ-3 A-floor and m-2's monotonic-MAX as the *direction* rule on top of the *floor*, enforced at the same fill/submit point (m-7). **Known-A detector:** a system-side membership list of categories that force **≥A regardless of agent pick** — = the §J2 A-set (the A-floor table's anchoring `gate_category`s, `c4-cq-gateconfig/DESIGN-planner-20260702-015800`). A record whose agent-picked category ∈ known-A but is routed/absorbed as B is **rejected and raised to A**, category recorded — never silently orchestrator-absorbed. The membership is **operator-tunable config** (same pattern as the A-floor table + §J2), with the `other`→A fail-safe hardcoded. This closes the most direct operator-not-surfaced vector (an A-worthy decision mis-tagged B). *(m-7 enforcement flagged, fixture registration owed: the direction-invariant NF fixture — a B-pick / B-absorb over a known-A category ⇒ raised to A + `gate_category` recorded — is an owed m-7-side artifact bound to Step-1 build, NOT yet registered in m-7's design-of-record (locked 2026-07-02, before this ③ fold); until it lands this doc does not claim a registered m-7 fixture. §J ratifies the direction-invariant as a J1/J2-analogous addition — no new gate class, no new mechanism, a direction constraint on the existing HUMAN_GATE monotonic-raise pick.)*

**Terminal-token → bucket map (CQ-4 settled: `delivery_state {accepted, rejected, held}`, `bounced` retired — m-2 §17.1 `014626`; m-7 §6 enum; m-6.impl approve `024620`).** `accepted` → normal delivery (into the addressed mailbox; if the record is an A-gate → ODB **park**, §3/§4). `rejected` → **bucket D** author-return (`failing_edge` names the acceptance-stage veto). `held` → **bucket A** — the ODB/park A-lane (operator-visible, **non-suppressible**), the fault / fail-closed disposition only (m-7 §6 internal-fault-on-authority + CQ-2 class-conditional fail-closed); ordinary `HUMAN_GATE_REQUIRED` A-gate parking is `accepted`, **not** `held` (m-2 §17.1); **`held` never routes to bucket D.** An egress-block on an already-`accepted` A-gate's away-mirror is a distinct outbox event (`egress_scan_result=blocked` → `egress_blocked`, §4 → A local resummon), not a terminal token — the precedence below is unchanged.

**Egress-failure precedence — A vs D, no double-assignment [resolves m-6.implementer must-revise F2 `131702`].** Both an egress block and an observe-veto are "vetoes" in m-3's general sense, so the buckets key on the **`failing_edge` + pipeline stage** to keep them mutually exclusive:
- **D (author bounce)** = a relay rejected at **acceptance**: `delivery_state=rejected`, `failing_edge ∈ {form-validation, lineage, observe-predicate, declared-vs-observed integrity-veto}` (m-3 §3.2(a)(b)). The relay never delivered; the **author** repairs.
- **A (operator, local resummon)** = an **already-accepted A-gate** whose **external away-mirror** was blocked by the egress scan: `failing_edge = egress` at the conductor-governed egress chokepoint (m-3 §3.2(c)/§7; governance-surface, not system-level sole-egress — D5 residual). The gate is **valid and still parked** awaiting the operator's verdict; only its external delivery was blocked ⇒ it **parks locally + raises an A resummon** (the `egress_blocked` state, §4) — **never auto-redact/send, never a D author-bounce** (the author did nothing wrong).
- **Precedence:** egress is evaluated **only** at the external-send chokepoint on an **already-A** gate; the D bounces are evaluated at **acceptance**. Different stage + different `failing_edge` ⇒ a single external-send egress failure resolves to **A (local resummon)**, never D. Local-only relays never reach the egress stage. (An egress-blocked **B**-digest mirror simply **stays local** — no external send, no operator decision, since B is orchestrator-absorbed.)

**Surface realization (Step-2 mechanism / Step-4 UX):** buckets = **saved queries over tags, not exclusive containers** (an item may be `gate`+`m-4`+`urgent`); Gmail-taxonomy preset; additive-score→sort (oldest/stalled first). The **§J A/B map + protected-branch set are operator-config** (`ARCHITECTURE.md §J2`); `other`→A fail-safe inherited (`ARCHITECTURE.md §J2 fail-safe`), never surface-overridden.

---

## 3. Owner Decision Brief — render + capture (promote schema; design surface)

**Render (promoted 7-field bundle + locked evidence):** plain-language change + who-benefits · why-now · **`completed_proof` = conductor-observed `evidence_ref`** with **`record_integrity {observed|self_reported|mixed}` shown** (`m-3 …design.md:107-112`) · tradeoffs/residual-risk/scope · **opinionated recommendation** (agent-scripts `:69`) · **exact bounded choices**. Plus envelope/provenance (`subject_ref`, FROM/TO/PARENT/DISPATCH_ID, thread context — `m-2 …design.md:271-276`) and the actionable-brief enrichments (blast-radius, severity tier, length-capped headline).

**Model-name egress carve-out (operator decision ⑤ — folded c5, `c5-fold-decision-5`; joint m-3 scan + m-6 ODB + m-4 R2-guard).** When an ODB legitimately surfaces a model-name to the operator (e.g. a routing/deviation brief), m-6 renders it in a **typed, exempt-marked field** (`record_kind = ODB`, operator-facing, field = model-name) so m-3's egress scan applies the **confidentiality-class carve-out only here** — every other external send still blocks model-names, and the **safety/content class is NOT exempted**. Transport gating is unchanged: **non-away, the ODB renders locally and never egresses**; **away-mode, the away-bridge opt-in + egress gate still govern the send** (the carve-out only relaxes the *confidentiality* scan on this one operator-facing field, not the opt-in or the safety scan). **R2 is UNTOUCHED** — the model-name stays bookkeeping/payload in the render, **never a gate input** (this is an *egress-confidentiality* relaxation, not a *gate-referenceability* change; m-4 co-confirms). *(flagged to m-7; egress-fixture registration owed — the ⑤ carve-out fixture (the ODB model-name field passes the confidentiality carve-out; all other model-name egress still blocked) is an owed m-7-side artifact bound to Step-1 build, NOT yet registered in m-7's design-of-record; carried in the same build-carry ledger as decision-④ until it lands. Mirrors the m-6-F1/③ "fixture registration owed" fix, §2.)*

**Capture (GRILL_LOCK G2) — buttons + elaborate-more, never a bare decision:**
- The bounded choices render as **buttons** (one per enumerated choice). The choice-set is a **self-describing, capability-scoped gate** realized as m-2 **`agent_enum_pick`** (per the pre-build design-state export, not vendored) — the operator can pick **only a legal verdict**.
- A picked button → an **operator-FROM verdict relay** (R1, `ARCHITECTURE.md §4 R1`) that wakes the parked lane (§4). Optional **free-text note rides *with* a choice** — never a new hidden choice, never a gate-lowering.
- **J1 throughout:** `hold_and_resummon`; **refresh-before-resummon** (agent-scripts `:57` — never re-ask an answered/stale gate; re-observe ground-truth); **never auto-approve** (`ARCHITECTURE.md §J1`).

**"Elaborate more" = a context-preserving collaboration fork (operator-initiated governance→collaboration bridge):**
- When the operator chooses **elaborate more** instead of a verdict button, the conductor **forks the parked lane's stream/context** into an **interactive elaboration session** — seeded with the lane's **running context** (shared context/prompt-cache; the `runForkedAgent` prior art, `jcode-ux-notes.md:48-53`) so the operator elaborates with full context, no re-explaining.
- The **original gate stays parked** (`parked_waiting_human`, J1 — not resolved, not lost, **never auto-approved**) while the fork runs. This is the operator-initiated bridge into the **collaboration lane** (§5).
- The fork is **read-only** (a discussion to *shape* the verdict; it must not mutate the parked lane's state — avoids two writers). **The fork's ONLY output is the decision, relayed back to the original session — it is NOT a work-stream to merge** (operator-confirmed 2026-06-30; a write-capable fork was considered and **declined**).
- On conclusion, the verdict rides the **same bounded-choice → operator-FROM relay** (possibly now better-informed, or the gate is re-shaped/re-issued) → wakes the parked lane → **m-3 re-observes** (newest-authoritative) before the lane proceeds.
- **Ownership:** the **fork primitive is runtime-owned** (same primitive as the Seam-B side-question fork, §6); m-6 hosts the affordance + the forked-session surface. **m-5 seam follow-up (§10):** whether the fork inherits the parent `seat_archetype`/`human_mode` or runs under an `interactive`/`review_checkpoint` collaboration posture — a bounded COORD note, non-blocking.

---

## 4. Scheduler park/wake + away-mode bridge

**Positioning:** the durable **signal + await + timer** shape (Temporal/durable-execution), NOT a blocking prompt — and the conductor's **governed-write append-only store IS the durable checkpointer** (governed-write per m-1's c5 relabel — confusion-resistant + D5 residual) (a parked lane consumes nothing; woken by a durably-recorded verdict relay). **Polling is read-only status only** (export `:34`); the verdict path is the durable mutating channel.

**The 7-state machine (on the durable store):**
`active → parked_waiting_human` (A-gate accepted, ODB emitted) `→ resummon_due` (cadence point, no valid reply) `→ replied_pending_validation` (operator reply captured; form/lineage/choice + operator-FROM channel validated) `→ resumed` (valid reply appended + delivered to the woken lane; **m-3 re-observe on wake**). Plus `bounced_repair` (an acceptance-stage veto → `delivery_state=rejected` terminal → bucket D author-return; the FSM label stays m-6-local, the terminal token is `rejected`) and `egress_blocked` (away-mode outbound scan blocked → park locally + resummon, **never auto-redact/send**).

A `delivery_state=held` terminal (the fault / fail-closed disposition only — m-7 §6 internal-fault-on-authority + CQ-2 class-conditional fail-closed; ordinary A-gate parking is `accepted`, **not** `held`) is **not an m-6 park/wake FSM transition** — it is an m-7-owned runtime disposition that m-6 **renders on bucket A** (operator-visible, non-suppressible, never auto-accept) and resolves by re-run or operator verdict. The 7-state machine stays as above (CQ-4 alignment: `held`→A surface, `rejected`→D).

**Resummon cadence (GRILL_LOCK G4):** two timers (no-response vs answered-but-stalled); **escalate the SUMMON CHANNEL** (local → louder local → away-bridge if opted-in) → **never escalates to a verdict** (J1). **No hard deadline by default**; the gate waits indefinitely for the operator rather than ever auto-approving. Optional per-gate block-ceiling (a gate may make itself more conservative — block, not approve; monotonic, `ARCHITECTURE.md §J1`).

**Away-mode external-inbox bridge (GRILL_LOCK G1; opt-in; the one new design):**
- **Trigger:** operator declares **`away` posture** (m-5-declared, the activation condition).
- **Destination policy (operator-chosen):** **A-gates mirror live** + **B-items mirror as a frequent digest** (operator-config interval, **default sub-daily**); **C stays local**. Opt-in; default = local in-app inbox (`ARCHITECTURE.md §J1 away-mode bridge`).
- **First external send ⇒ egress-gated fail-closed** by m-3 §7 (`m-3 …design.md:116-122`); any block → `egress_blocked` → local resummon.
- **`away_bridge_eligible` = an m-6-owned per-gate boolean**, defaulted by the locked policy (A-bucket ∧ opt-in ∧ egress-clear); m-5 declares only the `away` posture (seam-converged `123022`⊕`131856`). A hard per-archetype never-bridge ceiling is a **reserved future m-5 hook** (§10 OQ-3), not needed for the initial release.
- **✅ RESOLVED (Seam C = A) — the inbound verdict-token bridge.** The bridge converts **untrusted SMTP/IMAP-inbound → a trusted operator-channel verdict record** via a conductor-minted **signed, one-time, per-`(decision_id, seat, choice)` token**; the token signature **is** the operator-channel stamp brought inbound (no new trust *model*; **POST/interstitial, never GET**; replay-nonce ≠ validity-window; fail-closed). **m-1 answered (A)** (`c3-design-m5-m6-coord/COORD-planner-20260630-182600`; orchestrator-accepted `183008`): the mint/verify surface is **TCB-owned**, *forced* by **DI-1** (nonce-burn = an atomic conductor-owned store append; confused lanes have no store-write *tool* — D5 residual: same-uid direct store write out of scope, per m-1's c5 relabel) + **DI-2** (signing key = a TCB secret; m-6 custody would be a forbidden second identity authority). It is the **first activation of m-1's reserved `certification` (DKIM-analog) field** — additive, zero schema change, no new public verb, **not a c1 reopen** (email is the channel-stamp-unavailable case that seam was reserved for). **The A boundary:**
  - **m-1 (TCB) owns:** `mint(decision_id, seat, choice, expiry)` on egress → conductor-signed one-time token (key OS-isolated from lanes, DI-2); `verify(token)` on return → the **five checks** (sig → audience → expiry → nonce-unused → seat-matches-expected), **fail-closed**, nonce-burn = atomic conductor-owned store append (DI-1), seat-match to m-1's minted address space; on pass, stamp `FROM: operator` on the operator-relay channel.
  - **m-6 (bridge, outside the TCB) owns + CALLS the above:** the egress **trigger** + email render carrying the minted token, the **POST-not-GET** receipt endpoint, bucket routing, away-mode UX; m-6 **supplies** `expiry` (**validity-window = m-6/operator policy**), m-1 **enforces** it.
  The away-mode token cell now **locks** with the rest of §4 (the additive later-step build carries are in §12 — **four** from `COORD-182600` + the **c5 decision-④** carry).

---

## 5. Email-governance + meeting-collaboration surface (contracts now; UX Step-4)

**(a) Email-governance surface (governance lane UX).** A calm, keyboard-driven, **low-chrome bucketed local inbox** over the relay store; buckets = §2, ODB = §3. **Looks (proof requirement):** **jcode negative-look** — no GUI noise/filler, **no persistent panels that don't earn their pixels**, explicitly **not** the deferred `/btw` side-panel (`jcode-ux-notes.md`); **codex positive-look** — stable-scrollback + transient-tail (calm under load), **Tab-queues-while-busy** composer (never drop operator input), reserved space (`codex-notes.md:34-79`). Digest/batch/throttle vocabulary distinguished; the **A-gate is the non-suppressible `critical` class** (monotonic floor, cannot be silenced).

**(b) Meeting-collaboration surface (the under-served lane — DQ-4 resolved).** The export: route **design/open/grill** gates to the **meeting/attach** surface and **MUST NOT compress them to a brief**; "PROTECT the collaboration lane as deliberately as the governance one" (`:51`). Contract locked now:
- **Gate→lane routing (m-6-owned, off LOCKED atoms, no new field):** a gate-bearing record with `phase ∈ {DESIGN, DESIGN-REVIEW}` (the locked `phase` atom is the **primary key**) → **collaboration/meeting lane** (conversational, not compressed); otherwise (verdict-shaped) → **governance/ODB lane**. `GRILL_REQUIRED=yes` **corroborates** the meeting-lane route, but it is **not yet an m-2 registry field** — it is a ported upstream protocol header (agent-raised). **Owed m-2 dependency (c6-F6):** m-2 declares GRILL_REQUIRED's FieldSpec row (owner/type/values — the ported upstream header) before the routing binds it as a field; until then the route keys on the locked `phase` atom alone. (Seam-converged; reuses the conductor-derivation discipline on locked atoms.)
- **Re-observe-on-resume (export (E), `:49`):** a meeting / an "elaborate more" fork is an out-of-band channel ⇒ conductor state may be stale ⇒ **re-observe before the woken lane acts; newest authoritative** (the m-3 observe-gate re-fired on resume).
- **Attach mechanism DEFERRED to Step-3/4 runtime** (the live-session attach + the fork primitive are runtime-owned); this doc locks the **routing + no-compress + re-observe contract**, not the attach plumbing.

---

## 6. Interjection host (Seam B; Claude-Code three-mechanism)

m-6 owns the **surface**: the inline composer + the **steer / side-question / interrupt** choice + the rendering of the side-question answer on a **separate, non-lane surface** (the lane is NOT interrupted; the answer never injects into the lane transcript — the explicit fix vs jcode's deferred side-panel, `jcode-ux-notes.md:37-77`). m-5 owns the **sensor archetype**; m-4 **routes** it `fast-cheap` (locked); the **runtime** owns boundary-injection / soft-cancel / fork.

**Affordance binding (seam-converged `123022`⊕`131856`, off m-5's `accepts_interjection`-by-longevity):** the composer offers `{steer, interrupt, side_question_target}` **scoped to the target archetype's declared capability** — full 3-way for long-lived seats (`implementer`/`solo_worker`/`planner`/`orchestrator_lead`); side-question+interrupt for `reviewer`; the 1-turn **`sensor` is not a steerable target**; the **`actuator` `interrupt`-before-commit cell is a named placeholder** pending m-5's grill (declare-before-bind). steer = queue-to-next-safe-boundary (inline, no panel); side-question = read-only tool-blocked 1-turn fork (answer = `surface_intent=advisory`, the Seam-A↔B tie, content `self_reported`, never gate-bearing, routed `fast_cheap`); interrupt = soft-cancel + redeliver. The **API floor** (no true mid-generation injection) is a documented constraint, not a gap (`jcode-ux-notes.md:23-29`).

---

## 7. The m-5 seam resolution (cite COORD converged state)

Resolved in `c3-design-m5-m6-coord` — **m-6 bind `123022` ⊕ m-5 confirm `131856`** (the converged model; the interim `125604`/`131747` `{verdict,fyi,collaborate}` excursion crossed `123022` and was retracted both sides — m-5.implementer must-revise `131617` + m-5 SITREP `132314`). **m-5 DECLARES; m-6 BINDS** (declare-before-bind, F2). The two-layer vocabulary: **`human_mode` posture {interactive, away, unattended}** (on the locked F2 per-assignment home, `ARCHITECTURE.md §C2.4/F2`) ⊥ **`surface_intent` {progress, review_checkpoint, advisory, result}** (**conductor-derived**, total, **non-gate** records only — like `record_integrity`; **no new m-2 field**, C2.4 holds). Gate-bearing records carry **no** `surface_intent` — they bind off the **locked** `gate_category`/HUMAN_GATE/J1 (no duplicate-mechanism-as-vocabulary).

**The m-6 bind matrix (over (posture × surface_intent) + the locked gate fields):**
| `surface_intent` | m-6 binding | lane |
|---|---|---|
| `progress` | ambient/ephemeral; low retention; batchable; no park | governance(ambient) |
| `review_checkpoint` | route to meeting/collaboration; **no-compress**; re-observe-on-resume | collaboration |
| `advisory` | separate non-lane side-question surface (§6); never gate-bearing | governance(advisory) |
| `result` | terminal; retained + navigable; may feed a verdict if also gate-bearing | governance(result) |
| *(gate-bearing)* | off LOCKED `gate_category`/HUMAN_GATE → A-bucket ODB + park + J1; lane via §5 routing | governance / collaboration |

Posture binds the channel/cadence/away-eligibility intensity. `away_bridge_eligible` = m-6-owned boolean (§4). **Consumer requirement on the derivation: it must be TOTAL** (`progress ⇐ otherwise` satisfies it).

---

## 8. Consumer boundary contract (named consumed fields)

m-6 **consumes** (reads, never writes): m-1 — the addressing graph + the operator-FROM stamp/channel (`m-1 …design.md:107-118`); §J defaults (J1/J2). m-2 — `gate_category` (A/B), `HUMAN_GATE_REQUIRED` (monotonic floor), `delivery_state`/`failing_edge`, the ODB field slots (`:250-276`). m-3 — `observe_result` veto + `egress_scan_result` + `evidence_ref`/`achieved_evidence` + `record_integrity` (`:41-63,107-122,146-151`). m-4 — the routing escalation (`human_decision_required`/`routing_unavailable` → the routing-raise, `:333-350`). m-5 — `human_mode` + `surface_intent` (derived) + `accepts_interjection` + the sensor archetype (the seam, §7).
m-6 **writes** (its own surface/scheduler artifacts): the bucket projection/config, the ODB render, the park/wake state, the away-bridge policy + the token-bridge trigger/render/POST-receipt (which **calls** m-1-owned `mint`/`verify` — Seam C = A), the meeting-routing, the interjection composer.
**Downstream:** none (terminal design domain; m-7..m-12 runtime/product are later). **No-consumer flag:** every bucket/surface in this doc has a locked writer behind it (§2 table); the away-mode token bridge's upstream is now **settled** — m-1 (TCB) owns `mint`/`verify` (Seam C = A, §4/§10) and m-6 owns the bridge that calls it. **No binding is left without a settled upstream.**

---

## 9. GRILL_LOCK (operator decisions, 2026-06-30)

Operator-grilled (HUMAN_GATE_REQUIRED:yes); folded here; all operator-overridable §J config except where noted.
- **G1 — away-mode destination:** A-gates mirror **live** + B-items mirror as a **frequent digest** (operator-config interval, **default sub-daily**); **C stays local**; opt-in; egress-gated. *(Operator chose more remote visibility than the A-only default.)*
- **G2 — reply grammar:** bounded choices as **buttons** + optional **"elaborate more"** → a **context-preserving read-only fork** of the parked lane (gate stays parked; the fork's only output is the **decision relayed back** to the original session — not a work-stream to merge; write-capable fork **declined**; verdict still bounded-choice → operator-FROM; re-observe-on-resume). Free-text note rides *with* a choice; never a hidden choice / gate-lowering.
- **G3 — Bucket-B visibility:** **live local digest, non-interrupting** (audit autonomy without noise).
- **G4 — resummon cadence:** **escalate the channel, never auto-resolve** (J1); no hard deadline by default; optional per-gate conservative block-ceiling.
- **G5 (recommend-and-default):** delegation **DEFERRED** (solo-operator); meeting attach-mechanism **deferred to Step-3/4**; D stays author-facing in the taxonomy; the §J map/protected-branch/away-opt-in/egress-whitelist **config surface** is m-6's, the **values** are the operator's.

---

## 10. Open questions (no remaining blockers — OQ-1 resolved; OQ-2/OQ-3 non-blocking)

- **OQ-1 [RESOLVED — Seam C = A] — the m-1 confirm-or-gap (away-mode inbound verdict-token).** m-1 answered **(A): m-1 (TCB) owns mint/verify; m-6 owns the bridge over it** (`COORD-planner-20260630-182600`; orchestrator-accepted `183008`). *Forced* by DI-1 + DI-2; additive — first activation of m-1's reserved `certification` seam; **no c1 reopen**. The away-token cell is folded in §4 and now **locks**. The A-boundary + the build carries (**four** from `COORD-182600` + the **c5 decision-④** carry) are recorded in §4/§12.
- **OQ-2 [m-5 seam follow-up, non-blocking] — the "elaborate more" fork posture.** Does the context-preserving elaboration fork inherit the parent `seat_archetype`/`human_mode`, or run under an `interactive`/`review_checkpoint` collaboration posture (read-only ceiling)? Bounded COORD note to m-5; does not block the m-6 lock (the fork primitive is runtime-owned; the surface is m-6's).
- **OQ-3 [operator-declined / reserved] — write-capable elaboration fork is DECLINED** (operator 2026-06-30: the elaboration fork stays read-only; its only output is the decision relayed back). A hard per-archetype **never-bridge ceiling** (beyond content-egress) remains a reserved future hook — not needed for the initial release.

---

## 11. Local-first / egress fail-closed compliance + what locks now

**PASS.** Every Step-2 mechanism (§2 buckets, §3 ODB, §4 park/wake, §5 email surface, §6 interjection) is **local-only** over the append-only relay store. The **only conductor-governed external send** is the §4 away bridge (governance-surface claim, not system-level sole-egress — D5 residual: same-uid shell/curl bypass out of scope) — **opt-in**, **egress-gated** by m-3 §7 (the dormant fail-closed chokepoint activates on exactly this first external send). No surface needs external send before the egress gate exists.

**Locks now — no held cell remains:** the bucket taxonomy (§2), the ODB render + capture incl. the elaborate-more fork contract (§3), the full 7-state park/wake machine + resummon cadence **including the away-mode token-bridge cell over m-1-owned mint/verify (Seam C = A, §4/§10)** (§4), the email+meeting surface contract + gate→lane routing + re-observe-on-resume (§5), the interjection host (§6), the m-5 bind matrix (§7). Seam C = A is folded (`182600`/`183008`/`183345`); the mint/verify **build** is an additive later-step PLAN carry (§12), not a design blocker.

## 12. PLAN carry-forwards (for the build cycle — NOT this phase)
- ODB render + the buttons/elaborate-more capture as a Step-2 local mechanism; the operator-FROM verdict relay write-path.
- The park/wake state machine as durable-store transitions (TOCTOU on wake; the m-3 re-observe-on-wake binding).
- The away-bridge egress activation (m-3 rule-set); the signed-token bridge **build** (Seam C = A: mint/verify is m-1-owned per the §12 inbound-token carries — OQ-1 answered, not pending); the digest cadence scheduler.
- The fork primitive (runtime-owned) for elaborate-more + side-question; the meeting attach mechanism (Step-3/4).
- The §J config surface (gate_category map / protected-branch / away opt-in / egress whitelist).
- The Step-4 TUI realization (codex stack; notmuch-style tag/saved-search buckets; mutt additive-score sort).

**Inbound verdict-token mint/verify (Seam C = A; additive later-step build carries — the **four** from `COORD-planner-20260630-182600` below, + the **c5 decision-④** carry as the final bullet):**
- **m-1:** signing-key custody OS-isolated from all lanes (DI-2); the key never resides in an m-6 / any lane process.
- **m-1:** nonce-burn = an atomic conductor-owned store append (verify's nonce-unused → burn closes the same TOCTOU as `submit()`).
- **m-1:** `certification`-seam activation scoped **inbound-verdict-only** for this step; general DKIM/Merkle courier-hardening stays the deferred seam (do not widen here).
- **m-6:** POST-not-GET is m-6's transport invariant (anti-scanner); m-1's `verify` stays fail-closed / method-agnostic and rejects any auto-GET-derived replay. m-6 supplies `expiry` (validity-window policy); m-1 enforces it.
- **m-6 resummon edge (operator decision ④, ROTATE + RE-OBSERVE — recorded §2C build-carry, `c5-fold-decision-4`; NOT design-locked here):** a resummon triggers **rotate `decision_id` + burn prior-cycle nonces** (m-1-owned burn) and a **`verify` re-observe that bounces the approval if state changed** since the operator last saw it — closing the stale-approval / TOCTOU window. The 7-state FSM's refresh transition carries it. **Mechanism + fixture + full-pair adversarial review are owed at build-step (d)** (before park/wake or the away-bridge ships), not design-locked here; this **subsumes/pairs with** the m-7 `re-mint-supersedes` carry + my CQ-6 co-sign (`c4-cq-m1/DESIGN-planner-20260702-020100`). **Residual (dormant in Step-1):** until (d) builds it, the base `(decision_id, seat)` sibling-burn + the never-auto-resolve-on-expiry FSM (§4) hold.

---

## 13. Fold-log (post-c3-lock, review-driven consistency folds — lock invariants unchanged)

The c3 design-lock (`c3-lock` VP co-sign `20260630-191315`) is the frozen contract. The folds below are **claim-text / consistency** consequences of downstream cycles applied INTO the locked doc; **no mechanism, no lock invariant changed.**
- **c4-cq-gateconfig** (pair-approved `024620`): CQ-3 pure-judgment A-floor table authored (canonical home `c4-cq-gateconfig/DESIGN-planner-20260702-015800`; NF-S8 binds there); CQ-4 terminal-token alignment (`delivery_state {accepted, rejected, held}`; `bounced`→`rejected`; `held`→bucket-A fault-disposition, ordinary A-gate parking = `accepted`) folded §2/§4; CQ-4b config-composition confirm.
- **c4-cq-m1** (pair-approved): CQ-6 away-token park/wake co-sign + the re-mint-supersedes carry (§12).
- **c5-claim-sweep-light** (approve `134742`): 3 relabels (`sole-writer`/`sole-external-send`→governed-write/conductor-governed-egress + D5 residual), 1 DI-1 mirror, 1 "structural"→bucket-grammar tighten (§2/§4/§11); claim-text only, four sanctioned by-construction claims + `{accepted, rejected, held}` preserved.
- **c5-fold-decision-3** (approve `134743`): decision ③ RAISE-ONLY classification-direction invariant + known-A detector (§2). *(m-6-F1: the m-7 NF fixture is "registration owed", CTO-applied §2.)*
- **c5-fold-decision-4** (approve `134744`): decision ④ ROTATE+RE-OBSERVE recorded as a non-locking §2C/§12 build-carry (mechanism at step d).
- **c5-fold-decision-5** (approve `134745`): decision ⑤ ODB model-name egress carve-out (§3); R2 untouched.
- **c6-fix-m-6** (this pass): **F3** — decision-⑤ m-7 note reworded to "fixture registration owed" (§3). **F6** — meeting-lane routing keys on the locked `phase` atom; `GRILL_REQUIRED` corroborates but is an owed m-2 registry declaration (§5). **F7** — STATUS + domain README updated to c3 LOCKED. **F8** — ARCHITECTURE line-anchors converted to stable section anchors (`§J1`, `§J2 A-set`/`B-set`/`fail-safe`, `§4 R1`, `§C2.4/F2`) to survive re-baseline line drift. **F9** — Seam-C carry counts split (four `COORD-182600` + one c5 decision-④) at §4/§10/§12. **F1/F2** were CTO-applied pre-dispatch (§2:44 ③-fixture-owed, §2/§4 `held`-gloss narrowing) — verified, not redone.
