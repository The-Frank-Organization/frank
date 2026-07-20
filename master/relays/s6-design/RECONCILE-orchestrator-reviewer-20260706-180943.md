## RECONCILE -- VP pre-handoff review of s6 transport-fix design dispatch

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- VP pre-handoff review only; operator handoff remains separate
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator
IN_REPLY_TO: s6-design/PLAN-REVIEW-orchestrator-planner-20260706-180315.md
SUBJECT: s6 transport-fix design pre-handoff VP review -- must revise F14 ownership before pair handoff

VERDICT: must-revise

## Blocking Finding

1. F14 is mis-homed to m-7 as sole owner, contrary to the transport ledger and standing domain charter. The s6 dispatch assigns "the store lockfile (F14)" to m-7 under engine liveness and ops. But the spec seed's owner split says m-1 owns "store/lineage: anchor model, lock, retraction," and the standing charter says m-1 owns the store isolation / governed-writer contract while m-7 hosts and sequences that contract. F14 is not just local process hygiene; it guards the single-writer/store-isolation invariant. Before handoff, revise the ownership split so m-1 owns the F14 store-lock invariant and lock semantics, with m-7 owning any runtime enforcement, process choreography, startup/bounce behavior, and fixture execution needed to make that invariant live. If master wants joint handling, name it explicitly as an m-1/m-7 seam rather than routing it only to m-7.

## Non-Blocking Review Notes

1. The in-step translation is faithful. The dispatch keeps s6 inside Step-1 because the operator-as-transport goal is not honestly closed while the live conductor can livelock, and it moves the step-exit test onto the fixed conductor with the §7 registry-application leg plus the F11 dogfood replay regression.

2. The m-1 parenting fork gate geometry is correct. The dispatch treats conductor-computed PARENT as a candidate, not a pre-decided answer, and requires decision packet -> master -> operator GRILL -> GRILL_LOCK before the m-1 fork locks.

3. The total-disposition exit bar is the right design-phase bar. Every F1-F17 item must carry amendment / in-slice fix / wontfix-with-rationale, with F8 preserved as the numbering gap, before master integrates and asks for VP co-sign.

4. The design-amendment cadence is right for this package. These fixes touch locked pair contracts, so the owning pairs should amend and pair-review their designs before a fresh build slice is dispatched. Do not collapse this into one design+build slice team.

## Required Revision Shape

- Move F14's durable owner from m-7 to m-1, or state a joint seam where m-1 owns the invariant and m-7 owns the runtime enforcement.
- Keep the design-phase exit requirement that every F item is dispositioned; F14's disposition should cite both the lock invariant and the runtime fixture obligation.
- Preserve the current no-code/no-frank-edits/no-build-slice-work constraint.

## Verification

- Source review-request lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/PLAN-REVIEW-orchestrator-planner-20260706-180315.md` -> OK.
- Held dispatch lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/DESIGN-orchestrator-planner-20260706-180315.md` -> OK.
- Held dispatch-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s6-design` -> OK.
- Basis check: `master/TRANSPORT-FINDINGS-2026-07-06.md:70-73` assigns m-1 the store/lineage lock, and `CLAUDE.md:45-54` assigns m-1 store isolation while m-7 hosts/sequences policy-domain contracts.
- Step-1 ruling check: `ROADMAP.md:83-90`, `master/README.md`, and `master/RECONCILE.md` record the fixed-conductor exit test and the 2026-07-06 operator in-step ruling.
- Baseline check: `git -C frank status --short --branch` -> `## main`; `git -C frank rev-parse --short HEAD` -> `7e5c527`; `git -C frank rev-parse --short s5-close^{}` -> `7e5c527`.
- Filed relay lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/RECONCILE-orchestrator-reviewer-20260706-180943.md` -> OK.
- Filed dispatch-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s6-design` -> OK.
- INDEX row check: `tail -n 5 master/relays/INDEX.md` shows the `20260706-180943` must-revise row at EOF.
- `git -C frank status --short --branch` after filing this relay -> `## main`.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no code edits and no `frank/` edits.

FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` status before filing this relay was `## main`.
