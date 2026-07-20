## HOLD (stop-work) — `step3-amend-m7-cred` is STOPPED; the credential contract cannot continue as an m-7 contract; r3 + its reviews are provisional audit input only

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-hold-m7
PARENT_DISPATCH_ID: step3-arch-reframe
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — containment ahead of an operator-ratified architecture re-cut; this relay stops work, it does not re-design
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-amend-m7-cred-credential-contract
IN_REPLY_TO: master/relays/step3-arch-reframe/RECONCILE-orchestrator-reviewer-20260715-013000.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, m-1.planner, m-8.planner, master.orchestrator-reviewer, operator
SUBJECT: STOP-WORK on step3-amend-m7-cred — provider credentials are app-side (connector/control-plane boundary, m-1-governed), not a conductor-core contract; r3 + three paired reviews = provisional audit input; no r4/lock

m-7 — **stop work on `step3-amend-m7-cred` immediately.** The operator architecture-of-record correction (`step3-arch-reframe/…-011000`), VP-dispositioned `human-decision-required` (`…-013000`), removed the premise of this amendment: **provider credentials do not live in conductor-core.** The conductor never dials a provider and holds no provider secret; credentials attach at the **connector / app-control-plane boundary**, m-1-governed. The current m-7 credential design **cannot continue as an m-7 contract.**

**This hold requires:**
1. **Stop** all semantic design edits, grills, paired review loops, consumer confirmations, and any move toward lock, PLAN, or implementation.
2. **Preserve every current artifact byte** — do not rewrite or delete work produced under the old framing.
3. **Return only a bounded status handoff** naming: your current artifact(s) + path(s); current verdict/lock state; unresolved findings; and the dependencies the reframe must disposition.
4. **Treat the m-7 credential r3 (`…-233829`) + its three paired implementer reviews as PROVISIONAL AUDIT INPUT ONLY** to a re-owned connector/control-plane credential contract — **no r4, no lock may issue.** Its useful threat / census / activation / immutable-authorization work is salvage, subject to fresh owner review.
5. Make **no** source, credential, provider-call, external-send, merge, deployment, or live-store action.

**Do not resume on the correction's prose alone.** m-7 remains the **conductor host** (its store/commit-loop/recovery/trusted-config/seat-guardrail are untouched and not reopened); it is simply not the owner of the provider-credential contract. Where that contract re-homes (m-8 connector actuator vs the new m-10 control plane, with m-1 as secret/provenance boundary) is decided in the operator-ratified architecture-amendment packet.

ACTIONS_GIT_REF: none — a stop-work hold; no `frank/` edit, no code, no credential action. Artifacts: this relay + one INDEX.md row timestamped 20260715-013530.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-7.planner returns the bounded status handoff TO master.orchestrator-planner (CC VP); no further design/r4/lock until the ratified reframe re-owns and re-dispatches the credential contract.
