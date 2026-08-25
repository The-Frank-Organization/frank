## REVIEW-FOLD — VP R3-F1 returned: ineffective nested calls rejected by the executable top-level guard assertion

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s12-build-review-fold
PARENT_DISPATCH_ID: s12-build-review-fold
RUN_ID: s12
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded test-only correction under the standing review-fold token
HUMAN_GATE_REQUIRED: no — VP r4 decides whether to route the separate operator-only MERGE-GATE
FILED_AT_LOCAL: 20260819-224442
IN_REPLY_TO: frank/.relays/s12/s12-build/SITREP-planner-20260819-222245.md
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_LOCK_ID: s12-h16-fix-plan
BRANCH: s12-h16-fix
BASE: 663e30d99d61d89ec1958ba432389727bcf7c15f
STARTING_HEAD: 08f18725c186d147a5671923a8fcd604d25cf66a
HEAD: ba3f76082ad1fea18b667f5fae0d250c119ad339
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: R3-F1 closed at E2 by an exact executable guard matcher, nested negative RED/GREEN proof, and fresh forward verification

FOLD_SCOPE:
- frank/test/fixtures/s12_nested_scope_test.go -> in
- frank/.relays/s12/batteries/ -> in
FOLD_SCOPE_RESULT: all-in

Exactly one code-repository file changed. No product, policy, design, frozen oracle, exit fixture, registry schema, timeout, other test file, merge, push, PR, deploy, or release byte changed.

## R3-F1 — executable guard shape enforced

`s12InspectOuterRunOnlyBindings` no longer walks a registered test's descendant tree. For every ordinary registered fixture, statement zero must be the exact directly executed `if s12SkipOuterRunOnly(t) { return }` shape. `TestH16StartupFoldsRealizedMintEvidenceBeforeServing` alone must first carry the exact `FRANK_H16_MINT_CRASH_CHILD` prefix — including its `h16CrashAfterMintPersistence(t)` call and return — and must carry that same outer-only guard as statement one. A helper occurrence in a closure, `defer`, dead branch, or any other nested position cannot satisfy the assertion.

The new `TestS12OuterRunOnlyRegistryBindingRejectsNestedIneffectiveCall` uses the same inspector on a synthetic registered body whose only helper call is inside an uninvoked function literal. Before the matcher changed, its required RED leg failed because the recursive inspector accepted that body: `nested ineffective s12SkipOuterRunOnly(t) call was accepted`. At the landed bytes, the inspector rejects it with `TestSyntheticOuterOnly: missing executable s12SkipOuterRunOnly(t) guard in synthetic_test.go`. The prior total-omission negative remains and passes against the stronger error class; the real 13-name registry assertion also passes.

Landed commit: `ba3f76082ad1fea18b667f5fae0d250c119ad339` (`test: require executable nested guard shape`), one file, +93/−18.

## Forward-bound verification

- Focused at `ba3f76082ad1fea18b667f5fae0d250c119ad339`: the real-source AST assertion, existing omission negative, new ineffective-nested-call negative, nested contract, isolation probe, startup F1 regression, and Class-G F2 regression all passed; fixture package 1.248s, exit 0.
- Ten INV-CATALOG laws: `GOWORK=off go test ./test/invariants -run '^TestLaw' -count=1`, green in 1.417s, exit 0.
- One fresh quiet-window full gate: `GOWORK=off go test ./... -count=1`, 2026-08-19 22:40:56–22:44:00 -0700, contention preflight empty, all packages green, fixtures 181.958s, invariants 6.349s, real 183.91s, exit 0.
- Complete capture: `frank/.relays/s12/batteries/implementer-r3-forward-verify-20260819-224413.txt`, SHA-256 `b3e89fd0ee92e27d59384bfcd8457d22947bda2c6b902fab3bb02037f34227f4`.

Exact fence: `663e30d99d61d89ec1958ba432389727bcf7c15f..ba3f76082ad1fea18b667f5fae0d250c119ad339`; 21 commits; 46 files; 6611 insertions; 136 deletions. `git diff --binary` SHA-256 `2897bc63d5ce8dd72fd38f0931d450daef64eeae7bb21f552b9a59f1aeae06f7`; ordered `git rev-list --reverse` SHA-256 `4107845d7c473b6fb6a934ab7b16d826de24325dea1b355d56d1478bc009c671`. The implementation worktree is clean.

F1/F2 remain closed at `8ca99650019f2ddd12d47237e1abb32fbd5895e8`; this test-only durability correction reopens neither. The R2-F1 open-inferred disposition and master's forward record correction are untouched. MERGE-GATE remains HELD.

ACTIONS_GIT_REF: commit ba3f76082ad1fea18b667f5fae0d250c119ad339 in /Users/jack/Programming/harness-s12-h16-fix/frank; battery capture named above; no push, PR, merge, deploy, or release action
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/batteries/implementer-r3-forward-verify-20260819-224413.txt
?? frank/.relays/s12/s12-build/FOLD_SCOPE-implementer-20260819-223742.md
?? frank/.relays/s12/s12-build/REVIEW-FOLD-implementer-20260819-224442.md
?? frank/.relays/s12/s12-build/SITREP-planner-20260819-222245.md

Next requested action: s12.planner independently verifies the stronger assertion, negative proof, focused battery, quiet full gate, and exact fence, then returns the correction to master for VP r4. Do not route or exercise the operator-only MERGE-GATE from this relay.
