## REVIEW-FOLD report — T7 E1 machinery-fault classification folded

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s8-build-t7-fold
PARENT_DISPATCH_ID: s8-build-t8-review-verdict
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded must-revise fold under the standing s8-build-impl token; no merge authority
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T7
IN_REPLY_TO: /Users/jack/Programming/harness/master/relays/s8-build-t8-review/SITREP-planner-20260711-231500.md
FROM: s8.implementer
TO: s8.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-7.planner, m-2.planner
SUBJECT: F1-F3 folded at 590afae; E1 timeout and IO breakage take the machinery row while valid missing-file citations reject as observed false
FOLD_SCOPE:
- internal/observe/checks_base.go -> in
- internal/observe/registry.go -> in
- test/fixtures/s8_decision2_test.go -> in
- test/fixtures/s8_check_registry_e1_test.go -> in
- .relays/s8/ -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: s8-observe-spine@590afae; commit 590afae fix(observe): classify E1 machinery faults
FINAL_GIT_STATUS_SHORT: none — clean code worktree at 590afae before this report artifact

Summary:
- F1: `git-status` timeout now returns `skipped`/`blocked`, timing `timeout`, detail `check-machinery-git-status-timeout`; other command failure returns `check-machinery-git-status`. Neither uses policy-only `unsafe`.
- F2: non-absence `read-file` IO errors now return `skipped`/`blocked` with `check-machinery-read-file`.
- The registry's internal machinery edge recognizes both symbolic families: `executor-*` and `check-machinery-*`. The locked `CheckVerdict` shape is unchanged.
- F3: a lexically valid missing final component now reaches `os.ReadFile` and returns `fail`/`read-file-absent`, so the gate rejects it as `observed-false`.
- Parent symlinks are resolved before containment. Existing escape symlinks and broken final-component escape symlinks remain refused.
- The default E1 timeout remains five seconds. `RegistryEnv` accepts a test-only executable and shorter positive timeout so the fixture exercises the real `CommandContext` timeout path deterministically.

Decision-row proof:
- Each new E1 machinery fixture is paired with a non-authority `blocked` no-vantage control that remains `accepted`.
- E1 git timeout and read-file IO failure both produce authority `held`+escalate and non-authority `rejected`; both name `observe-machinery-fault` and neither can reach accepted+self_reported.
- A valid missing-file citation produces non-authority `rejected`, failure class `observed-false`, failing predicate `missing`.
- Existing executor machinery behavior remains covered by the prior Decision2 fixture.

Sequence-honest RED to GREEN:
- RED after adding the new fixture legs: focused build failed because `RegistryEnv` had no `GitExecutable` or `ReadTimeout` seam; the exact errors were four `unknown field` diagnostics across the two named fixtures. No implementation classification edit had yet been made.
- GREEN after the fold: `go test ./test/fixtures -run 'TestS8(E1Distinguishes|Decision2E1)' -count=1` — PASS.
- Expanded focused: `go test ./test/fixtures -run 'TestS8(Decision2|CheckRegistry|ReadFile|GitStatus|E1)' -count=1` — PASS.

Fresh verification before commit:
- `go test ./test/invariants -run '^(TestInvariantCatalog|TestLawPathHygiene)$' -count=1` — PASS.
- `go test ./... -count=1` — PASS: 25 packages reported `ok`, zero FAILs; fixture package 23.432s.
- `git diff --check` — PASS.
- INV-CATALOG bytes are unchanged: the fold adds no bounce-format sink or other catalog-counted site.

Out-of-scope and gates:
- T8 remains approved and untouched. No T9/T10 implementation, catalog-marker/descriptor/status-byte edit, merge, push, PR, or activation action occurred.
- m-7 has still supplied no reply to `s8-build-fidelity-m7/COORD-implementer-20260711-220000.md`; T7 close remains conditioned on that T5 executor-fidelity verdict after this fold review.
- m-7 has still supplied no reply to `s8-build-fidelity-m7/COORD-implementer-20260711-213500.md`; T9/T10 remain hard-gated on the reviewed s8-v1 catalog bytes.

Next requested action: s8.planner adversarially reviews this F1-F3 fold at `590afae`; T8 stands approved; no T9 work begins until the separate m-7 catalog-byte gate closes.
