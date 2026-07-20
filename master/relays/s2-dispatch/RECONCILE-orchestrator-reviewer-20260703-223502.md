## RECONCILE -- VP review of S2 dispatch

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s2-dispatch
PARENT_DISPATCH_ID: s2-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- requested review only; operator has already directed S2 sequencing per the dispatch
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, s2.orchestrator-planner, m-7.planner, m-1.planner, m-1.implementer, m-7.implementer
IN_REPLY_TO: s2-dispatch/PLAN-orchestrator-planner-20260703-223146.md
SUBJECT: S2 dispatch review -- revise for dispatch-root lint; scope/routing substance otherwise sound

VERDICT: revise

## Required fix

F-S2-1 -- The dispatch is exact-file lint-clean but not dispatch-root lint-clean.

Evidence:
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s2-dispatch/PLAN-orchestrator-planner-20260703-223146.md` -> OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s2-dispatch` -> FAIL:
  `PLAN-orchestrator-planner-20260703-223146.md: relay claims a merge/merge commit without an earlier MERGE-GATE authorization relay with the same DISPATCH_ID`.

Required change: file a superseding `s2-dispatch` PLAN copy that is dispatch-root lint-clean, or obtain an explicit operator waiver in the relevant authority trail. Cheapest likely fix is to reword or fence the S1 close/baseline and separate human-gate prose that trips the root-mode claim detector. Because this is a new authority-bearing dispatch root, do not leave it dirty by inheriting the S1 trail-cleanliness waiver; that waiver was scoped to the S1 close trail.

## Substantive review

No semantic blocker found in the S2 dispatch after the lint issue above is corrected.

1. The guide refinement is acceptable. The Step-1 plan listed S2 guide as m-1 with m-7 on the loop, but the live S2 content is dominated by the engine substrate: recovery phases 0-4, durable FIFO, GC/genesis, and the owed-item projection. The locked m-7 domain owns that engine substrate, while m-1 remains consulted and keeps fidelity authority for store-record/API touches. This preserves the m-1 contract boundary rather than reassigning it.

2. The S2 scope maps to locked C4 engine surfaces. `master/ARCHITECTURE.md` §C4.1 names durable FIFO, phases 0-4 recovery, genesis/GC, and derived projections; the m-7 lock package names the same recovery/FIFO/GC substrate. The owed-item projection also matches the Step-1 kickoff's materialize-first rule.

3. The S1 close prerequisite is present. The cited S1 close record exists and lints exact-file OK; `frank` has `s1-close` pointing at f0dcb85 and current HEAD at 674c844 with a clean tracked status. `master/RECONCILE.md` and `master/README.md` both record S1 closed and S2 dispatched.

4. The MCP live-adapter / fuller FieldSpec deferral is honestly stated, not hidden. The dispatch names the work as deferred due no current live testbed and keeps it out of S2 scope. That matches the live dashboard and avoids overclaiming S1/S2 as live agent integration.

5. The F2 delegated-dispatch conditions are the right shape for S2+. The relay keeps S2 off the bootstrap guide+VP plan gate and names escalation triggers: scope or boundary deviation, hard trigger, cross-slice collision, locked-contract touch, or design-of-record amendment. That matches the VP-approved Step-1 r2 model.

## Watchpoints For The Superseding Dispatch

- Keep m-1's authority explicit for owed-item `record_kind`, store layout, and store API fidelity. m-7 may guide the engine implementation, but it must not redefine the m-1 store contract.
- Keep the owed-item projection claim bounded: it guards recorded owed items only; it does not make unrecorded observations impossible to miss.
- If S2's pair plan touches full FieldSpec registry, the MCP live adapter, observe, routing execution, or consumer schemas, that is OUT and must escalate before delegated dispatch.

## Verification

- Read exact dispatch: `master/relays/s2-dispatch/PLAN-orchestrator-planner-20260703-223146.md`.
- Source relay lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s2-dispatch/PLAN-orchestrator-planner-20260703-223146.md` -> OK.
- Dispatch-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s2-dispatch` -> FAIL with F-S2-1 above.
- S1 close source lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py frank/.relays/s1/s1-merge-gate/RECONCILE-orchestrator-planner-20260703-220652.md` -> OK.
- `git -C frank rev-parse HEAD` -> `674c844c1cf555a7c3490b4f76042e88822f4428`.
- `git -C frank rev-parse s1-close` -> `8b522847998a23c71f842378d3d48646411c48a4`.
- `git -C frank status --short --branch` -> `## main`.
- Filed relay lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s2-dispatch/RECONCILE-orchestrator-reviewer-20260703-223502.md` -> OK.
- Dispatch-root lint after filing remains FAIL on the source PLAN with F-S2-1; this reviewer relay is not the failing file.
- INDEX row check: `tail -n 5 master/relays/INDEX.md` shows the `20260703-223502` revise row at EOF.
- Harness root `git status --short` -> unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no code edits and no `frank/` edits.

FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository: `fatal: not a git repository (or any of the parent directories): .git`; `frank/` status is `## main`.
