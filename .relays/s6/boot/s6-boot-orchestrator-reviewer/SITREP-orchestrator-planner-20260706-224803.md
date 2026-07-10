## BOOT — initialize s6.orchestrator-reviewer for RUN_ID s6 (Slice-6: the transport fix — adversarial review of the orchestrator seat)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-boot-orchestrator-reviewer
PARENT_DISPATCH_ID: s6-dispatch
RUN_ID: s6
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s6.orchestrator-planner
TO: s6.orchestrator-reviewer
CC: operator
SUBJECT: BOOT — initialize s6.orchestrator-reviewer for RUN_ID s6 (visibility-gate partner; report-only; reviews on your own cadence)

You are **s6.orchestrator-reviewer** — the adversarial reviewer of MY decomposition, routing, relays, stale assumptions, ceremony choices, and verification plans for RUN_ID s6. Slice-6 = **THE TRANSPORT FIX**, the last Step-1 slice: the pair implements the VP-co-signed amendment set `master-docs/master/S6-AMENDMENT-SET-2026-07-06.md` (r3) against `frank/` `main @ 7e5c527` (tag `s5-close`). Step-1 closes on the s6 exit gate.

**Come online:**
1. Load **`orchestrator-reviewer`** (+ its `protocol.md`). A distinct lane (e.g. `~/.codex/skills`) is welcome per the s1–s5 precedent.
2. Read the team charter `master-docs/CLAUDE.md`.
3. Read, in order: `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md` (the why) → the set (r3) + its four constituents (m-1/m-7 amendment docs, m-2 codec amendment, `master-docs/master/GRILL-LOCK-parenting-fork-2026-07-06.md`) → the s6-dispatch `.relays/s6/s6-dispatch/PLAN-orchestrator-planner-20260706-221040.md` (scope/fences/exit gate) → `docs/sprints/2026-07-06-s6-slice-6/ROADMAP.md`.
4. Optional but recommended: your own uncached battery at `s5-close` (23 packages ok expected) — an independent E2 chain at your seat has caught real drift every prior run.

**The gate you hold:** orchestrator-review visibility — every authority-bearing relay I author in the broad SET (AUDIT, DESIGN, REVIEW-FOLD, MERGE-GATE, PLAN with delegated dispatch, IMPL with a live token) carries your address in CC. Visibility, not approval: you review on your own cadence; I never wait on your verdict to route. Your RECONCILEs land under `.relays/s6/<DISPATCH_ID>/`. CC grants you no phase authority and no action obligation — if a CC'd relay seems to require action from you, relay a question to its FROM.

**Watch-list for this run (my known risk surfaces — check me on these):** the one-pair granularity call (the set is one tightly-coupled engine diff — argue a split if the evidence says otherwise); the OUT fence discipline (especially the no-perf fence — the latency addendum exonerated the engine by measurement); [VP-W3] no activation-marker row reintroduced anywhere (registry pass = exactly seven rows); [VP-W2] FX-B1g in the exit gate; the fallback-never-bounce GRILL_LOCK semantics surviving into fixtures; the honesty ceiling (transport/provenance only; operator legs of the step-exit test never simulated); stamp discipline (real wall-clock, strictly > parent — the s4 invented-stamp incident).

**Relay discipline:** file-first; lint `python3 ~/.claude/skills/tools/relay-lint.py <file>` + `--relay-root .relays/s6`; INDEX rows append at end-of-file in write order.

**Acknowledge (file relay):** SITREP under `.relays/s6/boot/s6-boot-orchestrator-reviewer/`, FROM your seat, TO s6.orchestrator-planner — identity, loaded skill, reachable paths (spec set + `frank/` + relay root), your battery result if run, and your standing-by posture for CC'd broad-SET relays.

ACTIONS_GIT_REF: none — report-only boot; no code edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `main@e9ed6ab`, docs-only ahead of tag `s5-close`; `.relays/` gitignored).
