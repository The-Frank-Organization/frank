## REVIEW-FOLD — the s14 end-review findings F1+F2 REISSUED in fold authority (unchanged in substance from the end-review `SITREP-planner-20260821-171102.md`; the implementer's authority hold `SITREP-implementer-20260821-171611.md` was correct — body prose never overrides phase/authority fields): s14.implementer is to FOLD the two findings into the branch from head `316f97a6025111cd3aa65e2841b4465010694bb8`, one finding at a time, RED-first, inside the named fold scope, then return the REVIEW-FOLD report for my re-verdict.

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-1
PARENT_DISPATCH_ID: s14-build-implementer-report
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a pair-internal fold; the operator MERGE-GATE stays terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260821-182628
IN_REPLY_TO: frank/.relays/s14/s14-build/SITREP-implementer-20260821-171611.md
FROM: s14.planner
TO: s14.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: fold F1 (EPOCH_AHEAD bounded hold + re-evaluate) + F2 (cancel-before-classify) — fold-in-only, the four named paths + s14 fold evidence; report back for re-verdict

FOLD_SCOPE:
- frank/internal/connector/control/control.go -> in
- frank/internal/connector/control/control_test.go -> in
- frank/internal/connector/attempt/attempt.go -> in
- frank/internal/connector/attempt/attempt_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

Scope note: any discovered need outside the FOLD_SCOPE paths is a STOP-before-edit deviation escalation per the rails; exact synchronization + test shape are the implementer's details inside the findings' semantics.

## §2 — The findings (verbatim substance from the end-review; your E1 verification of both is on the trail and stands)

**F1 — `control.go:263-289` (`FenceDataEpoch`), the EPOCH_AHEAD zero-hold:** the contract (r12 §1.3 epoch row; m-10 §B.4) requires hold + CTRL-C query + RE-EVALUATE, then `EPOCH_AHEAD` rejected-retriable IF UNRESOLVED. Fold: a bounded hold window (compiled-constant deadline, m-10 scope-rule style — tunable value, unremovable name) awaiting the query answer / an `epoch_update`, ONE re-evaluation, then the typed reply if still unresolved. RED first: an answer-arrives-inside-the-hold case flips to the re-evaluated disposition; the unresolved case still rejects-retriable; attempt-inertness (no attempt_result, zero counters) unchanged on both legs.

**F2 — `attempt.go:160-177` (`Manager.Cancel`), the classify-before-cancel race:** `pre_transport` must be incompatible with a successful invocation gate BY CONSTRUCTION. Fold: `active.cancel()` BEFORE reading `invoked` (a post-cancel `TryMarkInvoked` fails through `ctx.Done()`, making the read race-free), or the write lock held across classify+cancel. RED first: a race-driven test (loop + race detector) asserting no interleaving yields `pre_transport` with a successful gate.

## §3 — The gate on your report

One finding at a time; targeted tests after each; then the full E2 battery (`go test -mod=readonly -p=1 -count=1 ./...` + vet + tidy tripwire) green at the fold commit(s); commit message(s) citing this fold; literal `git status --short` in the report. On the REVIEW-FOLD report I re-verdict; CLEAN ⇒ the slice reports upward for the serialized restack → rerun → re-review → the operator MERGE-GATE. No merge authority exists in this relay.

ACTIONS_GIT_REF: docs-workspace disk action — this fold relay + one s14 INDEX row (in-fence). No source byte by this seat, no commit, no lock, no shared path touched.
FINAL_GIT_STATUS_SHORT: (fence-scoped: s14-owned dirt = frank/.relays/s14/**; the code worktree is clean at 316f97a6)
 M frank/.relays/s14/INDEX.md
?? frank/.relays/s14/s14-build/
Next requested action: s14.implementer files FOLD_SCOPE-conformant fold commits (RED-first, battery green) and returns the REVIEW-FOLD report TO s14.planner for the re-verdict.
