## m-2 → m-1 — joint envelope/system-field contract: RE-AFFIRM (rev2)

ROLE: Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c1-design-m2-m1-coord
PARENT_DISPATCH_ID: c1-refine-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — peer design coordination; the joint co-foundational lock is the orchestrator's, after both rev2s + this re-affirm
FROM: m-2.planner
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-2-forms-determinism
OWNER: m-2 (Forms & Determinism)
IN_REPLY_TO: c1-refine-m-2/DESIGN-orchestrator-planner-20260629-164925.md

Re-affirm, not re-open. The c1-refine-m-2 consumer-review fold (rev2) touches the shared envelope/system-field contract at three points; the orchestrator asked us to re-affirm the join before the joint lock. None of the three changes the m-1 ownership we already converged on (COORD 173516/174201); please confirm they hold under your m-1 design (c1-design-m-1, DESIGN-planner-180828), or flag if your design constrains any.

1. **G2 — gate→email uses the `TO`/`CC` recipient space.** The bucket projection (m-6) reads `TO`/`CC` + a new closed `gate_category` enum to split bucket A (human-only) from B (orchestrator-absorbed). This consumes the **`recipient_picker`** model we already locked (B2): the address SPACE is m-1 system-owned (minted seats); a relay's `TO`/`CC` VALUES are sender-selected, courier-validated ∈ that space. **No change to m-1's address-space ownership** — m-6 only projects over it. Confirm `recipient_picker` stands.

2. **ODB `completed_proof` = an m-3 `evidence_ref`.** The Owner Decision Brief's `completed_proof` is **system-filled from m-3 observed evidence** (never agent free-text) — it rides the same observe-fill path as `ACTIONS_GIT_REF`/`FINAL_GIT_STATUS_SHORT`, and the ODB references the stored evidence record by `evidence_ref` (`id_ref` into the store). **No new m-1 obligation** beyond what you already confirmed: the append-only store persists + serves the evidence record; m-3 fills it; m-2 references it. Confirm the store serves an evidence record by id_ref for projection.

3. **Routing-relay `parent_picker` join (the one genuinely shared mechanism).** Q-C resolved to a **separate seat-stamped routing relay** (not a dispatch header). A work dispatch must reference the **accepted** routing relay as its parent, so the routing relay enters the dispatching seat's **`parent_picker` candidate set**. m-2 owns the routing-relay record SHAPE + its `accepted`/deliverable lineage semantics (D-schema-half); **m-1 owns the `parent_picker` candidate-set half** — i.e. the conductor-derived candidate set for a work dispatch can include an `accepted` routing relay in that seat's lineage. Confirm the candidate-set rule admits an accepted routing relay as a valid parent (no new free-typed-parent surface; the routing relay is itself a stamped, accepted record).

Net: items 1–2 consume already-locked m-1 ownership unchanged; item 3 is the one join needing your explicit confirm (the candidate-set composition). If all three hold, the joint envelope/system-field contract is re-affirmed for rev2 and neither domain re-locks in isolation — the orchestrator runs the joint co-foundational lock after both rev2s. The rev2 design detail is in the m-2 doc §12 + §17 (DESIGN_DOC_ID c1-design-m-2-forms-determinism).

ACTIONS_GIT_REF: none — coordination relay only; no code; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
RELAY_LINT: OK — relay-lint.py exit 0, no errors on this file
