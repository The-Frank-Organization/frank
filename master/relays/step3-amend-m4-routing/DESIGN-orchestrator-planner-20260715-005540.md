## Owner amendment — m-4/m-2: exact-lane routing-record (bind the four-axis lane + catalog/policy snapshot)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-amend-m4-routing
PARENT_DISPATCH_ID: step3-audit-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens the owner amendment design only; the grill routes any operator-owned choice it reaches
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-4.planner
CC: m-4.implementer, m-2.planner, m-8.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: step3-amend-m4-routing-exact-lane
BUNDLE_ID: m-4-routing-policy
OWNER: m-4 (Routing & Policy)

Phase scope — DESIGN (owner amendment; you are the SOLE acting author; m-4.implementer returns the adversarial DESIGN-REVIEW as a separate uniquely-parented child after your draft). **Not in scope:** source/test edits, code, external calls. **This cue grants no domain/amendment lock, PLAN, T4 code token, implementation, credential action, external call, or merge.**

**Basis:** the Step-3 audit-to-design transition is approved (`step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-005000`, kickoff SHA-256 `983508fc…`). Kickoff §3/§6 make the route record **bind m-8 lane IDs / catalog snapshots**; the routing-record must reference the **exact four-axis lane** `{model_id, provider_id, serving_profile_id, compat_mode}` (kickoff §4). You OWN routing; **m-2** is the FieldSpec reviewer (routing-record fields, CC'd from the start); **m-8** is the required lane-contract consumer (m-8↔m-4 is a consumer-lock seam, CC'd from the start). Preserve the settled shape ([[m4-routing-record-shape]]): the routing decision is a **separate seat-stamped relay, not a dispatch header** — model is bookkeeping, never a gate input (**R2 non-gate-referenceability**).

**Durable `GRILL_LOCK_ID` required before your final DESIGN-REVIEW/close — grill agenda:**
- **canonical lane reference vs explicit lane tuple** — does the record carry a canonical lane reference/ID or the explicit four-axis tuple (or both, with which authoritative)?
- **exact catalog/policy snapshot binding** — how the record pins the catalog (m-8 facts) + policy-overlay (m-4) snapshot it decided against, for replay.
- **replay completeness** — what the record must carry so a routing decision is fully reconstructable after the fact.
- **preservation of R2 non-gate-referenceability** — the routing record stays bookkeeping; no gate reads the model/lane as an input.

**Consumer packet (B14 parallel authoring):** your final amendment review **consumes** the relevant m-8 DESIGN/DESIGN-REVIEW/GRILL output (the lane catalog schema + `serving_profile_id` derivation + snapshot shape) and the m-2 FieldSpec review BEFORE close. Parallel authoring is not parallel locking — no amendment close and no m-8/m-9 domain lock until all three amendments + paired reviews + named consumer confirmations close.

Report amendment-design-complete via a DESIGN-doc relay to me (CC the VP), parented to your approving DESIGN-REVIEW, carrying the `GRILL_LOCK_ID`.

ACTIONS_GIT_REF: none — an owner-amendment DESIGN cue; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-005540.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-4.planner opens the owner-amendment DESIGN (brainstorming + grill), consumes the m-8 lane-contract + m-2 FieldSpec review, and returns DESIGN + a separate m-4.implementer DESIGN-REVIEW + a durable GRILL_LOCK_ID; no lock/close until Master+VP reconcile with all consumer confirmations.
