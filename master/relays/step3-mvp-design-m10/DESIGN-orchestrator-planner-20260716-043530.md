## DESIGN SUPPLEMENT to `step3-mvp-design-m10` (VP F69, `step3-arch-packet/…-043205`) — m-3 joins your CC + consumer-confirmation set (it consumes your `run_manifest_digest` in the external-E3 tuple); and you are a named confirmer of m-2's F58 components. Append-only: the `…-041640` dispatch stands except as supplemented here

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a consumer-routing supplement per the VP review; no scope or policy change
GRILL_REQUIRED: no — unchanged for THIS stage-1 lane; your grill rides the stage-5 m-10 control-plane DESIGN
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-043205.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-3.planner, m-3.implementer, m-2.planner, m-2.implementer, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-1.planner, m-7.planner, master.orchestrator-reviewer, operator
SUBJECT: supplement — consumer-set corrections: m-3 CONFIRMS your run-manifest-digest producer seam (the external-E3 tuple binds it); you CONFIRM m-2's relay-verb digests + mapping version (the components your serve gate verifies)

m-10 — the VP's dispatch review (`043205`, F69) corrected the consumer-confirmation routing on your lane. Your `…-041640` dispatch otherwise stands unchanged; this supplement is append-only.

### The two edge corrections
1. **m-3 joins your CC + consumer-confirmation set** (alongside m-9, m-8, m-7): m-3's external-E3 tuple binds your **`run_manifest_digest`** — the F62 applicability evaluator turns on it. Your DESIGN's digest production (what is digested at run freeze, canonical encoding, when it changes) must be confirmable by m-3 as the evidence consumer. (The reciprocal direction already stands — the m-3 dispatch names you a consumer.)
2. **You are a named confirmer of m-2's F58 components** (their supplement `step3-mvp-design-m2/…-043520` adds you to their consumer set): confirm that the relay-verb tool-schema digests + mapping version, as m-2 canonically encodes them, are verifiable by your exact-set-equality serve gate against the locked identity vector.

Everything else in `…-041640` (the app-IPC/supervision interface, the manifest/dispatch-seam contract, the F59 ticket protocol, the durable app-state schema, the §6 wake contract, the no-policy/no-seat/opaque-refs bounds) is unchanged. No affected final-byte closure or interface-lock readiness may be claimed until this supplement is consumed; DESIGN-only authoring may continue meanwhile.

ACTIONS_GIT_REF: none — supplemental dispatch relay + one INDEX row only; the `…-041640` historical bytes untouched; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner consumes this supplement; the DESIGN's consumer set = m-9 + m-8 + m-7 + m-3; the m-2-components confirmation is owed when m-2's bytes land; pair review + SITREP as dispatched.
