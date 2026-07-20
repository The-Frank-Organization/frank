## Owner amendment — m-3: provider-request egress class (a NEW egress disposition, distinct from away-email `egress_blocked`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-amend-m3-egress
PARENT_DISPATCH_ID: step3-audit-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens the owner amendment design only; the grill routes any operator-owned choice it reaches
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-7.planner, m-8.planner, m-9.planner, m-1.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: step3-amend-m3-egress-provider-request
BUNDLE_ID: m-3-observation-evidence
OWNER: m-3 (Observation & Evidence)

Phase scope — DESIGN (owner amendment; you are the SOLE acting author; m-3.implementer returns the adversarial DESIGN-REVIEW as a separate uniquely-parented child after your draft). **Not in scope:** source/test edits, code, credentials, external calls. **This cue grants no domain/amendment lock, PLAN, T4 code token, implementation, credential action, external call, or merge.**

**Basis:** the Step-3 audit-to-design transition is approved (`step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-005000`, kickoff SHA-256 `983508fc…`). Kickoff §1 (`STEP-3-KICKOFF.md:13-16`) is explicit: the landed m-3 mechanism is the dormant **away-email local-outbox scanner** — it **cannot front a provider request as-is.** You OWN the new amendment defining a **distinct provider-request egress class** that does **NOT inherit** the away-email `egress_blocked` behavior or the model-name confidentiality rule. This is m-3-authored, m-7-hosted (m-7 hosts the send path; you own the egress disposition contract).

**Locked floors (design WITHIN these, do not re-open them):** a denial ⇒ **zero provider network send** · a no-send path emits **no provider-wire event** · **no fourth relay `delivery_state`** is introduced · each intake still reaches **exactly one existing** relay delivery-state through the mapping you design.

**Durable `GRILL_LOCK_ID` required before your final DESIGN-REVIEW/close — grill agenda:**
- **final-wire-only vs a specified pre/post pair** — is the egress decision a single final-wire authorization, or a specified pre-authorization + post-authorization pair?
- **retry/idempotency relative to authorization** — how re-attempts relate to a prior authorization (consumes the m-9 Q4 packet).
- **no post-authorization mutation** — the m-8 final-authorization / no-post-authorization-mutation constraint.
- **the typed-denial mapping** — the still-open disposition: what a denial maps to among the existing relay delivery-states (not a new token, not the away-email park).

**Consumer packet (B14 parallel authoring):** your final amendment review **consumes** the relevant m-8/m-9 DESIGN/DESIGN-REVIEW/GRILL output (m-9 Q4 retry/idempotency; m-8 final-authorization/no-post-authorization-mutation) BEFORE close. Parallel authoring is not parallel locking — no amendment close and no m-8/m-9 domain lock until all three amendments + paired reviews + named consumer confirmations close. Confirm the m-7 host seam (`step3-amend-m7-cred`) and the m-1 store/provenance boundary.

Report amendment-design-complete via a DESIGN-doc relay to me (CC the VP), parented to your approving DESIGN-REVIEW, carrying the `GRILL_LOCK_ID`.

ACTIONS_GIT_REF: none — an owner-amendment DESIGN cue; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-005520.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-3.planner opens the owner-amendment DESIGN (brainstorming + grill), consumes the m-8/m-9 packet, and returns DESIGN + a separate m-3.implementer DESIGN-REVIEW + a durable GRILL_LOCK_ID; no lock/close until Master+VP reconcile with all consumer confirmations.
