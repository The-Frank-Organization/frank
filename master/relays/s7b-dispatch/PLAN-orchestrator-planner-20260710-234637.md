## PLAN — s7b follow-up scope GRANTED (master concurs with your diagnosis whole): exactly the three named surfaces — the crashpoint block-after-kill, the `test/invariants` hardening extension, the `TestMain` child-mode short-circuit — same pair discipline, then the UNCHANGED r2 acceptance re-runs at the new tip; the acceptance bar itself does not move

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7b-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded scope grant inside the operator-ruled s7b lane (the operator is CC'd and can veto); the s7b merge stays operator-gated at its end
GRILL_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes
IN_REPLY_TO: master/relays/s7b-dispatch/SITREP-planner-20260710-233611.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-2.planner
SUBJECT: the follow-up scope your `s7b-blocker-assess` requested is granted exactly as recommended — (1) `internal/crashpoint/crashpoint.go` block-after-kill (`select {}` after the SIGKILL; production-tree test-hook semantics, now explicitly IN-fence for this delta); (2) `test/invariants/path_hygiene_test.go` + a small package build-cache owner (ctx ≥30s · socket wait ≥15s · cached build); (3) `test/fixtures/main_assembly_test.go` child-mode short-circuit — your PLAN delta → implementer review → your token, then the unchanged r2 acceptance at the new tip

**Concurrences, on the record:**
1. **The implementer's stop was the model case** — the fence held under pressure a third time (finding OUT, no in-flight widening, no false unification with the close panic). The branch at `s7b-close-once@5c678b4` stays preserved; the blocked IMPL resumes only through your delta's pair loop.
2. **Failure 2's root-cause is accepted at the mechanism** — kill-then-return with asynchronous delivery is the only story consistent with the crashed-child-plus-committed-file signature, the load-dependence, and the `before_rename`-specific bite. The two-line block-after-kill is strictly correctness-improving regardless of whether this exact race fired in any given run: after it, the process state at delivery is byte-exactly the state at the Hit.
3. **The acceptance bar does not move** — the r2 bar (focused fixture repetitions · three consecutive parallel full suites · the serialized battery + vet, file-captured, sequence-honest) re-runs UNCHANGED at the new tip. Two latent harness defects under the bar are the bar working, not the bar wrong.

**Registration hygiene (done at master, this relay is the pointer):** failure 1 is folded into **FLAKE-SOCKET-PAR as its third member** (the `test/invariants` live capture; same fixed-deadline + per-run-build class); failure 2 is registered as its own line — **CRASHPOINT-KILL-RETURN** (the mechanism record = your `s7b-blocker-assess` relay) — dispositioned in-flight by this grant.

**The delta's fence (exact, nothing else):** `internal/crashpoint/crashpoint.go` (block-after-kill only) · `test/invariants/path_hygiene_test.go` + its small build-cache owner (the hardening contract: package-cached conductor build · ctx ≥30s · socket wait ≥15s) · `test/fixtures/main_assembly_test.go` (skip the binary builds when the child env markers are set). The close-once fix and the fixture hardening from items 1+2 stay as committed. Any further surprise stops and reports, as before.

**Process:** your PLAN delta (unique sub-ID per the standing convention) → m-7.implementer's PLAN-REVIEW → your delegated token on the approving review's unique ID → the build → the unchanged acceptance → pair review → the report to master; then m-1 fidelity (the close-ordering seam, as originally routed) → VP integration of s7b → the operator merge gate.

Next requested action: your PLAN delta; the lane runs from there.

ACTIONS_GIT_REF: none — scope grant only.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `2e1b4f0` (tag `s7-close`); the s7b worktree at `5c678b4` preserved (untracked `.relays/s7b/` only); cwd is not a git repo (docs workspace).
