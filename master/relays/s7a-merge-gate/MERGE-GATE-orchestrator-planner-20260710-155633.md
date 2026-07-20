## MERGE-GATE — the operator's s7a merge grant, EXECUTOR = m-2.implementer (the original build implementer, operator-named): merge `s7a-colgrain@2bc0763` into the private `frank/main` `--no-ff`, run the serialized battery at the merge commit, push to `frank-dev`, report — exactly these steps, nothing else

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s7a-merge-gate
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator's merge grant is EXERCISED by this dispatch (granted 2026-07-10 against the merge-decision relay `PLAN-…-155040`, VP final approve `…-154754`; the operator named the original implementer as executor); nothing further is requested of the operator
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-dispatch/PLAN-orchestrator-planner-20260710-155040.md
FROM: master.orchestrator-planner
TO: m-2.implementer
CC: operator, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-7.implementer
SUBJECT: you are the operator-named merge executor for s7a — the bounded steps below, in the MAIN checkout `/Users/jack/Programming/harness/frank` (not your worktree); any conflict, failure, or surprise = STOP and report, resolve nothing in-flight

**The grant of record:** the operator granted the s7a merge (2026-07-10, in-session against `PLAN-…-155040`) and named **you — the seat that built the change — as the executor**, per the s4/s6 executor precedent. The VP's final approve (`RECONCILE-…-154754`) is the standing integration verdict; its conditions ride below.

**The bounded execution (exactly these, in order):**
1. In `/Users/jack/Programming/harness/frank` (main checkout; expect `main` clean at `1d3e92c`):
   `git merge --no-ff s7a-colgrain -m "s7a: column-grain any_row gate (F-S7-R2-COLGRAIN closed)"`
2. At the merge commit: `go vet ./...` and the **serialized uncached battery** `go clean -testcache && go test -count=1 -p=1 ./...` — report the **observed sequence honestly** per the standing evidence rule (`OI-S7A-CLOSE-ONCE-RACE` is open; if you also run a parallel attempt, report each attempt's outcome, never a flattened summary).
3. Push the merged `main` to the private remote: `git push origin main` (origin = `frank-dev`; the pre-push guard permits it — the public repo is untouched by construction).
4. **No tag** — s7a is a micro-lane, not a slice close (s7's own close comes later). **Do not delete the `s7a-colgrain` branch or touch either worktree.** No other ref, file, or config action.
5. **Report back** (your MERGE-GATE execution report, TO master, CC the operator + the VP + both planners): the merge commit SHA · first-parent = `1d3e92c` check · the vet + battery outputs (sequence-honest) · `git status --short` clean · the push confirmation (`git ls-remote origin main` matches local).

**Conditions preserved into this trail (the VP's requirement):** the two s8 pre-flight obligations stand undiminished by the merge — the **genesis condition** (dogfood genesis from these exact `s7a-fieldspec-v5` bytes, member SHA `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`, with the recorded composite digest) and **`OI-S7A-CLOSE-ONCE-RACE`** (owner m-7; blocks live MCP dogfood startup until the idempotent-close acceptance passes).

**After your report:** master verifies at the merge commit, folds the record, and dispatches the s7 pair's rebase + row-3 fold (the m-2+m-4 re-confirms and the s7 integration follow; none of that is authorized here).

ACTIONS_GIT_REF: none — no git action by this relay (an authority grant only; the merge/push are the executor's from here, bounded as above).
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` clean at `1d3e92c` at dispatch time; `s7a-colgrain` clean at `2bc0763`; cwd is not a git repo (docs workspace).
