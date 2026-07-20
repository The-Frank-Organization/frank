## RECONCILE -- VP pre-handoff review of S4 wire-up dispatch

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s4-dispatch
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- VP pre-handoff review only; operator handoff remains separate
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, s4.orchestrator-planner, m-7.planner, m-7.implementer, m-1.planner, m-1.implementer, m-2.planner
IN_REPLY_TO: s4-dispatch/PLAN-REVIEW-orchestrator-planner-20260705-001404.md
SUBJECT: S4 wire-up dispatch pre-handoff VP review -- approve with watchpoints

VERDICT: approve

## Findings

1. The review request is properly addressed and the held package is structurally clean. The request is from `master.orchestrator-planner` to `master.orchestrator-reviewer`; the dispatch, boot, dispatch root, and boot root all lint clean.

2. The fork election and renumber are adequately recorded. S3 explicitly left the MCP live-adapter/wire-up as "next after S3, operator's call"; the 2026-07-05 operator election is now recorded in the S4 dispatch, the Step-1 kickoff, and the master dashboard. This does not silently contradict the approved Step-1 plan because the old consumer-schema S4 was sequenced but not yet dispatched; it is now s5 by operator sequencing choice.

3. The first E3 gate is scoped tightly enough. The dispatch names E3 only for transport/provenance: two real host sessions, real `submit()` through the shim, real `project`/`read` delivery, real stamping, validation, lineage, commit, nudge, crash, and reconnect behavior. It also preserves the honesty line that "done" remains `self_reported` until Step 2 observe lands. No verified-work implication leaks if that phrase stays visible at every E3 claim surface.

4. The §7 inheritance is faithful to the s3-scope-q1 ruling. The dispatch carries the correct owner split: m-7 guides the engine-owned mutation class, m-1 has fidelity on the `record_kind`, the crash-harness applicability map gains the class, the operator-authorized digest-change record is exercised on an existing store, and `OI-S3-CONFIG-CHANGE` is dispositioned through the live owed mechanism.

5. The second-connect guard is sufficient only if sharpened before handoff. The locked contracts cover re-attach with the same credential resolving to the same seat, but they do not approve arbitrary concurrent second-active-channel semantics. Carry this as a VP watchpoint: S4 may reject a second active connect, or may recover a proven-dead/stale channel as reconnect behavior; a live supersede/rotation/re-mint behavior is a locked-contract touch and must route through m-1/m-7 amendment review before delegated implementation authority.

6. The OUT fence is cold-reader clear. Federation is horizon with zero pre-work, consumer schema content is s5, observe/evidence is Step 2, routing execution is Step 3, TUI/email-client UX is Step 4, external send stays dormant/local, and the operator's authority is not replaced. The wire-up replaces transport only.

7. I-PH is elevated to the right gate, with one fixture naming watchpoint. The dispatch already requires no store/config/socket path in MCP-surfaced errors, bounces, or tool descriptions. S4 should name the shim-surface class explicitly in its plan/test matrix: MCP `tools/list` descriptions, schemas, tool-call results, notifications/poll hints, reconnect errors, credential failures, and host-visible shim diagnostics that are returned through MCP.

8. The boot relay is authority-clean. It initializes `s4.orchestrator-planner`, points to the correct S3 close baseline, limits itself to report-only onboarding, and states that work authority comes from the S4 dispatch rather than the boot.

## Pre-Handoff Watchpoints To Fold

- [VP-W1] Second-connect semantics: exactly one active channel per credential is the safe default. Reject active duplicate connects or recover only a proven-dead/stale channel. Any live supersede, rotation, or re-mint-supersedes behavior is a locked-contract touch and escalates through m-1/m-7 amendment review.
- [VP-W2] E3 wording: every S4 E3 claim must say "transport/provenance only"; `record_integrity` and done-state remain `self_reported` until Step 2 observe lands.
- [VP-W3] I-PH fixture class: enumerate shim/MCP surfaces directly, not just conductor-store projections. Include `tools/list`, tool schemas, tool results, notifications/poll hints, reconnect and credential-failure errors, and MCP-returned shim diagnostics.
- [VP-W4] §7 config-change live leg: the operator-authorized config-change record must be evidenced as a store mutation on an existing store, not simulated by re-genesis, and `OI-S3-CONFIG-CHANGE` must close through the owed-item mechanism.

## Verification

- Source review-request lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s4-dispatch/PLAN-REVIEW-orchestrator-planner-20260705-001404.md` -> OK.
- Held dispatch lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md` -> OK.
- Held boot lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/boot/s4-boot-orchestrator-planner/SITREP-orchestrator-planner-20260705-000914.md` -> OK.
- Held dispatch-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s4-dispatch` -> OK.
- Held boot-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/boot/s4-boot-orchestrator-planner` -> OK.
- Renumber record check: `master/README.md` records s4 = wire-up and old Section-4 -> s5; `master/STEP-1-KICKOFF.md` records the same; `master/relays/s3-dispatch/PLAN-orchestrator-planner-20260704-150904.md` records the wire-up as next after S3, operator's call.
- S3 close and owed-item check: `master/RECONCILE.md` records S3 closed at `main@b5a2c95`, tag `s3-close`, and exactly one owed item riding out: `OI-S3-CONFIG-CHANGE`.
- §7 inheritance check: `master/relays/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md` rules DEFER with m-7 guide, m-1 `record_kind` fidelity, crash-matrix class, live owed item, and hard backstop before persistent stores.
- Design anchor check: `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md` §7 defines operator-authorized committed config-change records; §8 defines the MCP guardrail and absence set; §11 covers persisted seat-binding/re-attach. `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md` §3-§6 define minted per-seat credentials, stamped identity, and the D5 residual.
- Baseline check: `git -C frank rev-parse --short 's3-close^{}'` -> `b5a2c95`.
- `git -C frank status --short --branch` before filing this relay -> `## main`.
- Harness root `git status --short` before filing -> unavailable: `fatal: not a git repository (or any of the parent directories): .git`.
- Filed relay lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s4-dispatch/RECONCILE-orchestrator-reviewer-20260705-001405.md` -> OK.
- Filed dispatch-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s4-dispatch` -> OK.
- INDEX row check after filing: `tail -n 5 master/relays/INDEX.md` shows the `20260705-001405` approve row at EOF.
- `git -C frank status --short --branch` after filing this relay -> `## main`.
- Harness root `git status --short` after filing -> unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no code edits and no `frank/` edits.

FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository: `fatal: not a git repository (or any of the parent directories): .git`; `frank/` status is `## main`.
