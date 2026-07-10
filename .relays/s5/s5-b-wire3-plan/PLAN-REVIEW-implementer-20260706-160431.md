## PLAN-REVIEW - s5-b wire3 delta approved with loader/test constraints

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-b-wire3-plan-review
PARENT_DISPATCH_ID: s5-b-wire3-plan
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-wire3
IN_REPLY_TO: .relays/s5/s5-b-wire3-plan/PLAN-planner-20260706-155345.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: delta plan approved for T10 live wiring and T11 archive replay; carry explicit error-returning config loader and binary-path cmd proof
PLAN_REVIEW_VERDICT: approve
ACTIONS_GIT_REF: wrote `.relays/s5/s5-b-wire3-plan/PLAN-REVIEW-implementer-20260706-160431.md` and appended `.relays/s5/INDEX.md`; no code edits; no implementation dispatch issued
FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/

Verdict: approve.

I reviewed the addressed plan relay, the updated plan document, the locked design, the s5 orchestrator GO relay, and the three master semantics relays. The delta is inside the orchestrator-granted hop: T10 wires the already-landed detector mechanism into `cmd/frank/main.go`, and T11 runs the existing zeroloss archive leg against the operator copy.

Named defaults:
- S1 empty a_floor default is approved. The merged registry has 83 rows, and the live `record_kind` tokens are `genesis`, `owed_item`, `owed_disposition`, `gate_resolution`, `disposition`, `diagnostics`, and `config_change`; for non-operator seats only `diagnostics` is renderable. The CQ-3 committed-action vocabulary has no live Step-1 atom to bind by default, so shipping an empty table avoids inventing operator values. Fixture-scoped rows are the right way to prove the mechanism.
- S3 input-atom-pending default is approved. The merged registry has no target-branch atom; the only fields matching target/branch/merge/protected are `EVIDENCE_TARGET`, `target_surface`, and `target_gap_result`. Leaving `target_branch_field` unset keeps `s3MergeSplit` inert and avoids pointing config at an undeclared header.

Implementation constraints for the dispatch:
- `DetectorConfigFromPinned` must be error-returning or otherwise error-propagating at startup. The plan's "load-fail loudly" requirement is binding: invalid detector JSON or an `a_floor.member` outside `reg.GateCategory["A"]` must make `run` fail before serving. A config-only return shape that silently drops or panics outside the startup error path is not approved.
- Keep `internal/config/*` untouched. The current `config.Pinned` shape already preserves raw engine member bytes in `Members`, and `EngineConfig` ignores unknown JSON keys through normal `json.Unmarshal`; parse the optional detector section from `pinned.Members["engine"]` on the engine side.
- `cmd/frank/main.go` must bind `env.KnownA` per submit from the fresh `liveTables.Snapshot()` used at the submit handler. `KnownA` does not participate in `Render` or form digesting today, so the existing describe/submit digest shape is compatible.
- The live-path fixture in `test/fixtures/s5_wire3_test.go` must prove the `cmd/frank` production path through the same binary/socket style used by `test/fixtures/main_assembly_test.go`, or an equivalently real `run` path that does not bypass `cmd/frank/main.go`. A fixture that only calls `engine.SubmitHandlerWithRender` would not prove T10's cmd wiring.
- T11 remains read-only against `~/frank-s5-team/replay-store-dogfood-20260706`; an archive incompatibility is a reported finding, not license to mutate the copy.

Scope review:
- In scope for the coming dispatch: `internal/engine/detector.go`, `internal/engine/detector_test.go`, `cmd/frank/main.go`, new `test/fixtures/s5_wire3_test.go`, and `test/replay/zeroloss/*` only if the archive leg needs an in-surface harness fix.
- Out of scope remains hard: registry bytes/tests, `internal/config/*`, `cmd/frank-mcp/*`, and any unrelated production surface. Any need there escalates before dispatch.
- `s5-b-wire3` is not present in my local branch/worktree list at review time; create or verify the fresh branch from `main @ b30df4d` before implementation. Do not implement on `main`.

Verification performed:
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s5/s5-b-wire3-plan/PLAN-planner-20260706-155345.md` -> OK.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s5/s5-b-wire3/PLAN-orchestrator-planner-20260706-154556.md` -> OK.
- Full `.relays/s5` root-mode lint returns known run-level noise for `INDEX.md` and the earlier s5-a merge artifact, while still reporting OK for the target plan relay.
- `find ~/frank-s5-team/replay-store-dogfood-20260706 ...` confirmed the archive copy exists with journal, records, outbox, config, projections, binding, and mailboxes.
- Static probes reviewed `config.Pinned.Members`, current `DetectorConfig`/`KnownADetector`, `cmd/frank/main.go` submit wiring, `Render` digest inputs, and the landed registry row/field set.

Next requested action: if you accept these constraints, issue the scoped `s5-b-wire3-impl` dispatch with SCOPE_DIFF all-in. This relay does not authorize implementation by itself.
