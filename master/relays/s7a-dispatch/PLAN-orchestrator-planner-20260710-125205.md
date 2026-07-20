## PLAN — the s7a guard-lane dispatch to m-2.implementer (operator ruling "A", the second B10 application): close F-S7-R2-COLGRAIN — the `any_row` column-grain model-identity guard in your own validator, red-first, off BASE `frank/main@1d3e92c`; the dispatch token stands bare below

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7a-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator's staffing ruling ("A", 2026-07-10) is recorded in `PROTOCOL-DEVIATIONS.md` B10 (second application); the s7a merge stays operator-gated at its end
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/SITREP-orchestrator-planner-20260710-113841.md
FROM: master.orchestrator-planner
TO: m-2.implementer
CC: operator, master.orchestrator-reviewer, m-2.planner, m-4.implementer, m-4.planner, m-7.planner, m-7.implementer
SUBJECT: s7a — the bounded production fix for F-S7-R2-COLGRAIN (a model-identity COLUMN predicate compiles: `internal/fieldspec/predicate.go:136-152` validates only the parent array; the C1 column-grain trigger is already met) — you write, m-2.planner pair-reviews; requirements = the two convergent fidelity verdicts (`s7-fidelity-m2/SITREP-implementer-20260710-113112.md` row 3 · `s7-fidelity-m4/SITREP-implementer-20260710-113340.md`); the token stands bare below

DISPATCH IMPL

m-2.implementer — under the operator's B10-second-application ruling you are the **sole s7a code writer**; **m-2.planner is your pair Planner** (shape concurrence before code + adversarial implementation review before the report reaches master). The requirements are already adversarially specified by the finding record — the two fidelity verdicts named in the SUBJECT are your requirements documents (point-not-restate); no separate master plan round.

**Work order:**
1. Branch **`s7a-colgrain`** off **BASE `frank/main@1d3e92c`** (NOT the s7 branch — s7a is its own micro-lane; the s7 test fold happens after your guard merges).
2. **Shape concurrence first (inside the pair):** a short PLAN note from you, m-2.planner concurs — the mechanism choice is yours (per-column validation at `any_row` resolution vs first-class per-column model-identity metadata), the acceptance below is not.
3. **Red-first:** land the failing fixtures BEFORE the fix — synthetic `required_when` AND `visible_when` predicates `any_row:routing_assignments.chosen_model` must be shown accepted by the current parser (the defect demonstrated red), then rejected at registry load after the guard (green).
4. **The guard:** `any_row:<array>.<column>` resolution validates the **column**, not just the parent array — a model-identity column (and any column that is not legitimately gate-referenceable at column grain) is rejected at load with a typed error. This is the C1 carry's guard at the grain the live grammar already reaches.
5. **Regression legs:** the legal atom `any_row:routing_assignments.declared_deviated == yes` (both justification fields, `registry.json:175-176`) must be UNCHANGED — the declared-policy path stays green; the full uncached battery + `go vet` green at your tip; the shipped `registry.json` needs no data change if your mechanism choice avoids it (if it does need one, flag to master before landing it — repo-default registry edits are in-scope, but say so; there is no live store right now).
6. **Fence:** `internal/fieldspec/` (+ its tests) only. No engine/store/seat-surface change; no s7-branch touch; anything wider escalates to master.

**Return (your IMPL report, TO master, CC the VP + operator + m-2.planner + m-4 + m-7):** branch + tip SHA · the red-then-green fixture transcript · the regression-leg evidence · full uncached battery + vet output · the diffstat. Then: **m-4 fidelity** (the R2 routing face of the guard) + **m-7 fidelity** (registry-load/hosting surface) → VP integration review of s7a → **operator merge of s7a into `main`** → the s7 pair rebases and folds the row-3 `any_row` negatives (test-only, now green against your mechanism) → m-2 + m-4 re-confirm row 3 → the full s7 package to VP integration.

Blocked or ambiguous → relay to master; do not widen scope to resolve ambiguity.

ACTIONS_GIT_REF: none by this relay — authority grant only; the branch/commits are yours from here.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `1d3e92c`; the s7 worktree at `81dce49` untouched by this lane; cwd is not a git repo (docs workspace).
