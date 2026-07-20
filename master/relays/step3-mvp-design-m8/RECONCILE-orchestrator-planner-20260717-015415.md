## RECONCILE — consumer-review routing to m-9 (§7 stage-2, per the m-8 return path): review m-8's r0 provider contract @ `9d86c165…` as the NAMED consumer — the five §10 review asks + the worker-loop consumability surfaces — returned INTO this lane; m-8 folds, then m-8.implementer reviews the FINAL bytes

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a §7 stage-2 consumer-review routing; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — consumer review; the m-9 grill rides its stage-4 DESIGN
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/DESIGN-planner-20260717-023000.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-8.implementer, master.orchestrator-reviewer, operator
SUBJECT: m-9's stage-2 consumer review of m-8 r0 (`master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` @ SHA-256 `9d86c165db6e48fe33e93d58f00d55218bb15213103c670db97e5aabc55b0705`) — the five named asks + the general worker-loop consumability check; findings or CONSUMER-REVIEW-CLEAN returned into this lane

m-9 — m-8's stage-2 provider contract is authored (r0, hash above; the three stage-1 inputs consumed byte-bound at the hashes you already confirmed). Per the §7 stage-2 return path, **your consumer review comes BEFORE m-8's implementer review**: read r0 at its exact bytes and return either named findings or a clean consumer review INTO this lane (TO master, CC the m-8 pair + VP). m-8 then folds (any byte change re-hashes r0) and routes the final bytes to m-8.implementer.

Review as the CONSUMER (the worker that drives this contract):
1. **The five §10 named review asks** in m-8's doc — answer each explicitly.
2. **The DATA-P request/event stream** (`m8.llm_request.v1` / `m8.provider_event.v1`): consumable by your turn state machine — block identity + fragments-inert-until-assembled tool calls (your C7/amendment-§4 posture), terminal-exactly-once semantics, the opaque `reasoning_replay` round-trip (you carry it back without interpretation), and `egress_denied` as a DATA-P REPLY (never a stream event) fitting your typed no-call dispositions.
3. **The §2 enforcement pipeline** from your side: the freeze/attach/send invariants change nothing about what you send (you mint `LLMRequest`; the frozen core is m-8's derivation); one-attempt-per-invocation + user-retry-= -new-`attempt_id` composes with your §2a bounds accounting.
4. **The §6 E0 outcome table** vs your carriage: every disposition you must carry as the worker seat maps to exactly one `m3.app_event.v1` shape you can submit; `phase`/`deny_reason` values fit the m-3 schema you confirmed.
5. **The §3 S-A custody + §0 claim ceiling**: nothing in the contract asks your worker to touch, see, or carry a provider secret (the donor-hazard schema-absence law from your side).

Bounds: your review is byte-bound to `9d86c165…`; a dispute is a named finding in this lane, never a local re-design; no lock/PLAN/T4/code. (Your stage-3 lifecycle-half dispatch rides in parallel at `step3-mvp-lifecycle-m9/…-015425` — sequence your two items as your pair prefers.)

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the consumer review into this lane; m-8 folds → m-8.implementer final-fold review → the m-8 SITREP names the approved bytes.
