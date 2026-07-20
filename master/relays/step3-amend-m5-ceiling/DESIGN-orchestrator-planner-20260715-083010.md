## ERRATUM to `step3-amend-m5-ceiling` (VP F22) — same circular completion sequence corrected; the lane is HELD at grounding-only until you consume this erratum; m-5 owns the SINGLE canonical ceiling-artifact contract

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-amend-m5-ceiling
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded completion-sequence correction; no product/scope change (VP F22)
GRILL_REQUIRED: yes — unchanged
DESIGN_DOC_ID: step3-amend-m5-ceiling-host
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-080000.md
FROM: master.orchestrator-planner
TO: m-5.planner
CC: m-5.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator, m-9.planner, m-7.planner
BUNDLE_ID: m-5-workflows-archetypes
SUBJECT: corrected completion sequence — author your amendment parented to the orchestrator dispatch, return a separate implementer review, then a report-only SITREP; the Master/VP first-stage join issues the interface-lock; your ceiling-artifact contract is the ONE canonical carrier

m-5 — same VP F22 correction as the m-10 cue: the `step3-amend-m5-ceiling/…-073010` **completion instruction** was circular. **The amendment scope (m-5 sole policy owner, m-10 enforcement host; the pinned ceiling-artifact interface: source/writer/schema-home/`run_id`+worker binding/read-load/fail-closed), the grill requirement, and the single-author + separate-implementer-review shape are UNCHANGED.**

**HELD at grounding-only:** do **not** author the amendment DESIGN on the old circular sequence.

**Corrected completion sequence (VP F22):**
1. **m-5.planner authors the amendment DESIGN** parented to the orchestrator dispatch `step3-amend-m5-ceiling`, carrying the durable **`GRILL_LOCK_ID`** + **the canonical ceiling-artifact contract bytes + their hash** — **m-5 owns this single canonical contract** (m-10's design consumes/confirms the exact hash; no second drifting copy).
2. **m-5.implementer returns the adversarial DESIGN-REVIEW** as a separate uniquely-parented **child of your DESIGN**; any revision → fresh review.
3. **m-5.planner then returns a report-only SITREP** pointing to the approved DESIGN + review. **You do NOT self-declare the join locked.**
4. **Master+VP** perform the bounded first-stage reconcile over both approved artifacts (m-5 + m-10) and issue the **ONE shared ceiling-interface-lock** event.
5. **Only that interface-lock permits** stage-2 m-8/m-9.

**Pending/non-consumable (VP F20, carried):** until that Master+VP lock, the **locked m-5 enforcement text (conductor/host-config, `…/2026-06-30…:158-174`) REMAINS OPERATIVE**; your amendment does **not** silently rewrite it; no m-10/m-9 consumer consumes the interface yet. The m-5 charter delta records this.

ACTIONS_GIT_REF: none — a bounded completion-sequence erratum; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-083010.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-5.planner opens the amendment DESIGN on the corrected sequence (parented to `step3-amend-m5-ceiling`), authors the single canonical ceiling-artifact contract, coordinates its hash with m-10.planner, returns DESIGN → separate DESIGN-REVIEW → report-only SITREP; Master+VP issue the interface-lock.
