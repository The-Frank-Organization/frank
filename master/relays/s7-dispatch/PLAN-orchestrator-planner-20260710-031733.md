## PLAN — the corrected s7 implementation dispatch, to the dedicated build seat s7-core.implementer: build the INV-CATALOG per the VP-approved r2+r3 plan at BASE `frank/main@1d3e92c`, branch `s7-inv-catalog`, test-only; the dispatch token stands bare below

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the VP phase-opener gate is passed (`RECONCILE-orchestrator-reviewer-20260710-030737`, approve); the operator's role correction is folded (`PLAN-…-031728`); merge stays operator-gated later
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-030737.md
FROM: master.orchestrator-planner
TO: s7-core.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-7.implementer, m-1.planner, m-2.planner, m-4.planner
SUBJECT: s7 INV-CATALOG build dispatch to s7-core.implementer (a fresh operator-booted pair-Implementer session; the SOLE s7 code writer) — binding docs: r2 `PLAN-…-023635.md` + the four r3 replacements in `PLAN-…-030148.md` (rows 6 and 9 use the r3 text) + the VP approval `RECONCILE-…-030737.md` + the retraction/correction `PLAN-…-031728.md`; the token stands bare below

DISPATCH IMPL

s7-core.implementer — you are a fresh build seat running the **agent-pair-implementer** role, the **sole s7 code writer**. Your guide for the m-7-owned mechanism rows (8–10) is **m-7.planner** (domain questions route there); master (this seat) owns coordination and integration; **no planner seat and no master domain-pair seat edits `frank/`** (the operator's standing role fence — the prior dispatch to m-7.implementer is retracted by `PLAN-…-031728.md`).

**Onboard (before any code):**
1. Read the charter (`CLAUDE.md` at the workspace root) and the four binding documents named in the SUBJECT. The effective plan = r2 with the r3 replacements; where they differ, **r3 wins** (rows 6 and 9).
2. Scaffold your slice tree in `frank/` per `sprint-doc-setup` — `docs/sprints/2026-07-10-s7-inv-catalog/` with its `.relays/s7/` substrate; your own boot/IMPL relays live there (Cardinal rule 2: build relays live with the code; only master's dispatch acts live in `master/relays/`).
3. **Intake acknowledgment (your plan-review affordance):** before writing code, relay any blockers/ambiguities you see in the binding plan back to master — the plan was VP-gated at master level (three rounds), so this is your chance to flag, not a re-plan.

**Work order:**
1. Branch `s7-inv-catalog` off **BASE `frank/main@1d3e92c`** (clean; battery-verified 24 ok + 2 no-test-files, vet clean).
2. Build the single `test/invariants` package: the **ten named checks exactly as contracted** (the r2 table; r3 text for rows 6 and 9) + the versioned catalog artifact (law list, owners, the single-writer/owner-fidelity-on-change convention recorded in the artifact — **no §7/governance claim**; that property completes at s8).
3. Rows 4 and 5 carry the claim-grain bounds **verbatim** in their row text (derived-only = the seat-lifecycle invariant; I1-P = the sole *governed* write path with the D5 direct-store residual stated).
4. **Row 6 acceptance behavior (VP finding 4 — acceptance, not commentary):** the seat-delivered-family census must be mechanical — an unregistered seat-visible family must turn the check red until the census is updated; integration will reject a hand-maintained list that cannot demonstrate that. Include the planted-leak negative proving the scanner bites.
5. **Row 9 uses the r3 at-most-one grain** — all five clauses; a pending intake with zero outcomes is legitimate state, never a violation.
6. The **red-battery demo**, command-pinned: weaken one selected law on a scratch branch → `go test -count=1 ./test/invariants` FAILS naming the law → discard the scratch → the same command green. Capture the transcript as evidence.
7. Scope fence: **test-only** — zero production-code change, no registry edit, no new record kinds. A genuine defect exposed while naming a law is a finding relayed to master, never an in-slice fix.

**Return (your IMPL report, TO master, CC the VP + operator + m-7.planner + m-7.implementer):** branch + tip SHA · `go test -count=1 ./test/invariants` output (ten rows named, green) · the red-then-green transcript · full uncached `go test ./...` + `go vet` at your tip · the diffstat proving test-only · the catalog artifact path. On your report, master routes the fidelity reviews — m-1 (rows 2/4/5/6) · m-2 (rows 1/3/6) · m-4 (row 3) · **m-7.implementer (rows 8–10 + the harness, the pair's contract-fidelity function)** — then the package goes to the VP for integration review. **No merge on your authority** — integration → the operator merge gate.

Blocked or ambiguous → relay the question to master; do not widen scope to resolve ambiguity.

ACTIONS_GIT_REF: none by this relay — authority grant only; the branch/commits are yours from here.
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main` clean at `1d3e92c` at dispatch time; cwd is not a git repo (docs workspace).
