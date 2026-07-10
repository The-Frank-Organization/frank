## COORD — gate rulings in: deliverable ACCEPTED at master; deviation 1 concurred (owed item ledgered); deviation 2 = ONE pre-concurred ~10-line fixture — fold it and the gate closes

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-core-impl
PARENT_DISPATCH_ID: s1-exit-gate
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — after this fold the ONLY remaining S1 gate is the operator's merge decision (optional VP confirmatory pass first, operator's call)
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: s1-core.implementer, s1.orchestrator-reviewer, operator
IN_REPLY_TO: s1-exit-gate/SITREP-planner-20260703-200827.md
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
SUBJECT: fold the guide's pre-concurred deviation-2 fixture (spec verbatim below) + optionally the dedupe-grain robustness note; deviation-1 owed item is already ledgered by me; cite the guide relay + your fold commit in the gate-close record — no second guide round-trip

**Rulings (verdict copies filed in .relays/s1/s1-exit-gate/; ledger entry 11):**
- The CTO **ACCEPTED the S1 charter deliverable at the master seat** after an independent
  uncached re-run (15 ok, their own E2) — three independent battery verifications now on
  record. Scope discipline and honesty framing explicitly confirmed.
- **Deviation 1 (F11 breadth): CONCURRED, S1-sufficient** — the guide verified on code that
  every charter window is covered, the cross-product is structurally redundant at S1, the
  registry-live check holds, and the at-most-one watch-surface is satisfied (including
  chasing — and clearing — a suspected outbox-counter hole: correct-by-construction, "worth
  keeping"). The condition (typed owed-item record for the un-run class×point sweep,
  S2-dispositioned) is **already written by me** — RECONCILE.md entry 11, `OI-S1-F11-SWEEP`.
  Nothing for you to fold on this one.
- **Deviation 2 (C7 mid-Complete): NARROW BOUNCE with PRE-CONCURRENCE.** The composition
  argument fails at one point the guide proved from code: `gate.Complete` converges from
  partial state, but the partial state a mid-Complete crash actually leaves (park committed,
  outbox missing) is exercised by NO test — the double-run proves full-state idempotence
  only. The property is currently true-by-code-reading; the gate's standard is executable
  claims, and a per-source-dedupe refactor would regress it silently.

**Required fixture (the guide's spec, verbatim contract — deterministic, no crash machinery,
~10 lines in `internal/gate/derived_test.go`):**
1. Commit a gate-bearing accepted record + its park record ONLY (no outbox item).
2. Run `gate.Complete`.
3. Assert exactly one outbox item appears AND the park record is not duplicated.
Optional mirror leg: outbox record present, park missing → Complete → park appears once, no
outbox dup. On green + spec-match, deviation 2 is CONCURRED by pre-agreement — cite the guide
relay (`s1-exit-gate/SITREP-planner-20260703-200827.md`) + your fold commit in your gate-close
report; no second guide pass unless the fixture diverges from this spec.

**Optional ride-along (guide's non-blocking robustness note — fold now if convenient):**
`completeOutbox` dedupes on the projection file (`os.Stat outbox/<item>.json`) rather than the
canonical `outbox-<item_id>` record. Safe today (recovery rebuilds projections at step 2
before derived-work completion at step 5; `store.Commit` rejects duplicate ids fail-safe).
Recommended: switch to a canonical-record scan symmetric with `completePark`, OR a code
comment pinning the rebuild-before-Complete ordering dependency. Your call which; if you fold
the canonical-scan variant, the new fixture must still pass unchanged.

**Process:** this is a REVIEW-FOLD-class edit on main per your established fold discipline —
FOLD_SCOPE pre-filed (expected rows: `internal/gate/derived_test.go` -> in; optionally
`internal/gate/derived.go` -> in for the robustness fold), fold commit, battery re-run, your
gate-close report back to me citing guide relay + commit. On that report verified, the S1
gate CLOSES; the operator's merge decision (with an optional VP confirmatory pass first) is
the only remaining S1 gate.

ACTIONS_GIT_REF: wrote this relay + INDEX rows + filed the two master verdict copies under .relays/s1/ (gitignored); ledger entry 11 committed on main (see git log)
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: pair folds the fixture (+ optional robustness note), re-runs the battery, returns the gate-close report; I verify, declare the gate closed, and hand the merge question to the operator.
