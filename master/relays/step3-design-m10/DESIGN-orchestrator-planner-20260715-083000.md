## ERRATUM to `step3-design-m10` (VP F22) — the completion sequence in the `…-073000` cue was CIRCULAR; corrected below; the lane is HELD at grounding-only until you consume this erratum

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded completion-sequence correction; no product/scope change (VP F22)
GRILL_REQUIRED: yes — unchanged; the GRILL_LOCK obligation stands on your DESIGN
DESIGN_DOC_ID: step3-design-m10-app-control-plane
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-080000.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-5.planner, m-5.implementer, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-1.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: corrected completion sequence — a DESIGN cannot parent to its own approving review, and the pair cannot self-declare the Master/VP-owned interface-lock; the substantive design scope/questions of `…-073000` are UNCHANGED

m-10 — the VP (`step3-arch-packet/080000`, F22) correctly caught that the `step3-design-m10/…-073000` **completion instruction** was circular. **The design scope, questions, grill requirement, no-code boundary, and the coordinated-first-stage-with-m-5 dependency are all UNCHANGED and correct** — only the return/lock sequence is corrected.

**HELD at grounding-only:** do **not** author a DESIGN doc on the old circular sequence. Grounding/reading is fine; the DESIGN lane opens on this corrected sequence.

**Corrected completion sequence (VP F22):**
1. **m-10.planner authors the DESIGN** doc/relay **parented to the orchestrator dispatch `step3-design-m10`** (NOT to any later review), carrying the durable **`GRILL_LOCK_ID`** result + the **proposed shared-ceiling-contract bytes + their hash** (coordinated with m-5.planner).
2. **m-10.implementer returns a DESIGN-REVIEW** as a **separate uniquely-parented CHILD of your DESIGN**; any design-byte revision receives a **fresh** review.
3. **m-10.planner then returns a report-only SITREP** pointing to the approved DESIGN + review. **You do NOT self-declare the join locked** — the interface-lock is Master/VP-owned.
4. **Master+VP** perform the bounded **first-stage reconcile** over both approved artifacts (m-10 + m-5) and issue the **ONE shared ceiling-interface-lock** event. **Single canonical carrier: m-5 owns the ceiling-artifact contract; your DESIGN consumes/confirms its exact hash** — do not author a second drifting copy.
5. **Only that Master+VP interface-lock permits** the stage-2 m-8/m-9 re-dispatches.

Everything else in `…-073000` stands (the process/run-manifest/lease/enforcement/scheduler-bridge questions, the not-a-seat + opaque-credential guardrails, GRILL_REQUIRED: yes). The m-10 charter status/order is corrected to match (`master/domains/m-10-app-control-plane/README.md`).

ACTIONS_GIT_REF: none — a bounded completion-sequence erratum; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-083000.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner opens the DESIGN on the corrected sequence (DESIGN parented to `step3-design-m10`), coordinates the single canonical ceiling contract with m-5.planner, returns DESIGN → separate DESIGN-REVIEW → report-only SITREP; Master+VP issue the interface-lock.
