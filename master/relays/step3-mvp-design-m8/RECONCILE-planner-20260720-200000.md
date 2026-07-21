## RECONCILE — m-8 stage-4 attempt-consumer confirmation (leg 2 of 4): CONFIRM the worker realization at r7 × frozen r12; and the `sent`-token RULING — "authorized+started" is honest at my bytes, no stronger per-attempt wire-crossing fact exists, a true wire-crossing fact is NOT needed for the MVP (the F4-fallback amendment is mine to author if ever required)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a byte-bound consumer confirmation over pair-approved bytes; the operator gates at the stage-6 lock
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260720-192928.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-3.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260720-200000.md

**CONFIRM.** m-9's stage-4 worker realization (`2026-07-19-mvp-full-worker.md`, r7 @ `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`, recomputed this session) consumes the m-8 provider contract (frozen r12 @ `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`) exactly as my bytes intend. No conflict; not a finding. The worker consumes m-8 r12 byte-bound (§Consumed:11). My r12 stays frozen and unedited.

### Locus 1 — worker stream consumption vs my normalized event grammar — CONFIRM
m-9 §2.2: per attempt, `attempt_open` on CTRL-W → DATA-P `LLMRequest` **only after `attempt_open_ok`** → **consume `m8.provider_event.v1` to exactly one terminal (`completed`|`failed`|`cancelled`)** → `attempt_stream_end{disposition}`; tool-call fragments packaged only at completion (§3.2); authoritative usage = terminal `completed.usage`. This realizes my §1.2 grammar exactly — `attempt_started` → typed blocks → terminal-exactly-once, tool-call-inert-until-`tool_call_end`, `completed.usage` authoritative (my C-2 pin) — and my §1.1 request shape (§2.1: `m8.llm_request.v1` with `api_key`/`headers`/`baseUrl`/model-override schema-forbidden). ✓

### Locus 2 — `stream_lost`/`stream_failed`/no-`attempt_stream_end` enum (their flag 2) — CONFIRM
m-9 §2.2 + flag 2: `attempt_stream_end{disposition}` over the closed enum `{stream_completed, stream_failed, stream_cancelled, stream_lost}`, "a no-stream outcome emits none." Flag 2 states it is "mine [m-9's], restated byte-identical, cited owner-real by m-8." This is the exact three-loss-fact distinction my r12 §1.2:67 authored (r11/R10-F1): `stream_failed` = an observed `failed{transport}` terminal; `stream_lost` = channel death of an existing stream with no terminal; a no-stream loss ⇒ no `attempt_stream_end`. m-9 restated it byte-identical and attributes the source correctly. ✓

### Locus 3 — the §6.1 disposition set as my CTRL-C views expect — CONFIRM
m-9 §6.1's total table maps every m-8 r12 `attempt_result` disposition to its E0 terminal `phase`: `completed`→`completed`; `transport_failed`→`failed`; `egress_denied{token}`→`denied`(+`deny_reason`); `rejected_local{4 tokens}`→`failed` (deny_reason absent); `cancelled(pre_transport)`→`cancelled`; `cancelled(post_invocation)`→`cancelled`; live-observer loss→`unknown` (conditional per my r12 §1.3/§6); worker-crash + epoch-inert → **no E0**. This is total and correct over my §1.3 disposition set — cancellation never `failed`, the loss/absence cuts honest, the conditional-emission rule preserved (m-9 correctly cites my liveness-sets-value-not-emission rule). The `phase` tokens themselves are m-3-owned (their leg); the disposition→phase mapping realizes-against my bytes faithfully. ✓

### The `sent`-token RULING (my half; m-3 owns the token label)
**The "authorized + started" reading is HONEST and CORRECT at my bytes, and my r12 carries NO stronger per-attempt wire-crossing fact deliverable to m-9.** Verified:
1. m-9 binds `sent` to observing my `attempt_started{attempt_id}`, which my §1.2 emits **only for attempts that passed authorize** (my §1.2:71 — `egress_denied`, `rejected_local`, and `cancelled(pre_transport)` produce **no `attempt_started`, no stream**). So observed `attempt_started` = **"the connector authorized the request and began the provider attempt."**
2. **`attempt_started` is provably NOT a wire-crossing attestation:** a post-authorize dial failure still emits `attempt_started` then terminates `failed{transport}` (my §1.2 grammar + §2.3 no-retry census — a fresh dial can fail with `{1,0,0,0}` counters after the attempt opened). m-9 cites this exactly.
3. **My wire counters `request_write_started`/`request_write_completed`/`dial_attempts` are m-8-INTERNAL fixture instrumentation** (my §5.1 attempt-accounting boundary + §8 fixture 4) — they are **not on the DATA-P/CTRL-C surface** delivered to m-9. m-9's reading ("m-8-internal fixtures, not on the m-9 surface") is correct; my surface delivers no discrete "bytes crossed" boolean.
4. **A true per-attempt wire-crossing fact is NOT needed for the MVP.** E0 is deliberately coarse — self-reported, never gate-satisfying, never promoting (m-3 §0.1); the authoritative wire-truth is the E3 external observer's job, not an E0 token. Adding a wire-crossing signal to my event contract would have no MVP consumer.
5. **IF a true wire-crossing meaning is ever required, that is the F4-fallback AMENDMENT — mine to author, master-routed — NOT a silent fold and NOT widenable from m-9.** I state it explicitly per your instruction; I judge it **not required for the MVP**.
6. **Boundary with m-3:** whether the *token named `sent`* may mean "authorized+started" is m-3's semantic call (they own `m3.app_event.v1.phase`). My bytes SUPPORT that meaning and carry nothing stronger. **No conflict from my side.** The only conflict path — m-3 ruling that `sent` MUST denote true wire-crossing — would trigger my F4-fallback amendment; if m-3's parallel leg returns that, master routes it to me. I do not anticipate it (their leg proposes the same "authorized+started" reading).

### Verification
- m-9 stage-4 worker r7 SHA-256 recomputed: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` (the confirm target).
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` (unchanged, unedited).
- Confirmed loci: m-9 §2.1:36 · §2.2:37 · §3.2:44 · §6.1:62-78 · §11 flags 2/5:164,167 · census E8:188/E9:190/E11:194 (E8 authority = m-8 provider-policy authorization immediately before attach/send, VP F87 — matches my freeze→authorize→attach→send).
- My r12 wire counters (`request_write_*`/`dial_attempts`) confirmed fixture-internal (2 occurrences, both §5.1/§8).

Claims:
- All three consumer loci verified verbatim against the r7 bytes and my frozen r12; the `sent`=`attempt_started` observability verified honest at my §1.2:56/71 — evidence E1.
- No conflict found; my r12 stays frozen and unedited at `4b670a79…`; no amendment authored (the F4-fallback is named-not-built) — evidence E1.
- No lock, PLAN, T4 token, code, credential use, provider call, or deploy performed or claimed — evidence E1.

ACTIONS_GIT_REF: docs-only in non-git cwd — this confirmation relay + one INDEX.md row; the approved design doc UNCHANGED at `4b670a79…`; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 6e4d657
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree (fresh status, empty output)

Next requested action: master carries this leg-2 attempt-consumer confirmation + the `sent` ruling into the stage-6 lock packet, and reconciles with m-3's parallel `sent`-token-semantics leg; route to me ONLY if m-3 rules `sent` must mean true wire-crossing (⇒ my F4-fallback amendment). m-8.planner holds for the Master+VP interface-lock.
