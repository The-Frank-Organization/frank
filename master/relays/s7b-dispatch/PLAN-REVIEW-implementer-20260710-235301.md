## PLAN-REVIEW - s7b blocker delta approved; all four paths remain inside the master-granted fence

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s7b-plan-delta-review
PARENT_DISPATCH_ID: s7b-plan-delta
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - pair Planner may issue the delegated delta implementation token; fidelity, VP integration, and merge remain downstream
DESIGN_RECORD_KIND: audit-record
PLAN_LOCK_ID: s7b-close-once-plan-delta
IN_REPLY_TO: master/relays/s7b-dispatch/PLAN-planner-20260710-234943.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: verdict approve - crashpoint block-after-kill, invariants hardening/child short-circuit, and fixture child short-circuit are mechanically all-in; unchanged r2 acceptance remains binding

PLAN_REVIEW_VERDICT: approve

### Review

- The delta answers the two observed blockers at their mechanisms without weakening acceptance or introducing a parallelism pin.
- `internal/crashpoint.Hit` currently calls `SIGKILL` and can return to the immediately following mutation while signal delivery is pending. Blocking forever after the kill syscall makes execution-past-Hit impossible; the process still dies at the same named/count-selected crashpoint. The name list, target parsing, count map, trace, and public API remain untouched.
- The accepted proof shape is proportionate: existing F10, s2 crash-window, and s7 pivot consumers exercise the changed hook, while the unchanged full batteries detect any broader crash semantics regression.
- `test/invariants` has the same per-run build and fixed-deadline class already hardened in fixtures. A package-process cached binary, 30-second live-capture context, and 15-second socket wait are the exact master-granted extension.
- `FRANK_S7_PIVOT_CHILD` is the only invariants child marker. Its test re-execs the test binary and enters an early child branch before parent orchestration; the new owner-file short-circuit prevents recursive conductor builds and is inside the granted build-cache-owner fence.
- The fixture child-marker census is complete at `5c678b4`: `FRANK_F10_CHILD`, `FRANK_F11_CHILD`, `FRANK_F11_TRACE_CHILD`, and `FRANK_F9_CHILD`. Each marker belongs to an env-guarded re-exec branch. Skipping package binary builds in those children removes amplification while the loud cached-binary error preserves the assumption as a tripwire.
- Items 1 and 2 remain committed and unchanged. The r2 race, named-fixture, three-consecutive-parallel, serialized uncached, vet, file-capture, and sequence-honesty requirements all re-run at the delta tip.

### Mechanical scope diff

SCOPE_DIFF:
- internal/crashpoint/crashpoint.go -> in
- test/invariants/path_hygiene_test.go -> in
- test/invariants/testmain_test.go -> in
- test/fixtures/main_assembly_test.go -> in
SCOPE_DIFF_RESULT: all-in

All rows are explicitly covered by `PLAN-orchestrator-planner-20260710-234637.md`. The invariants child short-circuit resides in the granted small build-cache owner file and is not a fifth surface.

### Gates and residuals

- No hard trigger or unresolved operator decision blocks the delegated delta. Master has accepted the crashpoint mechanism record and explicitly granted this small-tier fence.
- No red-first delta is owed under the master grant; existing focused crash consumers plus the full unchanged acceptance are binding.
- Any additional child marker, harness failure, source path, acceptance change, or pin is out of scope and must stop-and-report.
- A failed `syscall.Kill` is not a newly handled contract in this delta; implementation must follow the granted exact block-after-kill shape and preserve current platform assumptions.

Not authorized / not done: no worktree/source/test edit, commit, push, resumed implementation, pair verdict, fidelity, merge, deployment, or live verification. This review contains no implementation token and grants no merge authority.

ACTIONS_GIT_REF: wrote this approving PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; preserved worktree remains `s7b-close-once@5c678b4` with only untracked `.relays/s7b/`
FINAL_GIT_STATUS_SHORT:
## s7b-close-once
?? .relays/s7b/
Next requested action: m-7.planner may issue `s7b-impl-dispatch-2`, parenting to this unique review ID, with the exact all-in delta scope and bare implementation token. Delta implementation remains blocked until that relay exists.
