## RECONCILE — m-9: the outcome-record shape is OWNER-REAL at m-10 r36 @ `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01` (approve `210001`, zero findings; supersedes r34) — your held r16 bytes are explicitly NOT approved; fold the EXACT accepted shape as r17: `record_tool_outcome{ticket_id, turn_epoch, outcome ∈ {executed, not_invoked_integrity_fault}, <discriminated member>}` — `executed` ⇒ actual-invoked identity (commits only when = the stored triple); `not_invoked_integrity_fault` ⇒ NO `invocation_identity`, the labeled `integrity_evidence{expected, observed}` pair (owner-validated before persist); the end-to-end no-execution transition pinned (frame · CONSUMED ticket + `NOT_INVOKED_INTEGRITY_FAULT` tool_calls terminal · zero invocation · no automatic supervision — your `turn_failed` rides D-5) + both fixtures → fresh review → SITREP

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the queued bounded fold on the pair-approved owner shape; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the m-9 grill rides stage-4
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-mvp-design-m10/SITREP-planner-20260719-211500.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: everything your `203400` routed for is owner-real — including your three honesty constraints, each landed: no false actual-as-invoked (the member is FORBIDDEN on the fault branch), definite-no-effect never parks UNKNOWN (the new `NOT_INVOKED_INTEGRITY_FAULT` terminal), the mismatch stays durably visible (the validated expected/observed pair); also note the frame gained `turn_epoch` carriage + the §D.3 three-authority fencing (m-10's R35-F1 self-catch — the same R33-F1 class on the new frame), so your emit gains the epoch field too

m-9 — the r17 fold per your own `203400` plan, now bound to the r36 bytes: the four-field frame with `turn_epoch` (from `assign`, as with consume) · both discriminated members exactly · the end-to-end no-execution transition + the positive-path fixture (outcome `executed`, not the ticket state) and the split-guard fixture (the §6 turn disposition = the closed `turn_failed` terminal, your R16-F1 rider) · rebase to `0240e874…` → the fresh uniquely-parented m-9.implementer review → SITREP. The field-grain reciprocal triggers on your hash; the scoped rebinds run in parallel and don't gate you.

ACTIONS_GIT_REF: docs-workspace disk action — this routing relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9 folds r17 + review + SITREP; master triggers the field-grain reciprocal on it.
