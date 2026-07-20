## PLAN — the s7 implementation dispatch to m-7.implementer (the sole s7 code writer): build the INV-CATALOG per the VP-approved r2+r3 plan at BASE `frank/main@1d3e92c`, on branch `s7-inv-catalog`, test-only; the dispatch token stands bare below

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the VP phase-opener gate is passed (`RECONCILE-orchestrator-reviewer-20260710-030737`, VERDICT: approve); merge stays operator-gated later
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-030737.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-1.planner, m-2.planner, m-4.planner
SUBJECT: s7 INV-CATALOG implementation dispatch; the binding plan = r2 `PLAN-orchestrator-planner-20260710-023635.md` PLUS the four r3 replacements in `PLAN-orchestrator-planner-20260710-030148.md` (rows 6 and 9 use the r3 text, NOT the superseded r2 rows), under the VP approval `RECONCILE-orchestrator-reviewer-20260710-030737.md`; you are the sole s7 code writer; the dispatch token stands bare below

DISPATCH IMPL

m-7.implementer — the s7 phase-opener plan is VP-approved; implementation authority is granted to you alone by this relay. The three binding documents (cite them in your report): the r2 plan `s7-dispatch/PLAN-orchestrator-planner-20260710-023635.md` (the ten-row law→test→mechanism contract, execution model, fence, acceptance) · the r3 delta `s7-dispatch/PLAN-orchestrator-planner-20260710-030148.md` (**rows 6 and 9 replaced** — I-PH exhaustive-with-census + planted-leak negative; A-2 at-most-one grain with pending-zero legitimate — plus the staged-governance statement and the bound fidelity scopes) · the VP approval `s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-030737.md` (approval scope + integration bar).

**Work order:**
1. Branch `s7-inv-catalog` off **BASE `frank/main@1d3e92c`** (clean, battery-verified: 24 ok + 2 no-test-files, vet clean).
2. Build the single `test/invariants` package: the **ten named checks exactly as contracted** (r2 table; r3 text for rows 6 and 9) + the versioned catalog artifact (law list, owners, the single-writer/owner-fidelity-on-change convention recorded in the artifact — **no §7/governance claim**; that completes at s8).
3. Rows 4 and 5 carry the claim-grain bounds verbatim in their row text (derived-only = the seat-lifecycle invariant; I1-P = the sole *governed* write path with the D5 direct-store residual stated).
4. **Row 6 acceptance behavior (VP finding 4 — this is acceptance, not commentary):** the family census must be mechanical — an unregistered seat-visible family must turn the check red until the census is updated; the integration review will reject a hand-maintained list that cannot demonstrate that behavior. Include the planted-leak negative proving the scanner bites.
5. The **red-battery demo**, command-pinned: weaken one selected law on a scratch branch → `go test -count=1 ./test/invariants` FAILS naming the law → discard the scratch → the same command green. Capture the transcript as evidence.
6. Scope fence: **test-only** — zero production-code change, no registry edit, no new record kinds. A genuine defect exposed while naming a law is a finding relayed to master, never an in-slice fix.

**Return (your IMPL report, TO master, CC the VP + operator + m-7.planner):** branch + tip SHA · `go test -count=1 ./test/invariants` output (all ten rows named, green) · the red-then-green demo transcript · full uncached `go test ./...` + `go vet` at your tip · the diffstat proving test-only · the catalog artifact path. On your report I route the three fidelity requests per the bound scopes (m-1 → rows 2/4/5/6 · m-2 → rows 1/3/6 · m-4 → row 3), then the whole package goes to the VP for integration review. **No merge on your authority** — integration → operator merge gate, per the standing discipline.

Blocked or ambiguous → relay the question to master (m-7.planner is your guide for the engine rows 8–10); do not widen scope to resolve ambiguity.

ACTIONS_GIT_REF: none by this relay — authority grant only; the branch/commits are yours from here.
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main` clean at `1d3e92c` at dispatch time; cwd is not a git repo (docs workspace).
