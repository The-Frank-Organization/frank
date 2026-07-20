## MERGE-GATE — the operator's s7 merge grant, EXECUTOR = m-7.implementer (operator-named): merge `s7-inv-catalog@5e6bf83` into the private `frank/main` `--no-ff`, TAG `s7-close` at the merge commit, run the serialized battery, push main + the tag to `frank-dev`, report — exactly these steps, nothing else

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s7-merge-gate
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator's merge grant is EXERCISED by this dispatch (below, in the recognized form); nothing further is requested of the operator
GRILL_REQUIRED: no
HUMAN_MERGE_AUTHORIZATION: granted — the operator's in-session s7 merge grant of 2026-07-10 ("granted, tag it, and send it to m-7.implementer to merge"), issued against the merge-decision relay `PLAN-orchestrator-planner-20260710-183959.md` under the VP round-3 approval `RECONCILE-orchestrator-reviewer-20260710-183819.md`; executor = m-7.implementer; scope = `s7-inv-catalog@5e6bf83` → private `frank/main`, `--no-ff`, tag `s7-close` at the merge commit, push main + tag to `frank-dev` only
IN_REPLY_TO: master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-183959.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-2.planner, m-2.implementer
SUBJECT: you are the operator-named merge executor for the s7 slice close — the bounded steps below, in the MAIN checkout `/Users/jack/Programming/harness/frank` (not the worktree); any conflict, failure, or surprise = STOP and report, resolve nothing in-flight

**The grant of record is in this relay's `HUMAN_MERGE_AUTHORIZATION` header** (the recognized form, at grant time, earlier than any execution claim — the s7a-merge-gate convention applied). The VP's approval scope and conditions (`…-183819`) bind below.

**The bounded execution (exactly these, in order):**
1. In `/Users/jack/Programming/harness/frank` (main checkout; expect `main` clean at `54420db`, equal to `origin/main`):
   `git merge --no-ff s7-inv-catalog -m "s7 CLOSE: INV-CATALOG — ten named executable laws (test/invariants); F-S7-R2-COLGRAIN caught through the fence, fixed in s7a, guarded by the named law"`
2. **Tag the merge commit:** `git tag s7-close` (the `s1-close`…`s6-close` pattern continues; private dev repo only per the release-separation ruling).
3. At the merge commit (the VP's condition): `go vet ./...` and the **serialized uncached battery** `go clean -testcache && go test -count=1 -p=1 ./...` — **capture the output to a file and report counts from the file** (the pipe-count lesson); report the observed sequence honestly per the standing rule (`OI-S7A-CLOSE-ONCE-RACE` is open — expected shape at the merge commit: 25 tested packages ok + 2 no-test-files, exit 0).
4. Push: `git push origin main` and `git push origin s7-close` (origin = the private `frank-dev`; the pre-push guard permits it; the public repo is untouched by construction).
5. **Do not delete the `s7-inv-catalog` branch or touch either worktree.** No other ref, file, or config action.
6. **Report back** (your MERGE-GATE execution report, TO master, CC the operator + the VP + m-7.planner): the merge commit SHA · first-parent = `54420db` and second-parent = `5e6bf83` checks · `git tag --points-at HEAD` = `s7-close` · the vet + battery outputs (file-captured, sequence-honest) · the push confirmations (`git ls-remote origin refs/heads/main refs/tags/s7-close` matching local) · the exact final status with the tracked-vs-untracked distinction.

**Conditions preserved into this trail (unchanged by the merge):** the s8 §7-pinning/genesis carry (dogfood genesis from the exact `s7a-fieldspec-v5` bytes, member SHA `1ef6abab…2485`, with the recorded composite digest) · **`OI-S7A-CLOSE-ONCE-RACE`** (owner m-7; blocks live MCP dogfood startup until the idempotent-close acceptance passes) · `FLAKE-SOCKET-PAR` · the S7A-TRAIL-FINDINGS ledger.

**After your report:** master verifies at the merge commit and runs the s7 CLOSE fold (RECONCILE §s7 · dashboard · ROADMAP · the PLAYBOOK friction log); the pre-s8 package follows separately. Nothing beyond step 6 is authorized here.

ACTIONS_GIT_REF: none — no git action by this relay (an authority grant only; the merge/tag/push are the executor's from here, bounded as above).
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` clean at `54420db` at dispatch time; the s7 worktree at `5e6bf83` untouched by this relay; cwd is not a git repo (docs workspace).
