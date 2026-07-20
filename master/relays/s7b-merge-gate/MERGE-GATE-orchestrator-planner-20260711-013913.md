## MERGE-GATE — the operator's s7b merge grant, EXECUTOR = m-7.implementer: merge `s7b-close-once@e155aa6` into the private `frank/main` `--no-ff` with the pinned message, NO tag, run the serialized battery file-captured, push to `frank-dev`, report — exactly these steps, nothing else

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s7b-merge-gate
PARENT_DISPATCH_ID: s7b-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator's merge grant is EXERCISED by this dispatch (below, in the recognized form); nothing further is requested of the operator
GRILL_REQUIRED: no
HUMAN_MERGE_AUTHORIZATION: granted — the operator's in-session s7b merge grant of 2026-07-11 ("granted", accepting the recommended executor), issued against the merge-decision relay `PLAN-orchestrator-planner-20260711-013450.md` under the VP integration approval `RECONCILE-orchestrator-reviewer-20260711-013104.md`; executor = m-7.implementer; scope = `s7b-close-once@e155aa6` → private `frank/main`, `--no-ff`, NO tag, push `main` to `frank-dev` only
IN_REPLY_TO: master/relays/s7b-dispatch/PLAN-orchestrator-planner-20260711-013450.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-1.implementer
SUBJECT: you are the operator-named merge executor for s7b — the bounded steps below, in the MAIN checkout `/Users/jack/Programming/harness/frank` (not the worktree); any conflict, failure, or surprise = STOP and report, resolve nothing in-flight

**The grant of record is in this relay's `HUMAN_MERGE_AUTHORIZATION` header** (the recognized form, at grant time, earlier than any execution claim).

**The bounded execution (exactly these, in order):**
1. In `/Users/jack/Programming/harness/frank` (expect `main` clean at `2e1b4f0`, equal to `origin/main`):
   `git merge --no-ff s7b-close-once -m "s7b: idempotent channel close + deterministic crashpoint + startup hardening (OI-S7A-CLOSE-ONCE-RACE, FLAKE-SOCKET-PAR, CRASHPOINT-KILL-RETURN closed)"`
2. **No tag** (a micro-lane, the s7a precedent).
3. At the merge commit: `go vet ./...` and the **serialized uncached battery** `go clean -testcache && go test -count=1 -p=1 ./...` — **capture the output to a file and derive counts from the file**; report the observed sequence honestly (expected shape: 25 tested packages ok + 2 no-test-files, exit 0). A parallel-suite attempt is welcome as an additional leg (the mode is now expected green) but is not required by this gate.
4. Push: `git push origin main` (origin = the private `frank-dev`; the pre-push guard permits it; the public repo untouched by construction).
5. **Do not delete the `s7b-close-once` branch or touch any worktree.** No other ref, file, or config action.
6. **Report back** (your MERGE-GATE execution report, TO master, CC the operator + the VP + m-7.planner + m-1.implementer): the merge commit SHA · first-parent = `2e1b4f0` and second-parent = `e155aa6` checks · `git tag --points-at HEAD` = empty · the vet + battery outputs (file-captured) · the push confirmation (`git ls-remote origin refs/heads/main` matching local) · the exact final status with the tracked-vs-untracked distinction.

**Effects on the register (master folds them on your report — stated here so the trail carries them):** the merge closes `OI-S7A-CLOSE-ONCE-RACE` (its live-channel gate LIFTS), all three `FLAKE-SOCKET-PAR` members, and `CRASHPOINT-KILL-RETURN`. It does not touch the s8 design/genesis/config gates or any operator gate.

ACTIONS_GIT_REF: none — no git action by this relay (an authority grant only; the merge/push are the executor's from here, bounded as above).
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` clean at `2e1b4f0` (tag `s7-close`) at dispatch time; the s7b worktree at `e155aa6` untouched by this relay; cwd is not a git repo (docs workspace).
