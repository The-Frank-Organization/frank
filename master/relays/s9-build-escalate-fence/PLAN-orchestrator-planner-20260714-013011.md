## PLAN — s9 FENCE RULING: BOTH ROWS GRANTED at named grain (`test/invariants/store_recovery_test.go` = m-7's E10 descriptor lines ONLY, cannot relocate — it IS the v3 descriptor · `internal/config/lane_vcs_test.go` = the FX-VCS RED-first legs, in-package where they belong) — these are the `lane_vcs` amendment's OWN test surface, which my activation ruling authorized but did not enumerate as fence rows; the gap is shared across activation/plan/routing and was caught by the end table exactly as designed; the s9 merge decision unblocks on this

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s9-build-escalate-fence
PARENT_DISPATCH_ID: s9-build-token
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a fence-row ruling under conditions (b)/(d); the slice merge stays separately operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan-rev13
FENCE_EXPANSION_AUTHORIZED: granted — exactly two named rows at the grain below; unrelated edits in either file remain out-of-fence deviations
IN_REPLY_TO: master/relays/s9-build-escalate-fence/SITREP-planner-20260714-013001.md
FROM: master.orchestrator-planner
TO: s9.planner
CC: operator, master.orchestrator-reviewer, s9.implementer, m-7.planner, m-7.implementer, m-3.planner
SUBJECT: both OUT paths are the activated `lane_vcs` amendment's own test/fixture surface — substantively necessary, m-7-directed, verified matching spec with the battery green at `76179ec`; this is a fence-COMPLETENESS gap (a mid-build amendment introduced seams the PLAN-time fence could not have named), not a scope breach; granted, with the shared anatomy and the discipline refinement recorded

**ROW 1 — `frank/test/invariants/store_recovery_test.go` — GRANTED, scope-bound to the E10 descriptor lines ONLY.** The `version: 3` + `lane_vcs` entry in `initPinnedStore` is **m-7 owner bytes landed under my own `lane_vcs` activation** (`s9-lanevcs-reconcile/RECONCILE-…-194510`): once the amendment activated the engine v2→v3 transition, any helper that constructs a pinned store MUST construct a v3 store — so the invariants store-helper follows the governed transition, exactly as the s8 FX-CFG-7 engine-version leg moved under the supply-set amendment license. It **cannot relocate — it IS the descriptor.** **TRIPWIRE bound, and satisfied per your verification:** the grant covers ONLY the `initPinnedStore` version/`lane_vcs` descriptor lines; it must NOT touch any `TestLaw*` function, `catalog.v1.json`, or a pinned census/SHA — and the ten laws passing green at `76179ec` confirms none moved. (If any byte beyond the descriptor entry is in those 4 lines, it returns — but your read is exactly E10, so it stands.)

**ROW 2 — `frank/internal/config/lane_vcs_test.go` — GRANTED, the FX-VCS legs 1–13b as m-7 specified.** These are the build seat's **RED-first test work for the activated amendment**, landed as an in-package `internal/config` test beside the code they exercise — the natural Go home. The fence row was `config.go` at **file** grain when the amendment's test surface needed **package** grain; that is a fence-grain gap, not a scope breach. Relocating to `test/fixtures/` is *possible* (the legs use exported `config.Load`/`ValidateMemberTransition`) but strictly worse — it would break unit-tests-beside-code for no benefit. Granted in place.

**The shared anatomy (condition-g, owned across the chain — not pinned on you):** the deviation is the touched-but-unnamed class, same as the rev11 self-catch and s10 catch #1, this time entering through an **owner-bytes channel** — and the honest accounting spreads three ways: (i) **my share** — my `lane_vcs` activation ruling authorized the RED-first test surface ("both ways") but did **not enumerate its fence rows** (the FX-VCS file + the invariants descriptor helper); a mid-build amendment activation should name the seams it introduces, and mine didn't; (ii) **your share, which you already owned** — the routing relay asserted "both in-fence" from m-7's locus-naming without re-running the row check (owner-named loci are not fence-membership; they must be reconciled against the block); (iii) the build seat's diff→license table marked the test file "in" at task-seam grain rather than block grain. **The discipline refinement for the ledger:** *when a mid-build amendment activates, its test/fixture seams must be reconciled into the fence — by the activation ruling, the re-fenced PLAN, or the routing relay — and an owner return that names loci outside the block is itself an escalation trigger, not an in-fence assertion.* 

**The catch is the WIN here, and it belongs on the record as one:** the end review's **mechanical diff→block reconciliation — the day-one exit table you were required to build — caught this before any merge traffic**, with the substance already verified (both diffs match m-7's spec, full uncached battery green at `76179ec`, nothing else in the 29-file diff outside the block, `executor.go`/`fieldspec` zero-diff). Three seats' fence-completeness gap, converged and caught by the mechanical table pre-merge. That is the discipline working exactly as designed — the fence held because the table is load-bearing.

**On this grant, the s9 merge decision UNBLOCKS:** the diff→block table now reconciles fully (both rows licensed by this ruling), so cut the s9 merge-decision relay TO the operator with the two grant citations in the diff→license table and the fence anatomy in the dogfood-evaluation section. The MF-1 evidence fold and the two batched owner confirms ride in parallel as you have them. Merge stays operator-only (`HUMAN_MERGE_AUTHORIZATION` at grant time).

ACTIONS_GIT_REF: none — a fence ruling (disk refs: this relay + one INDEX.md row timestamped 20260714-013011).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0` (`s10-close`); the `s9-evidence-thicken` worktree clean at `76179ec` per your report.
Next requested action: operator carries this to s9.planner; the two rows land in the diff→license table under this license; s9.planner cuts the merge-decision relay TO the operator; master's next owed act is the s9 step-integration verification when that relay lands.
