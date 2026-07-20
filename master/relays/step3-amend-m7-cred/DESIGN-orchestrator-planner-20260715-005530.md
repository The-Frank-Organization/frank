## Owner amendment — m-7: provider credential / trusted-config contract (NEW)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-amend-m7-cred
PARENT_DISPATCH_ID: step3-audit-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens the owner amendment design only; the operator-owned credential-provisioning decision routes THROUGH this grill BEFORE any credential use or E3 call, not at design open
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, m-1.planner, m-8.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: step3-amend-m7-cred-credential-contract
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

Phase scope — DESIGN (owner amendment; you are the SOLE acting author; m-7.implementer returns the adversarial DESIGN-REVIEW as a separate uniquely-parented child after your draft). **Not in scope:** source/test edits, code, credential provisioning/use, external calls. **This cue grants no domain/amendment lock, PLAN, T4 code token, implementation, credential action, external call, or merge.**

**Basis:** the Step-3 audit-to-design transition is approved (`step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-005000`, kickoff SHA-256 `983508fc…`). Kickoff §1 records that m-7 has **no credential contract today**; provider credential + trusted-config handling is a **mandatory pre-build owner amendment.** You OWN it; m-1 owns the secret boundary (CC'd from the start); m-8 is the consumer (drives the lane it is handed, never handles a secret).

**Durable `GRILL_LOCK_ID` required before your final DESIGN-REVIEW/close — grill agenda:**
- **secret source** — where a provider credential originates and how it enters trusted config without crossing the seat surface (seats reach only `submit`/`project`/`read`).
- **endpoint binding / allowlist** — how a credential binds to an allowed endpoint set; confusion-resistance against a mis-pointed send.
- **rotation** — credential rotation without a store/genesis break.
- **redaction** — credential/secret redaction across logs, evidence, and the local outbox (composes with the m-3 provider-request egress class + model-name confidentiality).
- **the operator-owned provisioning decision before E3 credential use** — the decision that routes to the operator through this grill (and gates m-8 Q5 / E3), NOT resolved at design open.

**Consumer packet (B14 parallel authoring):** coordinate the m-1 secret-boundary seam and the m-8 consumer seam; the m-3 provider-request egress class (`step3-amend-m3-egress`) redaction/egress overlaps yours — reconcile in design. Parallel authoring is not parallel locking — no amendment close and no m-8/m-9 domain lock until all three amendments + paired reviews + named consumer confirmations close.

Report amendment-design-complete via a DESIGN-doc relay to me (CC the VP), parented to your approving DESIGN-REVIEW, carrying the `GRILL_LOCK_ID`. **Any credential provisioning/use or E3 call remains separately operator-gated after this design — this cue does not authorize it.**

ACTIONS_GIT_REF: none — an owner-amendment DESIGN cue; no `frank/` edit, no code, no credential action. Artifacts: this relay + one INDEX.md row timestamped 20260715-005530.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-7.planner opens the owner-amendment DESIGN (brainstorming + grill), and returns DESIGN + a separate m-7.implementer DESIGN-REVIEW + a durable GRILL_LOCK_ID routing the operator provisioning decision; no lock/close/credential action until Master+VP reconcile.
