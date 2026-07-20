## FOLD_SCOPE — s11 master-ruled engine-v4 cadence re-home

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s11-build-fold-v4-scope
PARENT_DISPATCH_ID: s11-build-fold-v4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — merge remains held for the reissued decision at the post-fold head
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-fold-v4/REVIEW-FOLD-planner-20260714-171200.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-7.implementer, m-6.planner, m-3.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: pre-edit scope for the master-ruled v4 re-home — carry m-7 F1-F6 verbatim, add A-G cadence/version fixtures RED-first, and change no non-config or non-cadence byte

FOLD_SCOPE:
- frank/internal/config/config.go -> in
- frank/test/fixtures/s11_cadence_test.go -> in
- frank/test/fixtures/ -> in
- frank/.relays/s11/fold-v4-red-green.md -> in
- frank/.relays/s11/mechanical-tables.md -> in
FOLD_SCOPE_RESULT: all-in

The current `e86644d` bytes match every OLD anchor in m-7's directly addressed owner spec `master/relays/s11-build-escalate-config-lock/SITREP-planner-20260714-171302.md`. The implementation is limited to F1–F6: reader/descriptor/transition ceilings become 4, the cadence key moves from the v3 arm to v4, the cadence fixture stamps v4, and A–G prove restored v3 rejection, v4 acceptance/optional residency, adjacent-forward transition, rollback/skip rejection, v5 reader-ceiling rejection, and unchanged cadence behavior. `validResummonCadenceShape`, `ResummonCadenceDelays`, the struct field, supply machinery, `main.go`, and every other slice surface remain byte-untouched.

ACTIONS_GIT_REF: none — pre-edit FOLD_SCOPE only; no frank edit yet; s11-comms-thicken clean at e86644ddf10ca9bbdc4c098f443ad3eab73c4e20
FINAL_GIT_STATUS_SHORT: none — clean s11-comms-thicken worktree at e86644ddf10ca9bbdc4c098f443ad3eab73c4e20
Next requested action: s11.implementer adds A–G tests and observes RED before carrying F1–F6, then runs the bounded targeted verification, commits and pushes the existing PR branch, and reports for m-7 countersign. No merge authority is requested or implied.
