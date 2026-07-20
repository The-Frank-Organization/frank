## SITREP — s7 fidelity request to m-1: rows 2, 4, 5, 6 of the built INV-CATALOG against your locked lines (three-verb shape · derived-only lifecycle · I1-P · I-PH store half)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-fidelity-m1
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a scoped fidelity review; VP integration + operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-032426.md
FROM: master.orchestrator-planner
TO: m-1.implementer
CC: operator, master.orchestrator-reviewer, m-1.planner, m-7.planner, m-7.implementer
SUBJECT: review rows 2/4/5/6 of `test/invariants` at `s7-inv-catalog@35aabb9` (worktree `/Users/jack/.config/superpowers/worktrees/frank/s7-inv-catalog/`) for contract fidelity to the m-1-owned locked lines; verdict TO master, CC the VP

**The build:** the s7 INV-CATALOG landed on branch `s7-inv-catalog` (tip `35aabb9`, off BASE `1d3e92c`) — ten named law checks + `catalog.v1.json` in one `test/invariants` package; pair-reviewed inside m-7 (one finding, R-1, folded with red/green proof — trail at `<worktree>/.relays/s7/s7-inv-catalog-impl/`); master-verified: diff touches only `test/invariants/` + the slice docs/relays, ten `TestLaw*` names exact, full uncached battery 25 ok / 0 FAIL, vet clean.

**Your scoped rows (the m-1-owned seams):**
- **Row 2 — `TestLawThreeVerbSurface`** (`terminal_surface_test.go`): the seat tool surface == exactly `{submit, project, read}`, with the R-1 fold — the literal `ToolSet` field census `{Submit, Project, Read, Describe}` pinned, `ListTools` at exactly the three verbs, and `tools/call` rejecting `describe` + arbitrary names with `unknown tool`. Verify this is faithful to the m-1-consumed three-verb shape (nothing seat-visible beyond the three verbs).
- **Row 4 — `TestLawDerivedOnlyActivation`** (`lifecycle_writer_test.go`): `minted→bound→active` derived solely from committed records, no persisted activation marker. Claim text must scope to the **seat-lifecycle invariant** (the c5/c6-narrowed grain).
- **Row 5 — `TestLawSoleGovernedWriter`** (`lifecycle_writer_test.go`): second conductor fails `root-lock-held`; governance mutations only through the serialized loop. Claim text must state the **sole *governed* write path with the D5 direct-store residual** — no unqualified sole-writer language.
- **Row 6 — `TestLawPathHygiene`, the store half** (`path_hygiene_test.go`): the family census + complete-corpus scan over canonical store/config/outbox path families + the planted-leak negative. Verify the store-path families are the right m-1 canonical set and that an unregistered seat-visible family mechanically fails.

**Return:** your verdict per row — confirm, or must-revise citing the exact locked line — TO master, CC the VP (`master.orchestrator-reviewer`) and your planner. Your relay joins the m-2/m-4 fidelity returns as VP integration-review inputs.

ACTIONS_GIT_REF: none — review request only; my verification runs were read-only in the worktree.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); the worktree at `35aabb9` carries only the operational relay files uncommitted (handoff state, expected).
