## REVIEW-FOLD — VP F1/F2 closed at `s12-h16-fix@8ca99650019f2ddd12d47237e1abb32fbd5895e8`; both focused regressions, the ten laws, and one explicitly uncontended full suite are green

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-impl
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded implementation-fidelity correction against the frozen rev21 design
HUMAN_GATE_REQUIRED: no — implementation correction is complete; the operator MERGE-GATE remains held
FILED_AT_LOCAL: 20260818-022818
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260818-021247.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_DOC_SHA256: cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan @ sha256 f1de9e91e2c7d7c87faa471886d186295bbf7251acae54e98919c97e1ade9696
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
HEAD: 8ca99650019f2ddd12d47237e1abb32fbd5895e8
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
SUBJECT: review fold complete — normal startup folds realized mint evidence before serve and failed Class-G retries gate all blind Class-D work

VERDICT: complete
BLOCKING_FINDINGS_CLOSED: 2
DESIGN_BYTE_FINDINGS: 0
PLAN_BYTE_FINDINGS: 0
MERGE_GATE_STATUS: held — no merge, push, PR, deploy, or release action was taken or is implied

## Corrections

1. **F1 — normal-startup pre-serve mint evidence fold.** `Loop.DrainClassG` now scans the canonical open set after the successful Class-G pass, selects only valid unresolved mint attempts, consults the installed `MintRealized` evidence callback, sorts matching relay IDs deterministically, and commits the already-validated `realized-undelivered` transition through the ordinary derived writer. Each transition's `commitDerived` completion rebuilds/publishes the fold before `DrainClassG` returns; `main.go` still opens the authenticated channel only after that synchronous barrier. Canonical failed state prevents duplicate terminalization and preserves no-re-execution.
2. **F2 — Class-G dirty barrier.** `retryClassGIfDirty` now reports whether the retry actually cleared dirty state. Top-of-process blind replay runs only on that success. Existing caller-present mint delivery independently retains its `ClassGDirty` check, so neither blind nor caller-present Class-D work can advance while Class-G remains dirty.

## RED → GREEN regressions

- F1 process crash cut: the child process durably executes `MintOrReplace` for the pivot and dies by `SIGKILL` before any cursor advance. The failure leg makes the terminal record unwritable and proves no channel opens; the normal restart leg proves one terminal, canonical `failed`, byte-identical binding (no remint), and duplicate-free second restart.
  - RED: `batteries/review-fold-f1-red-20260818-022111.txt` · sha256 `497b019bbf32dc8306814e2fb77e22cec918c19fdcc6a856c7647cc01a0f73d4`
  - GREEN at HEAD: `batteries/review-fold-f1-green-20260818-022312.txt` · sha256 `39f1a7e184806657f8ab631ad67c246b2259b9811ccdfc7383b85f1071758eed`
- F2 persistent Class-G failure: the open gate/approval item sees zero hooks and zero record/cursor growth through repeated failed retries; after Class-G succeeds, the trace is Class-G → gate → Class-G → approval → Class-G, followed by byte-count/hook-count idempotency.
  - RED: `batteries/review-fold-f2-red-20260818-022133.txt` · sha256 `0ce580472feb43ea6a4b8410b6af73fba7fb488d8322a9d122362a2b54f32f3a`
  - GREEN at HEAD: `batteries/review-fold-f2-green-20260818-022313.txt` · sha256 `27c2856489f164be75e1ea16e38f9b6f36558c2ea902fe2158036aaa916bd799`

## Verification at the returned tip

- Ten INV-CATALOG laws: `go test ./test/invariants -run '^TestLaw' -count=1 -v` → exit 0; `batteries/review-fold-invariants-green-20260818-022314.txt` · sha256 `3cbae0145ceda74baa4742d5e37cb48bff53ced4d48a2247ed0f03e78c821a3a`.
- Uncontended full suite: quiet window explicitly opened; preflight found `CONCURRENT_PROCESS_PREFLIGHT=none`; operator/planner were asked to run nothing in either worktree; this shell ran only `go test ./... -count=1` → exit 0. Fixtures passed in 184.625s. Capture: `batteries/review-fold-full-green-uncontended-20260818-022414.txt` · sha256 `a4c4f4bdb55150c74caa5f8941c68d335e13c91bd388016b3282f5b7c71a2e59`.
- Exact fence `663e30d99d61d89ec1958ba432389727bcf7c15f..8ca99650019f2ddd12d47237e1abb32fbd5895e8`: 18 commits, 43 changed paths, all under `frank/`; `git diff --check` is silent.
- Correction commit: `8ca99650019f2ddd12d47237e1abb32fbd5895e8` (`s12 REVIEW-FOLD: close H-16 startup and Class-G blockers`) — exactly five files, 310 insertions, 3 deletions.
- Implementation worktree `git status --short`: clean.

FOLD_SCOPE:
- frank/internal/engine/loop.go -> in
- frank/cmd/frank/main.go -> in
- frank/test/fixtures/main_assembly_test.go -> in
- frank/test/fixtures/h16_startup_evidence_test.go -> in
- frank/test/fixtures/h16_classg_barrier_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: committed the bounded five-file correction as 8ca99650019f2ddd12d47237e1abb32fbd5895e8 on local branch s12-h16-fix; filed forward RED/GREEN/invariant/full-suite captures plus this relay and one live-EOF INDEX row in the shared relay tree; no merge, push, PR, deploy, release, design, plan, registry, frozen-oracle, exit-fixture, INV-CATALOG-law, S8-precedence, H-12, or unrelated byte changed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/batteries/review-fold-f1-green-20260818-022312.txt
?? frank/.relays/s12/batteries/review-fold-f1-red-20260818-022111.txt
?? frank/.relays/s12/batteries/review-fold-f2-green-20260818-022313.txt
?? frank/.relays/s12/batteries/review-fold-f2-red-20260818-022133.txt
?? frank/.relays/s12/batteries/review-fold-full-green-uncontended-20260818-022414.txt
?? frank/.relays/s12/batteries/review-fold-invariants-green-20260818-022314.txt
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260818-021838.md
?? frank/.relays/s12/s12-build/REVIEW-FOLD-implementer-20260818-022818.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260818-021247.md

Next requested action: s12.planner independently verifies this returned tip and evidence set, then routes the refreshed merge-decision packet to master for VP re-review. MERGE-GATE remains HELD.
