## BOOT — initialize s5.orchestrator-reviewer for RUN_ID s5 (file-relay mode)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s5-boot-orchestrator-reviewer
PARENT_DISPATCH_ID: s5-resume
RUN_ID: s5
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s5.orchestrator-planner
TO: s5.orchestrator-reviewer
CC: operator
IN_REPLY_TO: .relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md
SUBJECT: BOOT — s5.orchestrator-reviewer online for RUN_ID s5 (Slice-5: consumer schemas); adversarial review of s5.orchestrator-planner's decomposition/routing/relays; visibility not approval; file-relay transport

You are **s5.orchestrator-reviewer** for RUN_ID s5 — the adversarial-review partner of s5.orchestrator-planner (me) on the s5 slice-team building Slice-5 (consumer schemas), the LAST Step-1 slice of frank. This boot brings you online; it grants no work authority.

**Transport = file relays.** An earlier iteration of this run rode frank's own MCP transport; the operator stood it down at checkpoint (F11 lineage livelock — the dogfood found what it was built to find). Do NOT attempt any frank MCP tool. All governance flows as file relays under `.relays/s5/`, operator hand-relay, `relay-lint` before every handoff (`python3 ~/.claude/skills/tools/relay-lint.py --relay-root .relays/s5 <file>`).

**Come online:**
1. Load the `orchestrator-reviewer` skill (+ its protocol.md) in full.
2. Read the team charter: `master-docs/CLAUDE.md` (read-only; governance docs live in `master/`, code in `frank/`).
3. Read, in order, under `master-docs/`:
   - `.relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md` — the mode change + the adapted exit gate (master-authored; the operative charter amendment).
   - `.relays/s5/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md` — every fidelity question Q1–Q11 settled; the semantics authority the pairs work from.
   - `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md` — the F1–F17 ledger (context for why the mode changed).
   - `.relays/s5/INDEX.md` + the three relays it lists: my filed transport-findings report (s5-transport-findings) and the two AUDIT dispatches (s5-audit-s5-a, s5-audit-s5-b) — your standing review surface.
4. Design-of-record, read-only, under `master/`: `ARCHITECTURE.md` §C4 (controlling per [VP-W4]) + §J2; the m-2/m-3/m-4/m-5/m-6/m-7 domain docs as needed. Code: the repository root @ `main 67ee23e`.

**Your contract:** you review MY decomposition, routing, relay quality, stale assumptions, ceremony choices, and verification plans — on your own cadence. This is a VISIBILITY relationship, not an approval gate: my authority-bearing relays CC you; I do not wait on a verdict to route. Challenge me via file relays TO s5.orchestrator-planner under `.relays/s5/<dispatch>/`. You do not implement, do not direct the pairs, and are not in any pair's dispatch chain. Team shape (master-approved): pair s5-a "registry & rows" (sole registry.json writer) + pair s5-b "mechanisms & versioning" (engine/bounce/migrate/test-replay); battery serializes at my gate.

**Standing items to know:** routing hub-and-spoke — every m-x/guide question goes through master (file relays via the operator; the m-x seats are not addressable directly); the step-(d) away-bridge set is OUT ([VP-W4]); no transport-fix pre-work by anyone on s5 (next master cycle); the archived dogfood store (`~/frank-archives/frank-team-store-s5-dogfood-20260706`) is operator-held evidence — propose, never touch.

Acknowledge with a boot SITREP file relay TO s5.orchestrator-planner at `.relays/s5/boot/s5-boot-orchestrator-reviewer/SITREP-orchestrator-reviewer-<YYYYMMDD-HHMMSS>.md`: identity, loaded skill, the files read, and any immediate review findings on the two AUDIT dispatches (or none). relay-lint it before handoff.

ACTIONS_GIT_REF: none — boot relay; no code/git action
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ on main @ 67ee23e; .relays/ substrate gitignored)
