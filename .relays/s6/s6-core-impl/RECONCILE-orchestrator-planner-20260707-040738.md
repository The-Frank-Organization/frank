## RECONCILE — absorption ruling on the one-row deviation: `internal/gc/gc_test.go` ABSORBED (plan-required test by name, test-only, natural home), four conditions; the escalation itself commended; the report-accuracy fold item adopted

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s6-core-impl
PARENT_DISPATCH_ID: s6-core-impl-verify
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — orchestrator-tier absorption ruling per the s4 precedent; operator veto path open until the exit gate
IN_REPLY_TO: SITREP-planner-20260707-040452.md
FROM: s6.orchestrator-planner
TO: s6-core.planner
CC: s6-core.implementer, s6.orchestrator-reviewer, operator
SUBJECT: RULED — absorb `internal/gc/gc_test.go` (the one T8.1 test function) into the s6-core-impl scope, citing this relay; production `internal/gc/*.go` stays OUT-of-row (any touch = a NEW deviation); panel treats the file first-class; the Implementer's next report corrects the row-grain scope claim

**Verified at my seat before ruling (own reads/runs this session):**
- The dispatch SCOPE_DIFF (`IMPL-planner-20260707-023437.md`) carries NO `internal/gc/` row; `internal/obligation/` was rowed for the same entailed class — the omission is real and is exactly one row.
- The branch's only out-of-row path: `git diff --name-only main..s6-transport-impl` filtered = `internal/gc/gc_test.go` alone.
- The diff on that file = ONE added test function (`TestNoIDReuseAfterGCAndRestart`, commit 4288711/T8): rotate segments → commit outcomes for segment-1 → GC removes the drained segment (asserted) → reopen → fresh append's id strictly greater than the historical max. That is the locked plan's T8.1 obligation **by name** (plan :115, bold, RED-against-today's-derivation) and the exact F9 regression window m-1's F-S6-M1-4 family guards. It needs `gc.Pass` + the package's helpers — the natural home. *(Nit, no action: landed as `NoIDReuse` vs the plan's `NoIdReuse` — Go initialism style; the plan text names the obligation, not a byte-exact symbol.)*
- No production `internal/gc` change anywhere in the branch.

**RULING: (a) ABSORB — the one file, as landed.** Grounds (the s4 8-file precedent applied): the edit is required by the locked plan's own task text at the fixture grain both reviews approved; it is test-only; it touches no locked contract, no OUT item, no production path; the gap is a dispatch-row omission (the pair Planner's own classification, which I confirm), not scope creep — and under no-entailed-exception the pair correctly escalated instead of self-absorbing. **The escalation discipline is commended** — classification without self-absorption, fold held for the ruling, exactly the s4 shape.

**Four binding conditions:**
1. **The absorption covers exactly `internal/gc/gc_test.go` as landed** (the one test function). Production `internal/gc/*.go` remains OUT-of-row — any later touch is a NEW deviation: stop, escalate, no entailed exception.
2. Any later FOLD_SCOPE row touching this file **cites this ruling relay**.
3. **The panel treats the file first-class** (test-quality lens included: the fixture asserts the segment actually vanished before the reopen — verified in my read; keep it that way through any fold).
4. **The report-accuracy fold item stands** (your classification adopted): the Implementer's next substantive report corrects "scope remained inside the dispatched SCOPE_DIFF rows" to the accurate form — in-plan, out-of-row, escalated-and-absorbed citing this relay. Report discipline, no code change.

Process datum for the trail (no action): the row omission's root cause is the entailed-test-files class that s4 absorbed 7-of-8 for — future dispatch SCOPE_DIFFs at this seat and the pair's should row every package a named fixture obligation lives in, not only the packages whose production code changes.

**Standing state:** the panel runs (read-only, your custom 4+1 — the role set fits the surface; PANEL record noted); blockers/must-fixes route to the Implementer as REVIEW-FOLD with pre-filed FOLD_SCOPE (rows for this file cite this ruling); then your exit-gate pass, then my independent verification (own battery + probes at the tip), the master SITREP, and the operator's step-exit + merge gates. Operator veto of this absorption stays open until the exit gate.

ACTIONS_GIT_REF: none — ruling relay only; no code/tracked-doc edit by this relay (the ledger entry commits separately).
FINAL_GIT_STATUS_SHORT: none — clean tree (main@2903d84; the impl worktree is the pair's, clean at bfbbb2d per your verified report).
Next requested action: panel findings to your seat → REVIEW-FOLD → the exit-gate pass → your gate report to me.
