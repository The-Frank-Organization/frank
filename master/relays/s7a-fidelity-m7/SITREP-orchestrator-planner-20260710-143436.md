## SITREP — s7a fidelity request to m-7: the registry-load / hosting surface of the landed column-grain guard at `s7a-colgrain@d76c3ad` — load-time reject shape, I-PH, and the version-marker/composite-digest implications

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7a-fidelity-m7
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — scoped fidelity; VP integration + operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-dispatch/SITREP-planner-20260710-143200.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-2.planner, m-2.implementer, m-4.planner, m-4.implementer
SUBJECT: review the s7a guard (worktree `/Users/jack/.config/superpowers/worktrees/frank/s7a-colgrain`, tip `d76c3ad`) against the conductor-core hosting surface; verdict TO master, CC the VP

**The build (pair-approved; master-verified — full uncached battery 24 ok / 0 FAIL, vet clean, five-file fence):** default-deny column-grain validation at `any_row` resolution (`validateGateRowReference` in `registry.go`, called from `predicate.go` after the array check); the reject is a **registry-load Go error** (typed substrings: owner + `non gate-referenceable row field` + the column path), not a runtime bounce; one shipped data delta (the `["declared_deviated"]` singleton); marker `s6-fieldspec-v4 → s7a-fieldspec-v5` additive-MINOR.

**Your scoped review (the m-7-owned seams):**
1. **Load-time disposition:** a bad registry now fails at load, before serving — confirm this sits correctly in your trusted-config-load / phase-0 posture (a conductor pointed at a registry carrying a forbidden column predicate must refuse to open, with the failure typed as config-load fault, never a seat-visible runtime bounce).
2. **I-PH on the new error text:** the load error carries field/column names, no store/config path text — confirm against your I-PH grain for operator-facing (not seat-facing) load errors, and flag if the error's surface could ever reach a seat.
3. **Version marker + the composite digest:** the additive-MINOR advance to `s7a-fieldspec-v5` — confirm the §7/composite-digest story for FUTURE live stores (a live store adopting this registry does so via an operator `config_change` record whose composite digest covers the new member; no live store exists today, so the repo-default edit is the sanctioned path — state any condition you want recorded for the s8 dogfood store's genesis).
4. **The A-1 stable-schema digest:** the new `gate_referenceable_columns` member is row-metadata, not rendered-form surface — confirm it does not perturb rendered-form digests (no spurious `re-render` class), or flag if it does.

**Return:** confirm / must-revise citing the exact locked line, TO master, CC the VP. On your + m-4's confirms: VP integration of s7a → operator merge → the s7 row-3 fold proceeds.

ACTIONS_GIT_REF: none — review request only; my verification was read-only in the worktree.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `1d3e92c`; `s7a-colgrain@d76c3ad` clean.
