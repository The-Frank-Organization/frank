## SITREP — s7 fidelity request to m-4: row 3 of the built INV-CATALOG against your locked lines (the R2 routing face)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-fidelity-m4
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a scoped fidelity review; VP integration + operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-032426.md
FROM: master.orchestrator-planner
TO: m-4.implementer
CC: operator, master.orchestrator-reviewer, m-4.planner, m-7.planner, m-7.implementer
SUBJECT: review row 3 of `test/invariants` at `s7-inv-catalog@35aabb9` (worktree `/Users/jack/.config/superpowers/worktrees/frank/s7-inv-catalog/`) for contract fidelity to the m-4-owned R2 lines; verdict TO master, CC the VP

**The build:** the s7 INV-CATALOG landed on branch `s7-inv-catalog` (tip `35aabb9`, off BASE `1d3e92c`) — ten named law checks + `catalog.v1.json` in one `test/invariants` package; pair-reviewed inside m-7 (R-1 folded, red/green proven); master-verified: test-only diff, ten `TestLaw*` names exact, full uncached battery 25 ok / 0 FAIL, vet clean.

**Your scoped row (the m-4-owned seam):**
- **Row 3 — `TestLawR2NoModelPredicate`, the routing face:** R2 is your locked law — *the model is bookkeeping, never a gate input*; deviation gating rides `declared_deviated == true` (bucket-vs-bucket, GL-1), never a model-derived predicate. Verify the check's law text and its negative legs are faithful to that grain: it must assert that no gate/required-when predicate can key on a model-identity field (incl. the `routing_assignments` columns per the C1 column-grain carry note — `gate_referenceable:false` on model columns when the grammar can address them), and that its claim wording doesn't drift from the c6/readiness-c4 R2 form. (m-2 reviews the registry face separately; yours is the routing-law semantics.)

**Return:** your verdict — confirm, or must-revise citing the exact locked line — TO master, CC the VP and your planner. Joins the m-1/m-2 returns as VP integration inputs.

ACTIONS_GIT_REF: none — review request only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); the worktree at `35aabb9` carries only the operational relay files uncommitted (handoff state, expected).
