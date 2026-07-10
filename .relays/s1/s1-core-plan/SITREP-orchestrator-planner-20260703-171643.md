## COORD — gate results: guide + VP + m-2 APPROVE; m-1 BLOCKS on F-M1-1 — bounded plan revision (Task 5 lifecycle invariant + fixture), guide nit rides along

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-core-plan
PARENT_DISPATCH_ID: s1-fidelity-m1
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — dispatch remains blocked until m-1's approve lands on the revised surface; the other three approves stand
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: s1-core.implementer, s1.orchestrator-reviewer, operator
IN_REPLY_TO: s1-fidelity-m1/SITREP-implementer-20260703-171028.md
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
SUBJECT: three of four gates green (guide 7/7; VP; m-2 incl. §J2 byte-custody confirmed); m-1 returns exactly one usage finding — fold it as a bounded plan r3 + narrow pair re-review, then I re-route to m-1

**Scoreboard (all verdicts filed in .relays/s1/):**
- m-7 guide APPROVE — `s1-plan-gate/PLAN-REVIEW-planner-20260703-171032.md` (checklist 7/7; your should-fix folds verified; lineage verified).
- master VP APPROVE — `s1-plan-gate/RECONCILE-orchestrator-reviewer-20260703-170942.md`.
- m-2 fidelity APPROVE — `s1-fidelity-m2/SITREP-implementer-20260703-171027.md` (all five questions faithful; §J2 byte-custody formally confirmed).
- m-1 fidelity BLOCKS — `s1-fidelity-m1/SITREP-implementer-20260703-171028.md`, single finding:

**F-M1-1 (fold required):** the m-1 PLAN carry (m-1 :228-229 — credential generation/rotation/
revocation) is acknowledged at design :185 but not landed in plan Task 5 (:120-123), which
asserts only "re-attach same credential = same seat, no re-mint" — no generation/epoch,
duplicate-mint, remint/recycle, or old-credential-rejection invariant or fixture. Required:
ONE explicit lifecycle invariant + E2 fixture in Task 5, in whichever shape is true to your
design (m-1 offered both):
  (a) one active credential generation per seat; remint/recycle atomically replaces the
      binding; old credentials reject before staging; E2 test; or
  (b) if S1 genuinely has no remint/recycle operation: `Mint(existing seat)` returns the
      existing binding or rejects without creating a second credential; E2 test that no
      stale/parallel credential survives.
Note m-1's four specific-question confirmations were all no-finding (DI-2 posture, submit
pipeline, operator address, TOCTOU carry) — this is genuinely the sole gap.

**Ride-along (fold in the same pass, cheap):** the guide's non-blocking nit — design D-7 :112
still keys the crash-window paragraphs by `gate_record_ref` while the r4 paragraph (:115)
makes `(source_kind, source_record_ref)` THE key; add the one-line parenthetical (for a
gate-sourced item, `gate_record_ref` IS `source_record_ref` with `source_kind=gate`).

**Also carry forward (no action now — guide's declared IMPL watch-surfaces, already specified
in the plan; do not weaken them in the fold):** F11 asserts EXACTLY one pivot per mutation
class; P1 captures push frames + tool descriptions, not just bounce/error text.

**Process (bounded — this is a fold, not a re-plan):**
1. Fold F-M1-1 into plan Task 5 (+ design if your chosen shape needs a design line) and the
   nit parenthetical into the design — plan r3 / design r5, fold-logged, committed.
2. Pair-Implementer NARROW re-review of the fold only (the r2 approve stands for everything else).
3. Report back to me; I re-route the revised surface to m-1.implementer for narrow re-review.
4. Guide/VP/m-2 approves are NOT refreshed — the revision touches only the m-1 surface + a
   doc parenthetical; if your fold somehow grows beyond that, stop and relay to me instead.
SCOPE_DIFF still waits until m-1's approve lands; the ledger entry is RECONCILE.md entry 7.

ACTIONS_GIT_REF: wrote this relay + INDEX rows (this file + the guide-verdict copy filed at s1-plan-gate/PLAN-REVIEW-planner-20260703-171032.md) under .relays/s1/ (gitignored); ledger entry committed on main (see git log)
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: pair folds F-M1-1 + nit (plan r3/design r5), narrow pair re-review, report to me; I then send the narrow re-review packet to m-1.implementer.
