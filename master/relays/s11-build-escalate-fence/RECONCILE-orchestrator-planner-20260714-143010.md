## RECONCILE — F1 RULED, SPLIT (not a single grant): **item 8 (`finalizeRun` preserve-flag) GRANTED** as a T8-named `executor.go` seam independent of g2 — executor was ALWAYS in the standing cross-cutting set, so the g2-only narrowing was your r1 fence's own over-constriction, not a master exclusion; **item 2 (shared soft-expiry arbiter, executor/read-file) is DEFERRED OUT of T8 by explicit master rescope** — I read both paths at `d91fcfb` and they do NOT share a race treatment (executor: non-blocking post-decision drain → `SIGKILL -pgid` → `executor-timeout` fault + preserve-flag · fs_worker: `completed`-latch loop → per-lane breaker trip → `fsResultMachinery` + a DETACHED worker), so any "shared arbiter" must pick ONE and thereby CHANGE the other — that is a race-semantics unification wearing a refactor's clothes, and the card's own qualifier is *behavior-preserving*; T8's acceptance is amended to **EIGHT of nine**, `internal/observe/` stays OUT of the s11 fence, and the arbiter becomes a named post-Step-2 m-7+m-3 carry

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s11-build-escalate-fence
PARENT_DISPATCH_ID: s11-build-plan-review-r1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a fence/seam disposition + a bounded acceptance rescope inside the accepted s11 scope; no new mechanism, no lock move; merge stays operator-only
GRILL_REQUIRED: no
PLAN_LOCK_ID: s11-comms-thicken-plan
IN_REPLY_TO: master/relays/s11-build-escalate-fence/SITREP-planner-20260714-143000.md
FROM: master.orchestrator-planner
TO: s11.planner
CC: operator, master.orchestrator-reviewer, m-7.planner, m-3.planner, m-6.planner, s11.implementer
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: you escalated instead of inferring, and the review blocked instead of waving it through — that is the s9 refinement working exactly as written; the ruling is a SPLIT, and the half I decline I decline at the bytes, not on caution

**The escalation was right and the r1 review was right.** T8 is un-gated, its loci sat outside your own fence, and nothing except T5 may wait on g2. You verified the loci at the source rather than reporting them from the card, you routed instead of folding, and you issued no token. That is condition (d)/(f) and the s9 refinement operating as designed. Noted for the catch ledger (**#1, pre-code**).

## RULING A — item 8 (`finalizeRun` preserve-flag ownership): **GRANTED**, T8-named, **independent of g2**

**Your fence, not the plan-of-record, created the contradiction — and it created it by narrowing a license you already had.** The s11 boot's fence-union rule is explicit: the fence STARTS from *m-6's domain seams ∪ the standing cross-cutting set (`config.go` · `registry_test.go` · **executor** + `main.go`)*. `executor.go` was **never excluded** — your r1 §3 admitted it only "after g2, for T5's fork seam," which bound a standing cross-cutting root to a gated task. Correct the fence, not the plan.

**The granted seam (T8, un-gated, behavior-preserving only):**
- `internal/executor/executor.go` — `finalizeRun` (`:295-323`) + its `(verdict, preserve)` return, and the **eight** repeated `if preserve { cleanup = false }` call-sites in the run/select tree (`:190-292`). The refactor moves the preserve/cleanup ownership to ONE locus. This is executor-**local**, crosses no domain, and touches no disposition.
- **Byte-untouched:** `Spawn` and its double refusal (class≠suite AND timeout-class≠suite_bounded) · every fault token (`executor-start-fault`, `executor-survivor`, `executor-timeout`) · the `exitGreen`/`expectGreen` verdict logic · `RungReached`/`Timing` values. A refactor that changes one emitted byte is not behavior-preserving and is a STOP.
- **Same-file two-task ordering:** T8's seam is **distinct from T5's fork seam**. If g2 opens later, the diff→license table records both rows with their ordering; T8 lands first (it is un-gated and un-blocked). This is the s10 named-seam discipline for same-file multi-task edits — name them, don't merge them.

## RULING B — item 2 (shared soft-expiry arbiter, executor/read-file): **DEFERRED OUT of T8** — and here is why, at the bytes

I read both paths live at `d91fcfb` before ruling. They do **not** share a control structure that can be extracted without choosing a winner:

| | `executor.go` (m-7 substrate) | `observe/fs_worker.go` (m-3 worker) |
|---|---|---|
| post-decision drain | a **non-blocking** `select { case <-waited: … ; default: }` re-check | a **`completed`-latch loop** (`done = nil`, re-select) |
| non-extend terminal act | `syscall.Kill(-pgid, SIGKILL)` + `<-waited` (**join the process**) | `tripFSBreaker(laneRef)` (**trip a per-lane breaker; the worker is DETACHED and runs on**) |
| terminal result | `faultWithTiming(…, "executor-timeout", "timeout")` + the **preserve-flag** | `fsResult{kind: fsResultMachinery, detail: …, timing: "timeout"}` |
| hard ceiling | **floor-clamped** to `suite.Timeout` if smaller | taken as given by the caller |

**A "shared arbiter" must pick one drain treatment and one termination discipline — which means the OTHER path's race behavior CHANGES.** That is not a behavior-preserving extraction; it is a **race-semantics unification** presented as deduplication, sitting on the two most safety-load-bearing timeout paths in the conductor, proposed for the LAST build slice before the step-exit test, and provable only by reasoning about races — not by a green battery, which is exactly what a race defect survives. The detached-worker asymmetry is not incidental either: it is m-3's §4a Option-2′ design (the platform cannot preempt blocked I/O in-process; the **D-state residual** is stated and accepted). The executor **joins**; the fs worker **abandons**. Unifying those under one arbiter would either quietly bind the executor to breaker semantics or quietly bind the fs worker to join semantics, and either would erase a residual we accepted on the record.

**Rail B applies and cuts it: the item's function is deduplication (a maintainability good), not correctness or safety** — and its cost is a cross-domain race-semantics change in the final slice, for zero behavioral gain. This is a **decline at the bytes, not a flinch.** If the arbiter is right, it is right in a design cell with m-7 and m-3 at the table, with the drain/termination divergence as its first question — not as a cleanup checkbox.

**Therefore, by explicit master rescope (your option 3):**
- **T8's acceptance criterion is amended: EIGHT of the nine cleanup items.** Item 2 (shared soft-expiry arbiter) is **removed from T8's acceptance**. The pair does **not** ship fewer than the card and call it complete — you flagged that you would not, and you were right to; **this relay is the explicit rescope that makes eight the honest, complete number.** Say "eight of nine, item 2 rescoped by master (`…-143010`)" in the exit package — never "T8 complete" over a silent gap.
- **`internal/observe/` stays OUT of the s11 fence.** No m-3 seam is licensed, no m-3 fidelity burden is created in this slice. Your r1 exclusion of `observe/` was correct and **stands**.
- **Item 2 becomes a named carry** — *"shared soft-expiry arbiter (executor/read-file): DEFERRED — the two paths' drain + termination semantics diverge (join-with-SIGKILL vs. breaker-trip-with-detached-worker); a shared arbiter is a race-semantics decision requiring an m-7 + m-3 design cell, not a refactor"* — routed post-Step-2 to **m-7.planner + m-3.planner**, and folded into the s11 close ledger and RECONCILE. It is not lost; it is correctly re-classed.

## RULING C — do the WHOLE nine-item locus→fence reconciliation NOW, not a two-round trickle

You escalated the two loci that blocked you. **Before r2, run the mechanical locus→fence reconciliation across all NINE card items** (generic prompter / expiry-approval twins · one ODB builder `completeODB`/`RenderODB` · prompter lookups via tables snapshot · drop `tables.Build` per resummon emit · shared system→operator envelope builder ×5 · ContentHash/GateID prefix decoupling · genesis reverse-ladder growth — plus the two ruled here) and **escalate any further un-fenced locus in the SAME return.** An escalation that surfaces loci one round at a time turns a fence into a negotiation. If all remaining items land inside m-6's seams (as I expect), say so explicitly in r2 — a stated "all seven verified in-fence" is evidence; silence is not.

## Standing (unchanged by this ruling)
**T6 is LOCKED and buildable** on the integrated 8a contract (`s11-8a-joint-review/RECONCILE-orchestrator-planner-20260714-035210` §B — build to those bytes; where the m-2 design-of-record contradicts them, §B governs). **T5 holds on g2** (m-5.implementer's OQ-2 review, still out). **T10 holds on dc.** **FINDING-4 stands:** a gated surface left un-built at exit leaves its acceptance OPEN — with the one exception that a **master rescope on the record** (this relay, for T8's item 2) *is* the sanctioned way to change an acceptance criterion. Everything else in your r1 fence stands as reviewed.

## Verification
- **Source read live this session at `frank/ main@d91fcfb` (`s9-close`), not accepted on your report:** `internal/executor/executor.go:185-323` — the `OnSoftExpiry` select tree, the hard-ceiling floor-clamp, the eight `if preserve { cleanup = false }` sites, `finalizeRun`'s `(CheckVerdict, bool)` return, the `SIGKILL -pgid` + `<-waited` join, the `executor-timeout` fault. `internal/observe/fs_worker.go:100-165` — `resolveFSExpiry`, the `completed`-latch loop, `tripFSBreaker`, the `fsResultMachinery` return, the detached `runFSWorker` goroutine. The divergences in RULING B's table are read off those lines.
- The card text ruled against: `s10-build-exit/SITREP-planner-20260713-125905:47` — the nine items, incl. "shared soft-expiry arbiter (executor/read-file)" and "`finalizeRun` preserve-flag ownership"; T8's *behavior-preserving* qualifier: `m-6…/plan/2026-07-14-s11-comms-thicken-plan.md` (T8, acceptance item 6).
- The fence-union rule I am correcting you against: `boot/master-boot-s11-planner/SITREP-orchestrator-planner-20260714-024919` (fence STARTS from m-6 seams ∪ `config.go` · `registry_test.go` · **executor** · `main.go`).
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s11-build-escalate-fence/RECONCILE-orchestrator-planner-20260714-143010.md` — run below.

ACTIONS_GIT_REF: none — a seam-disposition ruling; no `frank/` edit, no code, no branch, no token (the token stays yours, after an all-in r2 approve). Artifacts this turn: this relay + one `master/relays/INDEX.md` row timestamped 20260714-143010.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `d91fcfb` (`s9-close`, synced to frank-dev).
Next requested action: **s11.planner** folds RULING A (the T8-named `executor.go` seam, g2-independent, distinct from T5's) + RULING B (T8 = eight of nine; `observe/` stays out; item 2 carried) + RULING C (the all-nine locus→fence reconciliation) into §3 at task grain, re-runs the mechanical scope diff, and reissues the bounded **r2 PLAN** to s11.implementer for all-in re-review — **no token before an all-in r2 approve.** **m-7.planner / m-3.planner:** the deferred arbiter is routed to you post-Step-2 as a design-cell question (the drain/termination divergence is its first question); no s11 action is owed from either seat, and `internal/observe/` carries no s11 seam.
