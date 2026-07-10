## MERGE-GATE — s4 wire-up integration authorized + tag `s4-close` (operator-directed · VP-confirmed · master-accepted)

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s4-merge-gate
PARENT_DISPATCH_ID: s4-exit-gate
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: yes — this relay carries the human/operator-directed merge grant; the executor runs the four bounded steps and reports
IN_REPLY_TO: .relays/s4/s4-merge-gate/MERGE-GATE-orchestrator-planner-20260705-230526.md
FROM: master.orchestrator-planner
TO: s4-wire.implementer
CC: operator, s4.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-2.planner
REPO: frank
BASE: main
TARGET_BRANCH: main
BRANCH: s4-wire-impl
SUBJECT: integrate the s4 wire-up into main (--no-ff) + annotated tag s4-close — the exit gate is master-accepted (…-231116) and VP-confirmed (…-232713); operator-directed 2026-07-05; four bounded steps, no extras

**Grant basis (the gate is clean at every station).** Master-acceptance `s4-exit-gate/RECONCILE-orchestrator-planner-20260705-231116` (battery 21 ok uncached at `6a23cf0`, F-GATE-2 fix verified, live-host E3 verified at the store) · VP confirmatory pass `s4-vp-confirm/RECONCILE-orchestrator-reviewer-20260705-232713` (`confirm`, no pre-merge blocker) · operator-directed this session. The s4 close-packet's five decisions (`…-230526`) are resolved with its own recommendations: **authorize · executor `s4-wire.implementer` · tag `s4-close` · `--no-ff` · VP pass (done)**. This is the **single** integration authorization — do not also route the s4-orchestrator token path; one grant.

**Why master issues it (not the slice orchestrator this time):** the operator directed master, and `main` is the shared trunk master owns; trunk integration is master's architectural authority with the operator as the human gate. Equivalent-validity to the S2/S3 slice-orchestrator-issued pattern; chosen for directness.

**The four bounded steps (executor: `s4-wire.implementer`; no extra edits, no fix-forward, no push):**
1. **Verify-then-integrate:** on `main` (`a47381a`), `git merge --no-ff s4-wire-impl` (tip `6a23cf0`). No conflicts expected (main advanced by ledger-docs-only since base `28dfa33`). If a real conflict appears, **STOP and relay** — do not resolve under this grant.
2. **Post-integration battery:** `go vet ./...` clean + `go test -count=1 ./...` = all packages ok, uncached, on the integration commit. A red battery **voids** this grant — stop and relay.
3. **Tag:** annotated `git tag -a s4-close` on the integration commit (message may cite this grant + the live-host E3 relay `relay-4a33925b…`).
4. **Execution report:** a MERGE-GATE/SITREP action report with `PARENT_DISPATCH_ID: s4-merge-gate`, the integration commit id + its two parent ids (one being the branch tip), the post-integration battery result, and the tag — the lineage closes against this relay.

Suggested commands (adapt as needed; the executor owns the mechanics):
```
git -C frank checkout main
git -C frank merge --no-ff s4-wire-impl
go -C frank vet ./... && go -C frank test -count=1 ./...
git -C frank tag -a s4-close -m "s4 wire-up closed — operator-authorized, VP-confirmed; first live Claude↔Codex governed relay; Step-1 owed set empty"
```

**Scope of this grant:** exactly `s4-wire-impl@6a23cf0 → main` + tag `s4-close`. Not a push (no remote). Not a deploy (E3 = transport/provenance only; done-state stays `self_reported`). Not authorization for any further integration.

## The grant

DISPATCH MERGE

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s4/s4-merge-gate` — run below.
- Gate chain: master-accept `…-231116`, VP-confirm `…-232713`, s4 close `…-230525`, s4 decision-request `…-230526`. Branch `s4-wire-impl@6a23cf0` battery-verified at the master seat; `main@a47381a` (ledger-docs-only ahead of base `28dfa33`).

ACTIONS_GIT_REF: none — this relay grants authorization only and performs no git action. No `frank/` edit by this seat; cwd is not a git repo (docs workspace). The integration commit + tag are the executor's actions, reported under `PARENT_DISPATCH_ID: s4-merge-gate`.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main `a47381a`, branch `s4-wire-impl@6a23cf0` (pre-integration).
Next requested action: the executor (`s4-wire.implementer`, or the operator acting as the human authority at the terminal) runs the four bounded steps and files the execution report; on the merge + tag, s4 files its close record and master folds S4 into `RECONCILE.md` + dispatches s5.
