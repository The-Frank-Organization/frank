## Team m-10 — App Control Plane / Supervisor: PROCEED TO DESIGN (coordinated FIRST STAGE with the m-5 ceiling-host amendment)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens design only; the ratified reframe (`STEP-3-ARCH-AMENDMENT.md` @ `2d240eb6…`) is the spec-of-record
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-5.planner, m-8.planner, m-9.planner, m-1.planner, m-3.planner, m-7.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: step3-design-m10-app-control-plane
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)

Phase scope — DESIGN. Planner leads (brainstorming + design-grill); Implementer adversarially reviews. **Not in scope:** source/test edits, code, credentials, external calls. **Design-lock is terminal — no implementation / PLAN / T4 code token.**

**Basis:** the Step-3 architecture reframe is **RATIFIED** (`master/STEP-3-ARCH-AMENDMENT.md`, operator-ratified at SHA-256 `2d240eb6…`, VP-approved `step3-arch-packet/063000`; ratification `step3-arch-packet/070000`). Your charter is `master/domains/m-10-app-control-plane/README.md`. Design the m-10 App Control Plane/Supervisor **boundary** per packet §1/§8 + the charter — treat the packet as the spec-of-record, its open items as **design questions to resolve, not proven facts**.

**You are the COORDINATED FIRST STAGE** (packet §8 dependency graph, stage 1): m-10 boundary design **and** the m-5 ceiling-host amendment (`step3-amend-m5-ceiling`, issued in parallel) proceed together and **interface-lock their shared ceiling contract BEFORE any m-8/m-9 consumer lock**. m-8/m-9 design against your locked interfaces.

**Design questions (to a durable `GRILL_LOCK_ID`):**
- **Process boundary + app IPC** — the control-plane process, its local IPC to workers + connector (framed stdio / unix socket), backpressure; m-10 is **NOT a conductor seat** and holds no `submit` credential.
- **The run manifest** — the pinned-lane MVP run state (immutable m-8 lane ID + catalog digest; **not** an m-4 routing decision, not a conductor gate input); its store + writer + recovery.
- **Active-turn lease + supervision** — worker lifecycle; the one-active-worker lease invariant (distinct from m-9's one-active-turn); **fail-closed** recovery (never auto provider resend).
- **The app-side authority ENFORCEMENT point** — reads the **m-5-authored ceiling artifact** (the m-10↔m-5 shared contract), binds it to `run_id`+worker, **fails closed if absent/stale**; m-10 hosts, does NOT own the policy.
- **Connector supervision + opaque credential-reference orchestration** — never secret bytes (m-1 boundary; m-8 holds).
- **Scheduler bridge** — the app scheduler vs the conductor park/wake/ODB; the bridge **reuses the worker seat's existing verbs**, no new conductor address/event.

**Consumer-lock seams** (interface-lock-first): m-10↔m-5 (ceiling artifact — the shared first-stage contract) · m-10↔m-9 (lifecycle + enforcement) · m-10↔m-8 (connector supervision + credential-reference) · m-10↔conductor (only via the m-9 worker seat's verbs).

Report design-complete via a DESIGN-doc relay to me (CC the VP), parented to your approving DESIGN-REVIEW, carrying the `GRILL_LOCK_ID` + the **interface-locked shared ceiling contract** (jointly with m-5). No lock/PLAN/code until Master+VP reconcile.

ACTIONS_GIT_REF: none — a DESIGN dispatch; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-073000.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner opens the DESIGN lane (brainstorming + grill), coordinates the shared ceiling contract with m-5, and returns DESIGN + DESIGN-REVIEW + a durable GRILL_LOCK_ID; no lock/PLAN/code until reconcile.
