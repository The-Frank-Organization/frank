## SITREP — s7 fidelity request to m-2: rows 1, 3, 6 of the built INV-CATALOG against your locked lines (terminal-state vocabulary · R2 registry face · I-PH bounce/reason-text half)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-fidelity-m2
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a scoped fidelity review; VP integration + operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-032426.md
FROM: master.orchestrator-planner
TO: m-2.implementer
CC: operator, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-7.implementer
SUBJECT: review rows 1/3/6 of `test/invariants` at `s7-inv-catalog@35aabb9` (worktree `/Users/jack/.config/superpowers/worktrees/frank/s7-inv-catalog/`) for contract fidelity to the m-2-owned locked lines; verdict TO master, CC the VP

**The build:** the s7 INV-CATALOG landed on branch `s7-inv-catalog` (tip `35aabb9`, off BASE `1d3e92c`) — ten named law checks + `catalog.v1.json` in one `test/invariants` package; pair-reviewed inside m-7 (R-1 folded, red/green proven — trail at `<worktree>/.relays/s7/s7-inv-catalog-impl/`); master-verified: test-only diff, ten `TestLaw*` names exact, full uncached battery 25 ok / 0 FAIL, vet clean.

**Your scoped rows (the m-2-owned seams):**
- **Row 1 — `TestLawTerminalEnumByteExact`** (`terminal_surface_test.go`): the terminal set byte-exact `{accepted, rejected, held}`; a forged fourth token rejected typed at submit. Terminal-state vocabulary is your contract surface — verify the byte-exactness grain and the reject leg match your locked lines (incl. the s6 `bounced`→`rejected` unification).
- **Row 3 — `TestLawR2NoModelPredicate`**: no `required_when`/`visible_when`/gate predicate keys on a model-identity field in the live registry; the predicate grammar refuses a synthetic model-keyed atom. Verify against your §5 grammar + the c6/readiness-c4 R2 form (`declared_deviated`-shaped, never model-derived) — the registry face is yours (m-4 reviews the routing face separately).
- **Row 6 — `TestLawPathHygiene`, the bounce/reason-text half** (`path_hygiene_test.go`): bounce/reason text is one of the enumerated seat-delivered families in the census; verify the reason-text family coverage is faithful to your validator's bounce surfaces (D-2 detail parity included) and the planted-leak negative bites on that family.

**Return:** your verdict per row — confirm, or must-revise citing the exact locked line — TO master, CC the VP and your planner. Joins the m-1/m-4 returns as VP integration inputs.

ACTIONS_GIT_REF: none — review request only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); the worktree at `35aabb9` carries only the operational relay files uncommitted (handoff state, expected).
