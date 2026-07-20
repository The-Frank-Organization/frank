# PROTOCOL-DEVIATIONS — how the frank team framework extends / deviates from the stock agentic dev team protocol

**What this is.** A living register of *every* way our operating framework departs from the stock
**agentic dev team protocol** in the course of designing → planning → building **frank**. The stock protocol supplies
the physics (role skills, lint-clean file relays, authority/evidence/ceremony rules, pair adversarialism, VP
visibility); everything below is a **layer on top — never a change to an installed skill**. Two purposes: (1)
the honest process record — why our team runs differently than stock; (2) the **seed of frank's own product
role-model** — most of these deviations become first-class product features. **Discipline: append every new
deviation here as it is made.** (Sibling docs: `CYCLE-PLAYBOOK.md` = how a *design cycle* ran, worked-examples;
this doc = the *framework-level* deltas across all phases.)

Each entry: **Stock protocol → Ours → Why → Status → frank implication.**

---

## 1. Design-phase / standing-team deviations (the `master` governing team — the first big departure)

- **D1 · Standing team + persistent charter.** Stock: dated, ephemeral per-sprint teams. Ours: an undated
  standing team (`RUN_ID master`) whose charter lives in `CLAUDE.md`, auto-loads every session in the cwd, and
  survives compaction. *Why:* a multi-week effort outlives any one context window. *Status:* adopted.
  *frank:* the charter/persistence layer → frank's project-context primitive.
- **D2 · Durable domain ownership (m-1..m-7).** Stock: task-scoped bundles. Ours: N pairs own *durable domains*
  (each owns interfaces others consume; consumes/collision edges written down). *Why:* decompose by ownership,
  not task, for stable cross-domain contracts. *Status:* adopted. *frank:* → m-5 archetype/ownership model.
- **D3 · Visible, durable relay root.** Stock/sprint-doc-setup: hidden gitignored `.relays/` + dated
  `docs/sprints/…`. Ours: a *visible, durable* `master/relays/` — the relay trail IS the design-of-record.
  *Status:* adopted (deliberate sprint-doc-setup override).
- **D4 · Phase-band scope governor.** Stock: full lifecycle per team. Ours: banded cycles (AUDIT+DESIGN only;
  terminate at design-lock; MERGE stays gated). *Why:* stop a design team drifting into code. *Status:* adopted.
- **D5 · Role recast under the band — `.implementer` = adversarial design-reviewer.** Stock: implementer builds.
  Ours (design phase): the m-x `.implementer` is the *adversarial design-reviewer, not a builder*. *Why:* no code
  in the design band. *Status:* adopted — but the root of the role-name confusion later diagnosed (→ F1).
- **D6 · VP on a different vendor/lane than the CTO.** Stock: has a reviewer. Ours: the VP is a *different
  model/vendor* — adversarial review only bites without shared blind spots. *Status:* adopted.
- **D7 · Domains added mid-flight by evidence.** m-7 Conductor-Core stood up only after an adversarial review
  found the runtime substrate was nobody's domain. *Why:* the org is evidence-revisable. *Status:* adopted
  (sanctioned move: VP-approved decomp + charter + boots).
- **D8 · Adversarial pre-build review empowered to RETRACT.** A multi-lens × multi-verifier fleet over the whole
  design-of-record, allowed to *retract prior certifications* and return NO-GO → a bounded re-baseline, not a
  rewrite. *Status:* adopted (the 2026-07-01 NO-GO).
- **D9 · Review-driven design-of-record amendment (never silent re-design).** Locked-design changes go through a
  reviewed amendment (c6 / c6.1 / step1-prep / s2-amend-m-1), pair-approved + VP-co-signed. *Status:* adopted.

---

## 2. Build-phase deviations (the nested slice-teams — the second big departure)

- **B1 · Team-per-slice / nested orchestrator-team instances.** Stock: one flat team. Ours: each build slice
  (s1, s2, …) is its *own* orchestrator-team instance (own orchestrator-planner/reviewer + core pair), guided by
  an m-x leader. From S2 on: a *new* team per slice (not reused — new sprint = new team). *Why:* dogfoods the
  roadmap's Step-5 nested/recursive-team vision as the build method; fresh eyes per slice. *Status:* adopted.
  *frank:* the nested-team product capability (m-5 + the Step-3 spawn primitive).
- **B2 · The m-x guide relationship.** New in build: a master domain pair *guides* a slice-team — feeds the
  locked design, answers domain questions (→ **m-x.planner**), co-gates. *Status:* adopted.
- **B3 · The F3 fidelity edge.** The contract-owner reviews the slice's consuming-surface *fidelity* (correct use
  of a locked contract) as a precondition to the slice's dispatch. *Refinement (2026-07-04):* loop the m-x
  **planner/owner**, not just the implementer — the owner must see how their contract is consumed. *Status:*
  adopted + refined.
- **B4 · F2 risk-scaled plan-gate.** Stock: uniform gate. Ours: S1/bootstrap gets a guide+VP plan-gate; S2+ run
  on the normal pair Implementer plan-review + *conditioned delegated dispatch*, escalating to master only on
  named triggers (scope/boundary deviation · hard trigger · locked-contract touch · design-of-record amendment ·
  cross-slice collision). *Why:* pay bootstrap ceremony once, not every slice. *Status:* adopted.
- **B5 · Two-workspace relay split.** Governance relays in cwd `master/`; slice-team relays in the build repo
  `frank/` (via sprint-doc-setup). The master→slice dispatch bridges them. *Why:* build relays live *with* the
  code (Cardinal #2). *Status:* adopted.
- **B6 · Boot + dispatch, two artifacts per slice-team.** The master-boot pattern (report-only onboarding SITREP)
  applied one level down: a fresh slice-orchestrator gets a *boot* (comes online) distinct from its *dispatch*
  (the work). *Status:* adopted.
- **B7 · Escalate-on-locked-contract-touch → design-of-record amendment.** A build finding that needs to extend a
  *locked* contract escalates to the owner → is folded into the design-of-record (s2-amend-m-1: the
  conductor-internal `system` provenance). Ties the build back to the spec. *Status:* adopted (fired correctly, S2).
- **B8 · Governance-through-the-product (the s5 dogfood) — run, yielded, stood down.** Stock: relays are files the
  operator hand-carries. For s5 (2026-07-06) the slice-team's live governance ran **on frank itself** —
  boot/dispatch/SITREPs via `submit`/`project`/`read`, hub-and-spoke m-x routing via master (m-x seats unminted by
  design), an operator seat authoring waiver/§7 records. *Why:* generate real multi-seat usage data before building
  deeper on the transport. *Outcome:* it worked as an experiment precisely by failing as a transport — one day of
  real traffic surfaced 17 findings (`master/TRANSPORT-FINDINGS-2026-07-06.md`), headlined by the **F11 lineage
  livelock** (stop-the-line). Operator-directed stand-down at checkpoint; conductor decommissioned, store archived
  as evidence; s5 resumed on stock file-relay (`s5-resume`). *Status:* **retired at checkpoint — by design, the
  fallback was always file-relay.* *frank:* the findings ledger IS the transport-fix cycle's spec seed; re-adopt
  dogfooding after the fix (its relaunch gate: this same run's traffic pattern must not livelock).
  **[Sweep 2026-07-08, Step-1 close: the relaunch gate PASSED** — the archived traffic landed 14/14 on the fixed
  conductor (zero parent-class, zero livelock; `RECONCILE.md` § s6). Re-adoption rides the Step-2 relaunch.]
- **B9 · Executable-truth discipline (adopted 2026-07-07; source = an external case study — "How Agents Quietly Break
  Architecture", transcript at `agents-quietly-break-architecture-transcript.txt` — local reference copy, not vendored — cross-read against our
  own incident record).** The failure class: agents preserve *local* module
  contracts while quietly deforming *global* semantics (the case study's bug: a test made green by coalescing a
  point-like value into a range, erasing a design-named distinction); prose rules don't compile and go quietly stale.
  Our own worst defects were exactly this shape — the three-parsers-one-envelope divergence (F6/F7), the
  activation-marker cross-domain conflict (two locally-approved docs, contradictory global semantics — VP-caught),
  DEF-2 (lane-suppliable system headers: locally validated, globally dishonest). Three rules: **(a) "green-by-erasure"
  is a NAMED standing lens** in every code-review panel + pair-review checklist — *does this diff satisfy a test by
  widening/coalescing a type or semantic the locked design distinguishes?* (unnamed lenses get skipped under
  pressure); **(b) every locked law names its executable fixture** — a design-phase EXIT criterion, not a habit (s6
  practiced it: FX-B1g derived from the derived-only rule); **(c) point-not-restate** — dispatches + integration
  artifacts point at the authoritative doc + executable checks, never restate contracts in prose (restated prose is
  where the s6 co-sign found its conflicts). *Status:* adopted — (a)/(c) immediate, (b) rides every future design
  gate. *frank:* the **INV-CATALOG carry** (`ARCHITECTURE.md` §C4) — the laws become one named-invariant test
  package; horizon: a declared "law layer" (machine-checked global invariants beside the m-2 registry) as a product
  differentiator.
- **B10 · Lean slice — a domain pair MAY be the build agent-pair (operator-ruled 2026-07-10; s7-scoped precedent).**
  Stock-of-our-own-making: master domain pairs design/plan/review docs and *never type code* (the STEP-1-KICKOFF
  guide fence); every build slice spawns its own team. For **s7 INV-CATALOG** (a small, test-only phase-opener) the
  operator ruled the ceremony down: **no new sessions — the m-7 pair itself is the build agent-pair**, with the
  roles held straight *inside* the pair: `m-7.implementer` = the sole code writer; `m-7.planner` = the pair Planner
  (domain guide + adversarial implementation review before anything reaches master). The operator's stated floor:
  *any build involves at least one real agent-pair* — the pair primitive is the non-negotiable, not the seat
  pedigree or a fresh team. External gates unchanged (m-1/m-2/m-4 scoped fidelity · VP integration review · operator
  merge). *Why:* booting a session costs more than this slice's risk warrants; the catalog consolidates fixtures m-7
  designed, so domain-pair authorship is low-drift — while pair cross-review + the unchanged external gates keep the
  grounding. *Preserve:* the pairs-don't-code fence stays the DEFAULT — this is a per-slice operator ruling, re-earned
  each time (candidate trigger: test-only + small + the owning domain's own artifact); planner seats still never code
  (VP s7-F1 intact). *Second application (operator ruling "A", 2026-07-10):* the **s7a F-S7-R2-COLGRAIN guard lane** —
  extends the shape beyond test-only to a **small bounded PRODUCTION fix inside the owning domain's own mechanism**
  (the m-2 pair: implementer writes the `any_row` column-grain guard red-first in its own validator; planner
  pair-reviews; m-4 + m-7 fidelity; VP integration; operator merge). The trigger refines to: *small + the owning
  domain's own mechanism + the requirements already adversarially specified by the finding record.* *Third
  application (operator ruling "A", 2026-07-10):* the **s7b `OI-S7A-CLOSE-ONCE-RACE` disposition lane** — the m-7 pair
  fixes the client-lifecycle double-close race in its own channel code (planner-first with delegated dispatch, since
  the idempotent-close mechanism choice is real design freedom; acceptance pre-pinned by the VP's finding record).
  Same trigger, now stable across three runs. *frank:* the product's archetype layer (m-5) should carry exactly this knob — a
  ceremony-tier-scaled team topology where the floor is the governed pair, not a fixed org shape.

---

## 3. Model & capability deviations

- **M1 · Heterogeneous model→seat assignment; capability dictates topology.** The seats run on *different models*:
  planner = Claude Code (Agent Teams = real multi-agent orchestration), implementer = Codex (strong coder, weak
  orchestrator). A capability only one seat's model has (panel orchestration) *pins* that duty to that seat.
  *Status:* adopted. *frank:* m-4 routing must pin capability-dependent roles to capable models — a first-class
  routing constraint.
- **M2 · Code review is a multi-lens panel.** Stock: the planner reviews code (single pass). Ours: the planner
  *runs an adversarial panel* (e.g. the 5-lens panel on the s1 code) via Agent Teams. *Why:* multi-lens beats
  single-reviewer; no single blind spot. *Status:* adopted. *frank:* the "x panel" review archetype (m-5 fan-out),
  later native via the Step-3 governed-spawn primitive.

---

## 4. Planned / not-yet-adopted (design now, adopt at a clean boundary)

- **F1 · The 3-role split — planner / reviewer / implementer.** Split the stock-protocol implementer into a dedicated
  *reviewer* (design/plan review — was the implementer's) + a dedicated *implementer* (pure coder). Planner
  unchanged — design + plan + **runs the code-review panel** (kept there for the M1 capability *and* the
  design-intent grounding). *Why:* frees Codex to pure-code; gives design/plan review a dedicated seat; dissolves
  the D5 role-name confusion + the fidelity-routing asymmetry. *Preserve:* the mutual-review grounding — an
  implementer feasibility-flag + a planner design-fidelity flag, so the split doesn't trade grounding for
  independence. *Status:* **planned** — adopt at a clean slice boundary, not mid-slice. **[Sweep 2026-07-08: the
  Step-1 close IS the clean boundary — the adoption decision goes on the Step-2 planning agenda.]** *frank:* the
  `plan / review / build` product role-model.
- **F2 · frank's own governed agent-teams (Step-3 runtime primitive).** A native governed agent-spawn so any seat
  frank drives can spawn a panel/sub-team *inside* frank (each spawned agent channel-stamped + authority-ceilinged
  at spawn) — replacing the ride on Claude Code's Agent Teams and *lifting the M1 capability-dictates-topology
  constraint*. *Status:* **deferred** to Step-3 (runtime) → powers Step-5 (nested teams). Recorded in
  `ROADMAP.md` Step-3.

---

## 5. The through-line (what these collectively seed for frank)

Read top-to-bottom, the deviations are one arc: **the stock agentic dev team protocol is a single ephemeral pair-team; frank's
operating model is a *persistent, domain-owned, recursively-nested, capability-routed, governed* team-of-teams.**
Nearly every deviation becomes a frank product feature — the charter/persistence layer, durable domains +
archetypes (m-5), nested/recursive teams (Step-5), capability-aware routing (m-4), governed multi-agent spawn
(Step-3), and the plan/review/build role-model. **This doc is where the process we are *living* is captured
before it is the product we *ship*.**

- **B11 · Straight-through slice builds — per-task review gates priced out by hand-relay (operator-ruled 2026-07-13; s10-scoped, standing candidate).** Stock: every task lands through a per-task review leg before the next opens. Deviation, operator verbatim: *"we went just fine with going once all the way through and then review at the end, this is unsustainable with hand relaying, mayhaps ok with relay automation in the future when this is all built, we can take a bit more review churn at the end."* The slice builds T1→TN straight through; ONE end-of-slice adversarial review (the slice planner) covers the whole diff against the plan-of-record + every ruling; owner Step-5 confirms BATCH to the end with **stop-on-contradiction** (a late owner contradiction halts the build wherever it stands and unwinds through master — an operator-priced residual, stated at amendment time). **What survives untouched (the substance/ceremony line):** fence discipline (out-of-block ⇒ STOP/escalate), owner-byte fidelity (mismatch ⇒ bounce, never adapt), lock-pin tripwires, condition (f) blocker holds, self-executed per-task disciplines (RED-first, green-at-every-commit FILE-captured, the running seam/license table, Rails A/B, I-PH), and the operator-only merge gate. Register note: the deviation prices REVIEW LATENCY, not review DEPTH — the end review carries the full per-task checklist; and the product under construction is, not incidentally, the relay automation whose absence forced this. Revisit when the conductor carries the relays itself. Record: `s10-build-impl/SITREP-planner-20260713-004230.md` (the token amendment, master-consumed same day).
  - **B12 (design-side analogue) — OFFERED AND DECLINED (operator, 2026-07-13):** *"design churn is fine, needs to be precise before going to implementation."* The asymmetry is the ruling: design loops KEEP the full multi-round r1→rN ritual (precision before implementation is worth the carries); B11's straight-through cadence stays BUILD-scoped. Do not re-offer absent new evidence.
- **B13 · The three-tier slice-orchestrator operating model, run ON frank (operator-decided 2026-07-14; Step-3+; the deferred T4 tier activated).** Stock (and Steps 0–2): TWO tiers — one master orchestrator over flat domain pairs, with the operator as manual transport between independent sessions. Deviation: a THIRD tier is added and the transport moves onto frank. An **independent slice orchestrator team (T4)** — its own planner/implementer per slice or bundle — owns **local detail-design + plan + impl** against the **m-x-authored spec-of-record** (the m-x pairs act as PMs), and **escalates UP through master to the owning m-x planner** on (a) a mistake in the spec, (b) a better way found while filling details, or the standing `DELEGATED_DISPATCH_AUTHORITY` triggers (locked-contract touch · fence deviation · cross-slice collision · design-of-record amendment). Master is router + arbiter; the m-x PM rules the amendment/improvement/decline. **The team's relays flow through frank** (`submit`/`project`/`read`), not operator hand-relay — closing exactly the transport friction B11 priced out. This is the **deferred T4 tier** from the founding org decision (`master/relays/master-org-decomp/…-031111`, 2026-06-28), shelved then on the explicit trigger *"until v3 automates the relaying"* — frank is now that courier, so it activates as-scheduled, not as a pivot. **Spec of record: `CYCLE-PLAYBOOK.md` Part F** (the tier map · the local-vs-escalate authority boundary · the escalation ladder · the per-slice lifecycle). **Still-owed convention debt** (the half frank doesn't solve, VP's 2026-06-28 caution): nested-run lineage across three tiers · the T4 team's authority-ceiling-at-spawn · how master's arbitration reaches down — m-5's archetype layer (`orchestrator_lead` + authority-ceiling-at-spawn) is the eventual carrier; pinned as the first slice runs on frank or via a bounded design cell. Charter rule 1 carries the high-level pointer; the frank live relaunch + seat roster + shakedown is the separate execution artifact that stands it up.

- **B14 · Parallel owner-amendment authoring alongside DESIGN (Step-3 §6 amendment; CTO+VP sequencing authority 2026-07-14; `step3-audit-reconcile`).** Stock/kickoff-as-locked: the pre-build design sequence ran strictly serial — DESIGN → GRILL → OWNER AMENDMENTS/CONSUMER-REVIEW → LOCK (`STEP-3-KICKOFF.md` §6). Deviation (both m-8/m-9 audits recommended it; VP-concurred): after the AUDIT reconcile, the m-8/m-9 **DESIGN** and the three owner-amendment **draft/audit/design lanes** (m-3 provider-egress/m-7-hosted · m-7 credential · m-4/m-2 routing-record) run **concurrently**, compressing the critical path. **The lock gate is UNCHANGED — parallel authoring is NOT parallel locking:** an amendment may not lock/close on assumptions its consumer design hasn't exposed (the m-8/m-9 DESIGN+REVIEW+GRILL outputs feed its consumer packet before the amendment's final review), and no m-8/m-9 lock occurs until all three amendments + paired adversarial reviews + named consumer confirmations close. Recorded in `STEP-3-KICKOFF.md` §6 (amendment block, resulting hash logged in the status header) + the dashboard/charter pointers. No product/secret/risk election; no operator ratification owed. Record: `master/relays/step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260714-234500` (VP concurrence-with-boundary).
