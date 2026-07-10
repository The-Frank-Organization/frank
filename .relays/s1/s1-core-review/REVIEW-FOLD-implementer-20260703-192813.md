## REVIEW-FOLD REPORT - s1-core.implementer folded panel blockers and must-fixes

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-implementer-report
PARENT_DISPATCH_ID: s1-core-review
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - merge and exit-gate SITREP remain separate
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@9c1839e
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-review/REVIEW-FOLD-planner-20260703-190326.md
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
SUBJECT: review fold complete on main@7f66057; B1-B7 and M1-M7 addressed with E2 verification; no merge authority

Scope artifact:
- Filed before tracked edits: `s1-core-review/FOLD_SCOPE-implementer-20260703-191241.md`.
- Fold stayed within its in-scope rows. No edits remain in `internal/crashpoint/*`; delivery/outcome crashpoint names were made live from scoped engine/store files instead of changing the registry.

FOLD_SCOPE:
- cmd/frank/main.go -> in
- internal/channel/capability_test.go -> in
- internal/channel/server.go -> in
- internal/engine/fault_test.go -> in
- internal/engine/loop.go -> in
- internal/engine/loop_test.go -> in
- internal/engine/pipeline_test.go -> in
- internal/engine/submit.go -> in
- internal/fieldspec/fieldspec.go -> in
- internal/fieldspec/fieldspec_test.go -> in
- internal/fsio/fsio.go -> in
- internal/gate/derived.go -> in
- internal/gate/derived_test.go -> in
- internal/intake/journal.go -> in
- internal/recover/recover.go -> in
- internal/recover/recover_test.go -> in
- internal/seat/binding.go -> in
- internal/store/projections.go -> in
- internal/store/store.go -> in
- internal/store/store_test.go -> in
- test/fixtures/f11_test.go -> in
- test/fixtures/iph_test.go -> in
- test/fixtures/main_assembly_test.go -> in
- test/replay/classmap.go -> in
- test/replay/replay_test.go -> in
- test/replay/report.md -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF:
- Final branch head: `main@7f66057`.
- Fold commit: `7f66057` s1 REVIEW-FOLD: address panel blockers and must-fixes.
- Fold range: `9c1839e..7f66057`.

Findings folded:
- B1/B2: loop/channel trust-boundary errors now return path-free reason codes; P1 captures a real loop error path.
- B3/B7: authenticated channel handshake added, `cmd/frank` now assembles recovery, intake, loop, authenticated socket, and submit/project/read tools; binary fixture covers authenticated submit -> project -> read.
- B4: recovery has `RunWithProcessor` and replays `intake - outcomes` before listener startup in assembled runtime.
- B5: server-assigned collision-resistant relay IDs prevent overwrite; F11 now asserts every registered crashpoint has a live hit site, including delivery/outcome sites.
- B6: R1 replay enumerates the real upstream fixture corpus and regenerates `test/replay/report.md`.
- M1/M2: package-global verdict lock removed; operator verdict classification runs through the submit handler/loop path and uses the parent gate edge plus one-shot store state.
- M3: accepted submits emit index/render/mailbox projection intents; gate completion emits a durable park record plus outbox item.
- M4/M5: journal/redo append paths sync through `AppendFsync` with directory fsync; auto IDs are random/collision-checked and duplicate record IDs are rejected.
- M6/M7: fault authority classification uses command/stamped seat role, not payload role; FieldSpec grants render/validate for operator and orchestrator roles.
- O1/O3/O4/O6/O7 partials: errcheck close paths are clean, credential compare is constant-time, binding file writes chmod to 0600, push continues across client write failures by retaining the safe broadcast path, and loop reply sends respect context/timeout.

Verification evidence after commit `7f66057`:
- `go test ./...` -> PASS.
- `go test -race ./...` -> PASS.
- `go vet ./...` -> PASS.
- `golangci-lint run ./...` -> PASS, `0 issues`.
- `git diff --check` -> PASS before commit.
- `git status --short --branch --untracked-files=all` -> `## main` after commit and before this gitignored relay/index write.

Evidence levels:
- E2 local unit/fixture/binary verification only.
- No E3 runtime soak, E4 deployment, merge, or exit-gate SITREP is claimed.

Remaining risk:
- The F11 proof is stronger than the prior hollow registry test, but still an S1-local fixture suite rather than a full organic crash corpus.
- R1 now enumerates the historical corpus, but MVP-invisible classes are explicitly `uncovered-S3` rather than claimed as caught.
- Operator review, planner re-check, exit-gate SITREP, and merge gate remain outside this fold report.

FINAL_GIT_STATUS_SHORT: none - clean tracked tree at `main@7f66057` before this gitignored relay/index write.
Next requested action: s1-core.planner runs the targeted re-check named in `REVIEW-FOLD-planner-20260703-190326.md`.
