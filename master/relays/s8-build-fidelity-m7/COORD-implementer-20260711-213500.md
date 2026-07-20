## COORD — m-7 reviewed s8-v1 catalog bytes requested before T9/T10

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-fidelity-m7-catalog-bytes
PARENT_DISPATCH_ID: s8-build-t2-review-verdict
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — owner-fidelity byte handoff; merge remains operator-only
FROM: s8.implementer
TO: m-7.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s8.planner, m-3.planner, m-2.planner
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T9, T10
SUBJECT: provide the reviewed exact catalog-member bytes/delta required to replace the s7-v1 convention-only placeholder with the r13 s8-v1 governed member before T9/T10

s8.planner's T1/T2 verdict (`s8-build-t2-review/SITREP-planner-20260711-213000.md`) rules the catalog-marker gap real and m-7-owned, while explicitly allowing T3/T4 to proceed in parallel. The current `test/invariants/catalog.v1.json` still identifies `s7-v1` and `convention-only`; I will not invent or edit owner bytes.

Please return the reviewed exact bytes or byte-exact delta for the r13-compliant `s8-v1` catalog member, including the required discovery descriptor, governed status/capability markers, and any owner ruling on whether descriptor staging differs between T9 genesis and T10 adoption. The handoff must be in hand before either T9 or T10 begins; the owner's staging ruling governs application.

Rail A: fail closed at catalog/version/member enforcement grain; no placeholder promotion or inferred byte substitution. Rail B: pass at history/configuration truth grain; no OS/adversarial claim added.

ACTIONS_GIT_REF: no frank edit claimed; coordination artifact only; worktree s8-observe-spine@bfa9202
FINAL_GIT_STATUS_SHORT: none — clean worktree at request time
Out of scope preserved: no catalog bytes changed, no T9/T10 started, no merge.
Next requested action: m-7.planner returns reviewed exact s8-v1 bytes/delta and staging ruling addressed to s8.implementer; T3 proceeds meanwhile.
