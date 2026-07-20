## PLAN-REVIEW — s11 r2 must revise: master's split ruling is folded, and the 15-row fence is all-in, but the RULING-C locus table is not byte-true for items 5/6

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-build-plan-review-r2
PARENT_DISPATCH_ID: s11-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair plan-review; merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.2
IN_REPLY_TO: master/relays/s11-build-plan/PLAN-planner-20260714-143700.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: must-revise bounded to RULING-C table truth — RULINGS A/B are folded exactly and the file fence is all-in, but item 6's claimed x5 system-to-operator map omits engine/odb.go and substitutes two non-operator loop fault records, while item 5's per-resummon Build row sweeps unrelated global Build sites without a causal seam license; correct the locus table/count before token

PLAN_REVIEW_VERDICT: must-revise

I re-reviewed only r2's three deltas against `s11-build-escalate-fence/RECONCILE-orchestrator-planner-20260714-143010`, my r1 finding, and the live tree at `d91fcfb`.

RULING A passes: `executor.go` now has a g2-independent T8 `finalizeRun`/preserve seam distinct from gated T5, with the required byte-untouched boundaries. RULING B passes: item 2 is explicitly rescoped out, T8 honestly becomes eight-of-nine, `observe/` remains out, and the post-Step-2 m-7+m-3 carry is named. The original F1 is closed.

The file-level fence also passes:

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

Approval still cannot issue because the byte-grain locus reconciliation master required in RULING C contains two executable inaccuracies.

## F1 — item 6's “x5 system-to-operator” locus map is false at the bytes

The r2 row (`PLAN-planner-20260714-143700.md:46`) names:

- approval
- expiry
- resummon
- `loop.go:387,411`
- obligation

That is six cited composition loci if both loop records count, not five. More importantly, neither loop record is system-to-operator: `loop.go:384-395` constructs an internal-fault `held` record with no `Envelope.To` and no encoded `TO`; `loop.go:408-418` does the same for the rejected fault record. Refactoring them under the system-to-operator builder would change or confuse a different mechanism.

The actual five pre-refactor system-to-operator composition sites at `d91fcfb` are:

1. `internal/engine/approval.go:238-254`
2. `internal/engine/expiry.go:232-248`
3. `internal/engine/resummon.go:261-277`
4. `internal/engine/odb.go:73-109`
5. `internal/obligation/obligation.go:190-226`

The row omitted `engine/odb.go`, which is the live `RenderODB` half already named by item 3. Correct the x5 map to these bytes and state the item-3/item-6 order so the one-ODB-builder consolidation does not make the later call-site count ambiguous. This correction stays inside the existing `engine/` + `obligation/` fence; no master escalation is needed.

## F2 — item 5's per-resummon locus is broadened to unrelated global `tables.Build` calls

The locked card and the master ruling say **drop `tables.Build` per resummon emit**. At the live path, `ResummonScheduler.Emit` calls `outcomeForContentHash` (`resummon.go:202-205`), whose `tables.Build` is at `resummon.go:229-234`. That is the per-emit build.

The r2 row (`PLAN-planner-20260714-143700.md:45`) additionally names `resummon.go:108,174`, both `submit.go` builds, three `loop.go` builds, and both `obligation.go` sites. Those serve ArmParked, due checks, submit/revalidation, loop initialization/quarantine/fallback, and obligation fallback respectively; the PLAN does not say whether they are evidence-only context or authorized edits necessary to supply the resummon snapshot. As written, the row can be read as a license to remove every listed build, broadening a one-path cleanup into a global table-lifecycle refactor.

Correct item 5 to name the actual per-emit consumer and the exact upstream snapshot-supply seams it requires. If any additional `tables.Build` must change, state the causal supply edge and behavior-preserving boundary for that site; otherwise remove it from the executable locus row. Do not use a call-site census as an edit license.

## Count/summary cleanup in the same bounded fold

RULING C asked for seven previously-unruled items in-fence; item 8 was separately granted by RULING A. The r2 table contains eight retained rows, while its heading says “all seven remaining” and §6 asks me to confirm a “seven-item table” (`:38`, `:81`). Make the arithmetic explicit: **seven RULING-C items + one RULING-A item = eight retained T8 items; item 2 rescoped = nine total accounted for.** This is the plan's own verify-the-summary-line discipline, not a new acceptance criterion.

No other blocker found in the bounded r2 pass. The g1/g2/dc state, T6 §B contract, FINDING-4, execution order, evidence capture, conditions (a)–(g), operator-only merge, and r1 accepted ground remain unchanged.

## Verification

- Exact r2 PLAN lint: `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/s11-build-plan/PLAN-planner-20260714-143700.md` — OK.
- Plan-of-record hash remains `6af929c3af1496440078351c68ab1e67af4f3211f9396b5dc8227bfcaed02d3d`.
- Master ruling read in full: `s11-build-escalate-fence/RECONCILE-orchestrator-planner-20260714-143010.md`.
- Live envelope census: the five operator-addressed sites listed above; `loop.go:384-418` has no operator address.
- Live resummon path: `resummon.go:194-234`; unrelated build sites classified from their callers.
- Base/head before report: `frank/` `main@d91fcfb340b029c39c8493084ce2f227409aa546`, clean.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no branch, no commit, no implementation token, no merge.
FINAL_GIT_STATUS_SHORT: none — `frank/` clean at `main@d91fcfb`; cwd root is a docs workspace and not a git repo.
Next requested action: s11.planner performs one bounded r3 table-truth fold (correct item 6's five sites, narrow/justify item 5's supply path, true the seven-plus-one count), re-runs the same all-in scope diff, and returns r3 for review. Do not issue the delegated token from this review.
