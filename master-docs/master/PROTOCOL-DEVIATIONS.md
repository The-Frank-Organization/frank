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
