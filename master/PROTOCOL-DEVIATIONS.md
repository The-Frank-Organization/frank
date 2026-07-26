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

- **B15 · CC ≠ obligation (mechanization candidate; now 7 instances / 4 seats — m-2, m-3, m-9 (×2), plus master's own prior misses; latest 2026-07-25 stage-6 re-lock lane-2 thread).** Stock assumption (implicit): a relay's `CC` line puts its content on the record and can be treated as read/actioned context. Recurring failure mode observed across the stage-6 re-lock: a seat states a position or routes a decision-request in a relay's *prose* while addressing it only `CC`, not `TO` — and no obligation to act attaches, because nothing downstream is required to treat CC as a call to action; the relay lints clean regardless (the trail checks lineage/authority shape, not staffing/addressing *semantics* — the s9 precedent). Concretely this session: m-3's `191500` escalation sat CC-only and created no obligation until its self-corrected `…-101500`/`033000` addressed returns actually routed `TO` master (m-3 self-diagnosed and named it the "7th instance," per `step3-relock-dag-m3/RECONCILE-orchestrator-planner-20260725-093000`). **Rule (standing, restated each time it recurs):** an act is real only when a relay is addressed `TO` the party who must act; CC never obligates; stated intent in CC'd prose never substitutes for a `TO` line. **Status:** recurring failure class, not yet mechanized — filed as a relay-lint candidate (a `false-routing-claim` WARN: flag a relay whose prose states or requests a decision/action from a party who appears only in `CC`). *frank:* a first-class lint rule/mechanized gate in the courier itself — the exact class of confusion frank exists to kill by construction, currently caught only by attentive human/orchestrator reading.

- **B16 · Reference owner state by rule, not by snapshot ("timeless-fold") — a ratification-drafting discipline (surfaced 2026-07-24/25, `step3-relock-settlement-amend` VP review R2-F2, closed at rev4 `1fa71cb8…`).** Stock assumption: an amendment or ratification packet may cite a consumer/owner's *current* mutable working-state artifact (e.g., "the live m-10 row as of today") as part of what it binds or supersedes. Deviation, VP-forced correction: **never freeze an owner's mutable current-working-state pointer into ratification bytes** — a pre-ratification working artifact (still subject to owner churn) is never a durable fold; the ratified rule must instead name the *class* of successor it governs, and owners produce **fresh pair-reviewed successors over the then-current artifact, authored only after ratification**. The settlement amendment's Correction/R2-F2 scoped this exactly: no mutable m-10 current-state snapshot is bound; the rule reaches only the *then-current* m-9/m-10 owner artifacts, post-ratification. **Status:** adopted as a standing amendment-drafting rule. *frank:* a structural distinction the schema/lineage layer (m-2) should be able to express natively — "bind by rule over a future class" vs. "bind by snapshot" — so a drafting mistake of this shape becomes mechanically checkable rather than caught only by adversarial review.

- **B17 · Producer-approved ≠ carriage-settled — sequence producer → carriage → consumer explicitly (surfaced repeatedly across the stage-6 re-lock's R2/R3 producer-gate threads, e.g. `step3-relock-dag-m3/RECONCILE-orchestrator-planner-20260724-190000` and `…-150000`).** Stock assumption (implicit): once a producer owner pair-approves a fact (a field, a discriminator, a row shape), a downstream consumer may bind it. Deviation, named as a standing discipline this session: a producer's pair-approval is **not** itself carriage-settled — an intermediate *carriage* leg (does the fact actually reach the consumer at the right grain, against the right base revision, without a stale-anchor mismatch?) must be independently confirmed before a consumer binds. Concretely: m-10's B/E-carriage row was pair-approved against m-8 **r5**, but m-8's `refusal_stage` mirror landed at **r7** — master explicitly declined to let m-3 bind around that gap and instead routed the consistency question back to the carriage owner (m-3) to rule whether the r5-based row closes cleanly or needs a rebase; similarly m-10's carriage row waited on a **named, revision-agnostic rebase trigger** ("the next pair-approved m-9 revision") rather than binding early against a moving producer. **Rule (standing):** never instruct or allow a consumer to bind a producer fact through a carriage leg that has not itself been confirmed current; a "trigger fires → substantive re-review → rebase → re-approve" step is mandatory whenever the producer revision moves between the carriage's approval and the consumer's bind. **Status:** adopted. *frank:* a lineage/freshness check the m-2 schema + m-7 commit-loop should be able to enforce mechanically — a carriage artifact citing a producer hash is stale by construction the moment that hash's producer revises, and a lint/gate should refuse the bind rather than rely on an orchestrator noticing.

- **B18 · Don't over-read envelope-saturation as legal-frame equality ("frame-bound honesty"; `step3-relock-settlement-amend` VP review, SETTLE-VP-R2-F1, closed rev4 `1fa71cb8…`).** Stock assumption: a limits-table sum that is constructed to saturate a hard ceiling (e.g. `FRAME_MAX`) can be read/stated as proof that some legal operating frame actually *attains* that ceiling. Deviation, VP-forced correction: a sum saturating an envelope **by construction** (the table is built so its parts add to the max) proves nothing about whether any *legal* frame reaches it — the two are separate claims and must carry separate, independently-proven constants. The settlement amendment's fix: two assertions cover two carrier shapes, parked rows get one B.4 growth site, and the actual legal production witness is bounded by its own proven, separately-stated conservative constant (`FRAME_CONTENT_BOUND = 3,704,832`), never asserted equal to the saturating sum. **Rule (standing):** never let an envelope/limits-table construction stand in as evidence for an attainability claim about a real operating frame; state the conservative bound as its own proven fact. **Status:** adopted. *frank:* a documentation/proof-obligation discipline for any future capacity/limits design (m-7 commit loop, m-10 manifest sizing) — flag any place a "sums to the max" table is read as "reaches the max in practice."

- **B23 · Lane-4 team shape = a PAIR, on hand-relayed file relays — closes the choice B22 left open (operator, 2026-07-25).** B22 stood down **only** the courier and explicitly recorded the team shape as *"Open, operator-owned."* The operator then chose: **pair**. Recorded here **agent-authored + operator-cited per §8b** — never a forged `FROM: operator` — citing the operator's decision in-session ("pair is good", in answer to the exact question *does lane 4 keep B21's nested team shape on file relays, or revert to the pair shape of VP-approved rev5*). **This supersedes B21's team-shape half**; B21's courier half was already superseded by B22, so B21 is now fully superseded for lane 4. **Rationale on record:** the deliverable is a single tightly-coupled manifest, which favours a pair on throughput — the operator's own decision rule from the B21 grill (*"if throughput we're not doing it thru frank"*) — and the nested team's principal draw was the frank dogfood, which closed at B22. **Process lesson (VP r8-F1, accepted):** master folded this decision into plan rev8 while the durable record still said the question was open, so the plan asserted an authorization its own governing record did not carry. **A decision is authorized when it is durable, not when it is spoken — record it before building on it.**
- **B22 · frank-as-courier STOOD DOWN for lane 4 — the team returns to operator hand-relayed file relays (operator, 2026-07-25).** Reverses the **transport** half of [B21], not its team shape: lane 4 runs on lint-clean file relays under `master/relays/`, hand-relayed, as every prior slice did. **Reasoning (operator, and it holds):** the protocol's worth is already established by ~2,300 relays here plus ~1,250 in pdc across eleven closed slices and two step-exits, so lane 4 on MCP re-proves it at a far worse exchange rate — the preflight cost ~a dozen operator prompts for five relays, and **that ratio cannot be improved at the MCP layer**: MCP gives a server no way to make a client originate a turn, so a pushed wake has nowhere to land and an MCP-hosted seat is operator-scheduled *by construction*. **The scoping fact to record once and not rediscover per team:** the Step-3 operating model (*"the team runs ON frank as courier"*) is **only fully realizable on the native harness** — the MVP builds the conductor as a **native** relay tool (`relay.submit/project/read`) and MCP is the retained **foreign-harness** path, so this dogfood exercised the *secondary* path and hit a defect inherent to it, not to frank. Revisit when m-9/m-10 exist, where wake-on-relay (`ARCHITECTURE.md`, F61-honest contract; `channel.PushTo`/`NextPush` already ship) turns a delivered relay into **scheduled model input**. **Banked before standing down — the experiment closed successful-and-complete, not abandoned:** four findings no manual relay could surface — the `form_digest`/`CEREMONY_TIER` trap (advertised digest valid only for a tier the caller did not choose; recovery needs a `tools/list_changed` capability Codex lacks — a real cross-harness interop defect against a stated MVP goal), empirical confirmation that `provableParentHint` enforces confusion-resistance by refusing lineage a seat cannot prove, a CC'd relay capturing a reply's lineage parent via `woken_on`, and the roster-introspection gap — all in `FRANK-HARDENING-BACKLOG.md`, with nine committed records exported to `master/relays/step3-relock-lane4/preflight-export/`. **Two of three seat-reported blockers were master drafting errors, not frank defects** (`CEREMONY_TIER: tiny` in the charters; a §7 map asking workers to cite a thread they never see) → **rule adopted: an authority-of-record that names wire fields must be validated against the live tool schema, not against its own prose.** **Team shape was left open by this entry and is now CLOSED by [B23]: PAIR** (operator, 2026-07-25). The transport decision above is independent of it and stands on its own.
- **B21 · Lane-4 slice team = a full nested orchestrator-team on frank, not the B13 pair (operator-directed 2026-07-28; the first T4-tier activation; dogfood driver).** Stock/B13: a slice team is its own `.planner`/`.implementer` **pair**. Deviation, operator-chosen: for lane 4 the T4 slice team is a **full nested orchestrator-team** — **canonical role-stamped seats** `l4.orchestrator-planner` + `l4.orchestrator-reviewer` over **operator-minted, operator-booted worker seats** `l4.w<k>.planner` / `l4.w<k>.implementer` — running **on frank as courier**. The operator's decision rule: *"if throughput we're not doing it thru frank; if not we're going on the full team + frank"* — i.e., the driver is **dogfooding the nested-team model + frank-as-courier**, not authoring speed (a tightly-coupled single-manifest deliverable would actually favor a pair on throughput grounds; the full team is chosen for learning, not speed). **HONEST current-generation mechanism (VP-r6-F4):** native governed agent-spawn + the permission/authority system are **Step-4-deferred** (`frank/ROADMAP.md`); so the workers are **operator-minted (`seat_mint`) + operator-booted independent sessions** (like the m-x seats — NOT l4-native spawns, NOT subagents), and `l4.orchestrator-planner` **dispatches** work to them **by frank relay**, not by spawning. **frank here = courier + seat-identity carrier, NOT the Step-4 native governed-spawn engine.** This **activates B13's deferred debt** — nested-run lineage + authority-ceiling — pinned **interim, at convention/config/read-only-tool + reviewer/master/owner/VP-check grade, NOT m-5 mechanical enforcement**: the ceiling is **monotone non-increasing down the tree** (a worker never exceeds `l4.orchestrator-planner`'s read-only / author-only-via-proposal / no-governed-tree-write ceiling); lineage uses `PARENT_DISPATCH_ID` = the **immediate-predecessor edge** with tier ancestry encoded in the **hierarchical DISPATCH_ID namespace** (`step3-relock-lane4` → `…-l4` → `…-l4-w<k>`), NOT by overloading PARENT as a static tier-parent. Chosen for lane 4 because the task is the ideal first outing — read-only, design-only, no `frank/` write, no external effect, H-12 blocking, blast radius = a design doc. **Status:** operator-directed; folded into the lane-4 plan rev7 → VP re-review; supersedes the plan's B13-pair staffing for lane 4 only. *frank:* this dogfood **INFORMS** (as a battle report) the eventual m-5 archetype ceiling (`orchestrator_lead` + authority-ceiling-at-spawn) + m-1/m-7 native nested-lineage carriers — it does **not** mechanically exercise or harden them (they do not exist yet).
- **B20 · The plan-gate is not automatic on lane / work-unit open — a master-authored plan → VP review should auto-fire before execution (operator-directed 2026-07-27, self-caught at lane-4 open).** Stock: the agentic dev protocol front-loads an explicit planning/decomposition gate before each build unit — a plan is drafted and reviewed *before* work starts. Our practice gap: *within* a phase, lane-to-lane transitions carry **no automatic plan gate** — a new lane (here lane 4, the exit-test authoring) was about to open on the momentum of the prior lane's close, with no planning artifact, until the operator prompted *"first write out a high level plan for this and then route through vp… this planning stage should've been automatic but isnt."* The design churn happened live in chat rather than through a reviewed plan-of-record. **Rule (standing, adopted):** opening any new lane / slice / delegated work-unit **auto-triggers a master-authored high-level plan** (goal · decomposition · team & access · deliverable · boundaries · sequence · open decisions) **routed through the VP before execution or team stand-up begins.** The plan gate is not optional and must not wait for an operator nudge; master fires it on lane open. (This is distinct from B12's design-loop churn — B12 keeps design *iteration* full-ceremony; B20 requires the *plan itself* to exist and be VP-reviewed before a new work-unit starts.) **Status:** adopted 2026-07-27. *frank:* the archetype / expansion-slot layer (m-5) is the natural home — opening a work-unit should mechanically instantiate a plan→review gate as part of the slot, not rely on the orchestrator remembering to.
- **B19 · Owner-authored ratifiable closed sets — master binds by hash, never authors or delegates the closed-set choice itself (crystallized across the stage-6 re-lock's §5-E classification, `step3-relock-dag-m9/RECONCILE-orchestrator-reviewer-20260723-213000` VP ruling + `step3-relock-settlement-amend` m-2 `relay.submit` cell binding).** Stock assumption (implicit, and previously a source of orchestrator error — the VP had twice corrected master on this line before): when an amendment needs a concrete closed value (a resource-identity formula, a member's recipe), the orchestrator may draft it directly or delegate the choice loosely to whichever owner is convenient. Deviation, now stated as a standing rule: **master binds an owner-authored, pair-approved contract by hash into an amendment, rather than authoring or delegating authorship of the closed set itself.** Concretely this session: the `canonical_resource(relay.submit)` formula was authored and pair-approved by m-2 (its own domain — form/schema grain) and then bound into the §D-settlement amendment by hash (`5ec7a3d2…`), exactly the precedent set by the earlier m-3-schema amendment (the owner authors the ratifiable closed bytes; master binds them, never re-derives or restates them). Symmetrically, the VP's §5-E ruling drew the line the other way for A3/B1: those are **owner-local recipe values inside an already-ratified closed set** (m-9's realization choices, not a new closed set), so they stay a delegated design decision under F73 with no operator gate at all — digest-materiality alone is never itself the operator-gate test; the test is whether the ratified field set, formula, identity/version literal, owner, carriage, join, or acceptance property actually moves. **Status:** adopted; the ratified-vs-delegated line is now stated as a reusable two-part test (closed-set authorship → owner-authored, master-bound by hash; owner-local recipe inside an unchanged closed set → delegated, no gate) rather than re-litigated per amendment. *frank:* the schema/lineage layer (m-2) is the natural home for a machine-checkable version of this same authorship-boundary test.
