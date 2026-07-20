## SITREP — s7 fidelity roll-up + fold routing: rows 1/2/4/5 CONFIRMED (3 seats); two test-only row-6 folds routed to the s7 pair NOW; row 3's fold DEFERRED — the fidelity pass exposed a pre-existing PRODUCTION defect (F-S7-R2-COLGRAIN), which routes to its own bounded lane, never through s7's fence

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the F-S7-R2-COLGRAIN production-fix lane needs the operator's staffing ruling (a B10-shape m-2 pair build, or a fresh build seat); requested via the operator CC
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-fidelity-m4/SITREP-implementer-20260710-113340.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-1.implementer, m-2.implementer, m-4.implementer, m-2.planner, m-4.planner
SUBJECT: to the s7 pair — fold the two row-6 fidelity findings under your standing dispatch (test-only, fold-in authority per the REVIEW-FOLD loop); do NOT touch row 3 yet (its any_row negative goes red on current production — the fix is a separate lane); re-reviews after: m-1 row 6 only, m-2 row 6 now / row 3 later, m-4 row 3 later

**Fidelity roll-up at `s7-inv-catalog@35aabb9`:** rows 1 (m-2), 2, 4, 5 (m-1) — **CONFIRM**, no re-review needed unless changed. Row 6 — must-revise on both halves (m-1 + m-2), test-only. Row 3 — must-revise (m-2 + m-4, independently convergent), **blocked on production**.

**To m-7.implementer — fold these two under your standing s7 authority (test-only; the REVIEW-FOLD loop of `PLAN-…-032426`; pair review + red/green per your R-1 pattern):**
1. **Row 6, m-1 half (the canonical census is incomplete):** keep the six seat-delivered families, carve-outs, sink census, and existing negatives unchanged. Add the **complete canonical path-family census** — at minimum `records`, `staging`, `journal`, `projections`, `mailboxes`, `outbox`, `binding`, `quarantine`, `conductor.lock`, `config/engine.json`, `config/fieldspec/registry.json` — **mechanically tied to the production layout** (derive from `store.Open`'s created set / a live initialized scratch root exercising the lazy homes), so a newly introduced canonical home turns the check red until censused. Add a **table-driven planted-path negative for every censused family**, using a root different from the positive captures. (Source: `s7-fidelity-m1/SITREP-implementer-20260710-113129.md` §row-6.)
2. **Row 6, m-2 half (the bounce-reason family must consume the live surface):** add a **live rejected-submit capture** asserting `Outcome.Detail == the stored rejection Body` (D-2 parity consumed by the named row, not by a neighboring package), scan that `Detail` as the `bounce-reason` family, and add a **family-local planted negative** (canonical path bytes planted into that capture must fail `scanSurfaceCorpus`). (Source: `s7-fidelity-m2/SITREP-implementer-20260710-113112.md` §row-6.)
Then: rerun the focused laws + the full invariants package + the uncached battery, pair review, and report; I re-route **row 6 to m-1 + m-2 for narrow confirmation**.

**Row 3 — HOLD (do not fold yet).** Both m-2 and m-4 require synthetic `any_row:routing_assignments.chosen_model` required/visible negatives — and both verified those negatives would go **RED on current production**: `internal/fieldspec/predicate.go:136-152` validates only the parent array, `routing_assignments` is gate-referenceable, so a model-identity **column** predicate compiles. That is a **pre-existing production enforcement gap**, now registered as **F-S7-R2-COLGRAIN**: the C1 column-grain carry's trigger condition ("the moment the grammar can address a row-array column") is already met, so the model-identity column guard is due now, not at Step-3. Per your fence and the B9 discipline: the s7 branch stays test-only; **the production guard routes to its own bounded lane** (m-2-owned mechanism, m-4 co-sign; staffing = the operator's call, CC'd). The row-3 fold (the negatives + any catalog-text truthing) lands **after** the guard exists, then re-routes to **both m-2 and m-4**.

**Sequencing effect:** s7 integration (VP) waits on: the two row-6 folds + narrow m-1/m-2 confirms + the row-3 resolution (guard lane → fold → m-2+m-4 confirms). If the operator prefers, the VP could integration-review everything-but-row-3 with row 3 held open — my recommendation is to wait; the slice is small and the guard lane is bounded.

Next requested action — m-7.implementer: the two row-6 folds. Operator (CC): the F-S7-R2-COLGRAIN staffing ruling — my recommendation: **B10-shape, the m-2 pair as the build pair** (implementer writes the column-grain guard + its red-first fixtures; planner pair-reviews; m-4 + m-7 fidelity; master integrates; test fold back in s7 after). VP (CC): visibility now; your integration package will include the fold + confirm chain.

ACTIONS_GIT_REF: none — routing only; my verification runs were read-only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `1d3e92c`; the s7 worktree at `35aabb9` carries only the operational relay files uncommitted.
