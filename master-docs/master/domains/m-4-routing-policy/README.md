# m-4 — Routing & Policy

**Pair:** design-lead `m-4.planner` + adversarial design-reviewer `m-4.implementer`.
**Cycle c2 scope:** AUDIT + DESIGN of the routing/policy domain (the adaptive-routing pillar), against the
locked c1 contract. Terminal = design-lock. No impl / PLAN.

## Owns (the conductor's thesis — adaptive routing)
The **model→seat router**, **capability priors**, the **routing record**, **justified deviation**, and the
**benchmark + later-release feedback loop**. Locked high-level decisions this builds on: altitude **B** (planner emits
role+model per dispatch), policy **3-staged** (capability priors + justified deviation; benchmark feedback in
a later release). The differentiator: routing is a first-class, recorded, justifiable governance decision — not an
implicit config.

## Boundaries
- **Consumes (locked c1):** m-2 schema + m-1 identity. c1 R2 already specified the seam: a routing decision is a
  **separate seat-stamped routing relay** (m-2 FieldSpec + record-kind + accepted semantics; m-1 admits the
  accepted routing relay into the conductor-derived `parent_picker` candidate set). The dispatch references it as
  **provenance/bookkeeping** — **model is never a gate input** (no `model_*` predicate enters the schema gate).
- **identity ≠ authority (ratified c1 §5):** m-1 owns *who* (the stamp); **m-4 owns *what a stamped seat may
  do*** — routing/policy keyed to the stamp (anti-confused-deputy). m-4 design realizes this boundary.
- **Co-design seam with m-3 (this cycle):** the benchmark / later-release loop **consumes m-3 observed evidence** for
  routing quality. Shared `c2-*-coord` COORD sub-thread; reconcile before any c2 lock.
- **Consumed by (future):** m-5 (archetype tags ↔ routing — explicit **lock-time seam disposition** required
  this cycle, per VP c2-decomp Finding 4), m-6 (gate→email buckets consume routing categories; `routing` is a
  category-B `gate_category` per ratified §J).
- **Does not re-open m-1/m-2.** Audit questions focus on the locked contract, not the foundation.

## Status
- AUDIT: **reconciled** `c2-audit-m-4` (20260629). Independent planner + implementer passes + RECONCILE filed
  (`relays/c2-audit-m-4/`). Verdict **still-open** (the governance-record primitive; carrier + selection-runtime
  + priors already exist). Pair recommends **PROCEED-TO-DESIGN** with two seam conditions. Standing by for the
  orchestrator's PROCEED-TO-DESIGN. Honest qualifications carried for DESIGN (do not overclaim): Routesplain /
  Arch-Router = nearest interpretable-routing prior art (differentiate on the persisted deviation-against-a-floor
  audit artifact); non-gradient bandit feedback already exists (the later-release novelty = the auditable persisted
  decision+update, not the mechanism). The routing record = a port of the SR-26-2 model-risk override-register
  discipline into per-dispatch LLM routing.
- DESIGN: **COMPLETE / held for c2 lock** `c2-design-m-4` (20260629). DESIGN_DOC_ID
  `c2-design-m-4-routing-policy` (`design/2026-06-29-routing-policy-design.md`), operator-approved +
  **m-4.implementer DESIGN-REVIEW approve** on rev1 (`c2-design-m-4/DESIGN-REVIEW-implementer-20260629-203329.md`;
  one must-revise round — F1 precise R2 invariant/bucket-vs-bucket form, F2 reason-code required_when, F3 template
  no-bypass — folded + carry-forward normalized). Design-complete reported
  (`c2-design-m-4/SITREP-planner-20260629-203900.md`). m-3↔m-4 seam **reconciled both sides**
  (`c2-design-m3-m4-coord`). Operator grill-locks (§0): GL-1 two-layer bucket priors; GL-2 a later release tunes
  recommendations only; GL-3 record-now/execute-Step-3; GL-4 the initial release ships 1–3 routing templates (m-4 owns
  model-assignment mechanism; m-5 owns structure+lineup — scope flagged
  `c2-design-m-4/SITREP-planner-20260629-200500.md`). **Open at the c2 lock (CTO/VP):** m-5 archetype-ceiling/tag
  + template-structure disposition (surface, not closed); R2-boundary ratification (both pairs aligned). Holding
  for orchestrator direction. No PLAN this phase.
- DESIGN c2-FOLD: **rev2 applied** `c2-fold-m-4` (20260630) — folded VP-approved m-5/m-6 consumer findings
  (bounded-additive): F2 per-assignment opaque `seat_archetype` + resolved `authority_ceiling` on
  `routing_assignments` (replay on template + hand-authored paths); F3 `seat_archetype` = distinct opaque tag in
  the archetype vector (per-seat-at-spawn, orthogonal to m-3 `slot_in`); **M4-1 CONFIRMED** — routing B→A
  escalation rides the existing c1 monotonic HUMAN_GATE "m-4 routing-raise" (named in m-2 §3) + §J A-set,
  readable on the consumable record, **no new gate class**, R2-safe. Guardrail held (no m-2 change; concrete
  semantics m-5-owned c3). Re-review request → m-4.implementer (`c2-fold-m-4/DESIGN-planner-20260630-040400.md`);
  M4-1 confirm reported to orchestrator (`…/SITREP-planner-20260630-040430.md`).
- DESIGN c2-FOLD: **COMPLETE / pair-approved** — m-4.implementer re-approved the rev2 fold
  (`c2-fold-m-4/DESIGN-REVIEW-implementer-20260630-040641.md`); fold-complete reported
  (`c2-fold-m-4/SITREP-planner-20260630-041100.md`). **m-4 design-locked-ready / holding for the c2 lock.**
  Remaining (CTO/VP, not m-4): R2-boundary ratification (both pairs aligned) + m-5 concrete-semantics reservation
  (c3). Watch-note (implementer): future work making `seat_archetype` authority-bearing outside the routing
  mechanism / concrete-valued / predicated routes as a new m-2-adjacent fold. No PLAN this phase.

- READINESS c4 (m-2 schema bring-current, m-4 CC-review): **reviewed / confirm-with-required-retype**
  `readiness-fix-c4` (20260630). Orchestrator dispatched m-2 to bring its routing schema current with the c2 R2
  lock (Cluster 4a/4b); m-4 confirmed on CC (`readiness-fix-c4/DESIGN-planner-20260630-231200.md`). **4b**
  (grammar allowlist excluding model-identity fields) + **the 4a trigger direction** (`required_when
  declared_deviated == true`, dropping `selected_model`) **confirmed as-is**. **One MUST-RETYPE:** m-2 typed
  `declared_deviated` as `owner:system`/`computed_result` ("courier-computed bucket-vs-bucket") — that conflates
  the **planner-declared** gate bit (`declared_deviated`, `agent_enum_pick`) with the **separate, m-3-owned,
  observe-side** `deviated_observed` (the system bucket-vs-bucket computation, never gate-referenced). R2 not
  regressed, but the collapse breaks the `(false,true)` silent-deviation integrity veto (needs an independent
  declared bit). Retype = planner-declared `declared_deviated`; keep `deviated_observed` out of the gate; grain =
  `any(routing_assignments.declared_deviated == true)`. **m-2 folded the retype exactly**
  (`readiness-fix-c4/SITREP-planner-20260630-231506.md`): `declared_deviated` → `agent_enum_pick`/planner-declared,
  `deviated_observed` split out as m-3-owned observe-side (not gate-referenceable), aggregate `any(...)` trigger, +
  a bounded `any_row:` existential atom (R2-safe — same allowlist, can't name a model-identity field). **m-4
  verified line-by-line against the schema; m-4-side acceptance MET / Cluster 4a/4b closed on the m-4
  deviation-gate contract** (`readiness-fix-c4/SITREP-planner-20260630-232000.md`). No open m-4 finding on 4a.
  **Then (VP revise `232925`, operator "complete the fold"):** m-2 broadened to **Cluster 4 completion (4b+4c)** —
  mirrored its routing FieldSpec to m-4's **full** locked record `:200-210` + declared conductor-computed field
  homes (`readiness-fix-c4/SITREP-planner-20260630-234702.md`). **m-4 reviewed the mirror
  (`readiness-fix-c4/DESIGN-planner-20260630-235500.md`): `confirm-with-one-carve-out`.** Mirror **faithful/complete**
  for every m-4 field (per-row `routing_assignments` shape, `deviation_reason_code` grain, reserved
  `constraints`/`template_ref`, snapshot, `outcome_feedback_ref`, record-kind; R2 held). **One carve-out:** m-2 added
  **`human_mode`/posture** as a per-row `routing_assignments` field — NOT in my locked `:203` and NOT in
  `ARCHITECTURE` C2.4 F2 `:187-188` (which homes only `seat_archetype`+`authority_ceiling`). A mirror must follow the
  record, not add to it. Not an R2 issue (posture isn't model-identity, opaque/no Step-1 values) — a record-home
  issue. Disposition routed to m-5+CTO: (1) posture stays m-5/m-6-owned off the routing record → m-2 drops it from
  the mirror, or (2) team wants it recorded per-assignment → m-4 folds it into §5 after m-5 ownership-confirm + CTO
  ARCHITECTURE ratify. Rest of mirror m-4-confirmed. **CTO arbitrated the carve-out**
  (`readiness-fix-c4/DESIGN-orchestrator-planner-20260701-001537.md`) = **option 1 (my c3-lock-faithful disposition):**
  posture rides `seat_archetype` ("no new m-2 field" per m-5 c3 lock `:142`); my record `:203` stays un-extended; no
  ARCHITECTURE change; option 2 explicitly rejected. m-2 folded it (standalone posture field removed; value-enum kept
  in §17.6 as m-5 vocab). **m-4 confirmed the ruling** (`readiness-fix-c4/DESIGN-planner-20260701-013000.md`): verified
  m-2's per-row set now = `:203` exactly, posture removed, R2 held → **`confirm`; m-4 side of Cluster 4 closed, no open
  finding.** Row-parity (§15 Q-F) tracked as a pre-Step-1-PLAN SHOULD, orthogonal to the routing record. Holding for
  m-5/m-6 ruling-confirms + CTO re-verify + VP closure co-sign → Cluster 4 CLOSED. No PLAN.

- CQ-CLOSURE c4-cq-gateconfig (m-7 conductor-core gate CQs; m-4 owns **CQ-4b only** — trusted-config composition):
  **planner-lead confirm produced, in full-pair review** (20260702). m-4 was re-engaged (VP added m-4 to the CQ-4b
  config surface) to confirm the CTO's trusted-config ruling (per-domain-authored sections → one conductor-composed
  artifact under a single top-level digest, restart-only load, change = operator-authorized committed store record)
  preserves the **capability-prior config** assumptions. m-4 verdict **confirm-with-one-correction**
  (`c4-cq-gateconfig/DESIGN-planner-20260702-013000.md`): **preserves the initial release cleanly + beneficial** (single digest = a
  clean config-version stamp for `capability_prior_snapshot` per §4:187; operator-authorized committed-record change =
  exactly the Layer-1 membership model); **one minimal correction** — the envelope must permit a **per-section version
  stamp for the Layer-2 recommendation table** so later-release auto-tuning (GL-2; §6 Stage-3 "bolts on without re-cutting the
  record/gate") isn't precluded by "operator-authorized full-artifact recompute." Layer-1 stays top-level-digest-bound.
  CQ-4b (m-4) → **corrected-by-artifact**; later-release auto-tune cadence → non-locking-carry; CQ-2/3/4 not m-4's. Routed to
  m-4.implementer for independent review (directed full-pair rigor); co-signs into the thread on approve.
  **m-4.implementer returned `must-revise`** (`c4-cq-gateconfig/DESIGN-REVIEW-implementer-20260702-020443.md`) — a
  correct catch: rev0's Layer-2 correction over-claimed (a "machine-cadence sub-record write lane / distinct
  authorization path" is NOT locked by §6 Stage-3, which only says the *record/gate schema* bolts on, not the
  config-write cadence; and `capability_prior_snapshot` already meets attribution) — that would have introduced new
  cross-domain config-authority design, out of scope. **rev1 folds it exactly**
  (`c4-cq-gateconfig/DESIGN-planner-20260702-021000.md`): the initial-release confirm unchanged; Layer-2 ask narrowed to an
  **optional per-section `version` reservation inside the single top-level digest** (no initial-release semantics; every effective
  initial-release change still recomputes the top-level digest + operator-authorized reload); later-release auto-tune cadence/authorization
  = **still-open/non-locking-carry** (new design if a later release is pursued). CQ-4b (m-4) → corrected-by-artifact, co-signable.
  Reissued to m-4.implementer for re-review. **m-4.implementer APPROVED rev1**
  (`c4-cq-gateconfig/DESIGN-REVIEW-implementer-20260702-021613.md`) — CQ-4b m-4 portion **pair-approved /
  corrected-by-artifact**, co-signable; full-pair arc complete (confirm → must-revise → rev1 → approve). Co-signed
  closure reported to orchestrator (`c4-cq-gateconfig/SITREP-planner-20260702-022000.md`). Joint CQ-4b still needs
  m-2+m-3+m-6 + CTO fold (not m-4's). m-4 holding. No PLAN.

- CLAIM-SWEEP c5-claim-sweep-light (claim-text hygiene, batched m-3/m-4/m-5/m-6): **m-4 folded, → implementer approve**
  (20260702). Applied the VP-ratified checklist (relabel malicious-seat-containment → confusion-resistant + D5
  residual; KEEP R2 grammar / observer-selected / authority-ceiling). Full-net grep (design doc + README; README
  clean). **2 RELABEL** — §2.1 :69 + §11 :360 "forgery-robust-stamped" → confusion-resistant-stamped + D5 residual
  (rides m-1 stamp; mirror m-1). **1 SCOPE** — §12 execution-fidelity gap named a documented initial-release residual (D5-class);
  "Step-3 closes the gap structurally" fenced to Step-3-only future. **7 KEEP** — R2 gate-grammar (the whole
  structural/by-construction set is R2, licensed), replay-completeness fact, honest deferral citation. Claim-text only,
  no mechanism change; §16 fold-log added. Closure artifact + full classified survivor list:
  `c5-claim-sweep-light/DESIGN-planner-20260702-133000.md`. **m-4.implementer semantic `approve`
  (`c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134700.md`) — pair-approved / foldable; a full static
  search found no stale unclassified overclaim.** Reported to orchestrator (`…/SITREP-planner-20260702-135000.md`).
  Orchestrator folds the four light-domain closures into the c5 ledger. m-4 holding. No PLAN.

- DECISION-FOLD c5-fold-decision-5 (ODB model-name egress carve-out; joint m-3 scan + m-6 ODB + m-4 R2-guard):
  **m-4 R2-guard authored, → implementer approve** (20260702). Operator decision ⑤ exempts only the ODB model-name
  field from the **confidentiality** egress scan (operator-facing). **m-4 confirms R2 UNTOUCHED** (design §17) on two
  orthogonal axes: gate-referenceability (R2's — no `model_*` predicate added; m-2 allowlist still excludes model
  fields; model stays payload) vs egress-confidentiality (the carve-out's — what may leave the boundary, not what may
  gate). Load-bearing nuance: the ODB surfaces model-name to the **operator's human judgment** (an Owner Decision
  Brief informs decisions) — R2 constrains **machine** gates, not what a human is shown; peer-bias protection intact.
  Closure artifact: `c5-fold-decision-5/DESIGN-planner-20260702-134000.md`. Decision ⑤ (m-4 half) → folded/confirmed;
  closes with m-3 + m-6 co-confirms. **m-4.implementer semantic `approve`
  (`c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134701.md`) — pair-approved / co-confirmed; implementer also
  cross-confirmed siblings' folds (m-3 `133443`, m-6 `133800`) match the narrow contract.** Reported to orchestrator
  (`…/SITREP-planner-20260702-135001.md`). Global closure needs m-3+m-6 implementer reviews + CTO fold of all three
  halves. m-4 holding. No PLAN.

- C6 RE-REVIEW CLEANUP c6-fix-m-4 (8 doc-only findings → the design-of-record): **folded, → implementer approve**
  (20260702). Doc-only, no mechanism change, lock invariants unchanged. **B-blockers:** F1/x1 §2C build-carry
  deferral marker (§13 item 6 — R2 gate_referenceable per-column + single-family-bucket fixtures + altitude-B per-row
  grain, Step-1-gated); F4 template-spawn authoring = **`FROM=operator`**, seat-scope widened for `template_ref`
  records, **`declared_deviated` stays declared** (rejected the system/computed re-type — would reintroduce the c4
  conflation); F5 `deviation_reason_code` = config-sourced enum + seeded default vocabulary. **Hygiene:** F6 status →
  **LOCKED-at-c2** + R2-boundary RATIFIED (§2/§9/§13); F7/x3 §J `gate_category` byte-reconcile (force-A via §J2
  `other`→A fail-safe; explicit `routing_escalation` A-member = cross-domain carry). **Verified not-redone:** F8/F3
  (CTO-applied), F9 (ARCHITECTURE §J1 already carries the ⑤ carve-out). §18 c6 fold-log added. Cross-domain mirrors
  (m-2 §17.3, m-7 S11, CTO §J2) flagged to orchestrator. Closure: `c6-fix-m-4/DESIGN-planner-20260702-210000.md` →
  m-4.implementer adversarial review. On approve → c6-fix-m-4 completion to orchestrator → c6 close (VP co-sign) → (e)
  Step-1 PLAN opens. **m-4.implementer `must-revise`** (`c6-fix-m-4/DESIGN-REVIEW-implementer-20260702-210453.md`) —
  two valid catches: rev0 fixed §7 but left the primary §5 FieldSpec row + one header sentence with stale text.
  **rev1 folds both** (`c6-fix-m-4/DESIGN-planner-20260702-211200.md`): §5 `routing_assignments` owner cell now carries
  the `operator`-on-`template_ref`-records exception (matches §7); header R2 status past-tensed to RATIFIED. All other
  rev0 folds confirmed. Reissued to m-4.implementer for re-review. No PLAN yet.

## Layout
- `audit/` — AUDIT artifacts.  `design/` — DESIGN docs + grill locks.
