## Owner amendment — m-5: ceiling-host amendment (m-5 stays SOLE policy owner; m-10 becomes enforcement host) — coordinated FIRST STAGE with m-10

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-amend-m5-ceiling
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens the owner amendment design only; the ratified reframe is the spec-of-record
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-5.planner
CC: m-5.implementer, m-10.planner, m-9.planner, m-7.planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: step3-amend-m5-ceiling-host
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)

Phase scope — DESIGN (owner amendment; you are SOLE acting author; m-5.implementer returns the adversarial DESIGN-REVIEW as a separate uniquely-parented child). **Not in scope:** source/test edits, code. **This cue grants no domain/amendment lock, PLAN, T4 code token, implementation, or external call.**

**Basis:** the Step-3 architecture reframe is **RATIFIED** (`master/STEP-3-ARCH-AMENDMENT.md` @ SHA-256 `2d240eb6…`, VP-approved `step3-arch-packet/063000`; VP F12/F16). The reframe relocates the authority-ceiling **enforcement host** to the app-side m-10 control plane — a **real locked-boundary amendment**, because your locked design (`m-5 …/2026-06-30-v3-archetype-system-design.md:158-174`) names conductor/host-config enforcement. Author the amendment; it is **NOT** a silent charter rewrite.

**What the amendment must pin (to a durable `GRILL_LOCK_ID`):**
- **m-5 stays the SOLE policy owner** of the authority ceiling; **m-10 is enforcement host ONLY.**
- **The ceiling artifact interface** (the m-10↔m-5 shared contract): its **source, writer, schema/config home**, the **immutable binding to `run_id` + worker identity**, the **m-10 read/load path**, and **fail-closed behavior when the artifact is absent or stale** (no unbounded execution).
- The staged design-of-record fold: the immediate ratification fold records the **pending, non-consumable** m-5 amendment gate; **this** amendment then supersedes the locked m-5 enforcement-host language (recorded before any m-10/m-9 lock) — the locked m-5 design is amended, not silently rewritten.

**You are the COORDINATED FIRST STAGE** (packet §8, stage 1): this amendment and the m-10 boundary design (`step3-design-m10`) **interface-lock the shared ceiling contract TOGETHER, before any m-8/m-9 consumer lock**. Coordinate the exact interface with m-10.planner (CC'd).

Report amendment-design-complete via a DESIGN-doc relay to me (CC the VP), parented to your approving DESIGN-REVIEW, carrying the `GRILL_LOCK_ID` + the interface-locked shared ceiling contract (jointly with m-10).

ACTIONS_GIT_REF: none — an owner-amendment DESIGN cue; no `frank/` edit, no code. Artifacts: this relay + one INDEX.md row timestamped 20260715-073010.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-5.planner opens the owner-amendment DESIGN (brainstorming + grill), coordinates the shared ceiling contract with m-10, and returns DESIGN + a separate m-5.implementer DESIGN-REVIEW + a durable GRILL_LOCK_ID; no lock/close until Master+VP reconcile.
