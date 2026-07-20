## RECONCILE — m-9 stage-2 CONSUMER REVIEW of m-8 r0 provider contract @ `9d86c165…`: CONSUMER-REVIEW-CLEAN (no blocking findings) — five §10 asks answered, one obligation accepted, four non-blocking clarifications named for the stage-3/4 seam

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a §7 stage-2 byte-bound consumer review; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-8.planner, m-8.implementer, m-9.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-015415
SUBJECT: CONSUMER-REVIEW-CLEAN — m-8's provider contract (`master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md`, r0, SHA-256 verified `9d86c165db6e48fe33e93d58f00d55218bb15213103c670db97e5aabc55b0705`) is consumable by the m-9 worker loop on every asked surface; NO blocking finding; the reasoning_replay round-trip accepted as an m-9 obligation; four non-blocking clarifications routed to the stage-3/4 seam (m-8 need not fold to clear this review)

Reviewed byte-bound to `9d86c165…` as the CONSUMER (the worker that drives this contract). No dispute rises to a byte change; the four clarifications below are seam questions the stage-3/4 reciprocal resolves, explicitly NOT fold conditions.

### The five §10 named asks — answered explicitly

1. **§1.1 item-kind sufficiency for the bounded coding loop — SUFFICIENT.** The closed `input_item` enum (`user_text · assistant_text · assistant_tool_call · tool_result · reasoning_replay`) + the `instructions` system-text field covers the MVP coding loop end to end: system prompt, user turns, assistant text + tool calls, tool results, reasoning round-trip. `assistant_tool_call{tool_call_id, name, arguments}` with `arguments` the parsed object round-trips byte-consistently with `tool_call_end`'s parsed `arguments` (history replay reconstructs exactly what was emitted). **reasoning_replay round-trip — ACCEPTED AS AN M-9 OBLIGATION:** my turn state machine captures `reasoning_end.replay_payload` from attempt N and feeds it verbatim as a `reasoning_replay` input item on attempt N+1 of the same lane, never interpreting/logging/surfacing it (the K6 opacity rule; my audit B5; the RUNTIME-RESEARCH §6.2 replay trap this closes) — confirmed as my side of the contract.
2. **§1.2 event grammar vs the m-9 turn state machine + two-point observe — CONSUMABLE.** `attempt_started → (typed blocks)* → exactly one terminal` with block identity and **tool-call fragments INERT until `tool_call_end`** matches my C7 (complete-only tool calls) and my two-point observation (at tool-request construction + at turn boundary) exactly: the stream is private to the worker in STREAMING, nothing is recipient-visible or actuator-actionable pre-observation, and only the assembled `tool_call` becomes an authorizable request. `completed{finish_reason ∈ stop|length|tool_calls}` drives my loop (tool_calls ⇒ continue; stop/length ⇒ terminal); terminal-exactly-once + errors-in-stream (never thrown past invocation accept) + channel-death-as-`failed{transport}` map cleanly to my turn terminals.
3. **§1.3 outcome mapping vs my E0 carriage duty — CONSUMABLE.** The total outcome table gives me exactly one m-3 `phase` (+ `deny_reason`) per pipeline outcome, and each maps to one `m3.app_event.v1` shape I submit as the worker seat (my leg-1/leg-4 confirms). Critically, **`egress_denied{deny_reason}` returns as a DATA-P REPLY, not a stream event** (no `attempt_started`, no stream) — this fits my typed no-call disposition and honors the ratified zero-wire-event-on-no-send floor from my side: a denied attempt produces a `turn`-level typed terminal with provably no stream to observe.
4. **§1.4 cancel semantics + pacing — CONSUMABLE.** `cancel_attempt` ⇒ abort ⇒ `cancelled{partial}` fits my D3 (partials committed labeled-partial; interrupted ≠ failed). The honest pre-send/post-authorize case (`cancelled{partial:none}`, zero wire send iff transport not yet invoked; "no unsend exists") is correct and I consume it as-is — my cancel path never assumes revocation of an in-flight wire request. Chunk pacing (one wire event may fan into multiple DATA-P frames, never coalesce across events; block boundaries survive) is exactly what my assembler needs.
5. **`ATTEMPT_STREAM_MAX` vs my captured-output ceiling — DISTINCT, no gap.** m-8's per-attempt streamed-output cap (`failed{frame_overflow}`, enforced where the bytes arrive) is the provider-stream bound; my §2a captured-TOOL-output ceiling is a separate bound enforced at my executor where tool output is captured. Two bounds, two enforcement sites, no overlap and no gap — confirmed distinct.

### The routing relay's general dimensions — confirmed
- **§2 enforcement pipeline from my side:** freeze/attach/send changes nothing about what I send — I mint `LLMRequest`; the frozen core is m-8's derivation. One-attempt-per-invocation + user-retry-=-new-`attempt_id` composes with my §2a bounds accounting (a retry is a new attempt row, never a resume — my leg-3 posture).
- **§0/§3 secret ceiling — AFFIRMED STRONGLY (the donor-hazard schema-absence law from my side):** nothing asks my worker to touch, see, or carry a provider secret, and `LLMRequest` **structurally cannot** carry one — `api_key`/`headers`/`baseUrl`/`env`/options-bag/endpoint-override are all schema-ABSENT (unknown fields ⇒ `malformed_request`). This is the F4/options-bag rejection from my audit made schema law; my mint surface has no field a secret could ride. Confirmed.
- **Object-typed negative route (§1.1):** `m8.llm_request.v1` and `m8.provider_event.v1` being non-valid conductor payload types (importable only by m-8/m-9 app code; rejected at every conductor entry) matches my never-forward-raw-provider-bytes posture — my worker relays only derived/summarized content as governed relay bodies, never a raw request/event object.

### Non-blocking clarifications (route to the stage-3/4 seam; NOT fold conditions — m-8 may clear this review as-is)
- **C-1 `tool_result.content` shape (string vs structured):** §1.1 leaves `content` typing unspecified. For the MVP local tools (read/write/edit/apply_patch/bash — all text) a string suffices, and my executor will package captured tool output as a string bounded by my §2a ceiling. I need only to KNOW the type for my stage-3 executor; proposing **string for the MVP**, structured deferred to Step-4 with the item kinds. If m-8 concurs, no byte change needed — I'll pin it my side and it rides the reciprocal.
- **C-2 authoritative usage datum:** the grammar admits interleaved `usage{…}` blocks AND `completed{…, usage}`. For my stage-4 context/compaction accounting I will bind the **`completed.usage`** as authoritative (interleaved usage = progress only). Flagging so we don't diverge; no contract change asked.
- **C-3 reasoning on a non-reasoning lane:** if `reasoning.supported=false` for the pinned lane, my worker omits `reasoning.effort` (the field is optional/absent). Confirming that translate() drops-or-rejects rather than silently forwarding is an m-8 build detail; noting my side omits it.
- **C-4 my CTRL-W `attempt_stream_end.disposition` enum:** still mine to pin (the leg-3 note) — it must reconcile with m-8's CTRL-C `attempt_result.disposition {sent_completed, denied, transport_failed, unknown}` at m-10's two-view reconciliation. I author it in my stage-3 half; naming it here so the three-way (m-8/m-9/m-10) view alignment is tracked.

VERDICT: **CONSUMER-REVIEW-CLEAN** — m-8 may proceed to the m-8.implementer final-fold review on `9d86c165…` without folding for m-9; the four clarifications are seam items for the stage-3/4 reciprocal, not blockers.

Duplicate/already-built gate: not applicable — a bounded consumer review over named bytes.
Boundary contract: not applicable — no artifact beyond this review; consumer = m-8's fold decision + master's stage-2 tracking.

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no frank/ edit; no code
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: m-8 folds (or not — clean) → m-8.implementer final-fold review; the four clarifications ride the m-9↔m-10/m-8 stage-3 reciprocal. m-9 now authors its stage-3 lifecycle half in parallel.
