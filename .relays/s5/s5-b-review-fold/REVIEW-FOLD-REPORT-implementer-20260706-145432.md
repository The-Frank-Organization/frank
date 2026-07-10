## REVIEW-FOLD report - s5-b MF-5 tree-invariant §7 fixtures folded

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-b-review-fold
PARENT_DISPATCH_ID: s5-b-impl-fold
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-mechanisms
IN_REPLY_TO: .relays/s5/s5-b-review-fold/SITREP-planner-20260706-144638.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: MF-5 folded; §7 old side now pinned as pre-s5 fixture data and green on branch plus combined tree
FOLD_SCOPE:
- test/fixtures/s5_config_change_test.go -> in
- test/fixtures/testdata/s5_pre_registry.json -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: s5-b-mechanisms@82524f7; commit 82524f7 s5-b: fold tree-invariant config fixtures
FINAL_GIT_STATUS_SHORT: none - clean code worktree

Summary:
- MF-5 folded in the two allowed fixture surfaces only.
- Added `test/fixtures/testdata/s5_pre_registry.json`, the exact pre-s5 `internal/fieldspec/registry.json` bytes from main @ `67ee23e`.
- `s5ConfigChangeDeps` now seeds the §7 test store from that pinned fixture instead of the in-repo registry path, making the OLD side independent of the tree under test.
- The NEW side stays on the landed-registry resolver from MF-4.
- The full-map raise leg was left unchanged; it already consumes only NEW bytes.

Pinned data:
- OLD source commit: `67ee23e`.
- OLD SHA-256: `e31c4b1e72b69699df7e100a9264ee9c10f0d9107c2dd645ddd58107244d7363`.
- Fixture loader fatals on old-fixture drift; it never skips.
- NEW source remains `dd7d0b5`, SHA-256 `827d24dafd0c1bc47e0968c9596aeae2f1575ad4b6e8c2f46a483b4187f1a9db`.

Red evidence:
- Combined-tree red reproduced at this seat before the fix: detached `afddc56` preview + `git merge --no-commit --no-ff 78bda2e` + focused §7 test run failed with:
  - `TestS5ConfigChangeMovesOldToNewDigestOverLandedRegistry`: `digest did not move`.
  - `TestS5ConfigChangeStaleFormRerenderShowsRoutingEscalation`: state `accepted`, wanted `form_digest:re-render`.
- The preview worktree was removed.

Why we missed it earlier:
- Branch-tree runs had in-repo registry = old bytes, so OLD != NEW and §7 passed.
- The earlier merged-main reverify proved the NEW registry bytes were byte-identical to `dd7d0b5`, but it still executed on the branch tree; it did not execute with in-repo registry = NEW bytes. The forward-merge preview is the missing execution context.

Focused verification:
- `shasum -a 256 test/fixtures/testdata/s5_pre_registry.json` returned `e31c4b1e72b69699df7e100a9264ee9c10f0d9107c2dd645ddd58107244d7363`.
- Branch tree: `go test -count=1 ./test/fixtures -run 'TestS5ConfigChange|TestS5GateRaiseFullMapRoutingEscalationFromLandedRegistry|TestS5ALandedRegistryBytes'` passed.
- Combined tree: detached `afddc56` preview + `git merge --no-commit --no-ff 82524f7` + `go test -count=1 ./test/fixtures/` passed; preview removed afterward.

Full branch verification at `82524f7`:
- `go build ./...` exit 0.
- `go test -count=1 ./...` exit 0.
- `go vet ./...` exit 0.

Out-of-scope preserved:
- No production files changed.
- No in-repo registry bytes changed.
- No `cmd/*`, merge, push, PR, or live integration action was taken.
