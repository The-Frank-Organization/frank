## REVIEW-FOLD — VP R2 returned: AST wiring proof closed; unchanged-boundary reproduction did not fire

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-review-fold
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded test-only correction and reproduction under the standing token
HUMAN_GATE_REQUIRED: no — m-3 owns the R2-F1 sufficiency disposition; the operator MERGE-GATE remains held
FILED_AT_LOCAL: 20260818-152017
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260818-145648.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
STARTING_HEAD: 1ac14a019f4122a7e43a7f2d17c465ba331e1985
HEAD: 08f18725c186d147a5671923a8fcd604d25cf66a
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: R2-F2 closed by deterministic AST binding and negative proof; R2-F1 five-run historical-boundary budget did not reproduce and remains owner-open

FOLD_SCOPE:
- frank/test/fixtures/s12_nested_scope_test.go -> in
- frank/test/fixtures/s12_historical_boundary_repro_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

Exactly one test fixture file landed. The disclosed historical-boundary reproduction harness was removed before staging and does not exist in the commit. No product, policy, shipped timeout, script, other fixture, design, plan, merge, push, PR, deploy, or release byte changed.

## R2-F2 — closed at E2

`TestS12OuterRunOnlyRegistryBindsNamedTestBodies` parses `ceremony_test.go`, `ceremony_retry_test.go`, and `h16_startup_evidence_test.go` with Go AST. Every name in `s12OuterRunOnlyFixtures` must resolve to exactly one `FuncDecl`, and that named body must contain a direct `s12SkipOuterRunOnly(t)` call. Missing bodies, duplicate bodies, and missing calls fail deterministically.

The negative proof uses the same inspector on a synthetic registered `TestSyntheticOuterOnly` body with its helper call absent. Its RED leg was observed before implementation: the no-op scaffold accepted the body and `TestS12OuterRunOnlyRegistryBindingRejectsMissingCall` failed with `missing s12SkipOuterRunOnly(t) call was accepted`. After the AST implementation, the inspector returns `TestSyntheticOuterOnly: missing s12SkipOuterRunOnly(t) in synthetic_test.go`; both the real-source binding test and synthetic negative test pass. The registry-to-function wiring is now test-enforced.

Landed commit: `08f18725c186d147a5671923a8fcd604d25cf66a` (`test: bind nested registry to fixture bodies`), one file, +102/−0.

## R2-F1 — bounded attempt complete; detail not reproduced

The immutable predeclared budget was five fixed trials at the unchanged historical 1731ms live deadline, one prebuilt fixture binary, exactly 180 `yes` CPU competitors per trial, `nice -n 20`, raw verdict retention, and no outcome-responsive tuning. Budget capture: `frank/.relays/s12/batteries/implementer-r2-f1-budget-20260818-151140.txt`, SHA-256 `4848fc3ed36d7136ee6e9cc177402dc2e56670d3faa6a844467c7120b384cd38`.

The first wrapper exposed a zsh harness defect: scalar PID storage did not split during cleanup, so competitors accumulated after valid run 1. Runs 2–5 from that wrapper were invalidated regardless of partial green output; all 900 exact competitor PIDs were removed. This is recorded before correction in `implementer-r2-f1-harness-correction-20260818-151425.txt`, SHA-256 `5492a04a15aaf0a7e820f8bfcc174d8ed7b29e0eded1601b1b89596dbfd8a5f2`. The same prebuilt binary and conditions then completed only the four missing trials using a zsh PID array with exact 180-before/zero-after assertions.

All five valid trials passed under the 1731ms boundary. Raw verdicts were `Outcome:"pass"`, `RungReached:"E2"`, `Predicate:"pass"`, `Timing:"under-timeout"`, `FailingDetail:""`; test durations were 0.57s, 0.36s, 0.20s, 0.19s, and 0.25s. No raw failure fired. Result capture: `implementer-r2-f1-result-20260818-151506.txt`, SHA-256 `6184be12085a0af507e8afe4eeb1f455bfe96473aed092c27cd97ef79977e15e`.

Honest disposition: the historical failing detail remains unidentified; `executor-timeout` remains an inferred cause, not a confirmed cause; the flake item remains open pending m-3's parallel sufficiency pre-ruling. No later m-3 pre-ruling was visible in `master/relays/step3-h16-h26-lane/` at filing time. Reproduction stopped at the declared budget—no tuning-to-fire, no escalation, and no route to the operator gate.

## Forward-bound verification

- Focused at `08f18725c186d147a5671923a8fcd604d25cf66a`: the AST positive test, AST missing-call negative proof, existing nested contract, isolation probe, startup F1 regression, and Class-G F2 regression all passed; fixture package 1.369s, exit 0.
- Ten INV-CATALOG laws: `go test ./test/invariants -run '^TestLaw' -count=1`, green in 1.279s, exit 0.
- Fresh quiet-window full gate: `GOWORK=off go test ./... -count=1`, 2026-08-18 15:16:33–15:19:26 -0700, contention preflight `none` including zero competitor processes, all packages green, fixtures 171.072s, invariants 6.614s, real 173.08s, exit 0.
- Complete capture: `frank/.relays/s12/batteries/implementer-r2-forward-verify-20260818-151943.txt`, SHA-256 `afb4599301e33c3dddcb66b86ad0941d12b00187231d99b022454288a57ed3ff`.

Exact fence: `663e30d99d61d89ec1958ba432389727bcf7c15f..08f18725c186d147a5671923a8fcd604d25cf66a`; 20 commits; 46 files; 6536 insertions; 136 deletions. `git diff --binary` SHA-256 `d12bb5374c6e4e858b504548942a3c8759a3f5f34a6f3774a94c539554d92bb4`; ordered `git rev-list --reverse` SHA-256 `a4212aa4ff56cb42625e6c21201d915c95ff0c16edcd234334c821d9c790f99a`. The implementation worktree is clean.

F1/F2 remain closed at `8ca99650019f2ddd12d47237e1abb32fbd5895e8`; this fold reopens neither. MERGE-GATE remains HELD.

ACTIONS_GIT_REF: commit 08f18725c186d147a5671923a8fcd604d25cf66a in /Users/jack/Programming/harness-s12-h16-fix/frank; battery captures named above; no push, PR, merge, deploy, or release action
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/batteries/implementer-r2-f1-budget-20260818-151140.txt
?? frank/.relays/s12/batteries/implementer-r2-f1-harness-correction-20260818-151425.txt
?? frank/.relays/s12/batteries/implementer-r2-f1-result-20260818-151506.txt
?? frank/.relays/s12/batteries/implementer-r2-forward-verify-20260818-151943.txt
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260818-150926.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260818-145648.md
(literal governance-workspace status captured immediately before this relay existed; the implementation worktree is clean)

Next requested action: s12.planner independently verifies the R2-F2 binding and fresh gate, carries this exact non-reproduction outcome, and combines it with m-3's owner pre-ruling when available before returning to master for VP re-review r3. Do not route the operator MERGE-GATE from this relay.
