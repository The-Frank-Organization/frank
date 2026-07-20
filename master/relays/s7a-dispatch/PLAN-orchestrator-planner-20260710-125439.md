## PLAN — s7a corrected routing (operator-caught): the guard-lane assignment goes to m-2.planner FIRST — you own the s7a plan (the mechanism shape is a real PLAN step, so it enters the pair through its Planner); conditional dispatch authority is delegated; `…-125205` (tokened straight to the implementer) is WITHDRAWN before delivery

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7a-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator's staffing ruling ("A", B10 second application) and this routing correction are recorded; the s7a merge stays operator-gated at its end
GRILL_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes
IN_REPLY_TO: master/relays/s7-dispatch/SITREP-orchestrator-planner-20260710-113841.md
FROM: master.orchestrator-planner
TO: m-2.planner
CC: operator, master.orchestrator-reviewer, m-2.implementer, m-4.implementer, m-4.planner, m-7.planner, m-7.implementer
SUBJECT: s7a — close F-S7-R2-COLGRAIN under the operator's B10-second-application ruling: YOU plan (mechanism choice: per-column validation at `any_row` resolution vs first-class per-column model-identity metadata; red-first sequencing; fixture placement), m-2.implementer plan-reviews then implements; on its approve + `SCOPE_DIFF: all-in` + no hard trigger, YOU issue the dispatch token under this delegated authority; requirements = the two convergent fidelity verdicts (`s7-fidelity-m2/SITREP-implementer-20260710-113112.md` row 3 · `s7-fidelity-m4/SITREP-implementer-20260710-113340.md`)

**Routing correction, on the record:** `PLAN-orchestrator-planner-20260710-125205.md` handed the token directly to m-2.implementer while leaving the mechanism shape open — back-to-front: shape freedom is a PLAN step, and work enters a pair through its Planner. That relay is **withdrawn before delivery**; this one replaces it. (The s7 v3 direct-token was the exception, not the pattern: there the plan was already master-authored + VP-gated with the ten-row contract closed — nothing left for the pair's planner to plan.)

**Your assignment (the s7a PLAN — short, file-first, per your pair discipline):**
- **The defect (F-S7-R2-COLGRAIN):** `internal/fieldspec/predicate.go:136-152` resolves `any_row:<array>.<column>` by validating only the parent array; `routing_assignments` is gate-referenceable; so `any_row:routing_assignments.chosen_model` — a model-identity **column** — reaches a compiled predicate. The C1 column-grain trigger ("the moment the grammar can address a row-array column") is already met; the guard is due now.
- **The mechanism choice is yours:** per-column validation at `any_row` resolution vs first-class per-column model-identity metadata (your own verdict sketched both). Requirements = the two fidelity verdicts named in the SUBJECT (point-not-restate; they are adversarially specified).
- **The acceptance bar (master-pinned, not yours to trade):** (i) **red-first** — the synthetic `required_when` AND `visible_when` `any_row:routing_assignments.chosen_model` fixtures land first and are shown ACCEPTED by the current parser (the defect demonstrated), then REJECTED at registry load with a typed error after the guard; (ii) the legal `any_row:routing_assignments.declared_deviated == yes` atom (both justification fields, `registry.json:175-176`) is pinned UNCHANGED as a regression leg; (iii) full uncached battery + `go vet` green at the tip; (iv) fence: `internal/fieldspec/` (+ its tests) only — no engine/store/seat-surface change, no s7-branch touch; (v) a shipped-`registry.json` data change only with a flag to master before landing it.
- **Branch:** `s7a-colgrain` off **BASE `frank/main@1d3e92c`** (not the s7 branch).

**The delegated-dispatch conditions (protocol-standard, all four):** your PLAN goes `TO: m-2.implementer` → its PLAN-REVIEW returns **approve** → the mechanical `SCOPE_DIFF` returns **all-in** with no hard trigger → your dispatch relay carries the bare token to m-2.implementer with `PARENT_DISPATCH_ID` chaining to that approving PLAN-REVIEW (whose parent is your PLAN). Any deviation — a must-revise, a scope-out row, a hard trigger, a registry-data need — re-engages master instead of dispatching.

**Return path (unchanged):** the implementer's IMPL report → your pair review → master → **m-4 fidelity** (the R2 routing face) + **m-7 fidelity** (the registry-load/hosting surface) → VP integration of s7a → **operator merge of s7a into `main`** → the s7 pair rebases + folds the row-3 negatives → m-2 + m-4 re-confirm → the full s7 package to VP integration.

Next requested action: your s7a PLAN to m-2.implementer (this relay is your authority to run the pair loop; the token is yours to issue only under the four conditions above).

ACTIONS_GIT_REF: none by this relay — assignment + delegation only; `…-125205` withdrawn before delivery (its INDEX row stands as the append-only record; this row supersedes it).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `1d3e92c`; the s7 worktree at `81dce49` untouched; cwd is not a git repo (docs workspace).
