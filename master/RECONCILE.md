# master — RECONCILE ledger

Orchestrator reconciliation of paired audits/designs. Incoming reports are E0 until reconciled here against
file:line evidence. Append per cycle.

> **Claim-boundary / supersession note (applies to all pre-c5 entries below).** Strength vocabulary in the c1–c3
> historical sections — "forgery-robust by construction", "sole-writer", "closes all three by construction",
> "confusion-robust → forgery-robust" — predates the fork/relabel decision and now reads **per the c5 relabel**:
> the honest Step-1 claim is **confusion-resistant** (GRILL-LOCK D4; D5 accepted direct-store residual), with the
> four licensed by-construction classes enumerated in `ARCHITECTURE.md` §C4.3. The historical entries are preserved
> verbatim (append-only discipline); this banner supersedes their strength wording. See RECONCILE c5 / D4.

---

## Cycle c1 — Step-1 foundations AUDIT reconciliation (2026-06-28)

Sources: `c1-audit-m-1` (planner 152628 + implementer 160119) · `c1-audit-m-2` (planner 152808 + implementer
160122). All four E1-cited (file:line). **Verdict: both domains still-open over a large promote-and-formalize
base; strong cross-pair convergence; both PROCEED-TO-DESIGN.**

### m-1 — Trust & Identity

**Converged (both seats, E1):**
- still-open: the sole-writer stamping courier exists in neither the stock protocol nor jcode nor claude-code — all three
  ship *self-asserted* identity.
- The three identity gaps, identically located: stock-protocol agent-authored `FROM` (lint checks only ROLE/FROM
  self-consistency); jcode `assign_role` + `from_session` in the wire body; claude-code self-written `from`
  guarded by `msg.from==='team-lead'`. Channel-stamped FROM closes all three **by construction** (the sender
  can no longer write the trusted field).
- Minimal API: **`submit` / `project` / `read` / `mint_seat`** (conductor-only mint; `submit` the sole write path).
- Two invariants → forgery-robust by construction: **I1 sole-writer/store-isolation + I2 channel-isolation**.
  Merkle/tamper-evidence **deferred** (the courier is the TCB).
- Reuse-not-rebuild: upstream store shape + lineage checks; claude-code inbox+lock projection/poll. On-disk
  markdown becomes a *projection of a stamped record*, not a lane-authored file.
- Keep **one FROM** (no payload display-FROM — the DMARC two-FROM trap).
- Boundary contract to m-2/m-3/m-4/m-6 mapped + consumer-validated (no orphan primitive).

**Resolved nuances:**
- jcode credit: planner = "promote the connection→session binding jcode already has and *discards*
  (`ClientConnectionInfo`)"; implementer = "jcode is transport prior-art, NOT a near-miss identity primitive
  (body still carries `from_session`; role mutable via `assign_role`)." → **Promote the binding (the datum),
  not jcode's identity model.** Compatible.
- `attest()`: no public `attest()` call — fold attestation into `mint_seat`/connection-setup; `submit()` fails
  on an unbound connection. (Implementer; planner had raised it as open.)

**Open (DESIGN):** identity ≠ authority boundary (the stamp says *who*; m-4/m-5 policy says *what*) — confirm
m-1 owns identity + store only.

**OPERATOR-DECISION — Step-1 transport strength** (the one consequential call): operator-attested manual-relay
(confusion-resistant, *not* by-construction) vs minted-token-over-isolated-per-seat-connection (by-construction)
vs OS-peer/mTLS/SPIFFE (standalone). Does **not** block DESIGN; decides whether the initial release *claims* "forgery-robust by
construction" or "operator-attested." Both seats independently flagged it.

### m-2 — Forms & Determinism

**Converged (both seats, E1):**
- The frank declarative schema artifact is net-new (still-open) over a **large promote-and-formalize base**:
  the upstream ~50-field catalog + enums + required-when + lineage already exist (scattered across protocol.md prose
  + relay-lint.py). Promote, don't rebuild.
- Net-new = the **field-ownership model** (owner ∈ {system, seat-scoped-enum, agent-enum-pick, free-text}) +
  **fill-time authority** (a forbidden option is *absent* from the seat's form, not rejected post-hoc) + **one
  canonical schema source** the tool, courier, and linter all read.
- Dissolve-vs-survive: ~32 prose/markdown/token-lexical checks **dissolve** (a typed envelope has no
  fences/rows/bare-tokens to disambiguate); enum/required/consistency **survives** as fill-time form-validation;
  the cross-relay lineage walk **survives** as a separate engine.
- GATE-1 (a typed envelope alone cannot preserve): (a) cross-relay lineage = separate engine, **strengthened to
  forgery-robust by m-1's system-filled PARENT**; (b) `DISPATCH IMPL`/`DISPATCH MERGE` authority — the lexical
  token dissolves but the authority survives as a **seat-scoped + phase-scoped form field**.
- Keep required-when a **bounded** declarative predicate (PHASE × CEREMONY_TIER × field-values), never
  Turing-complete (CodeAct/determinism). Both name this the top risk.
- `FROM` = owner:system (never lane-supplied); bodies non-executable; **canonical-iff-consumed** (a field with no
  mechanical consumer → free-text/flag).
- Boundary contract to m-1/m-3/m-4/m-6 mapped (incl. porting agent-scripts' **Owner Decision Brief** for m-6).

**Resolved nuance:** verdict label — planner still-open (+large already-closed) vs implementer recommended-next
(promote) → same substance (net-new artifact over a promote base).

**Open (DESIGN):**
- Carrier: JSON-Schema core + `x-owner`/`x-seat-scope`/`x-consumer` extensions (planner) vs bespoke FieldSpec data
  (implementer; claude-code uses Zod, jcode typed Rust records).
- Sanctioned **overflow channel**: one bounded, never-gate-input free-text escape hatch (so unanticipated messages
  aren't blocked and the body doesn't silently re-grow prose).
- **Format-Tax** constraint: keep reasoning-bearing free-text genuinely free (don't force the model to reason
  inside a rigid format).
- Schema versioning/evolution + parked-lane compat (cross-cutting m-1 store + m-6 scheduler).
- Storage backing (JSON vs SQLite vs both; markdown view-only vs signed export; legacy-markdown strictness).

### Cross-cutting (m-1 ↔ m-2)
- **Co-foundational — lock together.** m-2's schema *declares* the system-owned fields (FROM, PARENT, evidence);
  m-1's stamper/courier + m-3's probe *fill* them. Both pairs independently reached "m-1 store + m-2 schema lock
  first; m-3/m-4/m-6 design against the locked schema."
- **System-filled PARENT (m-1 courier) strengthens the m-2 lineage engine from confusion-robust → forgery-robust.**
  Both pairs named this convergence — a strong signal the foundations cohere.

### Process notes
- `m-2.implementer` relay used DISPATCH_ID `c1-audit-m-2-implementer` (others used the thread id `c1-audit-m-2`)
  and addressed `TO: master.orchestrator-planner` instead of `TO: m-2.planner` (the reconcile partner), planner on
  CC. Substance intact; minor addressing/threading deviation — flagged for the VP and the m-2 reconcile.

### Disposition
Both domains → **PROCEED-TO-DESIGN** (still-open, medium tier, strong convergence). DESIGN dispatch to
`m-1.planner` + `m-2.planner` (co-foundational), `GRILL_REQUIRED: yes` for the operator-owned + cross-domain
questions above. Consumer lenses m-3/m-4/m-6 review the interface sketches **before** design-lock.

---

## Cycle c1 — CLOSED / LOCKED (20260629-181713)

**Outcome.** m-1 Trust & Identity + m-2 Forms & Determinism jointly locked as the frank Step-1 design-of-record.
VP `approve` close-confirm (`c1-joint-lock`, 20260629-180934); operator §J ratified; close declared
(`c1-joint-lock`, 20260629-181713). Authoritative detail in the two rev2 domain design docs; integration spine
in `ARCHITECTURE.md` §1–§5 + §J.

**Lifecycle run.** boot → audit (m-1, m-2) → audit-reconcile → design (co-foundational, grilled) →
consumer-review (m-3/m-4/m-6) → consumer-reconcile → refine (rev2, both pairs re-affirmed the contract) →
joint-lock → VP close-confirm → operator §J ratification → CLOSE. VP co-review at every authority-bearing gate.

**Contract closed in both directions.** R1 operator/special-address; R2 routing = separate seat-stamped relay
(model never a gate input); R3 observe-integrity (`evidence_integrity`, conductor-read under DI-5). PARENT
(m-1, system-filled) strengthens the m-2 lineage engine confusion-robust → forgery-robust. identity ≠ authority
ratified. §J: J1 `hold_and_resummon` (never auto-approve); J2 `gate_category` default set, operator-configurable
(forward), merge split by target branch + protected-branch set.

**Authority boundary.** No PROCEED-TO-PLAN / implementation / merge / live-verify granted. Consuming domains
design against the locked contract in later cycles. Forward requirements + PLAN carry-forwards recorded in
`ARCHITECTURE.md` §6 + §J — none reopens the lock. Pairs stood down; next cycle is the operator's call per
`ROADMAP.md`.

---

## Cycle c2 — CLOSED / LOCKED (20260630-044308)

**Outcome.** m-3 Observation & Evidence + m-4 Routing & Policy jointly locked as the frank **Step-1
runtime-intelligence layer** atop the locked c1 substrate. VP co-sign `c2-lock` 20260630-043859; close
declaration 044308. Detail in the two rev2 domain design docs; integration spine `ARCHITECTURE.md` §C2.1–C2.8.

**Lifecycle run.** decomp (VP) → audit (m-3, m-4) → audit-reconcile (VP) → design (co-design, grilled) +
the m-3↔m-4 COORD seam (reconciled both sides) → lock-prep (VP: R2 ratified, m-5 narrow engagement approved,
GL-4 roadmap-checked) → consumer-lens round (m-5 narrow + m-6) → consumer-reconcile (VP) → fold-confirm round
(F1/F3→m-3, F2/F3/M4-1→m-4, implementer-re-approved) → lock (VP revise on one lock-text ambiguity → de-lock →
VP co-sign) → CLOSE. VP at every authority-bearing gate.

**Key locks.** The R2-preserving seam (silent-deviation block via m-3's generic integrity-veto; no model-derived
predicate in any gate; bucket-vs-bucket; snapshot-provenance). The two-layer capability prior (GL-1; R2
structural at the deviation comparison). Two opaque archetype atoms — `slot_in` work-archetype (conductor-classified
at acceptance, non-lane-writable) + `seat_archetype` per-seat (per-assignment record home) — concrete semantics
reserved to m-5/c3. M4-1 via the c1 monotonic HUMAN_GATE routing-raise. GL-4 routing-templates (record mechanism;
pane-spawn rides existing multiplexer infra). Interjection forward-req folded into `ROADMAP.md`.

**Operator-judgment / process notes.** GL-4 (initial-release routing-templates) + the m-5 narrow engagement were
operator-directed and cited as **operator-directed by current session context** (VP-sanctioned; no `FROM:operator`
relay). One lock-text de-lock to m-3's §5.1 was **orchestrator-applied under operator direction** (narrow,
self-attributed, non-substantive) — VP-accepted for this closure only; **not a precedent** (substantive changes
to a domain's design semantics still go through the owning pair's relay/review path).

**Authority boundary.** No PROCEED-TO-PLAN / implementation / merge / live-verify granted. m-5's full
archetype-system design + m-6's full human-surface design are **c3**; the c2 narrow engagements reserve m-5's
full ownership. PLAN carry-forwards in `ARCHITECTURE.md` §C2.8 — none reopens design. Pairs stood down (m-5 holds
its c3 reservation); next cycle is the operator's call per `ROADMAP.md`.

---

## Cycle c3 — CLOSED / LOCKED (20260630-191525)  ✦ Step-0 design-of-record COMPLETE

**Outcome.** m-5 Workflows & Archetypes + m-6 Human Surface & Scheduler jointly locked as the **final two Step-0
design domains** — **completing the six-domain Step-0 design-of-record** (c1 foundations + c2 runtime-intelligence +
c3 human-surface). VP co-signed **both the lock and the C3.6 integration capstone** (`c3-lock` 20260630-191315);
close declaration 191525. Detail in the two design docs; integration spine `ARCHITECTURE.md` §C3.1–C3.7.

**Lifecycle run.** decomp (VP) → audit (m-5 focused, m-6 full) → audit-reconcile (VP; **F4 enforced** — m-5 owed +
filed its pair-reconcile) → grilled co-design + the m-5↔m-6 COORD seam (3 items: human-mode declare-before-bind /
interjection host / the m-1 confirm-or-gap) → design-complete (both pair-approved) → **Seam C: the first
conditional-upstream-contract-check** routed to compacted m-1 (answered A) → lock (VP revise on one stale-status
contradiction → m-6 cleared it + pair-re-approved → co-sign of lock + capstone) → CLOSE. VP at every
authority-bearing gate.

**Key locks.** The **archetype = one governed expansion-slot** binding {topology + gate-set + ceiling-at-spawn +
observe-invariants + routing-prior} (m-5); the **promote-and-bind human surface** (m-6); the **seam-of-record**
(posture × `surface_intent`, four-class, conductor-derived, non-gate — **no m-2 micro-fold**); **Seam C = A** (m-1
owns the inbound-token mint/verify via its reserved `certification` seam — **additive, no c1 reopen**); both
GRILL_LOCKs; the **C3.6 capstone** (the six domains compose — writer-backed + acyclic; three seams closed; locked
m-1..m-4 invariants intact; deferrals recorded as build-carries).

**Process notes.** (1) m-5 filed two independent audit passes without a pair-reconcile; the VP's F4 bar required it,
so the audit-reconcile **held** until m-5 filed it (two convergent reconciles) — the recurring .codex-lane gap,
enforced not waved. (2) The m-5↔m-6 seam had a mid-design **crossed excursion** (m-5's `125604` crossed m-6's bind
`123022`) — caught by m-5.implementer, resolved by m-5 conforming to m-6's bound model + retracting; the seam-of-record
is `123022`⊕`131856`. (3) The first **conditional-upstream-contract-check** (VP F3) re-engaged compacted m-1 for one
bounded question; m-1 answered **A** and proved it *forced* by DI-1 (nonce-burn needs the sole-writer store) + DI-2
(signing key = TCB secret), and **additive** via the reserved `certification` seam. (4) The VP's lock revise was a
**narrow stale-status contradiction** in m-6's doc (locked-vs-held); m-6 fixed it (full-doc sweep) + m-6.implementer
re-approved — **no orchestrator proxy-edit** (ownership held).

**Authority boundary.** No PROCEED-TO-PLAN / implementation / merge / live-verify granted. The **C3.7 build carries**
(incl. the **scoped** Seam-C-A m-1 inbound-token mint/verify — not to be widened) inherit to the future build cycle
only. **Step-0 is complete;** the PLAN phase / Step-1 conductor-core build is a **separate operator-opened gate** per
`ROADMAP.md`. Pairs stood down; next cycle is the operator's call.

---

## Cycle c4 — Conductor-Core substrate AUDIT reconciliation (2026-07-01)  ✦ re-baseline step (b)

Sources: `c4-audit-m-7` — `AUDIT-planner-20260701-160243` + `AUDIT-implementer-20260701-155145` (independent paired
audit) → `RECONCILE-planner-…-161306` + `RECONCILE-implementer-…-161137` (pair-reconcile) → merge artifact
`master/domains/m-7-conductor-core/audit/2026-07-01-audit-pair-reconcile.md`. Full findings:
`master/domains/m-7-conductor-core/audit/2026-07-01-substrate-audit.md`. **Verdict: PRIMARY_BUCKET still-open;
PROCEED-TO-DESIGN. VP co-signed `c4-audit-m-7/RECONCILE-orchestrator-reviewer-20260701-162319` (APPROVE-TO-DESIGN-START).**

### Convergence (both seats, independent — residual divergence: none)
- **still-open** — the entire §2A substrate is net-new; **no existing conductor to promote** (duplicate/already-built
  gate run both sides; the stock protocol = "coordination protocol, not orchestration runtime"). Donor PARTS promoted inside the
  net-new engine (not rebuilt): upstream protocol/lint layer; jcode tmp→fsync→rename write discipline + connection
  binding; claude-code mailbox lockfile + re-read; codex single-owner serialized writer + ToolExposure + config-lock;
  external SQLite super-journal / `rename(2)` / `fsync(2)` / Maildir.
- **Claim boundary held (verbatim both sides):** "by construction" licensed **only** for the serialized-loop
  two-honest-seats double-accept kill (a control-flow property of the trusted engine); **no** adversarial sole-writer /
  unbypassable / same-uid write-exclusion claim; an F8 claim-sweep fixture required. The exact line the NO-GO turned on —
  held independently. Interface guardrail = conductor-as-MCP presenting exactly `{submit, project, read}`, raw
  store/config/outbox paths absent, **confusion-resistant only**. (Only 4-bucket vocabulary differed; zero substantive
  delta — certified different-coverage-not-conflict.)

### Canonical audit record (reconciled to E1)
- **18-row seam matrix (S1–S18)** `{contract owner · doc/section · m-7 execution obligation · positive/negative fixture ·
  CQ?}` + merged fixture set (F1–F8 + guardrail negatives G(i–iii) + live-store E2 probes) — merge artifact §2–§3.
- **Self-referential finding:** the LIVE `master/` store has no substrate — 261 INDEX rows with out-of-order + duplicate
  timestamps from unserialized independent seat appends, no lock/tmp, seats hold raw store paths. Our own store is the
  counter-example m-7 must fix.

### Unified CQ list — design-LOCK gate (owners; NONE block design-START; all owned by STOOD-DOWN pairs)
- CQ-1 phase-split required-set → m-1 + m-2
- **CQ-2 decision-② fail-closed → m-3 fold (= re-baseline step (c)) + m-2 field-home** — the "decision recorded but never
  folded" pattern (fail-open still at m-3 §3.2:63); foregrounded
- CQ-3 pure-judgment A-floor table → m-2 + m-6
- **CQ-4 terminal-state token set → m-2 + m-3 + m-6** (Q-E `bounced`→`rejected` + new HELD); foregrounded
- CQ-4b trusted-config composition → **CTO arbitrates** (VP-confirmed; m-6/m-3/m-4 supply inputs, m-7 supplies
  load/integrity, CTO/VP own the composition contract + final load boundary)
- CQ-5 `slot_in` classification ordering → m-3 + m-5
- CQ-6 seat-binding + sibling-burn/restart → m-1 (+ m-6)
- CQ-7 observe row-parity → m-2 (**non-locking**; pre-Step-1-PLAN SHOULD)

### Disposition + forward plan (VP co-signed `…-162319`)
- **PROCEED-TO-DESIGN:** open `c4-design-m-7` **now** (design-START unblocked), GRILL_REQUIRED: yes,
  seam-matrix-before-lock, **hard no-lock gate on CQ-1..CQ-6 + CQ-4b**.
- **CQ closure = targeted-parallel COORDs** (VP Q2) re-engaging the specific stood-down owner
  (conditional-upstream-contract-check), NOT one serial mini-cycle; the design dispatch carries ONE CQ ledger updated as
  COORDs close. CQ-3/CQ-4/CQ-4b MAY group into one focused m-2/m-3/m-6/CTO reconcile IF they converge on the same
  gate/config surface — **not** a broad c1/c2/c3 re-open.
- **CTO action items:** (i) arbitrate CQ-4b config composition before lock; (ii) sequence re-baseline step (c) so
  decision-② (CQ-2) lands in locked m-3 before m-7 design-lock — the m-7 cycle and the re-baseline claim-sweep/fold
  converge here.
- **VP carry-forwards (lock-hygiene):** (1) patch/supersede the merge artifact's stale "implementer confirmation
  pending" line (`…-audit-pair-reconcile.md:4`) before design-lock — superseded now by `RECONCILE-implementer-…-161137`;
  (2) the design dispatch carries the explicit CQ ledger (owner · start/lock-blocking · expected closure artifact).

### Authority boundary
No DESIGN lock, PLAN, IMPL, code/`pcode/`, spike, or CQ resolution granted. No stood-down pair re-engaged by the reconcile
itself (the COORD plan follows). Step-1 PLAN remains a separate operator-opened gate after the conductor-core
design-of-record locks (re-baseline step (e)).

---

## Cycle c4 — CLOSED / LOCKED (2026-07-02)  ✦ re-baseline step (b) COMPLETE — the conductor-core substrate is designed

**Outcome.** The **m-7 Conductor-Core design-of-record is LOCKED** — the runtime substrate the six policy domains ride
on, whose absence produced the 2026-07-01 NO-GO. `DESIGN_LOCK_ID c4-design-m-7-lock` (design doc §22); **VP co-sign
`c4-design-m-7/RECONCILE-orchestrator-reviewer-20260702-040327` (`VP_DESIGN_LOCK_CO_SIGN: approve`, no blocking finding,
5 checks passed).** Detail in `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md`;
integration spine `ARCHITECTURE.md` §C4.

**Lifecycle run.** standup (VP approve-to-boot) → boot (both seats) → audit (independent paired, still-open) →
audit-reconcile (VP co-sign, PROCEED-TO-DESIGN) → grilled design (GRILL_LOCK `c4-grill-m-7`; 2 adversarial revise
cycles → r3 approve) → **CQ gate** (8 design-LOCK CQs closed via 3 COORD clusters, 6 pairs re-engaged full-pair; the
gate/config cluster took an r1→r2 revise cycle; certified `c4-cq-coord/…-031533`, CQ-6 re-scoped to base `…-032227`,
VP-approved `…-032843`) → lock-package assembly (r4 CQ folds + r5 stale-wording purge → m-7.impl approve `…-035245`)
→ CTO routing + certification (`c4-design-m-7/…-040011`) → **VP design-lock co-sign** → CLOSE. VP at every
authority-bearing gate; the closure gate held (no CQ folded without planner + implementer + required co-sign).

**Key locks.** The **engine** (§1–§11): durable-FIFO single-thread commit loop + atomic clear-on-pop; Package-A
canonical-record `rename` pivot + derived projections (INDEX layout unchanged — CQ-8); phases 0–4 recovery; byte-exact
`{accepted, rejected, held}` + HELD internal-fault disposition (`bounced` retired); trusted config = per-domain
sections under one top-level digest + per-section stamps, load-once (CQ-4b); MCP `{submit, project, read}` interface
guardrail + schema-as-form + pipe wake; conductor-governed local outbox; genesis/GC; persisted seat-binding +
decision-scoped `(decision_id, seat)` sibling-burn (CQ-6 base). The **seam matrix** (§12, 18 rows, biting negatives) +
fixtures F1–F10/G (§13). The **claim boundary**: confusion-resistant (GL D4 verbatim); the *sole* licensed "by
construction" is the §2.4 serialized-loop double-accept kill; D5 accepted-risks restated; semantic claim-sweep clean.

**The eight CQ dispositions.** CQ-1 (a) step-gate; CQ-2 decision-② class-conditional fail-closed = `held`; CQ-3 m-6
A-floor table on the monotonic MAX; CQ-4 byte-exact terminal tokens; CQ-4b single-digest per-domain-section config;
CQ-5 slot_in post-gate/pre-observe/atomic-bind; **CQ-6 base only**; CQ-8 INDEX derived-authority (layout unchanged).

**Process notes.** (1) The VP caught four routing/certification errors before they propagated — m-4 omitted from the
CQ-4b owner set, the implementer/co-sign action authority left as CC-only (CC ≠ action authority), and my CQ-6
over-certification (base-approval ≠ add-on-approval). Each folded before firing/locking. (2) `re-mint-supersedes`
correctly re-scoped OUT of CQ-6's closure to a §2C away-bridge build-carry (m-1-confirmed-fit, adversarial review owed
at its build step; dormant in Step-1).

**Authority boundary.** No PROCEED-TO-PLAN / IMPL / code / `pcode/` / spike / build granted (VP co-sign scope: design
lock only). **Build-carries** (non-locking): `re-mint-supersedes` (§2C), CQ-7 row-parity, the operator-gated runtime
spikes (RUNTIME-RESEARCH §12). **Re-baseline status:** step (b) COMPLETE; (c) global claim-sweep + (d) §2C-at-build-step
remain (several items — CQ-2 decision-② fold, byte-exact tokens — already discharged in c4); **(e) Step-1 PLAN** is the
operator-opened gate. Pairs stood down; next is the operator's call.

---

## Cycle c5 — CLOSED (2026-07-02)  ✦ re-baseline step (c) COMPLETE — the docs tell the honest story end-to-end

**Outcome.** The **global claim-sweep + the remaining operator-decision folds** are complete and byte-consistent across
all seven design-of-record docs (m-1..m-6 + `ARCHITECTURE.md`). A **doc-text pass, not a mechanism change** (§2B): every
"by-construction / sole-writer / forgery-robust / unbypassable / tamper-resistant / non-lane-writable / sole-external-sender"
adversarial-strength claim is relabeled to the honest **confusion-resistant** line + its **D5 accepted-risk**; the genuine
trusted-engine invariants are kept. **VP close-approve `c5-decomp/RECONCILE-orchestrator-reviewer-20260702-143205`.**

**Lifecycle run.** decomp (VP revise → owner-pair review + decision-④ split → approve) → lane 1 CTO ARCHITECTURE sweep
(VP revise: under-grepped survivor set → r2 complete + hardened checklist → **VP-ratified exemplar**) → 6 lanes dispatched
(3 sweeps + 3 decision-folds, VP-reviewed) → owner-pair execution (3 must-revise→approve cycles) → CTO pieces (§J ③ + §C4
④ ledger) + byte-consistency re-verify → VP close-review (revise: stale live `ARCHITECTURE.md:305 delivery_state=bounced`
→ patched to `rejected` + complete re-run → **approve**). VP at every gate; the adversarial pass caught real defects at
each stage.

**The 6 lanes (all owner-fold + own classified survivor list + implementer semantic approve).**
- **Claim-sweep** — m-1 (37 candidate hits, the TCB), m-2 (18), light m-3/m-4/m-5/m-6 (2–9). Relabeled adversarial-strength
  → confusion-resistant + D5; **KEPT** the licensed class: R2 gate-grammar invariant, observer-selected control properties
  (the F1 archetype invariants), authority-ceilings, already-scoped "no tool" claims. The two-sided discipline — relabel the
  overclaims, don't under-claim the genuine control-flow/grammar strength.
- **Decision-folds** — **①** attach + confusion-resistant rides the m-1 sweep (identity conductor-owned; runtime identity
  fields never accepted as FROM). **② was c4** (CQ-2). **③ RAISE-ONLY** A/B + known-A detector (m-6 §2 + **CTO-ratified
  §J**). **⑤** ODB model-name egress carve-out (m-3 scan + m-6 ODB + m-4 R2-guard; confidentiality-scoped only, R2 untouched).
  **④ away-token ROTATE + RE-OBSERVE — RECORDED as a non-locking §2C build-carry** (m-1/m-6 + the integrated §C4 ledger);
  **mechanism / fixture / adversarial review = a step-(d) gate**, not folded/locked. **Routing-lane §2C (m-4/m-2), RECORDED
  non-locking:** R2 `gate_referenceable` per-column FieldSpec + `chosen_model`/single-family-bucket-proxy negative fixtures,
  and altitude-B per-row `justified_deviation`/`deviation_reason_code` grain — mechanism + fixtures = a **step-(d) gate**.
- **§2A.6 A-floor + §2A.7 decision-②** guardrails intact (c4 CQ-3/CQ-2).

**Byte-consistency (CTO cross-doc, VP-verified clean).** Zero unclassified raw overclaim survives; terminal-token vocabulary
byte-exact `{accepted, rejected, held}` across m-2/m-6/m-7/`ARCHITECTURE` with `bounced` **retired** as a value token
(surviving only as documented-retirement / the "bounce" verb / the m-6-local FSM label `bounced_repair`, terminal token
`rejected`). The stale `ARCHITECTURE.md:305` bucket-D `delivery_state=bounced` (a c3-era line, VP-caught) is patched.

**Process note.** The VP's adversarial review caught, across c4+c5: the m-4 CQ-4b owner gap, the CC-≠-authority addressing
miss, my CQ-6 over-certification, the ARCHITECTURE-sweep incompleteness (under-grepped survivors), and the stale
`bounced` token — each folded before it propagated. Two of the c5 misses were my own **sample-not-exhaustive** verification
error; the byte-consistency + survivor-list-matches-file discipline is the standing guard.

**Authority boundary.** No mechanism change (claim-text + recorded-decision folds only), no design-lock reopen, no PLAN /
IMPL / code / `pcode/` / spike granted. **Re-baseline status: (a) ✅ (b) ✅ (c) ✅.** **(d)** §2C-at-build-step (the
away-token / away-bridge mechanisms — decision ④ + m-7 `re-mint-supersedes` — **and the routing-lane R2 `gate_referenceable`-per-column
+ altitude-B per-row carries** — land at their build step) and **(e) Step-1
PLAN** (operator-opened) remain. Pairs stood down; next is the operator's call.

---

## Cycle c6 — adversarial RE-REVIEW of the design-of-record + doc-only cleanup (CLOSED 2026-07-02, VP co-sign `c6-close/RECONCILE-orchestrator-reviewer-20260702-213836`)

**Why.** Operator re-ran the adversarial pre-build review against the CURRENT (post-re-baseline) design-of-record — coarser grain (10 lanes: 7 domain + x1 FATAL-resolution / x2 claim-honesty / x3 seam-byte; Fable 5 on max, single adversarial verifier) with the locked boundary briefed in. **Verdict: CONDITIONAL-GO** — 90 confirmed findings, **0 FATAL**, 1 refuted; the 2026-07-01 NO-GO is **discharged at the structural level** (the m-7 substrate held, the serialized-loop double-accept kill survived, attach/confusion-resistant held). Root cause of the survivors: the c4 locks + c5 sweep were **scoped to the 7 design docs**, so retired vocabulary + un-folded decision-tails leaked at that boundary (`CLAUDE.md`, domain READMEs, dashboard, RECONCILE; the ③ / §2C / token-convergence tails). Full review-of-record: `DESIGN-REREVIEW-2026-07-02.md`.

**Fix (c6, doc-only).** 90 = **44 CTO single-hand** (gov surfaces + seam-token convergence + §2C ledger restore + charter; VP-approved apply half `c6-apply/…-204236`) + **45 pair-dispatched** (7 lanes 5/7/10/8/3/5/7, **7/7 pair-approved** — several via revise→approve) + **1 subsumed** (x3-F5 anchors, covered by m-1-F10 + m-7-F4 repoints). The 8 apply-declines re-routed to pairs (the apply agents refused to guess genuine design calls). **Seams converged:** CQ-2 `{self_reported, mixed}` across m-2/m-3/m-7 + m-6 `held`→bucket A / `rejected`→bucket D (the VP's two-axis blocking amendment, `c6-decomp/…-192059`); archetype **propose-vs-stamp** (m-2 per-column `system_only` ⊕ m-4 seam ⊕ m-5 registry); `deviated_observed` **GL-1 bucket-vs-bucket** folded into m-3 §9 to match m-4/ARCH; trusted-config **author-set → m-2/m-3/m-4/m-5/m-6** (§C4.1). Evidence: `c6-apply.diff` (18 files / 103 hunks / +433 / −177, clean).

**Owed carries enrolled in the §C4 ledger (nothing silently dropped).** step-(d) build-carries: ④ away-token/away-bridge, `re-mint-supersedes`, R2 negative fixtures (the `gate_referenceable` **attribute is now declared first-class** c6; fixtures deferred), altitude-B per-row, **m-5-F2** away-trigger posture-expressibility (CTO ruling → rides the away-bridge). Owed Step-1-build fixtures: **③** known-A/RAISE-ONLY direction-invariant NF, **⑤** ODB model-name egress. All are step-(d)/Step-1-**build** items — **not** Step-1-PLAN blockers.

**Process note.** The coarse re-review caught the c5 sweep's own **scope hole** (it certified "honest end-to-end" while the charter/dashboard/READMEs still carried retired claims). The VP then caught, on my CTO work: the **m-7 CQ-2 convergence miss** (NF-S7 + CQ-2 ledger keyed `self_reported`-only), the **corrupted diff artifact** (`diff --color` alias + self-inclusion), and the **close-record accounting** (headline 90 ≠ 52+38+4; stale +398/−133 diff stats) — each folded before it propagated. Standing lesson reinforced: when the check IS the deliverable, the record must be self-consistent and every hit inspected, not sampled.

**Authority boundary.** Doc-only; no mechanism change; no design-lock reopen (review-driven consistency folds recorded in each pair's fold-log, lock invariants unchanged); the four sanctioned by-construction claims + confusion-resistant/D5 vocabulary + byte-exact `{accepted, rejected, held}` preserved; `pcode/` untouched; no PLAN/IMPL/spike. **Re-baseline status: (a) ✅ (b) ✅ (c) ✅ + c6 re-review cleanup ✅.** **(e) Step-1 PLAN** (operator-opened) is the sole remaining gate; the (d) / owed-fixture carries land at their build step. Standing scope-guard extended: `CLAUDE.md`, domain READMEs, `master/README.md`, `RECONCILE.md` now sit inside the claim-sweep + byte-consistency net (the sweep's scope was the c5 bug). Pairs stood down; next is the operator's call.

### c6.1 — differential seam-hardening (CLOSED 2026-07-03, VP co-sign `c6-close/RECONCILE-orchestrator-reviewer-20260703-012327`)

**Why.** Operator ran a scoped pre-PLAN **differential** (6 lanes: 4 cross-pair seam + regression + residual; single adversarial verifier) over the c6-closed docs. **HOLD — 11 blocking / 0 refuted → 5 seam divergences:** 3 c6-*introduced* regressions the pairs' implementers had approved (m-3 `deviated_observed` mis-fold `chosen_bucket ≠ declared_bucket`; m-3 egress → terminal `held`; m-5 "every send observes") + 2 lagging/dropped mirrors (m-2 author-set, m-2 held-shape). The c6 close's "seams converged / nothing dropped" was **wrong on these** — presence-checks + VP-sampling pass while the *formulas/tokens* diverge. A **flag census** over the c6 relay bodies caught **3 more** dropped cross-domain flags (m-7 S11 author; m-2 `deviation_reason_code` value-set; §J2 `routing_unavailable` + GRILL_REQUIRED).

**Fix (c6.1, doc-only, convergence-to-locked-targets).** CTO drafted the corrections; the VP required **focused owner-confirms** before re-close (correctly — re-closing on CTO edits + sampling would repeat the failure at smaller scale). m-2/m-5/m-7 pair-confirmed first pass; **m-3 caught that my egress fix was incomplete** (I fixed the §3.3 row but left §3.2(c)/CQ-4 map still asserting egress → terminal `rejected`) and folded the completion in its own doc: **egress is never a terminal token — the outbound block is the non-terminal `egress_blocked` park + A local resummon** (converges m-6 §46/§50/§51 + m-7 NF-S9 + ARCHITECTURE :309-310). All 4 owner-confirmed + implementer-approved + CTO spot-verified. Evidence: `c61-fix.diff` (6 files / 15 hunks / +35/−17, clean).

**Process lesson (folded into playbook A.6 + Part D).** Seam convergence needs a **dedicated adversarial differential** — presence-checks and sampling pass while the formulas diverge; and an **owed-item-as-typed-record projection** (the m-7 `held`/burn pattern generalized) is the standing form of the flag-census, to build in Step-1. The differential → owner-confirm chain caught what **four** prior review layers missed (pair-review, orchestrator-sweep, VP-close-sampling, and my own c6.1 self-verify — the last missed the incomplete egress fix), **twice**.

**Re-baseline status: (a) ✅ (b) ✅ (c) ✅ + c6 ✅ + c6.1 seam-hardening ✅.** The design-of-record is now **verified-at-the-seams, not asserted**. **(e) Step-1 PLAN** (operator-opened) is the sole remaining gate; the owed-fixture / step-(d) carries land at their build step. Pairs stood down; next is the operator's call.

**c6.1a — differential-caught §J2 slip (CLOSED 2026-07-03, VP co-sign `c6-close/RECONCILE-orchestrator-reviewer-20260703-023723`).** A final tight differential over the c6.1 delta (5 lanes) returned the **4 seam lanes CLEAN** (the c6.1 corrections cohere) + **1 blocking**: my c6.1 §J2 dropped-flag edit mis-named the A-member `routing_unavailable` (the route_dispatch **outcome state**) where m-4 §7:363-369 prescribes a **distinct `routing_escalation`** member and states no §J2 change is required for correctness (force-A holds via `other`→A). **Reverted** §J2 to the locked 8-member A-set (converges to m-4); **recorded** the explicit `routing_escalation` member as an owed, correctness-safe cross-domain carry (§C4). `c61-fix.diff` regenerated to reconcile the evidence artifact (a VP-caught record-integrity repair). CTO revert-to-locked only — no pair edit, no pair re-confirm. **This was the 4th differential-caught defect at a CTO-edited cross-domain seam** — the standing implication (playbook A.6 / Part D): CTO cross-domain seam edits are the recurring weak point; the seam-differential + owed-item-as-typed-record-projection disciplines are the guardrails, to build into the conductor at Step-1. Re-baseline status unchanged: **(a) ✅ (b) ✅ (c) ✅ + c6 ✅ + c6.1 ✅ + c6.1a ✅** — the design-of-record is clean; **(e) Step-1 PLAN** is the operator-opened gate.

---

## Step-1 · Slice-1 (S1) — FIRST BUILD-CYCLE CLOSE reconciliation (CLOSED 2026-07-03; operator-ratified, VP confirmatory pass)

**The design→build transition is realized — the first `frank` code exists.** S1 = the thin end-to-end conductor relay (`mint→connect→submit→stamp→validate→lineage→append(crash-atomic)→project→deliver→gate-outbox`), built in `frank/` by the `s1` slice-team (its own orchestrator + one `s1-core` pair — its granularity call), guided by m-7, per the `s1-dispatch` charter (`master/relays/s1-dispatch/…-130634`, VP-approved + operator-ratified). **Close baseline: `frank` main@f0dcb85 (code) · ledger main@674c844 · annotated tag `s1-close`→f0dcb85 · tree clean.** Verdict for the S1 scope: **complete at E2** (E3/E4 deploy/live-verify explicitly out-of-S1 per the honesty framing — stated, not skipped).

**Master reconciliation (incoming reports E0 until reconciled here — I re-ran, not relayed).** My own uncached runs this session, at the exit gate (main@0b9cf86) and the close baseline (main@674c844): `go test -count=1 ./...` = **15 packages green, zero failures**; `go vet ./...` clean; every charter-named crash window is a real instrumented crashpoint (`internal/crashpoint/crashpoint.go:11-26`); the C7 partial-state fixture present **both legs** (`internal/gate/derived_test.go:51,72`); tag `s1-close` derefs to f0dcb85; `OI-S1-F11-SWEEP` materialized (`docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md:161`). The s1 SITREPs' E2 claims **reconcile clean** against the repo.

**Sign-off chain (five independent battery verifications on record):** the `s1-core` pair (5 internal review rounds, 2 red-first bounce folds) · `s1.orchestrator-reviewer` (×2 approve) · **m-1 fidelity** (block→fold→approve; F-M1-1) · **m-2 fidelity** (approve incl. §J2 byte-custody) · **m-7 guide** (advisory read + plan-gate 7/7 `s1-plan-gate/…-171032` + deviation ruling + pre-concurrence `s1-exit-gate/…-200827`) · **master CTO** (independent reconcile + acceptance `s1-exit-gate/RECONCILE-orchestrator-planner-…-200929`) · **master VP** (plan gate + confirmatory pass `s1-vp-confirm/…-215730`) · **operator** (ratification, verbatim in `s1-merge-gate/…-220652`).

**The two exit-gate deviations (guide-ruled).** (1) **F11 breadth** — CONCUR / S1-sufficient (charter windows all covered; the class×point cross-product is structurally redundant at S1 — every class flows the same `store.Commit`→`fsio.WriteFileAtomic` path); condition satisfied via the typed owed-item `OI-S1-F11-SWEEP`. (2) **C7 mid-Complete re-crash** — NARROW BOUNCE for one fixture, which **caught a real bug red-first** (commit 33ee910): the composition proved full-state idempotence but not the partial state a mid-Complete crash actually leaves. The executable-claims standard earned its keep — the property is now *tested*, not *read*. The **F3 fidelity edge** also worked: m-1's F-M1-1 (seat-credential lifecycle) folded (shape (b): no remint in S1 → typed reject) **without touching the locked m-1 contract** — usage-fidelity, not re-design.

**Owed to S2 (materialize-first, nothing silently dropped):** `OI-S1-F11-SWEEP` = the full F11 class×point crash sweep under the S2 recovery machinery, dispositioned to the S2 exit gate; + the guide's note that S2 re-runs F9/F11 under the added recovery phases. One operator lint-waiver on record (a root-mode merge-claim-lineage false-positive on the filed CTO-acceptance copy — VP-classified trail-cleanliness; single-file lint OK; no merge commit exists — single-branch `main`, no remote, so the "merge gate" = the operator's ratification + tag).

**Process validation.** The whole governance apparatus ran end-to-end on real code and held: the guide plan-gate (F2 bootstrap), the m-1/m-2 fidelity edge (F3), the S1-scoped hardened exit gate, the guide's code-level deviation ruling, the VP confirmatory pass, and the operator's human merge gate — with the executable-claims standard catching a genuine bug the composition argument missed. **The design→build method is validated on its first slice.** Next: **S2** (thicken store/engine + the owed-item-as-typed-record projection — for which `OI-S1-F11-SWEEP` is the real first customer).

---

## S2 amendment — conductor-internal provenance folded into the m-1 design-of-record (`s2-amend-m-1`, ratified + **VP-co-signed** 2026-07-04, `RECONCILE-orchestrator-reviewer-040631` `approve`)

**Review-driven design-of-record amendment (mirrors the c6-fix close pattern) — the escalate-on-locked-contract-touch trigger firing correctly.** During S2, the slice introduced the first **conductor-authored, non-`submit`** store records (genesis, gc_marker, recovery/incident, derived-outbox — the engine's *own* machinery writes them; they don't arrive over a seat connection). The **m-1.implementer S2 fidelity review flagged (F-M1-1)** that stamping these is **new m-1 stamping-contract surface** — `FROM`/`ROLE` stamping is m-1's owned mechanism, and locked §6 defined it only for the governed-`submit` path. Per the charter (no silent re-design; amendments are review-driven), it was folded into the m-1 **design-of-record**, not left in an S2 relay. Operator-authorized; routed to **m-1 the owner** (the F2 escalate-on-locked-contract-touch trigger working as designed — and the contract-owner-in-the-loop fix from the fidelity-routing discussion).

**What landed** (`master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md`): a new **§6 bullet** — conductor-internal provenance: `FROM = ROLE = "system"` (reserved token, neither seat nor lane) for the engine's own records; `DeliveryState` within byte-exact **`{accepted, rejected, held}`** (`held` for incident/quarantine, per m-1-F12); `schema_version` in its `system_only` home; **`system` never accepted from the public `submit` path** (extends I2 reject-unbound — confusion-resistant D4, D5 residual) — plus a **§0.e fold-log**.

**Pair-approved + CTO-verified (E1).** m-1.planner fold (`DESIGN-planner-035030`) + **m-1.implementer `approve`** (`DESIGN-REVIEW-implementer-035323`, findings none). CTO independently reconciled: §6 + §0.e present as described; **byte-exact enum preserved 5×** (was 3, +2 from the two edits, no new/variant token); observer-selected 9× unchanged; **no c1 / design-lock reopen**; `frank/` code untouched by the amendment. **Lane discipline clean** — scoped to m-1's *stamping* mechanism; `record_kind` = m-2's header slot, the internal-record catalog/on-disk shapes = m-7/S2 — neither claimed. Folded into `ARCHITECTURE.md` §C4.1.

**Separate track (NOT discharged by this):** the S2 core fold/re-review gate stays open — s2.orchestrator-planner → s2-core.planner must fold **F-M1-1/2/3** into the s2 design/plan and pass the **m-1.implementer S2 re-review** before any S2 `DISPATCH IMPL`; s2's F-M1-1 fold cites this m-1 §6 clause as authoritative. This amendment records the *contract*; the S2 *build* fix proceeds independently via s2.orchestrator-planner. m-1.planner holds (no self-advance). *(Track closed in-slice: m-1 narrow confirm APPROVE preceded the s2 delegated dispatch — see the S2 close below.)*

---

## Step-1 · Slice-4 (s4 = THE WIRE-UP) — FOURTH BUILD-CYCLE CLOSE (CLOSED 2026-07-05; operator-authorized merge; VP-confirmed; the operator-as-transport ended, live)

**The founding goal of Step-1 is met — demonstrated, not asserted.** s4 = the per-seat **MCP shim** (a live host session's tool surface becomes `submit`/`project`/`read`; one shim = one credential = one channel; DI-2 preserved across the bridge) + live **seat lifecycle** (reconnect/catch-up, second-connect = one-active-channel-per-credential) + the **§7 config-change record** (registry/config evolution on an existing store, operator-authorized digest-change, m-7-guided). Built by a new slice-team on `s4-wire-impl` (16 commits, +4301/−126); the operator elected it over the old Section-4 (which became **s5**). **Close baseline: merge commit `main@fb61fda` (`--no-ff`, parents `a47381a`+`6a23cf0`) · annotated tag `s4-close` · authorized via `master/relays/s4-merge-gate/MERGE-GATE-orchestrator-planner-20260705-233708` (operator-directed) → executed by `s4-wire.implementer` (four bounded steps, no extras).** Verdict for the s4 scope: **complete at E2**, with a genuine **E3 live-host** centerpiece.

**THE MILESTONE:** a governed relay (`relay-4a33925b…`) flew from a live **Claude Code** seat (`s4-wire.host-a`) to a live **Codex** seat (`s4-wire.host-b`) — filed via `submit`, received via `project`/`read`, **conductor-stamped `from: s4-wire.host-a` (not agent-supplied), no human transport.** Two vendors: provider-agnosticism exercised. Step-1's "remove the operator-as-transport" is delivered live.

**Master reconciliation (my own runs):** gate-day store reconciliation (`s4-gateday/…-221608`: the 8-record ledger, the live relay stamp+checksum, the §7 re-render `11ecf52b→7b14f9c1`, owed cycle, `OPEN.md` empty); exit-gate accept (`s4-exit-gate/RECONCILE-…-231116`: battery 21 uncached + the F-GATE-2 registry fix verified at `6a23cf0`); post-merge (`main@fb61fda` graph + tag + **battery 21 green uncached this seat**). Five verification stations + the VP confirmatory pass.

**Gate-day findings — all dispositioned:** **F-GATE-1** (MCP handshake `serverInfo.version` omitted → live Claude Code rejected in 13ms; the E3 catch three E2 stations were structurally blind to) fixed + class-pinned. **F-GATE-2** (the owed fill-time-authority gap I raised: `owner`/`source`/`target_surface`/`disposition_path`/`disposes_owed` code-required but FieldSpec-undeclared → the rendered form never prompted them) fixed = **five registry rows with `required_when: record_kind_in` predicates**, red-first, m-2-confirmed, 3-file fold — *required, and now rendered*. **F-GATE-3** (shim process-stderr names the socket path) ruled OUT of I-PH scope (delivered surface path-free; host-side stderr under the attach model), m-7-ratified + VP-concurred.

**All s3-scope-q1 conditions discharged; [VP-W1..W4] all shown-not-asserted** (channel-active live · transport/provenance-only agent-volunteered · W3 ruled · existing-store never re-genesis). **`OI-S3-CONFIG-CHANGE` discharged → Step-1's mandatory owed set is EMPTY.**

**Honesty scope (VP-ruled sufficient):** the centerpiece (live relay + second-connect) = genuine two-vendor live-host E3; the adversarial/crash/§7/owed legs = live-store, master-scaffolding-driven, master-verified (the three scaffolding bugs were harness, not frank). The VP ruled the caveat-stated evidence set does not require a clean procedure-of-record re-run before merge.

**Owed out of s4: one discretionary follow-on** — `OI-S4-TOKEN-SCOPE` (review narrowing owed/`genesis` record_kind authoring from `*` toward operator; a non-operator owed filing grants no authority, so hygiene not a hole; **incl. `genesis` per VP+m-7**). Operator authors at discretion; does not gate anything.

**Step-1 status after four slices:** spine (s1) → engine (s2) → the real form system (s3) → **wired + carrying live cross-vendor traffic (s4)**. The one remaining slice is **s5** (consumer schema slices + the §C4 fixtures) — which builds *over the wired conductor*, its registry rows landing as real §7 config-change records, its relays the first usage data. After s5 + the roadmap step-exit test, **Step-1 closes.**

---

## Step-1 · Slice-5 (s5 = CONSUMER SCHEMAS) — FIFTH BUILD-CYCLE CLOSE (CLOSED 2026-07-06; operator close-gate exercised; the dogfood slice — found the transport's floor, then finished clean on file relay)

**Scope delivered in full, against the s5-resume-adapted gate.** The registry is **complete + versioned**: 47→**83 rows @ `s5-fieldspec-v3`**, 24 named_enums — the four consumer domains' fields (m-3 observe · m-4 routing-record · m-5 archetype · m-6 gate/ODB) declared **STEP-GATED OFF** via the m-2-blessed `layer_present:observe` step-gate idiom, proven dormant by the enumerated negative fixture (full seat-class × phase × tier sweep, 334 subtests) [VP-W3]. The owed §C4 fixtures: **③** raise mechanics (token REWRITE + `gate_category_raised:"yes"` + `gate_category_pick` provenance per MR-1) + **live wiring** (S1 configured · S2 verdict-path live · S3 input-atom-pending = inert, exactly the confirmed set; fail-before-serving on invalid policy config — the run's first fail-closed startup evidence; the claim boundary **in the code** at both wiring sites) · **⑤** real scanner code, m-3's (a)/(b)/(c), at the real outbox-drain leg, present-but-dormant per R-2 · **GRILL_REQUIRED empty** (both legs pre-existing — a scope *shrink* found by the fresh team) · **`routing_escalation`** landed exactly per the m-2 delta (14 members, pre-`other`; §J2 edit CTO-applied) · **I-PH extended** over every new surface. **Versioned + migratable per Q4-final**: version label bumped, zero-loss replay + negotiation/refusal legs, **NO envelope migrator** (`migrate.Current = 1` — the R-1 two-axes ruling held), **and the real-store leg green at three seats against the archived dogfood store**. **Close baseline: work tip `main@f31d43a` · close record `7e5c527` (docs, 6 files +685) · annotated tag `s5-close` · three integrations each under the WRITTEN authorization chain (grant → reviewer approve → execution → orchestrator verification), close docs+tag operator-authorized in-session.** Verdict: **complete at E2.**

**The dogfood arc (the run's second deliverable).** s5 began as **the first team to run its governance ON frank** (dispatch `relay-45e263fa` submitted, boot-ack + full onboarding returned through `submit`/`project`/`read`; 5 sub-seats + an operator seat minted; hub-and-spoke m-x routing enforced by construction). One day of real multi-seat traffic surfaced **17 transport findings** — headlined by **F11, the lineage livelock** — and the operator ruled stop-the-line: conductor decommissioned, store archived, **`master/TRANSPORT-FINDINGS-2026-07-06.md` = the ledger and the transport-fix cycle's spec seed** (PROTOCOL-DEVIATIONS **B8**). s5 resumed on stock file relay (`s5-resume`, gate adapted: the §7-live-store leg rides the transport-fix relaunch; the dogfood-record exit item closed early by the ledger) and lost nothing: the archived store then served the replay's real-store leg.

**The fidelity machinery at full stretch (the run's third deliverable — the process proved out).** The complete consumer-semantics packet (Q1–Q11) answered by **all six domain owners** with zero unresolved contradictions; the M-1/M-2 blockers ruled (the dormancy idiom blessed; the ③ signal set S1+S2+S3 composed with the token-REWRITE mechanics); **three adversarial implementer approves** (m-4 (f)+(a) w/ C1+C2 registered as §C4 Step-3 carries; m-6 signal-set w/ the claim boundary pinned load-bearing; m-1's dual-confirm — owed rows operator-only as a Step-1 scope posture, **`genesis` removed from EVERY scope** as incoherent-by-construction). Master rulings G-1..G-5, MR-1 (pick provenance), MR-2 (ceremony floor on narrowings), R-1 (compat-class ≠ envelope-migration).

**The fresh-team bar paid a SIXTH time:** five live defects, triple-confirmed at file:line — DEF-1 (the raise stamp's `"true"` into a `["no","yes"]` enum), **DEF-2 (system/computed headers lane-suppliable — closed as the general typed-reject rule)**, DEF-3 (EVIDENCE_TARGET not actually required), DEF-4 (③ stamp path untested), DEF-5 (`resolves_gate` consumed-but-undeclared). All fixed red-first in-slice. Plus the run's process catches: the shared-id lineage defect (s5's own, caught by both pairs → per-hop re-threaded plan gates) and the MF-5 combined-tree red (the run's first integration-gate catch).

**Master reconciliation (my own runs at `7e5c527`):** vet clean · **battery 23 packages ok, uncached, zero failures** · `-race` green (channel/store/recover/engine/egress) · probes: registry 83 @ `s5-fieldspec-v3`, `gate_category` 14 w/ `routing_escalation` pre-`other` + A-set updated + `routing_unavailable` nowhere, `gate_category_pick`/`surface_intent`/`resolves_gate`/`attestation_source` present, record_kind scopes narrowed (`genesis` NO scope · owed/`gate_resolution`/`disposition` operator-only · `*` = `diagnostics`) — **this also closes `OI-S4-TOKEN-SCOPE`** · scanner zero production callers · `migrate.Current = 1` · claim boundary at `detector.go:25` · tag `s5-close` peels to `7e5c527`. Prior stations: my `b30df4d` probes; s5's own uncached runs at every tip; the reviewer's independent approves at each merge gate.

**Registered follow-ons (none s5's):** S3 activation (target-branch atom + operator config, a future registry pass) · the away-bridge carries · the step-(d) set (§C4) · **the transport-fix cycle — NEXT** (the F1–F17 ledger its input; the fixed conductor's first live act = applying s5's registry via §7 per the adapted [VP-W7]) · O-3/O-4 (m-5's, with master) · the 092547 known-dirty lint artifact (operator's waiver option stands).

## Step-1 · Slice-6 (s6 = THE TRANSPORT FIX) — SIXTH BUILD-CYCLE CLOSE **and the STEP-1 CLOSURE** (CLOSED 2026-07-08; operator-granted merge; VP-confirmed; the step-exit test passed LIVE on the fixed conductor)

**Close baseline: merge commit `main@6a1198a` (`--no-ff`, parents `73116e9`+`58f2233`) · annotated tag `s6-close` · executor report `frank/.relays/s6/s6-merge-gate/MERGE-GATE-implementer-20260708-030603` · battery 24-ok uncached at the merge commit, re-run at THIS seat.** The VP-co-signed amendment set (`S6-AMENDMENT-SET-2026-07-06` r3) implemented WHOLE: branch-A conductor-computed PARENT + fallback hints · the ONE envelope codec · A-1 stable-schema digest · A-2 idempotent intake · A-3 live mint (`seat_mint`) · A-4+§D the I1-P store lock · §B `project()` default-accepted · §C scoped waivers + `waiver_retraction` · F13 three-layer record_kind · D-1/D-2 · the B-1/B-2/B-3 boot stage with DERIVED-ONLY activation. Registry: +7 transport rows +2 boot fields, `ORCH_REVIEW_WAIVER` retired, zero marker rows — [VP-W3] verified at the byte grain at three independent recomputations. **Ten independent verification chains** (the pair ×2, s6 orchestrator, the s6 reviewer, the 4-lens panel, m-1/m-2/m-7 fidelity+guide, master's two stations, the VP's own battery, the operator's authorship legs).

**THE STEP-1 EXIT TEST — PASSED LIVE (2026-07-08, gate day, the operator + master, the s4 division of labor):** (a) ROADMAP:83-85 ×3 + a gate → a local outbox item, store-verified · (b) **the operator §7-apply of the registry as the fixed conductor's FIRST live act** — the adapted [VP-W7] discharged (with two correct operator-seat HOLDS en route: the digest-canonicalization refusal + the missing-context diagnosis) · (c) **the F11 redrive: all 14 archived s5-dogfood records accepted in order — ZERO parent-class bounces, ZERO same-context re-renders; the 9 archived seat-picked parents all fell back with `parent_hint_honored: no` + verbatim hints (the GRILL_LOCK triple, live)** · (d) the live boot walk: a no-restart `seat_mint` → a real session → the boot form → `minted→bound→active`, activation ref = the boot relay, derived from records alone. **The traffic that livelocked s5 lands whole on the conductor built from its wreckage.**

**Mid-slice design integrity:** the R1 generation-boundary catch (in-flight commands crossing a re-mint pivot → typed `credential-superseded`) · the panel's re-mint crash-window find (m-1 option A, eight redlines, structurally verified) · a voided first dispatch caught at the pair's own lint gate. **The fence held at the finish line:** gate day's two product findings ride out as typed OIs (`OI-S6-BOUNCE-CLASS-UX` · `OI-S6-ENVELOPE-KEY-HYGIENE`) because they sit outside the co-signed set. Gate-day F-GATE-s6-2/4 folded docs-only (the procedure wording; the hosted-seat `tools/list_changed` caveat → `docs/ops.md`).

**STEP-1 IS CLOSED.** The goal line — "remove the operator-as-transport" — is delivered on a transport that survives its own governance load: proven by carrying the exact traffic that broke its predecessor. Riding out of the step: the two s6 OIs · the step-(d)/away-bridge §C4 carries · the C1/C2 Step-3 routing carries · **the INV-CATALOG follow-on (registered at §C4; the FIRST item of the post-close queue)** · the relaunch ops notes. The dogfood relaunch (frank carrying its own governance again) rides Step-2 planning — its relaunch gate (the F11 traffic pattern must not livelock) is now literally a passing test.

---

**Step-1 status after five slices: THE BUILD QUEUE of the original decomposition IS EMPTY.** spine → engine → forms → wire → **consumer schemas + the owed fixtures (s5)**. **s5 close RATIFIED (operator, 2026-07-06).** **Operator ruling, same session: Step-1 does NOT close until the transport is fixed** — the goal line is "remove the operator-as-transport," and a transport that livelocks under multi-seat load (F11) has not durably removed it (s5 finished on hand-relay). **s6 = the transport fix, IN Step-1:** design-amendment phase (m-1 lineage/parenting = THE fork, grilled; m-7 engine liveness; m-2 single-codec) → the s6 build slice → **the step-exit test on the fixed conductor** (the ROADMAP:83-85 legs + §7-applying s5's registry as the first live act [the adapted VP-W7] + the F11 regression leg: the archived dogfood traffic pattern replayed without livelock). Then Step-1 closes.

---

## Step-1 · Slice-2 (S2) — SECOND BUILD-CYCLE CLOSE reconciliation (CLOSED 2026-07-04; operator-authorized merge; first real branch integration)

**The engine is thickened and the S1 debt is paid — through the product itself.** S2 = full crash-recovery (phases 0–4, reified phase machine, `Ready`/`Diagnostics` split) + durable FIFO (single intake-writer — closing 2 latent S1 races the fresh team's audits found; segmented crash-safe journals) + GC/genesis (store-root-pinned config; marker-first GC, off-by-default, drained-segments-only) + **the owed-item-as-typed-record projection** (one obligation mechanism; materialize-first; no auto-completer). Built by a **new** slice-team (B1: new sprint = new team) on branch `s2-core-impl` (17 commits, 45 files, +4475/−263), guided by m-7 with m-1 fidelity (must-revise → prescribed folds → narrow APPROVE; spawned the `s2-amend-m-1` design-of-record amendment above). **Close baseline: merge commit `main@b322b6d` (second parent = the authorized `18bd62e`) · annotated tag `s2-close` · ledger through `f3f66be`.** Verdict for the S2 scope: **complete at E2.**

**Master reconciliation (my own runs):** exit gate — battery **18 packages ok uncached + vet clean at `18bd62e`**, branch/base/diffstat exact, **the real store inspected directly** (`$HOME/frank-s2-store`: genesis + operator-stamped owed record + disposition; `OPEN.md` = empty table). Close — merge graph verified (`b322b6d` parents `b964fd0` + `18bd62e`), tag peels correct, **battery 18 ok uncached at main post-merge (this seat)**. Exit-gate acceptance: `s2-exit-gate/RECONCILE-orchestrator-planner-20260704-143728`.

**Firsts on record:** (1) **`OI-S1-F11-SWEEP` CLOSED as chartered — frank's first real governance transaction**: an operator-authored, channel-stamped owed record → the executed full F11 class×point sweep (clean-completion cells falsified red-first) → an operator-authored disposition → open set empty *by projection on a real store*. The S1 owed item is hereby **closed in this ledger** on that record. (2) **The first real `git merge`** under the layered merge authority: operator authorization (verbatim on record) → token-bearing MERGE-GATE (`s2-merge-gate/…-151500`, TO exactly one implementer) → implementer execution report (four steps, no conflicts, no extras) — the upstream merge-claim lineage resolving end-to-end on real git history. (3) **The F2 escalate-on-locked-contract-touch trigger fired correctly mid-slice** (→ `s2-amend-m-1`).

**Owed items riding out of S2: none.** Deferred unchanged: the MCP live-adapter/wire-up slice (operator's call, awaiting a testbed) · S3 registry/linter · Step-2 observe · Step-3 routing · S4 consumer schemas. Trail hygiene: two operator scoped waivers (superseded-file root-lint residue; one-file-scoped each, no standing rule); all live relays lint clean both modes.

**Process validation (second data point):** the slice lifecycle held under a *new* team — paired audits (found real S1 fragility) → grill → guide confirms → fidelity cycle → conditioned delegated dispatch → 3-round adversarial review → hardened gate → layered merge. The fresh-team-per-slice model is now evidence-backed, not just principled. Next: **S3** (thicken forms/lineage — full FieldSpec registry + 62-check linter dissolution + the FULL dissolved-linter replay + `schema_version`/migrators; guide **m-2**; new slice-team).

---

## Step-1 · Slice-3 (S3) — THIRD BUILD-CYCLE CLOSE reconciliation (CLOSED 2026-07-04; operator-authorized `--no-ff` integration; VP pre-integration confirm)

**frank now speaks the real protocol.** S3 = the **full FieldSpec registry** (the S1 6-enum MVP dialect deleted; the team's actual header vocabulary — ROLE/PHASE/AUTHORITY/lineage — renders, validates, commits, projects end-to-end on fresh stores) + **fill-time authority by negatives** + **the 62-check linter dissolution PROVEN BY EXECUTION** (115 anchor-grain disposition rows: 109 dissolved / 1 retained / 5 evidence-grounded obsolete; the frozen 146-case oracle over the 243-file corpus — 96 fail-side caught-or-genuinely-obsolete, 50 pass-side non-overblocking, zero uncovered, structurally enforced) + **`schema_version`/migrators** (v1→v2 fixture walk + 3 typed refusal legs) + **both §C4 owed carries discharged** (R2 per-column negatives · the GRILL_REQUIRED row) + the **§10c lineage engine** over incrementally-maintained tables with the S1 grant-narrowing carry landed. Built by a new slice-team (guide **m-2**; m-7 config-seam consults; the VP's 4 pre-handoff watchpoints folded as [VP-W] dispatch rows) on `s3-form-impl` (15 commits, 38 files, +5713/−749). **Close baseline: `main@b5a2c95` (`--no-ff`, parents `91a8a26`+`fe7308e`) · tag `s3-close` · ledger `50290e1`.** Verdict: **complete at E2** (fresh-store qualifier stated on every claim surface).

**Master reconciliation (my own runs):** exit gate — battery **20 ok uncached + vet clean** in a clean worktree at `fe7308e`; `dispositions.json` probed directly (115 rows · 110 distinct anchors · the `:840-873` census range rowed · zero `uncovered`); **the [VP-W] obsolete-adjudication rule checked row-by-row** (all 5 obsolete rows ground on real vanished surfaces/replaced invariants — none on a design-of-record change); MVP-dialect deletion confirmed; the real S2 store untouched (3 records). Close — merge graph + tag + ancestry verified; **battery 20 ok uncached at `main@b5a2c95` (this seat)**. Acceptance: `s3-exit-gate/RECONCILE-orchestrator-planner-20260704-214606`. **Five independent verification chains at close** (implementer · pair planner · s3 orchestrator · VP pre-integration confirm `s3-vp-confirm/…-215937` · the four-lens panel).

**Mid-slice governance of record:** the **s3-scope-q1 escalation** (the m-7 §7 config-change record — S2's "(S3)" forward-pointer vs the dispatch IN list) was escalated correctly, **ruled DEFER** with five conditions (`s3-scope-q1/RECONCILE-…-171608`): fresh-store posture ratified (restart-with-new-store IS the true semantics under no-hot-reload), `OI-S3-CONFIG-CHANGE` materialized, owner = the wire-up slice, hard backstop = before any store is declared persistent, (b)'s guide-split conditions carried forward, the stale pointer superseded on record. One dispatch-trail judgment endorsed (r3 PLAN-REVIEW parent across two non-content folds; the content-changing-folds-need-fresh-approves rule unweakened). One trail anomaly dispositioned (a vanished INDEX-row target in a hand-relay window; content preserved; VP-concurred no-hold).

**Riding out of S3: exactly one owed item** — `OI-S3-CONFIG-CHANGE` (→ the wire-up slice). **Step-1 status after three slices:** the machine is whole on fresh stores — spine (S1) + engine (S2) + the real protocol (S3). The wire-up prerequisite is MET; the deferred MCP slice is now "connect the real thing." The remaining decomposition item is **S4** (consumer schema expression + the ③/⑤/routing_escalation fixtures); the next-slice fork (S4 vs wire-up) is the operator's, as chartered in the s3-dispatch.

---

## Step-2 · Slice-7 (s7 = INV-CATALOG) — FIRST STEP-2 BUILD-CYCLE CLOSE (CLOSED 2026-07-10; operator-granted merge with the recognized authorization field; VP-approved after three integration rounds; the constitution became executable)

**Close baseline: merge commit `main@2e1b4f0` (`--no-ff`, parents `54420dbc`+`5e6bf83`) · tag `s7-close` at the merge commit · executor = m-7.implementer (operator-named, per the "original implementer" precedent) · executor report `master/relays/s7-merge-gate/MERGE-GATE-implementer-20260710-190100` · battery at the merge commit: serialized uncached, FILE-captured, 25 ok + 2 no-test-files / exit 0 — at the executor's station and re-run at THIS seat.** The merge-authority trail lints clean end to end: the operator's grant rode the execution dispatch as `HUMAN_MERGE_AUTHORIZATION` at grant time (`MERGE-GATE-orchestrator-planner-20260710-185615`) — the s7a-merge-gate lesson applied, no scar this time.

**What landed:** `test/invariants` — the standing global laws as ten NAMED executable checks (`TerminalEnumByteExact` · `ThreeVerbSurface` · `R2NoModelPredicate` incl. the `any_row` column negatives · `DerivedOnlyActivation` [seat-lifecycle grain] · `SoleGovernedWriter` [sole *governed* path, D5 stated] · `PathHygiene` [live-root canonical census + the tree-wide-discovered egress boundary + per-family planted negatives] · `CanonicalWins` · `OnePivotPerMutation` · `IntakeOutcomeOneToOne` [executes `recover.RunWithProcessor`, at-most-once grain] · `RebuildBeforeOpen`) + `catalog.v1.json` (the law registry: owners + the single-writer/owner-fidelity convention; the registered "governed like `registry.json`" property is STAGED to the s8 §7-pinning carry and deliberately not claimed). **Effect, live in every battery from this merge on: weakening a global law = a red battery naming the law; the only path through is the amendment ritual.**

**The build shape (the operator's B10 lean model, its first run):** the m-7 pair as the build pair — no slice team; the plan VP-gated THREE rounds before a line of code (seat authority · executable-rows-not-pointers · staged governance); the fidelity matrix closed by the owning seats — rows 1 (m-2) · 2/4/5 (m-1) · 3 (m-2+m-4, twice, final against the merged guard) · 6 (m-1+m-2, two fold rounds) · 7–10 (the pair's claimed half). The VP's integration gate then refused ASSERTED checks until they EXERCISED their mechanisms — three escalations on row 6 alone (catalog-seeded symbols → boundary-derived symbols → the boundary itself discovered tree-wide, proven by a genuine-new-file scratch red) and one on row 9 (the recovery re-enqueue executed by its production owner, not simulated). The final claim boundary is explicit and VP-accepted: portable idioms are detected tree-wide; genuinely novel wire idioms are not scanner-definable a priori and route through the locked §8 design-amendment + grill ritual.

**The catches (the slice paid for itself before it merged):** ① **F-S7-R2-COLGRAIN** — writing the R2 law honestly exposed a REAL pre-existing production gap (the predicate grammar compiled `any_row:routing_assignments.chosen_model`; the C1 column-grain trigger was already met); the s7 fence held (finding OUT, never an in-slice fix) → the **s7a guard lane** (m-2 pair, B10 second application, its own full pair/fidelity/VP/operator chain) landed the default-deny column-grain guard at `main@54420dbc`, now guarded by the named law that found it. ② The VP's verification exposed **`OI-S7A-CLOSE-ONCE-RACE`** — a real, pre-existing client-lifecycle double-close race (`Client.Close`/`readLoop`), reproducible under focused `-race` on the untouched baseline — registered with owner m-7 and a HARD gate: dispositioned before the s8 dogfood opens live MCP channels. ③ **S7A-TRAIL-FINDINGS** (kickoff 2c′): five trail-integrity failures across the arc — an orphan parent, a false byte-identical claim, a DISPATCH_ID collision, master's own merge-authorization shape defect (caught by the executor AFTER an irreversible push; dispositioned honestly, the scar kept), and the `one_by_id` earliest-wins footgun — every one at the layer the conductor kills by construction, while design and code judgment stayed sound throughout. The cleanest dogfood argument yet; the s8 dogfood evaluation names which classes the conductor demonstrably kills.

**Conventions pinned by this slice (now standing):** pair loops under a master thread mint unique sub-`DISPATCH_ID`s · every merge-grant relay carries `HUMAN_MERGE_AUTHORIZATION: granted — <context>` at grant time · battery evidence is FILE-captured and sequence-honest (never pipe-counted, never flattened while a race item is open) · status reports state the tracked-vs-untracked distinction exactly.

**Riding out of s7 (none of it moved by this close):** the s8 §7-pinning + genesis condition (the dogfood store initializes from the exact `s7a-fieldspec-v5` bytes, member SHA `1ef6abab…2485`, with the recorded composite digest; the discovery's walk roots + boundary-file set join the governed censuses) · `OI-S7A-CLOSE-ONCE-RACE` (the pre-dogfood hard gate) · `FLAKE-SOCKET-PAR` · the two s6 OIs (`BOUNCE-CLASS-UX`, `ENVELOPE-KEY-HYGIENE` → s9, unchanged) · the step-(d) and C1-representational/C2 Step-3 carries. **Next: the pre-s8 package** — the race-disposition staffing ruling, the layer-activation knob + catalog-pinning design rounds (m-2+m-7), FLAKE-SOCKET-PAR, executor isolation (m-3+m-7) — then the s8 observe-spine dispatch (fresh team, m-3 guide) and the dogfood relaunch on its store.

---

## Step-2 · s7b micro-lane (the close-once disposition) — CLOSED (2026-07-11; operator-granted merge, recognized-field trail; VP first-pass approve; three registered defect classes closed BY MECHANISM)

**Close baseline: merge commit `main@691d034` (`--no-ff`, parents `2e1b4f0`+`e155aa6`) · NO tag (micro-lane, the s7a precedent) · executor = m-7.implementer (operator-named) · report `master/relays/s7b-merge-gate/MERGE-GATE-implementer-20260711-014419` · serialized uncached battery at the merge commit 25 ok + 2 no-test-files / exit 0, file-captured at the executor's station and re-run at THIS seat · the merge-authority trail lint-clean end to end (`HUMAN_MERGE_AUTHORIZATION` at grant time).**

**What landed (three commits, B10 third application — the m-7 pair, planner-first with delegated dispatch):** `a2a6966` — the **idempotent channel close**: `close(c.done)`/`close(s.done)` owned solely by their `sync.Once` (`server.go:230/:527`), killing the double-close panic the VP proved on the untouched baseline (`OI-S7A-CLOSE-ONCE-RACE`) · `5c678b4` — fixture startup hardening (cached builds + honest deadlines) · `e155aa6` — the master-granted delta: the ONE-LINE crashpoint block-after-kill (`select {}` after the self-SIGKILL — the process state at delivery = the state at the Hit, deterministically), the `test/invariants` hardening (the FLAKE class's third member), and the `TestMain` child-mode short-circuits.

**The lane's distinguishing conduct:** the implementer's **model stop** on two unexpected acceptance failures (no widening, no pin, no false unification) → the pair planner's **mechanism-grade diagnosis** (the kill-then-return race root-caused from one call boundary + asynchronous signal delivery — the only story consistent with the crashed-child-plus-committed-file signature) → a bounded master grant → both classes closed **by removing the cause**, proven by the sequence-honest flip: fail·pass·fail pre-delta → three consecutive default-parallel full-suite greens post-delta, corroborated independently at the pair, master, and VP stations. **No parallelism pin was needed; none was checked in.** m-1's five-point lifecycle confirm: the close ordering moved nothing observable — B-3 bind/rebind, §8.5 re-attach, one-active-channel, and supersession byte-identical; the auth interval stays fail-closed server-side.

**Register effects (this close):** `OI-S7A-CLOSE-ONCE-RACE` **CLOSED — the s8 live-MCP-channel gate LIFTS** · `FLAKE-SOCKET-PAR` **CLOSED** (all three members, by mechanism) · `CRASHPOINT-KILL-RETURN` **CLOSED**. The honest-sequence evidence rule (no flattened battery summaries) **stays standing** — it proved its worth twice in this lane and costs nothing. Unmoved: the s8 design/genesis/config gates, the design-lane reconciles, every operator gate.

## Cross-domain — the CONFUSION-FIREWALL DIRECTIVE (OPERATOR-RATIFIED 2026-07-11; a scope amendment, not a build close)
Operator-originated (the m-2.planner session), scribed `frank-threat-model-scope/SITREP-planner-20260711-162331.md`, master-assessed with two interpretation rails (`…-162826`), **ratified verbatim the same day**. Text of record: `ARCHITECTURE.md` § CROSS-DOMAIN SCOPE AMENDMENT. Impact at adoption: **all four s8 design locks SURVIVE by function analysis** (digest = drift/one-truth · adoption atomicity = crash-safety/history-truth · A2 = anti-confidently-wrong-machinery · executor ceiling/sandbox-shelf = the directive pre-applied); claims were already confusion-graded (F5(a), FX-CFG-4, §7 ceilings) — no lock reopened, no rip-out pass. Deltas routed: the s8 addendum (SEQ-1 fresh-store-first steer + Rails A/B as the build-time review criterion) · **the adjudication rung = kickoff design item 9** (s9, m-3-owned, the NAMED egress-fence constraint: a cloud-judge call is an external send) · the m-1 identity-seam rail (structured stamp, could-carry-a-signature, no work now). The positioning line binds the publication track at the public flip.

## Step-2 · Slice-8 (s8 = THE OBSERVE SPINE) — SECOND STEP-2 BUILD-CYCLE CLOSE (CLOSED 2026-07-12; operator-granted merge with the recognized field at grant time; two-seat slice-exit APPROVE + master's independent third-seat verification; the phase-opener slice — the conductor now OBSERVES done-ness)
**Merge `main@8941889` (`--no-ff`, parents `691d034`+`b2c2062`), tag `s8-close` (lightweight, the s7-close form), pushed to private frank-dev (both refs remote-matched); executor s8.implementer under `s8-merge-gate/MERGE-GATE-orchestrator-planner-20260712-203020.md` (HUMAN_MERGE_AUTHORIZATION at grant time — the convention's third clean run); merged-not-deployed; master verified: merge tree BYTE-IDENTICAL to the approved candidate + own uncached serialized battery 25 ok/0 FAIL/exit 0 at that content + the executor's file-captured run at the merge commit (capture SHA `13a30a08…b52198` reproduced) + vet clean + parents/tag/ls-remote exact.**
**What landed (the whole arc):** observe-as-send LIVE inside the atomic submit through GOVERNED SUPPLY ONLY — the four s8 design locks (m-2 grammar · m-3 registry+§4a+§12+§13 · m-7 config-host r13 · m-7 executor-host r4) + two ritual amendments forced by build reds (claim-input: the executable_claims declaration surface, fieldspec v6→v7; supply-set: engine v1→2 `supply` member — descriptors/lane-roots/schema-refs pinned, ambient cwd + silent timeout fallback KILLED) + the five-row hardening (genesis profile enforced both entrances incl. the bless door · §4a read mechanism w/ descriptor-proven confinement + detach/breaker · the §13 absence floor w/ the honest E0 claimless degrade + live string-veto · stage-once/hash-staged executor · Row-5). Exit proven at the REAL SOCKET at three seats: fresh production genesis → live v5→v6→v7 + observe-flip config_changes → false-done REJECTED typed (read-file AND suite legs) → true-done accepted w/ observed stamps → decision-② authority holds → E2 ceiling labeled everywhere.
**The slice's own trail = the product's evidence case:** EIGHT failure classes caught in-slice, the catch layer moving earlier every round (r5 production RED → r6 independent review → r7 pre-exit sweep → r8 the build's OWN review); the reader-with-no-writer/label-without-mechanism ledger ran to seven instances, ALL resolved by ritual amendment, zero improvisation; ONE fence breach in ~9 gated rounds (`registry.go` Parse — granted exact/non-precedential, recorded as the scar) answered by the diff→license table that makes its class mechanical; TWO master overspecifications (Rail-A additive; the r7 wall-deadline) corrected by independent owner review — the review structure catching the orchestrator is the system working. The two mechanical exit tables (17-row consumption→supply; 52-path diff→license) ship IN the package.
**Register effects (this close):** the s8 exit-gate legs 1+2 of the roadmap step-exit are LIVE and demo-able · carries filed: s9 = adjudication rung (item 9) + turn-baseline fence (item 10) + red→green-differential labeling · s10 = the two ratified sunsets (timeout park/prompt; static gate → live prompt) · public-release = live-store bless activation · non-blocking in flight = m-3's two §13 doc-alignment folds · ledger watches (linkname portability; Row-1 fixture strengthening; `Outcome:"unsafe"` clarity; the Row-4 accepted consequence) ride the slice records. **The LIVE RELAUNCH is the operator's separate checklist act** (lane-root pin via governed config · catalog bytes at `943f07bb…e209d` · honest E3/E4 reporting) — NOT performed by this close.

## Step-2 · Slice-10 (s10 = THE COMMS SPINE) — THIRD STEP-2 BUILD-CYCLE CLOSE (CLOSED 2026-07-13; operator-authorized merge with a disclosed trail residual + ratification offered; B11 straight-through cadence's first full slice; ★ ROADMAP EXIT LEG 3 PROVEN LIVE — all three step-exit legs now exist)
**Merge `main@39474d0` (fast-forward from `8941889`, 13 commits T1–T11 + two review folds), tag `s10-close`; **pushed to private frank-dev 2026-07-13 (origin main 8941889→39474d0 + tag, remote-matched — a PRIVATE dev sync, NOT the public flip which stays a separate remote/decision);** master verified independently: lineage/tag/count/clean exact + own uncached serialized battery 25 ok / 0 FAIL / exit 0 + vet clean.** **Trail residual, disclosed by the pair at close (the honest-scar pattern):** the merge ran on the operator's DIRECT conditional authorization (verbatim-quoted; the condition — a second review — satisfied first) without an operator-FROM grant relay carrying the recognized field; the one-line post-hoc ratification was OFFERED in `s10-merge-decision/MERGE-GATE-planner-20260713-041930` — the s7a class, self-disclosed pre-close — and **operator-COUNTERSIGNED ("ratified", 2026-07-13; recorded w/ the recognized field in `MERGE-GATE-orchestrator-planner-20260713-125915`): INCIDENT CLOSED**, ordering scar kept as the permanent honest artifact.
**What landed:** the minimum A-gate path END TO END on the fresh v8 dogfood store — ODB render/capture (`odb` + `resummon_command` interpreter-bearing record kinds via the governed v7→v8 transition, byte-exact to the m-2/m-7 owner deltas, marker-first old-reader refusal green, pinned-v5 tripwire held) → park → the operator's validated reply (operator-FROM, T5) → local re-observe on wake → EXACTLY-ONCE wake (the committed exit fixture counts observations exactly 2, wakes exactly 1; the operator's own authenticated `approve` on the production binary) → deterministic resummon commands deduping through A-2. **Both operator sunsets demonstrated GONE with behavioral negatives:** the silent auto-kill → soft-expiry park+ODB `{kill, extend}` with the block-only hard ceiling (the accepted m-7 two-sentence amendment); the static-only side-effecting gate → live-prompt pre-spawn default-DENY with typed refusal and GATE-LIFT-with-zero-spawn under the binding claim-text rider.
**Review structure under B11 (the cadence's first datum):** ONE end-of-slice review (planner; independent battery + 3 lenses → MF-1 fold `f481042`) + an OPERATOR-ORDERED second independent review (8 finder angles + verify → 7 must-fixes + 2 hardenings folded at `39474d0`: genesis v8 integrity bypass · commit-time TOCTOU on snapshot uniqueness · completed-then-killed lane bricking · executor `waited` drop · resummon production wiring · scoped registry-fault/single-load · root-ctx shutdown · prompter collision guard · nested-reply timeout — each RED-first + race batteries) + three owner CONFIRMs (m-2 byte-exact · m-7 four-cell · m-3 three-edge; stop-on-contradiction never fired). Playbook datum: **the two-layer end review caught what fidelity confirms and a green battery structurally cannot — it earned its churn.** Fence integrity: zero post-dispatch deviations; `Spawn:83-95` untouched; the I-PH census caught + repaired one real sink leak during the second fold (the invariant catching live).
**Register effects:** ROADMAP exit leg 3 SATISFIED → **all three step-exit legs live** (legs 1+2 at s8, leg 3 here); the step-exit test itself runs AFTER s11 per the ratified order. Carries filed (kickoff): the s11 design cell (non-Accepted-resolution/blocked-prompter re-prompt + the claimless authority-floor Held edge) · m-7's A-2 verdict-cache invalidation + two accepted residuals · m-6's G4 resummon-cadence config surface · the s11 cleanup card (9 items) · the standing shelves unchanged (side-effecting EXECUTION beside the OS sandbox · the 8a branch). Catch ledger (one series): s10 = #1 fence expansion pre-code · #2 F-b owner refutation pre-code · two correct mid-build holds · #3 the second-review seven post-build/pre-merge.

## Step-2 · Slice-9 (s9 = THE EVIDENCE THICKEN) — FOURTH STEP-2 BUILD-CYCLE CLOSE (CLOSED 2026-07-14; operator-granted merge with the recognized field at grant time — the convention's FOURTH clean run; two-seat exit APPROVE + master's third-seat verification; a FRESH slice pair after the Option-A staffing correction)
**Merge `main@d91fcfb` (`--ff-only` fast-forward from `39474d0`, 9 linear commits), annotated tag `s9-close` (obj `841ae9d`, peels to `d91fcfb`), PUSHED to private frank-dev in the same gate (both refs remote-matched); executor s9.implementer under `s9-merge-gate/MERGE-GATE-orchestrator-planner-20260714-020520` (HUMAN_MERGE_AUTHORIZATION at grant time); merged-not-deployed; master verified: the FF tree is BYTE-IDENTICAL to the candidate (own uncached serialized battery 26 ok / 0 FAIL / exit 0 at `d91fcfb` covers it; the executor's capture SHA `2ac1161a…` reproduced at my seat), tag/parent/ls-remote exact, status clean.**
**What landed (the observed evidence layer thickened):** the shared detachable descriptor-rooted FS worker + per-lane breaker generalized (§4a seam) · `find-references` E1 on it (declared textual domain, numeric ceilings, fail-closed incomplete-scan, lexical-not-semantic label) · the §6.1 conductor verdict pass (identity from Selection, total tuple matrix, output redaction, re-derived MachineryFault, derived `signal_class` rows — **CheckVerdict shape untouched, the F1 avoid-amendment ruling**) · the **master-activated `lane_vcs:none` opaque-accept branch** (branch-only token scope; git/nil NEVER accepts; the total input table; degradation labeling; **the locked Option-2 E0 floor restored**; the serialized-path FS probe removed) · m-7's E1–E10 `lane_vcs` v3 owner bytes carried verbatim · the FX-VCS v3/v2 transition matrix · the attestation negative (lane-cannot-forge-operator; positive deferred to B4/item-9) · the runnable exit-fixture set + the ⑤ ODB egress pair · MF-1 evidence-table fold. B-opaque built + fixture-proven both ways; the blocked ledger governance-only, no `t.Skip` stubs.
**The build's own record:** a FRESH `s9.planner`/`s9.implementer` pair (the Option-A staffing correction — the mis-seated same-owner token was voided before harm), m-3 as guide + owner-fidelity, B11 straight-through, ONE end-of-slice adversarial review (code APPROVE, zero blockers), two owner confirms clean (m-7 blob-verified, m-3 byte-verified), AO-1=CARRY.
**s9's catch-ledger entries (condition g, one series):** (1) **the mis-seated staffing token — LINTED CLEAN, caught only by the operator's read** — the standing "the trail checks lineage/authority SHAPE, not staffing SEMANTICS" limit, and the honest "what the gates do NOT catch" line for the pitch; (2) the T3 totality gap — caught TYPED + PRE-COMMIT by the fail-closed tuple collapse + the wide regression battery (the machinery catching its own plan's false claim); (3) **the fence anatomy** — the mid-build `lane_vcs` activation introduced test seams no fence had named (my activation ruling didn't enumerate them · the routing asserted in-fence from owner locus-naming · the build's table licensed at task grain) — **caught by the mechanical diff→block table PRE-MERGE, exactly as designed**; discipline refinement recorded (a mid-build amendment activation must reconcile its test/fixture seams into the fence; an owner return naming un-fenced loci is an escalation trigger, not an in-fence assertion). **And the standing reconciliation value:** SIX over-scopings caught this build at the byte-grain domain boundary (mine ×3–4, m-2's token-granularity improvement, m-7's consumer-contract reach into m-3's observation) — whoever opines past their owned surface is reconciled back to it.
**Carries into master's step-exit queue (named at close, not s9 failures):** T7/T8 (report-and-hold — the m-2/m-1 co-signs didn't arrive in-slice, no code shadow) · C-1 (the pre-existing `git status` exec on the serialized commit loop — the one un-detached path; natural home item-10) · R-1 (hardlink confinement narrowed to "beneath the root by path, not by inode provenance", D5-class, labeled §4a/§11) · AO-1 (a bounded count/saturation field when a real consumer exists) · the blocked ledger (B1 scope_paths — **the §F m-1 leg STILL pending** — · B2 · B3 · B4, governance-only).
**Register effects:** the s9 evidence layer is thickened and merged; **the ONLY remaining Step-2 build slice is s11** (the comms thicken, serialized after s9) → then the step-exit test (all 3 legs live on the dogfood store + the INV-CATALOG red-battery demo + uncached green) closes Step-2.

## Step-2 · Slice-11 (s11 = THE COMMS THICKEN) — FIFTH STEP-2 BUILD-CYCLE CLOSE **and the LAST Step-2 build slice** (CLOSED 2026-07-14; operator-granted merge with the recognized field at grant time — the convention's FIFTH clean run; a FRESH Option-A slice pair; four-lens end review unanimous + batched STOP-ON-CONTRADICTION owner confirms; master's independent third-seat battery at the merged candidate)
**Close:** tag `s11-close` (annotated obj `0462c24`) at `main@502e06c`, pushed frank-dev (main + tag remote-verified; PR #1 merged=true). Fast-forward `d91fcfb..502e06c`, 18 linear commits. Merge executed by s11.implementer on the operator-granted bounded dispatch (`s11-merge-gate/MERGE-GATE-orchestrator-planner-20260714-175210`, operator in-session "write me a dispatch merge"), reported `…-181240`; master's own uncached serialized battery green at `502e06c` (26 ok / 0 FAIL / vet clean / ten INV-CATALOG laws / fixtures ~130s), fold verified spec-exact.
**Surface merged:** the B/C/D bucket projections (B non-interrupting/raise-only · C CC-FYI/no-obligation · D author-return with egress/D-vs-A precedence) · the complete 7-state FSM (`bounced_repair` live; `egress_blocked` local-park-resummon only, away send unbuilt per step-(d)) · **the full g1 §B 8a hardening** (both reason tokens byte-exact — `stale_schema` A/`held` migration-fault · `stale_choice_set` D/`rejected` author-return — the frozen-π guard fail-closed, no-wake, new decision identity, real-process crash-replay to the same replacement) · the 14-row bucket/terminal/edge matrix + the ③ known-A NF fixture · the T8 cleanup (**eight of nine**) · the G4 resummon cadence on operator-config, re-homed at **engine v4** with the day-one v3→v4 migration path fixture-exercised.
**The three gates:** **g1 (8a) CLOSED → T6 built** on the integrated three-member contract (`s11-8a-joint-review/RECONCILE-orchestrator-planner-20260714-035210`; the bucket-D token ruled `stale_choice_set`, registration settled at the bytes — `failing_edge` is `system_only` text with no `enum_set`, so an open stamped value is the only shape). **g2 (OQ-2 fork ceiling) → T5 HELD OPEN** (m-5.implementer review + completion never returned). **dc (re-prompt/claimless-`held`) → T10 HELD OPEN** (the m-3+m-6 design cell never returned). T5/T10 are acceptance-OPEN per FINDING-4 — the merge makes no completion claim; their disposition is master's step-exit call.
**Two mid-build master rulings (each a real boundary event, ruled at the bytes):**
- **The T8 fence split** (`s11-build-escalate-fence/RECONCILE-orchestrator-planner-20260714-143010`): item 8 (`finalizeRun` preserve-flag) GRANTED as a T8-named `executor.go` seam independent of g2 (the pair's r1 fence had bound a standing cross-cutting root to a gated task — a self-inflicted contradiction, corrected); item 2 (shared soft-expiry arbiter, executor/read-file) DEFERRED OUT of T8 by explicit master rescope — read at the bytes, the two paths' drain + termination semantics DIVERGE (executor joins-with-SIGKILL + preserve-flag; fs_worker trips-a-breaker + DETACHED worker = m-3's §4a D-state residual), so a "shared arbiter" is a race-semantics unification, not a behavior-preserving refactor (Rail B: dedup is a maintainability good, not a safety one). T8 acceptance amended to eight of nine; `internal/observe/` stays OUT of the s11 fence; the arbiter is a named post-Step-2 m-7+m-3 design-cell carry.
- **The T9 config-lock contradiction** (`s11-build-escalate-config-lock/RECONCILE-orchestrator-planner-20260714-170510`): the `resummon_cadence` key landed on m-7's r13-locked engine-v3 schema surface with no version bump and no owner routing (an r13 + condition-(c) violation). **Owner-caught** by the batched confirm after three plan-review rounds and a four-lens panel all missed it. Ruled **(a) re-home at engine v4 via the adjacent hop, folded on-branch** — NOT (b) retroactive descriptor widening, which would make "engine v3" name two admitted-key-sets in the wild (the exact confusion the version marker prevents, and antithetical to the day-one public versioning+migration promise). The v4 fold RESTORES v3's honest meaning (v3 uniformly rejects the key; v4 uniformly accepts it), costs marginal (zero live v3 stores ⇒ descriptor arm + adjacent-forward transition + fixtures, no data migration), and m-7 countersigned the realized v4 bytes hunk-by-hunk ("engine v3 once again names exactly one schema surface; the owner path has been paid in full"). Ruling (a) is r13 CONFORMANCE — no PROTOCOL-DEVIATIONS entry.
**s11's catch-ledger (condition g, one series):** #1 pre-code (the r1 executor-license fence contradiction — pair-caught, master-ruled) · #2 pre-code (the r2 locus-table truth errors — pair-caught at the bytes) · #3 in-build (evidence-prose ordering — self-caught) · **#4 post-build, OWNER-CAUGHT (the r13 silent landing — ruled, folded, countersigned before merge).** The cycle-datum pair is now complete: *the end review catches what a green battery structurally cannot (s10); the owner confirm catches what the review panel cannot (s11).* **Binding discipline refinement (standing):** a fence row licenses a FILE; it never substitutes for the OWNER PATH on a locked contract living inside it — the PLAN must name the lock and route the owner countersign BEFORE the edit.
**Carries into master's step-exit queue (named at close, not s11 failures):** T5/T10 (gate-bound behind g2/dc — a small gate-bound completion leg after the returns, or an explicit rescope, at the step-exit fold) · the item-2 soft-expiry-arbiter design-cell carry (m-7+m-3 post-Step-2, the drain/termination divergence its first question) · the s9 carry queue still open (T7/T8 · C-1 · R-1 · AO-1 · **B1 scope_paths §F m-1 leg still pending**).
**Register effects:** s11 merged; **the Step-2 BUILD is COMPLETE — all five slices closed (s7 · s8 · s10 · s9 · s11).** The ONLY remaining Step-2 act is the master-owed **step-exit test** (all three legs live on the dogfood store + the INV-CATALOG red-battery demo + uncached green) — on its pass, Step-2 closes.

## STEP-2 CLOSE — OPERATOR-RATIFIED (2026-07-14, in-session): the phase is CLOSED; the charter phase line flips to Step-3
The master-owed step-exit test PASSED (`step2-step-exit/RECONCILE-orchestrator-planner-20260714-182600`; three legs green live on the dogfood store + the INV-CATALOG red-battery demo naming `TestLawTerminalEnumByteExact` + uncached green battery 26 ok/0 FAIL at `main@502e06c`). **The operator RATIFIED the Step-2 close in-session (2026-07-14)** — channel-stamped via current-session context per the standing convention (operator = final authority). Actions on ratification: the charter phase line (CLAUDE.md rule 1) flips **Step-2 BUILD → Step-3 OPENING**; the ROADMAP + dashboard mark Step-2 CLOSED. **The Step-2 BUILD delivered all five slices (s7·s8·s10·s9·s11): unfakeable gates (observe-as-send + evidence ladder), self-pacing lanes (park/wake + buckets), and the constitution as ten executable laws.**
**One outstanding formal leg:** the **VP adversarial close-confirm** (named in the step-exit relay TO `master.orchestrator-reviewer`) has not yet returned; the operator elected to proceed. If the VP surfaces findings, they are handled as an amendment — the step-exit PASS + the operator ratification stand.
**T5/T10 disposition — RESCOPED as Step-3 carries** (operator-ratified recommendation): both stayed acceptance-OPEN behind un-returned gates (T5/g2 = m-5.implementer OQ-2 review; T10/dc = the m-3+m-6 design cell); they are additive surfaces beyond the minimal-comms floor the exit certifies, and carry into Step-3 (a bounded gate-bound completion leg when g2/dc return, or an explicit rescope).
**OPERATING MODEL for Step-3+ (operator-decided 2026-07-14):** the team RUNS ON frank as courier; an **independent slice orchestrator team per slice/bundle** owns local detail-design + plan + impl against the **m-x-authored spec-of-record** (m-x = PMs), escalating UP through master to the m-x planners on (a) a spec mistake · (b) a better way · the standing `DELEGATED_DISPATCH_AUTHORITY` triggers. This is the deferred **T4 tier** (`master-org-decomp`, 2026-06-28) activated now that frank automates the relaying — the exact precondition on which it was shelved. First Step-3 acts: the operating-model spec + the frank live relaunch + shakedown; build slices open on frank after.
**Step-3 carry queue:** T5/T10 (g2/dc) · the item-2 soft-expiry arbiter design cell (m-7+m-3) · the s9 queue (T7/T8 · C-1 · R-1 · AO-1 · B1 scope_paths §F m-1 pending) · the operating-model / nested-run conventions design (the tier authority boundary; the up-escalation relay convention).

## STEP-3 OPENED — the kickoff LOCKED + Division II stood up (2026-07-14)
The Step-3 kickoff-of-record (`master/STEP-3-KICKOFF.md`) is **LOCKED**: operator-ratified (org expansion + vertical-first scope + kickoff, in-session) + **VP-co-signed** (`step3-prep/RECONCILE-orchestrator-reviewer-20260714-222000`, approve at SHA `cf480081…`) after **FOUR adversarial rounds** — r1 six directional folds → r2 four contract-consistency corrections (against the locked bytes) → r3 two stale-clause fixes → r4 one review-count fix. The convergence (6→4→2→1) is the design-churn-is-precision discipline (B12) working as intended.
**What locked:** frank owns the normative provider contract (pi/opencode = prior-art + conformance fixtures only) · m-8/m-9 single-writer boundaries + design-only boot + the full consumer-lock set · the four-axis lane key + pinning + the **unconditional** m-4/m-2 routing-record amendment · the m-3 provider-request-egress + m-7 credential amendments (all three mandatory, sequenced pre-lock with consumer reviews) · **vertical-first V1→V2→V3** with an **E3 live independently-bound-provider evidence floor** + named zero-send/zero-exec/no-fallback negatives · the T4-token gate (full frank roster + `T4→master→m-x→master→T4` round-trip + Part-F mechanics + durable export) · the spine explicitly **non-terminal** for Step-3.
**Executed on co-sign:** charter amended (org chart ×2 rows + domains table m-8/m-9); domain charters authored (`master/domains/m-8-provider-adapters/README.md`, `…/m-9-model-runtime/README.md`); **design-only AUDIT boots issued** to both greenfield pairs (`boot/master-boot-m-8/…-223000`, `boot/master-boot-m-9/…-223010`).
**The honest shape of Step-3's opening (the VP's real finding):** the conductor's locked contracts do NOT yet support governed provider-send — so Step-3 opens with **three mandatory owner amendments** (m-3 provider-egress · m-7 credential/config · m-4/m-2 routing-record) closing real gaps, THEN the m-8/m-9 design cycle (audit → design → GRILL_LOCK → consumer-review → reconcile → lock), THEN a vertical T4 build with an E3 live floor. No design-lock/PLAN/T4-code/credential/external-call is authorized yet.
**Register:** Step-2 CLOSED; Step-3 DESIGN PHASE OPEN. The greenfield m-8/m-9 pairs are DESIGN-only until their domains lock; they become T4-PMs only after (Part F). The Step-3 carry ledger (benchmark/spawn/steer + T5/T10 + the s9 queue) rides per kickoff §7.

## STEP-3 ARCHITECTURE REFRAME — ratified; the conductor is one service in a larger app shell (2026-07-15)
Mid-Step-3-DESIGN, an **operator architecture-of-record correction** re-cut the topology: **the conductor is ONE isolated governed-relay service, NOT the app's central hub**; the app is a **modular monolith + supervised workers** (not networked microservices) with the conductor as one component. Surfaced by the operator (the m-x lanes had drifted toward "add HTTPS to the conductor"), VP-dispositioned `human-decision-required` (`step3-arch-reframe/013000`), operator-ratified across four framing questions + a one-at-a-time grill, VP-approved after a six-round packet review.
**Architecture-of-record:** `master/STEP-3-ARCH-AMENDMENT.md`, **operator-ratified 2026-07-15 at SHA-256 `2d240eb6…`** (`step3-arch-packet/070000`), VP-approved (`…/063000`, VERDICT approve). It is standalone, **outside the locked kickoff bytes**; the reframed `STEP-3-KICKOFF.md` **§§1–3 AND §§5–8** are **SUPERSEDED** (V2/V3→Step-4, amendment set re-cut, spine-non-terminal re-cut; V1/E3/T4-gate survive; old kickoff hash `983508fc…` preserved as the historical lock).
**What ratified:** conductor = governed relay plane for stamped participants (seats, orchestrators, the operator channel, system records) + own store/writer, Steps 1–2 unchanged · app shell = **m-10 App Control Plane/Supervisor (NEW domain)** + **m-8** connector (app-side, holds provider creds + does the wire) + **m-9** worker (app-side runtime) · provider traffic + credentials + tool payloads + PTY + run state NEVER transit the conductor · **Step-3 = the one-governed-turn MVP** (single pinned **app-side run manifest**, no routing execution, live E3; m-4 routing + benchmark/spawn/steer defer to Step-4) · provider-send **mechanism** app-side, **policy stays m-3**; app-side send = **E0 self_reported attestation** carried in an existing `SITREP` body (no conductor change) · the **m-5 ceiling-host amendment** (m-5 sole policy owner, m-10 enforcement host only) · the **operator direct route** = authority-bearing by construction, non-transitive, no forced relay-authoring, cross-seat effects via the landed typed-grant grammar.
**The review chain (design-churn-is-precision, B12):** `030000` (F1–F10 structural must-revise) → `043000` (F11–F14 authority-mechanism/m-5-host/E0-carrier) → `053000` (F15–F17 source reconciliation) → `063000` **approve**. Convergence 10→4→3→0.
**Five held lanes:** `step3-design-m-8`/`-m-9` re-dispatch app-side; `step3-amend-m3-egress` re-dispatch (m-3 keeps policy); `step3-amend-m7-cred` RE-OWNER (connector-side, m-1 boundary; m-7 r3 = provisional audit input, m-7 conductor-host scope untouched); `step3-amend-m4-routing` DEFER to Step-4. All stopped clean (no lock/PLAN/code, bytes preserved).
**Register:** the reframe is the architecture-of-record. Source fold **F27–F29 folded + F31 dispositioned** (`step3-arch-packet/134000`→VP, folding VP `100000`/`101000`/`123753` revise). **First-stage live state (precise):** the m-5 amendment **proposal bytes are pinned + pair-approved but NON-CONSUMABLE** — canonical contract @ `643dd7c2…`, GRILL_LOCK closed, report-only; **m-10 has hash-confirmed those bytes in COORD `091500` + received the clarification `092000`, but has returned NO DESIGN/GRILL_LOCK/implementer review/completion SITREP** — COORD/hash convergence is NOT an approved m-10 artifact. **MVP SCOPE CONVERGED (operator, 2026-07-15) → `step3-arch-packet/152000`.** The operator re-cut the MVP to a **barely-enough coding agent on the governed courier** and **deferred the entire permission/authority system to Step-4**. So the whole Branch-B thread above is **superseded**: there is **no config-derived ceiling in the MVP**, hence nothing to be fresh/stale about — the **seam-13 knot dissolves**, and the **m-5 amendment STANDS DOWN** (`step3-amend-m5-ceiling/152000` withdraws `145500`; the VP F36–F38 `150756` corrections to it are OBE). **MVP = local tools `read/write/edit/bash/apply_patch` (m-9, app-side) + the conductor as a NATIVE relay tool** (`submit/project/read` over `internal/channel`; the **MCP server retained** for foreign harnesses; a **shared conductor-client refactor** de-dupes both frontends) + a **built-but-EMPTY permission seam** (authorization = a trivial static run-manifest allow-list; m-10's enforcement point present, permissive default) + a **push-based wake-on-relay stretch**. `bash` = ambient host authority, operator-accepted (no sandbox; trusted executor = Step-4 H-12). **m-5's ceiling `643dd7c2…` (untouched) + `config_generation` freshness + per-role permissions + the registry/manifest/carousel are the STEP-4 basis.** First-stage now reduces to: **(1)** the **m-10 DESIGN** (empty seam + run manifest + worker supervision + wake loop, consuming the trivial allow-list) → implementer review → SITREP · **(2)** the **m-9 DESIGN** (coding-agent worker + native conductor tool + shared-client extraction) · **(3)** the Master+VP first-stage interface-lock. Full spec: `master/ARCHITECTURE.md` "Step-3 MVP" §. Five holds stand. No design-lock/PLAN/T4-code/credential/provider-call authorized.

## STEP-3 MVP AMENDMENT — RATIFIED (2026-07-16): the operative MVP architecture; the first-stage DESIGNs dispatch per its §7
The explicit MVP architecture amendment (`master/STEP-3-MVP-AMENDMENT.md`) is **OPERATIVE**: **operator-ratified 2026-07-16 at SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d` (r7)** — in-session direct-channel ratification in direct response to the presented hash, recorded `step3-arch-packet/RECONCILE-orchestrator-planner-20260716-040405` per the §8b convention — after the **VP byte-bound APPROVE** (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-035505`, fold-review r15). It amends the ratified reframe packet `2d240eb6…` at exactly **four §1-named fragments** (the packet file stays byte-exact as the historical lock): the Sequence-A ceiling clause (→ the fixed tool-dispatch set), the first-stage order (→ the §7 graph), the `:29` credential-address-space phrase (→ the F57 narrow non-injection claim), and the `:27` worker-as-principal phrase (→ the F66 logical-seat model).
**The review chain (design-churn-is-precision):** r0 `a524bcbf` (F39–F44) → r1 `02e9da1c` (F45–F51) → r2 `3db3eb96` (F52–F55) → r3 `e25bce10` (F56; VP-approved, voided by design at r4) → r4 `57aa3170` (operator-directed external-review errata) → r5 `e47d514d` (F57–F62 + the grill) → r6 `5d66bf24` (F63/F64) → **r7 `2f75f2a1` (F65/F66; VP APPROVE)**. Convergence 6→2→2→0 across F57–F66.
**The three-decision amendment-level grill (operator-locked 2026-07-16, durable records):** **#1 process topology** (`…-023557`): conductor = own service · **m-10 = a MODULE in the app main process** (not a separate daemon; manifest/ticket/epoch/IPC seams designed as-if process-separated) · m-9 = the supervised worker process · m-8 = a separate connector process; field evidence from `references/claude-code` adjudicated at principle level (desktop = monolith; the hosted-container pattern validates the m-8 chokepoint + the F57 narrowing; no field harness runs a local supervisor process). **#2 F59 = Option B** (`…-024350`): the durable ONE-SHOT authorization ticket `{run_id, turn_id, turn_epoch, tool_call_id, canonical_tool_name, canonical_args_digest}`, exactly-once atomic consume, crash windows park-not-replay, m-10 authors the protocol / m-9 the executor half, actual-invocation acceptance proof. **#3 F60** (`…-025642`): **one broker-held credential per LOGICAL seat**, never copied into worker generations; `turn_epoch`-fenced replacement, NO implicit new identity; m-1 authors identity/credential-lifecycle, m-7 the channel/broker contract (placement outside the replaceable worker generation, F66); worker-per-seat rejected (mint-before-serve makes per-generation minting an operator act).
**Other operative rulings landed r4–r7:** the honest §2 credential boundary (narrow non-injection; same-user inspection an explicit unsandboxed residual; OS boundary = Step-4 H-12) · §2a turn-vs-attempt (one attempt per provider INVOCATION, no auto-retry anywhere in the stack; a turn holds multiple recorded attempts; compiled bounds) · §3 object-typed negative route + the exact product claim + evidence scoped to the two app-side/external artifacts (E0 worker-carried report — "attestation" avoided; EXTERNAL E3 with named writer/digest-producers/m-3 applicability evaluator; **F65: the release/E3 vector = the app/provider vertical ONLY; the conductor service identity is bound separately in the exit-test record for the relay-exchange leg**) · §4 the tool-DISPATCH seam (the operator-ratified **8 canonical NAMES** = policy identity; build identity = per-tool vectors with named producers, bound at the interface-lock + a **post-build RELEASE-BINDING event** before the live E3, F63; set-equality over identity; fail-closed deny) · §5 the shared conductor-client 3-way seam (m-7 transport / m-2 mapping / m-9 consumer; MCP retained off the critical path) · §6 wake = best-effort advisory push + durable rediscovery + at-most-once `UNIQUE(relay_id)` scheduling (no at-least-once claim, F61) · §7 the acyclic single-owner graph + named DESIGN requirements (seat topology RESOLVED; m-10 durable state + monotonic `turn_epoch` fencing + UNKNOWN/PARTIAL park-not-replay) · §10 the 14-row acceptance annex + sentinel caveat + build order.
**Executed on ratification (this fold):** the §7 graph + §1 fragment-supersessions reconciled into `ARCHITECTURE.md` (the Step-3 MVP § + the matrices), `ROADMAP.md` Step 3, `CLAUDE.md` rule 1 + domain table, and the m-1/m-2/m-3/m-5/m-7/m-8/m-9/m-10 charters (m-9 carries the F66 logical-seat wording verbatim; m-4/m-6 verified no-contradiction, no delta owed); the README dashboard flipped to RATIFIED/OPERATIVE; the ordered 15-file manifest recomputed (the fold relay carries the hashes). **The five §7 stage-1 DESIGN dispatches issued** (m-2 mapping · m-7 transport+broker · m-10 IPC/manifest-seam · m-1 secret-boundary+seat-identity · m-3 egress/E0/E3-evaluator); stage 2 (m-8) dispatches when stage-1 artifacts land; stages 4–5 (m-9/m-10 designs) follow their inputs; stage 6 = the Master+VP first-stage interface-lock; then the post-build release-binding event precedes the live E3 + exit gate.
**Register:** Step-3 DESIGN is now governed by the packet **as amended** (`2d240eb6…` + `2f75f2a1…`). No DESIGN-lock, PLAN, T4 code token, credential provisioning, provider call, release-binding execution, merge, or deploy is authorized by the ratification itself — those stay behind the §7 stage gates.

## §7 stages 1–3 close (2026-07-19)

**VP close-confirm ISSUED** (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260719-224500`, close-review r6) over the corrected close supplement (`…-223821`). Six VP review rounds (`021858` F70–F74 · `065204` F75–F79 · `071626` F80–F81 · `191718` F82–F83 · `215549` F84 · `224500` approve); every finding verified real and folded owner-real; the last three were escalating-grain interface defects on the F59 seam (census → family → field → value-source/timing). The seven closed hashes: m-1 `7c8b09a6…` · m-2 `83d8e63e…` · m-3 r4 `009df607…` · m-7 r11 `9331ea88…` · m-10 r36 `0240e874…` · m-8 r12 `4b670a79…` · m-9 r19 `2a96a07b…`. Evidence: the 16-edge/13-carrier census · the three-identity reciprocal (`confirm-m10/223500`) · the stage-2 four-verdict chain. N1–N4 = permanent lock-record errata. Next: stage-4/5 dispatches (issued same day) → stage-6 Master+VP interface-lock.
