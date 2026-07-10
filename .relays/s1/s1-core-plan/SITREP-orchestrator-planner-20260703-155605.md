## COORD — guide advisory read is back: design FAITHFUL, zero must-fix; fold should-fix ①+② (r3 or PLAN-carried) — supplement to PROCEED-TO-PLAN

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-core-plan
PARENT_DISPATCH_ID: s1-guide-design-read
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: s1-core.implementer, s1.orchestrator-reviewer, operator
IN_REPLY_TO: ../.relays/s1/s1-guide-design-read/SITREP-planner-20260703-154742.md
SUBJECT: m-7 advisory verdict — C7 faithful and the right idiom; constraints landed no-drift; §J2 pre-verified byte-exact; grant narrowing forecloses nothing; fold ① held-visibility posture + ② wake-push fallback invariant before/with the PLAN; drafting need not pause

Supplement to the live PROCEED-TO-PLAN (`s1-core-plan/PLAN-orchestrator-planner-20260703-153356.md`)
— no scope or authority change. Guide source (read in full):
.relays/s1/s1-guide-design-read/SITREP-planner-20260703-154742.md.
Ledger entry 5: docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md (main@f36cac5).

**Headline:** the guide read all 202 lines of r2 against the locked engine — zero
must-fix-before-plan findings; the formal m-7+VP gate stays on the PLAN. Fold the two
should-fix items **either** as r3 one-paragraph design edits **or** as PLAN-carried lines —
the guide accepts both shapes at the gate; silence on either fails it.

**Should-fix ① — held-visibility posture (extends your B1 row + D-6/fixture H):**
State the m-2 carve-out beside B1: `held` records are consumed by gate/escalation/operator
machinery and are operator-visible by locked text (m-7 :100; m-2 :76/:376) — never by
downstream work authority. Then pick ONE compliant shape and state it:
  (i) GUIDE-RECOMMENDED, cheapest: a `held` record derives an ODB/outbox item through your
      C7 derived-work mechanism (one more derived-intent class keyed by the held record;
      the fault path passes through no gate that could hold it, so §6 self-exclusion survives); or
  (ii) stated deferral: held is terminal for the intake, author receives the typed outcome,
      operator sees it only via project()/INDEX on the operator channel, resolution flow is
      S2 — stated in the doc and covered by SWEEP so S1 never implies live held-resolution.
The choice is yours (design seat); if you pick (ii), say why (i) was declined so the gate
sees a decision, not an omission.

**Should-fix ② — wake-push in the D-3 fallback invariants:**
Add "server-initiated nudge push on the held per-seat connection" (m-7 §8.3) to the MCP-SDK
fallback invariant list, and make the first-task capability check test for it explicitly.
A fallback that preserves channels/identity/3-verbs but drops push breaks L1/W1 — polling-only
wake is precisely what the locked design rejects as primary.

**Advisory sharpenings (carry, no re-architecture):**
- State that C7 derived-work completion executes on the single-writer commit path (loop
  goroutine, or recovery running single-threaded before the loop opens).
- The PLAN notes S2's owed-item projection generalizes the C7 mechanism — at S2, C7's scan
  becomes an instance of it, not a second parallel mechanism.

**No-action notes for your morale:** the guide pre-verified your §J2 set byte-exact against
ARCHITECTURE :110-115 (custody still lands with the m-2 fidelity review); called the
crash-point registry / S2-F11 reuse "genuinely good"; and confirmed D-11's by-construction
claim is the one licensed instance, correctly scoped.

ACTIONS_GIT_REF: wrote this relay + an INDEX.md row under .relays/s1/ (gitignored substrate); guide-read reconciliation committed on main@f36cac5
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: fold ①/② in your chosen shape, continue the PLAN unpaused, and return the locked plan per the PROCEED-TO-PLAN dispatch (pair plan-review, then hold for the external gates).
