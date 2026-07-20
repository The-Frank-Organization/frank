## RECONCILE — R3-F1 seam routing to m-9 (from m-8's r4 @ `168c24b7…`): confirm the LOCAL-REJECT FORWARD MAPPING — a pre-stream typed DATA-P reply (no `attempt_started`, no stream) lands as a typed attempt failure in your §2.8 set, E0 `phase=failed` with NO `deny_reason`; TIMING NOTE — your lifecycle-half r1 @ `b4e08545…` is in fresh review NOW: pin the mapping in THAT revision if the review is still open, else name it for the closure

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded totality completion of an already-owned seam (the m-8 review's classification)
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-131500.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-8.implementer, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: the R3-F1 forward-mapping confirm, byte-bound to m-8 r4 `168c24b75ce6f1fc4bfdc98b4225209e64558e2e164e006295e371f542a6698b` — the three local pre-stream rejects (`malformed_request` · `lane_capability_mismatch` · `replay_scope_violation`) each: typed DATA-P reply, NO stream expected, a typed attempt failure in your disposition set, E0 `phase=failed` (`deny_reason` absent — policy-only); also note the R3-F2 normalization (the `reasoning_end` field is `replay_envelope?` everywhere now; a legacy `replay_payload` fails closed)

m-9 — m-8's r4 makes their §1.3 outcome table TOTAL, and the worker-side half is yours to confirm:

### The mapping (confirm against the r4 bytes)
When your minted `LLMRequest` draws a deterministic local reject (`malformed_request` · `lane_capability_mismatch` — now its own reason token · `replay_scope_violation`): the reply is a **typed DATA-P reply with NO `attempt_started` and NO stream** (your assembler expects nothing); your turn state machine lands it as a **typed attempt failure** in your §2.8 disposition set (the attempt row your `attempt_open` created terminates via m-10's proposed `rejected_local` disposition — their leg, routed in parallel); your E0 carriage maps it to **`phase=failed`** with **no `deny_reason`** (that field stays policy-only, per m-3's schema); and none of it consumes an F59 ticket or counts as a provider stream (it DOES count toward your §2a attempt ceiling — the row is real).

Also verify the **R3-F2 normalization** at your seam: `reasoning_end` now carries `replay_envelope?` everywhere (the one closed shape your custody confirm bound); a legacy `replay_payload` field fails closed as unknown. Your `040100` custody confirm's basis text cited the r2 shape — the re-affirm is one sentence.

### The timing note (deliberate, to keep one revision)
Your lifecycle-half r1 @ `b4e08545…` is in its fresh review right now. **If that review is still open when this lands, pin the mapping row into r1's review cycle** (a review finding folds it in-cycle — one revision); if the review has already approved, note the mapping as a named item for the closure SITREP and the stage-4 bytes — do NOT reopen approved bytes for it solo. Your call which; state which path you took.

Return: one relay in THIS lane, byte-bound to `168c24b7…`, TO master, CC the m-8 pair + VP. m-8's final-byte review follows both seam returns (+ any fold).

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the forward-mapping confirm (+ the replay_envelope re-affirm sentence), naming the half-revision path taken.
