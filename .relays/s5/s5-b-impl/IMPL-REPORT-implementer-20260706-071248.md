## IMPL report - s5-b mechanisms

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s5-b-impl
PARENT_DISPATCH_ID: s5-b-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-mechanisms
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: implementation report for s5-b-mechanisms-plan T1-T6 and T8-T9; T7 remains orchestrator-gated
ACTIONS_GIT_REF: s5-b-mechanisms@81e0551
FINAL_GIT_STATUS_SHORT: none - clean tree

Summary:
- Implemented T1-T6 and T8-T9 from `s5-b-mechanisms-plan`.
- T7 was not executed; it remains gated on the orchestrator's s5-a integration instruction.
- No merge, PR, push, or cleanup action was taken.

PR:
- none; merge/PR was not authorized by this relay.

Plan lock:
- DESIGN_LOCK_ID: s5-b-mechanisms-design
- PLAN_LOCK_ID: s5-b-mechanisms-plan
- Dispatch relay: `.relays/s5/s5-b-impl/IMPL-planner-20260706-065059.md`

Files changed:
- `internal/fieldspec/validate.go`
- `internal/fieldspec/render.go`
- `internal/fieldspec/fieldspec_test.go`
- `internal/fieldspec/validate_test.go`
- `internal/engine/detector.go`
- `internal/engine/detector_test.go`
- `internal/egress/egress.go`
- `internal/egress/render.go`
- `internal/egress/rules.go`
- `internal/egress/egress_test.go`
- `test/replay/zeroloss/zeroloss.go`
- `test/replay/zeroloss/zeroloss_test.go`
- `test/fixtures/s5_submit_guard_test.go`
- `test/fixtures/s5_gate_raise_test.go`
- `test/fixtures/s5_egress_test.go`
- `test/fixtures/s5_iph_test.go`

Task commits:
- T1: `8618ad9` s5-b: guard seat-supplied system headers
- T2: `22aa3cd` s5-b: add known-a raise detector mechanics
- T3: `5074259` s5-b: cover gate raise fixture legs
- T4: `87f34a3` s5-b: add dormant egress scanner package
- T5: `f5300cf` s5-b: cover dormant egress acceptance
- T6: `dc7e903` s5-b: add zero-loss replay harness
- T8: `81e0551` s5-b: extend iph sweep for new surfaces

Acceptance criteria status:
- DEF-2 guard: implemented; lane and operator supplied system/computed headers reject as `system-owned`.
- ③ raise mechanics: implemented; detector-injected B-pick/B-absorb rewrites, DEF-1 `yes` byte, no A lowering, and claim-boundary negative are covered.
- ⑤ egress: implemented as present-but-dormant package; acceptance fixture proves pass/block/fail-closed paths through `Drain`.
- Replay: implemented; constructed zero-loss replay, canonical-wins, optional archive skip, and Reader.Read refusal legs are covered.
- I-PH: implemented; new Field:Class strings and formatter valve are covered.
- T7: not done by design; gated on the s5-a integration instruction.

Boundary contract proof:
- Writes are limited to the approved fieldspec mechanics, engine detector, dormant egress package, zero-loss replay package, and s5 fixture files.
- Untouched hard-OUT surfaces: `internal/fieldspec/registry.json`, `registry_test.go`, registry-content fixtures, `internal/bounce/formatter.go`, `internal/migrate/migrate.go`, `cmd/*`, `internal/lineage/*`, and `internal/store/*`.
- Guardrail surface remains `submit`/`project`/`read`; delivery-state enum files were not changed.

Tests/verification:
- Baseline before edits: `go build ./...`, `go test -count=1 ./...`, `go vet ./...` all passed at `67ee23e`.
- Per-task battery was run after every task commit: `go build ./...`, `go test -count=1 ./...`, `go vet ./...`.
- Final fresh battery at `81e0551`: `go build ./...` exit 0; `go test -count=1 ./...` exit 0; `go vet ./...` exit 0.
- Forbidden-claim sweep: no active "live scan" or "structurally unbypassable" claim; the only "every content mis-pick" hit is the required negated claim-boundary comment in `test/fixtures/s5_gate_raise_test.go`.
- Report disposition strings `egress_ready` and `egress_blocked` are non-terminal report values, not delivery-state enum members.

Evidence levels:
- E1: file diff and commit inventory above.
- E2: full local build/test/vet battery and fixture package coverage.
- E3: not claimed.
- E4: not claimed.

Out-of-scope preserved:
- T7 not executed.
- No live egress activation.
- No live detector config wiring under `cmd/*`.
- No registry-content edits.
- No transport, lineage, store, migrate, or formatter edits.

Remaining risk:
- T7 and ③ integration wiring remain dependent on the orchestrator's s5-a/m-6 integration instruction.
- ⑤ remains dormant by design until the away-mode bridge imports it.

Next requested action:
- Planner review of this implementation report and downstream orchestration of the T7/integration gate when s5-a is ready.
