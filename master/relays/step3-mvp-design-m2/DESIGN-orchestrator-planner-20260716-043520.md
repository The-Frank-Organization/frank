## DESIGN SUPPLEMENT to `step3-mvp-design-m2` (VP F69, `step3-arch-packet/…-043205`) — m-10 joins your consumer-confirmation set (it hosts the exact-set identity verification that consumes your relay-verb schema digests + mapping version); and you are a named confirmer of m-7's transport contract. Append-only: the `…-041620` dispatch stands except as supplemented here

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m2
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a consumer-routing supplement per the VP review; no scope or policy change
GRILL_REQUIRED: no — unchanged for this lane (stage-1; grills ride the stage-4/5 build lanes)
DESIGN_DOC_ID: step3-mvp-design-m2-mapping
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-043205.md
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-2.implementer, m-7.planner, m-7.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: supplement — consumer-set corrections: m-10 CONFIRMS your F58 components (the digests/mapping-version its serve gate verifies); you CONFIRM m-7's transport contract (the mapping module must be neither absorbed nor stranded)

m-2 — the VP's dispatch review (`043205`, F69) corrected the consumer-confirmation routing on your lane. Your `…-041620` dispatch otherwise stands unchanged; this supplement is append-only.

### The two edge corrections
1. **m-10 joins your consumer-confirmation set** (alongside m-9 and m-7): m-10 hosts the exact-canonical-set-equality verification over tool IDENTITY and emits `run_manifest_digest` — it directly consumes your **relay-verb tool-schema digests + mapping VERSION** as members of the locked identity vector. Your DESIGN's canonical encoding + applicability rules must be confirmable by m-10 as the verifying host, not just by the producers beside you.
2. **You are a named confirmer of m-7's transport contract** (their supplement `step3-mvp-design-m7/…-043459` adds you to their consumer set): confirm that the shared transport half **neither absorbs nor strands** your mapping module across the 3-way seam.

Everything else in `…-041620` (the mapping-module contract, the F58 components, the parity vectors, the no-registry-byte bound) is unchanged. No affected final-byte closure or interface-lock readiness may be claimed until this supplement is consumed; DESIGN-only authoring may continue meanwhile.

ACTIONS_GIT_REF: none — supplemental dispatch relay + one INDEX row only; the `…-041620` historical bytes untouched; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-2.planner consumes this supplement; the DESIGN's consumer set = m-9 + m-7 + m-10; the m-7-transport confirmation is owed when m-7's bytes land; pair review + SITREP as dispatched.
