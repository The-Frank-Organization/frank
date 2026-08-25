## REVIEW-FOLD — F1/F2/Q3 closed at `s13-m10-module@7bbdb31181038dc36bea0425605ca91a8a18d115`; durable turn replay, fresh-only wake consumption, and owner-final lifecycle are E2 green

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s13-build-impl-fold-f1-f2-1
PARENT_DISPATCH_ID: s13-build-impl
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
GRILL_REQUIRED: no — bounded implementation-fidelity correction against pair-approved design r5
HUMAN_GATE_REQUIRED: no — this fold used the standing IMPL token; planner re-review and the operator MERGE-GATE remain held
FILED_AT_LOCAL: 20260822-143440
IN_REPLY_TO: frank/.relays/s13/s13-build-design/DESIGN-planner-20260822-135724.md
DESIGN_DOC_ID: DS-s13-m10-module
DESIGN_DOC_SHA256: d762cb2f2097f824e1c49422395a4fc9d9ac6024740005fbdde931cff833a2ea
DESIGN_LOCK_ID: DS-s13-m10-module
PLAN_LOCK_ID: s13-build-plan @ sha256 00b1f60368d2e227f48a1bf22d8cb0e1ebeefb800bcc316aaab1bc2e5c9ad161
BRANCH: s13-m10-module
STARTING_HEAD: f090868fb28347de93464fd56df4514928f479cd
HEAD: 7bbdb31181038dc36bea0425605ca91a8a18d115
TARGET_BRANCH: main (operator-only merge; gate HELD)
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: F1/F2/Q3 fold complete at 7bbdb311 — lifecycle/replay/wake regressions and all final gates green; planner delta re-review requested

VERDICT: complete
BLOCKING_FINDINGS_CLOSED: 2
DESIGN_BYTE_FINDINGS: 0
PLAN_BYTE_FINDINGS: 0
MERGE_GATE_STATUS: held — no merge, push, PR, restack, deploy, publication, or release action was taken or is implied

## Corrections

1. **F1 — wake-origin continuation and durable replay.** Continuation is now derived from durable predecessor
   lineage, decodes inherited `admission_ref` into a fresh value, and never re-consumes a wake. Each continuation
   transaction persists its immutable session-log/settlement snapshot, inherited run-scoped `create_auth_id`,
   admission reference, and disclosure rows. `ReemitActive` reconstructs a committed ACTIVE turn from those
   durable sources through the ordinary post-commit emission path without mutating `state_seq` or wake state.
   The regression injects a failed first post-commit emission, then proves recovery re-emits the same continuation
   payload and leaves `resume_disposition=PENDING` plus the already-ADMITTED wake untouched.
2. **F2/Q3 — owner-final lifecycle and replay sources.** Schema v1 now admits only lowercase
   `{created,create_authorized,established}` phases; terminality remains solely in `runs.state`.
   `runs.session_log_path` is written with manifest admission and is the initial-turn replay source.
   `turns.create_auth_id` is NOT NULL 32-lowercase-hex: the fresh admission transaction mints it internally,
   rejects mint failure/noncanonical/collision, commits the guarded `created→create_authorized` edge, and later
   admissions copy it from predecessor lineage rather than caller input or reminting. The additive CTRL-W
   `genesis_committed` family drives the exact ordered six-branch receiver; only the current
   `create_authorized→established` branch mutates.
3. **Exact command-test ruling.** The only moved byte surface under `cmd/frank-app/**` is
   `main_test.go`'s literal `RUNNING→established` phase seed. That fixture seeds an ACTIVE run with an existing
   epoch and broker control, so its genesis has committed and it is past both create edges. No other command byte
   moved.

## RED → GREEN evidence

- Initial RED: `go test -count=1 ./internal/appctl/scheduler ./internal/appipc` failed on the intentionally absent
  `GenesisCommittedBody`, `ReemitActive`, and `RecordGenesisCommitted` mechanisms.
- Focused final: `go test -count=1 ./internal/appctl/scheduler ./internal/appctl/store ./internal/appipc ./cmd/frank-app`
  → exit 0; scheduler `0.856s`, store `0.913s`, appipc `1.250s`, frank-app `1.130s`.
- Reduced-limit first pass exposed and corrected two stale fresh-admission fixture phases; the final
  `go test -tags frank_test_reduced_limits -count=1 ./internal/appipc/... ./internal/appctl/...` → exit 0;
  scheduler `2.637s`, store `1.545s`, supervisor `3.706s`.
- Complete sequence-honest capture: `frank/.relays/s13/batteries/FOLD-F1-F2-20260822.txt` in this fold commit.

## Verification at the returned tip

- Exact full gate: `go build ./... && go test -count=1 ./... && go vet ./...` → exit 0; engine `8.553s`,
  fixtures `212.099s`, invariants `5.997s`, vet silent.
- Disclosed intermediate datum: the preceding unchanged-tree full attempt had one timing-sensitive foreign
  `internal/engine/TestLoopNestedReplySendTimesOutWhenReceiverAbandons` failure; its exact isolated reproduction
  passed 10/10, no `internal/engine` byte moved, and the subsequent exact aggregate gate above is green.
- Compile matrix: `go test -run '^TestLimitsCompileMatrix$' -count=1 ./internal/appipc` → exit 0, `0.779s`.
- Race stability: `go test -race -count=3 ./internal/appctl/... ./internal/appipc/... ./cmd/frank-app` → exit 0;
  scheduler `7.056s`, store `8.724s`, appipc `9.118s`, frank-app `7.931s`.
- Fresh ten-law confirmation: `go test -count=1 ./test/invariants` → exit 0, `7.968s`.
- `go mod verify` reports `all modules verified`; `go mod tidy -diff` and `git diff --check` are silent.
- Source worktree `git status --short` is empty at exact commit `7bbdb31181038dc36bea0425605ca91a8a18d115`.

## Carriage, scope, and evidence boundaries

Carriage-v4 remains exact at the returned tip:

- `internal/executor/executor.go` = `05b529fce363efbe6c22b0b164efb423870e1f32a8a1033cd24811a0387cda76`.
- `internal/executor/executor_test.go` = `858e28cfd1179bc215a3795d1eabbefaa9fa41a0a56f05e3a8a102a10a0b73cb`.
- `test/fixtures/s8_exit_gate_test.go` = `c339bc0f2b2467d29b05d1a273acd905cc1ec66b8922d741698ae05d03c1af8f`.
- `test/fixtures/s8_executor_test.go` = `93f1e94b734b3b35d05cc4fcb856d3a25f1ab151621a2f1327ef569f599be6ab`.

FOLD_SCOPE:
- frank/internal/appctl/** -> in
- frank/internal/appipc/** -> in
- frank/cmd/frank-app/main_test.go -> in
- frank/.relays/s13/batteries/FOLD-F1-F2-20260822.txt -> in
- frank/.relays/s13/s13-build-impl/SITREP-implementer-20260822-135441.md -> in
- frank/.relays/s13/s13-build-impl/REVIEW-FOLD-implementer-20260822-143440.md -> in
- frank/.relays/s13/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

Evidence is E2 component evidence. No live composed m-8/m-9 turn, provider credential, provider call, E3/E4,
restack, merge, push, PR, CI/CD, deployment, publication, release, operator decision, policy ownership, or sibling
lane closure is claimed. The independent planner delta re-review and the serialized restack/rerun/re-review path
remain required before the operator's terminal MERGE-GATE.

ACTIONS_GIT_REF: committed the bounded correction as 7bbdb31181038dc36bea0425605ca91a8a18d115 (`s13 IMPL: fold lifecycle and durable turn replay corrections`) on local branch s13-m10-module in /Users/jack/Programming/harness-s13-m10-module/frank; filed the fold battery, this relay, and one append-only live-EOF s13 INDEX row; no push, PR, restack, merge, CI/CD, deployment, publication, or release
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/docs/sprints/active/2026-08-20-s13-m10-module/designs/DS-s13-m10-module-20260820.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260822-140014.md
?? frank/.relays/s13/s13-build-design/DESIGN-planner-20260822-135724.md
?? frank/.relays/s13/s13-build-impl/REVIEW-FOLD-implementer-20260822-143440.md
?? frank/.relays/s13/s13-build-impl/SITREP-implementer-20260822-135441.md
(literal governance-workspace status after this relay and its INDEX row exist; source worktree is clean)

Next requested action: `s13.planner` independently reviews the exact `f090868f..7bbdb311` fold and returns a delta verdict. MERGE-GATE remains HELD; no merge authority is requested or implied.
