## PLAN-REVIEW — s11 pair PLAN must revise: the proposed all-in fence makes un-gated T8 impossible by licensing `executor.go` only for gated T5 and excluding the live read-file soft-expiry locus

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-build-plan-review-r1
PARENT_DISPATCH_ID: s11-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair plan-review; merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.1
IN_REPLY_TO: master/relays/s11-build-plan/PLAN-planner-20260714-141100.md
FROM: s11.implementer
TO: s11.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: must-revise — adoption/hash/lineage/gates pass, but §3's named-seam fence contradicts T8: the cleanup card requires the shared executor/read-file soft-expiry arbiter plus `finalizeRun` preserve-flag ownership, while `executor.go` is licensed only after g2 and only for T5, and `internal/observe/` is explicitly OUT; the mechanical result is deviation-present, so no token may issue

PLAN_REVIEW_VERDICT: must-revise

I reviewed the directly addressed PLAN against the byte-pinned plan-of-record, its m-6 approval, the s11 dispatch and boots, the current g1/g2/dc trail, the m-2 rev4 editorial fold, conditions (a)–(g), and the live `frank` tree at `d91fcfb`.

The adoption and gate handling pass: the plan-of-record hashes to the pinned `6af929c3af1496440078351c68ab1e67af4f3211f9396b5dc8227bfcaed02d3d`; g1 is closed and T6 accurately consumes integrated §B; g2 and dc remain open and T5/T10 correctly hold under FINDING-4; merge stays operator-only; no implementation token is present. One blocking scope contradiction remains.

## Mechanical scope check

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
- frank/internal/executor/executor.go :: T8 shared-soft-expiry arbiter and finalizeRun preserve-flag ownership -> OUT
- frank/internal/observe/fs_worker.go :: T8 read-file side of the shared soft-expiry arbiter -> OUT
SCOPE_DIFF_RESULT: deviation-present

The first fifteen rows are the PLAN's proposed block and exist or are valid create-targets at the stated base. The two OUT rows are not new work inferred by this review; they are executable loci of two items already inside T8's locked nine-item cleanup card.

## Blocking finding F1 — T8 is inside scope but its required seams are outside the license

The plan-of-record requires T8 to complete all nine cleanup items, including the **shared soft-expiry arbiter** and **`finalizeRun` preserve-flag ownership** (`master/domains/m-6-human-surface-scheduler/plan/2026-07-14-s11-comms-thicken-plan.md:77-80`). The originating s10 carry makes the first item's boundary explicit: **shared soft-expiry arbiter (executor/read-file)** (`master/relays/s10-build-exit/SITREP-planner-20260713-125905.md:43-47`).

At `d91fcfb`, those are not abstract labels:

- `frank/internal/executor/executor.go:195-292` owns the executor soft-expiry arbitration, and `:295-323` owns `finalizeRun` plus its preserve return.
- `frank/internal/observe/fs_worker.go:107-160` owns the read-file worker's parallel soft-expiry arbitration path (called by `read_file_worker.go`).

But the PLAN's named-seam license says `executor.go` has **no licensed seam until g2 completes**, then licenses it **only for T5 fork spawn/feed/retire**; it also says `internal/observe/` is **not in the fence** (`PLAN-planner-20260714-141100.md:62`). That creates two contradictions:

1. T8 is un-gated and ordered before T11, yet its executor cleanup work is accidentally made dependent on g2, contrary to the locked rule that nothing except T5 waits on g2.
2. Even if g2 returned, its T5-only executor license still would not authorize either T8 executor edit, and the read-file half remains outside the block entirely.

This cannot be treated as a small integration hook: the boot and plan both make seam-grain licensing binding, and conditions (d)/(f) require cross-domain surprises to route and hold. An approval would falsely convert a mechanically dirty fence into a tokenizable `all-in` result.

## Required revision

Before re-review:

1. Route the T8 executor/read-file loci through master under the standing cross-domain rule, because the correction touches m-7's executor substrate and m-3's observe worker.
2. Fold the returned exact file/test seams into §3 at task grain, including a T8 license for `executor.go` that is independent of g2 and a named T8 read-file/observe seam. If the returned ruling names different loci, reconcile those exact loci instead; do not infer them in-fence.
3. Name the same-file ordering explicitly if g2 later opens: T8's executor cleanup seam and T5's fork seam must remain distinct. Re-run the mechanical scope diff; only an `all-in` result may advance to a token.

The alternative is an explicit master rescope removing or deferring the affected cleanup items from T8 and its acceptance criterion. The pair cannot silently ship fewer than nine while claiming T8 complete.

No other blocker found in this pass. T6's §B bytes, the g2/dc holds, FINDING-4, the out-of-scope list, byte-exact terminal enum, R2, I-PH, sequence-honest captures, per-task RED/GREEN discipline, and operator-only merge are preserved.

## Verification

- Exact PLAN lint: `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/s11-build-plan/PLAN-planner-20260714-141100.md` — OK.
- Plan hash: `shasum -a 256 master/domains/m-6-human-surface-scheduler/plan/2026-07-14-s11-comms-thicken-plan.md` — `6af929c3af1496440078351c68ab1e67af4f3211f9396b5dc8227bfcaed02d3d`.
- Base/head: `frank/` `main@d91fcfb340b029c39c8493084ce2f227409aa546`; clean before relay write.
- Gate trail: live `INDEX.md` through the addressed PLAN has g1 closed at `…-035210`, m-2 rev4 folded at `…-140700`, and no later g2 implementer review/completion or dc return; T5/T10 remain held.
- Source proof: `executor.go:195-323`; `observe/fs_worker.go:107-160`; originating cleanup boundary `s10-build-exit/…-125905.md:43-47`; proposed license `PLAN-planner-20260714-141100.md:62`.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no branch, no commit, no implementation token, no merge.
FINAL_GIT_STATUS_SHORT: none — `frank/` clean at `main@d91fcfb`; cwd root is a docs workspace and not a git repo.
Next requested action: s11.planner routes F1 to master for exact T8 executor/read-file seam disposition, folds the ruled fence correction or explicit rescope, and returns a bounded r2 PLAN for re-review. Do not issue the delegated token from this review.
