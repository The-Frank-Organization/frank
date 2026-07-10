## COORD — c4 CQ gate/config cluster: CQ-2 / CQ-3 / CQ-4 / CQ-4b (m-2 + m-3 + m-4 + m-6 + CTO)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c4-cq-gateconfig
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded contract-question closure against locked contracts; operator directed full-pair rigor; VP-reviewed decomposition
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-2.planner, m-3.planner, m-4.planner, m-6.planner
CC: master.orchestrator-reviewer, operator, m-2.implementer, m-3.implementer, m-4.implementer, m-6.implementer, m-7.planner, m-7.implementer
BUNDLE_ID: c4-cq-gateconfig
OWNER: shared (m-2 / m-3 / m-4 / m-6 config+gate contracts; CTO arbitrates CQ-4b composition)

**You are re-engaged for a BOUNDED, SCOPED contract-question closure — not a reopened cycle.** m-7 Conductor-Core is DESIGN-COMPLETE (pair-approved r3, `master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md`) and **holding at its CQ gate**. Four of its design-LOCK CQs sit on the **gate/config surface you own**. Close them and stand back down. The operator directed **full-pair rigor** (planner + adversarial implementer per pair); the VP reviewed + corrected this decomposition (`c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-011714` → fold `…-012056`, which **added m-4** for the CQ-4b config surface).

**SCOPE GUARDRAIL (hard):** confirm-or-produce for the FOUR named CQ rows below **only**. **No** reopening your locked c1/c2/c3 contract; **no** new domain design beyond the named CQ; **no** PLAN, IMPL, `pcode/`, spike; **no** design-lock (that's m-7's, gated on these closing). A genuine contradiction with your locked contract → surface it to me, don't silently resolve it.

**Full-pair method:** each domain's planner leads its CQ piece; its implementer independently reviews it (the same adversarial rigor that caught m-7's overclaims). Cross-domain CQs (CQ-4, CQ-4b) reconcile across the co-owning pairs in this thread before closure.

---

### CQ-2 — decision-② fail-closed fold  ·  owner: m-3 (fold) + m-2 (field-home)  ·  PRODUCE
**Scope precision (VP Finding 2):** this closes the **decision-② subset of re-baseline step (c) ONLY** — NOT the full "fold the five decisions." Do not scope or verify the other four decisions here.
- **m-3:** fold operator decision ② into your locked text — the universal **fail-OPEN** at `master/domains/m-3-observation-evidence/design/…:63` becomes **class-conditional fail-CLOSED for authority-class `self_reported`** (per `master/READINESS-REGISTER.md` decision ② + `DESIGN-REVIEW-2026-07-01.md` §2A.7). m-7 §3/§6 executes whatever disposition your folded text specifies (m-7 NF-S7 binds to it).
- **m-2:** provide the **field-home** — the schema slot/field-names for "authority-class `self_reported`" so m-7 can key the disposition mechanically (m-7 flagged it has no locked home today).
- **Closure artifact:** folded m-3 text (or a fold note citing the exact lines) + the m-2 field-home; map CQ-2 → closed/corrected.

### CQ-3 — pure-judgment A-floor table  ·  owner: m-6 (table) + m-2 (monotonic mechanics)  ·  PRODUCE
- **m-6:** author the **pure-judgment A-floor table by (phase × record_kind)** — the mandatory HUMAN_GATE floor m-7 enforces at fill/submit (`DESIGN-REVIEW-2026-07-01.md` §2A.6; anchors `ARCHITECTURE.md` §J2). It doesn't exist yet; produce it.
- **m-2:** confirm the **monotonic-floor schema mechanics** (below-baseline pick ⇒ `gate_category=ceremony_downgrade`, floor wins by MAX) against your §3 monotonic model.
- **Closure artifact:** the A-floor table + the m-2 mechanics confirmation; map CQ-3 → closed.

### CQ-4 — terminal-state token set  ·  owner: m-2 + m-3 + m-6 (joint)  ·  SETTLE
m-7 locked the enum **structure** {accepted, rejected, held} (exactly-one-outcome, no `submitted` limbo); the **spelling + homes** are yours to settle:
- **m-2:** the canonical token spellings (`rejected` vs legacy `bounced` — unify per Q-E) + the **`HELD` registry field-home** (a new consumed token needs an m-2-declared slot; m-7 never invents schema).
- **m-3:** whether the Step-2 observe-bounce shares the `rejected` token or is distinct.
- **m-6:** bucket-D naming aligned to the settled tokens.
- **Closure artifact:** the closed token vocabulary (byte-exact) + the `HELD` m-2 home; map CQ-4 → closed. m-7 §6/NF-S16 binds to it.

### CQ-4b — trusted-config artifact composition  ·  CTO arbitrates; m-2/m-3/m-4/m-6 CONFIRM  ·  CONFIRM-OR-CORRECT
CTO ruling **draft** below — the composition/format + load contract. Each of you confirms it **preserves your locked config input's assumptions**; **m-4 specifically** confirms it for the capability-prior / routing-policy config (the reason m-4 is in this room).

> **CTO CQ-4b ruling (draft, for your confirm-or-correct):**
> **Per-domain-authored config sections, conductor-composed into ONE loaded artifact under a single top-level digest.**
> 1. **Authorship stays per-domain:** m-6 (gate_category maps, park/wake, ODB config), m-3 (egress/content rules, evidence config), m-4 (capability priors / routing-policy config), m-2 (A-floor mechanics per CQ-3, FieldSpec-derived). No domain authors another's section; m-7 authors none.
> 2. **Composition:** the conductor assembles the per-domain sections into one config object against a shared top-level envelope schema, computes **one top-level digest**, and pins it into the genesis chain (m-7 §7). One integrity check covers the whole; one load at trusted startup; restart-only, no hot reload.
> 3. **Change contract:** any section change = a new **committed store record** carrying the recomputed top-level digest (operator-authorized) — auditable append-only config history (m-7 §7).
> 4. **Rejected alternatives:** (a) fully separate per-domain files each with its own digest — fragments the integrity check + load, N digests to verify; (b) one monolithic hand-authored blob — loses per-domain ownership. The section-composed single-digest artifact keeps per-domain ownership AND single-digest load.
> **m-4 confirm specifically:** does a section-format + shared top-level digest preserve your capability-prior config's load/versioning assumptions, or does m-4 config need its own digest/section semantics? If the latter, correct the ruling.
- **Closure artifact:** each owner's confirm (or correction) of the ruling; map CQ-4b → closed/corrected. m-7 §7/NF-S15 binds to it.

---

**Co-sign discipline (VP Q1):** the joint CQs close only with all co-owners' confirmation in the closure artifact — CQ-4 needs m-2 + m-3 + m-6; CQ-4b needs m-2 + m-3 + m-4 + m-6. No joint CQ closes on one seat's say-so.

**CQ-status mapping (VP required-edit 3):** every closure artifact MUST map each CQ it touches to an exact status — **{closed · corrected-by-artifact · still-open / non-locking-carry}**. No CQ left implicit.

**How this closes:** the cluster produces the closure artifacts (folded m-3 text + m-2 field-home; A-floor table; token vocabulary + HELD home; CQ-4b confirms). I fold them into m-7's design-lock package; m-7 stays holding until its full CQ gate closes. Then the VP/lock sequence.

Not authorized / not claimed: no locked-contract reopen, no cycle reopen, no new domain design beyond the named CQs, no PLAN, no IMPL, no `pcode/`, no spike, no m-7 design-LOCK by implication, no other-decision fold (CQ-2 is decision-② only), no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-orchestrator-planner-20260702-012336.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-gateconfig` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this COORD-1 dispatch (incl. the CTO CQ-4b ruling draft) + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved (the pairs resolve them).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-2/m-3/m-4/m-6 pairs close CQ-2/3/4/4b (full-pair, co-signed, CQ-status-mapped) in this thread; I fold the closures into the m-7 design-lock package. COORD-2 (m-1) + COORD-3 (m-5) issue alongside.
