## RECONCILE -- VP re-review of r2 s6 transport-fix amendment set

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- VP co-sign review only; operator and build dispatch remain separate
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-7.planner, m-4.planner
IN_REPLY_TO: s6-design/RECONCILE-orchestrator-planner-20260706-214751.md
SUBJECT: s6 r2 amendment set re-review -- activation model fixed, but r2 package still needs stale-marker/r5 fixture cleanup before co-sign

VERDICT: must-revise

## Blocking Finding

1. The activation-marker conflict is resolved in substance, but the r2 package is not yet byte-faithful to that resolution. The integrated set now correctly removes the activation-marker row and uses m-1's derived-only activation rule, and the m-7 pair confirms the marker withdrawal on the merits. However, three stale package surfaces still point at the old marker/r4 shape:
   - The integrated set's artifact table still names the m-7 amendment as **r4**, while the resolved no-marker source is **r5**.
   - The m-7 B-1.2a allowlist text still says the build plan carries the allowlist through "m-2's marker row," even though r5 withdraws that row.
   - The integrated set's fixture roll-up still says m-7's fixture set ends at **FX-B1f (17)**, but r5 adds **FX-B1g** for the re-mint/new-generation activation leg. That generation fixture is the executable proof of the derived-only model's mint-boundary behavior and cannot be dropped from the build-slice roll-up.

This is a narrow integration-cleanliness blocker, not a rejection of the derived-only activation model. The design lock should not freeze a set that simultaneously says "no marker row" and "carry through m-2's marker row," nor should it dispatch the build slice while omitting the r5 generation fixture.

## Required Revision Shape

- Update `master/S6-AMENDMENT-SET-2026-07-06.md` to cite the m-7 amendment as r5/current, not r4.
- Update the integrated fixture roll-up to include FX-B1g and the correct fixture count/scope.
- Update `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md` B-1.2a so the allowlist is carried through the B-2 boot form / m-7 admission check, not through an m-2 marker row.
- Tidy the m-7 status line so it no longer says r5 is awaiting implementer co-sign after `DESIGN-REVIEW-implementer-20260706-213621` approved it. This one is status text, but it should be fixed in the same touch.

## Non-Blocking Confirmations

1. The original VP blocker is discharged on the merits. The r2 set removes the persisted marker row; m-1's derived-only activation rule owns semantics; m-7's accept-time classification is transient; m-2's no-marker B-2 shape is used.

2. The m-7 pair's r5 confirm is adequate. The implementer withdrew the prior marker integration condition, confirmed the exact-form admission fix still stands, verified that v3.0 GC does not collect accepted canonical records, and preserved the future-retention route-back to m-1.

3. F14, the grilled parenting fork, m-4's `routing_ref_honored` condition, and the F1-F17/B-1..B-3 total disposition bar remain otherwise co-signable.

## Verification

- Incoming re-request lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/RECONCILE-orchestrator-planner-20260706-214751.md` -> OK.
- Dispatch-root lint before filing this relay: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s6-design` -> OK.
- r2 set check: `master/S6-AMENDMENT-SET-2026-07-06.md:45-60` removes the marker row and states the derived-only model, but still omits FX-B1g in the fixture roll-up and cites the m-7 artifact as r4 at `:17-23`.
- m-7 r5 check: `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:96-105` adopts derived-only activation and transient classification; `:127-129` adds FX-B1g.
- stale marker-row text: `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:104` still says the build plan carries the allowlist through m-2's marker row.
- m-1 boundary check: `master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:99-113` remains the governing derived-only boundary.
- GC evidence check: `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:134-137` and `frank/internal/gc/gc.go:46-102` support the no-GC-failure finding for accepted canonical records.
- Filed relay lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/RECONCILE-orchestrator-reviewer-20260706-215052.md` -> OK.
- Filed dispatch-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s6-design` -> OK.
- INDEX row check: `tail -n 6 master/relays/INDEX.md` shows the `20260706-215052` must-revise row at EOF.
- `git -C frank status --short --branch` after filing this relay -> `## main`.
- Harness root `git status --short` after filing -> unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no code edits and no `frank/` edits.

FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository: `fatal: not a git repository (or any of the parent directories): .git`; `frank/` status is `## main`.
