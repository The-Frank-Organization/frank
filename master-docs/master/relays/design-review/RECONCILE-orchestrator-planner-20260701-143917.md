## RECONCILE — deployment fork DECIDED (attach + confused-threat + shelve adversarial); review must-fix reframed; VP: pressure-test the bucketing

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: design-review
PARENT_DISPATCH_ID: design-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — records an operator grill decision + reframes the review scope; does not open PLAN
GRILL_REQUIRED: no — the grill is complete (`GRILL-LOCK-deployment-fork-20260701`)
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner — you flagged the deployment fork as the operator's to decide. It's decided, via a `design-grill`
(`master/GRILL-LOCK-deployment-fork-20260701.md`), and it **materially reframes the review's must-fix scope** — so I
want your adversarial eye on the reframe before it becomes re-baseline task text.

**The decision (operator, grilled):**
- **D1** Step-1 threat model = **confused/hallucinating, not adversarial** (the operator's call: confused agents are the
  real, frequent failure mode; deliberate malicious agents are rare + later; and even frontier-lab containment is
  defeatable — our own §14 found `srt` foolable + the Ona self-disable).
- **D2** Step-1 = **ATTACH**, guardrail **interface-level** (seats act only through `submit()`/`project()`/`read()`;
  config is conductor-owned, not a seat tool). Resolved by inspection: a confused agent fails *through* the gate, not
  *around* it, so wrap's DI-1 isolation defends a threat that's out of Step-1 scope.
- **D3** wrap / adversarial-containment / **"by-construction" SHELVED indefinitely** (research-gated) — not a roadmap rung
  (kills the "false gradient").
- **D4/D5** honest claim-set ("confusion-resistant; malicious code-executing agent explicitly out of scope; re-open
  trigger = containment research matures or threat escalates") + config/store/operator-FROM recorded as **documented
  accepted-risks**, not silent holes.

**The reframe (what I need you to break if it's wrong):** I bucketed the review's ~12 FATALs into **(i)
adversarial-only → COLLAPSE** to a global claim-sweep + accepted-risks (no architecture change), vs **(ii)
threat-independent → REMAIN** as Step-1 re-baseline MUSTs.
- **COLLAPSE (adversarial-only):** I1 sole-writer, operator-FROM forgeability, config-lane-writable, sole-external-sender,
  fill-time-authority, the three-rung ladder. *(A confused agent won't `echo >> store` or hand-edit the A/B map.)*
- **REMAIN (threat-independent — bite honest agents too):** (1) **conductor-core** has no design/owner; (2)
  **serialization + crash-atomicity** (two honest concurrent submits double-accept; a crash mid-multi-file-commit loses
  a record) — *note: a single-threaded serialized conductor kills the concurrency half by construction, leaving
  crash-recovery*; (3) the **Step-1/Step-2 phase-split** dead-end (unbuildable); plus the two **confused-agent
  guardrails** the operator wants — **pure-judgment A-floor** (a confused agent mislabels an A-decision as "minor" and
  it vanishes) and **decision-② fail-closed** (a confused agent `self_reports` a merge-done that never happened).

**Your job (the R2-catch again):** did I mis-bucket any finding — i.e., is there an "adversarial-only" item I collapsed
that actually bites *confused/honest* agents too (and thus must REMAIN)? Pure-judgment absorption and ② I already moved
to REMAIN for exactly that reason; check I didn't miss another. And do you concur the honest claim-set (D4) + the
accepted-risk framing (D5) is the right resolution of your "global claim sweep" finding?

**Governance state:** `DESIGN-REVIEW-2026-07-01.md` + `READINESS-REGISTER.md` updated to reflect the fork resolution +
the reframed scope; `GRILL-LOCK-deployment-fork-20260701.md` written. Still NO-GO; re-baseline steps (b)–(d)
(stand up conductor-core → fold the 5 decisions + fix mechanical FATALs → then Step-1 PLAN) remain open. No PLAN, no
code/pcode/spike, no domain-design edit.

ACTIONS_GIT_REF: ran the `design-grill` (operator interview); wrote `master/GRILL-LOCK-deployment-fork-2026-07-01.md`; updated `master/DESIGN-REVIEW-2026-07-01.md` + `master/READINESS-REGISTER.md`; wrote this relay + appended `master/relays/INDEX.md`; no domain-design edits, no code/source/pcode, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP pressure-tests the finding-bucketing (any adversarial-only item that actually bites confused agents?) + concurs the honest claim-set; then the operator directs re-baseline step (b) — stand up conductor-core with a named owner.
