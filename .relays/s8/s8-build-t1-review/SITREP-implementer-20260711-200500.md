## SITREP — s8.implementer / T1 config foundation review gate

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-t1-review
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — task review only; merge remains operator-only
FROM: s8.implementer
TO: s8.reviewer
CC: s8.planner, operator, m-7.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T1
SUBJECT: adversarial T1 review requested at d87dc21 before T2

Summary: T1 is implemented test-first on `s8-observe-spine`. The pinned engine member now carries `version` and `present_layers`; one config-derived layer context reaches render, validate, grant-digest, and both production `RenderEnv` sites; production `-init` accepts and pins the catalog as the third member. Legacy two-member fixture stores remain readable so the separately-designed adoption route still has a pre-adoption state.

Rail A: `EngineConfig.version`, `present_layers`, and the catalog member are closed/fail-closed at the governed-reader boundary because ignoring them changes activation or enforcement meaning; the `RenderEnv.PresentLayers` carrier is additive internally but has one config-derived production writer.
Rail B: pass — this is configuration generation/history truth and deterministic predicate context, not an adversarial-only mechanism; claims remain confusion-grade.

Boundary contract proof:
- Produces `config.PresentLayers(*Pinned)` = core layers plus configured optional layers, copied from pinned config.
- Both production construction sites in `cmd/frank/main.go` receive the same value.
- Three-member genesis records a digest equal to `config.Load(...).Digest`; the pinned fieldspec hash is `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`.
- Stale forms suppress activated predicate-requiredness while independent enum/subject diagnostics remain available.

Tests / verification:
- RED: focused T1 fixture initially failed to compile on absent `EngineConfig.Version`, `config.PresentLayers`, and `RenderEnv.PresentLayers`.
- GREEN: `go test ./test/fixtures -run 'TestS8FXCFG|TestS8ProductionInit' -count=1`.
- Regression/invariants: `go test ./... -count=1` passed, including `test/invariants`.

ACTIONS_GIT_REF: s8-observe-spine@d87dc21
FINAL_GIT_STATUS_SHORT: ?? .relays/s8/
Out-of-scope preserved: no registry bytes changed; no observe/executor implementation started; no live-store migration; no merge.
Remaining risk: m-7 fidelity confirmation is still required for T1. T2 is held pending m-2's reviewed changeset bytes and this adversarial review.
Next requested action: return an addressed T1 review verdict naming blockers or approval; do not grant merge authority.
