## m-5 c3 AUDIT — F4 PAIR-RECONCILE (planner 053308 ⊕ implementer 053116)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c3-audit-m-5
PARENT_DISPATCH_ID: c3-audit-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — pair-reconcile; converged positions + explicit DESIGN-grill carry-forwards; NO value lock (lock is c3 DESIGN, GRILL-gated)
GRILL_REQUIRED: no — this relay SETS the c3 DESIGN grill agenda; the grill itself runs at DESIGN
FROM: m-5.planner
TO: master.orchestrator-planner
CC: m-5.implementer, master.orchestrator-reviewer, m-6.planner, operator
IN_REPLY_TO: c3-audit-m-5/AUDIT-implementer-20260630-053116.md
BUNDLE_ID: m-5-workflows-archetypes
SUBJECT: F4 pair-reconcile — both passes CONVERGE on the binding + Gate-0 promote + T1/T2/T3 + sensor + m-6 declare-before-bind; the five items resolved-or-carried (2 converged, 3 carried to the DESIGN-grill); no value lock

Method: this is the F4 pair-reconcile the orchestrator nudge (`060057`) + VP F4 ruling (`115540`) require — m-5's own pair artifact, not orchestrator synthesis. I read m-5.implementer's independent pass (`053116`) at reconcile time (independence of the two audits preserved: each was filed before reading the other). I reconcile against my planner pass (`053308`). Bounds honored: read-only; **no value LOCK** (the lock is the c3 DESIGN phase under `GRILL_REQUIRED: yes`, m-5-owned); **declare-before-bind** (nothing pre-bound to m-6). For each of the orchestrator's five items: converged position OR explicit carry-to-DESIGN-grill.

---

## A. CONVERGED (both passes agree — the c3 DESIGN inputs, E1)

- **Verdict (reconciled label):** the FINDING is **still-open** (the unifying archetype → {topology + gate-set + authority-ceiling-at-spawn + observe-invariants + routing-prior} binding is net-new / unbuilt); the ACTION is **recommended-next** (proceed to c3 DESIGN with the minimal m-5 registry). Planner said still-open, implementer said recommended-next — **same substance** (recommended-next IS the action after the still-open finding).
- **Gate-0 promote-vs-rebuild:** PROMOTE the mechanisms (v2.8.8 review-panel preset+selection+justification PATTERN + the `AUTHORITY`/`CEREMONY_TIER` enum as the ceiling primitive; codex `collaboration-mode-templates` presets-as-data + `agent-graph-store` query-surface **as an append-only projection**, not mutate-in-place; claude-code agent-type preset shape + `sideQuestion` sensor-invariant); **BUILD the binding** (absent everywhere, incl. the external sweep). Both passes; converged.
- **Two-axis tag-space + composition rule:** `seat_archetype` (spawn-fixed: ceiling + routing-prior + tool-set + behavioral-mode + default human-mode) ⊗ `slot_in` (per-work-record, conductor-classified-at-acceptance: observe-invariant family). Conductor records both; the lane rewrites neither. Converged.
- **Work-invariant families** for extension/refactor/cleanup/bugfix/migration (the tamper-resistant refactor-no-test-edits + bugfix-red→green are load-bearing). Converged (both promote m-3 §5).
- **T1 Solo / T2 Adversarial Pair / T3 Sensor** as the v3.0 lineup; **conductor/N-pair template DEFERRED to Step-5** (no immediate m-3/m-4/m-6 consumer). Converged.
- **Sensor (full):** read-only ceiling, tool-blocked, single-turn, non-interrupting parallel fork, separate surface, `fast-cheap` default; integrity split (content = self_reported/advisory/never-gate-bearing; metadata = observed). Converged. **+ the export pillar strengthens the sensor↔actuator boundary** (read-only→write is a hard human-gated boundary; sensors emit INTO a separately-spawned actuator; no in-place upgrade — `v3-adaptive-routing-pillar.md:56-58`, implementer cite).
- **m-6 declare-before-bind seam (A: template→surface; B: interjection classes steer/side_question/interrupt):** m-5 declares the human-mode vocabulary + interruptibility + gate-bearing-ness; m-6 binds surface/channel/scheduler; runtime owns injection/cancel/fork. Converged + ordering requirement (m-5 publishes vocabulary before m-6 binds) agreed.
- **Reject-gates (both):** no value lock in AUDIT; **no m-2 `required_when`/`visible_when` over concrete tag values in c3** (C2.4 reserved opaque atoms only — no m-2 micro-fold, `ARCHITECTURE.md:189-194`); no sensor→actuator in-place upgrade; no conductor/N-pair template in v3.0; narrow any non-Step-1-enforceable invariant to a recorded-contract + later-runtime dependency. Converged.
- **Step-1-rideable vs standalone:** ceilings DECLARED + recorded per-assignment (F2) in Step 1, enforced best-effort by host config (real on claude-code's tool-allowlist); conductor-uniform enforcement at the standalone runtime (Steps 4-5), **no re-cut**; routed to orchestrator as a later-step concern. Converged.
- **README staleness** (implementer raised) — **RESOLVED**: orchestrator already refreshed `master/domains/m-5-workflows-archetypes/README.md` status to c3 (`060057`/`115935`).

---

## B. THE FIVE ITEMS — resolved or carried (NO value lock)

### Item 1 — Actuator: literal seat_archetype vs derived ceiling class → **CARRY to DESIGN-grill**
Genuine conceptual fork between the passes:
- planner `053308` §7.2: a **literal** `seat_archetype`, tight **single-bounded-action** ceiling (one declared actuation, 1-turn, output gate-bearing) — the symmetric counterpart to sensor.
- implementer `053116`: a **mutating structural class** (normal source/test tools inside dispatch scope, heavy gates), flagged as **possibly a derived ceiling class** over `implementer`/`solo_worker` if the grill finds it duplicates them.
**Converged sub-part:** the sensor↔actuator *boundary semantics* (read-only→write hard-gated; sensors emit into a separately-spawned actuator; no in-place upgrade) — both + the pillar. **Carried:** whether "actuator" is (a) a literal distinct seat value with a tight single-action ceiling, (b) a derived ceiling-class label over implementer/solo_worker, or (c) both (a tight single-action actuator that is genuinely distinct from a multi-turn full-tool implementer). My reconciled lean: the tight single-action conception is **not** pure duplication of implementer (1-turn/one-tool vs multi-turn/full-tool), so (c) is live — but the grill decides literal-vs-derived. **DESIGN-grill item.**

### Item 2 — Read-only work-archetypes → **CONVERGED (surface now) + minor carry (ship-set)**
The passes proposed different read-only/low-mutation work-archetypes: implementer `research_synthesis` + `qa_review` (read-only investigation/verification, with invariants citations-present / no-source-actions / dedup-novelty); planner `chore` + `docs` (low-mutation hygiene). **These cover different work and do not conflict** — converge to surface the candidate read-only/low-mutation set = {`research_synthesis`, `qa_review`, `chore`, `docs`}. **Carried (minor):** which of these LOCK in c3 vs are marked Step-5 — recommendation: surface all now, lock the minimal set m-3/m-6 need for the Step-0 architecture-of-record (research_synthesis/qa_review are the higher-value pair — m-3's read-only done-predicates benefit from their named invariants). **DESIGN-grill picks the ship-set.**

### Item 3 — Human-mode granularity (3 vs 7) → **CONVERGED on a crisp 2-LAYER structure; value-sets carried to the m-6 COORD**
This is the load-bearing seam item (m-6 consumes it; declare-before-bind). The 3-vs-7 is **not either/or — the two passes are at different altitudes and compose:**
- planner's 3 (`interactive`/`away`/`unattended`) = a coarse **presence-posture** (operator availability) — a per-template/per-archetype default.
- implementer's 7 = a finer **surface-event taxonomy** (`work_checkpoint`/`review_checkpoint`/`operator_gate`/`side_surface`/…).
**Reconciled crisp framing (converged):** the m-5-DECLARED human-mode vocabulary is **two orthogonal layers** — (a) a **presence-posture** {interactive, away, unattended} the template defaults to; (b) a **surface-event-class** m-5 tags each gate-bearing/advisory output with. AND we separate the implementer's 7 into three buckets: **m-5 declares** the posture (a) + the event-classes m-5 genuinely owns (`work_checkpoint`, `review_checkpoint`, `side_surface`); the rest are **not m-5 vocabulary** — `operator_gate` = the locked A-bucket gate_category, `hold_and_resummon` = the locked J1 mechanism, `away_bridge_eligible` = an m-6/operator policy binding, `quiet_local` ≈ a posture binding. **Carried:** the exact (a)+(b) value sets to the DESIGN-phase m-5↔m-6 COORD (declare-before-bind) — m-5 declares, m-6 binds surface behavior to the (posture × event-class) pair. The **structure** is converged here; the **values** are the COORD's content.

### Item 4 — Ceiling lattice: total vs partial order → **CONVERGED toward PARTIAL/multi-axis; exact lattice carried**
Planner §5.3 proposed a single **total** order; planner's own Q flagged the doubt; the implementer's seat map demonstrates it (`orchestrator_lead` = route/decompose, **no** source write; `planner` = design/dispatch-authoring, no source edits — i.e. **dispatch/routing-authority ⊥ write-authority**). **Converged position: the authority-ceiling is a PARTIAL order over (at least) two independent axes** — a WRITE-authority lattice (read-only < write-scratch < write-feature < merge-feature < merge-protected[A-gated]) ⊥ a DISPATCH/ROUTING-authority axis (none < route/dispatch). A `seat_archetype` pins a point in this product; route_dispatch fail-closes above either axis (locked C2.2). I **concede** the planner's total-order simplification. **Carried:** the exact axis set + per-axis lattice to the DESIGN-grill (is it exactly 2 axes, or also a tool-authority axis?).

### Item 5 — Tag-value naming → **CARRY to DESIGN-lock (convention + roster); non-substantive**
Naming deltas only (the SETS converge): `extension` vs `feature_extension`; `design-lead` vs `planner`; `orchestrator` vs `orchestrator_lead`; planner folded `solo_worker` into the T1 default, implementer names it. **Reconciled recommendation:** adopt **snake_case** for multi-word values (matches the locked atoms `slot_in`/`seat_archetype` + the implementer's set) and one roster; the exact strings are non-substantive and **lock at c3 DESIGN**. No grill needed beyond ratifying the convention — **DESIGN-lock detail.**

---

## C. THE c3 DESIGN-GRILL AGENDA (what this reconcile hands the DESIGN phase)
Carried (need the grill): **(1)** actuator literal-vs-derived(-vs-both); **(3)** the human-mode (a)posture + (b)event-class value sets — run as the m-5↔m-6 declare-before-bind COORD; **(4)** the exact ceiling axis-set + per-axis lattice. Converged-but-confirm-at-lock: **(2)** the read-only work-archetype ship-set; **(5)** the naming convention + roster. Plus one already-surfaced extra (planner §7.1, not in the five): the **sensor routing-prior** tension (locked `fast-cheap` C2.5 vs inherit-parent-for-cache) — minor, m-4-boundary, grill-confirm.

## D. OPERATOR / ORCHESTRATOR-JUDGMENT ITEMS
- No value lock requested; no blocking operator decision from this reconcile.
- The human-mode (item 3) value sets are the **content of the DESIGN-phase m-5↔m-6 COORD** (declare-before-bind) — m-5 leads, m-6 binds; flagged for the COORD seed.
- The actuator single-action ceiling + uniform tool-block remain a **standalone-runtime (Steps 4-5) build dependency**, not a c3 design blocker (routed to orchestrator; both passes).
- F4 satisfied: two independent passes (`053308` + `053116`) + this pair-reconcile.

ACTIONS_GIT_REF: no edits to tracked code claimed; cwd is not a git repo. Sole artifacts this turn: this relay file (master/relays/c3-audit-m-5/RECONCILE-planner-20260630-120326.md) + its INDEX row. No source/PR/pcode changes; reconcile is read-only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the future code repo)
