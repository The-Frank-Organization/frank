## COORD — spec-surface supplement to the s1-core AUDIT dispatch (absolute paths + the guide-gate checklist)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-core-audit
PARENT_DISPATCH_ID: s1-dispatch
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1.orchestrator-planner
TO: s1-core.planner, s1-core.implementer
CC: s1.orchestrator-reviewer, operator
IN_REPLY_TO: s1-core-audit/AUDIT-orchestrator-planner-20260703-133046.md
SUBJECT: read WITH your AUDIT dispatch — exact locked-spec paths/sections from the m-7 guide's context-brief; the pre-published guide-gate checklist your later PLAN is reviewed against

Supplement only — changes no scope, grants no authority; the AUDIT dispatch
(`s1-core-audit/AUDIT-orchestrator-planner-20260703-133046.md`) remains the authorizing relay
and wins on any disagreement. This replaces its relative/glob spec pointers with the exact
surface the m-7 guide published after that dispatch was written
(`.relays/s1/s1-dispatch/COORD-planner-20260703-134029.md`).

Audit-target paths (read-only, ABSOLUTE; never edit — escalate spec problems to s1.orchestrator-planner):
1. Charter dispatch: .relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md
2. Build strategy + hardened exit gate: master-docs/master/STEP-1-KICKOFF.md
3. Engine spec: master-docs/master/ARCHITECTURE.md — §C4.1, §C4.2 (18-row seam matrix), §C4.3 (claim boundary + I-PH, :450-463), owed fixture ledger (:477-482)
4. m-7 design-of-record (DESIGN-LOCKED): the m-7 conductor-core design-of-record (2026-07-01) — S1-relevant: §2, §3, §4 (Package-A rename pivot; presence=committed), §6 (fault→held), §8 (guardrail + wake), §12, §13 (F1–F11), §16 (claim sweep)
5. FROZEN m-1 contract: the m-1 trust/identity design-of-record (2026-06-28)
6. FROZEN m-2 contract: the m-2 forms/determinism design-of-record (2026-06-28)
7. Replay corpus (upstream baseline): <the archived upstream release>/
8. Guide-gate checklist (your later PLAN is reviewed against it; summarized in the sprint ROADMAP): .relays/s1/s1-dispatch/SITREP-planner-20260703-133102.md

Audit-relevant additions the checklist makes concrete (fold into your spec-to-exit-gate map):
- The **rename-pivot shape** is gate-item #5: canonical-record `rename()` = the single commit pivot, fsync-before-rename, presence=committed, projections derived, outcome records reference `intake_id`. Your audit should extract the exact §4/§13 requirements (incl. F9 no-stale-re-emission, F11 one-pivot-per-mutation).
- **`bounced` is not a value token** — the terminal enum is byte-exact `{accepted, rejected, held}`.
- The **D5 residual** must be stated beside every exclusivity-shaped claim in anything seat- or user-facing (§16 sweep).

The sprint ROADMAP (docs/sprints/2026-07-03-s1-slice-1/ROADMAP.md) now carries this full spec
table + the checklist; it was updated after your dispatch was written.

ACTIONS_GIT_REF: wrote this supplement relay + an INDEX.md row under .relays/s1/ (gitignored substrate); ROADMAP.md spec-table/checklist update committed separately on main (see git log)
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: read this with your AUDIT dispatch; proceed with the audit against the absolute paths above.
