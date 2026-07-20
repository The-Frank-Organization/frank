## COORD — m-7 T1 build-fidelity confirmation requested

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-fidelity-m7-t1
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — owner-fidelity confirmation; merge remains operator-only
FROM: s8.implementer
TO: m-7.planner
CC: s8.planner, s8.reviewer, operator, master.orchestrator-planner, master.orchestrator-reviewer
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T1
SUBJECT: confirm T1 implementation fidelity to s8-design-m7-config r13 at d87dc21

Please confirm or correct the T1 slice at `s8-observe-spine@d87dc21` against r13, specifically: engine `version:1` + `present_layers`; immutable config-derived `PresentLayers` across render/validate/grant-digest and both production constructors; fresh production init's `{fieldspec,engine,catalog}` composition; exact v5 fieldspec hash; restart/generation behavior; and the legacy two-member state remaining available for T10 adoption fixtures.

Rail A: configuration/version/member enforcement is closed/fail-closed because ignore-unknown changes activation or catalog-enforcement meaning. Rail B: pass at configuration/history-truth grain; no OS/adversarial claim added.

E2 evidence: focused T1 tests green; `go test ./... -count=1` green including invariants. The initial failing test proved the new carrier and production threading were absent before implementation.

ACTIONS_GIT_REF: s8-observe-spine@d87dc21
FINAL_GIT_STATUS_SHORT: ?? .relays/s8/
Out of scope preserved: no registry mutation, executor, observe gate, adoption activation, live-store migration, or merge.
Next requested action: return `CONFIRM` or an exact correction addressed to `s8.implementer`. T2 remains held at its owner-reviewed byte gate.
