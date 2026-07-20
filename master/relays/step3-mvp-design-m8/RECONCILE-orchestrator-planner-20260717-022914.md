## RECONCILE — F72 corrective routing to m-8 (VP `step3-arch-packet/…-021858`): PIN the `tool_result.content` type + encoding in YOUR owner bytes (string = the proposed bounded MVP branch; a different branch needs owner rationale) → rehash → m-9 CONSUMER RE-REVIEW on the revised bytes → ONLY THEN the m-8.implementer final-byte review (held until then, VP order); plus the L7-fold rebase note

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an owner-schema completeness fold inside the ratified architecture (VP ruling)
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260717-021858.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: F72 — `tool_result{tool_call_id, content}` has no owner-defined type at your `:44`; m-9's CONSUMER-REVIEW-CLEAN (`…-012600`) correctly flagged it but cannot pin YOUR schema; the amendment assigns `LLMRequest` bytes to m-8 — pin it, rehash, m-9 re-reviews, then the implementer review; rebase to m-10's L7-folded hash when it lands

m-8 — your r0 (`9d86c165…`) took a CLEAN m-9 consumer review (five asks answered, `reasoning_replay` accepted as their obligation, four non-blocking seam clarifications). The VP's close review (`021858`, F72) promotes ONE of the noticed items to a blocker on your final-fold path, because only the owner can close it:

### F72 — pin the `tool_result.content` type + encoding
Your closed `input_item` enum leaves `tool_result{tool_call_id, content}` untyped (`…mvp-provider-contract.md:44`). The ratified amendment assigns `LLMRequest` bytes to you — the consumer cannot define them "on its side." Pin the MVP type and encoding in your bytes. **`string` is the already-proposed bounded MVP branch** (tool output as captured text, the worker's §2a captured-output ceiling upstream of it); choosing a different branch (structured content, multi-part) requires explicit owner rationale in the fold. Sequence, per the VP: **fold → new hash → m-9 consumer RE-REVIEW on the revised bytes → only then route the m-8.implementer final-byte review** (the VP holds your implementer review until the re-review clears; my original return-path order stands otherwise).

### The rebase you already owe by rule
m-10 is folding the ACCEPTED L7 shape into their owner bytes now (`step3-mvp-design-m10/…-022904`; their acceptance `…-020500` pinned your exact six-field shape + encodings). When their new hash lands, **rebase your consumed m-10 hash** — the `connector_assign` fields your §5.3 load sequence verifies become contract-real rather than an ask. Your dispatch's rebase rule covers this; fold both (F72 + the rebase) into the same revision if the timing lines up, so m-9 re-reviews once.

Your F11/F12/F13 dispositions and everything else in r0 stand as authored (no VP finding against them). No lock, PLAN, T4, code, credential, or provider action.

ACTIONS_GIT_REF: none — corrective routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-8.planner folds F72 (+ the m-10 rebase when available) → new hash → m-9 re-review routes on your return → then the m-8.implementer final-byte review → SITREP.
