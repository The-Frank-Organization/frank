## PLAN — the s7 implementation dispatch v3, under the operator's B10 ruling: the m-7 pair IS the s7 build agent-pair — m-7.implementer writes (sole writer), m-7.planner pair-reviews; no new sessions; supersedes both `…-031733` (unbooted seat, withdrawn) and the stand-down clause of `…-031728`; the dispatch token stands bare below

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the VP phase-opener gate is passed (`…-030737`, approve); the execution-seat question is settled by the operator's B10 ruling (2026-07-10, `PROTOCOL-DEVIATIONS.md`); merge stays operator-gated later
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-030737.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-1.planner, m-2.planner, m-4.planner
SUBJECT: s7 INV-CATALOG build dispatch v3 to m-7.implementer as the Implementer half of the s7 build agent-pair (operator ruling B10: a lean slice may use the owning domain pair as the build pair; the floor = at least one real agent-pair) — m-7.planner is the pair Planner (guide + adversarial implementation review); binding docs: r2 `PLAN-…-023635.md` + the r3 replacements `PLAN-…-030148.md` + the VP approval `RECONCILE-…-030737.md` + `PROTOCOL-DEVIATIONS.md` B10; the token stands bare below

DISPATCH IMPL

m-7.implementer — the operator ruled (B10, registered): for this lean, test-only phase-opener the **m-7 pair is the s7 build agent-pair**, no new sessions. You are the **Implementer half and the sole s7 code writer**; **m-7.planner is your pair Planner** — domain guide for the engine rows and **adversarial reviewer of your implementation before it reaches master**. This supersedes the dispatch to the never-booted `s7-core.implementer` (`…-031733`, withdrawn) and, by the operator's ruling, the stand-down clause of the retraction (`…-031728`); the pairs-don't-code fence remains the default outside this ruling.

**Supersession chain, explicit (VP r2's own condition — no executor follows superseded text):** the effective plan = r2 `PLAN-…-023635.md` with the four r3 replacements in `PLAN-…-030148.md` (**rows 6 and 9 use the r3 text**), executed under the VP approval `RECONCILE-…-030737.md` and this v3 seat assignment. Prior seat-assignment text in r2 ("m-7.implementer sole writer" → retracted by `…-031728` → restored here by B10) resolves to: **you write; m-7.planner reviews; no planner seat and no master seat edits `frank/` besides you.**

**Work order (unchanged from the gated plan):**
1. Branch `s7-inv-catalog` off **BASE `frank/main@1d3e92c`** (clean; 24 ok + 2 no-test-files, vet clean).
2. Scaffold the slice tree in `frank/` per `sprint-doc-setup` — `docs/sprints/2026-07-10-s7-inv-catalog/` with `.relays/s7/`; the pair's own relays (your intake-ack, the pair review, your IMPL report) live there (Cardinal #2).
3. **Intake-ack before code**: read the binding docs; relay blockers/ambiguities to master; m-7.planner co-signs the intake (the pair's plan-review affordance — the plan itself is already VP-gated, three rounds).
4. Build the single `test/invariants` package: the **ten named checks exactly as contracted** (r2 table; r3 text for rows 6 and 9) + the versioned catalog artifact (law list, owners, the single-writer/owner-fidelity convention in the artifact — **no §7/governance claim**; that completes at s8).
5. Rows 4/5 carry the claim-grain bounds **verbatim** (derived-only = seat lifecycle; I1-P = sole *governed* write path, D5 residual stated).
6. **Row 6 acceptance behavior (VP finding 4):** the seat-delivered-family census is mechanical — an unregistered seat-visible family turns the check red until the census is updated; a hand-maintained list that can't demonstrate that fails integration. Planted-leak negative included.
7. **Row 9 = the r3 at-most-one grain**, all five clauses; pending-zero is legitimate state.
8. The **red-battery demo**, command-pinned: weaken one law on scratch → `go test -count=1 ./test/invariants` fails naming the law → discard → green. Keep the transcript.
9. Fence: **test-only** — zero production-code change, no registry edit, no new record kinds; defects found are findings to master, never in-slice fixes.

**Return path:** your IMPL report (branch + tip SHA · the ten-rows-green output · the red-then-green transcript · full uncached `go test ./...` + `go vet` at tip · the test-only diffstat · the catalog path) goes **first to m-7.planner for the pair review**, then with its review relay to master. Master then routes the external fidelity set — m-1 (rows 2/4/5/6) · m-2 (rows 1/3/6) · m-4 (row 3) — and the package goes to the VP for integration review. **No merge on the pair's authority** — integration → the operator merge gate.

**VP (CC):** the seat identity moves by explicit operator ruling (B10, registered in `PROTOCOL-DEVIATIONS.md`); your F1 principle stands — exactly one Implementer-role writer, no planner seat coding; the pair adds the cross-review your F1 was protecting. Flag if you consider a re-gate required; absent a flag, the approved plan runs with this seat assignment.

Blocked or ambiguous → relay to master; do not widen scope to resolve ambiguity.

ACTIONS_GIT_REF: none by this relay — authority grant only; the branch/commits are yours from here.
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main` clean at `1d3e92c` at dispatch time; cwd is not a git repo (docs workspace).
