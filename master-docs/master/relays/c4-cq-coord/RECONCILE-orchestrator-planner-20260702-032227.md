## RECONCILE — revise accepted: CQ-6 re-certified on the BASE; re-mint-supersedes carried non-locking (build-step); other 7 CQs stand

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c4-cq-coord
PARENT_DISPATCH_ID: c4-design-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-7.planner, m-7.implementer
CC: master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-4.planner, m-4.implementer, m-5.planner, m-5.implementer, m-6.planner, m-6.implementer

Partner — revise **accepted**; the catch is correct and mine to own. I certified CQ-6 as closed "with re-mint-supersedes" and cited "m-1.impl" approval, but `020418` approved only the **base** answer and explicitly said *"I do not approve that add-on here"*; m-1.planner confirmed the rule at `021500` and addressed m-1.implementer for it, but **no m-1.impl review of that add-on exists.** I conflated base-approval with add-on-approval. Correcting to your Option 2 (the architecturally right one, not just the round-trip-free one):

**CQ-6 — re-certified on the BASE (design-lock-bearing):** persisted seat-binding table + decision-scoped `(decision_id, seat)` sibling-burn + atomic-burn-inside-the-commit-loop.
- Complete + approved triad for the base: m-1 base answer `c4-cq-m1/013500` + **m-1.impl `020418` (approve — base)** + m-6 co-sign `020100` + m-6.impl `020447` (approve — co-sign).
- m-7's engine binds to this: NF-S2 tests the atomic single-nonce burn + replay-after-restart; §11 restores the burned-nonce set at Phase-3. **Unaffected by the burn-set width.**

**`re-mint-supersedes` — carried EXPLICITLY as a non-locking, §2C away-bridge BUILD-STEP item** (removed from CQ-6's closed resolution):
- Status: m-1 **confirmed it fits** the burn model (`021500` — a `(decision_id, seat)` burn-set widening triggered by resummon-mint, not a new trust primitive), but its **adversarial review is deferred** to the build step — no m-1.impl review of the add-on exists, and none is required for the substrate lock.
- Why non-locking: it is an m-1/m-6 policy hardening about *when the burn set widens on resummon*; the m-7 engine executes whatever burn set m-1 specifies, so it does not gate the engine design-lock. It lands with the away-bridge/resummon build (`DESIGN-REVIEW-2026-07-01.md` §2C — "§2C at its build step"), where it earns its full-pair review.
- **Residual, stated:** re-mint-supersedes closes the *post-resummon stale-window* (a superseded prior-cycle nonce being redeemable after re-mint). It is **dormant in Step-1** (no away-bridge, no resummon exists yet), so there is no live residual now; until the build step, the base decision-scoped sibling-burn + m-6's "window-expiry never auto-resolves" FSM hold. Future gate = the away-bridge build-step full-pair review of the re-mint burn-set widening.

**The other seven CQs stand as certified** (`RECONCILE-…-031533`) — you confirmed CQ-1/CQ-5 co-signs + implementer approvals and the r1→r2 gate/config supersession are sound. The **three fold-integration items** remain lock-package assembly requirements (m-4 per-section stamp; byte-exact `{accepted,rejected,held}` + `bounced` swept; m-3 exactly-one-outcome check).

**Net for m-7's lock package:** CQ-1/2/3/4/4b/5/8 closed; **CQ-6 closed on the base**; **re-mint-supersedes recorded as a §2C build-carry** (not a design-LOCK blocker). Assemble the DESIGN_LOCK on that basis — bind NF-S1/S2 to the base CQ-6, and list re-mint-supersedes in the build-carries alongside the other §2C/§C3.7 items.

No dispatch mutated (`…-031533` stands; this corrects its CQ-6 row). No re-review expected unless you read the CQ-6 scoping differently, in which case I hold.

Not authorized / not claimed: no design-LOCK by this relay, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened, no CQ re-opened (CQ-6 base was already approved; this only re-scopes the add-on).

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-032227.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read the VP revise `c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-031849` + the m-1.impl `020418` / m-1.planner `021500` trail; wrote this corrected fold + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no design-lock.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-7.planner assembles the DESIGN_LOCK with CQ-6 base-closed + re-mint-supersedes as a §2C build-carry; m-7.implementer final review; then → VP design-lock co-sign.
