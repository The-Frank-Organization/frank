## RECONCILE -- VP spot-check of s1 dispatch

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s1-dispatch
PARENT_DISPATCH_ID: s1-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- fix relay-location/charter handoff before operator relays to s1
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, s1.orchestrator-planner, m-7.planner, m-1.implementer, m-2.implementer, m-7.implementer
SUBJECT: s1-dispatch review -- revise relay substrate before fresh s1 handoff

VERDICT: revise

This review is an operator-requested VP spot-check. The incoming relay addresses `master.orchestrator-reviewer` in `CC`, not `TO`, and says the VP is copied for visibility rather than re-gate. I therefore do not claim acting authority over `s1.orchestrator-planner` or `m-7.planner`; this relay reports the handoff issue back to the master planner before the operator relays the dispatch onward.

### Finding

1. **The dispatch changes the relay substrate without updating the durable charter or providing a fresh `frank/` charter.**

The approved `step1-plan` r2 still said slice-team relays live under `master-docs/master/relays/`, namespaced by `s<N>` (`step1-plan/PLAN-orchestrator-planner-20260703-125536.md:31-34`), and the VP approval scope approved that convention (`step1-plan/RECONCILE-orchestrator-reviewer-20260703-125826.md:31-37`). The new dispatch changes the convention to slice-team relays in `frank/`, governed by `sprint-doc-setup` (`s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md:22-25,77-85`), and `STEP-1-KICKOFF.md` / `master-docs/master/README.md` have been updated to match.

That may be the right operator product call, but the durable charter loaded by sessions in this cwd still says governance docs and relays live in `master/`, Step-1 code lives in `frank/`, and every substantive handoff is a lint-clean relay under `master-docs/master/relays/` (`CLAUDE.md:19-21`). A fresh `frank/` session currently has no `AGENTS.md`, `CLAUDE.md`, or README carrying the s1 seat charter or the `master/` spec pointers. Evidence: `find frank -maxdepth 2 \( -name AGENTS.md -o -name CLAUDE.md -o -name README.md -o -name SKILL.md \) -print` returned no files; `frank/` contains only `.git`.

Without a durable charter/pointer in `frank/`, the s1 session can start in an empty repo, run `sprint-doc-setup`, and create a private `.relays/<RUN_ID>` substrate without inheriting the master seat rules, locked spec pointers, CC/fidelity obligations, or the master INDEX visibility expectations. That is a handoff hazard, not a mechanism disagreement.

### Required amendment

Before the operator relays this to a fresh `s1.orchestrator-planner`, do one of:

1. Keep the approved r2 convention: s1 operational relays remain under `master-docs/master/relays/` with `s1-*` dispatch IDs, while code lands in `frank/`.
2. Or make the `frank/` convention durable: update the standing charter or add a `frank/AGENTS.md` / `frank/CLAUDE.md` bootstrap pointer that names the s1 seat, points to the master design-of-record, states that build relays live under `.relays/<RUN_ID>` via `sprint-doc-setup`, and preserves the required master visibility points: master dispatch in `master-docs/master/relays/INDEX.md`, s1 exit SITREP back to master, m-7+VP plan gate before any implementation dispatch, and m-1/m-2 fidelity review before implementation authority.

If the operator deliberately wants `frank/.relays` to be the slice substrate, I do not object to that product structure. The blocker is that the durable bootstrap surfaces must agree before a fresh team session relies on them.

### Non-blocking concurrence

- The incoming relay is lint-clean and contains no live `DISPATCH IMPL`.
- The S1 technical scope matches the approved r2 plan: thin end-to-end relay, MVP FieldSpec, minimal validation/lineage, crash-atomic append, S1-minimal dissolved-linter replay, I-PH, liveness, and park/wake.
- The dispatch preserves the S1/bootstrap guide+VP plan gate and the m-1/m-2 consuming-surface fidelity edge.
- `frank/` remains empty on unborn `main`; no code exists to review yet.

## Verification

- Read exact relay: `.relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md`.
- Planner relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md` -> OK.
- Dispatch-root lint before filing: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s1/s1-dispatch` -> OK.
- Checked prior VP-approved r2 plan: `master-docs/master/relays/step1-plan/PLAN-orchestrator-planner-20260703-125536.md` and `master-docs/master/relays/step1-plan/RECONCILE-orchestrator-reviewer-20260703-125826.md`.
- Checked live charter and kickoff: `CLAUDE.md:19-21`, `master-docs/master/STEP-1-KICKOFF.md:13-19`, `master-docs/master/README.md:9,137`.
- Checked `sprint-doc-setup` relay substrate rules: `~/.codex/skills/sprint-doc-setup/SKILL.md`.
- Checked `frank/` bootstrap state: no `AGENTS.md`, `CLAUDE.md`, `README.md`, or `SKILL.md`; `frank/` currently contains only `.git`.
- New relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-dispatch/RECONCILE-orchestrator-reviewer-20260703-131452.md` -> OK.
- Dispatch-root lint after filing: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s1/s1-dispatch` -> OK.
- `master-docs/master/relays/INDEX.md` tail contains the appended `20260703-131452` row.
- `git -C frank status --short --branch` -> `## No commits yet on main`; `git -C frank symbolic-ref --short HEAD` -> `main`; `git -C frank rev-list --count --all` -> `0`.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master-docs/master/relays/INDEX.md`; no design-doc/source/code edits and no `frank/` edits.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository: `fatal: not a git repository (or any of the parent directories): .git`; `frank/` status remains `## No commits yet on main`.
