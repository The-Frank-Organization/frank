## IMPL REPORT - s1-core.implementer completed Tasks 0-12 for s1-slice-1

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s1-core-impl-implementer-report
PARENT_DISPATCH_ID: s1-core-impl
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - review panel and merge gate remain separate; this report grants no merge authority
DESIGN_LOCK_ID: s1-slice-1-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@03efa72
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-impl/IMPL-planner-20260703-182721.md
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
SUBJECT: implementation report - Tasks 0-12 complete on main@9c1839e; E2 suite passes; no merge authority

Result:
- Completed the standing dispatch from `s1-core-impl/IMPL-planner-20260703-175416.md`, resumed by `s1-core-impl/IMPL-planner-20260703-182721.md`.
- Implemented Tasks 0-12 in order against `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md`.
- Task 13 remains process work for the pair Planner and was not executed.
- Task 0 shipped the D-3 fallback path: minimal self-hosted JSON-RPC 2.0 over Unix socket with exactly `submit`, `project`, `read`, plus unsolicited server-to-client nudge support. The task-close commit records fallback.
- No PR was opened, no merge was performed, and no live-verify claim is made.

ACTIONS_GIT_REF:
- Final branch head: `main@9c1839e`.
- Implementation range: `03efa72..9c1839e`.
- Commits:
  - `139aaa9` s1 IMPL: channel transport + capability check (fallback)
  - `282aff9` s1 IMPL: crashpoints + atomic fsio + sealed records (F10)
  - `f748268` s1 IMPL: store pivot + redo projections (F2-shape, C5)
  - `254c3e3` s1 IMPL: durable intake + serialized loop (F9, C6)
  - `d0448db` s1 IMPL: FieldSpec registry + render + validate (B3, A2-enum, G-enum, 3)
  - `a8a0650` s1 IMPL: mint/bind/stamp + 3-tool guardrail (B2, A1, G(i)(ii))
  - `3219839` s1 IMPL: lineage + authority superset + fault to held + I-PH formatter (A3, H-core, F5-shape)
  - `037e670` s1 IMPL: submit pipeline e2e (B1, C1, C2)
  - `29349da` s1 IMPL: project/read + push delivery (L1)
  - `22c4e01` s1 IMPL: gate/park/wake/outbox + held visibility (B4, W1, C7, A4, H-r3)
  - `9860c7a` s1 IMPL: recovery + crash-matrix + F11 sweep
  - `5b451f0` s1 IMPL: R1 replay harness + disposition report
  - `eaa5183` s1 IMPL: I-PH + claim-sweep fixtures + honest README (P1, SWEEP)
  - `9c1839e` s1 IMPL: mark plan tasks 0-12 complete

Tracked implementation surfaces:
- Runtime and core packages: `cmd/frank/`, `internal/channel/`, `internal/crashpoint/`, `internal/fsio/`, `internal/record/`, `internal/store/`, `internal/intake/`, `internal/engine/`, `internal/fieldspec/`, `internal/seat/`, `internal/lineage/`, `internal/bounce/`, `internal/gate/`, `internal/recover/`.
- Fixture and replay surfaces: `test/fixtures/`, `test/replay/`, `test/seatproc/`.
- Documentation and plan surface: `README.md`, `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md` checkbox updates only.
- Relay surface: this IMPL report and `.relays/s1/INDEX.md` row only.

Task and acceptance evidence:
- Task 0: capability test covers exact 3-tool list, submit call, unsolicited push, and socket isolation.
- Task 1: F10 and atomic persistence covered by crashpoint, fsio, record, and fixture tests.
- Task 2: F2-shape and C5 covered by store commit/projection tests.
- Task 3: F9 and C6 covered by intake journal and serialized-loop tests.
- Task 4: B3, A2-enum, G-enum, and gate-category raise covered by FieldSpec tests.
- Task 5: B2, A1, G(i), G(ii), and single credential generation covered by seat/channel tests.
- Task 6: A3, H-core, and F5-shape covered by lineage, bounce, and engine fault tests.
- Task 7: B1, A2 reject leg, C1, and C2 covered by submit pipeline tests.
- Task 8: L1 and project/read scoping covered by store project tests.
- Task 9: B4, W1, C7, A4, and H-r3 covered by gate and engine tests.
- Task 10: assembled recovery and F11 sweep covered by recovery and fixture tests.
- Task 11: R1 disposition report committed at `test/replay/report.md`.
- Task 12: P1 and SWEEP claim-honesty fixtures pass; README committed with S1 scope limits.

Verification evidence captured after final commit:
- `go test ./...` -> PASS.
- Passing packages included `internal/bounce`, `internal/channel`, `internal/crashpoint`, `internal/engine`, `internal/fieldspec`, `internal/fsio`, `internal/gate`, `internal/intake`, `internal/lineage`, `internal/record`, `internal/recover`, `internal/seat`, `internal/store`, `test/fixtures`, and `test/replay`.
- No-test packages reported by Go: `cmd/frank` and `test/seatproc`.
- `git status --short --branch --untracked-files=all` -> `## main` before this gitignored report and INDEX write.

Boundary and out-of-scope confirmation:
- No writes were made to `../master` or `../extracted`.
- No drain, egress, migrator, garbage collection, genesis, quarantine, phase-machine, or `model_name` surface was added.
- The fallback channel remains limited to the three seat tools named by D-3.
- `test/replay` uses representative typed-equivalent fixtures and records uncovered S3 disposition instead of claiming full organic corpus closure.
- All proof here is E2 local fixture proof. E3, E4, live runtime qualification, external review panel, and merge remain outside this IMPL report.

FINAL_GIT_STATUS_SHORT: none - clean tracked tree at `main@9c1839e` before this gitignored relay/index write.
Next requested action: s1-core.planner reviews this implementation report and runs the planned review panel before any exit-gate SITREP or merge gate.
