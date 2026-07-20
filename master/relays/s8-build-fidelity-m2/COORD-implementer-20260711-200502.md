## COORD — m-2 reviewed T2 registry changeset bytes requested

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-fidelity-m2-t2
PARENT_DISPATCH_ID: s8-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — owner-fidelity byte handoff; merge remains operator-only
FROM: s8.implementer
TO: m-2.planner
CC: s8.planner, s8.reviewer, operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T2
SUBJECT: provide the sole-reviewed v5-to-successor registry changeset bytes before T2

T1 is committed at `s8-observe-spine@d87dc21`, with the v5 fieldspec still byte-exact at SHA-256 `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`. Per the locked sole-writer/fidelity boundary, I have not touched `internal/fieldspec/registry.json`.

Please provide or confirm the exact T2 changeset bytes and successor marker for the one governed transition:
- add `adoption` and `catalog` to `config_member`;
- add both to `member.seat_scope.operator`;
- remove `surface_intent`'s static `required_when {layer_present:observe}` and `visible_when`, leaving applicability to the step-4.5 producer/profile manifest;
- preserve every unrelated registry byte/contract.

Rail A: the registry remains additive/open for future check entries, while the config-member enum/fill gate is closed because ignore-unknown changes acceptance meaning. Rail B: pass — deterministic grammar/history truth only.

ACTIONS_GIT_REF: no registry edit claimed; implementation branch s8-observe-spine@d87dc21
FINAL_GIT_STATUS_SHORT: ?? .relays/s8/
Next requested action: return the exact reviewed byte delta (and marker, if advanced) addressed to `s8.implementer`; any expansion beyond these scheduled sites escalates to master before work.
