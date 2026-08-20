## REVIEW-FOLD — m-3 flake ruling complete: nested headroom restored and executor timeout diagnosed

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-review-fold
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded test-only correction under the ratified m-3 ruling
HUMAN_GATE_REQUIRED: no — this fold used the standing IMPL token; the operator MERGE-GATE remains held
FILED_AT_LOCAL: 20260818-044833
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260818-042436.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
STARTING_HEAD: 8ca99650019f2ddd12d47237e1abb32fbd5895e8
HEAD: 1ac14a019f4122a7e43a7f2d17c465ba331e1985
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: test-only flake remedies complete with confirmed executor-timeout, restored numeric margins, green uncontended gate, and clean branch

FOLD_SCOPE:
- frank/test/fixtures/s8_adversarial_test.go -> in
- frank/test/fixtures/s8_exit_gate_test.go -> in
- frank/test/fixtures/s12_nested_scope_test.go -> in
- frank/test/fixtures/ceremony_test.go -> in
- frank/test/fixtures/ceremony_retry_test.go -> in
- frank/test/fixtures/h16_startup_evidence_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

Exactly the six scoped test fixture files changed; `scripts/`, product bytes, policy bytes, design, plan, merge, push, PR, deploy, and release bytes did not change.

## Diagnosis and remedy

The diagnosis instrumented the existing real executor probe to retain the raw `EvaluateClaims` verdict in a failure message. The historical 1.74-second failure sat on the former live 1.731-second sentinel. A controlled boundary RED reproduced the same public `isolation-probe` / `observe-machinery-fault` shape and exposed `Timing:"timeout"` plus `FailingDetail:"executor-timeout"`. The m-3 hypothesis is confirmed; the no-`executor-*` block branch did not fire, and neither a gate write-allowlist fault nor a product-side race was observed. Full details are in `frank/.relays/s12/batteries/implementer-flake-diagnosis-20260818-044741.txt`.

The probe now uses the unique 17.311-second test sentinel. `timeoutSentinel.String()` remains in the forbidden-leak list, so uniqueness and leak detection are preserved; the raw-verdict failure diagnostic remains. The existing product cap at `internal/executor/executor.go:116` is still 120 seconds.

The nested narrowing is executable in `test/fixtures/s12_nested_scope_test.go`: one authoritative 13-test outer-run-only registry gives one reason per crash-cut/ceremony/startup test, the exact m-3 attestation is a checked constant, and the pure decision contract was observed RED before implementation. All listed tests call the same `FRANK_DOGFOOD_NESTED` decision. The exit-gate test carries that same guard and asserts/logs the outer reachability argument immediately after it: reaching the exit-gate body proves the variable was unset in the outer invocation's shared parent environment, licensing every same-guard fixture's outer execution. The in-process retry-authority-delta test remains in the nested suite.

`nested-green attests the courier's gate adjudication over the nested suite; the process-level crash/ceremony cuts are proven ONLY by the outer run.`

## Forward-bound evidence

- Final uncontended full gate: `GOWORK=off go test ./... -count=1`, source tip `1ac14a019f4122a7e43a7f2d17c465ba331e1985`, 2026-08-18 04:43:21–04:46:11 -0700, contention preflight `none`, exit 0, real 169.85 seconds; fixtures green in 168.575 seconds and invariants green in 6.521 seconds.
- Final uncontended nested leg: `FRANK_DOGFOOD_NESTED=1 GOWORK=off go test -p=1 ./... -count=1`, same source tip, 2026-08-18 04:46:34–04:47:27 -0700, contention preflight `none`, exit 0, real 52.45 seconds.
- Numeric margin: unchanged 150-second exit-gate context minus 52.45 seconds = 97.55 seconds restored headroom; unchanged 120-second executor cap minus 52.45 seconds = 67.55 seconds restored headroom.
- Fresh invariant check: `GOWORK=off go test ./test/invariants -count=1`, same source tip, exit 0, green in 1.258 seconds at 04:47:49–04:47:50 -0700.
- Complete capture: `frank/.relays/s12/batteries/implementer-forward-verify-20260818-044741.txt`.

## Tip, fence, and status

New commit: `1ac14a019f4122a7e43a7f2d17c465ba331e1985` (`test: harden nested dogfood timing`). The implementation worktree is clean.

Exact fence: `663e30d99d61d89ec1958ba432389727bcf7c15f..1ac14a019f4122a7e43a7f2d17c465ba331e1985`; 19 commits; 46 files changed, 6434 insertions, 136 deletions. `git diff --binary` SHA-256 is `f5242e9b1cedcdc23ff3769aff0b2578c0c7af7ad794f10f476722cb47d636bb`; the ordered `git rev-list --reverse` SHA-256 is `0943f3af172f48dbeda04d6a96b3bd7679e225aec7bd795a45c327d90c868d59`.

Condition 4 / anti-precedent: option (b), increasing the 150-second test context or the 120-second executor cap, is expressly declined. This ruling does not license moving either value.

ACTIONS_GIT_REF: commit 1ac14a019f4122a7e43a7f2d17c465ba331e1985 in /Users/jack/Programming/harness-s12-h16-fix/frank; no push, PR, merge, deploy, or release action
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/batteries/implementer-flake-diagnosis-20260818-044741.txt
?? frank/.relays/s12/batteries/implementer-forward-verify-20260818-044741.txt
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260818-043218.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260818-042436.md
(literal governance-workspace status captured immediately before this relay existed; the implementation worktree is clean)

Next requested action: s12.planner independently verifies in its own quiet window and returns the corrected whole to master with VP and m-3 copied. MERGE-GATE remains HELD.
