## SITREP — the two operator legs are ON RECORD (the lock-blockers for the approved m-2 grammar): (1) activation-authorization RATIFIED; (2) m-3's three check-policy defaults RATIFIED AS INTERIM with a MANDATORY s10 SUNSET — the silent auto-kill and the static gate are testing-phase only and MUST be replaced by operator-in-the-loop shapes at s10, as a named s10 exit condition

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step2-plan
PARENT_DISPATCH_ID: s6-exit-gate
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this relay RECORDS the operator's in-session rulings of 2026-07-11; nothing further is requested
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s8-design-m2-grammar/SITREP-planner-20260711-003600.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: operator, master.orchestrator-reviewer, m-2.planner, m-2.implementer, m-3.implementer, m-6.planner, m-7.planner
SUBJECT: operator rulings recorded — (1) activation of the observe layer = an operator-authored §7 `config_change`, restart-effective, both directions (RATIFIED as elaborated); (2) your grill's three defaults (E1 ≤5s / E2 ~120s hard-kill · the empty `side_effecting` allowlist · the static config gate) RATIFIED FOR THE TESTING PHASE with the operator's hard condition: **"the recommendations are ok for now for testing purposes but MUST BE REMOVED by s10"** — the s10-sunset shape below binds your design record and s10's exit

**Ruling 1 — activation authorization: RATIFIED (2026-07-11).** The observe layer turns on — and off — only by an operator-authored §7 `config_change` under the composite digest, restart-effective, no hot reload, no code default. This satisfies the first operator leg of m-2's approved grammar design (`s8-design-m2-grammar`, technical approve `…-003000`).

**Ruling 2 — the three check-policy defaults: RATIFIED AS INTERIM, with the s10 sunset (the operator's words carried verbatim in the SUBJECT).** The operative content:
1. **The 120s/5s hard-kill timeouts stand for s8/s9** — but the silent-kill disposition is testing-phase only. **At s10, the timeout path gains the operator-in-the-loop shape the operator described: a long-running check ALERTS the operator with the decision (kill / extend) instead of dying silently** — i.e., the park + ODB + operator-verdict machinery s10 builds, applied to check execution. The hard ceiling remains as the fail-safe backstop behind the prompt (the exact composition — J1's no-hard-deadline posture vs a wedged submit — is an s10 DESIGN question, named now, designed then).
2. **The empty `side_effecting` allowlist stands** ("sure") — the fail-closed floor is not sunset; what changes at s10 is only HOW additions/approvals arrive.
3. **The static config gate stands for s8/s9 and is REPLACED by the live operator prompt at s10** (as your grill already shaped it; the operator explicitly ratified the wait-for-s10 sequencing).
**The sunset is a NAMED s10 EXIT CONDITION, not a preference:** s10 does not close while the silent auto-kill disposition or the static-only gate survives. m-6 (CC'd) inherits this into the s10 comms-spine scope beside wake-on-reply; m-3 (TO) carries it in the check-policy design record so the interim defaults are marked interim in the artifact itself; master registers it in the kickoff.

**Effect on the board:** both operator legs of the m-2 grammar lock are satisfied; the lock now waits only on the master reconcile with `s8-design-m7-config` (in flight at m-7). m-3's own design (`s8-design-m3-registry`) should annotate the interim/sunset status on its §4 timeout/gate sections — a one-line fold, no re-review needed unless your implementer judges otherwise.

ACTIONS_GIT_REF: none — a ruling record; the kickoff registration rides the same fold (`STEP-2-KICKOFF.md` s10 bullet + design items).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; cwd is not a git repo (docs workspace).
