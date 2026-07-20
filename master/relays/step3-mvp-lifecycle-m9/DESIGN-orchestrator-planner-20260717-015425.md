## DESIGN dispatch — §7 STAGE-3: the m-9 LIFECYCLE HALF — the seat receiver + turn/session state + the F59 EXECUTOR half, authored as m-9's owned contract against the closed stage-1 byte set; m-10's reciprocal confirmation routes on return (per the RATIFIED MVP amendment r7 @ `2f75f2a1…`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a stage-3 owned half under the ratified amendment; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the m-9 grill rides the stage-4 full worker DESIGN (§7 item 4); this half carries pair review + m-10's reciprocal confirmation
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260717-015405.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-10.planner, m-10.implementer, m-7.planner, m-1.planner, master.orchestrator-reviewer, operator
SUBJECT: author the m-9 stage-3 half: the seat receiver (rediscovery + wake forwarding + push receipt via the broker capability) · turn/session state (one-active-turn; the §2a bounds accounting incl. denied-authorize counting) · the F59 executor half (consume-then-execute + invocation-identity capture + `record_tool_outcome` + crash-window behavior) — pair-reviewed final bytes; m-10 reciprocally confirms

m-9 — stage 1 is closed at the five byte-bound hashes (all thirteen confirmations CONFIRM, incl. your four legs), and the gating m-7↔m-10 edge is closed. Your **§7 stage-3 owned half** opens now. Your four stage-1 confirmations are the grounding — this dispatch asks you to AUTHOR the half those confirms said was implementable.

### Author (you own these bytes; m-9.implementer pair-reviews the FINAL bytes; m-10 reciprocally confirms on return)
1. **The seat receiver:** the worker-side consumption of the broker capability surface (the 3 canonical relay verbs + typed `Describe` + push receipt, per m-7 §2.8 / m-1 §1.4b at their approved bytes) — the **durable-rediscovery loop** (catch-up `project`/`read` on startup/reconnect; the record is truth), **wake forwarding** (`wake_forward{relay_id}` to m-10, duplicate-safe, per their §E), and the capability-hygiene discipline your leg-4 confirm stated (the capability appears in no model context, tool argument, tool-subprocess env, or inherited FD).
2. **Turn/session state:** the one-active-turn invariant (yours — distinct from m-10's supervision lease, the mutual non-re-ownership both confirms noted); the turn state machine over the §2a model/tool loop (attempt lifecycle rows consistent with m-10 §B.1 `attempt_open`/`attempt_stream_end`; a user retry = a NEW `attempt_id`, never a resume); the **§2a bounds accounting** as compiled constants — max attempts/tool calls/wall-clock/output — **including denied-authorize requests counting toward the max-tool-calls ceiling** (the hole-closure your leg-3 confirm accepted); cancellation + the fail-closed-on-EOF rule (their §B.3) as a binding executor obligation.
3. **The F59 EXECUTOR half (your owned side of the ticket protocol):** consume-then-execute against m-10's §D.3 single-transaction conditional consume — `consume_ok` ⇒ execute **exactly the digested call** (the `canonical_args_digest` re-derived over the JCS bytes of the assembled validated immutable args object); the three typed rejections terminally handled (`DUPLICATE_CONSUME`/`STALE_EPOCH` ⇒ no execution, fail closed; `IDENTITY_MISMATCH` ⇒ no execution, internal fault surfaced); **invocation-identity capture at the executor boundary** + `record_tool_outcome{ticket_id, outcome, invocation_identity}`; both crash windows honest (`ISSUED`-never-consumed dies with turn/epoch; `CONSUMED`-no-outcome parks `UNKNOWN_TOOL_OUTCOME` — never replay); streamed fragments INERT until assembly + validation.
4. **The §10 acceptance rows your half carries:** authorized==executed (the actual invocation equals the ticket, + the four negatives), stale-worker denial at the executor, and the crash-outcome honesty row — named as RED-first build targets, not claimed properties.

### Inputs (byte-bound; all already confirmed by you)
m-10 `79fcf742…` (the reciprocal half + F59 record/consume) · m-7 `f072bd99…` (the capability surface) · m-1 `7c8b09a6…` (the identity frame) · m-2 `83d8e63e…` (the mapping API where the receiver touches forms). The rebase rule applies if any input re-hashes (the L1/L7 lock-round candidates are the known possibilities — batched there precisely so your authoring window stays stable).

### Boundaries + return path
No FieldSpec/conductor/store change; no code, PLAN, T4 token, lock, credential, or provider action. This half ≠ the stage-4 full worker DESIGN (local tools + native tool + model loop + catalog build — that dispatch follows m-8's stage-2 close and consumes this half). Return: m-9.planner authors the DESIGN parented to THIS dispatch → m-9.implementer DESIGN-REVIEW (uniquely parented, fresh review per byte revision) → report-only SITREP naming the approved bytes + hash → **m-10's reciprocal confirmation routes on master's direction**. The Master+VP interface-lock is the gate; no self-declared lock.

ACTIONS_GIT_REF: none — dispatch relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner authors the stage-3 half (sequencing vs the routed m-8 consumer review is the pair's choice); pair review; SITREP; m-10 reciprocal confirmation on return.
