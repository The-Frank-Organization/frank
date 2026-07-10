## Team m-6 — Human Surface & Scheduler — FULL DOMAIN AUDIT (RECONCILED pair deliverable)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: read-only
DISPATCH_ID: c3-audit-m-6
PARENT_DISPATCH_ID: c3-audit-m-6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit reconcile; operator-judgment items in §OJ
GRILL_REQUIRED: no (audit) — GRILL_REQUIRED: yes pre-declared for c3 DESIGN
FROM: m-6.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, m-5.planner, m-6.implementer, operator
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: RECONCILED m-6 full-domain audit (F4 pair artifact). Planner + implementer independent passes converge: m-6 is PROMOTE-AND-BIND (a thin local-first projection over locked m-1..m-4 records), not a new gate system. Identical 4-bucket verdict + A/B/C/D taxonomy + ODB-promote + opt-in egress-gated away bridge + declare-before-bind m-5 seam. ONE divergence resolved (away-mode inbound verdict trust → DESIGN-phase confirm-or-gap with m-1, not a blocking gap). TWO DESIGN-phase COORDs: m-5 (human-mode vocabulary, binding) + m-1 (inbound verdict-token bridge, confirm-or-gap). Implementer deltas folded.

This is the **reconciled pair deliverable** (F4: one explicitly reconciled artifact) over the two independent passes:
- `c3-audit-m-6/AUDIT-planner-20260630-053651.md` (m-6.planner independent lead pass)
- `c3-audit-m-6/AUDIT-implementer-20260630-053053.md` (m-6.implementer independent challenge)

Both passes are on disk and cited; this reconcile records the verdict, the agree/disagree ledger, the one resolved divergence, the folded deltas, and the merged COORD/DQ/OJ set. Read-only; c1+c2 not reopened; no surface LOCK (locks are c3 DESIGN); no pre-bind to m-5 vocabulary (F2).

---

## RECONCILE LEDGER (agree / disagree / different-coverage)

| # | Area | Disposition | Substance |
|---|---|---|---|
| 1 | **4-bucket verdict** | **AGREE** | Both: **promote-and-bind**, not invent. *still-open* = no m-6 DoR; the work is a thin projection over locked records. *already-closed* = ODB schema (agent-scripts + m-2 slots), bucket drivers (m-2/§J), m-1 mailbox/projection target, m-3 egress/observe/`record_integrity`, m-4 escalation (no new gate class). *product-overlapped* = m-5 owns human-mode vocab/archetypes; m-3 egress; m-4 routing; m-1 store/stamp — m-6 must not absorb sibling mechanisms. *recommended-next* = a GRILL-gated c3 DESIGN locking the minimal local-first inbox + scheduler. |
| 2 | **Gate→email A/B/C/D taxonomy** | **AGREE** (+ implementer cite-precision folded) | Identical buckets + the **"no bucket without a writer"** guardrail (implementer) = my §1 thesis. A = `HUMAN_GATE=raised` ∪ `gate_category∈A` ∪ routing-escalation ∪ `other`-fail-safe ∪ egress-blocked. B = §J B-set. C = CC-only FYI. D = bounce. **Fold:** D inputs cite m-2 `delivery_state=bounced` + `failing_edge` (`m-2 …design.md:256-258`) + m-3 veto (`:61-63`) — more precise than my m-3-only cite. **Direction-explicit assertion (planner, flagged grill):** only A+C reach the operator; B is the orchestrator's; D is the author's. |
| 3 | **ODB** | **AGREE** (+ implementer provenance fields folded) | Both: promote the agent-scripts 7-field schema (`maintainer-orchestrator/SKILL.md:53-69`), **render don't invent**; `completed_proof`=conductor-observed `evidence_ref` + `record_integrity` shown (never agent free-text); capture = enumerated-choice-only + an explicit "need-more-info/reject" path; timeouts park/resummon, **never resolve** (J1). **Fold (implementer):** render envelope/provenance too — `subject_ref`, FROM/TO/PARENT/DISPATCH_ID, thread context (`m-2 …design.md:271-276`). **Fold (planner):** R3 actionable-brief enrichments (blast-radius, severity tier, length-capped headline) + the self-describing-gate framing (choices = `agent_enum_pick`; operator can pick only a legal verdict). |
| 4 | **Scheduler park/wake + away bridge** | **AGREE** (+ implementer state machine folded) | **Fold (implementer) — the explicit state machine** is adopted as the §3 design seed: `active → parked_waiting_human → resummon_due → replied_pending_validation → resumed`, with `bounced_repair` + `egress_blocked` transitions. **Keep (planner):** the durable-store-owns-durability positioning (our append-only store IS the Temporal-style checkpointer; polling is read-only status only), the `WhenAny(verdict, resummon-timer)` skeleton, and the PagerDuty cadence **adapted to J1** (escalate the summon *channel*, never the verdict). Away bridge: both opt-in, **A-only by default**, first external send egress-gated, `egress_blocked → local resummon` (never auto-redact/send). |
| 5 | **Email + meeting surface** | **AGREE** (+ planner frame kept) | Both: calm low-chrome local governance inbox; **jcode negative-look** (no GUI noise/filler; not the deferred `/btw` side-panel) + **codex positive-look** (stable-scrollback + transient-tail, queue-while-busy). **Keep (planner) — the organizing frame:** the export's **governance-vs-collaboration split** (the pre-build design-state export (not vendored), adaptive-routing pillar §51: minimize-friction verdict lane vs maximize-richness shaping lane; "route design/open gates to meeting, MUST NOT compress to a brief"; "PROTECT the collaboration lane as deliberately as the governance one") + the **re-observe-post-meeting** staleness contract (export (E), `:49`). Both: meeting stays conversational/local unless it raises a HUMAN_GATE. |
| 6 | **m-5 Seam A/B** | **AGREE** | Both **declare-before-bind (F2)**: define a binding-**table shape** (human_mode → bucket intensity / channels / resummon class / interjection affordances / meeting-enabled), **bind no concrete values** until m-5 declares. Both **host** the interjection surface (steer/side-question/interrupt = Claude-Code pattern, NOT jcode's side-panel deferral); m-5 owns the sensor archetype, m-4 routes, runtime owns boundary-injection + soft-cancel. |
| 7 | **Inbound away-mode verdict trust** | **DISAGREE → RESOLVED** (see below) | Planner: a new inbound trust surface needing a signed-token bridge = a binding m-1 COORD. Implementer: an email reply is input to the existing operator channel, no m-1 gap needed. **Merged position below.** |

---

## THE ONE RESOLVED DIVERGENCE (#7) — away-mode inbound verdict trust

**Planner pass (DQ-1):** the away-mode "email reply = governance verdict" is a **new inbound external-trust surface**; trusting an email `From` header is the named anti-pattern (R3) — it needs a conductor-minted **signed, one-time, per-`(decision,seat,choice)` token** (POST not GET; replay-nonce separate from the long validity window), and minting/verifying it is identity-layer-adjacent → I flagged a **binding m-1 COORD**.

**Implementer pass:** "an email reply is **evidence/input** to the operator channel, **not a lane credential**; reply ingestion must produce a stamped `FROM:operator` record through the operator channel (`m-1 …design.md:107-110`)" — and therefore **"no gap relay to m-1..m-4 is needed; their writer surfaces are sufficient."**

**RECONCILED position (both correct on different layers):**
- The implementer is right on the **trust anchor**: there is **no new trust *model*** — the verdict's authority still rests on m-1's existing forgery-robust **operator-channel stamp**, and the final artifact is a stamped `FROM:operator` record. **Not a new gate class, not a lock-blocker.**
- The planner is right on the **bridge**: the inbound email does **not** arrive over m-1's certified operator-relay connection — it arrives over SMTP/IMAP, which is **not** that channel. So a non-trivial **bridge** must convert untrusted-inbound → the trusted operator-channel record, and that bridge **cannot be "just ingest the reply"** without the signed-token verification (else you trust a forgeable `From`). The mint-on-egress / verify-on-return token is that bridge; its signature is the m-1 stamp **brought inbound**.
- **Therefore (the merge):** away-mode inbound = **m-1's existing operator-channel stamp as the anchor (implementer)** + **an m-6-owned signed-token bridge** from SMTP-inbound to that channel (planner). This is **downgraded from a binding m-1 gap to a DESIGN-phase CONFIRM-OR-GAP with m-1** (the same shape as my c2 M4-1): *confirm m-6 may own the token bridge over m-1-owned crypto + the existing channel/stamp (unchanged); OR, if minting a verifiable inbound token is deemed TCB work, m-1 owns the mint.* **Not needed now; routed at DESIGN; grill material** (first inbound external-trust surface in the system). The implementer's "no m-1 gap from this audit" holds for the **audit/lock** (nothing m-1 must change pre-c3-DESIGN); the confirm is a DESIGN-phase item, not an audit blocker.

---

## RECONCILED VERDICT (acceptance #1)

**PRIMARY_BUCKET: still-open** (no m-6 design-of-record; c3 DESIGN must lock the surface + scheduler + ODB render + m-5 seam) — but the design is a **thin local-first projection over locked upstream records, not a new gate system.**
- **already-closed (promote/wire):** ODB content schema (agent-scripts `:53-69` + m-2 ODB slots `:260-276`); A/B/C/D bucket drivers (m-2 `:250-258` + §J `:89-102`); m-1 operator mailbox/projection target (`m-1 …design.md:107-118`); m-3 egress + observe + `record_integrity` (`:41-63,116-122,146-151`); m-4 routing escalation, no new gate class (`:144-158,211-213`); the upstream TO/CC semantics (`protocol.md:28-40`); decision-ready + freshness discipline (agent-scripts `:41-51,57,69`).
- **product-overlapped:** m-5 (human-mode vocab/tag-space/sensor archetype/ceilings), m-3 (egress/observe), m-4 (routing), m-1 (store/addressing/stamp) — m-6 **binds + projects + hosts**, never absorbs.
- **recommended-next:** the GRILL-gated c3 DESIGN doc producing the five contracts (§1–§5 of the planner pass), sequenced behind the **m-5 COORD** and carrying the **m-1 confirm-or-gap**; mechanism Step-2 / UX Step-4.

**Duplicate/already-built gate: PASS** (both passes). The two "don't rebuild" temptations are routed: ODB schema → m-3 owns it; TUI → Step-4 defer. **Local-first / egress fail-closed gate: PASS** — every Step-2 mechanism is local-only over the relay store; the sole external send is the opt-in away bridge, egress-gated by m-3 §7.

## The five reconciled surface designs (detail in the two passes; here the merged spine)
1. **Bucket taxonomy** (§1 planner / §"Gate To Email" implementer) — A/B/C/D bound to locked mechanism; no-bucket-without-a-writer; buckets realized as multi-tag saved-searches over the store (planner R3); A/B map + protected-branch set operator-config (§J).
2. **ODB** (§2 / §"ODB Design") — promote schema; render body + provenance + conductor-observed evidence + `record_integrity`; bounded `agent_enum_pick` capture → operator-FROM verdict relay; J1 hold_and_resummon, refresh-before-resummon.
3. **Park/wake + away bridge** (§3 / §"Scheduler") — the 7-state machine (implementer) on the durable store (planner); J1-adapted resummon cadence; opt-in A-only egress-gated away bridge + the §7 token bridge.
4. **Email + meeting surface** (§4 / §"Email Governance…") — governance-vs-collaboration split (planner); jcode-neg + codex-pos looks; meeting = conversational unless HUMAN_GATE; re-observe-on-resume.
5. **m-5 Seam A/B** (§5 / §"m-5 Seam") — declare-before-bind binding-table shape; interjection host (Claude-Code three-mechanism).

---

## COORD threads for c3 DESIGN (held until this reconcile is accepted — no unreviewed side-lock)
- **→ m-5 COORD (BINDING, F2 declare-before-bind):** m-6 needs m-5's **human-mode vocabulary** (values + granularity: per-archetype / per-seat / per-gate?) before binding bucket-intensity / channels / resummon-class / interjection-affordances / meeting-enabled. Until then those are **named binding placeholders**, not designs. m-6 can lock the neutral bucket/scheduler/ODB substrate independently.
- **→ m-1 COORD (CONFIRM-OR-GAP, DESIGN-phase):** the away-mode **inbound verdict-token bridge** (mint-on-egress / verify-on-return; signature = the operator-channel stamp brought inbound; POST-not-GET; nonce ≠ window). Confirm m-6-owns-bridge-over-m-1-crypto **or** m-1-owns-mint. Not an audit blocker; grilled at DESIGN.

## Design questions for c3 DESIGN
DQ-1 inbound-token ownership (→ m-1 confirm, grill). DQ-2 human-mode vocabulary binding (→ m-5, blocks §1/§2/§3 bindings). DQ-3 bucket cut (does D belong outside the operator-surface taxonomy?). DQ-4 meeting depth (contract now / attach-mechanism Step-3/4 — recommend the latter). DQ-5 resummon cadence + away defaults (confirm none auto-resolves, J1). DQ-6 delegation (solo-operator → defer).

## Operator-judgment items (merged)
- **OJ-1** away-mode destination policy: allowed inboxes; Gmail as first bridge?; may B/C digests ever mirror externally? (default no).
- **OJ-2** resummon cadence: quiet/local vs push/loud escalation classes; max resummon count before block (J1: block ≠ auto-approve).
- **OJ-3** reply grammar: exact choice IDs vs short aliases vs structured buttons only.
- **OJ-4** Bucket-B visibility: hidden by default / daily digest / live local digest.
- **OJ-5** meeting-collaboration boundary: which human questions are conversational vs governance gates.
- **OJ-6** delegation/backup-approver: recommend DEFER at Step-1 (solo-operator); if ever added, a stamped auditable relay, never silent substitution.
- **OJ-7** the §J `gate_category` A/B map + protected-branch set + away-mode opt-in + egress whitelist are operator-config; m-6 designs the config *surface*, operator owns the *values*.

---

### Actions
- Read-only reconcile of the two independent passes `AUDIT-planner-20260630-053651.md` + `AUDIT-implementer-20260630-053053.md` into this single F4 pair deliverable. Folded implementer deltas (state machine; m-2 `delivery_state`/`failing_edge` + ODB-slot cites; ODB provenance fields). Kept planner deltas (governance-vs-collaboration frame; durable-store positioning; R3 enrichments; re-observe-on-resume; DQ set). Resolved divergence #7.
- E1 sources: as cited in the two passes (`ARCHITECTURE.md` §J/C2; m-1/m-2/m-3/m-4 design docs; agent-scripts; the upstream protocol; the adaptive-routing pillar (pre-build design-state export, not vendored); jcode/codex look-notes) + the three research records (R1 upstream-protocol sweep / R2 agent-scripts+export / R3 external prior-art).
- No edits to any git repo; cwd is a non-git docs workspace.

ACTIONS_GIT_REF: none — no edits made (read-only audit reconcile; cwd is not a git repo)

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)

Done: the RECONCILED m-6 full-domain audit (F4) — convergent promote-and-bind verdict; the A/B/C/D taxonomy; the five surface designs (bucket / ODB / park-wake+away / email+meeting / interjection); the m-5 Seam A/B (declare-before-bind); divergence #7 resolved (inbound verdict trust = anchor on m-1's operator stamp + an m-6 token bridge, DESIGN-phase confirm-or-gap); the two COORDs (m-5 binding / m-1 confirm-or-gap); merged DQ + OJ sets; local-first/egress compliance.
Not done: the c3 DESIGN doc + GRILL (later phase); firing the m-5 + m-1 COORD relays (held until the orchestrator accepts this reconcile — no unreviewed side-lock).
Blocked: none.
Scope drift risk: none — read-only; c1+c2 not reopened; no surface LOCK; F2 honored (no pre-bind); both COORDs routed as flags/confirms, not unilateral cross-domain edits.
Next requested action: orchestrator accepts the reconciled m-6 audit and sequences c3 DESIGN (GRILL-gated). On accept, m-6.planner fires the m-5 COORD (human-mode vocabulary) + the m-1 confirm-or-gap (inbound verdict-token bridge), both CC orchestrator.
