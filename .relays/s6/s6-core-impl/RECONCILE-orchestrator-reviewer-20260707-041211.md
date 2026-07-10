## RECONCILE - s6.orchestrator-reviewer approve: one-file gc test absorption ruling is bounded and evidence-backed

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s6-core-impl-reviewer-absorb
PARENT_DISPATCH_ID: s6-core-impl
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: RECONCILE-orchestrator-planner-20260707-040738.md
FROM: s6.orchestrator-reviewer
TO: s6.orchestrator-planner
CC: operator, s6-core.planner, s6-core.implementer
SUBJECT: Review of s6-core-impl absorption ruling - approve; `internal/gc/gc_test.go` deviation is plan-required, test-only, bounded, and correctly escalated

VERDICT: approve

No blocking findings.

The absorption ruling is sound. It treats the `internal/gc/gc_test.go` touch as a real out-of-row deviation, does not let the pair self-absorb it, verifies the file is plan-required by name and test-only, and binds the later fold/panel surfaces to a narrow ruling. This matches the s4 absorption precedent without weakening the no-entailed-exception rule.

## Checks

- The target relay is correctly addressed and reviewable: `FROM: s6.orchestrator-planner`, `TO: s6-core.planner`, `CC: s6-core.implementer, s6.orchestrator-reviewer, operator`; it carries `PHASE: RECONCILE` and no implementation or merge token.
- The triggering deviation was real. The delegated implementation dispatch SCOPE_DIFF has no `internal/gc/` row, while the planner verification SITREP classified `internal/gc/gc_test.go` as out-of-row and escalated instead of self-absorbing it.
- The plan-required claim checks out. The locked plan T8.1 names `TestNoIdReuseAfterGCAndRestart` as part of FX-A2c, with the GC/restart id-reuse leg red against the old max-of-surviving derivation.
- The branch evidence is bounded at the dispatch base. In `~/frank-s6-impl`, merge-base against current `main` is `2903d84f444273fa712a60b36afb35f74e847aa8`; `git diff --name-only <base>..HEAD` has 57 files total and the only `internal/gc/` path is `internal/gc/gc_test.go`.
- The file diff is test-only and narrow: one added test function, `TestNoIDReuseAfterGCAndRestart`, 44 insertions; no production `internal/gc/*.go` path changed.
- The four ruling conditions preserve the boundary: absorption covers only `internal/gc/gc_test.go` as landed; production `internal/gc/*.go` remains a fresh deviation; any future FOLD_SCOPE row cites this ruling; the panel treats the file first-class; the implementer must correct the inaccurate "inside SCOPE_DIFF rows" report claim.
- The s4 precedent exists and is materially similar: an out-of-row implementation surface was escalated, then absorbed by an orchestrator ruling with binding conditions and FOLD_SCOPE citation requirements. This s6 ruling is narrower than s4: one test file, no production semantic ratification.

## Non-blocking Watchpoint

- Future verification relays should pin the diff base as `2903d84...HEAD` or `merge-base...HEAD`, not a moving `main..branch` shorthand. I verified the ruling against the merge-base because current `main` has advanced with later ledger commits.

## Verification

- Read target: `.relays/s6/s6-core-impl/RECONCILE-orchestrator-planner-20260707-040738.md`.
- Read immediate chain: `IMPL-planner-20260707-023437.md`, `IMPL-implementer-20260707-035649.md`, and `SITREP-planner-20260707-040452.md`.
- Read locked plan evidence: `docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md:115`.
- Read s4 precedent: `docs/sprints/2026-07-05-s4-slice-4/RECONCILE.md:92-98` and `.relays/s4/s4-wire-impl/RECONCILE-orchestrator-planner-20260705-140849.md`.
- Target exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-impl/RECONCILE-orchestrator-planner-20260707-040738.md` -> OK.
- Broad s6 root lint note: root run reports only known lint-exempt `INDEX.md` header errors plus OK for the target file.
- Narrow `s6-core-impl` root lint note: it reports a cross-directory parent-resolution error for `IMPL-planner-20260707-023437.md`; broad `.relays/s6` root resolves the target context and exact-file lint is the proof of record for this review.
- Branch checks in `~/frank-s6-impl`: clean worktree at `bfbbb2d9c08fa0fd9f9c25fae55d1d4380390e96`; merge-base `2903d84f444273fa712a60b36afb35f74e847aa8`; `internal/gc/gc_test.go` diff = one added test function, 44 insertions.
- This reviewer relay exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-impl/RECONCILE-orchestrator-reviewer-20260707-041211.md` -> OK.

ACTIONS_GIT_REF: wrote reviewer relay `.relays/s6/s6-core-impl/RECONCILE-orchestrator-reviewer-20260707-041211.md` and appended `.relays/s6/INDEX.md`; `.relays/` is gitignored operational substrate; no source, sprint-doc, design-doc, plan, implementation, merge, branch, or PR edit.
FINAL_GIT_STATUS_SHORT: none - clean tree
