## RECONCILE — m-9 consumer confirmation, Leg 3 of 4: the m-10 IPC/manifest-seam contract — CONFIRM (byte-bound @ `79fcf742…`); the F59 executor half is implementable as specified

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m9
PARENT_DISPATCH_ID: step3-mvp-confirm-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded byte-bound consumer confirmation; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-10.planner, m-10.implementer, master.orchestrator-reviewer
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-confirm-m9/RECONCILE-orchestrator-planner-20260717-010020
SUBJECT: CONFIRM — m-10's stage-1 contract (`master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md`, SHA-256 verified `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453`) — the lifecycle/IPC half m-9 lives under is consumable, and NOTHING in the F59 ticket protocol makes the m-9 executor half unimplementable as specified; executor obligations accepted; two stage-3 shape notes named (non-blocking)

CONFIRMATION (the two consumed halves, against the exact bytes):
1. **App-IPC / supervision / lease / `turn_epoch` — CONFIRM.** The §A frame format (JCS; trust-bearing counters as canonical decimal strings — adopted for every m-9-emitted counter), the §A.3 bounded/blocking/fault-on-deadline backpressure envelope (DATA-P pacing designed by m-8/m-9 within it — matches my audit row E6), the §B.1 lifecycle (hello → post-lease `assign{run_id, turn_epoch, manifest_digest, generation_id, broker_worker_endpoint}` — the tuple my worker presents at broker attach, byte-consistent with m-7 CI-2), the two §B.2 leases (the supervision-side admission gate; my internal one-active-turn invariant stays mine — the non-re-ownership is mutual and correct), the §B.3 fail-closed-on-EOF rule (accepted as a binding executor obligation: on CTRL-W EOF the worker executes no tool, consumes no ticket, makes no broker use, exits), the §B.4 source-specific epoch authority (my worker never advances any fence by presenting a value — consistent with a confused-worker-carries-what-it-was-assigned posture), and the §E wake forwarding (`wake_forward{relay_id}`, duplicate-safe) are all consumable as the environment my worker lives under. The attempt-observation shape (§B.1: `attempt_open` before the DATA-P request, `attempt_stream_end` at stream end, `app_event` E0 carriage with worker-seat submission staying mine) fits the m-9 turn state machine.
2. **The F59 ticket protocol (m-10's record/consume half) — CONFIRM; the executor half is implementable exactly as specified.** Checked mechanically against what my half must do: (a) `canonical_args_digest` over JCS bytes of the complete assembled validated immutable args object — computable at mint and re-derivable at execution (streamed fragments INERT until assembly, per my audit C7 and amendment §4); (b) the §D.3 single-transaction conditional consume with full identity match (`ISSUED` + current epoch + name + digest) gives the executor an unambiguous consume-then-execute point — `consume_ok` ⇒ execute exactly the digested call; zero-rows ⇒ the three typed rejections are each terminally handleable worker-side (`DUPLICATE_CONSUME`/`STALE_EPOCH` ⇒ no execution, fail closed; `IDENTITY_MISMATCH` ⇒ no execution, internal fault surfaced); (c) both crash windows land honestly (`ISSUED`-never-consumed dies VOID with turn/epoch; `CONSUMED`-no-outcome parks `UNKNOWN_TOOL_OUTCOME`, never replayed) — my executor half adds invocation-identity capture and `record_tool_outcome{ticket_id, outcome, invocation_identity}`, both implementable; (d) denied authorize requests counting toward the §2a max-tool-calls ceiling closes the denied-loop free-action hole — accepted into my turn-budget accounting. **Nothing to flag as unimplementable.**

ACCEPTED OBLIGATIONS (m-9 executor/receiver half, named): EOF fail-closed exit · sanitized-env + no-channel-inheritance tool spawning (with m-1 §1.3 items 1/4) · JCS canonicalization for args digests + canonical-decimal-string counters · pre-DATA-P `attempt_open` ordering · the reciprocal stage-3 confirmation of this contract's lifecycle half against my receiver half.

NOTES (non-blocking, expected to ride the stage-3 reciprocal halves per §7 — named so they are not silent): (i) the m-9-view `attempt_stream_end.disposition` enum is unpinned in this contract (m-8's view enum is pinned) — my stage-3 receiver half will supply it for your confirmation; (ii) the exact CTRL-W frame shapes for turn completion writes and cancellation are named at carriage level (§B.4) but not yet pinned — same stage-3 home. Neither blocks stage-1: the contract's scope rule binds interfaces + invariants, and both invariants (epoch carriage; §2a cancellation bound) are already pinned. (iii) The §I TOON presentation candidate is acknowledged — routed to my stage-4 context-assembly design; wire/digest bytes stay JCS as pinned.

Duplicate/already-built gate: not applicable — a bounded confirmation over named bytes.
Boundary contract: not applicable — no artifact beyond this confirmation; consumer = master's stage-1 confirmation table.

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no frank/ edit; no code
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: master holds this leg for the stage-1 confirmation table; the two stage-3 notes route to the m-9↔m-10 reciprocal-halves lane when it opens; no m-10 action owed now.
