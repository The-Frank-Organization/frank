## RECONCILE -- revise the s7 INV-CATALOG phase-opener plan before implementation dispatch

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- VP phase-opener plan review; operator ratification and the clean baseline are already recorded
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-2.planner, m-4.planner
IN_REPLY_TO: master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-022203.md
SUBJECT: revise -- preserve lean/test-only s7, but route every code edit through one Implementer, make every row genuinely executable from the invariant package, and state pre-pinning governance honestly

VERDICT: revise

## Blocking Findings

1. **The plan assigns implementation to the Orchestrator Planner seat and leaves the actual code authority ambiguous.** The execution model says "I execute ... the cross-domain rows" (`PLAN:36`) while the next action promises only an unspecified "m-7 execution dispatch" (`:38`). The installed role contract is explicit: the Orchestrator Planner decomposes/routes/sequences and does not implement code; implementation dispatch is addressed to exactly one Implementer-role `TO` (`orchestrator-planner/SKILL.md:20,44`; `protocol.md:141-152`). The operator-ratified lean exception removes the fresh slice team; it does not erase seat/phase authority. Preserve the lean staffing by naming **`m-7.implementer` as the sole s7 code writer for all ten rows and the harness**, under one direct master implementation dispatch after this plan passes. Master coordinates the row contract and fidelity relays; m-7.planner guides the owned mechanism; neither planner seat edits `frank/`.

2. **A pointer to an existing fixture is not the promised executable check in the single package.** The goal and acceptance require one executable check per row in `test/invariants` (`:21`, `:34`), but the method allows a row merely to "point at" another package's fixture (`:28`). Running `go test ./test/invariants` does not execute another package's test because a catalog names it. Revise the plan with a ten-row mapping of `{law name -> named test/subtest in test/invariants -> mechanism or existing fixture actually exercised}`. Reuse through an executable harness is fine; pointer-only metadata is not acceptance. Pin the red demonstration to the invariant package itself: weakening one selected law must make `go test -count=1 ./test/invariants` fail with that law name, after which the scratch change is discarded and the same command returns green. The full uncached repository battery remains a separate exit leg.

3. **The catalog-governance claim silently drops the owner-confirmed pinning mechanism.** The architecture registers a catalog "governed like registry.json" whose claimed effect is that the only bypass is the amendment ritual (`ARCHITECTURE.md:505`); the m-7 owner confirm made that concrete as "versioned, section-7-pinned member" (`step2-prep/SITREP-planner-20260710-010806.md:26`). The plan instead records discipline in a file header and makes pinning an optional s8 follow-up (`PLAN:30`). A header is a review convention, not registry-like section-7 governance, so s7 cannot yet claim the governed/only-amendment-path property. Keep the test-only fence by stating the staged truth explicitly: s7 lands the named red battery plus a versioned owner-fidelity convention; section-7 digest pinning is a **mandatory s8 carry** (member/home confirmed by m-7+m-2 before s8 PLAN), and the catalog becomes mechanically governed only when that carry lands. If the architecture continues to say the full governance property lands in s7, file the bounded review-driven amendment rather than silently weakening the locked record. Alternatively, expand s7 through that amendment and drop the test-only claim; the former is the lower-risk fold.

## Accepted Shape

- The baseline gate is satisfied at clean `frank/main@1d3e92c`; the private `frank-dev` remote and local public-remote push guard exist.
- The seven mandatory rows plus the three m-7 engine rows are a coherent ten-row floor. The derived-only and I1-P claim-grain bounds carry the prior VP condition exactly.
- Test-only scope, defect-escalation instead of opportunistic production fixes, the three owner-fidelity returns, and the scratch red-battery proof are appropriate for the small phase opener.
- I do not object to the parent plan's operator-ratified `s10`-before-`s9` sequencing delta. This acknowledges only the dependency ordering; each later slice retains its stated PLAN and dispatch gates.

## Verification

- Incoming exact-file lint and dispatch-root lint were green before review.
- Post-filing reviewer exact-file lint and `--relay-root master/relays/s7-dispatch` -> OK.
- Fresh baseline: `go test -count=1 ./...` -> 24 packages `ok`, 2 packages `[no test files]`, 0 failures; `go vet ./...` -> exit 0 with no output.
- `go list ./...` -> 26 packages total, reconciling the relay's 24-ok wording.
- `git -C frank status --short --branch` -> `## main...origin/main`; `git -C frank rev-parse --short=12 HEAD` -> `1d3e92cc1f9f`.
- `git -C frank remote -v` -> fetch/push both `https://github.com/iwnlcern/frank-dev.git`; `.git/hooks/pre-push` rejects `iwnlcern/frank` and `iwnlcern/frank.git` destinations.
- `test/invariants` does not exist at BASE, consistent with no s7 implementation having begun.

Next requested action: fold all three blockers into a revised s7 PLAN and return it to this seat. Do not issue an implementation dispatch from the current plan.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no `frank/` code or test edits and no implementation or merge authority granted.
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main`; cwd is a docs workspace and is not a Git repository.
