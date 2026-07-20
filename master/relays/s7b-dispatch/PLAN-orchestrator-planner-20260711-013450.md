## PLAN — MERGE DECISION to the operator: s7b `s7b-close-once@e155aa6` → private `frank/main`, `--no-ff`, no tag — VP-approved first pass (`…-013104`); the merge closes three registered items and lifts the s8 live-channel gate; your grant names the executor

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7b-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — this relay IS the merge-decision request; nothing merges until you grant
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7b-dispatch/RECONCILE-orchestrator-reviewer-20260711-013104.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-7.planner, m-7.implementer, m-1.implementer
SUBJECT: requesting your merge grant for s7b — three commits (`a2a6966` the idempotent close · `5c678b4` fixture hardening · `e155aa6` the granted delta) into the private `frank/main@2e1b4f0`; on the merge: `OI-S7A-CLOSE-ONCE-RACE` + all three FLAKE-SOCKET-PAR members + CRASHPOINT-KILL-RETURN close, and the s8 dogfood's live-MCP-channel gate LIFTS; the release tree and public repos untouched

**The decision:** merge the s7b lane — the double-close race fixed at its owner (`sync.Once`, both client and server `done` channels), the three flaky-startup surfaces hardened by cache + honest deadlines, and the crashpoint's kill-then-return window closed by one line — into `main`. `--no-ff`, **no tag** (a micro-lane, the s7a precedent).

**The chain (point-not-restate):** your ruling "A" (B10 third application) → the planner-first lane with delegated dispatch → the VP's visibility approval with two conditions → the pair's r2 loop → the model stop on two blockers → the mechanism diagnosis → my bounded three-surface grant → the delta loop → pair APPROVE → m-1's five-point lifecycle CONFIRM → **VP integration APPROVE first-pass (`…-013104`)**, its own runs green across every mode including the formerly-flaky default-parallel.

**The honest test record at `e155aa6`:** three consecutive default-parallel full suites green in one uninterrupted sequence (implementer) · independent parallel greens at the pair, master, and the VP (25 ok / 0 FAIL each, file-captured) · the serialized uncached battery green (VP) · both race stations `-race -count=20` green at three stations · vet clean. No parallelism pin was needed; none was checked in.

**Execution options:**
1. **You run it** (in `/Users/jack/Programming/harness/frank`): `git merge --no-ff s7b-close-once -m "s7b: idempotent channel close + deterministic crashpoint + startup hardening (OI-S7A-CLOSE-ONCE-RACE, FLAKE-SOCKET-PAR, CRASHPOINT-KILL-RETURN closed)"` then `git push origin main` — I verify at the merge commit and fold the record.
2. **You grant and name the executor** — the precedent says the original implementer (**m-7.implementer**); on your word I cut the MERGE-GATE dispatch with `HUMAN_MERGE_AUTHORIZATION` carrying your grant at grant time, the bounded steps (merge with the pinned message · vet + the serialized uncached battery FILE-captured at the merge commit · push `main` to `frank-dev` · no tag · report).

Next requested action: your grant — e.g. "granted, m-7.implementer executes" — or option 1.

ACTIONS_GIT_REF: none — the decision request only; no merge performed.
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` clean at `2e1b4f0` (tag `s7-close`); the s7b worktree at `e155aa6` clean (untracked `.relays/s7b/` only); cwd is not a git repo (docs workspace).
