## PLAN-REVIEW — s11 r3 approved: master's split ruling, corrected byte loci, and the 15-row fence are all-in

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-build-plan-review-r3
PARENT_DISPATCH_ID: s11-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair plan-review; merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-plan/PLAN-planner-20260714-144500.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: approve r3 — the bounded table-truth fold closes both r2 findings, all nine cleanup-card items are honestly accounted for, and every licensed implementation seam is inside the standing 15-row fence; issue only the lineage-valid mechanical token

PLAN_REVIEW_VERDICT: approve

I re-reviewed r3 against the live bytes at `frank/` `main@d91fcfb340b029c39c8493084ce2f227409aa546`, master's split ruling `s11-build-escalate-fence/RECONCILE-orchestrator-planner-20260714-143010.md`, and both prior review findings. The plan is now specific, executable, verifiable, and within the ruled boundary.

## Findings closed

1. **Item 6 is byte-true.** The plan names exactly the five system-to-operator envelope composition sites: `internal/engine/approval.go:238-254`, `internal/engine/expiry.go:232-248`, `internal/engine/resummon.go:261-277`, `internal/engine/odb.go:73-109`, and `internal/obligation/obligation.go:190-226`. The two `loop.go` fault records are correctly excluded because they carry no operator address and belong to a different mechanism. Item 3 precedes item 6, followed by a fresh five-site census.
2. **Item 5 is causally bounded.** The executable per-resummon path is `ResummonScheduler.Emit` -> `outcomeForContentHash` -> `tables.Build` at `internal/engine/resummon.go:229-234`, with only the scheduler constructor/state and, if injection requires it, `cmd/frank/main.go` licensed as the upstream snapshot-supply seam. The dedupe outcome must remain byte-identical.
3. **The reconciliation arithmetic is true.** Seven RULING-C items plus item 8 granted by RULING A equals eight retained T8 items; item 2 is explicitly rescoped by RULING B, so all nine original cleanup-card items are accounted for. `internal/observe/` remains out and the shared soft-expiry arbiter remains a named post-Step-2 m-7+m-3 design-cell carry.

## Scope result

SCOPE_DIFF:
- frank/internal/engine/ -> in
- frank/internal/obligation/ -> in
- frank/internal/store/projections.go -> in
- frank/internal/store/genesis.go -> in
- frank/internal/bounce/ -> in
- frank/internal/tables/ -> in
- frank/internal/migrate/ -> in
- frank/internal/crashpoint/ -> in
- frank/internal/recover/ -> in
- frank/internal/config/config.go -> in
- frank/internal/fieldspec/registry_test.go -> in
- frank/internal/executor/executor.go -> in
- frank/cmd/frank/main.go -> in
- frank/test/fixtures/ -> in
- frank/.relays/s11/ -> in
SCOPE_DIFF_RESULT: all-in

## Approval boundary

The other `tables.Build` call sites enumerated by r3 are evidence-only context, not edit licenses. If the named `resummon.go` scheduler/state seam plus conditional `cmd/frank/main.go` wiring cannot supply the snapshot while preserving behavior, implementation must stop, report the causal insufficiency to `s11.planner`, and obtain an amended reviewed plan before touching another build site. A file's presence in the broad scope table does not self-license an unnamed edit.

All other locked boundaries stand: g1/T6 is closed and builds to integrated section B; T5 remains held on g2; T10 remains held on dc; FINDING-4 requires report-and-hold for any open exit gate; the T8 executor seam is independent of g2 and distinct from T5; `Spawn`, fault tokens, verdict logic, and the ruled byte-untouched paths remain untouched; merge is operator-only.

This approval advances the pair only to token issuance. It is not implementation authority by itself. `s11.planner` may now issue the lineage-valid mechanical implementation token under `s11-build-token`, carrying the reviewed lock, all-in scope result, conditions, gates, and stop rules.

## Verification

- Exact r3 PLAN lint: `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/s11-build-plan/PLAN-planner-20260714-144500.md` — OK.
- Plan-of-record hash: `6af929c3af1496440078351c68ab1e67af4f3211f9396b5dc8227bfcaed02d3d`.
- Live envelope census: exactly the five operator-addressed sites listed above; `loop.go:384-418` has no operator address.
- Live resummon consumer: `internal/engine/resummon.go:194-234`; `Emit` reaches the single per-emit build through `outcomeForContentHash`.
- Base/head before report: `frank/` `main@d91fcfb340b029c39c8493084ce2f227409aa546`, clean.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no branch, no commit, no implementation token, no merge.
FINAL_GIT_STATUS_SHORT: none — `frank/` clean at `main@d91fcfb`; cwd root is a docs workspace and not a git repo.
Next requested action: s11.planner issues the lineage-valid mechanical `DISPATCH IMPL` token under `s11-build-token`, preserving this review's scope and approval boundary. s11.implementer remains review-only until directly addressed by that token.
